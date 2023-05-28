# File: dnsconfig_windows.go

dnsconfig_windows.go文件是Go语言标准库中net包下面的一个文件，负责在Windows系统上实现DNS的配置。

DNS（Domain Name System，域名系统）是将域名转换为IP地址的一种系统，它是互联网的重要基础设施之一。在Windows上，DNS的配置可以通过多种方式实现，比如手动设置DNS服务器、自动获取DNS服务器等等。而dnsconfig_windows.go文件就是为了方便Go语言在Windows系统上对DNS进行配置而设计的。

具体来说，dnsconfig_windows.go文件实现了以下功能：

1. 获取当前计算机上DNS服务器的配置信息，包括DNS服务器的IP地址、是否启用DNS故障转移、是否启用NetBIOS等信息。

2. 设置DNS服务器的配置信息，包括设置DNS服务器的IP地址、启用/禁用DNS故障转移、启用/禁用NetBIOS等。

3. 更新当前计算机上DNS服务器的配置信息，包括添加/删除DNS服务器的IP地址、启用/禁用DNS故障转移、启用/禁用NetBIOS等。

4. 获取当前计算机上的网络适配器的配置信息，包括适配器的名称、状态、连接类型、默认网关等。

在实现上述功能时，dnsconfig_windows.go文件使用了Windows系统提供的一些API，比如GetAdaptersAddresses、DnsQueryConfig、DnsAcquireContextHandle等。这些API使得Go语言可以方便地访问Windows系统的DNS配置信息，并对其进行修改和更新。

总之，dnsconfig_windows.go文件提供了对Windows系统DNS配置的操作接口，方便Go语言程序在Windows平台上进行网络编程。

## Functions:

### dnsReadConfig

dnsReadConfig函数的作用是在Windows系统上读取DNS配置。Windows系统中的DNS配置存储在注册表中，这个函数会读取注册表中的配置信息，并将其转换为Go语言中的结构体类型。

具体来说，dnsReadConfig函数会读取以下注册表项：

- HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\Tcpip\Parameters\NameServer
- HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\Tcpip\Parameters\DhcpNameServer
- HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\Tcpip\Parameters\Domain
- HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\Tcpip\Parameters\SearchList

这些注册表项存储了DNS服务器的IP地址、域名和搜索列表等信息。读取这些信息后，dnsReadConfig函数将其存储在一个结构体中，并返回该结构体的指针，供其他函数使用。

需要注意的是，dnsReadConfig函数只能在Windows系统上使用，如果在其他系统上调用该函数，则会返回错误信息。



