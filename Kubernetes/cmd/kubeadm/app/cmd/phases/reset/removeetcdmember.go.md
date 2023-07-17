# File: cmd/kubeadm/app/cmd/phases/reset/removeetcdmember.go

文件 removeetcdmember.go 是 kubeadm 命令行工具中一个重要的文件，它负责定义和实现 kubeadm reset 过程中移除 etcd 成员的阶段。

`NewRemoveETCDMemberPhase` 函数的作用是创建一个新的 kubeadm reset 过程中移除 etcd 成员的阶段。此阶段实现了 kubeadm reset 命令的 "remove-etcd-member" 步骤。它负责卸载当前 etcd 成员，将其从 etcd 集群中移除，并最终将 etcd 数据目录清除。

`getEtcdDataDir` 函数的作用是获取 etcd 成员的数据目录。它首先检查 etcd 的运行模式，然后根据运行模式返回相应的 etcd 数据目录。

`runRemoveETCDMemberPhase` 函数是整个移除 etcd 成员阶段的执行逻辑。该函数首先通过调用 `getEtcdDataDir` 函数获取 etcd 数据目录，然后关闭 etcd 服务，并卸载该成员。接下来，它使用 `wait.Port` 函数等待 etcd 服务的关闭，然后调用 `etcdutil.IsMember` 函数来验证 etcd 成员是否已从集群中移除。最后，该函数删除 etcd 数据目录。

简而言之，该文件的作用是实现了 kubeadm reset 过程中移除 etcd 成员的阶段，包括卸载 etcd 成员、从集群中移除该成员、删除数据目录等操作。

