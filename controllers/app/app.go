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

type AppController struct {
	base.BaseController
}

// CreateApp CreateApp
// @Title CreateApp server
// @Description create app
// @Success 201		{object}	models.App
// @Param	body		body 	models.App		true	"body for user content"
// @router /apps [post]
func (c *AppController) CreateApp() {

}

// DeleteApp DeleteApp
// @Title DeleteApp server
// @Description create namespace
// @Success 201		{object}	models.App
// @Param	body		body 	models.App		true	"body for user content"
// @router /apps/:app [delete]
func (c *AppController) DeleteApp() {

}

// StartApp StartApp
// @Title StartApp server
// @Description start app
// @Success 201		{object}	models.App
// @Param	body		body 	models.App		true	"body for user content"
// @router /apps/:app/start [put]
func (c *AppController) StartApp() {

}

// StopApp StopApp
// @Title StopApp server
// @Description stop app
// @Success 201		{object}	models.App
// @Param	body		body 	models.App		true	"body for user content"
// @router /apps/:app/stop [put]
func (c *AppController) StopApp() {

}

// ListApp ListApp
// @Title ListApp server
// @Description stop app
// @Success 201		{object}	models.App
// @Param	body		body 	models.App		true	"body for user content"
// @router /apps [get]
func (c *AppController) ListApp() {

}

// ReDeployApp ReDeployApp
// @Title ReDeployApp server
// @Description ReDeploy app
// @Success 201		{object}	models.App
// @Param	body		body 	models.App		true	"body for user content"
// @router /apps/:app/redeploy [put]
func (c *AppController) ReDeployApp() {

}
