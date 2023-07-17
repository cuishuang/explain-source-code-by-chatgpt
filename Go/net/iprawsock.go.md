# File: iprawsock.go

iprawsock.go是Go语言中net包的一部分，主要实现了一个IPv4和IPv6的RAW（原始）套接字封装器，用于向操作系统发送和接收原始IP数据包。

IPRAW套接字提供了一种比普通TCP和UDP套接字更底层、更灵活的网络编程接口，可以用于实现一些与网络协议相关的高级功能，如路由、防火墙等。同时，IPRAW套接字还可以用于网络安全领域的研究和实验，如网络扫描、漏洞分析等。

在iprawsock.go中，主要实现了以下几个功能：

- IPRAW套接字的创建和关闭：通过调用系统API创建原始套接字，并返回一个IPConn类型的对象，用于发送和接收数据。
- 套接字选项的设置和获取：可以设置和获取IPRAW套接字的各种选项，如协议版本、是否启用多播等。
- 数据的发送和接收：可以向指定的目标IP地址和端口发送原始IP数据包，也可以从套接字中接收到来自指定地址和端口的IP数据包，并进行解析和处理。

总之，iprawsock.go文件实现了一组API，用于方便地进行IPRAW套接字编程，同时结合其他net包中的功能，可以实现丰富的网络编程应用程序。




---

### Structs:

### IPAddr

IPAddr结构体是一个包含IP地址和端口号的结构体，用于表示一个网络地址。它的定义如下：

type IPAddr struct {
    IP   IP
    Zone string // IPv6 scoped addressing zone
}

其中，IP是一个IP类型的结构体，用于表示IP地址。Zone是IPv6地址中的可寻址范围标识符。IPAddr结构体可以通过调用ResolveIPAddr函数获取，例如：

addr, err := net.ResolveIPAddr("ip4", "127.0.0.1")

IPAddr结构体主要用于网络编程中的IP通信，可以用于表示本地或远程的IP地址和端口号。它经常用于创建TCP连接、发送UDP数据包、获取本地IP地址等操作。在网络通信中，IPAddr结构体还常常和net包中的其他结构体一起使用，如net.Dial、net.Listen等。



### IPConn

IPConn是net包中表示IP层面数据传输连接的结构体，它继承自net.PacketConn。IPConn类型提供ReadFromIP()和WriteToIP()方法，可以用于从IP层面读取或写入数据。

IPConn结构体的实例可以通过以下函数创建：

func IPConn(fd uintptr) *IPConn

其中fd参数是底层网络文件描述符。

此外，IPConn结构体也实现了PacketConn接口中定义的所有方法：

func (c *IPConn) ReadFrom(b []byte) (int, net.Addr, error)
func (c *IPConn) WriteTo(b []byte, addr net.Addr) (int, error)
func (c *IPConn) Close() error
func (c *IPConn) LocalAddr() net.Addr
func (c *IPConn) SetDeadline(t time.Time) error
func (c *IPConn) SetReadDeadline(t time.Time) error
func (c *IPConn) SetWriteDeadline(t time.Time) error

其中，ReadFrom()和WriteTo()方法是从PacketConn继承而来的，可以用于在IP层面读写数据。LocalAddr()方法返回本地地址，Close()方法用于关闭连接，SetDeadline()、SetReadDeadline()和SetWriteDeadline()方法用于设置读写超时。

在实际使用中，IPConn结构体常用于实现自定义的IP层数据传输协议。可以通过创建IPConn实例，并使用ReadFromIP()和WriteToIP()方法，实现自定义的协议的数据读写操作。



## Functions:

### Network

在iprawsock.go文件中，函数Network的作用是返回一个IP层原始数据报套接字的协议名称，通常是"ip"或"ipv6"。

该函数的定义如下：

```
func (so *ipRawSock) Network() string {
    if so.family == syscall.AF_INET6 {
        return "ipv6"
    }
    return "ip"
}
```

