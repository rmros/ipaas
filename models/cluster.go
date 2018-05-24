// Copyright © 2017 huang jia <449264675@qq.com>
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

// Create insert cluster to db
func (cluster *Cluster) Create() error {
	return mysql.GetDB().Create(cluster).Error
}

// Get get one cluster by id
func (cluster *Cluster) Get() error {
	return mysql.GetDB().First(cluster, cluster.ID).Error
}

// GetByTeamID get cluster by teamID
// func (cluster *Cluster) GetByTeamID() ([]*Cluster, error) {
// 	var clusters []*Cluster
// 	err := mysql.GetDB().Where("team_id=?", cluster.TeamID).Find(&clusters).Error
// 	return clusters, err
// }

// Update update cluster
func (cluster *Cluster) Update() error {
	return mysql.GetDB().Model(cluster).Updates(cluster).Error
}

// Delete delete cluster from db
func (cluster *Cluster) Delete() error {
	return mysql.GetDB().Delete(cluster).Error
}

// ListAll get all cluster from db
func (cluster *Cluster) ListAll() ([]*Cluster, error) {
	var clusters []*Cluster
	err := mysql.GetDB().Find(&clusters).Error
	return clusters, err
}
