# File: switch.go

switch.go文件是Go语言中net包中的一个文件，该文件主要用于实现网络连接的多路复用器（Multiplexer）。

作为网络编程中的常见需求，程序需要同时监听多个连接，以便立即处理任何可读、可写或错误事件。该文件中的多路复用器可同时管理多个网络I/O事件，接收到事件后将其发送到合适的处理器中。当网络事件到达时，多路复用器调用相应的处理器，以便程序可以处理该事件。

switch.go文件的核心是Switch结构体，该结构体包括了许多成员变量和方法，用于管理网络连接。其中最重要的是Add、Remove和Poll方法。Add方法用于向多路复用器中添加一个新的连接，Remove方法用于从多路复用器中删除一个连接，Poll方法用于超时轮询网络事件并返回活动的网络连接。

多路复用器在网络编程中扮演着至关重要的角色，能够提高程序的效率和性能。switch.go文件中的多路复用器代码是Go语言中实现高效网络编程的关键部分之一。




---

### Structs:

### Switch

Switch是一个结构体类型，定义在go/src/net/switch.go文件中。该结构体用于网络转发根据网络特征进行选择的机制。

网络通信中的转发是指将数据包从一台设备或网络节点传输到另一台设备或网络节点的过程。在实际情况中，数据包通常需要经过多个路由器或交换机进行转发才能到达目的地。网络转发的根本目的是将数据从一个网络子网传输到另外一个网络子网，使数据能够到达目标主机。

Switch结构体定义了网络转发机构体，它管理网络所有接口的状态和网络包的传输。Switch结构体中包含了多个成员变量和成员方法，通过这些成员变量和成员方法，可以实现网络数据包从一个网络子网传输到另外一个网络子网的转发过程。 

Switch结构体的主要作用包括：

1. 提供了一种基于网络中的通道进行最佳路径选择的机制，它能够根据检测到的异常条件选择最优的路径。

2. 网络管理，包括维护网络接口和接口状态，处理接口的配置信息，控制网络流量等。

3. 根据接口、流量、数据类型等特征进行转发选择，确保数据包的高速、高效和可靠地到达目标。

4. 支持多种转发技术，如虚拟局域网、链路聚合、虚拟路由器等。

5. 维护与本地和远程设备通信所需的距离矩阵，通过计算这些距离选择最短的路径。

总之，Switch结构体是网络转发机制的核心所在，它协调了各种不同层面的网络状况，确保数据能够快速、可靠、安全地传输。



### Cookie

在go/src/net中的switch.go文件中，Cookie结构体是用于解析HTTP请求头中的Cookie字段内容的。当客户端发送请求时，会将请求中的Cookie字段信息发送给服务端，Cookie结构体可以帮助我们解析出Cookie字段中包含的所有cookie值。

Cookie结构体的定义为：

```go
type Cookie struct {
    Name  string
    Value string

    Path       string    // optional
    Domain     string    // optional
    Expires    time.Time // optional
    RawExpires string    // for reading cookies only

    // MaxAge=0 means no 'Max-Age' attribute specified.
    // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
    // MaxAge>0 means Max-Age attribute present and given in seconds
    MaxAge   int
    Secure   bool
    HttpOnly bool
    SameSite SameSite // Go 1.11

    Raw string
    Unparsed []string // Raw text of unparsed attribute-value pairs
}
```

Cookie结构体包含了名称、值、路径、域名、有效期等信息，并提供了一些方法来设置这些值。使用Cookie结构体可以帮助我们简化处理cookie的过程，并提高代码的可读性和复用性。



### Status

在net包中，Status结构体是用于表示网络连接状态的。具体来说，它用于描述从底层网络库（比如操作系统的网络API）返回的错误信息的状态码和原因，以便应用程序能够更好地理解网络连接的状态。

Status结构体包含了以下字段：

- Code int：表示错误状态码，通常是一个非负整数；
- Message string：表示错误的人类可读的描述信息；
- Temporary bool：表示错误是否是暂时性的，如果是，则应用程序可以尝试重新连接或重新发送请求以解决问题；如果不是，则应用程序需要更彻底地处理错误。

Status结构体还定义了一些方法，如IsTimeout()和IsTemporary()等，用于方便地判断错误是否是超时、暂时性等特定类型的错误。

在net包中，许多函数和方法都会返回一个Status结构体以表示网络连接的状态。应用程序可以根据这个结构体来处理连接错误、重新连接等操作。



### Stat

在go/src/net中，switch.go文件中定义了一个叫做Stat的结构体。这个结构体在net包中用于表示网络地址的状态信息。

Stat结构体包含了以下成员：

