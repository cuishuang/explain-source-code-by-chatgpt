# File: netpoll.go

netpoll.go是Golang运行时库中的一个文件，它的作用是实现网络轮询（network polling）。

网络轮询是一种通过多路复用的方式同时管理多个网络连接的技术。在Golang中，我们可以使用select语句实现网络轮询，但是select语句仅适用于文件描述符或通道，而无法直接使用socket句柄。因此，需要实现一种网络轮询的机制来管理socket连接。

netpoll.go文件中定义了netpollDesc类型，并提供了一些方法，用于将socket文件描述符与本地端点（local endpoint）或远程端点（remote endpoint）相关联，管理网络连接的读写事件，并在goroutine之间传播错误。

具体来说，netpoll.go文件中定义了以下几个函数：

1. netpollinit()

该函数会初始化网络轮询机制，并为每个CPU核心创建一个pollDescriptor对象。

2. netpollOpenDescriptor(fd uintptr) (*pollDescriptor, error)

该函数会用socket文件描述符fd创建并返回一个新的pollDescriptor对象。

3. (pd *pollDescriptor) evict()

该方法会从pollDescriptor数组中移除与其相关联的网络连接，并关闭这些连接。

4. (pd *pollDescriptor) wait(netpollWakeReason, seq uintptr) int

该方法会监听网络连接上的读写事件，并返回事件数量。如果发生错误，则向goroutine中传播错误。

5. (pd *pollDescriptor) resume(netpollWakeReason, seq uintptr)

该方法用于重新恢复网络连接上的读写事件监听。

通过这些函数和方法，netpoll.go文件实现了一个高效的网络轮询机制，可以用于管理多个并发的socket连接，提高网络通信的效率和吞吐量。




---

### Var:

### netpollInitLock

netpollInitLock是一个用来保护netpoll初始化过程的互斥锁。在runtime/netpoll.go文件中，有一个init函数，在程序启动时会被调用。该函数主要的作用就是初始化netpoll。

在init函数中，会首先尝试获取netpollInitLock锁，如果获取成功，就进入到初始化过程中。如果获取失败，则代表已经有其他协程正在初始化netpoll，当前协程需要等待直到锁被释放才能继续执行。

这个锁的作用是确保在初始化过程中只有一个协程在进行，避免并发导致的竞争问题。netpollInitLock保护了netpoll的初始化流程，因为在初始化期间，所有的goroutine都必须能够正常使用netpoll。如果在初始化期间出现了竞争条件，就可能会导致死锁或其他类似问题的发生。

总之，netpollInitLock起到了控制并发初始化的作用，保证了程序正常运行。



### netpollInited

netpollInited是一个布尔变量，用于标记网络轮询器是否已经初始化。

在Go语言中，每个操作系统线程都有一个网络轮询器（netpoller），该轮询器负责监视该线程上所注册的网络事件（如套接字可读/可写事件）和出现的其他事件（如计时器事件），并在发生事件时通知相应的协程。

当一个新的线程被创建时，该线程的netpoller需要被初始化。而在初始化之前，我们需要保证其他的goroutine不会去尝试使用该线程的netpoller。因此，当netpollInited为false时，其它goroutine会陷入轮询状态，直到该变量变为true。

具体来说，当一个新线程被创建时，会先调用procPin函数将其绑定到某个P（逻辑处理器）上，并在这个P的下一个周期调用procUnpin函数的时候执行网络轮询器的初始化。

网络轮询器的初始化包括以下几个步骤：

1. 创建一个epoll实例，并调用epollctl将该线程绑定到该epoll实例上，这样所有的网络事件都会交给该线程处理。

2. 初始化pollSize，并给netpolldescs分配一定空间。

3. 调用netpollOpen函数打开/proc/sys/net/core/somaxconn文件，并读取其内容，用来给somaxconn字段赋值。

4. 将netpollInited变量置为true，使得其他goroutine可以使用该线程上的网络轮询器。

综上所述，netpollInited变量用于保证网络轮询器的安全初始化，在网络轮询器初始化之前，其他goroutine会一直等待，直到初始化完成。



### pollcache

在go/src/runtime/netpoll.go文件中，pollcache是一个内部缓存，用于缓存网络轮询结果，避免在需要重用网络轮询结果时频繁调用操作系统的网络轮询接口，以提高网络I/O性能。

具体来说，pollcache是一个数组(slice)类型，每个元素都是一个结构体类型，包含如下成员：

- fd: 文件描述符
- mode: 文件描述符的模式，一般为“读”或“写”
- pv: 网络轮询结果缓存，存储操作系统对该文件描述符的网络轮询结果

当执行网络I/O操作时，代码会首先尝试在pollcache中查找是否已经缓存了该文件描述符的网络轮询结果。如果存在，则直接使用缓存中的结果进行网络I/O操作，避免频繁调用操作系统的轮询接口。

当然，由于pollcache是有限的，当缓存已满或者不存在所需的网络轮询结果时，代码会重新调用操作系统的网络轮询接口，并将结果缓存到pollcache中以备后续使用。

总之，pollcache作为一个内部缓存，可以提高网络I/O性能，减少频繁调用操作系统的网络轮询接口带来的开销。



### netpollWaiters

netpollWaiters是一个SyncMap类型的变量，它的作用是用于存储网络轮询器(netpoll)的等待列表。

在Go程序中，每一个非阻塞的网络IO操作都会通过netpoll来进行调度和观察。当某个IO操作无法立即完成时，轮询器将会把该操作对应的goroutine挂起，直到有数据可读或可写时再唤醒它。

如果有多个goroutine在等待轮询器的IO操作结果，那么它们将会被存储在netpollWaiters中。当IO操作完成后，轮询器会遍历netpollWaiters，将每一个等待中的goroutine都唤醒。

由于netpollWaiters的类型为SyncMap，它支持并发访问。因此，多个goroutine可以同时等待同一个IO操作的结果，安全地存储在netpollWaiters中。同时，由于SyncMap的高效性能，网格员可以快速地在等待列表中查找和唤醒goroutine。

