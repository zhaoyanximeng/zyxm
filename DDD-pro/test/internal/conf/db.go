package conf

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	dsn := "user1:!19981216sc@tcp(rm-wz9w7yc051ryt28xgvo.mysql.rds.aliyuncs.com:3306)/course?charset=utf8mb4&parseTime=True&loc=Local"
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	mysqlDB,err:=db.DB()
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)
	DB = db
}