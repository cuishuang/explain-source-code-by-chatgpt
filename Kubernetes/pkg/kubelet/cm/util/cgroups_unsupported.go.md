# File: pkg/kubelet/cm/util/cgroups_unsupported.go

在 Kubernetes 项目中，pkg/kubelet/cm/util/cgroups_unsupported.go 文件的作用是提供了一些未受支持的功能函数，用于在不支持 Cgroups 的系统上进行 Cgroups 相关操作。

该文件中定义了 `GetPids`, `GetCgroupRoot`, `PidsInCgroup` 等几个函数，它们分别有以下作用：

1. `GetPids`: 该函数用于获取某个 Cgroup 中的所有进程号（PID）。在不支持 Cgroups 的系统上，该函数会返回一个空的进程号列表。
2. `GetCgroupRoot`: 该函数用于获取 Cgroups 的根目录路径。在不支持 Cgroups 的系统上，该函数会返回空字符串。
3. `PidsInCgroup`: 该函数用于获取某个 Cgroup 的所有子进程号（PID）。在不支持 Cgroups 的系统上，该函数会返回一个空的进程号列表。

这些函数的目的是在不支持 Cgroups 的系统上提供一些功能的替代方法，以保持代码的兼容性和可移植性。它们是通过在不支持 Cgroups 的情况下返回无效的或空的数据来达到这个目的的。

