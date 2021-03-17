package common

import "github.com/jinzhu/gorm"

var dbIns *gorm.DB

func GetDB() *gorm.DB {
	return dbIns
}

func SetDB(db *gorm.DB) {
	dbIns = db
}
