# File: mheap.go

mheap.go文件是Go语言运行时包中（runtime）的一个文件，它的作用是实现Go语言的堆内存管理。具体来说，它定义了mheap结构体和相关的方法，用于在Go语言的运行时环境中跟踪、分配和释放堆内存。

mheap结构体是Go语言堆内存管理的核心，它包含了多个字段，如arena、free、allspans、busybits、cache、central、sweepgen、scavengestat等等。这些字段共同组成了堆的状态，并提供了一些方法来实现堆内存的分配、释放、回收等操作。

除了实现堆内存管理，mheap.go文件还负责实现了一些与堆内存有关的操作，如创建heapArena、从heapArena中分配内存、设置heapArena状态、通过堆内存索引访问内存块等等。这些操作的实现，可以让Go语言的运行时系统更加高效地管理堆内存，提高程序的性能和稳定性。

总的来说，mheap.go文件是Go语言运行时系统中负责实现堆内存管理的核心文件之一。通过对该文件的深入理解，可以帮助开发者更好地理解Go语言的内存管理机制，从而有效地编写高质量的Go语言程序。




---

### Var:

### mheap_

mheap_变量是Go语言运行时系统中的一个结构体变量，它是内存分配器（heap）的主要数据结构，用于管理堆内存的分配、回收等操作。

mheap_变量中包含了许多重要字段，如freelist等，它们分别用于记录当前空闲的内存块、已分配内存块的地址范围等信息。mheap_变量还包含了一些方法，用于堆内存的分配和回收。其中最重要的是mheap_.alloc函数，它用于申请和分配一块内存。

mheap_变量的作用非常重要，它直接决定了Go程序的内存使用和性能表现。通过对mheap_变量进行优化，可以提高Go程序的内存分配效率，避免内存泄露等问题的发生，提升程序的整体性能表现。

总之，mheap_变量是Go语言内存分配器的核心之一，它的作用不可忽略。了解它的结构、字段和方法，对于深入了解Go语言的内存管理机制和性能优化非常有帮助。



### mSpanStateNames

mSpanStateNames这个变量在runtime包下的mheap.go文件中，它是一个保存堆中内存块状态名称的映射表。该表提供了各种堆内存块状态的名称，包括空闲、已分配、已扫描、已标记、已清除等状态。

这个变量的主要作用是方便调试和排除堆相关的问题。例如，在调试过程中，当需要查看某个内存块的状态时，可以通过访问这个映射表来查找其相应的状态名称。同时，这个映射表还可以帮助程序员更好地理解程序中与堆相关的处理逻辑。

除了在调试过程中使用之外，这个映射表还可以在代码中进行条件判断和处理。例如，当内存块状态为已分配时，程序可以进行相应的内存释放操作，以避免出现内存泄漏等问题。

总之，mSpanStateNames这个变量在runtime包的mheap.go文件中扮演着非常重要的角色，它是一个方便调试和处理堆相关问题的常量映射表。



### gcBitsArenas

在go语言中，gcBitsArenas变量表示从堆中分配给MHeap管理的page的数量。在具体实现中，gcBitsArenas是一个32位的原子变量，用于记录分配给MHeap的page数量，其中最高的4位用于表示当前正在进行垃圾回收的goroutine的数量，低28位则表示分配给MHeap的page数量。垃圾回收时，gcBitsArenas用于计算堆大小和下一次垃圾回收的时间。

具体来说，当gcBitsArenas的值低于最大值时，系统将继续为MHeap分配新的page，直到值达到最大值。上限的初始值为系统的总RAM的一半或mheap.max，取两者的最小值。当gcBitsArenas的值达到了上限后，分配新的page时只会从空闲page池中获取。在垃圾回收时，gcBitsArenas用于计算需要回收的内存大小，以及下一次垃圾回收的时间，以便在系统负载较轻的时候进行垃圾回收。

总之，gcBitsArenas变量是MHeap管理的page数量的记录器，在分配内存和进行垃圾回收时都会参与计算，是Go语言运行时的重要组成部分。






---

### Structs:

### mheap

mheap结构体是Go的运行时系统中堆的主要组成部分之一，用于管理动态分配的内存。其主要作用有以下几个：

1. 维护堆的状态：mheap结构体中包含了一些成员变量来记录堆的状态，例如当前堆的大小、分配的字节数和分配的次数等。

2. 管理内存分配：Go程序中的内存分配都是由运行时系统来管理的，mheap结构体中定义了一些方法来分配内存、释放内存和调整内存的大小。

3. 实现垃圾回收机制：Go使用一种叫做Mark-and-Sweep的垃圾回收算法来自动回收不再使用的内存。mheap结构体中的一些方法用于将内存块标记为“已使用”或“未使用”，以便垃圾回收器识别哪些内存可以被回收。

4. 实现堆的扩容操作：当分配的内存超过了堆的大小时，mheap结构体会触发堆的扩容操作，从而保证程序可以继续分配更多的内存。

总之，mheap结构体是Go运行时系统中非常重要的一部分，它负责管理程序的内存分配和垃圾回收，是整个系统的核心。



### heapArena

heapArena是Go语言运行时的一种数据结构，它表示了操作系统在虚拟内存中分配给Go语言堆的一段连续内存区域。heapArena内部包含了一些管理这段内存区域和其中对象的信息，如内部的bitmap用来记录哪些内存块被使用，以及空闲内存块列表等等。

在Go语言程序运行时，一开始Go语言运行时会从操作系统申请一些内存块作为堆空间，这些内存块会被分成多个heapArena，作为Go语言堆的子区域使用。这些heapArena之间可以相互独立进行内存分配和回收，从而提高堆的性能和可用性。

同时，heapArena还可以进行线程间的内存管理和通信，以及协调GC的过程。比如，在进行GC的时候，heapArena会被分成多个小块进行扫描和标记，这些操作会涉及到多个线程的协同，并且需要考虑内存分配和回收的顺序，以及对GC的最终影响等问题。因此，合理地设计和使用heapArena结构体可以对整个Go语言运行时系统的性能和可靠性产生重要影响。



### arenaHint

arenaHint结构体在Go的内存管理中用于帮助调整堆的布局，以优化内存使用效率。在Go中，堆是通过mheap结构来管理的，而arenaHint结构体是mheap中的一个字段。

具体来说，arenaHint结构体包含的是一个指针，这个指针指向的是最近使用的堆空间的位置。当需要分配新的堆空间时，会尽量在这个位置的附近分配，以避免在堆的不同区域反复跳转，从而提高内存访问效率。

另外，arenaHint结构体还包含了一个记录上次GC清理时堆的总大小的字段，以及一个标记堆是否处于扩展状态的标志。这些信息也会被用于优化堆的布局。

总之，arenaHint结构体在Go的内存管理中具有重要作用，它通过记录最近使用的堆空间位置以及当前堆的状态等信息，来优化堆的布局，从而提高内存使用效率。



### mSpanState

mSpanState是一个标识位的枚举类型，在mheap.go中的mSpan结构体中使用，用于描述一个mSpan对象所处的状态。mSpan对象表示一小块内存，用于分配给对象使用。mSpanState枚举类型的定义如下：

type mSpanState int

const (
    _ mSpanState = iota
    mSpanDead
    mSpanInUse
    mSpanManual
    mSpanFree
)

mSpanDead表示这个mSpan对象已经不能再使用，因为其中存储的对象已经被释放了。

mSpanInUse表示这个mSpan对象当前正在被使用，其中存储的对象是活动状态的。

mSpanManual表示这个mSpan对象是由用户手动分配和释放的。

mSpanFree表示这个mSpan对象可以被再次使用，其中存储的对象已经被释放。

通过使用mSpanState枚举类型，mheap.go可以轻松地判断一个mSpan对象当前的状态，从而有效地管理内存的使用。



### mSpanStateBox

mSpanStateBox是用于表示一组连续的内存块，通常称之为"span"。在Go语言中，内存的分配和释放都是通过mheap来进行管理的。mSpanStateBox结构体是用于描述一个span当前的状态，包含了与这个span相关的统计信息、指向其他span的指针等。

具体来说，mSpanStateBox结构体的定义如下：

type mSpanStateBox struct {
    addr      uintptr        // span的起始地址
    sizeclass uint16         // span的sizeclass
    state     uint16         // span的状态（如 "MSpanInUse"、"MSpanFree"等）
    spanclass uint8          // span的所在对象的类别（如"spanClassHeap"、"spanClassStack"等）
    elemidx   uint8          // span内部可用空间偏移量
    unused    uint32         // 无用字段，保留
    startobj  uintptr        // span内部第一个可用对象地址
    nelems    uintptr        // span内部可用对象的总个数
    limitobj  uintptr        // span内部最后一个可用对象的地址
    next      *mspanStateBox // 指向下一个span的指针
    allocCache uintptr       // 存储上一次分配对象时的指针，实现连续分配机制
    _         [12]byte      
}

其中的主要字段含义如下：

- addr：span的起始地址
- sizeclass：span的对齐大小（内存块的大小不一定是恰好等于请求的大小）
- state：span的状态，如"MSpanInUse"表示正在使用，"MSpanFree"表示空闲等
- spanclass：span所在的对象类别，如"spanClassHeap"表示动态分配内存，"spanClassStack"表示栈内存等
- elemidx：指示span内部可用空间偏移量
- startobj：span内部第一个可用对象的地址
- nelems：span内部可用对象的总个数
- limitobj：span内部最后一个可用对象的地址
- next：指向下一个span的指针，用于形成链表以便快速操作
- allocCache：记录上一次分配对象时的指针，实现连续分配机制。



### mSpanList

mSpanList结构体是用来管理span的双向链表。span是一段连续的内存区域，被分割成大小相同的块，用于分配小对象的内存池。

mSpanList结构体包含以下字段：

- head: 指向链表头部的span指针。当链表为空时，head为nil。
- tail: 指向链表尾部的span指针。当链表为空时，tail为nil。

mSpanList结构体的方法包括以下几个：

- empty(): 返回链表是否为空。
- push(): 将span添加到链表头部。
- remove(): 从链表中移除指定的span。
- pop(): 从链表头部移除一个span，如果链表为空则返回nil。

mSpanList结构体在内存分配过程中，用于维护不同大小的span链表。当需要分配某个大小的内存块时，会先检查对应的span链表是否为空。如果为空则需要从内存池中获取一段新的span，并添加到对应的span链表中。如果链表不为空，则从链表头部移除一个span进行内存分配。当一个span中的所有块都被分配完后，它会从链表中移除，并返回给内存池中，等待下一次使用。



### mspan

mspan结构体在Go语言的运行时系统中扮演了很重要的角色，它是一个Go语言运行时堆管理机制中管理内存的基本单元，表示一个内存区域。

具体来说，mspan结构体存储了一段连续的内存空间的元数据信息，包括该内存区域的起始地址、大小、状态、绑定的对象等等。相邻的多个mspan结构体可以组成一个连续的内存区域，称为一个heap，用于分配和管理内存。

mspan结构体的定义如下：

```go
type mspan struct {
    // 数据起始地址
    startAddr uintptr
    // 数据大小
    npages uintptr

    // 代表该mspan所属的heap
    heap    *mheap

    // 代表该mspan的下一个和上一个mspan，用于构成链表
    next    *mspan
    prev    *mspan

    // span 使用的的栈（解释：运行时栈）
    stack   [(_StackCacheSize / _PageSize)]uintptr
    stackIdx int32
    // Span 中的堆栈缓存是否可用
    unusedsince int64

    // 处于当前mspan之上的mspan，比它们地址更高
    above *mspan

    // GC 相关信息
    state spanState // 所在的状态
    // span 所属的堆类别
    // 正在被分配的语言对象所在的堆。堆栈对象属于哪种类型的堆，使用的 span 也属于这个堆
    // 不同类别的 heap 在不同的 goroutine 之间分配，所以 HeapArena 可以在使用的 goroutine 之间传递
    heapid    uintptr
    // 当前 span 中的已分配对象大小
    allocCount uint32
}
```

mspan结构体中的一些重要字段注释已经写在了代码中，这里重点解释一下state字段和heapid字段。

- state字段：代表了当前mspan所处的状态，包括：

  - mSpanDead：该mspan已经死亡，可以被回收
  - mSpanFree：该mspan为空闲状态，可以被分配给新对象
  - mSpanManual：该mspan状态由用户管理，不参与垃圾回收
  - mSpanInUse：该mspan正在被使用，其中包括mSpanInUse、mSpanManual和mSpanStack类别的mspan

