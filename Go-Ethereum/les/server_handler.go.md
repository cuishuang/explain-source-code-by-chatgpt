# File: les/server_handler.go

les/server_handler.go这个文件是go-ethereum项目中实现Light Ethereum Subprotocol (LES)服务器的关键文件。它定义了一个ServerHandler结构体，该结构体实现了LES服务器的所有功能。

- errTooManyInvalidRequest是一个错误变量，表示请求的无效次数过多。
- serverHandler结构体具有以下作用：
  - 包含与LES服务器相关的配置参数，如块链（BlockChain）和交易池（TxPool）。
  - 在服务器启动时，通过调用newServerHandler函数进行初始化。
  - 处理LES协议消息的接收和发送。提供了与LES协议相关的一些功能函数。

- newServerHandler函数用于创建一个新的ServerHandler实例，并初始化其相关的参数。
- start函数用于启动LES服务器，开始监听和接收来自对等节点的LES协议消息。
- stop函数用于停止LES服务器，关闭连接并进行清理。
- runPeer函数是一个goroutine，用于处理一个与对等节点的连接。
- handle函数根据接收到的消息类型，将消息分发给不同的处理函数进行处理。
- beforeHandle函数在处理消息前执行一些预处理操作。
- afterHandle函数在处理消息后执行一些后处理操作。
- handleMsg函数用于实际处理接收到的LES协议消息。
- BlockChain提供了与块链相关的功能函数，如获取最新区块、获取区块头等。
- TxPool提供了与交易池相关的功能函数，如添加交易、获取交易等。
- ArchiveMode用于判断当前节点是否运行在归档模式下。
- AddTxsSync函数用于将交易添加到交易池中，并等待交易同步完成。
- getAccount函数用于从状态数据库中获取账户信息。
- GetHelperTrie函数用于获取助记词（Merkle Trie）。
- broadcastLoop函数用于不断广播消息给已连接的对等节点。

总之，les/server_handler.go文件中的这些函数和结构体定义了LES服务器的关键功能，包括处理LES协议消息、与块链和交易池交互等。

