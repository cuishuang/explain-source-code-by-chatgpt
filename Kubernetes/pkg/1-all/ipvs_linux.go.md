# File: pkg/proxy/ipvs/util/ipvs_linux.go

pkg/proxy/ipvs/util/ipvs_linux.go文件是Kubernetes项目中用于与Linux内核ipvs模块交互的工具。IPVS（IP Virtual Server）是Linux内核提供的一种负载均衡技术，用于将请求从前端VIP（Virtual IP）地址转发到后端真实服务器的一种网络模块。

该文件中定义了一些结构体和函数，用于封装对ipvs模块的操作。

- runner结构体：用于执行命令和处理输出结果。
- Protocol结构体：表示IP协议，包括TCP和UDP。
- New函数：创建一个runner结构体实例。
- AddVirtualServer函数：添加一个虚拟服务器到ipvs。
- UpdateVirtualServer函数：更新已存在的虚拟服务器的配置。
- DeleteVirtualServer函数：删除一个虚拟服务器。
- GetVirtualServer函数：获取指定的虚拟服务器配置。
- GetVirtualServers函数：获取所有虚拟服务器的配置。
- Flush函数：清空所有虚拟服务器和真实服务器的配置。
- AddRealServer函数：添加一个真实服务器到指定的虚拟服务器。
- DeleteRealServer函数：删除指定虚拟服务器下的一个真实服务器。
- UpdateRealServer函数：更新已存在的真实服务器的配置。
- GetRealServers函数：获取指定虚拟服务器下的所有真实服务器的配置。
- ConfigureTimeouts函数：配置ipvs的超时时间参数。
- toVirtualServer函数：将虚拟服务器配置转换为对应的ipvs服务结构体。
- toRealServer函数：将真实服务器配置转换为对应的ipvs目标结构体。
- toIPVSService函数：将ipvs服务结构体转换为对应的字符串。
- toIPVSDestination函数：将ipvs目标结构体转换为对应的字符串。
- stringToProtocol函数：将字符串表示的协议转换为Protocol结构体。
- protocolToString函数：将Protocol结构体转换为字符串表示的协议。