- heapid字段：代表当前mspan所属的heap的ID，即其所在的堆类别。一个heap代表了一组大小相同的mspan，属于同一堆类别的mspan之间可以互相分配和释放对象，不同堆类别之间的mspan则不能相互影响。Go语言运行时系统通过动态调整heap的个数和每个heap所管理的内存大小，来适配不同大小和使用情况的内存分配需求。



### spanClass

spanClass是一个用于描述span（内存块）分类的结构体，它定义了不同大小的span所属的分类，用于管理内存分配和回收。

在golang的堆内存管理中，使用了类似于操作系统的内存分配算法。将整个堆空间按固定大小的小块划分，其中每个小块称为span，每个span的大小都是2的幂次方。spanClass结构体用于管理每个span的分类，记录每个大小的span的数量和可用列表，方便内存分配和回收时快速找到空闲的span。

具体而言，spanClass结构体的定义如下：

```go
type spanClass struct {
    sizeclass  uint8  // span大小分类
    nelem      uint16 // span中元素数量
    npage      uint16 // span占据的页数
    run        struct {     // 可用的span列表
        lock   mutex
        head   *mspan  // span链表头
        tail   *mspan  // span链表尾
        nelem  uintptr // span列表中元素数量
    }
}
```

其中，sizeclass表示span的大小分类，nelem表示一个span中元素的数量，npage表示span占据的页数，run表示当前span大小的可用span列表。

当需要内存分配时，根据需要分配的内存大小确定所需的span大小分类，然后在对应的spanClass可用列表中查找是否有空闲的span可供分配。如果列表为空，则需要从中心缓存（central）或全局堆（heap）中获取可用的span，如果还没有可用的span，则需要扩大堆大小以分配新的span。

当回收内存时，可以根据span的大小将已经使用的span归还到对应的spanClass可用列表中进行复用。这样可以避免频繁地向操作系统申请和释放内存，提高内存分配和回收的效率。



### arenaIdx

在Go语言的运行时中，mheap.go文件中定义了一个名为"arenaIdx"的结构体。这个结构体的作用是表示堆区域的索引，并且在运行时中用于跟踪和管理堆区域。

具体来说，arenaIdx结构体是一个数组，其中的每个元素都代表了一组堆区域（即一个连续的内存块），每组堆区域都有固定的大小。在程序启动时，堆区域会被初始化，并且各个索引值会被设置为相应的大小。每次需要增加或者释放堆区域时，arenaIdx结构体的索引值会被更新，以确保堆区域能够被正确地分配和释放。

除了上述作用之外，arenaIdx结构体还可以用于优化堆区域的分配。在Go语言中，堆区域的分配是通过按照固定大小分配内存块来实现的。由于不同大小的内存块可能是有不同的用途的，因此可以通过使用arenaIdx结构体来调整堆区域的分配策略，以适应不同的应用场景。例如，对于需要大量小内存块的应用程序，可以将较小的堆区域分配给较小的内存块，从而避免浪费内存。



### spanAllocType

spanAllocType是runtime中用于描述MHeap中span的分配情况的数据结构，其定义如下：

```go
type spanAllocType uint8

const (
    spanAllocated spanAllocType = iota // The span is allocated for some use.
    spanFreelist                      // The span is on a free list.
    spanStack                         // The span has a stack allocated for it.
    _                                 // skip (was spanActive)
    spanManual                        // The span was allocated via sysAlloc, but is managed by the Go heap.
    spanDead                          // The span is no longer being used.
    spanMax                           // The end of the spanAllocType enum values.
)
```

该结构体用于描述span在MHeap中的分配状态，具体包括以下几种类型：

- spanAllocated：span已被分配，用于某些目的；
- spanFreelist：span在一个空闲列表中，等待被分配；
- spanStack：span已经为其分配了一个新的栈；
- spanManual：span是通过sysAlloc分配的，但由Go堆管理；
- spanDead：span不再被使用。

MHeap中的spanAlloc数组存储了每个span的分配情况，可以快速地了解整个MHeap中可用的span数量，以及哪些span正在使用或等待分配。在调用内存分配器相关的函数时，会根据这些span的分配情况进行调度和分配。



### special

在Go的运行时系统中，每个线程都关联着一个M（Machine）结构体，其中包含了一些关键的信息和状态，例如goroutine的调度和执行（G）、内存分配器（mheap）、追踪系统（trace）等。其中，mheap.go这个文件中定义了mheap结构体，它是内存分配器的核心组件，负责为程序中的大多数对象分配和回收内存。

在mheap.go中，还定义了一个名为special的结构体，用于管理特殊的对象，例如栈空间、堆空间和mspan结构体。这些特殊的对象是不同于一般的对象，需要针对它们进行特殊的处理和管理。特殊对象的管理由mheap.special中的一些方法实现，例如acquiremspan()和findP()等。

特别地，heapBits和spanAllocCount这两个字段是mheap.special结构体的重要组成部分。heapBits是一个用于标记堆空间的位图，每个位代表一个固定大小的内存块，用于记录该内存块是否已被分配。spanAllocCount是一个计数器，用于记录当前span结构体的分配数量。

总而言之，mheap.special结构体起着重要的作用，负责管理特殊的对象，其中heapBits和spanAllocCount字段是其核心组成部分，帮助实现内存管理器的各种功能。



### specialfinalizer

在Go语言的垃圾回收机制中，当一个对象不再被引用时，其所占用的内存就会被回收。而有些对象需要在被回收之前执行某些操作，比如释放掉与其相关的资源。这时候就需要用到Finalizer（终结器）机制。Finalizer是一个回调函数，当一个对象被垃圾回收器标记为可回收时，其对应的Finalizer函数就会被调用。

在mheap.go文件中，specialfinalizer结构体定义了Finalizer的相关信息，包括回调函数、对象大小、类型标记等。这些信息将会在对象被标记为可回收时，与GC相关的代码一起进行处理。specialfinalizer提供了一种在GC期间自动执行Finalizer的机制，可以帮助开发者更方便地释放资源，避免内存泄漏。

特别需要注意的是，由于Finalizer的执行时间是不确定的，因此不能在其中进行任何有副作用的操作，比如访问全局变量或共享数据结构等。否则可能导致不可预测的错误。另外，因为Finalizer只在垃圾回收器执行时才会被调用，因此不能依赖Finalizer来释放资源。最好的做法是使用defer或者在对象不再使用时主动释放占用的资源。



### specialprofile

specialprofile是runtime中的一个结构体，用于描述特殊的内存分配情况，并提供对应的统计信息。具体来说，specialprofile包含了以下三个成员：

1. n：表示特殊内存分配的次数；
2. bytes：表示特殊内存分配的字节数；
3. name：表示特殊内存分配的类型的名称。

特殊内存分配指的是在runtime执行过程中，由于某些原因而发生的不同于正常情况下内存分配的情形，例如：

1. 程序中使用了cgo调用，需要向C库申请内存；
2. 调用了runtime库中的特殊函数，例如mmap和munmap；
3. 调用了GC(垃圾回收)相关的特殊函数，例如gcSweep、gcMark、startCleanupGoroutine等。

特殊内存分配虽然在程序中占比较小且不太常见，但由于其特殊性，需要对其进行特殊的处理和统计。因此，runtime使用specialprofile结构体来记录这些特殊内存分配的情况，并提供给开发者便利的接口，以便他们能够了解和优化程序的内存分配情况。



### specialReachable

specialReachable结构体是用于表示特殊可达对象的数据结构。特殊可达对象是指那些不能通过一般的指针分析算法来判断是否可以回收的对象。比如，在Go语言中，因为有finalizer机制，一个对象即使在程序中没有被引用也可能不会被回收，只有当其finalizer执行完毕后才会被回收。这种对象就是特殊可达对象。

specialReachable结构体中保存了特殊可达对象的相关信息，包括对象的地址和finalizer信息（如finalizer函数和对象大小等）。这些信息可以协助垃圾回收器判断对象是否需要被回收。

在垃圾回收算法中，当发现一个对象无法通过指针分析算法确定是否可以被回收时，便会将其标记为特殊可达对象，并将其信息保存到specialReachable结构体中。垃圾回收器在进行可达性分析时，会先处理特殊可达对象，然后再对其它对象进行分析。

通过使用specialReachable结构体，可以更准确地识别和处理特殊可达对象，提高垃圾回收器的效率和准确性。



### specialsIter

specialsIter结构体是用于迭代和管理mheap数据结构中特殊（special）堆块的迭代器。特殊堆块是指不常用的或者不常见的堆块，对于它们的管理需要一些特殊的处理。

specialsIter结构体中包括了以下字段：

- mheap：指向mheap结构体的指针；
- typ：表示特殊堆块的类型；
- span：指向特殊堆块对应的span结构体的指针；
- start：特殊堆块的开始地址；
- end：特殊堆块的结束地址；
- objIndex：表示迭代器当前指向的特殊堆块在其所在的span结构体中的位置。

specialsIter结构体的作用在于，它提供了一种迭代和管理特殊堆块的方式，可以让代码更加清晰和简洁。通过迭代器，可以逐个处理所有的特殊堆块，完成它们的管理和释放，同时不会漏掉任何一个堆块。

在实现中，specialsIter结构体被广泛使用在mheap的gcSweep函数中，用于扫描不常用的或者不常见的堆块，并将它们加入到空闲堆块链表中，以便后续的内存分配使用。



### gcBits

gcBits结构体是用来描述heap中每个block和span的gc标记信息的，它是一个可变长度结构体，具体实现如下：

```
type gcBits struct {
    // 省略其他字段

    // gcmarkBits用于标记该spans中每个对象是否被标记过
    gcmarkBits    *gcBitsArena // bits for GC marking; Same size as maxSpans*objectsPerSpan/8.

    // gcmarkBitsForSweep用于标记该spans中每个对象在清扫阶段是否需要被清理
    gcmarkBitsForSweepCache *gcBitsArena // gcmarkBits leaf cache for sweeping; nil if there is no active sweep
    gcmarkBitsForSweep      *gcBitsArena // gcmarkBits for sweeping; Same size as the largest size class. If there are
    // too many size classes to allow for a reasonable size allocation, then
    // this field is nil and we use gcmarkBitsForSweepCache instead.

    // 描述span的状态和属性
    state   mSpanState
    spanclass spanClass
    sizeclass uint8
    // 省略其他字段
}
```

gcBits中最重要的字段是gcmarkBits和gcmarkBitsForSweep，它们用于标记heap中每个对象的gc标记信息。在gc标记阶段，gcmarkBits表示该span对应的所有对象是否被标记过；在清扫阶段，gcmarkBitsForSweep表示该对象是否需要被清理。

采用这种方式来表示gc标记信息，可以让gc标记和清扫的时间和空间复杂度都得到优化。同时，由于heap中对象数量非常大，因此使用可变长度的gcBits结构体可以更好地节约内存空间，提高效率。

除此之外，gcBits还有其他一些字段用于描述span的属性和状态，比如state、spanclass和sizeclass等。这些字段用于帮助runtime管理heap中的span和block，以提高内存的利用率。



### gcBitsHeader

gcBitsHeader是runtime/mheap.go中定义的一个结构体，用于描述堆对象的GC元数据信息。每个堆对象在分配时都会被标记为"black"或"white"两种状态之一，当垃圾回收程序进行扫描时，就会根据这些状态来判断哪些对象是垃圾，需要被回收。

gcBitsHeader结构体包含了以下字段：

```
type gcBitsHeader struct {
   data  *uint8 // 指向元数据位图的指针
   marked uintptr // 该对象是否被标记为黑色（已扫描）的标记位
}
```

其中，data字段指向了该对象对应的元数据位图的起始地址，而marked字段则表示该对象是否被标记为黑色（已扫描）的标记位。当垃圾回收程序需要扫描某个堆对象时，会首先根据该对象的地址查找对应的gcBitsHeader结构体，然后再根据data字段找到该对象在元数据位图中的对应位，从而判断该对象的状态。

通过gcBitsHeader结构体，垃圾回收程序不仅能够准确地判断堆对象的状态，同时还可以高效地检索和更新这些信息，提高了垃圾回收的效率。



### gcBitsArena

