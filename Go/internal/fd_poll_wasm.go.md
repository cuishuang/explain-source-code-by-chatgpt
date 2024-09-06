# File: fd_poll_wasm.go

fd_poll_wasm.go 是 Go 语言的标准库中 internal 包下的一个文件，主要用于实现 WebAssembly 操作系统环境下的文件描述符的异步 I/O 事件轮询。

WebAssembly 是一种基于堆栈的虚拟机，允许在浏览器中运行编译后的代码。在 WebAssembly 的操作系统环境下，存在一些系统调用和文件 I/O 等操作，而 fd_poll_wasm.go 就是利用 WebAssembly 的系统调用来实现文件描述符的异步 I/O 事件轮询的。

具体来说，fd_poll_wasm.go 实现了一种名为 poll 的事件轮询机制，其功能是检查一组文件描述符上的 I/O 事件是否就绪，并将就绪的文件描述符返回。这种机制可以用于实现异步 I/O 系统调用，例如 read、write 等，并提高系统的性能和响应速度。

总之，fd_poll_wasm.go 是 Go 语言的一个内部文件，用于实现 WebAssembly 操作系统环境下的文件描述符的异步 I/O 事件轮询。




---

### Structs:

### pollDesc

fd_poll_wasm.go文件是Go编程语言的内部包，它提供了与操作系统IO接口相关的实现，其中pollDesc结构体是其中的一个用于轮询文件描述符的结构体。

pollDesc结构体主要用于在操作系统上请求和异步等待文件描述符的可读/可写事件并通过Go的运行时系统进行跟踪。这个结构体中包含了若干个字段，常用的有：

- fd int: 文件描述符
- closing int32: 用于检测文件描述符的关闭状态
- rseq uintptr: 读同步序列标识
- wseq uintptr: 写同步序列标识
- user uint32: 用户标记，用来标记pollDesc是否被使用
- coWait *sudog: 用于等待协程的调度

pollDesc结构体的作用是让Go运行时系统能够跟踪文件描述符的状态，并且在需要时通知等待的协程继续执行。在WebAssembly平台，Go编程语言并没有实现真正的系统调用，因此不能像在PC或者服务器端那样直接调用操作系统提供的API来进行文件IO操作，而是通过JS的WebAPI来模拟文件IO，这就需要借助pollDesc结构体等相关类型在WebAssembly中实现类似的机制来模拟IO操作。



## Functions:

### init

在go/src/internal/fd 目录中，fd_poll_wasm.go文件是用于WebAssembly平台的轮询器实现。init函数是一个特殊的函数，当程序启动时会被自动调用。它主要的作用是初始化某些变量或执行一些必要的操作以便程序运行。具体到fd_poll_wasm.go文件，init函数主要做了以下两件事：

1. 初始化一个socket数组

这个socket数组实际上是一个全局变量，用于保存所有需要轮询的socket。通过调用函数fd.GlobalSet(n, socketArray)来初始化。其中n是WebAssembly的全局变量索引，socketArray是一个用于保存socket的数组。

2. 注册轮询器函数

调用函数syscall/js.FuncOf来创建一个用于轮询的JavaScript函数poll。然后通过调用函数syscall/js.Global().Set("poll", poll)将该函数注册到全局变量poll上。poll函数主要做了以下两件事情：

- 调用JavaScript的cb()函数，并将socket数组作为参数传入。
- 将发生事件的socket的索引数组返回给JavaScript，以便JavaScript回调函数可以处理这些事件。

总之，init函数的作用是初始化需要轮询的socket数组和注册WebAssembly平台的轮询器函数。它们是WebAssembly平台实现轮询器的关键。



### close

在WASM平台上，文件描述符的实现方式与其他平台不同。文件描述符在WASM中实际上是浏览器中的XMLHttpRequest对象，而不是标准的文件描述符。因此，close函数在WASM中的作用是关闭正在使用的XMLHttpRequest对象，并在必要时调用回调函数。

具体来说，close函数会获取包含XMLHttpRequest对象的pollFD结构体，并调用XMLHttpRequest对象的abort方法来关闭连接。如果pollFD结构体中有回调函数，它们将在关闭连接后被调用。最后，pollFD的值被重置，以便下次使用。

close函数的目的是确保在使用完XMLHttpRequest对象后，正确释放资源并避免内存泄漏。



### evict

