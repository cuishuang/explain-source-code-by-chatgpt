# File: netlink_linux.go

netlink_linux.go文件是Go语言中syscall包的一部分，是针对Linux下使用Netlink通信机制的实现文件。Netlink是Linux内核提供的一种消息传输机制，可以用于内核与用户空间进程之间的通信。

该文件主要实现了Netlink的相关接口，如创建socket、发送和接收消息等。其中，包含了多个结构体用于表示Netlink消息头、数据报文等信息，并且对这些数据结构进行了具体的字段解析。

此外，该文件还提供了一些实用的函数，如发送Netlink消息、接收Netlink消息、解析Netlink消息等。这些函数可以帮助用户更方便地使用Netlink进行数据传输。

总体来说，netlink_linux.go文件的作用就是提供了Go语言与Linux内核进行通信的基础工具，使得程序员可以更加方便地使用Netlink机制进行进程间通信。




---

### Var:

### pageBufPool

在netlink_linux.go文件中，pageBufPool是一个变量，用于存储用于Netlink消息的缓冲区。Netlink协议是用于在Linux内核和用户空间之间进行通信的协议。在这种通信中，内核将数据发送到用户空间的一个套接字中，用户空间可以使用该套接字来读取和写入数据。

由于Netlink消息可能非常大，因此需要使用缓冲区来存储它们。这个缓冲区可以使用malloc函数分配，但是这样做会导致内存碎片和额外的开销。为了避免这些问题，Go语言使用了一个叫做pageBufPool的缓冲池来管理Netlink消息的缓冲区。这个缓冲池使用了内存页（通常是4KB）作为缓冲区的大小，以确保缓冲区的大小与内存页的大小匹配。

当需要分配Netlink消息的缓冲区时，使用pageBufPool.Get()方法从缓冲池中获取一个缓冲区。使用完缓冲区后，可以使用pageBufPool.Put()方法将缓冲区返回到缓冲池中以供重复使用。这样做可以避免使用malloc函数导致的内存碎片和额外的开销，因此可以显著提高程序的性能和效率。






---

### Structs:

### NetlinkRouteRequest

NetlinkRouteRequest是一个结构体，用于表示Netlink协议族中的路由请求。路由请求是在Linux系统中管理路由表的一种方式。该结构体中包含了路由请求的各种参数和选项，例如路由表类型、目标地址、网关地址、操作类型等等。

具体来说，NetlinkRouteRequest结构体的定义如下：

type NetlinkRouteRequest struct {
    Header NlMsghdr
    Rtmsg  RtMsg
    // ...
}

其中，Header字段表示Netlink协议消息头，包含消息类型、长度等信息；Rtmsg字段表示路由请求的消息体，包含路由表类型、目标地址、网关地址、操作类型等信息。

通过Netlink协议族发送和接收路由请求，可以实现添加、删除、修改系统中的路由表。这对于网络编程和系统管理很有用，例如可以通过这个功能实现NAT、VPN、路由选择等网络技术。

总之，NetlinkRouteRequest结构体是在Linux系统中管理路由表时使用的重要工具。它通过Netlink协议族实现了路由表的增删改查等操作，为网络编程和系统管理提供了便利和灵活性。



### NetlinkMessage

NetlinkMessage是一个结构体，用于表示在Linux内核与用户空间之间进行通信时使用的网络链接消息。

在Linux内核中，Netlink协议是用于内核和用户空间之间传递网络协议消息的一种机制。NetlinkMessage结构体定义了接收和发送消息时使用的格式。它包括以下字段：

- Length: 消息的总长度（包括头部和负载内容）。
- Type: 消息类型，指定发送的消息是控制消息还是数据消息等。
- Flags: 控制消息和数据消息操作的标志（例如，请求或响应）。
- Sequence: 发送消息的序列号，用于指示收到的响应是对哪个请求的响应。
- PID: 发送消息的进程ID，用于回复响应。

NetlinkMessage结构体还有一个Payload字段，它包含实际的消息负载。在发送NetlinkMessage时，应发送有效负载的字节数量。在接收NetlinkMessage时，可以使用Payload字段来解码这些字节。

总之，NetlinkMessage结构体是在Linux内核和用户空间之间传递的网络协议消息的基本表示，它定义了消息的格式和发送/接收时需要使用的字段。



### NetlinkRouteAttr

NetlinkRouteAttr这个结构体是用来表示路由信息的属性的。路由信息是用于管理网络连接的信息，包括目的地网络地址、下一跳地址、优先级等等参数。

NetlinkRouteAttr包含了一些重要的字段，如Type、DataLen和Data，它们表示了路由信息的类型、数据的长度和具体的数据内容。

在网络编程中，NetlinkRouteAttr可以被用来获取或设置路由信息，例如添加一个新的路由规则、删除已有的路由规则等等操作。具体的操作需要使用系统调用来实现，例如用NetlinkRouteAttr来设置规则，可以使用netlink.RouteAdd()方法。

