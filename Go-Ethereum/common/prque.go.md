# File: common/prque/prque.go

在go-ethereum项目中，common/prque/prque.go文件是一个优先级队列（Priority Queue）的实现。优先级队列是一种数据结构，其中每个元素都有一个与之关联的优先级。

Prque文件中定义了三个结构体：entry、Prque、itemHeap。

- entry结构体是一个队列中的条目，包括了一个优先级和一个项目。
- Prque结构体是优先级队列的主要结构，它有一个entry组成的slice和一个map，用于快速查找和删除条目。
- itemHeap结构体是一个被Prque结构体使用的优先级队列的堆表示。

以下是Prque结构体中定义的几个方法及其作用：

- New：创建一个新的优先级队列。
- Push：将一个条目添加到队列中。
- Peek：查看优先级最高的条目，不删除它。
- Pop：返回并删除优先级最高的条目。
- PopItem：返回并删除一个指定的条目。
- Remove：删除一个指定的条目。
- Empty：检查队列是否为空。
- Size：返回队列中的条目数。
- Reset：将队列重置为空。

这些方法能够让用户使用优先级队列，根据优先级来管理和操作队列中的条目。例如，可以使用Push方法将条目添加到队列中，使用Pop方法删除并返回优先级最高的条目，使用Remove方法删除指定的条目等等。这样，可以方便地实现各种基于优先级的功能和算法。

