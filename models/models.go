package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"ws/marketApi/pkg/setting"
)

var db *gorm.DB

func init() {
	db, err := gorm.Open(setting.DB.Type,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&collation=%s&parseTime=True&loc=Local",
			setting.DB.User,
			setting.DB.Password,
			setting.DB.Host,
			setting.DB.Port,
			setting.DB.Database,
			setting.DB.Charset,
			setting.DB.Collation))
	if err != nil {
		log.Panicf("数据库连接失败:%s", err)
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	log.Println("数据库链接成功")
}

func CloseDB() {
	defer db.Close()
}
