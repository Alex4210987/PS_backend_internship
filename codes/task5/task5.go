package task5

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Preference struct {
	ID         int `gorm:"primaryKey"`
	UserID_P   int `gorm:"column:userid_p"`
	Mode       string
	OutputMode string `gorm:"column:outputmode"`
	Tactics    string
}

type Route struct {
	ID         int    `gorm:"primaryKey"`
	StartPoint string `gorm:"column:startpoint"`
	EndPoint   string `gorm:"column:endpoint"`
}

type Alias struct {
	ID       int `gorm:"primaryKey"`
	UserID_A int
	Location string
	Alias    string
}
type User struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

func DbManager() {
	// 连接到 MySQL 服务器
	dsn := "root:123456@tcp(127.0.0.1:3306)/"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建数据库
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS task5")
	if err != nil {
		log.Fatal(err)
	}

	// 连接到 task5 数据库
	dsn = "root:123456@tcp(127.0.0.1:3306)/task5"
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 创建表
	err = gormDB.AutoMigrate(&User{}, &Route{}, &Preference{}, &Alias{})
	if err != nil {
		log.Fatal(err)
	}

	// 创建用户
	user := User{ID: 1, Name: "default"}
	result := gormDB.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}
func CloseDB() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/task5"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 删除表
	err = db.Migrator().DropTable(&User{}, &Route{}, &Preference{}, &Alias{})
	if err != nil {
		log.Fatal(err)
	}
}
