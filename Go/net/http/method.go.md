# File: method.go

`method.go`是Go语言的标准库`net`包中的一个文件，主要定义了`net`包的中的`Conn`和`PacketConn`两个接口类型的方法。这些方法是`net`包用于网络通信连接和数据传输的核心方法，包括读取和写入数据、设置和获取连接属性以及关闭连接等操作。

具体来说，`Conn`接口定义了基础的网络连接操作，包括读取和写入数据、设置和获取连接属性等操作。`PacketConn`接口则是基于数据报的网络连接的特化接口，它定义了发送和接收数据包、设置和获取连接属性等操作。

`method.go`文件中定义的方法包括：

#### Conn接口

- `Read(b []byte) (n int, err error)`：从连接中读取数据。
- `Write(b []byte) (n int, err error)`：向连接中写入数据。
- `Close() error`：关闭连接。
- `LocalAddr() Addr`：返回本地网络地址。
- `RemoteAddr() Addr`：返回远程网络地址。
- `SetDeadline(t time.Time) error`：设置连接的读写截止时间。
- `SetReadDeadline(t time.Time) error`：设置连接的读取截止时间。
- `SetWriteDeadline(t time.Time) error`：设置连接的写入截止时间。

#### PacketConn接口

- `ReadFrom(b []byte) (n int, addr Addr, err error)`：从连接中读取数据包及其来源网络地址。
- `WriteTo(b []byte, addr Addr) (n int, err error)`：向指定网络地址发送数据包。
- `Close() error`：关闭连接。
- `LocalAddr() Addr`：返回本地网络地址。
- `SetDeadline(t time.Time) error`：设置连接的读写截止时间。
- `SetReadDeadline(t time.Time) error`：设置连接的读取截止时间。
- `SetWriteDeadline(t time.Time) error`：设置连接的写入截止时间。

总之，`method.go`文件定义了`net`包中网络通信连接的核心方法，提供了网络编程中必不可少的方法，是Go语言网络编程的重要组成部分。

