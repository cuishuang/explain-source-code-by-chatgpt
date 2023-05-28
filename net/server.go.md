# File: server.go

server.go文件是Go语言标准库中net包的核心文件之一，它定义了net包的主要类型和函数实现，主要用于实现网络服务器的基本功能。

具体来说，server.go文件主要包括以下几个方面的内容：

1. 类型定义：定义了多个类型，包括Listener、Conn、TCPListener、 TCPConn等。这些类型主要用于表示网络监听器、连接等概念，并提供了相关的方法来实现网络通信功能。

2. 函数实现：server.go文件中定义了多个函数，包括Listen函数、ListenTCP函数、Dial函数等。这些函数主要用于实现网络连接、监听等功能，是构建基于TCP/IP协议的服务器应用的基础。

3. 内部实现：server.go文件还包含了net包内部实现的一些重要内容，例如goroutine调度、I/O多路复用等技术实现。这些内容为网络通信提供高效稳定的支持，是Go语言网络编程的特色之一。

总之，server.go文件是Go语言网络编程的核心代码之一，它为基于TCP/IP协议的网络服务器提供了基础的类型和函数实现，是构建高效稳定的网络应用的重要组成部分。




---

### Var:

### serveFlag

serveFlag变量是用来控制网络服务的行为的标志变量。它是一种标识位，用于确定网络服务器在运行时的行为。

serveFlag变量是一个int类型的变量，在代码中的初始化值为1。

主要作用：

1. 控制网络服务器是否提供Keep-Alive。

如果serveFlag为1，则网络服务器将启用HTTP Keep-Alive。如果为0，则不启用HTTP Keep-Alive。

HTTP Keep-Alive是一种HTTP协议的扩展，它允许在同一连接上发送多个请求和响应，从而减少了网络服务器和浏览器之间连接的数量。

2. 控制网络服务器是否允许TCP快速ACK。

如果serveFlag为1，则网络服务器将允许TCP快速ACK。如果为0，则不允许TCP快速ACK。

TCP快速ACK是一种TCP协议的扩展，它允许TCP立即发送一个ACK响应，而无需等待一段时间。

3. 控制网络服务器是否提供带有大量请求的限流措施。

如果serveFlag为1，则网络服务器将启用请求限制。如果为0，则不限制请求。

请求限制是一种利用网络服务器资源的方法，以确保服务器在遇到大量请求时不会过载。

总之，serveFlag变量是用来控制网络服务器行为的标志变量。它可以控制网络服务器是否提供Keep-Alive、是否允许TCP快速ACK和是否提供请求限制等功能，以确保服务器在某些情况下的高效运行。






---

### Structs:

### Server

在go/src/net中，server.go这个文件中的Server结构体是表示一个网络服务器的对象。它的作用是提供一个通用的框架，用于构建各种类型的网络服务。Server结构体包含了很多参数和方法，可以配置服务器的一些属性、启动服务器、处理连接请求、并发处理客户端请求等。

Server结构体是一个聚合类的实现，将底层网络协议的实现和业务逻辑的实现分离，并提供了一个通用的框架使得用户可以方便地实现自己的网络应用。

Server结构体中的主要方法包括:

1. ListenAndServe：启动服务器并开始监听用户的连接请求

2. Serve：接受连接，处理请求，返回响应

3. SetKeepAlivesEnabled：设置是否启用保持连接

4. SetReadTimeout：设置读取超时时间

5. SetWriteTimeout：设置写入超时时间

6. SetConnState：设置连接状态的回调函数

7. SetErrorLog：设置错误日志

8. SetKeepAlive：设置TCP keepalive参数

通过Server结构体的方法，用户可以轻松配置服务器的属性、启动服务器、处理请求，并发处理客户端请求，保持连接等，实现自己的网络应用。



### closeIdleTransport

closeIdleTransport这个结构体是用于关闭空闲的HTTP/1连接的。

