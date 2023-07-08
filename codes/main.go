//all in one file first; spilt later
/*
using gin framework
使用百度地图开放平台api
AK 不可以明文存放在代码中，必须通过环境变量等方式进行获取。
通过正确构建 HTTP Request, 实现 轨迹重合率分析 的API的调用，
并通过正确的解析 HTTP Response，获取 API 返回结果中的 "status" 和
"similarity"，并输出。
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// 定义请求体结构体
type CompareRequest struct {
	Track1 []string `json:"track1"`
	Track2 []string `json:"track2"`
}

// 定义响应体结构体
type CompareResponse struct {
	Status     int     `json:"status"`
	Similarity float64 `json:"similarity"`
}

func main() {
	// 获取AK
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ak := os.Getenv("BAIDU_MAP_AK")
	if ak == "" {
		fmt.Println("AK error")
		return
	}

	// 创建HTTP服务
	router := gin.Default()
	api_path := os.Getenv("API_PATH")
	// 创建API路由
	router.POST(api_path, func(c *gin.Context) {
		// 解析请求体
		req, err := parseCompareRequest(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 调用百度地图API
		apiResp, err := callBaiduMapAPI(api_path, ak, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 返回结果
		c.JSON(http.StatusOK, apiResp)
	})

	// 启动HTTP服务
	runServer(router)
}

// 解析比较请求
func parseCompareRequest(c *gin.Context) (*CompareRequest, error) {
	var req CompareRequest
	if err := c.BindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

// 调用百度地图API
func callBaiduMapAPI(api_path string, ak string, req *CompareRequest) (*CompareResponse, error) {
	url := fmt.Sprintf("http://api.map.baidu.com%s?ak=%s", api_path, ak)
	body := fmt.Sprintf("track1=%s&track2=%s", strings.Join(req.Track1, ";"), strings.Join(req.Track2, ";"))
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp CompareResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// 启动HTTP服务
func runServer(router *gin.Engine) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(fmt.Sprintf(":%s", port))
}
