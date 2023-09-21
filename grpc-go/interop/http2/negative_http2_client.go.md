# File: grpc-go/interop/http2/negative_http2_client.go

文件`negative_http2_client.go`位于`grpc-go/interop/http2/`目录下，用于实现一个负面场景的HTTP/2客户端。该客户端通过与gRPC服务器进行通信来测试gRPC与HTTP/2相关功能的兼容性，特别是在一些异常情况下的行为。

下面是对变量和函数的详细介绍：

变量：
- `serverHost`：gRPC服务器的主机名/ip地址。
- `serverPort`：gRPC服务器的端口号。
- `testCase`：测试用例，用于指定需要运行的负面场景。
- `largeReqSize`：大请求的大小，用于指定要发送的大请求的大小。
- `largeRespSize`：大响应的大小，用于指定接收的大响应的大小。
- `logger`：用于日志记录的logger对象。

函数：
- `largeSimpleRequest`：发送一个大的简单请求，该请求的大小由`largeReqSize`指定。
- `goaway`：发送一个GOAWAY帧，通知服务器不再接受新的流。
- `rstAfterHeader`：发送一个RST_STREAM帧，在接收到响应头后立即中止流。
- `rstDuringData`：发送一个RST_STREAM帧，在接收响应数据时立即中止流。
- `rstAfterData`：发送一个RST_STREAM帧，在接收完整个响应后立即中止流。
- `ping`：发送一个PING帧，用于检测与服务器的连接状态。
- `maxStreams`：尝试打开比gRPC服务器最大可接受流数目更多的流。
- `main`：程序入口函数，根据`testCase`的值来选择相应的功能进行测试。

通过设置不同的`testCase`值，可以测试不同的负面场景。这些测试用例包括发送大的简单请求、发送GOAWAY帧、发送RST_STREAM帧等，以验证gRPC的服务器是否能正确处理这些异常情况，并且能够根据HTTP/2协议规范进行适当的处理。

日志记录器`logger`用于记录程序的运行日志，以便进行故障排查和问题定位。通过阅读该日志，可以了解测试的具体过程和运行结果。

