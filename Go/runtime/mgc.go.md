# File: mgc.go

mgc.go 是 Go 语言 runtime 包中的一部分，主要负责 Go 语言的垃圾回收机制 (Garbage Collector) 的实现。

Go 的垃圾回收机制采用了标记 - 清除 (mark and sweep) 算法，其过程主要包括以下几个步骤：

1. 标记阶段（Marking Phase）：从根对象开始，遍历所有对象标记活动对象，即那些在程序中仍然可达的对象。这个过程使用了三色标记法（Three-color marking algorithm）来提高效率。
2. 制作工作列表（Making Worklist）：将待回收的对象放入工作列表中。这个过程采用了 Write Barrier 策略来避免漏标活动对象，并同时减少在标记阶段遍历整个堆的时间。
3. 清除阶段（Sweeping Phase）：释放那些未被标记的对象的内存空间。此时程序可以重新分配这些内存空间给其他对象。
4. 内存压缩（Memory Compaction）：对于一些碎片化严重的内存区域进行整理，以提高程序的内存使用效率。

mgc.go 文件中记录了垃圾回收器的状态，以及负责整个垃圾回收过程中的进程管理，包括在垃圾收集中分配堆内存，清除不再使用的对象等等。它还包含了一些内存管理相关的函数，例如： gc() 函数、bgsweep() 函数等等。这些函数负责执行垃圾回收算法的各个步骤，确保 Go 程序在垃圾回收时的性能和稳定性。




---

### Var:

### gcphase

gcphase是一个表示垃圾回收器所处阶段的变量。在Go语言的垃圾回收过程中，gcphase用于判断垃圾回收器的具体状态，以确定是否可以执行某些操作或运行特定阶段的垃圾回收。

这个变量有以下几种状态：

1. gcNone：没有进行垃圾回收。
2. gcMark：进行标记阶段。
3. gcMarkTermination：标记阶段结束，准备进入清扫阶段。
4. gcSweep：执行清扫阶段。
5. gcSweepWait：清扫阶段结束，等待goroutine的结束。
6. gcFinalize：执行最终清扫和某些其他操作。

gcphase的作用是用于同步垃圾回收器的不同阶段，确保垃圾回收器在不同阶段之间的正确转换。同时，它也可以用于检查垃圾回收器是否已经完成，以便其他代码可以安全地使用不再需要清理的内存。



### writeBarrier

在 Go 语言中，当进行垃圾回收时，一个重要的操作是标记和清理不再使用的内存。标记操作可以通过遍历程序中的对象进行标记，而清理操作可以将被标记为不再使用的对象回收。在进行标记操作时，需要确保对象在被标记后不会再被修改，以避免标记的误判。

在任何时刻，如果程序中的 Go 程序对指针进行了赋值、传参、成员访问或数组访问操作等，就可能会改变一个对象的地址或者创建新对象，这种情况下就可能出现标记误判。为了避免这种误判，Go 语言引入了 write barrier 机制，也称为写障碍。

write barrier 机制会在程序对指针进行修改时，触发一个拦截器，在记录修改前后指针所指向的对象，并标记其为“灰色”。

在 Go 语言中，write barrier 实现是在 mgc.go 文件中的 writeBarrier 变量中，这个变量是一个函数类型的值，它会被插入到 Go 语言的运行时系统中的指针赋值和传递操作中。每次进行这些操作时，writeBarrier 函数就会被调用，用于实现标记和跟踪指针操作并记录它们的变化。

简而言之，writeBarrier 变量的作用是实现 Go 语言中的写障碍机制，保证垃圾收集器程序在标记时不会误判。



### gcBlackenEnabled

在 Go 语言中，gcBlackenEnabled 变量是垃圾回收器的一个开关变量，它的值会影响到 GC 的执行行为。具体来说，当 gcBlackenEnabled 为 true 时，GC 会默认执行黑色染色算法来标记内存对象的状态；当 gcBlackenEnabled 为 false 时，则会禁用黑色染色算法，改用灰色染色算法来标记内存对象状态。这个算法是用于标记内存对象是否已经被扫描，从而避免对同一内存对象重复扫描的。

gcBlackenEnabled 变量的初始值为 true，如果用户程序想要关闭黑色染色算法，可以在程序的运行时环境中通过环境变量 GOGC=off 来禁用。禁用黑色染色算法会带来一些运行时的开销，但可以避免由于内存对象的繁多而导致的 GC 暂停过长，从而提高系统的性能。因此，在实际开发中，我们需要根据具体情况，来权衡开启或者关闭 gcBlackenEnabled 变量。



### gcMarkWorkerModeStrings

gcMarkWorkerModeStrings是用于描述垃圾回收器标记阶段的工作模式的变量。垃圾回收器的标记阶段是指在程序运行过程中识别出哪些内存块可以被回收的阶段，是垃圾回收器中的重要组成部分。

gcMarkWorkerModeStrings变量的作用是提供一种可读性更高的方式来设置和描述垃圾回收器标记阶段的工作模式。它允许用户通过设置不同的字符串来改变垃圾回收器的默认行为，并以一种更加友好的方式来了解垃圾回收器的工作模式。

具体来说，gcMarkWorkerModeStrings变量包含了垃圾回收器标记阶段的三种不同工作模式，分别是:

1. "local"：表示使用本地缓存来存储标记的对象，减少线程间同步操作的次数，从而提高标记效率。

2. "central"：表示使用中央缓存来存储标记的对象，可以加快标记速度，但需要进行更多的线程间同步操作。

3. "global"：表示使用全局缓存来存储标记的对象，在大型程序中可以更好地利用系统资源，但可能会对程序的运行速度造成一定影响。

通过设置gcMarkWorkerModeStrings变量，用户可以更细致地控制垃圾回收器标记阶段的工作模式，从而优化程序的性能和内存使用。



### work

在Go语言的runtime包中，mgc.go文件是与垃圾回收（GC）相关的代码文件之一。其中，work变量是一个全局变量，主要用于管理和跟踪GC的工作状态。

work变量包括了以下字段：

