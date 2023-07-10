package task2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	//"strings"
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
	Distance         int    `json:"distance"`
	Duration         int    `json:"duration"`
	Price            int    `json:"price"`
	LinePrice        []Line `json:"line_price"`
	Steps            []Step `json:"steps"`
	TrafficCondition int    `json:"traffic_condition"`
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
	StartLoc    LocInfo `json:"start_location"`
	EndLoc      LocInfo `json:"end_location"`
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
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

func RoutePlanning(mode string, origin string, destination string, outputmode string, tactics string) (body []byte) {
	// 获取用户输入
	_ = godotenv.Load()
	ak := os.Getenv("BAIDU_AK")
	var (
		steps_info int
	)
	//仅输出路线时间、额外输出转站点、形式化的路线输出
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
	//params.Set("vehicle", vehicle)
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
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		os.Exit(1)
	}
	//写道example.json里
	f, _ := os.Create("example1.json")
	l, _ := f.WriteString(string(body))
	fmt.Println(l, "bytes written successfully")
	// 解析JSON响应
	if mode == "4" {
		RoutePlanningBus(body, outputmode)
		return body
	}
	var route Route
	err = json.Unmarshal(body, &route)
	fmt.Println("route is", route)
	if err != nil {
		fmt.Println("解析响应失败:", err)
		os.Exit(1)
	}
	switch outputmode {
	case "1":
		OnlyTime(route)
	case "2":
		OnlyTransferCar(route)
	case "3":
		AllInfo(route)
	default:
		fmt.Println("无效的输出模式编号")
	}
	return body
}
func OnlyTime(route Route) {
	fmt.Println("路线规划模式：仅输出路线时间")
	i := 0
	for _, route := range route.Result.Routes {
		i++
		fmt.Println("no", i, "预计耗时：", route.Duration, "秒")
	}
	//taxi duration
	//如果存在这个字段：
	if route.Result.Taxi.Duration != 0 {
		fmt.Println("预计打车耗时：", route.Result.Taxi.Duration, "秒")
	}
}
func AllInfo(route Route) {
	fmt.Println("路线规划模式：输出所有信息")
	i := 0
	for _, route := range route.Result.Routes {
		i++
		fmt.Println("no", i, "预计耗时：", route.Duration, "秒")
		fmt.Println("预计距离：", route.Distance, "米")
		fmt.Println("预计价格：", route.Price, "元")
		fmt.Println("路线：")
		for _, step := range route.Steps {
			fmt.Println(removeHTMLTags(step.Instruction))
		}
	}
	if route.Result.Taxi.Duration != 0 {
		fmt.Println("预计打车耗时：", route.Result.Taxi.Duration, "秒")
		fmt.Println("预计距离：", route.Result.Taxi.Distance, "米")
		fmt.Println("预计价格：", route.Result.Taxi.TotalPrice, "元")
		fmt.Println("路线：")
		fmt.Println(route.Result.Taxi.Detail)
	}
} //输出各个路线和打车的耗时、距离、价格、路线、转站点
func OnlyTransferCar(route Route) {
	var stations []string
	i := 0
	for _, route := range route.Result.Routes {
		i++
		fmt.Println("no", i, "预计耗时：", route.Duration, "秒")
		fmt.Println("预计距离：", route.Distance, "米")
		fmt.Println("预计价格：", route.Price, "元")
		fmt.Println("路线：")
		for _, step := range route.Steps {
			//在instruction字段中寻找被<b>\<\b>包围的字符串并存在一个数组里
			re := regexp.MustCompile(`<b>(.*?)<\/b>`)
			matches := re.FindAllStringSubmatch(step.Instruction, -1)
			for _, match := range matches {
				if len(match) > 1 {
					stations = append(stations, match[1])
				}
			}
		}
		for _, station := range stations {
			fmt.Println(station)
			//如果不是最后
			if station != stations[len(stations)-1] {
				fmt.Print("->")
			}
		}
	}
}
func removeHTMLTags(html string) string {
	re := regexp.MustCompile("<[^>]*>")
	plainText := re.ReplaceAllString(html, "")
	return plainText
}
