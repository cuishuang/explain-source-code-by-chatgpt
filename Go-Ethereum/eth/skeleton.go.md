# File: eth/downloader/skeleton.go

eth/downloader/skeleton.go文件是go-ethereum项目中负责处理区块链同步的骨架模块。这个模块实现了下载、处理和验证区块链的功能。

以下是对每个变量和结构体的详细介绍：

- errSyncLinked：用于在同步期间标记链链接错误。
- errSyncMerged：用于在同步期间标记链合并错误。
- errSyncReorged：用于在同步期间标记链重组错误。
- errTerminated：用于在同步期间标记链同步被终止错误。
- errReorgDenied：用于在同步期间标记链重组被拒绝错误。

- subchain：表示一个子链，由一组区块组成。
- skeletonProgress：表示骨架模块的同步进度，用于跟踪同步的状态。
- headUpdate：用于更新头部区块。
- headerRequest：表示一个请求获取头部区块的任务。
- headerResponse：表示头部区块的响应。
- backfiller：用于回填区块数据。
- skeleton：整个骨架模块的主体结构体。

以下是对每个函数的详细介绍：

- init：初始化骨架模块。
- newSkeleton：创建一个新的骨架模块实例。
- startup：启动骨架模块，开始同步区块链。
- Terminate：终止骨架模块的同步。
- Sync：进行区块链同步的入口函数。
- sync：处理同步请求的主循环。
- initSync：初始化同步状态。
- saveSyncStatus：保存同步状态到磁盘。
- processNewHead：处理新的区块头。
- assignTasks：分配任务给下载器。
- executeTask：执行任务。
- revertRequests：回滚未完成的请求。
- scheduleRevertRequest：调度回滚请求。
- revertRequest：执行回滚请求。
- processResponse：处理来自下载器的响应。
- cleanStales：清除过时的请求。
- Bounds：获取骨架的区块链范围。
- Header：获取指定ID的区块头信息。

这些函数和变量的组合保证了go-ethereum项目中的区块链同步工作的顺利进行。

