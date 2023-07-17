# File: pkg/kubelet/cm/memorymanager/memory_manager.go

在kubernetes项目中，pkg/kubelet/cm/memorymanager/memory_manager.go文件的作用是实现Kubelet的内存管理功能。该文件定义了MemoryManager接口及其具体实现。内存管理器的主要职责是跟踪和管理每个节点上容器和Pod的内存使用情况，确保不超出节点的内存限制，并根据容器的请求分配和释放内存资源。

下面是文件中提到的各个变量和结构体的作用：

- _ 变量: 用于包引用，表示忽略这个包的返回值。

- ActivePodsFunc 变量: 一个函数指针，用于获取活跃的Pod列表。

- runtimeService 变量: 运行时服务的接口，用于与容器运行时进行通信。

- sourcesReadyStub 变量: 用于生成带有特定状态的文件系统源对象。

- Manager 结构体: 内存管理器的主要结构体，实现了MemoryManager接口的所有方法。

- manager 结构体: 实际的内存管理器实例，通过Manager结构体的AddSource、AllReady、NewManager等方法来创建并操作。

下面是文件中提到的各个函数的作用：

- AddSource 函数: 添加一个用于监视内存增长的来源。

- AllReady 函数: 检查所有内存源是否都已准备就绪。

- NewManager 函数: 创建一个新的内存管理器实例。

- Start 函数: 启动内存管理器。

- AddContainer 函数: 向内存管理器添加一个容器，并为其分配内存资源。

- GetMemoryNUMANodes 函数: 获取节点上所有内存NUMA节点的列表。

- Allocate 函数: 为指定容器分配内存资源。

- RemoveContainer 函数: 从内存管理器中移除一个容器，并释放其占用的内存资源。

- State 函数: 获取当前内存管理器的状态。

- GetPodTopologyHints 函数: 获取Pod的拓扑提示，用于内存分配。

- GetTopologyHints 函数: 获取给定容器和Pod的拓扑提示。

- removeStaleState 函数: 移除过时的内存状态信息。

- policyRemoveContainerByRef 函数: 根据引用移除一个容器。

- getTotalMemoryTypeReserved 函数: 获取节点上特定内存类型所保留的总内存。

- validateReservedMemory 函数: 验证保留内存的合法性，并返回有效的保留内存值。

- convertReserved 函数: 根据NodeMemoryManager的配置值，将内存保留转换为低数组。

- getSystemReservedMemory 函数: 获取节点上已保留给系统使用的内存。

- GetAllocatableMemory 函数: 获取可分配的内存资源。

- GetMemory 函数: 获取指定容器的内存使用情况。

- setPodPendingAdmission 函数: 设置Pod为待审批状态，表示内存资源还未被分配。

这些函数共同实现了内存管理器的各个功能，包括资源的管理、分配、释放以及状态的维护等。