- wbuf：用于分配工作缓冲区的slice；
- wbuf1：工作缓冲区的备用slice；
- wbuf2：另外一个工作缓冲区的slice；
- ptr：当前正在使用的工作缓冲区指针；
- n：当前分配的工作缓冲区大小；
- barrier：用于保护工作缓冲区的锁（mutex）。

在Go语言的垃圾回收算法中，工作缓冲区是用于存储指向可回收对象的指针的，用于标记这些对象以进行回收。work变量的作用是确保GC的每个阶段都处理了所有指向可回收对象的指针，避免出现漏标（对象没有被标记进行回收）或误标（无需回收的对象被标记）的情况。

工作缓冲区的大小和分配方式是动态调整的，以充分利用系统资源和避免内存溢出。work变量的字段n和ptr是用于管理工作缓冲区大小和分配的指针。同时，备用缓冲区wbuf1和wbuf2也保证了当GC需要在缓冲区溢出时，总能够快速地进行切换和分配可用的缓冲区。

总之，work变量是Go语言垃圾回收算法中至关重要的一个全局变量，它的作用是管理和跟踪GC的工作状态，确保GC能够准确地标记和回收内存中的可回收对象，避免内存泄漏和程序崩溃。



### gcMarkDoneFlushed

gcMarkDoneFlushed是在Go语言的垃圾回收机制中标记已完成的标志变量，它主要用于标记已经完成的垃圾回收扫描操作，并告知系统可以继续执行下一个阶段。

具体来说，gcMarkDoneFlushed变量的作用体现在以下两个方面：

1.标记扫描结束：
在gc扫描过程中，垃圾回收器需要扫描根据当前goroutine的堆栈、静态变量和全局变量等信息，并跟踪对象的引用关系来决定哪些对象是存活的，哪些是垃圾对象。当扫描结束时，gcMarkDoneFlushed变量会被设置为true，表示当前阶段的扫描已经完成。

2.触发垃圾回收：
同时，当已经完成了垃圾回收扫描操作时，系统会检查gcMarkDoneFlushed变量的状态。如果gcMarkDoneFlushed为true，说明当前阶段的根据引用关系扫描已经完成，垃圾对象已经被标记。这时候垃圾回收机制会启动垃圾清理过程对标记的垃圾对象进行回收，释放内存空间。反之，如果gcMarkDoneFlushed为false，说明垃圾回收扫描操作尚未完成，垃圾回收机制会继续扫描直到所有对象全部扫描完成。

综上所述，gcMarkDoneFlushed变量是垃圾回收机制的一个标记，用于标记垃圾回收扫描阶段已完成，并且可以触发垃圾清理操作，从而实现有效的垃圾回收和内存管理。



### poolcleanup

变量名：poolcleanup

作用：这个变量用于指示是否应该在每次垃圾收集后清理池。垃圾收集是指在 Go 程序运行期间由 go runtime 进行的自动内存管理过程。

详细介绍：

在 Go 中，当我们需要使用大量的临时对象时，使用池是一种有效的方式来减少内存分配和垃圾回收的开销。池是带有同步机制的对象池，它使我们可以重用先前分配的对象以提高性能。

go runtime 中的 poolcleanup 变量控制池的清理时间。如果 poolcleanup 的值为 true，则在每次垃圾收集后会清理池。这意味着池中的对象会在 gc 期间不会被处理，因此需要在垃圾收集后进行清理以回收资源。

如果 poolcleanup 的值为 false，则在每次垃圾收集中不会清理池。这可以提高性能，因为不需要花费额外的时间来清理池，但这也可能会导致池中的对象在过多时间内没有被回收，从而占用更多的内存。

在实际编程中，我们可以根据具体的场景来选择是否需要开启 poolcleanup。如果池中的对象的生命周期很短，可能没有必要在每次垃圾收集后清理池。但是，如果池中的对象的生命周期很长，例如在长时间运行的服务器程序中，建议在垃圾收集后清理池，以避免内存泄漏和过多的内存占用。



### boringCaches

变量boringCaches是runtime中的一个缓存结构。它主要用于加速垃圾回收器中的一部分操作。

在垃圾回收器中，需要频繁从堆上分配和回收内存。通过boringCaches，可以实现对部分对象的复用，从而减少分配和回收的次数，提高效率。

boringCaches的具体实现是一个数组，数组中的每个元素都是一个指向C空间的指针。在垃圾回收器的操作中，可以将一些对象存放到boringCaches中，然后在需要使用时直接从boringCaches中获取对应的对象，而无需重新分配内存。当对象不再需要时，可以将其放回boringCaches中以便复用。

boringCaches在不同的垃圾回收器中有不同的用途。在Go 1.5及之前的版本中，boringCaches主要用于保存已分配的对象以供下一次使用。而在Go 1.6及之后的版本中，boringCaches也可以用于保存垃圾回收器中的内部数据结构，从而提高运行时的性能。

总之，boringCaches通过复用一些对象来减少分配和回收的次数，从而提高垃圾回收器的效率。它是垃圾回收器中的一个重要优化措施，也是runtime中重要的一个缓存结构。






---

### Structs:

### gcMarkWorkerMode

gcMarkWorkerMode是Go语言中垃圾回收器中的一种工作模式，用于在并发标记阶段中协调多个标记器的工作。

具体来说，当Go语言的垃圾回收器需要回收内存时，它会启动一个并发的标记阶段，即前面提到的阶段二。在这个阶段中，垃圾回收器会使用多个goroutine并发扫描程序中正在使用的堆中的对象，标记活跃对象和垃圾对象。在多核CPU上，这个并发扫描的过程更加高效。

而gcMarkWorkerMode结构体则是用来控制并发标记过程中的多个标记器的工作，它的主要作用有以下几个：

1. 维护标记队列。在并发标记过程中，标记器会使用work.markfor扫描堆中的对象，标记活跃的对象并放入标记队列中。gcMarkWorkerMode结构体中的markfor属性会存储这个标记队列的指针，用来协调多个标记器之间的工作，避免冲突和重复工作。

2. 控制标记器的运行状态。gcMarkWorkerMode结构体中的stop和shutdown等属性可以控制标记器的运行状态，根据需要中途停止标记器的扫描工作。

