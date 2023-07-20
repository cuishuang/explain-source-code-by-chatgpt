# File: eth/downloader/queue.go

在go-ethereum项目中，eth/downloader/queue.go文件的作用是实现以太坊区块的下载和管理功能。它是用于处理区块同步时的队列管理的核心文件。

下面是每个变量的作用：
- blockCacheMaxItems：用于设置缓存中最多存储的区块数量
- blockCacheInitialItems：用于设置缓存初始化时的区块数量
- blockCacheMemory：标识是否只将区块头部存储在缓存中
- blockCacheSizeWeight：用于衡量区块的大小
- errNoFetchesPending：表示没有要下载的区块错误
- errStaleDelivery：表示已经过期的交付错误

下面是每个结构体的作用：
- fetchRequest：表示一个要下载的区块请求
- fetchResult：表示一个已下载的区块结果
- queue：是一个区块下载队列，负责管理待下载、正在下载和已下载的区块

下面是每个函数的作用：
- newFetchResult：用于创建一个新的区块下载结果
- SetBodyDone：用于标记下载结果中的区块体已完成
- AllDone：判断是否所有的区块都已下载完成
- SetReceiptsDone：用于标记下载结果中的区块交易收据已完成
- Done：用于标记某个区块下载任务已完成
- newQueue：用于创建一个新的下载队列
- Reset：用于重置下载队列的状态
- Close：关闭下载队列
- PendingHeaders：返回待下载的区块头信息
- PendingBodies：返回待下载的区块体信息
- PendingReceipts：返回待下载的交易收据信息
- InFlightBlocks：返回正在下载的区块信息
- InFlightReceipts：返回正在下载的交易收据信息
- Idle：判断下载队列是否空闲
- ScheduleSkeleton：用于调度待下载区块的头信息
- RetrieveHeaders：获取指定高度和数量的区块头信息
- Schedule：将待下载的区块加入下载队列
- Results：返回已下载的区块结果
- Stats：返回下载队列的统计信息
- stats：用于更新下载队列的统计信息
- ReserveHeaders：预定待下载的区块头信息
- ReserveBodies：预定待下载的区块体信息
- ReserveReceipts：预定待下载的交易收据信息
- reserveHeaders：优化待下载的区块头信息列表
- Revoke：撤销对某个区块的下载预定
- ExpireHeaders：清除已经过期的区块头信息
- ExpireBodies：清除已经过期的区块体信息
- ExpireReceipts：清除已经过期的交易收据信息
- expire：清除已经过期的区块信息
- DeliverHeaders：将已下载的区块头信息进行交付
- DeliverBodies：将已下载的区块体信息进行交付
- DeliverReceipts：将已下载的交易收据信息进行交付
- deliver：将已下载的区块信息进行交付
- Prepare：准备下载指定区块的头信息及相关信息

