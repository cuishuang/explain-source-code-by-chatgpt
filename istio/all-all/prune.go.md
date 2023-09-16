# File: istio/operator/pkg/helmreconciler/prune.go

在istio/operator/pkg/helmreconciler/prune.go文件中，主要实现了通过Helm Chart进行Istio部署的资源清理功能。

ClusterResources、ClusterCPResources 和 AllClusterResources是用来存储集群级资源（cluster-level resources）的变量。它们分别表示所有资源、仅控制面相关资源和所有资源（包括控制面和数据面）的列表。这些变量被用于删除集群中多余的资源。

NamespacedResources 是用来存储命名空间级资源（namespace-level resources）的变量，它表示所有的命名空间级资源的列表。它同样被用于删除多余的资源。

Prune 是一个方法，它负责执行资源清理操作。它会删除不再需要的资源，以确保Istio的部署状态与Helm Chart中定义的期望状态一致。

PruneControlPlaneByRevisionWithController 是一个方法，它会根据指定的控制面版本和控制器（controller）名称在数据面和控制面中删除多余的资源。

pilotExists 是一个方法，用于检查集群中是否存在Pilot服务。

DeleteObjectsList 是一个方法，用于删除指定的资源对象列表。

GetPrunedResources 是一个方法，用于获取需要被删除的资源列表。

getIstioOperatorCR 是一个方法，用于获取当前运行的Operator的自定义资源的对象。

DeleteControlPlaneByManifests 是一个方法，用于根据指定的清理策略和对象清单列表来删除控制面。

runForAllTypes 是一个方法，负责遍历所有资源类型进行指定操作。

deleteResources 是一个方法，负责根据清理策略删除资源。

deleteResource 是一个方法，负责删除指定的资源。

removeFromObjectCache 是一个方法，用于从对象缓存中删除指定的资源对象。