gcBitsArena是用于标记清除（mark-and-sweep）垃圾回收算法中的一个数据结构。

在垃圾回收过程中，需要知道哪些内存区域正在使用，哪些已经被释放。为了方便表示和操作这些内存区域的状态，每个内存区域都被划分为若干个页面（page），每个页面对应一个bit。如果该页面正在被使用，则对应的bit为1，否则为0。这些bit被组织为一组二进制数，称为位图（bitmap）。

gcBitsArena结构体就是用于存储这些位图的数据结构。具体来说，每个arena对应着一段内存区域，arena中包含了多个页表（page table），每个页表包含了多个页（page）对应的bit信息。还有一些其他的元数据信息，比如arena起始地址和大小等。

在垃圾回收过程中，通过访问gcBitsArena中的bit来判断内存区域的状态，从而实现内存的标记和清除。



## Functions:

### set

在Go语言中，mheap是用来管理堆(heap)的一个数据结构。set函数是mheap的一部分，它的作用是将给定的arena作为新的heap arena加入到mheap中。

具体来说，set函数接收三个参数：addr、size和spanClass。其中，addr是新的arena的起始位置，size是arena的大小，spanClass是指arena中的span的大小类别。

在set函数中，主要做以下四个操作：

1. 根据新的arena的大小和spanClass，计算出需要多少个span。

2. 将这些span的信息加入到mheap的spanAlloc中，以便后续的内存分配可以使用这些span。

3. 将新的arena的信息加入到mheap的arenaAlloc中，以便mheap管理arena的使用情况。

4. 根据新的arena的地址和大小，将其加入到mheap的free中，以便后续的内存分配可以从这里进行。

总的来说，set函数的作用是将新的arena加入到mheap中，并更新mheap中有关arena和span的信息。这些信息将被用来支持后续的内存分配操作。



### get

get函数是mheap的主要函数之一，用于从堆中获取一个能够存储 n 个字节的内存块，函数签名如下：

func (h *mheap) get(n uintptr) *mspan

具体的作用如下：

1. 根据 n 对齐获取一个 sizeclass。sizeclass 是将小于 32KB 的内存按固定尺寸分成的若干类别的集合。

2. 从 spanAlloc 对应的 mspan 列表中取出一个 mspan。

3. 如果 spanAlloc 为空，那么就通过 h.refill 和 h.allocSpan 函数来重新获取mcentral 数据结构的 span 缓存。

4. 设置 mspan 的 state 为 mSpanInUse 表示已经被使用。

5. 将 mspans 和它所对应的内存块的 bitmap 更新。

6. 统计当前堆区占用的字节数，如果当前占用的字节数大于了最大堆区容量的 4/5，那么就会触发 mheap.grow 函数用于扩容堆内存。

总的来说，get 函数的作用是从堆中获取一块内存块来分配给应用程序使用，这是分配内存的核心实现之一，而且这个函数的执行效率对内存使用的好坏会产生重大的影响，因此应用程序应对它的调用进行最优化的操作。



### base

在Go的运行时(runtime)中，mheap.go文件中的base函数是用于计算堆的起始地址的函数。堆是程序运行时用于分配内存的区域，可以理解为一个动态的内存池。调用base函数可以获取堆的基址，也就是起始地址。

这个函数主要用于将内存地址转换为堆索引，从而便于分配和回收内存。堆索引是一个与具体物理地址无关的数值，通过对堆索引的计算和操作，可以实现对内存的动态管理。

具体来说，base函数内部首先获取了mheap对象的base地址。然后通过对当前内存地址减去base地址的计算，得到了堆索引。这样就可以直接对该堆索引进行内存分配操作。

在Go语言中，内存的回收是通过垃圾回收(garbage collection)机制来实现的。垃圾回收主要是将不再使用的内存进行清理和回收。通过base函数获取到堆的基址，可以保证回收过程不会对其他的内存区域产生影响，并且可以高效地进行内存回收操作。



### layout

在Go语言中，mheap.go文件中的layout函数是用于获取内存分配器（mheap）的布局信息的函数。在Go语言中，内存分配器使用一个类似于小块堆的数据结构来管理内存分配和回收。然而，这个小块堆的布局信息是动态变化的，因此需要进行实时更新和计算。

layout函数中的主要作用是封装了内存分配器的堆布局信息，该布局信息中包含了对内部小块质量和内部小块可用性的描述。具体来说，layout函数会遍历内存分配器中的所有链表（heaps）并收集有关每个堆块（heap span）的信息。然后，它将这些信息存储在一个数据结构中，并返回给调用者。这个数据结构中包含有关每个堆块的大小、大小类别、指向下一个堆块的指针和堆块中每个对象的大小等信息。

此外，layout函数还可以根据分配器在堆中的地址范围计算堆布局的指针大小和对齐设置，这在实际的内存分配过程中非常重要。在Go语言中，内存分配器的设计是非常精细和高效的，因此layout函数在调用时能够快速计算出内存分配器的堆布局信息。

总之，layout函数是Go语言内存分配器中一个关键的组成部分，它提供了有关内存分配器布局的重要信息，并且可以帮助我们更好地了解Go语言中的内存管理机制。



### recordspan

`recordspan`函数的主要作用是将一块内存分配给堆，以供运行时使用。它针对给定的堆区域（`*_heap`），创建一个记分卡，用于跟踪当前区域的统计信息，例如当前使用、空闲大小等等。这些统计数据可以在图形界面中可视化，以便更好地了解运行时堆的使用情况。

当recordspan函数成功记录了给定区域的统计信息之后，它将再次调用runtime.MCentral_CacheSpan函数将该区域添加到运行时的central空间中。这个central空间是一个类似于内存池的缓存区，用于快速分配内存和回收内存。稍后，这个区域就可以被使用并且经由central空间来管理。这有助于提高运行时的效率，减少内存分配和回收的次数和复杂性。

总之，`recordspan`函数是整个运行时系统非常重要的一个功能，它为堆管理和内存分配提供了关键的支持。



### makeSpanClass

在go语言中，mheap.go文件定义了内存管理的堆(heap)数据结构。makeSpanClass是该文件中定义的一个函数，用于将给定大小的堆块(span)分类到合适的span class中，并返回相应的class id。该函数的作用包括以下几个方面：

1. 为了更高效地管理内存，go堆里的span被分成多个不同的大小类(class)，每个大小类都包含若干固定大小的span。makeSpanClass函数就是用于确定一个span应该属于哪个大小类。

2. makeSpanClass函数根据span的大小，计算出对应的span class id。这个id将用于在内存分配时快速定位所需的span class，以减少搜索时间。

3. makeSpanClass函数还计算出了当前span class所需要的堆块数量及大小，并将它们返回给调用者。这样，内存管理器就可以预分配足够的堆块，以避免频繁调整堆大小带来的开销。

总之，makeSpanClass函数是go语言堆内存管理的关键函数之一，它帮助我们更有效地管理内存，提高程序性能。



### sizeclass

在 Go 语言运行时系统中，mheap.go 文件中的 sizeclass 函数定义了用于分配内存的大小类别。大小类别是一组固定大小的内存块，使得每个内存块都足够小，以便分配给小型对象，同时还能够更好地管理内存。

sizeclass 函数的作用是计算给定大小对象所需的大小类别。它采用一个整数参数 size，表示请求分配的对象大小，并返回一个代表该对象所属的大小类别的整数。

该函数实现了“round-to-power-of-two”算法，它将对象大小调整为最接近且小于等于 2 的幂的大小，并使用简单的公式来计算大小类别。这个公式对于每个大小类别都是不同的，因为每个大小类别涵盖的对象的大小不同。

sizeclass 函数的结果将用于查找与此类别匹配的空闲内存块，并将其分配给请求的对象。该函数的实现是 Go 运行时系统中的一个重要组成部分，用于管理内存各方面的内部细节。



### noscan

mheap.go文件中的noscan函数是用于在堆上分配对象时标记对象为非扫描状态的函数。当一个对象被标记为非扫描状态时，GC将不会扫描该对象，因为该对象已被标记为不再需要进行垃圾回收。这可以提高GC的效率，因为GC可以跳过这些对象而不必检查它们是否需要进行垃圾回收。

具体地说，noscan函数使用了一个变量runtimeType，它记录了对象的类型信息，包括对象大小、指针数量、GC标记等信息。当一个对象被标记为非扫描状态时，noscan函数会将该对象的runtimeType中的gcflag字段设置为类型不需进行扫描（GC不需要扫描该对象），然后标记该对象为已分配状态。

当GC需要进行垃圾回收时，它会根据对象的runtimeType中的gcflag字段来判断对象是否需要进行扫描和回收。如果gcflag为0，表示对象需要进行回收；如果gcflag为GC不需要扫描，则可以跳过该对象而不进行扫描和回收，从而提高GC的效率。

总之，noscan函数的作用是标记对象为非扫描状态，从而减少GC的扫描和回收时间，提高程序的性能和效率。



### arenaIndex

在Go语言的运行时系统中，mheap.go文件定义了分配堆内存的相关函数和数据结构。arenaIndex()是其中一个辅助函数，用于计算分配的内存块大小所对应的arena的编号。

在Go语言中，内存分配是以arena为单位进行的。一个arena是一个内存连续的区域，大小为8MB，其中包含多个内存块。每个内存块的大小都是2的幂次方（最小为8字节，最大为arena大小），并且内存分配是按照堆栈的方式进行的。

arenaIndex()函数的作用是计算给定的内存块大小所处的arena编号。这个函数用一个简单的公式来计算：对于一个大小为n的内存块，它所处的arena编号为i = log2(n) - 3。其中，log2(n)表示以2为底n的对数，并减去3是为了将前面的8字节内存块计算在内。

例如，如果要分配一个16字节的内存块，arenaIndex()函数会计算出它所处的arena编号为1。如果要分配一个4096字节（4KB）的内存块，arenaIndex()函数会计算出它所处的arena编号为6。

这个函数的作用在于帮助Go语言的运行时系统快速地确定分配内存块所处的arena，从而更加高效地进行内存分配。



### arenaBase

在Go语言中，为了管理内存的分配和回收，运行时系统使用了一些数据结构来维护内存的使用情况。其中包括了一个叫做mheap的结构体，它表示了整个堆（heap）的状态，包括空闲内存块链表、已分配内存块链表、各种参数等。

而arenaBase函数就是mheap结构体中的一个方法，它的作用是将一个地址转换为对应的堆区域的起始地址。

具体来说，Go语言中的内存分配使用了多个arena，也就是堆区域，每个arena都是一个固定大小的内存区域，包含了可用于分配的内存块。对于每个arena，都有一个起始地址，而该函数的作用就是给定一个任意地址，找出其所在的arena的起始地址。

在具体实现上，该函数使用了一些运算来确定输入地址所处的arena的起始地址，同时还对输入地址的有效性进行了检查，以避免程序出现意外错误。这样的设计也便于对内存的管理，以及对堆区域的划分和分配。



### l1

在Go语言的runtime包中，mheap.go文件是与内存堆分配和管理器相关的。在mheap.go文件中，l1这个func是内存堆的第一级缓存，它主要的作用是通过分离器的基本大小和区间来维护一个对象的池。

具体来说，l1函数的作用如下：

1. 初始化l1缓存：将缓存池的结构体数组初始化成一个固定的长度，为每个元素分配一段内存区域。

2. 获取l1缓存：从l1缓存池中获取适当大小的对象。如果缓存为空，则会从下一级缓存或者堆上分配内存，并将分配的对象放入l1缓存中。

3. 将对象返还给l1缓存：当一些对象不再被使用时，l1会将这些对象返还给缓存池中，并将其标记为空闲状态，以便下一次使用。

总的来说，l1函数的作用在于减轻内存分配器的压力，提高内存分配的效率，并且缓存池的管理方式也是内存池技术的一种实现。它使得Go语言的runtime包中的内存堆分配与管理器更加高效、健壮和稳定。



### l2

在Go语言的运行时环境中，mheap.go文件是堆（heap）内存分配器的实现。而其中的l2函数则是用来计算堆空间二级索引表的大小，在heap初始化时会调用l2函数来计算堆空间二级索引表的大小，以便为堆空间分配足够的内存。

