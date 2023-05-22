# File: mcache.go

mcache.go文件的作用是管理和分配MCache的内存管理器。MCache是一种专门为每个线程维护一部分内存分配空间的数据结构，它存储了一些小的、短周期的分配对象。在Go语言中，每个线程都有一个MCache用于缓存分配的内存块，这样可以减少对于堆的访问，提高分配性能。

MCache是由一些无锁的内存分配器和一些元数据组成的。mcache.go文件通过一些关键函数来管理这些元数据和内存分配器。例如，mcacheinit函数用于初始化MCache，mcachealloc函数用于从MCache中分配一定大小的内存，mcacheallocspan函数用于从堆上分配足够大的内存块，并将其分成多个小块存储在MCache中，mcachedestroy函数用于销毁MCache以释放其内存空间，等等。这些函数的实现在底层使用了一些高效的算法，例如CAS和Spinlock等，以确保性能和并发性。

综上所述，mcache.go文件是Go语言中一个非常重要的内存管理器，它的实现涉及到了底层的并发算法和内存管理的细节。理解这个文件可以帮助我们了解Go语言中的内存管理机制，并能够更好地优化和调试代码中的内存使用。




---

### Var:

### emptymspan

在go/src/runtime/mcache.go中，emptymspan是一个指向空的mspan结构体的指针，它的作用是作为初始化时的一个默认值。mspan结构体是 Go 语言中用于管理堆内存的数据结构，代表了一段连续的内存空间。

在mcache.go中，emptymspan变量被用于初始化mcache结构体中的mspancache数组，这个数组是用于管理不同大小的内存块的缓存。在初始化时，所有的mspan都被认为是空的，因此mspancache数组中的每个元素被初始化为指向emptymspan变量的指针。在程序运行过程中，当需要分配某个特定大小的内存块时，会先检查对应大小的mspancache中是否有可用的mspan，如果有，则直接从缓存中分配；如果没有，则需要从堆中分配新的内存空间。

通过使用emptymspan变量作为默认值，可以避免在初始化mcache时每次都需要创建新的空的mspan，从而提高程序的性能和效率。






---

### Structs:

### mcache

mcache结构体是golang中的内存缓存，主要用于缓存小的对象（小于32KB），避免过度使用系统的内存分配器。每个线程(M)都有自己的mcache，因为每个线程需要有自己的内存分配器。

mcache结构体的成员变量包括：

1. allocator: 内存分配器，用于分配小于32KB的内存块。

2. stackcache: 栈式缓存，用于快速分配小的内存块，避免频繁地调用内存分配器。

3. stat: 内存使用情况的统计信息，包括缓存中的对象数目和大小。

通过使用mcache结构体，可以有效地提高golang程序的性能和效率，避免过度的内存分配和回收。在实际开发中，建议对内存分配的使用进行优化，合理地使用mcache的功能，以避免内存分配器的开销和碎片化问题。



### gclink

在go/src/runtime/mcache.go文件中，gclink是一个结构体，代表了一个缓存块的内存区域。该结构体有以下几个作用：

1. 存储缓存块的内存区域信息：该结构体中包含了指向上下文相关的缓存块的指针，以及该缓存块的大小和已使用的空间大小等信息。

2. 用于在不同的mcache之间传递缓存空间：在多个线程之间共享工作时，需要一个数据结构来传递缓存空间。gclink结构体就是用来传递缓存空间的一个数据结构。

3. 参与垃圾回收：当需要回收一块内存时，垃圾回收器需要知道哪些内存块不再使用，从而将它们释放回内存池中。gclink结构体用于管理缓存块的内存信息，如果该内存块不再使用，则可以将其标记为垃圾并进行回收。

总之，gclink结构体是在内存分配和回收过程中起关键作用的一个数据结构，它管理着缓存块的内存和垃圾回收等相关信息，同时也支持在不同线程之间共享缓存空间，从而提高了系统的性能和效率。



### gclinkptr

gclinkptr 结构体定义如下：

```go
type gclinkptr struct {
	next *mspan
	prev *mspan
}
```

在 Go 的垃圾回收机制中，`gclinkptr` 结构体用于连接 `mspan` 结构体，来维护内存分配的信息。

`mspan` 结构体表示一段连续的内存区域，`gclinkptr` 结构体表示这段连续内存区域当前的状态：

- `next` 指向下一段空闲的 `mspan` 区域；
- `prev` 指向上一段空闲的 `mspan` 区域。

通过这种方式，可以将多个 `mspan` 区域连接在一起，形成一个链表结构。当需要分配内存时，可以遍历这个链表，找到第一段足够大的空闲 `mspan` 区域。同时，当有一段 `mspan` 区域被释放后，也可以将其插入到合适的位置，以避免内存碎片化。



### stackfreelist

stackfreelist是一个自由栈列表，主要用于储存一定大小的空闲栈，以供后续的调用。