在HTTP/1协议中，一次HTTP请求和响应的交互需要通过一个TCP连接来进行。当一个HTTP请求完成后，如果TCP连接没有被关闭，那么这个连接就可以被用于发送下一个HTTP请求，这样可以避免建立和关闭TCP连接的开销。但是，如果连接一直没有被复用，就会成为空闲连接，这时候如果不关闭它，会浪费系统资源。

closeIdleTransport结构体的作用就是定期检查所有的空闲连接，如果超过一定时间没有被复用，就将其关闭。这样可以避免空闲连接浪费系统资源，并且可以让系统的网络安全更加健康。

closeIdleTransport结构体的实现利用了Go语言的time包，定期触发一个time.Tick操作，使用连接池中的连接发送一个心跳请求，如果连接回复正常，则表示连接可用，如果连接长时间没有回复，则表示连接已过期，需要关闭。



## Functions:

### newLocalListener

newLocalListener 函数是用来创建本地（使用 IP 地址 "127.0.0.1" 和随机端口号）监听器的函数。该函数将会创建并返回一个未被使用的 TCP/IP 网络地址，用于在本地服务器上开始监听 TCP 连接请求。

在 Go 标准库中，该函数主要用于一些测试相关的场景。比如在测试中需要模拟一个本地服务，那么就可以使用该函数创建出一个监听器，并将该监听器绑定到被测试代码中的某个端点。

对于该函数的具体实现而言，它主要依赖于 Go 标准库中的 `net.ListenTCP()` 函数来创建 TCP 监听器，并通过该函数的返回值实现在本地启动服务。在该方法内部，首先调用 `net.TCPAddr{}` 方法创建 TCP 监听的地址，然后将其传递给 `net.ListenTCP()` 方法进行实际的监听器创建，并最终返回监听器对象即可。

值得注意的是，在该函数中，我们使用了一些基本的错误处理和异常检测，以确保在监听器创建必要的过程中，不会出现任何的非预期行为。



### init

在Go语言中，init函数是一个特殊的函数。每个包可以有多个init函数，它们在包被导入时自动执行，但执行顺序不能保证。init函数不能被显式调用。

在net包的server.go文件中，init函数主要有以下作用：

1.注册HTTP handler回调函数

在init函数中，通过调用`http.HandleFunc`函数注册了一些HTTP请求的回调函数。这些回调函数定义了HTTP服务器如何处理来自客户端的请求，并返回响应结果。这些回调函数的实现在其他文件中，如http/server.go等。

2. 初始化DNS解析器

在init函数中，通过调用`go lookupIPEnabled()`函数初始化了DNS解析器。该函数检查操作系统是否支持IP地址解析，然后设置了全局变量enableDNS，用来决定DNS解析器是否可以使用。

3. 初始化Unix socket性能参数

在init函数中，通过调用`initUnixSocket()`函数初始化了Unix socket的性能参数。Unix socket是一种本地进程间通信机制，而该函数则设置Unix socket的缓冲区大小等性能参数，以提高其性能。

总而言之，init函数是Go语言中一个特殊的函数，它在包被导入时自动执行。在net包的server.go文件中，init函数主要用于初始化一些全局变量和可执行的特殊函数。



### strSliceContainsPrefix

strSliceContainsPrefix函数用于判断字符串切片中是否存在以指定前缀开头的字符串。

该函数的实现很简单，它遍历了给定的字符串切片，对每个字符串都检查它是否以指定前缀开头。如果找到了这样的字符串，则返回true，否则返回false。

这个函数在net包中的server.go文件中被用来检查HTTP请求头中是否包含指定的字符串前缀。如果包含，则认为这个请求是一个websocket握手请求。换句话说，这个函数的作用是实现了websocket协议的解析器。



### NewServer

NewServer函数是net包中的函数之一，其作用是创建一个新的服务器对象。具体而言，该函数返回一个*Server的指针，该指针表示一个TCP服务器。在使用NewServer函数时，需要提供一个ServerOptions结构体作为参数，其中包括要监听的地址和端口以及其他一些可选的配置参数。

NewServer函数内部会根据传入的ServerOptions创建一个新的Server对象，并返回指向该对象的指针。这个Server对象一般会在程序的入口处被创建，并在服务器主循环中被使用。

