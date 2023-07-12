//用gin实现各种功能的调用入口
package main

import (
	"github.com/gin-gonic/gin"

	"codes/task1"
	"codes/task2"
	"codes/task3"
)

func main() {
	router := gin.Default()

	// 定义轨迹匹配的POST路由
	router.POST("/trackmatch", task1.TrackMatchProcess)
	// 定义路况查询的POST路由
	router.POST("/trafficstatus", task2.TrafficStatusProcess)
	//定义路线规划的POST路由
	router.POST("/routeplanning", task3.RoutePlanningProcess)
	// 启动服务
	router.Run(":8080")
}
