# File: cgocall.go

cgocall.go文件是Go语言运行时（runtime）的一部分，它的作用是提供一个统一的接口，使Go代码能够调用C语言函数。它是实现Go和C语言互操作的一个重要模块。

具体来说，cgocall.go文件中的函数有以下作用：

1. 将Go语言栈转换为C语言栈。因为Go语言和C语言的调用约定不同，Go语言函数的栈和C语言函数的栈并不兼容，所以需要进行转换。

2. 保存和恢复Go语言的goroutine状态。因为在调用C语言函数时，当前Go语言线程会转换为C语言线程，并且恢复时需要重新设置Go语言线程的状态。

3. 处理Go语言和C语言的内存管理。当Go语言调用C语言函数时，需要注意内存的分配和回收。

4. 实现goroutine的抢占。在Go语言中，goroutine的调度是由运行时系统负责的。当一个goroutine调用了C语言函数而被阻塞时，其他goroutine可以被抢占执行。

总之，cgocall.go文件是Go语言运行时系统中非常重要的一个模块，为Go语言和C语言之间的互操作提供了非常方便和高效的接口。




---

### Var:

### ncgocall

在Go语言中，CGO（C语言调用Go语言函数的接口）是非常重要的功能之一。cgocall.go是CGO的核心文件之一，其中ncgocall变量是用来记录当前CGO调用次数的。

当Go程序在调用CGO的时候，它会创建一个新的线程，然后将CGO调用的参数与函数指针传递给该线程。每个线程都会在执行CGO调用之前增加ncgocall的值，以表示它正在执行一个CGO调用。当CGO调用结束时，该线程将ncgocall的值减少，并将控制权返回给Go运行时。

该变量的作用很重要，因为它使得Go程序可以限制CGO调用的并发数。这是因为在许多应用程序中，同时进行太多的CGO调用可能会降低性能并可能导致一些未定义的行为。因此，通过限制同时进行的CGO调用数，可以确保CGO调用的稳定运行。



### racecgosync

racecgosync是一个用于Go语言运行时系统中的Goroutine之间同步的变量。这个变量是一个布尔类型，用于指定Goroutine在使用Cgo调用时是否应该启用竞态检测（Race Detection）。

竞态检测是Go语言提供的一种机制，可以帮助开发者发现程序中存在的并发问题。通过在Goroutine中使用竞态检测，可以检查程序是否存在数据竞争、死锁等问题，从而提高程序的可靠性和稳定性。

当racecgosync被设置为true时，Go语言运行时系统会启用竞态检测机制，对Cgo调用中的Goroutine进行检查。这个变量默认为false，需要手动设置为true才能启用竞态检测。

需要提醒的是，启用竞态检测会影响程序的性能，因为它会增加程序运行时的开销。开发者应该根据具体情况选择是否启用竞态检测。对于一些关键的程序和应用，启用竞态检测可以帮助提高程序的质量和可靠性。






---

### Structs:

### cgoCallers

在go的runtime中，cgoCallers这个结构体的作用是用于记录当前的goroutine堆栈信息。这个结构体中包含了一个指向Goroutine的指针(g)，表示当前堆栈所属的goroutine；一个pc指针，表示当前堆栈的程序计数器(PC)；以及一个sp指针，表示当前堆栈的栈顶指针(SP)。

这个结构体在C和Go之间的调用中非常重要。当Go函数调用C函数时，runtime需要记录当前Go堆栈的信息，以便C函数在执行完毕后重新获得Go的堆栈信息。这些信息中包括了当前程序计数器(PC)和栈顶指针(SP)的位置。而cgoCallers结构体就是用于记录这些信息的，以便在C函数执行完毕后能够正确恢复Go的堆栈信息。

另外，cgoCallers结构体还可以用于调试，它可以将堆栈信息记录下来，以便用户在调试时识别问题所在。在Go语言中，堆栈调用信息非常重要，在出现问题时可以帮助我们快速定位代码中的错误。cgoCallers结构体提供了记录和检索堆栈信息的功能，为调试工作提供了很大的帮助。



