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

/*
Package account include user、team、space model's basic operation of database
*/
package models

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
	return mysql.GetDB().Model(user).Where("name=?", user.Name).Update(user).Error
}

// Delete delete user from db
func (user *User) Delete() error {
	return mysql.GetDB().Delete(user, "name=?", user.Name).Error
}

// ListAll get all user from db
func (user *User) ListAll() ([]*User, error) {
	var users []*User
	err := mysql.GetDB().Find(&users).Error
	return users, err
}

func (user *User) Exsit() bool {
	return !mysql.GetDB().Model(user).Where("name=?", user.Name).RecordNotFound()
}
