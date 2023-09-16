# File: istio/pkg/kube/kclient/crdwatcher.go

crdwatcher.go文件是在Istio项目中用于监视和管理Kubernetes的自定义资源定义（Custom Resource Definition，CRD）的文件。

该文件中定义了几个结构体用于处理CRD的监视操作：

1. crdWatcher：用于存储一个CRD资源的信息，包括资源的类型、版本和监视回调函数等。
2. crdWatcherSet：用于管理多个crdWatcher实例，提供添加、删除和获取监视器等功能。

以下是这些结构体的详细说明：

- crdWatcher：
  - `init()`：用于初始化crdWatcher的内部状态。
  - `newCrdWatcher()`：根据给定的资源类型和版本创建一个新的crdWatcher实例。
  - `HasSynced()`：检查crdWatcher是否已经与Kubernetes服务器同步完成。
  - `Run()`：启动资源的监视操作，在该函数中会创建一个Kubernetes的watcher来监视CRD资源的变化。
  - `WaitForCRD()`：等待CRD资源的创建和同步完成。
  - `KnownOrCallback()`：根据给定的资源版本和类型检查该crdWatcher是否已知，如果已知则直接返回，否则执行回调函数。
  - `known()`：检查指定的资源版本和类型是否已经在crdWatcher中标记为已知。

这些函数的作用如下：

- `init()`：为crdWatcher对象初始化内部状态，比如设置其资源类型和版本等。
- `newCrdWatcher()`：根据给定的资源类型和版本创建一个新的crdWatcher实例，并返回该实例。
- `HasSynced()`：检查crdWatcher是否已经与Kubernetes服务器同步完成，返回一个布尔值。
- `Run()`：启动资源的监视操作，该函数会创建一个Kubernetes的watcher来监视CRD资源的变化。
- `WaitForCRD()`：等待CRD资源的创建和同步完成，直到所有CRD资源都已经被Kubernetes服务器创建和同步完成。
- `KnownOrCallback()`：根据给定的资源版本和类型检查该crdWatcher是否已知，如果已知则直接返回，否则执行回调函数。
- `known()`：检查指定的资源版本和类型是否已经在crdWatcher中标记为已知，返回一个布尔值。

总之，crdwatcher.go文件中的这些结构体和函数用于帮助Istio项目监视和管理Kubernetes的CRD资源的创建和同步操作。

