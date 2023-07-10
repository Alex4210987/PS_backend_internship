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