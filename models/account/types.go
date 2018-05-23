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
Package account include user、team、space model's basic operation of database
*/
package account

import (
	"time"

	"ipaas/models/infrastructure"
	"ipaas/pkg/tools/storage/mysql"
)

// Role defines some enums for the role field in users
type Role int

const (
	// RoleNormal the normal user
	RoleNormal = iota
	// RoleTeam the team manager
	RoleTeam
	// RoleAdmin the system manager
	RoleAdmin
	// RoleSuper the system super manager
	RoleSuper
)

const (
	// SuperUserName default super user name for now
	SuperUserName = "admin"
)

// User user info
type User struct {
	ID             int32 `gorm:"primary_key"`
	Name           string
	Displayname    string
	Password       string
	Email          string
	Phone          string
	CreationTime   time.Time
	LastLoginTime  time.Time
	LoginFrequency int
	Active         int8
	APIToken       string
	Role           int32
	Type           int
	Company        string
	Teams          []*Team `gorm:"many2many:user_teams;"`
}

// TableName return user model's  table name
func (user *User) TableName() string {
	return "users"
}

// Team team info
type Team struct {
	ID           string `gorm:"primary_key"`
	Name         string
	Description  string
	CreatorID    int32
	CreationTime time.Time
	Users        []*User `gorm:"many2many:team_users;"`
}

// TableName return team model's  table name
func (team *Team) TableName() string {
	return "teams"
}

// Space space info
type Space struct {
	ID           string    `json:"id,omitempty" gorm:"primary_key"`
	Name         string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	TeamID       string    `json:"teamID,omitempty"`
	CreationTime time.Time `json:"creationTime,omitempty"`
	Type         int       `json:"type,omitempty"` // 1 personal namespace 2 team's namespace
}

// TableName return Space model's  table name
func (space *Space) TableName() string {
	return "spaces"
}

// Company company info
type Company struct {
	ID          string `gorm:"primary_key"`
	Name        string
	Description string
}

// TableName return Company model's  table name
func (company *Company) TableName() string {
	return "companys"
}

func init() {
	mysql.GetDB().SingularTable(true)
	mysql.GetDB().CreateTable(
		new(User),
		new(Team),
		new(Space),
		new(infrastructure.Cluster),
	)
}
