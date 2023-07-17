# File: pkg/scheduler/internal/queue/scheduling_queue.go

pkg/scheduler/internal/queue/scheduling_queue.go文件是Kubernetes项目中调度器模块的一部分，用于管理调度队列。主要功能是将待调度的Pod加入队列，并按照一定的规则进行调度。

- defaultPriorityQueueOptions：默认的优先级队列选项，用于创建优先级队列。
- _：占位符，用于忽略某个变量。
- pendingPodsSummary：用于跟踪待调度的Pod的摘要信息。

- PreEnqueueCheck：表示一个前置检查的策略函数，用于决定将Pod加入队列前是否需要进行额外的检查。
- SchedulingQueue：调度队列的主要数据结构，包含待调度的Pod队列、优先级队列和一些相关的配置信息。
- PriorityQueue：优先级队列，存储待调度的Pod，并按照一定的优先级进行排序。
- QueueingHintFunction：表示一个队列提示函数，用于根据Pod的属性生成一个队列提示，以帮助调度器选择合适的队列。
- priorityQueueOptions：优先级队列的配置选项，可以设置队列的大小和并发度。
- Option：调度队列的配置选项，可以设置队列的大小、初始和最大延迟时间等。
- QueueingHintMapPerProfile：将队列提示函数按照调度器配置文件进行分组存储。
- QueueingHintMap：将队列提示函数按照节点名称进行映射存储。
- UnschedulablePods：存储未调度成功的Pod的队列。
- nominator：表示一个认可者接口，用于管理Pod的认可者。

- NewSchedulingQueue：创建一个新的调度队列。
- NominatedNodeName：获取Pod的已认可Node节点名称。
- WithClock：使用指定的时钟参数配置调度队列。
- WithPodInitialBackoffDuration：设置Pod的初始回退延迟时间。
- WithPodMaxBackoffDuration：设置Pod的最大回退延迟时间。
- WithPodLister：将Pod列表程序插入调度队列。
- WithPodMaxInUnschedulablePodsDuration：设置Pod在未能调度成功队列中的最大持续时间。
- WithQueueingHintMapPerProfile：设置根据调度器配置文件分组的队列提示函数。
- WithPreEnqueuePluginMap：设置预先检查插件函数。
- WithMetricsRecorder：设置度量记录器。
- WithPluginMetricsSamplePercent：设置插件度量采样百分比。
- newQueuedPodInfoForLookup：为Pod创建一个新的队列信息，用于查找。
- NewPriorityQueue：创建一个新的优先级队列。
- Run：启动调度队列的主循环。
- isPodWorthRequeuing：判断Pod是否值得重新入队列。
- runPreEnqueuePlugins：运行预先检查插件。
- runPreEnqueuePlugin：运行一个预先检查插件。
- addToActiveQ：将Pod添加到活跃队列中。
- Add：将Pod添加到调度队列中。
- Activate：激活调度队列。
- activate：激活调度器，为队列中的Pod重新计算优先级。
- isPodBackingoff：判断Pod是否正在回退中。
- SchedulingCycle：调度周期的主要循环。
- AddUnschedulableIfNotPresent：将Pod添加到未调度成功队列中。
- flushBackoffQCompleted：将完成回退的Pod从队列中移除。
- flushUnschedulablePodsLeftover：清空未调度成功队列中的Pod。
- Pop：从调度队列中弹出下一个Pod。
- isPodUpdated：判断Pod是否有更新。
- Update：更新调度队列中的Pod。
- Delete：删除调度队列中的Pod。
- AssignedPodAdded：通知调度器已分配Pod。
- isPodResourcesResizedDown：判断Pod的资源是否被缩小。
- AssignedPodUpdated：通知调度器已分配的Pod进行了更新。
- moveAllToActiveOrBackoffQueue：将调度队列中的所有Pod移动到活跃队列或回退队列中。
- MoveAllToActiveOrBackoffQueue：将所有等待队列中的Pod移动到活跃队列或回退队列中。
- requeuePodViaQueueingHint：根据队列提示重新将Pod入队。
- movePodsToActiveOrBackoffQueue：将Pod移动到活跃队列或回退队列。
- getUnschedulablePodsWithMatchingAffinityTerm：获取未调度成功队列中满足亲和性条款的Pod。
- PendingPods：获取待调度的Pod数量。
- Close：关闭调度队列。
- DeleteNominatedPodIfExists：如果存在，删除已认可的Pod。
- deleteNominatedPodIfExistsUnlocked：如果存在，删除已认可的Pod（无锁）。
- AddNominatedPod：添加已认可的Pod。
- NominatedPodsForNode：获取指定节点的已认可的Pod。
- podsCompareBackoffCompleted：比较Pod的回退状态和完成状态。
- newQueuedPodInfo：创建一个新的队列信息。
- getBackoffTime：获取Pod的回退时间。
- calculateBackoffDuration：计算Pod的回退持续时间。
- updatePod：更新Pod的状态。
- addOrUpdate：添加或更新Pod的状态。
- delete：删除Pod。
- get：获取Pod的状态。
- clear：清空调度队列。
- newUnschedulablePods：创建一个新的未调度成功Pod的队列。
- addNominatedPodUnlocked：添加已认可的Pod（无锁）。
- UpdateNominatedPod：更新已认可的Pod。
- updateNominatedPodUnlocked：更新已认可的Pod（无锁）。
- NewPodNominator：创建一个新的Pod认可者。
- newPodNominator：创建一个新的Pod认可者。
- MakeNextPodFunc：创建下一个Pod的函数。
- podInfoKeyFunc：用于计算Pod信息的键值的函数。

以上是各个变量和函数的简要介绍，这个文件中的代码负责管理和调度Pod的队列，以及提供相关的操作和功能。