Server对象包含了TCP服务器的一些基本属性和方法，例如Listen函数会开始监听传入的连接，Serve函数则会启动服务器的主循环，接受客户端连接并处理请求。这些方法都可以通过Server对象调用。

总的来说，NewServer函数是net包中创建TCP服务器的入口函数，其可以创建一个新的Server对象，帮助程序员创建TCP服务器，以便与客户端进行通信。



### NewUnstartedServer

NewUnstartedServer函数是net包中的一个函数，用于创建一个未启动的Server对象。本函数的作用是创建一个未启动的网络Socket服务器对象，该对象可以被用户自己进行配置，例如设置监听地址和监听端口等，最终通过调用该对象的ListenAndServe()方法启动服务器。

在实现过程中，NewUnstartedServer函数会返回一个指向未启动Server对象的指针，用户需要在执行ListenAndServe方法之前对该Server对象进行配置，例如通过调用Server对象的Serve方法来设置服务器的监听地址和监听端口等。

总之，NewUnstartedServer函数的作用是创建一个未启动的Server对象，并返回一个指向该对象的指针，用户可以通过该对象的方法来配置服务器的参数和启动服务器等操作。



### Start

Start这个func实现了TCP服务器的启动功能，作用如下：

1. 创建TCP Listener对象并进行TCP监听：通过调用net.Listen函数创建Listener对象，并指定协议类型、IP地址及端口号，实现对TCP客户端的监听。

2. 处理TCP客户端请求：在for循环中，不断调用Listener对象的Accept方法，监听并等待TCP客户端的连接请求，获取到连接请求后，将连接封装成Conn对象，然后交给新的goroutine处理。

3. 处理TCP连接请求：在新的goroutine中，通过调用Server的Handler方法，将Conn对象传递给处理函数进行处理。处理函数可以是自定义的处理函数，也可以是标准库中提供的HTTP处理函数。

4. 关闭TCP Listener对象：在defer函数中调用Listener对象的Close方法，用于关闭TCP监听器，释放相关资源。

总之，Start这个func是实现TCP服务器启动的核心代码，通过不断地监听TCP客户端的连接请求并创建TCP连接对象，然后将连接对象交给处理函数进行处理，从而实现对TCP客户端的响应处理。



### StartTLS

StartTLS这个函数是用来启用TLS协议的。TLS协议是一种安全的传输层协议，用于保护数据通信的安全性和完整性，防止数据被窃听、篡改或者伪造。

在server.go文件中，StartTLS函数被调用时，它会根据传入的tls.Config配置开启TLS。它首先会调用Conn方法获取一个net.Conn接口类型的对象，然后再以此对象为参数调用tls.Server函数，用来启动TLS协议。最后，它会返回一个tls.Conn对象，这个对象可以被用来进行加密的读写操作。

StartTLS函数是在TCP服务器监听器已经监听了端口，并已经创建了TCP连接之后，才会被调用。它将TCP连接和TLS协议结合起来，从而提供了一个安全的通信通道。当客户端向服务器发送请求时，服务器收到请求之后就会通过TLS协议加密传输数据，保证数据的机密性和完整性。



### NewTLSServer

在go/src/net/server.go文件中，NewTLSServer函数用于创建一个新的TLS服务器。它接受一个TCP地址和一个TLS配置作为参数，并返回一个指向TLS服务器的指针。

具体来说，该函数执行以下步骤：

1. 调用net.ListenTCP(addr)函数创建一个TCP监听器。

2. 使用传入的TLS配置，调用tls.NewListener(listener, config)函数创建一个TLS监听器。该监听器将自动在接受连接的时候完成TLS握手过程。

3. 创建一个新的服务器实例，并将其绑定到TLS监听器上，然后将其返回。这个服务器实例将在监听器上接受连接并处理传入的请求。

使用NewTLSServer函数可以方便地创建一个安全的TLS服务器，在保证数据传输安全的同时提供高效的服务。通过传入不同的TLS配置，我们可以实现不同的TLS握手和认证机制，以满足不同的安全需求。



### Close

