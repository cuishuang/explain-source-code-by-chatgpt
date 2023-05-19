# File: defs_linux_loong64.go

defs_linux_loong64.go是Go语言运行时系统中的一个文件，它的作用是为了支持龙芯64位（Loong64）架构的操作系统。

在Go语言中，每种操作系统架构都需要定义一些常量和类型，以便在运行时对不同的硬件和操作系统进行适配。特别是在跨平台开发中，操作系统和硬件架构的不同会对代码的编写和运行带来很大的影响。

因此，在Go语言运行时系统中，每种操作系统架构都需要有一个defs_操作系统架构.go的文件来定义一些常量和类型。defs_linux_loong64.go就是为了支持Loong64架构的Linux操作系统而设计的。

其中，defs_linux_loong64.go定义了一些系统调用的常量和类型，如Syscall、RawSockaddr、SockaddrInet、SockaddrInet6、SockaddrLll等，这些常量和类型用于在Loong64架构的Linux操作系统上执行Linux系统调用。

总之，defs_linux_loong64.go的作用是为了让Go语言在Loong64架构的Linux操作系统上顺利地运行，以充分发挥这种硬件和操作系统的特性优势。




---

### Structs:

### timespec

timespec是一个用于表示时间的结构体，常用于操作系统中。在Go语言的运行时(runtime)中，defs_linux_loong64.go这个文件中定义了一个timespec结构体，用于在Linux龙芯64位平台上表示时间。

具体来说，这个结构体有两个字段：tv_sec和tv_nsec。tv_sec是一个long int类型的字段，表示秒数；tv_nsec是一个long int类型的字段，表示纳秒数。因此，一个完整的timespec结构体实例表示了一个精确到纳秒的时间。

在Go语言的运行时中，这个结构体常用于实现时间相关的函数和操作，例如定时器（timer）和休眠（sleep）操作。由于在Linux龙芯64位平台上表示时间的方式与其他平台可能有所不同，因此需要特别定制一个defs_linux_loong64.go文件，定义适用于该平台的时间结构体和相关的操作函数。



### timeval

在defs_linux_loong64.go文件中，timeval结构体是用来表示时间的结构体。在操作系统中，由于需要处理很多时间相关的任务，因此定义了一些结构体或函数来方便处理时间。

timeval结构体定义如下：

```
type timeval struct {
    tv_sec  int64
    tv_usec int64
}
```

其中，tv_sec表示秒数，tv_usec表示微秒数，即秒数的小数部分（1秒=1000000微秒）。这个结构体主要用于一些系统调用中，例如gettimeofday函数会返回当前的时间信息，并将其存储到timeval结构体中。

在Go语言的runtime中，也会用到timeval结构体，例如在time包中，通过cputicks函数获取处理器时钟计数器的值时，会将计数器的值存储到timeval结构体中。这个时钟计数器的值可以用来计算程序的运行时间，或者用来实现一些计时功能。

总之，timeval结构体是一个通用的表示时间的结构体，在操作系统中或者其他需要处理时间的场景中都有广泛的应用。



### itimerspec

在go/src/runtime中的defs_linux_loong64.go文件中，itimerspec是一个结构体，用于定义Linux系统上使用的时间结构。具体来说，itimerspec结构体包含了两个成员变量：it_interval和it_value。

it_interval表示在定时器事件到达后，为了支持周期性操作，系统应该再次设置定时器所需的时间间隔。而it_value表示定时器应该到期的相对时间。

在Go语言中，itimerspec结构体的主要作用是用于定时器（timer）的操作和管理。定时器是一种操作系统提供的基本机制，用于交互式的操作和异步事件处理。通过使用itimerspec结构体，程序可以设置定时器的触发时间和间隔，并在定时器到期时执行特定的操作。

总之，itimerspec结构体是Go语言运行时环境下的一个重要组成部分，用于管理定时器的触发时间和周期性操作，从而支持程序的异步事件处理和交互式操作。



### itimerval

在Go语言的runtime包中，defs_linux_loong64.go文件定义了一些与操作系统底层相关的常量、类型、函数等内容，其中itimerval结构体是用来表示时间间隔的。

具体来说，itimerval结构体包含了两个成员：it_interval和it_value。其中，it_interval表示定时器的周期性间隔，也就是每隔多久触发一次定时器。而it_value表示定时器的启动时间，即定时器第一次触发的时间。

当程序需要使用定时器时，可以利用itimerval结构体和相关的系统调用来实现。具体步骤是：首先创建一个itimerval结构体，并设置好它的成员值；然后调用系统调用setitimer()或timer_settime()来启动定时器；最后在程序中等待定时器触发，并处理相应的事件。

