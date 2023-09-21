# File: grpc-go/test/rawConnWrapper.go

在grpc-go项目中，`rawConnWrapper.go`文件的作用是提供了对原始TCP连接进行包装的功能，以在gRPC传输层的底层字节流与拆分为帧的消息之间建立桥梁。

`listenerWrapper`结构体是对`net.Listener`接口的包装，用于在接受新连接时将其包装为带有`rawConnWrapper`的`Conn`对象。

`dialerWrapper`结构体是对`net.Dialer`接口的包装，用于将新连接包装为带有`rawConnWrapper`的`Conn`对象。

`rawConnWrapper`结构体是对`grpc.RawConn`接口的实现，用于在底层TCP连接上发送和接收字节流，并将其封装为gRPC帧。

下面是对一些重要函数的介绍：

- `listenWithConnControl`: 这个函数在监听器上启用连接控制，一旦有新连接建立，它会使用`rawConnWrapper`为连接创建一个新的`Conn`对象。这样，通过调用`getLastConn`方法可以得到这个新连接。

- `Accept`: 这个方法会阻塞等待新连接的到来，并返回一个带有`rawConnWrapper`的`grpc.ClientConn`对象。

- `getLastConn`: 返回最近创建的`Conn`对象。

- `dialer`: 创建一个带有`rawConnWrapper`的`grpc.ClientConn`对象，并使用传入的地址进行连接。

- `getRawConnWrapper`: 返回原始的TCP连接的`rawConnWrapper`对象。

- `newRawConnWrapperFromConn`: 从传入的连接创建一个新的`rawConnWrapper`对象。

- `Close`: 关闭连接并释放所有相关资源。

- `encodeHeaderField`: 编码gRPC消息头字段。

- `encodeRawHeader`: 编码原始的gRPC消息头。

- `encodeHeader`: 编码gRPC消息的头部。

- `writeHeaders`: 在连接上写入gRPC消息头。

- `writeRSTStream`: 写入用于重置流的消息帧。

- `writeGoAway`: 写入用于关闭连接的消息帧。

- `writeRawFrame`: 在连接上写入原始的gRPC消息帧。

