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

//Insert insert app to db
func (app *App) Insert() error {
	return mysql.GetDB().Create(app).Error
}

//GetByID get app by id
func (app *App) GetByID() (*App, error) {
	err := mysql.GetDB().First(app, app.ID).Error
	return app, err
}

//GetAll get all app
func (app *App) GetAll() ([]*App, error) {
	var apps []*App
	err := mysql.GetDB().Find(&apps).Error
	return apps, err
}

//GetByNamespace get  storage by namespace
func (app *App) GetByNamespace() ([]*App, error) {
	var apps []*App
	err := mysql.GetDB().Find(&apps, "user_name=?", app.UserName).Error
	return apps, err
}

// DeleteByNameAndNamespace delete app by name and namespace
func (app *App) DeleteByNameAndNamespace() error {
	return mysql.GetDB().Delete(App{}, "name=? and user_name=?", app.Name, app.UserName).Error
}
