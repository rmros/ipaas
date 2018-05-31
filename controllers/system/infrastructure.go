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

package system

import (
	base "ipaas/controllers"
)

type SystemController struct {
	base.BaseController
}

// func (c *SystemController) ListCluster() {
// 	map[string]interface{}{}
// }

// func listSystemComponent() interface{} {
// 	component := interface{}
// 	return component
// }

// func listSystemCPU() interface{} {
// 	cpu := interface{}
// 	return cpu
// }

// func listSystemMemory() interface{} {
// 	memory := interface{}
// 	return memory
// }

// func listSystemPod(clusterID string) interface{} {
// 	clusterContainer := struct {
// 		Total     int64
// 		Operation int64
// 		Error     int64
// 	}{}

// 	pod := map[string]interface{}{}
// 	return pod
// }

// // map[string]
// func queryBasic() map[string]interface{} {
// 	basic := map[string]interface{}{}
// 	return basic
// }

// clusterNodes := struct {
// 	Total     int
// 	Scheduler int
// 	Heathy    int
// }{
// 	Total:     len(nodeList.Items),
// 	Scheduler: len(nodeList.Items),
// 	Heathy:    len(nodeList.Items),
// }

// clusterCpu := struct {
// 	Total            int64
// 	Allocated        int64
// 	AllocatedPersent int64 //需要前端自己去算百分比
// }{}

// clusterMemory := struct {
// 	Total            int64
// 	Allocated        int64
// 	AllocatedPersent int64 //需要前端自己去算百分比
// }{}

// clusterContainer := struct {
// 	Total     int64
// 	Operation int64
// 	Error     int64
// }{}