在Go语言的内存分配器中，堆被分成多个连续的、大小相等的span（类似于内存页）以便有效地利用内存空间。而span的大小取决于堆空间大小，为了保证span的大小可以被整除，堆空间大小会被调整为2^k的形式。一级索引表通过将span按照大小从小到大依次链接起来，以便在分配和回收时能够快速地找到符合要求的span。而二级索引表则是在一级索引表的基础上进一步优化空间分配，通过按照span大小将堆空间划分为连续的段，使得空间分配更加灵活和高效。l2函数就是用来计算二级索引表的大小的。

具体来说，l2函数根据堆空间总大小计算出span大小，然后再根据span的大小来计算出堆空间总共需要多少个span。为了让每一段的大小尽可能大，l2函数会将堆空间划分为不同的大小段，然后对每一段进行按大小存储的span的数量的计算，以此来计算出二级索引表的大小。最终，l2函数就返回了二级索引表的大小信息。

总之，l2函数在Go语言的运行时环境中被用来计算堆空间二级索引表的大小，是一项非常重要的工作，可以让堆分配更加高效和灵活。



### inheap

inheap是一个用于检查指针是否指向堆内存的函数。

在Go语言中，内存分配是通过堆来实现的，当程序需要分配内存时，会向堆申请一段连续的内存块，然后将这段内存分配给程序使用。而在Go语言中，采用了垃圾回收机制来自动回收不再使用的内存，经常需要检查指针是否指向堆内存。

inheap函数就是用于检查指针是否指向堆内存的函数。它接收一个指针作为参数，然后通过遍历内存块的方式，查找该指针是否指向了堆内存。如果指针指向了堆内存，则返回true，否则返回false。

inheap函数的实现主要依赖于mheap结构体中的arenaStart、arenaUsed和arenaEnd等字段。arenaStart和arenaEnd字段分别指向堆内存的起始地址和结束地址，arenaUsed字段用于记录堆中已经分配的内存大小，通过这些字段，inheap函数可以遍历整个堆内存并查找指定的指针是否存在于堆内存中。

总的来说，inheap函数的作用是用于检查指针是否指向堆内存，这在垃圾回收机制中是非常重要的。



### inHeapOrStack

在Go语言中，程序在运行时需要动态地分配和释放内存，其中内存是由堆(heap)和栈(stack)两个部分构成的。堆是一块动态分配的内存区域，用于存放程序中所有的动态分配对象。栈是线程私有的内存区域，用于存储函数的局部变量、函数参数和返回值等。

在mheap.go文件中，inHeapOrStack这个函数的作用是判断指定的地址addr是否位于堆(heap)或者栈(stack)中。这个函数的实现比较简单，它首先检查addr是否在当前的堆(heap)中，如果是，则直接返回true。否则，它再检查addr是否在当前线程的栈(stack)中，如果是，则也返回true。如果同时不在堆(heap)和栈(stack)中，则说明addr不是有效地址，这时候返回false。

具体实现细节可以参考下面的代码：

```
// 判断给定的地址是否是堆或栈中的地址
func (h *mheap) inHeapOrStack(addr unsafe.Pointer) bool {
    if addr == nil {
        return false
    }
    if uintptr(addr) >= h.arena_start && uintptr(addr) < h.arena_used {
        return true // 在arena中，即在堆(heap)中
    }
    return mspanOf(addr) != nil // 在线程栈(stack)中
}
```

在Go语言中，内存管理是由运行时系统(runtime)负责的，而mheap.go文件中的代码就是实现堆(heap)相关的内存管理逻辑的。对于一个开发者来说，我们通常不需要关心这些细节，只需要使用Go语言提供的内存分配和释放函数(new和make等)即可。



### spanOf

在Go语言中，mheap.go文件中的spanOf函数用于查找给定地址处所属的堆span的指针。span是管理堆内存的重要数据结构，它负责管理一段连续的内存，包括分配、释放等操作，是堆上所有内存分配的基本单位。

具体而言，spanOf函数的输入参数是一个指向要查找地址的指针，它首先确定该地址所属的堆区间，并检查该区间是否已经对齐到堆的最小单位，如果不是则将其对齐。然后它通过计算该地址所在区间与堆起始地址之间的距离，以及区间所占用的span大小，计算出该地址在span内的偏移量，并返回该span的指针。

具体代码实现如下：

```
// spanOf returns the span containing page p or nil if p does not
// map to the heap.
func (h *mheap) spanOf(p heapAddr) *mspan {
    // If p is outside the range of the heap, ignore it.
    if p < heapStart || p >= heapEnd {
        return nil
    }
    // Round p down to the span boundary and compute how far it is from
    // the start of the heap.
    pIndex := uintptr(p) >> _PageShift
    spIndex := pIndex &^ (pagesPerSpan - 1)
    offset := pIndex - spIndex
    sp := h.spans[spIndex/pagesPerSpan]
    if sp == nil {
        return nil
    }
    // The span may not have been swept yet and thus may not be accounted
    // for in h.heap_live; we need to safely add an appropriate amount.
    // sp.bytes accounts for all bytes, including any bytes
    // that are not currently allocated.
    spBytes := uintptr(sp.npages) << _PageShift
    if sp.state == mSpanManual && sp.manualFreeList != 0 {
        spBytes -= _PageSize
    }
    heapLive := atomic.Load64(&h.heap_live) + spBytes
    if heapLive > h.heap_live {
        atomic.Store64(&h.heap_live, heapLive)
    }
    // Check that offset is in range, accounting for rounding.
    if size := spanClass(sp.npages << _PageShift).pageSize(); uintptr(offset) >= size {
        return nil
    }
    return sp
}
```

总的来说，spanOf函数的作用是根据给定地址找到其所在的堆span，方便对该地址所对应的内存进行操作。



### spanOfUnchecked

func spanOfUnchecked(addr uintptr) *mspan

作用：

spanOfUnchecked函数用于将给定的地址转换为mspan对象。在Go运行时系统中，mspan代表一块连续的内存。Go运行时系统使用mspan来管理堆内存，并跟踪哪些内存块是空闲的，哪些是已经被占用的。

在mheap.go文件中，spanOfUnchecked函数被用于获取某个地址所在的mspan对象。这个函数不会检查输入地址是否在堆上，因此它被称为"unchecked"函数。如果输入地址不在堆上，则函数可能导致未定义的行为。

简单来说，spanOfUnchecked函数的作用就是将一个地址转换为mspan对象，进而可以管理和跟踪这块内存的状态。



### spanOfHeap

spanOfHeap函数是Go语言运行时系统中mheap（内存堆管理）模块中的一个函数，其作用是找到地址所在的span（内存块）。

具体来说，当用户程序调用malloc、free等内存分配和释放函数时，Go语言运行时系统会调用mheap模块中的函数进行内存管理。其中，mheap模块会维护一张内存分配表，该表中记录了所有内存块的信息，包括每个内存块的大小、状态、对应的mcache等信息。

spanOfHeap函数的作用就是在内存分配表中查找给定地址所在的内存块（即span）。函数首先计算指定地址所在的内存块的起始地址，然后在内存分配表中进行查找并返回对应的span结构体对象。经过spanOfHeap函数处理，Go语言运行时系统能够更准确地控制内存分配和释放，从而保证程序的可靠性和稳定性。

总之，spanOfHeap函数是Go语言运行时系统内存管理模块中重要的内存查找函数，其作用是找到指定地址所在的内存块（span）以便进行后续内存管理操作。



### pageIndexOf

在Go语言的运行时系统中，内存管理是一个重要的组成部分。Mheap.go文件中的pageIndexOf函数用于计算由给定地址表示的内存页的索引。在堆管理中，内存被组织成固定大小的块，并按页对齐。一个页面通常是4KB或8KB，并且通常包含多个堆块。

pageIndexOf函数的作用是将给定地址映射到所属的页面索引。它使用右移位运算将地址按页大小进行对齐，并将结果除以页大小来计算页面索引。具体实现如下：

```go
func pageIndexOf(addr uintptr) uintptr {
    return addr / pageSize
}
```

该函数返回一个无符号整数，表示地址所属的页面索引。由于页面大小是固定的，因此此函数的结果也是固定的。在使用堆管理系统时，pageIndexOf函数可以帮助确定内存块的位置，并为垃圾回收器提供有用的信息。



### init

mheap.go是Go语言运行时（runtime）中的一个文件，控制着内存分配的堆（heap）部分。init函数是Go语言中一个特殊的函数，在程序启动时自动执行，用来初始化一些必要的数据。而在mheap.go文件中的init函数则用来初始化一个名为mheap的堆结构体。

该init函数主要完成以下几个任务：

1. 设置mheap的初始状态，包括堆的大小、span（堆内内存块的最小单位）的大小、空闲堆块链表等。

2. 向操作系统申请一块用于堆的内存空间，并将其映射到虚拟内存中。

3. 初始化一些用于锁定和解锁堆的互斥锁。

4. 初始化一些统计信息，如记录分配、释放堆内存的总次数、总量等。

总之，这个init函数用来初始化堆的相关信息，为后续程序的内存分配提供基础。在程序运行过程中，堆会不断地按需增长或缩小，重新分配内存空间，进行垃圾回收等操作，以满足程序的内存需求。



### reclaim

func reclaim(mheap *mheap) {
	reclaimed := mheap_sweep(mheap)
	if reclaimed == 0 {
		// Free the lowest order span from each nonempty list to try
		// to trigger scavenging of unused spans. This isn't always
		// possible, though, because the write barrier can prevent
		// reclamation of span bodies, in which case the spans will
		// remain unscavenged until the next time they're made
		// unreachable.
		for i := range mheap.free {
			 // 不断调用spanalloc_m，然后又flushmcache，又自动触发scavenge清理未使用的页，从而迅速释放和回收内存
			s := mheap_freeSpanLocked(mheap, i)
			if s != nil {
				mheap_->pages.scavengeStartGen -= s.npages
				goto done
			}
		}

		// We were unable to free anything at all. Force
		// mheap_.sweepgen to the empty past-the-end state.
		atomic.StoreUint64(&mheap.sweepgen, mheap_.sweepdone)
		// no spans swept, reset sweepers so they don't hold pointers to the free buffer.
		for i := 0; i < int(mheap.sweepers); i++ {
			mheap.sweepgen[i] = mheap_.sweepdone
		}
	}

done:
	// Note that since we don't hold the sweepgen lock,
	// it's possible that the sweepgen has changed since we
	// read it earlier. This is fine; it just means that we
	// may have missed some things to reap. We can just keep
	// operating on the current sweepgen.
	if debug.gcpacertrace > 0 || trace.enabled {
		getg().m.traceReclaimed(reclaimed)
	}
}

这是Go语言中runtime包中mheap.go文件中的reclaim()函数。它的作用是在垃圾回收期间回收内存。在垃圾回收的过程中，可能会有一些对象已经被标记为垃圾，但是它们仍然占用着内存，这些对象所占用的内存就可以通过reclaim()函数回收。该函数调用mheap_sweep()函数完成内存回收的过程。

如果没有任何可回收的内存，该函数会从非空的自由列表中获取最低位的span，以尝试触发清除未使用的span页的操作。但是这并不总是可能的，因为写入障碍可能会阻止span页体的回收，此时这些span将保持未使用状态，直到下次它们变得不可达为止。

在获取最低位的span后，函数会调用mheap_freeSpanLocked()函数释放和回收内存。该函数在内部调用spanalloc_m()函数来获取新的span，然后调用flushmcache()函数来刷新mcache，最后触发scavenge()函数来清理未使用的页，从而迅速释放和回收内存。如果释放成功，mheap.pages.scavengeStartGen 将减去释放的npages值。

如果在此过程中没有释放任何东西，函数将强制将sweepgen设置为空状态。对于没有扫描的spans，它们的扫描器会重置，这样它们就不会持有指向自由缓冲区的指针。

最后，如果启用了gcpacertrace或跟踪，则函数将调用getg().m.traceReclaimed()函数来追踪管理的内存。





### reclaimChunk

reclaimChunk函数是在运行时进行垃圾回收的过程中，针对堆内存进行的一种清理机制。其主要作用是将无法使用的连续内存块标记为可用状态，以便后续的内存分配可以复用这些空间，以减少系统内存的占用率。

reclaimChunk函数是由mheap的freeSpan和sweepSpans函数调用的，它的作用是处理freeSpan和sweepSpans中的空闲簇。freeSpan是一个链表，跟踪着所有空闲簇，每个空闲簇由连续的空闲页组成。sweepSpans则是跟踪着所有正在使用的簇，每个簇包含了连续的已分配页或不连续的前缀空闲页。

