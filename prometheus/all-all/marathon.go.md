# File: discovery/marathon/marathon.go

在Prometheus项目中，discovery/marathon/marathon.go文件的作用是实现与Marathon（一个容器编排工具）的集成，通过监听Marathon的事件来自动发现和监控Marathon中的应用程序。

DefaultSDConfig是用于定义默认的服务发现配置的变量。其中包含了以下几个变量：

- SDConfig：用于配置服务发现的相关信息，如Marathon的API地址、刷新间隔等。
- Discovery：服务发现的接口，定义了发现和获取目标的方法。
- authTokenRoundTripper：用于进行HTTP请求认证的RoundTripper。
- authTokenFileRoundTripper：用于从文件中获取认证令牌的RoundTripper。

下面是几个重要的结构体及其作用：

- task：表示Marathon中的一个任务或者容器。
- ipAddress：表示任务的IP地址。
- portMapping：表示任务的端口映射。
- dockerContainer：表示任务所在的Docker容器。
- container：表示容器的配置信息。
- portDefinition：表示端口定义。
- network：表示容器的网络信息。
- app：表示Marathon中的一个应用程序。
- appList：表示应用程序列表。
- appListClient：用于获取应用程序列表的客户端。

下面是几个重要的函数及其作用：

- init：用于初始化服务发现并注册相关的配置项。
- Name：返回服务发现的名称。
- NewDiscoverer：根据配置信息创建一个新的Marathon发现实例。
- SetDirectory：设置Marathon的API目录。
- UnmarshalYAML：从YAML配置中解析并更新服务发现配置。
- NewDiscovery：根据配置信息创建一个新的服务发现实例。
- newAuthTokenRoundTripper：创建一个新的进行HTTP请求认证的RoundTripper实例。
- RoundTrip：执行HTTP请求。
- newAuthTokenFileRoundTripper：创建一个新的从文件中获取认证令牌的RoundTripper实例。
- refresh：刷新应用程序列表。
- fetchTargetGroups：从Marathon获取目标组信息。
- isContainerNet：检查是否是容器网络。
- fetchApps：从Marathon获取应用程序列表。
- randomAppsURL：随机选择一个应用程序的URL。
- appsToTargetGroups：将应用程序列表转换为目标组列表。
- createTargetGroup：创建一个新的目标组。
- targetsForApp：获取目标组列表中与特定应用程序相关的目标组。
- targetEndpoint：获取目标组的端点。
- extractPortMapping：从容器配置中提取端口映射信息。

