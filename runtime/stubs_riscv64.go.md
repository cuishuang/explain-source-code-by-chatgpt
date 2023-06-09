# File: stubs_riscv64.go

stubs_riscv64.go是Go语言运行时的一个核心文件，它提供了对RISC-V 64位架构的本地二进制接口（ABI）的支持。该文件包含了一些被称为stub（存根）的函数，这些函数在编写Go语言代码时可能会被调用，但在Go语言运行时中并没有实现。这些存根函数提供了对于各种硬件的基本支持，包括内存管理，线程管理，调度和垃圾回收等等。

对于RISC-V 64位架构上的Go语言，开发人员可以通过直接或间接地调用stubs_riscv64.go中的函数来实现简化硬件抽象接口访问，从而方便实现Go语言程序并降低Go语言程序运行时的性能开销。

同时，stubs_riscv64.go文件还充当了Go语言与RISC-V 64位架构兼容性的桥梁。如果Go语言被编译为RISC-V 64位架构上的二进制文件，那么就需要用到stubs_riscv64.go文件中定义的存根函数，并且这些存根函数在运行时会进行链接和动态加载。

## Functions:

### load_g

load_g函数的作用是从寄存器中读取G（goroutine）指针，然后返回该指针的值。在RISC-V架构上，G指针存储在寄存器g的值中。

在Go程序中，每个goroutine都有其自己的G结构体。G结构体是一个协程的上下文，包括其状态、栈指针、计数器、调度信息等等。在Go程序中，每个G结构体都由运行时系统负责分配和管理。

load_g函数的作用就是为了让Go程序可以访问当前协程的G结构体。由于G指针存储在寄存器中，因此需要通过load_g函数来读取并返回它的值。

load_g函数是在运行时系统初始化的时候生成的。它的具体实现依赖于特定的架构和编译器。在RISC-V架构上，load_g函数的实现是通过寄存器g来完成的。



### save_g

save_g函数是用来保存当前goroutine的上下文，包括程序计数器，栈指针等信息，以便于在切换到其他goroutine时可以恢复当前goroutine的执行状态。具体来说，save_g函数会将当前的goroutine中的一些重要寄存器的值保存到goroutine内部存储和栈中，其中包括：

1. 保存X0到X30寄存器的值，在RISC-V架构中，这些寄存器用于存放通用数据，函数调用时会用来保存参数、局部变量和返回值等；
2. 保存X31寄存器的值，在RISC-V架构中，这个寄存器用来存放link寄存器，函数调用时会用来保存返回地址，以便在函数返回时跳转到正确的位置；
3. 保存FP、GP、TP、RA和SP寄存器的值，这些寄存器在RISC-V架构中分别用来存放帧指针、全局指针、线程指针、返回地址和栈指针。

由于goroutine在Go语言中是轻量级线程，因此需要在多个goroutine之间进行快速、高效的切换。save_g函数的作用就是在每次切换goroutine时保存当前goroutine的执行状态，以便下次继续从上次中断的位置继续执行。同时，save_g函数还可以避免不同goroutine之间的干扰，因为每个goroutine都会有自己的内存空间，不同goroutine之间的数据不会相互影响。这样可以确保每个goroutine都能够独立运行，不会被其他goroutine干扰。



### spillArgs

在riscv64架构下，函数调用时需要按照一定的规则将参数传递给被调用函数。其中，前八个参数可以通过寄存器传递，而超出八个的参数则需要通过栈传递。spillArgs函数的作用就是将超出八个的参数保存到栈中。

具体来说，spillArgs函数会根据参数的类型和大小，将其按照栈的高地址向低地址依次存储。比如，对于一个16字节的结构体参数，会分成两部分分别存储到连续的两个8字节的栈空间中。同时，spillArgs函数还会根据参数的对齐规则，将栈的起始地址调整到下一个对齐边界。

需要注意的是，spillArgs函数仅在参数个数超过八个时被调用。如果参数个数不足八个，则所有参数均通过寄存器传递。此外，spillArgs函数中的参数个数和类型与被调用函数所需的参数个数和类型是完全匹配的，因此调用者和被调用者在计算参数时不需要进行额外的转换和调整。



### unspillArgs

unspillArgs是一个用于恢复寄存器数据的函数，目的是将函数参数从栈帧中溢出到寄存器中。它在RISC-V 64位架构下的Go语言运行时中实现。

具体来说，Go语言中函数参数和局部变量都是存储在栈中的。当一个函数调用另一个函数时，它会将自己的参数存储到自己的栈帧中，并将控制权转移到被调用的函数。被调用的函数也有自己的栈帧，在其中存储由调用者传入的参数。在RISC-V 64位架构下，还有一些额外的寄存器，被用作参数传递和返回值。

在函数调用过程中，为了提高寄存器的利用率，由于所有参数不能同时放置在寄存器中，有时需要将一些参数从寄存器存储到栈中。在函数执行完毕后，unspillArgs函数就将这些参数从栈中恢复到寄存器中，方便后续的操作。

具体实现细节和算法请参考代码注释。



### getcallerfp

在Go语言程序中，每个函数都可以通过调用runtime.Callers函数获取当前函数调用栈的信息，但是如果想要获取函数调用栈的更详细信息，比如想要知道当前函数的调用者的信息，就需要使用getcallerfp函数获取当前函数的调用者的帧指针（frame pointer），然后通过解析该帧指针来获取调用者的信息。

具体来说，在stubs_riscv64.go文件中的getcallerfp函数实现了以下操作：

1. 获取当前函数调用栈的帧指针(fram pointer)的地址fp。

2. 通过对fp指向的内存进行解析，获取当前函数的栈帧信息，包括该函数的返回地址ra、上个函数的帧指针(fpb)、当前函数的参数等信息。

3. 取出当前函数的返回地址ra，然后再通过解析上个函数的帧指针(fpb)，获取上个函数的参数和返回地址，并将上个函数的返回地址作为getcallerfp函数的返回值返回，这样就得到了当前函数的调用者的地址。

需要注意的是，由于栈帧信息是按照特定的格式被放置在内存中的，所以解析栈帧信息需要遵循一定的规则和约定，否则可能会导致解析出错。



