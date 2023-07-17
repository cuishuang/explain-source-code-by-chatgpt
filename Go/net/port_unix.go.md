# File: port_unix.go

port_unix.go文件是Go语言标准库中net包的一个组成部分，主要用于处理Unix域socket。 Unix域socket是一种在本地计算机上使用套接字进行数据通信的机制，它可以在同一台计算机中的不同进程之间传输数据。该文件提供了Unix域socket处理所需的代码和函数。

该文件中定义了一系列函数，包括了Unix域socket的建立、连接、监听和接受等基本操作。其中，Unix域socket建立与TCP/IP网络socket建立有些不同，通过unix.Listen函数可以创建一个Unix域socket，而不是使用net.Listen函数创建TCP/IP网络socket。通过unix.Dial函数可以连接到另一个Unix域socket，同样不需要使用net.Dial函数。

与TCP/IP网络socket相比，Unix域socket的传输速度更快，因为它不需要进行网络传输。在跨平台开发中，使用Unix域socket也可以帮助程序更轻松地在不同操作系统和编程语言之间进行数据传输。

总的来说，port_unix.go文件的主要作用是提供Unix域socket的基本操作函数，帮助程序员更方便地使用该机制进行本地进程间的数据传输。




---

### Var:

### onceReadServices

在 Go 语言中，`net` 包中的 `port_unix.go` 文件定义了 Unix 平台下的网络端口相关的函数和结构体。其中 `onceReadServices` 是一个 `sync.Once` 类型的变量，其作用是用来确保只能读取一次系统服务文件 `/etc/services`。

在 Unix 系统中，`/etc/services` 文件定义了各种网络协议的端口号和对应的服务名称，`net` 包需要读取该文件来确定这些信息。由于该文件的内容相对稳定，因此只需要读取一次。而 `onceReadServices` 变量的作用就是保证只执行一次 `readServices()` 函数对 `/etc/services` 文件的读取，并将结果缓存下来以供后续使用，避免了重复读取和解析服务文件的开销。这样，每次访问 `/etc/services` 时就可以直接访问缓存了的数据，从而提高了程序的效率。



## Functions:

### readServices

readServices函数是用来读取/etc/services文件并解析其中的网络服务名称和对应的端口号的函数。它的作用是将服务名和端口号的对应关系保存在服务表中，以便客户端程序可以通过服务名获取对应的端口号。

具体来说，readServices函数打开/etc/services文件，并按行读取每一行的内容。对于每一行，它会忽略以#开头的注释行，并将服务名和端口号解析出来。将服务名和端口号保存到服务表中，以便日后使用。

在服务表中，服务名和端口号的对应关系以map[string]map[string]struct{}的形式保存。即服务名为一级键，端口号为二级键，value为一个空结构体。这样，在客户端程序中，可以通过服务名和端口号的组合来找到对应的端口号。同时，服务表中还保存了已经解析过的/etc/services文件的路径和最近修改时间，以便日后可以判断是否需要重新读取/etc/services文件。

总之，readServices函数的作用是解析/etc/services文件，并将服务名和端口号的对应关系保存到服务表中，以便客户端程序可以使用服务名来获取相应的端口号。



### goLookupPort

goLookupPort函数的作用是将给定的网络名称(network)和服务名称(service)对应到对应的TCP/IP端口号(port)。

该函数首先通过调用goLookupProtocol函数获取给定网络名称(network)的protocol number，即该网络所使用的协议的编号。然后根据给定的服务名称(service)以及获取到的protocol number，到/etc/services或者其它系统文件中查找对应的端口号(port)。

如果服务名称(service)为整数，则将其直接作为端口号(port)返回。如果找不到对应的端口号，则返回0。

如果网络名称(network)为"unix"，则服务名称(service)表示的是Unix套接字文件路径，因此直接返回该路径对应的文件地址。



