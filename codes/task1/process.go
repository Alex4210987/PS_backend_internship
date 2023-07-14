package task1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TrackMatchProcess(c *gin.Context) {
	// 读取请求体
	var reqData Request
	err := c.BindJSON(&reqData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "请求参数错误",
		})
		return
	}
	fmt.Println(reqData)
	// 调用轨迹匹配函数
	respData, err := TrackMatch(reqData)
	fmt.Println(respData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"msg":    "轨迹匹配失败",
			"data":   reqData,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "轨迹匹配成功",
		"data":   respData,
	})
}