它会返回字符串"ip"或者"ipv6"，这个名称表示创建的IP层原始数据报套接字的协议名称。如果创建的套接字是IPv6的，则返回"ipv6"，否则返回"ip"。

Network函数是一个方法，有一个名为so的指向ipRawSock类型的指针作为接收者，所以在调用该函数时需要通过该类型的实例来调用它，例如：

```
rawConn, err := net.ListenIP("ip:tcp", &net.IPAddr{IP: net.ParseIP("127.0.0.1")})
if err != nil {
    log.Fatal(err)
}
defer rawConn.Close()

network := rawConn.(*net.IPConn).Network()
fmt.Println(network)
```

在这个例子中，我们创建了一个IP层原始数据报套接字，并将套接字绑定到本地地址127.0.0.1上。然后我们将该套接字转换为net.IPConn类型，并通过调用Network方法来获取创建的套接字协议名称。在这种情况下，网络名称将是"ip"。



### String

`String`是一个go语言的内置函数，用于将某个对象转换为字符串表示形式。在`iprawsock.go`文件中，`String`函数是一个为`IPRawConn`类型定义的方法，用于返回`IP`原始套接字连接的字符串表示形式。

`IPRawConn`类型表示一个`IP`原始套接字连接，它可以通过`IPConn()`函数创建，并且可以直接使用`WriteTo()`和`ReadFrom()`方法进行数据的读写操作。

而`String`函数给出的字符串表示形式将包含连接的本机IP地址和端口号，以及远程IP地址和端口号。这样可以方便用户在调试期间查看当前连接的信息，以便更好地了解和处理数据的传输。

具体而言，`String`函数返回的字符串的格式为：`{local-address} -> {remote-address}`，其中`{local-address}`和`{remote-address}`都是包含IP地址和端口号的字符串表示形式。

总之，`String`函数的作用是返回`IPRawConn`类型连接的字符串表示形式，以方便调试和检查连接状态。



### isWildcard

isWildcard这个函数的作用是判断一个IP地址是否是通配符地址（通配符地址是指IP地址中某一段以0表示，如0.0.0.0或者192.168.0.0/16）。

具体的实现逻辑为，首先判断IPv4地址是否为0或255，如果是则返回true，因为这两个地址是保留的特殊地址。如果不是，再判断是否是子网掩码形式的地址，如果是则返回true，因为这种地址也是通配符地址。最后，再判断IP地址是否为0.0.0.0或者127.0.0.1，如果是则返回true。

这个函数通常用于网络编程中的一些特殊场景，例如网络套接字创建时需要指定本地IP地址。如果指定的是通配符地址，则表示可以使用任意可用的地址，通常会被系统自动分配一个可用的地址。



### opAddr

在go/src/net/iprawsock.go中，opAddr是一个函数，它的作用是将一个网络地址转换为通用地址。在IPRawConn类型的方法中经常使用它来处理源和目的地址。

具体而言，opAddr函数的实现基于地址的类型，它将IPv4和IPv6网络地址转换为通用地址。然后，它使用网络字节序将地址转换为字节串，并返回一个IPAddr类型的结构体，请看下面的代码：

```
func opAddr(a net.Addr) (net.Addr, error) {
    switch x := a.(type) {
    case *IPAddr:
        return &UDPAddr{IP: x.IP}, nil
    case *UDPAddr:
        return &UDPAddr{IP: x.IP}, nil
    case *TCPAddr:
        return &TCPAddr{IP: x.IP}, nil
    default:
        return nil, &OpError{Op: "write", Net: a.Network(), Addr: a, Err: errInvalidAddr}
    }
}
```

如果传入的地址是IPAddr、UDPAddr或TCPAddr类型的，则opAddr函数会返回相应的通用地址类型。如果传入的地址类型无效，则该函数将返回一个错误。

总之，opAddr函数是一个用于处理网络地址的工具函数，在处理IPRawConn等类型的网络连接时非常有用。它将不同类型的地址转换成一个通用地址，这在实现网络连接时是很常见的需求。



### ResolveIPAddr

