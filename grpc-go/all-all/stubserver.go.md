# File: grpc-go/internal/stubserver/stubserver.go

`stubserver.go`文件是gRPC框架中的一个文件，主要用于实现gRPC服务器端的Stub Server功能。

1. `StubServer`结构体是Stub Server的核心结构体，它包含了Stub Server的相关信息和状态。它有以下字段：
   - `addr string`：Stub Server的监听地址
   - `opts []grpc.ServerOption`：gRPC服务器的选项
   - `hinthandler csrf.healthCheckServiceServer`：health check服务的处理器，用于健康检查
   - `handlers []grpc.StreamHandler`：一组处理RPC流的处理器
   - `status int`：服务器状态（0表示未启动，1表示正在运行，-1表示已停止）

2. `registerServiceServerOption`结构体是一个辅助结构体，用于注册gRPC服务端的选项。它有以下字段：
   - `serviceName string`：要注册的服务名
   - `impl interface{}`：服务的实现对象

3. `EmptyCall`、`UnaryCall`和`FullDuplexCall`等是一组用于处理gRPC调用的函数。它们分别对应不同类型的调用方式，并处理请求和响应。

4. `Start`函数用于启动Stub Server，并监听指定的地址。

5. `RegisterServiceServerOption`是一个辅助函数，用于注册服务及其实现对象到Stub Server中。

6. `StartServer`函数用于启动gRPC服务器。

7. `StartClient`函数用于启动gRPC客户端。

8. `NewServiceConfig`函数用于创建一个新的服务配置。

9. `waitForReady`函数用于等待RPC调用准备就绪。

10. `Stop`函数用于停止Stub Server。

11. `parseCfg`函数用于解析服务配置文件。

12. `StartTestService`函数用于启动测试服务。

综上所述，`stubserver.go`文件是gRPC框架中实现Stub Server功能的关键文件，它提供了一系列函数和结构体，用于注册服务、处理RPC调用、启动服务器等操作。

