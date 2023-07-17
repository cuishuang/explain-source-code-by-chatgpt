# File: pkg/kubelet/pleg/generic.go

pkg/kubelet/pleg/generic.go文件是Kubernetes项目中kubelet的Pod Lifecycle Event Generator（PLEG）的实现。PLEG负责监视容器运行时的变化，并生成相应的事件。

- GenericPLEG：包含PLEG的实例，它维护了一系列Pod的运行时状态，并根据容器的状态更新这些状态。
- plegContainerState：表示Pod中容器的运行时状态，包括容器的ID、容器的状态和容器所属的Pod ID。
- podRecord：表示一个Pod的运行记录，包含Pod ID、Pod标签和运行时状态。
- podRecords：存储了一组PodRecord。

下面是这些方法的功能描述：

- convertState：将容器的运行状态转换为PLEG容器状态，如running、terminated等。
- NewGenericPLEG：创建一个新的GenericPLEG实例。
- Watch：监听容器事件和状态变化。
- Start：启动PLEG的监控。
- Stop：停止PLEG的监控。
- Update：更新PLEG的状态。
- Healthy：检查PLEG是否健康。
- generateEvents：生成容器事件。
- getRelistTime：获取下一次重新列举Pod的时间间隔。
- updateRelistTime：更新重新列举Pod的时间间隔。
- Relist：重新列举所有的Pod。
- getContainersFromPods：从Pods中获取所有容器。
- computeEvents：计算容器状态变化所产生的事件。
- cacheEnabled：检查是否启用缓存。
- getPodIPs：获取Pod的IP地址。
- updateCache：更新缓存。
- UpdateCache：更新缓存中的Pod状态。
- updateEvents：根据容器状态变化，更新事件。
- getContainerState：获取容器的运行状态。
- updateRunningPodAndContainerMetrics：更新Pod和容器的度量指标。
- getOld：获取之前的Pod运行时的状态。
- getCurrent：获取当前的Pod运行时的状态。
- setCurrent：设置当前的Pod运行时的状态。
- update：根据容器状态变化，更新PLEG中的状态。
- updateInternal：根据Pod的变化更新PLEG的内部状态。

