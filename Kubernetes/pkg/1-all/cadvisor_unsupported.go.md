# File: pkg/kubelet/cadvisor/cadvisor_unsupported.go

pkg/kubelet/cadvisor/cadvisor_unsupported.go这个文件是用于处理不支持的cadvisor功能的。

其中，变量`_, errUnsupported`用来在函数返回时表示不支持的错误信息。

cadvisorUnsupported结构体是一个空结构体，用于表示不支持的cadvisor功能。

下面是对每个函数的作用的详细介绍：

- `New`: 创建一个新的cadvisor对象，并返回一个未支持错误。

- `Start`: 启动cadvisor，并返回一个未支持错误。

- `DockerContainer`: 获取Docker容器的信息，并返回一个未支持错误。

- `ContainerInfo`: 获取容器的信息，并返回一个未支持错误。

- `ContainerInfoV2`: 获取容器v2版本的信息，并返回一个未支持错误。

- `GetRequestedContainersInfo`: 获取请求的容器信息，并返回一个未支持错误。

- `SubcontainerInfo`: 获取子容器的信息，并返回一个未支持错误。

- `MachineInfo`: 获取机器的信息，并返回一个未支持错误。

- `VersionInfo`: 获取版本信息，并返回一个未支持错误。

- `ImagesFsInfo`: 获取镜像的文件系统信息，并返回一个未支持错误。

- `RootFsInfo`: 获取根文件系统的信息，并返回一个未支持错误。

- `WatchEvents`: 监听事件，并返回一个未支持错误。

- `GetDirFsInfo`: 获取目录的文件系统信息，并返回一个未支持错误。

以上这些函数都是返回一个未支持错误，表示所使用的cadvisor版本不支持相关功能。这个文件的作用就是处理这些不支持的情况。