reclaimChunk函数的具体实现如下：

首先，它会检查Golang堆的可达性标记。如果标记为0，那么代表当前堆已经被标记为未被使用，即所有的对象都可以被回收。这时，reclaimChunk函数会处理所有空闲簇，并将它们添加到可用空闲簇的链表中。

其次，它会遍历所有的空闲簇，检查它们的页帧的头部是否被标记为可达性。如果是，说明该空闲簇中的内存空间还有一些对象被引用，所以这个空闲簇不能被回收。否则，该空闲簇将被标记为空闲，并加入到可用空闲簇的链表中。

最后，reclaimChunk函数会将可用空闲簇按照大小进行排序，并设置为对应的、可用于分配的span。使用这些空闲簇，堆内存就可以复用旧的空间，使得内存占用更加高效。

总结：reclaimChunk函数处理Golang堆内的空闲簇，将它们标记为可用，并添加到可用空闲簇的链表中。它将被使用的空间和空闲空间分开，并根据空闲空间的大小对其进行排序以提高内存利用率。



### manual

mheap.go 文件中的 manual 函数是用于手动触发垃圾回收的函数。垃圾回收是 Go 语言中的一种内存管理机制，它会在程序运行过程中自动回收不再使用的内存。然而，有时候我们希望手动触发垃圾回收，这时候就可以使用 manual 函数。

manual 函数的作用就是触发一次垃圾回收，它会将程序中不再使用的内存标记为可回收的垃圾，然后将这些垃圾进行回收。手动触发垃圾回收有时可以提高程序的性能，因为它可以清除不再使用的内存，释放内存空间。

在 manual 函数内部，会调用 mheap 类型的 manualGC 方法来进行垃圾回收。manualGC 方法会遍历整个堆，并对其中的对象进行标记和回收，这个过程跟自动垃圾回收的过程类似。

需要注意的是，在大多数情况下，应该让程序自动进行垃圾回收，手动触发垃圾回收只在必要的情况下使用。因为手动触发垃圾回收会导致程序的性能下降，同时也会增加垃圾回收器的负担，影响整个程序的运行效率。



### alloc

在go/src/runtime中，mheap.go文件中的alloc函数是用于从堆中分配一块内存的函数，它的作用是为正在执行的程序动态分配内存，满足程序的内存需求。

该函数的具体实现涉及到了golang运行时中的堆管理机制。在程序运行过程中，堆会维护一组分配了但是未被释放的对象的内存块。当程序通过代码逻辑请求新的内存时，alloc函数会根据当前堆的状态，从堆中取出一块内存，并且分配给程序。

需要注意的是，alloc函数在分配内存时，不会直接从系统申请内存，而是会从堆的空闲内存块中分配一块足够大的内存。如果当前堆中没有足够的内存块，则alloc函数会调用mheap.grow函数从系统中申请一块新的内存块，并添加到堆中，然后再从中分配一块内存片段给程序。

总而言之，alloc函数提供了一个动态内存分配的机制，用于满足程序在运行时的内存需求，支持程序的正常运行和数据的存储。



### allocManual

allocManual函数是Golang运行时系统中的一个重要函数，它用于手动分配内存，主要用于一些特殊场景下的内存分配，例如内存池、GC小对象等。

该函数的具体实现是在内存堆对象中实现的。在调用该函数时，会首先检查是否有可使用的内存块，如果有可用的内存块，则从内存块中分配指定大小的内存；否则，会从操作系统中分配一块新的内存页作为新的内存块，然后从新的内存块中分配指定大小的内存。

该函数通常被内存分配器中的其他函数所调用，以满足不同的内存分配需求。相比于其他内存分配函数，allocManual函数更加灵活，可以手动控制内存的分配和释放。但需要注意的是，由于该函数是手动分配内存，因此需要开发者自己掌握内存分配和释放的管理，避免内存泄漏或悬垂指针等问题的出现。

总之，allocManual函数是一个非常重要的Golang运行时函数，用于实现手动内存分配，为其他内存分配器和GC机制提供支持。



### setSpans

setSpans函数在runtime中的mheap.go文件中，用于为给定的页框设置相应的span。

在分配内存时，heap会把物理内存分割成大小为1-32个页框的区域。为了管理这些区域，heap会为每个区域分配一个span，用于跟踪该区域的使用情况，包括已分配的块数和未使用的块数。setSpans函数的作用就是将给定页框的span中的描述信息设置为新值。

setSpans函数的参数包括一个堆heap、一个起始页框号base和页框数n，以及4个相关信息的集合：free和scav的bitmap和spans和large等描述信息的指针。具体来说，它首先根据base计算出要修改的span在spans的偏移量，然后更新相应的描述信息。对于一个最常见的情况，free bitmap和scav bitmap将在同一个span中，因此一个单独的调用就可以应对这两个bitmap。更新large bitmap需要单独调用。

总之，setSpans函数是管理heap中内存分配的重要组成部分，它确保所有的内存使用都可以追踪并避免内存泄漏。它主要作用是管理span的描述信息，包括free bitmap和scav bitmap等。



### allocNeedsZero

在Go语言中，allocNeedsZero这个函数位于runtime的mheap.go文件中，它的作用是判断是否需要对新分配的内存进行初始化。当我们使用Go语言的new关键字或make函数进行内存分配时，新分配的内存是未初始化的，即其内容是不确定的。

为了确保在分配的内存区域被使用之前可以正确地初始化，Go语言的allocator会使用该函数检测是否需要对这块内存进行清零操作，以确保所有空间的位都是从0开始的。

该函数的实现比较简单，它接收两个参数size和flags，其中flags是一个包含了一些标记的参数，这些标记用于指定内存分配的属性。该函数主要是判断flags中是否包含了_MZero标记，如果包含则认为需要对新分配的内存进行清零，否则认为无需进行清零操作。

需要注意的是，由于清零操作会对内存性能造成一定的影响，所以只有在必要的情况下才进行清零，以避免不必要的开销。在一些特定的场景下，如在大量重复分配相同大小的内存块时，也可以使用缓存来提高内存分配的效率。



### tryAllocMSpan

tryAllocMSpan函数是Golang runtime中mheap（memory heap）模块的一部分，它的作用是尝试从给定的mspan中分配一个或多个内存块。

在Golang runtime中，内存块通常以mspan的形式进行管理。每个mspan代表一段连续的虚拟内存区域，由连续的页框组成。每个mspan都有一个对应的bitmap，用于跟踪分配的内存块和空闲的内存块。

tryAllocMSpan函数的主要工作如下：

1. 从给定的ms内部bitmap中查找连续的空闲内存块。如果存在，则更新bitmap，将其标记为已分配。
2. 如果未找到连续的空闲内存块，函数返回nil，提示调用者需要尝试其他的mspan。

该函数是内存管理的核心部分，它负责在申请内存时分配空闲的内存块。因为内存分配是Golang并发架构的必要组件，因此该函数必须高效地运行。通常情况下，tryAllocMSpan函数能够快速地判断出哪些内存块已经被分配，而不需要扫描整个内存区域。



### allocMSpanLocked

mheap.go中的allocMSpanLocked函数的作用是在堆中分配一个新的span并将其添加到span列表中。具体的步骤如下：

1. 首先，该函数检查中间子堆是否可以分配所需的span。这里中间子堆是一组heapArena的集合，用于分配堆空间。如果中间子堆可以分配新的span，则将分配的span添加到该子堆的span列表中，并返回该span的起始地址。

2. 如果中间子堆不能分配所需的span，则从heapArena列表中查找空闲的arena。如果找到了，则分配新的span并将其添加到该arena的span列表中，并返回该span的起始地址。否则，如果没有找到空闲的arena，则调用grow方法，向操作系统请求更多的虚拟内存。

3. 在grow方法中，首先会尝试从操作系统中申请一个huge page，在Linux系统上这一般是2MB或1GB大小的物理内存页。如果成功，则将huge page切分为多个span，并添加到heapArena中。如果申请huge page失败，则会申请一小块内存，然后将其分配为一个span。

总的来说，allocMSpanLocked的主要作用是分配新的span，为堆的管理提供基础的支持。它是Golang内部堆分配机制的重要组成部分，是支撑Golang应用程序高效运行的关键。



### freeMSpanLocked

freeMSpanLocked函数的作用是将已释放的mspan对象归还给堆空间池。

在Go的堆管理实现中，由于内存分配所需的空间大小不一，为了能够高效地分配内存，Go使用了分级的mcentral和mspan对象。当一个mspan被分配出去使用后，当需要释放这部分空间时，就需要将该mspan对象归还给堆空间池，这样以后再需要分配相同大小的空间时，就可以直接从堆空间池中获取已经分配好的mspan对象，从而提高内存分配的效率。

freeMSpanLocked函数主要有以下几个操作：

1. 检查该mspan是否为中央缓存mcentral中的的span，如果是，则将该mspan从mcentral中移除；
2. 检查该mspan是否有未分配出去的空间块，如果有，则将这些空间块都归还给堆空间池；
3. 将该mspan对象标记为freed，并将其加入空闲mspan链表中。

其中，空闲mspan链表中的mspan对象可以用于后续的内存分配。由于在多线程环境中，需要保证多个goroutine能够安全地访问该空闲mspan链表，因此该链表不仅用于存储空闲的mspan对象，同时还包含了用于并发访问的锁对象。

通过这些操作，freeMSpanLocked函数实现了将已使用的mspan对象及其中的空间块归还给堆空间池，从而帮助提高了Go的内存管理效率。



### allocSpan

allocSpan是Golang运行时中mheap.go文件中的一个函数，用于在堆上分配连续的内存空间（称为span）。

这个函数会首先从heap中找到一个大小合适的span块，如果找到的span块的大小超过需要的大小，则会将这个span进行分割，剩下的部分会重新放回heap中。

然后，对于找到的span，会根据需要进行一些处理。比如，如果需要的大小超过了span的大小，那么会在span的周围创建一个新的span，先将原来的span的数据复制到新的span中，再将原来的span放回heap中，然后使用新的span作为需要的内存空间。

最后，如果有需要的话，allocSpan函数还会在span上加入一些元数据，比如标记这个span是否已经被分配等信息。

总之，allocSpan函数的作用是在堆上分配连续的内存空间，并进行一些处理。它是Golang运行时的一个重要组成部分，对于程序的内存管理和性能优化起到了重要的作用。



### initSpan

initSpan函数是在堆上初始化span的函数。在Go语言的自动内存管理中，堆是由mheap管理的。而span是mheap对堆的抽象，表示可用内存块的一段连续内存区域。

initSpan函数的具体作用是初始化一个新的span，对其进行清零操作，并将其插入到相应的span列表中。具体实现如下：

- 首先，该函数会将所给的span的所有内存块都初始化为0。
- 然后，会根据span的大小，将其插入到相应的span列表中。每个span列表表示一段内存大小相同的span，不同大小的span列表按照大小递增排列。
- 如果插入到span列表中的span的大小超过了一定阈值（由小于32页到大于等于1024页不等），则会将其根据spansizeidx缓存下来，以便后续重复使用。这样可以提高内存分配的速度。

总之，initSpan函数的作用是在堆上初始化一段新的连续内存区域，将其加入到相应的span列表中，并缓存下来以提高内存分配速度。



### grow

函数grow位于mheap.go中，它的作用是将堆的大小扩展到至少所需的大小。

在Go运行时中，mheap表示管理内存分配的堆。堆是由一系列MemoryBlock组成的，MemoryBlock可以分配给程序使用。

当需要更多的内存时，可能需要先增加堆的大小。Grow函数做的就是这个，它将堆的大小增加到至少所需的大小。具体来说，函数将扩展堆意味着将分配更多的MemoryBlock。这些MemoryBlock可能来自操作系统分配的新区域或者是再利用空闲的部分。在完成扩展步骤后，所有新的MemoryBlock将被添加到mheap中。

grow函数的过程涉及几个步骤，其中最主要的是增加堆的大小并更新相关的指针和标志。当堆需要被扩展时，函数调用sysAlloc来为堆分配新的内存区域。接着，函数使用updateHavocRatio方法来更新有害比率，这个比率反映了堆内存碎片的程度。最后，如果排除来自sysAlloc的新区域之后，空闲的内存还不足以满足需求，grow将会调用allocSlow方法来查找足够大的内存块。

