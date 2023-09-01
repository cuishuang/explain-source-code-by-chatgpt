# File: client-go/tools/cache/thread_safe_store.go

在client-go项目中，thread_safe_store.go文件实现了一个线程安全的、可以进行索引操作的存储结构，主要用于缓存API对象。

ThreadSafeStore是一个线程安全的存储结构体，处理并发访问的情况。storeIndex是一个索引结构体，用于保存键值对及其对应的索引。threadSafeMap则是一个线程安全的映射，提供了对存储和索引的基本操作方法。

- reset：重置存储结构，清空所有数据和索引。
- getKeysFromIndex：根据索引获取指定键的集合。
- getKeysByIndex：根据索引值获取所有的键集合。
- getIndexValues：根据索引名称获取指定键的索引值集合。
- addIndexers：向存储结构中添加索引器。
- updateIndices：更新存储结构中的索引。
- addKeyToIndex：将键添加到指定的索引中。
- deleteKeyFromIndex：从指定的索引中删除键。
- Add：向存储结构中添加键值对。
- Update：更新存储结构中的键值对。
- Delete：从存储结构中删除指定的键值对。
- Get：根据键获取指定的值。
- List：返回存储结构中的所有值。
- ListKeys：返回存储结构中的所有键。
- Replace：用新的键值对覆盖原有的键值对。
- Index：为存储结构添加索引并返回索引值。
- ByIndex：根据索引值返回存储结构中的值。
- IndexKeys：根据索引名称返回存储结构中的键集合。
- ListIndexFuncValues：根据索引名称和过滤函数返回符合条件的值集合。
- GetIndexers：返回存储结构中的所有索引器。
- AddIndexers：向存储结构中添加多个索引器。
- Resync：重新同步存储结构中的数据。
- NewThreadSafeStore：创建一个新的线程安全的存储结构。

通过这些方法和结构体，可以方便地在Kubernetes中实现对API对象的缓存和索引操作，提高性能和效率。

