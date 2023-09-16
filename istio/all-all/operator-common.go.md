# File: istio/operator/cmd/mesh/operator-common.go

文件 `istio/operator/cmd/mesh/operator-common.go` 是 Istio 项目中 Operator 的一个公共文件，用于处理 Operator 的公共操作。

`operatorCommonArgs` 结构体是用于定义 Operator 的公共参数的。它包含了一些配置选项，如命名空间、部署名称、日志配置等。

`isControllerInstalled` 函数用于检查 Istio Operator 控制器是否已经安装到 Kubernetes 集群中。它会查询指定的命名空间中是否存在 `istio-operator` 的 Deployment 资源，如果存在则认为已经安装。如果未安装，则会返回 `false`。

`renderOperatorManifest` 函数用于创建 Istio Operator 控制器的 Kubernetes Manifest（配置文件）。它会根据传入的参数和默认值，生成一个包含部署、服务和服务账号等资源的 Kubernetes 配置。

`deploymentExists` 函数用于检查指定部署在指定命名空间中是否已经存在。它会查询指定命名空间中的 Deployment 资源，并检查部署名称是否匹配。存在则返回 `true`，否则返回 `false`。

这些函数都是用于处理和管理 Istio Operator 控制器的常见操作。

