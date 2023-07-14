package main

import (
	"codes/task1"
	"codes/task2"
	"codes/task3"
	"codes/task5"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 配置静态文件路由
	router.Static("/static", "./static")

	// 设置HTML模板引擎
	router.LoadHTMLFiles("static/index.html")

	// 定义HTML路由
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 定义轨迹匹配的POST路由 武汉市洪山区华中科技大学
	router.POST("/trackmatch", task1.TrackMatchProcess)

	// 定义路况查询的POST路由
	router.POST("/trafficstatus", task2.TrafficStatusProcess)

	// 定义路线规划的POST路由
	router.POST("/routeplanning", task3.RoutePlanningProcess)

	//定义别名设置的POST路由
	router.POST("/alias", task5.AliasProcess)

	//定义偏好设置的POST路由
	router.POST("/prteference", task5.PreferenceProcess)

	// 设置gin服务器的静态文件路径
	router.StaticFile("/trackmatch.js", "./static/trackmatch.js")
	router.StaticFile("/trafficstatus.js", "./static/trafficstatus.js")
	router.StaticFile("/routeplanning.js", "./static/routeplanning.js")
	// 启动服务
	router.Run("localhost:8080")
}
