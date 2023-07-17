# File: pkg/volume/fc/fc.go

在Kubernetes项目中，pkg/volume/fc/fc.go文件的作用是定义了与Fibre Channel (FC) 存储卷相关的逻辑和功能。

_ (下划线) 这几个变量通常用于忽略某个值，表示这个值并不会被使用或者不重要。

- fcPlugin结构体是FC卷的插件接口，定义了FC卷相关的方法。
- fcDisk结构体表示FC卷，在节点上表示一个FC卷设备。
- fcDiskMounter结构体表示用于挂载FC卷的方法和属性。
- fcDiskUnmounter结构体表示用于卸载FC卷的方法和属性。
- fcDiskMapper结构体表示用于映射FC卷的方法和属性。
- fcDiskUnmapper结构体表示用于取消映射FC卷的方法和属性。

以下是这些函数的作用：

- ProbeVolumePlugins用于检测并返回支持FC卷的插件列表。
- Init用于初始化FC插件。
- GetPluginName返回FC插件的名称。
- GetVolumeName根据FC卷的规范返回卷的名称。
- CanSupport判断指定的卷是否为FC卷。
- RequiresRemount判断指定的卷是否需要重新挂载。
- SupportsMountOption判断指定的卷是否支持指定的挂载选项。
- SupportsBulkVolumeVerification判断指定的卷是否支持批量卷验证。
- SupportsSELinuxContextMount判断指定的卷是否支持SELinux上下文挂载。
- GetAccessModes返回FC卷支持的访问模式列表。
- NewMounter创建FC卷的挂载器。
- newMounterInternal根据指定的参数创建FC卷的挂载器。
- NewBlockVolumeMapper创建FC卷的块卷映射器。
- newBlockVolumeMapperInternal根据指定的参数创建FC卷的块卷映射器。
- NewUnmounter创建FC卷的卸载器。
- newUnmounterInternal根据指定的参数创建FC卷的卸载器。
- NewBlockVolumeUnmapper创建FC卷的块卷取消映射器。
- newUnmapperInternal根据指定的参数创建FC卷的取消映射器。
- ConstructVolumeSpec根据给定的参数构建一个表示FC卷的VolumeSpec对象。
- ConstructBlockVolumeSpec根据给定的参数构建一个表示FC块卷的VolumeSpec对象。
- GetPath返回FC卷的挂载路径。
- fcGlobalMapPath返回FC卷全局设备映射文件的路径。
- fcPodDeviceMapPath返回FC卷Pod设备映射文件的路径。
- GetAttributes返回FC卷的属性。
- SetUp在节点上设置FC卷，包括映射设备、挂载和设置权限。
- SetUpAt在指定路径上设置FC卷。
- TearDown清理FC卷相关的设备和挂载。
- TearDownAt在指定路径上清理FC卷相关的设备和挂载。
- TearDownDevice在节点上清理FC卷设备。
- UnmapPodDevice取消FC卷在Pod设备映射文件中的映射。
- GetGlobalMapPath返回FC卷全局设备映射文件的路径。
- GetPodDeviceMapPath返回FC卷Pod设备映射文件的路径。
- getVolumeSource返回FC卷的VolumeSource对象。
- createPersistentVolumeFromFCVolumeSource根据FC卷的VolumeSource对象创建一个持久卷。
- getWwnsLunWwids返回FC卷的WWN和LUN WWID属性。

这些函数和结构体定义了在Kubernetes中使用FC卷时的相关功能和方法，用于管理和操作FC卷的挂载、卸载、映射等操作。

