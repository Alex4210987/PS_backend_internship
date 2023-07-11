package task5

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Route struct {
	ID         int `gorm:"primaryKey"`
	StartPoint string
	EndPoint   string
}

type Preference struct {
	ID         int `gorm:"primaryKey"`
	UserID_P   int
	Mode       string
	OutputMode string
	Tactics    string
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
	dsn := "root:123456@tcp(127.0.0.1:3306)/task5"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 创建表
	err = db.AutoMigrate(&User{}, &Route{}, &Preference{}, &Alias{})
	if err != nil {
		log.Fatal(err)
	}

	// 创建用户
	user := User{ID: 1, Name: "default"}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	dbTest(db)
}
func dbTest(db *gorm.DB) {
	// 插入数据
	db.Create(&Route{ID: 1, StartPoint: "A", EndPoint: "B"})
	db.Create(&Route{ID: 2, StartPoint: "B", EndPoint: "C"})

	db.Create(&Preference{ID: 1, UserID_P: 1, Mode: "Option A", OutputMode: "Option B", Tactics: "Option C"})

	db.Create(&Alias{ID: 1, UserID_A: 1, Location: "Work", Alias: "Office"})
	db.Create(&Alias{ID: 2, UserID_A: 1, Location: "Home", Alias: "My House"})

	// 查询数据
	fmt.Println("查询路线数据：")
	var routes []Route
	db.Find(&routes)
	for _, route := range routes {
		fmt.Printf("ID: %d, StartPoint: %s, EndPoint: %s\n", route.ID, route.StartPoint, route.EndPoint)
	}

	fmt.Println("\n查询用户偏好数据：")
	var preferences []Preference
	db.Find(&preferences)
	for _, preference := range preferences {
		fmt.Printf("ID: %d, UserID: %d, Mode: %s, OutputMode: %s, Tactics: %s\n", preference.ID, preference.UserID_P, preference.Mode, preference.OutputMode, preference.Tactics)
	}

	fmt.Println("\n查询地点别名数据：")
	var aliases []Alias
	db.Find(&aliases)
	for _, alias := range aliases {
		fmt.Printf("ID: %d, UserID: %d, Location: %s, Alias: %s\n", alias.ID, alias.UserID_A, alias.Location, alias.Alias)
	}
}
