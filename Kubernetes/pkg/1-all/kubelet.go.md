# File: pkg/kubelet/kubelet.go

pkg/kubelet/kubelet.go文件是 Kubernetes kubelet 组件的主要入口文件，它定义了 Kubelet 结构体和一些供 Kubelet 使用的函数。

Kubelet 结构体是 kubelet 组件的核心结构，它包含了 kubelet 组件的运行时状态和相关的配置信息，可以用于控制 kubelet 的行为。其中的几个重要字段包括：
- ContainerLogsDir: 容器日志目录的路径，用于存储容器的日志文件。
- etcHostsPath: etc/hosts 文件的路径，用于保存 Pod 的 DNS 记录。

SyncHandler, Option, Bootstrap, Dependencies, serviceLister, Kubelet 几个结构体分别用于实现 kubelet 的不同功能：
- SyncHandler: 用于同步 Pod 的处理逻辑。
- Option: 包含了 kubelet 的配置选项。
- Bootstrap: 用于初始化 kubelet 的运行环境。
- Dependencies: 维护了 kubelet 的依赖模块，如容器运行时、网络插件等。
- serviceLister: 用于监听和获取 service 的相关信息。
- Kubelet: Kubelet 实例，用于执行 kubelet 的主要逻辑。

getContainerEtcHostsPath: 获取容器的 etc/hosts 文件路径。

makePodSourceConfig: 创建 Pod 的源配置。

PreInitRuntimeService: 预初始化容器运行时服务。

NewMainKubelet: 创建 Kubelet 的主实例。

ListPodStats, ListPodCPUAndMemoryStats, ListPodStatsAndUpdateCPUNanoCoreUsage: 获取 Pod 的运行时状态和资源使用情况。

ImageFsStats, GetCgroupStats, GetCgroupCPUAndMemoryStats, RootFsStats, GetContainerInfo, GetRawContainerInfo, RlimitStats: 获取容器相关的统计信息。

setupDataDirs: 设置数据目录。

StartGarbageCollection: 开始垃圾回收。

initializeModules: 初始化模块。

initializeRuntimeDependentModules: 初始化与容器运行时相关的模块。

Run: 运行 kubelet 组件的主循环。

SyncPod, SyncTerminatingPod, SyncTerminatingRuntimePod, SyncTerminatedPod, getPodsToSync, deletePod, rejectPod, canAdmitPod, canRunPod, syncLoop, syncLoopIteration, handleProbeSync, HandlePodAdditions, HandlePodUpdates, HandlePodRemoves, HandlePodReconcile, HandlePodSyncs, isPodResizeInProgress, canResizePod, handlePodResourcesResize: 这些函数是 kubelet 主循环过程中用于处理 Pod 同步的不同阶段和各种操作的核心逻辑。

LatestLoopEntryTime: 获取最近一次进入 kubelet 主循环的时间。

updateRuntimeUp: 更新当前容器运行时是否可用的状态。

GetConfiguration: 获取 kubelet 的配置。

BirthCry: kubelet 启动时的初始化逻辑。

ResyncInterval: 控制 kubelet 的 resync 间隔。

ListenAndServe, ListenAndServeReadOnly, ListenAndServePodResources: 启动 kubelet 的 HTTP 服务。

cleanUpContainersInPod: 清理 Pod 中的容器。

fastStatusUpdateOnce: 快速更新 Pod 状态。

CheckpointContainer: 容器checkpoint的管理。

ListMetricDescriptors: 列出度量指标的描述信息。

ListPodSandboxMetrics: 列出 Pod 的沙箱度量指标。

supportLocalStorageCapacityIsolation: 支持本地存储容量隔离。

isSyncPodWorthy: 判断是否值得同步 Pod。

PrepareDynamicResources, UnprepareDynamicResources: 准备和清理动态资源。


