# File: pkg/kubelet/pleg/evented.go

文件pkg/kubelet/pleg/evented.go是Kubernetes项目中kubelet中PLEG（Pod Lifecycle Event Generator）的实现。PLEG用于监视和处理容器和Pod的生命周期事件。

eventedPLEGUsage和eventedPLEGUsageMu是用于标记evented PLEG是否正在使用的变量。eventedPLEGUsage是一个计数器，用于记录使用evented PLEG的数量。eventedPLEGUsageMu是一个互斥锁，用于保护eventedPLEGUsage变量的并发访问。

EventedPLEG是一个结构体，用于将evented PLEG的状态和操作封装在一起。它包含了一个KubeletClient用于与kubelet API进行交互，一个已经监视的Pod状态的缓存，一个事件队列，以及一些用于记录metric和处理event的辅助函数。

- isEventedPLEGInUse用于检查evented PLEG是否在使用中。
- setEventedPLEGUsage用于设置evented PLEG的使用状态。
- NewEventedPLEG用于创建一个新的evented PLEG对象。
- Watch用于开始监听Pod和容器事件。
- Relist用于重新列出已知的Pod和容器，更新缓存。
- Start用于启动evented PLEG的处理循环。
- Stop用于停止evented PLEG的处理循环。
- updateGlobalCache用于更新缓存中的全局Pod状态。
- Update用于更新Pod状态。
- Healthy用于检查evented PLEG的健康状态。
- watchEventsChannel是evented PLEG的事件处理循环，监听事件队列并处理事件。
- processCRIEvents用于处理CRI（Container Runtime Interface）事件。
- processCRIEvent用于处理单个CRI事件。
- getPodIPs用于获取Pod的IP地址。
- sendPodLifecycleEvent用于发送Pod的生命周期事件到KubeletClient。
- getPodSandboxState用于获取Pod Sandbox的状态。
- updateRunningPodMetric用于更新运行中的Pod的metric。
- getContainerStateCount用于获取容器状态的数量。
- updateRunningContainerMetric用于更新运行中的容器的metric。
- updateLatencyMetric用于更新metric的延迟。
- UpdateCache用于更新缓存中的Pod和容器状态。

