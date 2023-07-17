# File: initAsanGlobals.go

initAsanGlobals.go文件是Go语言的源码文件，它的作用是初始化ASAN全局变量。

ASAN（AddressSanitizer）是一种内存安全框架，它可以检测和报告应用程序中的内存错误，例如堆缓冲区溢出、Use-After-Free、堆栈缓冲区溢出等。在Go语言中，ASAN被用于编写更加健壮和安全的代码。

initAsanGlobals.go文件包含一些全局变量的声明和初始化，这些变量用于启用和配置ASAN框架。其中一些变量包括：

1. __asan_default_options：一个字符串，用于指定ASAN运行时参数的默认值。

2. __asan_initialized：一个布尔值，指示ASAN是否已经初始化。

3. __asan_option_detect_container_overflow：一个布尔值，用于指定ASAN是否检测容器溢出错误。

4. __asan_option_quarantine_size_mb：一个整数，指定ASAN隔离器的大小，可以减少内存泄漏。

通过初始化这些变量，initAsanGlobals.go文件为ASAN提供了必要的配置和启用。这使得Go语言在运行时检测内存安全问题成为可能。




---

### Var:

### InstrumentGlobalsMap

在Go语言中，InstrumentGlobalsMap变量是用于记录ASan（AddressSanitizer）全局变量的映射的。ASan是一种内存错误检测工具，它可以帮助开发者在程序运行时检测出内存错误的问题，例如使用未初始化指针、缓冲区溢出等。在Go语言中，ASan工具的实现是基于LLVM编译器的插件形式，称为go-llvm-asan。

在ASan插件中，InstrumentGlobalsMap变量被用于记录全局变量的地址映射信息。具体来说，它是一个map类型，键是一个全局变量的包路径，值是一个InstrumentedGlobal类型的结构体。InstrumentedGlobal结构体包含了该全局变量的地址、类型和大小等信息。

在程序运行期间，ASan会通过InstrumentGlobalsMap变量来追踪全局变量的内存使用情况，以检测出内存错误问题。例如，在检测未初始化指针的时候，ASan会在InstrumentGlobalsMap中查找未初始化的全局变量，并设置对应的标记，以便在程序运行时检测到未初始化使用该全局变量的情况。因此，InstrumentGlobalsMap变量对于ASan全局变量检测的实现非常重要。



### InstrumentGlobalsSlice

在Go编译器中，InstrumentGlobalsSlice变量是用于跟踪已标记为污染的全局变量的列表。 全局变量是在程序启动时初始化的变量，可以在程序中的任何地方访问。 有时，全局变量可能会被多个协程访问，无法保证在不同协程之间的互斥性。 这可能导致数据竞争和其他问题。

为了解决这个问题，Go编译器实现了一个称为“污点分析”的技术，可以跟踪变量在程序中的传递和使用。 具体来说，每个可执行文件都在程序的每个汇编时刻都有一个唯一的污点分析版本，对于所有已标记为污染的全局变量，都会在运行时插入额外的指令来进行污点传播和检查，以确保程序在运行时的正确性。

InstrumentGlobalsSlice变量是一个全局的记录已标记为污染的全局变量的字符串切片。在程序汇编时刻，它会被初始化，并在每个导入包的汇编时刻中都会被更新。 对于每个已标记为污染的全局变量，都会生成一条注释为“go.asan.globals”的汇编指令，以表示该变量已被污染。 当程序运行时，污点传播和检查指令会检查所有已污染的全局变量的状态，以确保程序在运行时的正确性。



## Functions:

### instrumentGlobals

initAsanGlobals.go文件中的instrumentGlobals函数是用于插入代码来保护全局变量的函数。全局变量是在整个程序中都可以访问的变量，如果程序中的线程同时访问同一个全局变量，就可能会产生数据竞争，导致程序出现不可预料的行为。

