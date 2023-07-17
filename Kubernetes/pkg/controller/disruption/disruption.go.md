# File: pkg/controller/disruption/disruption.go

pkg/controller/disruption/disruption.go文件是kubernetes项目中的一个控制器，用于管理Pod的失效和恢复。具体来说，它允许管理员定义要移除的Pod集合，以便进行计划的维护和升级，同时确保这些Pod的缩放等级保持不变。

接下来让我们介绍一下这个文件中的各个变量和结构体的作用：

1. controllerKindRS, controllerKindSS, controllerKindRC, controllerKindDep

这些变量分别代表ReplicaSet、StatefulSet、ReplicationController和Deployment这几种控制器的类型，以便区分不同的控制器。

2. updater

updater是这个控制器的核心逻辑。它用于获取更新Pod集合的请求并执行相关操作。

3. DisruptionController

DisruptionController是一个控制器对象，用于启动更新逻辑和初始化状态。

4. controllerAndScale

controllerAndScale结构体存储控制器的名称和缩放大小，以便后续更新Pod集合。

5. podControllerFinder

podControllerFinder是一个Finder接口，它提供了查找Pod所属控制器的方法。

6. NewDisruptionController, NewDisruptionControllerInternal, finders, getPodReplicaSet, getPodStatefulSet, getPodDeployment, getPodReplicationController, getScaleController, implementsScale, verifyGroupKind

这些方法和结构体用于初始化DisruptionController对象，并提供基本的查找和验证功能。

7. Run, addDb, updateDb, removeDb, addPod, updatePod, deletePod, enqueuePdb, enqueuePdbForRecheck, enqueueStalePodDisruptionCleanup, getPdbForPod, getPodsForPdb, worker, processNextWorkItem, recheckWorker, processNextRecheckWorkItem, stalePodDisruptionWorker, processNextStalePodDisruptionWorkItem, sync, trySync, syncStalePodDisruption, getExpectedPodCount, getExpectedScale, countHealthyPods, buildDisruptedPodMap, failSafe, updatePdbStatus, writePdbStatus, nonTerminatingPodHasStaleDisruptionCondition

这些方法实现了整个控制器的逻辑，包括更新Pod集合、处理任务队列、同步状态等操作。

总的来说，这个控制器是kubernetes项目中一个非常重要的组件，它可以帮助管理员进行计划的维护和升级而不影响现有的工作负载。

