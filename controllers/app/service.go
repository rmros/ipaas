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

package app

import (
	"fmt"

	base "ipaas/controllers"
	"ipaas/models"
	"ipaas/pkg/tools/log"
	"ipaas/pkg/tools/parse"
	"ipaas/pkg/tools/validate"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/golang/glog"
)

// ServiceController the service controller
type ServiceController struct {
	base.BaseController
}

// CreateService CreateService
// @Title CreateService server
// @Description create app
// @Success 200		{object}	models.Service
// @Param	body		body 	models.Service		true	"body for user content"
// @router / [post]
func (c *ServiceController) CreateService() {
	clusterID := c.GetString(":cluster")
	namespace := c.GetString(":namespace")
	svc, err := validate.ValidateService(c.Ctx.Request)
	if err != nil {
		c.Response400(err)
		c.FinishAudit(clusterID, namespace, models.ServiceResourceType, svc.Name, "创建", 2, true)
		return
	}
	service := svc.TOK8SService(namespace)
	deployment := svc.TOK8SDeployment(namespace)
	result, err := base.DelpoyService(clusterID, service, deployment)
	if err != nil {
		log.Error("deploy app where named %q err: %v", svc.Name, err)
		c.Response500(fmt.Errorf("deploy app where named %q err: ", err))
		c.FinishAudit(clusterID, namespace, models.ServiceResourceType, svc.Name, "创建", 2, true)
		return
	}
	c.FinishAudit(clusterID, namespace, models.ServiceResourceType, svc.Name, "创建", 1, true)
	c.Response(200, result)
}

// DeleteService DeleteService
// @Title DeleteService server
// @Description create namespace
// @Success 200
// @router /:service [delete]
func (c *ServiceController) DeleteService() {
	clusterID, namespace, name := c.GetString(":cluster"), c.GetString(":namespace"), c.GetString(":service")
	if err := base.DeleteService(name, namespace, clusterID); err != nil {
		log.Error("delete service where named %q err: %v", name, err)
		c.Response500(fmt.Errorf("delete service where named %q err: %v", name, err))
		c.FinishAudit(clusterID, namespace, models.ServiceResourceType, name, "删除", 2, true)
		return
	}
	c.FinishAudit(clusterID, namespace, models.ServiceResourceType, name, "删除", 1, true)
	c.Response(200, "ok")
}

// ListService ListService
// @Title ListService server
// @Description stop app
// @Success 200		{object}	models.Service
// @router / [get]
func (c *ServiceController) ListService() {
	clusterID, namespace := c.GetString(":cluster"), c.GetString(":namespace")
	services, deployments, err := base.ListService("", namespace, clusterID)
	if err != nil {
		glog.Errorf("get service in namespace [%v] err: %v", namespace, err)
		c.Response500(fmt.Errorf("get service in namespace [%v] err: %v", namespace, err))
		return
	}
	result := map[string]interface{}{
		"services":    services,
		"deployments": deployments,
	}
	c.Response(200, result)
}

// ListServiceEvents ListServiceEvents
// @Title ListServiceEvents server
// @Description list service events
// @Success 200		{object}	models.Event
// @router /:service/events [get]
func (c *ServiceController) ListServiceEvents() {
	clusterID, namespace, name := c.GetString(":cluster"), c.GetString(":namespace"), c.GetString(":service")
	events, err := base.GetServiceEvents(name, namespace, clusterID)
	if err != nil {
		glog.Errorf("list service %v's envets err: %v", name, err)
		c.Response500(fmt.Errorf("list service %v's envets err: %v", name, err))
		return
	}
	c.Response(200, events)
}

// OperatorService OperatorService
// @Title OperatorService server
// @Description start stop reqploy restart scale expansion
// @Success 200
// @router /:service/:verb [put]
func (c *ServiceController) OperatorService() {
	clusterID, namespace, name := c.GetString(":cluster"), c.GetString(":namespace"), c.GetString(":service")

	svc, deploy, exist := base.ServiceExist(name, namespace, clusterID)
	if !exist {
		c.Response500(fmt.Errorf("service %v not found", name))
		return
	}

	verb := c.GetString(":verb")
	switch verb {
	case "start", "stop", "restart", "redeploy":
		if err := base.OperatorService(svc, deploy, verb, clusterID); err != nil {
			c.Response500(fmt.Errorf("%v service %v err: %v", verb, deploy.Name, err))
			return
		}
		c.Response(200, "ok")
	case "rollingupdate":
		c.rollUpdateService()
	case "scale":
		c.scaleService()
	case "expansion":
		c.expansionService()
	default:
		c.Response(200, nil)
	}
}

func (c *ServiceController) rollUpdateService() {
	clusterID, namespace, name, image := c.GetString(":cluster"), c.GetString(":namespace"), c.GetString(":service"), c.GetString("image")
	deploy, err := base.GetDeployment(name, namespace, clusterID)
	if err != nil {
		c.Response500(fmt.Errorf("when roll update, get service %v err: %v", name, err))
		return
	}
	deploy.Spec.Template.Spec.Containers[0].Image = image
	result, err := base.UpdateDeployment(deploy, clusterID)
	if err != nil {
		c.Response500(fmt.Errorf("roll update service %v err: %v", name, err))
		return
	}
	c.Response(200, result)
}

