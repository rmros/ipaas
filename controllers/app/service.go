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

package app

import (
	base "ipaas/controllers"
)

// ServiceController the service controller
type ServiceController struct {
	base.BaseController
}

// CreateService CreateService
// @Title CreateService server
// @Description create app
// @Success 201		{object}	models.Service
// @Param	body		body 	models.Service		true	"body for user content"
// @router /services [post]
func (c *ServiceController) CreateService() {

}

// DeleteService DeleteService
// @Title DeleteService server
// @Description create namespace
// @Success 201		{object}	models.Service
// @Param	body		body 	models.Service		true	"body for user content"
// @router /services/service [delete]
func (c *ServiceController) DeleteService() {

}

// StartService StartService
// @Title StartService server
// @Description start app
// @Success 201		{object}	models.Service
// @Param	body		body 	models.Service		true	"body for user content"
// @router /services/:service/start [put]
func (c *ServiceController) StartService() {

}

// ReStartService ReStartService
// @Title StartService server
// @Description start app
// @Success 201		{object}	models.Service
// @Param	body		body 	models.Service		true	"body for user content"
// @router /services/:service/start [put]
func (c *ServiceController) ReStartService() {

}

// StopService StopService
// @Title StopService server
// @Description stop app
// @Success 201		{object}	models.Service
// @Param	body		body 	models.Service		true	"body for user content"
// @router /services/service/stop [put]
func (c *ServiceController) StopService() {

}

// ListService ListService
// @Title ListService server
// @Description stop app
// @Success 201		{object}	models.Service
// @Param	body		body 	models.Service		true	"body for user content"
// @router /services [get]
func (c *ServiceController) ListService() {

}

// ReDeployService ReDeployService
// @Title ReDeployService server
// @Description ReDeploy app
// @Success 201		{object}	models.Service
// @Param	body		body 	models.Service		true	"body for user content"
// @router /services/:service/redeploy [put]
func (c *ServiceController) ReDeployService() {

}
