# File: grpc-go/stream.go

stream.go文件是grpc-go项目中负责处理流式请求和响应的核心文件。

StreamHandler结构体是用于定义流式请求处理函数的接口，其中包含一个ServeStream函数，用于处理请求流并返回响应。

StreamDesc结构体是对流式方法的描述，它包含了请求和响应的类型信息、方法名称以及流式处理函数。

Stream结构体是一个通用的流式方法，它包含了请求和响应的类型信息、方法名称、服务和流式处理函数。

ClientStream结构体是客户端流式方法的抽象表示，它包含了请求和响应的类型信息、方法名称、服务和底层的流式处理。

clientStream结构体是用于客户端流式方法的实现，它实现了Stream接口，包含了客户端流式请求的状态和底层的流。

csAttempt结构体是一个连接状态的封装，用于追踪客户端的重试尝试。

addrConnStream结构体是一个客户端流式请求的封装，包含了底层的流和请求相关的状态和数据。

ServerStream结构体是服务端流式方法的抽象表示，它包含了请求和响应的类型信息、方法名称、服务和底层的流式处理。

serverStream结构体是用于服务端流式方法的实现，它实现了Stream接口，包含了服务端流式请求的状态和底层的流。

NewStream函数用于创建一个新的流式方法。

NewClientStream函数用于创建一个新的客户端流式方法。

newClientStream函数用于创建一个新的客户端流。

newClientStreamWithParams函数用于创建一个带有参数的新的客户端流。

newAttemptLocked函数用于创建一个新的客户端重试尝试。

getTransport函数用于获取流的传输通道。

newStream函数用于创建一个新的流。

commitAttemptLocked函数用于提交客户端重试尝试。

commitAttempt函数用于提交客户端重试尝试。

shouldRetry函数用于判断是否应该重试。

retryLocked函数用于进行重试操作。

Context函数用于获取流的上下文。

withRetry函数用于包装流的上下文，添加重试功能。

Header函数用于获取流的请求头。

Trailer函数用于获取流的响应尾。

replayBufferLocked函数用于回放缓冲区中的请求。

bufferForRetryLocked函数用于在重试之前在缓冲区中缓存请求。

SendMsg函数用于发送请求消息。

RecvMsg函数用于接收响应消息。

CloseSend函数用于关闭发送。

finish函数用于完成流的发送。

sendMsg函数用于发送消息。

recvMsg函数用于接收消息。

newNonRetryClientStream函数用于创建一个不带重试功能的客户端流。

SetHeader函数用于设置流的请求头。

SendHeader函数用于发送流的请求头。

SetTrailer函数用于设置流的响应尾。

MethodFromServerStream函数用于获取服务端流的方法名称。

prepareMsg函数用于准备消息。

