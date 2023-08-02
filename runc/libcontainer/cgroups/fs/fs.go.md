# File: runc/libcontainer/cgroups/fs/fs.go

在runc项目中，runc/libcontainer/cgroups/fs/fs.go文件是libcontainer包中实现cgroup操作的核心文件之一。它定义了与文件系统层次结构相关的操作，其中包括使用不同的cgroup子系统（如cpu、memory、blkio等）进行资源限制、监控和控制的函数。

subsystems和errSubsystemDoesNotExist变量是用于存储可用的cgroup子系统集合和错误信息的字符串。subsystems是一个字符串切片，其中包含了所有可用的cgroup子系统名称，在初始化时被填充。errSubsystemDoesNotExist是一个预定义的错误信息，当cgroup子系统不存在时会被返回。

subsystem结构体是对每个cgroup子系统的抽象，它包含了子系统名称、控制器路径和实现的函数用于对相应的cgroup进行操作。

Manager结构体是对cgroup的管理器的抽象，它包含了对cgroup子系统的实例集合以及一些操作和属性的通用方法。

以下是fs.go文件中一些重要函数的介绍：

- init函数用于初始化cgroup的文件系统，它会在环境启动时被调用。

- NewManager函数返回一个新的cgroup管理器实例，用于管理给定的init进程的cgroup。

- isIgnorableError函数根据错误类型判断是否可以忽略对应的错误。

- Apply函数将进程归类到对应的cgroup中，并应用相应的资源限制。

- Destroy函数销毁指定的cgroup。

- Path函数返回给定cgroup的文件系统路径。

- GetStats函数返回给定cgroup的统计信息。

- Set函数用于设置给定cgroup的特定控制器的属性值。

- Freeze函数用于暂停给定cgroup中的所有进程。

- GetPids函数返回指定cgroup中的所有进程ID。

- GetAllPids函数返回所有cgroup中的进程ID。

- GetPaths函数返回给定cgroup及其子cgroup的路径切片。

- GetCgroups函数返回当前系统上所有的cgroup信息。

- GetFreezerState函数返回给定cgroup的冻结状态。

- Exists函数检查给定cgroup是否存在。

- OOMKillCount函数返回给定cgroup中的OOM（内存不足）杀死的进程数量。

这些函数提供了对cgroup进行操作的接口，用于实现资源限制、监控和控制等功能。

