package task2

import (
	"encoding/json"
	"fmt"
	"os"
	//"strings"
	//"regexp"
)

type RouteBus struct {
	Status int       `json:"status"`
	Result ResultBus `json:"result"`
}

type ResultBus struct {
	Routes []RouteInfoBus `json:"routes"`
	Taxi    Taxi   	  `json:"taxi"`
}

type RouteInfoBus struct {
	Distance         int      `json:"distance"`
	Duration         int      `json:"duration"`
	Price            int      `json:"price"`
	LinePrice        []Line   `json:"line_price"`
	Steps            [][]Step `json:"steps"`
	TrafficCondition int      `json:"traffic_condition"`
}

// type StepOption struct {
// 	Steps []Step
// }

func RoutePlanningBus(body []byte, outputmode string) (route RouteBus) {
	// 解析JSON响应
	err := json.Unmarshal(body, &route)
	fmt.Println("busroute is", route)
	if err != nil {
		fmt.Println("解析响应失败:", err)
		os.Exit(1)
	}
	switch outputmode {
	case "1":
		OnlyTimeBus(route)
	case "2":
		OnlyTransferBus(route)
	case "3":
		AllInfoBus(route)
	default:
		fmt.Println("无效的输出模式编号")
	}
	return route
}
func OnlyTimeBus(route RouteBus) {
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
func AllInfoBus(route RouteBus) {
	fmt.Println("路线规划模式：输出所有信息")
	i := 0
	for _, route := range route.Result.Routes {
		i++
		fmt.Println("No.", i, "预计耗时：", route.Duration, "秒")
		fmt.Println("预计距离：", route.Distance, "米")
		fmt.Println("预计价格：", route.Price, "元")
		fmt.Println("路线：")
		for _, stepoptions := range route.Steps {
			for _, step := range stepoptions {
				fmt.Print(removeHTMLTags(step.Instruction))
			}
		}
	}
	if route.Result.Taxi.Duration != 0 {
		fmt.Println("预计打车耗时：", route.Result.Taxi.Duration, "秒")
		fmt.Println("预计距离：", route.Result.Taxi.Distance, "米")
		fmt.Println("预计价格：", route.Result.Taxi.TotalPrice, "元")
	}
} //输出各个路线和打车的耗时、距离、价格、路线、转站点
func OnlyTransferBus(route RouteBus) {
	end := ""
	fmt.Println("路线规划模式：仅输出转站点")
	i := 0
	for _, route := range route.Result.Routes {
		i++
		fmt.Println("no", i, "预计耗时：", route.Duration, "秒")
		fmt.Println("预计距离：", route.Distance, "米")
		fmt.Println("预计价格：", route.Price, "元")
		fmt.Println("路线：")
		for _, stepoption := range route.Steps {
			for _, step := range stepoption {
				if step.Vehicle.StartName != "" {
					fmt.Print(step.Vehicle.StartName, "->")
				}
				//输出最后一个不为空的endlocation
				if step.Vehicle.EndName != "" {
					end = step.Vehicle.EndName
				}
			}
		}
		fmt.Println(end)
	}
	if route.Result.Taxi.Duration != 0 {
		fmt.Println("预计打车耗时：", route.Result.Taxi.Duration, "秒")
		fmt.Println("预计距离：", route.Result.Taxi.Distance, "米")
		fmt.Println("预计价格：", route.Result.Taxi.TotalPrice, "元")
	}
} //输出换乘站（很tricky，因为除了公交之外是没有换乘站的）