ResolveIPAddr是用于将IP地址字符串解析为IP地址类型的函数。它将传递的字符串参数解析为IP地址类型，并创建并返回一个IPAddr类型的实例，该实例包含了解析的IP地址和端口。如果传递的字符串参数不是有效的IP地址，函数将返回一个错误。

在网络编程中，需要将字符串形式的IP地址转换为可操作的IP地址类型。例如，在使用RawSocket时，可能需要将本地主机或远程主机的IP地址转换为IPAddr类型，以便可以创建RawConn或RawPacketConn。

以下是ResolveIPAddr函数的签名和参数说明：

```go
func ResolveIPAddr(net, addr string) (*IPAddr, error)
```

其中net参数是网络类型，如"ip"或"ip6"，addr参数是表示IP地址的字符串。

此外，ResolveIPAddr还可以解析主机名并返回其对应的IP地址。如果传递的addr参数是主机名而不是IP地址字符串，则函数将使用DNS解析器查找主机名并返回其IP地址。例如：

```go
ipaddr, err := net.ResolveIPAddr("ip", "google.com")
if err != nil {
    // handle error
}
fmt.Println(ipaddr.IP) // prints "216.58.194.174"
```

上述代码将使用DNS解析器查找"google.com"主机名，并返回其IP地址216.58.194.174的IPAddr类型实例。



### SyscallConn

SyscallConn是net包中的一个方法，它返回一个用于操作系统原始套接字的Conn接口。操作系统原始套接字是一种可以与网络协议栈交互的接口，它可以用来接收和发送任意类型的网络数据包，这些数据包不会经过协议栈的解析和处理。

在iprawsock.go文件中，SyscallConn方法是网络层协议处理中非常关键的一部分。它首先创建一个系统原始套接字，并将其绑定到指定的网络接口和协议类型上。然后，它封装了这个原始套接字，并将其返回给调用方。

返回的Conn接口可以用来实现与操作系统原始套接字进行通信。调用者可以使用它来发送和接收原始套接字数据，同时还可以配置一些选项和一些属性，如超时和缓冲区大小等。

总之，SyscallConn方法提供了对操作系统原始套接字的底层访问，它是实现网络协议和数据包处理的关键步骤之一。



### ReadFromIP

ReadFromIP函数是net包中的一个方法，实现了从IPv4或IPv6原始套接字读取数据的操作。具体功能如下：

1. 从指定的IPv4或IPv6原始套接字读取数据。
2. 该方法将有效负载作为字节切片返回。
3. 返回值中的第二个参数是目的地址，表示从哪个地址接收到数据。
4. 返回值中的第三个参数是错误信息。

该方法的函数签名如下：

func (c *IPRawConn) ReadFromIP(b []byte) (int, *IPAddr, error)

其中：

- c类型是一个指向IPRawConn的指针，表示要读取的IPv4或IPv6原始套接字。

- b是一个字节切片，表示读取操作的接收缓冲区。

- 返回值int表示已读取的字节数。

- *IPAddr表示接收到数据的地址。

- error表示读取操作中可能发生的错误。



### ReadFrom

iprawsock.go中的ReadFrom函数是IP协议原始套接字的读取函数，它的作用是从IP协议原始套接字中读取数据报文，并返回数据报文的来源地址。具体介绍如下：

函数原型：
```go
func (c *IPRawConn) ReadFrom(b []byte) (int, net.Addr, error)
```

参数说明：
- b：用于存储读取的数据的缓冲区。
- 返回值：
  - int：读取的字节数。
  - net.Addr：数据报文的来源地址，类型为net.Addr。
  - error：读取过程中发生的错误。

函数实现：
- ReadFrom函数会调用操作系统提供的底层API从IP协议原始套接字中读取数据报文，并将读取的数据写入参数b指向的缓冲区。
- 返回值中的net.Addr表示数据报文的来源地址，它的实际类型是实现了net.Addr接口的某种地址类型，如IPv4Addr、IPv6Addr等。因为IP协议是无连接的，所以在接收数据报文时需要指定来源地址。
- 如果读取过程中发生错误，ReadFrom函数会返回一个非nil的error对象，表示读取失败的原因。