func (c *ServiceController) expansionService() {
	clusterID, namespace, name, cpu, memory := c.GetString(":cluster"), c.GetString(":namespace"), c.GetString(":service"), c.GetString("cpu"), c.GetString("memory")
	deploy, err := base.GetDeployment(name, namespace, clusterID)
	if err != nil {
		c.Response500(fmt.Errorf("when expansion service, get service %v err: %v", name, err))
		return
	}
	deploy.Spec.Template.Spec.Containers[0].Resources = v1.ResourceRequirements{
		Limits: v1.ResourceList{
			v1.ResourceCPU:    resource.MustParse(cpu),    //TODO 根据前端传入的值做资源限制
			v1.ResourceMemory: resource.MustParse(memory), //TODO 根据前端传入的值做资源限制
		},
		Requests: v1.ResourceList{
			v1.ResourceCPU:    resource.MustParse(cpu),
			v1.ResourceMemory: resource.MustParse(memory),
		},
	}
	result, err := base.UpdateDeployment(deploy, clusterID)
	if err != nil {
		c.Response500(fmt.Errorf("expansion service %v err: %v", name, err))
		return
	}
	c.Response(200, result)
}

func (c *ServiceController) scaleService() {
	clusterID := c.GetString(":cluster")
	namespace := c.GetString(":namespace")
	name := c.GetString(":service")
	replicas := c.GetString("replicas")
	deploy, err := base.GetDeployment(name, namespace, clusterID)
	if err != nil {
		c.Response500(fmt.Errorf("when scale service, get service %v err: %v", name, err))
		return
	}
	deploy.Spec.Replicas = parse.StringToInt32Pointer(replicas)
	result, err := base.UpdateDeployment(deploy, clusterID)
	if err != nil {
		c.Response500(fmt.Errorf("scale service %v err: %v", name, err))
		return
	}
	c.Response(200, result)
}

func (c *ServiceController) addServicePorts() {
	clusterID := c.GetString(":cluster")
	namespace := c.GetString(":namespace")
	name := c.GetString(":service")
	servicePorts, err := validate.ValidatePorts(c.Ctx.Request)
	if err != nil {
		c.Response400(err)
		return
	}
	containerPorts := []v1.ContainerPort{}
	for _, p := range servicePorts {
		containerPorts = append(containerPorts, v1.ContainerPort{ContainerPort: int32(p.TargetPort.IntVal)})
	}
	deploy, err := base.GetDeployment(name, namespace, clusterID)
	if err != nil {
		c.Response500(fmt.Errorf("when add service  port, get service %v err: %v", name, err))
		return
	}
	if deploy == nil {
		c.Response500(fmt.Errorf("serivce %v not found", name))
		return
	}
	deploy.Spec.Template.Spec.Containers[0].Ports = containerPorts

	svc, err := base.GetK8SService(name, namespace, clusterID)
	svc.Spec.Ports = servicePorts
	resultsvc, resultdp, err := base.UpdateService(svc, deploy, clusterID)
	if err != nil {
		c.Response500(fmt.Errorf("add service %v's port err: %v", name, err))
		return
	}
	c.Response(200, map[string]interface{}{"serivce": resultsvc, "deploy": resultdp})
}

func (c *ServiceController) addServiceEnvs() {
	clusterID := c.GetString("cluster")
	namespace := c.GetString("namespace")
	name := c.GetString("name")
	servicePorts, err := validate.ValidatePorts(c.Ctx.Request)
	if err != nil {
		c.Response400(err)
		return
	}
	containerPorts := []v1.ContainerPort{}
	for _, p := range servicePorts {
		containerPorts = append(containerPorts, v1.ContainerPort{ContainerPort: int32(p.TargetPort.IntVal)})
	}
	deploy, err := base.GetDeployment(name, namespace, clusterID)
	if err != nil {
		c.Response500(fmt.Errorf("when add service  env, get service %v err: %v", name, err))
		return
	}
	if deploy == nil {
		c.Response500(fmt.Errorf("serivce %v not found", name))
		return
	}
	deploy.Spec.Template.Spec.Containers[0].Ports = containerPorts
	result, err := base.UpdateDeployment(deploy, clusterID)
	if err != nil {
		c.Response500(fmt.Errorf("add service %v's env err: %v", name, err))
		return
	}
	c.Response(200, map[string]interface{}{"deploy": result})
}

// GetServiceMetric Get service Metric
// @Title GetServiceMetric server
// @Description Get service Metric
// @Success 200		{object}	map[string]interface{}
// @router /:service/metrics [get]
func (c *ServiceController) GetServiceMetric() {
	clusterID := c.GetString(":cluster")
	namespace := c.GetString(":namespace")
	serviceName := c.GetString(":service")
	metricsName := c.GetString("type")
	glog.Info(fmt.Sprintf("%v=%v", models.MinipaasServiceName, serviceName))
	listOptions := metav1.ListOptions{LabelSelector: fmt.Sprintf("%v=%v", models.MinipaasServiceName, serviceName)}
	podList, err := base.ListPod(namespace, clusterID, listOptions)
	if err != nil {
		glog.Errorf("when get service %v's metric, get service's pod err: %v", serviceName, err)
		c.Response500(fmt.Errorf("when get service %v's metric, get service's pod err: %v", serviceName, err))
		return
	}
	result := map[string]interface{}{}
	for _, pod := range podList {
		metrics, err := base.GetPodMetrics(namespace, pod.Name, metricsName)
		if err != nil {
			glog.Errorf("get container %v's metric %v err: %v", pod.Name, metricsName)
			c.Response500(err)
			return
		}
		result[pod.Name] = metrics
	}
	c.Response(200, result)
}

// GetOperation Get service operation
// @Title GetOperation server
// @Description Get service operation
// @Success 200		{object}	[]models.Audit
// @router /:service/audits [get]
func (c *ServiceController) GetOperation() {
	// clusterID := c.GetString(":cluster")
	// namespace := c.GetString(":namespace")
	// serviceName := c.GetString(":service")
	// new(models.Audit).GetAll()
}