总之，mheap.grow函数的作用是将堆扩展到至少所需的大小，使得内存分配可以继续进行。



### freeSpan

freeSpan函数(自由段函数)在go程序运行期间的内存管理中起到了非常重要的作用。它的主要作用就是将一个空闲的内存段(即未被占用的内存段)加入到内存池(即堆)中，以便后续的新对象分配使用。

具体来说，freeSpan函数的详细流程如下：

1. 判断输入freeSpan的参数是否合法（这里的参数是指一个包含了空闲内存段信息的mSpan指针）。
2. 获取freeSpan的heap指针，并根据此指针计算出该mSpan指针对应的地址空间的起始位置和大小。
3. 检查此地址空间的大小是否符合预期，即不能太大或者太小（一般来说，内存段的大小应该在[PAGESIZE, MaxMem/2]的范围内）。
4. 将当前的mSpan指针加到heap中空闲的段列表中，同时更新内存池中相应大小的空闲内存段的计数器。
5. 根据mheap.scheduler或mheap.central的状态，决定是否唤醒GC等其他的goroutine。

其中，mheap是go语言中的内存池，scheduler和central是它的两种状态，分别对应并发和串行的内存分配管理方式。此外，该函数还会检查内存池中的内存段是否已满，并根据具体情况采取不同的扩展策略来保证内存池能够满足程序运行期间新对象的分配需求。

总体来说，freeSpan函数的作用就是动态管理内存池中的内存段，以方便后续的新对象分配，并且在一定程度上保证了整个程序的内存使用效率。



### freeManual

`freeManual`函数是在堆上手动释放内存的函数，它的作用是释放一块未被标记为已分配的内存。

具体而言，当应用程序需要释放一块内存时，它会调用`free`函数。但是，在某些情况下，应用程序可能需要手动释放内存，例如当应用程序使用`unsafe`包中的指针时。在这种情况下，应用程序可能需要直接释放指针指向的内存而不是使用`free`函数。

`freeManual`函数提供了这个功能。它接收一个指针和一个大小，并将指针指向的内存标记为未分配，从而将内存释放回堆中。但需要注意的是，如果应用程序不小心释放了已经被分配的内存，会导致内存泄漏或内存损坏。

因此，在使用`freeManual`函数时需要特别小心，并仅在必要时使用。并且，使用`freeManual`函数时，要确保传递给它的指针确实是指向未分配的内存。



### freeSpanLocked

freeSpanLocked函数的作用是将一个空闲的span添加到对应的page heap中。当一个span上的所有对象被释放时，这个span就成为空闲状态了。这时候就可以通过freeSpanLocked将其添加到page heap中，以便后续的分配使用。

该函数的主要实现逻辑如下：

1. 从对应的page heap中获取需要添加的空闲span列表。
2. 遍历空闲span列表，如果找到与目标span相邻的空闲span，则将两个span合并成一个。如果没有则直接将目标span添加到空闲span列表中。
3. 将更新后的空闲span列表更新到对应的page heap中。

通过freeSpanLocked函数将空闲span添加到对应的page heap中，可以使分配算法尽可能地充分利用空闲的内存。同时，通过合并相邻的空闲span，还可以避免内存碎片化的问题，提高内存的使用效率。



### scavengeAll

在Go语言中，每个goroutine都有其自己的栈空间。栈空间的大小是固定的，当栈空间用完了，就会触发栈的扩展，即将栈的大小加大一些。如果需要不停地扩展栈大小，会增加运行程序的内存开销。因此，Go语言的运行时系统采用了一种内存回收技术，压缩内存空间以回收已死对象占用的内存，进而减少内存占用。

mheap.go文件中的scavengeAll函数就是其中的一个回收内存的函数。该函数的作用是将内存池中的空闲内存块归还给操作系统，以便其他应用程序使用。scavengeAll函数的具体实现过程如下：

1. 启动一次全局垃圾回收，以便将内存中的已死对象进行回收。

2. 遍历内存池中所有的空闲内存块，找出那些可以被回收的内存块。

3. 将可以被回收的内存块从内存池中移除，并释放其占用的内存。

4. 将已回收的内存块归还给操作系统。

通过这种方式，scavengeAll函数可以在程序运行过程中自动回收不再使用的内存，并将这些内存返还给操作系统，从而提高了程序的内存利用率，并减少了内存占用。



### runtime_debug_freeOSMemory

函数名称：runtime_debug_freeOSMemory

作用：该函数的作用是尝试清除并释放Golang程序占用的一些操作系统内存，包括自由堆内存，堆分配器的全局元数据和垃圾收集器的非引用对象。它主要用于对调试程序的内存占用情况进行监测，以及启示垃圾回收机制执行的条件。该函数可以在程序运行时（runtime）运行时执行。

函数定义：

func runtime_debug_freeOSMemory()

函数说明：

当Golang程序运行时，它会占用一些操作系统内存，在某些情况下，Golang程序可能会持续占用这些内存，这会导致系统性能降低。可以使用runtime_debug_freeOSMemory函数来清除并释放这些操作系统内存，以改善系统性能。该函数主要释放以下内存：

1. 自由堆内存

自由堆内存是堆分配器可以分配但未被分配的内存空间。当运行Golang程序时，会产生自由堆内存，随着Golang程序的运行，此类内存会增多。使用runtime_debug_freeOSMemory函数可以清除并释放自由堆内存，从而加速程序运行。

2. 堆分配器的全局元数据

堆分配器的全局元数据是指堆分配器中的一些数据结构，用于跟踪堆中的对象。当堆中的对象被释放时，堆分配器的全局元数据可能会稍微增加一些内存。使用runtime_debug_freeOSMemory函数可以清除并释放堆分配器的全局元数据，从而进一步减少内存使用。

3. 垃圾收集器的非引用对象

在Golang中，垃圾收集器（GC）会自动回收堆中不再使用的对象。当垃圾收集器回收堆中的对象时，它会产生一些非引用对象，这些对象在堆中不再使用但却还占用内存。使用runtime_debug_freeOSMemory函数可以清除并释放垃圾收集器的非引用对象，从而减少内存使用。

总结：

runtime_debug_freeOSMemory函数是一种用于释放Golang程序占用的操作系统内存的方法。它可以清除自由堆内存，堆分配器的全局元数据和垃圾收集器的非引用对象，并显著改善系统性能。该函数主要用于调试程序的内存占用情况和启示垃圾回收机制执行的条件。



### init

在 Go 语言中，init 函数是一个特殊的函数，它会在程序启动时自动执行。每个包都可以定义一个或多个 init 函数，它们会被自动执行，而且按照它们在源文件中的顺序依次执行。

mheap.go 文件中的 init 函数主要是用来初始化堆空间的。堆空间是一种动态分配内存的方式，当程序需要分配内存时，堆空间会自动为程序分配一块空闲的内存，并返回内存地址。Go 语言的垃圾回收器就是基于堆空间实现的。

在 init 函数中，mheap 包会初始化一些全局变量和数据结构，并调用 runtime 包中的一些函数来初始化堆空间。具体来说，它会调用 runtime.mheap_init 函数来初始化 mheap 结构体，这个结构体主要用来管理堆空间，包括分配和回收内存等操作。

除了初始化堆空间，init 函数还会调用一些其他函数来初始化数据结构。例如，它会调用 initSizes 函数来初始化所有大小类的数据结构，其中包括每个大小类的大小、对齐方式、空闲链表等信息。还会调用 initSpan 函数来初始化 Span 数据结构，Span 是一个大小固定的连续内存块，它用来管理堆空间中的一段连续内存。

总之，mheap.go 文件中的 init 函数主要是用来初始化堆空间和相关数据结构，在 Go 语言中，它是在程序启动时自动执行的，而且调用了一些其他函数来初始化数据结构。



### inList

inList函数用于判断一个堆对象是否在空闲堆列表中。

在Go语言中，堆是用来管理动态分配内存的数据结构，当程序需要分配一块内存时，堆会从空闲堆列表中取出一块大小合适的内存进行分配。当这块内存不再使用时，堆将其加入空闲堆列表中以供下次使用。

在mheap.go文件中，inList函数的作用是遍历空闲堆列表，查找是否有与指定对象大小相同的堆对象。如果查找到了，则返回该堆对象的地址。如果未查找到，则返回0。

具体实现是：先从mheap中获取空闲堆列表的mutex锁，然后遍历空闲堆列表，将每个堆对象的地址与指定对象地址进行比较，如果大小相同则返回该堆对象地址。最后释放mutex锁。

这个函数的作用是为了提高分配内存的效率，当程序需要分配一块内存时，先查看是否有大小合适的空余内存可用，如果有，则可以避免从操作系统申请新内存，提高内存分配效率。



### init

在Go语言的运行时(runtime)中，mheap.go这个文件中的init()函数的主要作用是初始化堆(heap)的数据结构。堆是存储程序运行期间分配的内存的地方。它由一个或多个内存段组成，其中每个内存段的大小都相等。init()函数还会为堆分配一部分常规内存，称为“固定大小的内存池”。该内存池用于存储运行时所需的内部数据结构。此外，init()函数还会初始化堆内存管理器(mheap)的几个常用参数( 如mspansize，随机种子等)。最后，堆内存管理器会设置垃圾收集器所需的参数，并分配初始化堆模板，以及分配线程抢占工具所需的内存。总的来说，init()函数在Go语言的运行时中起到了很重要的作用，因为它确保了程序的堆分配和垃圾收集器的正常运行。



### remove

remove函数是在Go语言的runtime系统中mheap.go文件中的一个方法，其作用是从堆中删除指定的对象。堆是用于动态分配和管理内存的一种数据结构，其中存储的是程序需要的对象，以及中间数据。这个操作对于提高程序的内存使用效率有重要意义。remove方法是用于将堆内部的某个对象删除的方法。

remove方法接收两个参数，第一个参数是指向heap结构体的指针；第二个参数是指针类型*mspan。这个方法首先检查mspan指向的对象是否在heap中，然后将其从heap中删除。如果这个对象在heap中，那么remove函数将执行以下操作：

1. 搜索堆中的对象，找到mspan指向的地址

2. 将mspan从对象的链表中删除

3. 如果对象的链表为空，将对象从free[]数组中删除

4. 将对象从堆的页表和spanmap中删除

5. 返回被删除对象的大小

remove方法的作用非常重要，在Go语言的内存管理中扮演着关键的角色。它能够清理不需要的内存，保证Go代码的执行效率和快速性。



### isEmpty

isEmpty函数是用来判断mheap是否为空的，也就是判断mheap中是否还存在未分配的空间。isEmpty函数会检查以下几个条件：

1. mheap有没有被初始化
2. mheap是否已经空了，也就是laget的长度是否为0
3. mheap中是否还存在未分配的页，也就是mheap中的未分配空间是否大于 ArenaBitmapBytes

如果以上三个条件都满足，则返回true，否则返回false。

isEmpty函数的作用在于判断mheap是否为空，如果是空的，则可以执行一些清理工作，如释放整个mheap的内存。同时，在执行分配空间操作的时候，也可以通过检查mheap是否为空，来判断是否需要重新申请内存来扩大堆空间。



### insert

在Go语言中，当程序需要分配内存时，会调用runtime.mallocgc函数。该函数首先会检查从上一次GC之后已经申请的内存大小是否超过了指定的阈值，如果已经超过了，则会触发一次GC，回收无用的内存。

如果当前还有空闲的内存块没有被使用，那么mallocgc函数将直接返回其中的一个内存块。如果没有空闲的内存块可以被使用，那么就需要从堆中分配一个新的内存块。

堆是Go语言中用于管理内存的一个数据结构，它由mheap结构体来表示。当程序需要从堆中分配内存时，会调用mheap结构体中的insert函数。该函数的作用是往堆中插入一个新的span（一个连续的内存块），并将其加入到相应的链表中。

具体来说，insert函数会首先从mheap结构体中查找一个可用的span链表，如果找不到，则会分配一个新的span链表。然后，它会将新的span插入到相应的链表中，并更新堆的统计信息，包括已分配的内存大小、空闲的内存大小、已经申请的内存大小等。

总之，insert函数是Go语言中堆管理的核心函数之一，其主要作用是将新的span插入到相应的链表中，并更新堆的统计信息。



### insertBack

InsertBack是运行时系统中mheap中的一个函数，它的作用是在堆中插入一个新分配的span并将其添加到span列表的末尾。

