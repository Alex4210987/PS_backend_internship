package task2

import (
	"encoding/json"
	"fmt"
	"os"
)

type RouteBus struct {
	Status int       `json:"status"`
	Result ResultBus `json:"result"`
}

type ResultBus struct {
	Routes []RouteInfoBus `json:"routes"`
	Taxi   Taxi           `json:"taxi"`
}

type RouteInfoBus struct {
	Distance         int      `json:"distance"`
	Duration         int      `json:"duration"`
	Price            int      `json:"price"`
	LinePrice        []Line   `json:"line_price"`
	Steps            [][]Step `json:"steps"`
	TrafficCondition int      `json:"traffic_condition"`
}

func RoutePlanningBus(body []byte, outputmode string, output Output) (route RouteBus, returnoutput Output) {
	// 解析JSON响应
	err := json.Unmarshal(body, &route)
	fmt.Println("busroute is", route)
	if err != nil {
		fmt.Println("解析响应失败:", err)
		os.Exit(1)
	}
	switch outputmode {
	case "1":
		OnlyTimeBus(route, &output)
	case "2":
		OnlyTransferBus(route, &output)
	case "3":
		AllInfoBus(route, &output)
	default:
		output.Msg = append(output.Msg, "无效的输出模式编号")
	}

	// 设置其他字段的值
	output.RouteCount = len(route.Result.Routes)
	output.Distance = route.Result.Routes[0].Distance
	output.Price = route.Result.Routes[0].Price

	// 将输出的结构体转换为JSON格式
	jsonData, err := json.Marshal(output)
	if err != nil {
		fmt.Println("转换为JSON失败:", err)
		os.Exit(1)
	}

	// 将JSON数据写入文件
	err = os.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		fmt.Println("写入文件失败:", err)
		os.Exit(1)
	}

	return route, output
}

func OnlyTimeBus(route RouteBus, output *Output) {
	output.Msg = append(output.Msg, "路线规划模式：仅输出路线时间")
	output.RouteCount = len(route.Result.Routes)
	for _, route := range route.Result.Routes {
		routeInfo := RouteInfo{
			Duration: route.Duration,
			Steps:    FlattenSteps(route.Steps),
		}
		output.Routes = append(output.Routes, routeInfo)
	}
	// 如果存在这个字段：
	if route.Result.Taxi.Duration != 0 {
		output.Msg = append(output.Msg, fmt.Sprintf("预计打车耗时：%d 秒", route.Result.Taxi.Duration))
	}
}

func AllInfoBus(route RouteBus, output *Output) {
	output.Msg = append(output.Msg, "路线规划模式：输出所有信息")
	output.RouteCount = len(route.Result.Routes)
	for _, route := range route.Result.Routes {
		routeInfo := RouteInfo{
			Duration: route.Duration,
			Steps:    FlattenSteps(route.Steps),
		}
		output.Routes = append(output.Routes, routeInfo)
	}
	if route.Result.Taxi.Duration != 0 {
		output.Msg = append(output.Msg, fmt.Sprintf("预计打车耗时：%d 秒", route.Result.Taxi.Duration))
		output.Msg = append(output.Msg, fmt.Sprintf("预计距离：%f 米", route.Result.Taxi.Distance))
		output.Msg = append(output.Msg, fmt.Sprintf("预计价格：%f 元", route.Result.Taxi.TotalPrice))
	}
}

func OnlyTransferBus(route RouteBus, output *Output) {
	output.Msg = append(output.Msg, "路线规划模式：仅输出转站点")
	output.RouteCount = len(route.Result.Routes)
	for _, route := range route.Result.Routes {
		routeInfo := RouteInfo{
			Duration: route.Duration,
			Steps:    FlattenSteps(route.Steps),
		}
		output.Routes = append(output.Routes, routeInfo)

		end := ""
		output.Msg = append(output.Msg, fmt.Sprintf("预计耗时：%d 秒", route.Duration))
		output.Msg = append(output.Msg, fmt.Sprintf("预计距离：%d 米", route.Distance))
		output.Msg = append(output.Msg, fmt.Sprintf("预计价格：%d 元", route.Price))
		output.Msg = append(output.Msg, "路线：")
		for _, stepOptions := range route.Steps {
			for _, step := range stepOptions {
				if step.Vehicle.StartName != "" {
					output.Msg = append(output.Msg, step.Vehicle.StartName)
				}
				// 输出最后一个不为空的endlocation
				if step.Vehicle.EndName != "" {
					end = step.Vehicle.EndName
				}
			}
		}
		output.Msg = append(output.Msg, end)
	}
	if route.Result.Taxi.Duration != 0 {
		output.Msg = append(output.Msg, fmt.Sprintf("预计打车耗时：%d 秒", route.Result.Taxi.Duration))
		output.Msg = append(output.Msg, fmt.Sprintf("预计距离：%f 米", route.Result.Taxi.Distance))
		output.Msg = append(output.Msg, fmt.Sprintf("预计价格：%f 元", route.Result.Taxi.TotalPrice))
	}
}

func FlattenSteps(steps [][]Step) []Step {
	var flattened []Step
	for _, stepOptions := range steps {
		flattened = append(flattened, stepOptions...)
	}
	return flattened
}
