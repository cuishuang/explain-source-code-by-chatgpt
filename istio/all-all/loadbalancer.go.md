# File: istio/pilot/pkg/networking/core/v1alpha3/loadbalancer/loadbalancer.go

istio/pilot/pkg/networking/core/v1alpha3/loadbalancer/loadbalancer.go是Istio项目中负责负载均衡的文件。它定义了一些结构体和函数，用于处理负载均衡策略。

该文件中的主要结构体是WrappedLocalityLbEndpoints。WrappedLocalityLbEndpoints是用于将LocalityLbEndpoints结构体包装在一起的数据结构。它包含了包装后的LocalityLbEndpoints和其他一些辅助信息，例如权重、优先级等。

以下是文件中重要函数的作用：

1. GetLocalityLbSetting：根据传入的负载均衡设置和服务名称，获取本地负载均衡的配置。返回一个包装了LocalityLbEndpoints的WrappedLocalityLbEndpoints结构体。

2. ApplyLocalityLBSetting：将LocalityLbSetting应用到传入的网络负载均衡器实例。这个函数会根据设置配置每个权重、故障转移等。

3. applyLocalityWeight：根据LocalityLbEndpoints的权重设置，将权重应用到传入的网络负载均衡器实例。根据权重，决定转发请求到哪个终端节点。

4. applyLocalityFailover：在发生故障时，根据故障转移策略，将请求转发到其他可用的终端节点。

5. applyPriorityFailover：在发生故障时，在不同优先级的终端节点之间进行故障转移。

6. priorityLabelOverrides：返回一个根据优先级标签的权重映射。用于根据优先级调整终端节点的服务。

7. applyPriorityFailoverPerLocality：为每个本地性调整优先级故障转移。

以上这些函数一起工作，通过配置权重、故障转移策略和优先级，实现了Istio的负载均衡功能。通过这些函数，可以根据传入的设置，将传入的请求分发到不同的终端节点，以实现服务的负载均衡和故障恢复。

