# File: discovery/hetzner/hetzner.go

discovery/hetzner/hetzner.go文件是Prometheus项目中的一个实现Hetzner云平台服务发现的组件。它的主要作用是与Hetzner云平台进行交互，从而发现和监控在Hetzner云环境中运行的目标。

下面是对DefaultSDConfig变量的作用进行详细描述：

1. SDConfig：这是一个结构体，用于存储Hetzner服务发现配置的各种参数，如API密钥、服务器类型、标签等。
2. refresher：一个函数，用于定期刷新从Hetzner云平台获取的目标信息。
3. role：用于标识Hetzner实例的角色，可以是server、loadBalancer或者firewall。
4. Discovery：这是一个Hetzner服务发现实例。

下面是对SDConfig、refresher、role和Discovery结构体的作用进行详细描述：

1. SDConfig结构体：它存储了Hetzner服务发现的配置信息，即用户在配置文件中定义的相关参数。
2. refresher函数：它是一个定期刷新函数，用于从Hetzner云平台获取最新的目标信息。
3. role：它标识了Hetzner实例的角色，即该实例的功能是server、loadBalancer还是firewall。
4. Discovery结构体：这是一个表示Hetzner服务发现的实例，它包含了与Hetzner云平台交互的方法和属性。

下面是对这些函数的作用进行详细描述：

1. init函数：在包被导入时，执行一些初始化操作。
2. Name函数：返回Hetzner云平台的名称，用于标识服务发现组件。
3. NewDiscoverer函数：创建一个新的Hetzner服务发现实例。
4. UnmarshalYAML函数：将YAML格式的配置文件解析成对应的SDConfig结构体。
5. SetDirectory函数：设置Hetzner服务发现的目录，用于存储目标信息。
6. NewDiscovery函数：创建一个新的Hetzner服务发现实例。
7. newRefresher函数：创建一个新的刷新函数，用于定期获取最新的目标信息。

这些函数一起实现了Hetzner云平台服务发现的各种功能，包括配置解析、目标刷新和与Hetzner云平台的交互。

