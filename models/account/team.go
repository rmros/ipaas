/*
Package account include user、team、team model's basic operation of database
*/
package account

import "ipaas/pkg/tools/storage/mysql"

// Create insert team to db
func (team *Team) Create() error {
	return mysql.GetDB().Create(team).Error
}

// Get get one team by id
func (team *Team) Get() error {
	return mysql.GetDB().Where("team_id=?", team.ID).First(team).Error
}

// GetTeamUsers get team's users by teamID
func (team *Team) GetTeamUsers() ([]*Team, error) {
	var teams []*Team
	err := mysql.GetDB().Where("team_id=?", team.ID).Related(&teams).Error
	return teams, err
}

// Update update team
func (team *Team) Update() error {
	return mysql.GetDB().Model(team).Updates(team).Error
}

// Delete delete team from db
func (team *Team) Delete() error {
	return mysql.GetDB().Delete(team).Error
}

// ListAll get all team from db
func (team *Team) ListAll() ([]*Team, error) {
	var teams []*Team
	err := mysql.GetDB().Find(&teams).Error
	return teams, err
}
