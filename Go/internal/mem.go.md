# File: mem.go

mem.go文件是Go语言标准库内置的一个文件，位于/go/src/internal/syscall/unix/目录下，主要负责封装系统调用mmap()和munmap()，以便Go语言的运行时可以使用 Unix系统上的内存映射。

具体来说，mem.go文件提供了以下功能：

1. mmap()函数：该函数封装了系统调用mmap()，实现了内存映射功能，主要用于开辟一段指定大小的虚拟内存，并返回指向该内存的指针。此处的内存可以是匿名的，也可以基于已经存在的文件。

2. munmap()函数：该函数封装了系统调用munmap()，实现了对进程已映射内存的解除映射。主要用于释放之前通过mmap()映射的内存空间。

3. Copy函数：该函数用于将源内存空间的数据拷贝到目标内存空间，主要用于将高位地址的内存数据拷贝到低位地址的内存空间中。

4. Zero函数：该函数用于将指定内存空间清零，主要用于保护敏感内存数据并减少内存信息泄漏的风险。

以此，mem.go文件是Go语言标准库中非常重要的内置文件之一，它提供了内存映射的基本实现，并为Go语言提供了更好的内存管理机制，提高了程序的性能和安全性。




---

### Structs:

### sharedMem

在 Go 语言中，`sharedMem` 结构体定义在 `mem.go` 文件中，是一个虚拟内存映像结构体。它的作用是在不同的 goroutine 之间共享内存，实现数据在共享内存上的通信。具体来说，结构体中定义了两个关键属性：

- `fd` ：表示操作系统建立的一个虚拟文件，在内存中映射一段连续的空间。
- `pagesize` ：表示操作系统中一页的大小，常用的大小通常为4KB。

在内存映像建立之后，程序可以通过修改内存映像中的数据来实现协作。由于多个协程可以并发地读写同一块内存，所以需要使用并发安全的方式读写数据，防止出现数据竞争的情况。

因此，在共享内存时需要使用原子操作或锁的方式防止数据竞争，这是 Go 语言中的高级主题，需要仔细考虑才能避免问题。



### sharedMemHeader

在go/src/internal/mem.go中，sharedMemHeader结构体被用作共享内存块的元数据结构，用于跟踪内存块的大小、内部布局和已使用的字节数。

该结构体定义如下：

```
type sharedMemHeader struct {
	size    uintptr // 共享内存块的总大小
	offset  uintptr // 已使用字节数
	ref     uint32 // 引用计数
	kind    uint32 // 内存块的种类
	id      uintptr // 共享内存块的唯一ID
	version uint32 // 共享内存块的版本号
	_       uint32 // 用于内存对齐
}
```

下面是各字段的作用：

- size：共享内存块的总大小，以字节为单位。
- offset：已使用字节数，表示当前共享内存块中已经使用的字节数。
- ref：引用计数，表示当前有多少个进程或线程正在使用该内存块。每当一个新进程或线程使用该内存块时，该计数值会加1，释放时会减1。当该计数值为0时，该内存块将被销毁。
- kind：内存块的种类，表示该内存块的用途，比如是用于共享缓存还是共享数据等。该值通常由使用该内存块的进程或线程在创建时指定。
- id：共享内存块的唯一ID，用于在多个进程或线程之间唯一标识共享内存块。
- version：共享内存块的版本号，用于在修改共享内存块时进行协调其它进程或线程的同步。每次修改时，该值将递增。

通过在共享内存块的开头创建该结构体，并将其作为共享内存块的元数据，可以有效地跟踪内存块的使用情况，从而保证进程或线程之间的共享内存块的正确性和安全性。



## Functions:

### sharedMemSize

sharedMemSize这个函数的作用是计算共享内存的最大大小。

在Go语言中，共享内存是指多个并发的goroutine可以同时访问的内存区域。在使用共享内存时，需要考虑如何保证数据的同步和一致性，避免数据竞争问题。

sharedMemSize函数计算的是系统支持的共享内存的最大大小。这个值的计算与操作系统有关，通常是通过系统调用获取的。

在Linux系统中，共享内存大小的限制由内核参数shmmax控制。sharedMemSize函数通过读取shmmax参数的值来计算共享内存的最大大小，具体实现代码如下：