总之，netpollWaiters变量在网络I/O操作的等待中发挥着重要作用。它实际上是Go程序实现非阻塞IO的核心数据结构之一，为网络轮询器提供了高效的等待和唤醒机制。



### pdEface

在go/src/runtime/netpoll.go文件中，pdEface变量是一个实现了PollDesc接口的Eface类型。该接口描述了一个fd（文件描述符）与其对应读写事件以及发生的回调函数的关系。pdEface变量的作用是保持PollDesc接口实现的具体类型，以便在实现网络I/O复用时使用。通过将pdEface变量作为参数传递给相关函数，可以避免需要在多个地方重复定义PollDesc类型的实例。由于pdEface变量是一个接口类型，因此可以轻松地使用在不同的实现中使用它来满足不同的需求。



### pdType

pdType是一个枚举类型的变量，表示pollDesc的类型。在Go语言中，网络和文件IO都使用了非阻塞的方式，因此需要使用一个轮询模型，由操作系统通知程序有哪些文件描述符可以读或写，这个模型称为poller。在Go语言中，poller被封装在netpoll.go这个文件中。

pollDesc是一个轮询描述符，其内部包含一个关联的文件描述符和事件等待队列。poller使用pollDesc来表示等待的事件，并在事件发生时通知程序。在netpoll.go文件中，存在两种类型的pollDesc：netpollDesc和fdMutex。

pdType变量用于区分不同类型的pollDesc。目前，pdType变量包含以下两个值：

1. PdNetpoll：表示netpollDesc类型的pollDesc。
2. PdFdmutex：表示fdMutex类型的pollDesc。

通过pdType变量的值，可以清晰地了解一个pollDesc是哪种类型的。这在处理具体的网络事件或文件事件时非常有用。






---

### Structs:

### pollDesc

在go语言的运行时(runtime)中，netpoll.go文件中的pollDesc结构体起到了重要作用。该结构体主要用于表示I/O事件的状态和相关处理方法。在网络编程中，需要不断地监听socket文件描述符上的I/O事件（例如可读事件、可写事件等），同时根据不同的I/O事件类型，进行不同的处理操作。pollDesc结构体就提供了这样的一种表示I/O状态的机制。

pollDesc结构体中包含以下几个重要字段：

1. fd：表示文件描述符。

2. closing：表示是否正在进行关闭操作。

3. ev：表示I/O事件类型（例如可读事件、可写事件等）。

4. rtick：记录上一次检查该事件后的时间，用于定时器相关操作。

5. wd：表示等待的写事件。

6. rd：表示等待的读事件。

pollDesc结构体中还包含了一些方法，用于对I/O事件进行处理。例如，该结构体中的read和write方法分别用于表示读取和写入操作。同时，该结构体还提供了wait方法，用于等待I/O事件发生，并且根据事件类型，调用不同的处理方法。

总之，pollDesc结构体是go语言运行时中的一个重要机制，用于表示I/O事件的状态和相关处理方法。它可以帮助开发者实现高效的网络编程，并且提供了灵活的事件处理机制。



### pollInfo

pollInfo结构体是用于存储网络I/O多路复用的相关信息的。具体来说，它包含了以下几个字段：

- runtimeCtx: 与g批次相关联的运行时环境上下文。
- idlePoll阈值：阈值是I/O检查通道（”/Produce/-10“）的通知器 goroutine 的空闲时间。如果无通知，则监听器会面临睡眠。注意，如果在不操作文件描述符的情况下重新计时闹钟，则应该确认计时器仍然在运行。
- mode：网络I/O多路复用的模式，通常是阻塞或非阻塞模式。
- pd：与描述符（其文件描述符或I/OCPOLT与端口关联）相关的pollDesc结构。
- index：pollDescslice中的索引，用于反向查找。
- rtDesc：与描述符相关联的运行时描述符；在syscall.DescriporType时变为非空。

pollInfo结构体的作用是提供一种方式，使得可以轻松地管理和操作多种网络I/O多路复用的实现。在Go语言中，使用网络I/O多路复用来同时处理多个连接的输入和输出通常是一个非常常见的任务。通过使用pollInfo，开发人员可以更方便地访问这些通知器，同时也可以更好地管理和调试网络I/O多路复用代码。



### pollCache

pollCache是一个用于缓存系统发起的网络poll操作的结构体。在高并发的网络编程中，系统经常需要在多个socket文件描述符上进行poll操作，这会消耗大量的系统资源。为了减小系统资源的消耗，runtime/netpoll.go中维护了一个pollCache结构体。

pollCache结构体中有三个字段：
1. data：一个指向uintptr类型数组的指针，用于存储pollfd结构体
2. size：data指向的数组的大小
3. lastget：最后一次从data数组中取出的pollfd结构体的序号

pollCache结构体的作用是，在系统执行网络poll操作时，根据需要创建或扩展一个pollfd数组，并将该数组存储在pollCache中。这样，在下一次进行poll操作时，就可以直接从pollCache中取出pollfd数组，而无需重复申请内存和初始化。

pollCache结构体还有一个优化，在使用完pollfd数组后，并不需要立即释放。因为在高并发的情况下，很可能在下一次poll操作时会重新使用这个pollfd数组。所以，将pollfd数组保留在pollCache中，可以提高下一次poll操作的效率。

综上所述，pollCache结构体的作用是缓存系统发起的网络poll操作，减小系统资源的消耗，并提高下一次poll操作的效率。



## Functions:

### closing

在 Go 中，每个 Goroutine 都有自己的栈和寄存器。当 Goroutine 被创建时，它绑定到调度器并等待分配处理器。调度器通过将 Goroutine 线程放置在线程池中来执行它。因此，Goroutine 与底层操作系统线程之间存在一对多关系。

