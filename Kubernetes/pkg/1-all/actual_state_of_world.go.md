# File: pkg/controller/volume/attachdetach/cache/actual_state_of_world.go

pkg/controller/volume/attachdetach/cache/actual_state_of_world.go文件的作用是维护一个表示当前实际状态的数据结构，以及管理已附加的卷和待附加的卷之间的映射。它允许卷附加控制器查询节点和卷的状态，并与期望状态进行比较。

以下是几个重要的结构体和其作用：

1. ActualStateOfWorld：表示挂载和卸载卷的当前状态，并提供访问和修改此状态的方法。

2. AttachedVolume：表示已附加的卷。

3. AttachState：表示卷的附加状态。

4. actualStateOfWorld：表示当前的实际状态。

5. attachedVolume：表示已附加的卷。

6. nodeAttachedTo：表示卷附加的节点。

7. nodeToUpdateStatusFor：表示需要更新状态的节点。

以下是几个重要的函数和其作用：

1. String()：输出ActualStateOfWorld的字符串表示形式。

2. NewActualStateOfWorld()：创建一个新的ActualStateOfWorld实例。

3. MarkVolumeAsUncertain()：将卷标记为未知状态。

4. MarkVolumeAsAttached()：将卷标记为已附加状态。

5. MarkVolumeAsDetached()：将卷标记为已分离状态。

6. RemoveVolumeFromReportAsAttached()：从附加报告中删除指定卷。

7. AddVolumeToReportAsAttached()：将卷添加到附加报告中。

8. AddVolumeNode()：将节点添加到附加的卷中。

9. SetVolumeMountedByNode()：设置节点所附加的卷处于已挂载状态。

10. ResetDetachRequestTime()：将卷的分离请求时间重置为零。

11. SetDetachRequestTime()：设置卷的分离请求时间。

12. getNodeAndVolume()：获取附加卷的节点和卷。

13. removeVolumeFromReportAsAttached()：从附加报告中删除指定的卷。

14. addVolumeToReportAsAttached()：将卷添加到附加报告中。

15. updateNodeStatusUpdateNeeded()：更新需要更新状态的节点。

16. SetNodeStatusUpdateNeeded()：设置需要更新状态的节点。

17. DeleteVolumeNode()：从附加的卷中删除节点。

18. GetAttachState()：获取卷的附加状态。

19. InitializeClaimSize()：初始化卷的声明大小。

20. GetClaimSize()：获取卷的声明大小。

21. GetAttachedVolumes()：获取所有已附加的卷。

22. GetAttachedVolumesForNode()：获取附加到给定节点的所有卷。

23. GetAttachedVolumesPerNode()：获取每个节点所附加的卷。

24. GetNodesForAttachedVolume()：获取附加到给定卷的所有节点。

25. GetVolumesToReportAttached()：获取所有需要附加报告的卷。

26. GetVolumesToReportAttachedForNode()：获取给定节点需要附加报告的卷。

27. GetNodesToUpdateStatusFor()：获取需要更新状态的所有节点。

28. getAttachedVolumeFromUpdateObject()：从状态更新对象中获取附加的卷。

29. getAttachedVolume()：获取已附加的卷。

