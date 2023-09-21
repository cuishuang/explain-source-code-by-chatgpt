# File: grpc-go/stats/opencensus/server_metrics.go

在grpc-go项目中，grpc-go/stats/opencensus/server_metrics.go文件用于定义和导出用于监视gRPC服务器性能的指标。该文件提供了一组OpenCensus测量指标，可用于跟踪和记录服务器端gRPC操作的各种性能指标。

以下是这些变量的详细介绍：

1. keyServerMethod: 用于标识gRPC服务器方法的标签键。

2. keyServerStatus: 用于标识gRPC服务器状态的标签键。

3. serverReceivedMessagesPerRPC: 统计每个RPC请求接收的消息数。

4. serverReceivedBytesPerRPC: 统计每个RPC请求接收的字节数。

5. serverReceivedCompressedBytesPerRPC: 统计每个RPC请求接收的压缩字节数。

6. serverSentMessagesPerRPC: 统计每个RPC请求发送的消息数。

7. serverSentBytesPerRPC: 统计每个RPC请求发送的字节数。

8. serverSentCompressedBytesPerRPC: 统计每个RPC请求发送的压缩字节数。

9. serverStartedRPCs: 统计服务器端启动的RPC请求数。

10. serverLatency: 统计RPC请求的延迟时间。

11. ServerSentMessagesPerRPCView: 描述了每个RPC请求发送的消息数的视图。

12. ServerReceivedMessagesPerRPCView: 描述了每个RPC请求接收的消息数的视图。

13. ServerSentBytesPerRPCView: 描述了每个RPC请求发送的字节数的视图。

14. ServerSentCompressedMessageBytesPerRPCView: 描述了每个RPC请求发送的压缩字节数的视图。

15. ServerReceivedBytesPerRPCView: 描述了每个RPC请求接收的字节数的视图。

16. ServerReceivedCompressedMessageBytesPerRPCView: 描述了每个RPC请求接收的压缩字节数的视图。

17. ServerStartedRPCsView: 描述了服务器端启动的RPC请求数的视图。

18. ServerCompletedRPCsView: 描述了服务器端完成的RPC请求数的视图。

19. ServerLatencyView: 描述了RPC请求延迟时间的视图。

20. DefaultServerViews: 一组默认的服务器视图，包含上述的视图。

这些指标可以通过OpenCensus的监控和指标导出功能，被导出到不同的监控系统中，如Prometheus、Stackdriver等，以便进行性能分析和故障排查。通过收集和展示这些指标，可以更好地了解gRPC服务器的性能情况，并及时进行问题处理和性能优化。

