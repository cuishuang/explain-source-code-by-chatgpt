# File: p2p/server.go

p2p/server.go文件是go-ethereum项目中实现了P2P（peer-to-peer）网络通信协议的服务器部分。它负责管理节点之间的连接、握手、消息传递等功能。

以下是对每个变量和结构体的详细说明：

- `errServerStopped`：表示服务器已停止的错误。
- `errEncHandshakeError`：表示加密握手过程中出错的错误。
- `errProtoHandshakeError`：表示协议握手过程中出错的错误。

下面是一些重要的结构体和变量的说明：

- `Config`：定义了服务器的配置参数，包括节点密钥、p2p协议版本、P2P网络ID等。
- `Server`：代表P2P服务器实例，封装了网络传输层、节点信息管理等功能。
- `peerOpFunc`：是一个函数类型，用于定义执行特定操作的处理函数。
- `peerDrop`：定义了当连接被断开时要执行的函数。
- `connFlag`：表示当前连接的状态，用于表示连接是否已停止。
- `conn`：表示一个P2P连接。
- `transport`：封装了底层的网络传输协议，如TCP或UDP。
- `sharedUDPConn`：表示共享的UDP连接。
- `NodeInfo`：表示节点的基本信息，如公钥、IP地址、端口等。

以下是一些重要的函数的说明：

- `String`：将NodeInfo结构体转换为字符串表示。
- `is`：检查节点是否具有特定的标志。
- `set`：设置节点的标志。
- `LocalNode`：返回当前节点的信息。
- `Peers`：返回当前节点的所有对等节点。
- `PeerCount`：返回当前节点连接的对等节点数。
- `AddPeer`：向当前节点添加对等节点。
- `RemovePeer`：从当前节点中移除特定的对等节点。
- `AddTrustedPeer`：向当前节点添加受信任的对等节点。
- `RemoveTrustedPeer`：从当前节点中移除受信任的对等节点。
- `SubscribeEvents`：订阅节点事件，如新对等节点连接、断开连接等。
- `Self`：返回当前节点的NodeInfo。
- `Stop`：停止P2P服务器并关闭相关的网络连接。
- `ReadFromUDP`：从UDP连接中读取数据。
- `Close`：关闭连接。
- `Start`：启动P2P服务器并开始监听连接。
- `setupLocalNode`：设置当前节点的基本信息。
- `setupDiscovery`：设置P2P服务器的发现服务。
- `setupDialScheduler`：设置拨号计划任务。
- `maxInboundConns`：设置入站连接的最大数量。
- `maxDialedConns`：设置已发起连接的最大数量。
- `setupListening`：设置P2P服务器的监听地址和端口。
- `setupUDPListening`：设置UDP监听地址和端口。
- `doPeerOp`：执行特定的对等节点操作。
- `run`：启动P2P服务器并接收连接。
- `postHandshakeChecks`：在握手完成后执行一些检查。
- `addPeerChecks`：在添加对等节点之前执行一些检查。
- `listenLoop`：监听并接受新的连接。
- `setupConn`：设置一个新的连接。
- `nodeFromConn`：根据连接创建一个新的节点实例。
- `checkpoint`：检查连接是否是可接受的节点。
- `launchPeer`：为新的连接创建一个对等节点。
- `runPeer`：启动对等节点。
- `NodeInfo`：返回节点的基本信息。
- `PeersInfo`：返回对等节点的基本信息。

这些函数和结构体一起实现了P2P服务器的各种功能，包括节点连接管理、握手验证、消息传递等。

