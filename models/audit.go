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

import (
	"ipaas/pkg/tools/storage/mysql"
)

// Insert insert audit
func (audit *Audit) Insert() error {
	return mysql.GetDB().Create(audit).Error
}

// GetAll get all audit
func (audit *Audit) GetAll() ([]*Audit, error) {
	var audits []*Audit
	err := mysql.GetDB().Find(&audits).Error
	return audits, err
}

// Delete delete audit
func (audit *Audit) Delete() error {
	return mysql.GetDB().Delete(audit).Error
}

// Get get all audit
func (audit *Audit) Get() (*Audit, error) {
	var tk *Audit
	err := mysql.GetDB().Where("id=?", audit.ID).Find(tk).Error
	return tk, err
}

// Update update audit
func (audit *Audit) Update() error {
	return mysql.GetDB().Update(audit).Error
}
