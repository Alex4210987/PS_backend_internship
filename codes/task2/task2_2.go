package task2

import (
	//"strings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

type Route struct {
	Status int    `json:"status"`
	Result Result `json:"result"`
}
type Taxi struct {
	Detail     []Detail `json:"detail"`
	KmPrice    float64  `json:"km_price"`
	StartPrice float64  `json:"start_price"`
	TotalPrice float64  `json:"total_price"`
	Distance   float64  `json:"distance"`
	Duration   int      `json:"duration"`
	Remark     string   `json:"remark"`
}

type Detail struct {
	Desc       string  `json:"desc"`
	KmPrice    float64 `json:"km_price"`
	StartPrice int     `json:"start_price"`
	TotalPrice int     `json:"total_price"`
}
type Result struct {
	Routes []RouteInfo `json:"routes"`
	Taxi   Taxi        `json:"taxi"`
}

type RouteInfo struct {
	Distance         int     `json:"distance"`
	Duration         int     `json:"duration"`
	Price            int     `json:"price"`
	LinePrice        []Line  `json:"line_price"`
	Steps            []Step  `json:"steps"`
	TrafficCondition int     `json:"traffic_condition"`
	JamIndex         float64 `json:"jam_index"`
}

type Line struct {
	LineType  int `json:"line_type"`
	LinePrice int `json:"line_price"`
}

type Step struct {
	Distance    int     `json:"distance"`
	Duration    int     `json:"duration"`
	StepType    int     `json:"type"`
	Instruction string  `json:"instruction"`
	Vehicle     Vehicle `json:"vehicle"`
	Path        string  `json:"path"`
	//StartLoc    LocInfo `json:"start_location"`
	//EndLoc      LocInfo `json:"end_location"`
	Status int `json:"status"`
}

type Vehicle struct {
	Name          string  `json:"name"`
	DirectionText string  `json:"direction_text"`
	LineID        string  `json:"line_id"`
	StartName     string  `json:"start_name"`
	EndName       string  `json:"end_name"`
	StartTime     string  `json:"start_time"`
	EndTime       string  `json:"end_time"`
	StopNum       int     `json:"stop_num"`
	TotalPrice    float64 `json:"total_price"`
	Type          int     `json:"type"`
	ZonePrice     float64 `json:"zone_price"`
}

type LocInfo struct {
	Lng string `json:"lng"`
	Lat string `json:"lat"`
}

type FloatLocInfo struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}
type Output struct { //对应result
	Msg        []string          `json:"msg"`
	RouteCount int               `json:"route_count"`
	Distance   int               `json:"distance"`
	Price      int               `json:"price"`
	Routes     []ReturnRouteInfo `json:"routes"`
}
type ReturnRouteInfo struct {
	Distance int     `json:"distance"`
	Duration int     `json:"duration"`
	JamIndex float64 `json:"jam_index"`
}

func OnlyTime(route Route, output *Output) {
	output.Msg = append(output.Msg, "路线规划模式：仅输出路线时间")
	if len(route.Result.Routes) == 0 {
		output.Msg = append(output.Msg, "没有找到路线信息")
		return
	}
	output.Msg = append(output.Msg, fmt.Sprintf("预计时间：%d", route.Result.Routes[0].Duration))
	output.RouteCount = len(route.Result.Routes)
}

func AllInfo(route Route, output *Output) {
	output.Msg = append(output.Msg, "路线规划模式：输出所有信息")
	if len(route.Result.Routes) == 0 {
		output.Msg = append(output.Msg, "没有找到路线信息")
		return
	}
	output.RouteCount = len(route.Result.Routes)
	for i, route := range route.Result.Routes {
		output.Msg = append(output.Msg, fmt.Sprintf("路线%d：", i+1))
		output.Msg = append(output.Msg, fmt.Sprintf("预计距离：%d 米", route.Distance))
		output.Msg = append(output.Msg, fmt.Sprintf("预计价格：%d 元", route.Price))
		output.Msg = append(output.Msg, "路线：")
		for _, step := range route.Steps {
			output.Msg = append(output.Msg, removeHTMLTags(step.Instruction))
		}
	}
}

func OnlyTransferCar(route Route, output *Output) {
	var stations []string
	output.RouteCount = len(route.Result.Routes)
	output.Msg = append(output.Msg, "路线规划模式：仅输出换乘信息")
	if len(route.Result.Routes) == 0 {
		output.Msg = append(output.Msg, "没有找到路线信息")
		return
	}
	for i, route := range route.Result.Routes {
		output.Msg = append(output.Msg, fmt.Sprintf("预计距离：%d 米", route.Distance))
		output.Msg = append(output.Msg, fmt.Sprintf("预计价格：%d 元", route.Price))
		output.Msg = append(output.Msg, "路线：", fmt.Sprint(i+1))
		for _, step := range route.Steps {
			// 在instruction字段中寻找被<b>\<\b>包围的字符串并存在一个数组里
			re := regexp.MustCompile(`<b>(.*?)<\/b>`)
			matches := re.FindAllStringSubmatch(step.Instruction, -1)
			for _, match := range matches {
				if len(match) > 1 {
					stations = append(stations, match[1])
				}
			}
		}
		for _, station := range stations {
			output.Msg = append(output.Msg, station)
			// 如果不是最后
			if station != stations[len(stations)-1] {
				output.Msg = append(output.Msg, "->")
			}
		}
	}
}

