# File: sockcmsg_unix.go

sockcmsg_unix.go是Go语言中syscall包中的一个文件，该文件实现了Unix环境下使用SCM_RIGHTS socket控制消息的发送和接收。

SCM_RIGHTS是Unix环境下的一种socket控制消息，它用于将发送方中一个已打开的文件描述符传递给接收方的某个进程。sockcmsg_unix.go中的函数实现了将文件描述符发送给另一进程并在另一进程中接收该文件描述符的过程。

具体来说，sockcmsg_unix.go中的函数实现了以下功能：

1. 接收方调用recvmsg函数接收包含SCM_RIGHTS控制消息的socket消息，并从中获取发送方传递的文件描述符。

2. 发送方调用sendmsg函数将包含SCM_RIGHTS控制消息的socket消息发送给接收方，并将文件描述符传递给接收方。

3. 接收方调用Cmsghdr函数对接收到的控制消息进行解析，并获取发送方传递的文件描述符。

总之，sockcmsg_unix.go的作用是实现了Unix环境下进程间传递已打开的文件描述符的功能，可以方便地实现进程间的文件共享或进程间的文件传输。




---

### Structs:

### SocketControlMessage

SocketControlMessage结构体用于表示控制消息。在Linux系统中，控制消息通常用于在进程间传递辅助数据或操作控制信息，例如在Unix域套接字（Unix domain socket）中，控制消息可用于传递文件描述符。

SocketControlMessage结构体定义如下：

```
type SocketControlMessage struct {
    Header ControlMessageHeader
    Data   []byte
}
```

其中，Header字段是ControlMessageHeader类型，用于表示控制消息头部信息，如消息类型、消息长度等。Data字段是字节切片类型，用于表示控制消息的数据部分。

这个结构体属于syscall包，通常应用于低层次的网络编程中。在传输控制协议（Transmission Control Protocol，TCP）和用户数据报协议（User Datagram Protocol，UDP）等协议中，SocketControlMessage可用于传递外带数据（out-of-band data），以及其他控制信息。

在Unix域套接字中，SocketControlMessage可用于传递文件描述符。在此种情况下，Data字段中存储的是文件描述符对应的整数值，在接收端可以通过将该整数值转换为文件描述符来获取相应的文件描述符句柄。



## Functions:

### CmsgLen

在 Unix 系统中，控制消息（Control Message）是一组用于传递额外数据的消息，这些消息与普通的数据包不同，通常用于传递底层网络协议的元数据或其他系统信息。CmsgLen（Control Message Length）是一个函数，它用于计算控制消息的实际长度，以便正确分配内存和传递给底层的系统调用函数。

在 sockcmsg_unix.go 文件中，CmsgLen 函数接收一个或多个控制消息数组，然后迭代每个控制消息以计算它们的长度，并将它们的长度相加。具体而言，该函数会在迭代期间检查每个控制消息的套接字协议（即“协议控制字”，Protocol Control Word，PCW）中是否包含设置了 CMSG_LEN 选项的数据。如果是，函数则读取数据的 CMSG_LEN 字段并将其添加到总长度中；否则，函数将使用 sizeof 函数计算该控制消息的长度并添加到总长度中。

总之，CmsgLen 函数是一个方便的工具函数，它简化了在使用控制消息时计算它们长度的过程。在 Unix 套接字编程中，具有了 CmsgLen 函数，开发者可以更轻松地实现高级网络协议和通信功能。



### CmsgSpace

CmsgSpace是一个函数，用于计算在Unix域套接字协议中传递控制消息时需要的缓冲区大小。

Unix域套接字协议允许应用程序在套接字之间发送附加信息，这些信息称为控制消息。控制消息由 ancillary data（辅助数据）和 control message header（控制消息头）组成，并且用于支持一些高级功能，例如带外数据、路径负载均衡和身份验证等。

CmsgSpace函数通过计算附加数据的大小，以及确定是否需要添加空间来对齐控制消息的头信息，来确定传递控制消息所需的缓冲区大小。该函数接收附加数据大小、控制消息头类型和控制消息头长度作为参数，然后返回一个整数表示所需的缓冲区大小。

CmsgSpace函数被用于syscall包中的Unix域套接字相关函数（例如Recvmsg和Sendmsg），以确保在传递控制消息时，缓冲区足够大，以便传递所有相关信息。



### data

在syscall包的sockcmsg_unix.go文件中，data函数的作用是将控制消息中的数据部分解析为一个给定类型的切片。控制消息是传递给Unix域Socket API的附加信息，这些信息可能包含文件描述符、交换数据等。

该函数的原型如下：

```go
func data(m *SocketControlMessage, typ uintptr) ([]byte, error) 
```

其中包含两个参数：

- m：指向SocketControlMessage结构体的指针，表示要解析的控制消息。
- typ：uintptr类型，表示要解析的数据类型。

函数返回两个值：

- 一个切片，包含控制消息中的数据部分。
- 一个错误，表示在解析过程中是否出现了错误。

该函数会遍历控制消息中的所有元素，直到找到类型与typ匹配的元素并返回其数据部分。如果在控制消息中没有找到匹配的元素，则返回一个错误。

该函数与其他控制消息处理函数（如UnixRights）一起使用，可以实现在Unix域Socket API中进行高级操作，如文件描述符交换等。



### ParseSocketControlMessage

ParseSocketControlMessage是一个函数，它的作用是解析UNIX域socket控制消息。在UNIX域socket通信中，控制消息是可选的数据部分，用于携带一些特殊的信息，如文件描述符等。

