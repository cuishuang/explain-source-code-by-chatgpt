# File: lookup_plan9.go

lookup_plan9.go是Go语言标准库中net子包的一部分，它提供了在Plan 9操作系统上进行网络地址解析的功能。

具体来说，这个文件包含了一个函数lookupIPPlan9，用于将一个主机名解析为IP地址。这个函数在Plan 9操作系统上使用名字空间进行解析，能够处理Plan 9操作系统上的特殊名称格式以及网络服务的别名。

除了lookupIPPlan9函数之外，lookup_plan9.go文件还包含了一些内部使用的类型和函数，用于支持地址解析。这些类型和函数不是公开的API，一般情况下不需要用户直接使用。

总的来说，lookup_plan9.go文件是Go语言标准库中提供Plan 9操作系统上网络地址解析功能的一部分。它为在Plan 9操作系统上运行的Go程序提供了方便的网络通信能力。

## Functions:

### query

在Go语言标准库中，`go/src/net/lookup_plan9.go`文件中的`query`函数用于执行DNS查询。该函数所在的文件主要用于实现在Plan 9操作系统上执行的DNS查询。

函数签名如下：

```go
func query(name string, qtype uint16) (rrs []dnsRR, err error)
```

其中，`name`参数表示要查询的域名，`qtype`参数表示查询的记录类型，例如A记录、MX记录等。

函数首先会检查是否在`/srv/dns`中配置了可用的DNS服务器地址，并获取对应的地址列表。如果获取不到地址列表，则直接返回错误。如果有可用的地址，函数会依次向这些地址发送DNS请求，直到其中某个响应成功返回结果。

在发送DNS请求时，函数会先创建一个UDP连接，然后发送DNS请求报文，在等待一定时间后，如果没有收到响应，则会尝试几次重试，直到超过设定的超时时间。在收到响应后，函数会将响应报文解析为DNS记录（RR），然后返回结果。

需要注意的是，由于该函数与Plan 9操作系统上的DNS查询有关，因此在其他操作系统上运行代码时可能会出现不兼容的情况。建议使用其他类似的函数，例如`net.LookupIP`等，来执行DNS查询。



### queryCS

func queryCS(name string, deadline time.Time) ([]net.IPAddr, error) 是go/src/net中lookup_plan9.go这个文件中的一个函数。

queryCS函数是在Plan 9环境下进行DNS查找的函数。它通过打开/dev/dns文件来查找主机名的DNS记录。它使用了Plan 9的DNS查询实现，并对结果进行解析，以返回一个包含IP地址的切片。

该函数的参数name是需要查找的主机名，deadline是DNS查询的截止时间。如果deadline在当前时间之前，该函数将会立即返回一个超时错误。

queryCS返回一个包含全部IP地址的切片。如果没有找到IP地址，则返回一个“not found”错误。如果无法打开/dev/dns文件，则返回一个“network unreachable”错误。

该函数在标准情况下是不会由普通的网络代码使用的，因为它专门针对Plan 9环境进行了优化。它只在特定的嵌入式系统或网络环境中使用。



### queryCS1

本回答为机器翻译，可能存在语言表述不准确或不清晰之处，仅供参考。

在go/src/net/lookup_plan9.go文件中，queryCS1是一个函数，用于查询计算机系统1（CS1）的IPv4地址和端口。当程序需要连接到远程服务器时，它需要知道服务器的IP地址和端口号。这些信息通常由域名解析器获得，但在Plan 9操作系统上，网络协议使用了一种不同的机制。

在Plan 9操作系统中，所有网络访问都通过名为CS1的文件服务器进行。因此，查询CS1函数检查该服务器的命名空间以查找主机名，并返回相应的IP地址和端口号。如果主机名无效或操作系统版本不允许使用CS1，则会返回错误。

此函数使用的主机名可以是一个DNS名称（例如“example.com”）或一个Plan 9网络名称（例如“srv”）。如果给出的主机名包含一个“!”字符，则会将其视为一个目录名称，而不是计算机系统1的主机名。函数还尝试通过查询服务名称文件（/etc/quickdns）来解析DNS名称，并返回匹配的IP地址和端口号（如果存在）。

