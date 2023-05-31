# File: iprawsock_posix.go

iprawsock_posix.go是Go语言标准库中net包中的一个文件，它实现了基于原始套接字的网络通信功能。具体来说，它提供了IP层和通用数据报协议（UDP）层的原始套接字接口，用于实现IP和UDP协议的自定义网络应用。

该文件主要定义了两个类型：ipRawConn和udpPayload。

ipRawConn类型是基于原始套接字的IP连接，它包含连接的源地址和目标地址等信息，支持读取和写入原始的IP数据报文。使用ipRawConn可以直接获取到网络层的原始数据，对于一些特殊的网络应用非常有用。

udpPayload类型是基于UDP协议的数据载荷，它封装了UDP的源IP地址、目标IP地址、源端口号和目标端口号等信息，以及具体的UDP数据。使用udpPayload可以方便地构造和发送UDP数据包。

此外，该文件还定义了一些变量和函数，如listenPacket、readFromIP、writeToIP、readFromUDP和writeToUDP等，用于实现IP和UDP层的网络读写操作。这些操作可以实现基于原始套接字的网络编程，为开发者提供了灵活的系统级网络操作能力。

总之，iprawsock_posix.go文件实现了基于原始套接字的IP和UDP协议传输，并提供了对原始网络数据的读写等操作，为Go语言网络编程提供了强大的底层支持。

## Functions:

### sockaddrToIP

sockaddrToIP函数可将一个sockaddr结构的IP地址转换为net.IP类型的IP地址。在网络编程中，使用sockaddr结构表示一个通用的网络地址，这个地址可能是ipv4或ipv6的地址结构。sockaddrToIP函数将这个通用的地址结构解析为特定的IP地址，以便进一步的操作。

具体地，sockaddrToIP函数的作用如下:

1. 接收一个sockaddr结构以及这个结构的大小。
2. 获取结构中的IP地址部分，以及该地址的协议类型（ipv4/ipv6）。
3. 根据协议类型将IP地址转换为net.IP类型，并返回这个IP地址。

在网络编程中，经常需要将sockaddr结构的IP地址转换为实际可用的IP地址，这个转换过程就是通过sockaddrToIP函数来实现的。



### family

在 go/src/net 中的 iprawsock_posix.go 文件中，family 函数用于根据给定的 IP 地址确定网络协议族。它的作用是将一个字符串形式的 IPv4 或 IPv6 地址转换为一个相应的协议族常量。

该函数的实现利用了 Go 语言的 net 库中的 ParseIP 方法，该方法会分析一个 IP 地址，并返回一个 net.IP 类型的值。接着，如果该 IP 地址是 IPv4 地址，该函数会返回 syscall.AF_INET，否则返回 syscall.AF_INET6。

家族是一种通信协议，它定义了一组规则，以确定如何在计算机网络中传输数据。协议族常量是一个整数值，用于指定所使用的协议族。在 Posix 系统中（如 Linux 和 macOS），网络协议族由一个整数值来表示，如 AF_INET（IPv4）和 AF_INET6（IPv6）。函数的作用是实现根据 IP 地址选择协议族的功能，以确保连接到正确的网络协议族。



### sockaddr

在go/src/net中的iprawsock_posix.go文件中，sockaddr是一个函数。该函数的作用是将IP地址和端口号与网络协议相关联的结构体（即Sockaddr结构体）转换为操作系统特定的Sockaddr结构。

具体来说，Sockaddr结构包含了一个通用的地址结构体，它可以同时支持IPv4和IPv6地址。而不同的操作系统实现了不同的Sockaddr结构体，因此，当使用go语言操作IP网络时，需要将通用的Sockaddr结构体转换为特定操作系统下的Sockaddr结构体。

sockaddr函数接受两个参数，第一个参数是IP地址和端口号相关联的Sockaddr结构体，第二个参数是特定操作系统下的Sockaddr结构体的内存地址。它使用系统相关的函数将通用Sockaddr转换为操作系统特定的Sockaddr结构体，并将结果存储在第二个参数所指向的内存地址中。

总之，sockaddr函数的作用是实现在不同操作系统下处理IP地址和端口号相对应的Sockaddr结构体，使网络编程实现更加简单和可移植。



### toLocal

toLocal函数的作用是将目标IP地址转换为本地IP地址。具体来说，它在IPv4地址和IPv6地址之间进行转换，以适应不同的应用程序需求。

在IPv4中，IP地址由32位二进制数表示，而IPv6地址由128位二进制数表示。因此，在进行socket编程时，需要进行适当的转换，以便在两种协议之间进行通信。