具体来说，ParseSocketControlMessage函数接受一个[]byte类型的参数msg，并返回一个SocketControlMessage类型的值。这个函数首先会根据msg中的数据解析出控制消息的头部部分，包括长度和消息类型。

接着，这个函数会根据消息类型将控制消息的数据部分解析成不同的格式。例如，如果消息类型是SCM_RIGHTS，则表示消息携带了一个或多个文件描述符，这个函数就会将这些文件描述符逐一解析并封装到SocketControlMessage结构体中返回。

总之，ParseSocketControlMessage函数的作用是对UNIX域socket的控制消息进行解析，方便用户在socket通信中传输和接收一些特殊的数据。



### socketControlMessageHeaderAndData

socketControlMessageHeaderAndData函数是用于解析套接字控制消息的函数，它读取一段长度为n的字节数组，然后解析出控制消息头和数据，并返回它们的指针和长度。

这个函数的参数包括：

- b []byte: 用于存储套接字控制消息的字节数组。
- sa: 用于传递控制消息的通用类型的指针。
- typ: 控制消息的类型。
- isIPv6 bool: 是否是IPv6套接字。

函数首先会根据isIPv6判断传入的sa是IPv4还是IPv6类型的，然后根据传入的typ类型来初始化不同的控制消息头。接着，它使用二进制序列化的方式从b中读取出控制消息头的内容，并进行处理。

如果控制消息头中包含了数据，则函数会继续从b中读取出这些数据，并返回控制消息头和数据部分的指针和长度。

最后，函数会检查b中是否还有未处理的字节，如果有，则说明发生了错误，返回一个错误。如果没有，则返回解析后的控制消息头和数据。



### UnixRights

UnixRights函数是一个帮助函数，用于将Unix域套接字的文件描述符数组封装成控制消息（Control Message）。控制消息（Control Message）是在Unix域套接字（Unix Domain Socket）之间发送额外信息的一种机制，可以用于传递文件描述符，选项信息等等。

UnixRights函数将文件描述符数组打包成控制消息的辅助函数。其原型如下：

```go
func UnixRights(fd int) []byte
```

其中，fd是要传递的文件描述符数组的第一个文件描述符（fd[0]），因为控制消息只能携带一个文件描述符数组。

UnixRights函数返回一个字节切片（[]byte），这个字节切片的内容是打包好的控制消息。打包好的消息可以直接作为Unix域套接字的writev函数（参照Linux 2.6版本以上）的第3个参数（iov[2].iov_base）发送。

控制消息的格式为：

```c
struct cmsghdr {
    socklen_t cmsg_len;   // 控制消息的总长度
    int       cmsg_level; // 协议层（SOL_SOCKET等）
    int       cmsg_type;  // 控制类型（SCM_RIGHTS等）
} cm;

FDs := []int{} // 传递的文件描述符数组

msg := []byte{} // 整个控制消息的二进制形式

hdrlen := C.CMSG_LEN(len(FDs) * 4) // 实际控制头部长度（一个int=4byte）
mlen := hdrlen + len(FDs) * 4   // 总长度
msg = make([]byte, mlen)
cm := (*C.struct_cmsghdr)(unsafe.Pointer(&msg[0]))
cm.cmsg_len = hdrlen + len(FDs) * 4
cm.cmsg_level = C.SOL_SOCKET
cm.cmsg_type = C.SCM_RIGHTS

// 传递的文件描述符数组写入控制消息
copy(msg[hdrlen:],(*[len(FDs)]byte)(unsafe.Pointer(&FDs[0]))[:])
```

UnixRights的实现代码如下：

```go
func UnixRights(fd int) []byte {
    msg := make([]byte, 1)
    p := uintptr(unsafe.Pointer(&msg[0]))
    _, _, err := syscall.RawSyscall(
        unix.SYS_IOCTL,
        uintptr(fd), 
        uintptr(unix.TIOCUCNTL(UnixRightsType)), 
        p,
    )
    if err != 0 {
        panic(err)
    }
    return msg[:1]
}
```

该实现使用了UnixRights类型的ioctl控制命令，将fd封装成控制消息并返回。在具体实现中，RawSyscall函数用于执行这个ioctl命令。

其中，syscall.RawSyscall是一个低级别的API，用于调用系统调用（系统级别函数）。

UnixRightsType是指控制消息类型常量SCM_RIGHTS。

总的来说，UnixRights函数的作用是将一个文件描述符数组封装到控制消息中，以便在Unix域套接字（Unix Domain Socket）之间传递文件描述符。



### ParseUnixRights

ParseUnixRights是一个函数，它的作用是从传入的unix控制消息中解析出文件描述符列表。unix控制消息是一个通信协议，它允许在进程之间传递描述符和其他控制信息。

在Unix系统中，文件描述符是一种标识符，用于唯一标识在进程中打开的文件或资源。通过传递文件描述符，一个进程可以向另一个进程传递打开文件的权限，以便后者可以访问这些文件。

在网络编程中，Unix域套接字（Unix Domain Socket）是一种常见的通信方式。当使用这种套接字进行进程间通信时，可以使用Unix控制消息来传递文件描述符。

ParseUnixRights函数解析传入的控制消息中包含的描述符，并返回一个整数列表，其中每个整数表示一个文件描述符。这些文件描述符可以用于后续的操作，例如读取或写入数据。

总之，ParseUnixRights函数是一个解析unix控制消息中文件描述符的功能，它在网络编程和进程间通信中扮演重要角色。