3. 提供线程安全的接口。gcMarkWorkerMode结构体中的pushgrey和commitgrey等方法可以提供多个标记器之间的线程安全的接口，避免冲突和数据竞争。

总之，gcMarkWorkerMode结构体是Go语言中垃圾回收器中实现并发标记的重要组成部分，它能够协调多个标记器之间的工作，提高垃圾回收的效率。



### workType

workType结构体是垃圾收集器的核心数据结构之一，主要用于保存和管理工作线程（G），用于协调和执行垃圾收集的任务。该结构体定义如下：

```
type workType struct {
    victim      uintptr           // 下一个被扫描对象的地址
    nobj        uintptr           // 当前扫描对象的数量
    obj         uintptr           // 当前正在扫描的对象地址
    bytesMarked uint64            // 当前线程刚完成标记的 bytes 数量
    scanp       uintptr           // 当前指向的对象或者下一个对象的地址
    gcw         *gcWork           // 当前工作线程的 gcWork 结构指针
    scanner     *scanState        // 扫描过程中的状态信息
    nest        int32             // 嵌套的 for 循环层数（用于调试）
    // ...（其他字段省略）
}
```

其中，主要有以下几个字段：

- victim：表示下一个要被扫描的对象的地址。在Concurrent Mark阶段，扫描线程会根据victim指针逐个扫描heap中的对象。
- obj：表示当前正在被扫描的对象的地址。如果workType结构体中其他字段没有被设置，则obj指向victim所指向的对象，否则obj指向其他对象。
- bytesMarked：表示当前线程刚完成标记的字节数。在Concurrent Mark阶段，各工作线程在扫描heap时会逐渐标记各对象是否是存活对象，bytesMarked则记录已经标记的存活对象占用的内存大小，注意它只记录当前线程标记的对象大小，因此需要用atomic.AddUint64原子加操作统计所有线程标记的内存大小。
- gcw：表示当前工作线程的gcWork指针，用于追踪该工作线程的扫描进度。
- scanner：表示扫描过程的状态信息，包括扫描模式（是Concurrent还是Fallback模式）、标记过程中遇到的全局对象（用于发现重要存活对象，如Goroutine、栈等），以及用于增量式标记的跟踪状态（包括指针上一个标记的层数、扫描时上次扫描到的地址、是否正在扫描根集等）。

总的来说，workType结构体在垃圾收集器中发挥着至关重要的作用，它是维持工作线程和垃圾收集器协同工作的关键所在。通过该结构体，垃圾收集器可以跟踪工作线程的扫描进度、协调各工作线程的扫描工作、发现全局存活对象等，从而实现对整个堆中存活对象的准确标记和回收。



### gcMode

gcMode结构体定义了运行时的垃圾回收模式以及一些相关参数，用于控制垃圾回收的行为。具体来说，gcMode结构体包含了以下字段：

- kind：表示垃圾回收的模式，有“off”、“m”、“ms”、“mw”四种取值。分别代表关闭垃圾回收、标记-清理模式、标记-整理模式、标记-整理模式（针对只有一个内存堆的情况）。
- triggerRatio：表示堆使用率超过多少时触发垃圾回收。默认为0.5。
- heapMinimum：表示堆最小大小。默认为8192KB。
- heapMaximum：表示堆最大大小。默认为0，即不限制。
- stackGuardMultiplier：表示协程的栈空间大小比例，默认为0.875。

通过调整gcMode结构体的字段值，可以灵活地控制垃圾回收的触发和行为，进而影响程序的性能和稳定性。例如，可以通过修改triggerRatio和heapMaximum来调整垃圾回收的频率和堆大小，从而优化程序的内存使用和垃圾回收的效率。



### gcTrigger

在 Go 语言中，gcTrigger 结构体用于控制 GC 触发的条件。在 Go 运行时，GC 需要在某些情况下触发，以回收不再使用的内存。gcTrigger 结构体中定义了一些用于控制 GC 触发条件的属性和方法。

gcTrigger 结构体包含以下属性：

1. memstats：用于存储内存统计信息的结构体。

2. gcSystemStartTime：系统启动时间。

3. lastGCPauseEndTime：上一次 GC 暂停结束时间。

4. triggerRatio：可被 GC 回收的堆内存占总内存的比例阈值。

5. triggerBytes：可被 GC 回收的堆内存大小的阈值。

6. lastTriggerRatio：上一次垃圾回收时可被回收的堆内存占总内存的比例。

7. lastTriggerBytes：上一次垃圾回收时可被 GC 回收的堆内存大小。

gcTrigger 结构体包含以下方法：

1. setTriggerRatio(ratio float64)：设置 gcTrigger 的 triggerRatio 属性。

2. setTriggerBytes(bytes uint64)：设置 gcTrigger 的 triggerBytes 属性。

3. trigger()：判断是否需要触发 GC。当可被 GC 回收的堆内存大小超过 triggerBytes 或可被回收的堆内存占总内存的比例超过 triggerRatio 时，触发 GC。

总之，gcTrigger 结构体中的这些属性和方法用于控制 GC 的触发条件，以确保 GC 能够在适当的时间回收内存，提高 Go 程序的性能和稳定性。



### gcTriggerKind

在Go语言中，GC（垃圾回收）被用来自动管理内存。有时，垃圾回收程序会在应用程序正在运行时突然触发，这可能会导致不必要的性能开销。

为了避免这种情况，Go语言在runtime/mgc.go中定义了一个名为gcTriggerKind的结构体，用于描述垃圾回收触发的类型。

gcTriggerKind结构体的作用是描述垃圾回收的类型，包括它是由内存分配还是使用量触发的、它是否由手动触发、以及它的严格程度等等。

具体来说，gcTriggerKind结构体包含以下属性：

- kind：触发垃圾回收的类型。包括以下四种类型：gcTriggerCycle、gcTriggerFraction、gcTriggerHeap、gcTriggerManual。

- gcTriggerCycle：当前GC在固定间隔（gcController的cycle时间）之后触发。也就是说，触发GC是根据一定的时间间隔来触发的。

