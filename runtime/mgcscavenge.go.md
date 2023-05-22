# File: mgcscavenge.go

mgcscavenge.go 是 Go 语言标准库中的 runtime 包中一部分垃圾回收器（GC）的实现。它的作用是在程序运行过程中定期进行垃圾回收，以清理被程序中断的对象占用的内存，使得这些内存可以被重新使用。

具体来讲， mgcscavenge.go 实现了一种名为 "Scavenger GC" 的 GC 策略。在Scavenger GC策略中，程序首先将内存空间分为两个区域：新生代和老年代。新生代是具有短寿命的对象的存储区域，一般使用eden，from，to三个空间进行管理，新对象从eden开始分配，并尽可能长时间停留在eden空间中，而存活到一定期限的对象会被移动到from空间中。当from空间中的对象数量或大小达到一定阈值后，GC就会触发，此时将会把from空间中存活的对象复制到to空间中，不存活的对象直接清除， from与eden进行角色调换，这样eden变为新的待分配空间。老年代是用于存放具有长寿命的对象的区域。老年代中的对象不再被GC直接回收，而是通过另一种GC策略来回收。

mgcscavenge.go 实现的Scavenger GC 策略在实现上相对较简单，一般适用于实时性要求不高且内存占用不大的场景。但相比其他GC策略来说，它的效率稍低。如果需要更高的GC效率，程序可以考虑使用更复杂的GC策略。




---

### Var:

### scavenge

在Go语言的垃圾回收器中，"scavenge"（扫描）是指扫描heap的过程，找到所有可以被清理的垃圾对象并将它们回收。"mgcscavenge.go"文件中的"scavenge"变量用于标识当前是否正在进行扫描操作。

具体来说，"scavenge"变量是一个bool类型的变量，用于记录当前是否正在进行扫描。如果"scavenge"为true，说明垃圾回收器正在扫描操作，否则为false。

在执行垃圾回收时，首先会先检查"scavenge"的值，如果为true，则直接返回，否则将"scavenge"赋值为true，并开始执行扫描操作。当扫描操作完成后，将"scavenge"重新赋值为false。

通过"scavenge"变量，可以有效地避免多个扫描操作同时执行，从而提高垃圾回收的效率和减少内存占用。



### scavenger

在go/src/runtime/mgcscavenge.go中，scavenger是一个结构体类型变量，它存储了垃圾回收控制器的一些状态和信息。具体来说，scavenger包含以下字段：

1. Lock：一个互斥锁，用于对scavenger进行加锁，以在多线程编程时确保它的访问安全。

2. work：一个Work pool结构体类型变量，用于管理所有需要执行的工作任务。

3. startTime：采用Unix时间戳格式记录垃圾回收开始的时间。

4. first：一个*mspan指针，指向当前堆中的第一个mspan。

5. last：一个*mspan指针，指向当前堆中的最后一个mspan。

6. gcController：一个gcController类型变量，存储垃圾回收控制器的状态和信息。

7. pagesInUse： 一个用于记录当前扫描的内存页的slice。

8. free：一个用于记录没有使用的内存页的slice。

9. sparedSafe分配器可以安全地保留的余数页面数

上述字段的作用如下：

1. Lock确保对scavenger的访问是安全的，以免出现意外的并发问题。

2. work跟踪需要执行的工作任务，确保垃圾回收器完成所有必要的任务。

3. startTime记录垃圾回收开始的时间，以便监控垃圾回收的持续时间。

4. first和last字段用于跟踪堆的范围，使垃圾回收器能够扫描整个堆。

5. gcController记录垃圾回收器的状态和信息，包括时间间隔、目标内存使用量等等。

6. pagesInUse和free分别记录当前扫描的和未使用的内存页的信息，以帮助垃圾回收器确定哪些内存已经被使用，哪些可以回收。

7. sparedSafe记录分配器可以安全地保留的剩余页面数，以确保堆的扩展与系统内存的可用性之间的平衡。

综上所述，scavenger变量在实际的垃圾回收过程中具有重要的作用，能够帮助实现有效的内存回收和管理。






---

### Structs:

### scavengerState

scavengerState 是一个结构体，用于存储垃圾回收器（GC）的状态信息。在 Go 语言中，GC 由多个并发组成，Scavenger 是其中一个负责清除内存中垃圾的组件。scavengerState 在 Scavenger 组件的执行过程中记录了几个关键的状态：

1. scavenging：表示当前 Scavenger 是否正在执行清除过程中。当 scavenging 为 true 时，表示 Scavenger 正在运行，否则表示 Scavenger 没有在运行。

2. pMask、resultMask、penaltyMask：这三个 mask 变量用于存储 CPU 的掩码（CPU set mask），它们记录了当前 Scavenger 内部会用到的 CPU。

3. spanBytes、unscavengedBytes、scavengedBytes：这三个变量存储了 Scavenger 正在处理的 Span 的信息，如 spanBytes 表示当前 Scavenger 正在处理的 Span 总共包含的字节数，unscavengedBytes 和 scavengedBytes 则分别表示未清除和已清除的字节数。

4. scavTraversal：记录了已经遍历过的 Span 的数量，表示当前 Scavenger 已经遍历过多少个 Span。

这些状态信息可以帮助 Scavenger 组件更好地管理内存回收的过程。Scavenger 的执行过程中会不断更新这些状态，以保证垃圾回收的完整性和准确性。另外，这些状态信息也可以用于监控和调优 GC 的性能。



### scavengeIndex

scavengeIndex结构体是用于垃圾回收的索引结构体，它的作用是记录待回收的内存块的位置和大小信息，以便于在垃圾回收过程中快速地定位和回收这些内存块，从而提高垃圾回收的效率。

具体来说，scavengeIndex结构体包含三个字段：

1. addr：记录待回收内存块的起始地址。
2. size：记录待回收内存块的大小。
3. next：用于构建链表，记录下一个待回收内存块的位置。

在mgcscavenge.go文件中，scavengeIndex结构体是在scavengelist结构体中使用的。scavengelist结构体是一个链表结构，用于记录所有待回收的内存块，而scavengeIndex结构体则是链表结构的节点。

当运行垃圾回收时，会首先遍历所有的堆对象，将所有待回收的内存块添加到scavengelist链表中，并构建成一个索引结构。然后，在回收过程中，只需要遍历scavengelist链表，根据其中的scavengeIndex结构体的信息，定位到待回收的内存块，并将其回收。

