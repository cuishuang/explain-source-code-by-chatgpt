# File: client-go/transport/transport.go

client-go/transport/transport.go文件是Kubernetes client-go库中定义了与API服务器进行通信的传输层功能的核心文件。它包含了一些结构体和函数，用于构建与API服务器进行通信的HTTP请求。

- WrapperFunc: 这是一个函数类型，用于包装和修改每个请求发送之前和接收之后的传输过程。可以使用WrapperFunc在传输层上实施一些额外的功能，例如添加认证信息、修改HTTP头等。

- contextCanceller: 这是一个结构体，用于实现取消请求的功能。它包含一个通道用于接收取消请求的信号，并将该信号传播给底层的HTTP传输。

- certificateCacheEntry: 这是一个结构体，用于缓存证书的信息。它包含证书、密钥、证书过期时间等信息，并提供一些方法用于检查证书是否过期。

- New: 这是一个函数，用于创建一个新的Transport对象。它接收一组参数，用于配置传输层的行为，例如TLS配置、代理设置等。

- isValidHolders: 这是一个函数，用于检查给定的TLS证书和密钥是否有效。

- TLSConfigFor: 这是一个函数，用于根据传输层的配置创建TLS配置。它针对不同的TLS配置选项生成符合要求的TLS配置。

- loadTLSFiles: 这是一个函数，用于从文件中加载TLS证书和密钥。

- dataFromSliceOrFile: 这是一个函数，用于将数据从切片或文件中读取并返回。

- rootCertPool: 这是一个函数，用于创建根证书池，并从文件或切片中加载根证书。

- createErrorParsingCAData: 这是一个函数，用于创建一个解析证书数据时出错的错误信息。

- Wrappers: 这是一个函数，用于将多个WrapperFunc函数组合成一个WrapperFunc函数链。

- ContextCanceller: 这是一个函数，用于将取消请求的功能包装在WrapperFunc函数中。

- RoundTrip: 这是一个函数，用于执行HTTP请求并返回响应。它将处理整个请求的生命周期，包括添加认证信息、处理重定向、取消请求等。

- tryCancelRequest: 这是一个函数，用于尝试取消请求。它在接收到取消请求的信号后，会调用底层的HTTP传输的CancelRequest方法来取消请求。

- isStale: 这是一个函数，用于检查缓存的证书是否过期。

- newCertificateCacheEntry: 这是一个函数，用于创建一个新的证书缓存条目。

- cachingCertificateLoader: 这是一个函数，用于从缓存中加载证书。如果证书已过期或不存在，则将调用loadTLSFiles函数从文件中加载证书。

