//路况查询：支持输入地名，然后调用经纬度api，再返回路况信息。
package task2

import (
	"codes/task3_1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TrafficStatusProcess(c *gin.Context) {
	//读取location参数
	location := c.PostForm("location")
	//调用经纬度api
	lng, lat, err := task3_1.Geocode(location)
	//错误处理
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "请求参数错误",
		})
		return
	}
	//调用路况api
	status, statusdesc, err := GetTrafficStatus(lat, lng)
	//返回路况信息
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"msg":    "路况查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "路况查询成功",
		"data":   status,
		"desc":   statusdesc,
	})
}
