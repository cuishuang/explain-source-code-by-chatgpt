# File: discovery/kubernetes/kubernetes.go

discovery/kubernetes/kubernetes.go文件是Prometheus项目中的一个文件，其主要作用是提供Kubernetes的服务发现功能。该文件中定义了一些变量、结构体和函数，用于实现与Kubernetes进行交互并获取服务监控的目标。

变量介绍：

1. userAgent：用于标识发送请求的客户端。
2. eventCount：用于计算事件的数量。
3. DefaultSDConfig：默认的服务发现配置。

结构体介绍：

1. Role：定义Prometheus的角色，可以是Prometheus Server、Alertmanager等。
2. SDConfig：服务发现的配置。
3. roleSelector：用于选择所需的角色。
4. SelectorConfig：标识了Prometheus的选项，如metricRelabelConfigs、interval等。
5. resourceSelector：选择Kubernetes资源的标识。
6. AttachMetadataConfig：将特定元数据附加到目标配置的配置。
7. NamespaceDiscovery：用于发现Kubernetes的命名空间。
8. Discovery：定义了服务发现的接口。

函数介绍：

1. init：用于初始化服务发现的配置。
2. UnmarshalYAML：将YAML格式的输入数据解析为结构体。
3. Name：返回服务发现的名称。
4. NewDiscoverer：创建一个新的服务发现器。
5. SetDirectory：设置服务发现的目录。
6. getNamespaces：获取Kubernetes中的命名空间。
7. New：创建一个新的服务发现器实例。
8. mapSelector：将选择器字符串转换为结构体。
9. Run：运行服务发现器。
10. lv：输出日志。
11. send：发送事件到Prometheus。
12. retryOnError：在出现错误后重试操作。
13. checkNetworkingV1Supported：检查是否支持NetworkingV1版本的API。
14. newNodeInformer：创建一个新的节点信息提供器。
15. newPodsByNodeInformer：创建一个新的Pod信息提供器。
16. newEndpointsByNodeInformer：创建一个新的Endpoints信息提供器。
17. newEndpointSlicesByNodeInformer：创建一个新的EndpointSlices信息提供器。
18. checkDiscoveryV1Supported：检查是否支持DiscoveryV1版本的API。

总的来说，discovery/kubernetes/kubernetes.go文件实现了Prometheus与Kubernetes进行集成，通过与Kubernetes API进行交互来发现并监控Kubernetes中的服务。它定义了获取命名空间、资源选择器、服务发现配置等功能，并提供了相应的函数来初始化、创建和运行服务发现器。