- gcTriggerFraction：当堆大小增加了指定的分数（gcController的heapMinimum和heapGoal参数）时，将触发GC。根据该参数，GC触发的频率将会随着堆的增长而加速，这样可以根据堆的使用情况来触发GC。

- gcTriggerHeap：当堆大小达到指定的值（gcController的heapSizeGoal参数）时，将触发GC。

- gcTriggerManual：这个是人工触发的GC。

- gcTriggerScale：触发GC的严格性级别。级别越高，GC触发的频率也越高，但是程序的性能也会更差。共有四个级别，从0到3。

- gcTriggerDisable：禁用垃圾回收

通过gcTriggerKind结构体，我们可以灵活地控制垃圾回收的触发类型和频率，从而平衡应用程序的性能和内存使用。



### gcBgMarkWorkerNode

在go/src/runtime/mgc.go文件中，gcBgMarkWorkerNode结构体是用于表示与BGMarkWorker相关的状态的一种类型。它的定义如下：

```go
// A gcBgMarkWorkerNode is a node in the concurrent work queue of
// the background mark worker. Thus it is an element of the doubly-
// linked list defined by the gcBgMarkWorkerState.
//
// One gcBgMarkWorkerNode is used to represent both the state of being
// enqueued for work and the state of being processed but not yet
// enqueued for the next phase of the mark worker.
//
// Depending on the state of the gcBgMarkWorkerState, it is mutable
// from one of potentially many concurrent GC workers or the mark worker
// itself; however, each transition MUST be made under the appropriate
// lock.
//
// Rather than label each and every field below as mutable or not,
// it is crucial to keep in mind the following invariant:
//
// The fields are only modified when the state is actively enqueued 
// or being dequeued. This doesn't overlap with the currently active 
// gcBgMarkWorkerState because each phase of the mark worker has a 
// private immutable snapshot of the work queue that is shuffled 
// independently of any other phase. If two phases are executing 
// concurrently (which would happen during "on-the-fly" single-P 
// marking or D phase), then they split the work hand over fist 
// rather than interleaving.
type gcBgMarkWorkerNode struct {
    // A node struct must start with a gcWork.
    work gcWork

    // Popped off the current mark queue.
    next *gcBgMarkWorkerNode
    prev *gcBgMarkWorkerNode

    // live == live objects; dead == garbage.
    leaf bool
    live uintptr
    dead uintptr
    target  *gcBgMarkWorkerState
}
```

在并发垃圾回收器中，每个后台标记worker都维护着一个并发工作队列（concurrent work queue）。gcBgMarkWorkerNode结构体就是这个队列中的元素。它有以下几个成员：

- work：表示要被处理的GC Work（例如存活对象的标记工作，死对象的清理工作）。
- next / prev：表示下一个和前一个gcBgMarkWorkerNode元素。
- leaf：表示该元素处理的是可达的存活对象（即“绿子树节点”），还是无法访问的垃圾对象（即“灰子树节点”）。
- live / dead：分别表示其处理的存活对象和垃圾对象的数量。
- target：表示这个元素所在的众多GC work队列的指向。

gcBgMarkWorkerNode的主要作用是：

- 代表了一个需要后台标记worker处理的GC Work。
- 作为后台标记worker的并发工作队列中的一个元素。
- 记录了已标记的存活对象和已确定为垃圾的对象的数量。



## Functions:

### gcinit

gcinit函数是Go语言垃圾回收机制的初始化函数，它会在程序启动时被调用。gcinit会负责以下几个任务：

1. 设置全局变量gcphase的初始值为_gcoff。gcphase用来记录当前垃圾回收的阶段，具体取值可以是_gcoff、_gcmark、_gcmarktermination、_gcSweep、_gcSweepDone等，不同阶段对应不同的垃圾回收操作。

2. 初始化工作线程。在Go语言的垃圾回收过程中，需要使用多个工作线程来并行标记和清理内存。gcinit会创建一组工作线程，并将它们初始化，准备好执行垃圾回收任务。

3. 设置全局变量gcController。gcController用来控制垃圾回收的流程，包括阶段转换和线程调度等，gcinit会初始化gcController并将其设置为正在进行的垃圾回收的控制器。

4. 初始化P的本地缓存。在Go语言的垃圾回收过程中，每个工作线程都会维护一个本地缓存（Local Cache），用来保存一些临时变量和对象。gcinit会将所有工作线程的本地缓存初始化，并将它们和全局变量gcController进行关联。这样，每个工作线程都可以将本地缓存中的内容同步到全局缓存中，保证垃圾回收的正确执行。

总的来说，gcinit函数的作用是初始化垃圾回收机制的各个组件，为垃圾回收做好准备工作，保证其正确、高效地执行。



### gcenable

在Go语言中，垃圾回收是由运行时系统负责的。 gcenable 是运行时系统中的一个函数，其作用是启用垃圾回收。

具体来说，gcenable 函数会将标记位从 _GCoff 修改为 _GCmark，并触发一次强制执行的垃圾回收操作。此外，gcenable 函数还会检查其他相关参数，如 gcpercent（触发垃圾回收的最低堆使用量）和 maxprocs（最大并发执行的进程数），以确保垃圾回收系统正常运行。

在启用垃圾回收之前，应用程序可能需要完成一些初始化工作。例如，创建一些 goroutine、分配内存等等。一旦这些初始化完成，应用程序可以调用 gcenable 函数，以启用垃圾回收系统。

总之，gcenable 函数是 Go 运行时系统中的一个关键功能，它使得垃圾回收可以正常工作，确保了 Go 代码的稳定性和可靠性。



### setGCPhase

setGCPhase函数是用于设置当前的垃圾回收阶段的函数。垃圾回收阶段包括STW (Stop-The-World)、Mark、MarkTermination、Sweep、SweepTermination等。该函数如果被调用时未处于STW阶段，会抛出错误。

具体来说，当程序需要进行垃圾回收时，Go runtime会按照垃圾回收阶段的顺序进行不同的操作。setGCPhase函数就是用来在不同阶段中切换的函数，例如在Mark阶段时调用setGCPhase函数来切换到MarkTermination阶段。另外，在切换不同阶段时需要进行一些涉及到全局状态的操作，例如更新P的状态、清空缓存等，经常会调用一些相关的函数，以便在不同阶段中进行必要的全局操作。