在Go语言中，堆被用于在运行时动态地分配内存。在堆中，内存被分配为一个或多个span。每个span都包含一个或多个对象。当所有对象都被释放并且span变为空闲时，它将被添加到一个空闲span列表中以供后续分配使用。

InsertBack函数的主要工作是将一个新分配的span添加到末尾的span列表中。它首先判断span列表是否为空，如果为空，则创建一个新的span列表，并将新分配的span添加到其中。如果span列表不为空，则找到span列表的最后一个span，并将新分配的span添加到其后面。最后，InsertBack更新span的前驱和后继指针，以便在未来的span操作中更容易地访问此span。

总之，InsertBack函数是一个用于在堆中添加新分配的span到span列表的末尾的辅助函数，它使得运行时系统可以更有效地管理堆内存。



### takeAll

在Go语言中，mheap.go文件是runtime包中的一个重要的文件，它主要包含了堆内存管理器（mheap）的实现。

takeAll函数是堆内存管理器中的一个函数，它的作用是将堆内存管理器中的所有可用内存块取出并返回。在堆内存管理器中，当一个对象被释放时，它所占用的内存块会被添加到空闲内存块的链表中。takeAll函数的作用就是将这些空闲内存块的链表取出，并返回给调用者。这个函数通常在GC周期结束时被调用，用于清理所有可用的内存块，并将它们返回到系统默认的堆池中，以便被重复利用。

具体来说，takeAll函数的实现过程如下：

1. 首先从堆内存管理器（mheap）中取出可用内存块的链表，检查这个链表是否为空。

2. 如果链表不为空，则遍历链表，将所有的内存块取出，同时更新内存管理器中的状态，并设置内存块的状态为已分配。

3. 将取出的内存块添加到一个内存块数组中，并将数组返回给调用者。

4. 如果链表为空，则直接返回一个空数组。

总之，takeAll函数是一个用于释放堆内存管理器中所有可用内存块的函数，它在GC周期结束时起到了非常重要的作用。通过这个函数，系统可以有效地重复利用内存，以提高运行效率和资源利用率。



### spanHasSpecials

func spanHasSpecials(s *mspan) bool

这个函数的作用是判断给定的mspan中是否存在特殊类型的对象，如大对象或实时对象。这个函数在垃圾回收过程中用于判断哪些对象应该被处理和移动。

mspan是一个内存区域的描述符，它包含了描述该区域的元数据，如大小和状态。每个mspan描述的区域被划分成多个固定大小的块（对象大小）以容纳分配的对象。当对象被释放时，它们返回到可用空间中供后续分配使用。

特殊类型的对象是指超过一定大小或需要实时分配的对象。由于这些对象在内存使用方面存在一些特殊的要求，因此需要特殊的处理。当垃圾回收器找到一个包含特殊类型对象的mspan时，它需要采取相应的措施，以确保这些对象得到适当的处理。

spanHasSpecials函数遍历mspan中的每个块（对象）并检查它们的大小是否超过了特定的阈值，以判断该mspan是否包含特殊类型的对象。如果存在特殊类型的对象，则返回true；否则返回false。

在垃圾回收器的标记阶段，它需要遍历所有的mspan并判断它们是否包含特殊类型的对象。如果包含，垃圾回收器需要采取不同的标记方法，以确保这些对象能够被正确处理。因此，spanHasSpecials函数在垃圾回收器的运行中发挥了重要作用。



### spanHasNoSpecials

spanHasNoSpecials函数位于mheap.go文件中，其作用是判断给定的span是否具有特殊的内存分配需求。特殊分配需求包括大对象分配、对齐分配或禁用分配，如果一个span具有这些需要，那么它将不适合普通的内存分配。

这个函数对垃圾回收器非常重要，因为它需要知道哪些span可以被用于堆的普通内存分配。如果一个span不能被用于普通内存分配，那么它将被归类为特殊span，并被垃圾回收器进行特殊处理。

在具体实现上，函数检查了span的allocation bitmap，如果bitmap中只有0或1，那么该span被认为没有特殊的内存分配需求，可以用于普通内存分配。如果bitmap中存在其他数字，那么该span就不适合普通内存分配，并被归类为特殊span。

总之，spanHasNoSpecials函数的作用是检查span是否具有特殊的内存分配需求，以便确定它是否适合用于普通内存分配，这对垃圾回收器非常重要。



### addspecial

addspecial函数是用来为heap分配器特殊分配器添加一些空间的。特殊分配器是在heap中分配小对象的分配器，它们的大小小于等于32K，它们被精心设计，以便可以最小化内存碎片，并且具有快速分配和释放速度。

在特殊分配器中分配内存时，如果没有足够的可用空间，就需要调用addspecial函数。addspecial函数的作用是将特殊分配器的空间增加到足够大小以容纳新的内存块。这个函数的行为取决于所使用的特殊分配器类型，但通常它会向heap头部添加新的元数据列表或新的分配缓存。

在addspecial函数中进行了多个检查和计算，以确保新的空间分配正确，例如：

1. 检查请求的大小是否小于特殊分配器的最大尺寸，否则返回错误

2. 计算新的空间大小，确保它至少能容纳单个内存块

3. 计算新的元数据页的数量

4. 为新的元数据页和分配缓存分配内存

5. 初始化新的元数据页，并将它们添加到特殊分配器的元数据列表中

总之，addspecial函数的作用是确保特殊分配器有足够的内存来管理小对象的分配和释放。



### removespecial

removespecial是Golang运行时包中mheap.go文件中的一个函数，用于将特殊对象从堆中移除。该函数主要用于在堆结构中删除一些特定的对象，如mspan、mcentral等，以确保整个堆结构的正确性和一致性。

在Golang的堆结构中，每个堆都由多个mspan（memory span）组成，每个mspan表示一段内存。每个mspan都拥有对应的mcentral（memory central）和mcache（memory cache），这些对象都是用于维护和管理堆的数据结构。

removespecial函数将特殊对象（如mspan、mcentral等）从堆中移除，主要是因为这些对象可能被其他对象所引用，如果不进行特殊处理，就会导致堆结构出现较大的问题，如内存泄漏等，最终导致程序崩溃。因此，当这些对象不再需要时，就需要使用removespecial函数将其从堆中移除，从而避免潜在的问题。

removespecial函数的实现比较复杂，需要考虑很多细节，如并发情况、内存管理等。在函数执行过程中，会对相关的数据结构进行修改，以使整个堆结构保持一致性，同时还需要处理一些边界情况，确保操作的正确性。

总之，removespecial函数是Golang运行时包中mheap.go文件中的一个重要函数，用于将特殊对象从堆中移除，以确保整个堆结构的正确性和一致性。



### addfinalizer

addfinalizer函数的作用是向堆中的对象添加一个finalizer（终结器）函数，该函数在该对象被回收前被调用。

在Go语言中，垃圾回收器是通过标记清除（Mark and Sweep）算法实现的。当一个对象在堆中不再被引用时，垃圾回收器会将其标记为可回收的，并等待下一次回收时将其清除。但是，在某些情况下，我们希望在对象被回收之前执行一些清理工作，比如释放该对象占用的资源，关闭该对象持有的文件等。这就是finalizer函数的作用。

addfinalizer函数的实现涉及到了mheap结构体的相关函数和各种锁的使用。具体步骤如下：

1. 通过mheap_.lock()获取全局堆锁，避免多个goroutine同时修改堆结构发生竞争。

2. 将对象和finalizer函数绑定成一个finalizer对象，然后将该对象添加到堆中。如果finalizer对象的大小大于等于_64KiB，那么该对象必须使用mheap.alloc_huge函数分配，否则使用mheap.alloc函数分配。

3. 如果finalizer对象已经被标记为大对象（即使用mheap.alloc_huge函数分配），那么需要在mcentral中释放该对象。mcentral是由多个mcache组成的全局对象池，用于存储中等大小的对象并避免昂贵的系统调用（例如brk）。

4. 释放全局堆锁。

需要注意的是，在使用finalizer函数时需要小心，因为如果finalizer对象本身也被回收了，那么finalizer函数将不会被执行。此外，finalizer函数的执行时间是不确定的，因为垃圾回收器可能会在任何时间回收对象。因此，建议使用finalizer函数时要尽可能的简单并且不能依赖于时间敏感的行为。



### removefinalizer

在Go语言中，通过runtime.SetFinalizer函数可以为一个对象设置finalizer，即该对象的垃圾回收时会自动执行对应的finalizer函数。而removefinalizer函数则是用来取消对一个对象的finalizer设置。

具体来说，removefinalizer函数中的实现逻辑如下：

1. 判断对象是否为nil，如果是则直接返回。
2. 获得对象的finalizer信息（finalizerCleanup结构体），如果没有设置finalizer，则直接返回。
3. 清空对象的finalizer属性，即将对象的finalizerCleanup结构体置为空。
4. 将对象从finalizer queue（维护所有需要执行finalizer的对象队列）中移除。
5. 将对象从finalizer map（维护所有设置了finalizer的对象映射表）中移除。

总体来说，removefinalizer函数的作用就是取消对一个对象设置finalizer的操作，并从维护finalizer信息的数据结构中移除该对象，以避免影响垃圾回收的正常进行。



### setprofilebucket

setprofilebucket函数的作用是为堆分配器的分配器档案设置桶信息。

在Go语言中，分配器档案用于跟踪和统计内存分配和释放。Mheap中的setprofilebucket函数接受桶和minsize参数，将堆分配器中的桶信息更新为该参数指定的值。对于每个桶，它保存可分配的对象大小范围，并且存储在mheap的profile字段中。

具体来说，setprofilebucket函数负责设置每个桶的范围和名称，以便对内存分配进行计数和跟踪。这些桶对于发现分配方面的问题和优化内存分配非常重要。

堆分配器会利用桶将小的分配请求合并以节省内存，从而避免小碎片和对性能的影响。

总之，setprofilebucket函数是堆分配器的一部分，负责设置内存分配桶信息，是Go语言实现内存管理的重要组成部分。



### newSpecialsIter

newSpecialsIter这个函数是用于创建一个用于特殊堆块迭代的特殊迭代器的函数。在Go语言运行时，特殊堆块是指一些内部使用的、较小的堆块，它们通常比普通堆块更小，并且具有不同的内存管理特性。

newSpecialsIter函数的作用是创建一个特殊堆块迭代器，用于在特殊堆块中遍历内存块。特殊堆块迭代器可以从一个特殊堆块的起始位置开始，按照一定的顺序遍历特殊堆块中的每个内存块。这个函数的返回值是一个指向特殊堆块迭代器的指针。

在Go语言运行时中，特殊堆块迭代器的主要作用是在垃圾回收器的工作过程中对内存进行分类和管理。垃圾回收器会将内存分为几个不同的类别，并使用特殊堆块迭代器来遍历这些内存块，进行相应的内存清理和管理操作。

总之，newSpecialsIter函数的作用是创建一个特殊堆块迭代器，用于在特殊堆块中遍历内存块，并支持Go语言运行时的内存管理和垃圾回收。



### valid

valid这个func的作用是检查malloc堆中的数据结构是否正确，并通过打印日志、调用panic等方式来报告错误。

在Go语言中，malloc堆是存放动态分配的内存的区域，它由MHeap结构体维护，而valid函数就是用来检查这个结构体是否正确的。具体来说，valid函数会检查以下几个方面是否正确：

1. 验证MHeap的arena_start和arena_used成员变量是否指向正确的地址。arena_start是malloc堆的起始地址，arena_used是已经使用的堆空间大小。

2. 验证中心缓存页的数量是否正确。中心缓存是每个P使用来缓存小对象的地方，其中P是指处理器的抽象概念。

3. 验证MHeap的free和busy treap的高度和后继关系是否正确。Treap是一种能够同时提供二叉搜索树和堆操作的数据结构，被用来维护malloc堆的空闲块和已用块的树状结构。

4. 验证所有free和busy treap节点所指向的内存块是否满足在堆空间区域中、内存对齐、大小正确，并且相邻节点之间没有重叠。

如果valid函数检测到的错误无法修复（例如指向错误的内存地址、空闲块内存已经被破坏、小对象数量错误等），它将报告错误信息并调用panic函数，以防止程序进一步运行下去。



### next

