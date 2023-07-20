# File: les/odr.go

在go-ethereum项目中，les/odr.go文件是用来实现LES (Light Ethereum Subprotocol)的ODR (On-Demand Retrieval)功能的。ODR是一种协议，用于按需从Ethereum全节点网络中检索区块和交易，以便轻客户端（如轻钱包）可以更高效地同步和查询区块链数据。

该文件中定义了几个结构体：

1. LesOdr：LES ODR对象，用于管理接收到的区块和交易的索引和缓存。它还提供了通过该协议进行数据检索的功能。

2. Msg：LES ODR消息对象，用于在网络中传输ODR请求和响应。

3. peerByTxHistory：交易历史记录索引，用于对LES ODR请求进行筛选，并将请求与可以提供所需数据的网络节点进行匹配。

以下是该文件中的一些函数的作用：

1. NewLesOdr：创建一个新的LES ODR对象。

2. Stop：停止LES ODR对象的操作。

3. Database：从现有的数据库返回LES ODR对象，用于存储和检索区块和交易数据。

4. SetIndexers：配置LES ODR对象使用的索引器。

5. ChtIndexer：返回一个支持复合合同历史记录（CHT）索引的LES ODR对象。

6. BloomTrieIndexer：返回一个支持布隆过滤器和Trie索引的LES ODR对象。

7. BloomIndexer：返回一个支持布隆过滤器索引的LES ODR对象。

8. IndexerConfig：索引器的配置。

9. Len：返回LES ODR对象中缓存的区块和交易数量。

10. Less：比较两个LES ODR项的顺序。

11. Swap：交换两个LES ODR项的位置。

12. RetrieveTxStatus：检查交易在LES ODR缓存或索引中的状态。

13. Retrieve：从网络中检索一个区块或交易。

这些函数提供了LES ODR功能的核心实现，包括创建对象、配置索引器、数据检索和处理等。通过使用这些函数，LES ODR对象可以与Ethereum网络交互，并提供按需检索数据的功能。