总之，setGCPhase函数主要作用是在垃圾回收过程中切换不同的阶段，并进行一些涉及到全局状态的操作。它是Go runtime中垃圾回收部分的关键函数之一。



### pollFractionalWorkerExit

pollFractionalWorkerExit函数是用于检查并退出Goroutine的函数。在并发垃圾回收（Concurrent Garbage Collection，简称CGC）中，有专门的Goroutine用于协助GC完成其中的一些任务，这些Goroutine称为helper goroutine，它们是在应用程序正常执行的情况下运行的。但是，由于这些helper goroutine不会自行退出，它们可能会占用应用程序中的一些资源并降低性能。pollFractionalWorkerExit函数的作用就是用于检查这些helper goroutine是否有退出的机会，如果可以退出，就会强制让它们退出。

在函数实现中，pollFractionalWorkerExit会获取helper goroutine的一份副本，通过检查副本的状态来判断是否有退出的机会。如果helper goroutine的副本状态处于waitexit状态（这意味着helper goroutine已经准备好退出了），则将其从goroutine队列中移除，并设置helper goroutine的状态为exited。如果helper goroutine的状态不是waitexit，则仅仅是将其副本状态设置为pollidle状态，以便后续轮询。

pollFractionalWorkerExit函数一般是由pollFractionalWorker函数调用的，用于轮询并退出helper goroutine。通过这种方式，可以保证helper goroutine不会无限制地占用资源，从而保持应用程序的性能稳定。



### GC

GC是Golang的垃圾回收机制，它的作用是在程序运行的过程中对内存进行自动化管理，保证程序能够自动回收不再使用的内存，防止内存泄露和程序崩溃。

在Go语言的GC实现中，mgc.go文件中的GC函数是其中最为核心的部分，它负责对程序的内存进行回收。GC函数会使用可达性分析算法来检查哪些内存是可达的，哪些是不可达的，对于不可达的内存会自动回收。

GC函数的流程大致如下：

1. 通过遍历G、P、M等数据结构获取所有的根对象，根对象是程序中所有被引用的对象的入口点。

2. 对所有根对象进行遍历，标记所有可达的对象为活动对象，不可达的对象则是垃圾对象。

3. 对于所有不可达的对象，调用对象的finalizer，让其执行垃圾回收前的清理工作。

4. 对于所有不可达且没有finalizer的对象，直接将其回收。

5. 经过垃圾回收后，将回收的内存重新分配给空闲列表，以便下一次的内存分配使用。

需要注意的是，在执行GC函数时，程序需要停止所有的goroutine并暂停程序的正常运行，等待GC完成后再恢复程序的执行，这个过程称为“Stop The World”（暂停世界）。

GC函数会在程序的一些关键点自动触发，例如当堆大小达到一定阈值、或当程序调用了runtime.GC()函数时，GC函数会对内存进行自动回收。



### gcWaitOnMark

gcWaitOnMark是runtime中的一个函数，主要目的是等待标记阶段的完成。

在Go语言的垃圾回收中，标记阶段是GC的一部分。在这个阶段中，GC会标记所有已分配对象中仍然在使用的对象。标记完毕后，GC就会知道哪些对象可以被回收，哪些对象仍然在使用中。但是，等待标记阶段的完成可能会降低应用程序的性能，因此，gcWaitOnMark函数的作用就是提高并行性，加快标记阶段的完成。

具体来讲，gcWaitOnMark函数的工作是阻塞当前goroutine，等待标记阶段完成。如果标记阶段已经完成，那么函数立即返回。否则，函数会调用gcController、gcBgMarkWorker和gcMarkDone等函数来进行协调和等待。当标记阶段完成后，函数会返回并继续执行后续操作。

总之，gcWaitOnMark函数的作用就是等待GC标记阶段的完成，并最大限度地提高并行性，加快标记阶段的完成。



### test

在go/src/runtime/mgc.go文件中，test func主要是用于：
1. 验证小对象内存分配是否正常
2. 确保可达性分析过程不会出现死循环，保证垃圾回收不会影响应用程序的正常运行

具体来说，test func会创建一些小对象（allocs）并且在子协程里分配和释放它们，测试系统是否可以正常扩展堆空间。如果分配和释放正常完成，则说明内存分配系统正常工作。在分配和释放这些对象时，堆大小也会动态调整以适应需要。

测试还会模拟垃圾回收过程，即执行强制垃圾回收，然后检查回收后的堆大小是否正确。如果大小正确，则说明垃圾回收器正常工作，并且不会对应用程序的正常运行产生影响。

测试还会验证标记过程是否正常完成，即在分配和释放这些小对象后，会标记所有可达对象，以确保可达性分析过程不会进入死循环状态。

总之，test func是用于验证垃圾回收器正常工作，并确保不会对应用程序的正常运行产生影响的重要函数。



### gcStart

gcStart是Go语言运行时底层的垃圾回收器（GC）开始执行的函数。它主要有以下几个作用：

1. 初始化GC相关数据结构：gcStart会初始化标记位图、堆栈缓存、roots集合等数据结构。这些数据结构是GC执行过程中必要的。

2. 标记根对象：GC开始前需要明确哪些对象是可达的根对象，即哪些对象是从外部引用的或者是全局变量。gcStart会标记这些根对象以确保它们不会被回收。

3. 调用标记阶段的回调函数：GC执行过程中有一些回调函数会被调用，比如记录对象信息的回调、处理finalizer的回调等。gcStart会调用标记阶段的回调函数。

4. 标记完整个堆：经过前面的准备工作后，gcStart会开始执行标记阶段，遍历整个堆，将可达对象标记为活动对象，未被标记的对象为垃圾对象。

5. 执行清理阶段：当标记阶段结束后，gcStart会执行清理阶段，将所有未被标记的对象（即垃圾对象）回收。清理过程包括与其它goroutine协作完成的事宜，例如中断时间的清理，将被释放的内存归还给操作系统等。

