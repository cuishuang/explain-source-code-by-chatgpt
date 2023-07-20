# File: light/nodeset.go

在go-ethereum项目中，light/nodeset.go文件的作用是实现一个用于存储和管理轻量级节点的集合。该文件定义了NodeSet和NodeList两个结构体，并提供一些相关操作的函数。

首先介绍NodeSet结构体，它是一个轻量级节点集合的主要数据结构。NodeSet使用一个map来存储节点的信息，每个节点都有一个唯一的键值和相关数据。使用键值可以快速访问和操作节点集合中的节点。NodeSet结构体的定义如下：

```go
type NodeSet struct {
	nodes       map[string][]byte
	keyCount    int // 节点数量
	dataSize    int // 数据大小
	memoryLimit int // 内存限制
}
```

而NodeList结构体则是用来表示一个节点列表的数据结构。NodeList包含了NodeSet结构体的引用，以及一个按照节点插入顺序的节点列表（切片）。NodeList结构体的定义如下：

```go
type NodeList struct {
	set   *NodeSet
	nodes []string // 按照插入顺序存储节点键值
}
```

下面介绍一些关键的函数功能：

- `NewNodeSet(memoryLimit int) *NodeSet`：创建一个新的NodeSet对象，并指定内存限制。
- `Put(key string, value []byte) bool`：将一个节点以指定的键值存储到NodeSet中。成功存储返回true，否则返回false。
- `Delete(key string) bool`：根据键值从NodeSet中删除一个节点。成功删除返回true，否则返回false。
- `Get(key string) (value []byte, exists bool)`：根据键值获取一个节点的数据，如果节点存在则返回节点数据和true，否则返回nil和false。
- `Has(key string) bool`：判断NodeSet中是否存在某个节点，存在返回true，否则返回false。
- `KeyCount() int`：返回NodeSet中节点的数量。
- `DataSize() int`：返回NodeSet中所有节点数据的总大小。
- `NodeList() *NodeList`：返回一个按照插入顺序的节点列表对象。
- `Store(path string) error`：将NodeSet的数据存储到指定的路径下的文件中。
- `NodeSet(path string, memoryLimit int) (*NodeSet, error)`：从指定的路径读取数据，创建一个新的NodeSet对象，并指定内存限制。

这些函数提供了NodeSet结构体的常见操作，例如添加、删除、获取节点，判断节点是否存在以及存储和加载节点数据等。NodeList结构体则用于按照插入顺序记录节点的列表信息。通过使用这些结构体和函数，可以方便地管理和操作轻量级节点集合。

