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

package app

import (
	"fmt"

	base "ipaas/controllers"
	"ipaas/models"
	"ipaas/pkg/tools/validate"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/golang/glog"
)

// ContainerController the container controller
type ContainerController struct {
	base.BaseController
}

// ListContainer query all pod in current namespace
// @Title ListContainer server
// @Description query all pod in the current namespace
// @Success 200		{object}	[]models.Container
// @router / [get]
func (c *ContainerController) ListContainer() {
	clusterID, namespace := c.GetString(":cluster"), c.GetString(":namespace")
	pods, err := base.ListPod(namespace, clusterID, metav1.ListOptions{})
	if err != nil {
		glog.Errorf("list container in namespace [%v] err: %v", namespace, err)
		c.Response500(fmt.Errorf("list container in namespace [%v] err: %v", namespace, err))
		return
	}
	containers := []*models.Container{}
	for _, pod := range pods {
		container := &models.Container{}
		container.AppName = pod.Labels[models.MinipaasAppName]
		container.CreateAt = pod.CreationTimestamp.Time
		container.Image = pod.Spec.Containers[0].Image
		container.Name = pod.Name
		container.Status = string(pod.Status.Phase)
		container.URL = pod.Status.PodIP
		containers = append(containers, container)
	}
	c.Response(200, containers)
}

// ReCreateContainer recreate pod
// @Title ReCreateContainer server
// @Description recreate pod
// @Success 200
// @Param	names	body	[]string	true	"the storage names who need to delete"
// @router / [put]
func (c *ContainerController) ReCreateContainer() {
	clusterID, namespace := c.GetString(":cluster"), c.GetString(":namespace")
	names, err := validate.Array(c.Ctx.Request)
	if err != nil {
		c.Response400(err)
		return
	}
	errs := []error{}
	for _, name := range names {
		if err := base.DeletePod(name, namespace, clusterID); err != nil {
			glog.Errorf("recreate pod %v err: %v", name, err)
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		c.Response(200, errs)
		return
	}
	c.Response(200, "ok")
}
