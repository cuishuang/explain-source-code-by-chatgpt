# File: /Users/fliter/go2023/sys/unix/affinity_linux.go

在Go语言的sys项目中，/Users/fliter/go2023/sys/unix/affinity_linux.go文件的作用是实现Linux平台上的CPU亲和力相关的操作。

CPUSet是一个表示CPU集合的数据结构。它由两个字段组成：mask和count。其中，mask是一个位掩码，每个位表示一个CPU的可用性；count表示CPU集合中可用的CPU数量。

- schedAffinity函数用于设置进程或线程的CPU亲和力；它接收一个PID作为参数；
- SchedGetaffinity函数用于获取进程或线程的CPU亲和力；它接收一个PID和一个CPUSet作为参数，用于存储获取到的CPU亲和力；
- SchedSetaffinity函数用于设置进程或线程的CPU亲和力；它接收一个PID和一个CPUSet作为参数，用于设置新的CPU亲和力；
- Zero函数用于将一个CPUSet清零，即将mask字段置为0；
- cpuBitsIndex函数用于获取给定位置的位在mask字段中的索引；
- cpuBitsMask函数用于获取给定位置的位掩码；
- Set函数用于将给定位置的位设置为1；
- Clear函数用于将给定位置的位设置为0；
- IsSet函数用于判断给定位置的位是否被设置为1；
- Count函数用于统计CPU集合中设置为1的位的数量。

