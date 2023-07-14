// /*
// 我的想法是：将路径偏好作为请求参数放到task2里面。task2的路径规划函数返回route

// 通过tactics参数提供的偏好，提供多种偏好的路线
// 通过给出的时间修正公式，给出修正后的最短路径
// 要完成这些，需要：

// 给task2加一个参数
// 在task3中加一个控制函数
// 补充各种偏好选项的函数
// */
package task3

import (
	"fmt"
	"strconv"
	"strings"

	"codes/task2"
	"codes/task3_1"
)

func PersonaliazeRoutePlanning(mode string, oriname string, destname string, outputmode string, tactics string) (rtnoutput *task2.Output) {
	var orign *task2.LocInfo = &task2.LocInfo{}       // 为 orign 分配内存
	var destination *task2.LocInfo = &task2.LocInfo{} // 为 destination 分配内存
	var err error
	lngfloat, latfloat, err := task3_1.Geocode(oriname)
	orign.Lng = strconv.FormatFloat(lngfloat, 'f', -1, 64)
	orign.Lat = strconv.FormatFloat(latfloat, 'f', -1, 64)
	if err != nil {
		fmt.Println("起点地名错误", err)
	}
	lngfloat, latfloat, err = task3_1.Geocode(destname)
	destination.Lng = strconv.FormatFloat(lngfloat, 'f', -1, 64)
	destination.Lat = strconv.FormatFloat(latfloat, 'f', -1, 64)
	if err != nil {
		fmt.Println("终点地名错误", err)
	}
	//locinfo to string
	originStr := fmt.Sprintf("%s,%s", orign.Lat, orign.Lng)
	destinationStr := fmt.Sprintf("%s,%s", destination.Lat, destination.Lng)

	//调用task2的路径规划函数
	route, output := task2.RoutePlanning(mode, originStr, destinationStr, outputmode, tactics)
	//调整duration
	dur, jam := ModifidedTime(route)
	outputJamIndex(dur, jam, &output)
	//三种功能： 2.拥堵最少，4.加权时间最短，5.换乘最少。均是针对公交
	if mode != "4" {
		output.Msg = append(output.Msg, "不是公交模式，无法进行个性化路线规划")
	}
	switch tactics {
	case "2":
		leastJam(&route, jam, &output)
	case "4":
		leastTime(&route, dur, &output)
	case "5":
		leastTransfer(&route, dur, &output)
	default:
		fmt.Println("tactics参数错误，默认为最短时间")
		leastTime(&route, dur, &output)
	}
	return &output
}
func ModifidedTime(route task2.Route) ([]float64, []float64) {

	duration := make([]float64, len(route.Result.Routes))
	allJamIndex := make([]float64, len(route.Result.Routes))

	for i, r := range route.Result.Routes {
		var stepJamIndex float64
		var noneZeroPath int

		for _, step := range r.Steps {
			pathArray := Path2Array(step.Path)

			for _, path := range pathArray {
				trafficStatus, _, err := task2.GetTrafficStatus(path[1], path[0])
				if err != nil {
					fmt.Println("获取交通状态失败:", err)
					continue
				}

				jamIndexFloat := float64(trafficStatus)
				if jamIndexFloat >= 1.0 && jamIndexFloat <= 4.0 {
					stepJamIndex += jamIndexFloat
					noneZeroPath++
					//fmt.Println(jamIndexFloat, stepJamIndex, noneZeroPath)
				}
			}
		}

		if stepJamIndex > 0.5 {
			stepJamIndex /= float64(noneZeroPath)
			duration[i] = float64(r.Duration) * (0.8 + stepJamIndex*0.2)
		} else {
			duration[i] = float64(r.Duration)
		}
		allJamIndex[i] = stepJamIndex
		fmt.Println("r.Duration:", r.Duration)
		fmt.Printf("Route %d - Duration: %f, Jam Index: %f\n", i, duration[i], allJamIndex[i])
	}

	return duration, allJamIndex

}

