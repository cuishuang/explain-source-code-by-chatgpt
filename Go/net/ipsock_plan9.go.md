# File: ipsock_plan9.go

ipsock_plan9.go文件是Go语言的标准库中net包下的一个文件，主要实现了Plan9操作系统下的IP网络套接字的相关操作。该文件包含了 Plan9系统下 IP地址、端口号、套接字协议族的解析和转换方法。

具体来说，ipsock_plan9.go文件中封装了如下函数：

- func parseIPv4(s string) IP：用于将IPv4地址字符串转换为IP类型。
- func listenPlan9(net string, laddr *SockaddrInet4) (fd *netFD, err error)：用于创建并监听Plan9套接字。
- func parsePlan9IP(s string) (IP, error)：用于将 Plan9操作系统下的IP字符串转换为IP类型。
- func plan9AddrToIP(addr *NetAddr) IP：用于将Plan9套接字的地址转换为IP类型。
- func plan9IPCmp(a, b IP) int：用于比较两个Plan9 IP地址的大小。

这些函数的实现让Go语言支持在Plan9操作系统下基于IP的网络通信。同时，该文件的存在也展示了Go语言在跨平台网络编程方面的能力。

## Functions:

### probe

ipsock_plan9.go这个文件中的probe()函数是用来检查IP地址是否可以连接到特定的端口。通过使用plan9的net.Dial()函数，probe()函数会先尝试通过TCP连接到指定的地址和端口，并检查是否成功建立连接。如果建立连接失败，则probe()函数会尝试通过UDP协议进行连接，同样检查是否运作正常。如果UDP连接也失败，则probe()函数会返回错误信息。

该函数的作用是在网络中探测某个IP地址上某个端口是否可用，以便网络程序能够准确、快速地确定哪些IP地址和端口是可以连接的，哪些是不可用的。这对于网络应用程序是非常重要的，因为它们需要快速和准确地建立与其他计算机的通信连接。



### probe

在 Go 语言的 net 包中，ipsock_plan9.go 文件中的 probe 函数的作用是检查本地网络和 IP 源地址的可用性。

具体来说，probe 函数会使用 IP 和端口号创建一个本地地址，并通过向该地址发送 ICMP 报文来检查本地网络连接是否正常。如果连接正常，则函数会返回一个非空地址，否则返回空地址。

此外，probe 函数还会检查 IP 源地址是否可用。它在内部调用 net.DialIP 函数来尝试通过该地址创建一个网络连接。如果连接成功，则说明该地址可用；否则，说明该地址不可用。

总之，probe 函数的主要作用是检查本地网络和 IP 源地址的可用性，以便在网络编程中选择合适的地址来建立连接。



### parsePlan9Addr

`parsePlan9Addr`函数是用于解析Plan9风格的IP地址的函数。Plan9是一个类Unix操作系统，它的IP地址格式和其他操作系统有所不同。

具体来说，`parsePlan9Addr`函数将传入的字符串解析为一个IP地址和端口号，返回一个`TCPAddr`类型的结构体。如果解析失败，函数将返回一个错误。

以下是`parsePlan9Addr`函数的代码实现：

```go
func parsePlan9Addr(s string) (*TCPAddr, error) {
    var ip IP
    var zone string
    var port int
    var err error

    if len(s) == 0 {
        return nil, EINVAL
    }

    switch s[0] {
    case '/':
        addr := s
        var cc string
        if addr[1] == 'e' && addr[2] == '/' {
            addr = addr[1:]
        } else {
            cc, addr = split(addr[1:], '/')
            if cc == "" {
                cc = "net"
            }
        }
        if addr == "" {
            return nil, EINVAL
        }
        port, err = parsePort(addr)
        if err != nil {
            return nil, err
        }
        ip = IPv6unspecified
        zone = cc
    case '#':
        addr := s[1:]
        var cc string
        if addr[0] == 'e' && addr[1] == '#' {
            addr = addr[1:]
        } else {
            addr, cc = split(addr, '#')
        }
        if len(addr) == 0 {
            return nil, EINVAL
        }
        port = 0
        ip, err = parseIPv6(addr)
        if err != nil {
            return nil, err
        }
        zone = cc
    default:
        addr, c, err := splitHostPort(s)
        if err != nil {
            return nil, err
        }
        ip = ParseIP(addr)
        if ip == nil {
            return nil, &AddrError{Err: "cannot parse address", Addr: s}
        }
        if c != "" {
            port, err = strconv.Atoi(c)
            if err != nil || port < 0 || port > 0xFFFF {
                return nil, &AddrError{Err: "invalid port", Addr: s}
            }
        }
    }
    return &TCPAddr{IP: ip, Port: port, Zone: zone}, nil
}
```

