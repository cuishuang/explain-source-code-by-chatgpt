# File: fstatfs_zos.go

fstatfs_zos.go是Go语言标准库中的文件之一，用于在IBM z/OS操作系统上获取文件系统的信息。

在Linux和其他类Unix操作系统上，可以使用statfs系统调用来获取文件系统的信息，但在IBM z/OS操作系统上，没有类似的系统调用。因此，fstatfs_zos.go实现了一个自定义的函数fstatfs来获取文件系统的信息。

fstatfs函数使用Unix System Services (USS)命令dfp来获取文件系统的信息。dfp命令在z/OS上用于显示文件系统的使用情况和配额信息。

fstatfs函数返回一个包含文件系统信息的结构体Statfs，其中包括文件系统块大小、总大小、可用大小、文件节点数量等信息。

该文件的作用是为在IBM z/OS操作系统上运行的Go程序提供文件系统信息。

