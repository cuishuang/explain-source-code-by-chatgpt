# File: les/fetcher.go

les/fetcher.go文件是go-ethereum项目中实现Light Ethereum Subprotocol (LES)的网络请求和数据传输的代码文件。LES是以太坊网络用于支持轻客户端的一个子协议。

在该文件中，有几个重要的结构体，包括：
1. announce：用于向peers广播自己拥有的区块范围。
2. request：用于向peers请求区块数据。
3. response：用于存储peer返回的区块数据。
4. fetcherPeer：用于跟踪和管理与fetcher通信的peer的状态。
5. lightFetcher：LES协议的核心结构体，负责管理请求和数据传输。

下面分别介绍这些结构体的作用和相关的函数：

1. addAnno：将广播的announce消息添加到本地待处理的消息队列。
2. forwardAnno：向所有的peer广播announce消息。
3. newLightFetcher：初始化并返回一个新的lightFetcher实例。
4. start：启动lightFetcher，开始与peers进行交互。
5. stop：停止lightFetcher，关闭与peers的连接。
6. registerPeer：注册一个新的peer到lightFetcher。
7. unregisterPeer：从lightFetcher中注销一个peer。
8. peer：根据peer ID获取对应的fetcherPeer实例。
9. forEachPeer：对所有的peers执行一些操作，如发送请求等。
10. mainloop：lightFetcher的主循环，处理各种消息和事件。
11. announce：处理接收到的announce消息，更新本地区块链状态。
12. trackRequest：跟踪某个请求的状态，如已完成、超时等。
13. requestHeaderByHash：根据区块hash向peers请求区块头。
14. startSync：启动区块同步操作。
15. deliverHeaders：处理请求到的区块头，更新本地区块链状态。
16. rescheduleTimer：重新调度定时器，用于检查请求超时。

这些函数共同实现了LES协议中的请求、数据传输和区块同步等功能。fetcherPeer结构体用于跟踪每个peer的状态，并维护了与每个peer的连接。lightFetcher结构体作为LES协议的核心，负责管理所有的请求和数据传输，并处理各种消息和事件。

