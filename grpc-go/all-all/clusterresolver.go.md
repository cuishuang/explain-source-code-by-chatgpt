# File: grpc-go/xds/internal/balancer/clusterresolver/clusterresolver.go

clusterresolver.go文件的作用是实现了一个gRPC负载均衡器的集群解析器。这个解析器依赖于服务发现系统的配置，并根据最新的资源信息动态调整负载均衡的策略。

- errBalancerClosed变量：表示负载均衡器是否已关闭。
- newChildBalancer变量：用于创建子负载均衡器的函数。

- bb结构体：负责管理负载均衡器的所有子负载均衡器。
- ccUpdate结构体：表示`balancer.ClientConn`的最新状态。
- exitIdle结构体：表示是否处于退出空闲状态。
- clusterResolverBalancer结构体：负责管理与集群解析器相关的所有状态。
- ccWrapper结构体：用于包装`balancer.ClientConn`，以便捕捉状态变化。

- init函数：用于初始化clusterResolverBalancer的状态。
- Build函数：用于创建集群解析器。
- Name函数：返回集群解析器的名称。
- ParseConfig函数：用于解析配置文件。
- handleClientConnUpdate函数：处理`balancer.ClientConn`的状态更新。
- handleResourceUpdate函数：处理资源更新。
- updateChildConfig函数：更新子负载均衡器的配置。
- handleErrorFromUpdate函数：处理由资源更新引起的错误。
- run函数：开启集群解析器的定时任务。
- UpdateClientConnState函数：更新`balancer.ClientConn`的状态。
- ResolverError函数：处理解析器错误。
- UpdateSubConnState函数：更新`balancer.SubConn`的状态。
- Close函数：关闭负载均衡器。
- ExitIdle函数：退出空闲状态。
- ResolveNow函数：立即解析资源。

