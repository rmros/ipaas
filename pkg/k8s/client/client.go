/*
Copyright [huangjia] [name of copyright owner]

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"ipaas/models"
	iappv1 "ipaas/pkg/k8s/typed/apps/v1"
	iv1beta1 "ipaas/pkg/k8s/typed/apps/v1beta1"
	iav1 "ipaas/pkg/k8s/typed/autoscaling/v1"
	iv1 "ipaas/pkg/k8s/typed/core/v1"
	iextensions "ipaas/pkg/k8s/typed/extensions/v1beta1"
	isv1 "ipaas/pkg/k8s/typed/storage/v1"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/golang/glog"
)

var fakes = make(map[string]*clientset)

func init() {
	clusters, err := new(models.Cluster).ListAll()
	if err != nil {
		glog.Fatalf("query cluster err: %v", err)
	}
	config := filepath.Join(homeDir(), "config")
	file, err := os.Create(config)
	if err != nil {
		glog.Fatalf("create k8s config file err: %v", err)
	}

	for k := range clusters {
		if clusters[k].APIToken == "" {
			if err = file.Truncate(0); err != nil {
				glog.Fatalf("clean k8s tmp config file err: %v", err)
			}
			if _, err = file.Write([]byte(clusters[k].Content)); err != nil {
				glog.Fatalf("write k8s tmp config file err: %v", err)
			}
			cs, err := newClientsetByConfile(config)
			if err != nil {
				glog.Fatalf("create k8s client err: %v, who's clusterID is %q", err, clusters[k].ID)
			}
			fakes[clusters[k].ID] = &clientset{Clientset: cs}
		} else {
			cs, err := newClientsetByToken(clusters[k].ClusterName, clusters[k].APIProtocol, clusters[k].APIHost, clusters[k].APIToken, clusters[k].APIVersion)
			if err != nil {
				glog.Fatalf("create k8s client err: %v, who's clusterID is %q", err, clusters[k].ID)
			}
			fakes[clusters[k].ID] = &clientset{Clientset: cs}
		}
	}
	os.RemoveAll(config)
}

func newClientsetByConfile(config string) (*kubernetes.Clientset, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", config)
	if err != nil {
		return nil, err
	}
	cs, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, err
	}
	return cs, nil
}

func newClientsetByToken(clusterName, apiserverProtocol, apiserverHost, apiserverToken, apiVersion string) (*kubernetes.Clientset, error) {
	cfg, err := NewConfig(clusterName, apiserverProtocol, apiserverHost, apiserverToken, apiVersion)

	if err != nil {
		return nil, err
	}

	cs, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	return cs, nil
}

type clientset struct {
	*kubernetes.Clientset
}

//GetClientset return clientset
func GetClientset(clusterID string) *clientset {
	if fake, ok := fakes[clusterID]; ok {
		return fake
	}
	return nil
}

//GetClientsets return clientsets
func GetClientsets() map[string]*clientset {
	return fakes
}

// AddClientset add clientset by clusterID
func AddClientset(c *models.Cluster) error {
	if c.APIToken == "" {
		config := filepath.Join(homeDir(), "config")
		file, err := os.Create(config)
		if err != nil {
			glog.Fatalf("create k8s config file err: %v", err)
			return err
		}
		if err = file.Truncate(0); err != nil {
			glog.Fatalf("clean k8s tmp config file err: %v", err)
			return err
		}
		if _, err = file.Write([]byte(c.Content)); err != nil {
			glog.Fatalf("write k8s tmp config file err: %v", err)
			return err
		}
		cs, err := newClientsetByConfile(config)
		if err != nil {
			glog.Fatalf("create k8s client err: %v, who's clusterID is %q", err, c.ID)
			return err
		}
		fakes[c.ID] = &clientset{Clientset: cs}
		os.RemoveAll(config)
	} else {
		cs, err := newClientsetByToken(c.ClusterName, c.APIProtocol, c.APIHost, c.APIToken, c.APIVersion)
		if err != nil {
			glog.Fatalf("create k8s client err: %v, who's clusterID is %q", err, c.ID)
			return err
		}
		fakes[c.ID] = &clientset{Clientset: cs}
	}
	return nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func NewConfig(clusterName, apiserverProtocol, apiserverHost, apiserverToken, apiVersion string) (*rest.Config, error) {
	config := clientcmdapi.NewConfig()
	config.Clusters[clusterName] = &clientcmdapi.Cluster{Server: fmt.Sprintf("%s://%s", apiserverProtocol, apiserverHost), InsecureSkipTLSVerify: true}
	config.AuthInfos[clusterName] = &clientcmdapi.AuthInfo{Token: apiserverToken}
	config.Contexts[clusterName] = &clientcmdapi.Context{
		Cluster:  clusterName,
		AuthInfo: clusterName,
	}
	config.CurrentContext = clusterName

	clientBuilder := clientcmd.NewNonInteractiveClientConfig(*config, clusterName, &clientcmd.ConfigOverrides{}, nil)

	return clientBuilder.ClientConfig()
}

var timeout int64 = 518400

//Stream
type Stream struct {
	cs    *clientset
	k8sNs string
}

//GetStreamApi
func GetStreamApi(cs *clientset, namespace string) *Stream {
	return &Stream{
		cs:    cs,
		k8sNs: namespace,
	}
}

//WatchResource
func (s *Stream) WatchResource(resourceType string) (watch.Interface, error) {
	options := metav1.ListOptions{
		Watch:          true,
		TimeoutSeconds: &timeout,
	}
	var result watch.Interface
	var werr error
	if resourceType == "pod" {
		result, werr = s.cs.CoreV1().Pods("").Watch(options)
		// result, werr = s.cs.RESTClient().Get().Prefix("watch").Resource("pods").VersionedParams(&options, scheme.ParameterCodec).Watch()
	}
	if resourceType == "deployment" || resourceType == "app" {
		result, werr = s.cs.ExtensionsV1beta1().Deployments("").Watch(options)
		// result, werr = s.cs.RESTClient().Get().Prefix("watch").Resource("deployments").VersionedParams(&options, scheme.ParameterCodec).Watch()
	}
	if werr != nil {
		return nil, werr
	}
	return result, nil
}

func (s *Stream) FollowLog(podName, containerName string, tail int64) (io.ReadCloser, error) {
	logOption := &v1.PodLogOptions{
		Container:  containerName,
		Follow:     true,
		Timestamps: true,
		//SinceTime: &unversioned.Time{
		//	Time: time.Now(),
		//},
	}
	if 0 == tail {
		tail = 100
	}
	logOption.TailLines = &tail
	reader, err := s.cs.CoreV1().Pods(s.k8sNs).GetLogs(podName, logOption).Stream()
	return reader, err
}

func CheckClusterHealthz(clusterID string) bool {
	healthClient, exist := fakes[clusterID]
	if !exist {
		return false
	}
	healthResult := healthClient.Discovery().RESTClient().Get().AbsPath("/healthz").Do()
	if healthResult.Error() != nil {
		return false
	}
	rawHealth, err := healthResult.Raw()
	if err != nil {
		return false
	}
	if string(rawHealth) != "ok" {
		return false
	}
	return true
}

func CheckClusterVserion(clusterID string) (string, error) {
	healthClient, exist := fakes[clusterID]
	if !exist {
		return "", fmt.Errorf("cluster %v's client doesn't exist", clusterID)
	}
	versionInfo, err := healthClient.Discovery().ServerVersion()
	if err != nil {
		return "", err
	}

	return versionInfo.GitVersion, nil
}

func (c *clientset) DeploymentsV1() iappv1.DeploymentInterface {
	return iappv1.Deployments(c.Clientset)
}

func (c *clientset) DeploymentsExtensions() iextensions.DeploymentInterface {
	return iextensions.Deployments(c.Clientset)
}

func (c *clientset) HPAs() iav1.HPAInterface {
	return iav1.Hpas(c.Clientset)
}

func (c *clientset) Statefulsets() iappv1.StatefulsetInterface {
	return iappv1.Statefulsets(c.Clientset)
}

func (c *clientset) StatefulsetsV1beta1() iv1beta1.StatefulsetInterface {
	return iv1beta1.Statefulsets(c.Clientset)
}

func (c *clientset) Pods() iv1.PodInterface {
	return iv1.Pods(c.Clientset)
}

func (c *clientset) Services() iv1.ServiceInterface {
	return iv1.Services(c.Clientset)
}

func (c *clientset) ConfigMaps() iv1.ConfigMapInterface {
	return iv1.ConfigMaps(c.Clientset)
}

func (c *clientset) Events() iv1.EventInterface {
	return iv1.Events(c.Clientset)
}

func (c *clientset) PersistentVolumeClaims() iv1.PersistentVolumeClaimInterface {
	return iv1.PersistentVolumeClaims(c.Clientset)
}

func (c *clientset) StorageClasses() isv1.StorageClassInterface {
	return isv1.StorageClasses(c.Clientset)
}

func (c *clientset) Nodes() iv1.NodeInterface {
	return iv1.Nodes(c.Clientset)
}
