import requests
import json

# 构建请求参数
params = {
    "origin": "中国湖北省武汉市洪山区华中科技大学",
    "destination": "中国湖北省武汉市洪山区武汉大学",
    "mode": "1",
    "outputmode": "1",
    "tactics": "1"
}

# 发送HTTP请求
response = requests.post("http://127.0.0.1:8080/routeplanning", data=params)
#把结果输入到文件中
with open('result.json', 'w') as f:
    f.write(response.text)
f.close()