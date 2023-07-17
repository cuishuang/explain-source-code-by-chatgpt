# File: route_freebsd_32bit.go

route_freebsd_32bit.go文件是用于提供 Freebsd 操作系统的路由接口的功能的。该文件实现了针对 32 位系统的相关路由操作，包括增加、删除和获取路由表中的路由信息以及设置网关等。具体来说，该文件定义了以下函数：

1. SysctlRtsock()：该函数用于查询和设置系统的路由信息。可以通过指定不同的操作码实现获取路由信息（RTM_GET）或设置路由信息（RTM_ADD、RTM_DELETE、RTM_CHANGE 等）。
2. parseRouteMessage()：该函数用于解析路由信息，将系统返回的二进制数据转换为可读取的结构体。
3. addRoute()：该函数用于添加新的路由信息到系统路由表中。
4. deleteRoute()：该函数用于从系统路由表中删除指定的路由信息。
5. getRoutes()：该函数用于获取系统当前路由表中的所有路由信息。获取到的路由信息会被返回为一个结构体切片。

总的来说，route_freebsd_32bit.go 文件提供了一些基本的路由操作函数，方便在 Go 语言程序中使用 Freebsd 的路由功能来实现网络相关的操作。

## Functions:

### parseRouteMessage

parseRouteMessage这个函数是用于解析FreeBSD系统中的路由消息的函数。在FreeBSD系统中，路由消息是通过AF_ROUTE socket类型发送和接收的。

该函数的主要作用是解析路由消息中的各个字段，例如路由表项的目标网络地址、掩码、网关、优先级等信息，并将这些信息存储在结构体RouteMessage中返回。同时，该函数还会对路由表变化事件进行分类，目前支持以下事件类型：

- RTM_ADD：添加一个路由表项
- RTM_CHANGE：修改一个路由表项
- RTM_DELETE：删除一个路由表项
- RTM_GET：获取指定的路由表项
- RTM_LOSING：告警，路由器正在失去连接
- RTM_REDIRECT：该路由表项被重定向到另一个网关
- RTM_MISS：该路由表项匹配不到目标

在路由表变化事件发生时，该函数还会调用回调函数，以便应用程序可以对它们做出响应，例如更新本地路由表或发送路由变更通知。

总之，parseRouteMessage函数对于实现FreeBSD系统中路由管理功能至关重要，它负责解析路由消息，提取路由表项信息并进行分类处理，使得应用程序可以及时地响应路由表变化事件。



### parseInterfaceMessage

parseInterfaceMessage函数用于解析FreeBSD系统接口消息的数据结构ifMsghdr32和ifMsghdr64，并将它们转换为统一的接口消息ifMsghdr类型。该函数的作用是在读取系统接口状态信息时，将不同版本的数据结构统一为同一种数据结构，方便后续的数据处理和使用。

具体来说，该函数首先判断接口消息的数据结构类型，如果是32位的ifMsghdr32，就将其中的字段转换为ifMsghdr类型中的字段，然后返回ifMsghdr类型的接口消息。如果是64位的ifMsghdr64，也会进行类似的转换。在转换过程中，该函数还会根据接口消息中的标志位（flags）来设置接口消息中的掩码值（mask），从而方便后续对接口消息的处理和使用。

总之，parseInterfaceMessage函数主要是用于将不同版本的接口消息数据结构统一为同一种数据结构，从而简化接口状态信息的处理过程。



