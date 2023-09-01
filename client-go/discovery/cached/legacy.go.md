# File: client-go/discovery/cached/legacy.go

client-go 是 Kubernetes 官方提供的 Go 语言客户端库，用于与 Kubernetes API 进行交互。在 client-go 中的 `discovery/cached/legacy.go` 文件提供了一些用于缓存发现（discovery）结果的功能。

该文件主要定义了两个结构体，一个是 `CachedDiscoveryClient`，另一个是 `MemCacheClient`，其中包含以下几个重要的函数和变量：

1. 函数 `NewMemCacheClient`：这个函数创建了一个新的 `MemCacheClient` 对象，并返回它。`MemCacheClient` 是 `CachedDiscoveryClient` 的一个内部实现，用于维护缓存的资源发现结果。

2. 函数 `CachedDiscoveryClient.ServerGroups`：该函数实现了 `discovery.DiscoveryClient` 接口中的方法，用于获取 API Server 支持的 Group 列表。它首先检查缓存是否有结果，如果有则直接返回；否则，它会调用真实的 `discovery.DiscoveryClient.ServerGroups()` 方法获取结果，并将结果保存到缓存中。

3. 函数 `CachedDiscoveryClient.ServerResourcesForGroupVersion`：该函数实现了 `discovery.DiscoveryClient` 接口中的方法，用于获取指定 Group 和 Version 的资源列表。它也会首先检查缓存是否有结果，如果有则直接返回；否则，它会调用真实的 `discovery.DiscoveryClient.ServerResourcesForGroupVersion()` 方法获取结果，并将结果保存到缓存中。

4. 变量 `ErrCacheNotFound`：这个变量主要用于表示在缓存中没有找到相应的缓存结果时返回的错误。它是 `errors.New("discovery cache not found")` 的一个实例，用于标识缓存没有找到的情况。

上述的 `NewMemCacheClient` 函数会创建一个 `MemCacheClient` 对象，该对象实现了 `discovery.ServerResourcesInterface` 接口，用于对资源进行缓存。在创建 `MemCacheClient` 对象时，它会创建一个内存缓存（`MemCache`）用于存储缓存数据。

这些函数和变量的作用是为了提高资源发现的性能。通过使用缓存，可以避免频繁地向 API Server 发送查询请求，从而提高应用程序的性能和效率。

