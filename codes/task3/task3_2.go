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

// import (
// 	"codes/task2"
// 	"fmt"
// )

// func PersonaliazeRoutePlanning(mode string, oriname string, destname string, outputmode string, tactics string) {
// 	var orign *task2.LocInfo
// 	var destination *task2.LocInfo
// 	var err error
// 	orign.Lng, orign.Lat, err = Geocode(oriname)
// 	if err != nil {
// 		fmt.Println("起点地名错误", err)
// 	}
// 	destination.Lng, destination.Lat, err = Geocode(destname)
// 	if err != nil {
// 		fmt.Println("终点地名错误", err)
// 	}
// 	//locinfo to string
// 	originStr := fmt.Sprintf("%s,%s", orign.Lat, orign.Lng)
// 	destinationStr := fmt.Sprintf("%s,%s", destination.Lat, destination.Lng)
// 	//调用task2的路径规划函数
// 	route := task2.RoutePlanning(mode, originStr, destinationStr, outputmode, tactics)
// 	if tactics ==
// }
