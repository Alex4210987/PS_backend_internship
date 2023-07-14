//路径规划：输入起点终点地名，转成经纬度，调用路径规划api，返回路径规划信息。
package task3

import (
	"codes/task5"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RoutePlanningProcess(c *gin.Context) {
	//打印请求体
	fmt.Println("请求体是：", c.Request.Body)
	// 获取请求参数
	origin := c.PostForm("origin")
	destination := c.PostForm("destination")
	origin = task5.TryAlias(origin)
	destination = task5.TryAlias(destination)
	mode := c.PostForm("mode")
	outputmode := c.PostForm("outputmode")
	tactics := c.PostForm("tactics")
	//open file
	fmt.Println("参数是：", origin, destination, mode, outputmode, tactics)
	// 调用路径规划函数
	output := PersonaliazeRoutePlanning(mode, origin, destination, outputmode, tactics)
	// 返回结果
	c.JSON(200, gin.H{
		"status":  "OK",
		"message": "成功",
		"result":  output,
	})
}