总之，queryCS1函数是用于查询计算机系统1的IP地址和端口号的函数，在Plan 9操作系统上使用，可以用于代替传统的DNS解析器。



### queryDNS

queryDNS 函数是用来查询 DNS 记录的，它接受两个参数：hostname 和 qtype。其中，hostname 表示要查询的主机名，qtype 表示要查询的 DNS 记录类型。当 queryDNS 函数被调用时，它会对 hostname 和 qtype 进行一些初始化处理，并且创建一个 DNS 查询请求包。请求包中包含了一些标准的 DNS 查询信息，例如查询类型、查询类和查询域。

接下来，queryDNS 函数调用 net.Dial 函数与本地 DNS 服务器进行通信，将请求包发送给 DNS 服务器。DNS 服务器接收到请求包后，会根据其中的查询信息进行相应的查询操作，并且将查询结果封装到响应包中返回给客户端。queryDNS 函数会从响应包中解析出目标主机的 IP 地址，并将其返回给调用者。如果无法查询到目标主机的 IP 地址，则会返回一个错误对象。

总体来说，queryDNS 函数的作用是查询 DNS 记录，获取指定主机的 IP 地址。这是网络编程中比较常见的一种操作，例如在使用 Go 语言编写网络应用程序时，我们需要通过 DNS 解析获取目标主机的 IP 地址，然后才能进行网络通信。



### toLower

toLower是一个函数，它的作用是将字符串中所有大写字母转换为小写字母。

在go/src/net/lookup_plan9.go文件中，toLower函数被用于规范化主机名和服务名。规范化是为了方便后续处理，避免因大小写不同导致的匹配错误。因此，toLower函数被用于将主机名和服务名中所有大写字母转换为小写字母，从而使它们具有相同的大小写格式，便于比较和匹配。

以下是toLower函数的定义：

func toLower(s string) string {
    b := make([]byte, len(s))
    for i := 0; i < len(s); i++ {
        c := s[i]
        if c >= 'A' && c <= 'Z' {
            c = c - 'A' + 'a'
        }
        b[i] = c
    }
    return string(b)
}

函数的实现使用了一个循环来遍历字符串中的每个字符。如果某个字符是大写字母（'A'到'Z'之间的字符），则将其转换为相应的小写字母（'a'到'z'之间的字符）。最后，返回一个包含所有小写字母的新字符串。

这样就可以确保主机名和服务名的大小写格式相同，从而可以成功地进行比较和匹配。



### lookupProtocol

lookupProtocol函数的作用是返回TCP或UDP协议的编码。它是net包中实现网络的底层功能的一部分。

lookupProtocol函数的输入参数是一个字符串协议名，例如"tcp"或"udp"，并返回一个表示该协议的常量。如果输入的协议名无效，函数会返回0。

lookupProtocol函数实现了一种查找协议编码的算法，该算法首先寻找给定的协议名，如果找到则返回该协议的编码。如果没有找到，函数会检查输入字符串是否是一个数字，如果是，它会将其解释为协议的编码，并返回该编码。如果输入字符串既不是有效的协议名也不是数字，则返回0。

lookupProtocol函数是net包中许多其他函数和类型使用的重要组成部分，例如Dial、Listen和ResolveTCPAddr函数。



### lookupHost

lookupHost函数是Go语言中net包中的一个函数，用于实现域名解析功能。在lookup_plan9.go文件中，这个函数被用于在Plan 9操作系统中进行DNS解析。

具体来说，lookupHost函数接受一个主机名作为输入参数，然后从本地的名字服务（通常是DNS服务器）获取该主机名对应的IP地址列表。当本地名字服务无法解析该主机名时，该函数会尝试使用Plan 9操作系统提供的文件系统接口来解析主机名。

如果主机名解析成功，则返回一个字符串类型的IP地址列表，否则返回一个非nil的错误。如果输入参数为空，则返回本机的所有IP地址。

该函数的内部实现过程涉及到多个底层系统调用和文件操作，包括gethostbyname、ipconfig、ndb等，这些操作都是对于底层网络协议栈的封装和处理。

总之，lookupHost是一个非常重要的网络函数，它可以帮助Go程序在Plan 9系统中高效地进行DNS解析，从而实现各种网络服务。