ASan（Address Sanitizer）是一种内存错误检测工具，可以检测并报告程序运行时发现的内存错误和数据竞争。instrumentGlobals函数使用随机token为全局变量生成唯一的属性名，然后为每个全局变量创建一个对应的结构体，该结构体含有原始全局变量的指针和一些元数据，例如是否已被访问等。接着，instrumentGlobals函数会在程序中插入代码，来检测对全局变量的读写操作，并在运行时进行跟踪。当发现数据竞争时，ASan会触发一个错误，记录相关的堆栈信息等，帮助程序员调试问题。

总而言之，instrumentGlobals函数的作用是通过插入代码来保护全局变量，避免数据竞争导致的内存错误和不可预料的行为，并提高程序的稳定性和安全性。



### createtypes

在Go语言中，ASan（Address Sanitizer）是一种内存错误检查工具，它可以检测内存相关的错误，如访问已释放的内存、缓冲区溢出、使用未初始化的内存等等。

initAsanGlobals.go是Go语言的一个初始化文件，其中包含了创建ASan类型的函数creatypes。该函数的作用是创建ASan需要用到的一些类型，如asanInternalGlobals、asanOptions等。这些类型是ASan内部使用的，用于存储ASan的全局变量和参数等信息。

具体来说，creatypes函数的实现过程如下：

1. 创建asanInternalGlobals类型，该类型包含了ASan内部需要用到的全局变量，如flags、reportCounters等。

2. 创建asanOptions类型，该类型用于存储ASan的参数信息，如verbosity、replaceMalloc等。

3. 将创建的asanInternalGlobals和asanOptions类型赋值给全局变量internalGlobals和options。这样，在整个程序运行过程中，就可以通过这两个全局变量来访问ASan的内部变量和参数信息。

总之，creatypes函数是initAsanGlobals.go文件中的一个关键函数，它创建了ASan需要用到的一些类型，为ASan的运行提供了必要的支持。



### GetRedzoneSizeForGlobal

GetRedzoneSizeForGlobal函数的主要作用是计算全局变量的red zone大小。red zone是指内存中的一段区域，用于检测堆栈溢出和内存越界等错误。

该函数首先判断传入的全局变量类型是否为指针类型，如果是，则将red zone大小设置为8个字节（64位系统），否则将其设置为全局变量类型的大小加上8个字节，这8个字节用于存储red zone是否被破坏的标志位。

如果全局变量是一个数组，该函数会递归调用自身来计算每个数组元素的red zone大小并累加。最后，该函数返回计算出来的总red zone大小。

该函数是作为编译器的一部分实现的，是用于在构建时插入污点检测和错误检测代码的。通过计算全局变量的red zone大小，编译器可以在全局变量访问时插入检测代码以保证内存安全。



### canInstrumentGlobal

initAsanGlobals.go文件中的canInstrumentGlobal函数的主要作用是判断一个全局变量是否需要进行ASan插桩（instrumentation）。ASan是一种内存安全工具，它可以检测程序中的内存错误（如缓冲区溢出、空指针引用等），将其定位并提出警告。canInstrumentGlobal函数的实现细节如下：

- 在开头使用了一些常量和类型别名的定义，例如uintptr和userSymbolicInfo。
- canInstrumentGlobal函数的输入参数global是一个全局变量的元数据表示。该元数据包括变量的名称、类型、存储位置等信息。
- canInstrumentGlobal函数首先判断变量是否是一个指针类型或指向函数的指针类型。如果不是，则直接返回false，表示该变量不需要进行ASan插桩。
- 如果变量是指针类型，则判断该指针是否指向栈内存。如果是，则返回false；否则，继续执行。
- 如果变量是指向函数的指针类型，则需要进一步判断该函数是否是一个动态链接库（DLL）中的全局函数。如果是，则返回false；否则，继续执行。
- 对于其他类型的指针变量，或指向堆内存的指针变量，都需要进行ASan插桩。因此，canInstrumentGlobal函数返回true，表示该变量需要进行ASan插桩。

因此，canInstrumentGlobal函数的主要作用是为ASan工具提供指导，指示需要进行哪些全局变量的插桩。在程序运行时，ASan插桩会自动对这些变量进行检测和报告，增强程序的内存安全性。



