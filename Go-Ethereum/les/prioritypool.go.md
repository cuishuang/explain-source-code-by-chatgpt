# File: les/vflux/server/prioritypool.go

在Go-Ethereum项目中，les/vflux/server/prioritypool.go文件的作用是实现了网络连接池的优先级分配逻辑。该连接池用于管理网络连接，并根据连接的优先级来分配资源。

以下是对每个结构体及其作用的详细介绍：

1. priorityPool：连接池的主结构体，用于管理连接和分配资源。
2. ppNodeInfo：保存关于节点连接的重要信息，如连接状态、优先级和临时状态等。
3. capUpdate：用于更新连接容量的结构体，保存容量变化的信息。
4. capacityCurve：连接容量曲线结构体，用于根据优先级分配资源的曲线算法。
5. curvePoint：曲线上的点，保存了特定优先级和对应容量的信息。

以下是对每个函数及其作用的详细介绍：

1. newPriorityPool：创建一个新的连接池实例。
2. requestCapacity：根据节点优先级和容量需求，请求分配连接容量。
3. SetLimits：设置连接池的上限和下限。
4. setActiveBias：设置活动连接偏好。
5. Active：激活指定的连接节点。
6. Inactive：将指定的连接节点设置为非活动状态。
7. Limits：获取连接池的上限和下限。
8. inactiveSetIndex：获取非活动队列的索引。
9. activeSetIndex：获取活动队列的索引。
10. invertPriority：反转连接节点的优先级。
11. activePriority：获取活动连接节点的优先级。
12. activeMaxPriority：获取活动连接节点的最大优先级。
13. inactivePriority：获取非活动连接节点的优先级。
14. removeFromQueues：将连接节点从活动和非活动队列中移除。
15. connectNode：将连接节点添加到连接池中。
16. disconnectNode：将连接节点从连接池中移除。
17. setTempState：设置连接节点的临时状态。
18. unsetTempState：取消连接节点的临时状态。
19. setTempCapacity：设置连接节点的临时容量。
20. setTempBias：设置连接节点的临时偏好。
21. setTempStepDiv：设置连接节点的临时步长。
22. enforceLimits：根据连接池的上下限约束对连接节点进行限制。
23. finalizeChanges：在容量更新后，对连接节点进行最终调整。
24. updateFlags：更新连接节点的标记。
25. tryActivate：尝试激活一个连接节点。
26. updatePriority：更新连接节点的优先级。
27. getCapacityCurve：获取连接容量曲线。
28. exclude：排除指定的连接节点。
29. getPoint：获取连接容量曲线上的指定点。
30. maxCapacity：获取连接池的最大容量。

这些函数和结构体共同协作，实现了连接池的优先级分配逻辑，管理和调度网络连接的资源分配。

