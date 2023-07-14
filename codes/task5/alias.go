package task5

import (
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func AliasProcess(c *gin.Context) {
	// 解析请求中的 JSON 数据
	var alias Alias
	if err := c.ShouldBindJSON(&alias); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/task5"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 根据 Location 查询是否已存在
	var existingAlias Alias
	if err := db.Where("location = ?", alias.Location).First(&existingAlias).Error; err == nil {
		// 别名已存在，更新别名数据
		existingAlias.Alias = alias.Alias
		if err := db.Save(&existingAlias).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update alias"})
			return
		}
		c.JSON(http.StatusOK, existingAlias)
	} else {
		// 别名不存在，创建新的别名数据
		if err := db.Create(&alias).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create alias"})
			return
		}
		c.JSON(http.StatusCreated, alias)
	}
}

func TryAlias(location string) string {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/task5"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 根据地点查询别名
	var alias Alias
	if err := db.Where("location = ?", location).First(&alias).Error; err == nil {
		// 找到别名，返回别名值
		return alias.Alias
	}

	// 没有找到别名，返回原始地点
	return location
}