在Go语言的goroutine调度过程中，每个goroutine都会拥有自己的栈，当goroutine中的函数执行完毕后，栈空间就会被释放，并加入到stackfreelist中，以供后续的goroutine使用。这样做可以避免频繁申请和释放内存空间，提高了程序的效率。

stackfreelist中存储的对象是stack，它是一个struct类型，包括以下属性：

- size：表示栈的大小。
- stk：存储栈空间的指针，指向一块连续的地址空间。
- freeindex：表示空闲位置的索引，指向下一个可用的位置。

当程序需要一个新的栈空间时，会从stackfreelist中取出一个栈，将其从列表中删除，并返回给goroutine使用。如果stackfreelist中没有可用的空余栈，程序会根据需要重新分配一块内存空间来创建新的栈。

总之，stackfreelist结构体的作用是优化goroutine的调度过程，提高程序的性能和效率。



## Functions:

### ptr

mcache.go文件中的ptr函数是用于获取大小为size的对象的指针的。该函数首先尝试从本地mcache的列表中获取大小为size的对象，如果列表中没有可用对象，则尝试从mcentral获取对象。如果mcentral中没有可用对象，则尝试从全局堆分配对象，并将其分配给mcentral或本地mcache。

该函数的作用是优化内存分配，以避免频繁地使用全局堆。通过使用本地mcache和mcentral，可以减少锁的竞争和内存碎片，提高程序的性能和稳定性。此外，该函数可以在堆和本地mcache之间移动对象，以优化内存的使用。



### allocmcache

allocmcache是一个函数，用于为M（machine，即线程）分配mcache（memory cache，内存缓存），它存在于go/src/runtime/mcache.go文件中。

在Go语言中，有两个重要的内存缓存，一个是全局缓存，一个是线程本地缓存。全局缓存存储由mallocgc分配的大块内存，以及大型对象，它的大小可以根据需要动态增长和缩小。线程本地缓存则存储线程私有的小对象，每个线程都有自己的mcache，用于缓存小对象分配和归还操作的中间结果。它能够减少内部锁的争用，提高并发性能。

allocmcache的作用是创建线程的mcache，它首先从mcachealloc缓存中尝试获取一个可用的mcache，如果获取到了则直接返回。否则，它将使用malloc进行堆分配，并将结果与freemcache链表（空闲mcache）进行连接。它还会调用reflectcall函数注册一个线程在从全局堆中分配内存时用于分配内存的函数，以及一个线程在将内存返回给全局堆时使用的函数。

该函数的实现细节非常重要，它需要比较高的并发性能和低的内存开销，因为每个线程都需要分配自己的mcache。因此，该函数使用多个内存缓存以减少锁的争用，并使用一些技巧来减少内存开销，并尽可能地在mcache直接分配内存。它还需要与其他函数紧密协作，以确保mcache在正确的时间分配和回收。



### freemcache

在Go语言中，每个线程都会有一个自己的本地缓存，用于分配内存，该缓存被称为mcache。这个缓存包含了一系列由操作系统分配的arena，在这些arena中分配小对象。

freemcache函数的作用是将mcache中的arena清空，并将它们释放回到系统中。这是在线程终止时调用的，以确保所有内存都已释放回操作系统。

在这个函数中，首先会获取线程的mcache，然后逐一将mcache中的arena释放，并将对应的span归还给central pool，最后调用mSys.free函数将arena的内存块释放回到系统中。

因此，该函数的作用是确保线程退出时，所有的内存都被释放回系统，以减少内存泄漏的可能性。



### getMCache

getMCache这个func的作用是从当前的P（Processor）中获取MCache（Memory Cache），即获取当前P的私有缓存。

在Go语言中，每个P代表着一个Processor，它是由操作系统线程（goroutine）和一些额外的堆栈等信息组成的执行单元。每个P负责管理一定数量的goroutine任务，并配备了一个私有的MCache来缓存一些常用的对象或是已经分配过的对象，以便快速的分配/回收内存。

getMCache这个func的具体实现如下：

1. 检查当前P是否已经有了私有的MCache，如果有则直接返回。

2. 如果当前P尚未拥有MCache，则进入lockOSThread临界区，并再次检查当前P是否有私有的MCache，以免在临界区准备分配MCache的时候，其他线程进行了干扰。

3. 如果当前P确实没有MCache，则从centralg中心缓存中获取需要的MCache。注意，在获取时需要加锁。

4. 如果centralg中心缓存中没有可用的MCache，则直接新分配一个MCache。

5. 分配好了MCache之后，需要将MCache绑定到当前P上，以便以后的使用。

6. 最后退出临界区并返回当前P的私有MCache。

通过这样的一套机制，我们可以保证每个P都有自己的MCache，从而可以快速地分配/回收内存，提高应用程序的执行效率。



### refill