### argset

argset结构体是用于管理调用C函数时的参数的结构体。它的作用在于将Go函数传递给C函数时，将Go类型转换为C类型，并将它们作为参数传递给C函数，并在C函数返回时将结果转换为Go类型。

具体来说，argset定义了一个数组args，它包含了所有要传递给C函数的参数。它还包含一个len字段，它表示args数组中的元素数目。当调用C函数时，argset使用cfunc字段指向的C函数（类型为funcptr）以及将要传递给C函数的参数列表来调用该函数。它还使用ret字段来存储C函数返回的结果。

argset还定义了一些方法，例如add，用于将Go类型添加到args数组中，以便将它们传递给C函数。它还定义了get方法，用于从args数组中获取Go类型的参数。

总之，argset结构体允许将Go类型转换为C类型，并在C函数返回时将它们重新转换为Go类型。



## Functions:

### syscall_cgocaller

syscall_cgocaller是一个函数指针，用于在C代码中调用Go函数。它的作用是将C语言中的函数调用转换为Go语言中的函数调用，从而实现跨语言调用。

在Go语言中，每次进行系统调用时，需要先将参数转换为C语言中的对应类型，然后再调用C语言中的系统调用函数。而在通过cgo调用C语言函数时，也需要先将参数传递给syscall_cgocaller，再通过syscall_cgocaller间接调用Go语言函数。

syscall_cgocaller的具体实现方式是通过汇编语言来实现的，它可以将C语言参数封装成Go语言中的参数结构体，并使用Go语言中的汇编代码来进行调用。因此，使用syscall_cgocaller可以方便地实现Go和C语言之间的交互，提高了程序的灵活性和可移植性。

总之，syscall_cgocaller是一个关键的函数指针，它的作用在于将C语言中的函数调用转换为Go语言中的函数调用，从而实现跨语言调用，是cgo机制的核心之一。



### cgocall

cgocall函数是Go语言运行时系统和C代码之间交互的关键函数，它的作用是在Go语言和C语言之间执行函数调用。在Go语言调用C函数的时候，需要通过cgocall函数将控制权交给C语言，并在C语言执行完毕后将控制权还给Go语言。

具体来说，cgocall函数的作用是：

1. 将当前Goroutine暂停，保存Goroutine上下文。

2. 调用Go和C之间的桥接函数，将C函数的参数和函数指针传递给C语言代码。

3. 在C函数调用结束后，再次调用桥接函数来获取C函数的返回值。

4. 恢复Goroutine上下文，将执行流程返回给Go语言程序。

因此，cgocall函数的重要作用是协调Go语言和C语言之间的调用关系，确保程序的正确性和可靠性。同时，在cgocall函数内部，Go语言会采取一些措施来保护C代码不受Go语言的垃圾回收和协程调度等影响，从而保证整个程序在运行过程中的稳定性和可靠性。



### cgocallbackg

在go语言中，cgocallbackg函数是用于将go函数包装成C语言函数，以便在C语言环境中被调用的函数。具体来说，它将一个go函数和相关的参数打包成一个结构体，然后将其作为参数传递给C语言函数（通过指针形式），这样就可以在C语言环境中执行相应的操作。

这个函数在执行过程中涉及到一些重要的参数和机制，比如goroutine的调度、堆内存申请等。其中，最重要的是它使用了go的调度器（scheduler）来控制goroutine的调度，从而保证了go函数能够在C语言环境中正常执行。

具体来说，cgocallbackg函数会启动一个新的goroutine来执行go函数，并将其中任意一个goroutine挂起，以便在之后的调度过程中继续执行。这保证了go函数在C语言环境中的可调用性和可重入性。

另外，cgocallbackg函数还需要申请堆内存来存储参数和返回值，并释放这些内存的操作，以保证在go函数执行完毕后内存能够被及时回收，避免内存泄露等问题。

总的来说，cgocallbackg函数在go语言中扮演着极其重要的角色，它保证了go函数能够在C语言环境中正常运行，并且在操作上做了很多的优化和处理，使得该函数更加高效和灵活。



