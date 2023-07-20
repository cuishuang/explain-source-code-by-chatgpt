# File: les/utils/limiter.go

在go-ethereum项目中，les/utils/limiter.go这个文件的作用是实现一个通用的限制器（Limiter）。

结构体的作用如下：
1. Limiter：限制器的主结构体，用于保存限制器的内部状态信息。
2. nodeQueue：用于存储节点对象的队列，按照节点优先级排序。
3. addressGroup：用于存储地址与节点对象的映射关系。
4. request：用于表示一个请求的结构体，在添加到节点队列中时会带有权重属性。
5. dropListItem：表示在限制器中被丢弃的请求的结构体。

函数的作用如下：
1. flatWeight：将权重值平均分配到所有节点上。
2. add：将节点对象添加到地址组和节点队列中。
3. update：更新限制器中节点对象对应的权重值。
4. remove：从地址组和节点队列中移除指定的节点对象。
5. choose：根据节点的权重随机选择一个节点。
6. NewLimiter：构造一个新的限制器对象。
7. selectionWeights：根据地址组中的节点对象的权重，生成节点对象的选择权重数组。
8. Add：将请求添加到限制器中。
9. addToGroup：将请求按照地址分组添加到地址组中。
10. removeFromGroup：将指定地址对应的请求从地址组中移除。
11. processLoop：限制器的核心工作循环，负责从地址组和节点队列中取出请求进行处理。
12. Stop：停止限制器的工作循环。
13. dropRequests：将指定数量的请求从限制器中丢弃。

