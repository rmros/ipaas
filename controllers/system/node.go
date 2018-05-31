package system

import (
	"fmt"

	base "ipaas/controllers"
	"ipaas/models"
	"ipaas/pkg/tools/validate"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/golang/glog"
)

// NodeController node api server
type NodeController struct {
	base.BaseController
}

// Scheduler Scheduler pod
// @Title ReCreateContainer server
// @Description make node to scheduleable or not
// @Success 200
// @Param	scheduleable	query	bool	true	"true or false"
// @router /:node [put]
func (c *NodeController) Scheduler() {
	clusterID, name := c.GetString(":cluster"), c.GetString(":node")
	scheduleable, err := c.GetBool("scheduleable", false)
	if err != nil {
		c.Response400(err)
		return
	}

	node, err := base.GetNode(name, clusterID)
	if err != nil {
		glog.Errorf("when scheduler node [%v], get node err: %v", err, name)
		c.Response500(err)
		return
	}

	node.Spec.Unschedulable = scheduleable
	node, err = base.UpdateNode(node, clusterID)
	if err != nil {
		c.Response500(err)
		return
	}
	c.Response(200, "ok")
}

// LabelOperator LabelOperator pod
// @Title LabelOperator server
// @Description add delete label
// @Success 200
// @Param	scheduleable	query	bool	true	"schedule or not"
// @router /:node/labels/:verb [put]
func (c *NodeController) LabelOperator() {
	clusterID, name, verb := c.GetString(":cluster"), c.GetString(":node"), c.GetString(":verb")

	switch verb {
	case "add":
		labels, err := validate.Map(c.Ctx.Request)
		if err != nil {
			c.Response400(err)
			return
		}
		node, err := base.GetNode(name, clusterID)
		if err != nil {
			glog.Errorf("when add node [%v] label, get node err: %v", err, name)
			c.Response500(err)
			return
		}
		for k, v := range labels {
			glog.Info(k, "=", v)
			if node.Labels != nil {
				node.Labels[k] = v
			} else {
				node.Labels = labels
			}
		}
		node, err = base.UpdateNode(node, clusterID)
		if err != nil {
			glog.Errorf("when add node [%v] label, update node err: %v", err, name)
			c.Response500(err)
			return
		}
	case "delete":
		node, err := base.GetNode(name, clusterID)
		keys, err := validate.Array(c.Ctx.Request)
		if err != nil {
			glog.Errorf("when add node [%v] label, get node err: %v", err, name)
			c.Response500(err)
			return
		}
		for _, v := range keys {
			if node.Labels != nil {
				delete(node.Labels, v)
			}
		}
		node, err = base.UpdateNode(node, clusterID)
		if err != nil {
			glog.Errorf("when add node [%v] label, update node err: %v", err, name)
			c.Response500(err)
			return
		}
	}

	c.Response(200, "ok")
}

// ListContainer list node pod
// @Title ListContainer server
// @Description list node po
// @Success 200		{object}	[]models.Container
// @router /:node/containers [get]
func (c *NodeController) ListContainer() {
	clusterID, name := c.GetString(":cluster"), c.GetString(":node")
	listOptions := metav1.ListOptions{FieldSelector: fmt.Sprintf("spec.nodeName=%v", name)}
	pods, err := base.ListPod(v1.NamespaceAll, clusterID, listOptions)
	if err != nil {
		glog.Errorf("list container in node [%v] err: %v", name, err)
		c.Response500(fmt.Errorf("list container in node [%v] err: %v", name, err))
		return
	}
	containers := []*models.Container{}
	for _, pod := range pods {
		container := &models.Container{}
		container.AppName = pod.Labels[models.MinipaasAppName]
		container.Namespace = pod.Namespace
		container.CreateAt = pod.CreationTimestamp.Time
		container.Image = pod.Spec.Containers[0].Image
		container.Name = pod.Name
		container.Status = string(pod.Status.Phase)
		container.URL = pod.Status.PodIP
		containers = append(containers, container)
	}
	c.Response(200, containers)
}

// func (c *NodeController) GetMetric() {
// 	clusterID := c.GetString(":cluster")
// }

// ListNode list node
// @Title ListNode server
// @Description list node
// @Success 200		{object}	[]*models.Node
// @router / [get]
func (c *NodeController) ListNode() {
	clusterID := c.GetString(":cluster")
	nodeList, err := base.ListNode(clusterID, metav1.ListOptions{})
	if err != nil {
		glog.Errorf("list cluster [%v] node err: %v", clusterID, err)
		c.Response500(fmt.Errorf("list cluster [%v] node err: %v", clusterID, err))
		return
	}

	var nodes []*models.Node
	for _, item := range nodeList {
		n := base.TranslateK8sNode(clusterID, item)
		nodes = append(nodes, n)
	}
	c.Response(200, nodes)
}

// GetNode get node
// @Title GetNode server
// @Description get node
// @Success 200		{object}	models.Node
// @router /:name [get]
func (c *NodeController) GetNode() {
	clusterID, name := c.GetString(":cluster"), c.GetString(":name")
	knode, err := base.GetNode(name, clusterID)
	if err != nil {
		glog.Errorf("get cluster [%v] node [%v] err: %v", clusterID, name, err)
		c.Response500(fmt.Errorf("get cluster [%v] node [%v] err: %v", clusterID, name, err))
		return
	}
	node := base.TranslateK8sNode(clusterID, *knode)
	c.Response(200, node)
}
