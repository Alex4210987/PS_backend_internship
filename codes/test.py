import requests

url = 'http://127.0.0.1:8080/trafficstatus'
data = {
    'location': '武汉市江岸区中山大道',
}

response = requests.post(url, data=data)

print(response.status_code)
print(response.json())
