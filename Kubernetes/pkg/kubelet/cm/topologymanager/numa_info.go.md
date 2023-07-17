# File: pkg/kubelet/cm/topologymanager/numa_info.go

`numa_info.go`文件是Kubernetes项目中kubelet的代码文件，其作用是管理与NUMA（Non-Uniform Memory Access，非一致性内存访问）相关的拓扑信息。

1. `NUMADistances`结构体：用于表示NUMA节点之间的距离信息。它包含一个矩阵，表示各个NUMA节点之间相对距离的程度。

2. `NUMAInfo`结构体：用于表示NUMA节点的拓扑信息。它包含了节点的ID、总内存和CPU信息，以及节点之间的距离信息。

3. `NewNUMAInfo`函数：用于创建一个NUMAInfo结构体，并根据传入的节点信息和距离信息进行初始化。

4. `Narrowest`函数：用于获取所有NUMA节点中最窄的节点。最窄的节点是指具有最小可用内存的节点。

5. `Closest`函数：用于获取给定节点与其他节点的距离，根据NUMADistances中的信息计算。

6. `DefaultAffinityMask`函数：用于计算默认的NUMA亲和性掩码，即选择NUMA节点的首选顺序。

7. `CalculateAverageFor`函数：根据NUMAInfo和Pod的资源需求，计算可以用于Pod的平均NUMA节点。

这些函数主要用于NUMA拓扑管理和调度。它们根据节点的拓扑结构和距离信息，提供了一些有助于任务分配和负载均衡的功能。