使用scavengeIndex结构体可以大大提高垃圾回收的效率，减少扫描的次数，加快扫描的速度。



### atomicScavChunkData

atomicScavChunkData是一个结构体，用于管理垃圾回收器（GC）在进行垃圾回收时分配和回收的内存块（chunk）。该结构体使用了原子操作来保证多线程环境下对chunk的操作的同步。

具体来说，一个chunk可以被分配为几个不同的用途，如存储堆对象、存储栈对象、存储标记位节点等。atomicScavChunkData结构体的主要作用是保存和管理这些不同用途的chunk，以确保GC的高效性和正确性。

该结构体包含以下字段：

- scav：这是一个unsafe.Pointer类型，指向所管理的chunk的起始地址。
- stack：这是一个uintptr类型，表示chunk分配给栈使用的内存量。
- heap：这是一个uintptr类型，表示chunk分配给堆使用的内存量。
- markBits：这是一个uintptr类型，表示chunk中标记位节点的使用情况。
- abitmap：这是一个数组，表示chunk中的可用位图节点，以用于分配堆对象。

此外，atomicScavChunkData还包含了一个Lock字段，用于在多线程环境下保护chunk的操作，以确保线程安全性。

总之，atomicScavChunkData结构体在GC中充当了重要的角色，通过维护和管理内存块的使用情况，以确保GC的正确性和高效性。



### scavChunkData

scavChunkData结构体是用于描述分配给scavenger线程的内存块信息的数据结构。

它包含了以下字段：

- start：指向内存块起始地址的指针；
- end：指向内存块结束地址的指针；
- bitmap：一个指向位图的指针，用于表示内存块中每个字节的使用情况；
- pageCount：内存块中包含的页面数；
- span：一个指向mspSpan结构体的指针，表示该内存块所属的span。

scavChunkData结构体的作用是为scavenger线程提供内存块的详细信息，帮助scavenger线程有效地遍历和回收内存块中的垃圾对象，从而减少内存占用和提高程序性能。在mgcscavenge.go文件中，scavChunkData结构体会被用于构建scvg的chunk管理器，并通过maxPagesPerChunk和maxObjsPerChunk这两个参数配置chunk的最大页面数和最大对象数。



### scavChunkFlags

scavChunkFlags是Golang中用于垃圾回收的标志位的结构体。其作用是将一些标志位打包到一个uint32类型的字段中。

在Golang的垃圾回收过程中，需要对不同的内存块进行标记、扫描和移动等操作，而这些操作需要针对不同的情况设置不同的标志位。scavChunkFlags结构体中的字段包括以下几个：

- hasPointers: 指示内存块是否包含指针。如果包含指针，则需要进行指针的扫描和移动。
- isFree: 指示内存块是否已经释放。如果已经释放，则无需再进行任何操作。
- hasSwept: 指示内存块是否已经进行过标记和扫描。如果已经进行过，则无需重复操作。
- scavenged: 指示内存块是否已经进行过移动。如果已经移动过，则无需重复移动。
- needzero: 指示内存块是否需要进行清空操作。需要清空的情况包括内存块刚分配时、内存块从其他线程的缓存中获取时、以及内存块在调用者释放时。

这些标志位在垃圾回收的不同阶段中起到不同的作用，可以提高垃圾回收的效率和减少开销。例如，在标记阶段，只需要扫描具有指针的内存块，而不需要扫描不包含指针的内存块，可以减少扫描时间和消耗。在移动阶段，只需要移动未移动过的内存块，可以减少移动时间和消耗。



### piController

在 Go 语言的垃圾回收机制中，GC 时间和 CPU 利用率之间有一个平衡点，这个平衡点的计算和调整就是由 piController 结构体完成的。

piController 结构体中有下列几个字段：

- pgc：指向当前的 GC 变量；
- pgobuf：指向上一次的 GC 变量，以备回滚使用；
- lastTick：上一次 GC 操作的时间点；
- tgClockDiv：每次 GC 操作增加的时间，单位是 nanoseconds；
- tgInertia：调整时间的比率， 参数越大，调整结果就越缓慢；
- Kp, Ki, Kd：控制时间的 PID 参数；
- tgPhase：目标相位时间，单位为 nanoseconds，即 piController 的目标调节值。

piController 结构体的主要作用是通过对 GC 时间和 CPU 利用率的监测和分析，来计算出一个可接受的平衡点，并相应地调整 GC 时间以达到这个平衡点。这样可以在尽可能短的时间内完成垃圾回收，并尽可能地避免对正常程序执行的影响。



## Functions:

### heapRetained

在Go语言中，mgcscavenge.go文件中的heapRetained函数主要用于计算当前堆上仍被使用的内存大小，即已经分配但还没有被释放的内存大小。该函数在进行强制垃圾回收之前被调用，以确定是否有足够的空闲空间来进行垃圾回收。

具体来说，heapRetained函数首先会计算出堆上当前已经分配的内存大小（即堆的总大小），然后通过遍历堆上所有已经分配的对象，计算出仍然被使用的内存大小。最后，heapRetained函数将已分配但未释放的内存大小（即堆的总大小减去被使用的内存大小）作为返回值返回。

在进行强制垃圾回收之前，如果堆上未释放的内存大小超过了一定的阈值，就需要进行垃圾回收以回收未使用的内存，避免出现内存泄漏和应用程序崩溃等问题。因此，heapRetained函数的计算结果可以帮助判断是否需要进行垃圾回收，以及需要释放多少内存空间，从而保证应用程序的稳定性和性能。



### gcPaceScavenger

gcPaceScavenger是 Go 语言中 GC 的一个函数，位于 runtime/mgcscavenge.go 文件中，主要负责在垃圾回收过程中对垃圾回收器的行为进行调整，使其能够更好地符合应用程序的需求。

具体来说，gcPaceScavenger 的作用包括：

1. 根据应用程序的内存使用情况动态调整垃圾回收器的运行速度，确保垃圾回收能够尽可能地减少对应用程序的影响。

2. 控制垃圾回收的频率，避免频繁的垃圾回收对应用程序的性能造成负面影响。

3. 控制垃圾回收的深度，避免无限递归导致程序崩溃。

4. 调整垃圾回收器的并发度，避免某些操作会造成饥饿问题，影响程序的性能。

