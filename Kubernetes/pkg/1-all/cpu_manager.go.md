# File: pkg/kubelet/cm/cpumanager/cpu_manager.go

在kubernetes项目中，pkg/kubelet/cm/cpumanager/cpu_manager.go文件的作用是实现CPU管理器的功能。CPU管理器负责管理和分配Pod的CPU资源。

现在来介绍每个变量和结构体的作用：

1. `_`：在Go语言中，`_`表示一个匿名变量或匿名成员，表示不关心该变量的具体值，只是用于占位或忽略某个返回值。

2. `ActivePodsFunc`：一个函数类型的变量，其返回值是一个`PodLister`接口对象，用于获取当前活动的Pod列表。

3. `runtimeService`：一个接口类型的变量，用于与运行时（如Docker或Containerd）进行交互，执行容器相关的操作，如创建、删除、重启容器等。

4. `policyName`：字符串类型的变量，表示CPU管理策略的名称。

5. `Manager`：一个接口类型的变量，表示CPU管理器的接口，在实际运行中的类型为`internalcpu.Manager`。

6. `manager`：一个结构体类型的变量，表示CPU管理器的实例。

7. `sourcesReadyStub`：一个用于触发事件的函数。

8. `reconciledContainer`：一个代表已协调容器的映射表。

接下来是每个函数的作用：

1. `AddSource`：向CPU管理器中添加新的资源源（如NUMA节点）。

2. `AllReady`：检查CPU管理器是否准备就绪，所有源的状态都必须为就绪。

3. `NewManager`：创建一个新的CPU管理器。

4. `Start`：启动CPU管理器，开始进行资源管理。

5. `Allocate`：根据Pod的CPU需求，分配合适的CPU资源。

6. `AddContainer`：向CPU管理器添加一个新的容器。

7. `RemoveContainer`：从CPU管理器中移除一个容器。

8. `policyRemoveContainerByID`：根据容器的ID从CPU管理器中移除容器。

9. `policyRemoveContainerByRef`：根据容器的引用从CPU管理器中移除容器。

10. `State`：返回CPU管理器的当前状态。

11. `GetTopologyHints`：获取CPU拓扑的提示信息。

12. `GetPodTopologyHints`：获取特定Pod的CPU拓扑的提示信息。

13. `GetAllocatableCPUs`：获取可分配的CPU资源列表。

14. `removeStaleState`：移除过时的状态信息。

15. `reconcileState`：根据当前的Pod状态和容器状态，协调CPU管理器的状态。

16. `findContainerIDByName`：根据容器名称查找容器的ID。

17. `findContainerStatusByName`：根据容器名称查找容器的状态。

18. `updateContainerCPUSet`：更新容器的CPU分配。

19. `GetExclusiveCPUs`：获取独占的CPU资源列表。

20. `GetCPUAffinity`：获取CPU亲和性。

21. `setPodPendingAdmission`：设置Pod等待许可。

