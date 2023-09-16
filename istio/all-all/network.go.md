# File: istio/pilot/pkg/model/network.go

在istio项目中，`istio/pilot/pkg/model/network.go`文件定义了与网络相关的结构体和函数，主要用于管理网关和网络配置。

1. `MinGatewayTTL`和`MaxGatewayTTL`变量分别用于定义网关的最小和最大生存时间。
2. `NetworkGatewayTestDNSServers`变量定义了用于测试网络网关的DNS服务器地址列表。

以下是`networkAndCluster`结构体中的变量的作用：
- `NetworkGateway`定义了网络网关规范。
- `NetworkGatewaysWatcher`用于监听网络网关的变化。
- `NetworkGatewaysHandler`用于处理网络网关的变化。
- `NetworkGateways`存储已知的网络网关列表。
- `NetworkManager`用于管理网络网关的增删改查等操作。
- `networkAndCluster`将网络和集群信息关联起来。
- `NetworkGatewaySet`定义了网络网关的集合。
- `networkGatewayNameCache`用于缓存网络网关的名称和地址信息。
- `nameCacheEntry`定义了网络网关名称和地址的缓存条目。
- `dnsClient`用于与DNS服务器进行通信。

以下是`networkAndCluster`结构体中的函数的作用：
- `AppendNetworkGatewayHandler`用于向网络网关处理列表中添加处理器。
- `NotifyGatewayHandlers`用于通知网络网关处理器执行网关相关的操作。
- `NewNetworkManager`创建一个新的网络管理器实例。
- `reloadGateways`用于重新加载网关列表。
- `reload`用于重新加载网络网关并通知监听器。
- `update`用于更新网络网关。
- `resolveHostnameGateways`用于解析主机名所关联的网关。
- `IsMultiNetworkEnabled`判断是否启用了多网络模式。
- `GetLBWeightScaleFactor`用于获取负载均衡的权重因子。
- `AllGateways`返回所有的网络网关。
- `allGateways`返回所有的网络网关。
- `GatewaysForNetwork`根据网络名称返回相应的网络网关。
- `GatewaysForNetworkAndCluster`根据网络名称和集群名称返回相应的网络网关。
- `networkAndClusterForGateway`根据网络网关名称返回相应的网络和集群信息。
- `networkAndClusterFor`根据网络名称和集群名称返回相应的网络和集群信息。
- `SortGateways`对网关进行排序。
- `gcd`计算最大公约数。
- `lcm`计算最小公倍数。
- `newNetworkGatewayNameCache`创建一个新的网络网关名称缓存。
- `newNetworkGatewayNameCacheWithClient`创建一个新的网络网关名称缓存，并指定DNS客户端。
- `Resolve`用于解析网络网关的名称和地址。
- `cleanupWatches`用于清理非法的观察器。
- `resolveFromCache`从缓存中解析网络网关的名称和地址。
- `resolveAndCache`解析网络网关的名称和地址，并缓存到缓存中。
- `refreshAndNotify`刷新并通知网络网关的信息。
- `resolve`用于解析网关的名称和地址。
- `minimalTTL`用于计算最小的生存时间。
- `newClient`创建一个新的DNS客户端。
- `getReqNames`获取网络网关名称的请求列表。
- `Query`用于查询网络网关的名称和地址。

