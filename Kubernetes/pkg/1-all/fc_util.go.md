# File: pkg/volume/fc/fc_util.go

pkg/volume/fc/fc_util.go文件是Kubernetes项目中用于处理Fibre Channel (FC) 存储卷的工具函数文件。它包含了一些帮助函数和类型定义，用于在Kubernetes集群中管理和操作FC存储卷。

ioHandler是FC存储卷的输入/输出处理程序接口。它定义了多个方法，包括读取目录、获取文件信息、解析链接、写入文件等。

osIOHandler是ioHandler接口的实现，它使用标准的操作系统IO函数来执行FC存储卷的输入/输出操作。

fcUtil是用于管理、操作和配置FC存储卷的工具函数。它提供了一些常用的功能，例如查找磁盘、刷新设备、从SCSI子系统中删除设备等。

下面是对一些重要函数的介绍：

- ReadDir: 读取指定目录中的所有文件和子目录的名称。
- Lstat: 获取指定路径的文件信息，包括文件类型、大小、权限等。
- EvalSymlinks: 解析指定路径上的所有符号链接，并返回最终目标路径。
- WriteFile: 将指定的数据写入到指定路径的文件中。
- findDisk: 查找指定的FC存储磁盘。
- findDiskWWIDs: 根据指定的WWIDs（World Wide Identifiers）查找FC存储磁盘。
- flushDevice: 刷新指定设备，使更改生效。
- removeFromScsiSubsystem: 从SCSI子系统中移除指定设备。
- scsiHostRescan: 扫描SCSI主机，以刷新扫描过的设备列表。
- makePDNameInternal: 根据指定的Fibre Channel主机号和端口号生成FC存储PD名称。
- makeVDPDNameInternal: 根据指定的Fibre Channel主机号、端口号和存储虚拟机名生成FC存储VD PD名称。
- parsePDName: 解析FC存储PD名称，提取Fibre Channel主机号和端口号。
- MakeGlobalPDName: 根据指定的Fibre Channel主机号和端口号生成全局唯一的FC存储PD名称。
- MakeGlobalVDPDName: 根据指定的Fibre Channel主机号、端口号和存储虚拟机名生成全局唯一的FC存储VD PD名称。
- searchDisk: 在指定路径下搜索符合条件的磁盘。
- AttachDisk: 将FC存储磁盘附加到指定的节点上。
- DetachDisk: 将FC存储磁盘从指定节点上分离。
- detachFCDisk: 分离指定的FC存储磁盘。
- DetachBlockFCDisk: 分离指定的FC块存储磁盘。
- deleteMultipathDevice: 删除指定的多路径设备。
- checkPathExists: 检查指定路径是否存在。

这些函数和结构体提供了在Kubernetes中管理和操作FC存储卷所需的底层功能和工具。通过这些函数和结构体，可以实现FC存储卷的附加、分离、删除等操作，并对FC存储卷进行读写和管理。

