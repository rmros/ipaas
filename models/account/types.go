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
	ID           string `gorm:"primary_key"`
	Name         string
	Description  string
	TeamID       string
	CreationTime time.Time
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
