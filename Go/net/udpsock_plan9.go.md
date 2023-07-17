# File: udpsock_plan9.go

udpsock_plan9.go是Go语言中net包中的一个文件，它的作用是在Plan 9操作系统上实现UDP socket。

在Plan 9操作系统中，UDP协议的实现方式与其他操作系统不同，因此需要单独实现UDP socket的相关函数和结构体。udpsock_plan9.go文件提供了Plan 9系统下UDP socket的相关实现。

具体来说，udpsock_plan9.go文件中定义了UDPConn类型，它包含了Plan 9系统下UDP socket的文件描述符fd、本地地址laddr、远程地址raddr等属性。该文件还实现了UDPConn相关的方法，如ReadFromUDP、WriteToUDP、Close等，以便用户在Plan 9系统下使用UDP socket进行网络通信。

总的来说，udpsock_plan9.go文件为Go语言在Plan 9操作系统下提供了UDP socket的实现，方便用户进行网络编程。




---

### Structs:

### udpHeader

在go/src/net中的udpsock_plan9.go文件中，udpHeader是一个结构体，用于表示UDP协议头。

UDP（User Datagram Protocol，用户数据报协议）是一种无连接的、轻量级的传输层协议，它不保证可靠传输，但具有低延迟和高吞吐量等特点。UDP协议头包含以下字段：

- Source Port：源端口，表示发送端口号；
- Destination Port：目的端口，表示接收端口号；
- Length：长度，表示UDP数据报的总长度，包括UDP协议头和UDP数据；
- Checksum：校验和，用于检测数据在传输过程中的完整性和正确性。

udpHeader这个结构体就是对UDP协议头的一个封装，包含了以上四个字段，并提供了一些方法可以方便地进行序列化和反序列化操作。具体而言，该结构体定义如下：

```go
type udpHeader struct {
    src  uint16
    dst  uint16
    len  uint16
    chk  uint16
}
```

其中，src表示源端口，dst表示目的端口，len表示UDP数据报长度，chk表示校验和。

除了以上字段，该结构体还提供了以下方法：

- marshal函数：将udpHeader序列化为长度为8字节的二进制数据；
- unmarshal函数：将长度为8字节的二进制数据反序列化为udpHeader结构体；
- setChecksum函数：计算UDP数据报的校验和，并将结果存入chk字段。



## Functions:

### readFrom

readFrom函数是UDP连接的读取方法，它从UDP套接字接收数据并将其放入缓冲区中。该函数的作用是从连接读取数据包并将其写入缓冲区。如果从连接中读取的数据超过了缓冲区的容量，那么超出部分将被丢弃。同时，该函数还返回接收到数据包的远程地址和端口号，以及任何错误信息。

具体来讲，readFrom函数会首先检查缓冲区是否已满，如果已满，则会直接返回错误信息。如果缓冲区未满，则会通过Unix系统调用从套接字中读取数据，并将读取到的数据存储在缓冲区中。同时，readFrom函数还会返回读取到的数据包的长度和远程地址和端口号等信息。

在UDP协议中，数据包不可靠且可能会丢失，因此程序需要不断接收数据并进行处理。readFrom函数就可以帮助程序在不停止程序的情况下接收UDP数据包。



### readFromAddrPort

readFromAddrPort是一个函数，用于从UDP套接字中读取数据时，同时获取发送方的IP地址和端口号。具体来说，它的作用是从UDP套接字中读取一个数据包并解析其中的源地址和端口号信息，然后返回读取到的数据和源地址及端口号。

这个函数的定义如下：

```
func (c *UDPConn) readFromAddrPort(b []byte) (int, *UDPAddr, error) {
    …
}
```

它的输入参数是一个指向字节数组的切片b，表示读取到的数据，输出参数是读取到的数据长度、发送方地址和端口号的指针以及可能出现的错误。这个函数实际上是UDPConn中的一个方法，因此它只能在UDPConn类型的对象上调用。

