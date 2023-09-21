# File: grpc-go/internal/transport/http_util.go

在grpc-go项目中，`grpc-go/internal/transport/http_util.go`文件提供了HTTP/2的实现细节，包括客户端建立连接、通信数据的编码和解码，以及错误处理等功能。下面逐一介绍相关的变量和函数的作用：

### 变量：
1. `clientPreface`：HTTP/2客户端连接握手阶段所需的client preface。
2. `http2ErrConvTab`：HTTP/2错误码与gRPC内部错误码的映射表。
3. `HTTPStatusConvTab`：HTTP状态码与gRPC状态码的映射表。
4. `writeBufferPoolMap`：用于缓存写入缓冲区的数据的池子（pool）。
5. `writeBufferMutex`：保护写缓冲区池的互斥锁。

### 结构体：
1. `timeoutUnit`：gRPC中的超时时间单位的定义。
2. `bufWriter`：封装了`http.ResponseWriter`接口的缓冲写入器。
3. `ioError`：封装了IO错误及其对应的gRPC错误码。
4. `framer`：对HTTP/2的帧进行封装和解包的对象。

### 函数：
1. `isReservedHeader`：检查给定的HTTP头字段是否是gRPC保留的标准头字段。
2. `isWhitelistedHeader`：检查给定的HTTP头字段是否允许在gRPC协议中使用。
3. `encodeBinHeader`：将二进制数据编码为字符串，并写入HTTP头字段。
4. `decodeBinHeader`：从HTTP头字段中解码二进制数据。
5. `encodeMetadataHeader`：将gRPC元数据编码为HTTP头字段的格式。
6. `decodeMetadataHeader`：从HTTP头字段中解码gRPC元数据。
7. `decodeGRPCStatusDetails`：解码gRPC状态详情。
8. `timeoutUnitToDuration`：将gRPC中的超时时间单位转换为Go中的`time.Duration`。
9. `decodeTimeout`：从HTTP头字段中解码gRPC调用的超时时间。
10. `encodeGrpcMessage`：将gRPC消息进行编码，包括消息头和消息体。
11. `encodeGrpcMessageUnchecked`：和`encodeGrpcMessage`类似，但不进行输入验证和错误检查。
12. `decodeGrpcMessage`：解码gRPC消息的头部和消息体。
13. `decodeGrpcMessageUnchecked`：和`decodeGrpcMessage`类似，但不进行输入验证和错误检查。
14. `newBufWriter`：创建一个`bufWriter`对象。
15. `Write`：将数据写入`bufWriter`中。
16. `Flush`：将`bufWriter`中的数据刷新到`http.ResponseWriter`处理器中。
17. `flushKeepBuffer`：类似于`Flush`函数，但保持写入的缓冲区。
18. `Unwrap`：返回底层的`http.ResponseWriter`对象。
19. `isIOError`：检查给定的错误是否是IO错误。
20. `toIOError`：将给定的错误信息转化为符合gRPC标准的IO错误。
21. `newFramer`：创建一个新的`framer`对象，用于对HTTP/2的帧进行封装和解包操作。
22. `getWriteBufferPool`：根据给定的大小返回对应的写缓冲区池。
23. `parseDialTarget`：解析gRPC调用的目标地址。

