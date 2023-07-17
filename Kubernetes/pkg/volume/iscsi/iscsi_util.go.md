# File: pkg/volume/iscsi/iscsi_util.go

pkg/volume/iscsi/iscsi_util.go这个文件是Kubernetes项目中用于支持iSCSI存储卷的工具函数和结构体。

作用如下：
- 实现了一系列的功能函数，用于管理iSCSI存储卷的连接和断开连接，以及设备的挂载和卸载。
- 定义了多个结构体，用于封装iSCSI相关的操作，如检测设备状态、扫描LUN、等待设备路径出现等。

下面是一些重要变量和结构体的介绍：

1. 变量：
- chapSt: iSCSI CHAP信息的状态变量。
- chapSess: iSCSI CHAP会话的状态变量。
- ifaceTransportNameRe: 用于提取iSCSI接口的传输名称的正则表达式。
- ifaceRe: 用于提取iSCSI接口的正则表达式。

2. 结构体：
- StatFunc: 函数类型，用于执行系统命令并返回执行结果的统计信息。
- GlobFunc: 函数类型，用于在文件系统中查找匹配指定模式的文件路径。
- ISCSIUtil: 用于封装iSCSI相关操作的工具类，包括设备挂载、断开连接、设备状态检测等。

下面是一些重要函数的介绍：

- updateISCSIDiscoverydb: 更新iSCSI发现数据库。
- updateISCSINode: 更新iSCSI节点信息。
- waitForPathToExist: 等待设备路径出现。
- waitForPathToExistInternal: 内部函数，用于实现等待设备路径出现的逻辑。
- makePDNameInternal: 内部函数，用于生成持久化存储设备的名称。
- makeVDPDNameInternal: 内部函数，用于生成虚拟持久化存储设备的名称。
- MakeGlobalPDName: 生成全局持久化存储设备的名称。
- MakeGlobalVDPDName: 生成全局虚拟持久化存储设备的名称。
- persistISCSIFile: 持久化iSCSI连接信息到文件系统。
- loadISCSI: 从文件系统加载iSCSI连接信息。
- scanOneLun: 扫描一个LUN（逻辑单元号）。
- waitForMultiPathToExist: 等待多路径设备路径出现。
- AttachDisk: 挂载iSCSI存储卷到指定节点上。
- persistISCSI: 持久化iSCSI连接信息。
- deleteDevice: 删除iSCSI设备。
- deleteDevices: 删除多个iSCSI设备。
- DetachDisk: 卸载iSCSI存储卷。
- DetachBlockISCSIDisk: 卸载块形式的iSCSI存储卷。
- detachISCSIDisk: 卸载iSCSI设备。
- getDevByPath: 根据设备路径获取设备信息。
- extractTransportname: 从iSCSI接口中提取传输名称。
- extractDeviceAndPrefix: 从iSCSI设备路径中提取设备和前缀信息。
- extractIface: 从iSCSI设备路径中提取接口信息。
- extractPortalAndIqn: 从iSCSI设备路径中提取目标地址和iSCSI Qualified Name。
- removeDuplicate: 去除重复的元素。
- parseIscsiadmShow: 解析iscsiadm show命令的输出。
- cloneIface: 克隆iSCSI接口。
- isSessionBusy: 判断iSCSI会话是否处于繁忙状态。
- getVolCount: 获取iSCSI卷的数量。
- ignoreExitCodes: 忽略指定的命令执行结果。
- execWithLog: 执行系统命令，并将命令执行日志记录下来。