该函数通过调用Linux平台下的内核系统调用recvfrom系统调用来从UDP套接字中读数据，并利用UDP类型的go结构类型来解析数据包，并返回读取到的数据、源地址和端口号。通过此函数，可以将从UDP套接字读取的数据和发送方的IP地址和端口号分别获取到，方便后续的数据处理和应用程序对通信行为的统计。



### readMsg

readMsg函数是用于从UDP套接字中读取数据的函数。具体来说，它会读取UDP套接字中的数据并返回一个UDPConnRead对象，其中包含了读取的数据和发送方的地址信息。

readMsg函数的参数有三个：buf，control，和from。buf是一个byte类型的切片，用于接收读取到的数据；control是一个用于接收控制信息的切片；from是一个用于接收发送方地址的UDPAddr对象。

readMsg函数会将UDP套接字中的数据读取到buf切片中，然后将发送方的地址信息填充到from对象中。如果控制信息不为空，则也会将控制信息写入control切片中。

func readMsg(fd *netFD, buf []byte, oob []byte) (int, *UDPAddr, *ControlMessage, error) {

    var cm *ControlMessage
    if oob != nil {
        cm = new(ControlMessage)
    }
    // 读取数据
    n, sa, err := fd.readFrom(buf, oob)
    if err != nil {
        return 0, nil, nil, err
    }
    // 解析地址
    addr, err := anyToUDPAddr(sa)
    if err != nil {
        return 0, nil, nil, err
    }
    return n, addr, cm, nil
}

总之，readMsg函数是用于从UDP套接字中读取数据的函数，它可以读取数据并返回一个UDPConnRead对象，其中包含读取到的数据和发送方的地址信息。



### writeTo

在Go语言的标准库net包中，udpsock_plan9.go文件中的writeTo函数用于将数据报写入UDP网络连接中。

具体来说，该函数会向指定的地址发送数据，并返回已发送数据的字节数以及可能出现的错误。如果发送失败，将会返回相应的错误信息。该函数的定义如下：

```
func (c *UDPConn) writeTo(b []byte, addr net.Addr) (int, error)
```

其中，参数b表示待写入的数据报，参数addr表示数据报需要发送到的地址。该函数会根据addr参数的值发送UDP数据报，如果addr为nil，则数据报会默认发送到连接的远程地址。

在使用该函数时，需要注意的是，因为UDP是无连接的，所以如果写入的数据超出了MTU（最大传输单元），则该函数会将数据分成多个数据包进行发送，这会增加数据包的延迟、丢包的可能性，并且可能导致UDP协议的性能下降。

总之，writeTo函数是在UDP网络连接中用于向指定地址发送数据报的重要函数。



### writeToAddrPort

在go/src/net中的udpsock_plan9.go文件中，writeToAddrPort函数用于向指定的IP地址和端口发送UDP数据包。

具体来说，该函数获取一个UDP连接和一个IP地址和端口，从连接中获取数据包并将其发送到指定的IP地址和端口。如果发送成功，则返回已发送的字节数和零值的错误。

该函数的实现过程如下：

1.转换IP地址和端口为网络字节序。

2.获取UDP连接中的数据包，如果数据包为空，则返回错误。

3.使用Plan 9原生的UDP协议进行数据包的发送，发送数据包到指定的IP地址和端口。

4.如果发送成功，则返回已发送的字节数和零值的错误。

总之，writeToAddrPort函数是将UDP数据包发送到指定IP地址和端口的实现方法。



### writeMsg

writeMsg函数是用于UDP的发送消息功能。它的作用是将一个UDP消息封装成一个packet（也叫做UDP数据包），并发送到相应的对端。packet 内容由消息类型和消息数据组成。消息类型是UDP消息的参数。

函数原型如下：

```
func (c *UDPConn) writeMsg(p []byte, addr *UDPAddr, cm *controlMessage) (int, error) {}
```

其中：