总之，gcPaceScavenger 的作用就是通过动态调整垃圾回收器的行为，使其能够更好地适应应用程序的内存使用情况，尽可能地减少对应用程序的影响，提高程序的性能。



### init

init函数是Go语言中一个特殊的函数，用于初始化程序或库的特定部分。mgcscavenge.go文件中的init函数的作用是，在垃圾收集器启动时初始化一些变量和参数，为垃圾回收做一些准备工作。具体来说，该init函数执行了以下操作：

1. 初始化gcController。gcController是垃圾收集器的控制器，它包含了垃圾收集器的状态、参数和配置信息等。init函数初始化了gcController的一些参数，如gcpercent（GC触发的比例阈值）、demonsWait（等待后台线程的时间）等。

2. 初始化sweep。sweep是垃圾收集器中的一个子任务，负责清除不再被使用的内存对象。init函数初始化了sweep的一些参数，如spanBytes、maxPagesPerSweep、pageNotInUse、pagesPerSweep等。这些参数影响了sweep任务的执行效率和效果。

3. 初始化spineLock。spineLock是一个互斥锁，用于保护span框架的修改。init函数初始化了这个锁，确保在多线程并发环境下，对span框架的操作是线程安全的。

总的来说，mgcscavenge.go文件中的init函数是垃圾收集器的初始化函数之一，它为垃圾回收做了一些准备工作，为后续的垃圾回收任务提供了必要的环境和参数。



### park

park是一个runtime函数，主要用于将goroutine挂起，即将其放入等待队列中等待信号，直到对应的信号到达其所在的线程，才会被唤醒继续执行。

在mgcscavenge.go文件中，park函数的主要作用是在执行gc过程中防止goroutine在堆栈上分配对象，通过将goroutine阻塞，将它从当前工作线程上移除，从而让gc程序可以遍历堆栈扫描对象。

park函数主要包含以下步骤：

1. 将当前goroutine的状态标记为Gwaiting，表示它等待信号的到来而被挂起。
2. 将当前goroutine放入等待队列中，等待信号的到来唤醒。
3. 调用goparkunlock函数将goroutine挂起并释放锁，等待外部信号的到来唤醒。
4. 当信号到达并唤醒goroutine时，park函数会返回，并将goroutine状态设置为Grunnable，表示它可以被重新执行。 

总的来说，park函数的作用是实现goroutine的调度和管理，帮助保证gc程序的正常运行。



### ready

在 Go 语言中，内存管理是自动的，而垃圾回收是 Go 语言中最为重要的特性之一。当程序运行时，内存空间被分配给变量和数据结构。在某些情况下，我们可能需要释放一些内存以使其可再利用。Go 编译器和运行时系统负责管理程序的内存，以便在不再使用的变量和数据结构的情况下回收内存空间。

在 Go 语言中，垃圾回收器负责回收不再使用的内存，以保持程序在运行期间的性能和可扩展性。在 mgcscavenge.go 中，ready 函数是垃圾回收器的一部分。它的作用是准备分配在工作时需要使用的框架。这个函数首先检查存活对象的列表，然后使用 scav 堆栈框架准备好 gcWork 对象。

在 ready 函数中，gcFlushWork 和 gcDrainWork 函数将已经完成的gcWork对象推送到全局队列中去，方便队列中其他的 worker 线程取走并执行任务。如果全局队列已经满了，该函数会等待其他的 worker 线程取走一些任务，然后将 gcWork 对象推送到队列中。这个过程维护了垃圾回收器的平稳运行，确保没有 goroutine 因为等待而阻塞了父进程。

总之，ready 函数是 Go 语言中垃圾回收器的一部分，它负责准备分配在工作时使用的框架。这个函数的主要作用是将已经完成的 gcWork 对象推送到全局队列中，方便队列中其他的 worker 线程取走并执行任务。这有助于维护程序的性能和可扩展性，并确保程序中没有 goroutine 因为等待而阻塞了父进程。



### wake

在Go语言中，GC（垃圾回收）是自动执行的，当系统中没有可用的内存空间时，GC就会在运行时清理不再需要的对象。通常情况下，GC的运行不需要程序员干涉，但是如果我们在程序运行的过程中需要手动触发GC，那么就需要使用`runtime.GC()`函数。

mgcscavenge.go文件中的`wake`函数是GC的一部分，它的作用是唤醒处于休眠状态下的P（处理器）。在GC运行期间，处理器P会被临时停止工作，以便GC可以扫描所有的内存并确定哪些对象可以被垃圾回收。处理器P被暂停时，其上的goroutines也会被暂停。

当GC结束时，处理器P需要被唤醒，以便继续处理后续的任务。这是`wake`函数的作用。具体来说，`wake`函数会从休眠的队列中选择一个P，并使其进入工作状态。此时，P将开始处理在其上获取的goroutine。

需要注意的是，在进入wait状态的时候需要先获取g0和m的锁，以保证线程在wait sleep的同时，g0不会退出，m也不会被销毁。在被唤醒后，需要再次释放锁，以保证程序的正常执行。另外，在程序运行时，wait等操作常常会阻塞线程，导致程序的性能下降，因此应该谨慎使用。



### sleep

在 Go 语言的运行时环境中，mgcscavenge.go 文件实现了垃圾回收器的扫描过程。其中的 sleep() 函数是一种睡眠机制，它用于等待当前线程完成一个任务后，暂时挂起该线程，以便其他线程可以尽快运行。

具体而言，sleep() 函数的作用是让线程进入阻塞状态，并设置一个最小休眠时间。该函数使用的是系统级别的阻塞操作，因此当前线程会被放入一个等待队列中，并被标记为“睡眠”状态，直到等待时间到达或等到一个被唤醒的信号量时才会重新激活。

在 mgcscavenge.go 文件中，sleep() 函数通常用于等待闲置的堆栈扫描器，以便这些扫描器可以重新运行并扫描内存。这样可以减少垃圾回收器的运行时间，并提高程序的性能。



### controllerFailed

controllerFailed是一个在垃圾回收期间用于控制GC控制器状态的函数。它的作用是将控制器状态从"running"状态转换为"waiting"状态，以便其他goroutine可以在该控制器执行期间执行其他任务。

在mgcscavenge.go文件中，controllerFailed函数是由controller函数调用的。当垃圾收集器发现它无法继续运行，例如当所有可用的P都在等待某些资源或其他原因导致GC暂停时，控制器函数就会调用controllerFailed函数。

