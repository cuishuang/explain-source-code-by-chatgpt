# File: runc/libcontainer/intelrdt/mbm.go

在runc项目中，runc/libcontainer/intelrdt/mbm.go文件的作用是实现Intel Resource Director Technology（RDT）中的Memory Bandwidth Monitoring（MBM）功能。RDT是一种硬件特性，它允许操作系统监视和控制任务对共享处理器资源的使用。

该文件中的mbmEnabled变量用于标识MBM功能是否启用。mbmNumCounters和mbmExactCnt表示MBM的参数配置。

- mbmEnabled：标志变量，指示MBM功能是否启用。当RDT硬件支持MBM且文件系统支持MBM特性时，mbmEnabled将设置为true，否则设置为false。
- mbmNumCounters：表示每个逻辑号码（logical processor）上可用的MBM计数器数量。
- mbmExactCnt：表示每个计数器中每个bucket大致能容纳的最大计数数目。它是衡量内存带宽的计量单位。

下面是一些相关的函数介绍：

- IsMBMEnabled函数：检查MBM功能是否已启用。如果mbmEnabled标志变量为true，则表示MBM功能已启用，返回值为true；否则返回false。

- getMBMNumaNodeStats函数：用于获取MBM的NUMA节点统计数据。NUMA（Non-Uniform Memory Access）是一种内存访问模式，它允许计算机访问分布在多个物理内存节点上的内存。该函数通过指定的节点ID，从sysfs中读取和解析MBM统计数据，返回MBM节点统计信息。

这些函数和变量用于实现CPU的资源监视和控制，特别是内存带宽的监视。通过使用这些函数和变量，runc可以利用RDT硬件特性中的MBM功能来监视和控制容器的内存带宽使用情况。这对于确保容器之间的资源隔离和性能调优非常重要。

