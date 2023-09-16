# File: istio/pkg/cache/ttlCache.go

在Istio项目中，istio/pkg/cache/ttlCache.go文件定义了一个TTL（Time-To-Live）缓存实现，用于存储具有过期时间的键值对。

ttlWrapper结构体是一个包装器，用于将键值对添加到缓存中，并设置各自的过期时间。ttlCache结构体是具体的缓存实现，它通过保存entry（键值对）的映射关系和过期时间戳来进行缓存管理。EvictionCallback结构体是用于在缓存项被删除时执行回调函数的容器。

以下是这些结构体的作用：

- ttlWrapper: 一个包装器，用于将键值对添加到缓存中，并设置过期时间。
- ttlCache: 缓存的主要实现，管理缓存项的存储和过期时间戳。
- entry: 缓存项的结构体，包含键、值和过期时间戳。
- EvictionCallback: 用于在缓存项被删除时执行的回调函数。

以下是这些函数的作用：

- NewTTL: 创建一个新的TTL缓存实例。
- NewTTLWithCallback: 创建一个新的TTL缓存实例，并指定在缓存项被删除时执行的回调函数。
- evicter: 定期删除过期的缓存项。
- evictExpired: 删除过期的缓存项。
- EvictExpired: 定期调用evictExpired函数，删除过期的缓存项。
- Set: 向缓存中添加一个键值对。
- SetWithExpiration: 向缓存中添加一个键值对，并指定过期时间。
- Get: 根据键从缓存中获取对应的值。
- Remove: 根据键从缓存中删除对应的缓存项。
- RemoveAll: 清空缓存，删除所有缓存项。
- Stats: 返回缓存的统计信息，如缓存项的数量、命中率等。

ttlCache的作用是提供一个可自动过期的缓存机制，用于临时存储和访问键值对，允许用户根据需要设定过期时间，并提供回调函数以响应缓存项的删除事件。这对于需要缓存临时数据、缓解服务负载或者提高访问效率等场景非常有用。