controllerFailed函数执行以下操作：

1.设置控制器状态为"waiting"。

2.将控制器状态转换为"running"状态。

3.重新计算GC周期(如果需要)。

4.重置暂停时间戳以允许下一次GC暂停。

5.通知所有阻塞在控制器上的goroutine，以便它们可以继续执行。

6.返回控制器状态。

总之，controllerFailed函数通过控制器状态来管理GC控制器的行为，以确保GC操作可以正确流程化执行并且不会对其他goroutine的执行造成不利影响。



### run

在Go语言中，垃圾回收是非常重要的，因为它可以自动地回收不能访问的内存，从而避免内存泄露和程序崩溃。mgcscavenge.go文件中的run函数是Go语言的垃圾回收器（GC）中的一部分，主要用于执行扫描堆的操作。具体来说，该函数会从堆中的任何未标记对象开始，扫描有指针的对象对它们进行标记，并且将未被标记的对象放置到新的堆区域中，然后清空原有堆的所有对象，并将新的堆作为当前堆。这个过程是垃圾回收器的一次完整的扫描过程，也称之为“标记-清除”算法。

在run函数的执行过程中，会使用一个 workbuf 结构体来保存所有待处理的对象，并使用一个 Buffer 对象来作为临时存储区域。在GC的过程中，当一个对象被扫描并标记后，它的指针会被添加到workbuf中，以供后续扫描使用。同时，如果在扫描过程中出现了无法识别的指针，则会将该指针存储到缓冲区中，以待后续的扫描。

在run函数执行完成之后，将清空原有的堆，以便于后续的内存分配操作。同时，新分配的空间也会被赋值为零值，以避免UAR（未初始化的自动变量）的存在。在整个过程中，GC将会持续进行，并且会持续运行直到分配器无法再为堆分配内存为止。

总的来说，run函数是Go语言垃圾回收器的核心部分之一，它的目的是实现“标记-清除”算法，并使用workbuf和缓冲区来协助执行GC的过程。通过这个函数的执行，可以避免内存泄露和程序崩溃等问题，保证Go程序的稳定运行。



### bgscavenge

在Go语言中，mgcscavenge.go文件是垃圾回收器的实现文件之一，其中的bgscavenge函数是后台扫描函数。

具体来说，bgscavenge函数是采用Scavenger的方式进行并发扫描和回收不再使用的对象。它可以异步执行GC并充分利用空闲的CPU资源。

当执行bgscavenge函数时，它会扫描堆中的存活对象，并且将不再被使用的对象标记为“垃圾”，并加入到待回收队列中。同时，它还会监听在后台的gcTrigger信道中，当有GC请求时，它就会立即启动扫描和回收。

bgscavenge函数还有其他一些重要的作用，包括：

1. 管理扫描过程中对象的黑白图信息，方便后续引用关系分析和增量扫描。

2. 调用SweepPrune函数完成对标记为“垃圾”的对象的内存回收工作，释放闲置的内存空间。

总之，bgscavenge函数是Go语言中GC的核心实现之一，它可以无缝地与并发程序协同工作，有效地管理内存资源，提高程序性能。



### scavenge

在Go语言中，mgcscavenge.go文件中的scavenge函数主要负责执行垃圾回收过程中的“扫描”操作。该函数会遍历所有的对象，并尝试对那些被标记为垃圾的对象进行回收。

在具体实现中，scavenge函数会先遍历所有的根对象，然后递归扫描这些根对象引用的对象，并将这些对象标记为“活跃”的。接着，函数会对那些未被标记为“活跃”的对象进行回收。

与常见的“标记-清除”垃圾回收算法不同，Go语言的垃圾回收器采用的是“标记-扫描”算法。因此，在回收过程中，scavenge函数会扫描整个内存堆，找到所有存活的对象，并对未存活的对象进行回收。

总的来说，scavenge函数在Go语言的垃圾回收算法中扮演着非常重要的角色，它通过扫描所有的对象，释放未被使用的内存，从而实现了内存自动管理和垃圾回收的功能。



### printScavTrace

printScavTrace函数的主要作用是打印垃圾收集器线程的调试信息。

在进行垃圾回收过程中，垃圾收集器需要扫描整个堆，以确定哪些对象是可以释放的。为了完成这个过程，垃圾收集器使用多个线程同时扫描堆，并打印这些线程正在扫描的对象的信息。printScavTrace函数就是在这个过程中被调用来打印线程的调试信息。

具体来说，printScavTrace函数会记录每一个扫描线程的ID，并输出正在被扫描的内存地址和对象的类型等信息。这些信息对于定位垃圾收集过程中出现的问题非常有用。

值得注意的是，printScavTrace函数只用于调试垃圾收集器，而不应该用于正式生产环境中。如果在生产环境中开启了printScavTrace函数，可能会带来额外的开销和性能损失。



### scavengeOne

scavengeOne函数是Golang运行时垃圾收集器的一部分，它用于标记和回收内存中的垃圾对象。 

具体而言，scavengeOne函数的作用是扫描当前的并发标记根集，收集存活的垃圾对象并标记其header为非垃圾状态，同时将未被标记的对象回收。该函数是一个循环函数，每次迭代都会检查扫描过的对象是否有指向未标记的对象的指针，如果有则需要对其进行标记并把该对象放入待扫描队列中。 

此外，scavengeOne函数还调用了markAndEnqueue函数来标记根和待扫描对象，以及evacuate函数来将可回收的对象移动到新的内存页中，以便将其空间回收并汇总。 

总的来说，scavengeOne函数是垃圾回收器的核心组件之一，它通过标记和回收内存中的垃圾对象来确保Golang程序的稳定性和性能。



### fillAligned

fillAligned函数的作用是在扫描和清扫内存时，将从当前指针位置开始的一段内存填充为0。该函数被调用的场景包括：

1. 在scavengeWork函数中，当从当前span中分配对象时，需要将新分配的内存填充为0，以防止出现空指针等问题。

2. 在sweepFast函数中，当遇到CacheLineSize的整数倍地址时，需要将前面的内存填充为0，以确保下一个对象的头指针正确地对齐。

fillAligned函数首先计算出当前指针地址的偏移量，然后使用memclrNoHeapPointers将从当前地址开始的一段内存清零。memclrNoHeapPointers是个基于汇编实现的内置函数，用于清空内存中不包含堆指针的部分，以提高清空速度。

