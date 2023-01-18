package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Database(construing string) {
	db, err := gorm.Open("mysql", construing)
	if err != nil {
		fmt.Println(err)
		panic("数据库连接错误")
	}
	fmt.Println("数据库连接成功")
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(20)  //连接池
	db.DB().SetMaxOpenConns(100) //连接数
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
}
