# File: lookup_unix.go

lookup_unix.go是Go语言标准库中net包的一个文件，主要包含Unix（包括Linux、macOS等）操作系统下域名解析相关的函数和结构体。具体作用包括以下几个方面：

1. 实现了net包中的lookupIP函数，该函数通过调用Unix系统调用getaddrinfo和getnameinfo实现了将主机名或者IP地址转换为IP地址列表和主机名（反向查询）的操作。

2. 定义了addrinfo和netent两个结构体。addrinfo结构体是Unix系统调用getaddrinfo所需要的参数，包括主机名（node）、服务名（service）、地址族（family）、连接类型（socktype）和协议类型（protocol）等信息；netent结构体则是getnetbyname和getnetbyaddr系统调用的返回结果，包括网络名（name）、网络号（net）、地址长度（netlen）以及网络协议（proto）信息。

3. 实现了Unix系统下对主机名解析的缓存机制，缓存了最近解析过的主机名和IP地址，避免重复解析带来的性能开销。

总之，lookup_unix.go文件实现了与Unix系统下域名解析相关的底层功能，为net包中的高级API（如Dial和Listen函数）提供了支持。




---

### Var:

### onceReadProtocols

在go/src/net中的lookup_unix.go文件中，onceReadProtocols变量是一个用于单次读取协议的锁。它的作用是确保只有一个goroutine可以同时读取协议，从而避免竞争条件和可能的数据竞争。这个锁是一个sync.Once类型的值，它在第一次调用时会执行指定的函数并将其标记为已执行，后续的调用则会直接返回之前的结果。

在这个文件中，onceReadProtocols变量用于初始化protocols数组。protocols存储了操作系统支持的所有网络协议的名称和ID。这些协议名称和ID可以用于调用socket()和bind()函数来创建和绑定网络套接字。但是，由于协议信息是在程序运行时动态获取的，因此需要一定的初始化过程来读取并保存这些协议信息。onceReadProtocols变量就是用来控制这个初始化过程的锁。

在lookupProtocols函数中，当protocols数组还没有被初始化时，会调用onceReadProtocols.Do(func() {...})函数来获取锁并执行指定的函数。这个函数会逐个读取系统支持的所有协议信息，并将它们保存为一个map类型的数据结构。读取完成后，这个函数会释放锁并将保存的协议信息赋值给protocols变量。在后续的调用中，如果protocols已经被初始化，则直接返回缓存的结果。

总之，onceReadProtocols变量用于保证protocols数组只被初始化一次，并且在初始化期间只有一个goroutine可以读取协议信息，以避免并发问题和数据竞争。这个锁的使用可以提高代码的可靠性和性能。



## Functions:

### readProtocols

readProtocols这个func的作用是读取/etc/protocols文件并返回一个map，其中包含了协议名称和对应协议号的关联。在Unix系统中，这个文件主要被用来给协议号赋予可读的名称。

具体来说，这个func会首先尝试读取整个/etc/protocols文件并将其分成行。然后它会遍历每一行并检查其格式。如果格式正确，就会将名称和协议号存储到一个map中，其中名称是map的键，协议号是map的值。最后，这个func会返回完整的map。

这个func的作用在于为网络通信提供更好的可读性。在基于网络的程序中，协议号通常需要与网络协议栈交互，例如在设置套接字选项时。但是，使用协议名称可以使代码更易于理解和维护。因此，这个func将协议号和对应的名称存储为一个map，以便程序可以在需要时将协议号转换为可读的协议名称。



### lookupProtocol

lookupProtocol函数是在Unix系统下进行DNS查询并返回协议类型的函数。它的作用是将给定的协议名（如“tcp”、“udp”、“icmp”等）转换为对应的协议编号（如IPPROTO_TCP、IPPROTO_UDP、IPPROTO_ICMP等），以便在进行Socket操作时使用。

具体来说，该函数会通过syscall包提供的getprotobyname系统调用函数从系统中获取给定协议名的协议编号。如果该协议名不存在或获取失败，则返回默认协议编号（通常为0）。在返回结果前，该函数还会将结果存入cacheProtocolMap中，以便在下一次查询相同协议名时可以直接从缓存中获取。

该函数对于网络通信中的协议选择十分关键，因为不同协议可能会有不同的性能表现和特定的用途，通过将协议名转换为对应的协议编号后，程序可以更方便地操作Socket并使用相应的协议。



### lookupHost

lookupHost函数是net包中用于进行DNS解析的函数之一，它的作用是将给定的主机名解析为该主机的IP地址列表。具体来说，lookupHost函数将会依次查询主机名对应的IPv4和IPv6地址记录，然后返回一个包含所有查询结果的切片。

