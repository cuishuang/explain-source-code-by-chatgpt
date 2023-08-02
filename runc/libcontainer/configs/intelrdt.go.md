# File: runc/libcontainer/intelrdt/intelrdt.go

在runc项目中，runc/libcontainer/intelrdt/intelrdt.go文件的作用是实现了与Intel Resource Director Technology (RDT)相关的功能。该文件包含了一些函数和结构体，用于管理和监控RDT信息。

下面是对每个提到的变量、结构体和函数的详细介绍：

变量：
1. catEnabled：表示Cache Allocation Technology (CAT)是否启用。
2. mbaEnabled：表示Memory Bandwidth Allocation (MBA)是否启用。
3. initOnce：确保只初始化一次的sync.Once实例。
4. errNotFound：表示找不到相关资源时返回的错误。
5. intelRdtRoot：Intel RDT的根路径。
6. intelRdtRootErr：获取Intel RDT根路径时可能出现的错误。
7. rootOnce：确保只获取一次Intel RDT根路径的sync.Once实例。

结构体：
1. Manager：用于管理和配置RDT的结构体，提供对资源分配和监控的功能。
2. newManager：用于创建新的RDT管理器的调用构造函数。

函数：
1. featuresInit：初始化Intel RDT相关的特性。
2. findIntelRdtMountpointDir：查找并返回Intel RDT的挂载点目录。
3. Root：获取并返回Intel RDT的根路径。
4. getIntelRdtParamUint：从指定的文件中读取并返回无符号整数型参数。
5. getIntelRdtParamString：从指定的文件中读取并返回字符串型参数。
6. writeFile：将指定的数据写入到指定文件中。
7. getL3CacheInfo：获取并返回L3缓存信息。
8. getMemBwInfo：获取并返回内存带宽信息。
9. getLastCmdStatus：获取最后一次执行命令的状态信息。
10. WriteIntelRdtTasks：将任务的PID写入到指定的Intel RDT文件中。
11. IsCATEnabled：检查CAT是否启用。
12. IsMBAEnabled：检查MBA是否启用。
13. getIntelRdtPath：获取与指定资源相关的Intel RDT路径。
14. Apply：应用指定的资源配置到指定的任务上。
15. Destroy：销毁指定任务的RDT资源配置。
16. GetPath：获取与指定任务相关的资源路径。
17. GetStats：获取指定任务的资源统计信息。
18. Set：设置指定任务的资源配置。
19. newLastCmdError：创建一个存储命令执行错误的结构体的实例。

