# File: pkg/controller/util/endpointslice/endpointset.go

pkg/controller/util/endpointslice/endpointset.go文件是为了管理EndpointSlice的EndpointSet结构体。EndpointSet是一个存储Endpoint的集合，Endpoint必须连接到同一个EndpointSlice。而EndpointSlice是kubernetes中用于描述服务访问的一个资源对象。

在EndpointSet中，有三个重要的结构体：endpointHash、endpointHashObj和EndpointSet。

- endpointHash: 用于存储Endpoint的哈希表，它将Endpoint转换为可哈希的值。对于每个Endpoint，我们可以使用它的地址作为它的哈希值。
- endpointHashObj: 哈希表中的内部对象，它存储要哈希的Endpoint，是一个非导出结构体。
- EndpointSet: 用于存储Endpoint的结构体，提供了对Endpoint的添加、删除、查找和列表等操作。

同时，EndpointSet提供了一些重要函数，如：

- hashEndpoint: 计算Endpoint的哈希值。
- Insert: 添加Endpoint到EndpointSet中。
- Delete: 从EndpointSet中删除特定的Endpoint。
- Has: 判断EndpointSet中是否存在特定的Endpoint。
- Get: 获取特定哈希值的Endpoint。
- UnsortedList: 返回EndpointSet中的不排序列表。
- PopAny: 随机弹出一个Endpoint。
- Len: 获取EndpointSet中存储Endpoint的数量。

总之，EndpointSet是kubernetes中管理EndpointSlice的一个重要工具，使用哈希表存储Endpoint，提供了对Endpoint的添加、删除、查找和列表等操作。

