# 1
仔细阅读了返回的json文件，发现可以给出三种个性化输出：
1. 仅给出时间（duration字段）
2. 给出时间、完整的instruction（instruction字段）
3. 仅给出路线（star_name字段）
那么现在就可以完成这个函数

# 2
通过正则表达式，将instruction中的html标签去除，得到纯文本的instruction，同时提取出开车的“站点”

但是如果是公共交通会出json解析错误：

```
40452 bytes written successfully
解析响应失败: json: cannot unmarshal array into Go struct field RouteInfo.result.routes.steps of type task2.Step
```

暂时先看下一个任务吧

# 3 

在看下一个任务之前好好沉淀规划一下

## 1

task3的第一个部分实际上就是要把地名通过api转换成经纬度，再作为参数发到task2中

这个part还是比较好搞

## 2

第二个部分则是要通过加权平均计算修正后的时间。需要注意的是要向task2的函数里面加一个路线偏好参数。这个部分可能不太好弄。

我的想法是：将路径偏好作为请求参数放到task2里面。task2的路径规划函数返回route

1. 通过tactics参数提供的偏好，提供多种偏好的路线
2. 通过给出的时间修正公式，给出修正后的最短路径

要完成这些，需要：
1. 给task2加一个参数
2. 在task3中加一个控制函数
3. 补充各种偏好选项的函数

# 4

在学长的提示下我发现公交的api暗藏玄机，他的step是个二维数组（不知道为什么），折腾好久改过来了。但是这也要求在task3中要对公交的step进行特殊处理。头疼。

学习到了有在线工具可以json->golang struct。大意了没有想到。