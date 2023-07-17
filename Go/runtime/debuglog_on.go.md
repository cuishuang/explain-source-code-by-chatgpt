# File: debuglog_on.go

debuglog_on.go是Go语言标准库中runtime包的一个文件，它的作用是启用debuglog DEBUG级别的日志输出。

在Go运行时系统中，debuglog是一种调试工具，用于打印运行时系统中的一些调试信息，例如goroutine的调度、内存管理、垃圾回收等等。这些信息可以帮助开发人员进行性能优化、错误调试等工作。

在debuglog_on.go中，有一个名为debugLogEnabled的变量，该变量用于表示是否启用debuglog的DEBUG级别日志输出。如果该变量的值为true，则会打印DEBUG级别的日志；否则，不会打印。

debugLogEnabled的值在程序启动时被初始化为false，用户可以通过命令行参数“-d”来启用它。如果用户使用了“-d”参数，则debugLogEnabled的值将被设置为true，同时将打印DEBUG级别的日志输出。

总之，debuglog_on.go的作用是允许用户在运行时启用debuglog的DEBUG级别日志输出，从而更好地进行调试和优化工作。




---

### Structs:

### dlogPerM

在debuglog_on.go文件中，定义了一个dlogPerM结构体，它表示debug log的存储方式。

具体来说，dlogPerM结构体有以下4个字段：

1. m - int类型，表示每个log消息的大小，单位是字节。默认值是1024。

2. n - int类型，表示log消息的数量。默认值是8192。

3. f - string类型，表示存储log消息的文件名。默认值是"default.log"。

4. r - int类型，表示在写文件的时候是否开启日志轮替（Log Rotation）。如果r的值大于0，那么表示开启日志轮替。日志轮替是指在写日志时，如果当前文件大小达到了一定的阈值，就会自动将当前文件命名为旧日志文件，并创建一个新的文件来继续写入日志。默认值是0，表示不开启日志轮替。

总的来说，dlogPerM结构体定义了debug log存储的详细信息，包括文件名、每个log消息大小、消息数量等等。这些信息会被传递到实际的log存储实现中，并对其进行相关的设置和配置。



## Functions:

### getCachedDlogger

文件debuglog_on.go是 Go 语言运行时库（runtime）的一部分。getCachedDlogger函数的主要作用是获取调试日志记录器（dlogger），以便在对程序进行调试时打印调试日志信息。

该函数的具体作用如下：

1. 它首先尝试从缓存中获取调试日志记录器（dlogger），以避免重复创建记录器。

2. 如果没有找到缓存的记录器，则创建一个新的 dlogger。

3. 创建新的dlogger时，使用Mutex确保同一时刻只有一个线程在创建。

4. dlogger是 runtime 包中的一种日志记录器，用于记录Goroutine或其他运行时事件的日志信息。

5. 该函数提供了一种方便的方法，让使用这个调试日志功能的程序可以更容易地记录运行时信息和调试信息。

总之，getCachedDlogger函数是debuglog_on.go文件中的一个有用的函数，用于获取运行时调试日志记录器，是Go语言程序在调试过程中记录日志信息的重要工具。



### putCachedDlogger

putCachedDlogger函数是用于将新的dlogger添加到已经存在的缓存中的函数。在runtime中，dlogger是一个结构体，它负责输出debug log信息。在putCachedDlogger函数中，它会检查是否已经存在一个相同的dlogger，如果存在则返回，否则将新的dlogger添加到缓存中。

这个函数的作用是尽可能的重用已经存在的dlogger，避免不必要的创建和销毁。这样可以提高程序的性能，因为创建对象和销毁对象都是比较耗费时间的操作。通过缓存dlogger对象，可以减少这些操作的次数，从而提高程序的效率。

在debuglog_on.go文件中的其他函数中，都会调用putCachedDlogger函数，尝试把新的dlogger缓存起来，以便后续能重用它。因此，putCachedDlogger函数可以说是整个debug log功能的核心之一。



