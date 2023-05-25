# File: asm.s

sync包是Go语言提供的用于实现并发同步的包，其中的asm.s文件是一个汇编语言写的文件，主要用于实现了一些系统调用，提高了sync包中的一些函数的效率。

在sync包中，有一些函数需要调用系统原语，例如互斥锁、读写锁等，这些函数需要在操作系统层面进行实现。然而，由于操作系统的不同，这些原语的实现方式也不同，因此需要通过汇编语言来编写相应的实现。

具体来说，asm.s文件中实现了以下几个函数：

1. sync·atomic·AddInt32、sync·atomic·AddInt64、sync·atomic·AddUint32、sync·atomic·AddUint64：原子加操作。

2. runtime·mcall：用来阻塞当前的goroutine，并唤醒其它的goroutine。

3. runtime·osyield：系统调用，请求调度器切换到其它的goroutine。

4. sync·runtime·parksys：用来阻塞当前goroutine，并设置自旋锁，将来唤醒时会再次尝试获取锁。

总之，asm.s文件是sync包中至关重要的一部分。通过汇编语言的实现，可以提高一些函数的效率，增加程序的性能。

