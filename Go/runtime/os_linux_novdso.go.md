# File: os_linux_novdso.go

os_linux_novdso.go文件位于Go语言标准库的/runtime目录中，该文件定义了一个函数getPageSize()，该函数用于获取页面大小。在Linux中，程序可以通过vdso（virtual dynamic shared object）的方式来访问内核的一些系统调用，这样可以避免进入内核态的开销，从而提高程序的性能。但是在某些环境下，比如在某些容器中，使用vdso可能会有问题，因此需要禁用vdso。os_linux_novdso.go文件的作用就是在禁用vdso的情况下仍然可以获取页面大小。具体来说，getPageSize()函数实现了以下几步操作：

1. 通过调用syscall.Syscall()获取系统的页面大小。

2. 如果Syscall()函数返回错误（比如当前系统不支持该系统调用），则通过使用mmap()手动映射一个内存区域，然后获取该区域的大小作为页面大小。

3. 如果mmap()也失败了，则使用一个常量作为页面大小（这个常量是在编译时根据系统架构确定的）。 

总之，os_linux_novdso.go文件的作用是提供了一种在禁用vdso的情况下获取页面大小的方法，从而保证程序的正常运行。

## Functions:

### vdsoauxv

在 Linux 上，vDSO (virtual Dynamic Shared Object) 是一个虚拟的共享对象，由内核提供，包含一些系统调用的实现。这样，在执行系统调用时，可以直接跳转到 vDSO 而无需经过用户态和内核态之间的切换，从而提高了系统调用的性能。os_linux_novdso.go 文件中的 vdsoauxv 函数就是用于获取 vDSO 辅助向量的函数。

vDSO 的实现会给每个进程创造一个“auxiliary vector”数据结构，记录了 vDSO 在该进程中的基地址和其他信息。vdsoauxv 函数通过遍历进程的“auxiliary vector”，找到其中包含 vDSO 基地址的条目，并返回该基地址。

在一些情况下，vDSO 可能会被禁用，此时 vdsoauxv 函数会返回 0。此时，glibc 库会回退到使用传统的方式来执行 Linux 的系统调用。但是这种方式会使得系统从用户态进入内核态，造成额外的 CPU 开销。因此，有些时候会选择在无法使用 vDSO 时终止程序，而不是回退使用传统方式。



