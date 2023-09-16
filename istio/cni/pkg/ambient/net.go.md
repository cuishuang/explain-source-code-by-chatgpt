# File: istio/cni/pkg/ambient/net.go

在 Istio 项目中，`istio/cni/pkg/ambient/net.go` 文件的作用是定义了与网络相关的方法和函数。这些方法和函数用于处理 Istio CNI 插件中的网络操作，例如添加和删除 Pod 的网络规则、获取主机 IP 地址等。

以下是一些变量和函数的详细介绍：

1. `log`、`addPodToMeshLog`、`delPodFromMeshLog`：这些变量是用于日志记录的。

2. `RouteExists`：检查指定的路由规则是否存在。

3. `AddPodToMesh`：将指定的 Pod 添加到 Istio mesh 中，并设置相应的网络规则。

4. `addPodToMeshWithIptables`：使用 iptables 添加指定 Pod 的网络规则。

5. `delPodFromMeshWithIptables`：使用 iptables 删除指定 Pod 的网络规则。

6. `delIPsetAndRoute`：删除指定的 IPset 和路由规则。

7. `GetHostIPByRoute`：通过路由规则获取主机的 IP 地址。

8. `getOutboundIP`：获取出站网络流量所使用的 IP 地址。

9. `GetHostIP`：获取主机的 IP 地址。

10. `DelPodFromMesh`：从 Istio mesh 中删除指定的 Pod。

11. `SetProc`：设置 Proc 文件。

12. `cleanStaleIPs`：清理过时的 IP 地址。

13. `cleanupPodsEbpfOnNode`：在节点上清理与 Pod 相关的 eBPF 规则。

这些函数和方法的作用是确保 Pod 在 Istio mesh 中的网络配置正确，并提供与网络规则相关的操作，以便实现 Istio CNI 插件的网络功能。

