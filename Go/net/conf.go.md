# File: conf.go

conf.go文件是net包实现网络通信的配置文件。该文件中定义了网络连接的多种配置和选项，如网络层次结构、TCP连接设置、HTTP代理设置和DNS服务器设置等。

具体而言，该文件中包含了一些常量和变量，用于表示网络保障、域名服务器、重试策略等。这些变量可以被其他网络相关程序引用和调用，以实现网络连接的配置。

例如，该文件中定义了一些TCP连接的相关参数，如TCP连接超时时间、TCP KeepAlive间隔和最大空闲连接数等。这些参数可以被指定为选项，并在实际连接时使用，以确保TCP连接的可靠性和性能。

此外，该文件还实现了一些网络通信的相关功能，如解析和处理curl-style URL地址、验证IP地址和端口号等。这些功能可以帮助开发人员更方便地实现网络通信。

总之，conf.go文件对于实现网络通信的程序来说是非常重要的，可以提供网络连接的关键配置和选项，并且实现一系列网络通信的相关功能。




---

### Var:

### confOnce

在go/src/net/conf.go文件中，confOnce是一个sync.Once类型的全局变量，用于确保在程序运行期间confInit函数只被调用一次。

confOnce变量使用了Go语言的sync包中的Once结构体，该结构体能够确保提供给它的方法在程序运行期间只调用一次。当调用Once.Do(f)方法时，它会只执行f函数一次，而在调用Do(f)之后其他调用Do(f)的goroutine将会阻塞，直到f函数被执行完毕。

在conf.go中，confOnce保证了confInit函数只会被执行一次，因此能够保证在Net包的使用中只会初始化一次网络的配置信息，避免了因为重复初始化网络配置而产生的问题和性能浪费。当然，这也意味着如果需要重新加载网络配置，需要重新启动应用程序。

总的来说，confOnce变量在Net包中起到了非常重要的作用，它确保了Net包在运行期间只会进行一次网络配置的初始化，减少了不必要的重复操作和资源浪费，从而提高了应用程序的性能和稳定性。



### confVal

在go/src/net/conf.go文件中，confVal变量是一个指向值为 nil 的接口类型变量。它的作用是存储网络配置。

网络配置（network configuration）包括如何连接到网络、如何连接到主机、如何连接到端口等信息。在网络编程中，我们需要读取、修改、设置这些网络配置信息。

confVal变量是一个接口类型变量，它和具体类型的值相分离。这个变量可以指向任何类型的值，只要满足这个值实现了net.Flag中定义的方法即可。这些方法是get和set，它们用于获取和设置网络配置。

confVal变量是在net包的init()函数中定义的，它首先会注册一些标志（flags），然后初始化这些标志所对应的网络配置信息。例如，在网络编程中，我们经常需要设置网络超时时限（timeout），那么就可以使用如下代码：

`var dialTimeout = flag.Duration("dial_timeout", 30*time.Second, "duration of dial timeout")`

这里，我们创建了一个标志"dial_timeout"，它的默认值为30秒。这个标志的类型是一个Duration类型的指针。当程序运行时，我们可以使用命令行参数来修改这个标志的值，例如：

`./mypgm -dial_timeout=10s`

这个命令行参数将会把dial_timeout标志的值修改为10秒。

在net包中，confVal变量就是用来存储这些标志的值的。例如，我们可以使用如下代码来获取"dial_timeout"标志的当前值：

`timeout := confVal.Duration("dial_timeout")`

这个方法会返回"dial_timeout"标志的当前值，如果没有设置过，就返回默认值30秒。

总之，confVal变量是用来存储网络配置的，这些网络配置由一些标志组成。标志的值可以通过命令行参数来修改。confVal变量的值可以由标志的默认值或者命令行参数来确定。



### netdns

netdns变量是一个用于DNS配置的全局变量，它存储了Go语言中网络DNS的相关配置信息。具体来说，netdns变量是一个dnsConfig类型的指针，该类型定义在conf.go文件中。

