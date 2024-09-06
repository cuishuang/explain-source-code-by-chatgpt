# File: util_gccgo.go

util_gccgo.go是Go语言的编译器源代码中的一个文件，其主要作用是提供对于gccgo编译器的支持实用工具。

具体来说，该文件包含许多用于与gccgo交互的函数和类型定义。其中，最重要的函数是execGCCGo函数，它可以启动gccgo编译器，并将编译器的输出转储为缓冲区。其他函数包括：

- goToolPath：返回gccgo安装目录下的工具路径。
- gccgoCmd：用于设置gccgo编译器的命令行参数。
- gccgoSupportedGOARCH、gccgoSupportedGOOS：返回支持的架构和操作系统。
- linkObjFiles：用于将GCC生成的目标文件链接到一个单独的可执行文件。
- nmExec：用于执行nm命令，以查看符号表信息。
- objdumpExec：用于执行objdump命令，以查看汇编代码。

总之，util_gccgo.go文件为Go语言的编译器提供了与gccgo编译器的支持实用工具，方便用户进行编译和链接等操作。

## Functions:

### useVFPv1





### useVFPv3





### useARMv6K





