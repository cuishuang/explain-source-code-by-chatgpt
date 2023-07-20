# File: les/downloader/downloader.go

在go-ethereum项目中，les/downloader/downloader.go文件是实现了下载和同步以太坊区块链的下载器模块。下面将对其中的变量和函数逐一进行介绍：

变量作用：
- MaxBlockFetch: 单次请求最大区块数量。
- MaxHeaderFetch: 单次请求最大头部数量。
- MaxSkeletonSize: 单次请求最大骨架大小。
- MaxReceiptFetch: 单次请求最大Receipt数量。
- MaxStateFetch: 单次请求最大状态数据大小。
- maxQueuedHeaders: 下载队列中最大保存的头部数量。
- maxHeadersProcess: 最大处理的头部数量。
- maxResultsProcess: 最大处理结果数量。
- fullMaxForkAncestry: 全节点最大Fork祖先数量。
- lightMaxForkAncestry: 轻节点最大Fork祖先数量。
- reorgProtThreshold: 收到一个无效区块，需要向邻居请求区块的阈值。
- reorgProtHeaderDelay: 等待收到BlockBodies的时间窗口。
- fsHeaderSafetyNet: 用于保护状态根不发生变化的头部数量。
- fsHeaderContCheck: 完整同步模式下检查状态一致性的头部数量。
- fsMinFullBlocks: 完整同步模式下至少需要的区块数量。
- errBusy: 下载器忙的错误提示。
- errUnknownPeer: 未知对等节点错误提示。
- errBadPeer: 错误的对等节点错误提示。
- errStallingPeer: 阻塞的对等节点错误提示。
- errUnsyncedPeer: 未同步的对等节点错误提示。
- errNoPeers: 无可用对等节点错误提示。
- errTimeout: 超时错误提示。
- errEmptyHeaderSet: 空头部集错误提示。
- errPeersUnavailable: 对等节点不可用错误提示。
- errInvalidAncestor: 无效的祖先错误提示。
- errInvalidChain: 无效链错误提示。
- errInvalidBody: 无效的区块体错误提示。
- errInvalidReceipt: 无效的Receipt错误提示。
- errCancelStateFetch: 取消状态数据请求错误提示。
- errCancelContentProcessing: 取消内容处理错误提示。
- errCanceled: 取消错误提示。
- errNoSyncActive: 没有激活的同步错误提示。
- errTooOld: 区块过旧错误提示。
- errNoAncestorFound: 找不到祖先错误提示。

结构体作用：
- Downloader: 下载器结构体，实现了以太坊区块链的下载和同步功能。
- LightChain: 轻客户端链结构体，用于保存轻客户端的链状态。
- BlockChain: 区块链结构体，用于保存和管理完整节点的链状态。

函数作用：
- New: 创建一个下载器实例。
- Progress: 返回当前下载进度。
- Synchronising: 返回是否正在同步中。
- RegisterPeer: 注册一个完整节点对等。
- RegisterLightPeer: 注册一个轻节点对等。
- UnregisterPeer: 取消注册一个对等。
- Synchronise: 启动同步过程。
- synchronise: 执行同步逻辑。
- getMode: 返回当前下载器模式（完整同步或快速同步）。
- syncWithPeer: 与对等节点执行同步。
- spawnSync: 同步节点。
- cancel: 取消同步过程。
- Cancel: 外部调用的取消同步方法。
- Terminate: 终止同步过程。
- fetchHead: 请求最新头部。
- calculateRequestSpan: 根据最近的一个块的头部信息计算请求的起始高度和结束高度。
- findAncestor: 查找某一高度的祖先区块。
- findAncestorSpanSearch: 使用遍历搜索查找祖先区块。
- findAncestorBinarySearch: 使用二分搜索查找祖先区块。
- fetchHeaders: 请求区块头部。
- fillHeaderSkeleton: 填充头部骨架。
- fetchBodies: 请求区块体。
- fetchReceipts: 请求Receipts。
- fetchParts: 请求部分区块数据。
- processHeaders: 处理接收到的头部数据。
- processFullSyncContent: 处理完整同步内容。
- importBlockResults: 导入区块结果。
- processFastSyncContent: 处理快速同步内容。
- splitAroundPivot: 根据主链分割轻节点链。
- commitFastSyncData: 提交快速同步数据。
- commitPivotBlock: 提交主链的轻链块。
- DeliverHeaders: 传递头部数据给上层。
- DeliverBodies: 传递区块体数据给上层。
- DeliverReceipts: 传递Receipts数据给上层。
- DeliverNodeData: 传递节点数据给上层。
- DeliverSnapPacket: 传递快照数据给上层。
- deliver: 传递数据给上层。