在运行时环境中，netpoll.go 文件中的关闭函数负责停止 Goroutine 并确保与此 Goroutine 相关的所有系统资源都得到释放。closing 函数的方法是通过将 Goroutine 从处理器集中删除，然后将其标记为关闭状态，该操作会防止 Goroutine 接收新任务，并等待当前任务完成。在此方法中，Goroutine 在关闭之前必须完成当前正在处理的任何请求，继续运行将导致死锁和其他不良的行为。

此外，closing 函数还执行关闭 Goroutine 的一些清理操作，例如释放可能持有的锁，关闭可能打开的文件描述符，取消 Goroutine 可能在等待的任何同步操作等。在所有这些清理操作完成之后，Goroutine 允许被垃圾回收并释放它的所有资源，例如对其分配的堆内存。



### eventErr

在 Go 语言中，`netpoll.go` 文件中的 `eventErr()` 函数用于处理套接字操作中的错误事件，并根据错误类型采取相应的处理方式。

该函数首先将指定的 `netFD` 结构体的 `lock` 字段加锁，以确保在 `eventfd` 函数中可能会访问到同一个 `netFD` 时不会出现问题。

接下来，该函数会通过判断传入的错误参数 `ev` 的类型，采取不同的处理方式：

- 如果 `ev` 是 `epoll.ErrPollHangup`（表明轮询被中止），则函数会调用 `fdclose()` 函数关闭对应 `netFD`。
- 如果 `ev` 是 `syscall.EPIPE`（表明管道破裂），则函数会关闭与该 `netFD` 关联的套接字。
- 如果 `ev` 是 `syscall.ECONNRESET` 或 `syscall.ENETRESET`（表明连接被重置或网络被重置），则函数会将该 `netFD` 的 `readystate` 字段设置为 `_RendezvousBroken`。
- 如果 `ev` 是其他一些错误类型，则函数会将该 `netFD` 的 `readystate` 字段设置为 `_RendezvousFailed`。

无论哪种情况，该函数都会将当前 `netFD` 的 `wrclosed` 和 `rdclosed` 标志位设置为 `true`，并使阻塞等待与该 `netFD` 关联的协程重新唤醒，以告知它们该套接字已经关闭。最后，该函数会释放该 `netFD` 的 `lock` 字段，并返回该 `netFD` 的状态。



### expiredReadDeadline

在go/src/runtime/netpoll.go文件中，expiredReadDeadline()函数的作用是计算已过期的读取截止日期。

在Go中，网络I/O操作通常使用非阻塞I/O，这意味着我们需要使用轮询或选择来检查I/O操作的状态。Netpoll就是一个这样的工具，它使用内核级事件通知来监视I/O操作的状态。

expiredReadDeadline()函数用于跟踪所有等待读取的goroutine，并检查它们的读取截止日期是否已过期。通过计算goroutine的读取截止日期与当前时间的差异，如果差异小于等于0，则表示读取截止日期已经过期，这意味着该goroutine需要被唤醒。

需要注意的是，expiredReadDeadline()函数只是负责检查读取截止日期是否已过期，而不是实际的唤醒操作。实际的唤醒操作是在netpoll函数中完成的，该函数使用IO复用技术（如select或epoll）来监听网络事件，并处理已过期的读取截止日期。

因此，expiredReadDeadline()函数是Netpoll的一个重要组成部分，通过跟踪所有等待读取的goroutine，确保I/O操作的高效运行。



### expiredWriteDeadline

在go语言中，goroutine和网络IO都是由操作系统内核来调度的。在发生网络IO的情况下，往往需要等待对方的响应，这个等待的过程需要占用CPU资源，如果等待时间过长，就会造成浪费。所以，在网络IO中，一般都会设置一个超时时间，如果等待时间超过了这个时间，就会被视为失败，释放占用的资源，以便后续处理。

expiredWriteDeadline这个函数是Netpoll包中的一个处理超时的函数，主要是用来检查网络写的超时情况。在网络写的时候，如果写入数据、刷新缓冲区等待对方响应的时间超过了设置的超时时间，就会发生超时，expiredWriteDeadline函数会将对应的网络文件描述符从写缓存中删除，并返回一个错误信息，以便提醒程序进行异常处理。

expiredWriteDeadline是在网络文件描述符注册时，设置超时时间的函数之一，在网络文件描述符调度时，Netpoll包会定期调用这个函数来检查是否发生超时，以便及时释放资源，避免浪费，保证系统的正常运行。



### info

`info`函数在`netpoll.go`文件中定义，主要用于打印网络轮询器的相关信息，包括轮询器正在监听的文件描述符的数量、当前等待的`goroutine`的数量、当前`goroutine`的状态等。该函数被用于调试和分析网络性能问题。

具体来说，`info`函数会打印以下信息：

- 当前网络轮询器正在监听的文件描述符的数量。
- 当前等待网络事件的`goroutine`的数量。
- 当前被阻塞在网络IO操作上的`goroutine`的数量。
- 当前处理网络IO事件的`goroutine`的数量。
- 当前正在执行与网络IO操作相关的`syscall`的`goroutine`的数量。
- 当前正在创建或销毁网络连接的`goroutine`的数量。

通过打印这些信息，可以帮助开发人员更好地理解和调试网络模型的性能和行为。在使用`Go`开发网络应用程序时，使用`info`函数进行调试和性能评估是一种非常有用的手段。



### publishInfo

在Go的运行时环境中，network poller（网络轮询器）是负责管理网络事件并将其通知给Go程序的线程的组件。 publishInfo()是一个函数，它提供了一种将事件通知给网络轮询器的机制。

当Go程序中产生一个新的网络事件时，例如一个新的连接建立或者一个已经存在的连接变成可读/可写，它需要通知网络轮询器来处理这个事件。publishInfo()函数为网络轮询器提供了通知机制，当一个事件需要被通知时，它会将事件的相关信息打包到一个结构体中，并通过发布者-订阅者模式向网络轮询器注册事件。