总之，itimerval结构体是在Go语言的runtime包中用来表示时间间隔的结构体，它在实现定时器等功能时非常重要。



### sigeventFields

在Go语言运行时的defs_linux_loong64.go文件中，sigeventFields结构体是用于封装sigevent事件结构体的字段，以便在使用时更加方便和直观。sigeventFields结构体中封装的字段与sigevent事件结构体中的字段一一对应，包括sigevent的通知方式、通知信号、私有数据等信息。

sigeventFields结构体的作用是为sigevent事件结构体提供一种更加友好的封装方式，使得程序员在使用sigevent时更容易理解和操作。此外，由于Go语言是跨平台的，因此该结构体还可以方便地适配不同的操作系统和硬件平台，以提高代码的可移植性和兼容性。

总的来说，sigeventFields结构体在Go语言的运行时中具有非常重要的作用，是Go语言运行时处理信号事件时的一个重要组成部分。



### sigevent

在Go语言的Runtime实现中，defs_linux_loong64.go文件定义了Linux系统下的一些常量、类型和系统调用。其中，sigevent结构体定义如下：

```
type sigevent struct {
	value           [4]byte
	signo           int32
	sigev_notify    int32
	sigev_tid       uint64
	sigev_un        [56]byte
}
```

sigevent结构体的主要作用是在Linux系统下进行信号通知（Signal Notification）。当某个事件发生时，操作系统会发送一个信号通知，告知进程该事件已经发生。而sigevent结构体是用来描述这个信号通知的。其各个字段的含义如下：

- value：一个长度为4的byte数组，可以用来传递一些额外的信息。
- signo：表示发送的信号编号。
- sigev_notify：表示信号通知的方式，可以是SIGEV_NONE、SIGEV_SIGNAL、SIGEV_THREAD、SIGEV_THREAD_ID四种之一。
- sigev_tid：如果sigev_notify字段的值为SIGEV_THREAD_ID，那么该字段表示接收信号通知的线程ID。
- sigev_un：一个长度为56的byte数组，可以用来传递一些额外的信息。

以上就是sigevent结构体的详细介绍，它在Go语言的Runtime中用来描述Linux系统下的信号通知。



### sigactiont

在Go语言的运行时中，defs_linux_loong64.go文件定义了一些系统相关的常量，变量和数据结构，其中包括sigactiont结构体。sigactiont结构体是用来描述信号处理的参数的，在捕获信号时需要使用该参数。

具体来说，sigactiont结构体用于指定信号的处理方式，包括信号的处理函数、信号的掩码、处理标志等。它包括以下字段：

1. sa_handler：表示处理该信号的函数或动作；
2. sa_flags：表示处理信号的标志，如SA_RESTART表示当进程被信号中断时，系统会尝试重新启动被中断的系统调用；
3. sa_mask：表示在信号处理函数执行期间新的信号会被阻塞的信号集合；
4. sa_sigaction：表示使用sa_sigaction指定的函数处理信号，该函数的原型为void func(int sig, siginfo_t *info, void *context)。

通过使用sigactiont结构体，程序可以指定信号的处理方式，保证信号在处理过程中不会被其他信号打断，从而保证程序的正常运行。



### siginfoFields

在Go语言中，由于需要处理信号来管理进程状态、处理错误、中断执行等，因此需要对操作系统信号进行处理。在Linux系统上，处理信号需要使用sigaction()函数，并使用siginfo_t结构体来传递信号处理的相关信息。为了便于处理信号，Go语言在运行时库中定义了siginfoFields结构体，用于存储siginfo_t结构体中的各个字段。具体来说，siginfoFields结构体包含了如下字段：

- sig: 信号编号
- errno: 与信号相关的错误码
- code: 与信号相关的一个值，用于进一步定义信号的含义
- pid: 发送信号进程的ID
- uid: 发送信号进程的用户ID
- status: 发送信号的线程状态（仅在SIGNAL_CLD时使用）
- addr: 与信号相关的地址（如SIGSEGV信号的地址）
- band: 文件描述符中发生事件的相关位图（如SIGPOLL信号的相关位图）
- value: 与信号相关的一个值，用于传递额外信息（如SIGFPE信号中的浮点异常类型）

通过将siginfo_t中的各个字段存储在siginfoFields结构体中，Go语言可以更方便、安全地处理信号，并且可以在不同的平台上保持一致性。



### siginfo

