# File: rawconn.go

rawconn.go文件是Go语言标准库中net包下的一个文件，主要用于提供原始TCP/UDP连接的操作和管理。

该文件中的主要结构体是rawConn，它封装了一个连接，包含了连接的文件描述符、本地地址、远程地址等信息。rawConn中定义了一些基本的I/O方法，如读取和写入数据的方法，以及一些其他的方法，如设置连接超时时间、获取连接状态等。同时，rawConn还实现了io.ReadWriteCloser接口，可以方便地用于Go语言中的各种I/O操作。

除此之外，rawConn还提供了一些底层的方法，如setUserTimeout、setNoDelay等，用于高级的TCP选项设置。同时，rawConn还支持使用IP地址或者主机名来进行网络连接，并且可以设置多种套接字选项。

总之，rawConn.go文件提供了基本的TCP/UDP连接操作和管理的功能，并且支持多种高级套接字选项，是Go语言网络编程中非常重要的一个组成部分。




---

### Structs:

### rawConn

rawConn结构体是net包中一个底层的连接结构体。它包括了TCP或UDP连接的socket描述符、本地地址和远程地址等信息。

其主要作用是提供底层的网络连接操作，如读写数据、设置连接参数等。它还可以根据底层连接类型的不同，提供TCP或UDP相关的操作和方法。

除了提供连接操作，rawConn还可以被其他更高层次的连接结构体组合使用，如TCPConn、UDPConn等，以提供更高层次的连接功能。

在网络编程中，rawConn是一个非常基础的连接结构体，为更高层次的连接结构体提供了基础支持，可以说是网络编程中的基础组件之一。



### rawListener

在 Go 语言的 net 包中，rawListener 是一个结构体，它用于表示一个原始连接监听器。原始连接是指没有经过任何协议处理，直接传输原始数据的连接。

rawListener 结构体的作用是用于创建并管理原始连接监听器。它定义了一些方法，包括：

- Accept：接受新连接并返回一个原始连接（rawConn）。
- Close：关闭原始连接监听器。

rawListener 结构体实现了 net.Listener 接口，因此可以使用标准的 net 包中的函数来创建和管理原始连接监听器，例如：

```go
rawListener, err := net.ListenPacket("ip:icmp", "0.0.0.0")
if err != nil {
    // 处理错误
}
defer rawListener.Close()

for {
    rawConn, err := rawListener.Accept()
    if err != nil {
        // 处理错误
    }
    // 处理原始连接
}
```

上述示例中，使用 net.ListenPacket 函数创建一个 IP 包监听器（"ip:icmp"）并指定监听地址为 "0.0.0.0"。接着使用 rawListener.Accept 方法接受新连接并返回 rawConn，再进行处理。最后使用 defer 关键字在函数退出时关闭原始连接监听器。

总之，rawListener 是 net 包中用于创建和管理原始连接监听器的一个结构体，它提供了接受连接和关闭监听器等方法，可以方便地进行原始连接的管理和处理。



## Functions:

### ok

在go/src/net中的rawconn.go文件中，ok这个func的作用是检查rawConn是否仍然处于活动状态。这个函数的定义如下：

func ok(err error) bool {
    if err == nil {
        return true
    }
    e, ok := err.(*OpError)
    if !ok {
        return false
    }
    err = e.Err
    if err == syscall.EINPROGRESS {
        return true
    }
    if op, ok := e.Op.(*netFD); ok && op.isBroken() {
        return false
    }
    return err == nil || err == syscall.EAGAIN || err == syscall.EWOULDBLOCK
}

ok函数有一个参数，即err error。它会检查err是否为nil，如果是，则返回true。否则，会将err转换为*OpError类型，并检查其是否是syscall.EINPROGRESS。如果是，则说明rawConn仍在等待网络操作完成，在这种情况下，ok函数应返回true。

如果不是syscall.EINPROGRESS，则会检查e.Op是否为*netFD类型，并调用其isBroken()方法，如果isBroken()方法返回true，则说明rawConn已经不再处于活动状态，应返回false。

