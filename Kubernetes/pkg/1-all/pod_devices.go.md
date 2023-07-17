# File: pkg/kubelet/cm/devicemanager/pod_devices.go

pkg/kubelet/cm/devicemanager/pod_devices.go文件的作用是管理Kubernetes集群中的设备分配情况。它定义了一系列结构体和函数，用于维护设备的分配信息和状态。

1. deviceAllocateInfo结构体用于表示设备的分配信息，包括设备ID和设备路径等信息。
2. resourceAllocateInfo结构体用于表示资源的分配信息，包括设备的数量和资源限制等信息。
3. containerDevices结构体用于表示容器的设备分配情况，包括容器ID和分配的设备列表等信息。
4. podDevices结构体用于表示Pod的设备分配情况，包括Pod ID和容器设备列表等信息。
5. DeviceInstances结构体表示设备实例的分配情况，包括设备ID和容器ID等信息。
6. ResourceDeviceInstances结构体表示资源设备实例的分配情况，包括设备ID和资源分配信息等。
7. newPodDevices函数用于创建一个新的podDevices对象。
8. pods函数返回podDevices中维护的所有Pod的ID列表。
9. size函数返回podDevices中维护的Pod数量。
10. hasPod函数检查podDevices中是否包含指定的Pod ID。
11. insert函数将Pod设备信息插入到podDevices中。
12. delete函数从podDevices中删除指定的Pod设备信息。
13. podDevices函数返回指定Pod ID的设备分配信息。
14. containerDevices函数返回指定容器ID的设备分配信息。
15. addContainerAllocatedResources函数将容器的设备资源分配信息添加到podDevices中。
16. removeContainerAllocatedResources函数从podDevices中移除指定容器的设备资源分配信息。
17. devices函数返回所有分配的设备ID列表。
18. toCheckpointData函数将podDevices转换为可保存的检查点数据格式。
19. fromCheckpointData函数从检查点数据中恢复podDevices。
20. deviceRunContainerOptions函数返回指定容器的设备运行配置。
21. getContainerDevices函数返回指定容器ID的设备分配信息。
22. NewResourceDeviceInstances函数创建一个新的ResourceDeviceInstances对象。
23. Clone函数对podDevices进行深拷贝。
24. Filter函数根据指定的过滤条件过滤设备分配信息。