dnsConfig类型包含以下字段：

- goos：操作系统类型。
- goarch：基于的CPU架构。
- cgo：表示是否使用cgo库来做DNS查询。
- canonicalName：表示是否需要返回完整的域名（而不是别名）。
- dnssecOK：用于启用DNSSEC支持。
- hostLookupOrder：与操作系统相关的主机名查找顺序。
- preferGo：表示是否首选Go解析器来做DNS查询。
- forcedNS：一个强制使用的DNS服务器地址。
- forcedAddrFamily：强制使用IPv4或IPv6地址。

在Go语言中，当应用程序需要进行域名解析时，会默认使用系统提供的DNS解析器进行解析。然而，在某些时候，我们可能需要自定义DNS解析器的行为，这时候就需要对netdns变量进行适当的配置。例如，我们可以使用dnssecOK字段来启用DNSSEC支持，以提高域名解析的安全性。






---

### Structs:

### conf

conf结构体在net包中主要用于存储IP协议中的配置信息，包括网络层（IP）、传输层（TCP、UDP）和控制层（ICMP）的各种相关参数，例如MTU（最大传输单元）、TTL（生存时间）、窗口大小、拥塞控制等。

conf结构体定义如下：

type conf struct {
    goos           string
    arch           string
    family         string
    andress        func(IPAddr) bool
    sockaddr       func(IPAddr) sockaddr
    scopeMap       map[int]string
    system         string
    soShipit       bool
    ip_dontfragment bool
    udpChecksum   bool
}

其中，各个字段的作用如下：

- goos：目标系统操作系统名称
- arch：目标系统CPU架构
- family：网络协议族（IPv4或IPv6）
- andress：将IP地址转换为bool值，用于校验地址的有效性
- sockaddr：将IP地址转换为网络地址结构体（如sockaddr_in或sockaddr_in6）
- scopeMap：IP地址作用域映射表
- system：目标操作系统名称
- soShipit：在传输数据包时是否发送了SHIPT包标志
- ip_dontfragment：IP包是否禁止分片标志
- udpChecksum：在传输UDP数据时是否计算UDP校验和

conf结构体中的这些参数是通过调用net包中的函数进行设置和修改的。如果没有明确指定，则默认使用系统默认值。conf结构体的主要作用是在创建和管理网络连接时提供默认参数，并且可通过修改各个字段来自定义网络连接参数，确保网络连接符合特定的应用场景需求。



### mdnsTest

在 go/src/net/ 中的 conf.go 文件中存在 mdnsTest 结构体，其作用是进行 mDNS 测试的配置。

mDNS（Multicast DNS）是一种基于广播的域名解析协议，用于在小型局域网中发现和查找设备和服务，例如在同一家庭网络中连接多个设备，例如打印机、音频设备、智能电视，等等。当设备加入网络时，它将会发出一个询问广播，其他设备若实现了 mDNS 协议，则能够响应该询问，并告知设备自己的 IP 地址和提供服务的信息。

因此，在 conf.go 文件中，mdnsTest 结构体的字段用于配置 mDNS 的测试参数，包括测试的域名、服务类型、查找方式等。同时，该结构体还包含了一些测试结果的字段，例如找到的服务的 IP 地址和端口号等。

在 net 包中的其他模块中，可以使用该结构体的实例进行 mDNS 测试，以确保在网络中正确发现和连接服务。



## Functions:

### systemConf

在go\src\net\conf.go文件中，systemConf函数被用来获取系统的网络配置参数。具体来说，该函数检索操作系统中的网络配置文件，然后返回一个结构体，该结构体包含了包括DNS服务器地址、本机IP地址和默认网关等重要的网络配置参数。

因此，该系统函数在创建网络连接时非常有用。通过读取系统配置，程序可以正确地配置网络连接，以确保它们遵循操作系统的网络设置，并能够与其他设备正常通信。

总之，systemConf是一个非常有用的函数，它可以为程序提供关键的系统级网络配置参数，并确保网络连接的正确配置。



