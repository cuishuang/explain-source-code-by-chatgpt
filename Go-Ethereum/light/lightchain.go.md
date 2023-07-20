# File: light/lightchain.go

在go-ethereum项目中，light/lightchain.go文件是实现轻客户端链的核心文件。它定义了LightChain结构体以及与轻客户端相关的函数。

bodyCacheLimit和blockCacheLimit是用于配置缓存限制的参数。bodyCacheLimit限制了在轻客户端中缓存的区块体的数量，而blockCacheLimit限制了缓存的完整区块的数量。当超过这些限制时，较早的区块将从缓存中删除。

LightChain结构体是轻客户端链的核心数据结构，用于管理轻客户端的区块链状态和链的操作。它包含了一些重要的成员变量和函数。

- NewLightChain是一个初始化LightChain结构体的函数，用于创建一个新的轻客户端链。
- getProcInterrupt返回LightChain结构体中的链头进程中断信号。
- Odr是一个用于排序的排序函数。
- HeaderChain返回LightChain结构体中的区块头链。
- loadLastState用于加载最近状态。
- SetHead用于设置链的头块。
- SetHeadWithTimestamp用于设置具有给定时间戳的链的头块。
- GasLimit返回链的当前燃料限制。
- Reset用于重置链数据，以清除缓存和重置内部状态。
- ResetWithGenesisBlock用给定的创世区块重置链数据。
- Engine返回轻客户端链的区块引擎。
- Genesis返回链的创世区块。
- StateCache返回链的状态缓存。
- GetBody用给定的区块哈希返回完整的区块体。
- GetBodyRLP用给定的区块哈希返回序列化的区块体。
- HasBlock检查给定的区块哈希是否在链中。
- GetBlock返回给定区块哈希的区块。
- GetBlockByHash返回给定区块哈希的区块。
- GetBlockByNumber返回给定区块号的区块。
- Stop停止区块链处理。
- StopInsert停止插入新区块。
- Rollback回滚链到给定区块号。
- InsertHeader插入一个区块头。
- SetCanonical将给定块标记为链的正文。
- InsertHeaderChain按顺序插入区块头链。
- CurrentHeader返回链的当前头块。
- GetTd返回给定区块哈希的区块总难度。
- GetTdOdr返回给定区块哈希的区块总难度，通过排序返回。
- GetHeader返回给定区块哈希的区块头。
- GetHeaderByHash返回给定区块哈希的区块头。
- HasHeader检查给定的区块哈希是否在链中。
- GetCanonicalHash返回链的正文哈希。
- GetAncestor返回给定块号和块哈希的祖先区块。
- GetHeaderByNumber返回给定块号的区块头。
- GetHeaderByNumberOdr返回给定排序的块号的区块头。
- Config返回链的配置。
- LockChain锁住链，以防止并发修改。
- UnlockChain解锁链。
- SubscribeChainEvent订阅链事件。
- SubscribeChainHeadEvent订阅链头事件。
- SubscribeChainSideEvent订阅链侧事件。
- SubscribeLogsEvent订阅日志事件。
- SubscribeRemovedLogsEvent订阅移除的日志事件。

这些函数一起构成了轻客户端链的核心功能，用于管理轻客户端的区块链状态并提供各种链操作的方法。

