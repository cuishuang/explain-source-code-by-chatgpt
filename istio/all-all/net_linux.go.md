# File: istio/cni/pkg/ambient/net_linux.go

在istio/cni/pkg/ambient/net_linux.go文件中，定义了一些用于Linux环境的网络操作的函数。下面是这些函数的详细介绍：

1. IsPodInIpset() - 检查一个Pod是否存在于指定的IP集合中。

2. buildEbpfArgsByIP() - 根据给定的IP地址构建eBPF规则所需的参数。

3. updateNodeProxyEBPF() - 更新节点上的eBPF规则，以便代理网络流量。

4. delZtunnelEbpfOnNode() - 删除节点上的ZeroTier网络隧道的eBPF规则。

5. updatePodEbpfOnNode() - 更新节点上与特定Pod相关的eBPF规则。

6. delPodEbpfOnNode() - 删除节点上与特定Pod相关的eBPF规则。

7. getLinkWithDestinationOf() - 获取具有指定目标地址的网络链路。

8. getVethWithDestinationOf() - 获取具有指定目标地址的veth设备。

9. getDeviceWithDestinationOf() - 获取具有指定目标地址的网络设备。

10. GetIndexAndPeerMac() - 获取指定网络设备的索引和对等MAC地址。

11. getMacFromNsIdx() - 从指定的网络命名空间索引获取MAC地址。

12. getNsNameFromNsID() - 从指定的网络命名空间ID获取命名空间名称。

13. getPeerIndex() - 获取指定网络设备的对等索引。

14. CreateRulesOnNode() - 在节点上创建网络规则。

15. determineDstPortForGeneveLink() - 确定Geneve链路的目标端口。

16. CreateEBPFRulesWithinNodeProxyNS() - 在节点代理命名空间中创建eBPF规则。

17. createTProxyRulesForLegacyEBPF() - 为传统的eBPF规则创建TProxy规则。

18. CreateRulesWithinNodeProxyNS() - 在节点代理命名空间中创建网络规则。

19. ztunnelDown() - 关闭ZeroTier隧道。

20. cleanupNode() - 清理节点的网络配置。

21. getEnrolledIPSets() - 获取已注册的IP集合。

22. addTProxyMarkRule() - 添加TProxy标记规则。

23. addOrgSrcMarkRule() - 添加源组织标记规则。

24. disableRPFiltersForLink() - 禁用链路上的反向路径过滤器。

25. setProc() - 修改/proc文件中的某个值。

26. buildRouteForPod() - 为Pod构建路由。

27. flushAllRouteTables() - 清空所有路由表。

28. routeFlushTable() - 清空指定路由表。

29. deleteIPRules() - 删除IP规则。

30. deleteTunnelLinks() - 删除隧道链接。

31. routesDelete() - 删除指定路由。

