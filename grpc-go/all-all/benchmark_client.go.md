# File: grpc-go/benchmark/worker/benchmark_client.go

benchmark_client.go这个文件是用于实现gRPC的benchmark客户端的。

首先，文件中的`caFile`、`serverHostOverride`、`serverHostPort`等变量用于指定gRPC连接的相关参数。其中：
- `caFile`是指定服务器端证书的根CA文件的路径。
- `serverHostOverride`是指定用于在TLS握手时检查服务器证书中的CommonName字段。
- `serverHostPort`是指定服务器地址和端口号。

文件中定义了两个结构体`lockingHistogram`和`benchmarkClient`。

- `lockingHistogram`是一个带有锁的直方图结构体。它用于记录并统计请求的延迟情况。
- `benchmarkClient`是benchmark的客户端结构体，它存储了所有benchmark所需的状态和参数。

接下来有一些函数，每个函数都有不同的作用：

- `add`函数用于将两个直方图合并。
- `swap`函数用于交换两个直方图的数据。
- `mergeInto`函数用于合并两个benchmark客户端的状态信息。
- `printClientConfig`函数用于打印benchmark客户端的配置信息。
- `setupClientEnv`函数用于设置客户端的环境，例如TLS证书、连接参数等。
- `createConns`函数用于创建和初始化一组gRPC连接。
- `performRPCs`函数用于执行benchmark中定义的RPC调用。
- `startBenchmarkClient`函数用于启动benchmark客户端。
- `unaryLoop`函数用于在benchmark过程中执行单一(非流式)RPC调用。
- `streamingLoop`函数用于在benchmark过程中执行流式RPC调用。
- `poissonUnary`函数用于生成泊松分布的非流式RPC调用。
- `poissonStreaming`函数用于生成泊松分布的流式RPC调用。
- `getStats`函数用于获取benchmark客户端的统计数据。
- `shutdown`函数用于将benchmark客户端优雅地关闭并输出最终的统计信息。

这些函数共同实现了benchmark客户端的初始化、配置、连接、RPC调用和性能统计等功能。