server.go文件中的Close函数是用于关闭一个Listener的，它的作用是停止接受新的连接请求，并关闭所有已经建立的连接。

具体来说，Close函数会将Listener的状态改变为closed，阻止后续的Accept函数调用，并关闭Listener底层所使用的网络连接。同时，它还会关闭所有已经建立的连接，即调用每个已经建立连接的Conn的Close函数。

Close函数的实现比较简单，主要是对一些内部状态的修改和调用底层网络连接的Close方法。当调用Close函数后，监听者就不能再接受新的连接请求，已经连接上的客户端也会收到一个关闭的消息，这样就可以安全地关闭TCP连接，并释放所有相关资源。

总之，Close函数是用来关闭Listener和其中所有的连接，并释放所占用的资源的，是网络编程中非常重要的操作。



### logCloseHangDebugInfo

logCloseHangDebugInfo是一个函数，用于记录关闭超时的连接的调试信息。该函数在server.go文件中被定义。

在网络通信中，有时候会出现关闭超时的情况，即客户端或服务器发送关闭连接的请求，但是对方没有回应，连接一直处于关闭状态，这就会导致连接无法正常释放，一定程度上影响网络通信的效率和性能。

为了解决这个问题，logCloseHangDebugInfo函数会记录关闭超时的连接的相关信息，包括连接地址、接收和发送数据的缓冲区大小、连接的开始和结束时间等，并将这些信息记录到系统日志中。这样一来，运维人员或开发人员就能够更加方便地分析连接关闭超时的原因，进而采取相应的措施来解决这个问题。

总之，logCloseHangDebugInfo函数是一个非常重要的调试函数，它可以帮助保证网络通信的高效性和稳定性。



### CloseClientConnections

CloseClientConnections是net包中server.go文件中的一个函数，其作用是关闭服务器端已建立的所有客户端连接。

在网络通信中，服务器端为了处理来自客户端的请求，通常需要建立多个客户端连接。当服务器端需要停止服务或者更新代码时，需要关闭这些客户端连接并释放相关资源，以便让新的请求能够被处理。

CloseClientConnections函数就是用于完成这个任务的。它通过遍历server中所有的 activeConns 切片，逐个调用conn.Close()关闭每个客户端连接。同时，它还会清空所有connReadLimit的限制，并把server的activeConns切片清空。

需要注意的是，CloseClientConnections函数只会关闭已建立的连接，并不会阻止客户端再次连接进来。如果需要停止所有连接请求，需要在外部使用listener.Accept函数返回的net.Conn对象调用Close函数，关闭服务端监听。



### Certificate

Certificate是net包中tls.Config结构体的一个方法，用于返回当前设置的证书列表。当使用TLS协议建立加密通信时，服务器需要提供一个证书来证明自己的身份。Certificate方法用于返回当前设置的服务器证书，即tls.Certificate类型的对象。

tls.Certificate类型的对象包含一个或多个公钥证书和对应的私钥。服务器使用私钥来生成数字签名，客户端使用公钥验证签名以确保通信的安全性。

在使用Certificate方法时，需要首先创建tls.Certificate类型的对象，然后调用SetCertificate方法将其设置到tls.Config结构体中，最后调用ListenAndServeTLS方法启动TLS加密服务器。

除了服务器证书，Certificate方法还可以返回客户端证书列表。客户端证书通常用于验证客户端的身份，以保护服务器免受欺诈或中间人攻击。



### Client

server.go文件中的Client函数实现了TCP连接的操作。它接受一个net.Conn类型的参数，该参数表示已建立的TCP连接。Client函数将该连接封装为一个*Client类型并返回。

*Client类型是一个TCP客户端连接的封装。它提供了一些方法，如读取和写入数据，关闭连接等。

在使用net包建立TCP服务器时，每当有客户端连接到服务器时，服务器会创建一个新的goroutine来处理该连接。在该goroutine中，服务器会调用Client函数将连接封装为一个*Client类型，并使用该类型执行一些具体的读写操作。

通过Client函数，我们可以访问连接的本地地址、远程地址、读写缓冲区等信息。这样，我们可以对连接进行更加细粒度的控制和调整，以满足不同的需求。



