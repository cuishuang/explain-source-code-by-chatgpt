# File: errno_windows.go

errno_windows.go是Go语言标准库中internal包下的一个文件，这个文件主要的作用是定义了在Windows平台下的系统错误码。

在Unix/Linux下，错误码一般定义在errno.h头文件中，而Windows平台下没有类似的标准头文件。因此，在Windows平台下，Go语言标准库通过定义errno_windows.go文件中的常量来代替errno.h文件中的定义。在这个文件中，定义了许多常见的Windows系统错误码，如ERROR_FILE_NOT_FOUND、ERROR_INVALID_PARAMETER等，并将其定义为了Go语言中的常量。

errno_windows.go文件主要包含以下内容：

1. 定义了Windows系统的错误码常量，常量名与系统错误码名称保持一致。例如：ERROR_FILE_NOT_FOUND、ERROR_INVALID_PARAMETER等，这些常量可以被Go语言中的其他模块进行引用。

2. 定义了系统调用相关的错误码常量，这些常量通常表示在进行系统调用时出现的错误，例如：WSAEACCES、ERROR_ACCESS_DENIED等。

3. 定义了函数名与错误码的映射关系，这些映射关系可以让调用者看到与错误码相关的函数名称，例如：MapLastError、MapSyscallError等。

总体来说，errno_windows.go文件的作用是定义了Windows平台下的系统错误码常量，这些常量可以帮助Go语言标准库中其他的模块进行错误处理和调用系统相关的API。




---

### Var:

### errERROR_IO_PENDING

在 Windows 操作系统中，当一个 I/O 操作（如文件读取或写入）被提交时，它不会立即完成。相反，系统通常将 I/O 操作放入队列中，然后允许程序继续执行其他操作。当 I/O 操作完成时，系统会通知程序并返回结果。

在 Go 语言中，当执行某些 I/O 操作时，如果操作无法立即完成，则会返回一个 "pending" 错误。在 Windows 中，这个错误的错误代码为 ERROR_IO_PENDING。这个错误表示 I/O 请求已经提交，但是操作尚未完成。程序可以继续执行，等待操作完成后系统会自动通知程序。

errno_windows.go 文件中的 errERROR_IO_PENDING 变量就是用来表示这个 pending 错误的。当 I/O 操作返回这个错误码时，程序可以通过判断错误码是否等于 errERROR_IO_PENDING 来确定当前操作是仍然在进行中还是已经完成了。



## Functions:

### errnoErr

errno_windows.go文件是Go语言内置库中的一个文件，其中errnoErr这个函数的作用是将Windows API返回的错误码映射为一个Go的error类型。

在Windows下，系统调用返回的错误码有时与Unix和Linux下不同，因此Go语言需要将Windows API的错误码转换为与Unix和Linux相对应的错误。errno_windows.go文件中定义了一个函数errnoErr来处理这一转换过程。

errnoErr函数接受一个Windows返回的错误码作为参数，并返回一个对应的Go错误类型。该函数定义了一个名为errMap的映射表，它将Windows错误码映射为Go错误类型。如果函数无法找到映射表中的对应项，则会返回一个默认的错误信息。

这个函数的作用相当于把Windows API的错误码转换为了Go语言中的错误类型，方便Go程序员使用这些错误码来进行错误处理。通过这个函数，Go程序员可以避免自己去查找Windows API的错误码，并且可以使用Go自带的错误处理机制来处理这些错误。



