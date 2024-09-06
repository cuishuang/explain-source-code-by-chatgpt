# File: zgoarch_mips64le.go

该文件的作用是定义了MIPS64LE架构下的相关常量和函数，用于支持Go代码在MIPS64LE架构CPU上的编译和运行。

其中包含的常量有:

1. _MIPS_ARCH_MIPS64R6：该常量用于指示MIPS64LE架构下的MIPS64R6指令集。

2. _MIPS_SIM_ABI64：该常量用于指示MIPS64LE架构下的64位ABI。

包含的函数有:

1. func archInit()：该函数用于初始化MIPS64LE架构下的相关参数。

2. func getcallerpc() uintptr：该函数用于获取当前函数调用者的程序计数器。

3. func getcallersp() uintptr：该函数用于获取当前函数调用者的栈指针。

通过这些常量和函数的定义，实现了在MIPS64LE架构CPU上的Go代码编译和运行。