这个函数对于Go程序的网络性能和可靠性非常重要，因为它提供了一种高效的机制来提醒网络轮询器处理事件，从而使程序能够充分地利用网络资源，并保持与其他程序和系统的连接的稳定性。



### setEventErr

setEventErr函数是在网络轮询器中的事件处理函数之一，主要的作用是将网络错误设置为事件中的错误码，以便上层函数处理该错误。

在网络轮询器中，当一个网络事件发生时，例如一个新的连接被建立，或者一个连接上有数据可读或可写，就会调用对应的事件处理函数来处理该事件。setEventErr函数就是在处理连接上的数据读写事件时被调用的函数之一。

具体来说，当网络事件处理函数向系统注册了某个连接的读写事件，如果此时网络出现了异常，例如连接被重置、连接被关闭等，那么就需要将该事件设置为一个错误事件，并返回给上层函数处理。setEventErr函数就是负责把该事件设置为一个错误事件的函数。

setEventErr函数的参数有三个，分别是epollevent、errno和closing。其中epollevent表示当前正在处理的网络事件，errno表示网络出现异常的错误码，closing表示当前连接是否正在关闭。setEventErr函数的主要作用是根据这些参数，将事件设置为一个错误事件，并返回给上层函数处理。

总之，setEventErr函数在网络轮询器中是非常重要的一个函数，它的主要作用是将网络错误设置为事件中的错误码，以便上层函数处理该错误。



### poll_runtime_pollServerInit

poll_runtime_pollServerInit函数是用来初始化网络轮询服务器的。这个函数会创建一个锁和一组epoll事件，用于处理各种网络事件。该函数还会创建一个goroutine来处理触发的事件。

具体来说，poll_runtime_pollServerInit函数有以下作用：

1. 创建一个互斥锁，用于保护内部状态。

2. 创建一个用于监听文件描述符的epoll对象，加入epoll事件。

3. 创建一个新的goroutine，该goroutine将负责监视epoll事件。

4. 为处理网络IO事件提供一个通用的通道。

5. 初始化所有必须的参数，如允许的最大连接数、读写超时等。

总的来说，poll_runtime_pollServerInit函数是为上层应用程序提供底层网络支持的重要组件之一。它提供了一种高效的方式来处理网络IO事件，从而使应用程序能够快速响应用户请求。



### netpollGenericInit

netpollGenericInit是Go语言运行时（runtime）中的一个函数，其作用是初始化网络轮询器（netpoller）并启动一个goroutine来处理网络事件。

在Go语言中，网络I/O通常是通过goroutine和网络轮询器实现的。具体来说，当我们需要对一个网络连接进行读取或写入操作时，我们不会直接进行I/O操作，而是会创建一个goroutine来负责该操作，并将goroutine和相应的网络描述符（文件句柄）注册到网络轮询器中，等待网络事件的触发。

netpollGenericInit就是负责初始化这个网络轮询器的函数。在该函数中，首先会创建一个epoll描述符（Linux系统中的一种高效的I/O事件通知机制），并对该描述符进行设置，以便能够处理网络事件。然后，会创建一个goroutine来负责网络事件的处理，具体来说，会调用runtime·netpoll函数对网络轮询器中的事件进行处理。最后，会将该网络轮询器的相关信息保存到全局变量中，以供后续使用。

需要注意的是，netpollGenericInit是一个平台无关的函数，其实现并不依赖于具体的操作系统或硬件平台。在实际使用中，不同平台的网络轮询器可能会有所不同，因此Go语言运行时会根据不同的平台选择相应的网络轮询器来进行初始化，具体的实现可以参考其他相关的文件。



### netpollinited

netpollinited是一个用于初始化网络轮询器的函数。在Go语言程序中，网络轮询器是用于判断网络事件是否发生的一个核心组件。通过使用网络轮询器，Go语言程序可以同时监听多个网络连接，实现高并发的网络通信。

具体来说，netpollinited主要完成以下两个任务：

1. 初始化网络轮询器的数据结构

在netpollinited函数中，会调用goepollinit函数初始化网络轮询器的数据结构。其中，网络轮询器的数据结构具体是由操作系统提供的。在Linux系统中，网络轮询器使用epoll实现。在Windows系统中，网络轮询器使用I/O completion ports实现。goepollinit函数会根据操作系统的不同，调用相应的系统接口进行网络轮询器的初始化。

2. 启动网络轮询器的goroutine

除了初始化网络轮询器的数据结构外，netpollinited函数还会启动一个goroutine，用于监听网络事件。具体来说，该goroutine会在一个无限循环中调用goepollwait函数，等待网络事件发生。一旦网络事件发生，goroutine就会调用netpollready函数，处理对应的网络事件。

总的来说，netpollinited函数是Go语言网络轮询器的初始化函数，它会完成网络轮询器的数据结构的初始化，并启动一个goroutine，用于监听网络事件。



### poll_runtime_isPollServerDescriptor

函数名称：poll_runtime_isPollServerDescriptor

函数功能：该函数用于判断当前文件描述符是否属于网络轮询器（netpoll）管理。

函数实现：

该函数的具体实现如下：

```
// 判断fd是否属于网络轮询
func poll_runtime_isPollServerDescriptor(fd uintptr) bool {
    // 判断该fd与pollDesc数组中的pollfd是否一致
    for _, pd := range netpoll.pollDesc {
        if pd.runtimeCtx != 0 && pd.runtimeCtx == pd.readable && fd == pd.fd {
            return true
        }
        if pd.runtimeCtx != 0 && pd.runtimeCtx == pd.writable && fd == pd.fd {
            return true
        }
    }
    return false
}
```

该函数先遍历了netpoll中的pollDesc数组，判断该fd是否属于网络轮询器管理的文件描述符范围。具体来说，函数判断文件描述符fd和每一个pollDesc的读写描述符是否相等，如果相等的话，就说明该文件描述符是由网络轮询器管理的。

