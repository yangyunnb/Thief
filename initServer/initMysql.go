package initserver

import (
	"github.com/Thief.git/common"
	//nolint
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitMysql() {
	db, err := gorm.Open("mysql", Conf.Mysql.DSN)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(Conf.Mysql.MaxIDConn)
	db.DB().SetMaxOpenConns(Conf.Mysql.MaxOpenConn)
	common.SetDB(db)
}