/*
它提供了三种拥堵status，1、2、3、4，那么事实上可以认为H=H*(1+(status-1)*20%)
*/
func Path2Array(path string) (pathArray [][]float64) {
	// 将 path 转化为二维数组
	// path 的格式为：lat1,lng1;lat2,lng2;lat3,lng3;...

	// 初始化 pathArray 为空切片
	pathArray = make([][]float64, 0)

	// 添加第一个坐标对到 pathArray
	pathArray = append(pathArray, make([]float64, 2))

	// 将一维数组中的每个元素按照逗号分割，得到一个二维数组
	// 逐个处理路径中的坐标对
	for _, pair := range strings.Split(path, ";") {
		// 分割经度和纬度
		coordinates := strings.Split(pair, ",")
		if len(coordinates) == 2 {
			// 将经度和纬度转化为 float64 类型并添加到 pathArray
			lat, _ := strconv.ParseFloat(coordinates[0], 64)
			lng, _ := strconv.ParseFloat(coordinates[1], 64)
			pathArray = append(pathArray, []float64{lat, lng})
		}
	}

	return pathArray
}
func outputJamIndex(duration []float64, jamIndex []float64, output *task2.Output) {
	fmt.Println("我他妈进来了！")
	for i := range duration {
		//增加一个新的空output.routes，设置duration和jamIndex
		output.Routes = append(output.Routes, task2.ReturnRouteInfo{})
		output.Routes[i].Duration = int(duration[i])
		output.Routes[i].JamIndex = jamIndex[i]
		output.Msg = append(output.Msg, fmt.Sprintf("第%d条路线的拥堵指数：%f", i+1, jamIndex[i]))
		output.Msg = append(output.Msg, fmt.Sprintf("第%d条路线的修正时间：%d", i+1, int(duration[i])))
	}
}

func leastJam(route *task2.Route, jamIndex []float64, output *task2.Output) {
	//找到最小的拥堵指数
	if len(jamIndex) == 0 {
		output.Msg = append(output.Msg, "没有找到拥堵信息")
		return
	}
	minJamIndex := jamIndex[0]
	minIndex := 0
	for i := 1; i < len(jamIndex); i++ {
		if jamIndex[i] < minJamIndex {
			minJamIndex = jamIndex[i]
			minIndex = i
		}
	}
	output.Msg = append(output.Msg, fmt.Sprintf("最小拥堵指数的路线：%d", minIndex+1))
	output.Msg = append(output.Msg, fmt.Sprintf("拥堵指数：%f", minJamIndex))
	output.Msg = append(output.Msg, fmt.Sprintf("预计时间：%d", route.Result.Routes[minIndex].Duration))
}

func leastTime(_ *task2.Route, modifidedTime []float64, output *task2.Output) {
	//找到最小的时间
	if len(modifidedTime) == 0 {
		output.Msg = append(output.Msg, "没有找到时间信息")
		return
	}
	minTime := modifidedTime[0]
	minIndex := 0
	for i := 1; i < len(modifidedTime); i++ {
		if modifidedTime[i] < minTime {
			minTime = modifidedTime[i]
			minIndex = i
		}
	}
	//存储最小时间的路线到output结构体的Msg字段
	output.Msg = append(output.Msg, fmt.Sprintf("最小时间的路线：%d", minIndex+1))
	output.Msg = append(output.Msg, fmt.Sprintf("预计时间：%f", modifidedTime[minIndex]))
}

func leastTransfer(route *task2.Route, modifidedTime []float64, output *task2.Output) {
	// 找到最小的换乘次数
	minTransfer := 1000
	minIndex := 0
	if len(route.Result.Routes) == 0 {
		output.Msg = append(output.Msg, "没有找到换乘信息")
		return
	}
	for i := 0; i < len(route.Result.Routes); i++ {
		// 统计换乘次数
		transfer := 0
		for _, step := range route.Result.Routes[i].Steps {
			if step.Vehicle.StartName != "" {
				transfer++
			}
		}
		if transfer < minTransfer {
			minTransfer = transfer
			minIndex = i
		}
	}
	// 存储最小换乘次数的路线到 output 结构体的 Msg 字段
	output.Msg = append(output.Msg, fmt.Sprintf("最小换乘次数的路线：%d", minIndex+1))
	output.Msg = append(output.Msg, fmt.Sprintf("换乘次数：%d", minTransfer))
	output.Msg = append(output.Msg, fmt.Sprintf("预计时间：%f", modifidedTime[minIndex]))
}
