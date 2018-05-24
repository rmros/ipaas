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

import "ipaas/pkg/tools/storage/mysql"

//Insert insert storage to db
func (storage *Storage) Insert() error {
	return mysql.GetDB().Create(storage).Error
}

//GetByID get storage by id
func (storage *Storage) GetByID() (*Storage, error) {
	err := mysql.GetDB().First(storage, storage.ID).Error
	return storage, err
}

//GetAll get all storage
func (storage *Storage) GetAll() ([]*Storage, error) {
	var storages []*Storage
	err := mysql.GetDB().Find(&storages).Error
	return storages, err
}

//DeleteByName delete storage by name
func (storage *Storage) DeleteByName() error {
	return mysql.GetDB().Delete(Storage{}, "name=?", storage.Name).Error
}

//GetByNamespace get  storage by namespace
func (storage *Storage) GetByNamespace() ([]*Storage, error) {
	var storages []*Storage
	err := mysql.GetDB().Find(&storages, "namespace=?", storage.Namespace).Error
	return storages, err
}