### cgocallbackg1

cgocallbackg1是Go语言调用C/C++函数时，通过go的调度器将调用函数的控制权转移给另一个go routine去执行的函数。

当在Go语言程序中调用C/C++函数时，会通过cgo将控制权转移到C/C++函数执行，这时在cgo代码中会通过CGO当前执⾏的线程找到一个可⽤的G（go routine），然后将C/C++函数执行的控制权以G的形式托付给Go调度器，由Go调度器⾃⾏决定该G在哪条系统线程上执⾏，然后当该G在相应系统线程获取到CPU时间片时，会被继续执⾏。

在调⽤G操作系统线程前和调用C/C++函数后都需要有一个回调函数，而这就是cgocallbackg1的作用，它就是负责执行这个回调函数的。实际上，⽤户级代码并不直接调用cgocallbackg1，而是通过CGO调用它的包装函数_cgo_runtime_cgocallback_gofunc。这个函数就是为了调用cgocallbackg1并传递回调函数的参数的。cgocallbackg1只是被这个函数调用执行回调函数的。 它的作用是：将 C/C++调用函数的执行函数体转换到 Go的goroutine中执行，从而避免在 C/C++函数中阻塞过久影响协程的执行，实现了Go和C/C++之间的调用交互。



### unwindm

在 Go 的运行时环境中，goroutine 是一种轻量级的线程，由 Go 的运行时系统调度。当一个 goroutine 执行到某个调用了 C 语言库函数的程序行时，Go 运行时系统会使用 `cgocall` 函数调用一个对应的 Go 函数，这个函数会负责创建一个新的线程来执行 C 函数。

当 C 函数执行完成后，Go 函数会被回调并且负责清理线程。但是，如果 C 函数因为出现了长时间的 I/O 操作，如网络请求、磁盘读写等等，而被阻塞了，那么在此期间，Go 的运行时系统不能等着这个 C 函数结束，否则可能会导致整个程序出现阻塞。这时候，Go 的运行时系统需要在新的线程上恢复当前 goroutine 的执行，这个过程就是通过 `unwindm` 函数实现的。

`unwindm` 函数的主要作用是将当前 goroutine 所在线程的状态保存到栈中，并切换到与当前 goroutine 相关的 M（Machine）。M 是 Go 运行时系统中的一个概念，它代表了一个系统线程和一组 goroutine。M 负责管理 goroutine 的调度和执行，同时也处理与操作系统的交互。

在执行 `unwindm` 函数之前，M 所在线程将执行预添加函数（preemptive）并释放持有的锁。在 `unwindm` 函数执行结束后，当前 goroutine 被放回到调度队列中等待下一次执行机会。这个过程和 Go 的协程调度机制是紧密相关的，Go 运行时系统会动态地分配和回收 M 和 goroutine，并根据负载自动调整 goroutine 的执行位置和线程数量，从而提高程序的并发性和性能。



### badcgocallback

在go/src/runtime/cgocall.go文件中，badcgocallback函数的作用是处理CGO回调中的错误。

CGO（C语言与Go语言间的互操作）回调指的是用C代码定义的函数，它在Go程序中被调用。当C语言中定义的函数被作为Go语言函数的参数传递时，Go语言必须确保它们是安全的，以避免给CGO带来不必要的麻烦。

在某些情况下，如果CGO回调函数没有正确地使用，则可能会导致Go程序崩溃或出现其他问题。为了防止这种情况发生，当发现CGO回调函数出现错误时，badcgocallback函数会被调用并抛出一个恐慌（panic），以确保程序安全地退出。

具体而言，badcgocallback函数会打印出错误信息并终止程序的执行，以避免潜在的安全漏洞。在这个过程中，它还会记录跟踪信息，以帮助调试和解决CGO回调中的问题。



### cgounimpl

在Go语言中，Cgo调用是一种让Go和C语言互操作的机制。在运行时，当需要调用C函数时，Go语言会在当前的Goroutine中创建一个新的系统线程，然后调用C函数。在该线程中，C函数可以操作C语言数据类型，还可以调用其它C函数。一旦C函数返回，Go语言会将其结果传递给调用方，并且释放该系统线程。

