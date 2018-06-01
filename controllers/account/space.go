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

/*
Package account include user、team、space's controller basic operation of logic
*/
package account

import (
	base "ipaas/controllers"
	"ipaas/pkg/tools/validate"
)

// SpaceController space controller
type SpaceController struct {
	base.BaseController
}

// CreateSpace CreateSpace
// @Title CreateSpace server
// @Description create namespace
// @Success 201		{object}	models.Space
// @Param	body		body 	models.Space		true	"body for user content"
// @router / [post]
func (c *SpaceController) CreateSpace() {
	space, err := validate.ValidateSpace(c.Ctx.Request)
	if err != nil {
		c.Response400(err)
		return
	}
	ns := space.TOK8sNamespace()
	if err = base.CreateNamespace(c.GetString(":cluster"), ns); err != nil {
		c.Response500(err)
		return
	}
	c.Response(201, "ok")
}

// DeleteSpace DeleteSpace
// @Title DeleteSpace server
// @Description delete namespace
// @Success 200
// @router /:namespace [delete]
func (c *SpaceController) DeleteSpace() {
	name := c.GetString(":namespace")
	cluster := c.GetString(":cluster")
	if err := base.DeleteNamespace(cluster, name, ""); err != nil {
		c.Response500(err)
		return
	}
	c.Response(200, "ok")
}

// ListSpace ListSpace
// @Title ListSpace server
// @Description list namespace
// @Success 200		{object}	[]models.Space
// @router / [get]
func (c *SpaceController) ListSpace() {
	cluster := c.GetString(":cluster")
	namespaces, err := base.ListNamespace(cluster)
	if err != nil {
		c.Response500(err)
		return
	}
	c.Response(200, base.DecodeNamespaces(namespaces))
}

// GetSpace GetSpace
// @Title GetSpace server
// @Description get namespace by name
// @Success 200		{object}	models.Space
// @router /:namespace [get]
func (c *SpaceController) GetSpace() {
	cluster := c.GetString(":cluster")
	name := c.GetString(":namespace")
	namespace, err := base.GetNamespace(cluster, name)
	if err != nil {
		c.Response500(err)
		return
	}
	c.Response(200, base.DecodeNamespace(namespace))
}
