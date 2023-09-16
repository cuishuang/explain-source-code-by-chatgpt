# File: istio/pilot/pkg/serviceregistry/kube/controller/ambientindex_external.go

在Istio项目中，`istio/pilot/pkg/serviceregistry/kube/controller/ambientindex_external.go`文件的作用是实现了外部服务索引的控制器，用于将来自Kubernetes集群中的服务和工作负载信息注册到Pilot中，以供Istio进行流量管理和服务发现。

以下是对这些函数的详细介绍：

1. `handleServiceEntry`：用于处理ServiceEntry对象的事件，根据对象的创建、更新或删除操作，对Pilot中的外部服务索引进行相应的操作。

2. `getWorkloadEntriesInPolicy`：根据指定的策略名称，从外部服务注册表中获取与之相关的工作负载条目。

3. `extractWorkloadEntry`：从外部服务注册表中提取并返回指定的工作负载条目。

4. `extractWorkloadEntrySpec`：从工作负载条目中提取并返回与规范相关的内容，例如Selector，ServiceAccount和Labels等。

5. `handleWorkloadEntry`：用于处理WorkloadEntry对象的事件，根据对象的创建、更新或删除操作，对Pilot中的外部服务索引进行相应的操作。

6. `constructWorkloadFromWorkloadEntry`：根据工作负载条目构建一个工作负载对象。

7. `updateWaypointForWorkload`：更新工作负载的路由规则，将其添加到Pilot的Waypoint中。

8. `getWorkloadEntryServices`：获取与指定的工作负载条目相关联的服务列表。

9. `findPortForWorkloadEntry`：根据工作负载条目中定义的端口名称，查找并返回与之关联的端口信息。

10. `getWorkloadEntriesInService`：根据Service条目的名称，获取与之关联的工作负载条目列表。

11. `getSelectedWorkloadEntries`：获取经过筛选的、与指定标签匹配的工作负载条目列表。

12. `getControllerWorkloadEntries`：获取与控制器相关的工作负载条目列表。

13. `generateWorkloadEntryUID`：生成工作负载条目的唯一标识符。

14. `generateServiceEntryUID`：生成ServiceEntry的唯一标识符。

15. `cleanupOldWorkloadEntriesInlinedOnServiceEntry`：在ServiceEntry对象中清理旧的工作负载条目。

16. `getWorkloadServicesFromServiceEntries`：从ServiceEntry对象中获取工作负载服务列表。

17. `getVIPsFromServiceEntry`：从ServiceEntry对象中获取虚拟IP地址列表。

18. `getPortsForServiceEntry`：从ServiceEntry对象中获取端口信息列表。

这些函数通过与Kubernetes集群中的服务和工作负载对象进行交互，完成了将服务和工作负载信息注册到Pilot的过程，并提供了一些功能用于获取关联信息、更新路由规则等操作。

