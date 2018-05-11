package infrastructure

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