- Count：表示当前连接的数量；
- Bytes：表示当前连接的字节数；
- Errors：表示当前连接的错误数；
- Milliseconds：表示当前连接的毫秒数。

通过这些状态信息，我们可以对当前网络连接的情况有一个更加详细的了解。在进行网络编程时，这些信息可以帮助我们进行网络优化和监控。

Stat结构体的作用类似于Linux中的netstat命令，可以帮助我们实时查看网络连接状态。在Go语言的net包中，Stat结构体是一个非常有用的工具。



### stats

在go/src/net中，switch.go这个文件中stats结构体的作用是用于统计网络连接的数量和状态，包括正在使用的连接数、连接的总数、连接的最大数、空闲连接的数量等等。

stats结构体的定义如下：

```go
type stats struct {
    maxOpen      int64
    open         int64
    inuse        int64
    idle         int64
    waitCount    int64
    waitDuration time.Duration
}
```

其中，maxOpen表示最大连接数，open表示当前连接数，inuse表示正在使用的连接数，idle表示空闲连接数，waitCount表示等待连接的数量，waitDuration表示等待连接的总时间。

在net包中，通过stats结构体来记录连接池中的连接数和状态，从而可以更好地管理和优化网络连接。例如，在创建新的连接时，如果已经达到了最大连接数，则会将连接放入等待队列中，当有连接释放时，会从等待队列中获取连接。通过这样的方式，可以避免连接池中的连接过多导致性能下降，同时也可以避免连接不足导致请求阻塞等问题。



### FilterType

FilterType是一个枚举类型，定义了一些常量，用于表示网络过滤器的类型。

具体来说，网络过滤器可以由用户自定义，用于在网络连接的不同阶段进行筛选和处理数据包。这个过程和网络中的路由功能类似，是网络协议栈中非常重要的一个组成部分。

FilterType枚举类型提供了一些常量，用于表示不同类型的网络过滤器。将这些常量定义在一个枚举类型中，可以方便过滤器的管理和使用。同时，这种枚举类型也可以作为函数的参数类型，从而在不同的网络过滤器中进行判断和处理。

例如，在网络协议栈的TCP连接建立过程中，可以使用不同的过滤器对数据包进行筛选和处理。如果要添加一个新的过滤器类型，只需要在FilterType枚举类型中添加相应的常量即可，非常方便和灵活。



### Filter

在Go语言的net包中，Filter是一个结构体，用于对网络连接进行过滤。在net包中，有多种连接类型，例如TCP连接、UDP连接、Unix socket等等，而Filter结构体可以根据特定的过滤条件，筛选出符合要求的连接。

Filter结构体中包含以下字段：

- Network：表示所要过滤的网络协议，例如"tcp"、"udp"等等。
- Address：表示所要过滤的网络地址，可以是一个IP地址或者一个主机名。
- Dialer：一个Dialer类型的指针，表示用于发起连接的Dialer实例。

当一个连接被传递到Filter结构体的过滤函数中时，该函数会首先检查连接是否符合所设定的过滤条件。如果符合条件，该函数会将连接返回；否则，该函数会返回nil，表示连接不符合条件。

Filter结构体主要用于网络编程中的连接管理和路由控制等方面。例如，可以使用Filter结构体来过滤掉不符合特定IP地址或端口号的连接，或者通过Filter结构体来实现多个网络链路之间的路由控制等。



### AfterFilter

在go/src/net/switch.go文件中，AfterFilter是一个结构体，它被用作网络交换机中的过滤器。

具体来说，网络交换机是一个网络设备，用于连接多个网络设备，并转发数据包。在这种情况下，过滤器是一种机制，可用于在将数据包转发到目标设备之前，对数据包进行检查或修改。

AfterFilter就是一个过滤器，它是网络交换机中的一个重要组件，用于在数据包转发之后处理一些任务。例如，AfterFilter可以用于收集网络统计信息或检查传输层协议中的错误。

当一个数据包从网络交换机的输入端口进入时，可以通过AfterFilter进行检查或修改，然后将数据包转发到目标设备。在数据包离开交换机的输出端口之前，AfterFilter还可以进行一些操作，例如对数据包进行记录或修改。

总之，AfterFilter是网络交换机中的一个重要组件，用于管理和控制数据包的流动，以确保网络通信的正常和高效运行。



## Functions:

### init

在go/src/net/switch.go中的init函数主要有以下几个作用：

1. 注册网络协议：
init函数中首先会使用net包提供的RegisterProtocol函数，将TCP、UDP等网络协议注册到net包的protocolMap中。这些协议通过这种方式注册后，可以在后续的网络编程中直接使用。