```
func sharedMemSize() (s int64, err error) {
        const (
                sysShmMax = 0x1fffFFFF // 64-bit only
                sysPage   = 1 << 12    // Assume 4KB pages
        )
        const testSize = sysShmMax / 2

        // use sysconf to get page size and getrlimit for SHARED_MEMORY_OBJECTS
        page := int64(syscall.Getpagesize())

        rl := &syscall.Rlimit{}
        if err := syscall.Getrlimit(syscall.RLIMIT_MEMLOCK, rl); err != nil {
                return 0, err
        }

        // Rounded to page size.
        if uint64(testSize) > rl.Cur {
                testSize = int64(rl.Cur) & ^(page - 1)
        }

        // Try to create a shared memory object of testSize.
        // We cannot use os.Create because we need a file descriptor.
        fd, err := syscall.Shmget(syscall.IPC_PRIVATE, uintptr(testSize), 0600|syscall.IPC_CREAT|syscall.IPC_EXCL)
        if err != nil {
                // EINVAL means testSize was too big for available memory.
                if err == syscall.EINVAL {
                        return int64(rl.Cur) / 2, nil
                }
                return 0, err
        }

        // remove the shared memory so we can re-use the test name if link fails.
        defer syscall.Shmctl(fd, syscall.IPC_RMID, nil)

        // try to map it
        p, _, e := syscall.Mmap(fd, 0, int(testSize), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
        if e != nil {
                return 0, e
        }
        defer syscall.Munmap(p)

        // limit rlimit to something we can actually use
        if rl.Cur > int64(sysShmMax) {
                rl.Cur = int64(sysShmMax)
        }

        for s = rl.Cur; s > testSize; s -= s / 8 {
                // try to map object
                _, _, e := syscall.Mmap(fd, 0, int(s), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
                if e == nil {
                        return s, nil
                }
                if e != syscall.ENOMEM {
                        return 0, e
                }
        }

        return testSize, nil
}
```

该函数首先通过调用系统调用获取系统的页大小和内存锁定限制（rlimit），然后尝试创建一个大小为testSize的共享内存对象，如果创建失败，则逐渐减小大小（每次减小1/8）再次创建，直到成功为止。如果无法成功创建，则返回已经授权的内存锁定限制的一半。最后，返回计算出的共享内存的最大大小。



### sharedMemTempFile

sharedMemTempFile函数是一个临时文件生成函数，主要用于在内存中创建一个临时文件来共享数据。具体的作用包括以下几点：

1. 在内存中创建一个文件对象。

2. 将文件对象映射到共享内存。

3. 返回文件对象和共享内存句柄。

通过创建一个临时文件并将其映射到共享内存，可以方便地在多个进程之间共享数据。由于共享内存是在操作系统层面实现的，因此比其他 IPC 方式更高效，可以在高并发情况下提供更好的性能。

在实现方面，sharedMemTempFile函数会先创建一个独立的临时文件对象，然后调用 mmap 函数将文件对象映射到共享内存中，最后将文件对象和共享内存句柄一起返回。使用完毕后，可以通过 munmap 函数取消映射并删除临时文件，以释放内存和文件系统空间。

总的来说，sharedMemTempFile函数是一个方便且高效的共享内存实现方式，在需要跨进程或跨机器共享数据时可以考虑使用。



### header

header函数是Go语言内置包internal/mem中的一个函数。该函数的作用是返回给定切片的指针，以及该切片的长度和容量。具体来说，header函数的代码实现如下：

```
// header returns the slice header for the given data pointer, length, and capacity.
func header(data unsafe.Pointer, length, capacity int) []byte {
    return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
        Data: data,
        Len:  length,
        Cap:  capacity,
    }))
}
```

函数的输入参数包括切片的底层数据指针data、切片的长度length以及切片的容量capacity。函数的输出结果是一个类型为[]byte的切片，该切片的指针指向data，长度为length，容量为capacity。

从代码实现中可以看出，该函数使用了unsafe和reflect两个包。unsafe包提供了操作指针的功能，而reflect包则提供了类型信息的反射功能。在header函数内部，使用unsafe.Pointer将data指针转换为*reflect.SliceHeader类型的指针，然后再将该指针转换为[]byte类型。这样一来，就可以得到底层数据指针、长度和容量组成的切片。

