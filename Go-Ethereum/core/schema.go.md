# File: core/rawdb/schema.go

core/rawdb/schema.go文件的主要作用是定义了以太坊区块链数据库的Schema（模式），其中包含了各种数据库键（key）的定义和相关操作函数。

以下是对其中提到的一些关键变量和函数的详细介绍：

变量：
1. databaseVersionKey: 存储数据库版本的键。
2. headHeaderKey: 存储当前区块链头部的键。
3. headBlockKey: 存储当前区块链最新区块的键。
4. headFastBlockKey: 存储当前区块链最新快速同步区块的键。
5. headFinalizedBlockKey: 存储当前区块链最新确认区块的键。
6. lastPivotKey: 存储上一个区块链分叉点的键。
7. fastTrieProgressKey: 存储快速同步过程中Trie的进度的键。
8. snapshotDisabledKey: 存储是否禁用数据库快照功能的键。
9. SnapshotRootKey: 存储数据库快照根节点的键。
10. SnapshotJournalKey: 存储数据库快照日志的键。
11. SnapshotGeneratorKey: 存储数据库快照生成器的键。
12. SnapshotRecoveryKey: 存储数据库快照恢复器的键。
13. SnapshotSyncStatusKey: 存储数据库快照同步状态的键。
14. SkeletonSyncStatusKey: 存储骨架同步状态的键。
15. txIndexTailKey: 存储交易索引尾部的键。
16. fastTxLookupLimitKey: 存储快速交易查询限制的键。
17. badBlockKey: 存储坏区块的键。
18. uncleanShutdownKey: 存储非正常关闭状态的键。
19. transitionStatusKey: 存储区块链状态过渡状态的键。
20. headerPrefix, headerTDSuffix, headerHashSuffix, headerNumberPrefix: 区块头相关键的前缀和后缀。
21. blockBodyPrefix, blockReceiptsPrefix: 区块主体和交易收据相关键的前缀。
22. txLookupPrefix: 交易查询键的前缀。
23. bloomBitsPrefix: Bloom过滤器位图键的前缀。
24. SnapshotAccountPrefix, SnapshotStoragePrefix: 数据库快照中账户和存储键的前缀。
25. CodePrefix: 字节码键的前缀。
26. skeletonHeaderPrefix: 骨架头部键的前缀。
27. trieNodeAccountPrefix, trieNodeStoragePrefix: Trie节点中账户和存储键的前缀。
28. PreimagePrefix: 预映像键的前缀。
29. configPrefix: 配置键的前缀。
30. genesisPrefix: 创世区块键的前缀。
31. BloomBitsIndexPrefix: Bloom过滤器位图索引键的前缀。
32. ChtPrefix, ChtTablePrefix, ChtIndexTablePrefix: Cht（Consecutive Hash Tries）相关键的前缀。
33. BloomTriePrefix, BloomTrieTablePrefix, BloomTrieIndexPrefix: Bloom Trie相关键的前缀。
34. CliqueSnapshotPrefix: Clique快照的前缀。
35. preimageCounter, preimageHitCounter: 预映像计数器和命中计数器的键。

结构体：
1. LegacyTxLookupEntry: 用于存储已弃用的交易查询条目的结构体。

函数：
1. encodeBlockNumber: 将区块号编码为字节数组。
2. headerKeyPrefix: 根据区块头哈希前缀获取相应的键。
3. headerKey, headerTDKey, headerHashKey, headerNumberKey: 根据区块头的哈希、总难度、哈希、高度获取相应的键。
4. blockBodyKey: 根据区块哈希获取区块主体键。
5. blockReceiptsKey: 根据区块哈希获取区块交易收据键。
6. txLookupKey: 根据交易哈希获取交易查询键。
7. accountSnapshotKey, storageSnapshotKey, storageSnapshotsKey, bloomBitsKey: 根据哈希等信息获取相应的快照和Bloom位图键。
8. skeletonHeaderKey: 根据骨架头部哈希获取键。
9. preimageKey: 根据哈希获取预映像键。
10. codeKey: 根据代码哈希获取代码键。
11. IsCodeKey: 判断给定键是否为代码键。
12. configKey: 根据配置名获取配置键。
13. genesisStateSpecKey: 获取创世状态规范键。
14. accountTrieNodeKey, storageTrieNodeKey: 根据节点哈希获取账户和存储Trie节点键。
15. IsLegacyTrieNode, IsAccountTrieNode, IsStorageTrieNode: 判断给定键是否为已弃用的Trie节点类型。

