# 1
获得了ak

# 2
速通了golang语法、http、gin框架，不会的再查或者上ai

# 3
<<<<<<< HEAD
照猫画虎写了task1的代码，但是还没有测试。接下来看看api的文档并debug。
=======
照猫画虎写了task1的代码，但是还没有测试。接下来看看api的文档并debug。

# 4
阅读api文档，这参数好多看得头大

我的想法是把参数写到一个json中，然后用main.go读这个json再发送请求。

哦官方文档给了json例子啊，我是笨比

啊啊啊啊我是sb，一开始写的是对的，改了之后反而错了，现在要回滚了

# 5

python脚本测试时返回status 101

说明：

- 数据包成功传输到了百度服务器
- 但是没有携带ak，所以返回101

考虑哪儿出了问题:

- 加入print ak后，发现ak是对的
- 突然想到没必要肉眼debug，加了log

根据log判断，python发送的原始请求没问题：
<pre>
原始请求数据：
POST http://localhost:8080/task1
{'Content-Length': '139', 'Content-Type': 'application/json'}
b'{"standard_track": ["36.2716924153,120.401133898,1542702871,10,0,182,10"], "track": ["36.2716924153,120.401133898,1542702871,10,0,182,10"]}'
status:  1005
similarity:  0
</pre>

但是go服务端收到的请求压根就不存在：
<pre>
2023/07/08 19:32:40 Received POST request with body: 
2023/07/08 19:32:40 请求地址: https://api.map.baidu.com/trackmatch/v1/track?ak=*******************
2023/07/08 19:32:40 请求参数: track1=&track2=
2023/07/08 19:32:41 响应状态: 1005
2023/07/08 19:32:41 响应结果: 0.000000
</pre>
玉玉了

更抽象的是这个状态码根本不存在？
>>>>>>> 572d2cb (deleted)