### preferGoOverPlan9

preferGoOverPlan9函数的作用是在主机名解析中确定使用Go语言还是Plan 9的DNS服务。

在Posix系统中，通过调用getaddrinfo来解析主机名。而在Plan 9环境中则使用了特殊的DNS服务来解析主机名。当Go程序在运行在Plan 9系统中时，它会自动识别并使用Plan 9的DNS服务进行主机名解析。

但是，在某些情况下，我们需要优先选择使用Go语言的方式来解析主机名，例如：

- Plan 9 DNS服务器并不总是可靠的，可能会出现故障或者不可用的情况。
- 使用Plan 9的DNS服务会导致程序运行缓慢。

因此，当我们需要避免Plan 9的DNS服务时，可以使用preferGoOverPlan9函数来进行设置并选择使用Go语言的方式进行主机名解析。

这个函数基本上是一种策略，它会从以下列表中选择一个最佳的策略：

- 正在运行的操作系统是否为Plan 9
- 是否设置了开关preferGoOverPlan9
- 是否存在Plan 9的DNS服务

如果存在Plan 9的DNS服务，preferGoOverPlan9函数会返回false，并且使用Plan 9的DNS服务进行主机名解析操作。如果不存在Plan 9的DNS服务或者preferGoOverPlan9被设置为true，则返回true并使用Go语言的方式进行主机名解析操作。



### preferGoOverPlan9WithOrderAndConf

func preferGoOverPlan9WithOrderAndConf(addrList []Addr, order []string, conf func(string) bool) []Addr 函数的作用是对地址列表进行排序，以便在选择网络地址时更倾向于使用 Go 语言内置的实现而非 Plan 9 的实现。

通过检查每个地址的网络类型，并在地址列表中按以下顺序安排它们，该函数选择适当的实现：

1. IPv4 地址，仅当有 Go 实现时才会选择。
2. IPv6 地址，仅当有 Go 实现时才会选择。
3. Unix 套接字地址（如果存在）。

在其中，Go 实现具有更高的优先级，这意味着即使 Plan 9 实现是可用的，Go 实现也将首选。这是因为 Go 实现通常更稳定和更快速。

order 参数指定了一组可接受的协议，它们将在 Plan 9 实现逐个尝试，直到找到一个支持的协议。例如，如果该列表包含 "tcp" 和 "udp"，则 Plan 9 实现会首先尝试 tcp，并且仅在失败时才尝试 udp。

conf 参数是一个函数，它将接收由网络类型构成的字符串作为参数，并返回一个布尔值，表示网络是否被禁用或启用。网络类型字符串对应于 net 包中支持的网络类型。

最后，该函数将返回一个经过排序的地址列表。



### lookupIP

该文件中的lookupIP函数是用于查询指定主机名的IP地址列表的函数。具体介绍如下：

函数签名：

```go
func lookupIP(host string) ([]IP, error)
```

参数host是要查询IP地址的主机名。

返回值是一个由IP类型的元素组成的切片和一个error类型的错误信息。如果查询成功返回IP地址列表，反之返回错误信息。

函数实现：

该函数的实现分为两个阶段，第一阶段是调用Plan 9中的iplookup命令查询主机名对应的IP地址列表，第二阶段是对查询结果进行解析并返回IP地址列表。

第一阶段调用方法：

该函数首先尝试根据host参数构建Plan 9的iplookup命令行参数，并使用exec.Command函数来调用iplookup，如果执行成功返回查询结果，如果执行失败返回错误信息。

```go
func lookupIP(host string) ([]IP, error) {
    // 构建命令行参数
    cmd := exec.Command("/bin/ip/iplookup", host)
    out, err := cmd.Output()
    if err != nil {
        return nil, &DNSError{Err: err.Error(), Name: host}
    }
    return parseIP(out, host), nil
}
```

第二阶段解析方法：

第二阶段主要是对查询结果进行解析并返回IP地址列表。