这个结构体还有很多其他的应用，例如在iptables等防火墙软件中，也会使用NetlinkRouteAttr来控制网络连接的访问权限。总体来说，NetlinkRouteAttr是网络编程中的一个重要的数据结构，用于处理路由信息。



## Functions:

### nlmAlignOf

nlmAlignOf函数用于计算给定类型的字节对齐方式。它的作用是确保在使用网络套接字的时候，结构体的成员的布局能够正确地对齐，以避免出现数据不对齐的问题，从而提高数据的读取和写入效率。

具体来说，nlmAlignOf函数接收一个类型作为参数，然后使用类型的对齐方式和大小来计算给定类型的字节对齐方式。这个字节对齐方式可以用来确保任何一个包含该类型成员的结构体都按照正确的字节对齐方式进行布局。这在网络通信中十分重要，因为如果数据不按照正确的字节对齐方式进行布局，可能会导致性能下降或数据损坏等问题。

总之，nlmAlignOf函数的作用是提供一个通用的计算字节对齐方式的方法，以确保网络套接字的数据传输效率和正确性。



### rtaAlignOf

在netlink_linux.go文件中，rtaAlignOf是一个函数，其作用是计算给定数据类型的内存对齐方式。

在Linux中，内存对齐是十分重要的一个概念。内存对齐是指在数据结构中，把数据元素存放在某个类型所规定的地址处，使数据存取更高效。在访问存储器时，内存访问对齐会大大提高CPU的效率，降低内存负荷和磁盘I/O开销。

在网络编程中，经常需要处理各种控制数据。这些控制数据是按照特定的格式组织的，并且包含了各种不同大小和类型的元素。因此，为了正确地解析这些数据，必须了解每个元素的大小和偏移量。

在这种情况下，rtaAlignOf函数是用来确定网络控制块数据结构中某个元素的偏移量和对齐方式的。它通过计算一个元素的大小和对齐方式，返回一个适当的偏移量，以便正确地访问控制数据结构中的这个元素。它的具体实现会根据不同的数据类型，计算出该类型所需的对齐方式，这样可以确保在对数据进行访问时，每个数据元素都具有相同的内存对齐方式，从而提高访问效率。



### toWireFormat

toWireFormat是用于将netlink消息转换为字节序列的函数。在netlink通信中，消息必须转换为特殊的二进制格式才能在系统之间传递。

该函数接收一个netlink消息结构体并将其转换为字节序列。该结构体包含消息的各个字段，例如消息类型、flags、PID等。 在转换时，它使用了一些特定的编码规则来确保消息的正确格式，包括字节顺序、长度、填充等。

toWireFormat函数的作用是将netlink消息转换为字节序列，以便可以通过套接字在系统之间传递。 它是在syscall包的netlink_linux.go文件中定义的，用于在Linux系统上使用netlink通信的应用程序中编码和解码netlink消息。



### newNetlinkRouteRequest

newNetlinkRouteRequest是一个函数，用于创建一个新的Netlink Route请求。Netlink Route请求是用于向内核发送路由表更新的请求，以便更新路由表信息。这个函数主要用于创建路由请求的数据包，填充正确的数据包头和数据包负载，这些数据将被发送到内核。如果创建成功，函数将返回一个指向Netlink Route请求的指针。具体来说，这个函数的作用包括以下几个方面：

1. 创建Netlink Route请求：函数将创建一个新的Netlink Route请求，并填充正确的数据包头和数据包负载。

2. 发送路由表更新请求：Netlink Route请求是用于向内核发送路由表更新的请求，以便更新路由表信息。这个函数将创建路由请求的数据包，填充正确的数据包头和数据包负载，这些数据将被发送到内核。

3. 填充正确的数据包头和数据包负载：函数将根据参数中指定的操作类型和消息类型，填充正确的数据包头和数据包负载。

4. 返回Netlink Route请求指针：如果创建成功，函数将返回一个指向Netlink Route请求的指针，可以将该指针分配给其他变量或用于其他对Netlink Route请求的操作。

总之，newNetlinkRouteRequest是一个重要的函数，用于创建一个新的Netlink Route请求，以便更新路由表信息，它为路由表的操作提供了底层支持。



### NetlinkRIB

NetlinkRIB是一个用于查询Linux系统路由表信息的函数。它使用Netlink协议与内核通信，发出查询请求并接收响应，获取与路由有关的信息，例如路由表、网关、子网掩码、接口等。

在Linux系统中，路由表用于决定哪个网关将被用于发送给定目标的数据包。路由表的查询和修改是网络管理员进行路由设置和故障排除的重要工作。NetlinkRIB函数使得程序能够从内核中读取路由表的信息，便于进行网络配置和管理。

