# File: stubs_arm64.go

stubs_arm64.go文件是用于Go语言运行时在ARM64芯片架构上的桩代码文件，用于实现Go语言运行时对于特定硬件平台的底层操作和实现，该文件中包含了一些桩代码，用于实现Go语言运行时中无法使用标准库实现的一些底层操作。

具体而言，stubs_arm64.go文件中定义了一些桩代码函数，用于实现Go语言运行时在ARM64芯片架构上的一些底层操作，如获取CPU周期计数器、设置以及清除某个地址处的内存、获取当前执行指令的地址、暂停程序执行等底层操作。这些底层操作是必须的，因为Go语言需要保证在不同平台下的可移植性以及实现高效的垃圾回收、并发等机制，这需要底层操作的支持。

除此之外，stubs_arm64.go文件还包含一些与调试有关的桩代码函数，用于在进行调试时更容易诊断和分析程序中的问题，如打印栈帧、检查goroutine等，这些桩代码有助于开发人员更高效地进行调试和错误定位。

因此，stubs_arm64.go文件是Go语言运行时实现在ARM64芯片架构上的一部分，是Go语言保证底层可移植性、实现高效垃圾回收和并发机制、进行调试和错误定位的必要组成部分。

## Functions:

### load_g

load_g函数是用来加载当前Goroutine的函数。在ARM64架构中，Goroutine的结构体中有一个指针字段m，指向当前Goroutine所在的M（machine，代表一个线程）

load_g函数通过ASM汇编实现，其主要步骤如下： 

1. 内联汇编指令将当前指令地址存储到LR寄存器中； 

2. 通过内联汇编指令将当前M的地址存储到x18寄存器中； 

3. 调用runtime·getg函数，将当前Goroutine的指针存储在x19寄存器中。

通过以上步骤，load_g函数就能够加载当前Goroutine的指针了。这个指针会被存储在g变量中，方便后续程序使用。

load_g函数在不同的地方有不同的作用。例如，在调用协程切换的时候，需要使用load_g函数获取当前Goroutine的指针，以便在调度器中进行切换。在实现锁操作时，也需要使用load_g函数获取当前Goroutine的指针，以便在解锁时将其设置为等待状态的Goroutine的指针。



### save_g

save_g是一个用于ARM64平台的汇编函数，主要作用是将当前goroutine（即当前线程）的上下文保存到栈帧中，以便在稍后恢复时使用。具体来说，函数的实现会将以下内容保存到栈帧中：

1. 寄存器：将所有的通用寄存器（R0-R30）以及特殊寄存器（SP, PC, PSTATE）的值保存到栈帧中。这样做可以确保在恢复goroutine时，所有的寄存器都能够保持之前的状态，从而不会丢失任何重要的信息。

2. 栈指针：将栈指针（SP）的值也保存到栈帧中。在恢复goroutine时，该栈指针将被用来恢复栈中的数据，保证函数能够继续执行下去。

总之，save_g的主要作用是将当前goroutine的状态保存到栈帧中，以便在后续恢复时能够继续执行。这样做可以保证程序的正确性和稳定性，从而提高系统的可靠性。



### asmcgocall_no_g

asmcgocall_no_g is a function in stubs_arm64.go file in the go/src/runtime package. This function is used to call a Go function from an assembly routine or a foreign function written in C or C++. 

When a Go program performs an external call or a system call, it needs to switch from the Go runtime to the external function, and then switch back to the Go runtime when the external function is finished. This switch is handled by the Go runtime using a technique called "CGO" (C Go). 

In some cases, the external function may not need to access the Go runtime, for example when it is a simple function that performs some computation and then returns a result. In such cases, a special CGO function can be used to avoid the overhead of switching between the Go runtime and the external function. This function is called "asmcgocall_no_g".

The "asmcgocall_no_g" function is a special CGO function that can be used to call a Go function without the overhead of switching to the Go runtime. This function takes a function pointer to a Go function, as well as a list of arguments to pass to the function, and calls the Go function directly from C or assembly code.

