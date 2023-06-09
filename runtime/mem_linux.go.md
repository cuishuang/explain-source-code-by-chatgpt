# File: mem_linux.go

mem_linux.go是Go语言运行时的一个源码文件，它的主要作用是为Linux系统提供内存管理功能。

具体来说，mem_linux.go定义了一些函数，这些函数用于获取和设置内存限制、分配和释放内存、查看内存使用情况等。其中一些函数包括：

1. maxMHeap 和 getPageSize：用于获取最大可用堆内存和页面大小。

2. initSizes 和 sysAlloc：用于初始化内存分配器，并分配和释放内存。

3. DebugMemoryAllocate 和 DebugMemoryFree：用于在DEBUG模式下跟踪内存分配和释放。

除了函数之外，mem_linux.go还定义了一些与内存相关的全局变量，这些变量用于跟踪内存使用情况、限制分配器的内存消耗等。

总的来说，mem_linux.go的主要作用是为Go语言在Linux系统上提供内存管理功能，这是实现高效、稳定运行的必要条件。




---

### Var:

### adviseUnused

在Go语言的运行时（runtime）中，mem_linux.go是负责内存管理的代码文件，其中的adviseUnused变量用于控制Linux系统上内存管理的策略。

具体来说，adviseUnused的值通常为1，表示告诉操作系统在释放内存时使用MADV_FREE标志，即将内存页标记为未使用，并通知系统它们可以立即回收，但仍然保留在进程地址空间中。这可以在稍后重新访问时避免昂贵的页表操作，同时如果程序需要使用这些内存，也可以更快地进行重新分配。

换句话说，通过使用adviseUnused变量来控制操作系统对内存的释放方式，可以在一定程度上提高程序的性能。但需要注意的是，该变量的使用也可能会影响程序的内存占用和性能，因此应针对具体情况进行调整。



## Functions:

### sysAllocOS

sysAllocOS是Go语言中runtime包中的一种函数，位于runtime/mem_linux.go文件中。它的作用是向操作系统请求一段内存空间并返回其地址。

在Go语言中，sysAllocOS函数通常用于在堆上分配一些未被使用的内存空间。在Go语言中，一个程序的内存由三部分组成：代码区、数据区和堆区。堆区是程序可用内存最大的存储区域，因为它没有固定的空间大小，而是按照需要进行动态分配和释放。通过sysAllocOS函数分配的内存空间一般会用于堆的扩张和内存泄露的补偿等操作。

在操作系统中，请求内存空间一般需要系统调用，而sysAllocOS就是通过系统调用实现内存分配的。在Linux系统中，sysAllocOS函数通过调用mmap函数申请一段空间来满足内存分配的需求，并且将这段空间与程序的虚拟地址空间进行对齐。

需要注意的是，Go语言实现的堆的分配器不仅仅依赖于sysAllocOS，还依赖于一些其他机制，如对空闲内存的管理、垃圾回收、内存分配的算法和统计等。因此，如果你想深入理解Go语言的内存管理机制，不仅需要了解sysAllocOS的实现，还需要了解runtime包中其他函数的作用。



### sysUnusedOS

sysUnusedOS函数是在linux系统上清理未使用内存的函数，它的作用是将操作系统中的闲置内存释放回来。这个函数会使用系统调用来获取闲置内存的大小，然后使用系统调用madvise使操作系统将这些内存释放。这个函数是在runtime中使用的，用于管理和控制内存的分配和释放。

具体地说，sysUnusedOS函数的实现流程如下：

1. 调用系统调用mincore，获取所有内存页面的状态（是否被使用）以及闲置内存页面的数量。
2. 根据闲置内存页面的数量，计算需要释放的内存数量。
3. 调用系统调用madvise，告诉操作系统释放这些内存。

这个函数的目的是尽可能地将未使用的内存返回给操作系统，以便其他进程或系统可以利用这些内存。这也有助于减少内存泄漏的风险，并提高系统的稳定性和性能。



