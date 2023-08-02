# File: runc/libcontainer/cgroups/v1_utils.go

在runc项目中，runc/libcontainer/cgroups/v1_utils.go文件的作用是提供一些辅助函数用于操作Linux cgroups的v1版本。

以下是每个变量的作用说明：

- errUnified：表示统一层次控制组的错误。当操作系统不支持统一层次控制组时，将返回该错误。
- ErrV1NoUnified：表示不支持v1版本的错误。当操作系统不支持v1版本的cgroups时，将返回该错误。
- readMountinfoOnce：用于保证readMountinfo函数只被调用一次的标志。
- readMountinfoErr：readMountinfo函数执行时的错误返回结果。
- cgroupMountinfo：表示cgroup的挂载信息。它是一个slice，每个元素代表一个挂载点。

以下是每个结构体的作用说明：

- NotFoundError：表示未找到错误。当无法找到指定的cgroup路径时，将返回该错误。

以下是每个函数的作用说明：

- Error：自定义错误类型，用于包装其他错误。
- NewNotFoundError：创建一个NotFoundError实例。
- IsNotFound：检查给定的错误是否为NotFoundError类型。
- tryDefaultPath：尝试获取默认的cgroup路径。
- readCgroupMountinfo：从/proc/self/mountinfo文件中读取cgroup的挂载信息。
- FindCgroupMountpoint：查找指定cgroup名称的挂载点。
- FindCgroupMountpointAndRoot：查找指定cgroup名称的挂载点和根路径。
- findCgroupMountpointAndRootFromMI：从指定的mountinfo中查找cgroup的挂载点和根路径。
- GetOwnCgroup：获取当前进程所属的cgroup路径。
- getCgroupMountsHelper：从指定的mountinfo中获取cgroup的挂载信息。
- getCgroupMountsV1：获取cgroup的v1版本的挂载信息。
- GetOwnCgroupPath：获取当前进程所属的cgroup路径。
- getCgroupPathHelper：根据给定的controller和cgroup名称获取cgroup路径。
- getControllerPath：根据给定的controller名称获取cgroup控制器的路径。

这些函数和变量提供了操作cgroups的基本功能，例如找到cgroup的挂载点、获取当前进程所属的cgroup路径等。