在Go语言中，fd_poll_wasm.go文件是实现在WebAssembly平台上支持文件描述符轮询的功能。该文件中的evict函数的作用是从轮询的文件描述符列表中删除指定的文件描述符。在WebAssembly平台中，poll函数仅支持Simple I/O模型，即每次只能处理一个文件描述符。在处理多个文件描述符时，需要轮询这些文件描述符，fd_poll_wasm.go实现了这个轮询功能。但是，如果轮询的文件描述符列表中有文件描述符已经被关闭，则需要将其从列表中删除，这就是evict函数的作用。

evict函数的实现较为简单，首先通过for循环遍历轮询的文件描述符列表，找到需要删除的文件描述符的位置。然后将该位置对应的元素删除，并将列表的长度减1。最后，如果删除的元素不在列表的最后一个位置，则将列表的最后一个元素移动到删除元素的位置。

总之，evict函数保持轮询文件描述符列表的正确性，确保在从列表中读取或写入数据时，不会操作已关闭的文件描述符，提高了轮询的效率和稳定性。



### prepare

fd_poll_wasm.go是Go语言中用于WebAssembly平台的文件，其中的prepare函数是用来准备文件描述符（file descriptor）以进行轮询操作。

在WebAssembly平台上，由于浏览器的限制，Go语言程序并不能直接访问底层系统资源，比如说网络文件描述符。为了解决这个问题，Go语言在WebAssembly平台上使用了 Javascript 封装了底层的轮询事件，而prepare就是一个用来将 JavaScript 中的轮询事件适配成文件描述符的函数。

prepare 函数的主要作用：

- 读取函数调用参数，包括文件描述符和轮询事件类型；
- 将文件描述符转换成 JavaScript 对象；
- 调用 JavaScript 函数（或回调）并附带参数，实现文件描述符的轮询操作；
- 将轮询操作结果返回给 Go 语言程序。

总的来说，prepare 函数让 Go 语言程序在 WebAssembly 平台上可以以与其他平台相同的方式操作文件描述符，从而扩展了 Go 语言在 WebAssembly 平台上的使用场景。



### prepareRead

文件fd_poll_wasm.go中的prepareRead函数是针对WebAssembly平台的实现，其主要作用是准备进行读取操作。在WebAssembly平台上，无法使用标准的Unix文件描述符，因此需要进行一些额外的处理。

具体来说，prepareRead函数会先检查文件描述符是否有效，然后将其转换为一个WebAssembly的JavaScript对象。接下来，函数会将读取请求打包成一个JavaScript对象，包括文件描述符、读取长度、读取位置等信息。最后，将这个JavaScript对象传递给JavaScript环境中的预定义函数，以便进行实际的读取操作。

总的来说，prepareRead函数的作用是为在WebAssembly平台上对文件进行读取操作做好必要的准备工作。



### prepareWrite

prepareWrite函数在Go语言中的WebAssembly平台下，用于准备在文件描述符上进行写操作。

在WebAssembly平台下，Go语言的文件读写操作是通过JavaScript代码来实现的。因此，准备写操作需要进行以下步骤：

1. 准备一个消息发送队列，用于发送写操作请求到JavaScript代码中。

2. 将数据写入到一个缓冲区中，并将缓冲区添加到消息发送队列中。

3. 发送写操作请求到JavaScript，请求JavaScript将缓冲区中的数据写入到文件描述符中。

prepareWrite函数就是用于完成这些步骤的。它接收一个*FD类型的参数，表示需要准备写操作的文件描述符。

首先，prepareWrite函数会判断当前文件描述符是否已经指定了写缓冲区。如果没有，则会创建一个新缓冲区，并分配一块内存区域用于存放缓冲区的数据。然后，它会将这个新的缓冲区添加到文件描述符的缓冲区列表中，并返回缓冲区的索引。

如果文件描述符已经指定了写缓冲区，则直接返回缓冲区的索引。

最后，prepareWrite函数将会发送一个写操作请求到JavaScript代码中，请求JavaScript将缓冲区中的数据写入文件描述符中。具体而言，它会将一个writeOp类型的指针添加到文件描述符的写操作队列中。

当JavaScript代码完成写操作后，将会调用JavaScript的回调函数，回调函数会从写操作队列中弹出writeOp类型的指针，并将其状态设置为完成状态。这个机制可以保证WebAssembly平台下的文件读写操作是异步的。



### wait

