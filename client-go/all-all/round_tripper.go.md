# File: client-go/discovery/cached/disk/round_tripper.go

在client-go的cached/disk/round_tripper.go文件中，有几个结构体和函数用于实现对Discovery API的缓存，并将其存储在磁盘上。

1. cacheRoundTripper结构体：实现了http.RoundTripper接口，并用于在请求过程中对请求进行缓存。
   
2. sumDiskCache结构体：用于维护磁盘缓存的摘要，其中包括缓存的大小，过期时间等信息。

3. newCacheRoundTripper函数：用于初始化并返回一个新的cacheRoundTripper结构体。

4. RoundTrip函数：实现了http.RoundTripper接口的RoundTrip方法，用于执行HTTP请求并返回响应。在这个函数中，会首先检查是否有缓存存在，如果有则从缓存中获取响应，否则会执行实际的HTTP请求，并将响应缓存到磁盘上。

5. CancelRequest函数：实现了http.RoundTripper接口的CancelRequest方法，用于取消HTTP请求。

6. WrappedRoundTripper函数：用于在原始的RoundTripper上包装一个新的RoundTripper，并返回一个新的cacheRoundTripper结构体。

7. Get函数：用于从缓存中获取特定键的值。

8. Set函数：用于将一个键值对写入缓存。

9. Delete函数：用于从缓存中删除特定键的值。

10. sanitize函数：用于对URL进行清理和归一化处理。

总的来说，这个文件实现了一个能够在本地磁盘上缓存Discovery API响应的功能。通过使用cacheRoundTripper结构体，可以将HTTP请求的响应存储在磁盘上，并在后续的请求中直接从磁盘缓存中获取响应，以提高性能和减少对API服务器的请求。

