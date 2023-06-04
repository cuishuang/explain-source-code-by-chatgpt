# File: threadpanic.go

go/src/runtime/ threadpanic.go文件的作用是处理Go语言中的异常情况，比如panic和recover。

具体来说，threadpanic.go文件定义了一些关键的结构和函数，包括：

1. panicFrame结构体：用于保存当前执行上下文的信息，包括函数调用栈、goroutine ID等等。

2. panic的实现函数：runtime·panic、runtime·gopanic、runtime·plainPanic等，这些函数都是用来触发panic的。

3. recover的实现函数：runtime·recover，用来从当前函数调用栈中恢复panic。

4. 调用栈的处理函数：runtime·printcreatedframes、runtime·callers、runtime·pcvalue等函数，用来处理函数调用栈等信息。

同时，threadpanic.go文件还提供了一些预定义变量和常量，比如panicCode、panicError等。

总的来说，threadpanic.go文件主要是用来处理Go语言中的异常情况，提供了一套完整的处理方案，包括panic、recover和异常信息处理等功能。

## Functions:

### init

init函数是Go语言中的一个特殊函数，用于初始化包，并在包被引入时自动执行。在go/src/runtime中的threadpanic.go文件中，init函数的作用是为每个m（goroutine）初始化panic信息结构体。当一个goroutine发生panic时，就会把panic信息保存到该结构体中，然后由runtime包处理panic。

具体来说，init函数会为每个m预分配一个panic信息结构体panicbuf，并把该结构体和m绑定在一起。当m发出panic信号时，runtime会将panic信息保存到该结构体中，并通过goroutine的控制流向上传递panic，直到被某个recover函数处理。

在Go语言中，panic和recover是一对特殊的函数，用于处理程序的错误和异常情况。当程序出现不可恢复的错误时，可以使用panic函数抛出一个异常，然后通过recover函数捕获该异常并进行处理。init函数的作用是为panic和recover函数的正常运行打下基础，保证程序能够正确地处理异常情况。



### CgoExternalThreadPanic

CgoExternalThreadPanic函数的作用是处理由外部线程（即非Go线程）引发的panic情况。在Go语言的并发模型中，所有的Goroutine都运行在Go runtime的调度下，而不是直接运行在操作系统的线程上。因此，如果一个外部线程调用了Go语言中的一个函数，并在其中触发了panic，这个panic是没有被Go runtime掌控的，需要进行处理。

具体来说，CgoExternalThreadPanic函数会创建一个新的Goroutine来处理由外部线程引发的panic，并通过调用runtime.gopanic函数来抛出一个由runtime.sigpanic包装的panic。这个新的Goroutine执行一个特殊的函数，即“non-Go code”，负责进行panic的处理逻辑，最终会将控制权交给runtime.sigpanic函数来进行后续处理。

需要注意的是，CgoExternalThreadPanic函数只能在开启了CGO的情况下才会被编译到程序中。在纯Go程序中是不需要的，因为所有的Goroutine都是由Go runtime调度的。因此这个函数的实际使用场景比较局限，但是在一些需要与非Go语言交互的场景下仍然非常有用。



### gopanic

gopanic()函数是在Go中用来抛出panic的函数，它的作用是在当前Goroutine中抛出一个panic的异常值，一旦异常值被抛出，当前的Goroutine将会被终止，并在调用栈向上传递，直到被recover捕捉到，或者程序完全结束。

gopanic函数的主要任务是向当前的Goroutine抛出一个panic异常，该异常会在当前Goroutine内部不断向上抛，直到被捕获或者崩溃。它的实现主要依赖于runtime包中的一些函数，如getcallersp()和getcallersp()函数，用于获取当前Goroutine调用栈的信息。

在调用gopanic函数时，会将panic信息压入当前Goroutine的栈帧中，并进入执行栈的panic流程。具体而言，在gopanic函数内部，会首先获取当前调用栈的信息，用于构建panic信息，然后构造对应的panic结构体，并通过_g_.m对象的throw变量将panic结果传递给当前M线程。然后，当前Goroutine将被设置为P标记为Panic，表示当前Goroutine已经处于panic状态，接下来的调用将会在该Goroutine上调用系统栈处理器。

总之，gopanic函数的作用是将panic异常抛出到当前Goroutine的调用栈上，并将控制流传递到系统栈处理器，该处理器在捕获到panic时可以处理异常并在程序崩溃之前执行一些过程。



