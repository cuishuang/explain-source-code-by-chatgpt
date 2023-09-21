# File: grpc-go/internal/balancergroup/balancergroup.go

在grpc-go项目中，`balancergroup.go`文件的作用是实现负载均衡器组。

`subBalancerWrapper`结构体是负载均衡器的包装器，它实现了`Balancer`接口并包装了子负载均衡器。`subBalancerWrapper`具有以下函数：
- `UpdateState`：更新负载均衡器的状态。
- `Start`：启动负载均衡器组。
- `ExitIdle`：将负载均衡器组设置为空闲状态。
- `ExitIdleOne`：将单个负载均衡器设置为空闲状态。

`BalancerGroup`结构体是一个负载均衡器组，它维护一个负载均衡器的集合并将流量分配给它们。`BalancerGroup`具有以下函数：
- `New`：创建一个负载均衡器组。
- `Start`：启动负载均衡器组。
- `AddWithClientConn`：向负载均衡器组中添加一个负载均衡器，并与`ClientConn`关联。
- `Add`：向负载均衡器组中添加一个负载均衡器。
- `UpdateBuilder`：更新负载均衡器组的构建器。
- `UpdateSubConnState`：更新子连接的状态。
- `UpdateClientConnState`：更新客户端连接的状态。
- `ResolverError`：报告给定的解析器错误。
- `Close`：关闭负载均衡器组。

`Options`结构体是负载均衡器组的选项，用于配置负载均衡器组的行为。

`NewSubConn`函数创建子连接。
`updateBalancerStateWithCachedPicker`函数更新负载均衡器的状态并使用缓存的选择器。
`startBalancer`函数启动负载均衡器。
`exitIdle`函数将负载均衡器设置为退闲状态。
`updateClientConnState`函数更新客户端连接的状态。
`resolverError`函数报告解析器错误。
`gracefulSwitch`函数优雅切换负载均衡器组。
`stopBalancer`函数停止负载均衡器组。
`cleanupSubConns`函数清理子连接。
`connect`函数建立连接。
`updateSubConnState`函数更新子连接的状态。
`UpdateSubConnState`函数更新子连接的状态。
`UpdateClientConnState`函数更新客户端连接的状态。
`ResolverError`函数报告给定的解析器错误。
`newSubConn`函数创建子连接。
`updateBalancerState`函数更新负载均衡器的状态。
`Close`函数关闭负载均衡器组。
`ExitIdle`函数将负载均衡器组设置为退闲状态。
`ExitIdleOne`函数将单个负载均衡器设置为退闲状态。

