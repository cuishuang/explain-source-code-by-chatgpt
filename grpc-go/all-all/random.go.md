# File: grpc-go/internal/wrr/random.go

在grpc-go项目中，`random.go`文件的作用是实现基于加权随机轮询算法（Weighted Random Round-Robin）的负载均衡器。该负载均衡器用于在调用gRPC服务时选择合适的服务器。

`grpcrandInt63n`变量是一个`randInt63n`函数，它是一个返回不超过给定参数的非负随机int64值的函数。它由`math/rand`包提供。

`weightedItem`结构体用于表示具有权重的项。它包含一个`Value`字段，用于存储实际的项，以及一个`Weight`字段，用于表示该项的权重。

`randomWRR`结构体是加权随机轮询的负载均衡器。它包含一个`items`字段，用于存储所有的`weightedItem`，以及一个`sum`字段，用于表示所有项的总权重。

`String`函数用于将负载均衡器的状态以字符串形式返回，用于调试目的。

`NewRandom`函数返回一个初始状态的负载均衡器。

`Next`函数根据加权随机轮询算法从负载均衡器中选择下一个项，并返回该项。

`Add`函数向负载均衡器中添加一个新的项，并更新总权重的计数。

