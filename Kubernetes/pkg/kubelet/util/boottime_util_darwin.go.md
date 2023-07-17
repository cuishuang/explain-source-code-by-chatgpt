# File: pkg/kubelet/util/boottime_util_darwin.go

在Kubernetes项目中，`pkg/kubelet/util/boottime_util_darwin.go`文件的作用是提供获取主机启动时间的功能。

该文件中的`GetBootTimeSeconds`函数使用了`hostBootTime`全局变量保存主机的启动时间，并在第一次调用时通过调用`C.sysctl`函数获取主机的启动时间，并将其保存到`hostBootTime`中。之后，每次调用`GetBootTimeSeconds`函数时，将返回保存在`hostBootTime`中的启动时间。

`GetBootTime`函数调用`GetBootTimeSeconds`函数，并将启动时间转换为`time.Time`类型并返回。

因此，`pkg/kubelet/util/boottime_util_darwin.go`文件中的这些函数主要用于获取主机的启动时间。

