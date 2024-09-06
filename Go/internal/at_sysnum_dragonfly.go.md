# File: at_sysnum_dragonfly.go

at_sysnum_dragonfly.go 文件是 Go 语言标准库中 internal 包内部的一个文件。它的作用是提供一个函数，用于转换系统调用的名称为系统调用号，目前仅支持在 DragonFly BSD 操作系统上进行。

在 DragonFly BSD 操作系统中，系统调用的名称与 Linux 或其他操作系统有所不同。因此，at_sysnum_dragonfly.go 文件中定义了一个名为 sysnameToSyscall 的函数，用于将操作系统中的系统调用名称映射为系统调用号。这个函数会在 os 包中的 syscall_unix.go 文件中使用，用于执行各种系统调用。

需要注意的是，由于 at_sysnum_dragonfly.go 文件定义在 Go 语言的 internal 包内部，因此它只能在 Go 标准库的内部使用，而不能在第三方库或应用程序中直接使用。这对于保护 Go 语言标准库的稳定性和正确性非常重要。

