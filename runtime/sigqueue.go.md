# File: sigqueue.go

sigqueue.go这个文件是Go语言运行时系统中处理信号队列的代码实现。当操作系统向Go程序发送信号时，该信号封装成一个siginfo结构体，并放入一个信号队列中。Go程序需要通过sigqueue.go中的代码来从信号队列中取出信号并处理它们。

具体来说，sigqueue.go中定义了一个sigqueue类型，它表示一个信号队列，其中包含了一个“信号池”（signalPool），用于存储可用的信号结构体，以及一个“等待池”（waiterPool），用于存储等待信号的goroutine。

sigqueue.go中的代码提供了以下几个函数来处理信号队列：

- signalAllocate：从信号池中获取一个可用的信号结构体。
- signalQueueInit：对信号队列进行初始化。
- signalQueueEnable：启用信号队列，使得信号可以被处理。
- signalQueueDisable：禁用信号队列，停止信号的处理。
- signalQueueDeliver：从信号队列中取出信号并进行处理。
- signalWaitUntilIdle：等待信号队列中的信号全部被处理完毕。
- signalNotifyGoroutine：通知等待池中的goroutine有新的信号可以处理了。

总之，sigqueue.go这个文件是Go语言运行时系统中非常重要的一部分，它保证了Go程序能够正确、及时地处理系统发来的各种信号，从而保证程序的稳定性和可靠性。




---

### Var:

### sig

在go/src/runtime/sigqueue.go这个文件中，sig这个变量的作用是维护一个带有缓冲的信号通道，用于按顺序传递接收的信号到所有等待的Goroutine。

在sigqueue.go的init()函数中初始化sig变量，并设置了一个信号处理程序，该处理程序会将接收到的信号写入sig通道中。

在runtime库中，当信号到达时，信号处理程序会将信号写入sig通道。任何等待接收信号的Goroutine都可以从该通道中读取信号。

这种方式可以确保所有接收的信号被按照顺序传递到等待的Goroutine，而不是随机传递或重叠传递到多个Goroutine。

总之，sig变量的作用是作为一个信号通道，确保所有接收到的信号被顺序地传递给等待的Goroutine。



## Functions:

### sigsend

sigsend函数的主要作用是向指定的进程发送信号。它通过调用系統调用sigqueue来完成这个任务。在Linux系统中，sigqueue系统调用是通过调用kill函数来实现的。sigsend使用sigqueue函数发送一个信号给一个指定的进程或者一个进程组。它还负责设置当前进程的信号掩码，以便在发送信号之前防止信号的不必要传递。此外，sigsend还会在发送信号时记录有关信号的一些信息，例如信号的发送时间和发送者的身份信息。

在go语言中，sigsend函数是为了解决在处理信号时，有些信号可能会被阻塞或者被忽略的问题。通过使用sigsend函数，我们可以确保指定的进程或进程组能够接收到信号，从而确保程序的正确性和稳定性。sigsend还可以在后台发送信号，并且可以处理来自多个goroutine的信号，并在必要时合并它们，从而减少可能出现的竞态条件和死锁等问题。

总之，sigsend函数是一个非常有用的工具，它可以用于向指定的进程或进程组发送信号，并处理多个goroutine之间的信号竞争等问题。它是运行时系统中的一个关键部分，确保程序在处理信号时的正确性和稳定性，从而保证程序的正常运行和优化。



### signal_recv

signal_recv函数主要的作用是接收并处理来自内核的信号。具体来说，它会不断地从sigincoming队列中读取信号，然后根据信号的类型采取不同的行动。对于大部分信号，该函数会将其添加到sigqueue中等待进一步处理，但是对于一些特殊的信号（例如SIGSEGV和SIGILL），该函数会触发运行时的panic或者终止程序。在处理完所有的信号之后，该函数还会根据全局变量sigdisable和sigdisable_all来设置信号处理器的状态，从而影响后续信号的处理。

总的来说，signal_recv函数是runtime包中非常重要的一个函数，它承担着接收和处理信号的核心功能，直接影响着程序的稳定性和正确性。



### signalWaitUntilIdle

signalWaitUntilIdle函数位于runtime/sigqueue.go文件中，主要作用是等待所有正在进行的Goroutine都进入空闲状态，然后再执行一些后续的操作。

在这个函数中，首先会进入一个for循环，不断地检查所有的Goroutine是否都处于空闲状态，如果某个Goroutine正在执行，则等待一段时间后再检查。这个等待时间是通过time.Sleep()函数实现的，每次等待时间会逐渐增加，最长不会超过5秒钟。

当所有Goroutine都处于空闲状态时，程序就会执行一些后续的操作，例如向所有Goroutine发送一个信号，让它们停止运行。这些后续操作的逻辑是在signalAfterWait函数中实现的。

signalWaitUntilIdle函数的主要作用是保证程序的并发安全性。如果在Goroutine还没有完全退出的情况下，就强制停止它们，可能会导致一些不可预测的行为。通过等待所有Goroutine都处于空闲状态后再执行后续操作，可以避免这种情况的发生，确保程序的安全性和正确性。



### signal_enable

signal_enable函数是Golang运行时(runtime)中sigqueue.go文件中的一个函数，其主要作用是启用指定的信号，使其可以被接收器所响应。