如果符合条件的pollfd是多个，或者数据队列中存在多个pollfd（即读写描述符一致，但runtimeCtx不一致），那么函数只返回第一个符合条件的pollfd所在pollDesc的信息。

如果找不到符合条件的pollDesc，那么该函数返回false，说明该fd不是轮询器管理的文件描述符。

使用场景：

这个函数在socket编程中很常见。当我们需要进行异步IO或者事件驱动的时候，我们需要利用select/poll/epoll等机制来实现非阻塞IO。而这些机制都是基于文件描述符的，所以我们需要将文件描述符加入到网络轮询器管理中。在加入或者删除文件描述符的时候，我们都需要使用到该函数来判断当前文件描述符是否已经被轮询器管理。



### poll_runtime_pollOpen

poll_runtime_pollOpen这个函数是用来打开一个新的文件描述符的，该文件描述符通常用于网络轮询事件。该函数的实现是通过在系统上调用epoll_create1()函数来创建一个epoll文件描述符，该文件描述符用于在运行时执行异步网络轮询操作。该函数一般会在程序初始化时调用一次，以初始化网络轮询事件的相关资源，并将其存储在runtime实例的poll字段中。

在具体实现中，该函数先检查运行时的poll字段是否已经初始化，如果已经初始化，则立即返回；否则，创建一个新的pollDesc变量，并将其赋值给运行时的poll字段。同时将其操作系统文件描述符设置为新创建的文件描述符，并将I/O兴趣设置为pollNone（即不轮询任何事件）。最后，该函数调用netpoll_init将poll字段交给网络轮询器初始化，从而为程序后续的网络轮询事件准备好资源。



### poll_runtime_pollClose

在Go语言中，网络轮询器（netpoll）是一个在操作系统的I/O复用机制之上构建的轻量级库，它使用内部的异步I/O和goroutine来实现高效的网络I/O操作。

在go/src/runtime/netpoll.go中的poll_runtime_pollClose函数是用来关闭一个文件描述符对应的网络轮询器的。当一个网络轮询器不再需要使用时，它应该被关闭以释放资源。

具体来说，poll_runtime_pollClose函数中会执行以下操作：

1. 通过runtime_pollDescFromGo获取与文件描述符关联的runtime_pollDesc结构体。runtime_pollDesc结构体是netpoll的核心数据结构，它包含了轮询器中文件描述符的相关信息，例如文件类型、队列等。

2. 通过调用netpollCloseDesc停止与该文件描述符关联的I/O操作，并从netpoll中删除该描述符。

3. 释放runtime_pollDesc结构体中的相关内存资源，包括文件描述符关联的文件状态、事件信号等。

总的来说，poll_runtime_pollClose函数是用来安全地关闭网络轮询器和它所关联的文件描述符，在操作系统层面上释放旧资源和占用内存。同时，通过netpollCloseDesc函数，它也确保了所有与文件描述符相关的I/O操作都被安全地停止。



### free

netpoll.go文件是Go语言运行时包中的一个文件，定义了网络轮询器（也称为I/O多路复用器）的实现。在这个文件中，free函数的作用是释放netpollDesc数组中的一个元素占用的内存，以便在未来重新使用。

具体地说，当一个网络轮询器不再需要与一个文件描述符关联时，调用这个函数可以将对应的netpollDesc元素标记为“未使用”，并从netpollDescFree链表中取下一个可用的元素并返回。由于netpollDesc的数量是有限的，所以通过重用这些元素可以节省内存。

这个函数的定义如下：

func (p *pollDesc) free() {
    if ptr := p.netpollDesc; ptr != nil {
        p.netpollDesc = nil
        netpollDescFree(ptr)
    }
}

其中，pollDesc是与一个文件描述符相关联的结构体，它包含了一个指向netpollDesc的指针。free方法会先检查这个指针是否为nil，如果不是，则将其置为nil，并调用netpollDescFree函数释放相关内存。netpollDescFree函数的定义如下：

func netpollDescFree(pd *pollDesc) {
    pd.setDeadline(nil, 0)
    if atomic.Loadint32(&pd.closing) == 0 {
        pd.fd.closeOnExec()
        pd.fd = nil
    }
    pd.rsema = 0
    pd.wsema = 0

    // Add the descriptor to the freelist, not the garbage collector.
    // See the comment in netpoll.go.
    for {
        old := netpollDescFreeList.Load()
        pd.next = old
        if atomic.CompareAndSwapPointer(&netpollDescFreeList, old, unsafe.Pointer(pd)) {
            return
        }
    }
}

这个函数会将传入的netpollDesc元素的成员变量都清空，并将其加入到一个链表中。这个链表中的元素都是已经被标记为“未使用”的netpollDesc元素，可以在未来重新使用。需要注意的是，这个链表并不交由垃圾回收器（GC）管理，因为这些元素虽然没有被使用，但由于它们的成员变量中仍然包含有指向其他对象的指针，所以它们仍然需要被垃圾回收器扫描。为了避免这种情况，这些元素被单独管理，并不交由GC管理。



### poll_runtime_pollReset

poll_runtime_pollReset函数的作用是重置一个描述符的事件状态位。在Go的网络库中，每个socket都需要注册到网络事件轮询器（往往是epoll或者kqueue）中，以便在网络事件（读/写等）发生时及时处理。在这个过程中，轮询器会为每个socket维护一个状态位，表示该socket关注的事件类型，这个状态位在每次poll调用时都会被检查，并决定是否触发相应的网络事件回调。

poll_runtime_pollReset函数在每次poll调用之前被调用，用于重置socket的事件状态位。具体来说，它会将socket的状态位重新设置为0（即不关注任何事件），然后重新注册到轮询器中，重新指定关注的事件类型。这个过程确保在下一次调用poll时，轮询器可以正确的检查socket的状态位，并且在合适的时候触发相应的网络事件回调。



### poll_runtime_pollWait

poll_runtime_pollWait是一个用于等待IO事件的系统调用，它实际上直接调用了系统的epoll_wait或者kqueue等函数来完成等待IO事件的操作。它是Go语言中网络轮询模型的关键函数之一。