- p：表示要发送的数据。类型为[]byte，即字节数组
- addr：表示数据的目标地址。类型为*UDPAddr，即指向UDP地址的指针
- cm：表示一个controlMessage对象，它是构建UDP包时使用的控制信息。类型为*controlMessage，即指向控制消息的指针
- 返回值：返回成功写入的字节数和可能发生的错误信息

在函数实现中，writeMsg先将UDP消息封装成packet，然后通过系统调用write向目标地址发送该packet。具体实现细节可以阅读函数源码。



### writeMsgAddrPort

writeMsgAddrPort是一个UDP连接对象的方法，它的作用是向指定的目标地址和端口发送UDP数据报文。

具体来说，该函数会将UDP数据报文写入一个特定的内存缓冲区，然后使用Plan 9网络协议栈的writeMsg系统调用将数据报文发送出去。

该函数接受以下参数：

- b []byte：要发送的UDP数据报文的字节数据
- addr *UDPAddr：目标地址和端口信息
- port int：本地端口号

其中，b和addr都是必须提供的参数。b表示要发送的数据，addr表示目标地址和端口号。而port则是可选的参数，它表示要使用的本地端口号。如果不提供该参数，则使用系统自动分配的空闲端口号。

需要注意的是，该函数并不保证数据报文能够成功到达目标地址。因为UDP是无连接的协议，因此可能会出现丢包、重复包等情况。因此，在使用该函数发送UDP数据报文时，需要根据具体场景进行判断和处理。



### dialUDP

dialUDP函数是用于创建一个UDP连接的函数，它位于go/src/net/udpsock_plan9.go文件中。

该函数的实现与平台相关。在Plan 9操作系统中，它使用一个syscall来创建一个UDP连接，并返回一个UDPConn类型的实例。

以下是dialUDP函数的参数和返回值：

func dialUDP(netProto string, laddr, raddr *UDPAddr) (*UDPConn, error) {}

具体来说，参数说明如下：

- netProto：指定使用的协议，只能是 "udp" 或 "udp4" 或 "udp6"。
- laddr：本地网址（Local Address），用于指定本地IP地址和端口。如果不需要指定，则可以传递nil。
- raddr：远程网址（Remote Address），用于指定远程IP地址和端口。

返回值说明如下：

- conn：表示创建的UDP连接的UDPConn类型的实例。
- err：表示创建UDP连接时可能发生的任何错误。

在创建UDP连接后，可以使用UDPConn的方法来读取和写入数据。例如，ReadFromUDP和WriteToUDP可以分别用于从UDP连接读取数据和将数据写入UDP连接。



### Bytes

在Go语言标准库的net包中，udpsock_plan9.go文件中的Bytes函数主要用于将UDP数据包的Header和Payload字节序列化。具体来说，该函数会将UDP数据包的Header和Payload字节序列化成一个[]byte类型的切片，并返回该切片。

函数源码如下：

func (b *UDPConn) Bytes() ([]byte, error) {
    if err := b.init(); err != nil {
        return nil, err
    }
    h := b.control[:]
    b.control[0] = byte(b.defaultCtrl >> 16)
    b.control[1] = byte(b.defaultCtrl >> 8)
    b.control[2] = byte(b.defaultCtrl)
    return h[b.nheader:], nil
}

其中，UDPConn是一个UDP连接的封装结构体，init方法用于初始化UDP连接，control是一个存放UDP控制信息的缓存，nheader表示UDP协议头的长度。

在该函数中，首先调用init方法初始化UDP连接。然后，将b.defaultCtrl这个16位无符号整数按照大端序序列化到control切片中，即分别将该整数的高8位、次高8位和低8位分别存放到control[0]、control[1]和control[2]中。最后，将control切片的第nheader个字节开始的部分作为UDP数据包的Payload序列化成一个新的切片并返回。

总的来说，Bytes函数的作用是用于将UDP数据包的Header和Payload序列化成一个[]byte类型的切片。该函数在UDP通信中起到重要的作用，可以帮助用户将UDP数据包按照网络字节序列化后发送到目标主机。



### unmarshalUDPHeader

