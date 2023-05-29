# File: interface_linux.go

interface_linux.go 文件是在 Linux 操作系统上实现网络接口相关操作的 Go 代码文件。它定义了一个结构体类型，代表了一个 Linux 网络接口，并提供了一些方法来实现对网络接口的操作，包括：

1. 获取网卡列表：该文件提供了一个函数 `interfaces`，可以返回当前系统上所有的网络接口列表。

2. 获取网络接口的属性：该文件提供了一个函数 `interfaceAddr`，可以返回指定网络接口的 IP 地址和子网掩码等属性。

3. 设置网络接口的属性：该文件提供了一个函数 `setInterfaceAddr`，可以给指定的网络接口设置 IP 地址和子网掩码等属性。

4. 获取网络接口的 MAC 地址：该文件提供了一个函数 `interfaceByIndex`，可以返回指定网络接口的 MAC 地址。

5. 开启或关闭网络接口：该文件提供了一个函数 `setInterfaceUp`，可以将指定的网络接口开启或关闭。

除了提供上述功能外，该文件还提供了一些底层的系统调用和内核数据结构的封装，使写网络相关的应用程序更加方便和高效。

## Functions:

### interfaceTable

interfaceTable是Go语言net包中的一个函数，在Linux平台下实现。这个函数的作用是获取网络接口列表，并将其存储在一个结构体数组中返回。该结构体中包含了网络接口的名称、地址、MAC地址等信息。

具体来说，interfaceTable函数会调用系统函数netlinkSocket以获取Linux内核中的网络接口信息。然后，它将这些信息转换成Go语言中的类型，并返回一个包含所有网络接口的结构体数组。

这个函数在Go语言中的net包中被广泛使用，尤其是在网络编程中经常用到。通过获取网络接口列表，网络编程可以更加方便地进行网络信息的获取和处理。



### newLink

newLink函数的作用是根据设备名称创建一个新的Netlink socket连接，并在该连接上监听与网络接口相关的消息。

在Linux系统中，内核通过Netlink协议向用户层应用程序提供了一种机制来获取系统状态信息，其中包括与网络接口相关的状态信息。newLink函数就是通过Netlink协议监听与网络接口相关的消息。

该函数的具体实现流程如下：

1. 创建一个新的Netlink socket连接，可通过该连接与内核进行通信；

2. 设置Netlink连接的文件描述符为非阻塞模式；

3. 向内核发送一个Netlink请求，请求内核通知用户层应用程序与网络接口相关的状态信息；

4. 轮询Netlink连接的文件描述符，等待内核发送响应消息；

5. 接收内核发送的Netlink响应消息，并解析消息中的网络接口相关的状态信息；

6. 对解析得到的状态信息进行处理，比如添加或删除网络接口。

总之，newLink函数的主要作用是建立一个用于获取网络接口状态信息的Netlink socket连接，并监听该连接上的消息，从而实现对网络接口状态的实时监控。



### linkFlags

在 Go 语言中，linkFlags 这个函数是在实现网络接口时用到的。它主要用于为 Linux 平台上的网络接口设置链接标志。

具体来说，linkFlags 函数会返回一个包含特定网络设备相关的链接标志的字符串。这些标志可以用于控制网络设备的行为，例如启用或禁用特定功能，以及定义其他网络接口参数。

linkFlags 函数实现基于 Linux 平台的网络接口，因此它会检查平台是否为 Linux，并在其中调用一些 Linux 特定的函数，以设置特定的链接标志。

总之，linkFlags 函数是一个在网络接口实现中用于设置特定链接标志的功能。它与 Linux 平台有关，并提供了一些特定于平台的功能。



### interfaceAddrTable

interfaceAddrTable函数主要用于获取给定网络接口的IP地址表。该函数的实现利用了Linux内核实现的netlink socket API，它可以获取Linux内核内网络接口和地址数据的各种详细信息，包括接口名称、MAC地址、IPv4地址、IPv6地址等。

具体来说，interfaceAddrTable函数首先创建一个netlink.Conn类型的netlink连接对象，通过这个连接对象与Linux内核进行通信。之后，它调用netlink.Dial函数来与内核通信，然后设置消息头部和过滤器，最后发送消息获取与给定接口相关的网络地址数据。

interfaceAddrTable函数返回一个Addrs类型的列表，其中每个元素均为该接口下的IP地址。这些IP地址不仅包括IPv4和IPv6地址，还包括IPv4和IPv6的广播地址、点对点地址等。

该函数的主要作用是充分利用内核提供的API接口，快速获取网络接口数据。在网络编程中，我们很常常需要获取本机的网络接口和地址信息，并对它们进行操作，例如设置SOCKET选项等，该函数就可以用来满足这些需求。



