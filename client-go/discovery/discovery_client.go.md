# File: client-go/discovery/discovery_client.go

在K8s组织下的client-go项目中，client-go/discovery/discovery_client.go是用于处理Kubernetes API的发现功能。它提供了与集群交互并获取有关可用的API组、版本、资源和服务器信息的方法。

下面是相关变量和结构体的作用：

- DiscoveryInterface：定义了执行与API发现相关操作的接口。
- AggregatedDiscoveryInterface：定义了聚合式API发现相关操作的接口。
- CachedDiscoveryInterface：定义了缓存API发现相关操作的接口。
- ServerGroupsInterface：定义了获取API组信息的接口。
- ServerResourcesInterface：定义了获取API组和版本相关资源信息的接口。
- ServerVersionInterface：定义了获取服务器版本信息的接口。
- OpenAPISchemaInterface：定义了获取OpenAPI模式信息的接口。
- OpenAPIV3SchemaInterface：定义了获取OpenAPI V3模式信息的接口。
- DiscoveryClient：实现了DiscoveryInterface接口，用于与集群进行API发现。
- ErrGroupDiscoveryFailed：表示组发现失败的错误。

以下是一些重要的函数和方法：

- apiVersionsToAPIGroup：根据API版本返回匹配的API组。
- GroupsAndMaybeResources：返回API组及其相关资源的信息。
- downloadLegacy：从服务器上下载API组信息的原始数据。
- downloadAPIs：从服务器上下载API组的信息。
- isV2Beta1ContentType：判断Content-Type是否为v2beta1。
- ServerGroups：获取所有API组的信息。
- ServerResourcesForGroupVersion：获取给定API组和版本的资源信息。
- ServerGroupsAndResources：获取所有API组及其相关资源的信息。
- Error：表示获取API失败的错误。
- IsGroupDiscoveryFailedError：判断是否为组发现失败的错误。
- GroupDiscoveryFailedErrorGroups：返回组发现失败的API组信息。
- ServerPreferredResources：根据API版本返回首选的资源列表。
- fetchGroupVersionResources：获取给定组和版本的资源信息。
- ServerPreferredNamespacedResources：根据API版本返回首选的命名空间资源列表。
- ServerVersion：获取服务器的版本信息。
- OpenAPISchema：获取服务器的OpenAPI模式信息。
- OpenAPIV3：获取服务器的OpenAPI V3模式信息。
- WithLegacy：使用遗留方式进行发现。
- withRetries：在遇到临时错误时进行重试。
- setDiscoveryDefaults：设置API发现的默认值。
- NewDiscoveryClientForConfig：基于给定的配置创建DiscoveryClient。
- NewDiscoveryClientForConfigAndClient：基于给定的配置和客户端创建DiscoveryClient。
- NewDiscoveryClientForConfigOrDie：基于给定的配置创建DiscoveryClient。如果失败则终止程序。
- NewDiscoveryClient：创建默认配置的DiscoveryClient。
- RESTClient：在给定的HTTP传输参数下创建一个进行REST操作的客户端。

这些函数和方法提供了从Kubernetes集群中获取API组、版本、资源和服务器信息的能力，并且可以让开发者更方便地与集群进行交互。