siginfo结构体用于存储信号的详细信息，包括信号的类型、地址、发送者、CPU等。这个结构体在runtime中的定义与操作主要是为了处理在程序运行过程中发生的各种信号，例如SIGSEGV（Segmentation Fault，段错误）、SIGBUS（总线错误）等。

在defs_linux_loong64.go中，siginfo结构体的定义如下：

type siginfo struct {
    signo int32
    errno int32
    code  int32
    pad   [1]byte // TODO: use siginfo_t.__pad field instead (when present)
    _data [32]byte
}

其中，signo表示信号的类型，errno表示错误码，code表示附加的信号代码信息。在runtime/sys_linux_loong64.go中，还定义了一个从信号处理函数中解析出siginfo结构体的函数：

func sigaction(sig uint32, new, old *sigactiont) int32 {
    //...
    var info siginfo
    //...
    if flags&(SA_SIGINFO|SA_ONSTACK) != 0 {
        if flags&SA_SIGINFO != 0 {
            // Extract siginfo from third arg to handler.
            // Must be an *os.Signal or *os.ProcessState.
            osSig, hasOsSig := args[2].(*os.Signal)
            procState, hasProc := args[2].(*os.ProcessState)
            if !hasOsSig && !hasProc {
                return EINVAL
            }
            t.addr = add
            if hasOsSig {
                info.fillin(osSig, sig)
            } else {
                info.fillin(nil, sig)
            }
        } else { 
            //...
        }
    } else { 
        //...
    }
    //...
}

从上面的代码可以看出，在处理信号的过程中，程序通过调用sigaction函数来获取siginfo结构体，其中fillin方法根据传入的参数填充siginfo。这样可以方便程序对于发生信号时的场景作出相应的处理，例如捕获SIGSEGV信号时，可以通过解析siginfo结构体来获取程序崩溃时的地址信息，从而更好地分析和调试问题。

总之，siginfo结构体在runtime中的作用是为程序处理各种信号提供了必要的信息，使得程序可以更加智能地在信号处理过程中做出反应，有利于程序的正常运行和调试。



### usigset

在 Linux 系统上，sigset_t 用于表示一个信号集合，它是一个包含所有信号的数据结构。当一个进程想要阻塞某些信号时，它可以将信号添加到自己的信号集合中。

在 Go 的 runtime 库中，defs_linux_loong64.go 文件中定义了一个 usigset 结构体，它是对 sigset_t 的简单封装。usigset 结构体包含两个字段：

- mask：一个 uint64 类型的值，表示一个信号集合的位图。这个位图中的每一位都对应一个信号的编号，如果某个信号对应的位为 1，则表示该信号被阻塞了，否则表示该信号未被阻塞。
- flags：一个 int32 类型的值，表示该信号集合的标志。目前定义了两个标志：SIG_BLOCK 和 SIG_UNBLOCK。SIG_BLOCK 标志表示将该信号集合中的信号阻塞，SIG_UNBLOCK 标志则表示解除该信号集合中的信号阻塞。

usigset 结构体提供了一些方法来操作信号集合。例如，usigset 的 set 方法可以将一个指定的信号添加到信号集合中；clear 方法可以将信号集合中的某个信号清除。此外，usigset 还提供了和 sigprocmask 系统调用类似的 mask 和 unmask 方法，用于阻塞和解除阻塞信号集合中的信号。

在 Go 库中，usigset 主要被用于实现信号处理。Go 程序可以使用信号处理函数来处理各种信号，而这些信号可能会在运行时发生。因此，Go 程序需要阻塞某些信号，以避免信号处理函数在执行时被打断。usigset 提供了一种简单的方式来阻塞和解除阻塞信号的流程。



### stackt

在 Go 语言中，每个 goroutine 都会有一个对应的 Go 栈（goroutine stack），用于保存函数调用、局部变量等信息。在 Linux 系统上，这个 goroutine stack 的实现使用了 pthread 库中的 `pthread_attr_setstack` 和 `pthread_attr_getstack` 函数来分配和管理内存。

在 `defs_linux_loong64.go` 文件中定义了 `stackt` 结构体，用于描述 goroutine stack 的信息，包括 stack 的起始地址、大小、保护页个数等。

```go
type stackt struct {
	ss_sp uintptr    // stack bottom
	ss_size uintptr  // stack size
	ss_flags int32   // flags
}
```

其中，`ss_sp` 表示 stack 的起始地址，`ss_size` 表示 stack 的大小，`ss_flags` 表示 stack 的标志位，用于描述 stack 的状态信息，比如是否需要保护页、是否可以执行等。

