# File: nih.go

nih.go是Go语言中runtime的一个重要组成部分，属于非内联函数(noinline)。它的主要作用是为Go语言的内存管理系统提供基础支持。

具体来说，nih.go文件中定义了大量的基础函数，用于实现Go语言的内存分配、GC（垃圾回收）等功能，包括以下几个方面：

1. 内存分配：nih.go文件中定义了一系列函数，如mallocgc、allocm、largeAlloc等，用于分配内存空间；

2. GC回收：nih.go文件中也定义了一些函数，如gc、scavenge、markroot、marknonlocal等，用于垃圾回收和标记、扫描等操作；

3. 对象分配：Go语言中的对象分配是由运行时环境负责的，nih.go中实现了包括指针类型、数组类型在内的常见对象的分配和释放；

4. 内存映射：Go语言的底层内存管理基于内存映射实现，nih.go也提供了相应的函数实现；

5. 安全栈保护：Go语言的协程（goroutine）采用用户态栈（stack）实现，nih.go实现了安全栈保护的相关函数。

总之，nih.go是Go语言中重要的一部分，支撑了整个Go语言的内存管理和安全运行环境，保证了Go语言的运行效率和稳定性。




---

### Structs:

### nih

nih这个结构体是Go语言中重要的垃圾回收器的核心数据结构之一。nih实现了Go语言的标记-清除垃圾回收算法，用于回收不再使用的内存。它主要包括以下几个重要字段：

1. arenas: 用于管理堆空间的一组内存区域，每个区域大小为2的30次方字节，可以在需要时动态增长。

2. tiny, small, large: 表示三种不同大小的对象分配缓存，来提高对象分配的效率。

3. free: 用于维护空闲内存块的链表，用于提高内存分配的效率。

4. marksweep: 表示标记-清除垃圾回收器中的标记和清除阶段，负责标记所有可达的对象以及回收不可达的对象。

除此之外，nih结构体还包括了许多其他的重要字段和方法，用于实现垃圾回收器的各种功能，如内存分配、对象扫描等等。由于垃圾回收是Go语言中非常重要的机制，因此nih结构体是整个引擎中最为核心的数据结构之一，对于理解Go语言的内存管理和垃圾回收机制有着非常重要的作用。



### NotInHeap

在Go语言中，GC的机制需要对内存地址进行扫描和标记，对于某些特殊情况下的内存地址，我们需要特殊标记一下，以确保GC不会对其进行扫描和标记。而在runtime包中，就定义了一个叫做NotInHeap的结构体，是用来做这种特殊标记的。

具体来说，当我们需要获取一块内存地址时，如果需要避免GC对其进行扫描和标记，就可以使用runtime包中的Malloc函数，这个函数会返回一个指针类型的内存地址，如果这个地址需要避开GC的扫描，可以将其强制转换为NotInHeap类型，如下所示：

```
p := runtime.Malloc(n)
np := (*runtime.NotInHeap)(unsafe.Pointer(p))
```

这样一来，np就是一个标记为NotInHeap的内存地址了，GC会自动忽略这个地址而不进行扫描和标记，保证了程序的正确性和性能。

在nih.go中，NotInHeap结构体就起到了这个作用，它没有任何字段，只是作为一个标记类型，用于表示某个内存地址需要避开GC扫描的情况。此外，还定义了几个函数，用来对NotInHeap类型的值进行操作和转换，如下所示：

```
func AllocNotInHeap(size uintptr) unsafe.Pointer {
    return unsafe.Pointer(sysAlloc(size, &memstats.other_sys))
}

func FreeNotInHeap(ptr unsafe.Pointer, size uintptr) {
    sysFree(ptr, size, &memstats.other_sys)
}

func NotInHeap(uintptr) NotInHeap {
    return NotInHeap{}
}

func (n NotInHeap) Add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
    return unsafe.Pointer(uintptr(p) + x)
}

func (n NotInHeap) Uintptr(p unsafe.Pointer) uintptr {
    return uintptr(p)
}
```

这些函数的作用，就是对NotInHeap类型的值进行操作和转换，以满足程序中对内存地址的相关需求。



