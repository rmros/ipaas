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

// ConfigController the config controller
type ConfigController struct {
	base.BaseController
}

// CreateConfig CreateConfig
// @Title CreateConfig server
// @Description create config
// @Success 200		{object}	models.Config
// @Param	body		body 	models.Config		true	"body for user content"
// @router / [post]
func (c *ConfigController) CreateConfig() {
	config, err := validate.ValidateConfig(c.Ctx.Request)
	if err != nil {
		c.Response400(err)
		return
	}
	clusterID := c.GetString(":cluster")
	namespace := c.GetString(":namespace")
	configMap := config.TOK8SConfigMap(namespace)
	result, err := base.CreateConfigMap(clusterID, configMap)
	if err != nil {
		c.Response500(fmt.Errorf("create config %v err: %v", config.Name, err))
		return
	}
	c.Response(200, result)
}

// DeleteConfig DeleteConfig
// @Title DeleteConfig server
// @Description delete config
// @Success 200
// @router /:config [delete]
func (c *ConfigController) DeleteConfig() {
	clusterID := c.GetString(":cluster")
	namespace := c.GetString(":namespace")
	name := c.GetString(":config")
	if err := base.DeleteConfigMap(name, namespace, clusterID); err != nil {
		glog.Errorf("delete configMap %v err: %v", name, err)
		c.Response500(fmt.Errorf("delete config %v err: %v", name, err))
		return
	}
	c.Response(200, "ok")
}

// AddConfigData add config data
// AddConfigData AddConfigData
// @Title AddConfigData server
// @Description add config data
// @Success 200
// @Param	body		body 	map[string]string		true	"body for user content"
// @router /:config [put]
func (c *ConfigController) AddConfigData() {
	data, err := validate.ValidateConfigData(c.Ctx.Request)
	if err != nil {
		c.Response400(err)
		return
	}
	clusterID := c.GetString(":cluster")
	namespace := c.GetString(":namespace")
	name := c.GetString(":config")
	configMap, err := base.GetConfigMapByName(name, namespace, clusterID)
	if err != nil {
		c.Response500(fmt.Errorf("when add config data, get config %v err: %v", name, err))
		return
	}
	for k, v := range data {
		if configMap.Data != nil {
			configMap.Data[k] = v
		} else {
			configMap.Data = data
		}
	}
	result, err := base.UpdateConfigMap(clusterID, configMap)
	if err != nil {
		c.Response500(fmt.Errorf("when add config data, get config %v err: %v", name, err))
		return
	}
	c.Response(200, result)
}

// DeleteConfigData delete config data
// AddConfigData AddConfigData
// @Title AddConfigData server
// @Description delete config data
// @Success 200
// @router /:config/:key [delete]
func (c *ConfigController) DeleteConfigData() {
	clusterID := c.GetString(":cluster")
	namespace := c.GetString(":namespace")
	name := c.GetString(":config")
	key := c.GetString(":key")
	configMap, err := base.GetConfigMapByName(name, namespace, clusterID)
	if err != nil {
		c.Response500(fmt.Errorf("when add config data, get config %v err: %v", name, err))
		return
	}
	if configMap.Data != nil {
		delete(configMap.Data, key)
	}
	result, err := base.UpdateConfigMap(clusterID, configMap)
	if err != nil {
		c.Response500(fmt.Errorf("when add config data, get config %v err: %v", name, err))
		return
	}
	c.Response(200, result)
}

// ListConfig ListConfig
// @Title ListConfig server
// @Description list config
// @Success 200		{object}	[]models.Storage
// @router / [get]
func (c *ConfigController) ListConfig() {
	clusterID, namespace := c.GetString(":cluster"), c.GetString(":namespace")
	configMaps, err := base.ListConfigMap(namespace, clusterID)
	if err != nil {
		glog.Errorf("list storage err: %v", err)
		c.Response500(fmt.Errorf("list storage err: %v", err))
		return
	}
	configs := []models.Config{}
	for _, cm := range configMaps {
		config := models.Config{}
		config.Name = cm.Name
		config.Namespace = cm.Namespace
		config.Data = cm.Data
		configs = append(configs, config)
	}
	c.Response(200, configs)
}
