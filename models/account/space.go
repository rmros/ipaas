/*
Package account include user、team、space model's basic operation of database
*/
package account

import (
	"ipaas/pkg/tools/storage/mysql"
)

// Create insert space to db
func (space *Space) Create() error {
	return mysql.GetDB().Create(space).Error
}

// Get get one space by id
func (space *Space) Get() error {
	return mysql.GetDB().First(space, space.ID).Error
}

// GetByTeamID get space by teamID
func (space *Space) GetByTeamID() ([]*Space, error) {
	var spaces []*Space
	err := mysql.GetDB().Where("team_id=?", space.TeamID).Find(&spaces).Error
	return spaces, err
}

// Update update space
func (space *Space) Update() error {
	return mysql.GetDB().Model(space).Updates(space).Error
}

// Delete delete space from db
func (space *Space) Delete() error {
	return mysql.GetDB().Delete(space).Error
}

// ListAll get all space from db
func (space *Space) ListAll() ([]*Space, error) {
	var spaces []*Space
	err := mysql.GetDB().Find(&spaces).Error
	return spaces, err
}
