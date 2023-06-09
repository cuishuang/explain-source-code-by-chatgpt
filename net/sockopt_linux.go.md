# File: sockopt_linux.go

sockopt_linux.go是Go语言标准库中net包的一个文件，它包含了对于Linux系统socket的选项设置与读取的相关代码。更具体地说，该文件定义了一些函数和变量，使得我们可以在Linux系统下使用Go语言来设置和读取socket相关的选项。

在Linux系统中，socket选项是用来控制socket的行为和属性的一些标志。例如，我们可以通过设置SO_REUSEADDR选项来允许一个端口在他申请者已经退出后立即和新的进程一起绑定。又或者我们可以设置TCP_NODELAY选项来禁用Nagle算法，加快数据传输速度。

所以说，在网络编程过程中，设置和读取socket选项是十分重要的。而在Go语言标准库中，就是由sockopt_linux.go这个文件来完成这些工作的。

具体来看，该文件中定义了SocketConn结构体，它是一个Go语言中socket连接的抽象类型。它包含一个文件描述符fd，以及一些方法（如SetSockOpt、GetSockOpt等）来进行socket选项的设置和读取操作。这些方法实际上是调用了Linux系统调用中的setsockopt和getsockopt函数。另外，该文件还定义了一些相关的常量、变量和错误信息等。

综上所述，sockopt_linux.go文件是Go语言标准库中的net包的一个重要组成部分，负责Linux系统下socket选项的相关操作。

## Functions:

### setDefaultSockopts

setDefaultSockopts函数位于文件sockopt_linux.go中，其作用是在创建一个新的Unix Socket时，设置其默认选项。

Unix Socket是一种进程间通信（IPC）机制，类似于网络Socket，但只用于本地通信。在创建一个新的Unix Socket时，核心库会使用当前进程的默认选项来设置其属性。setDefaultSockopts函数通过将某些选项设置为其最优值来改善Unix Socket的性能和安全性。

这个函数的具体实现包括以下步骤：

1. 设置TCP_NODELAY选项：启用该选项以减少延迟，这对于实时网络应用程序（如VoIP和视频会议）非常重要。

2. 设置SO_REUSEADDR选项：启用该选项可允许一个Socket在关闭后的一段时间内重新使用其IP地址和端口号，从而减少必要的等待时间。

3. 设置IPV6_V6ONLY选项：启用该选项可防止Socket同时使用IPv4和IPv6地址。这有助于避免网络攻击。

4. 设置IP_FREEBIND选项：启用该选项可使系统在所有可用接口上接受传入连接请求，而不仅仅是在默认网关上。

5. 设置TCP_MAXSEG选项：该选项设置Socket可以发送的最大TCP段大小。对于高延迟网络（如卫星链路）来说，这是一项重要的性能调整。

总之，setDefaultSockopts函数的主要目的是自动设置Unix Socket的默认选项，从而提高其性能和安全性。



### setDefaultListenerSockopts

setDefaultListenerSockopts是一个用于设置TCP监听器socket选项的函数。当程序创建TCP监听器时，可以调用此函数来设置一些默认的选项，以确保监听器表现良好并具有良好的性能。

该函数主要做了以下几件事情：

1. 设置TCP选项：将TCP_KEEPIDLE，TCP_KEEPINTVL，TCP_KEEPCNT和TCP_DEFER_ACCEPT选项设置为适当的值，以确保TCP连接处于连接状态，并且垃圾回收以及其他TCP连接管理任务得以正确执行。

2. 设置ReusedPort选项：将SO_REUSEPORT选项设置为1以实现端口复用。这样可以使多个进程使用同一端口进行网络通信。

3. 设置TCP FastOpen选项：如果系统支持TCP FastOpen，并且此选项已启用，则将TCP_FASTOPEN选项设置为适当的值，以在不进行3次握手的情况下加速初始数据传输。

总之，setDefaultListenerSockopts函数会设置一些重要的TCP监听选项，以确保TCP监听器的表现良好，并具有良好的性能和可靠性。



### setDefaultMulticastSockopts

setDefaultMulticastSockopts函数的作用是为多播套接字设置默认的选项。在Linux系统中，多播套接字是一个特殊的套接字，可以用于向多个目标主机发送数据。为了确保多播套接字的正确性和可靠性，需要设置一些选项。

具体来说，setDefaultMulticastSockopts函数会为输入的UDPConn类型的对象设置如下选项：

- IP_MULTICAST_LOOP：控制是否将发送的数据包回送到本地接口。默认为开启。
- IP_MULTICAST_TTL：控制数据包在多播发送时的生存时间。默认为1，即数据包只能传输到当前网段内。
- IP_MULTICAST_IF：控制多播数据的发送接口。默认为本地默认路由。

这些选项可以通过设置UDPConn对象的MulticastLoopback、MulticastTTL和MulticastInterface方法来改变。setDefaultMulticastSockopts函数的作用就是为确保这些方法的默认值正确。



