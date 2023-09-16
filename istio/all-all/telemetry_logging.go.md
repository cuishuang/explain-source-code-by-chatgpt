# File: istio/pilot/pkg/model/telemetry_logging.go

在Istio项目中，istio/pilot/pkg/model/telemetry_logging.go文件是一份用于配置和处理遥测日志的代码文件。它定义了一些变量和函数，用于为Istio网格中的服务记录访问日志和遥测事件。

下面是对每个变量和函数的详细介绍：

变量：
- clusterLookupFn：是一个用于查找集群的函数类型变量，用于将请求映射到相应的集群。
- EnvoyJSONLogFormatIstio：是一个代表将日志格式为JSON的标志（配置选项）。
- envoyWasmStateToLog：是一个函数变量，用于将Envoy的Wasm模块状态转换为日志格式。
- reqWithoutQueryFormatter：是一个函数变量，用于格式化不带查询参数的HTTP请求。
- metadataFormatter：是一个函数变量，用于格式化元数据。

函数：
- telemetryAccessLog：用于创建TelemetryAccessLog实例，它是一种用于记录访问日志的抽象。
- tcpGrpcAccessLogFromTelemetry：用于构建基于gRPC的TCP访问日志，用于将TCP层的遥测事件发送到远程服务器。
- fileAccessLogFromTelemetry：用于构建基于文件的访问日志，用于将遥测事件记录到本地文件。
- buildFileAccessTextLogFormat：用于构建文本格式的文件访问日志配置。
- buildFileAccessJSONLogFormat：用于构建JSON格式的文件访问日志配置。
- accessLogJSONFormatters：用于构建JSON格式的访问日志格式化器。
- accessLogTextFormatters：用于构建文本格式的访问日志格式化器。
- httpGrpcAccessLogFromTelemetry：用于构建基于gRPC的HTTP访问日志，用于将HTTP层的遥测事件发送到远程服务器。
- fileAccessLogFormat：用于构建访问日志的格式化配置。
- FileAccessLogFromMeshConfig：用于从网格配置中构建访问日志配置。
- openTelemetryLog：用于创建OpenTelemetry access log实例，它是一种将遥测事件发送到OpenTelemetry后端的抽象。
- buildOpenTelemetryAccessLogConfig：用于构建OpenTelemetry访问日志的配置。
- ConvertStructToAttributeKeyValues：用于将结构体转换为键值对的函数。
- LookupCluster：用于查找集群的函数，通过给定的服务和端口信息，返回匹配的集群配置。

这些变量和函数的目的是为了配置和处理遥测日志，包括记录访问日志、构建各种类型的访问日志配置，以及将遥测事件发送到远程服务器或本地文件等。