NetlinkRIB函数的具体实现细节包括发送Netlink消息、解析内核响应、处理错误状态等。它还提供了一些辅助函数，例如ParseRouteAttr用于解析内核返回的路由属性，GetLinkByName用于根据接口名称获取接口ID等。



### ParseNetlinkMessage

ParseNetlinkMessage是一个函数，它的作用是将从Netlink Socket接收到的消息解析为一个NetlinkMessage类型的变量。

Netlink是Linux内核的一种机制，用于在内核和用户空间之间传递网络协议消息。Netlink Socket是Linux中的一种特殊socket，它与内核通信。当内核需要将一个网络协议消息发送给用户空间时，它会通过Netlink Socket发送该消息。

ParseNetlinkMessage函数的参数是一个字节数组，它包含了从Netlink Socket接收到的消息。该函数首先会检查该消息的长度是否为合法值，并将其转换成一个NetlinkMessage类型的结构体变量。

NetlinkMessage结构体变量包含了所有与该消息相关的信息，包括消息类型、序列号、标志等信息。该函数还会解析消息中的所有NetlinkAttribute（可选的附加数据），并将其存储在NetlinkMessage变量的Attributes字段中。

通过调用ParseNetlinkMessage函数，我们可以方便地将从Netlink Socket接收到的消息解析为一个易于理解的格式，从而使我们能够更轻松地处理和分析该消息，进而实现网络协议的交互和数据传输。



### netlinkMessageHeaderAndData

netlinkMessageHeaderAndData是一个函数，它的作用是将一个netlink消息的头部和数据部分拼接成一个byte slice返回。

在Netlink协议中，每个消息由一个头部和一个数据部分组成，头部包含了消息的属性，如消息类型、源地址、目标地址等，而数据部分则包含了具体的消息内容。netlinkMessageHeaderAndData函数就是将这个头部和数据部分拼接成一个byte slice返回。代码实现如下：

```
// netlinkMessageHeaderAndData returns a byte slice containing the concatenated
// netlink message header and payload data.
func netlinkMessageHeaderAndData(hdr *NetlinkMessageHeader, data []byte) []byte {
    hlen := headerLength(hdr)
    buf := make([]byte, hlen+len(data))
    binary.LittleEndian.PutUint32(buf[0:4], uint32(hlen+len(data)))
    binary.LittleEndian.PutUint16(buf[4:6], uint16(hdr.Type))
    binary.LittleEndian.PutUint16(buf[6:8], uint16(hdr.Flags))
    binary.LittleEndian.PutUint32(buf[8:12], uint32(hdr.Seq))
    binary.LittleEndian.PutUint32(buf[12:16], uint32(hdr.Pid))
    copy(buf[hlen:], data)
    return buf
}
```

该函数首先计算出头部长度，并根据头部和数据部分的长度创建一个byte slice。然后，将头部的各个属性依次拷贝到byte slice中，并将数据部分也拷贝到byte slice中。最后，将拼接后的byte slice返回。

该函数的作用主要是将一个netlink消息转换成byte slice，方便进行网络传输。在Netlink协议中，消息通常需要进行网络传输才能达到目的地，而网络传输的数据是以byte slice的形式进行的，因此需要将netlink消息转换为byte slice。



### ParseNetlinkRouteAttr

ParseNetlinkRouteAttr这个函数的作用是解析Netlink路由信息的属性，可以解析出该路由的各种属性，比如目的网络地址、下一跳地址、路由表ID、权重、优先级等。

在Netlink路由消息中，各个属性以Netlink属性的形式进行传输，因此需要使用ParseNetlinkRouteAttr函数将这些属性解析出来，以便后续对路由信息进行处理。

该函数主要包括以下步骤：

1. 遍历属性列表，按照属性类型逐个解析属性值。

2. 根据属性类型，判断该属性值的类型和长度，并按照相应的格式解析出属性值。比如，对于目的网络地址属性，解析出该属性值的IP地址和子网掩码。

3. 将解析出的属性值存储到对应的变量中，最终返回一个RouteAttrs类型的结构体，其中包含解析出的各个属性值。

ParseNetlinkRouteAttr函数在Netlink路由模块中起着非常重要的作用，可以帮助开发人员更方便地获取和处理路由信息，提高路由处理效率。



### netlinkRouteAttrAndValue

netlinkRouteAttrAndValue函数是用来解析和处理netlink消息中的路由属性和对应的值的。它接收的参数包括一个字节数组，代表一个路由属性值对应的二进制数据，返回值包括一个路由属性和对应的值。

具体来说，netlinkRouteAttrAndValue函数从字节数组中读取路由属性和值，并将其解析成一个netlinkRouteAttr类型的结构体。该结构体包含了路由属性的类型、长度和具体的值，以及一些辅助方法用于处理和转换属性值。

在处理netlink消息的过程中，通常需要使用netlinkRouteAttrAndValue函数来解析路由属性和对应的值，以便正确地理解和处理完整的netlink消息。