lookupHost函数的实现主要依赖于操作系统的DNS解析服务。在Unix系统中，lookupHost函数会调用getaddrinfo函数进行DNS查询。getaddrinfo是一个标准的C函数，通常由系统的glibc库提供。它能够根据给定的主机名和协议类型返回一组与之对应的网络地址。lookupHost函数在调用getaddrinfo时，会根据主机名参数指定的不同形式（如IP地址、域名、别名等），以及其它参数（如网络类型、查询超时时间等），来选择适合的查询方式。如果查询成功，则会将返回所有查询到的地址信息。

在网络编程中，lookupHost函数通常用于在客户端中解析主机名，以便与服务器进行连接。它也可以帮助我们实现一些需要基于主机名的操作，比如日志收集、爬虫等。值得注意的是，在进行DNS查询时，由于网络不稳定性等原因，可能会导致查询失败或返回不完整的结果，因此我们需要做好错误处理和结果检查，以保证程序的健壮性。



### lookupIP

lookupIP是一个在Unix系统下实现的函数，用于将主机名解析为IP地址列表。它接受一个参数host string，其中包含要解析的主机名。lookupIP还返回一个切片，其中包含解析出的所有IP地址。

在Unix系统中，可以使用多种方式来进行主机名到IP地址的解析。lookupIP会根据当前系统的设置自动选择最佳的方式进行解析。通常情况下，这个函数会首先尝试使用/etc/hosts文件中的条目进行解析，然后会尝试查询DNS服务器或者NIS服务器等其他方式进行解析。

具体来说，lookupIP的实现会调用gethostbyname系统调用来进行解析。如果解析成功，会返回一个hostent结构体，其中包含了IP地址列表。lookupIP会将hostent结构体中的IP地址转换为Go语言中的net.IP类型，并返回一个切片，其中包含了所有解析出的IP地址。

总之，lookupIP这个函数提供了一种在Unix系统中简单方便地进行主机名到IP地址解析的方法。



### lookupPort

lookupPort函数的作用是根据传入的服务和协议找到对应的端口号。

在代码中，首先定义了一个ports数组，用来存储服务和协议对应的端口号。然后遍历数组，通过比较传入参数serviceName和protocolName与数组中的对应元素是否匹配，如果匹配则返回该端口号。如果没有找到匹配的端口号，则返回空字符串。

例如，调用lookupPort("http", "tcp")将会返回字符串"80"，代表HTTP协议的TCP端口是80。而调用lookupPort("ssh", "udp")则会返回空字符串，表示没有找到对应的端口号。

这个函数主要被其他函数调用，比如在lookupIP函数中，会根据传入的参数找到对应的端口号后，再去尝试连接该端口。



### lookupCNAME

lookupCNAME函数的作用是查询指定主机名的规范名。

在DNS系统中，一个主机名可能有多个别名，但只有一个规范名。lookupCNAME函数会查询给定主机名的规范名，并返回该规范名以及找到的任何别名。

该函数会首先查询主机名本身，如果没有找到，则会递归查询主机名的别名直到找到规范名为止。如果找到了规范名，则会将其缓存起来以供以后使用。

lookupCNAME函数的实现使用了Linux系统中的getnameinfo和getaddrinfo函数，它们分别用于查询主机名和IP地址的规范名称。

该函数还会处理一些错误情况，例如无法解析主机名或解析返回的结果不正确。如果有任何错误发生，该函数会返回一个错误值。

总之，lookupCNAME函数是net包中实现DNS查询的重要函数之一，它提供了查询主机名规范名的便捷方法，并且自动处理了别名的递归查询。



### lookupSRV

lookupSRV函数的作用是查找服务记录（Service Records），它是DNS中的一个记录类型，包含了提供某个服务的主机名和端口号。

在网络编程中，通常需要查找一个特定服务（如SMTP邮件服务器、FTP服务器、Web服务器等）所在的主机和端口号，以便建立连接。lookupSRV函数就是用来完成这个任务的。

具体来说，lookupSRV函数会向DNS服务器发送一个SRV记录查询请求，根据给定的服务名、协议、和域名来查询符合条件的服务主机和端口号。

例如，lookupSRV函数可以用来查询SMTP邮件服务器的主机名和端口号：

```
_, addrs, err := net.LookupSRV("smtp", "tcp", "example.com")
if err != nil {
    fmt.Println(err)
    return
}
if len(addrs) == 0 {
    fmt.Println("no srv records found")
    return
}
for _, a := range addrs {
    fmt.Printf("Host: %s, Port: %d\n", a.Target, a.Port)
}
```

