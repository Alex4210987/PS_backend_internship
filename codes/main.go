package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type Request struct {
	Track         []string `json:"track"`
	StandardTrack []string `json:"standard_track"`
}

type Response struct {
	Status int `json:"status"`
	Data   struct {
		Similarity float64 `json:"similarity"`
	} `json:"data"`
}

func main() {
	// 此处填写您在控制台-应用管理-创建应用后获取的AK
	ak := os.Getenv("BAIDU_AK")
	// 服务地址
	host := "https://api.map.baidu.com"

	// 请求地址
	uri := "/trackmatch/v1/track"

	// 读取包含请求数据的JSON文件
	file, err := ioutil.ReadFile("input.json")
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}

	var reqData Request
	err = json.Unmarshal(file, &reqData) // 解析JSON数据
	if err != nil {
		fmt.Println("解析JSON失败:", err)
		return
	}

	// 转换为字符串形式
	standardTrackStr := "[" + "\"" + joinStrings(reqData.StandardTrack, "\",\"") + "\"" + "]"
	trackStr := "[" + "\"" + joinStrings(reqData.Track, "\",\"") + "\"" + "]"

	// 设置请求参数
	params := url.Values{
		"ak":                []string{ak},
		"option":            []string{"need_mapmatch:1|transport_mode:driving|denoise_grade:1|vacuate_grade:1"},
		"standard_option":   []string{"need_mapmatch:1|transport_mode:driving|denoise_grade:1|vacuate_grade:1"},
		"coord_type_input":  []string{"bd09ll"},
		"coord_type_output": []string{"bd09ll"},
		"standard_track":    []string{standardTrackStr},
		"track":             []string{trackStr},
	}

	// 发起请求
	url := host + uri
	resp, err := http.PostForm(url, params)
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

	// 解析响应体
	var responseData Response
	err = json.Unmarshal(body, &responseData)
	//fmt.Println(string(body))
	//print status and similarity in json
	fmt.Println("Status:", responseData.Status, "Similarity:", responseData.Data.Similarity)
}

// 辅助函数：将字符串切片连接成一个字符串
func joinStrings(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	result := strs[0]
	for _, s := range strs[1:] {
		result += sep + s
	}
	return result
}
