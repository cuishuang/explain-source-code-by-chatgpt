# File: pkg/controller/nodelifecycle/node_lifecycle_controller.go

该文件是 Kubernetes 中节点生命周期控制器的实现。

在 Kubernetes 中，节点生命周期控制器的作用是确保节点在 Kubernetes 中可用，同时维护节点的状态和各种配置。

UnreachableTaintTemplate 和 NotReadyTaintTemplate 是用于标记不可达和不可用节点的 taint 模板。

nodeConditionToTaintKeyStatusMap 是将节点条件映射到 taint 状态的 map。

taintKeyToNodeConditionMap 是将 taint 键映射到节点条件的 map。

labelReconcileInfo 用于记录当前正在处理的节点标签的信息。

ZoneState 用于记录节点集群的各个 zone 的状态。

nodeHealthData 用于记录节点的健康数据。

nodeHealthMap 用于记录节点健康状态的 Map。

podUpdateItem 用于更新 Pod 的信息。

Controller 是节点生命周期控制器的数据结构，其中包含了所有的配置信息和各种操作函数。

init 初始化节点生命周期控制器。

deepCopy 深拷贝 Controller 结构体。

newNodeHealthMap 生成节点健康状态的 Map。

getDeepCopy 获取 Controller 的深拷贝。

set 设置 Controller 的属性值。

NewNodeLifecycleController 生成一个新的节点生命周期控制器。

Run 调用节点生命周期控制器的主运行函数。

doNodeProcessingPassWorker 节点处理 pass 的 Worker 并对集群的节点进行操作。

doNoScheduleTaintingPass 执行 NoSchedule tainting Pass。

doNoExecuteTaintingPass 执行 NoExecute tainting Pass。

monitorNodeHealth 监控节点健康状态。

processTaintBaseEviction 处理 Taint-Based Eviction。

isNodeExcludedFromDisruptionChecks 判断节点是否要排除在 Disruption Checks 之外。

tryUpdateNodeHealth 更新节点健康状态。

handleDisruption 处理节点不可用情况。

podUpdated 当 Pod 更新时进行处理。

doPodProcessingWorker 处理 Pod 的 Worker。

processPod 处理 Pod。

setLimiterInZone 在指定 Zone 中设置限制器。

classifyNodes 对节点进行分类。

HealthyQPSFunc 获取健康的 QPS 值。

ReducedQPSFunc 获取降级后的 QPS 值。

addPodEvictorForNewZone 为新的 Zone 添加 Pod Evictor。

markNodeForTainting 标记节点为 Tainting。

markNodeAsReachable 将节点标记为 reachable。

ComputeZoneState 计算 Zone 的状态。

reconcileNodeLabels 调节节点标签。

