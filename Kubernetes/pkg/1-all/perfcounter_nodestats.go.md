# File: pkg/kubelet/winstats/perfcounter_nodestats.go

文件pkg/kubelet/winstats/perfcounter_nodestats.go在Kubernetes项目中的作用是提供用于收集节点统计信息的性能计数器实现。

- modkernel32：是Windows系统库kernel32.dll的句柄，用于加载该库并获取其中的函数。
- procGlobalMemoryStatusEx：指向kernel32.dll库中的GlobalMemoryStatusEx函数，在Windows上获取系统内存状态信息。
- procGetActiveProcessorCount：指向kernel32.dll库中的GetActiveProcessorCount函数，用于获取系统活动处理器的数量。

结构体:
- MemoryStatusEx：存储系统内存状态信息，包括总物理内存、可用内存等。
- perfCounterNodeStatsClient：用于收集节点统计信息的性能计数器客户端，包括CPU利用率、内存使用量、系统UUID等指标。

函数:
- NewPerfCounterClient：创建一个新的性能计数器客户端。
- startMonitoring：启动性能计数器监控。
- getMachineInfo：获取机器信息，包括操作系统版本、系统架构等。
- ProcessorCount：获取系统中的逻辑处理器数量。
- getActiveProcessorCount：获取当前活动处理器数量。
- getVersionInfo：获取操作系统版本信息。
- getNodeMetrics：获取节点的指定度量指标，如CPU使用率、内存使用等。
- getNodeInfo：获取节点的一般信息，包括节点名称、操作系统、架构等。
- collectMetricsData：收集节点指标数据。
- convertCPUValue：将CPU利用率从类型为uint64的值转换为nanocores。
- getCPUUsageNanoCores：获取CPU使用率。
- getSystemUUID：获取系统UUID。
- getPhysicallyInstalledSystemMemoryBytes：获取物理安装的系统内存的字节数。
- getBootID：获取系统的引导ID。

这些函数的作用主要是用于收集Windows节点的系统统计信息，包括CPU、内存等指标，用于在Kubernetes集群中监控节点的健康状况和性能。