当一个goroutine需要等待IO事件时，它会通过netpoller对象将自己的相关信息（如文件描述符、事件类型等）注册到内核的IO复用机制中。当有IO事件发生时，内核会通知netpoller对象，并调用poll_runtime_pollWait来等待IO事件。

在poll_runtime_pollWait函数中，它会先检查goroutine是否已经被唤醒（在等待IO事件之前，可能已经有其他goroutine唤醒了该goroutine），如果已经唤醒，则直接返回，否则调用epoll_wait或者kqueue等函数来等待IO事件的发生。如果等待过程中被唤醒，则停止等待，如果等待超时或者出现错误，则返回相应的错误码。

该函数的显著特点是阻塞等待，当无法进入select阻塞才会起作用，它具有低延迟、高效、可扩展等优点，因此在Go语言中被广泛使用。



### poll_runtime_pollWaitCanceled

`poll_runtime_pollWaitCanceled`是 Go 语言运行时调度器中实现的一个函数，它的作用是在等待系统 I/O 事件发生时，阻塞 Go 协程的执行，并在有 I/O 事件发生或超时时解除阻塞，从而保证协程可以及时运行。

具体来说，当 Go 协程需要执行系统 I/O 操作时，runtime 会调用 `netpoll.pollDesc.rearm` 触发系统 epoll 或 kqueue 等 I/O 复用机制，在 `poll_runtime_pollWaitCanceled` 中会调用 `netpoll.pollDesc.wait` 方法等待 I/O 事件的发生，同时，利用了一个已经取消的 channel，确保可以在等待结束之前及时取消等待。

当有 I/O 事件发生时，`poll_runtime_pollWaitCanceled` 会调用 `netpoll.pollDesc.notifyReady` 方法通知 runtime，唤醒对应的协程继续执行；当超时时，也会调用 `netpoll.pollDesc.notifyReady` 方法通知 runtime，但此时唤醒的是一个通知已超时的 channel。

总之，`poll_runtime_pollWaitCanceled` 是 Go 并发模型中实现 I/O 复用机制非常重要的一部分，它能够高效地处理 I/O 事件和协程的阻塞与唤醒。



### poll_runtime_pollSetDeadline

poll_runtime_pollSetDeadline函数的作用是将给定的网络文件描述符（fd）的读写超时时间设置为指定的时间戳（deadline）。

具体地说，该函数会将fd注册到一个特定的I/O复用器（IOCP或epoll）的时间轮中，使其在deadline到达之前被监视。一旦该fd的I/O事件变得有效（如可读或可写），复用器将该文件描述符从时间轮中删除并通知运行时系统。运行时系统接着调用goroutines和重复这个过程，直到fd的读写操作被完成或发生超时。

这个函数在Go语言的网络编程中非常重要，因为它能够自动地关注一系列的网络事件，并且能够在发生事件时快速地唤醒相应的goroutines。通过使用该函数，Go语言可以实现高效的并发IO模型，而不需要显式地调用操作系统提供的复杂的IO接口。



### poll_runtime_pollUnblock

函数poll_runtime_pollUnblock是用于唤醒被网络I/O复用器(netpoll)阻塞的goroutine的函数，它会将传入的参数`pd`所代表的文件描述符从epoll中删除，以从网络I/O复用器中解除该文件描述符所导致的阻塞状态。

具体来说，poll_runtime_pollUnblock函数的作用可以概括为以下三个方面：

1. 从I/O复用器中删除文件描述符：调用epoll_ctl将`pd`所代表的文件描述符从epoll中删除，从而使其不再参与I/O复用的过程，也因此，其在网络I/O复用器中的阻塞状态因此被解除。

2. 唤醒等待的goroutine：在执行step2之前，用于等待文件描述符的goroutine必定已经挂起，等待网络I/O事件到来而进入阻塞状态。当删除文件描述符后，这些goroutine则会被同时唤醒，并根据之前阻塞的类型执行不同的操作。

3. 处理唤醒的goroutine：唤醒后的goroutine会返回一个错误(ECONNRESET)，表明在阻塞期间，文件描述符已经被关闭。这时，netpoll会在goroutine所在的系统线程上抛出一个异常，将goroutine重新放回到队列中，等待被再次唤醒（可能是在新的文件描述符上）。

总的来说，poll_runtime_pollUnblock所做的就是一个解除阻塞状态的操作，它将使得之前被网络I/O复用器阻塞的goroutine重新进入到可调度状态，并在之后的调度中争取到执行的机会。



### netpollready

netpollready函数是Go语言运行时的网络轮询函数之一，主要负责在操作系统底层的网络异步I/O接口发生可读或可写事件时，触发Go语言程序响应这些事件，并执行相应的回调函数。

具体来说，当程序需要监控某个文件描述符（通常为套接字）的读写事件时，它会通过netpollinit函数将该文件描述符注册到操作系统的异步I/O接口上，然后通过netpoll函数开启一个轮询循环，在循环体内不断监听异步I/O接口的可读可写事件。

当异步I/O接口发现某个文件描述符发生可读或可写事件时，它会向Go语言程序发送一个通知，此时netpoll函数就会通过调用netpollready函数来处理这个事件。netpollready函数会首先根据文件描述符对应的IO对象类型（netFD或file）来选择相应的回调函数，然后调用它们来执行实际的读写操作或其他处理。最后，netpollready函数将处理结果返回给netpoll函数，并立即进入下一个轮询循环。

总之，netpollready函数是Go语言运行时网络I/O模块的关键组成部分，它通过异步I/O接口等底层机制，实现了高性能的、基于事件驱动的网络编程模型。



### netpollcheckerr

netpollcheckerr这个函数是一个用于处理网络I/O错误的函数。它的作用是检查网络I/O操作返回的错误，如果错误不是暂时错误（例如EAGAIN或EWOULDBLOCK）且不是可恢复错误（例如EINTR），则将其记录到日志中并关闭相关的文件描述符。

