# File: runc/libcontainer/intelrdt/stats.go

在runc项目中，runc/libcontainer/intelrdt/stats.go文件的作用是实现对Intel Resource Director Technology（RDT）统计数据的处理和管理。

该文件定义了多个结构体来表示不同类型的RDT统计信息。这些结构体包括：

1. L3CacheInfo：表示L3缓存的统计信息，包括缓存大小、缓存占用、缓存限制等。
2. MemBwInfo：表示内存带宽的统计信息，包括读/写带宽、带宽限制等。
3. MBMNumaNodeStats：表示MBM（Memory Bandwidth Monitoring） NUMA节点的统计信息，包括每个NUMA节点的内存带宽、缓存带宽等。
4. CMTNumaNodeStats：表示CMT（Cache Monitoring Technology） NUMA节点的统计信息，包括每个NUMA节点的缓存占用、缓存限制等。
5. Stats：表示所有RDT统计信息的集合。

在stats.go中，还定义了一些用于创建新的RDT统计信息的函数，这些函数被称为newStats函数。这些函数包括：

1. newL3CacheInfo：用于创建一个新的L3CacheInfo结构体，初始化其字段并返回该结构体的指针。
2. newMemBwInfo：用于创建一个新的MemBwInfo结构体，初始化其字段并返回该结构体的指针。
3. newMBMNumaNodeStats：用于创建一个新的MBMNumaNodeStats结构体，初始化其字段并返回该结构体的指针。
4. newCMTNumaNodeStats：用于创建一个新的CMTNumaNodeStats结构体，初始化其字段并返回该结构体的指针。
5. newStats：用于创建一个新的Stats结构体，其中会调用上述函数创建并初始化各个子结构体，并返回该结构体的指针。

这些函数的作用是方便在代码中创建和初始化RDT统计信息的结构体，使得对RDT统计数据的处理和管理更加方便和灵活。

