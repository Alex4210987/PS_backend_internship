# 1

调了一上午小bug应该是差不多好了。。。但是就在此时api额度满了。。。先写后面的吧。

好的，又搞了一个ak，现在前三个任务都完成了。

# 2

现在考虑数据库的问题。

服务器现在提供了三个接口，分别是轨迹匹配，路况查询，路线规划。

而数据库中要存储的是：查询历史纪录，用户的偏好，地点的别名，某一个特定的路线。

我理解的使用场景是：后三者作为一种类似于cookie的东西存储在用户的浏览器中。

而查询历史纪录，我想可能有现成的解决方法？搜索一下。

# 3

考虑了这个问题，我觉得可以这样做：历史记录显然只能在前端解决，所以需要一个前端，我选择的是直接用js和html，因为不需要太花哨。

学习了基本的js，强行土法炼钢。

剩下两个数据库问题可以在后端解决。