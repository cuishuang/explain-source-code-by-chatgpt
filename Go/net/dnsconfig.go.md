# File: dnsconfig.go

dnsconfig.go 是 Go 语言中 net 包中的一个文件，其主要功能是管理 DNS 相关的配置信息。

在该文件中，定义了一个名为 DNSConfig 的结构体，用于存储 DNS 解析相关的配置信息，这些信息包括：域名解析超时时间、最大 DNS 报文长度、是否允许 IP 分片、是否启用 IPv6、是否启用 EDNS、是否启用 DNSSEC 等。

此外，dnsconfig.go 中还定义了一些函数，用于设置和获取 DNS 配置信息。例如，通过 `func SetDNSConfig(config *DNSConfig)` 函数可以设置 DNS 配置信息，而通过 `func DNSConfig() *DNSConfig` 函数可以获取 DNS 配置信息。

同时，dnsconfig.go 中还定义了一些常量，用于指定默认的 DNS 配置参数，例如：

```go
const (
    defaultTimeout            = 2 * time.Second
    maxDnsMessageSize         = 65536
    defaultUseIPv6            = true
    defaultSingleResolverFile = "/etc/resolv.conf"
)
```

这些常量为 DNS 配置参数提供了默认值，如果用户没有指定相应的参数，那么就会使用这些默认值。

总之，dnsconfig.go 是 net 包中非常重要的一个文件，它提供了对 DNS 配置信息的管理，使得 Go 语言中的网络连接可以更加稳定和可靠。




---

### Var:

### defaultNS

defaultNS是一个字符串切片，存储了默认的DNS服务器地址。在进行DNS查询时，如果未指定DNS服务器地址，则会使用defaultNS指定的地址进行查询。

该变量的默认值是空切片，也就是说如果用户未指定DNS服务器地址，则会使用操作系统的默认设置进行查询。

用户也可以通过net.Resolver中的Dial函数显式指定DNS服务器地址，如果指定了则会覆盖defaultNS的设置。

在实际的网络应用中，DNS查询是一个非常常见的操作。通过设置defaultNS，可以方便地指定默认的DNS服务器地址，避免在每次进行DNS查询时都手动指定DNS服务器地址，提高代码的可读性和简洁性。



### getHostname

在dnsconfig.go文件中，getHostname变量用于检索本地主机名。该变量是一个函数，它的作用是尝试使用一些不同的机制来获取本地主机名，例如：

1.从环境变量中获取主机名
2.从系统配置中获取主机名
3.尝试连接到本地的套接字并获取主机名
4.使用系统调用获取主机名

getHostname函数会按照上述顺序依次尝试这些机制，并且在任何一种机制成功获取主机名之后，就会返回主机名。如果所有机制都失败了，getHostname函数会返回一个空字符串，表示无法获取主机名。通过使用getHostname变量，代码可以轻松地获取本地主机名并将其用于DNS查询等目的。






---

### Structs:

### dnsConfig

dnsConfig结构体定义了一个DNS配置，其中包括了DNS服务器地址和是否从主机的/etc/hosts文件中查找域名解析。在go/src/net中的其他文件中，会使用这个结构体来设置网络连接的DNS配置。

具体来说，dnsConfig结构体有以下字段：

- servers []string：DNS服务器地址列表。在解析域名时，将首先尝试从这些地址中的任一地址获取IP地址。如果该列表为空，则使用默认的DNS服务器地址。
- searches []string：搜索列表。当解析不包含域名的主机名时，将使用这些搜索列表进行搜索。例如，如果主机名为"foo"，并且搜索列表为["example.com"]，则尝试使用"foo.example.com"解析域名。
- ndots int：n个点的阈值。当解析不包含"."的主机名时，如果该主机名中的非"."字符数少于ndots，则会自动尝试将该主机名追加到搜索列表中的每个域名后。如果该值为0，则禁用自动添加域名。
- useCgo bool：是否通过CGO使用C系统中的getaddrinfo函数进行DNS解析。默认情况下，go会使用纯go实现的解析函数。设置useCgo为true，则会使用系统C库中的实现，这可能会更快，但也有风险。

总之，dnsConfig结构体的作用是为网络连接提供DNS配置，以便在进行域名解析时使用。



## Functions:

### serverOffset

在 Go 的 net 包中，dnsconfig.go 文件中的 serverOffset 函数用于计算 DNS 解析时获取 DNS 服务器地址的偏移量。

在 DNS 报文中，DNS 服务器地址可以有多个（一般情况下有两个），通过 serverOffset 函数计算偏移量，程序才能正确地获取 DNS 服务器地址。

具体来说，该函数通过读取 DNS 报文的 Header 中的两个字段（Question、Answer），计算出 DNS 报文中包含的 DNS 服务器地址的数量。然后，该函数使用这个数量和报文头的长度来计算 DNS 报文中 DNS 服务器地址的偏移量。

在 DNS 报文中，DNS 服务器地址的偏移量往往是固定的，因此通过 serverOffset 函数计算出 DNS 服务器地址的偏移量后，就可以方便地通过解析 DNS 报文获取 DNS 服务器地址，并将其用于进行 DNS 请求。



