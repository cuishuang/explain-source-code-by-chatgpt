# File: pkg/kubelet/cm/container_manager.go

文件`pkg/kubelet/cm/container_manager.go`在Kubernetes项目中的作用是实现容器管理器的功能。该文件定义了`ContainerManager`结构体，该结构体实现了`CM`接口，并用于处理容器的创建、删除、查询等操作。

下面对提到的几个结构体和函数进行介绍：

1. `ActivePodsFunc`：这是一个函数类型的别名，用于定义获取活跃Pod的函数签名。在容器管理器中，可以使用这个函数获取当前节点上活跃的Pod列表。

2. `ContainerManager`：这是一个结构体，实现了`CM`接口。它包含了容器管理器的各种方法和属性，用于管理容器的生命周期、资源分配等。

3. `NodeConfig`：这个结构体定义了节点的相关配置信息，包括节点名称、节点IP地址等。

4. `NodeAllocatableConfig`：这个结构体定义了节点的可分配资源配置，包括CPU、内存、存储等。

5. `Status`：这个结构体定义了容器管理器的状态信息，包括节点名称、容器运行状态等。

下面对提到的几个函数进行介绍：

1. `parsePercentage`：这个函数用于解析字符串表示的百分比，并返回对应的浮点数。例如，可以使用该函数解析配置文件中的资源分配百分比。

2. `ParseQOSReserved`：这个函数用于解析字符串表示的预留资源量，并返回对应的整数值。例如，可以使用该函数解析配置文件中的QoS预留资源。

3. `containerDevicesFromResourceDeviceInstances`：这个函数用于获取容器实例的设备信息。它将资源设备实例中的设备路径、容器路径等信息进行转换，返回容器设备的相关配置信息。

总体而言，`container_manager.go`文件定义了容器管理器的相关结构体和函数，用于实现对容器的管理和操作。具体功能包括获取活跃的Pod列表、处理容器的创建和删除、解析资源配置信息等。

