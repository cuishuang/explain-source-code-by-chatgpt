# File: defs_linux.go

defs_linux.go这个文件是Go语言标准库中运行时runtime的一个子模块，主要负责定义和实现与Linux系统相关的底层操作。具体来说，它的作用包括：

1. 定义了运行时所需的各种类型和常量，如进程号、系统调用号、信号编号等等。这些类型和常量是运行时实现所必须的，而与系统相关，因此需要依据不同的操作系统分别定义。

2. 定义了Linux系统调用的封装函数。runtime在各种情况下需要与Linux内核进行系统调用交互，比如申请内存、释放内存、进行线程切换等等。为此，defs_linux.go中定义了一系列封装函数，如mmap、munmap、clone、epoll_create等等，以便runtime进行系统调用。

3. 定义了与Linux信号处理相关的函数。Linux系统中会向进程发送各种信号，如SIGINT、SIGTERM等等，用于通知进程执行某种操作或中断程序执行。defs_linux.go中定义了处理这些信号的一系列函数，如signal_recv、signal_enable、signal_disable等等。

总之，defs_linux.go这个文件是Go运行时系统在Linux平台上的一个重要组成部分，它通过定义操作系统相关的类型、常量和函数，为软件提供了与操作系统交互的底层基础，从而支持整个Go程序的正常执行。




---

### Structs:

### Sigset

Sigset结构体是用来表示Linux下的信号集的，可以理解为一个标记集合，用于标记哪些信号是处于阻塞状态的。在Linux系统中，进程可以通过信号与其他进程进行交流并处理异步事件，而Sigset结构体就是用于控制这些信号的。

Sigset结构体具体包含两个字段，分别是a和b，都是uint64类型的数字，每一位代表一个信号是否被阻塞。其中a字段表示64个信号中的第0-63个信号，b字段表示64个信号中的第64-127个信号。当一个信号被标记，表示该信号被阻塞，收到该信号时，进程不会立即处理，而是等待解除阻塞的时候再处理该信号。

Sigset结构体的作用主要有两个：一是用来控制信号的阻塞和解除阻塞，这在多线程环境下常用到；二是用于挂起和恢复对某些系统调用的中断，从而避免出现竞态条件或死锁。

在Go语言中，Sigset结构体常常被用来实现信号的处理和控制。例如，在runtime库中，Sigset结构体被用于控制Go程序在处理垃圾回收时阻塞所有的信号，以避免在执行垃圾回收的过程中被中断，从而保证程序的正确性和稳定性。



### Timespec

在Linux系统中，有一种叫做“文件描述符”的机制，用于唯一标识文件、网络连接等资源，这些资源都有一个特定的生命周期，需要在一定的时间内进行回收或者释放。对于这个问题，Linux系统提供了一种叫做“定时器”的机制，可以设置一个定时器，当时间到达时，系统会发出信号来提醒用户进程进行相应的操作，例如关闭文件、终止网络连接等。

在Go语言中，这个机制主要是由runtime包中的defs_linux.go文件中的Timespec结构体来完成。Timespec结构体是一个与Unix下的timespec结构体相对应的类型，用于描述时间戳和时间间隔，包括秒和纳秒级别的精度。在Go语言中，这个结构体被用来设置和读取定时器的时间间隔，以便于控制资源的使用。比如，在使用select语句时可以设置一个超时时间，如果在指定时间内没有接收到任何数据，就会执行默认操作或者执行超时操作，这就是利用定时器来完成的。

总之，Timespec结构体的作用主要是用于设置和读取定时器的时间间隔，以实现对资源的有效控制。



### Timeval

Timeval结构体在Linux系统中用于表示时间的秒数和微秒数。

该结构体在Go语言中的runtime包中被定义，并用于处理时间相关功能，如时间片的分配、定时器的管理等。在Go中，Timeval结构体被用于实现time.Ticker和time.Sleep等函数。

Timeval结构体的定义如下：

```go
type Timeval struct {
    Sec  int64
    Usec int64
}
```

其中，Sec表示秒数，Usec表示微秒数。这两个成员变量的数据类型都是int64，可以存储非常大的整数值。

在Linux系统中，该结构体也被广泛使用，用于表示时间值。比如，获取当前时间的系统调用gettimeofday就会返回一个包含Timeval结构体的结构体。在Linux系统编程中，通过Timeval结构体可以方便地进行时间的转换和计算。

总之，Timeval结构体在Go语言中的runtime包中和Linux系统编程中都有着重要的作用，用于表示时间值和实现时间相关的功能。



### Sigaction

Sigaction是一个结构体，用于设置和获取信号处理函数的属性，包括信号处理函数指针、信号掩码以及一些处理选项等。在Linux系统中，Sigaction主要用于管理信号的处理方式，可以通过Sigaction结构体来注册用户自定义的信号处理函数。当信号被触发时，系统会调用对应的信号处理函数来处理该信号的相关事件，从而对于程序的正确性和稳定性起到重要作用。

