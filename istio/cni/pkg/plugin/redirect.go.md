# File: istio/cni/pkg/plugin/redirect.go

redirect.go文件位于istio/cni/pkg/plugin目录下，是Istio CNI插件的一部分。该文件定义了一些常量、结构体和函数，用于配置和管理Pod网络流量的重定向。

以下是每个变量的作用：

- includeIPCidrsKey：用于配置可以访问Pod IPC的CIDR列表。
- excludeIPCidrsKey：用于配置被禁止访问Pod IPC的CIDR列表。
- excludeInboundPortsKey：用于配置禁止Pod接收流量的端口列表。
- includeInboundPortsKey：用于配置允许Pod接收流量的端口列表。
- excludeOutboundPortsKey：用于配置禁止Pod发出流量的端口列表。
- includeOutboundPortsKey：用于配置允许Pod发出流量的端口列表。
- excludeInterfacesKey：用于配置禁止Pod使用的网络接口列表。
- sidecarInterceptModeKey：用于配置Istio sidecar代理的拦截模式，控制流量的转发。
- sidecarPortListKey：用于配置Istio sidecar代理监听的端口列表。
- kubevirtInterfacesKey：用于配置Kubernetes虚拟网络接口列表。
- annotationRegistry：用于注册具有特定处理逻辑的注释的映射。

以下是每个结构体的作用：

- Redirect：定义了Pod网络流量的重定向规则，包括允许的CIDR列表、端口列表和网络接口列表等。
- annotationValidationFunc：定义了一个用于验证注释的函数类型。
- annotationParam：定义了一个注释的验证参数结构。

以下是每个函数的作用：

- alwaysValidFunc：一个注释验证函数，始终返回true，即始终认为注释是有效的。
- validateInterceptionMode：验证拦截模式注释的有效性。
- validateCIDRList：验证CIDR列表注释的有效性。
- splitPorts：将端口列表注释拆分为单独的端口字符串列表。
- dedupPorts：从端口列表中移除重复的端口。
- parsePort：将单个端口字符串解析为端口号和协议类型。
- parsePorts：解析多个端口字符串，返回端口列表。
- validatePortList：验证端口列表注释的有效性。
- validatePortListWithWildcard：验证带有通配符的端口列表注释的有效性。
- validateCIDRListWithWildcard：验证带有通配符的CIDR列表注释的有效性。
- getAnnotationOrDefault：获取指定注释的值，如果注释不存在，则返回默认值。
- NewRedirect：根据传入的注释创建一个Redirect结构体。

这些函数和结构体的目标是验证和处理Pod中的注释配置，创建重定向规则，并对Pod的网络流量进行必要的重定向操作，以便与Istio代理进行通信以实现流量控制。