在 Go 语言中，当一个 goroutine 被创建时，系统会自动为其分配一个对应的 goroutine stack，并将其存储在调用方的栈中。当 goroutine 结束时，其对应的 goroutine stack 会被释放回系统内存池。

因此，在 `defs_linux_loong64.go` 文件中，`stackt` 结构体的作用是描述 goroutine stack 的信息，便于系统或调度器对其进行管理和回收。



### sigcontext

在 Linux 操作系统中，当程序在运行过程中发生信号时，内核会将进程的执行状态保存到一个称为信号上下文（sigcontext）的结构体中。其中包括了程序当前 CPU 寄存器的状态、栈、代码指针等信息。

在 Go 语言的 runtime 包中，defs_linux_loong64.go 文件中的 sigcontext 结构体用于保存进程在接收到中断信号时的 CPU 寄存器、栈及指令等上下文信息。Go 程序在处理信号时，需要依赖于该结构体中保存的信息来确定进程的执行状态，并且从该结构体中获取信息，以便程序能够正确地处理信号。

该结构体的详细成员信息如下：

```go
type sigcontext struct {
    sc_regmask   uint64   // CPU寄存器标志
    sc_mask      uint64   // 信号掩码
    sc_acx       uint64   // X寄存器
    sc_scratch   uint64   // Scratch寄存器
    sc_pc        uint64   // 程序计数器（PC）
    sc_cpu       uint64   // CPU ID
    sc_npc       uint64   // 下一个程序计数器值
    sc_fpcsr     uint64   // 浮点控制状态寄存器（FP control and status register, FCSR）
    sc_used      uint32   // 该结构体被使用的标记位
    // 省略了一些其他成员
    sc_masked    [2]uint64 // 信号标志掩码
    sc_trapno    uint64   // 中断的trap号（与x86 architecture相似）
    sc_err       uint64   // 错误码（异常产生时）
}
```

因此，在 Go 语言的运行时（runtime）中，sigcontext 结构体提供了处理信号时所需要的上下文信息，这些信息是处理信号时必不可少的。



### ucontext

在Go语言中，ucontext结构体定义在defs_linux_loong64.go文件中，是用来表示线程的上下文信息的结构体。

具体来说，ucontext结构体记录了一个线程的当前状态，包括寄存器的值、处理器状态、信号信息等等。当线程因为某种原因（例如收到信号、申请系统调用等）被挂起时，系统可以通过保存和恢复线程的ucontext结构体来实现之前状态的记忆和还原，以保证线程能够从之前的状态继续执行。

在Linux系统中，ucontext结构体被广泛地应用于线程同步、错误处理和信号处理等方面。Go语言的运行时系统也充分利用了ucontext结构体的特性，通过保存和恢复线程的上下文信息来支持协程的切换、异步信号处理和错误恢复等功能。



## Functions:

### setNsec

在Go语言中，时间单位是纳秒（nanosecond，ns），setNsec函数用于将纳秒数设置为时间结构体指定时间的纳秒值。

该函数的定义如下：

```
func (t *Timespec) setNsec(nsec int64)
```

其中，Timespec结构体是一个封装了Linux系统调用timespec类型的时间结构体，具体定义如下：

```
type Timespec struct {
    Sec  int64 // 储存秒数
    Nsec int64 // 储存纳秒数
}
```

在实际使用中，可以通过setNsec函数将纳秒数设置到时间结构体的Nsec字段中，从而表示指定时间的纳秒值。例如：

```
var t Timespec
t.setNsec(123456789) // 设置t时间为1970-01-01 00:00:00.123456789 +0000 UTC
```

总之，setNsec函数是一个非常简单但功能实用的函数，它可以方便地将纳秒数设置到时间结构体中，使我们可以更加灵活地处理时间。



### set_usec

在go/src/runtime中的defs_linux_loong64.go文件中，set_usec函数的作用是设置操作系统级别的计时器分辨率（单位为微秒）。该函数通过调用系统调用setitimer设置用户空间定时器，以便在应用程序中进行高精度计时。

具体来说，set_usec函数以一个整数作为参数，将其转换为以微秒为单位的时间间隔，然后将该值分配给定时器的时间间隔。定时器上的信号被捕获并路由到go-routine进行处理，以处理需要高精度计时的任务，例如垃圾回收和调度。

总之，set_usec函数是操作系统级别的高精度计时器的设置函数，它允许Go程序在进行需要高精度计时的任务时获得更好的性能和精度。