由于内存填充是经常使用的操作，使用fillAligned可以优化垃圾回收的性能。



### findScavengeCandidate

在Go语言中，垃圾回收是自动管理的。其工作原理是通过标记未使用的内存块，然后将其清除并回收。在垃圾回收过程中，Go语言的运行时系统使用多个内存管理器来管理内存分配和回收。

mgcscavenge.go文件中具有findScavengeCandidate这个函数。这个函数的主要作用是在垃圾回收阶段中，找到能够被垃圾回收器扫描的内存块。这些内存块可能包含可回收的对象或它们的部分。

具体工作流程如下：

1. 遍历所有内存分配器，并返回一个包含所有可扫描对象地址的切片。

2. 根据每个地址指针的大小，按照从小到大的顺序进行排序。由于较小的地址指针在内存增长过程中往往会被覆盖，因此优先处理较小的地址指针。

3. 对于每个排序后地址指针，判断其是否在需要扫描的区域内。如果是，则将其添加到扫描队列中，以便垃圾回收器扫描。

总结一下，findScavengeCandidate函数的主要任务是找到需要扫描的内存块，以便垃圾回收器对其进行扫描和回收。这个过程是Go语言自动管理垃圾回收的关键步骤之一。



### init

init函数是Go编程语言中的一个特殊函数，它在包的导入时自动执行，无需显式调用。init函数可以用于执行一些初始化操作，例如初始化全局变量或注册函数等。在mgcscavenge.go文件中，init函数也是用于执行一些初始化操作。

具体来说，mgcscavenge.go文件中的init函数是用于初始化垃圾回收器(scavenger)的相关参数。垃圾回收器是Go语言运行时的一个重要组件，它可以自动管理内存，清理不再使用的对象。

在init函数中，首先会初始化一些全局变量，例如mheap、work和gcController等。这些变量是垃圾回收器的关键组成部分，用于存储堆、工作列表和垃圾回收控制器的信息。

接下来，还会初始化一些与垃圾回收相关的参数，例如scavengePercent、scavengingGOGC和scavengeRetries等。这些参数用于控制垃圾回收器的行为，例如何时启动扫描、何时停止扫描以及何时回收内存等。

总之，mgcscavenge.go文件中的init函数的作用可以归纳为两点：一是初始化垃圾回收器的相关组件和参数，二是为垃圾回收器的启动、运行和停止等行为提供支持。



### grow

在 Go 的运行时系统中，mgcscavenge.go 是垃圾回收的主要代码文件之一。其中的 grow() 函数的作用是在进行垃圾回收时扩大堆空间。

垃圾回收是 Go 语言中自动进行的过程，它会从不再使用的对象中回收内存空间并将其投入到可用内存池中以供以后使用。运行时系统使用一个堆来管理这些对象，但如果堆空间不足，则垃圾回收可能无法进行，会导致 Go 程序崩溃。因此，当堆空间不足时，grow() 函数会被调用来扩大堆空间。

具体来说，grow() 函数会进行以下操作：

1. 计算扩展后的堆空间的大小

2. 从操作系统中分配新的内存空间

3. 将原来的堆空间中的对象移动到新的内存空间中，并更新相关的指针

4. 将新的内存空间与原来的堆空间合并，形成一个更大的堆

这个过程涉及到大量的内存操作，因此需要特殊的处理。在实现上，grow() 函数会使用一些技术来最大限度地提高效率。例如，它会对新的内存页进行预分配，避免了频繁的系统调用。它还会使用并发技术来充分利用多核 CPU 的计算能力，从而加快扩大堆空间的速度。

总之，grow() 函数是 Go 语言运行时系统中非常重要的组成部分之一，为垃圾回收提供了必要的支持，帮助程序在更大的堆空间中顺利执行。



### find

find函数是mgcscavenge.go中的一部分，用于在堆中查找可以回收的对象。

在Go语言的垃圾回收机制中，主要分为两种回收方式：标记清除和标记整理。在标记整理算法中，需要找到并移动对象。find函数就是用于查找可以被回收的对象。

find函数的具体实现如下：

```go
func (s *mspan) find(size, npages uintptr) gclinkptr {
  	for q := s.freeindex; q != ~uintptr(0); q = atomic.Loaduintptr(&s.arena[q].next) {
  		if q+npages > s.nelems {
  			break
  		}
  		// we know that size is > 0 so we don't need to check if size - 1 is negative
  		i := int((size - 1) / s.elemsize) // 当前size所在的堆块索引位置（整数部分）
  		for ; i < int(npages); i++ {
  			if s.arena[q+uintptr(i)].isFree() {
  				goto Found
  			}
  		}
  	}
  	return 0
  Found:
  	atomic.Storeuintptr(&s.freeindex, q+uintptr(i))
  	sizep := uintptr(i)*s.elemsize + s.base()
  	if uintptr(i) == npages-1 {
  		return gclinkptr(sizep)
  	}
  	p := gclinkptr(sizep)
  	for i := uintptr(i) + 1; i < npages; i++ {
  		q := sizep + i*s.elemsize
  		*(*gclinkptr)(unsafe.Pointer(q)) = p
  		p = gclinkptr(q)
  	}
  	return p
  }
```

具体来说，find函数传入两个参数：size和npages，它们分别表示需要回收的对象的大小和页数。然后，函数会从当前mspan的freeindex处开始向后扫描堆（Arena），查找可以用于存储待回收对象的空闲区域。

在找到合适的空闲区域后，find函数会更新freeindex的值，并将空闲区域分割成与待回收对象大小相同的多个小块。然后，它会返回一个指向第一个小块的指针，并将其余的小块链接到一起。

总的来说，find函数是Go语言中垃圾回收器的重要组成部分。它的作用是在堆中查找可以回收的对象，并将其分割成更小的块，以便更高效地进行垃圾回收。



### alloc

The `alloc` function in `mgcscavenge.go` file in the `go/src/runtime` package is responsible for allocating memory during the scavenging phase of the Go garbage collector.

During this phase, the garbage collector identifies memory regions that are no longer in use and frees them up for new allocations. The `alloc` function is called when new memory is needed to replace the freed regions.

The function takes as input the size of the memory block to allocate and returns a pointer to the start of the allocated block. It first tries to allocate memory from the available span of memory. If there is no available span, it triggers the garbage collector to run another cycle to free up more memory.

