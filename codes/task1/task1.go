package task1

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

var ak string // 私有变量

func TrackMatch(filepath string) {
	err := godotenv.Load()
	ak = os.Getenv("BAIDU_AK")
	host := "https://api.map.baidu.com"
	uri := "/trackmatch/v1/track"

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}

	var reqData Request
	err = json.Unmarshal(file, &reqData)
	if err != nil {
		fmt.Println("解析JSON失败:", err)
		return
	}

	standardTrackStr := "[" + "\"" + joinStrings(reqData.StandardTrack, "\",\"") + "\"" + "]"
	trackStr := "[" + "\"" + joinStrings(reqData.Track, "\",\"") + "\"" + "]"

	params := url.Values{
		"ak":                []string{ak},
		"option":            []string{"need_mapmatch:1|transport_mode:driving|denoise_grade:1|vacuate_grade:1"},
		"standard_option":   []string{"need_mapmatch:1|transport_mode:driving|denoise_grade:1|vacuate_grade:1"},
		"coord_type_input":  []string{"bd09ll"},
		"coord_type_output": []string{"bd09ll"},
		"standard_track":    []string{standardTrackStr},
		"track":             []string{trackStr},
	}

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

	var responseData Response
	err = json.Unmarshal(body, &responseData)
	fmt.Println("Status:", responseData.Status, "Similarity:", responseData.Data.Similarity)
}

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
