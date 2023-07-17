# File: pkg/kubelet/qos/policy.go

在Kubernetes项目中，`pkg/kubelet/qos/policy.go`文件的作用是定义了Kubelet的QoS（Quality of Service）策略。

QoS是一种对容器进行分类和分配资源的方法，以确保不同优先级的容器在资源紧张的情况下能够得到适当的保护和分配。Kubernetes中，每个容器都被分为三个不同的QoS类别：Guaranteed、Burstable和BestEffort。这些类别根据容器对资源的需求和保证进行了不同的划分。

`policy.go`文件中的`GetContainerOOMScoreAdjust`这几个函数是用来获取容器OOM（Out of Memory）得分调整值的。OOM得分调整值是一个整数，它代表了容器相对于其他容器的内存资源使用情况。

具体来说，这几个函数的作用如下：

1. `GetContainerOOMScoreAdjust(pod *v1.Pod, container *v1.Container, nodeName string) (int, error)`：该函数负责根据Pod和容器的信息获取容器的OOM得分调整值。它在创建或更新Pod时被调用。该函数首先检查Pod的QoS类别，然后根据QoS类别和容器的限制条件计算OOM得分调整值。

2. `GetPodOOMScoreAdjust(pod *v1.Pod, staticPod bool) int`：该函数用于获取整个Pod的OOM得分调整值。它在创建或更新Pod时被调用。如果Pod是静态Pod，则会返回一个固定的得分调整值；否则，它会依据Pod中容器的OOM得分调整值来计算Pod的总得分调整值。

3. `GetQOSReservedOOMScoreAdjust(podQOS v1.PodQOSClass) int`：该函数根据Pod的QoS类别返回一个预留的OOM得分调整值。预留的OOM得分调整值用于给高优先级的Pod或容器分配额外的资源保护。

4. `GetOOMScoreAdj(pod *v1.Pod, container *v1.Container) (int, error)`：该函数用于获取容器的OOM Score Adj，它根据容器的请求和限制内存大小、容器的QoS类别等信息来计算得分。如果容器有明确的OOM得分调整值（通过`resources.Limits["memory"]`设置），则直接返回该调整值；否则，根据容器的QoS类别和资源设置来计算得分。

这些函数的作用是为了根据容器的资源使用情况和QoS类别来计算和获取适当的OOM得分调整值，以便Kubelet能够根据这些调整值来对容器的内存资源进行合理的保护和分配。

