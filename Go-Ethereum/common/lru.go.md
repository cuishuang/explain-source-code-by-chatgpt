# File: common/lru/lru.go

在go-ethereum项目中，common/lru/lru.go文件是一个实现了“最近最少使用”（Least Recently Used）缓存算法的包。LRU缓存通常被用来提高访问速度和节省系统资源，通过丢弃最长时间未被使用的元素来保持缓存大小的限制。

文件中定义了三个结构体：Cache、entry和node。这三个结构体的作用如下：

1. Cache：代表一个LRU缓存，其中可以存储任意的键值对。缓存的大小通过maxEntries参数进行控制，一旦超过了这个限制，最早的条目将被丢弃。

2. entry：代表缓存中的一个键值对，其中包含了键和对应的值。此外，entry还包含双向链表中的前驱和后继指针。

3. node：代表一个双向链表中的节点，其中包含了前驱和后继指针，用于维护缓存中元素的访问顺序。

以下是一些重要的函数和它们的作用：

1. NewCache(maxEntries int) *Cache：通过传入的maxEntries参数，创建并返回一个具有指定最大容量的新缓存。

2. Add(key interface{}, value interface{})：向缓存中添加一个键值对，如果缓存的大小超过了最大容量，将删除最久未使用的条目。

3. Contains(key interface{}) bool：检查缓存中是否包含具有给定键的条目。

4. Get(key interface{}) (value interface{}, ok bool)：根据键获取缓存中的值，并返回该值以及一个布尔值，表示该键是否存在于缓存中。

5. Len() int：获取缓存中当前的条目数。

6. Peek(key interface{}) (value interface{}, ok bool)：根据键获取缓存中的值，但不会更新该键的在缓存中的访问顺序。

7. Purge()：从缓存中删除所有的键值对。

8. Remove(key interface{})：根据键从缓存中删除对应的条目。

9. Keys() []interface{}：返回一个包含缓存中所有键的切片。

这些函数使得可以方便地使用LRU缓存来存储和管理数据，提高查询性能，并控制缓存的大小。