函数unmarshalUDPHeader的作用是将UDP数据包的头部信息反序列化，即将字节流转换为结构体。

具体而言，该函数接收一个字节切片作为参数，表示UDP数据包的头部信息，然后根据UDP协议规定的头部格式，将字节切片中的各个字段解析出来，存入一个udpHeader结构体中。这个结构体包括以下字段：

- SourcePort: 发送方端口号，2个字节
- DestPort: 接收方端口号，2个字节
- Length: UDP数据包总长度，2个字节
- Checksum: 校验和，2个字节

解析后的udpHeader结构体信息将作为输出结果返回给调用者。

在UDP协议中，头部信息是固定长度的，而且每个字段的大小和顺序都是规定好的。因此，unmarshalUDPHeader函数的作用在于将收到的UDP数据包头部信息按照规定的格式解析出来，以便后续的处理和分析。例如，在UDP服务中，需要根据接收方端口号将UDP数据包分发到不同的处理函数中。此时，就需要解析UDP数据包头部信息来获取接收方端口号，从而进行分发。



### listenUDP

listenUDP函数是UDPConn的内部函数，用于在Plan 9操作系统下创建一个UDP套接字并绑定到指定的本地地址和端口。它的作用是：

1. 创建一个新的UDP套接字并返回UDPConn对象；

2. 将UDPConn对象绑定到指定的本地地址和端口；

3. 在UDPConn对象的底层文件描述符上注册一个读取事件，以监听从网络上收到的UDP数据包，并将它们发送到处理程序；

4. 返回UDPConn对象，以供开发人员使用。

该函数具有以下参数：

1. laddr：本地地址，表示要绑定的IP地址和端口号；如果为nil，则使用默认的本地地址。

2. ifi：网络接口，表示要绑定到哪个网络接口；如果为nil，则使用默认的网络接口。

3. zone：地址范围，表示广播地址的范围；如果为nil，则使用默认的地址范围。

该函数的主要实现流程如下：

1. 调用systemstack函数获取系统栈，并将它作为函数的第一个参数；

2. 调用socket函数创建一个新的UDP套接字，并返回其文件描述符；

3. 调用bind函数将UDP套接字绑定到指定的本地地址和端口；

4. 调用setsockopt函数设置一些选项，如广播、多播等；

5. 调用NetpollInit函数创建并初始化pollDesc对象，用于注册读取事件；

6. 调用netpoll函数将pollDesc对象的文件描述符注册到网络轮询器中，用于监听读取事件；

7. 返回UDPConn对象，供开发人员使用。

总之，listenUDP函数是UDPConn的重要组成部分，它使得UDP数据包可以被正确地接收并处理。



### listenMulticastUDP

listenMulticastUDP函数是用于在Plan 9操作系统上创建一个UDP多播socket的函数。在Plan 9中，使用IP组播协议来实现UDP多播。该函数按照指定的IP地址和端口号创建一个UDP多播socket，并加入指定的多播组。 

函数签名如下：

```go
func listenMulticastUDP(network string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)
```

其中，network参数指定网络类型，应该使用"udp4"或"udp6"，ifi参数指定要使用的网络接口，gaddr参数是一个UDPAddr类型的地址，指定要加入到的多播组的IP地址和端口号。 

该函数返回一个UDPConn类型的指针和一个错误类型的值。如果成功创建了多播socket，则返回的UDPConn指针可以用于发送和接收消息。如果创建失败，则返回一个非nil的错误类型的值，包含错误信息。 

在函数内部，首先会通过net.ResolveUDPAddr函数将gaddr中的IP地址和端口号转换为一个*UDPAddr类型的地址，然后调用net.ListenUDP函数创建一个UDP网络连接，并将其绑定到本地系统上随机生成的端口号，在使用setsockopt函数将其加入到指定的多播组中。

总之，listenMulticastUDP函数是一个用于创建并加入UDP多播组的函数，可以让应用程序在Plan 9操作系统上发送和接收UDP多播消息。



