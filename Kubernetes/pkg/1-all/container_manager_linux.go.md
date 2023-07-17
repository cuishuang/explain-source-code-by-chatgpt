# File: pkg/kubelet/cm/container_manager_linux.go

pkg/kubelet/cm/container_manager_linux.go文件是Kubernetes项目中kubelet的实现之一。它主要负责Linux环境下容器管理器的实现。

_变量在Go语言中用作占位符，表示一个变量不会被使用。通常被使用在函数参数或返回值中。

以下是这些结构体的作用：

- systemContainer: 定义了系统容器的相关信息，包括容器名称、镜像、命令等。
- containerManagerImpl: 实现了容器管理器接口的具体实现。
- features: 定义了容器管理器支持的功能列表。
- KernelTunableBehavior: 定义了内核可调整参数的行为，如何获取和设置内核参数的策略。

以下是这些函数的作用：

- newSystemCgroups: 创建系统容器组的实例，并返回其引用。
- validateSystemRequirements: 检查系统是否满足容器管理器的要求，例如是否支持cgroups等。
- NewContainerManager: 创建一个新的容器管理器实例。
- NewPodContainerManager: 创建一个新的Pod容器管理器实例。
- InternalContainerLifecycle: 定义了容器的生命周期管理接口。
- createManager: 创建容器管理器的实现。
- setupKernelTunables: 设置内核可调整参数。
- setupNode: 在节点上设置容器管理器相关的配置。
- GetNodeConfig: 获取节点的容器管理器配置。
- GetPodCgroupRoot: 获取Pod的cgroup根路径。
- GetMountedSubsystems: 获取当前系统上已挂载的子系统列表。
- GetQOSContainersInfo: 获取容器的QoS信息，如CPU、内存等限制。
- UpdateQOSCgroups: 更新容器的QoS cgroups。
- Status: 返回容器管理器的状态。
- Start: 启动容器管理器。
- GetPluginRegistrationHandler: 返回容器管理器的插件注册处理器。
- GetResources: 获取节点上的资源信息，如CPU、内存等。
- UpdatePluginResources: 更新插件的资源信息。
- GetAllocateResourcesPodAdmitHandler: 返回Pod的资源分配处理器。
- SystemCgroupsLimit: 返回系统容器组的限制配置。
- buildContainerMapFromRuntime: 从runtime构建容器的映射关系。
- isProcessRunningInHost: 检查指定的进程是否运行在宿主机中。
- ensureProcessInContainerWithOOMScore: 确保具有指定OOM分数的进程运行在容器中。
- getContainer: 获取指定容器的信息和状态。
- ensureSystemCgroups: 确保系统容器组存在。
- isKernelPid: 检查指定的pid是否属于内核进程。
- GetCapacity: 获取节点的资源容量。
- GetDevicePluginResourceCapacity: 获取设备插件资源的容量。
- GetDevices: 获取节点上的设备列表。
- GetAllocatableDevices: 获取可分配设备列表。
- int64Slice: 类型定义，用于表示int64切片。
- GetCPUs: 获取节点上可用的CPU列表。
- GetAllocatableCPUs: 获取可分配的CPU列表。
- GetMemory: 获取节点的内存信息。
- GetAllocatableMemory: 获取可分配的内存信息。
- GetDynamicResources: 获取动态资源信息。
- ShouldResetExtendedResourceCapacity: 检查是否应该重置扩展资源容量。
- UpdateAllocatedDevices: 更新已分配的设备列表。
- containerMemoryFromBlock: 根据块设备信息获取容器的内存信息。
- PrepareDynamicResources: 准备动态资源。
- UnprepareDynamicResources: 取消准备动态资源。
- PodMightNeedToUnprepareResources: 检查Pod是否需要取消准备资源。

这些函数都是kubelet在管理容器时使用的一些工具函数，用于初始化容器管理器、获取资源信息、更新容器状态等操作。