总体来说，gcStart是Go语言运行时GC的引擎，它通过在标记阶段标记活动对象，清理阶段回收垃圾对象，并在整个垃圾回收过程中调用相关回调函数，保证了Go语言程序的内存安全和高效使用。



### gcMarkDone

在Go语言中，gcMarkDone函数是用于标记已完成的标记阶段的。在垃圾回收过程中，标记阶段的任务是从根对象出发，标记所有可达的对象。标记完成后，就可以进行垃圾回收。gcMarkDone函数的作用是将已完成的标记阶段状态设置为done，以便进行下一阶段的垃圾回收。在gcMarkDone函数中，通过将相关的标记位设置为已完成来标记标记阶段已完成。在标记阶段完成后，我们可以安全地清理任何未被使用的内存。同时，gcMarkDone函数还会调用gcSweep函数来执行扫描操作，并将未被使用的内存释放回系统。

总之，gcMarkDone函数是用于标记完成垃圾回收中的标记阶段并清理未使用的内存。它是Go语言垃圾回收流程中非常重要的一个环节。



### gcMarkTermination

gcMarkTermination函数是垃圾回收器在并发标记阶段结束时调用的函数之一，它的作用是将一些未被扫描的对象标记为黑色，并把它们添加到黑色对象列表中。具体来说，gcMarkTermination会做以下三件事情：

1. 扫描grayobject列表，将其中仍未被扫描的对象标记为黑色。

2. 扫描graySpecial列表，将其中的特殊对象（如g或m）标记为黑色。

3. 将一些未被扫描的堆对象（如跨度对象或大对象）标记为黑色，这些对象在扫描阶段可能被遗漏。

通过以上三个步骤，gcMarkTermination函数可以确保所有的可达对象都被正确标记为黑色，以使垃圾回收器能够正确识别哪些对象应该被保留，哪些对象应该被释放。

总之，gcMarkTermination函数是垃圾回收器并发标记阶段的一个重要组成部分，它通过扫描未被正确标记的对象，最大程度地减少垃圾回收的误判和漏判，从而更加高效地回收内存。



### gcBgMarkStartWorkers

gcBgMarkStartWorkers函数是Go语言标准库中runtime包中mgc.go文件中的一个函数，这个函数的主要作用是启动后台标记的工作。在go程序中，当堆中对象数量达到一定阈值时，需要进行垃圾回收，这个阈值就是GC触发器的阈值。在垃圾回收过程中，需要标记和回收不再使用的对象，这个过程需要占用CPU和内存资源，会影响应用程序的性能。

gcBgMarkStartWorkers函数的作用就是通过启动后台标记工人来降低GC对应用程序的影响。在这个函数内部，会检查当前是否有后台标记工人在运行，如果没有，则根据当前CPU数量创建n个后台标记工人，并启动它们，这样可以并发的运行标记过程，从而减少GC对应用程序造成的瓶颈。这些后台标记工人会在标记过程中协调采样器和全屏扫描工作，以及在标记完成后执行清洗、压缩和其他后处理工作，最终完成垃圾回收的过程。

(gcBg)MarkStartWorkers函数是gcBackgroundMarkRunLoop函数调用的子函数，而(gcBg)MarkStartWorkers函数起到类似创建并启动goroutine的作用，即为Mark State申请和释放P以及创建和运行G的工作。在输入参数的情况下，返回一组操作future（并且保证他们都进行，以及回到调用者）。对于未处理可能的错误，(gcBg)MarkStartWorkers处理了它们并且不会退出。从某种意义上说，这种方法实施了类似于实现操作系统中涉及异常和保护策略的方式。



### gcBgMarkPrepare

gcBgMarkPrepare是Go语言运行时中mgc.go文件中的一个函数，其作用为在后台标记阶段开始前准备标记。下面是该函数的详细解释：

1. 首先，该函数会获取gcController状态（gc背景标记控制器）并检查标记过程是否被取消（stopped）。

2. 接着，它会检查待处理的goroutine数量，如果没有，则更新gcMarkWorkerMode的状态为markDone并返回。

3. 如果仍有待处理的goroutine，它会首先检查当前P（处理器）是否已准备好使用，并且当前的gcMarkWorkerMode状态是否为“推送”（pullWait）。

4. 如果是，则将gcMarkWorkerMode更新为“推送”（pulling），并返回将该P置于推卡片队列中的信号。这意味着其它goroutine将在P准备好时推送任务。

5. 如果当前P没有准备好，则将gcMarkWorkerMode更新为“阻断”（blocking），并在进入阻塞之前重新检查是否进入状态stopped。如果状态未更改，则当前G（goroutine）将会进入休眠状态，直到某个P可用。

6. 如果当前gcMarkWorkerMode状态为“推送”，则检查是否有等待推卡片的goroutine。如果没有，则更新gcMarkWorkerMode为“推送完成”（markWait），返回并继续进行标记。

7. 如果有等待推卡片的goroutine，则将它们放置在卡片推送队列的末尾，并将gcMarkWorkerMode更新为“推送完成”（markWait），返回并继续进行标记。

总之，gcBgMarkPrepare函数用于在标记阶段启动前准备标记，根据当前状态（例如，准备好的处理器数量和待处理的goroutine数量）推送或阻塞goroutine以优化标记过程。



### gcBgMarkWorker

gcBgMarkWorker函数是Go语言的垃圾回收器背景标记阶段的工作者函数，它的作用是协调和执行背景标记阶段的任务。

在Go语言中，垃圾回收分为两个阶段：标记阶段和清理阶段。在标记阶段中，GC会标记哪些对象是可达的，而背景标记则是在程序继续执行的同时进行的。gcBgMarkWorker函数就是在背景标记阶段中发挥作用的。

gcBgMarkWorker函数的具体工作流程如下：

1. 获取P（processor）：gcBgMarkWorker函数会获取一个P（processor），该P将用于执行G（goroutine）中的任务。

2. 获取任务：gcBgMarkWorker函数会获取一个任务，该任务是从全局队列中获取的，用于标记已分配的对象。

3. 执行任务：gcBgMarkWorker函数会将获取到的任务分配给P执行。

4. 解锁：执行完任务后，gcBgMarkWorker函数会解锁其他任务，然后再次获取P和任务执行。

