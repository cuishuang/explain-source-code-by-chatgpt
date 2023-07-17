# File: pkg/registry/core/service/portallocator/allocator.go

pkg/registry/core/service/portallocator/allocator.go是Kubernetes项目中负责端口分配和管理的文件。它定义了端口分配器的接口和实现。

ErrFull是一个错误变量，表示端口分配器已满，无法分配新的端口。
ErrAllocated是一个错误变量，表示要分配的端口已经被分配给其他资源。
ErrMismatchedNetwork是一个错误变量，表示要分配的端口与所在网络不匹配。

Interface是一个接口类型，定义了端口分配器的方法。它包含了Allocate、AllocateNext、Free、Used等函数的声明，用于分配、释放和查询端口信息。

ErrNotInRange是一个错误变量，表示分配的端口超出了指定的范围。

PortAllocator是一个结构体，实现了Interface接口，用于管理端口的分配和释放。它包含了一个端口映射表，记录了已分配和未分配的端口信息。

Error是一个函数，用于创建一个PortAllocatorError类型的错误。

New是一个函数，创建并返回一个新的端口分配器，可以指定起始和结束端口。

NewInMemory是一个函数，创建并返回一个新的内存端口分配器。

NewFromSnapshot是一个函数，根据给定的端口映射表创建并返回一个新的端口分配器。

Free函数用于释放指定的端口。

Used函数用于查询指定端口是否已被分配。

Allocate函数用于分配一个端口。

AllocateNext函数用于分配下一个可用的端口。

ForEach函数用于迭代端口分配器中的端口和分配情况。

Release函数用于释放端口分配器及其资源。

Has函数用于检查端口分配器是否包含指定的端口。

Snapshot函数用于创建当前端口分配器的快照。

Restore函数用于从给定的端口映射表恢复端口分配器。

Contains函数用于检查给定的端口范围是否与指定的端口范围重叠。

Destroy函数用于销毁端口分配器及其资源。

EnableMetrics函数用于启用端口分配器的度量指标。

calculateRangeOffset函数用于计算IP范围和端口范围的偏移量。

