# File: cmd/kube-controller-manager/app/controllermanager.go

在 Kubernetes 项目中，cmd/kube-controller-manager/app/controllermanager.go 文件是 kube-controller-manager 二进制文件的入口点，它负责启动和管理一组控制器（也称为引导程序），每个控制器负责监控和调整 Kubernetes 集群的状态。

具体来说，该文件实现了以下功能：
- 初始化命令行标志和配置参数。
- 定义了一些变量和结构体，用于控制器的管理和配置。
- 定义了一些函数，用于初始化控制器和启动控制器管理器。

现在来详细介绍一下其中的变量和结构体：

1. `ControllersDisabledByDefault`：该变量是一个布尔值，用于标识是否禁用默认的控制器。

2. `ControllerLoopMode`：该结构体定义了控制器的循环模式，包括短循环模式和长循环模式。

3. `ControllerContext`：该结构体封装了控制器运行所需的上下文信息，包括客户端构建器、资源获取器等。

4. `InitFunc`：该结构体定义了控制器初始化的函数签名，用于在控制器启动前进行初始化操作。

5. `ControllerInitializersFunc`：该结构体定义了一组控制器初始化函数的集合。

6. `serviceAccountTokenControllerStarter`：该结构体用于启动 Service Account Token 控制器。

接下来让我们看一下一些重要的函数及其作用：

1. `init`：该函数会在程序启动时被自动调用，用于初始化一些全局变量。

2. `NewControllerManagerCommand`：该函数用于创建 kube-controller-manager 的命令。

3. `ResyncPeriod`：该函数确定了控制器的重新同步周期。

4. `Run`：该函数是主要的控制器管理器运行函数，它负责启动和管理控制器。

5. `IsControllerEnabled`：该函数用于检查指定的控制器是否启用。

6. `KnownControllers`：该函数返回了 Kubernetes 中已知的所有控制器列表。

7. `NewControllerInitializers`：该函数用于创建一组控制器初始化函数。

8. `GetAvailableResources`：该函数用于获取可用的资源列表。

9. `CreateControllerContext`：该函数用于创建控制器的上下文对象。

10. `StartControllers`：该函数用于启动控制器。

11. `startServiceAccountTokenController`：该函数用于启动 Service Account Token 控制器。

12. `readCA`：该函数用于读取根证书。

13. `createClientBuilders`：该函数用于创建客户端构建器。

14. `leaderElectAndRun`：该函数用于进行领导选举，并开始执行控制器循环。

15. `createInitializersFunc`：该函数用于创建控制器初始化函数列表。

总结一下，controllermanager.go 文件实现了 kube-controller-manager 的入口点，包括参数的解析、控制器的管理和控制器初始化等功能。它是控制器管理器的核心文件，负责启动和管理一组控制器，以确保 Kubernetes 集群的正常运行。