2. 注册地址解析协议：
除注册网络协议外，init函数还会使用net包提供的RegisterAddrResolver函数，将ipv4、ipv6地址解析协议注册到net包中的addrResolvMap中。这些协议注册后，可以在后续的网络编程中被直接调用。

3. 初始化出站网络类型：
在init函数中会通过net包提供的newDefaultResolver函数，获取系统默认的DNS解析器，并将其保存到net包的库中。同时，也会初始化出站网络类型，使得net.Dial等函数可以直接使用。

4. 加载网络环境变量：
在init函数中，通过net包提供的parseProxyEnv函数，加载系统的代理设置，并将其保存到net包中的proxyMap中。这样，在后续的网络操作中，可以自动使用系统设置的代理服务器。

总体来说，init函数在net包的初始化过程中，扮演了非常重要的角色。它完成了注册网络协议、初始化DNS解析器和出站网络类型、加载代理设置等任务。这些操作都在包初始化时自动完成，用户无需手动操作。



### Stats

在go/src/net中的switch.go文件中，Stats()是一个函数，用于获取网络I/O的统计信息，它返回一个net或者是netconn结构体，其中包含了许多值，如字节数、错误数、读取超时数、写入超时数、读取阻塞数、写入阻塞数，以及其它一些统计信息。

这些统计信息对于网络应用的性能分析和故障排除非常有用，可以帮助我们追溯数据包的丢失和传输性能的瓶颈。例如，在网络故障的排查过程中，可以通过比较两个时间点的统计信息，来判断网络状况的变化以及I/O的瓶颈位置。

此外，Stats()函数还可以为各种不同的网络协议提供不同的统计信息，方便开发者针对不同的网络协议进行调试和优化。

总之，Stats()函数是一个非常有用的工具，有助于开发者更加深入地了解和优化网络I/O的性能。



### Sockets

在go/src/net/switch.go文件中，Sockets函数是用于检查并返回所有当前已打开的套接字的列表，也就是获取所有已连接的网络连接。它的实现涉及到系统调用和底层协议栈，因此它的逻辑比较复杂。具体来说，该函数会将每个活动的套接字对象封装成一个Socket对象，并将所有活动的套接字对象放入一个数组中返回。

这个函数允许应用程序检查当前运行的网络连接的状态，以便更好的了解程序的网络通信情况，从而更好的进行调试和优化。该函数常用于对网络负载均衡器或其他网络服务器进行自动发现和配置。通过查看套接字的类型、地址和端口等信息，可以确定是哪些应用程序、服务或主机正在运行和通信，从而帮助开发人员更好的了解和优化网络通信的性能。



### Family

在go/src/net/switch.go中，Family是一个返回网络地址的地址族的函数。地址族定义了特定类型的网络地址，例如，IPv4和IPv6地址使用不同的地址族。

具体来说，Family函数接收一个网络地址（net.Addr）作为参数，并返回该地址的地址族（int类型）。地址族由最高位为1的无符号整数表示，例如，IPv4地址族为1，IPv6地址族为28。

Family函数使用了类型断言，以确定传入的地址类型。如果传入的地址为IPv4地址，则返回1，如果传入的地址为IPv6地址，则返回28。如果传入的地址不是IPv4或IPv6地址，则返回0。

这个函数的作用在于，它可以帮助网络编程者判断一个网络地址的类型，以便在不同的情况下使用不同类型的网络协议和处理方式。例如，在客户端编程中，如果需要连接到一个IPv6地址的服务器，就需要使用IPv6协议来建立连接。而Family函数就可以帮助我们确定网络地址的类型，从而决定使用IPv4协议还是IPv6协议。



### Type

在net包中，switch.go文件中的Type()函数用于获取给定错误的类型信息。它接受一个err类型的参数，并返回一个表示该错误类型的字符串。

Type()函数实现了error类型的Type()方法，即它是一个满足error接口的函数，可以用于处理错误信息。当进行网络编程时，可能会遇到各种各样的错误情况，如连接超时、网络不可达、地址已在使用等，这些错误信息都可以通过Type()函数进行分类和处理。

Type()函数会根据错误类型的不同返回不同的字符串，常见的字符串包括：

- "timeout"：表示连接超时或读写超时
- "temporary"：表示临时错误，可能会恢复正常
- "connection refused"：表示连接被拒绝
- "no route to host"：表示无法到达目标主机
- "address already in use"：表示地址已经被占用

通过这些字符串信息，程序可以根据错误类型采取不同的处理方式，也可以根据这些错误类型进行错误分析和定位。



### Protocol

