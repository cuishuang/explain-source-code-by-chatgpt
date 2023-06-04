# File: stubs_mipsx.go

stubs_mipsx.go是Go语言运行时的一个文件，它的作用是在MIPS架构上定义一些系统调用的存根函数。这些函数在运行时被动态链接器使用，以便Go程序可以调用底层操作系统提供的系统功能。

MIPS架构是一种常用于嵌入式系统的处理器架构，所以Go语言为MIPS架构提供了特定的运行时支持。stubs_mipsx.go文件中定义的存根函数（如open、fstat、write等）会在Go程序调用相关系统调用时被调用，以便Go程序可以在MIPS架构上实现与底层操作系统的交互与功能使用。

总的来说，stubs_mipsx.go文件在Go语言运行时中扮演了重要的角色，它实现了系统调用的存根函数，为Go语言在MIPS架构上构建可靠的应用程序提供了支持。

## Functions:

### load_g

文件stubs_mipsx.go中的load_g函数主要是用于在MIPS处理器上加载goroutine（G）的寄存器上下文的。G代表运行时系统对于应用程序所维护的执行上下文，其中包括程序计数器（PC）、栈指针（SP）、寄存器和其他与当前执行过程相关的信息。

load_g函数的作用是将G的寄存器上下文保存在函数栈中，并将它们存储在当前运行的goroutine的G结构中。具体来说，它将G的堆栈指针（stack pointer）保存在SP寄存器中，并将G的程序计数器（program counter）保存在RA（return address）寄存器中。然后，该函数使用汇编指令MEMCPY，在当前运行的goroutine的G结构中的寄存器上下文中存储G的寄存器上下文，该函数的最后一行使用JALr从存储G的R寄存器的函数返回。

G的寄存器上下文在切换goroutine时非常重要。每当一个goroutine调用它的另外一个goroutine或让出CPU时，运行时系统必须将当前正在执行的goroutine的寄存器上下文保存起来，以便在将来从该点继续执行。然后，它将要切换到的goroutine的寄存器上下文加载到寄存器中。

因此，load_g函数是运行时系统中的一个重要函数，用于在不同的goroutine之间进行寄存器上下文的切换，以确保正确的执行顺序。



### save_g

save_g是在MIPS32架构下实现的一个函数，它的主要作用是保存当前goroutine的寄存器状态。

在Go语言中，goroutine是轻量级线程的概念，每个goroutine都有自己的栈空间和寄存器状态。当goroutine被调度器切换出CPU时，需要将当前goroutine的寄存器状态保存到内存中，以便下次恢复时能够继续执行。

具体而言，save_g函数通过将当前寄存器状态保存到goroutine的g结构体中，实现了goroutine的切换，包括将尚未执行的指令保存到g结构体中，将寄存器状态保存到g结构体的栈空间中等操作。在下次调度该goroutine执行时，可以通过g结构体中保存的状态信息，恢复该goroutine到之前的执行状态。

需要注意的是，MIPS32架构下的save_g函数可能与其他架构下的实现方式不同，但其核心作用均为保存当前goroutine的寄存器状态，以支持goroutine之间的切换。



### getcallerfp

getcallerfp函数是一个用于在MIPSx架构的Go运行时中获取当前函数的调用者的帧指针（frame pointer）的函数。

在MIPSx架构中，函数调用的调用者需要将自己的帧指针（frame pointer）先保存到栈中，然后才能跳转到被调用函数的代码。因此，通过获取当前函数的调用者的帧指针，就可以找到当前函数的调用者，从而实现调用栈的回溯和调试等功能。

getcallerfp函数的具体实现如下：

```go
//go:noescape
//go:nosplit
func getcallerfp() uintptr {
	return uintptr(*(*uint64)(getcallerpc() + 16))
}
```

其中，getcallerpc是一个在MIPSx架构中获取当前函数的调用者的PC寄存器的函数，该函数在stubs_mipsx.go文件中也有实现。

实现原理：在MIPSx架构中，函数调用时，调用者需要将自己的帧指针先保存到栈中，然后跳转到被调用函数的代码。被调用函数在执行时，会将自己的帧指针保存到寄存器中，然后使用栈中保存的帧指针作为自己的上一级栈帧的帧指针。因此，通过当前函数的PC寄存器的地址加上偏移量16，可以找到当前函数的调用者的帧指针。