### sysUsedOS

sysUsedOS是用于获取操作系统级别的内存使用情况的函数。该函数会调用Linux系统下的getrusage系统调用，获取当前进程和其所有子进程的系统资源使用情况，包括CPU时间、内存大小、IO操作等信息。然后将该信息处理后返回当前进程的内存使用情况。

具体来说，sysUsedOS获取的是当前进程的系统内存使用情况，包括由进程已申请占用的内存和系统为该进程分配的缓存等内存。该函数会将这些内存数值相加，得到当前进程在系统中已占用的内存大小。

在Go语言运行时中，sysUsedOS函数会被Runtime调用，用于统计程序的内存使用情况，并将其用于垃圾回收等操作。通过sysUsedOS获取到的内存使用情况，运行时可以更好地控制程序内存的消耗，避免因为内存过度占用而导致程序异常或崩溃的情况发生。



### sysHugePageOS

在Go语言的runtime包中，mem_linux.go文件的sysHugePageOS函数主要用于在Linux操作系统下分配和释放大页面。

大页面是指更大的页面大小，通常是2MB、4MB甚至更大。这与常规页面的大小（通常为4KB）相比，可以显着提高内存访问速度和管理效率。在使用大页面时，可以减少内存碎片，提高缓存局部性和减少TLB（Translation Lookaside Buffer，翻译后备缓存）开销。

在sysHugePageOS函数中，使用了Linux系统调用（syscall）来请求操作系统为进程分配或释放大页面。这里使用的具体系统调用是mmap和munmap函数，其中mmap函数可以将一块物理内存映射到进程的内存空间中，而munmap函数可以取消对所有或一部分该进程地址空间中一块已映射内存区域的映射。

该函数还使用了内核命令sysctl来获取系统大页面的大小，以便正确分配和释放大页面。在函数的实现中，还包含了一些错误处理和调试信息输出的逻辑。



### sysNoHugePageOS

sysNoHugePageOS是一个在Linux系统上用来检查是否启用了HugePage的函数。HugePage是Linux内核中的一种内存管理机制，可以将大块的内存页映射到一个连续的虚拟地址区域中。这样的好处是可以减少页表的数量，提高内存访问的效率。

在Go中，通过sysNoHugePageOS函数，可以检查系统当前是否启用了HugePage的功能。如果启用了，Go运行时会将HugePage功能禁用。这是因为在Go运行时的内存管理中，页表的数量已经被优化到最少，再启用HugePage反而会对性能产生负面影响。

因此，sysNoHugePageOS函数在Go运行时的内存管理中发挥了重要作用，确保了内存管理的高效性和稳定性。



### sysFreeOS

sysFreeOS是Go语言运行时的一个函数，用于释放操作系统上的物理内存。它是在Linux操作系统下实现的，对应的文件是mem_linux.go。

在Go语言中，当程序申请内存时，运行时会通过操作系统向系统申请物理内存。当程序不再需要内存时，运行时会将它们标记为可回收的，但并不会立即释放给操作系统。相反，运行时将内存保存在一个特殊的“空闲列表”中，以便后续再次申请内存时可以直接使用。这种处理是为了性能优化和减少系统调用次数。

但是，当Go程序运行时间较长或者申请的内存比较大时，空闲列表会变得越来越长，可能导致运行时不能在物理内存中为新的内存请求提供足够的空间。此时，运行时就需要调用sysFreeOS函数来释放一部分物理内存，以便重新分配给当前或其它的进程使用。

sysFreeOS函数使用的核心操作是调用madvise系统调用，该调用会通知操作系统可以释放物理内存，并从操作系统的空闲内存池中删除这些内存页面。此时，运行时可以重新分配这些页面，以满足新的内存请求。

