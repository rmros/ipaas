/*
Copyright 2018 huangjia.

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

package system

import (
	"fmt"

	base "ipaas/controllers"
	"ipaas/models"
	"ipaas/pkg/k8s/util/node"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/golang/glog"
)

// ClusterController cluster api server
type ClusterController struct {
	base.BaseController
}

// Overview get cluster detail info
// @Title Overview server
// @Description get cluster detail info
// @Success 200		{object}	models.Overview
// @router /detail [get]
func (c *ClusterController) Overview() {
	clusterID := c.GetString(":cluster")

	clusert := &models.Cluster{ID: clusterID}
	if err := clusert.Get(); err != nil {
		glog.Errorf("when get cluster [%v] overview, get cluster err: %v", clusterID, err)
		c.Response500(err)
	}

	teams, err := new(models.Team).ListAll()
	if err != nil {
		glog.Errorf("when get cluster [%v] overview, get team err: %v", clusterID, err)
		c.Response500(err)
	}

	nodeOverview := models.NodeOverview{}
	cpuOverview := models.CPUOverview{}
	memoryOverview := models.MemoryOverview{}
	podOverview := models.PodOverview{}

	nodeList, err := base.ListNode(clusterID, metav1.ListOptions{})
	if err != nil {
		glog.Errorf("list cluster [%v] node err: %v", clusterID, err)
		c.Response500(fmt.Errorf("list cluster [%v] node err: %v", clusterID, err))
		return
	}
	nodeOverview.Total = len(nodeList)
	for _, item := range nodeList {
		if node.IsNodeReady(&item) {
			nodeOverview.Heathy++
		}
		if node.IsNodeSchedule(&item) {
			nodeOverview.Scheduler++
		}

		capacity := item.Status.Capacity
		allocatable := item.Status.Allocatable
		if cpuQuantity, ok := capacity["cpu"]; ok {
			cpuOverview.CPUCapacity += cpuQuantity.ScaledValue(resource.Milli)
		}
		if cpuAllocatable, ok := allocatable["cpu"]; ok {
			cpuOverview.CPUAllocatable += cpuAllocatable.ScaledValue(resource.Milli)
		}
		if memQuantity, ok := capacity["memory"]; ok {
			memoryOverview.MemoryCapacity += memQuantity.ScaledValue(resource.Kilo)
		}

		if memoryAllocatable, ok := allocatable["memory"]; ok {
			memoryOverview.MemoryAllocatable = memoryAllocatable.ScaledValue(resource.Kilo)
		}
	}

	podList, err := base.ListPod(v1.NamespaceAll, clusterID, metav1.ListOptions{FieldSelector: "status.phase!=Succeeded"})
	if err != nil {
		glog.Errorf("list cluster [%v] pod err: %v", clusterID, err)
		c.Response500(fmt.Errorf("list cluster [%v] pod err: %v", clusterID, err))
		return
	}
	for _, p := range podList {
		glog.Info(p.Status.Phase)
		if p.Status.Phase == v1.PodFailed || p.Status.Phase == v1.PodUnknown {
			podOverview.Error++
		} else if p.Status.Phase == v1.PodRunning {
			podOverview.Running++
		} else if p.Status.Phase != v1.PodSucceeded {
			podOverview.Operation++
		}
	}

	nsList, err := base.ListNamespace(clusterID)
	if err != nil {
		glog.Errorf("list cluster [%v] namespace err: %v", clusterID, err)
		c.Response500(fmt.Errorf("list cluster [%v] namespace err: %v", clusterID, err))
		return
	}
	namespaces := base.DecodeNamespaces(nsList)

	c.Response(200, models.Overview{
		Cluster:        *clusert,
		Teams:          teams,
		Namespaces:     namespaces,
		NodeOverview:   nodeOverview,
		CPUOverview:    cpuOverview,
		MemoryOverview: memoryOverview,
		PodOverview:    podOverview,
	})
}