Sigaction结构体有以下成员：

- Sa_handler：信号处理函数的指针，需要注意的是，如果该值为SIG_IGN，那么对应信号的处理函数就会被忽略。
- Sa_mask：指定在处理该信号的期间要阻塞的信号集，即当前信号的处理函数运行期间不能受到该信号及其它信号的干扰。
- Sa_flags：用于指定一些处理选项，如SA_RESTART，即在信号处理函数返回后，操作系统会自动重启系统调用，以便避免系统调用在信号处理函数被中断时被结束。

通过Sigaction结构体，可以实现对信号的注册、处理和回收等操作，使得程序能够及时准确地响应信号，保障了程序的安全和可靠性。



### Siginfo

Siginfo是一个结构体，用于在Linux系统中传递有关信号的信息。它定义了与正在处理的信号有关的许多属性，例如：

1. si_signo：表示正在处理的信号的类型。

2. si_errno：如果信号是由某个错误引起的，则表示与此错误相关的errno。

3. si_code：进一步说明了信号类型和源的原因。

4. si_pid：发送信号的进程的PID。

5. si_uid：发送信号的用户ID。

6. si_status：表示与被终止进程有关的状态信息。

7. si_value：可以包含有关信号值的附加信息。

在Go语言中，Siginfo结构体在处理信号时很有用。例如，在执行signal.Notify函数时，当注册时指定了Syscall或者User类型的信号处理函数，会返回一个channel，当收到相应信号时将会被写入该channel，同时该channel的类型为os.Signal。当程序收到信号并在处理函数内部时，常常需要访问Siginfo结构体中的信息以了解信号来源，状态和其他相关信息等。



### Itimerspec

在Go语言的运行时中，defs_linux.go文件中的Itimerspec结构体主要用于表示定时器的超时时间。Itimerspec结构体是由两个Timespec结构体组成的，Timespec表示了秒和纳秒，用于指定相对于定时器开始的超时时间。

Itimerspec结构体的具体定义如下：

type Itimerspec struct {
    Interval Timespec // 过期时间间隔
    Value    Timespec // 初始值
}

其中，Interval表示定时器超时的时间间隔，Value表示定时器的初始值。当定时器到达超时时间时，将会向操作系统发送SIGALRM信号，告知应用程序进行相应处理。

在Go语言中，通过调用timer_create()函数创建定时器，并将Itimerspec结构体作为参数传入。当定时器超时时，操作系统将向应用程序发送SIGALRM信号，Go语言的运行时会将该信号转换为一个异常，从而触发异常处理流程。



### Itimerval

Itimerval是一个用于设置定时器的结构体。在Linux系统中，它由time.h头文件定义，常用于设置SIGALRM信号的处理函数的间隔时间。

Itimerval结构体有两个成员变量：

1. Interval：表示定时器的间隔时间，以微秒为单位（1秒=1,000,000微秒）。

2. Value：表示定时器启动后第一次到期的时间，也以微秒为单位。

通过设置这两个成员变量的值，可以实现定时器的精确控制。在Go语言中，可以使用syscall.Setitimer和syscall.Getitimer函数来设置和获取定时器的值。

在Go语言的运行时（runtime）中，defs_linux.go文件中定义了Itimerval结构体，以及一些与其相关的常量和系统调用函数，用于支持Go语言的定时器功能。它提供了底层的操作接口，使得Go语言可以更好地与Linux系统进行交互。



### Sigevent

Sigevent结构体是定义在go/src/runtime/defs_linux.go文件中的一个结构体，它主要用于在Linux系统下设置信号事件的参数。

Sigevent结构体中包含以下字段：

- Value: 事件值，一般是0。

- Signo: 设置的信号值，比如SIGUSR1、SIGUSR2等。

- Notify: 定义事件触发时的通知方式。可以取值为SIGEV_SIGNAL、SIGEV_NONE、SIGEV_THREAD等。其中，SIGEV_SIGNAL表示事件触发后向指定进程发送信号；SIGEV_NONE表示事件触发后不做任何处理；SIGEV_THREAD表示事件触发后执行指定线程中设定的函数。

- _Syscall0: 保留字段。

- _Syscall1: 保留字段。

- _Syscall2: 保留字段。

- Pad_cgo_0: 保留字段。

Sigevent结构体主要是用于在Linux系统下设置信号事件处理函数的参数，其中Notify字段确定了事件触发后的处理方式。该结构体常用于使用系统调用timer_create()创建定时器时的参数设置。例如，在Go语言中使用time包里的time.AfterFunc()函数来创建定时器时就会使用Sigevent结构体。

总之，Sigevent结构体的作用是在Linux系统下用于设置信号事件参数，以实现相应的事件处理。



