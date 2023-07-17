# File: netpoll_stub.go

netpoll_stub.go文件位于Go语言的runtime包中，是为了在某些不支持系统调用的平台上提供网络IO的轮询服务。这个文件是一个桥接层，建立了底层网络轮询和Go语言的网络IO调用之间的连接。

在支持系统调用的平台上，网络IO的轮询是通过操作系统提供的select、epoll、kqueue等系统调用实现的。但是，在一些轻量级的嵌入式系统、手机平台等上，可能并不支持系统调用，这时候就需要使用不同的方式来实现网络IO的轮询。

netpoll_stub.go文件中提供了对网络IO轮询的常规操作的桥接，包括启动轮询、等待事件、处理IO事件等。 Go语言中的网络IO调用，如socket读写、接收、发送等，都是和这个文件中的底层轮询服务进行交互的。在轻量级的平台上，这个文件可以与自定义的网络轮询库一起使用来实现网络IO。

总之，netpoll_stub.go文件是Go语言实现网络IO调用的重要文件，它提供了在各种平台下实现网络轮询的框架，使得Go语言的网络IO能够在不同的嵌入式系统和手机平台上得到支持。




---

### Var:

### netpollInited

netpollInited是runtime/netpoll包中的一个全局变量，其作用是用来标记网络轮询器（netpoller）是否已经初始化完成。

在Go语言中，网络轮询器充当了一种实现异步I/O的关键角色。它基于操作系统提供的I/O多路复用机制，比如epoll、kqueue等，用来监控文件描述符或网络连接的事件，当有事件发生时，通知Go程序进行相应的处理。但在初始化阶段，网络轮询器需要完成一些和网络相关的初始化工作，如获取系统相关的参数或设置网络环境变量等。因此，需要有一个标志来表示网络轮询器的初始化状态。

具体来说，在netpoll_stub.go中，变量netpollInited被定义为一个bool类型的值，初始值为false，表示网络轮询器未初始化。在初始化完成后，会将其值设置为true。这个变量的主要作用是用于对调用netpoll开头的函数做出限制，因为这些函数需要在网络轮询器初始化完成的前提下才能使用，否则会导致程序意外退出。

总之，netpollInited这个变量的存在是为了确保网络轮询器的正确使用和运行，在Go中扮演着重要的角色。



### netpollWaiters

netpoll_stub.go文件是用来实现网络轮询机制（netpoll）的协程调度器。其中，netpollWaiters变量是一个存储等待轮询的Goroutine的列表。

它的作用是存储需要等待网络IO事件的协程（Goroutine），并在事件发生时通知它们恢复运行。在网络轮询器（netpoll）中，当发现没有任何IO事件需要处理时，就会从列表中取出一个等待的Goroutine恢复运行，使其继续执行。

具体实现中，当一个Goroutine需要等待一个网络IO事件时，它会将自己加入到netpollWaiters列表中，并调用gopark函数将自己挂起。当事件发生时，网络轮询器会将列表中的Goroutine重新放入调度队列中，使其继续运行。

因此，netpollWaiters变量的作用是存储等待网络IO事件的Goroutine，同时它也是网络轮询机制实现中的重要数据结构之一，可以使网络IO事件的处理更为高效和准确。



### netpollStubLock

netpollStubLock是一个互斥锁，用于保护netpollStub这个结构体的修改和访问。netpollStub是网络轮询的实现，它包含了一个epoll实例、网络事件队列和等待队列等属性，用于处理网络请求。由于多个goroutine可能会并发地调用网络轮询，因此需要一个互斥锁来确保并发访问的安全性。

具体来说，当一个goroutine需要访问或修改netpollStub的属性时，必须先获取netpollStubLock这个互斥锁，以确保没有其他goroutine同时访问或修改netpollStub。当该goroutine完成访问或修改后，必须释放netpollStubLock，以允许其他goroutine访问或修改netpollStub。

总之，netpollStubLock的作用是确保对netpollStub的并发访问的安全性和一致性。



### netpollNote

netpollNote是一个用于网络轮询通知的goroutine中的变量。在Go语言的网络轮询中，如果没有任何新的I/O事件到达，则轮询器将阻塞，直到有新的事件到达。但在发生I/O事件时，轮询器需要通知goroutine以便处理I/O事件。因此，程序会利用netpollNote构建一个通知链表，以便在goroutine中执行网络轮询时通知所有等待I/O事件的goroutine。

具体来说，当一个I/O事件到达时，调用netpollready函数向netpollNote链表中注册事件，通知等待该事件的goroutine。如果goroutine当前正在阻塞等待I/O事件，则它会从阻塞状态返回并开始处理事件。

netpollNote的值为0时表示链表为空。当向链表中注册一个I/O事件时，调用notewakeup函数会使等待该事件的goroutine从睡眠状态中恢复并开始处理事件，同时它会将自己从链表中移除，以便下一个goroutine可以被通知。当所有goroutine都没有等待I/O事件时，netpollNote的值会被重置为0，以便轮询器等待下一个I/O事件。

总之，netpollNote是用于在goroutine间传递网络轮询通知的变量，它使得goroutine可以更有效地等待I/O事件并及时响应它们。



### netpollBrokenLock

在Go语言的runtime包中，netpoll_stub.go文件是一个实现了网络轮询的代码。其中netpollBrokenLock变量的作用是用于异步网络轮询器的同步。当异步网络轮询发生错误或被关闭时，该变量会被设置为一个非零值，用于通知主线程关闭网络轮询器，并释放资源。

具体来说，该变量被定义为一个int类型的全局变量，它用于记录轮询器的状态。当轮询器遇到错误或者被关闭时，该变量会被设置为一个非零的值，表示轮询器需要被关闭。

