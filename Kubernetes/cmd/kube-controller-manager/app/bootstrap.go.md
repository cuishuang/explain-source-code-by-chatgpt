# File: cmd/kube-controller-manager/app/bootstrap.go

在Kubernetes项目中，`cmd/kube-controller-manager/app/bootstrap.go`文件是控制器管理器的启动文件，它负责初始化和启动一些关键的控制器。

这个文件主要进行一些启动时的初始化操作和创建控制器对象，然后通过调用这些控制器的`Run()`方法来启动它们的主循环。下面介绍一些关键的函数和它们的作用：

1. `startBootstrapSignerController`函数：该函数用于启动`bootstrapSignerController`控制器，其主要作用是为新创建的服务账户生成签名令牌。在Kubernetes中，服务账户用于在集群内的各个组件之间进行身份验证和授权。

2. `startTokenCleanerController`函数：该函数用于启动`tokenCleanerController`控制器，其主要作用是清理过期的身份验证令牌。Kubernetes中的身份验证令牌具有一定的有效期，为了保证安全性和性能，需要定期清理过期的令牌。

这些函数会在启动过程中被调用，并通过创建控制器对象来完成所需的操作。控制器会在主循环中监听和处理相关的事件，确保系统正常运行。通过将这些控制器集成到控制器管理器中，可以实现相关功能的自动化管理和维护。

总结起来，`cmd/kube-controller-manager/app/bootstrap.go`文件起到了初始化和启动关键控制器的作用，这些控制器在Kubernetes系统中扮演重要角色，包括服务账户令牌生成和过期令牌清理等功能。

