# File: stubs_loong64.go

stubs_loong64.go文件是Go语言运行时（runtime）的一个文件，用于充当Go语言在Loong64架构上的存根。在Loong64架构上运行Go程序需要调用特定的机器指令，但是Go语言本身并不直接支持这些指令。因此，在Loong64架构上运行Go程序需要利用一些不同于其他架构的实现方式，这就是stubs_loong64.go的作用。

具体而言，stubs_loong64.go文件定义了一些Loong64架构上需要用到的函数，这些函数实现了Go语言的运行时机制，包括栈调整、垃圾回收等。这些函数是Go语言的核心功能，也是其性能优秀的原因之一。

此外，stubs_loong64.go文件还实现了一些与平台相关的功能，比如对齐、分页等。这些函数在Go语言的运行时环境中非常重要，因为它们允许Go语言在Loong64架构上进行内存管理。

总之，stubs_loong64.go文件非常重要，是Go语言在Loong64架构上运行的必要组成部分，为Go语言提供了Loong64架构上的支持。

## Functions:

### load_g

load_g函数是用来加载当前goroutine的G结构体的函数。在Go语言中，每个goroutine都有一个对应的G结构体，其中包含了这个goroutine的状态和调度信息。load_g函数会在goroutine切换时被调用，用来加载下一个goroutine的G结构体。

load_g函数的主要作用是将当前线程状态的上下文切换到下一个goroutine的上下文。load_g函数首先获取下一个goroutine的G结构体指针，然后将这个指针存储到当前线程的TLS（线程本地存储）中，以便下一次线程调度时可以快速访问。load_g函数还会恢复下一个goroutine的栈指针和上下文，在切换后，线程将开始执行下一个goroutine的代码。

load_g函数的实现方式会与不同的系统架构和操作系统有关。在stubs_loong64.go这个文件中，load_g函数是针对loong64架构的具体实现。根据具体的架构和操作系统，不同的load_g函数实现方式可能会有所不同，但它们的核心作用都是相同的。



### save_g

save_g函数是在Go语言运行时中的一部分，它的作用是将当前的goroutine的上下文信息存储起来，方便后续恢复。

在Go语言中，goroutine是一种轻量级的线程实现。每个goroutine都有自己的栈、指令指针和一些额外的信息，如goroutine ID、锁状态等。当一个goroutine被抢占或阻塞时，需要将它的上下文信息存储下来，然后将CPU切换到其他goroutine上；当这个goroutine再次被唤醒时，需要将它之前保存的上下文信息恢复。

save_g函数的作用就是将当前goroutine的上下文信息存储到一个特定的数据结构中。具体来说，它会将当前goroutine的栈指针、寄存器内容和一些其他信息，如goroutine ID、锁状态等存储到一个结构体中。这个结构体称为G结构体，它是Go语言中goroutine运行时状态的核心数据结构。

G结构体的保存有助于实现goroutine的抢占和恢复机制。当一个goroutine被抢占或阻塞时，它的G结构体会被保存到一个运行时调度器的队列中。当这个goroutine再次被唤醒时，它的G结构体会被取出，上下文信息会被恢复，并且CPU会切换回这个goroutine的执行上下文中，继续执行。



### getcallerfp

函数getcallerfp用于获取调用当前函数的函数的帧指针（Frame Pointer, FP）。帧指针是一个指向函数栈帧的指针，栈帧是用于存储函数调用时的局部变量和参数的数据结构。

在函数调用过程中，每个函数都会为自己分配一个新的栈帧。栈帧中通常包括函数的参数、局部变量和返回地址等信息。通过帧指针，程序可以访问上一个栈帧中的信息，也可以返回到上一个函数。

在Go语言中，获取调用当前函数的函数的帧指针是非常常见的操作。它通常用于调试和错误处理等情况。在stubs_loong64.go文件中的getcallerfp函数中，底层采用了汇编代码实现获取当前函数的帧指针，这样可以更高效地实现该功能。