使用场景：
- IP协议原始套接字通常用于实现一些底层的网络应用，如网络抓包工具、路由器等。在这些应用中需要直接访问IP层数据，因此需要使用IP协议原始套接字进行通信。
- 在这些应用中，读取数据的来源地址通常是非常重要的，因为数据报文中的源IP地址和目的IP地址是网络协议中的重要标识符。因此，在使用IP协议原始套接字时需要使用ReadFrom函数来读取数据报文的来源地址。



### ReadMsgIP

ReadMsgIP函数是用于从IPv4或IPv6原始套接字中读取数据的函数。它采用一个IP消息结构体和一个可选的控制消息结构体作为参数，读取来自原始套接字的数据，并且填充IP消息结构体和控制消息结构体。如果读取的数据过长，会被截断，并且返回的错误值表明了是否发生了截断。这个函数的具体作用如下：

1. 从原始套接字接收原始数据包

2. 将原始数据包解密并转换为IP消息结构体

3. 如果可选的控制消息结构体不为空，则填充控制消息

4. 将IP消息结构体和控制消息结构体传递给调用者

5. 如果读取的数据过长，则返回错误信息，指示数据是否被截断

该函数主要用于通过IP层进行底层网络通信，以便实现更高层次和更灵活的网络应用。例如，可以在此基础上实现基于UDP或TCP协议的自定义传输协议，或者实现网络嗅探和数据包分析工具，以便监控网络流量和故障排除。



### WriteToIP

WriteToIP是net包中的一个函数，它的作用是向指定的IP地址发送数据。具体介绍如下：

函数定义：

```go
func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error)
```

函数参数：

- b：要发送的字节切片
- addr：要发送到的目标IP地址

函数返回值：

- sent：发送出去的字节数
- err：发送过程出现的任何错误

函数说明：

- WriteToIP函数可以向指定的IP地址发送数据，如果目标地址与创建连接时的目标地址不同，则会自动更改目标地址。
- 当写入的数据大小超过IP包大小时，WriteToIP函数会自动分片并发送多个IP包。
- 如果网络环境出现断开连接或者其他错误，WriteToIP函数会返回错误信息。

总结：

WriteToIP函数是net包中的一个用于发送IP数据包的函数。它的作用是向指定的IP地址发送数据，并且可以自动分片处理大数据包。该函数在网络编程中使用广泛，特别是在低层网络编程中，比如网络协议实现等场景。



### WriteTo

WriteTo是一个函数，它定义了一个RawConn的方法，用于将IPv4或IPv6协议数据包写入RawConn中。它的函数签名为：

func (c *RawConn) WriteTo(p []byte, addr syscall.Sockaddr) (int, error)

其中，p是待写入的数据包，addr是目标地址。

具体来说，WriteTo函数的作用是将数据包发送给网络协议栈，以便协议栈转发给目标主机。通过RawConn，我们可以底层地控制数据包的发送过程，包括设置IP头部的TTL值、选择IP头部的协议类型等。当然，这也要求我们对底层网络协议有一定的了解。

需要注意的是，WriteTo函数仅适用于IPv4和IPv6协议数据包，其他协议（如ICMP、TCP、UDP等）需要使用不同的函数进行发送。另外，WriteTo函数虽然提供了更精细的网络控制能力，但相对来说也比较繁琐，需要开发者具备相关的网络知识和经验。



### WriteMsgIP

WriteMsgIP是一个函数，用于将数据通过一个已经存在的IP连接(指定的IP地址和端口)发送出去。

具体来说，函数的参数包括：

- c：已经存在的IP连接
- cm：一个ControlMessage类型的指针，用于指定数据包的各种控制信息，比如TTL等
- dst：发送目的地的IP地址
- src：发送源的IP地址
- b：需要发送的数据
- oob：对应数据的Out-of-band数据，需用Unix的sendmsg特性传递出去
- to： 发送数据的目标地址

