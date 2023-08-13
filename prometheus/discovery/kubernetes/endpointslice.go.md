# File: discovery/kubernetes/endpointslice.go

discovery/kubernetes/endpointslice.go文件是Prometheus项目中实现Kubernetes的EndpointSlice监听和处理的主要文件。EndpointSlice是Kubernetes中的一种资源对象，用于替代旧的Endpoints对象，更具体地描述了Service的后端Pod。

在该文件中，有几个重要的变量用于统计操作次数：
- epslAddCount：用于统计添加EndpointSlice的次数。
- epslUpdateCount：用于统计更新EndpointSlice的次数。
- epslDeleteCount：用于统计删除EndpointSlice的次数。

EndpointSlice是定义在Kubernetes源代码中的结构体，它包括了以下几个重要的字段：
- Metdata：包含了EndpointSlice的元数据信息，如名称、命名空间、标签等。
- AddressType：定义了EndpointSlice的地址类型，可以是IPv4或IPv6。
- Ports：定义了服务的端口和协议。
- Endpoints：包含了实际的后端Pod的IP和Port信息。

文件中的几个函数的功能如下：
- NewEndpointSlice：用于创建一个新的EndpointSlice对象。
- enqueueNode：将节点加入到队列中，用于后续处理。
- enqueue：将EndpointSlice加入到队列中，用于后续处理。
- Run：运行EndpointSlice控制器，监听和处理事件。
- process：处理特定事件类型的函数。
- getEndpointSliceAdaptor：获取EndpointSlice适配器，用于处理EndpointSlice的增删改查操作。
- endpointSliceSource：封装了EndpointSlice的查询和监听操作。
- endpointSliceSourceFromNamespaceAndName：根据命名空间和名称创建EndpointSliceSource对象。
- buildEndpointSlice：根据给定的参数构建EndpointSlice对象。
- resolvePodRef：根据Pod的引用获取Pod名称和命名空间。
- addServiceLabels：添加Service的标签到EndpointSlice对象中。

以上函数和变量组合起来实现了Prometheus项目对Kubernetes的EndpointSlice监听和处理功能。

