# 1
看了一下gin框架，大概了解怎么把几个task的功能整合到服务器里面了。一开始做这个项目的时候还看的一头雾水不知道怎么用。但是啊啊啊啊好懒啊不想写代码qwq。

# 2
逐个需求分析怎么用gin实现:
## 2.1
轨迹匹配：这个直接用router把发过来的json解析出来丢到tackmatch里面就可以。
这个测试出来没问题
## 2.2/3
路况查询：支持输入地名，然后调用经纬度api，再返回路况信息。
结构体有点问题，给他修了一下调好了。

路径规划：输入起点终点地名，转成经纬度，调用路径规划api，返回路径规划信息。
这个写好了，但是还没测试。