```go
func parseIP(out []byte, host string) []IP {
    var ips []IP
    for _, line := range strings.Split(string(out), "\n") {
        if len(line) > 0 && line[0] != ' ' {
            // 从查询结果中解析出IP地址，并加入到切片ips中
            ip := parseIPv4(line)
            if ip != nil {
                ips = append(ips, ip)
            }
        }
    }
    return ips
}

func parseIPv4(s string) IP {
    // 解析IP地址
    if len(s) == 0 {
        return nil
    }
    ip := net.ParseIP(s)
    if ip == nil {
        return nil
    }
    return IPv4(ip[12], ip[13], ip[14], ip[15])
}
```

解析方法主要是依次读取查询结果的每一行，找到以IP地址开始的行，然后解析出IP地址，并加入到IP地址列表中。具体的IP地址解析由另一个函数parseIPv4完成，它首先调用net.ParseIP函数解析出IP地址的字节表示形式，然后再将其转换成IPv4类型的IP地址。最后返回IP地址列表。

总结：

lookupIP函数是实现了在Plan 9环境中查询指定主机名的IP地址列表的功能。它的实现主要是通过调用Plan 9中的iplookup命令来获取查询结果，并对查询结果进行解析，返回IP地址列表。



### lookupPort

在Go语言中，lookup_plan9.go文件中的lookupPort函数是用于查找网络端口的函数。它接收一个字符串参数，即要查找的端口号或服务名，返回一个整数型的端口号。

具体来说，该函数首先尝试解析参数字符串，判断其是否为数字端口。如果是数字端口，则直接将其转换为整数返回。如果不是数字端口，则会调用Plan 9的getaddrinfo函数进行DNS查找，获取服务名称或别名对应的端口号。

此外，lookupPort函数还会处理一些特殊的端口名，如“http”、“ftp”等常见的网络服务，将其转换为对应的端口号。如果无法查找到端口号，则返回0表示失败。

总之，lookupPort函数的主要作用是在Plan 9操作系统下查找网络服务名称或别名对应的端口号，以便进行网络连接。



### lookupCNAME

lookupCNAME函数在net包中的lookup_plan9.go文件中定义，用于查询指定主机名对应的CNAME记录，即主机名的规范名称。具体作用如下：

1. 查询主机名的CNAME记录。假设有一个主机名foo.example.com，lookupCNAME函数会首先查询foo.example.com对应的A记录。如果查询成功，就返回A记录中包含的IP地址。如果查询失败，就继续查询对应的CNAME记录。

2. 如果主机名存在CNAME记录，lookupCNAME函数将返回CNAME记录对应的规范名称。例如，假设主机名foo.example.com有一个CNAME记录指向bar.example.com，那么lookupCNAME函数将返回bar.example.com。

3. 如果主机名不存在CNAME记录，lookupCNAME函数将返回一个空字符串表示查询失败。

总之，lookupCNAME函数提供了一种查询主机名的规范名称的方式，可以帮助使用网络编程的开发者快速获得网络主机的地址。



### lookupSRV

lookupSRV函数是Go语言标准库net包中实现的一个DNS SRV记录查询函数，它用于根据服务类型、服务协议和服务名称查询DNS中的SRV记录信息。

在通信协议中，SRV记录是指指向服务的服务器名和端口号的一种DNS记录类型。通过SRV记录，可以实现服务的负载均衡、故障转移和服务发现等功能。

lookupSRV函数接受三个参数：服务类型（service）、服务协议（proto）和服务名称（name），其中服务类型和协议是必需的，名称可以为空字符串。它返回两个值：服务器名字的切片和端口号。

它会先向DNS服务器发起查询请求，获取到所有符合条件的SRV记录，然后按照一定的规则进行排序和筛选，最终返回一个可用的服务器地址和端口号。

该函数的实现方法和其他查询函数类似，使用了DNS解析器来查询DNS服务器，并返回解析结果。通过SRV记录查询实现了一种面向服务的服务发现过程，可以帮助程序自动根据服务类型、协议和名称查找服务器，从而实现更高效的网络通信。



### lookupMX

lookupMX是在net包中的lookup_plan9.go文件中定义的一个函数，用于在Plan 9上通过DNS协议查询给定域名的MX记录。MX（Mail Exchange）记录是指定邮件服务器的DNS记录类型。