在主线程中，当发现该变量被设置为非零值时，会关闭异步网络轮询器，并释放相关资源。这样做可以保证网络轮询器的安全退出，同时避免出现死锁或者资源泄漏等问题。

总之，netpollBrokenLock变量的作用是用于异步网络轮询器的同步，可以保证网络轮询器的安全退出和资源释放。



### netpollBroken

netpoll_stub.go这个文件是在GO语言的runtime包中，它包含了一些与网络轮询相关的代码。其中，netpollBroken是一个布尔型变量，用来表示网络轮询是否已经被破坏。

在GO语言中，在进行网络轮询的时候，会有一些IO操作会被阻塞，这些被阻塞的IO操作需要被注册到网络轮询器中，以便轮询器在IO操作可以进行的时候唤醒它们。如果网络轮询器本身出现问题，那么这些被阻塞的IO操作就有可能得不到正确的处理，导致整个程序出现故障。

因此，netpollBroken这个变量的作用就是用来表示网络轮询器是否已经出现了问题。如果网络轮询器已经出现了问题，那么在进行网络轮询的时候，就需要做出一些特殊处理，确保被阻塞的IO操作能够得到正确的处理，从而避免程序出现故障。



## Functions:

### netpollGenericInit

netpollGenericInit是一个初始化网络轮询器的函数，它的作用是在系统不支持IO复用机制时，使用非阻塞IO来模拟IO复用机制。具体来说，当系统调用epoll_create或kqueue时失败，就会调用这个函数。

该函数首先创建了一组读写事件的文件描述符，并使用fcntl设置为非阻塞模式。然后，它返回两个文件描述符的数组，第一个是读文件描述符，第二个是写文件描述符。之后，这些文件描述符将用于调用I/O操作以进行网络轮询。

在netpollGenericInit函数中，还会通过调用runtime.SetNonblock将所有文件描述符设置为非阻塞模式。这样做的目的是为了在进行I/O操作时不会阻塞整个进程。

需要注意的是，netpollGenericInit函数仅在某些Unix系统不支持epoll或kqueue时调用。在大多数情况下，都会使用操作系统提供的IO复用机制来实现网络轮询。



### netpollBreak

netpoll_stub.go这个文件是Go语言运行时包中的一个文件，它提供了一些用于网络轮询的桩函数，这些函数在操作系统上直接实现，但是在一些虚拟化或仿真环境中，可能需要进行适当的修改。

netpollBreak这个func主要作用是触发网络轮询器的中断，使其退出永久循环并返回到调用方。这个函数通常在某些情况下用于中断阻塞的网络轮询操作，比如在程序关闭时强制退出轮询器，或者在网络异常时及时中断轮询操作并进行相应的错误处理。

具体来说，该函数会向netpollBreaker管道中写入一个空字节，以触发管道上的读取操作，而读取操作会返回一个错误信号，通知网络轮询器退出永久循环。在网络轮询器中，该函数通常由runtime_pollWait函数调用，用于监测网络IO的事件并阻塞等待，以处理TCP/UDP连接的读写操作。



### netpoll

在Go语言中，netpoll_stub.go文件中的netpoll函数是模拟网络轮询的一个框架。它主要的作用是向内核注册关注的文件描述符，并将其添加到Epoll对象中，随后使用epoll_wait机制等待就绪事件的发生。

具体来说，netpoll函数首先会创建一个PollDesc对象，该对象中包含了一个文件描述符、一个Epoll对象、以及一些控制参数。接着，它会通过gopark函数进行阻塞，直到能够收到来自Epoll对象的待处理事件。

当待处理事件到来时，netpoll函数会将这些事件挨个分发给具体的处理程序。一般来说，这些处理程序会根据传入的事件类型，执行不同的操作，例如将事件添加到某个队列、将就绪的事件进行处理等。

总体来说，netpoll函数实现了网络轮询和事件处理的核心机制，让Go语言在网络编程方面表现出色。它不仅可以提高网络应用的效率，还能够处理复杂的网络通信情况，为高级网络编程提供了有力的支持。



### netpollinited

netpoll_stub.go文件中的netpollinited函数设计的目的是初始化网络轮询器。

在Go语言的网络模型中，网络轮询器是负责处理网络IO事件的一个子系统，它会在程序启动时被初始化，然后轮询网络IO事件并将事件分发到各个协程中去执行。在netpoll_stub.go中，通过调用netpollinited来初始化网络轮询器，核心代码如下：

func netpollinited() {
    initNetpoll()
}

在initNetpoll()中，会调用systemstack()函数创建一个新的 goroutine，并在其中启动 epoll 轮询器，同时设置全局变量netpollInited为true，表示网络轮询器已经被初始化完成，可以使用：

func initNetpoll() {
    if epollinited() {
        throw("initNetpoll on epoll platform not supported")
    }

    // 创建 epoll 实例
    fd, err := epollCreate1(epollFlags)
    if err != nil {
        throw("epollcreate1 failed")
    }
    netpollfd = fd

    // 构建 runtime.pollDesc 数组，这个数组用于记录 fd 与 netpollDesc 的关系
    for i := range pollcache {
        heapBitsForAddr(pollcache[i], &mheap).initSpan()
    }

    runtime_pollServerInit()

    // 启动 epoll 轮询器
    go netpoll(epollWaitFD)
    netpollInited = true
}

总之，netpollinited函数的作用就是初始化网络轮询器。这是Go语言协程模型中非常重要的一个操作，因为网络轮询器的高效性能可以大大提高Go语言网络程序的性能和并发能力。



