# File: istio/pilot/pkg/networking/core/v1alpha3/listener_builder.go

文件`listener_builder.go`的作用是构建Envoy的监听器配置，并根据Istio配置生成监听器的网络过滤器。

`ListenerBuilder`结构体是用于构建Envoy的监听器配置的主要结构体。`enabledInspector`结构体用于表示是否启用网络检查器。

`NewListenerBuilder`函数用于创建一个新的`ListenerBuilder`对象。

`appendSidecarInboundListeners`函数用于向Envoy配置中追加Sidecar的入站监听器。

`appendSidecarOutboundListeners`函数用于向Envoy配置中追加Sidecar的出站监听器。

`buildHTTPProxyListener`函数用于根据HTTP代理配置构建监听器。

`buildVirtualOutboundListener`函数用于构建虚拟服务的出站监听器。

`patchOneListener`函数用于修改单个监听器的配置。

`patchListeners`函数用于修改所有监听器的配置。

`getListeners`函数用于获取所有监听器的配置。

`buildOutboundCatchAllNetworkFiltersOnly`函数用于构建出站的网络过滤器。

`parseDuration`函数用于从字符串中解析出时间间隔。

`buildOutboundCatchAllNetworkFilterChains`函数用于构建所有出站的网络过滤器链。

`blackholeFilterChain`函数用于构建黑洞网络过滤器链。

`buildHTTPConnectionManager`函数用于构建HTTP连接管理器配置。

这些函数的作用是根据Istio配置构建Envoy配置中的监听器和网络过滤器，并以此来实现Istio的流量管理功能。