在cgocall.go文件中，cgounimpl函数是一个未实现的CGO调用。它是Cgo机制的异常处理函数，用于处理在Cgo调用期间引发的未处理异常。当在Cgo调用期间出现未处理异常时，调用该函数来进行必要的清理工作，并且将异常信息返回给调用方。该函数的代码如下：

```go
// cgounimpl is called when a cgo call is not implemented for a particular operation.
//go:nosplit
func cgounimpl() {
	typedmemclr(cgoTopFix, cgoCallSP-cgoTopFix) // clear what we pushed for the call.
	gp := getg()
	if gp.syscallsp == 0 {
		// We are not in a syscall, so there is no need to issue a SigPanic.
		// Instead, panic with a descriptive message.
		panic("cgo argument has Go pointer to Go pointer")
	}
	// We are in a syscall, so we can't panic. Issue the familiar "fatal error" message.
	print("fatal error: cgo argument has Go pointer to Go pointer\n")
	exit(2)
}
```

在函数实现中，首先调用typedmemclr函数来清除为CGO调用准备的栈空间。然后，获取当前Goroutine的指针，并判断是否在系统调用中。如果未在系统调用中，则表示出现了不支持的CGO操作，并且抛出一个带有详细信息的panic异常。如果在系统调用中，无法抛出异常，而是打印一个错误信息，并退出进程。

总之，cgounimpl函数是Cgo调用异常处理的关键部分之一，它负责在出现未处理异常时进行必要的清理工作，并返回异常信息。



### cgoCheckPointer

cgoCheckPointer函数是Cgo调用C函数时进行参数检查的函数之一。

在Cgo调用C函数时，需要将Go变量转换为C语言变量类型。但是，由于Go和C语言内存空间管理的机制不同，因此在进行这些内存空间转换时，需要进行参数检查以确保不会出现内存分配错误或其他内存相关的问题。

cgoCheckPointer函数主要用于检查指定的Go指针是否为空，如果为空，则向程序发送错误消息并返回false，否则返回true。

在将Go指针转换为C指针时，如果指针为空，程序将会崩溃或者出现其他严重的错误。因此，cgoCheckPointer函数的作用就是确保程序在调用C函数时不会使用空指针，以提高程序的稳定性和安全性。



### cgoCheckArg

在go程序中使用cgo调用C函数时，需要将go语言中的类型和C语言中的类型进行转换。cgoCheckArg函数的作用是在进行这些类型转换的过程中检查参数是否合法，并将其转换为C语言中对应的类型。

具体来说，cgoCheckArg函数会根据传入的参数类型，判断其是否是有符号或无符号整数、浮点数、指针、结构体等基本类型。它还会检查结构体中的每个字段是否符合C语言的对齐规则，并处理嵌套结构体的情况。除此之外，cgoCheckArg还会检查字符串、数组和切片等复杂类型的参数，并相应地转换为C语言中的类型。

总的来说，cgoCheckArg函数主要是为了保证go程序中调用C函数时参数类型和C语言中的类型匹配，从而避免类型不匹配造成的程序崩溃或内存泄漏等问题。



### cgoCheckUnknownPointer

cgoCheckUnknownPointer是一个函数，其目的是检查是否存在未知指针类型的CGo指针。在CGo中使用C代码时，我们需要将Go指针传递给C代码，因为C代码不能直接访问Go的堆栈。cgoCheckUnknownPointer函数的作用是检查传递给C的Go指针是否已经注册到C中，以下是该函数的实现：