最后，如果err是nil、syscall.EAGAIN或syscall.EWOULDBLOCK之一，则也应返回true。否则，返回false。



### Control

Control函数是net包中RawConn接口的一个方法。它的作用是将任意特定于网络连接的控制函数应用于底层连接。通过此方法，可以在底层连接上执行一些自定义的控制操作，例如启用/禁用TCP的某些选项，设置UDP的传输速率等。

这个方法的具体用法如下：

```go
func (c RawConn) Control(f func(fd uintptr)) error
```
其中，RawConn表示一个原始连接，f是一个函数，它接收一个uintptr参数，表示连接的文件描述符。

在Control函数内部，会调用底层的系统函数来执行指定的控制操作。例如，如果要启用TCP的拥塞控制算法（Cubic），可以使用以下代码：

```go
conn, err := net.Dial("tcp", "google.com:80")
if err != nil {
    log.Fatal(err)
}
rawConn, err := conn.SyscallConn()
if err != nil {
    log.Fatal(err)
}
err = rawConn.Control(func(fd uintptr) {
    err := syscall.SetsockoptInt(int(fd), syscall.IPPROTO_TCP, syscall.TCP_CONGESTION, 1)
    if err != nil {
        log.Fatal(err)
    }
})
if err != nil {
    log.Fatal(err)
}
```

在上面的代码中，我们通过SyscallConn方法获取了一个RawConn连接实例，然后使用Control方法来设置TCP的拥塞控制算法为Cubic。这个函数将fd（文件描述符）作为参数传递给了f函数，然后调用syscall包中的SetsockoptInt函数来设置TCP的选项。

总之，Control函数提供了一个灵活的接口，使得我们可以在底层连接上执行自定义的网络控制操作，让我们有更多的控制权来管理网络连接。



### Read

在net包的rawconn.go文件中，Read函数用于从底层原始连接读取数据。该函数的主要作用是从底层连接读取一个字节切片，以及指示是否有更多数据可用。

函数定义如下：
```go
func (c *rawConn) Read(b []byte) (int, error)
```

参数b是需要返回的数据的byte切片，返回值int表示实际读取的字节数，error表示是否出现了错误。

该函数的读取是阻塞的，也就是说，在没有数据可以读取之前，程序会阻塞并等待数据。如果到达了文件结束标记（EOF），则它会返回一个io.EOF错误。

如果读取数据时发生了错误，比如连接被关闭或超时，函数将返回该错误。

总的来说，该函数在Net编程中是非常重要的，它允许开发人员直接从底层连接读取原始数据，并能够处理错误和其他相关的数据传输操作。



### Write

rawconn.go中的Write函数是用于在底层的网络连接上写入数据的函数。底层的网络连接可以是TCP、UDP、IP、ICMP等不同的协议类型。该函数的作用是将数据写入底层的网络连接中，以便发送到远程主机。具体来说，该函数包括以下步骤：

1. 检查写入的数据的长度是否为0，如果是则直接返回。

2. 将要写入的数据添加到缓冲区中，此时缓冲区并没有真正写入到底层的网络连接中。

3. 如果缓冲区中的数据长度超过了缓冲区的阈值，则将缓冲区中的数据一次性写入到底层的网络连接中。

4. 如果缓冲区中的数据长度没有超过缓冲区的阈值，则等待下一次写入操作。

5. 当连接被关闭时，Write函数返回一个错误。

总之，Write函数是底层网络连接的写入接口，用于将数据传输到远程主机。



### PollFD

在Go语言中，PollFD函数位于net包中的rawconn.go文件中。这个函数主要用于操作系统级别的网络I/O事件轮询。PollFD函数执行的是底层的I/O多路复用，在进行网络I/O操作时，需要经常检查网络套接字是否有可用的数据或者是否需要等待。而PollFD函数就是用来监视网络套接字的状态，以便在可用时及时进行读写操作。

具体来说，PollFD函数的作用可以分为以下几个方面：

1. 注册文件描述符：PollFD函数会将指定的文件描述符注册到I/O多路复用的监视集合中，以便后续能够监视文件描述符的读写状态。