next函数是mheap结构体的一个方法，作用是从当前的mheap对象中获取下一个可用的span，如果当前buffer中没有可用的span，则该函数会切换到下一个buffer，直到找到一个可用的span为止。span是用于分配内存的基本块，每个span对应一个固定大小的内存块。

具体来说，该函数会按照一定的优先级规则遍历当前mheap对象中的buffer，从每个buffer中取出可用的span，直到找到一个符合要求的span为止。如果当前buffer中没有可用的span，则会尝试切换到下一个buffer，如果下一个buffer中仍然没有可用的span，则会进入到一个性能较低的模式，直接遍历所有的buffer，直到找到可用的span为止。

在遍历每个buffer中的span时，该函数会按照以下优先级规则进行尝试：

1. 从central存储池中获取span，central存储池中存储的span适用于任何大小的内存分配。
2. 从mcentral中获取span，mcentral存储的span适用于特定大小范围内的内存分配。
3. 从mheap中获取span，mheap存储的span适用与大内存分配。

next函数的作用在于保证mheap对象中始终有足够的span可用于内存分配。通过使用优先级规则和缓存机制，可以提高内存分配的效率和性能。



### unlinkAndNext

unlinkAndNext函数是mheap.go中的一个函数，用于从size类表中删除块，并返回指向在相同大小类别中的下一个块的指针。

该函数的作用如下：

1. 接受一个指向特定大小类别中的空闲块的指针。

2. 确定指针所指向的块的前一个和后一个块是否为空闲的。

3. 如果前一个块为空闲的，则合并这两个块。

4. 如果后一个块为空闲的，则链接前一个块和后一个块，并从大小类别中删除后一个块。

5. 如果前一个块和后一个块均不为空闲，则从大小类别中删除这个块。

6. 最后，返回后一个块的指针，以便比较它们的大小。

这个函数的作用是使mheap的空闲块列表保持正确的状态，以便后续的内存分配可以正确地使用空闲块。



### freeSpecial

mheap.go文件中的freeSpecial函数是用于释放特殊大小（special size）的堆内存块的。在Go的运行时（runtime）中，对于不同大小的对象，使用不同的分配策略。通常，小的对象会采用大小固定的堆内存块，而大的对象会采用按需分配的堆内存块。但是，有一些特殊大小的对象，它们的大小介于这两种情况之间，无法使用固定大小的堆内存块，也不适合按需分配堆内存块。为了解决这个问题，Go运行时使用了特殊大小的堆内存块，这些内存块的大小是固定的，但它们与常规的大小固定的堆内存块不同，因为它们只能分配特定大小的对象，并且它们的分配方式也与常规的大小固定的堆内存块不同。

freeSpecial函数的作用是释放特殊大小的堆内存块。当程序不再需要特殊大小的对象时，它们所占用的堆内存块就需要被释放。freeSpecial函数会将这些堆内存块标记为可用，并将它们添加到空闲列表中，以备下次分配使用。同时，如果空闲列表中的堆内存块数量过多，freeSpecial函数也会根据需要释放一些堆内存块，以尽可能减少内存的浪费。

需要注意的是，freeSpecial函数只能用于特殊大小的堆内存块，对于其他大小的堆内存块，应使用其他相应的函数进行释放。此外，在编写Go程序时，通常不需要直接使用freeSpecial函数，它是由Go的内存管理器自动调用的。



### bytep

bytep是mheap.go文件中的一个函数，它用于计算在堆中分配的字节数。具体来说，它通过对maxspansize，heap.sysmem和heap.maximumBlockCacheBytes进行简单的数学运算来计算分配的字节数。这些参数定义了堆的大小以及分块的大小，从而确定了可以分配的最大字节数。

该函数的返回值可以用于检测和诊断垃圾收集或内存问题。在调试和优化堆使用时，bytep可以帮助开发人员了解哪些部分消耗了堆空间以及如何管理内存。

总的来说，bytep函数的作用是计算在堆中分配的字节数，以便监测和优化内存的使用。



### bitp

bitp函数位于Go运行时的mheap.go文件中，是一个用于内存管理的二进制堆的实现的一部分。bitp函数的作用是找到与给定堆大小最接近的2的幂次方值，并返回该值。

二进制堆是一种用于分配和释放内存的数据结构，它将可用的内存块放入由2的幂次方大小组成的桶中。使用桶大小相乘的方式，可以快速地找到合适的桶大小，从而避免了不必要的分配和释放内存。

bitp函数通过将给定的堆大小向上取整到最近的2的幂次方，来找到最接近的2的幂次方值。由于内存块大小通常不会恰好等于2的幂次方，因此使用bitp函数可以更好地利用堆中的存储空间。

具体来说，bitp函数使用了一个算法，它首先将堆大小减去1，并将结果与该值的二进制补码进行或操作，以获得下一个最近的2的幂次方。然后，它将结果向右移动1位，以获得最接近的2的幂次方。

总之，bitp函数在Go运行时的mheap.go文件中是一个用于内存管理的重要功能，它帮助快速找到最接近给定堆大小的2的幂次方值。这有助于更好地利用可用的内存块，并避免不必要的内存分配和释放。



### tryAlloc

tryAlloc是runtime中mheap.go文件中的一个函数，它的作用是在堆上尝试分配内存。当程序需要分配内存时，系统会首先尝试在本地缓存中分配内存，如果本地缓存没有足够的空间，系统会将内存分配请求交给mheap.tryAlloc()来处理。

mheap.tryAlloc()函数会根据请求的内存大小，查找可用的空间，并根据描述符信息建立一个新的分配。它还会尝试在每个堆区域中查找空闲的大小来匹配要求的内存大小。如果找到了空闲的堆区域，会尝试将其合并成一个新的堆，以进行更好的内存使用。

如果mheap.tryAlloc()成功地分配了内存，它会返回一个描述符指针，并设置相应的元数据或记录以维护内存的使用情况。如果没有足够的可用空间，该函数将返回nil。

总之，tryAlloc函数是用于堆上的内存分配，它负责查找和分配可用的内存块，以满足程序的需求。它还负责维护内存使用情况，并将有关的元数据记录到堆数据结构中。



### newMarkBits

newMarkBits函数是用来创建一个新的标记位图的，它主要用于垃圾回收器中的标记阶段。在标记阶段，每个对象都需要被标记为已使用或未使用。为了实现这个标记，垃圾回收器需要使用标记位图来记录每个对象的状态。

在Go语言的垃圾回收器中，使用了分代垃圾回收算法。分代垃圾回收算法中，内存被分为若干代。每次触发垃圾回收时，只会对某些代进行回收。新生代中的对象通常会被更频繁地回收，而老年代中的对象则会较少被回收。

newMarkBits函数的主要作用就是创建一组标记位图，每组标记位图对应一个代。在新生代中，通常会创建两个标记位图：用于当前垃圾回收的位图和下一次垃圾回收的位图。在老年代中，通常会创建三个标记位图：用于当前垃圾回收的位图、下一次垃圾回收的位图和用于增量标记的位图。

在创建标记位图时，使用的是Go语言的slice结构，它能够动态地调整大小。标记位图本质上是一组uint8类型的数字，每个数字都对应于某个内存块的状态。标记位图的大小通常是根据当前代中的对象数量和内存块大小来确定的。  

总的来说，newMarkBits函数就是为垃圾回收器创建新的标记位图，以支持垃圾回收器中的标记阶段。



### newAllocBits

newAllocBits()函数是用来分配一段内存并初始化为“未分配”状态的位图，该位图用于表示堆上对应区域的分配状态。在Go语言的垃圾回收机制中，每个堆区域的分配状态都被表示为一个位图。当程序需要进行垃圾回收时，垃圾回收器会根据这些位图来判断哪些对象需要回收，哪些对象可以保留。

newAllocBits()函数的输入参数是heapBits类型的指针，该类型表示一个位图。该函数会为该位图分配足够的内存，使其可以表示堆上一个完整的区域（默认为8192字节）。函数会将位图所有的位都初始化为0，表示该区域全部是空闲的。

该函数的实现比较简单，具体代码如下：

```
func newAllocBits(hb *heapBits) {
    hb.n = _PageSize / heapBitDiv
    hb.bits = (*uint*)((*notInHeap)(unsafe.Pointer(sysAlloc(_PageSize))) + _PageSize - sys.PtrSize)
    for i := uintptr(0); i < hb.n; i++ {
        *hb.bits = 0
        hb.bits = (*uint)(unsafe.Pointer(uintptr(unsafe.Pointer(hb.bits)) + heapBitDiv))
    }
}
```

函数首先计算了需要分配的位图的大小n（实际上就是一页内存的大小除以heapBitDiv，heapBitDiv的默认值为8），然后调用sysAlloc函数分配一段连续的内存，该内存的大小为1页（默认大小为8KB）。

接下来，函数使用了一个指针技巧，将申请到的内存地址转换为一个notInHeap类型的指针，然后将该指针向前偏移一个页的大小（即_PageSize），得到该内存区域的起始地址。notInHeap类型是一个空类型，该类型的大小为0，所以不会占用内存。这种指针技巧可以避免将heapBits结构体存储在堆上，并且可以提高分配和释放内存的效率。

接下来的for循环遍历分配到的内存区域，并将每个uint类型的位图初始化为0。这里使用了指针运算，将bits指针向后偏移heapBitDiv个字节大小，以便指向下一个位图所在的地址。

总之，newAllocBits()函数是Go语言垃圾回收机制中的一个重要组成部分，它用于分配并初始化用于表示堆上分配状态的位图。虽然它的实现比较简单，但是它的作用却非常重要，能够确保垃圾回收器能够正常地工作，保障程序的稳定性和效率。



### nextMarkBitArenaEpoch

nextMarkBitArenaEpoch函数的作用是根据当前GC的进度和可用的堆空间大小，确定下一次标记位图扫描的阶段。

在Go语言的垃圾回收机制中，标记-清除算法被用来回收不再被程序使用的内存。在标记阶段，垃圾回收器会遍历堆内存中的所有对象，并标记出存活的对象，使得后续的清除操作可以跳过那些已经不再使用的对象。

nextMarkBitArenaEpoch函数定义在runtime/mheap.go文件中，它是垃圾回收器的一部分，专门负责管理堆内存的标记位图。标记位图是一个位序列，用于记录哪些内存块中的对象是存活的，垃圾回收器需要持续更新并扫描这个位图。

nextMarkBitArenaEpoch函数的实现流程如下：

1. 如果堆内存中已经没有空余空间可用，那么下一次标记位图扫描的阶段被设置为NullArena，表示标记阶段结束，垃圾回收器将对可回收的对象进行清理。

2. 如果当前的标记位图区域还未被完全扫描，则返回当前标记位图区域的下一个epoch，这个epoch用于收集新分配的内存块的标记位图信息。

3. 如果当前的标记位图区域已经扫描完毕，则先将当前的标记位图区域清空，再扫描下一个标记区域。如果下一个标记区域不存在，则将下一次标记位图扫描的阶段设置为NullArena。

通过nextMarkBitArenaEpoch函数的实现，垃圾回收器能够在程序运行时始终保持内存的健康状态，从而提高程序的性能和稳定性。



### newArenaMayUnlock

newArenaMayUnlock函数是Go语言运行时（runtime）内存管理模块中的一个函数，其主要作用是根据指定大小分配一个Arena（内存池）。

具体而言，newArenaMayUnlock函数的实现流程如下：

1. 先尝试从mheap的空余内存池（spans）中获取Arena
2. 如果从当前mheap的spans中获取到了Arena，则返回该Arena
3. 如果没有获取到Arena，则需要从全局mheap列表中查找一个合适的Arena
4. 从全局mheap列表中找到一个合适的Arena后，尝试从该Arena中分配一定数量的span，用于放置小对象
5. 如果成功获取span，则初始化该Arena，并返回可用的内存池


newArenaMayUnlock函数主要涉及到两个重要的数据结构：mheap和Arena。

mheap是Go语言运行时内存管理模块中的一个结构体，代表内存堆。其中，spans字段是一个空闲的span列表，其中包含了一定数量的Arena，可以用来分配内存。

Arena是内存池的最小单位，用于管理一定范围内的内存。Arena包含了多个span，用于存储小对象的内存块。

在实际应用中，newArenaMayUnlock函数被广泛应用于从内存池中获取空余的Arena，以便分配对象。该函数的实现非常高效，能够快速分配内存，提升程序的性能。



