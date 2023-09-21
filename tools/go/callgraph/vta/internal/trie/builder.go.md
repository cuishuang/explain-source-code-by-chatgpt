# File: tools/go/callgraph/vta/internal/trie/builder.go

文件`builder.go`的作用是实现构建trie（前缀树）的功能。该文件定义了用于构建trie的`Collision`、`Builder`和`MutMap`等结构体，以及一系列用于操作trie的函数。

`Collision`结构体表示trie中键的碰撞（冲突）情况，它包含两个键和对应的值。

`Builder`结构体用于构建trie，它包含一个`MutMap`，用于存储trie的节点和键值。

`MutMap`结构体表示一个可变的trie映射，它由多个叶子节点和分支节点组成。

下面是给出的函数的功能介绍：

- `TakeLhs`：获取给定键的左子树。
- `TakeRhs`：获取给定键的右子树。
- `NewBuilder`：创建一个新的`Builder`实例。
- `Scope`：在给定范围内创建一个新的`Builder`实例。
- `Rescope`：在给定范围内重新创建一个`Builder`实例。
- `Empty`：检查`Builder`是否为空。
- `InsertWith`：使用给定的合并函数插入键值对到`Builder`中。
- `Insert`：插入键值对到`Builder`中。
- `Update`：更新`Builder`中存在的键对应的值。
- `MergeWith`：将另一个`Builder`与当前的`Builder`进行合并。
- `Merge`：将另一个映射与当前`Builder`进行合并。
- `Clone`：克隆当前的`Builder`实例。
- `clone`：克隆给定的`Builder`实例。
- `Remove`：从`Builder`中删除给定键。
- `Intersect`：返回一个新的`Builder`，其中包含两个`Builder`的交集键值对。
- `IntersectWith`：将另一个`Builder`与当前的`Builder`进行交集操作。
- `MutEmpty`：检查`MutMap`是否为空。
- `Create`：创建一个新的空`MutMap`实例。
- `create`：创建一个非空的`MutMap`实例。
- `mkLeaf`：创建一个叶子节点。
- `mkBranch`：创建一个分支节点。
- `join`：将两个叶子节点合并为一个分支节点。
- `collide`：将两个叶子节点合并为一个碰撞节点。
- `insert`：将键值对插入到`MutMap`中。
- `merge`：将两个`MutMap`合并为一个。
- `remove`：从`MutMap`中删除给定键。
- `intersect`：返回两个`MutMap`的交集。
其中，`MutMap`是一个用于构建trie的核心数据结构，它包含叶子节点和分支节点。`Builder`则是对`MutMap`的封装，提供了一系列构建trie的操作方法。

