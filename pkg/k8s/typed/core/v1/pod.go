/*
Copyright [yyyy] [name of copyright owner]

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

package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ipaas/pkg/tools/configz"
	"ipaas/pkg/tools/log"
	"net/http"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
)

//PodInterface has methods to work with Pod resources.
type PodInterface interface {
	DeletePod(pod v1.Pod) error
	GetPod(name, namesapce string) (*v1.Pod, error)
	ListPods(namespace string) ([]v1.Pod, error)
	ListPodByDeploymentName(name, namespace string) ([]v1.Pod, error)
	GetPodLogs(name, namespace string, logOptions *v1.PodLogOptions) (string, error)
	GetPodMetrics(namespace, podName, metric_name string) (map[string]interface{}, error)
}

//pods implements PodInterface.
type pods struct {
	*kubernetes.Clientset
}

// Pods return pods.
func Pods(client *kubernetes.Clientset) PodInterface {
	return &pods{Clientset: client}
}

func (client *pods) DeletePod(pod v1.Pod) error {
DELETE_POD:
	deletePropagationForeground := new(metav1.DeletionPropagation)
	*deletePropagationForeground = metav1.DeletePropagationForeground
	if err := client.CoreV1().Pods(pod.Namespace).Delete(pod.Name, &metav1.DeleteOptions{PropagationPolicy: deletePropagationForeground}); err != nil {
		if errors.IsConflict(err) {
			goto DELETE_POD
		}
		return err
	}
	return nil
}

func (client *pods) GetPod(name, namesapce string) (*v1.Pod, error) {
	return client.CoreV1().Pods(namesapce).Get(name, metav1.GetOptions{})
}

func (client *pods) ListPods(namespace string) ([]v1.Pod, error) {
	list, err := client.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		return []v1.Pod{}, err
	}
	return list.Items, nil
}

func (client *pods) ListPodByDeploymentName(name, namespace string) ([]v1.Pod, error) {
	list, err := client.CoreV1().Pods(namespace).List(metav1.ListOptions{LabelSelector: "minipaas.io/name=" + name})
	if err != nil {
		return []v1.Pod{}, err
	}
	return list.Items, nil
}

func (client *pods) GetPodLogs(name, namespace string, logOptions *v1.PodLogOptions) (string, error) {
	req := client.CoreV1().RESTClient().Get().
		Namespace(namespace).
		Name(name).
		Resource("pods").
		SubResource("log").
		VersionedParams(logOptions, scheme.ParameterCodec)

	readCloser, err := req.Stream()
	if err != nil {
		return err.Error(), nil
	}

	defer func() {
		if err = readCloser.Close(); err != nil {
			log.Error("close readstream err:%v", err)
		}
	}()

	result, err := ioutil.ReadAll(readCloser)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

/*
**metricName参考：**
	[
	"network/tx",
	"network/tx_errors_rate",
	"memory/working_set",
	"network/tx_errors",
	"cpu/limit",
	"memory/major_page_faults",
	"memory/page_faults_rate",
	"cpu/request",
	"network/rx_rate",
	"cpu/usage_rate",
	"memory/limit",
	"memory/usage",
	"memory/cache",
	"network/rx_errors",
	"network/rx_errors_rate",
	"network/tx_rate",
	"memory/major_page_faults_rate",
	"cpu/usage",
	"network/rx",
	"memory/rss",
	"memory/page_faults",
	"memory/request",
	"uptime"
	]
*/

//GetPodMetrics get pod metric
func (client *pods) GetPodMetrics(namespace, podName, metric_name string) (map[string]interface{}, error) {
	path := fmt.Sprintf("%s/api/v1/model/namespaces/%s/pods/%s/metrics/%s", configz.GetString("apiserver", "heapsterEndpoint", "127.0.0.1:30003"), namespace, podName, metric_name)
	log.Info(path)
	heapsterHost := configz.GetString("apiserver", "heapsterEndpoint", "http://127.0.0.1:30003")
	log.Info("Creating remote Heapster client for %s", heapsterHost)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	v := map[string]interface{}{}
	json.Unmarshal(data, &v)
	return v, nil
}