该函数的输入参数是一个字符串类型的域名，输出是一个struct类型的数组，表示查询到的MX记录。如果查询失败，输出一个错误。

该函数首先通过调用syscall库中的_getmxrr函数查询给定域名的MX记录，然后将此记录格式化为一个struct类型的数组，具体格式为：

```go
type MX struct {
    Host string
    Pref uint16
}
```

其中Host表示MX记录中指定的邮件服务器域名，Pref表示MX记录中的优先级。

如果查询失败，则返回一个非nil的error值。

lookupMX函数的作用是充当net包中的DNS解析器，通过DNS协议查询给定域名的MX记录，并将结果解析为程序可以使用的格式。这个函数的实现对于开发邮件相关的应用程序非常有用，因为开发者可以通过它查询指定域名的邮件服务器，然后向该邮件服务器发送邮件。



### lookupNS

lookupNS函数是go/src/net/lookup_plan9.go文件中的一个函数，它的作用是向域名服务器查询指定主机名的DNS服务器记录。

具体来说，lookupNS函数会根据传入的host参数，从本机默认的域名服务器中查找与主机名匹配的DNS服务器记录。如果找到了记录，则返回记录中包含的DNS服务器的地址列表。

这个函数的实现方式是通过Unix系统调用“dnsquery”来查询本机的域名服务器记录。在调用dnsquery时，lookupNS会传入一个查询主机名的参数，如"_dns-sd._udp.local"，并传入一个包含查询结果的缓冲区。当dnsquery返回后，lookupNS会将缓冲区中的结果解析并返回DNS服务器地址列表。

lookupNS函数在网络编程中非常重要，它是实现DNS查询和域名解析的基础。除了在lookup_plan9.go文件中之外，该函数的实现还可以在其他类Unix/Linux系统和Windows系统上找到。



### lookupTXT

lookupTXT函数是 Go 语言中 net 包中的一个函数，它的作用是在 Plan 9 操作系统下查询指定域名的 TXT 记录，返回一个字符串数组。

该函数会在 Plan 9 的 /net/dns 文件中查找指定的域名对应的 TXT 记录，并将结果解析成字符串数组。在解析时，该函数会将每个记录中的字节串转换成 UTF-8 字符串。

该函数的定义如下：

```go
func lookupTXT(name string) ([]string, error)
```

其中参数 name 表示要查询的域名，返回值是字符串数组和一个错误信息。

需要注意的是，该函数只适用于 Plan 9 操作系统，其他操作系统下运行该函数会返回一个错误。如果要在其他操作系统下查询指定域名的 TXT 记录，可以使用 net 包中的其他函数，如lookupTXT函数。



### lookupAddr

lookupAddr函数是用于在Plan 9网络上查找特定IP地址的主机名的函数。它将一个IP地址作为参数，并使用系统命令执行查找操作。在Plan 9系统上，主机名存储在名字空间中，并且可以通过/ns命令访问。

lookupAddr函数首先检查给定的IP地址是否为空地址或私有地址。如果是，则函数会返回错误。如果不是，它会使用系统命令nslookup获取主机名。

如果主机名找到，则函数会将主机名作为字符串返回，并且错误为nil。否则，函数将返回一个空字符串和一个描述错误的非nil错误。

在实际应用中，lookupAddr函数可以帮助应用程序根据IP地址获取主机名，这对于记录日志和网络故障排除有很大帮助。



### concurrentThreadsLimit

在 go/src/net/lookup_plan9.go 中的 concurrentThreadsLimit 函数用于计算并发限制的数量，以控制最大并发线程数。

该函数的作用是根据平台的不同限制并发线程数，因为不同操作系统和硬件有不同的线程数限制。

在 Plan 9 中，可以使用 syscall.Nclock() 来检索当前计算机上可用的 CPU 数量，并使用该数字计算允许的最大并发线程数。

此函数确保在进行 DNS 查询时，系统不会因同时进行大量 DNS 查询而过载。因此，可以保持系统的稳定性和性能。

在执行 DNS 查询时，该函数还利用 Plan 9 的文件系统 API 获取查询主机的 IP 地址，并将 DNS 查询发送给 DNS 服务器以获取响应。并发限制的确立确保 DNS 查询不会超载系统。



