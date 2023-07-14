package task2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const defaultRadius = 200

type TrafficInfo struct {
	Evaluation Evaluation `json:"evaluation"`
}
type Evaluation struct {
	Status     int    `json:"status"`
	StatusDesc string `json:"status_desc"`
}

func GetTrafficStatus(latitude float64, longitude float64) (int, string, error) {
	// 加载环境变量
	err := godotenv.Load(".env")
	if err != nil {
		return 0, "", fmt.Errorf("加载环境变量失败: %v", err)
	}

	// 从环境变量中获取 AK
	ak := os.Getenv("BAIDU_AK")
	if ak == "" {
		return 0, "", fmt.Errorf("未设置 AK")
	}

	// 构造请求URL
	//https: //api.map.baidu.com/traffic/v1/around?ak=你的AK&center=39.912078,116.464303&radius=200&coord_type_input=gcj02&coord_type_output=gcj02
	url := fmt.Sprintf("https://api.map.baidu.com/traffic/v1/around?ak=%s&center=%f,%f&radius=%d&coord_type_input=gcj02&coord_type_output=gcj02", ak, latitude, longitude, defaultRadius)
	// 发起GET请求
	response, err := http.Get(url)
	if err != nil {
		return 0, "", fmt.Errorf("请求失败: %v", err)
	}
	defer response.Body.Close()
	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, "", fmt.Errorf("读取响应失败: %v", err)
	}

	// 解析JSON响应
	var trafficInfo TrafficInfo
	err = json.Unmarshal(body, &trafficInfo)
	if err != nil {
		return 0, "", fmt.Errorf("解析响应失败: %v", err)
	}

	return trafficInfo.Evaluation.Status, trafficInfo.Evaluation.StatusDesc, nil
}
