# File: istio/cni/pkg/ebpf/server/redirectionServer_linux.go

在Istio项目中，`redirectionServer_linux.go`文件是用于实现eBPF重定向功能的服务器代码。它负责监听来自其他组件的请求，并执行相应的操作以实现流量的重定向。

下面是对一些变量和结构体的详细介绍：

**变量：**

1. `log`：用于记录日志信息。
2. `isBigEndian`：表示当前系统是否使用Big Endian字节序。
3. `stringToLevel`：映射日志级别的字符串到相应的枚举类型，用于设置日志级别。

**结构体：**

1. `RedirectServer`：表示eBPF重定向服务器，包含用于处理请求的方法和必要的状态信息。
2. `eBPFObjects`：定义了一个eBPF对象的接口，包括加载和卸载eBPF程序的方法。
3. `eBPFObjectsImplOld`：`eBPFObjects`接口的旧实现，用于较旧版本的Linux内核。
4. `eBPFObjectsImplNew`：`eBPFObjects`接口的新实现，用于较新版本的Linux内核。
5. `mapInfo`：表示eBPF中的一张map表，包含表类型和名称等信息。

下面是对一些函数的详细介绍：

1. `Close`：关闭eBPF对象并清理资源。
2. `EBPFTProxySupport`：检查当前系统是否支持eBPF TProxy。
3. `SetLogLevel`：设置日志级别。
4. `UpdateHostIP`：更新主机的IP地址信息。
5. `AddPodToMesh`：将特定Pod添加到mesh中，并进行相应的eBPF配置。
6. `initBpfObjects`：初始化eBPF对象。
7. `NewRedirectServer`：创建一个新的eBPF重定向服务器实例。
8. `checkOrMountBPFFSDefault`：检查或挂载默认的BPFFS文件系统。
9. `setLimit`：设置RLIMIT_MEMLOCK的限制。
10. `Start`：启动eBPF重定向服务器。
11. `parseIPs`：解析IP地址和子网掩码。
12. `RemovePod`：从mesh中移除特定的Pod。
13. `handleRequest`：处理来自其他组件的请求。
14. `AcceptRequest`：接受请求并执行相应的操作。
15. `attachTCForZtunnel`：为Ztunnel流量附加TC规则。
16. `detachTCForZtunnel`：移除Ztunnel流量的TC规则。
17. `detachTCForWorkload`：移除工作负载流量的TC规则。
18. `attachTCForWorkLoad`：为工作负载流量附加TC规则。
19. `attachTC`：为指定的网络接口附加TC规则。
20. `delClsactQdisc`：删除指定网络接口上的CLSACT队列规则。
21. `dumpZtunnelInfo`：打印Ztunnel流量的信息。
22. `DumpAppIPs`：打印应用流量的IP地址。
23. `htons`、`htonl`：字节序转换函数。

这些函数用于实现eBPF重定向服务器的不同功能，包括配置eBPF程序、处理请求、附加和移除TC规则、进行流量信息的打印等。