在 go/src/net/switch.go 中，Protocol 函数是一个辅助函数，它用于将 TCP/UDP 协议的名称(例如"tcp4"、"udp6"等)转换为其对应的协议编号(例如IPPROTO_TCP、IPPROTO_UDP等)。

Protocol 函数的声明如下：

func Protocol(name string) (proto int, err error)

该函数接受一个名称参数并返回其对应的协议编号和可能发生的错误。如果找不到名称对应的协议编号，函数返回错误。

该函数内部实现使用了一个 switch 语句来对不同的名称做出不同的处理。例如，对于名称为 "tcp" 的协议，该函数返回常量值 IPPROTO_TCP，表示 TCP 协议的协议编号。对于名称为 "udp" 的协议，该函数返回常量值 IPPROTO_UDP，表示 UDP 协议的协议编号。此外，该函数还支持其他一些协议，如 SCTP、ICMP、IPv6 等。

由于网络编程中经常使用协议编号而不是协议名称，因此 Protocol 函数非常有用，可以方便地将协议名称转换为协议编号。



### cookie

cookie函数位于net包中的switch.go文件中。它的作用是解析和封装HTTP cookie。HTTP cookie是一小段服务器发送到用户Web浏览器并存储在用户计算机上的数据，以便下次访问时使用。

具体来说，cookie函数接收以下参数：

- header：HTTP请求头信息
- key：要解析的cookie的名称

函数的目的是在header中查找指定名称的cookie，并解析它的值。如果找到了匹配的cookie，则返回其值；否则返回一个空字符串。如果用户提供了一个可选的“基准时间”参数，函数还会解析cookie中的过期时间，并将其转换为一个时间点。如果cookie已过期，则返回一个零时间值。

cookie函数还提供一个名为SetCookie的嵌套函数，用于创建和封装一个新的HTTP cookie。该函数接收以下参数：

- w：HTTP响应的写入器接口
- name：cookie的名称
- value：cookie的值
- maxAge：cookie的过期时间（以秒为单位。如果为负数，则表示立即过期）
- path：cookie的服务器路径
- domain：cookie的服务器域
- secure：是否只在HTTPS连接上发送cookie
- httponly：是否在HTTP中向客户端发送cookie

SetCookie函数将创建一个新的HTTP Set-Cookie头信息，并将响应写入w。这个头信息将告诉浏览器在下次访问时发送一个新的cookie。这样，服务器和浏览器就可以在多个HTTP请求/响应周期之间共享数据。



### String

在Go语言中，所有的类型都可以实现String()方法，这个方法被称为字符串化方法。String()方法返回该类型的字符串形式。在switch.go文件中的String()方法是将Type类型转换成字符串形式。这个方法的作用是将类型转换成可读性更好的字符串形式，可以方便地进行调试和日志记录。

在switch语句中，对于每一个case子句，都需要遵循一个重要的规则，即case的表达式必须是常量，且类型必须与switch表达式的类型相同。所以，在将Type类型作为case的表达式时，需要将它转换成字符串形式，才能与其他字符串类型的常量进行匹配。

通过实现String()方法，可以将Type类型转换成字符串形式，从而在switch语句中使用。例如，在switch.go文件中的实现中，将Type类型转换成了字符串形式，以便与其他字符串类型的常量进行匹配。这个方法使用了fmt.Sprintf()函数将Type类型的信息格式化成字符串形式，以便更好地查看和调试。

总之，String()方法是将一个类型转换成字符串形式，方便调试和日志记录，同时也是让类型能在switch语句中使用的必要手段。



### String

在 go/src/net 中的 switch.go 文件中，String 函数是用于将一个TCP/IP协议常量的整数值（如IPV4, IPV6, TCP, UDP等）转换为其对应的字符串，以便于打印和输出时的易读性。

```go
func (p Protocol) String() string {
    switch p {
    case IP:
        return "ip"
    case ICMP:
        return "icmp"
    case TCP:
        return "tcp"
    case UDP:
        return "udp"
    case IPv6HopByHop:
        return "ipv6-hopbyhop"
    case IPv6Routing:
        return "ipv6-routing"
    case IPv6Fragment:
        return "ipv6-fragment"
    case IPv6ICMP:
        return "ipv6-icmp"
    case IPv6NoNextHeader:
        return "ipv6-nonxt"
    case IPv6DestinationOptions:
        return "ipv6-destopt"
    default:
        return "unknown"
    }
}
```

该函数接受一个类型为 Protocol 的参数 p，该类型是一个枚举类型，表示TCP/IP协议中的不同协议。通过 switch-case 语句，根据不同的协议常量，返回对应的协议名称的字符串。

