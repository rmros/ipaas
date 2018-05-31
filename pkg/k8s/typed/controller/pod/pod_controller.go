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

package controllers

import (
	"fmt"
	"server/pkg/api/apiserver/v1beta1"
	podUtil "server/pkg/k8s/util/pod"
	"server/pkg/utils/array"
	"server/pkg/utils/log"
	"time"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

var kubeSystemNamespace = []string{"kube-system", "kube-public", "default"}

// PodController run the logic of process pod
type PodController struct {
	Clientset     *kubernetes.Clientset
	PodController cache.Controller
	PodLister     cache.Indexer
	Queue         workqueue.RateLimitingInterface
}

// NewPodController return PodController
func NewPodController(client *kubernetes.Clientset, resyncPeriod time.Duration) *PodController {
	pc := &PodController{}
	podListerWather := cache.NewListWatchFromClient(client.CoreV1().RESTClient(), "pods", v1.NamespaceAll, fields.Everything())
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	resourceEventHandle := cache.ResourceEventHandlerFuncs{AddFunc: pc.OnAdd, DeleteFunc: pc.OnDelete}
	podIndex, podInformer := cache.NewIndexerInformer(podListerWather, &v1.Pod{}, resyncPeriod, resourceEventHandle, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
	pc.Clientset = client
	pc.PodController = podInformer
	pc.PodLister = podIndex
	pc.Queue = queue
	return pc
}

// OnAdd add event call func
func (pc *PodController) OnAdd(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	obj, exist, err := pc.PodLister.GetByKey(key)
	if err != nil {
		return
	}

	if !exist {
		log.Warning("pod %v doesn't exist", key)
		return
	}

	if p, ok := obj.(*v1.Pod); ok {
		if array.StringNotIn(kubeSystemNamespace, p.Namespace) {
			item := &v1beta1.PodLifeCycle{}
			labels := p.Labels
			item.ClusterID = labels["ClusterID"]
			item.ServiceName = labels["minipaas.com/svcName"]
			item.Namespace = p.Namespace
			item.PodName = p.Name
			item.CreateAt = p.CreationTimestamp.String()
			item.DeleteAt = ""
			if podUtil.IsPodReady(p) {
				item.Status = string(v1.PodRunning)
			} else {
				item.Status = "running"
			}
			item.Insert()
		}
	}
}

// OnUpdate update event call func
func (pc *PodController) OnUpdate(old, new interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(new)
	if err == nil {
		pc.Queue.Add(key)
	}
}

// OnDelete delete event call func
func (pc *PodController) OnDelete(obj interface{}) {
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err == nil {
		pc.Queue.Add(key)
	}

	obj, exist, err := pc.PodLister.GetByKey(key)
	if err != nil {
		return
	}

	if !exist {
		log.Warning("pod %v doesn't exist", key)
		return
	}

	if p, ok := obj.(*v1.Pod); ok {
		if array.StringNotIn(kubeSystemNamespace, p.Namespace) {
			item := &v1beta1.PodLifeCycle{}
			labels := p.Labels
			item.ClusterID = labels["ClusterID"]
			item.ServiceName = labels["minipaas.com/svcName"]
			item.Namespace = p.Namespace
			item.PodName = p.Name
			item.CreateAt = p.CreationTimestamp.String()
			item.DeleteAt = p.DeletionTimestamp.String()
			if podUtil.IsPodReady(p) {
				item.Status = string(v1.PodRunning)
			} else {
				item.Status = "running"
			}
			item.Insert()
		}
	}
}

// Run start pod controller
func (pc *PodController) Run(stopCh chan struct{}) {
	log.Info("Pod Controller is begining Running")
	pc.PodController.Run(stopCh)

	if !cache.WaitForCacheSync(stopCh, pc.PodController.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	wait.Until(pc.worker, time.Second, stopCh)

	<-stopCh
	log.Info("Pod Controller is ShutDown")
}

func (pc *PodController) worker() {
	for pc.processNextItem() {
	}
}

func (pc *PodController) processNextItem() bool {
	key, shutDown := pc.Queue.Get()
	log.Info(key, shutDown)
	if shutDown {
		return false
	}
	defer func() {
		pc.Queue.Done(key)
	}()
	log.Info("key == ", key)
	return true
}
