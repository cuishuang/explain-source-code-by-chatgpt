# File: tools/gopls/internal/lsp/lru/lru.go

在Golang的gopls源码中，tools/gopls/internal/lsp/lru/lru.go文件是实现LRU（Least Recently Used）缓存的核心文件。LRU缓存是一种常见的缓存算法，它根据最近使用的顺序来淘汰最不常用的缓存条目。

在lru.go文件中，定义了三个结构体：Cache、entry和queue。

1. Cache结构体：代表了一个LRU缓存，包含了一个字典映射（entries），映射的键是缓存的key，值是一个双向链表的节点，通过这个节点可以获取对应的缓存值。
   
2. entry结构体：代表了缓存的节点，其中包含了缓存的key和value。

3. queue结构体：是一个双向链表，用于按照最近使用的顺序排列缓存的节点。

接下来，我们来看一下lru.go文件中的一些主要函数的作用：

1. New(capacity int) *Cache：是一个构造函数，用于创建一个指定容量的LRU缓存实例。
   
2. Get(key string) (value interface{}, ok bool)：根据键获取缓存值。如果存在缓存，则将其移动到链表头部，并返回value和ok=true。
  
3. Set(key string, value interface{})：设置缓存值。如果缓存已存在，则更新其值，并将其移动到链表头部；如果缓存不存在，则在链表头部插入一个新节点。当缓存容量超限时，会删除链表尾部的节点。

4. Len() int：获取缓存的大小，即缓存中存储的条目个数。

5. Less(i, j int) bool：用于排序的辅助函数，比较两个节点的访问时间。

6. Swap(i, j int)：用于排序的辅助函数，交换两个节点的位置。

7. Push(x interface{})：用于排序的辅助函数，向队列中添加一个节点。

8. Pop() interface{}：用于排序的辅助函数，从队列中弹出一个节点。

这些函数配合使用，实现了对LRU缓存的管理和淘汰策略。通过LRU缓存的实现，可以有效地提升程序的性能，减少不必要的计算和IO操作。

