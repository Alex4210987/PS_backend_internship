package task5

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func PreferenceProcess(c *gin.Context) {
	// 从请求中获取数据
	var route Route
	if err := c.ShouldBindJSON(&route); err != nil {
		return
	}

	var preference Preference
	if err := c.ShouldBindJSON(&preference); err != nil {
		return
	}

	dsn := "root:123456@tcp(127.0.0.1:3306)/task5"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	// 增加或更新 Route
	//如果两个字段都相同，则return，否则增加
	if err := db.Where("start_point = ? AND end_point = ?", route.StartPoint, route.EndPoint).First(&route).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = db.Create(&route).Error
			if err != nil {
				return
			}
		} else {
			return
		}
	}

	// 增加或更新 Preference
	//如果三个字段都相同，则return，否则增加
	if err := db.Where("user_id_p = ? AND mode = ? AND output_mode = ? AND tactics = ?", preference.UserID_P, preference.Mode, preference.OutputMode, preference.Tactics).First(&preference).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = db.Create(&preference).Error
			if err != nil {
				return
			}
		} else {
			return
		}
	}

	return
}
