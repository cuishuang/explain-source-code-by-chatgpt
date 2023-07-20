# File: core/blockchain_reader.go

在go-ethereum项目中，core/blockchain_reader.go文件负责提供区块链数据的读取功能。下面逐个介绍各个函数的作用：

1. CurrentHeader：返回当前区块链的最新区块头。
2. CurrentBlock：返回当前区块链的最新区块。
3. CurrentSnapBlock：返回当前区块链的最新快照区块。
4. CurrentFinalBlock：返回当前区块链的最新确认的区块。
5. CurrentSafeBlock：返回当前区块链的最新安全的区块。
6. HasHeader：检查是否存在指定哈希的区块头。
7. GetHeader：根据区块号获取指定区块的区块头。
8. GetHeaderByHash：根据区块哈希获取指定区块的区块头。
9. GetHeaderByNumber：根据区块号获取指定区块的区块头。
10. GetHeadersFrom：从指定区块开始，获取一定数量的区块头列表。
11. GetBody：根据区块号获取指定区块的区块体。
12. GetBodyRLP：根据区块号获取被编码为RLP的指定区块的区块体。
13. HasBlock：检查是否存在指定哈希的区块。
14. HasFastBlock：检查指定哈希的区块是否已完成快速同步。
15. GetBlock：根据区块号获取指定区块。
16. GetBlockByHash：根据区块哈希获取指定区块。
17. GetBlockByNumber：根据区块号获取指定区块。
18. GetBlocksFromHash：根据区块哈希获取指定数量的区块列表。
19. GetReceiptsByHash：根据区块哈希获取指定区块的所有交易收据。
20. GetUnclesInChain：获取指定区块的所有叔块。
21. GetCanonicalHash：根据区块号获取指定区块的规范哈希。
22. GetAncestor：根据区块号和偏移量获取指定区块的祖先区块。
23. GetTransactionLookup：根据交易哈希获取交易查询结果。
24. GetTd：根据区块号获取指定区块的总难度。
25. HasState：检查是否存在指定哈希的状态。
26. HasBlockAndState：检查是否存在指定区块和状态。
27. TrieNode：获取指定节点哈希的Trie节点。
28. ContractCodeWithPrefix：根据合约哈希获取带有前缀的合约代码。
29. State：返回指定区块的状态。
30. StateAt：返回指定区块号的状态。
31. Config：返回区块链的配置参数。
32. Engine：返回用于区块链验证的共识引擎。
33. Snapshots：返回区块链当前状态的快照。
34. Validator：返回用于验证新区块的验证器。
35. Processor：返回用于处理新区块的处理器。
36. StateCache：返回用于缓存区块状态的状态缓存。
37. GasLimit：返回当前区块链的当前燃料限制。
38. Genesis：返回当前区块链的创世区块。
39. GetVMConfig：返回当前区块链的虚拟机配置。
40. SetTxLookupLimit：设置交易查询结果的限制。
41. TxLookupLimit：返回交易查询结果的限制。
42. TrieDB：返回与当前区块链相关的Trie数据库。
43. SubscribeRemovedLogsEvent：订阅已删除日志的事件。
44. SubscribeChainEvent：订阅区块链事件。
45. SubscribeChainHeadEvent：订阅区块链头事件。
46. SubscribeChainSideEvent：订阅区块链侧链事件。
47. SubscribeLogsEvent：订阅区块链日志事件。
48. SubscribeBlockProcessingEvent：订阅区块处理事件。

这些函数通过读取区块链的相关数据，提供了获取区块、交易、状态以及其他相关信息的功能。

