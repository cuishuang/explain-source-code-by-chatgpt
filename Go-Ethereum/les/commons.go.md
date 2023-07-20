# File: les/commons.go

在go-ethereum项目中，`les/commons.go` 文件是“Light Ethereum Subprotocol (LES)”的共享代码库。LES是以太坊的轻量级客户端协议，用于在轻量级客户端和全节点之间进行数据交互。该文件中定义了LES相关的一些结构体和函数。

1. `chainReader` 结构体：它是一个接口，定义了获取区块链状态和数据的方法。不同的区块链后端可以实现该接口，以满足LES客户端对区块链数据的请求。

2. `lesCommons` 结构体：它定义了LES协议的一些共享常量和参数。例如，它包含了LES协议的支持的最大区块接收数量、最大区块请求数量等。

3. `NodeInfo` 结构体：它保存了一个LES节点的信息，包括节点的IP地址、端口号、协议版本等。

`errResp` 函数：它创建了一个带有错误信息的LES响应消息。这个函数在发送错误响应时非常有用。

`makeProtocols` 函数：它根据指定的配置创建LES协议实例。

`nodeInfo` 函数：它返回一个节点信息结构体 `NodeInfo` 的实例，其中包含了节点的各种信息。

这些函数和结构体在go-ethereum的LES模块中扮演了关键的角色，实现了LES协议的各种功能和特性。

