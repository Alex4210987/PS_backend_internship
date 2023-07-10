/*
通过正确构建 HTTP Request，实现 实时路况查询 和四种 轻量级路线规划 ：
驾车、骑行、步行、公交路线的 API 的调用，
并通过正确的解析 HTTP Response，获取 API 返回信息，
并支持用户输入和个性化的输出格式。

用户输入指：Request 的构建可以通过用户的输入来改变参数。

个性化的输出格式指：提供多个输出模式，
比如仅输出路线时间、额外输出转站点、
形式化的路线输出（ A(起点)->B(换乘)-> D ->C(终点站)
这种输出格式即可）。
*/
package task2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

//路况查询
type Response struct {
	Status      int    `json:"status"`
	Message     string `json:"message"`
	Description string `json:"description"`
	Evaluation  struct {
		Status     int    `json:"status"`
		StatusDesc string `json:"status_desc"`
	} `json:"evaluation"`
	RoadTraffic []struct {
		CongestionSections []struct {
			CongestionDistance int     `json:"congestion_distance"`
			Speed              float64 `json:"speed"`
			Status             int     `json:"status"`
			CongestionTrend    string  `json:"congestion_trend"`
			SectionDesc        string  `json:"section_desc"`
		} `json:"congestion_sections"`
		RoadName string `json:"road_name"`
	} `json:"road_traffic"`
}

func GetRoadTrafficStatus(roadName string, city string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	apiKey := os.Getenv("BAIDU_AK")
	baseURL := "https://api.map.baidu.com/traffic/v1/road"
	params := url.Values{
		"road_name": []string{roadName},
		"city":      []string{city},
		"ak":        []string{apiKey},
	}

	url := baseURL + "?" + params.Encode()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("请求发送失败: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("响应解析失败: %v", err)
		return
	}

	var responseData Response
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Printf("JSON解析失败: %v", err)
		return
	}

	fmt.Println("状态:", responseData.Status)
	fmt.Println("信息:", responseData.Message)
	fmt.Println("描述:", responseData.Description)
	fmt.Println("评价状态:", responseData.Evaluation.Status)
	fmt.Println("评价描述:", responseData.Evaluation.StatusDesc)

	for _, road := range responseData.RoadTraffic {
		fmt.Println("道路名称:", road.RoadName)
		for _, section := range road.CongestionSections {
			fmt.Println("拥堵路段描述:", section.SectionDesc)
			fmt.Println("拥堵距离:", section.CongestionDistance)
			fmt.Println("速度:", section.Speed)
			fmt.Println("拥堵状态:", section.Status)
			fmt.Println("拥堵趋势:", section.CongestionTrend)
		}
	}
}
