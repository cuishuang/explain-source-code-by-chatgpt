# File: cgo_unix_cgo_res.go

cgo_unix_cgo_res.go这个文件是Go标准库中网络相关的文件，主要作用是在Unix系统上使用cgo调用底层的网络库，提供底层的网络支持。

在Unix系统上，Go语言使用cgo技术调用底层的系统函数来实现网络相关的操作，例如socket、bind、connect等。cgo_unix_cgo_res.go文件就是其中的一个用于支持cgo的文件，它定义了一系列用于调用底层函数的C语言函数，并使用Go语言对这些C函数进行了封装，提供了更为易用的函数接口。

具体来说，cgo_unix_cgo_res.go这个文件定义了以下内容：

1. 一系列C语言函数，例如socket、bind、connect、listen等，这些函数调用底层的系统函数来实现对网络的操作。

2. 一系列相应的Go函数，例如net.socket、net.bind、net.connect、net.listen等，这些函数封装了底层的C函数，提供更为易用的API接口。

3. 一些常量和变量定义，例如AF_UNIX、AF_INET、SOCK_STREAM、SOCK_DGRAM等，这些常量用于指定网络相关的选项、地址类型等。

cgo_unix_cgo_res.go这个文件的存在，使得Go语言在Unix系统上具有了强大而可靠的网络支持能力。因此，在编写Unix系统下的网络应用程序时，开发人员不必花费过多的时间和精力去实现底层的网络操作，而是可以直接使用Go标准库提供的API接口，从而更加高效地完成自己的任务。




---

### Structs:

### _C_struct___res_state

_C_struct___res_state结构体在Go的net包中主要用于存储DNS解析的状态和结果。其定义如下：

```go
// C struct for DNS resolver state.
// See /usr/include/resolv.h
type _C_struct___res_state struct {
    _C_long  retrans   int32
    _C_long  retry     int32
    _C_ulong options   uint32
    _C_long  nscount   int32
    _C_array [_C_int(_C__NSMAX)]_C_struct_in_addr nsaddr_list
    _C_array [_C_int(_C__MAXDOMAIN)]_C_char   defdname
    _C_struct___res_stateext ext
    _C_array [_C_int(_C__MAXRESOLVSORT)]_C_char  sort_list
    _C_array [_C_int(_C__MAXDNSRCH + 1)]_C_char   dnsrch
}
```

其中，各个字段的含义如下：

- retrans：重传DNS请求的等待时间。
- retry：在DNS请求失败后重试的次数。
- options：DNS请求的选项，如是否启用递归查询等。
- nscount：DNS服务器的数量。
- nsaddr_list：DNS服务器的IP地址列表。
- defdname：默认域名。
- ext：扩展的DNS状态信息。
- sort_list：DNS域名解析排序列表。
- dnsrch：DNS搜索列表，用于按照一定顺序向DNS服务器发出请求。

这些字段组成了一个完整的DNS解析状态和配置信息。在Go的net包中，该结构体通过C语言的数据结构实现，方便与系统底层DNS解析机制对接。



## Functions:

### _C_res_ninit

`_C_res_ninit`是一个在`net`包中使用的名为`cgo_unix_cgo_res.go`的C语言函数。其目的是初始化套接字地址和协议族的映射表，这些映射表用于解析主机名或服务名。

具体来说，该函数通过调用`getprotobynumber`、`getprotobyname`、`getservbyname`和`getservbyport`等函数，填充了具体的地址和协议族映射表。这样，在`net`包中调用诸如`net.Dial`和`net.Listen`等函数时，可以根据传递的主机名、服务名等参数，直接使用这些映射表来选择使用哪个协议族。

总之，`_C_res_ninit`函数的作用是初始化套接字地址和协议族的映射表，以实现DNS解析操作。



### _C_res_nclose

_C_res_nclose这个func是go/src/net/cgo_unix.go文件中定义的一个外部cgo库函数，作用是关闭一个关联到文件描述符的资源。

具体地说，在网络编程中，我们通常需要使用socket进行网络通信。socket是一个打开的文件描述符，使用完毕后需要显式地关闭。而_C_res_nclose就是用来关闭socket等文件描述符的资源的。

该函数接受一个整数参数，表示需要关闭的文件描述符。调用该函数后，它会使用系统调用close来关闭该文件描述符，并释放相关的资源。

在net包的一些方法中也会调用_C_res_nclose，以确保相关文件描述符被正确地关闭，防止产生内存泄漏等问题。



### _C_res_nsearch

在go/src/net中的cgo_unix_cgo_res.go文件中，_C_res_nsearch函数是用于从DNS服务器中查找与给定主机名相关联的资源记录的函数。

该函数使用C库函数res_nsearch()，其定义在<cresolv.h>头文件中。它需要一个主机名、类型和类，以及一个用于存储查找结果的空间。它将查询结果存储在空间中，并返回结果的数量，或者如果发生错误则返回-1。

在_C_res_nsearch函数中，go代码将参数转换为c语言可识别的类型，然后调用res_nsearch()函数。函数将得到的结果转换为go类型，然后返回给调用者。

该函数的作用是提供一种从DNS服务器中查找资源记录的方式，这是实现网络功能所必需的。