需要注意的是，sysFreeOS函数实际上只是建议操作系统可以释放内存，但不一定能够立即得到释放。操作系统会根据内存使用情况和系统负载等因素来决定是否立即释放内存，这与函数本身并没有关系。但sysFreeOS函数的调用可以提醒操作系统尽快处理内存释放请求，从而减少程序运行时间的等待。

总之，sysFreeOS是Go语言运行时中一个非常重要的函数，它可以帮助Go程序释放操作系统上的物理内存，提高系统的性能和稳定性。



### sysFaultOS

sysFaultOS函数是位于go/src/runtime/mem_linux.go文件中的函数，它的作用是将当前进程的地址空间映射到一个不可访问的页面，从而导致段错误（segmentation fault）。

在操作系统中，每个进程都有自己的虚拟地址空间，应用程序中的指针指向的是虚拟地址空间中的地址。当应用程序尝试访问一个未映射的或已被释放的内存地址时，操作系统会产生一个“段错误”或“访问冲突”。这种情况通常是程序出现bug或内存泄漏的结果。

sysFaultOS函数通过调用mmap系统调用将当前进程的地址空间映射到一个不可访问的页面，并将该页面的权限设置为不可读、不可写、不可执行，从而模拟出一个段错误的情况。这对于测试和排除虚拟内存管理相关的问题非常有用。

需要注意的是，sysFaultOS函数只能在Linux系统上使用，其他操作系统需要相应的实现。另外，由于该函数涉及到操作系统底层，不推荐在正式的应用代码中使用。



### sysReserveOS

sysReserveOS函数的作用是将虚拟地址空间的一段内存标记为保留并不能被使用，以便操作系统可以将其映射到物理内存。在Linux系统中，程序可以通过mmap系统调用将虚拟内存地址映射到物理内存地址，而sysReserveOS在程序初始化时为堆空间保留了一段虚拟地址空间，防止其他程序占用这一地址空间。

具体来说，sysReserveOS函数会调用mmap系统调用，在虚拟地址空间中标记一段内存为保留状态，并返回指向该内存区域的指针。该函数还会将保留区域的指针和大小存储在全局变量sysStat中，以备后续使用。

在Go运行时中，sysReserveOS函数是在堆空间的初始化过程中调用的。堆空间初始化时，将会调用该函数，并获得一段已保留的地址空间用于后续的堆分配。另外，该函数还会在子句处理的过程中为栈空间分配保留的虚拟内存空间。

总之，sysReserveOS函数是Go运行时中的一个关键函数，它保证了程序具有足够的虚拟地址空间和在物理内存上的映射，以支持堆和栈的分配。



### sysMapOS

在Go语言中，sysMapOS函数是位于runtime/mem_linux.go文件中的一个函数。它的主要作用是将OS特定的映射标志（如PROT_EXEC、PROT_READ、PROT_WRITE等）映射到对应的POSIX标识（如PROT_EXEC、PROT_READ、PROT_WRITE等）。它将操作系统的标志与POSIX标志进行了映射。这个函数的实现主要是因为不同系统之间的内存映射标志可能不同，需要做一些转换来保证内存映射的一致性。

该函数的实现细节如下:

- 首先，它根据给定的OS掩码，设置POSIX掩码的初始值。

- 然后，它使用一个可变数据类型（变量名为flag）迭代处理OS掩码和POSIX掩码的对应关系。在每次循环中，它会检查当前OS掩码是否等于POSIX掩码中最低位的掩码值，如果是，那么它就将POSIX掩码中最低位的掩码值与flag的当前值进行或运算，并将该POSIX掩码从POSIX掩码中删除。接着，它将OS掩码从OS掩码中删除，并以此重复该过程，直到所有的OS掩码都被删除。

- 最后，它将flag的值返回，这个值是将所有的POSIX掩码与OS掩码进行了转换后的结果。

总之，sysMapOS函数主要是将OS特定的映射标志映射到对应的POSIX标识，方便内存映射的使用，并保证在不同系统之间的内存映射标志的一致性。