在go/src/internal中，fd_poll_wasm.go文件中的wait func是用于等待并处理异步事件的函数，主要用于操作系统的异步I/O处理。该函数会一直等待直到发生一个异步事件，然后将其作为参数传递给回调函数。

具体来说，wait func会首先调用pollwait函数，该函数会阻止并等待I/O事件的发生。一旦事件发生，wait func会调用ready函数处理事件并将其作为参数传递给回调函数。

在处理事件时，wait func会遍历文件描述符列表并将每个文件描述符的事件状态传递给回调函数。回调函数可以根据事件状态决定如何处理该事件（例如读取或写入数据）。处理完所有事件后，wait func将重新开始等待下一个事件的发生。

总之，wait func是用于等待并处理异步事件的函数，可以帮助程序处理底层操作系统的I/O异步事件。



### waitRead

waitRead函数是用于等待读取操作完成的函数。在WebAssembly环境下，无法直接使用系统调用实现异步I/O，因此需要通过轮询使用JavaScript的API来模拟异步I/O操作。

该函数接受一个文件描述符fd和一个timeout时间，在timeout时间内等待文件描述符fd可以进行读取操作，如果可以进行读取操作则返回nil，否则返回错误。

在该函数内部，首先使用javascript.Global().Call("setTimeout"，callback，timeout)在指定的timeout时间后执行callback函数。callback函数则会检查fd是否可读，如果可以读取，则返回，否则再次调用javascript.Global().Call("setTimeout"，callback，timeout)函数进行轮询，直到timeout时间到或者可以进行读取操作为止。



### waitWrite

waitWrite这个函数是在WASM环境下实现的，用于等待套接字变为可写。

在WebAssembly中，套接字的写操作是异步执行的。当我们调用套接字的写方法时，写操作被提交到浏览器内部，然后我们的代码就可以继续往下执行。

为了检查写操作是否完成，我们需要使用Promise对象或事件监听器来监听套接字写操作完成的消息。由于这些机制的实现是异步的，因此我们需要等待写操作完成后才能继续执行代码。

waitWrite函数就是用来实现这个等待过程的。它在后台开启一个goroutine，并且监听套接字的写事件。一旦写操作完成，它就会在chan中写入一个标志位，然后等待调用它的函数接收这个标志位，以便继续执行后续操作。

总的来说，waitWrite函数的作用就是等待套接字的写操作完成，并且通知调用者代码可以继续执行后续操作了。



### waitCanceled

文件`fd_poll_wasm.go`是Go语言中网络轮询实现的文件，其中`waitCanceled`函数的作用是等待一个事件的发生并且监听是否被取消。

`waitCanceled`函数中使用了`select`语句实现了监听多个事件：

```go
select {
    case <-c.done:
        return ErrPendingIOCanceled
    case <-timeoutC:
        return ErrTimeout
    case <-fdC:
        return nil
}
```

`c.done`是一个`time.After`实例，它会在一定时间后发送一个值，用于快速取消监听。而`timeoutC`是一个超时信道，用于实现函数的超时机制，如果在一定时间内没有收到任何事件，函数会返回超时错误。最后的`fdC`则是实际监听网络连接事件的信道，如果连接上有可读或可写事件发生，会通过该信道向`select`语句发送一个值，然后函数返回`nil`表示可以执行IO操作。

总体来说，`waitCanceled`函数的作用是监听网络连接上的读写事件，当事件发生时返回nil。但在等待过程中，函数还会监听是否被取消或超时，如果发生则返回相应的错误。



### pollable

在Go语言的内部文件夹中，fd_poll_wasm.go文件是用于处理WASM平台的文件描述符轮询的代码。其中，pollable()函数是用于确定指定文件描述符是否可轮询的函数。

在WASM平台上，文件描述符通常是Socket或者类似于WebSocket的对象，这些对象不支持传统的I/O选择器（select/poll/epoll）操作。因此，就需要使用WASM平台特有的方法来确定一个文件描述符是否可轮询。

pollable()函数的工作原理是将文件描述符尝试传递给JavaScript端（也就是客户端浏览器上运行的JavaScript代码），看看JavaScript能否正确识别并处理该文件描述符，如果能够正确工作，则说明该文件描述符是可以轮询的。

此函数的返回值是一个布尔类型，如果文件描述符可轮询，则返回true，否则返回false。在Go语言中，这个函数会在轮询文件描述符前被调用，以确保在可能轮询到无效的文件描述符时不会出现崩溃或错误的行为。



### SetDeadline