函数首先检查传入的字符串是否为空，如果为空则返回`EINVAL`错误。然后，它根据IP地址的开头字符来判断地址的类型。

如果地址以`/`开头，则它是一个本地IPC地址，IP地址为IPv6的未指定地址，端口号从地址中解析出来，而区域ID为剩余部分的字符串。

如果地址以`#`开头，则它是一个IPv6地址，端口号为0，而区域ID从地址中解析出来。

如果地址不以`/`或`#`开头，则函数会尝试将地址和端口号分离，并将地址解析为IP地址。

最后，函数将解析出来的IP地址、端口号和区域ID封装到一个`TCPAddr`结构体中并返回。如果地址解析失败，则函数会返回一个错误。



### readPlan9Addr

readPlan9Addr函数的作用是从一个带有Plan 9风格地址的Unix域套接字中读取地址信息。Plan 9风格的地址格式是“{namespace} {address}”，其中namespace可以是“unix”或“tcp”，address则是命名空间特定的地址。该函数将地址信息解析为IP地址和端口号，以供后续使用。

具体的实现过程如下：

1. 从Unix域套接字中读取地址信息并解析它为一个UTF-8字符串。

2. 解析字符串中命名空间和地址部分，如果没有指定命名空间则默认为“unix”。

3. 根据命名空间的不同，进行相应的地址解析。如果是“unix”命名空间，则解析地址为Unix域socket路径；如果是“tcp”命名空间，则解析地址为TCP/IP地址和端口号。

4. 返回解析得到的IP地址和端口号作为一个Sockaddr类型的值。如果解析失败则返回nil。

总体上，这个函数对Unix域套接字进行了地址的解析和转换，方便后续的网络传输操作。



### startPlan9

ipsock_plan9.go文件中的startPlan9函数的作用是启动Plan 9协议栈并设置相关参数，准备在Plan 9上进行网络通信。该函数会通过调用netstack函数进行初始化，并设置一系列参数，包括IP地址、子网掩码、默认网关等等。同时，该函数还会监听网络接口状态，并根据状态变化来改变网络配置。

具体来说，startPlan9函数会调用netstack函数，创建一个Plan 9的网络栈，并存储在静态变量netstack中。在创建网络栈之后，函数会读取系统的网络配置，包括IP地址、子网掩码、默认网关等，在网络栈中进行设置。接着，函数会设置网络接口变化的通知，并将网络栈中网卡相关的信息设置好。最后，函数会开始监听网络接口状态，并在状态变化的时候更新网络栈中的相关参数，以保证网络栈总是处于最优的状态。

总之，startPlan9函数是Plan 9协议栈的起点，它负责启动、初始化和配置网络栈，以及监听网络接口状态的变化。它是网络通信过程中不可或缺的一环。



### fixErr

在ipsock_plan9.go文件中，fixErr函数的主要作用是将Windows和Linux系统中的IPv6错误代码转换为Plan 9系统中的错误代码，以便更好地处理网络连接问题。具体来说，fixErr函数接受一个error类型的参数，该参数可能是由IPv6连接产生的错误。如果错误代码在Windows或Linux系统中，则将其映射到Plan 9系统中的错误代码。这样，是当前程序在Plan 9环境下运行时，可以更方便地处理IPv6连接错误，使程序更加健壮和可靠。



### dialPlan9

dialPlan9函数是go/src/net中的一个函数，该函数的作用是在Plan 9操作系统中创建并返回一个连接到指定网络地址的IP套接字。当在目标地址上建立连接时，该函数将遍历可用的本地接口，选择最合适的本地IP地址以使用。如果找不到合适的本地IP地址，则返回一个错误。

