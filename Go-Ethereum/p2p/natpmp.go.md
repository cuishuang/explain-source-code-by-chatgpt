# File: p2p/nat/natpmp.go

在go-ethereum项目中，p2p/nat/natpmp.go文件的作用是实现对NAT-PMP（Network Address Translation Port Mapping Protocol）的支持。NAT-PMP是一种用于通过NAT设备进行端口映射的协议，允许将内部网络的服务暴露给外部网络。

该文件中定义了一些与NAT-PMP相关的结构体和函数：

1. `pmp`结构体：表示一个NAT-PMP客户端的实例，包含了与NAT-PMP服务器通信的相关参数和状态信息。

2. `String`方法：返回一个包含此`pmp`实例信息的字符串表示。

3. `ExternalIP`方法：向NAT-PMP服务器发起请求，获取外部IP地址。

4. `AddMapping`方法：向NAT-PMP服务器发起请求，创建一个端口映射规则，将外部端口映射到内部网络的指定端口。

5. `DeleteMapping`方法：向NAT-PMP服务器发起请求，删除指定的端口映射规则。

6. `discoverPMP`函数：通过向内网广播NAT-PMP请求，发现可用的NAT-PMP服务器。

7. `potentialGateways`函数：根据本地网络适配器配置，尝试发现可能是NAT设备的网关地址。

这些函数和结构体的作用是为了实现和管理与NAT-PMP服务器的通信，从而实现通过NAT设备进行端口映射的功能。

