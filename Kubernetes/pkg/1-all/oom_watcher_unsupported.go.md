# File: pkg/kubelet/oom/oom_watcher_unsupported.go

在kubernetes项目中，pkg/kubelet/oom/oom_watcher_unsupported.go文件的作用是提供一个在不支持OOM（out-of-memory）监控的平台上进行兼容性处理的备选方案。当平台不支持OOM监控时，该文件中定义的结构体和函数被用来模拟OOM监控的功能。

- `_`变量：在Go语言中，使用下划线_表示一个空标识符，用于忽略变量或导入包。在这个文件中，_变量被用于忽略某些返回值，表示不关心这些值。

- `oomWatcherUnsupported`结构体：该结构体用于模拟OOM监控的功能，当平台无法提供真正的OOM监控时，使用这个结构体进行兼容性处理。它包含了一个stop通道用于控制监控的停止，一个done通道用于通知监控已结束，并具有一些空的方法用于实现对应的接口。

- `NewWatcher`函数：该函数用于创建一个新的OOM监控器实例。在不支持OOM监控的平台上，它会返回一个`oomWatcherUnsupported`实例。

- `Start`函数：该函数用于启动OOM监控器。对于不支持OOM监控的平台，该函数仅返回nil。

通过使用`oomWatcherUnsupported`结构体和相关函数，kubernetes项目能够在不支持OOM监控的平台上提供一种兼容性处理的解决方案。但需要注意的是，这只是一个模拟方案，并不能真正提供OOM监控的功能。

