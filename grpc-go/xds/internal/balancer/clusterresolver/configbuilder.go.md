# File: grpc-go/xds/internal/balancer/clusterresolver/configbuilder.go

configbuilder.go文件的作用是构建配置，用于解析和构建集群解析器的配置信息。

priorityConfig结构体用于表示优先级配置，包括了名称、权重和本地性信息。

以下是各个函数的详细介绍：

1. buildPriorityConfigJSON：根据传入的权重和本地性信息，构建优先级配置的JSON格式。

2. buildPriorityConfig：根据传入的JSON格式优先级配置，构建优先级配置的结构体。

3. convertClusterImplMapToOutlierDetection：将传入的集群实现配置Map转换为异常检测的配置。

4. makeClusterImplOutlierDetectionChild：根据传入的异常检测配置和输入集群名字，构建异常检测子配置。

5. buildClusterImplConfigForDNS：根据传入的集群名字和DNS配置，构建DNS解析的集群实现配置。

6. buildClusterImplConfigForEDS：根据传入的集群名字和EDS配置，构建EDS解析的集群实现配置。

7. groupLocalitiesByPriority：根据传入的本地性信息，按照优先级进行分组，将本地性信息分为不同的优先级。

8. dedupSortedIntSlice：对传入的排序整数切片进行去重操作。

9. priorityLocalitiesToClusterImpl：根据传入的优先级本地性信息，构建集群实现配置。

