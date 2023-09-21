# File: grpc-go/benchmark/benchmark.go

grpc-go/benchmark/benchmark.go 文件的作用是实现了基于 gRPC 的性能测试。

logger 变量用于记录测试日志，分别代表普通日志（infoLogger）、错误日志（errorLogger）和详细日志（verboseLogger）。

testServer 结构体定义了测试服务器，包含了一个 gRPC 服务器和一些性能测试相关的参数和数据。

byteBufServer 结构体定义了基于 byte buffer 的测试服务器，用于性能测试。

ServerInfo 结构体用于传递测试服务器的相关信息，包括服务器地址和端口。

setPayload 函数用于设置有效载荷大小。

NewPayload 函数用于生成指定大小的有效载荷数据。

UnaryCall 函数实现了一元调用的性能测试。

StreamingCall 函数实现了流式调用的性能测试。

UnconstrainedStreamingCall 函数实现了无限制流式调用的性能测试。

StartServer 函数用于启动测试服务器。

DoUnaryCall 函数执行一元调用测试。

DoStreamingRoundTrip 函数执行流式调用测试。

DoStreamingRoundTripPreloaded 函数执行预加载的流式调用测试。

DoByteBufStreamingRoundTrip 函数执行基于 byte buffer 的流式调用测试。

NewClientConn 函数创建一个客户端连接。

NewClientConnWithContext 函数创建一个带上下文的客户端连接。

