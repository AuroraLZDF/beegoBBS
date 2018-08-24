package models

import (
	_ "github.com/go-sql-driver/mysql"//加载mysql
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

// 基本模型的定义
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

var cfg = beego.AppConfig

func DB() *gorm.DB {
	// database config
	dbUser := cfg.String("mysql_user")
	dbPass := cfg.String("mysql_pass")
	dbHost := cfg.String("mysql_host")
	dbPort := cfg.String("mysql_port")
	dbName := cfg.String("mysql_db")

	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10ms", dbUser, dbPass, dbHost, dbPort, dbName) + "&loc=Asia%2FShanghai"

	db, err := gorm.Open("mysql", dbLink)

	//defer db.Close()
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	return db
}