5. 循环执行：gcBgMarkWorker函数会一直循环执行2-4步，直到全局队列中没有待处理任务为止。

通过这些步骤，gcBgMarkWorker函数可以协调和执行背景标记阶段的任务，并进行高效的垃圾回收。



### gcMarkWorkAvailable

gcMarkWorkAvailable函数是Go语言中的垃圾回收机制（GC）的核心函数，在gcMarkWorkAvailable函数中，会检查是否有更多的垃圾标记工作需要执行。如果需要执行更多的垃圾标记工作，则它会将gcMarkWork.wakeup中的goroutine唤醒。

这个函数的主要作用是确保标记阶段中的所有垃圾标记工作都能够得到充分的处理，从而最大程度地减少垃圾回收的时间和频率。它在标记阶段中被多次调用，以确保所有的垃圾标记工作都能得到处理。

在Go语言中，GC的性能是非常重要的，因为它直接影响了代码的执行速度和响应时间。gcMarkWorkAvailable函数通过检查是否有更多的垃圾标记工作需要执行，让Go语言的垃圾回收机制更加健壮和高效。



### gcMark

gcMark函数是Go语言中 Garbage Collector 的其中一个核心函数，其主要作用是标记所有的存活对象。

在Garbage Collection算法中，标记阶段是非常重要的一环。简单来说，gcMark函数的作用就是通过遍历对象图，来标识哪些对象是被程序引用的，还需要被保留下来，而哪些是垃圾，需要被回收。

gcMark函数的执行流程如下：

1. 初始化根对象集合

该函数会在程序的全局变量、虚拟机栈、寄存器等计算机内存区域中找出所有的指针对象，并将其加入到根对象集合中。

2. 遍历对象图

接下来，该函数会从根对象集合中取出一个对象，进而遍历相关的所有对象，标记它们为活动对象，即便标记它们的颜色为灰色。

3. 标记灰色对象的相关对象

在上一步中，所有被标记为灰色的对象都是需要被细节扫描的。在这一步中，gcMark函数会遍历这些灰色对象引用的所有对象，并标记它们为活动对象。

4. 标记结束

gcMark函数将结束标记过程，并将被标记为活动对象的对象从灰色变为黑色，此时标记阶段结束。

总之，gcMark函数的作用是负责标记存活对象，使得垃圾回收器能够在后续的阶段中准确地回收无用对象，从而保证程序的内存安全和稳定性。



### gcSweep

gcSweep这个函数的作用是清除不再使用的对象，并将其返回给还未使用的空间池。gcSweep函数实现了垃圾回收中的标记-清除阶段，在清除过程中，它扫描程序中的所有对象，并标记那些还在使用的对象，将不再使用的对象清除。

在清除阶段，gcSweep函数会遍历所有的堆区，找到已经死亡的对象，然后将它们的空间返回给空闲内存池，以便能够为后续的对象分配更多的内存。gcSweep函数会使用一个指针列表，该列表存储所有被分配的对象的指针，以便在清除时进行遍历。

在清除阶段的最后，gcSweep函数会更新垃圾回收器的状态，并准备进入下一个阶段。该函数还会更新堆的统计信息，在垃圾回收的过程中记录已处理的对象数量，并更新堆的大小等信息，以便在下一次垃圾回收时使用。

总之，gcSweep函数是垃圾回收的关键之一，可以通过清除不再使用的对象，释放内存并提高程序的性能。



### gcResetMarkState

gcResetMarkState是在Go语言中垃圾回收器(gc)的运行过程中调用的一个函数，它的作用是重置垃圾回收器中与标记相关的状态，以便下一轮垃圾回收能够正确地进行。

在Go语言中的垃圾回收器执行过程中，需要分为两个阶段:标记阶段和清扫阶段。在标记阶段中，会从根对象出发，遍历所有能够到达的对象，并将这些对象标记为活动对象。在这个过程中，为了防止对象被多次标记或者不被标记的情况出现，需要记录一些状态，并在完成标记后进行清理。

gcResetMarkState函数就是负责重置这些状态的函数。具体来说，它会清理各种指针标记位，还会重置某些内存区域的状态，以防止垃圾回收器在下一轮回收时受到干扰。

总之，gcResetMarkState函数是垃圾回收器中关键的重置函数之一，它确保垃圾回收器在下一轮运行前的状态是正确的，这样才能准确地找到所有的垃圾对象进行回收。



### sync_runtime_registerPoolCleanup

sync_runtime_registerPoolCleanup是在Go语言中用于垃圾回收的一个函数，定义于runtime/mgc.go文件中。它的作用是向全局池（global pool）中注册一个清理函数（cleanup function），以便在每个垃圾回收周期结束后自动执行。

具体来说，垃圾回收器在回收过程中会创建一些临时对象，例如内存块和临时指针等。这些对象会存在于全局池中，等待下一轮垃圾回收周期结束后清理。如果不及时清理这些对象，将会导致内存泄漏和系统性能下降。

为了解决这个问题，Go语言提供了一个在垃圾回收周期结束后自动执行的全局清理函数机制。这个函数由sync_runtime_registerPoolCleanup注册，并在垃圾回收器的全局清理函数调度器中执行。它的作用是清理全局池中的临时对象，以便下一轮垃圾回收器能够重新使用它们。

总之，sync_runtime_registerPoolCleanup函数的作用是向全局池中注册一个清理函数，用于自动清理垃圾回收中产生的临时对象，以提高系统的性能和稳定性。



### boring_registerCache

boring_registerCache是Go语言的垃圾回收器（GC）中的一个函数，它的作用是注册一块新的缓存区域，以供垃圾回收器使用。

在Go语言中，当程序中的某个内存块不再被使用时，垃圾回收器会将这个内存块标记为垃圾，并将其回收。GC会扫描程序的内存空间，识别出那些已经不再使用的内存块，然后将它们释放回操作系统。为了进行这个操作，GC需要维护一些缓存区域，以记录哪些内存块可以被回收。

