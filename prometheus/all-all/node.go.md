# File: discovery/kubernetes/node.go

在Prometheus项目中，`discovery/kubernetes/node.go`文件的作用是实现从Kubernetes API获取节点信息并进行监控。该文件主要包含了节点发现相关的函数、变量和结构体。

- `nodeAddCount`、`nodeUpdateCount`、`nodeDeleteCount`是用于计数节点的添加、更新和删除操作的变量，用来记录节点发现的变化情况。
- `Node`结构体用于表示一个Kubernetes节点，包含了节点的名称、地址等信息。
- `NewNode`函数用于创建一个新的节点对象。
- `enqueue`函数将一个事件添加到队列中，以便后续处理。
- `Run`函数是事件处理的入口函数，用来启动节点发现的处理逻辑。
- `process`函数从队列中取出事件进行处理，根据事件的类型执行相应的操作。
- `convertToNode`函数用于将Kubernetes API返回的节点对象转换为自定义的`Node`结构体对象。
- `nodeSource`用于标识节点发现来源的常量。
- `nodeSourceFromName`函数通过名称获取节点发现来源。
- `nodeLabels`函数用于获取节点的标签信息。
- `buildNode`函数用于构建一个包含节点信息的`Node`对象。
- `nodeAddress`函数用于获取节点的地址信息。

总的来说，`discovery/kubernetes/node.go`文件实现了从Kubernetes API获取节点信息并进行监控的功能，包括节点的添加、更新和删除操作的处理以及节点对象的构建和转换。

