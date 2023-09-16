# File: istio/pilot/pkg/xds/endpoints/ep_filters.go

在Istio项目中，`istio/pilot/pkg/xds/endpoints/ep_filters.go` 文件负责实现一些用于筛选和处理终结点（endpoints）的过滤器。

以下是各个函数的详细介绍：

1. `EndpointsByNetworkFilter`：根据网络筛选终结点。它接受终结点的列表，并筛选出与所提供网络一致的终结点。这个函数允许将终结点按照网络进行分组，以便能够对特定网络的终结点应用特定的处理逻辑。

2. `selectNetworkGateways`：根据网络选择网关。在一个网络中可以有多个网关，这个函数会根据所提供的终结点和网关信息，选择在当前网络中运行的网关。

3. `scaleEndpointLBWeight`：调整终结点的负载均衡权重。它会根据网关配置和终结点的负载均衡权重，对终结点进行调整，以确保每个终结点的权重之和等于网关的总权重。

4. `splitWeightAmongGateways`：将权重在多个网关之间划分。当终结点被多个网关同时使用时，这个函数会根据每个网关所配置的权重，将终结点的权重按比例划分给每个网关。

5. `EndpointsWithMTLSFilter`：根据M-TLS（Mutual TLS）筛选终结点。它接受终结点的列表，并筛选出已经启用M-TLS的终结点。这个函数允许将只有启用了M-TLS的终结点用于特定的处理逻辑，如安全传输等。

这些函数在 Isto 中的用途是对终结点进行筛选、优化和适配，以确保终结点的选择和权重分配能够满足相关的网络需求，并提供一致的负载均衡和安全性。

