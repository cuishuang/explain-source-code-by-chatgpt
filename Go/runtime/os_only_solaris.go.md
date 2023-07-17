# File: os_only_solaris.go

os_only_solaris.go是Go语言的运行时包（runtime）中的一个文件，它的作用是在Solaris系统上编译Go语言程序时，提供一些Solaris特有的操作系统接口。

在Solaris系统上，操作系统接口和其他Unix-like系统有所不同，因此需要提供一些特殊的系统调用和系统常量给应用程序使用。os_only_solaris.go文件声明并实现了这些特殊的系统接口，包括：

1. ERESTARTSYS和ERESTARTRESTARTBLOCK常量：用于处理阻塞系统调用时的重新启动。这些常量定义了在Solaris系统上发生阻塞系统调用时，进程应该如何从阻塞状态恢复。

2. SetsockoptInt和Getsockname系统调用：用于设置和获取socket选项。在Solaris系统上，这些系统调用和其他Unix-like系统的实现有所不同。

3. Fattach和Fdetach系统调用：用于动态附加和分离文件系统。这些系统调用只在Solaris系统上存在。

通过os_only_solaris.go文件中实现的这些特殊接口，可以使Go语言应用程序在Solaris系统上正确地运行和执行，提高运行稳定性和可靠性。

## Functions:

### getncpu

os_only_solaris.go文件中的getncpu函数是用来获取系统的CPU核心数的。在Solaris操作系统中，这个函数使用的是top工具来获取CPU核心数。

具体来说，getncpu函数首先使用os/exec包中的Command函数来创建一个top命令的进程，然后使用命令行参数"-b"和"-d1"来让top命令每隔一秒输出一次CPU使用情况，这个输出会被写入到一个IO读取器中。接着，getncpu函数使用bufio包中的NewScanner函数来创建一个扫描器，用来读取IO读取器中的输出内容。然后，getncpu函数遍历扫描器的输出，查找CPU核心数相关的信息，最终返回系统CPU核心数。

总之，getncpu函数的作用是获取Solaris操作系统的CPU核心数。