如果发生可恢复错误，则将错误标记为Intr，并在PollDesc的err字段中存储该错误。

如果发生暂时错误，则将错误标记为Temporary，并在PollDesc的err字段中存储该错误。

如果发生不可恢复错误，则将其记录到stderr，并将文件描述符从poller中删除。如果文件描述符在网络效能模型中，则将其还原为可读/可写状态。 

总之，netpollcheckerr函数确保网络I/O操作在遇到错误时能够正确处理错误并避免在错误时一直阻塞。这是一个非常重要的函数，对于保障网络通信的稳定性和可靠性有着重要的作用。



### netpollblockcommit

函数名称：netpollblockcommit()

所在文件：go/src/runtime/netpoll.go

作用：将netpollg数组中的goroutine存放到调度器中的运行队列中。该函数会在全部的网络I/O文件描述符中都没有事件时被调用，用于将网络I/O操作阻塞掉的goroutine暂时从运行队列中移除，并将其存放到netpollblock数组中，以便后续继续执行阻塞读/写操作。

函数具体实现：

1.获取netpollgoroot标记的goroutine，并将其从调度器中的运行队列中移除。

2.将运行队列中剩余的goroutine存放到netpollextra数组中。

3.将运行队列中所有goroutine移除并存放到netpollblock数组中。

4.将所有网络I/O文件描述符的读/写事件取消，并将netpollWakeSigCh和netpollWakeErrCh的读/写事件也取消。

5.释放netpollWait队列的协程，以便它们可以执行网络I/O操作。

6.将netpollblock中存放的所有goroutine重新加入到调度器的运行队列中。

7.将netpollextra中存放的goroutine一并加入到调度器的运行队列中。

8.重新启用所有网络I/O文件描述符的读/写事件。

总结：netpollblockcommit()函数的主要作用是将因网络I/O操作而被阻塞的goroutine暂时从运行队列中移除，并将其存放到netpollblock数组中，以便后续继续执行阻塞读/写操作，同时将运行队列中的其他goroutine先放到netpollextra数组中，并将所有网络I/O文件描述符的读/写事件取消，释放netpollWait队列的协程，重新加入到调度器的运行队列中，然后重新启用所有网络I/O文件描述符的读/写事件。



### netpollgoready

在Go语言中，网络处理的事件往往是异步的，即使是网络I/O在执行时也会把它交给系统内核来处理。Go语言通过使用系统级别的select函数来处理异步事件，并将网络I/O操作添加到执行队列中。

在runtime/netpoll.go文件中，netpollgoready函数是用来接收网络事件的。当该函数被调用时，它会将当前goroutine添加到网络处理的goroutine队列中。这些goroutine会轮流地处理网络I/O事件。当有网络事件发生时，会通过该队列中的goroutine来处理这些事件。

此外，此函数还会调用netpollBlock函数，该函数会阻塞goroutine直到有网络事件发生，然后唤醒goroutine并开始处理事件。通过这种方式，Go语言的网络处理可以实现高效的异步I/O处理。



### netpollblock

Netpollblock是Go语言中实现网络IO复用的核心函数之一。它是Go语言运行时包中的一个函数，用于在网络I/O发生时，将发生I/O事件的Go程中的任务挂起，然后将其移到其他Go程中执行，以充分利用真实的多核处理器。

具体来说，Netpollblock函数在调用时会阻塞当前的Go程，等待网络I/O事件的发生。当网络I/O事件发生时，Netpollblock函数会将正在等待I/O事件的所有Go程批量唤醒，然后将唤醒的Go程移动到Runnable队列中，由系统调度器选择其他可运行的Go程执行。一旦发生I/O事件的Go程所需的I/O操作完成，它将从Runnable队列中移除并再次成为可调度状态。

Netpollblock函数的另一个重要特点是它可以避免在多核处理器上出现并发时出现的饥饿问题，这是由于Netpollblock函数将已发生网络I/O事件的Go程均匀地分配到多个系统线程上进行执行。

总的来说，Netpollblock函数是Go语言实现高效网络I/O复用的核心，它在并行执行和高性能方面有着显著的优势，可以充分利用系统的处理能力，提高应用程序的性能和响应速度。



### netpollunblock

netpollunblock是Go语言运行时的一部分，位于runtime/netpoll.go文件中。主要功能是解除阻塞的网络I/O操作。

在网络编程中，当一个goroutine阻塞在等待网络I/O操作时，可以使用netpollunblock来解除阻塞状态。该函数会从runtime的网络轮询器中移除当前goroutine的阻塞事件，使其重新可调度。

具体来说，当一个goroutine阻塞在等待网络I/O时，会将其注册到网络轮询器中进行监控，等待网络读写事件的通知。然而，一些情况下需要手动解除该goroutine的阻塞状态，比如出现网络错误或超时。这时就可以使用netpollunblock来取消该事件的监视，并将该goroutine重新加入到调度列表中等待下一次调度。

总结来说，netpollunblock的主要作用是在Go语言的网络编程中，提供一种手动解除阻塞状态的方法，从而提高程序的容错性和可靠性。



### netpolldeadlineimpl

netpolldeadlineimpl是Go语言中runtime包中netpoll.go文件中的一个函数，它的作用是在IO复用模型中用于设置下一次poll的deadline时间。

在IO复用模型中，应用程序需要等待某个socket的I/O事件，但是它不能无限期地等待，必须在一定的时间内得到响应。这就需要使用一个定时器来控制等待操作的时间，超时后会强制返回。

在这样的网络编程中，每次使用select系统调用都需要重新设置socket的timeout，这会增加系统调用的次数和时间开销。为了优化效率，在Go语言中会在一个goroutine中程序使用轮询的方式来监听所有的socket，而不需要反复地执行select系统调用。在这种情况下，当一个socket被检测到有I/O事件时，就会调用netpolldeadlineimpl函数设置其下一次poll的deadline时间，以便在规定的时间内得到响应。

