/*
Package account include user、team、space model's basic operation of database
*/
package account

import (
	"ipaas/pkg/tools/storage/mysql"
)

// Create insert user to db
func (user *User) Create() error {
	return mysql.GetDB().Create(user).Error
}

// Get get one user by id
func (user *User) Get() error {
	return mysql.GetDB().First(user, user.ID).Error
}

// GetByNameAndPassword get user by name and password
func (user *User) GetByNameAndPassword() error {
	return mysql.GetDB().Where("name=? and password=?", user.Name, user.Password).First(user).Error
}

// Update update user
func (user *User) Update() error {
	return nil
}

// Delete delete user from db
func (user *User) Delete() error {
	return nil
}

// ListAll get all user from db
func (user *User) ListAll() ([]*User, error) {
	var users []*User
	err := mysql.GetDB().Find(&users).Error
	return users, err
}
