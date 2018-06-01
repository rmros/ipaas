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
	"ipaas/pkg/tools/validate"

	"github.com/golang/glog"
	"k8s.io/api/core/v1"
)

// StorageController storage controller
type StorageController struct {
	base.BaseController
}

// CreateStorage CreateStorage
// @Title CreateStorage server
// @Description create storage
// @Success 200		{object}	models.Storage
// @Param	body		body 	models.Storage		true	"body for user content"
// @router / [post]
func (c *StorageController) CreateStorage() {
	storage, err := validate.ValidateStorage(c.Ctx.Request)
	if err != nil {
		c.Response400(err)
		return
	}
	clusterID, namespace := c.GetString(":cluster"), c.GetString(":namespace")
	pvc := storage.TOK8SPersistentVolumeClaim(namespace)
	result, err := base.CreatePersistentVolumeClaim(pvc, clusterID)
	if err != nil {
		c.Response500(fmt.Errorf("create storage %v err: %v", storage.Name, err))
		return
	}
	c.Response(200, result)
}

// ListStorage ListStorage
// @Title ListStorage server
// @Description list storage
// @Success 200		{object}	[]models.Storage
// @router / [get]
func (c *StorageController) ListStorage() {
	clusterID, namespace := c.GetString(":cluster"), c.GetString(":namespace")
	pvcs, err := base.ListPersistentVolumeClaim(namespace, clusterID)
	if err != nil {
		glog.Errorf("list storage err: %v", err)
		c.Response500(fmt.Errorf("list storage err: %v", err))
		return
	}
	storages := []models.Storage{}
	for _, pv := range pvcs {
		storage := models.Storage{}
		storage.Name = pv.Name
		storage.Namespace = pv.Name
		temp := pv.Spec.Resources.Requests[v1.ResourceStorage]
		storage.Size = (&temp).String()
		storage.AccessModes = string(pv.Spec.AccessModes[0])
		storage.Status = string(pv.Status.Phase)
		storage.CreateAt = pv.CreationTimestamp.Time
		storages = append(storages, storage)
	}
	c.Response(200, storages)
}

// DeleteStorage DeleteStorage
// @Title DeleteStorage server
// @Description delete storage
// @Success 200
// @Param	names	body	[]string	true	"the storage names who need to delete"
// @router / [delete]
func (c *StorageController) DeleteStorage() {
	clusterID, namespace := c.GetString(":cluster"), c.GetString(":namespace")
	names, err := validate.Array(c.Ctx.Request)
	if err != nil {
		c.Response400(err)
		return
	}
	errs := []error{}
	for _, name := range names {
		if err = base.DeletePersistentVolumeClaim(name, namespace, clusterID); err != nil {
			glog.Errorf("delete strorage %v err: %v", name, err)
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		c.Response(200, errs)
		return
	}
	c.Response(200, "ok")
}
