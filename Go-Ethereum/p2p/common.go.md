# File: p2p/discover/common.go

在go-ethereum项目中，p2p/discover/common.go是实现了以太坊节点之间互相发现的相关功能。该文件定义了一些基础的结构体和函数，以支持节点的发现和通信。

1. UDPConn结构体：表示一个UDP连接，用于发送和接收UDP数据包。
   - 字段包括：socket、读写锁和连接时的远程Peer地址等。

2. Config结构体：存储了节点的配置信息，包括本地节点的IP、端口等等。
   - 字段包括：IP、端口、协议版本、网络ID等。

3. ReadPacket函数：读取UDPConn中的数据包，如果存在则返回数据包和对端Peer，否则返回错误。

4. withDefaults函数：给一个Config结构体设置默认值，方便用户快速配置节点的参数。

5. ListenUDP函数：在指定的本地地址和端口上监听UDP数据包，返回UDPConn结构体和错误信息。
   - 这个函数在节点启动时用于监听和接收来自其他节点的消息。

6. min函数：返回两个整数中的最小值。

以上是对这些结构体和函数的简要介绍，它们在go-ethereum项目中的作用是搭建和管理节点互相发现和通信的基础。