### initConfVal

initConfVal函数的作用是初始化一些全局变量，这些变量对于网络连接的性能和可靠性十分关键。

具体来说，这个函数会根据不同的操作系统和网络环境来初始化一些变量，例如：

- 如果是Windows操作系统，会初始化netconf中的windowsVistaTCPConfig变量，该变量会影响TCP连接的行为，例如延迟确认和窗口缩小等；
- 如果是Linux操作系统，会根据是否启用了TCP Fast Open功能来初始化netconf中的enableTCPFastOpen变量，该变量决定了是否允许开启TCP Fast Open，从而加快TCP连接的建立速度；
- 如果是Android操作系统，会根据不同的版本号来初始化netconf中的androidControlsIPv6Egress和androidNoIPv6Interface变量，这些变量影响Android系统是否支持IPv6连接和IPv6路由控制等功能。

总的来说，initConfVal函数的作用是根据不同的操作系统和网络环境来初始化一些关键性能参数，从而提高网络连接的可靠性和速度。



### goosPrefersCgo

go/src/net/conf.go文件中的goosPrefersCgo函数用于判断当前操作系统是否更倾向于使用CGO库来实现网络功能。

CGO是Go语言的外部C语言调用库，可以让Go程序调用C语言库的函数。在实现网络功能时，例如TCP、UDP等协议的实现中，使用C语言库可以更好地利用操作系统提供的底层网络特性，从而提高Go语言程序的性能和可靠性。

goosPrefersCgo函数首先会判断当前操作系统是否已经显式地禁用了CGO库的使用，如果禁用了，则直接返回false，表示当前操作系统不适合使用CGO库。否则，再根据当前操作系统的类型，选择性地启用CGO库，返回一个bool值。

具体来说，Windows操作系统和Darwin（macOS）操作系统都默认启用CGO库，因为它们的底层网络实现采用了C语言库。而Linux和BSD等操作系统默认不启用CGO库，因为它们的底层网络实现已经被完全用Go语言实现了。当然，用户可以通过环境变量来手动启用或禁用CGO库的使用。

总之，goosPrefersCgo函数的作用是为了提高Go语言程序在各种操作系统下实现网络功能的性能和可靠性，以获得更好的用户体验。



### mustUseGoResolver

mustUseGoResolver函数是一个内部函数，用于判断是否应该使用Go的DNS解析器。

在默认情况下，Go使用系统的DNS解析器来处理所有网络请求。但有些情况下，系统的DNS解析器可能会出现问题。为了解决这些问题，Go提供了一种可以使用自己的DNS解析器的方式。

mustUseGoResolver函数的作用就是判断当前是否应该使用Go的DNS解析器。如果网络环境中配置了自定义的DNS服务器，或者是禁用了本地网络，那么就应该使用Go的DNS解析器。如果没有这些配置，则继续使用系统的DNS解析器。

具体来说，该函数会检查系统环境变量，包括环境变量“RES_OPTIONS”和“LOCALDOMAIN”，以及是否禁用了本地网络。如果发现任何一种情况，就返回true表示应该使用Go的DNS解析器。如果没有这些配置，则返回false表示继续使用系统的DNS解析器。

这个函数的结论会影响到后续的DNS解析操作。如果返回true，那么后续的DNS解析就会使用Go的DNS解析器。如果返回false，那么就会继续使用系统的DNS解析器。



### hostLookupOrder

函数`hostLookupOrder()`是用来决定主机名解析时搜索主机名的顺序。在Go语言中，当程序需要解析主机名时，常常会调用标准库的`net`包中的`LookupHost()`函数。这个函数会根据操作系统的配置和网络环境，在多个地方查找主机名的IP地址。如果在hostLookupOrder()函数中可以得到一个有序的解析顺序，就可以提高程序解析主机名的效率。

`hostLookupOrder()`函数的实现过程如下：

