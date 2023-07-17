# File: cmd/kubeadm/app/util/staticpod/utils.go

文件`cmd/kubeadm/app/util/staticpod/utils.go`是Kubernetes项目中`kubeadm`命令行工具的一个子模块。它提供了一些用于处理静态Pod的功能和工具函数。

下面是对变量和函数的详细介绍：

- `usersAndGroups`和`usersAndGroupsOnce`是两个全局变量，用于存储系统中的用户和组信息。`usersAndGroups`是`sync.Map`类型，用于存储用户和组的映射关系，`usersAndGroupsOnce`是`sync.Once`类型，用于在首次访问时加载用户和组信息。

- `ComponentPod`是一个函数，用于创建一个`v1.Pod`对象，表示一个组件的Pod。

- `ComponentResources`是一个函数，用于创建一个`corev1.ResourceRequirements`对象，表示组件Pod的资源需求。

- `NewVolume`是一个函数，用于创建一个`v1.Volume`对象，表示一个存储卷。

- `NewVolumeMount`是一个函数，用于创建一个`v1.VolumeMount`对象，表示一个存储卷的挂载点。

- `VolumeMapToSlice`是一个函数，用于将`v1.Volume`对象的映射关系转换为切片形式。

- `VolumeMountMapToSlice`是一个函数，用于将`v1.VolumeMount`对象的映射关系转换为切片形式。

- `GetExtraParameters`是一个函数，用于从静态Pod的注释中获取额外的参数。

- `PatchStaticPod`是一个函数，用于对静态Pod的配置进行补丁操作。

- `WriteStaticPodToDisk`是一个函数，用于将静态Pod的配置写入磁盘。

- `ReadStaticPodFromDisk`是一个函数，用于从磁盘中读取静态Pod的配置。

- `LivenessProbe`、`ReadinessProbe`和`StartupProbe`分别是用于创建不同类型的探测器（`v1.Probe`对象）的函数。

- `createHTTPProbe`是一个函数，用于创建一个HTTP类型的探测器。

- `GetAPIServerProbeAddress`、`GetControllerManagerProbeAddress`、`GetSchedulerProbeAddress`和`GetEtcdProbeEndpoint`分别是用于获取不同组件的探针地址的函数。

- `ManifestFilesAreEqual`是一个函数，用于比较两个Pod配置文件是否相等。

- `getProbeAddress`是一个函数，用于从Pod对象中获取探针的地址。

- `GetUsersAndGroups`是一个函数，用于获取系统中的用户和组信息。

- `DeepHashObject`是一个函数，用于计算给定对象的哈希值。