func RoutePlanning(mode string, origin string, destination string, outputmode string, tactics string) (returnroute Route, output Output) {
	// 获取用户输入
	_ = godotenv.Load()
	ak := os.Getenv("BAIDU_AK")
	var (
		steps_info int
	)
	// 根据用户输入构建API请求URL
	var (
		apiURL   string
		retCoord string
		//vehicle   string
		coordType string
		//tactics   int
	)
	switch mode {
	case "1":
		apiURL = "https://api.map.baidu.com/directionlite/v1/driving"
	case "2":
		apiURL = "https://api.map.baidu.com/directionlite/v1/riding"
	case "3":
		apiURL = "https://api.map.baidu.com/directionlite/v1/walking"
	case "4":
		apiURL = "https://api.map.baidu.com/directionlite/v1/transit"
	default:
		fmt.Println("无效的模式编号")
		os.Exit(1)
	}
	retCoord = "bd09ll"
	coordType = "bd09ll"
	steps_info = 1
	// 构建API请求参数
	params := url.Values{}
	params.Set("ak", ak)
	params.Set("origin", origin)
	params.Set("destination", destination)
	params.Set("ret_coordtype", retCoord)
	params.Set("coord_type", coordType)
	params.Set("steps_info", fmt.Sprint(steps_info))
	if mode == "1" {
		params.Set("tactics", fmt.Sprint(tactics))
	}
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
	//fmt.Println(string(body))
	// 解析JSON响应
	// 创建一个用于存储输出的结构体
	if mode == "4" {
		routeBus, output := RoutePlanningBus(body, outputmode, output)
		route := Bus2Other(routeBus)
		return route, output
	}
	var route Route
	err = json.Unmarshal(body, &route)
	//fmt.Println("route is", route)
	if err != nil {
		fmt.Println("解析响应失败:", err)
		os.Exit(1)
	}
	//fmt.Println("route is", route)
	switch outputmode {
	case "1":
		OnlyTime(route, &output)
	case "2":
		OnlyTransferCar(route, &output)
	case "3":
		AllInfo(route, &output)
	default:
		output.Msg = append(output.Msg, "无效的输出模式编号")
	}

	// 设置其他字段的值
	//output.Distance = route.Result.Routes[0].Distance
	//output.Price = route.Result.Routes[0].Price
	//output.Distance = route.Result.Routes[0].Distance
	return route, output
}
func removeHTMLTags(html string) string {
	re := regexp.MustCompile("<[^>]*>")
	plainText := re.ReplaceAllString(html, "")
	return plainText
}
func Bus2Other(routebus RouteBus) (route Route) {
	route.Status = routebus.Status
	route.Result.Taxi = routebus.Result.Taxi
	route.Result.Routes = make([]RouteInfo, len(routebus.Result.Routes))
	for i, routeinfobus := range routebus.Result.Routes {
		route.Result.Routes[i].Distance = routeinfobus.Distance
		route.Result.Routes[i].Duration = routeinfobus.Duration
		route.Result.Routes[i].Price = routeinfobus.Price
		route.Result.Routes[i].LinePrice = routeinfobus.LinePrice
		route.Result.Routes[i].TrafficCondition = routeinfobus.TrafficCondition
		route.Result.Routes[i].Steps = make([]Step, len(routeinfobus.Steps))
		for j, stepbus := range routeinfobus.Steps {
			route.Result.Routes[i].Steps[j].Distance = stepbus[0].Distance
			route.Result.Routes[i].Steps[j].Duration = stepbus[0].Duration
			route.Result.Routes[i].Steps[j].StepType = stepbus[0].StepType
			route.Result.Routes[i].Steps[j].Instruction = stepbus[0].Instruction
			route.Result.Routes[i].Steps[j].Vehicle = stepbus[0].Vehicle
			route.Result.Routes[i].Steps[j].Path = stepbus[0].Path
			//route.Result.Routes[i].Steps[j].StartLoc = stepbus[0].StartLoc
			//route.Result.Routes[i].Steps[j].EndLoc = stepbus[0].EndLoc
		}
	}
	return route
} //把Steps            [][]Step `json:"steps"`转换成Steps            []Step `json:"steps"`
