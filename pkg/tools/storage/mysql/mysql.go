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
}

//GetDB return the *gorm.DB
func GetDB() *gorm.DB {
	return db
}
