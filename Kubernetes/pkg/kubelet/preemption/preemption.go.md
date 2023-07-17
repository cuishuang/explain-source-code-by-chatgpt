# File: pkg/kubelet/preemption/preemption.go

在 Kubernetes 项目中，pkg/kubelet/preemption/preemption.go 文件的作用是实现 Kubernetes 节点上的 Pod 抢占逻辑。Pod 抢占是指在节点资源紧张的情况下，根据一定的策略和优先级，终止一些低优先级的 Pod，以保证高优先级 Pod 的资源需求。该文件中定义的函数和结构体实现了 Pod 抢占的逻辑。

下面解释一下几个关键变量和结构体的作用：

- `_` 变量：在 Go 语言中 `_` 是一个特殊的变量名，用于丢弃函数返回的不需要使用的值。在这里，`_` 变量主要用于忽略函数返回的错误信息。

- `CriticalPodAdmissionHandler` 结构体：定义了一个关键 Pod 入场处理程序，用于判断 Pod 是否为关键 Pod，可以允许在资源紧张的情况下被抢占。

- `admissionRequirement` 结构体：定义了一个 Pod 的入场要求，包括 CPU、内存等资源的请求和限制。

- `admissionRequirementList` 结构体：定义了入场要求的列表，每个入场要求对应一个 Pod。

下面是一些重要函数的介绍：

- `NewCriticalPodAdmissionHandler` 函数：创建一个新的关键 Pod 入场处理程序，并返回该实例。该函数内部实现了关键 Pod 入场要求的初始化。

- `HandleAdmissionFailure` 函数：用于处理 Pod 入场失败的情况。它会根据具体的错误信息和入场要求，执行不同的处理逻辑，如重试、停止重试等。

- `evictPodsToFreeRequests` 函数：根据节点上的资源使用状态和 Pod 的入场要求，计算需要被抢占的 Pod。

- `getPodsToPreempt` 函数：根据一定的优先级规则，从需要被抢占的 Pod 列表中选择一批 Pod 进行抢占。

- `getPodsToPreemptByDistance` 函数：根据一定的距离规则，从需要被抢占的 Pod 列表中选择一批 Pod 进行抢占。

- `distance` 函数：计算两个 Pod 之间的距离。距离的计算方式可以根据实际需求进行调整。

- `subtract` 函数：用于计算两个资源请求之间的差值。

- `toString` 函数：将一个资源请求对象转换为字符串表示。

- `sortPodsByQOS` 函数：根据 Pod 的资源需求和限制，对 Pod 进行按 QoS 等级（如 Guaranteed、BestEffort）排序。

- `smallerResourceRequest` 函数：用于比较两个资源请求的大小。

通过这些函数和结构体的组合和调用，可以实现 Pod 的抢占逻辑，保证节点资源的合理分配。