The `alloc` function also handles various edge cases such as providing an overcommit buffer for the operating system and handling alignment of memory blocks.

Overall, the `alloc` function is a critical component of the scavenging phase of the garbage collector and ensures efficient and safe allocation of memory.



### free

在Go语言中，内存管理器（Memory Manager）是由运行时系统（Runtime System）负责的。在运行时系统的内存管理机制中，垃圾收集器（Garbage Collector）是用来管理未被使用的对象的一种机制。

在runtime/mgcscavenge.go文件中，free()函数用于释放分配的内存区域。在Go语言中，当一个对象不再被使用时，垃圾收集器会将其标记为垃圾，并将其存放在一段指定的内存区域中。当这个垃圾收集器被触发时，它会进行扫描，找出这些垃圾对象，并将其释放。在这个过程中，就需要使用到free()函数。

free()函数的主要作用是将一个内存区域标记为空闲，并将其添加到堆的空闲列表里，待下次需要分配内存时，就可以从这个空闲列表里获取。在Go语言中，垃圾收集器是通过分代的方式进行工作的，可以根据对象的使用情况将其分为不同的代，而将分配的内存区域标记为空闲，也是为了便于下次分配内存时选择合适的代进行分配。

总结来说，free()函数是垃圾收集器机制中的一个关键步骤，它能够释放不再使用的内存，并将其标记为空闲，方便下次分配内存时进行使用。



### nextGen

在Go语言中，垃圾收集是指自动回收程序期间，Go运行时环境会跟踪和释放不再使用的内存。mgcscavenge.go是Go语言中垃圾回收器的源代码文件之一，其中包含了nextGen这个函数。

nextGen函数的主要作用是确定下一个需要进行垃圾回收的代数，并触发相应的垃圾回收器来完成回收操作。垃圾回收代表了内存中某个区域中的对象的集合，不同代之间对象的访问频率和生命周期不同，从而可以有效地控制内存的使用情况。

nextGen首先获取当前进行垃圾回收的代数，然后根据制定的回收策略和代数之间的映射来确定下一个需要进行垃圾回收的代数。通常情况下，垃圾回收器会根据根对象和存于堆上的对象的引用关系来选择下一个需要回收的代数。例如，标记-清除回收器就是一种常见的垃圾回收策略，它会在所有不可达的对象被标记并清除之后，使得所有未标记的对象进入下一代。

整个过程涉及到了大量的数据结构和算法知识，包括对象分配与回收、垃圾回收算法、堆管理、扫描和清除等内容。因此，nextGen函数的实现非常复杂，需要综合考虑多个因素，以控制垃圾回收的速度和效率。



### setEmpty

setEmpty函数在垃圾回收期间用于清空垃圾扫描缓存集合。在进行垃圾回收时，扫描工作必须在已经使用的集合对象中执行。为了加快扫描工作的速度，程序会缓存已经扫描过的对象集合。但是，当垃圾回收完成后，这些缓存集合需要被清空，以便在下次运行时重新填充新的已使用对象集合。

setEmpty函数的主要作用是在垃圾回收完成后，将垃圾扫描缓存集合中存储的对象清空，以便下次进行垃圾回收时，重新填充新的已使用对象集合。通过调用此函数，可以确保缓存集合不会在下一次回收过程中造成垃圾扫描错误或其他不良影响。因此，该函数是一个非常重要的功能，必须仔细地进行处理和管理。



### setNoHugePage

setNoHugePage是runtime中mgcscavenge.go文件中的一个函数，其作用是设置是否允许使用hugepage。 

在Linux系统中，hugepage是一种机制，可以将操作系统的虚拟内存页大小从4KB(默认)增加到大约2MB，从而提高内存的处理效率。然而，使用hugepage时，需要为其保留一定的物理内存空间，因此需要在启动时进行设置。

在setNoHugePage函数中，它通过将noHugePage设置为true或false来控制hugepage的使用。当noHugePage为true时，表示禁止使用hugepage，而当noHugePage为false时，表示可以使用hugepage(默认情况下)。在mgcscavenge.go文件的scavenger函数中，在开始标记阶段调用了setNoHugePage(true)，即禁用了hugepage的使用，以降低垃圾回收期间内存碎片的产生。而在结束标记阶段和标记清除阶段中又调用setNoHugePage(false)，表示恢复hugepage的使用，提高内存的处理效率。

总之，setNoHugePage函数的作用是控制是否允许使用hugepage，从而影响内存处理效率和内存碎片的产生。



### load

mgcscavenge.go文件中的load函数的作用是从指定的unsafe.Pointer指向的地址读取一个指向uintptr的指针并返回这个uintptr指针的值。在Go语言中，uintptr是一个与指针相似的类型，但它是一个无类型的整数值，可以作为指针运算的操作数。load函数在内存回收过程中被用来从堆栈和堆中读取指针地址，这些指针地址指向的对象需要被回收。 

具体来说，load函数的实现使用了unsafe包提供的指针操作，它首先将unsafe.Pointer类型的参数地址转化为*uintptr类型的指针，然后使用*uintptr类型的指针读取指向uintptr类型的指针指向的值，并返回这个uintptr类型的值。整个过程涉及到指针的转化和指针类型的读取操作，因此需要使用unsafe包，而不是在普通的程序中使用。在内存回收过程中，由于需要访问已经被释放或者已经迁移的对象，因此需要使用unsafe包和load函数。这个函数在Go语言的运行时中发挥着非常重要的作用，但在普通的程序中应该尽量避免使用unsafe包和指针操作，以避免潜在的风险和错误。



### store

在go/src/runtime中mgcscavenge.go这个文件中，store这个func是用于将指定的值存储到给定的内存地址中的函数。该函数是用于gc（垃圾回收）过程中的扫描操作，将堆栈、寄存器以及其他一些内存区域中引用的对象标记为“已使用”。

在gc过程中，由于还存在一些不可达对象，这些对象被标记为“已使用”会导致内存泄漏等问题。因此，store函数还会执行一些其他操作，如更新内存分配器中的信息、调用页错误处理程序等。通过这些操作，store函数能够确保已使用的对象得到正确的回收和清除，从而保证了内存管理的正确性和可靠性。

总之，store函数是一个重要的内存管理函数，它包括了多个关键的操作，可以确保gc过程中使用的内存得到正确的处理和回收。



### unpackScavChunkData

