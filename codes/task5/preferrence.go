package task5

import (
	"net/http"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"errors"

	"github.com/gin-gonic/gin"
)

func PreferenceProcess(c *gin.Context) {
	fmt.Println(c.Params)
	// 从请求中获取数据
	var preference Preference
	if err := c.ShouldBindJSON(&preference); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	fmt.Println(preference)
	// 增加或更新 Preference
	//如果三个字段都相同，则return，否则增加
	dsn := "root:123456@tcp(127.0.0.1:3306)/task5"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if err := db.Where("UserID_P = ? AND Mode = ? AND OutputMode = ? AND Tactics = ?", preference.UserID_P, preference.Mode, preference.OutputMode, preference.Tactics).First(&preference).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = db.Create(&preference).Error
			if err != nil {
				fmt.Println("Error creating preference:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create preference"})
				return
			}
			c.JSON(http.StatusCreated, preference)
			fmt.Println(preference)
		} else {
			fmt.Println("Error querying preference:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query preference"})
			return
		}
	}
}

func RouteProcess(c *gin.Context) {
	var route Route
	if err := c.ShouldBindJSON(&route); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	// 增加或更新 Route
	dsn := "root:123456@tcp(127.0.0.1:3306)/task5"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	//如果两个字段都相同，则return，否则增加
	if err := db.Where("StartPoint = ? AND EndPoint = ?", route.StartPoint, route.EndPoint).First(&route).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = db.Create(&route).Error
			if err != nil {
				fmt.Println("Error creating route:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create route"})
				return
			}
			c.JSON(http.StatusCreated, route)
			fmt.Println(route)
		} else {
			fmt.Println("Error querying route:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query route"})
			return
		}
	}
}
