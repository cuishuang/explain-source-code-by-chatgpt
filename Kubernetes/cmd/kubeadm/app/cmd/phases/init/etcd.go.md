# File: cmd/kubeadm/app/util/etcd/etcd.go

在Kubernetes项目中，`cmd/kubeadm/app/util/etcd/etcd.go`文件的作用是提供与etcd相关的功能函数和结构体，用于与etcd集群进行通信和管理。

`etcdBackoff`是一个用于退避重试的Backoff结构体，用于控制在发生错误时的重试策略。

`ErrNoMemberIDForPeerURL`是一个错误变量，表示找不到对应的etcd成员ID来匹配指定的Peer URL。

以下是几个重要的结构体及其作用：

1. `ClusterInterrogator`：用于与etcd集群进行交互和获取集群状态的结构体。

2. `etcdClient`：封装了etcd服务器的连接和操作接口的结构体。

3. `Client`：etcdClient的别名，用于向后兼容。

4. `Member`：表示etcd成员的结构体，包含成员的ID、名称和地址等信息。

以下是几个重要的函数及其作用：

- `New`：根据提供的参数创建etcdClient结构体的实例。

- `NewFromCluster`：通过读取kubernetes集群配置文件来创建etcdClient结构体的实例。

- `getEtcdEndpoints`：根据提供的kubeadm配置文件和etcd注释获取etcd的端点地址列表。

- `getEtcdEndpointsWithBackoff`：获取etcd的端点地址列表，在失败时进行退避重试。

- `getRawEtcdEndpointsFromPodAnnotation`：从Pod注释中获取未经处理的etcd端点地址列表。

- `getRawEtcdEndpointsFromPodAnnotationWithoutRetry`：在不进行重试的情况下从Pod注释中获取未经处理的etcd端点地址列表。

- `Sync`：检查etcd集群的健康状态，并等待所有成员准备就绪。

- `listMembers`：获取etcd集群的所有成员。

- `GetMemberID`：根据指定的Peer URL获取etcd成员的ID。

- `ListMembers`：获取etcd集群的所有成员，并返回成员的ID、名称和地址等信息。

- `RemoveMember`：从etcd集群中移除指定的成员。

- `AddMember`：将指定的成员添加到etcd集群。

- `AddMemberAsLearner`：将指定的成员作为学习者添加到etcd集群。

- `addMember`：将指定的成员添加到etcd集群。

- `isLearner`：检查指定的成员是否为学习者。

- `MemberPromote`：将指定的成员从学习者升级为正式成员。

- `CheckClusterHealth`：检查etcd集群的健康状态。

- `getClusterStatus`：获取etcd集群的状态。

- `WaitForClusterAvailable`：等待etcd集群可用。

- `GetClientURL`：获取etcd成员的客户端URL。

- `GetPeerURL`：获取etcd成员的Peer URL。

- `GetClientURLByIP`：根据指定的IP地址获取etcd成员的客户端URL。