该函数使用Plan 9操作系统提供的接口来创建套接字。在创建套接字之前，该函数会设置套接字使用的协议与网络类型。然后，它使用目标地址和端口号来构建一个套接字地址，并为套接字分配一个本地地址。最后，该函数启动套接字连接并将其返回给应用程序进行使用。

总的来说，dialPlan9函数是将应用程序的请求转换为具体的协议细节，以便在Plan 9操作系统中创建一个可用的IP套接字，连接到指定的网络地址。



### dialPlan9Blocking

dialPlan9Blocking是net包中ipsock_plan9.go文件中的一个函数，是 Plan 9 平台上的实现，用于与指定的网络地址建立连接，并返回一个与该连接关联的net.Conn接口。该函数主要的作用就是在 Plan 9 版本的操作系统上建立阻塞式的 IP 连接。

该函数的具体实现过程如下：

首先，该函数会根据连接协议（IPv4/IPv6）以及网络地址创建一个socket（文件描述符），并将其绑定到系统指定的端口，以及设置一些socket选项，例如 SO_REUSEADDR、SO_REUSEPORT。

然后，该函数会根据socket的文件描述符以及网络地址等参数，调用系统提供的connect系统调用函数，来连接指定的网络地址。

如果建立连接成功，该函数会返回一个与该连接关联的net.Conn接口。否则，该函数会返回对应的错误信息。

总之，dialPlan9Blocking函数是用于在 Plan 9 平台上建立阻塞式的 IP 连接的函数，它在net包中的使用范围广泛，例如在http、smtp、ftp等协议中都有使用到。



### listenPlan9

listenPlan9函数是net包中的一个函数，用于在Plan 9操作系统上创建一个IP网络监听器。它的作用是创建一个IP地址和端口号绑定到本地网络接口上的TCP/UDP server监听器。

在Plan 9操作系统上，网络设备驱动程序是由内核提供的，因此网络监听器的实现与其他操作系统不同。listenPlan9的实现方式是调用plan9.Listen()和plan9.Bind()来创建一个UDP或TCP的本地socket。它会根据参数指定的IP地址和端口创建一个本地的网络连接，并将socket绑定到本地网络接口上。

listenPlan9函数的参数包括network、address和service。network指定要监听的网络类型，可以是“tcp”或“udp”；address指定要监听的IPv4或IPv6地址；service指定要监听的端口号。

listenPlan9函数返回一个net.Listener对象，可以使用它来接受客户端连接，并启动一个goroutine来处理每个连接请求。该函数还会返回一个可能出现的错误，例如IP地址无效或端口已被占用等错误。



### netFD

netFD是一个函数，用于在Plan 9操作系统上创建网络FD（文件描述符）。在Plan 9中，网络和文件系统使用相同的接口和概念，因此要使用网络套接字，需要创建一个FD。这个函数是net包中的一个内部函数，实现了将一个新的文件描述符与套接字关联。

当调用net.Dial或net.Listen等函数时，会在内部调用netFD函数来创建网络FD。这个函数会使用操作系统的网络库来创建一个新的套接字，并把它关联到一个文件描述符上。然后，这个文件描述符就可以传递给其他网络函数，如Read和Write等，来进行数据的收发操作。

netFD函数还实现了非阻塞I/O和超时功能。在Plan 9中，I/O操作默认是阻塞的。但是，通过设置文件描述符的标志位，可以将它变成非阻塞模式。当读或写操作在非阻塞模式下进行时，netFD函数会轮询套接字，直到数据准备好或超时。如果超时，它会返回一个错误。

总之，netFD函数是net包中的一个重要组成部分，在Plan 9操作系统上实现了网络套接字的创建和操作，并支持非阻塞I/O和超时功能。



### acceptPlan9

acceptPlan9是一个在IP网络上面接受连接请求并返回一个新的连接的函数。它主要用于在Plan 9操作系统上使用。该函数会在指定的socket上等待新的连接请求，一旦有新的连接请求到达，它就会返回一个新的连接以供使用。

在具体实现中，acceptPlan9函数首先从指定的socket读取连接请求。然后，它会将新连接的本地地址和远程地址保存在一个新创建的net.Addr对象中，并将该对象和新连接的文件描述符一起作为返回值返回给调用者。