boring_registerCache就是用来注册这些缓存区域的函数。当程序创建一个新的对象时，GC会调用这个函数来为该对象分配一个缓存区域。当这个对象不再使用时，GC会将它从缓存区域中删除，并将它标记为垃圾。

这个函数的名字中包含了“boring”一词，意思是它是垃圾回收器中的一个平凡的函数。因为大多数情况下，程序员很少需要直接调用这个函数，因为Go语言的GC会自动进行垃圾回收，而无需程序员手动干预。



### clearpools

clearpools函数定义在mgc.go文件中，它的原型如下所示：

```go
// clearpools flushes all cached memory in the runtime.
// This includes mcache and deferred mspan frees.
// It must be called without any other locks held.
// Otherwise it is free to block or allocate memory itself.
// This function is not safe for concurrent invocation.
func clearpools() {
    // ...
}
```

该函数的作用是清空所有的内存池（memory pool），包括mcache和deferred mspan释放。内存池是在运行时（runtime）中由内存管理器（memory manager）分配和管理的一组预定义大小的内存块。这些内存块用于跟踪和分配堆对象和栈帧等运行时数据结构。内存池减少了在分配内存时的系统调用次数，从而提高了运行时的性能。

clearpools函数必须在没有其他锁被持有的情况下调用，否则它可能阻塞或分配内存。此外，该函数不适用于并发调用。因此，它通常在垃圾回收期间被调用，以确保垃圾回收器能够正确访问和管理所有内存池。它还可以在程序退出时调用，以释放所有未释放的内存池并确保程序退出时没有内存泄漏。



### itoaDiv

itoaDiv是一个用于将一个无符号整数转化成字符串的函数，它被用于扫描堆中对象，给对象分配编号，并将这些编号序列化成字符串。

itoaDiv函数的作用如下：

1. 将一个无符号整数num转换为字符串，使用的是基于10进制的转换方式。
2. 通过除法、取模来实现转换，divVal和divMod分别用来表示除数和模数。
3. 从低位到高位依次填充字符串，num被divVal整除的结果不断被填充到字符串的末尾，直至num被divVal整除的结果为0。
4. 序列化过程中默认使用的字符集是ASCII表中的数字字符。

这个函数主要用于在进行垃圾回收时，将对象的编号序列化成字符串，以便于对这些对象进行跟踪和调试，并在一定程度上提高了垃圾回收的效率。



### fmtNSAsMS

在go/src/runtime中，mgc.go文件中的fmtNSAsMS函数的作用是将纳秒时间格式的数字转换为人类可读的毫秒时间格式的字符串，例如将1000000转换为"1ms"。

该函数接收一个表示纳秒时间的int64类型参数，首先会根据这个数值来确定要转换成多少毫秒，然后再根据毫秒数的大小来选择不同的单位（ms, us, ns）。最后，将转换成的值和单位拼接起来并返回一个字符串。函数的实现过程中还包括了数值的四舍五入处理，以便得到更精确的结果。

这个函数通常用于在GC过程中输出日志和统计信息，帮助开发人员了解GC的执行情况和性能表现。



### gcTestMoveStackOnNextCall

gcTestMoveStackOnNextCall函数是一个用于测试的辅助函数，主要用于模拟在垃圾回收过程中，在函数调用时需要将当前堆栈移动到新的堆栈上的情况。

具体来说，当垃圾回收器对堆进行垃圾回收时，会扫描堆栈中的指针以确定哪些对象在使用中，哪些对象可以被释放。但是，在进行垃圾回收的时候，会出现堆栈需要移动的情况。因为垃圾回收的过程中，会存在压缩堆或者调整堆的情况，如果此时继续向堆栈上分配内存，会影响以后的垃圾回收，所以需要将已经分配的堆栈移动到新的地方。

gcTestMoveStackOnNextCall函数就是用于模拟这种场景的测试函数，当该函数调用时，会返回一个新的栈地址，然后将当前堆栈移动到这个新的地址上，从而模拟在垃圾回收过程中堆栈需要移动的情况。

可以看出，gcTestMoveStackOnNextCall函数在运行时垃圾回收的时候起到了非常重要的作用，它可以帮助开发人员测试代码在垃圾回收过程中的正确性，而不需要真的进行堆栈移动操作，从而提高了开发效率和代码质量。



### gcTestIsReachable

gcTestIsReachable是用来测试一个对象是否可以被标记为可达状态的函数。

在垃圾回收器中，需要扫描整个堆来标记所有可达的对象。对于每个需要扫描的对象，gcTestIsReachable函数会被调用来确定该对象是否是可达的。如果对象可达，它将保留在堆中，并继续被使用。否则，它将被垃圾回收器标记为不可达，并在下一次垃圾回收时被清除。

gcTestIsReachable函数的实现方式取决于对象的类型。对于常规对象，它将检查对象是否在某个根对象 (例如堆栈或全局变量) 中被引用。对于不同类型的对象，例如字符串、函数等，它有不同的实现方式。

总之，gcTestIsReachable函数是垃圾回收器中的重要组成部分，用于确定哪些对象需要保留在堆中，哪些对象可以被垃圾回收器清理。



### gcTestPointerClass

在Go语言中，垃圾回收是非常重要的一个功能，负责回收不再使用的内存，避免内存泄漏等问题。在Go语言中，垃圾回收是通过扫描内存中的指针来实现的，而GC中的指针分类就是根据指针的类型进行分类，以便于GC可以更好地识别和回收内存。

在Go语言的runtime/mgc.go文件中，gcTestPointerClass函数是用于测试一个指针所属的指针类别的函数。一个指针的指针类别决定了GC如何扫描它，以及垃圾回收期间是否可能需要复制它指向的对象。gcTestPointerClass函数使用一个特殊的标记位来确定指针的类型，判断该指针是否是指向堆对象的指针，从而将该指针标记为有效的指针，避免垃圾回收误删除该指针引用的对象。此外，gcTestPointerClass函数还通过检查指针所处的页的状态来确定该指针指向的对象是否在当前的堆页中，从而帮助垃圾回收器回收该对象。

总之，gcTestPointerClass函数在Go语言的垃圾回收中起到非常重要的作用，它帮助GC识别和管理指针，以确保内存的正确回收和利用。



