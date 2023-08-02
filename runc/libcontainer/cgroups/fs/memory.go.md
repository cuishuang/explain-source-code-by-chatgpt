# File: runc/libcontainer/cgroups/fs2/memory.go

在runc项目中，runc/libcontainer/cgroups/fs2/memory.go文件的作用是实现对内存（memory）cgroup的操作。该文件定义了一系列函数用于设置、获取和统计内存相关的信息。

- `numToStr(num uint64) string`函数用于将数值转换为字符串。
- `isMemorySet(cgroupPath string) bool`函数检查给定cgroup路径下的内存是否已设置。
- `setMemory(cgroupPath string, name string, val uint64)`函数用于设置指定cgroup路径下的内存相关参数的值。
- `statMemory(path string, statFile string) (MemoryData, error)`函数用于从给定路径下的内存统计文件中获取内存数据。
- `getMemoryDataV2(cgroupPath string) (*cgroupdata.CgroupData, error)`函数从给定cgroup路径中获取内存数据。
- `statsFromMeminfo(r io.Reader, stats *cgroups.MemoryStats) error`函数从读取器中读取meminfo文件并更新给定的内存统计数据。

这些函数的集合提供了对内存cgroup的管理和查询功能。在容器管理中，内存cgroup用于限制和监控容器的内存使用情况。这些函数可以帮助runc项目管理和控制容器的内存资源。

