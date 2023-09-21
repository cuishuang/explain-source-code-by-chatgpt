# File: grpc-go/internal/transport/http2_server.go

grpc-go/internal/transport/http2_server.go是gRPC-Go项目中的一个文件，主要负责实现HTTP/2服务器的功能。HTTP/2是gRPC的底层传输协议，该文件实现了与HTTP/2传输相关的逻辑。

以下是每个变量和结构体的作用：

1. ErrIllegalHeaderWrite：表示尝试写入非法首部的错误。
2. ErrHeaderListSizeLimitViolation：表示请求/响应头列表大小超出限制的错误。
3. serverConnectionCounter：用于计算服务器上的活动连接数。
4. goAwayPing：用于存储需要发送的带有goAway标志的ping。

5. http2Server：代表HTTP/2服务器的结构体，其中包含服务器的配置信息和状态等。
6. connectionKey：HTTP/2连接的键，用于唯一标识每个连接。

以下是每个函数的作用：

1. NewServerTransport：创建一个新的服务器传输实例。
2. operateHeaders：对接收到的请求进行预处理，处理请求头、处理帧类型等。
3. HandleStreams：处理连接上的流，根据接收到的HTTP/2帧类型执行相应的操作。
4. getStream：根据流ID获取对应的流对象。
5. adjustWindow：根据流的窗口大小调整全局窗口大小。
6. updateWindow：更新流的窗口大小。
7. updateFlowControl：更新流的流量控制机制。
8. handleData：处理接收到的数据帧。
9. handleRSTStream：处理接收到的RST_STREAM帧，关闭对应的流。
10. handleSettings：处理接收到的设置帧，根据接收到的设置参数更新服务器的配置。
11. handlePing：处理接收到的ping帧，发送相应的ack帧。
12. handleWindowUpdate：处理接收到的窗口更新帧，更新流的窗口大小。
13. appendHeaderFieldsFromMD：将元数据的字段添加到HTTP/2请求头。
14. checkForHeaderListSize：检查请求/响应头列表大小是否超出限制。
15. streamContextErr：获取与流相关的上下文和错误。
16. WriteHeader：写入响应头。
17. setResetPingStrikes：设置resetPingStrikes标志。
18. writeHeaderLocked：为流写入响应头。
19. WriteStatus：写入响应状态。
20. Write：写入响应数据。
21. keepalive：启动或停止保持活动状态的goroutine。
22. Close：关闭传输。
23. deleteStream：删除指定流。
24. finishStream：完成流的处理。
25. closeStream：关闭流并释放相应资源。
26. RemoteAddr：获取对端地址。
27. Drain：将服务器连接上的所有流标记为耗尽状态。
28. outgoingGoAwayHandler：处理传出的goAway帧。
29. ChannelzMetric：返回与服务器传输相关的度量值。
30. IncrMsgSent：增加已发送的消息计数。
31. IncrMsgRecv：增加已接收的消息计数。
32. getOutFlowWindow：获取输出流的流控窗口大小。
33. getPeer：获取对端的标识符。
34. getJitter：获取抖动值。
35. GetConnection：获取与传输关联的连接。
36. setConnection：设置与传输关联的连接。

