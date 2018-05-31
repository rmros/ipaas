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

	"github.com/golang/glog"
)

// AppController app api server
type AppController struct {
	base.BaseController
}

// CreateApp CreateApp
// @Title CreateApp server
// @Description create app
// @Success 200		{object}	models.App
// @Param	body		body 	models.App		true	"body for user content"
// @router / [post]
func (c *AppController) CreateApp() {
	app, err := validate.ValidateApp(c.Ctx.Request)
	if err != nil {
		c.Response400(err)
		return
	}
	if len(app.Items) == 0 {
		c.Response400(fmt.Errorf("app %v's service mustn't null", app.Name))
		return
	}
	svc := app.Items[0]
	svc.AppName = app.Name
	clusterID := c.GetString(":cluster")
	namespace := c.GetString(":namespace")
	service := svc.TOK8SService(namespace)
	deployment := svc.TOK8SDeployment(namespace)
	result, err := base.DelpoyService(clusterID, service, deployment)
	if err != nil {
		glog.Errorf("deploy app where named %q err: ", err)
		c.Response500(fmt.Errorf("deploy app where named %q err: ", err))
		return
	}

	app.ServiceCount = svc.InstanceCount
	app.InstanceCount = int(*(deployment.Spec.Replicas))
	if err = app.Insert(); err != nil {
		glog.Errorf("record to db err: %v", err)
		c.Response500(fmt.Errorf("record to db err: %v", err))
		go base.DeleteService(svc.Name, namespace, clusterID)
		return
	}
	c.Response(200, result)
}

// DeleteApp DeleteApp
// @Title DeleteApp server
// @Description delete config
// @Success 200
// @router /:app [delete]
func (c *AppController) DeleteApp() {
	clusterID := c.GetString(":cluster")
	namespace := c.GetString(":namespace")
	name := c.GetString(":app")
	if err := base.DeleteServiceByAppName(name, namespace, clusterID); err != nil {
		glog.Errorf("delete app %v err: %v", name, err)
		c.Response500(fmt.Errorf("delete app %v err: %v", name, err))
		return
	}
	app := &models.App{}
	app.UserName = namespace
	app.Name = name
	if err := app.DeleteByNameAndNamespace(); err != nil {
		glog.Errorf("delete %v's app %v from db err: %v", namespace, name, err)
		c.Response500(fmt.Errorf("delete %v's app %v from db err: %v", namespace, name, err))
		return
	}
	c.Response(200, "ok")
}

// OperationApp OperationApp
// @Title OperationApp server
// @Description start stop redploy app
// @Success 200
// @router /:app/:verb [put]
func (c *AppController) OperationApp() {
	clusterID, namespace, name := c.GetString(":cluster"), c.GetString(":namespace"), c.GetString(":service")
	verb := c.GetString(":verb")
	svcs, deploys, err := base.ListServiceByAppName(name, namespace, clusterID)
	if err != nil {
		glog.Errorf("%v app %v err: %v", verb, name, err)
		c.Response500(fmt.Errorf("%v app %v err: %v", verb, name, err))
		return
	}

	errs := base.OperatorServices(svcs, deploys, verb, clusterID)
	if len(errs) != 0 {
		glog.Errorf("%v app %v err: %v", verb, name, err)
		c.Response500(fmt.Errorf("%v app %v err: %v", verb, name, errs))
		return
	}
	c.Response(200, "ok")
}

// ListApp ListApp
// @Title ListApp server
// @Description stop app
// @Success 200		{object}	[]models.App
// @router / [get]
func (c *AppController) ListApp() {
	apps, err := new(models.App).GetAll()
	if err != nil {
		glog.Errorf("get apps err: %v", err)
		c.Response500(fmt.Errorf("get apps err: %v", err))
		return
	}
	c.Response(200, apps)
}
