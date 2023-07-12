//路径规划：输入起点终点地名，转成经纬度，调用路径规划api，返回路径规划信息。
package task3

import (
	"github.com/gin-gonic/gin"
)

func RoutePlanningProcess(c *gin.Context) {
	// 获取请求参数
	origin := c.PostForm("origin")
	destination := c.PostForm("destination")
	mode := c.PostForm("mode")
	outputmode := c.PostForm("outputmode")
	tactics := c.PostForm("tactics")
	// 调用路径规划函数
	output := PersonaliazeRoutePlanning(mode, origin, destination, outputmode, tactics)
	// 返回结果
	c.JSON(200, gin.H{
		"status":  "OK",
		"message": "成功",
		"result":  output,
	})
}