例如，在打印和输出时，可以使用该函数将协议常量（如 TCP/UDP/IPV4 等）转换为字符串便于阅读和理解。



### getLocked

getLocked是net包中switch.go文件中的一个函数。它的作用是从当前Dispatch的hash表中获取一个Key，该key被视为“锁定”。

在net包中，一组goroutine将负责处理所有的网络连接（或“事件”）。在服务器启动时，网格将配置一个指定数量的goroutine来执行此任务。在这一组goroutine中，一个goroutine负责对于每个socket上读取数据，一个goroutine负责发送数据到每个socket上，其他goroutine负责连接的accept、处理和关闭等操作。但是，任何一个goroutine都可以处理任何一个连接，它们共享输入/输出数据和状态。这个机制被称为goroutine池。

getLocked的作用就是保证每个goroutine在处理连接时不会与其他goroutine产生冲突，通过在一个连接上获取锁来实现。每一个连接会为其分配一个key，该key作为锁使用。在任何时候，只有一个goroutine可以获得一个与特定连接关联的锁，并且只有在持有锁的goroutine执行完连接处理或显式释放锁时，其他goroutine才能获取该锁并处理连接。

getLocked函数首先找到一个和当前goroutine关联的Conn中可用的key，并将其标记为锁定状态。如果没有可用的key，则呈现出繁忙状态。此外，该函数还尝试使用关联到获取的key的Conn作为参数启动goroutine，以在后台运行关联的Conn。此操作可以使Conn和正在获取或保持锁定的goroutine解耦。如果在获取锁和启动goroutine之间发生错误，该锁将被释放，调用goroutine将返回错误。



### apply

apply函数位于go/src/net/ switch.go中，其作用是将每个网络I/O操作执行时的参数与其结果写入Log中。具体来说，apply函数记录了每个网络操作的耗时、参数、结果以及错误信息等，以方便进行调试和分析。

在Go语言的net包中，有很多不同的网络I / O操作可供选择。例如，Dial（）函数用于建立到远程服务器的TCP连接，Listen（）函数用于在某个端口上侦听TCP连接请求，而Read（）和Write（）函数则用于读取和写入TCP连接数据。对每个网络I / O操作，apply函数将其参数与结果记录在日志中，以便更轻松地了解每个操作的执行情况。

当您调用net包中的任何函数时，apply函数会被调用，以记录该操作的参数和结果。例如，当您调用Dial（）函数时，apply函数将记录建立TCP连接所需的时间，以及连接是否成功。它还会记录途中发生的任何错误或异常情况。

总的来说，apply函数是Go语言内置的一种工具，用于记录网络I/O操作的执行情况和结果，以便更好地进行调试和问题排查。



### apply

在go/src/net/switch.go文件中，apply函数的作用是基于适当的协议和源/目标地址，将传入的IP数据包分配给正确的协议处理程序进行处理。该函数负责路由IP数据包，这些数据包可以是IPv4或IPv6，并将它们分发到正确的协议处理程序中。

具体来说，apply函数会读取传入的IP数据包的目标地址，并将其与已知的协议路由表中的条目进行匹配。如果存在匹配项，则将数据包分配给相应的协议处理程序。如果没有找到匹配项，则尝试使用默认路由进行匹配。

在apply函数中，还有一个称为midChan的中间通道。这个通道用于在路由数据包到适当协议的处理程序之前缓存数据包。这个通道的作用是防止可能发生的协议处理程序阻塞或延迟处理数据包。

总而言之，apply函数是net包中重要的路由数据包的方法，它负责将IP数据包路由到正确的协议处理程序进行调度处理，并通过中间通道从数据包源端缓存和缓存数据包，以防止协议阻塞或延迟。



### Set

Set函数是net包中的一个函数，它位于src/net/switch.go文件中。该函数主要用于配置网络接口的属性，以及更新网络接口的状态。它是Switch类型的方法之一，Switch是一个表示网络交换机的结构体，用于控制和管理网络端口的转发表、流量控制等相关功能。

具体来说，Set函数的作用有如下几个方面：

1. 配置网络接口的属性：例如设置MTU（最大传输单元）、MAC地址、IP地址、子网掩码等。

2. 更新网络接口的状态：例如开启/关闭网络接口、更新网络接口的速率、状态等。

3. 控制网络流量：例如实现流量控制、管理网络端口的转发表等。

4. 实现广播/组播功能：例如实现多播路由表的配置和维护等。

总之，Set函数是一个非常重要的函数，它充分体现了Switch结构体的作用和功能，它为网络接口的配置、状态更新和流量控制等方面提供了强大的支持和丰富的功能。