Overall, the "asmcgocall_no_g" function improves the performance of Go programs that call external functions in situations where the external function does not need to access the Go runtime.



### emptyfunc

在Go语言中，emptyfunc是一个空的函数（空函数），它没有任何操作或返回值。在stubs_arm64.go文件中，emptyfunc被用作ARM64架构上的汇编指令的占位符，以便在运行时进行填充。

在ARM64架构上，空指令（NOP）指令的长度为4个字节，因此在实现一些汇编操作时，需要插入一些NOP占位符来占用指令的位置。在这种情况下，emptyfunc可以用作一个标记，以便在编写汇编代码时插入所需的NOP指令。

在Go运行时中，emptyfunc的实现方式不需要任何参数或返回值。这简化了代码的实现，并避免了在处理汇编代码的时候出现冲突和错误。因此，emptyfunc在ARM64架构上的应用非常广泛，是Go语言成功运行的重要组成部分。



### spillArgs

在Go语言中，spillArgs函数的作用是将函数调用中的参数存储到栈上。在ARM64处理器架构中，函数调用时的参数是通过寄存器传递的。但是，如果参数太多，无法全部通过寄存器传递时，就需要通过栈来传递剩余的参数。spillArgs函数会将剩余的参数从寄存器中取出并存储到栈上，以供后续函数调用使用。

具体来说，spillArgs函数会将参数从通用寄存器（X0-X7）和浮点寄存器（V0-V7）中取出，并按照从右到左的顺序将其存储到栈上。这个过程需要注意的是，参数的存储顺序和寄存器的编号是相反的，因为ARM64处理器架构中的传参顺序是从右到左的。

在spillArgs函数完成后，函数调用就可以从栈上读取参数并使用它们了。这个过程结束后，函数将调用spillArgs函数的下一条指令的返回地址存储到栈上，最后将栈指针调整到新的栈顶位置，函数调用就可以继续执行了。



### unspillArgs

在ARM64架构中，函数参数通常被传递到寄存器中。但是，在某些情况下，寄存器可能会不够用来存储所有的参数，这时候需要将一些参数从堆栈中取出存储到寄存器中。这个过程称为“溢出”（Spill）。但是，在函数退出之前，这些参数需要被恢复到堆栈中，这个过程称为“反溢出”（Unspill）。

unspillArgs是在堆栈上分配的参数的恢复函数。在函数退出时，具有未使用的参数或备份参数的激活帧必须被恢复。该函数从堆栈中读取参数并将其放入相应的寄存器中，以便后续使用。这个函数的作用是恢复可能被溢出到堆栈中的函数参数到寄存器中，方便程序使用。它被用在runtime、syscall和Go汇编语言的代码中。



### getcallerfp

getcallerfp函数是用于获取调用者的帧指针的。在函数调用时，调用栈中每个函数的栈帧包含了函数的参数、局部变量、返回值和调用者的帧指针等信息。由于对于函数调用者而言，它需要在当前函数返回时，回到调用该函数的代码处继续执行。因此需要知道调用者的帧指针，以便能够获取它的局部变量和参数，并将控制流转回调用者处。

在ARMv8架构下，调用栈的结构和具体的实现有所不同。在栈底有一个四字节的标志位，表示该栈是返回栈还是答复栈。栈的剩余部分是由不同的函数调用所创建的，其中每个栈帧由16字节对齐的帧大小和该栈帧的返回地址组成。除此之外，每个栈帧还包含了保存和恢复寄存器的区域，包括参数、局部变量和frame pointer。

在getcallerfp函数中，它通过读取当前函数的帧指针、返回地址和参数区中的数据，来定位到调用者的帧指针位置并返回。这段代码通常在垃圾收集器等底层运行时库中使用，并且也用于实现Go语言中的defer机制。通过获取调用者的帧指针，Go语言的defer机制可以在函数返回时，正确地执行相应的defer语句。



