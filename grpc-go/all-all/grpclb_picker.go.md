# File: grpc-go/balancer/grpclb/grpclb_picker.go

grpc-go/balancer/grpclb/grpclb_picker.go 文件的作用是实现 gRPC 中负载均衡功能的 pickers（选项器）。该文件定义了用于负载均衡的结构体和函数。

- rpcStats：用于记录每个服务器的调用状态和延迟信息。
- rrPicker：轮询选项器，按照轮询的方式选择下一个可用服务器。
- lbPicker：负载均衡选项器，根据权重和调用状态来选择下一个可用服务器。

下面是这些结构体和函数的详细介绍：

- newRPCStats：创建一个新的 rpcStats 结构体，用于记录调用状态和延迟信息。
- isZeroStats：检查 rpcStats 是否为零值，即没有调用状态和延迟信息。
- toClientStats：将 rpcStats 转换为 balancer.ClientStats 结构体，用于报告给 balancer（负载均衡器）。
- drop：标记一个服务器不可用，在下一次负载均衡选择中将被剔除。
- failedToSend：标记一个服务器的调用未能发送成功，用于后续调用的负载均衡选择。
- knownReceived：标记一个服务器的调用已接收到响应，用于后续调用的负载均衡选择。
- newRRPicker：创建一个新的轮询选项器。
- Pick：选择下一个服务器，使用轮询方式从 rrPicker 中选择。
- newLBPicker：创建一个新的负载均衡选项器。
- updateReadySCs：更新可用的服务器列表，并根据权重和调用状态重新平衡负载，以便在后续调用中进行选择。

