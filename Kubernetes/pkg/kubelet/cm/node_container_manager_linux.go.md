# File: pkg/kubelet/cm/node_container_manager_linux.go

pkg/kubelet/cm/node_container_manager_linux.go文件是Kubernetes项目中kubelet组件的一个源代码文件，它负责管理Linux节点上容器的资源分配和控制。下面详细介绍该文件中的每个函数的作用：

1. createNodeAllocatableCgroups：
   - 作用：在Linux节点上创建或者重置节点上的Allocatable Cgroup。
   - 参数：无。
   - 返回值：无。

2. enforceNodeAllocatableCgroups：
   - 作用：根据配置文件，强制系统和容器为Node Allocatable Cgroups设置限制。
   - 参数：无。
   - 返回值：无。

3. enforceExistingCgroup：
   - 作用：强制现有的Cgroup保持在配置文件指定的限制范围内。
   - 参数：cgroupName和resourceSpec。
   - 返回值：error。

4. getCgroupConfig：
   - 作用：从kubelet的配置中获取节点的Cgroup配置信息。
   - 参数：无。
   - 返回值：CgroupConfig和error。

5. GetNodeAllocatableAbsolute：
   - 作用：获取指定节点的Allocatable资源的绝对值。
   - 参数：nodeName。
   - 返回值：NodeAllocatableAbsolute和error。

6. getNodeAllocatableAbsoluteImpl：
   - 作用：根据配置和运行时数据计算指定节点的Allocatable资源的绝对值。
   - 参数：nodeName、nodeInfo、imageGCHighThresholdPercent和imageGCLowThresholdPercent。
   - 返回值：NodeAllocatableAbsolute和error。

7. getNodeAllocatableInternalAbsolute：
   - 作用：从cgroup读取给定节点资源的内部计算结果。
   - 参数：nodeName。
   - 返回值：NodeAllocatableAbsolute和error。

8. GetNodeAllocatableReservation：
   - 作用：获取配置文件中指定节点的Allocatable资源的保留量。
   - 参数：nodeName。
   - 返回值：NodeAllocatableReservation和error。

9. validateNodeAllocatable：
   - 作用：校验节点的Allocatable资源是否满足配置要求。
   - 参数：nodeName、allocatable、hardLimits和resources。
   - 返回值：error。

这些函数共同负责在Kubernetes集群中管理Linux节点上容器的资源分配和控制，包括创建和重置Cgroup、强制Cgroup保持在限制范围内、获取节点资源的绝对值、计算资源使用率、校验资源是否满足配置要求等功能。