```go
func cgoCheckUnknownPointer(ctx uintptr, p unsafe.Pointer) {
    if trackCgoPointer == nil || !iscgoarm || cgoAlwaysFalse {
        return
    }
    // Check if this pointer is already registered.
    if ptr := trackCgoPointer(p, false); ptr != nil {
        // Already registered.
        return
    }
    // Check if this pointer is in a stack that hasn't been scanned yet.
    st := (*stack)(atomic.Loaduintptr(&curg.stack))
    if st == nil {
        return
    }
    // We're looking for pointers in the [sp, stacklo) range. Go
    // believes the stack is empty below stacklo, but we need to
    // scan it because we might find pointers there.
    //
    // Low priority TODO: Reduce stack scanning when we can prove
    // that safe points are not in the [sp, stacklo) range.
    pointerfind(st, st.sp, uintptr(unsafe.Pointer(&stackLo)), p, func() {
        cgocallbackg1(ctx, funcPC(cgoPanicUnknownPointer), noescape(unsafe.Pointer(&p)))
    })
}
```

如果指针被找到，则表示该指针未被注册并且未被分配。在这种情况下，将触发CGo运行时恐慌。因为未知的指针可能会导致C代码不知道如何释放它们，从而导致内存泄漏等问题。

此外，该函数还使用堆栈扫描算法来查找未被注册的指针。它会对Go栈中的每个指针进行扫描，并将扫描到的所有未被注册的指针传递给cgocallbackg1函数，触发CGo运行时恐慌。这种检查确保我们在向C代码传递指针时不会意外地泄漏Go堆栈中的指针，从而确保代码的安全性。



### cgoIsGoPointer

在go/src/runtime/cgocall.go中，cgoIsGoPointer函数用于判断传递给C代码的参数是否是Go指针。它采用一个unsafe.Pointer类型的参数，并检查它是否属于Go堆上的虚拟地址，即是否指向一个Go对象。

如果参数确实指向Go对象，则该函数返回1，否则返回0。根据返回结果，调用者可以进一步处理传递给C代码的参数，并确保它们能够正确地从C代码返回到Go代码中。

这个函数的作用是确保C代码可以安全地使用传递给它的指针，它对于跨语言调用来说很重要。由于Go和C代码是在不同的内存空间中运行的，因此涉及指针传递时需要特殊注意。这个函数帮助确保指针传递的正确性，进而保证应用程序的稳定性和可靠性。



### cgoInRange

cgocall.go中的cgoInRange函数的主要作用是关于cgo调用的范围检查。它的输入参数包括一个输入的uintptr类型地址指针和一个uintptr类型的大小参数。它的主要作用是检查这个地址和大小是否在当前的goroutine的stack的范围内，如果不在的话，就会panic并报错。这个函数主要是用来确保cgo调用时不会访问到不在堆栈范围内的内存，从而保证程序的健壮性和稳定性。

例如，如果在cgo调用期间试图访问超出其当前goroutine的栈空间，则cgoInRange函数将检查并抛出对应的异常，防止程序在使用时发生崩溃或者内存泄露等异常问题。

总之，cgoInRange函数是在进行cgo调用前检查栈空间范围，防止越界访问和内存泄漏，保证程序的安全性和可靠性。



### cgoCheckResult

cgocall.go文件中的cgoCheckResult函数是一个辅助函数，用于检查C函数调用的结果，并在出现错误时进行日志记录和处理。在cgo调用C函数时，与Go运行时的异常处理不同，C函数的错误通常通过返回值表示。因此，在调用C函数后，需要调用cgoCheckResult函数来检查它的返回值是否异常，并在检测到异常时进行日志记录和处理。

具体来说，cgoCheckResult函数会判断C函数返回的错误码是否为零。如果非零，则会记录错误日志，并通过runtime.throw函数抛出一个panic异常，以结束当前的cgocall。如果错误码为零，则返回C函数的返回值。该函数通常用于cgo调用C函数的wrapper函数中。

下面是cgoCheckResult函数的源代码实现：

func cgoCheckResult(code _Ctype_int32_t, fn string) int32 {
   if code != 0 {
      logCgoCall(fn, code)
      throw("cgo error")
   }
   return int32(code)
}

该函数接收两个参数，第一个是C函数的返回值，第二个是C函数的名称，用于输出错误日志。首先判断返回值是否为零，如果不是，则调用logCgoCall函数输出错误日志。然后调用throw函数抛出panic异常，以结束cgocall的执行。如果返回值为零，则直接返回该值。



