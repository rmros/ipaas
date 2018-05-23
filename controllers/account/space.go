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

/*
Package account include user、team、space's controller basic operation of logic
*/
package account

import (
	base "ipaas/controllers"
)

// SpaceController space controller
type SpaceController struct {
	base.BaseController
}

func (c *SpaceController) CreateSpace() {
	// space, err := validate.ValidateSpace(c.Ctx.Request)
	// if err != nil {
	// 	c.Response400(err)
	// 	return
	// }
	// ns := space.TOK8sNamespace()
	// createnamespace := func() {
	// 	for clusterID, client := range client.GetClientsets() {
	// 		glog.Info(clusterID)
	// 		_, err := v1.Namespaces(client.Clientset).Create(ns)
	// 		if err != nil {
	// 			glog.Errorf("when add user,create k8s namespace [%v] in cluster [%v] err: %v", ns.Name, clusterID, err)
	// 		}
	// 	}
	// }
	// go createnamespace()

}

func (c *SpaceController) DeleteSpace() {
	// name := c.GetString(":space")
	// deletenamespace := func() {
	// 	for clusterID, client := range client.GetClientsets() {
	// 		if err := v1.Namespaces(client.Clientset).Delete(name, &metav1.DeleteOptions{}); err != nil {
	// 			glog.Errorf("delete k8s namespace [%v] in cluster [%v] err: %v", name, clusterID, err)
	// 		}
	// 	}
	// }
	// go deletenamespace()
}

func (c *SpaceController) List() {

}
