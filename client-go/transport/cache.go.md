# File: client-go/transport/cache.go

在client-go的transport包中的cache.go文件主要定义了一些缓存数据的结构和方法，并提供了一些与缓存交互的函数。

1. DialerStopCh变量：是一个chan struct{}类型的通道，用于传递信号以停止Dialer的监听。

2. tlsCache变量：是一个sync.Map类型的缓存，用于存储TLS配置和HTTP传输层的缓存。

3. tlsTransportCache结构体：定义了一个TLS传输缓存的结构体，包含了一个读写锁和一个缓存，用于存储TLS传输层缓存的数据。

4. tlsCacheKey结构体：定义了一个TLS缓存的键结构，用来标识缓存数据。

5. String函数：是tlsCacheKey结构体的方法，返回一个描述唯一键的字符串。

6. get函数：是tlsTransportCache结构体的方法，根据tlsCacheKey获取对应的缓存项。

7. tlsConfigKey函数：是一个用于生成TLS配置键的函数，可根据给定的TLS配置生成唯一的键。

这些结构体和函数主要是为了方便在client-go中复用和管理TLS配置和HTTP传输层的缓存。例如，当需要创建一个TLS配置时，可以先检查缓存中是否已有对应的配置，如果有则直接使用缓存中的配置，避免重复创建。同时，这些缓存项也可以在传输层的HTTP请求中被复用。