unpackScavChunkData是gcScavenge模块中用来解压缩内存碎片的函数。在Go中，内存碎片是指没有归属对象的一小块内存空间，这些碎片可能会占用大量的内存，并影响系统的稳定性和性能。因此，Go运行时系统使用gcScavenge模块来收集和压缩这些内存碎片，以便系统可以更有效地使用内存。

unpackScavChunkData函数的主要作用是将压缩的内存碎片数据进行解压缩和重组，使其可以被垃圾回收器检测和清理。具体来说，该函数接收一个指向内存碎片数据的指针和一个作为输出缓冲区的指针。它首先解压缩内存碎片数据，然后根据内存对象的大小将其拆分成若干个内存块，并将它们放入输出缓冲区中。这些内存块可以被垃圾回收器识别和清理。

对于Go运行时系统来说，unpackScavChunkData函数的作用非常重要。通过使用该函数，系统可以更有效地回收内存碎片，减少内存消耗和提高系统性能。



### pack

在Go语言中，gcscavenge.go文件中的pack函数的作用是将堆中的对象紧密地打包，以产生最少的碎片。在垃圾收集时，个体采用打包算法使堆更好地压缩，并且具有更少的内存碎片和整体垃圾收集时间。

pack函数主要分为两部分：对象排队和对象压缩。在对象排队的过程中，函数将堆中的所有对象按照相对大小进行递增排序，并将它们排队。在对象压缩的过程中，函数会将递增排序后的对象按照最优策略进行压缩，以保证最少的碎片。

pack函数中还会执行一些其他操作，包括重置堆对象的存储标记、准备下一轮整体垃圾回收、释放空闲的堆空间等操作，以确保GC的正常执行。

总之，pack函数的作用是在垃圾收集器开始回收之前，对堆中的对象进行排队和压缩，以减少内存碎片并提高垃圾回收的效率。



### isEmpty

在go/src/runtime中mgcscavenge.go文件中的isEmpty函数用于判断当前span是否为空。span是表示内存分配的连续区域，每个span都有一定大小，例如8KB。

isEmpty函数的定义如下：

```go
func (s *mspan) isEmpty() bool {
    for i := uint32(0); i < s.nelems; i++ {
        if s.freeindex(i) != 0 {
            return false
        }
    }
    return true
}
```

该函数遍历span中的空闲位置，如果存在不为空的空闲位置，则返回false表示此span非空；如果所有空闲位置都为空，则返回true表示此span为空。

当GC回收时，如果找到的span属于空span，则将其加入到free或central空闲链表中，以便将来分配对象时能够利用这些内存空间。doubleValue函数就是在GC时调用isEmpty函数判断span是否为空。

isEmpty函数在GC时非常重要，判断span是否为空可以方便地将内存回收和再利用。它是Go语言垃圾回收机制的一部分，是提高Go语言内存管理效率的关键部分。



### setEmpty

setEmpty函数是在垃圾回收过程中用来清空一个对象的指针数组的函数。在Go语言中，垃圾回收主要是通过标记-清除和标记-整理两种方式来进行的。其中标记-清除会将标记为垃圾的对象直接释放掉，而标记-整理会将标记为垃圾的对象移动到一定的区域，在原区域留下一片连续的空间用于存放新的对象。对于标记-整理的方式来说，移动对象指针是非常困难的，因此需要在拷贝完对象之后，将原来的指针数组设置为空。

因为对象的指针数组可能非常大，它需要保存所有指向这个对象的指针。这些指针可能存在于其他对象中，也可能在栈、寄存器、全局变量或者其他地方。因此，在拷贝对象之后，需要将原来的指针数组标记为空，以保证下一次垃圾回收不会将这个已经处理过的对象再次拷贝一次。这样，就可以防止对象被拷贝多次，提高了垃圾回收的效率。

在setEmpty函数中，首先会将指针数组清空，然后再将这个对象的标志位设置为标记为空。这样，在下一次垃圾回收的时候，就可以跳过这个对象，不需要再次拷贝一次了。



### setNonEmpty

setNonEmpty是在垃圾收集器扫描一个对象时，标记该对象为非空的函数。它的作用是将对象的标记位设置为"已分配"，告诉垃圾收集器该对象已经在使用中，不可被回收。

在垃圾回收的过程中，垃圾收集器需要扫描堆中的所有对象。扫描时，它会检查每个对象的标记位，如果标记为非空，则说明该对象正在被使用。如果标记为空，则说明该对象不再被引用，可以被回收。

因此，setNonEmpty的作用就是将对象的标记位设置为非空的状态，避免被错误地回收，保证程序正常运行。具体实现时，该函数会调用 bitmapMark1.set 和 stack.addToMarkRoots 方法，将对象的标记位设置为"已分配"，同时添加到标记根中。

需要注意的是，setNonEmpty函数只在垃圾收集器扫描对象时调用，不会在程序中被其他函数直接调用。它是一个内部函数，仅由其他垃圾收集器相关的函数调用，用于保证垃圾回收的正确性。



### isHugePage

在go/src/runtime中的mgcscavenge.go文件中，isHugePage是一个函数，它的作用是判断所给定的地址是否是huge page，即是否是大页。

在Linux系统中，如果一个进程使用大量内存，那么会使用huge page机制来提高内存分配和使用的效率。通常情况下，操作系统将物理内存划分为一页一页的小块，称为普通页面，它们大小通常为4KB或8KB，但huge page的大小可以达到2MB或更大。使用大页的优点是可以减少虚拟地址空间的开销，提高内存分配和释放的效率。

在Golang中，如果一个对象的大小超过了某个阈值，它会被分配到Huge Page上。isHugePage函数就是用来判断一个给定地址是否是huge page的，它的实现方式是读取/proc/self/maps文件，寻找包含给定地址的内存映射，并判断该内存映射是否属于huge page。

在GC过程中，如果需要扫描某个内存块（比如根据goroutine的栈，扫描对应的堆栈信息），那么如果该内存块所在的页是huge page，那么就需要特别处理，以保证GC的正确性。所以在这种情况下，isHugePage函数就可以发挥它的作用。



### setHugePage

func setHugePage(addr unsafe.Pointer, n uintptr) {
	// 尝试将虚拟内存页转换为 huge page
	hp := uintptr(addr) >> hugePageShift
	endHp := (uintptr(addr) + n - 1) >> hugePageShift
	for i := hp; i <= endHp; i++ {
		hpAddr := unsafe.Pointer(i << hugePageShift)
		err := madvise(hpAddr, hugePageSize, _MADV_HUGEPAGE)
		if err == 0 {
			memstats.heap_huge += uint64(hugePageSize)
		}
	}
}

