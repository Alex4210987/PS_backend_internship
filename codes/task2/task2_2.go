package task2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	//"strings"

	"github.com/joho/godotenv"	
)

type Route struct {
	Status int    `json:"status"`
	Result Result `json:"result"`
}

type Result struct {
	Routes []RouteInfo `json:"routes"`
}

type RouteInfo struct {
	Distance int          `json:"distance"`
	Duration int          `json:"duration"`
	Steps    []Step   `json:"steps"`
}


type Step struct {
	Instruction string  `json:"instruction"`
	Distance    int     `json:"distance"`
	Duration    int     `json:"duration"`
	Path        string  `json:"path"`
	StepType    int     `json:"type"`
	StartLoc    LocInfo `json:"start_location"`
	EndLoc      LocInfo `json:"end_location"`
}

type LocInfo struct {
	Lng string `json:"lng"`
	Lat string `json:"lat"`
}

func RoutePlanning() {
	// 获取用户输入
	_ = godotenv.Load()
	ak := os.Getenv("BAIDU_AK")
	var (
		mode     string
		origin   string
		destination string
		outputmode string
	)
	//仅输出路线时间、额外输出转站点、形式化的路线输出
	fmt.Println("请选择路线规划模式(驾车 骑行 步行 公交)")
	fmt.Scanln(&mode)
	fmt.Print("请输入起点坐标（格式：纬度,经度）：")
	fmt.Scanln(&origin)
	fmt.Print("请输入终点坐标（格式：纬度,经度）：")
	fmt.Scanln(&destination)
	fmt.Print("请输入输出模式：(仅输出路线时间、额外输出转站点、形式化的路线输出)")
	fmt.Scanln(&outputmode)
	// 根据用户输入构建API请求URL
	var (
		apiURL    string
		retCoord  string
		vehicle   string
		coordType string
	)
	switch mode {
	case "1":
		apiURL = "https://api.map.baidu.com/directionlite/v1/driving"
		retCoord = "bd09ll"
		vehicle = "car"
		coordType = "bd09ll"
	case "2":
		apiURL = "https://api.map.baidu.com/directionlite/v1/riding"
		retCoord = "bd09ll"
		vehicle = "bike"
		coordType = "bd09ll"
	case "3":
		apiURL = "https://api.map.baidu.com/directionlite/v1/walking"
		retCoord = "bd09ll"
		vehicle = "walk"
		coordType = "bd09ll"
	case "4":
		apiURL = "https://api.map.baidu.com/directionlite/v1/transit"
		retCoord = "bd09ll"
		vehicle = "bus"
		coordType = "bd09ll"
	default:
		fmt.Println("无效的模式编号")
		os.Exit(1)
	}

	// 构建API请求参数
	params := url.Values{}
	params.Set("ak", ak)
	params.Set("origin", origin)
	params.Set("destination", destination)
	params.Set("ret_coordtype", retCoord)
	params.Set("vehicle", vehicle)
	params.Set("coord_type", coordType)

	// 发送HTTP GET请求并获取响应
	resp, err := http.Get(apiURL + "?" + params.Encode())
	if err != nil {
		fmt.Println("请求发送失败:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// 读取HTTP响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		os.Exit(1)
	}

	// 解析JSON响应
	var route Route
	err = json.Unmarshal(body, &route)
	if err != nil {
		fmt.Println("解析响应失败:", err)
		os.Exit(1)
	}

	// 处理解析后的响应
	switch outputmode {
	case "1":
		OnlyTime(route)
	case "2":
		OnlyTransfer(route)
	case "3":
		FormatOutput(route)
	default:
		AllOutput(route, mode, LocInfo{Lng: origin}, LocInfo{Lng: destination})
	}
}
func OnlyTime(route Route) {
	fmt.Println("路线规划模式：仅输出路线时间")
	fmt.Println("预计耗时：", route.Result.Routes[0].Duration, "秒")}
func FormatOutput(route Route) {
}
func OnlyTransfer(route Route) {
}
func AllOutput(route Route, modeName string, origin LocInfo, destination LocInfo) {
}