### addrTable

在go/src/net/interface_linux.go文件中，addrTable是一个函数，其目的是返回系统中的网络接口地址列表。

更具体地说，addrTable首先使用syscall包中的getifaddrs函数来获取系统中的接口列表。然后，它对每个接口进行循环处理，并使用syscall包中的getnameinfo函数来获取每个接口的IP地址和端口号。

最后，每个IP地址都被封装在一个结构体中，并将所有结构体存储在一个列表中，然后将该列表作为函数的返回值。

该函数的作用是提供一个简单的方法来获取系统中的网络接口地址列表，这对于实现高级网络应用程序非常有用。



### newAddr

在go/src/net/interface_linux.go文件中，newAddr函数是用来创建一个新的IP地址类型的对象，其功能是将传入的IP地址（IPv4或IPv6）封装到一个net.IP类型的对象中，以便可以被传递给网络协议的相关函数使用。

newAddr函数的定义如下：

```
func newAddr(ip net.IP) sockaddr {
  if ip.To4() != nil {
    sa := new(syscall.SockaddrInet4)
    copy(sa.Addr[:], ip.To4())
    return sa
  }
  sa := new(syscall.SockaddrInet6)
  copy(sa.Addr[:], ip.To16())
  return sa
}
```

其中，sa是一个syscall.Sockaddr类型的对象，表示一个网络地址。如果传入的IP地址是IPv4类型，那么就创建一个syscall.SockaddrInet4类型的对象，将IP地址存储到其Addr字段中；如果传入的IP地址是IPv6 类型，那么就创建一个syscall.SockaddrInet6类型的对象，将IP地址存储到其Addr字段中。最终函数返回这个对象，供调用者使用。

总之，newAddr函数的作用是将一个IP地址封装成一个Sockaddr类型的对象，以便于网络协议相关函数的使用。



### interfaceMulticastAddrTable

interfaceMulticastAddrTable函数的作用是获取本地网络接口的多播地址表。

在网络通信中，多播是一种一对多的通信方式，一个发送者可以同时向多个接收者发送同一份数据。多播地址是用来标识一个特定的多播组的。通过获取本地接口的多播地址表，可以知道本地接口所加入的多播组，并可以向这些组发送数据。

该函数的实现依赖于操作系统中的netlink接口，首先调用genetlink.Dial函数建立一个连接，然后调用netlink.Conn.LookupByFamily获取特定类型（AF_PACKET）的网络接口信息，并通过消息的类型（RTM_NEWADDR）和筛选条件（AddrFlags、IFA_MULTICAST）来查询多播地址信息。

最终，函数返回一个map，其中key为网络接口的索引，value为该接口所加入的多播组的IP地址列表。



### parseProcNetIGMP

parseProcNetIGMP函数是解析Linux系统/proc/net/igmp文件内容的函数，用于获取系统中IGMP组播协议相关的信息。

具体来说，该函数首先打开/proc/net/igmp文件，然后按行读取文件内容，每一行对应一个IGMP成员。对于每个成员，函数将其解析成IGMP成员结构体（struct igmpstat），该结构体包含了IGMP协议相关的一些信息，比如组播地址、成员状态、参与组播的接口等。

最后，函数将所有解析后的IGMP成员结构体添加到一个IGMP group list中，并返回该列表。这个列表可以用于进一步处理IGMP协议相关的信息，比如查询当前系统中哪些网络接口或组播组正在使用IGMP协议。

总之，parseProcNetIGMP函数是一个解析IGMP协议相关信息的函数，可以帮助我们更深入地了解系统中的组播网络配置、IGMP协议状态等信息。



### parseProcNetIGMP6

parseProcNetIGMP6函数是在net包中的interface_linux.go文件中实现的，它的作用是解析/proc/net/igmp6文件，提取出其中的IGMPv6组成员信息。

IGMPv6协议是Internet协议族中的一个组播协议，它允许一个主机或路由器向组播组中发送数据，组播组中所有成员都可以接收到这些数据。/proc/net/igmp6文件记录了主机或路由器加入的IGMPv6组，以及每个组中的成员信息。

parseProcNetIGMP6函数通过读取/proc/net/igmp6文件的内容，解析其中的IGMPv6组成员信息，并将其保存到一个IGMP6Group结构体的数组中。IGMP6Group结构体包含组播组的IP地址、接口索引、成员数量等信息，以及一个IP地址列表，其中存储了组中的成员IP地址。

通过调用这个函数，Go程序可以获取当前主机或路由器加入的IGMPv6组，以及每个组中的成员信息。这些信息可以用于网络监控、配置和故障排查等应用场景。



