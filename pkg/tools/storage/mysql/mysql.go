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

package mysql

import (
	"ipaas/pkg/tools/configz"
	"ipaas/pkg/tools/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

//Init init mysql database
func init() {
	db, err = gorm.Open("mysql", configz.GetString("mysql", "dsn", ""))
	if err != nil {
		log.Critical("init mysql connection err: %v", err)
	}
	db.LogMode(true)
}

//GetDB return the *gorm.DB
func GetDB() *gorm.DB {
	return db
}