总的来说，header函数的作用就是将底层数据指针、长度和容量组装成一个切片类型的数据结构。这个数据结构可以在某些场景下非常有用，例如需要手动管理内存时，可以使用该函数得到切片的底层数据指针，以及切片的长度和容量，从而进行更加精细的内存管理。



### valueRef

valueRef函数定义在go/src/internal/mem.go文件中，它的作用是返回一个指向值的指针。该函数使用了unsafe包中的Pointer()函数来实现，因此它具有较高的危险性，应该谨慎使用。

具体来说，当我们需要对某个值进行特定的操作时，有时需要使用指针来操作该值的内存空间。在一般情况下，如果我们要获取某个值的指针，写法通常是`&value`。但是，在一些特定的场合下，例如使用unsafe包中的函数，我们需要获取到一个指向值的指针。这时候，就可以使用valueRef函数来实现。

valueRef函数接收一个值作为参数，并返回一个指向该值的指针。它的实现方式非常简单，就是直接调用unsafe.Pointer()函数，将该值的地址转换为一个不安全的指针类型。这样，我们就可以使用该指针进行各种需要获取指针的操作了。

需要注意的是，valueRef函数只能用于操作本地变量或结构体中的字段，而不能用于操作数组、切片、映射、指针等类型。否则，可能会破坏内存安全性，产生不可预期的后果。因此，在使用valueRef函数时，必须要非常小心，并尽可能避免使用它。



### valueCopy

valueCopy函数的作用是将源切片中的元素复制到目标切片中。它是在内部包中定义的，说明该函数被用于Go语言内部的实现，而不是供外部程序使用的。

valueCopy函数实现了一个非常高效的算法来复制切片。它避免了在复制时分配新的切片，并且还可以跳过源和目标切片之间的额外空间，从而提高了效率。

valueCopy函数接收两个参数：src和dst。src是源切片，dst是目标切片。valueCopy函数根据src和dst的长度选择正确的复制策略，并将src中的元素复制到dst中。如果src和dst的长度都是零，则不执行任何操作。

总之，valueCopy函数是一个高效的切片复制函数，它在go内部使用，不面向用户，但对于理解Go语言切片背后的实现机制有很大的帮助。



### setValue

setValue这个函数是在go/src/internal/mem/mem.go文件中定义的。该函数的作用是将src指向的数据拷贝到dst指向的位置，同时返回拷贝的字节数。如果src或dst为nil，则此函数会panic。

具体来说，该函数使用了unsafe包中的指针类型，因为这个函数没有任何保证，可能会访问或修改指向任意内存地址的指针。因此它应该非常小心地使用。

该函数第一个参数是dst指针，第二个参数是src指针，第三个参数是拷贝的字节数。在函数的最开始，使用了非常小心的指针类型转换，将src和dst指针转换为unsafe.Pointer类型，这样就可以使用unsafe包中的函数进行内存操作。

然后，该函数通过判断src和dst之间的偏移量和长度大小，来判断是否会产生数据重叠或超出边界的情况。如果出现这种情况，就会调用runtime.memmove函数来复制数据。否则，就直接进行内存拷贝。

最后，该函数返回拷贝的字节数。这个返回值通常用在测试代码中，在测试时，该函数可以检查拷贝的字节数是否与预期相等。



### setValueLen

setValueLen是一个内部函数，用于将一个给定的无符号整数值编码为变长整数格式，并将结果写入缓冲区中。

在内存使用方面，变长整数格式是一种紧凑和节省空间的表示方法。它使用小于8个字节的空间来编码0到2^64-1之间的任意整数值。

当把一个无符号整数值编码为变长整数时，setValueLen函数将根据该值的大小选择一个特定的编码方案，并将它写入缓冲区。它采用的编码方案可以涵盖8种不同的大小类别，从一个字节到十个字节不等。

在Go语言的标准库中，setValueLen函数主要用于内部处理，比如在编码二进制协议时。它涉及到字节的移位、位掩码的计算等操作，因此较为复杂。但在使用时，我们只需知道该函数可以将整数值编码成紧凑的变长整数格式即可。