该函数的实现与其他操作系统上的实现略有不同，因为Plan 9操作系统使用了不同的系统调用和数据类型来处理网络连接。因此，该函数必须使用Plan 9特有的系统调用来处理网络连接。然而，它仍然遵循了在net包中定义的接口和行为，使得调用者可以方便地使用它来接受新的连接。



### isWildcard

isWildcard是一个函数，在go/src/net/ipsock_plan9.go文件中定义。这个函数的作用是用来检查一个IP地址是否是通配符地址（Wildcard Address）。

通配符地址是指IP地址中的一部分被设置为0，用来表示任意地址。在IPv4中，通配符地址通常表示为0.0.0.0，在IPv6中通配符地址通常表示为::。

isWildcard函数首先会将一个IP地址转换为IP类型，并判断它是IPv4还是IPv6地址。对于IPv4地址，如果IP的全部四个宽度都是0，那么这个IP就是通配符地址。对于IPv6地址，如果IP的全部16个宽度都是0，那么这个IP就是通配符地址。

isWildcard函数的返回值是一个布尔值。如果IP地址是通配符地址，则返回true，否则返回false。

这个函数主要用于在网络编程中进行地址绑定时的操作。如果绑定了通配符地址，那么这个套接字就会接收到所有发往该IP地址的数据包。



### toLocal

toLocal函数是net包中IPConn类型的一个私有方法，用于将一个IP地址转换为本地地址。具体来说，它先检查IP地址是否为IPv4的特殊本地地址（127.0.0.1），如果是则直接返回该地址；否则调用net包中的lookupIP函数查找本地主机的IP地址列表，然后返回其中一个和输入IP地址在同一子网的IP地址。

这个方法的主要作用是将一个IP地址与本地主机的IP地址进行比较，以确定是否需要在本地发送或接收数据。在Socket编程中，通常会使用本地地址（也称为绑定地址）来指定Socket绑定的网络接口和端口号。如果不使用本地地址，则操作系统将自动选择一个可用的地址和端口号进行绑定，这样可能会导致网络瓶颈、安全问题等不良后果。

总之，toLocal函数是net包中用于将IP地址转换为本地地址的基础函数，是Socket编程中非常重要的一部分。



### plan9LocalAddr

plan9LocalAddr函数是用来获取Plan 9操作系统中网络接口的本地IP地址的函数。

Plan 9是一种分布式操作系统，因此具有多个网络接口。每个接口都有一个本地IP地址。plan9LocalAddr函数的作用是获取当前计算机上的所有网络接口的本地IP地址，并将它们存储在一个数组中。

这个函数使用了Plan 9操作系统中和网络相关的系统调用来获取网络接口信息。首先，它调用了内核提供的socket系统调用，以创建一个socket对象。然后，它调用了getifaddrs系统调用，以获取连接到计算机的网络接口的列表。对于每个网络接口，函数会调用getsockname系统调用，以获取与该接口关联的本地IP地址，并将这些地址存储在一个数组中。

一旦plan9LocalAddr函数完成对所有网络接口的遍历，它会返回一个包含本地IP地址的数组。这些IP地址可以供其他程序使用，例如用于网络通信或者服务发现等。



### hangupCtlWrite

hangupCtlWrite是一个函数，它位于net / ipsock_plan9.go中。 它的作用是关闭一个给定的fd的写入端并通知所有正在等待以向其写入数据的线程。 它的代码如下：

```go
func hangupCtlWrite(fd *netFD) {
    fd.writeLock.Lock()
    defer fd.writeLock.Unlock()
    if !fd.shutw {
        fd.setAddr(nil)
        fd.eofError = errClosing
        fd.shutw = true
        fd.writeClosed.Set(true)
    }
}
```

首先，该函数获得一个互斥锁，以确保线程安全性。 如果fd的写入端尚未关闭，则函数会将其设置为关闭状态，并通知任何等待写入该fd的线程。

该函数的目的是确保安全关闭fd，以避免死锁和其他问题，同时仍允许写入fd的所有线程被通知该fd已关闭。