具体来说，netpolldeadlineimpl函数实际上是对epoll_wait系统调用进行封装，它会将下一次poll所需的等待时间转换为一个绝对时间，并将该时间作为参数传递给epoll_wait系统调用。如果指定的时间已经过期，那么返回的就是一个无限期的阻塞，否则就会在指定时间内阻塞等待I/O事件的到来。

总的来说，netpolldeadlineimpl函数实现了Go语言在Linux系统中的IO多路复用，提高了程序的响应效率。



### netpollDeadline

netpollDeadline是一个用来计算I/O操作截止时间的函数，它在Go语言运行时的runtime/netpoll.go源文件中。

在网络编程中，输入/输出（I/O）操作会导致阻塞，如果在一段时间内没有操作完成，则可能会使程序出现问题。为了避免这种情况，可以设置一个截止时间来控制I/O操作的处理时间，如果在这个时间内操作没有完成，则取消操作并返回错误。在Go语言中，可以使用SetDeadline方法来设置I/O操作的截止时间。

netpollDeadline的作用是根据指定的超时时间计算出I/O操作的截止时间。它首先获取当前时间，然后根据指定的超时时间计算出截止时间，返回一个Time类型的值表示操作的截止时间。如果超时时间为0，则直接返回0，表示操作不会超时。

计算截止时间的过程包括以下步骤：

1. 如果超时时间为0，则直接返回0，表示操作不会超时；
2. 获取当前时间；
3. 计算超时时间和当前时间之间的时间差，得到一个Duration类型的值；
4. 将时间差和当前时间相加，得到操作的截止时间。

通过这个函数，可以确保I/O操作不会超出指定的时间范围，从而保证程序的稳定性和可靠性。



### netpollReadDeadline

netpollReadDeadline函数的作用是设置网络轮询读取事件的超时时间。在Go语言中，网络轮询是通过epoll或者kqueue等系统调用实现的，它会将网络文件描述符上的读取事件注册到系统调用接口中，当有数据可读时，系统会通知Go语言的轮询机制，然后轮询机制将相应的读取事件传递给应用程序。

在网络轮询过程中，如果没有数据可读或者等待的时间过长，会导致应用程序阻塞或者超时。为了防止网络轮询事件过长时间阻塞，Go语言实现了netpollReadDeadline函数来设置网络轮询读取事件的超时时间。

具体来说，netpollReadDeadline函数会将轮询读取事件的超时时间设置为deadline参数指定的时间。当网络轮询事件等待时间超过deadline指定的时间时，该事件将被视为已超时，网络轮询机制将不再等待该事件的发生，并将事件传递给应用程序，由应用程序自行处理超时事件。

netpollReadDeadline函数的实现涉及了操作系统底层的系统调用和网络编程技术，需要对系统调用和网络编程有深入的了解才能更好地理解其功能和实现机制。



### netpollWriteDeadline

netpollWriteDeadline是在runtime源码中netpoll.go文件中定义的一个函数，它维护了关于网络IO write操作的超时信息，即在何时终止write操作。

具体来说，netpollWriteDeadline函数用于计算下一个网络IO write超时时间。它首先从write timers队列中取出最短时间的计时器，然后把这个最短时间减去当前时间，得到剩余时间。如果剩余时间小于等于0，则触发超时，返回true；否则，将剩余时间设置为netpollBlockPollDuration（在文件顶部定义的常量），并返回false。

写操作超时是很重要的，因为当网络IO write操作被阻塞时，应用程序的行为可能不可预知。对于服务器端应用，它可能导致客户端发生超时，从而关闭TCP连接。对于客户端应用，它可能导致响应等待时间过长，而造成用户体验糟糕。因此，维护网络IO write操作的超时信息是非常重要的。

总结：netpollWriteDeadline函数作用是维护网络IO write操作的超时信息，用于计算下一个网络IO write超时时间。



### alloc

在go/src/runtime中，netpoll.go文件中的alloc函数的主要作用是从堆中获取一块指定大小的内存，并返回其指针。在该函数中，会首先尝试从P的本地缓存中分配内存，如果本地缓存中没有足够的内存，则向全局堆申请。该函数主要涉及到以下步骤：

1. 检查P本地缓存中是否有内存可以使用，如果有，则直接返回该内存的指针。

2. 如果P本地缓存中没有可用内存，则检查全局堆的私有区域是否可以分配内存。如果私有区域没有足够的空间，则申请一块大的内存块，并将其切分成小块存入全局堆的公共区域中。

3. 如果私有区域中有足够的空间，则从私有区域中分配一块内存。

4. 如果全局堆的私有区域和公共区域都无法分配足够的内存，则会触发以下操作：

   - 调用M的nextFree函数来获取一个空闲的M。

   - 如果没有可用的M，则启动一个新的M。

   - 从新的M中获取一个新的P。

   - 将这个新的P的成员变量（scheduler等）初始化。

   - 重新检查P本地缓存和全局堆，直到可以分配一块足够的内存。

该函数在网络轮询中使用，主要用于分配文件描述符等资源。通过使用alloc函数，可以确保在需要资源的时候，总是可以分配到足够的内存，并且可以有效地利用内存池和P的本地缓存机制来提高效率。



### makeArg

makeArg函数在netpoll.go文件中被使用来生成pollDesc结构体对象所需要的arg字段。

具体来说，makeArg函数会创建一个类型为pollDescArg的结构体对象，然后将该对象的fields字段设置为传入的参数fields。接着，makeArg函数会使用unsafe.Pointer将pollDescArg对象转换为uintptr类型，作为arg字段的值。

在调用epollwait等网络相关函数时，系统会将arg字段传递给回调函数，以便回调函数可以使用该值来访问pollDesc结构体对象的其他字段。

因此，makeArg函数的作用就是创建pollDesc结构体对象时，为该对象的arg字段赋值，以便在之后的网络IO操作中可以使用该对象。