2. 等待I/O事件：PollFD函数会阻塞等待I/O事件的发生，直到有数据可读或者可写时才会返回。

3. 检查文件描述符状态：当有文件描述符状态变化时，PollFD函数会返回对应的事件类型，以便应用程序做出相应的处理。例如，当可读事件发生时，应用程序可以执行读取操作，当可写事件发生时，应用程序可以执行写入操作。

4. 跨平台支持：PollFD函数使用了底层操作系统提供的I/O多路复用技术，可以支持跨平台的网络I/O处理。

总之，PollFD函数是Go语言实现的底层网络I/O操作函数之一，它可以帮助应用程序进行高效的网络数据读写，并且具有跨平台的优点，是Go语言实现网络编程的重要组成部分。



### newRawConn

newRawConn是net包中的一个函数，其作用是创建一个基本的网络连接。这个函数返回一个rawConn结构体指针，这个结构体用来描述底层的网络连接。

具体而言，newRawConn函数会创建一个rawConn 结构体对象，并将其所有属性设为默认值。然后，它会打开一个新的套接字，并将该套接字绑定到指定的本地地址。接着，它返回这个rawConn 对象的指针。

需要注意的是，newRawConn函数是net包内部使用的辅助函数，并不适用于一般用户直接调用。它被其他许多函数所调用，比如TCP连接、UDP连接等。

总结一下就是，newRawConn函数的主要作用是创建一个基本的网络连接rawConn结构体对象。通过这个结构体，我们可以获取到底层网络连接的各种信息，并进行相应的操作，从而实现更高层次的网络通信。



### Read

在go/src/net/rawconn.go文件中，Read函数是用来从RawConn中读取数据的函数。RawConn是一个实现了Conn接口的原始底层网络连接，它提供了对TCP/IP和UDP网络传输协议的访问权限。

在RawConn中，Read函数用来读取网络传输协议数据流中的字节序列，将其读入到一个缓冲区中并返回读取的字节数。该函数有一个字节数组作为参数，读取的数据将存储在该字节数组中。

Read函数可以用来实现从RawConn中读取数据的操作。使用它可以将RawConn连接中的数据读取到应用程序中，从而实现数据交换和通信。同时，它还可以使用一些附加参数来控制读取操作的行为，例如使用超时参数来设置读取操作的超时时间，防止阻塞读取。

总之，Read函数是RawConn连接的一个重要组件，它提供了对底层网络连接中数据的读取、解码和处理的支持。在编写网络应用程序时，必须了解该函数的使用方法，以便实现高效、安全和可靠的通信。



### Write

Write是net包中的一个方法，定义在rawconn.go文件中。它的作用是向连接中写入数据。

在TCP连接中，Write方法将数据写入连接中，数据最终会被发送给远程主机。一个成功的Write方法调用返回的字节数表示发送成功的数据的长度，如果返回的字节数小于要发送的数据长度，则说明有部分数据没有成功发送出去。

在UDP连接中，Write方法也是将数据写入连接中，但是由于UDP是没有连接的，所以它是将数据发送给指定的远程主机。

需要注意的是，在网络编程中，数据传输往往会受到各种问题的影响，如网络拥塞、连接中断等，因此在使用Write方法时需要进行错误处理和重试等操作，以提高数据传输的可靠性。



### newRawListener

newRawListener函数是用于创建一个新的TCP监听器的函数。它采用一个网络地址（网络接口和端口号）作为参数，返回一个名为rawListener的新TCP监听器。

新创建的TCP监听器rawListener包含了一个TCP套接字（Socket），它与指定的地址绑定，并且可以监听到所有来自这个地址的TCP连接。接下来，这个TCP监听器可以用于接收来自客户端的TCP连接请求，并返回一个新的TCP连接对象Conn。

新的TCP连接的Conn对象可以用于在物理层面的TCP连接通道上进行读写操作。这个函数的作用是提供底层的TCP通信接口，为更高层次的应用协议（例如HTTP、FTP等）提供基础支持。



