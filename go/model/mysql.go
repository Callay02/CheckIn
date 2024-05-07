package model

import (
	"checkin/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", config.Application.Database.Mysql.Username, config.Application.Database.Mysql.Password, config.Application.Database.Mysql.Host, config.Application.Database.Mysql.Port, config.Application.Database.Mysql.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库:" + err.Error())
	}
	DB = db
}
