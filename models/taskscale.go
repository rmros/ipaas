// Copyright © 2017 huang jia <449264675@qq.com>
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

// Insert insert task
func (task *TickScaleTask) Insert() error {
	return mysql.GetDB().Create(task).Error
}

// GetAll get all task
func (task *TickScaleTask) GetAll() ([]*TickScaleTask, error) {
	var tasks []*TickScaleTask
	err := mysql.GetDB().Find(&tasks).Error
	return tasks, err
}

// Delete delete task
func (task *TickScaleTask) Delete() error {
	return mysql.GetDB().Delete(task).Error
}

// Get get all task
func (task *TickScaleTask) Get() (*TickScaleTask, error) {
	var tk *TickScaleTask
	err := mysql.GetDB().Where("id=?", task.ID).Find(tk).Error
	return tk, err
}

// Update update task
func (task *TickScaleTask) Update() error {
	return mysql.GetDB().Update(task).Error
}
