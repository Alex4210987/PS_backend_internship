package task3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	Status int    `json:"status"`
	Result Result `json:"result"`
}

type Result struct {
	Location      Location `json:"location"`
	Precise       int      `json:"precise"`
	Confidence    int      `json:"confidence"`
	Comprehension int      `json:"comprehension"`
	Level         string   `json:"level"`
}

type Location struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

func Geocode(address string) (string, string, error) {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		return "0", "0", fmt.Errorf("failed to load environment variables: %s", err.Error())
	}

	// 获取API密钥
	apiKey := os.Getenv("BAIDU_AK")
	if apiKey == "" {
		return "0", "0", fmt.Errorf("API key is missing")
	}

	// 构建请求URL
	baseURL := "https://api.map.baidu.com/geocoding/v3/"
	u, err := url.Parse(baseURL)
	if err != nil {
		return "0", "0", fmt.Errorf("failed to parse URL: %s", err.Error())
	}

	// 设置请求参数
	q := u.Query()
	q.Set("address", address)
	q.Set("output", "json")
	q.Set("ak", apiKey)
	u.RawQuery = q.Encode()

	// 发送HTTP GET请求
	resp, err := http.Get(u.String())
	if err != nil {
		return "0", "0", fmt.Errorf("failed to send request: %s", err.Error())
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "0", "0", fmt.Errorf("failed to read response body: %s", err.Error())
	}

	// 解析JSON响应
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "0", "0", fmt.Errorf("failed to parse JSON response: %s", err.Error())
	}

	// 检查API响应状态
	if response.Status != 0 {
		return "0", "0", fmt.Errorf("API request failed with status code: %d", response.Status)
	}

	// 提取经纬度（string形式）
	lng := fmt.Sprintf("%f", response.Result.Location.Lng)
	lat := fmt.Sprintf("%f", response.Result.Location.Lat)

	return lng, lat, nil
}