setHugePage函数的作用是将分配的内存页转换为巨大页面，以提高程序性能。在大内存（huge page）模式下，操作系统会以更大的内存页面来分配内存，这可以更好地利用硬件内存管理单元，并减少操作系统在处理大量页面时的开销。

setHugePage函数首先将传入的地址转换为指向最接近的巨大页面的地址。然后，它遍历分配的内存范围，并尝试将每个虚拟内存页面转换为巨大页面。如果转换成功，就会将memstats.heap_huge（表示已分配的巨大内存量）增加对应的数量。

需要注意的是，使用大内存页面的时候，程序会使用更多的物理内存，因为每个页面的大小会更大，所以在使用这种技术之前需要考虑系统的实际内存情况。



### setNoHugePage

setNoHugePage函数的作用是设置不使用巨页。
巨页是一种增加虚拟内存管理效率的技术。当应用程序请求大块内存时，操作系统可以分配巨页来满足这些请求。这些巨页比常规的小页更大，可以减少内存管理的开销和虚拟内存中页表的大小。
在Golang的垃圾回收算法中，如果应用程序使用了巨页，可能会导致分配内存时的性能瓶颈。因此，在mgcscavenge.go中，setNoHugePage函数可以禁止在垃圾回收扫描期间使用巨页。这能够提高垃圾回收的性能和效率。
具体来说，setNoHugePage函数会使用madvise系统调用，设置内存映射是否可以使用巨页。如果可以使用巨页，就会开启。如果不行，就禁用。 在垃圾回收扫描期间，可以使用这个函数禁用巨页，以避免影响性能。这个函数只在Linux平台上有用。



### shouldScavenge

shouldScavenge函数是Go语言中垃圾回收器（Garbage Collector）的一部分，其主要目的是确定是否需要进行新的扫描。

垃圾回收器是自动的，它工作在后台，定期扫描程序的内存，以查找不再使用的对象并释放它们。为了避免在程序执行期间影响性能，垃圾回收器只在必要的时候才会扫描内存。shouldScavenge函数就是用来判断当前是否需要进行扫描。

shouldScavenge的实现会根据当前的内存使用情况、时间等因素来决定是否需要执行垃圾回收。如果应该执行垃圾回收，shouldScavenge函数会返回true，否则返回false。如果shouldScavenge返回true，则会执行一次扫描以释放不再使用的对象。如果返回false，则不会进行扫描，从而避免不必要的性能影响。

总之，shouldScavenge函数是Go语言垃圾回收器的一部分，其作用是根据当前的内存使用状况和时间等因素来决定是否需要执行垃圾回收。



### alloc

alloc函数的作用是从MCache中分配一块内存块。MCache是Go语言中的一种缓存机制，用于缓存小型对象（小于32KB）的分配和归还，以避免频繁使用堆内存，提高性能。

在mgcscavenge.go文件中的scavengem方法中，当需要分配内存块时，会调用alloc函数从MCache中分配一块内存块。如果MCache中没有足够的空闲内存块可供分配，alloc函数会尝试从central缓存中获取内存块。如果central缓存中也没有足够的内存块可用，则会调用sysAlloc函数从操作系统中获取新的堆内存，并将其中一部分用于缓存中，以备下次使用。

alloc函数的实现方式非常简单，首先会将内存块尺寸对齐到8字节的倍数，然后判断MCache中是否有足够的内存块可用，如果有则返回其中一块内存块，如果没有则尝试从central缓存中获取内存块，如果获取成功则返回，否则调用sysAlloc函数分配新的内存块，将其中一部分放入central缓存中，并返回剩余的内存块。



### free

在Go语言中，GC（Garbage Collection，垃圾回收）是自动进行的，即无需手动管理内存。而这个free函数是专门用于垃圾回收的，其作用是释放内存。具体来说，它会将不再被使用的内存块标记为可回收，并将其添加到空闲列表中，以备下次分配内存时使用。

当Go应用程序运行时，内存可能会被分配和释放多次。为了避免出现内存泄漏（即已分配但不再被使用的内存），GC会定期扫描内存空间，找出不再被使用的内存并释放它们。在进行垃圾回收时，GC会遍历整个内存空间，找出标记为可回收的内存，并使用free函数释放这些内存。这可以减少内存的使用量，从而提高应用程序的性能和稳定性。

总之，free函数是GC中非常重要的一部分，它能够帮助我们自动管理内存空间，防止内存泄漏，提高应用程序的性能和稳定性。



### next

next函数是垃圾回收中scavenger的一部分。Scavenger是Go语言中垃圾回收器的一种类型，它是一种并发活动，它在程序运行期间不停地扫描内存并回收不再被引用的对象。

next函数的作用是返回下一个对象，它在scavenger并发扫描内存时用于遍历对象链表。它接受一个*sizeclass参数，并根据其返回一个未扫描的对象的地址。sizeclass参数是指对象所占用的字节数，这对于回收不同大小的对象非常重要。

next函数使用一个全局变量scav.next，该变量指向当前要扫描的对象。在每次调用next函数时，scav.next指针就会移动到下一个未扫描的对象。这个函数还会对指针进行边界检查，以防止它扫描超出堆的范围。

总的来说，next函数是Go语言垃圾回收器中的一个重要组成部分，它确保了内存中未被引用的对象能够被及时回收，并且实现了高效的并发扫描。



### reset

reset函数是在垃圾回收过程中用于重置扫描和标记状态的函数。它的主要作用是在垃圾回收之前，将标记为已扫描的对象和标记为已标记的对象都重新设置为未扫描和未标记状态。

具体而言，reset函数会遍历整个堆，将所有已经标记的对象都重新标记为未标记状态，同时清空所有协程的ScanState信息，以便在下一次扫描时可以从头开始扫描。此外，reset函数还会将存档(archived)对象的标记设置为未扫描状态，这些对象在前一次垃圾回收之后被存储在存档中，因此需要在下一次回收时重新扫描和标记。

总之，reset函数的作用是在垃圾回收过程中重新设置扫描和标记状态，以便下一次垃圾回收可以从头开始，不会受到上一次垃圾回收的影响。



