import requests

# 为了简化测试，可以将以下URL替换为Go服务所在的地址
BASE_URL = "http://127.0.0.1:8080"

def test_preference():
    url = BASE_URL + "/preference"
    data = {
        "UserID_P": 1,
        "Mode": "1",
        "OutputMode": "1",
        "Tactics": "1"
    }
    response = requests.post(url, json=data)
    print(response.json())

def test_route():
    url = BASE_URL + "/route"
    data = {
        "StartPoint": "1",
        "EndPoint": "2"
    }
    response = requests.post(url, json=data)
    print(response.json())

test_preference()
test_route()