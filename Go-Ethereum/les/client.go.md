# File: les/client.go

在go-ethereum项目中，les/client.go文件的作用是实现了与Light Ethereum Subprotocol（LES）节点进行通信的客户端。

LightEthereum结构体是LES协议的主要结构体之一，它用于表示一个LES节点。LightDummyAPI结构体是用于支持LES API的存根结构体。

以下是几个在client.go文件中定义的函数的作用：

- New函数: 创建一个新的LES客户端，接收参数为服务器地址、服务ID、客户端协议版本和客户端用户代理。
- VfluxRequest函数: 在现有块应该下载之前请求连续的块。
- vfxVersion函数: 用于请求远程节点的LES子协议版本。
- prenegQuery函数: 用于协商与远程节点之间的支持和配置。
- Etherbase函数: 返回远程节点的帐户地址。
- Coinbase函数: 返回远程节点挖矿奖励接收地址。
- Hashrate函数: 返回远程节点的哈希率。
- Mining函数: 返回远程节点的挖矿状态。
- APIs函数: 返回远程节点支持的API列表。
- ResetWithGenesisBlock函数: 重置本地状态，并从Genesis块开始重新同步。
- BlockChain函数: 返回与客户端关联的区块链。
- TxPool函数: 返回远程节点的事务池。
- Engine函数: 返回远程节点的区块链执行引擎。
- LesVersion函数: 返回LES协议的版本。
- Downloader函数: 返回远程节点的区块下载器。
- EventMux函数: 返回远程节点的事件处理器。
- Merger函数: 返回远程节点的块合并器。
- Protocols函数: 返回远程节点支持的协议列表。
- Start函数: 启动LES客户端。
- Stop函数: 停止LES客户端的运行。

这些函数提供了与LES节点的通信功能，以及获取该节点的相关信息和状态。

