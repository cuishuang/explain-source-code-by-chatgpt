# File: at_sysnum_fstatat64_linux.go

at_sysnum_fstatat64_linux.go是Go语言标准库中的一个文件，位于go/src/internal/syscall/syscall_linux_amd64.go中，主要的作用是为Linux系统下的系统调用fstatat64()提供一个轻量级的实现。该系统调用用于获取指定文件结构的文件状态信息，例如文件大小、文件修改时间等。

在Linux系统中，fstatat()和fstatat64()是两个版本的系统调用，区别在于fstatat64()支持大文件（64位文件大小）的操作，同时也是较新的系统调用。而在Go语言中，Go的标准库实现并没有在syscall包中提供fstatat64()的实现，因此at_sysnum_fstatat64_linux.go就是为了补充该空缺而存在的。

该文件中包含了atSyscall()和fstatat64()两个函数。其中atSyscall()用于调用Linux系统下的系统调用，fstatat64()则是对atSyscall()进行了封装，实现了对fstatat64()的调用，并返回了一个Unix文件的Stat结构体，其中包含了该文件的各种状态信息。

总的来说，at_sysnum_fstatat64_linux.go的作用是为Go语言提供了一个轻量级的接口，用于获取Linux系统下指定文件的状态信息。