具体来说，signal_enable函数接收一个整型参数signal，代表需要启用的信号。在函数内部，首先通过调用sigdisable信号禁用函数来禁用所有信号。然后调用sigaction系统调用来为指定的信号安装一个新的信号处理程序，并将旧的信号处理程序保存在sigtable中。

最后，初始化时已经被禁用的信号会通过循环再次发送给进程，以确保新的信号处理程序被安装并马上启用。

总之，signal_enable函数的主要作用是安装并启用指定信号的新的处理程序，以便被接受器所响应。它是Golang运行时中的一个重要函数，负责处理系统信号并通知其他线程来相应进行后续处理。



### signal_disable

在Go语言中，信号处理程序是非常重要的一部分。信号处理程序的目的是在程序运行时接收和处理来自操作系统的异步事件（例如键盘中断、控制+C等）。Go语言的运行时包中的sigqueue.go文件中，signal_disable()函数的主要作用是禁止信号的接收和处理。下面是signal_disable()函数的详细介绍：

函数签名：func signal_disable()

函数作用：

signal_disable()函数的主要作用是暂时禁止处理信号。该函数用于在goroutine（Go协程）开始执行一些临界区代码之前，将当前goroutine的信号处理程序暂时禁用。这样做的目的是为了防止信号在临界区代码中间发生不可预料的行为。

函数实现：

signal_disable()函数的实现非常简单，只需调用Go语言运行时包中的sigdisable()函数即可。该函数将当前goroutine的信号掩码设置为所有信号都被阻塞。

总结：

signal_disable()函数是Go语言中的信号处理程序的重要部分之一。它的作用是在goroutine开始执行临界区代码之前，暂时禁用当前goroutine的信号处理程序，以防止信号在临界区代码中间发生不可预料的行为。该函数的实现非常简单，只需调用Go语言运行时包中的sigdisable()函数即可。



### signal_ignore

signal_ignore函数是在runtime包的sigqueue.go文件中的一个私有函数。它的作用是设置一组处理器信号的处理方式，将它们设为SIG_IGN表示忽略。

具体来说，signal_ignore函数会将进程对指定信号的处理方式设置为SIG_IGN，即忽略该信号。这个函数的参数为一个信号的集合，也就是一个类似于signalSet的数据结构。在实际使用中，signal_ignore函数会在某些情况下用来处理一些不需要处理的信号，比如SIGQUIT、SIGPIPE等，将其忽略以避免对程序运行造成影响。

此外，signal_ignore函数还有一个重要的作用，就是为了兼容不同的平台和系统。在不同的操作系统或处理器架构下，信号的处理方式可能会有所不同。通过使用signal_ignore函数，可以统一处理器对信号的处理方式，确保程序在不同平台下都能够正常运行。

总之，signal_ignore函数是一个用来设置处理器信号的处理方式的函数，它能够忽略不需要处理的信号，并确保程序能够在不同平台下正常运行。



### sigInitIgnored

sigInitIgnored函数是一个初始化函数，它的主要作用是将SIGPIPE、SIGURG和SIGCONT三个信号的处理函数设置为忽略。

SIGPIPE信号是在向一个已经关闭的socket或管道写入数据时发出的。在这种情况下，收到SIGPIPE信号默认的处理方式是终止进程。但是，在某些情况下，我们可能想要忽略这个信号，例如在使用网络编程时可能会遇到这种情况。因此，sigInitIgnored函数将SIGPIPE信号的处理函数设置为SIG_IGN，即忽略该信号。

SIGURG信号是在进程接收到一个带外数据时发出的信号。在默认情况下，该信号不会被忽略，而是打印一条错误消息并终止进程。但是，在某些情况下，我们可能不想处理该信号，因此sigInitIgnored函数将SIGURG信号的处理函数设置为SIG_IGN，即忽略该信号。

SIGCONT信号是在进程从停止状态中恢复执行时发出的。在默认情况下，该信号不会被忽略，而是使进程从停止状态中恢复，并从停止的位置继续执行。但是，在某些情况下，我们可能想要忽略该信号，例如当我们需要手动控制进程的状态时。因此，sigInitIgnored函数将SIGCONT信号的处理函数设置为SIG_IGN，即忽略该信号。

总之，sigInitIgnored函数的作用是将一些默认情况下不需要处理的信号的处理函数设置为SIG_IGN，以实现将其忽略的目的。



### signal_ignored

在sigqueue.go文件中的signal_ignored函数是用于检查一个信号是否被忽略的函数。在Unix系统中，我们可以使用sigaction函数来控制信号处理程序的行为，包括处理程序的行为和信号的处理方式。在某些情况下，我们可能希望将某些信号忽略或屏蔽。在这种情况下，signal_ignored函数非常有用。

signal_ignored函数接受一个signal值，并返回一个布尔值来指示该信号是否被忽略。该函数首先获取与指定信号相关联的当前处理程序，然后检查处理程序是否为SIG_IGN，如果是，则返回true，表示该信号被忽略。如果处理程序不是SIG_IGN，signal_ignored函数将返回false，表示该信号未被忽略。

这个函数常用于处理信号的过程中，以决定是否要跳过某些信号。例如，在处理SIGCHLD信号时，我们可以使用signal_ignored函数来判断该信号是否被忽略，如果被忽略，我们可以跳过该信号的处理过程。这样可以避免在SIGCHLD信号频繁发生时，程序的性能受到影响。



