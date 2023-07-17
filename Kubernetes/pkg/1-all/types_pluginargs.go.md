# File: pkg/scheduler/apis/config/types_pluginargs.go

在Kubernetes项目中，`pkg/scheduler/apis/config/types_pluginargs.go`文件的作用是定义了调度器插件（scheduler plugin）的配置参数类型。

下面是对这些结构体的详细介绍：

1. `DefaultPreemptionArgs`：默认预抢占参数。预抢占是指当资源不足时，优先级较低的 Pod 可能会被强制调度到其他节点上，以便为优先级较高的 Pod 留出资源。

2. `InterPodAffinityArgs`：Pod 间亲和性参数。Pod 间亲和性是指通过标签或选择器来定义多个 Pod 之间的关系，以便它们被调度到同一节点上。

3. `NodeResourcesFitArgs`：节点资源适配参数。这些参数用于指定调度器如何评估节点上可用资源与 Pod 请求资源之间的匹配度，以便决定 Pod 是否可以调度到该节点上。

4. `PodTopologySpreadConstraintsDefaulting`：Pod 拓扑分布约束默认设置。用于定义 Pod 在集群中的拓扑分布限制，以防止它们过于集中在同一地区或机架上。

5. `PodTopologySpreadArgs`：Pod 拓扑分布参数。用于指定拓扑分布约束的策略和参数，以确保在调度期间 Pod 被适当地分布到不同的地区或机架上。

6. `NodeResourcesBalancedAllocationArgs`：节点资源平衡分配参数。用于指定在调度期间如何选择节点，以便在节点上实现资源的均衡分配。

7. `UtilizationShapePoint`：资源利用率曲线参数。用于定义资源利用率曲线中的一个数据点，该曲线表示节点资源的使用情况。

8. `ResourceSpec`：资源规格参数。用于指定要分配给 Pod 的资源的规格，如 CPU、内存等。

9. `VolumeBindingArgs`：卷绑定参数。用于指定如何将卷（volume）绑定到节点上的持久存储。

10. `NodeAffinityArgs`：节点亲和性参数。用于指定调度器如何根据节点上的标签或筛选器来选择合适的节点。

11. `ScoringStrategyType`：评分策略类型。用于指定调度器使用的评分算法的类型，如负载均衡、权重等。

12. `ScoringStrategy`：评分策略。用于指定调度器如何评估节点和 Pod 之间的匹配度，并为可行的调度候选项分配相应的分数。

13. `RequestedToCapacityRatioParam`：请求资源与节点容量比例参数。用于定义调度器如何根据节点的资源容量和 Pod 的资源需求来评估节点与 Pod 之间的匹配度。

