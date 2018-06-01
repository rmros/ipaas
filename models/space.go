/*
Copyright 2018 huangjia.

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
	"fmt"
	"ipaas/pkg/tools/storage/mysql"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func (space *Space) TOK8sNamespace() *v1.Namespace {
	labels := map[string]string{}
	if space.Type != 0 {
		labels["type"] = fmt.Sprintf("%v", space.Type)
	}
	if space.TeamID != "" {
		labels["teamID"] = space.TeamID
	}
	return &v1.Namespace{
		TypeMeta: metav1.TypeMeta{Kind: "Namespace", APIVersion: "v1"},
		ObjectMeta: metav1.ObjectMeta{
			Name:   space.Name,
			Labels: labels,
		},
	}
}
