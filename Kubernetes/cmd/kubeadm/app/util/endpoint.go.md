# File: cmd/kubeadm/app/util/endpoint.go

在Kubernetes项目中，cmd/kubeadm/app/util/endpoint.go文件的作用是提供一些实用函数来解析和格式化Kubernetes集群的终端点。

1. GetControlPlaneEndpoint函数用于获取控制平面的终端点。它首先检查配置文件中是否定义了controlPlaneEndpoint字段，如果存在则直接返回该地址。否则，它会调用GetLocalAPIEndpoint函数获取本地API服务器的终端点地址，并返回。

2. GetLocalAPIEndpoint函数用于获取本地API服务器的终端点。它会读取配置文件中的Kubernetes API服务器地址，并返回该地址。

3. ParseHostPort函数用于解析主机和端口。它接收一个字符串参数，例如"localhost:8080"，并将其分解为主机和端口两个部分。然后，它返回解析后的主机和端口值。

4. ParsePort函数用于解析端口号。它接收一个字符串参数，例如"8080"，并将其解析为整数值。然后，它返回解析后的端口号。

5. parseAPIEndpoint函数用于解析API服务器的终端点地址。它接收一个字符串参数，例如"http://localhost:8080"，并解析出协议、主机和端口号。然后，它返回解析后的终端点地址。

6. formatURL函数用于格式化URL。它接收多个参数，包括协议、主机、端口、路径等，并将它们组合成一个完整的URL字符串。然后，它返回格式化后的URL字符串。

这些函数是用于处理Kubernetes集群终端点地址的工具函数，用于在Kubeadm应用程序中解析和格式化相关的网络地址。