首先，`hostLookupOrder()`函数根据操作系统类型(`<windows>`或`<linux>`等)和网络环境(`ipv6Enabled()`，是否开启IPV6协议)来判断主机名解析的顺序。如果ipv6Enabled为true，则优先通过IPV6进行主机名解析。

其次，根据主机名的结构来决定搜索主机名的顺序。Go语言规定主机名的结构必须符合RFC952和RFC1123规范。RFC952规定主机名只能由字母、数字和连字符组成，长度不能超过24个字符。RFC1123则对主机名做了进一步的规定，包括了长度、字符集、连字符位置等等要求。因此，`hostLookupOrder()`函数根据主机名的结构来决定搜索的顺序，以保证程序尽可能快速地找到正确的IP地址。

最后，如果在操作系统配置中指定了DNS服务器，则会优先使用该DNS服务器进行主机名解析。这样可以减少网络环境对主机名解析的影响。

综上所述，`hostLookupOrder()`函数的作用是为了提高程序解析主机名的效率，通过操作系统和网络环境等多方面的判断，决定搜索主机名的顺序。



### goDebugNetDNS

goDebugNetDNS是一个名为debugNetDNS的变量的getter函数，在net包启动时被调用。它用于返回一个bool值，指示该包是否启用了网络DNS解析的调试信息。

当启用时，该变量将与其他调试标志组合，在DNS解析过程中打印更多的调试信息。如果没有启用，它将返回false，并且不会在解析模块中打印调试信息。

这个函数主要是用来帮助调试和分析网络DNS解析的行为，有助于定位可能存在的问题。它通常由网络系统管理员、开发人员或其他网络相关人员使用。



### isLocalhost

isLocalhost是一个用于判断IP地址是否为本地地址的函数。在网络通信中，为了方便和安全起见，通常会限制只有本地IP地址才能进行一些敏感操作，如回环地址127.0.0.1或::1等。

isLocalhost函数首先检查传入的IP地址是否为IPv4的回环地址127.0.0.1，如果是，则直接返回true，表示是本地IP地址。如果不是，则检查传入的IP地址是否为IPv6的回环地址::1，如果是，则也返回true，表示是本地IP地址。如果都不是，则继续判断是否为本机网卡地址，最后查询操作系统的hosts文件，判断IP地址是否为本地的localhost域名的解析结果，如果都不是本地IP地址，则返回false，表示不是本地IP地址。

总的来说，isLocalhost函数的作用就是判断一个IP地址是否为本机的本地地址，以便在网络通信中做出相应的处理。



### isGateway

func isGateway(mask, ip net.IP) bool

这个函数的作用是检查给定的IP地址和子网掩码是否匹配，并检查IP地址是不是该子网掩码所表示的子网中的默认网关。

如果IP地址和子网掩码匹配（即IP地址和子网掩码按位相与等于网络地址），则检查IP地址是否是默认网关。这样可以确保目标主机在同一子网中，并且有一个默认网关。

如果IP地址不是默认网关，则可能是目标主机或无法直接到达的路由器。



### isOutbound

isOutbound函数的作用是判断一个TCP连接是否为出站连接。在网络编程中，一个TCP连接可以分为两种类型：出站连接和入站连接。出站连接是主动从本机向远程主机发起的连接，而入站连接则是远程主机主动发起的连接。

isOutbound函数通过判断一个TCP连接的本地IP地址是否在本地网络接口的IP地址范围内，来确定该连接是否为出站连接。如果本地IP地址在本地网络接口的IP地址范围内，则该连接为出站连接，否则为入站连接。

具体来说，isOutbound函数首先获取本地网络接口的IP地址列表。然后遍历该列表，判断该连接的本地IP地址是否在列表中。如果在，就返回true，表示该连接为出站连接；如果不在，则返回false，表示该连接为入站连接。

实际上，isOutbound函数在net包中被多次调用，用于判断TCP连接类型。该函数对于网络编程来说非常重要，因为它能够帮助我们正确地处理TCP连接，从而优化网络通信的质量和效率。