toLocal函数的主要任务是对IPv6地址进行转换，将其转换为匹配的IPv4地址。它使用IPv6地址的最后32位，将其视为IPv4地址的IP地址部分，并将其分配给IPv4地址的低32位。这样，IPv6地址就可以转换为IPv4地址，以便在支持IPv6协议的主机和只支持IPv4协议的主机之间进行通信。

总的来说，toLocal函数的作用是帮助应用程序在不同的协议之间进行通信，确保网络连接的稳定和正确性。



### readFrom

在go/src/net/iprawsock_posix.go文件中，readFrom函数用于从底层的套接字读取数据，并返回读取的数据、从哪个地址读取到的数据、网路层协议号和错误信息。

更具体地说，readFrom函数的作用是：

1. 读取数据：readFrom函数调用read系统调用从原始套接字中读取数据，并返回读取到的字节数和错误信息。

2. 返回地址信息：同时，readFrom函数还会在读取到数据时获取数据的来源地址。它使用C语言的recvfrom系统调用读取数据，并将读取到的地址保存在一个syscall.Sockaddr结构体中。

3. 返回网络协议：通过原始套接字读取的数据不会被TCP或UDP处理，而是直接传输到应用程序中。因此，readFrom函数需要返回数据的网络层协议，以便应用程序根据不同的协议来处理接收到的数据。

总之，readFrom函数是一个非常重要的函数，因为使用原始套接字进行底层网络通信时，它是必须的。它可以用于接收网络层的各种协议的数据，并返回来源地址和错误信息，为高层网络协议的实现提供基础支持。



### stripIPv4Header

stripIPv4Header是一个用于去除IPv4数据包头部的函数，它的作用是将IPv4数据包中的头部信息去除，提取出数据部分。

在网络通信过程中，数据包被封装在不同的协议层中，头部信息用于描述数据包的协议、源地址、目标地址以及其他控制信息。在传输过程中，每个协议层的头部会被添加到数据包中，形成“协议栈”，而在接收端，需要逐层将头部去除，才能提取出数据部分。

在IPv4协议中，数据包头部是一个长度为20字节的固定格式的结构体，其中包含了源IP地址、目标IP地址、协议类型等信息。这些信息对于路由和网络控制非常重要，但却不需要在应用层得到保留。因此，在接收端，需要去除IPv4数据包头部，将数据部分提取出来，交给上层协议进行处理。

stripIPv4Header函数就是用于完成这个工作的。它会接收一个包含IPv4头部的字节数组，并将头部信息去除，返回剩余的数据部分。具体实现方式是，在IPv4头部中获取数据长度字段，然后从原始数据中截取该长度的数据作为返回值。

需要注意的是，stripIPv4Header函数只能用于IPv4数据包中去除头部，对于其他协议层的头部去除需要使用相应的函数。



### readMsg

iprawsock_posix.go文件中的readMsg函数是用于从IP原始套接字接收数据的函数。

IP原始套接字是一种特殊的网络套接字，它可以直接访问网络协议栈的底层数据包，而不必经过操作系统的高层抽象。这使得应用程序可以更直接地控制和处理网络数据。

readMsg函数的作用是读取从IP原始套接字接收到的数据，并将其封装成一个net.IPMessage类型的对象返回。net.IPMessage是一个表示IP数据包的结构体类型，其中包括数据包的头部信息和数据负载。

readMsg函数使用了底层系统调用recvmsg来接收数据。在接收数据之前，它会使用unix.SetNonblock函数将套接字设置为非阻塞模式，这样在没有数据可读的时候可以立即返回并避免阻塞。

在接收到数据后，readMsg函数会依次解析IP头部，ICMP头部和TCP/UDP头部（如果有的话），并将它们封装成一个net.IPMessage类型的对象返回。

总的来说，readMsg函数的作用是从底层IP原始套接字读取数据，并封装成net.IPMessage类型的对象，方便应用程序进行数据处理和分析。



### writeTo

writeTo是net/ipraw的一个函数，它的作用是将一个IP数据包写入到UDP协议栈中。在IP层中，IP数据包是通过网络传输的，而在UDP协议栈中，UDP协议把数据包封装成UDP报文并进行传输。

该函数接收一个Conn接口和一个地址作为参数，其中Conn接口代表了一个IP连接，通常是通过DialIP或ListenIP创建的。地址参数表示了发送目标，一般是一个IP地址和端口号。

在函数内部，它首先会将IP数据包封装成UDP报文，然后使用Conn接口发送该UDP报文。如果发送成功，则返回发送的字节数。如果发送失败，则会返回一个非nil的错误对象。

