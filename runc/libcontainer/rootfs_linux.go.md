# File: runc/libcontainer/rootfs_linux.go

在runc项目中的`runc/libcontainer/rootfs_linux.go`文件是用来处理Linux系统中的rootfs（root file system）相关操作的。

该文件中定义了两个结构体：`mountConfig`和`mountEntry`。

`mountConfig`结构体包含了一些用于挂载rootfs的配置信息，例如`rootPropagation`表示挂载类型，`readonly`表示是否以只读方式挂载。

`mountEntry`结构体则用来描述一个具体的挂载点，包括源路径`source`、目标路径`destination`、文件系统类型`fstype`等。

这些函数的作用如下：

- `src`：根据提供的路径返回一个`syscall.Stat_t`结构体，包含文件的属性信息。
- `needsSetupDev`：检查是否需要设置设备节点。
- `prepareRootfs`：准备rootfs，会进行一系列的准备工作，包括创建目录、挂载tmpfs等。
- `finalizeRootfs`：在准备rootfs的基础上进行一些后续处理，例如设置权限、复制文件等。
- `prepareTmp`：准备tmp目录。
- `cleanupTmp`：清理tmp目录。
- `prepareBindMount`：准备bind挂载点。
- `mountCgroupV1`：挂载Cgroup V1。
- `mountCgroupV2`：挂载Cgroup V2。
- `doTmpfsCopyUp`：复制tmpfs。
- `mountToRootfs`：将挂载点挂载到rootfs中。
- `getCgroupMounts`：获取Cgroup挂载点。
- `checkProcMount`：检查是否已经挂载了proc文件系统。
- `isProc`：判断路径是否是proc文件系统。
- `setupDevSymlinks`：设置/dev目录下的符号链接。
- `reOpenDevNull`：重新打开/dev/null。
- `createDevices`：创建设备节点。
- `bindMountDeviceNode`：绑定挂载设备节点。
- `createDeviceNode`：创建设备节点。
- `mknodDevice`：创建设备节点。
- `getParentMount`：获取父级挂载点。
- `rootfsParentMountPrivate`：设置rootfs的父级挂载点属性。
- `prepareRoot`：准备root目录。
- `setReadonly`：设置rootfs为只读。
- `setupPtmx`：设置ptmx。
- `pivotRoot`：切换rootfs。
- `msMoveRoot`：移动root到指定挂载点。
- `chroot`：切换root。
- `createIfNotExists`：如果路径不存在则创建。
- `readonlyPath`：设置路径为只读。
- `remountReadonly`：重新挂载为只读。
- `maskPath`：掩盖路径。
- `writeSystemProperty`：写入系统属性。
- `remount`：重新挂载。
- `mountPropagate`：挂载传播属性。
- `setRecAttr`：设置属性。

这些函数在runc项目中是用来完成对rootfs进行操作的，例如挂载文件系统、设置只读等操作。通过这些函数的组合使用，可以完成对rootfs的准备、设置和清理等工作。

