# File: eth/protocols/eth/handlers.go

在go-ethereum项目中，eth/protocols/eth/handlers.go文件扮演着处理以太坊协议消息的角色。该文件中包含了一系列不同的函数，每个函数都负责处理不同类型的以太坊协议消息。

1. handleGetBlockHeaders66：处理针对以太坊块头的查询请求，返回指定范围内的块头信息。

2. ServiceGetBlockHeadersQuery：提供块头查询服务，处理以太坊块头的查询请求，并返回相应的块头信息。

3. serviceNonContiguousBlockHeaderQuery：处理非连续的块头查询请求，返回指定范围内的块头信息。

4. serviceContiguousBlockHeaderQuery：处理连续的块头查询请求，返回指定范围内的块头信息。

5. handleGetBlockBodies66：处理对以太坊区块体的查询请求，返回指定区块的交易体信息。

6. ServiceGetBlockBodiesQuery：提供区块体查询服务，处理对以太坊区块体的查询请求，并返回相应的区块体信息。

7. handleGetNodeData66：处理对以太坊节点数据的查询请求，返回指定节点数据。

8. ServiceGetNodeDataQuery：提供节点数据查询服务，处理对以太坊节点数据的查询请求，并返回相应的节点数据。

9. handleGetReceipts66：处理对以太坊交易收据的查询请求，返回指定交易收据。

10. ServiceGetReceiptsQuery：提供交易收据查询服务，处理对以太坊交易收据的查询请求，并返回相应的交易收据。

11. handleNewBlockhashes：处理新块哈希消息，更新节点的最新块哈希信息。

12. handleNewBlock：处理新块消息，更新节点的最新区块信息。

13. handleBlockHeaders66：处理块头消息，根据请求返回相应的块头信息。

14. handleBlockBodies66：处理区块体消息，根据请求返回相应的区块体信息。

15. handleNodeData66：处理节点数据消息，根据请求返回相应的节点数据。

16. handleReceipts66：处理交易收据消息，根据请求返回相应的交易收据。

17. handleNewPooledTransactionHashes66：处理新的交易哈希消息，更新节点所维护的待处理交易哈希列表。

18. handleNewPooledTransactionHashes68：处理新的交易哈希消息，更新节点所维护的待处理交易哈希列表。

19. handleGetPooledTransactions66：处理待处理交易查询请求，返回待处理的交易信息。

20. answerGetPooledTransactions：回复待处理交易查询请求，返回待处理的交易信息。

21. handleTransactions：处理交易消息，根据交易消息执行相应的操作。

22. handlePooledTransactions66：处理待处理交易消息，将待处理交易添加到相应的处理队列中。

这些函数在以太坊协议中扮演着不同的角色，用于处理和响应特定类型的消息请求，从而支持以太坊网络的功能和交互。

