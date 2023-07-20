# File: core/rawdb/accessors_sync.go

在go-ethereum项目中，`core/rawdb/accessors_sync.go`文件的作用是提供了一些用于访问和操作区块链同步状态和头部的数据库函数。

下面是对每个函数的详细介绍：

1. `ReadSkeletonSyncStatus`: 此函数用于从数据库中读取区块链同步状态的骨架（Skeleton）。区块链同步状态骨架包含了当前区块链的同步信息，如最新的区块号、时间戳等。
   
2. `WriteSkeletonSyncStatus`: 此函数用于将区块链同步状态的骨架写入到数据库中。通过更新同步状态骨架，可以记录当前同步的进度和状态。

3. `DeleteSkeletonSyncStatus`: 此函数用于从数据库中删除区块链同步状态的骨架。

4. `ReadSkeletonHeader`: 此函数用于从数据库中读取区块链头部的骨架。区块链头部骨架包含了区块链中每个区块的头部信息，如区块的哈希、难度、时间戳等。

5. `WriteSkeletonHeader`: 此函数用于将区块链头部的骨架写入到数据库中。通过更新头部骨架，可以记录每个区块的头部信息，以便后续查询和验证。

6. `DeleteSkeletonHeader`: 此函数用于从数据库中删除区块链头部的骨架。

这些函数提供了对区块链同步状态和区块链头部的数据库访问和操作功能。通过读取和写入相关的骨架信息，可以更新和记录同步过程中的状态和区块头部数据，以支持区块链的同步和验证。

