import requests

URL = "http://127.0.0.1:8080/alias"

def test_alias():
    # 创建一个新的别名
    data = {"location": "武汉大学", "alias": "学校"}
    response = requests.post(URL, json=data)
    print(response.json())
    data = {"location": "华中科技大学大学", "alias": "学校"}
    response = requests.post(URL, json=data)
    print(response.json())
    data = {"location": "武汉大学", "alias": "武大"}
    response = requests.post(URL, json=data)
    print(response.json())

test_alias()