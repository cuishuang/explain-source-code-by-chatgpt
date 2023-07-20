# File: eth/fetcher/block_fetcher.go

在go-ethereum项目中，`eth/fetcher/block_fetcher.go`文件的作用是实现区块的获取和处理功能。该文件包含了一系列变量、结构体和函数，用于管理和执行区块的下载、过滤、验证和广播等任务。

以下是对每个变量的作用进行详细介绍：

- `blockAnnounceInMeter`、`blockAnnounceOutTimer`、`blockAnnounceDropMeter`、`blockAnnounceDOSMeter`：用于计量和监控区块广播的输入、输出、丢弃和拒绝服务等情况。
- `blockBroadcastInMeter`、`blockBroadcastOutTimer`、`blockBroadcastDropMeter`、`blockBroadcastDOSMeter`：用于计量和监控区块传播的输入、输出、丢弃和拒绝服务等情况。
- `headerFetchMeter`、`bodyFetchMeter`：用于计量获取区块头和区块体的速率。
- `headerFilterInMeter`、`headerFilterOutMeter`：用于计量区块头过滤的输入和输出数量。
- `bodyFilterInMeter`、`bodyFilterOutMeter`：用于计量区块体过滤的输入和输出数量。
- `errTerminated`：用于标志是否由于错误而终止区块下载过程。

以下是对每个结构体的作用进行详细介绍：

- `HeaderRetrievalFn`：定义了获取区块头的函数类型。
- `blockRetrievalFn`：定义了获取区块体的函数类型。
- `headerRequesterFn`：定义了向其他节点请求区块头的函数类型。
- `bodyRequesterFn`：定义了向其他节点请求区块体的函数类型。
- `headerVerifierFn`：定义了验证区块头的函数类型。
- `blockBroadcasterFn`：定义了区块广播的函数类型。
- `chainHeightFn`：定义了获取链的当前高度的函数类型。
- `headersInsertFn`：定义了将区块头插入区块链的函数类型。
- `chainInsertFn`：定义了将区块插入区块链的函数类型。
- `peerDropFn`：定义了当与某个节点的连接中断时的回调函数类型。
- `blockAnnounce`：区块广播的数据结构。
- `headerFilterTask`：用于表示需要过滤的区块头的任务。
- `bodyFilterTask`：用于表示需要过滤的区块体的任务。
- `blockOrHeaderInject`：区块或区块头注入的数据结构。
- `BlockFetcher`：区块下载器的主要结构体，包含了上述各种功能的实现。

以下是对每个函数的作用进行详细介绍：

- `number`、`hash`：用于返回当前区块的高度和哈希值。
- `NewBlockFetcher`：创建一个新的区块下载器实例。
- `Start`：启动区块下载器，开始执行区块获取任务。
- `Stop`：停止区块下载器，终止当前正在进行的任务。
- `Notify`：处理来自其他节点的新区块通知。
- `Enqueue`：将区块头或区块请求加入任务队列等待处理。
- `FilterHeaders`：过滤区块头。
- `FilterBodies`：过滤区块体。
- `loop`：循环执行区块下载任务。
- `rescheduleFetch`、`rescheduleComplete`、`enqueue`：用于重新安排区块下载任务的执行。
- `importHeaders`、`importBlocks`：导入区块头和区块。
- `forgetHash`、`forgetBlock`：用于从下载器的内存中删除已完成或取消的区块头或区块。

