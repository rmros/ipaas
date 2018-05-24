// Copyright © 2018 huang jia <449264675@qq.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"ipaas/pkg/tools/storage/mysql"
)

//GetAll get all app
func (pc *PodLifeCycle) GetAll() ([]*PodLifeCycle, error) {
	var pcs []*PodLifeCycle
	err := mysql.GetDB().Find(&pcs).Error
	return pcs, err
}

// GetByID get  PodLifeCycle  by clusterID
func (pc *PodLifeCycle) GetByID(ID string) (*PodLifeCycle, error) {
	err := mysql.GetDB().Where("cluster_id=?", pc.ClusterID).Find(pc).Error
	return pc, err
}

// Insert insert PodLifeCycle to db
func (pc *PodLifeCycle) Insert() (*PodLifeCycle, error) {
	err := mysql.GetDB().Create(pc).Error
	return pc, err
}