WriteMsgIP函数主要用于将数据通过Raw IP连接发送到指定IP地址，常用于一些高级网络操作、网络安全和数据包分析等场景。连接可以是TCP、UDP或ICMP连接。在底层，该函数调用了操作系统提供的底层网络接口，将待发送的数据包装并通过网络发送。

需要注意的是，由于Raw IP连接发送数据包时，完全依赖程序发送，因此需要自行指定一些必要的控制信息，并且需要具备一定的网络技术知识和对于底层网络接口的了解，才能够正确的使用该函数，并且实现特定的网络通信方式。



### newIPConn

newIPConn是一个用于创建新的IPConn实例的函数。IPConn是一个代表IP层网络连接的结构体，它实现了net.Conn接口。当我们需要在IP层上直接通信时，可以使用IPConn来进行数据传输。

newIPConn函数接受两个参数：一个是协议类型（即IP协议类型），另一个是本地地址。通过这两个参数，newIPConn可以创建一个新的IPConn实例，并将其返回给调用者。

在newIPConn函数内部，它先会通过net.ListenIP函数创建一个IP网络监听器（IPConn实际上是基于一个监听器进行创建的）。然后，它会调用监听器的Accept函数，将其返回的连接转换成一个IPConn实例，并将其返回给调用者。

总之，newIPConn函数的作用就是用于创建一个新的IPConn实例，它会在底层调用net.ListenIP和Accept函数，并将其返回的连接转换成一个IPConn实例。这样，我们就可以在IP层上进行直接通信。



### DialIP

IPrawsock.go文件是Go语言标准库中net包的子包之一，它包含了支持原始IP/socket的网络连接的函数。其中的DialIP函数是用于创建并返回一个UDP连接的函数，下面是该函数的签名：

func DialIP(netProto string, laddr, raddr *IPAddr) (*IPConn, error)

该函数接受三个参数：

netProto：表示网络协议类型，目前仅支持IPV4和IPV6协议。

laddr：表示本地IP地址，如果为nil，则系统会自动选择一个本地IP地址。

raddr：表示远端IP地址和端口号。

该函数的作用是创建一个UDP连接，该连接是通过IP协议进行通信，而不是TCP或UDP协议。

当用户调用DialIP函数时，系统会创建一个IPConn对象，并使用参数指定的本地IP地址和远程IP地址和端口号建立连接。IPConn对象表示一个UDP连接，可以通过Write方法向远程端发送数据，并通过Read方法从远程端接收数据。

使用DialIP函数建立的UDP连接可以直接发送原始IP包，也可以接收原始IP包，这使得用户可以更加细粒度地控制网络通信。但是，需要注意的是，使用DialIP函数建立的UDP连接可能会不受限制地接收所有的原始IP包，包括不符合协议规范或欺骗的IP包，这可能会导致网络安全问题，因此需要谨慎使用。



### ListenIP

func ListenIP(netProto string, laddr *IPAddr) (*IPConn, error)

ListenIP函数在指定的本地IP地址上监听指定的网络协议（如"ip:icmp"）。返回一个IPConn类型的指针和一个错误类型的值，其中IPConn类型的指针可以用于读写IP数据包，错误类型的返回值表示监听是否成功。

该函数接收两个参数，第一个参数是需要监听的网络协议，必须是小写字母开头的字符串，格式为"[protocol]:[icmp|icmp6|ip|ip4|ip6]”其中protocol是"ip"、"ipv6"等网络协议，而icmp、icmp6、ip、ip4、ip6指代不同的网络层协议。第二个参数是一个IPAddr结构体指针，表示需要监听的本地IP地址和端口号。如果端口号为0，则由系统自动选择一个空闲端口。

ListenIP函数会返回IPConn类型的指针和一个error类型的errno，如果没有错误，则errno为nil。如果函数返回时errno不为nil，则表示监听失败，其中errno包含了错误的具体信息。

使用ListenIP函数可以在指定本地IP地址上监听指定网络协议，实现IP层的数据通信。可以使用返回的IPConn类型的指针进行数据的读写操作。



