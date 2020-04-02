package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	DB *gorm.DB
)

func InitMySQL()(err error){
	dsn := "root:Zhaoxin..521@tcp(39.96.179.159:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	log.Info("数据库链接成功！")
	return DB.DB().Ping()
}

func Close(){
	DB.Close()
}