上面的代码表示查询example.com域名中提供SMTP服务的主机和端口号，查找到的结果存储在addrs中，遍历addrs可以获取每个服务主机的主机名和端口号。



### lookupMX

lookupMX函数是net包中的函数之一，它的作用是在域名系统中查找MX记录。MX记录用于将电子邮件发送到特定的邮件服务器。lookupMX函数接收一个domain参数，该参数是待查询MX记录的域名。

在执行lookupMX函数时，它会首先使用Go语言内置的net包中的函数进行DNS查询操作。并且在查询操作结束之后，lookupMX会返回一个包含多个MX记录的slice（切片），每个MX记录包括一个优先级（priority）和一个指向邮件服务器的域名。

如果查询期间出现了错误，例如DNS服务器无响应或查询超时，lookupMX将返回一个非nil的error对象，以指示查询操作失败。

在实际生产中，该函数可用于为邮件服务器检索MX记录，并确定应将电子邮件路由到哪个邮件服务器。



### lookupNS

在go/src/net中，lookup_unix.go文件中的lookupNS函数用于执行DNS查询以获取给定域名的权威域名服务器记录。它接受两个参数：域名和要使用的网络类型（默认为“ip”）。

该函数首先将给定的域名转换为问询格式，并将其编码为DNS消息。然后它使用系统调用发送该消息到本地DNS解析器，并等待响应。如果解析器成功返回响应，则lookupNS函数将从响应中提取权威域名服务器记录，并以字符串数组的形式返回它们的IP地址列表。

lookupNS函数是go/net包非常重要的一部分，因为它为TCP/IP协议栈提供了与DNS解析器的接口，它可以查询DNS记录以支持在互联网上进行网络通信。它可以帮助网络程序员编写与域名解析相关的代码，从而更轻松地处理网络连接和交互。



### lookupTXT

lookupTXT函数是net包中的一个函数，它对指定的域名执行TXT查询，并返回一个指定域名的TXT记录。在DNS中，TXT记录是在域名系统中存储与域名相关联的文本信息。通常，TXT记录用于验证域的身份，例如SPF（Sender Policy Framework）记录。

lookupTXT函数的定义如下：

```
func lookupTXT(host string) ([]string, error)
```

该函数接受一个字符串参数host，该参数表示要查询的域名。如果查询成功，则该函数会返回一个字符串类型的数组，其中包含指定域名的TXT记录。如果查询失败，则该函数会返回一个错误。

lookupTXT函数实际上是对net包中的dnsLookup函数的一层封装。dnsLookup函数负责从DNS解析器中查找指定的域名，并返回该域名的所有解析记录。但是，它不区分解析记录的类型，因此需要在结果中筛选出TXT记录，并将其返回。

总的来说，lookupTXT函数的作用是在DNS中查询指定域名的TXT记录，并返回结果。它可以帮助开发人员验证域的身份、获取域名相关的信息等。



### lookupAddr

lookupAddr是一个函数，它用于反向DNS查找给定IP地址的域名列表。

具体来说，它将给定的IP地址作为参数，然后对其进行反向DNS查询，以获取对应的域名列表。如果查询成功，则返回一个字符串切片，其中包含了所有找到的域名。

在实现上，lookupAddr函数先将给定的IP地址转换为网络字节序（Little-Endian），然后使用系统调用getnameinfo()进行反向DNS查询。查询结果会被保存在addrinfo结构体中，lookupAddr函数会将该结构体中所有的主机名（hostnames）提取出来，最终将它们存储到一个字符串切片中并返回。

举个例子，如果调用lookupAddr("8.8.8.8")，它将返回一个包含"google-public-dns-a.google.com"的字符串切片，因为这是Google公共DNS服务器在反向DNS查询中返回的域名。



### concurrentThreadsLimit

在go/src/net中的lookup_unix.go文件中，concurrentThreadsLimit函数是在Unix系统上启用并发域名解析的功能。它的主要作用是确定并发域名解析的线程数上限。

在Unix系统上，使用getaddrinfo（）函数批量获取多个IP地址的过程是非常慢的，因为它要等待每个请求都完成才能进行下一个。这时，为了提高并发解析速度，我们可以使用多线程技术来并发处理多个解析请求。但是，如果同时启动过多的线程，会导致系统资源的不足，从而影响整个系统的性能。

因此，concurrentThreadsLimit函数主要通过系统资源的限制来限制并发解析线程的数量，确保系统的性能可以得到合理的平衡。默认情况下，Unix系统限制并发解析线程的数量为512个。这个值可以根据实际情况进行调整。



