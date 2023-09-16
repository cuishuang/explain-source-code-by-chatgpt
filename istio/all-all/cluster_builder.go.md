# File: istio/pilot/pkg/networking/core/v1alpha3/cluster_builder.go

在Istio项目中，`cluster_builder.go`文件位于`istio/pilot/pkg/networking/core/v1alpha3`路径下，它的主要作用是构建网络集群（cluster）的配置。

`passthroughHttpProtocolOptions`是用于定义HTTP传递协议选项的变量。它用于配置传递给后端的HTTP传输协议的参数。

`ClusterWrapper`结构体是对Envoy集群配置的封装，包含了集群名称、类型和其他相关配置信息。`metadataCerts`结构体是用于管理元数据证书的配置。`ClusterBuilder`是用于构建集群配置的主要结构体。`mtlsContextType`是定义了mTLS（mutual Transport Layer Security）上下文的类型。

以下是`cluster_builder.go`文件中提供的一些主要功能和函数的简要描述：
- `NewClusterBuilder`：创建一个新的集群构建器。
- `String`：将集群配置转换为字符串。
- `newClusterWrapper`：创建一个新的集群封装。
- `sidecarProxy`：为Sidecar代理构建集群。
- `buildSubsetCluster`：构建带有子集的集群。
- `applyDestinationRule`：应用DestinationRule配置到集群。
- `applyMetadataExchange`：应用元数据交换配置到集群。
- `buildCluster`：构建集群。
- `buildInboundCluster`：构建入站集群。
- `buildLocalityLbEndpoints`：构建本地负载均衡端点。
- `addUint32`：添加32位无符号整数到集群配置中。
- `buildInboundPassthroughClusters`：构建入站直通集群。
- `buildBlackHoleCluster`：构建黑洞集群。
- `buildDefaultPassthroughCluster`：构建默认的直通集群。
- `setH2Options`：设置HTTP/2选项。
- `setUseDownstreamProtocol`：设置使用下游协议。
- `http2ProtocolOptions`：获取配置的HTTP/2协议选项。
- `isHttp2Cluster`：检查集群是否支持HTTP/2协议。
- `setUpstreamProtocol`：设置上游协议。
- `normalizeClusters`：规范化集群配置。
- `getAllCachedSubsetClusters`：获取所有缓存的子集集群。
- `build`：构建所有集群并返回。
- `CastDestinationRule`：将配置转换为DestinationRule类型。
- `maybeApplyEdsConfig`：可能应用EDS（Endpoint Discovery Service）配置到集群。
- `buildExternalSDSCluster`：构建外部SDS（Service-Discovery Service）集群。
- `addTelemetryMetadata`：为遥测元数据添加配置。

以上是对`cluster_builder.go`文件的概述，它在Istio中扮演着构建网络集群的重要角色，并提供了一系列函数和结构体来管理和配置集群。

