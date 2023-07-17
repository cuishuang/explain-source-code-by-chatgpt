# File: pkg/controller/garbagecollector/graph_builder.go

该文件的作用是构建 Kubernetes 对象之间的关系图，用于垃圾回收控制器（Garbage Collector Controller）从集群中删除不再需要的对象。通过 GraphBuilder 结构体对多个 monitor（监视器）进行协调并构建对象关系图，从而确保在删除一个对象时不会同时删除其他相关对象。

ignoredResources 是一个包含需要忽略的资源类型的切片，例如 ConfigMap 和 Secret，因为这些资源类型已经被 Kubernetes 的 API server 自动管理。

eventType 表示关系图中节点的事件类型，例如添加、删除或更新。

event 是一个带有对象元数据（metadata）和相关的对象事件数据（event data）的结构体。

GraphBuilder 是构建垃圾回收控制器的对象关系图的结构体，其中包括 monitors 列表，用于监视不同类型的 Kubernetes 对象之间的依赖关系。

monitor 是一个表示监视器的结构体，其中包括节点、监视器名称以及启动和停止监视器的函数。

ownerRefPair 是一个表示 Kubernetes 对象所有权对的结构体，其中包括名称和命名空间，并将每个对象与其所有者的元数据相对应。

String 函数将 GraphBuilder 对象转换为字符串。

Run 函数启动 Garbage Collector 控制器，该函数启动所有监视器并监视对象关系图的更改。

controllerFor 函数为指定的资源类型创建一个新的 Garbage Collector 控制器，并返回该控制器的指针。

syncMonitors 函数将每个监视器的状态与 Kubernetes 对象的当前状态同步。

startMonitors 函数启动每个监视器以开始监视依赖对象的更改。

IsSynced 函数检查是否已成功加载 Kubernetes 对象。

DefaultIgnoredResources 函数返回一个切片，其中包含默认情况下应该忽略的资源类型的列表。

enqueueVirtualDeleteEvent 函数将虚拟删除事件作为指定节点的事件转入 Event Queue。

addDependentToOwners 函数将依赖对象的所有权加入到所有者列表中。

reportInvalidNamespaceOwnerRef 函数报告命名空间不符合预期的 ownerRef 对象。

insertNode 函数将指定的节点插入到对象关系图中。

removeDependentFromOwners 函数从所有者列表中删除指定的依赖对象。

removeNode 函数从对象关系图中删除指定的节点。

referencesDiffs 函数比较两个 reference 列表，并返回资源类型的字符串切片，其中包含所有被删除或已添加到列表中的 resourceName。

deletionStartsWithFinalizer 函数检查对象是否需要删除依赖对象。

beingDeleted 是一个表示正在被删除的 Kubernetes 对象的 map，其中 key 是 Kubernetes 对象键，value 是一个布尔值，表示是否正在删除该对象。

hasDeleteDependentsFinalizer 函数检查对象是否具有删除依赖对象的 finalizer。

hasOrphanFinalizer 函数检查对象是否具有禁止垃圾回收或phans的 finalizer。

hasFinalizer 函数检查对象是否已设置 finalizer。

startsWaitingForDependentsDeleted 函数将正在等待依赖对象删除的对象重新标记为等待删除。

startsWaitingForDependentsOrphaned 函数将正在等待孤立的对象重新标记为等待删除。

addUnblockedOwnersToDeleteQueue 函数将可以被删除的所有者加入到删除事件队列中。

processTransitions 函数处理从一个节点到另一个节点的所有过渡并对所有的依赖者和所有者进行垃圾回收。

runProcessGraphChanges 函数运行对象关系图改变的过程，并根据需要对依赖对象和所有者进行垃圾回收。

identityFromEvent 函数返回指定事件的同一标识符。

processGraphChanges 函数处理从一个节点到另一个节点的所有过渡并对所有的依赖者和所有者进行垃圾回收。

partitionDependents 函数根据所依赖的资源以依赖关系将依赖者分为不同的分区。

referenceLessThan 函数返回两个引用的比较结果。

getAlternateOwnerIdentity 函数从给定的对象元数据列表中返回所有候选所有者，并在必要时使用给定的默认值。

