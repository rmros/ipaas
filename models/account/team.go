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