### goServe

goServe是net包中的一个函数，其作用是启动一个网络服务器，可以监听和接受客户端的连接请求。在 server.go 文件中，goServe 函数作为 Server 类型的方法被实现。

具体来说，goServe 函数会启动一个 TCP 监听器，并且通过循环接受客户端连接请求。每次接受到一个新的连接请求时，它会通过 Server 类型中的 handler 处理函数来处理该连接。handler 函数是在 Server 创建时指定的，用于处理客户端的请求并返回相应的响应。

当接受到一个连接时，goServe 函数会创建一个新的 goroutine 来处理客户端的请求。这样可以让服务器同时处理多个客户端连接请求，提高了服务器的并发性能。

另外，goServe 函数还会在处理客户端连接请求时，根据需要创建临时 goroutine 来处理一些复杂的任务，例如读写网络数据或者处理某些计算密集型的操作。这些任务在 goroutine 中异步执行，不会阻塞主线程，从而提高了服务器的响应速度。

综上，goServe 函数是一个非常重要的函数，它是net包中实现网络服务器的核心函数之一。通过它，可以方便地创建一个高效的、并发的网络服务器，用于处理客户端的请求。



### wrap

func wrap(f func(net.Listener) error) func(net.Listener) error {
	return func(l net.Listener) error {
		return f(tcpKeepAliveListener{l.(*net.TCPListener)})
	}
}

这个wrap函数的作用是生成一个新的func，并返回给调用者。新生成的这个func会执行传递进来的func（即参数f）并将传递给它的Listener对象（即参数l）转换成tcpKeepAliveListener类型的对象，再将转换后的对象作为参数传递给f进行处理。

在这里，wrap函数的作用就是将传递进来的func进行包装，使得该函数可以接收一个普通的Listener对象，并将其转换为TCP连接，使得该函数可以支持TCP的相关特性，比如TCP KeepAlive等。

在net包中，有许多函数都需要TCP连接对象作为参数，而wrap函数正是为了使得这些函数可以更加灵活地接收普通的Listener对象，并将其转换成TCP连接对象进行处理。



### closeConn

closeConn是一个用于关闭连接的函数，它的主要作用是在服务器处理完一个连接后，关闭该连接以释放资源。具体来说，closeConn会执行以下操作：

1. 将底层TCP连接关闭。

2. 从连接池中移除该连接。

3. 如果有Read和Write错误，那么会记录日志。

在服务器中，每个客户端连接都被封装为一个Conn对象。当服务器处理完一个连接后，需要使用closeConn函数关闭Conn对象，以确保不会产生资源泄漏。关闭连接还可以通知客户端连接已经断开，让客户端能够及时做出响应。

需要注意的是，closeConn函数并不是直接从内核中关闭底层TCP连接，而是将连接的关闭操作加入一个待处理的事件队列中，由事件循环线程进行处理。这是因为在多线程环境下，直接关闭连接可能会产生竞态条件或者死锁，所以需要通过事件驱动的方式来处理连接关闭操作。



### closeConnChan

在 go/src/net 中的 server.go 文件中，closeConnChan 函数的作用是关闭连接通道。它是在内部函数 serve 中被调用的，用于关闭连接通道，并在关闭通道后，通过向 done 队列中添加一个元素，来释放 serve 中的所有并发 goroutine。

在 serve 函数中，每个连接对应一个 goroutine，而 closeConnChan 函数则是用于关闭并结束这些 goroutine 的。当服务器需要关闭时，serve 函数会关闭连接通道，并阻止新的连接加入。然后，closeConnChan 函数会循环遍历所有连接，将它们从 server 队列中删除。在遍历过程中，如果某个连接已经关闭，则直接进行下一次循环。

当所有的连接都被删除后，closeConnChan 函数会向 done 队列中添加一个元素，来通知 serve 函数和其它相关的 goroutine 下线。最后，closeConnChan 函数返回一个空的 done 队列，以释放所有关联的 goroutine，最终结束整个服务进程的运行。



