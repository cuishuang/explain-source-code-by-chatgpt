# File: extern.go

在 Go 语言中，`extern.go` 是 runtime 包中的一个文件，它的作用是定义了大量的 Cgo 导出函数和 C 程序调用 Go 函数时的 trampoline。这些函数包括：

- 内存分配和释放函数：`malloc`、`calloc`、`realloc`、`free` 等；
- 同步和锁机制函数：`pthread_...`、`*Mutex`、`*Cond` 等；
- 计时器函数：`time_...`；
- 系统调用函数：`sys_...` 等。

`extern.go` 文件的另一部分则定义了 Go 语言中的一些底层接口，包括：

- `runtime/os`：操作系统相关的底层接口；
- `runtime/mem`：内存管理相关的底层接口；
- `runtime/panic`：异常（panic）相关的底层接口。

在 Go 语言的编译器和运行时实现中，`extern.go` 文件扮演着重要的角色，其中定义了很多与底层系统交互的接口。这些接口的实现保证了 Go 语言能够在不同的操作系统和硬件平台上运行，并且能够充分利用底层系统的资源。需要注意的是，一般情况下，开发者不需要直接修改或调用`extern.go` 中的函数和接口，除非开发者需要在底层系统和 Go 语言之间进行交互。




---

### Var:

### defaultGOROOT

defaultGOROOT是一个字符串常量，用于表示Go语言运行时的安装路径（即GOROOT环境变量指定的路径）。这个变量的作用是在运行时库初始化的时候提供一个默认的GOROOT值。

如果用户没有指定GOROOT环境变量，运行时库将使用defaultGOROOT作为默认值。如果用户已经设置了GOROOT环境变量，那么defaultGOROOT值将被忽略。

在初始化运行时库时，默认的GOROOT值被用于查找和加载运行时库所需的各种文件和资源，例如Go标准库的包和库文件。因此，defaultGOROOT确保了运行时库能够正常启动并访问所需的文件和资源。

在实际应用中，我们通常会将defaultGOROOT设置为Go语言安装路径。默认情况下，Go语言安装路径通常位于/usr/local/go（Linux/MacOS）或C:\go（Windows）下。



### buildVersion

变量buildVersion用于记录当前编译的Go版本号，可以通过打印该变量的值来确定当前Go二进制文件所对应的版本号。该变量在运行时代码外部定义，以便可以在包外部访问该变量。在main包的main函数中，可以通过使用runtime.Version()函数来获取当前Go版本号的字符串表示形式。这对于诊断与调试Go程序非常有用，特别是在跨平台使用Go代码时。通过检查Go版本号的变化，可以了解新功能、错误修复和性能优化等方面的改进，从而决定是否更新使用Go代码的应用程序。



## Functions:

### Caller

在Go语言中，Caller函数是用于获取调用函数的信息的函数，其定义如下：

```go
func Caller(skip int) (pc uintptr, file string, line int, ok bool)
```

它的参数skip表示要获取信息的调用栈深度（0表示调用Caller的位置，1表示其调用者的位置，以此类推），返回值为调用函数的PC值、源文件名、行号和是否获取成功的标记。

Caller函数一般用于调试和错误信息输出，常见的用法如下：

```go
func foo() {
    pc, file, line, ok := runtime.Caller(1)
    if ok {
        log.Printf("caller: %s:%d\n", file, line)
    }
    // do something
}

func bar() {
    foo()
}

func main() {
    bar()
}
```

这里定义了两个函数foo和bar，foo使用Caller函数获取其调用者的信息并打印到日志中，而bar调用foo。当main函数调用bar时，会依次调用foo和Caller函数，输出类似如下的信息：

```
caller: /path/to/file.go:10
```

因此，Caller函数可用于快速定位问题所在的代码行。



### Callers

Callers函数可以返回调用者的程序计数器(PC)。PC是一个指向程序中指令的地址，它指明了正在被执行的指令的位置。当一段代码在执行时，它所对应的PC值会随之变化。通过对PC进行分析，程序可以确定哪一段代码正在运行，从而可以更好地理解程序的行为。

Callers函数的作用是获取当前goroutine的调用栈信息。在调用此函数时，我们可以将一个从goroutine的调用栈的顶部开始的slice作为参数传入。函数将会从当前栈帧开始，一直持续向上搜索，将调用栈的程序计数器的值保存到slice中，并返回保存的个数。通过这种方式，我们可以得到当前goroutine的调用栈信息，从而能够更好地了解程序的运行情况。

这个函数通常用于排查程序出现panic或死锁等问题时，我们可以通过Callers函数获取goroutine的调用栈信息，定位到问题发生的地方，进而解决问题。在性能调优方面，这个函数也可以用于识别程序中的瓶颈，通过分析调用栈，找出程序中耗时较长的函数，然后进行优化。



### GOROOT

在 Go 语言中，GOROOT 是一个环境变量，它指向存放 Go 标准库和 Go 工具链的根目录。在 Go 程序运行时，运行时系统会根据 GOROOT 环境变量找到对应的库和工具链。

在 go/src/runtime/extern.go 文件中，GOROOT 是一个内部函数，它返回一个字符串，该字符串表示当前程序运行时的 GOROOT。具体来说，这个函数会先从一个全局变量（runtimeGOROOT）中读取 GOROOT 的值，如果该值为空，则会调用 getgoroot 函数从操作系统中获取。

有了 GOROOT，Go 编译器和运行时就能够准确地查找标准库的位置，同时也可以方便地在多个版本的 Go 之间切换。如果你需要使用新版本的 Go 标准库或者工具链，只需要更新 GOROOT 环境变量即可。



### Version

在go/src/runtime/extern.go中，Version()函数用于返回当前Go版本的规范版本字符串。它返回一个字符串，该字符串标识给定Go二进制文件的版本，并且它包含在构建Go语言时使用的特定Git修订版本的二进制程序的构建时Git哈希和日期。

对于像Go这样的开源项目，版本控制非常重要。众所周知，Go是一门快速发展的语言，每个版本都可能使一些库不兼容，因此为使用者提供一种方便的方法查找所使用的Go版本非常重要。Version()函数帮助开发人员和用户识别正在使用的版本，并且可以在程序的日志和错误报告中使用它。

例如，在Go 1.17中，Version()的输出结果为"go1.17"。