fd_poll_wasm.go这个文件是Go语言标准库内部的一个文件，是用来处理在WebAssembly上的socket连接的。

在这个文件中的SetDeadline这个func的作用是设置一个socket连接的读写超时时间。具体来说，它可以让应用程序在一定的时间内等待一个socket连接收到数据，或是发送数据成功，如果超过了这个时间，那么连接会被关闭，从而避免了连接处于无限等待的状态。

这个函数的实现方式是通过向WebAssembly中发送一个异步I/O调用来实现的。在这个调用中，我们可以指定一个超时时间，并在超过指定时间后返回一个错误码。接下来，我们可以通过这个错误码来判断是否需要关闭连接。



### SetReadDeadline

fd_poll_wasm.go是为WebAssembly平台提供网络I/O操作的代码文件。其中的SetReadDeadline函数是设置读取操作的截止时间。具体来说，它允许程序员向系统请求在指定时间之前完成读取操作，否则就会返回一个错误。

具体实现方法是使用JavaScript中的setTimeout函数，该函数可以在指定延迟后执行回调函数。在SetReadDeadline中，将通过JavaScript的setTimeout函数计算出截止时间与当前时间的差值，然后将其传递给JavaScript的setTimeout函数，以在到达截止时间时触发回调函数。如果读取操作在截止时间之前完成，则回调函数不会被触发，否则将返回一个错误。

需要注意的是，在WebAssembly中，网络I/O操作被封装在JavaScript API中，因此使用SetReadDeadline函数的前提是已经建立好了JavaScript API和WebAssembly的交互，这通常需要进行额外的配置和安装。



### SetWriteDeadline

Go语言中的SetWriteDeadline函数可以设置一个写入操作的截止时间（deadline），即如果在指定的时间内无法完成写入操作，写入操作将被取消并返回错误。该函数位于fd_poll_wasm.go文件中，是针对WebAssembly平台的文件描述符操作实现。

在WebAssembly平台中，文件系统通常是通过JavaScript运行时提供的，因此与JavaScript运行时的交互是必须的。SetWriteDeadline函数的作用是设置文件描述符的写入截止时间，然后将该操作添加到JavaScript运行时的事件队列中，JavaScript运行时会在指定的时间后将事件返回给Go语言程序。

实际上，在WebAssembly平台上，Socket和文件描述符的I/O操作都是通过JavaScript运行时实现的。因此，SetWriteDeadline函数在调用JavaScript运行时时设置了一个事件，然后等待JavaScript运行时返回事件，从而实现了对写入操作的截止时间的控制。

总之，SetWriteDeadline函数是fd_poll_wasm.go文件中的一个函数，主要作用是设置写入操作的截止时间，从而控制WebAssembly平台文件描述符的写入操作。它在构建针对WebAssembly平台的Golang应用程序时非常有用。



### setDeadlineImpl

文件fd_poll_wasm.go是Go语言的标准库中的一个文件，它提供了在WebAssembly环境下模拟poll函数的实现。setDeadlineImpl函数负责设置文件句柄的读取和写入超时时间。

具体来说，该函数将需要设置超时的文件句柄转换成底层的JavaScript对象，然后利用JavaScript提供的setTimeout函数设置读取和写入超时。当超过设置的时间时，JavaScript会自动取消文件句柄上的I/O操作，并将错误码返回给Go层。

在WebAssembly环境下，由于操作系统接口受到限制，无法直接使用传统的I/O模型。通过使用poll函数的模拟实现，可以实现基本的网络通信功能，从而在WebAssembly环境下运行Go程序。

综上所述，setDeadlineImpl函数在fd_poll_wasm.go文件中的作用是为文件句柄设置读取和写入超时时间，进而实现在WebAssembly环境下模拟poll函数的功能。



### IsPollDescriptor

IsPollDescriptor函数用于检查给定的文件描述符是否是支持poll系统调用的描述符类型。

在操作系统中，文件描述符是一个整数，用于唯一标识一个打开的文件、套接字或其他I/O资源。IsPollDescriptor函数通过检查文件描述符的类型，从而确定它是否支持poll系统调用。如果支持，则返回true，否则返回false。

这个函数对于实现跨平台的I/O操作非常重要。因为不同的操作系统可能有不同类型的I/O资源和不同的文件描述符类型。使用IsPollDescriptor函数可以轻松地确定是否可以在给定的平台上使用poll系统调用来实现I/O操作，从而提高了跨平台代码的可移植性。



