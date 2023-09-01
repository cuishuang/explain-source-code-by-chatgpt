# File: client-go/discovery/cached/disk/cached_discovery.go

在K8s组织下的client-go项目中，client-go/discovery/cached/disk/cached_discovery.go文件的作用是提供一个可缓存的发现客户端，用于向Kubernetes API服务器发出请求以获取关于集群中资源和服务的信息。

下面介绍文件中的重要部分：

- _这几个变量：这些是占位符变量，用于忽略某些返回值或参数。

- CachedDiscoveryClient：一个存储在磁盘上的缓存发现客户端。它实现了DiscoveryClient接口，并能够从磁盘读取缓存数据，代替每次请求向API服务器发出网络请求。

- ServerResourcesForGroupVersion：向API服务器请求获取指定Group和Version下的资源信息，并返回ResourceList对象。

- ServerGroupsAndResources：向API服务器请求获取集群中所有API组的信息，以及每个API组下的资源信息，并返回GroupResources对象。

- ServerGroups：向API服务器请求获取集群中所有API组的信息，并返回GroupList对象。

- getCachedFile：从磁盘读取缓存数据，返回字节数组。

- writeCachedFile：将字节数组写入磁盘，作为缓存数据。

- RESTClient：用于与API服务器进行通信的HTTP REST客户端。

- ServerPreferredResources：根据优先级返回推荐的资源列表。

- ServerPreferredNamespacedResources：类似于ServerPreferredResources，但只返回命名空间资源列表。

- ServerVersion：向API服务器请求获取集群的版本信息，并返回ServerVersionInfo对象。

- OpenAPISchema：向API服务器请求获取OpenAPI V2 Schema，并返回OpenAPISchema对象。

- OpenAPIV3：向API服务器请求获取OpenAPI V3 Schema，并返回OpenAPISchema对象。

- Fresh：检查缓存是否过期，如果过期则将缓存数据删除。

- Invalidate：使缓存数据无效，强制客户端重新加载数据。

- WithLegacy：将指定的发现客户端包装为CachedDiscoveryClient。

- NewCachedDiscoveryClientForConfig：使用给定的配置创建一个新的CachedDiscoveryClient。

- newCachedDiscoveryClient：根据给定的配置创建一个新的CachedDiscoveryClient的构造函数。

以上这些函数和结构体提供了与API服务器进行交互并提供缓存功能的能力，使得应用程序能够更高效地获取集群信息。

