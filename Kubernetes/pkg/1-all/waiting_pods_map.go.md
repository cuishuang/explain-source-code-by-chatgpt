# File: pkg/scheduler/framework/runtime/waiting_pods_map.go

文件pkg/scheduler/framework/runtime/waiting_pods_map.go在Kubernetes项目中的作用是实现了一个等待Pod的映射表。该映射表用于存储等待被调度的Pod，并提供了一些函数来添加、移除、获取和遍历等待Pod。

_ 这几个变量主要用于省略未使用的返回值（blank identifier）。

waitingPodsMap 是一个结构体，用于表示等待Pod的映射表。它内部包含了一个用于存储等待Pod的map，键是等待Pod的UID，值是对应的waitingPod结构体。

waitingPod 结构体用于表示一个等待Pod的信息，包含了Pod的UID、等待的插件列表以及一个条件变量用于在Pod被调度时进行通知。

newWaitingPodsMap 是一个函数，用于创建一个新的等待Pod的映射表实例。

add 函数用于将一个等待Pod添加到映射表中。

remove 函数用于从映射表中移除指定的等待Pod。

get 函数用于获取指定的等待Pod。

iterate 函数用于遍历映射表中的所有等待Pod。

newWaitingPod 是一个函数，用于创建一个新的等待Pod实例。

GetPod 通过Pod的UID从等待Pod的映射表中获取对应的waitingPod结构体。

GetPendingPlugins 用于获取指定等待Pod的等待的插件列表。

Allow 函数用于设置等待Pod的等待完成标志，并将其设置为可调度。

Reject 函数用于设置等待Pod的等待完成标志，并将其设置为不可调度。

