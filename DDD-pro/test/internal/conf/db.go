package conf

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/course?charset=utf8mb4&parseTime=True&loc=Local",
		Config.DBUser,Config.DBPassword,Config.DBHost,Config.DBPort)
	fmt.Println(dsn)
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatalf("init mysql fail:%s",err)
	}

	mysqlDB,err:=db.DB()
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)
	DB = db
}