mcache.go文件中的refill函数是用于重新填充mcache（memory cache）的。mcache是go语言中的一种内存分配器，用于管理小对象的分配和释放。在go语言中，堆内存的分配和释放需要涉及复杂的算法和数据结构，会消耗大量的时间和资源，而对于小对象，频繁的分配和释放只会浪费更多的资源。所以，go语言在运行时维护了一些小对象的缓存，即mcache，来提高内存分配和释放的效率。

refill函数的作用是将mcache重新填充为满状态。因为mcache是在运行时动态分配的，所以在使用一段时间后，可能会出现空余或者不足的情况。refill函数就是用于处理这种情况的。

refill函数中的主要步骤如下：

1. 检查mcache中的span是否已经用完（即alloc为空），如是则需要重新分配。

2. 调用nextFree函数找到下一个空闲的span。

3. 如果没有空闲的span，则调用central.allocSpan函数从central中获取一段新的span，并将其分配给当前的mcache。

4. 如果找到了空闲的span，则调用spanrefill函数填充alloc。

5. 最后，更新mcache中的其他字段（如本次填充的span数量、已填充的对象数量等）。

总的来说，refill函数的作用是保证mcache中总是有足够的对象可以使用。通过重复使用已经申请的span，可以减少内存分配和释放的开销，从而提高程序的性能。



### allocLarge

allocLarge是mcache.go文件中的一个函数，它的作用是用于大对象分配（即大小超过32KB的内存分配）。在Go语言中，大对象分配的过程与小对象分配的过程不同。当分配大对象时，需要从堆上直接分配，而不是从固定大小的空闲列表中选择一个小的区块来进行分配。

具体来说，当程序需要分配一个大对象时，allocLarge函数将会调用mheap_.alloc函数来从堆上分配一块足够大的内存，并返回分配的内存的起始地址。同时，allocLarge函数还会负责将大对象的大小放入mcache中的localLargeSizeList中，以便在之后的内存分配中进行检查以尽量减少内存浪费。

除了分配内存之外，allocLarge函数还会判断是否需要对内存进行清零处理，即是否需要将新分配的内存中的所有字节都初始化为零。这项操作可以提高程序的安全性，避免出现悬垂指针等问题。

总之，allocLarge函数是Go语言运行时系统中的一个重要功能，它负责管理大对象的内存分配，并确保程序的安全性。



### releaseAll

mcache.go文件中的releaseAll函数的作用是释放所有缓存的span，同时将对应的mcentral的准备列表(partialLists)转移到空闲列表(nonempty)中。MCache(一个线程缓存)中缓存了多个span，但是随着内存的使用，有些span可能不再使用而被释放，为了降低内存的占用，需要将这些未使用的span释放掉。 

releaseAll函数会遍历MCache中的缓存列表，将全部span列表中的状态改为MSpanDead(处于不活动状态的span)并加入到对应的mcentral的准备列表(partialLists)中。然后将准备列表(partialLists)中的span转移到空闲列表(nonempty)中。这样做的目的是为了在下次请求span时，可以直接从非空闲列表(nonempty)中取出对应大小的span，而不需要重新从中央页缓存(mcentral)中获取新的span，从而减少系统负担。 

另外，releaseAll函数也会释放MCache内存中缓存的span，减轻系统内存占用压力，提高系统性能。在MCache中缓存的span数量也是有限制的，如果超过了一定数量，就需要释放不再使用的span，从而避免系统崩溃。



### prepareForSweep

prepareForSweep函数是Go语言运行时系统中mcache.go文件中的一个函数，其作用是为垃圾回收器的扫描做准备。

在Go语言中，内存分为堆和栈。栈是由系统自动分配和回收的，而堆是由垃圾回收器进行管理。垃圾回收器会对堆中的对象进行标记和清除，将无用的对象回收并释放其占用的内存。在回收堆中的对象之前，需要先对这些对象进行扫描，以便将其与其他对象的引用关系找出来，这个扫描工作需要完成后回收器才能执行清除工作。

而prepareForSweep函数则是扫描对象之前的准备工作，它会对当前的内存缓存进行处理，让缓存中的对象可以在垃圾回收时被识别和扫描。prepareForSweep函数的工作流程如下：

1. 将当前mcache中的all和alloc对象切片交换，这样alloc中的对象就变为了all中的对象，all中的对象就变为了alloc中的对象；

2. 将alloc对象切片中的对象分为四类：标记、非指针、指针和特殊对象，标记的对象会被垃圾回收器标记并扫描，非指针的对象不会被回收器扫描，指针对象和特殊对象则需要进行特殊处理；

3. 通过字节对齐的方式对alloc中的对象进行重新排布，使得每个对象的首地址都对齐到8字节或更高的倍数，这个操作可以提高扫描性能；

4. 更新mcache中的top和本地的allocPool对象；

5. 将mcache中的all和alloc对象切片再次交换，保证prepareForSweep函数执行完毕时，mcache的alloc对象切片是最新的。

总之，prepareForSweep函数的作用是为垃圾回收器的扫描做准备，将所有需要回收的对象标记出来，并进行必要的操作以提高扫描效率。



