# File: pkg/scheduler/eventhandlers.go

pkg/scheduler/eventhandlers.go文件是Kubernetes调度器的事件处理程序。它包含一系列用于处理不同类型事件的函数。这些函数对于调度器的正常运行非常重要，它们会监视集群中发生的各种事件，并相应地更新调度器内部的数据结构和状态。以下是其中一些重要函数的详细介绍：

1. AdmissionResult结构体：它代表一次预选或者检查的结果，包含以下字段：
   - Allowed：一个布尔值，表示一个操作（例如调度一个Pod）是否被允许执行。
   - Reason：一个字符串，表示操作不被允许的原因。

2. onStorageClassAdd函数：这个函数用于处理存储类（StorageClass）添加事件。

3. addNodeToCache函数：这个函数用于将新节点添加到调度器的缓存中。

4. updateNodeInCache函数：这个函数用于更新已存在节点的缓存信息。

5. deleteNodeFromCache函数：这个函数用于从调度器的缓存中删除节点。

6. addPodToSchedulingQueue函数：这个函数将Pod添加到调度队列中，以待后续的调度过程。

7. updatePodInSchedulingQueue函数：这个函数更新调度队列中已存在的Pod。

8. deletePodFromSchedulingQueue函数：这个函数从调度队列中删除Pod。

9. addPodToCache函数：这个函数将Pod添加到缓存中，以供调度使用。

10. updatePodInCache函数：这个函数更新缓存中已存在的Pod。

11. deletePodFromCache函数：这个函数从缓存中删除Pod。

12. assignedPod函数：这个函数用于处理已被调度的Pod，包括发送事件、更新调度器内部状态和数据结构等操作。

13. responsibleForPod函数：这个函数用于判断调度器是否应该对某个Pod负责。

14. addAllEventHandlers函数：这个函数用于将所有的事件处理程序添加到调度器的事件处理队列中。

15. nodeSchedulingPropertiesChange函数：这个函数处理节点调度属性的变化事件。

16. nodeAllocatableChanged函数：这个函数处理节点可分配资源的变化事件。

17. nodeLabelsChanged函数：这个函数处理节点标签的变化事件。

18. nodeTaintsChanged函数：这个函数处理节点容忍性标签的变化事件。

19. nodeConditionsChanged函数：这个函数处理节点条件的变化事件。

20. nodeSpecUnschedulableChanged函数：这个函数处理节点未调度标志的变化事件。

21. preCheckForNode函数：这个函数在某个节点上进行调度前的预检查操作。

22. AdmissionCheck函数：这个函数进行调度前的权限检查操作，包括判断是否有足够的权限执行调度操作。