总之，writeTo函数是对IP数据包进行封装和发送的过程，将数据包发送到指定的目标地址，并返回发送的字节数或错误信息。



### writeMsg

在go/src/net中的iprawsock_posix.go文件中，writeMsg函数的作用是将UDP数据包写入给定的目标IP地址。

详细介绍如下：

函数原型：

```go
func writeMsg(fd int, b []byte, oob []byte, to *UDPAddr) (int, error)
```

参数说明：

1. fd：file descriptor，指定写入数据包的文件描述符

2. b：即将写入的UDP数据包的内容

3. oob：Out-of-Band data，为了在发送数据的过程中提供额外的信息。在该函数中，仅用于捕获源IP地址和端口信息。

4. to：目标IP地址

函数返回值：

成功写入数据包的字节数和可能发生的错误

函数执行流程：

1. 构造UNIX Kernel的struct msghdr对象，用于发送数据包的整体信息

2. 初始化msghdr对象的目标地址和控制信息

3. 发送UDP数据包

4. 将源IP地址和端口号填充到UDP数据包的OOB缓存区oob中，并将该信息添加到msghdr的控制信息中

5. 返回成功写入数据包的字节数和错误信息

总之，writeMsg函数是将UDP数据包写入给定的目标IP地址，并可以在控制信息中捕获源IP地址和端口信息以及其他所需的信息。



### dialIP

dialIP函数在net包中的iprawsock_posix.go文件中实现，用于建立基于IP协议的连接。该函数接收一个ip地址（IPv4或IPv6），目标地址端口号和本地地址端口号，然后创建一个基于IP协议的socket连接。在创建连接后，如果本地地址端口号为0，则将自动分配一个可用端口号。如果连接成功，则返回net.Conn实例和nil错误。否则，将返回一个nil连接和一个错误。通过调用返回的net.Conn实例就可以进行数据通信。

下面是dialIP函数的代码实现：

```
func dialIP(netProto, laddr, raddr *IPAddr) (*IPConn, error) {
    if !ipIsMulticast(raddr.IP) && !ipAllowed(raddr.IP) {
        return nil, &AddrError{Err: "invalid argument", Addr: raddr.String()}
    }
    s, err := sysSocket(family(netProto, raddr.IP))
    if err != nil {
        return nil, err
    }
    if err := setDefaultSockopts(s); err != nil {
        closeFunc(s)
        return nil, err
    }
    la := &syscall.SockaddrInet4{Port: 0}
    if laddr != nil {
        la.IP = laddr.IP.to4()
        la.Port = laddr.Port
        if ipVer(laddr.IP) != ipVer(raddr.IP) {
            return nil, &AddrError{Err: "mismatched IP protocol version", Addr: laddr.String()}
        }
    }
    if err := bindToSocket(s, la); err != nil {
        closeFunc(s)
        return nil, err
    }
    rsa := &syscall.SockaddrInet4{Port: raddr.Port}
    if ipVer(raddr.IP) == 6 {
        rsa := &syscall.SockaddrInet6{Port: raddr.Port}
        copy(rsa.Addr[:], raddr.IP.To16())
        if laddr != nil {
            copy(rsa.ZoneId[:], laddr.Zone)
        }
    } else {
        copy(rsa.Addr[:], raddr.IP.To4())
    }
    if err := syscall.Connect(s, rsa); err != nil {
        closeFunc(s)
        return nil, err
    }
    return newIPConn(s), nil
}
```

该函数的主要步骤：

1. 验证目标地址是否允许创建连接，如果不允许则返回错误
2. 创建IP连接的socket
3. 配置socket的默认选项
4. 绑定本地地址和端口号到创建的socket
5. 连接目标地址和端口号
6. 返回net.Conn实例和nil错误，如果连接失败则返回nil连接和一个错误



### listenIP

listenIP是一个函数，主要用于在指定的IP地址和端口上创建一个原始socket。原始socket是一种可以直接与网络协议栈交互的高级socket，可以用于定制化的网络应用程序或者测试网络协议栈。

在该函数中，首先通过系统调用socket创建一个原始socket，然后通过设置socket的地址属性（包括IP地址、端口号等）来指定原始socket的使用方式。这样就可以使用该原始socket来发送和接收特定IP地址和端口号的数据包。同时，该函数还考虑了多种错误情况，并在错误发生时返回对应的错误信息。

总的来说，listenIP函数主要用于创建一个原始socket，以便与指定的IP地址和端口进行交互。这对于定制化的网络应用程序或者测试网络协议栈非常有用。



