# File: pkg/controller/deployment/rolling.go

pkg/controller/deployment/rolling.go是kubernetes中处理Deployment Controller中rolling update的关键文件之一。rolling update是一种渐进式的更新方式，是通过逐步删除旧的Pod并逐步添加新的Pod来实现应用的无缝更新，从而保证应用持续可用。

这个文件实现了rolling update的主要逻辑。主要包括以下几个部分：

1. rolloutRolling: 更新rolling update状态，并触发reconcileNewReplicaSet函数，以创建新的ReplicaSet，同时触发rolloutOneBatch函数，对一批Pod进行rolling update操作。

2. reconcileNewReplicaSet: 根据当前Deployment容器的spec中的ReplicaSet数值，创建新的ReplicaSet，并将其更新到Deployment中的status部分

3. reconcileOldReplicaSets: 更新旧的ReplicaSet，并检查是否需要进行scale down操作。

4. cleanupUnhealthyReplicas: 清理不健康的Pod

5. scaleDownOldReplicaSetsForRollingUpdate: 对旧的ReplicaSet执行scale down。该函数会被调用多次，直到所有旧的ReplicaSet都被删除。

这些函数基于Deployment Controller的rolling update策略，实现了对rolling update操作的控制和处理，并保证了应用更新的平稳和可用。

