# File: light/odr_util.go

在go-ethereum项目中，light/odr_util.go文件的作用是提供了一系列用于通过轻客户端协议（OdrProtocol）获取轻量级区块数据的工具函数。

首先，errNonCanonicalHash是一个定义的错误变量，用于表示非规范哈希错误，当传入的哈希值不符合规范时，会返回这个错误。

接下来是一系列的函数，每个函数都是根据轻客户端协议中定义的请求类型，通过调用OdrProtocol的相关方法来获取具体的数据。它们的作用分别如下：

- GetHeaderByNumber：根据区块号获取相应区块的头部信息。
- GetCanonicalHash：根据区块号获取相应区块的规范哈希（canonical hash），用于唯一标识一个区块。
- GetTd：根据区块号获取相应区块的难度（total difficulty）值，表示挖矿所需的工作量。
- GetBodyRLP：根据区块号获取相应区块的二进制编码形式的区块数据，包含交易和本地状态数据。
- GetBody：根据区块号获取相应区块的区块数据，以方便列表形式进行访问。
- GetBlock：根据区块号获取完整的区块数据，包括区块头部和区块体。
- GetBlockReceipts：根据区块号获取相应区块的交易收据（receipts），用于验证交易的执行结果。
- GetBlockLogs：根据区块号获取相应区块中的日志记录。
- GetBloomBits：根据区块号获取相应区块的布隆过滤器位图，用于进行日志和事件相关查询。
- GetTransaction：根据交易哈希获取相应交易的详细信息。

这些函数是为了提供一个方便的接口，允许轻客户端快速获取区块链数据的部分内容，而无需下载整个区块链。这对于资源受限的设备（如移动设备）或带宽有限的网络环境是非常有用的。

