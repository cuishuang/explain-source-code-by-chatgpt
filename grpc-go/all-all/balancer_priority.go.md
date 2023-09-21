# File: grpc-go/xds/internal/balancer/priority/balancer_priority.go

`balancer_priority.go` 文件是 `grpc-go` 中负载均衡器优先级相关的代码实现。

`ErrAllPrioritiesRemoved` 是一个错误变量，表示所有优先级都被移除了。

`DefaultPriorityInitTimeout` 是一个默认的优先级初始化超时时间。

`syncPriority` 函数用于同步更新优先级，它会检查当前的优先级列表和最新的优先级列表，并进行相应的添加或移除操作。

`stopSubBalancersLowerThanPriority` 函数会停止所有低于指定优先级的子负载均衡器。

`switchToChild` 函数用于切换到指定的子负载均衡器，同时监听其状态变化。

`handleChildStateUpdate` 函数用于处理子负载均衡器状态的变化，根据不同的变化类型做出相应的操作，如添加、移除或更新子负载均衡器。

