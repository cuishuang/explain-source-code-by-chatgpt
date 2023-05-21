# File: lockrank_off.go

lockrank_off.go这个文件是Go语言的运行时（runtime）库中的一个文件，它的作用是禁用Goroutine锁竞争排序（lock rank）功能。

在Go语言中，作为一种并发编程的语言，Goroutine与锁是两个重要的概念。当多个Goroutine同时对同一个资源进行访问时，容易出现竞争条件（race condition）问题，需要使用锁来保护共享资源。在锁的实现中，如果锁的等待队列上存在多个等待线程，它们会按照先来先服务（FIFO）的顺序进行排队等待锁的释放。但实际上，在某些场景下，这种等待顺序并不是最优的，会导致锁竞争的效率降低。

针对这个问题，Go语言的运行时库提供了一种称为锁竞争排序（lock rank）的机制，它可以根据不同的场景，对等待锁的线程进行优先级排序，使得高优先级的线程可以更快地获取到锁，从而提高多线程的竞争效率。但是，锁竞争排序的机制也会带来一些额外的开销，因此有时候需要禁用这个功能，以便更好地适应特定的使用场景和优化性能。

lockrank_off.go这个文件就是用来禁用锁竞争排序功能的。它包含了一个调用了runtime lockRankOff函数的go:init函数，当程序启动时，会执行这个init函数，以便禁用锁竞争排序功能。这样，通过将lockrank_off.go文件链接到程序中，就可以禁用锁竞争排序功能，以适应个别的高性能场景。




---

### Structs:

### lockRankStruct

lockRankStruct是用于跟踪运行时系统中锁的等级以及锁的升级和降级的数据结构。具体来说，它包含了当前锁的等级、上一次锁的等级、以及需要升级或降级所需执行的函数。

在Go语言中，为了避免死锁的发生，锁是按照一定的顺序获取的。lockRankStruct就是用来维护这个获取锁的顺序的。当一个goroutine尝试获取一个锁时，它会获取锁的等级，并与其他goroutine持有的锁进行比较。如果其他goroutine持有的锁的等级小于当前锁的等级，则当前goroutine需要将自己的等级升级。如果其他goroutine持有的锁的等级大于当前锁的等级，则当前goroutine需要将自己的等级降级。

lockRankStruct中的升级和降级函数会被调用来执行这些操作。这些函数会更改锁的等级，并且会通知其他goroutine需要重新进行锁的升级或降级。通过这种方式，锁可以被有效地获取，并且可以避免死锁的发生。

总的来说，lockRankStruct是Go语言中用于维护锁的顺序以及锁的升级和降级的重要数据结构。它为Go语言提供了一种有效的机制来防止死锁的发生。



## Functions:

### lockInit

lockrank_off.go文件是Go语言运行时包（runtime）中的一个文件，其中的lockInit()函数用于初始化一个互斥锁（mutex）。

互斥锁是一种用于控制多个线程对共享资源的读写操作的同步机制。当一个线程占用锁时，其他线程必须等待其释放后才能继续执行。因此，互斥锁是非常重要的同步机制，可以用于保证共享数据在多线程环境下的正确性。

在lockrank_off.go文件中，lockInit()函数使用一个特殊的底层锁（lock）来初始化一个互斥锁。这个底层锁与标准互斥锁相比，优化了一些性能问题，并可以对锁的使用情况进行跟踪和分析。

lockInit()函数的代码如下：

```go
func lockInit(l *mutex, rank lockRank) {
    if race.Enabled {
        race.InitializeShareAddr(unsafe.Pointer(l), mutexSize)
    }
    l.rank = rank
    if debuglock {
        getcallerpc(unsafe.Pointer(&l.initpc))
    }
    lockInitProfile(l)
}
```

lockInit()函数的参数是一个mutex指针（l）和一个表示锁级别的rank枚举值。锁级别用于在分析锁的使用情况时提供更多的细节。

在函数中，首先会检查race.Enabled，如果race（data race detector）功能已启用，则会在共享内存中初始化一个与锁相关的地址。然后函数将rank赋值给锁的rank属性。

如果debuglock（调试锁）功能开启，则获取调用函数的程序计数器（PC）并存储到锁的initpc属性中。这个属性用于追踪使用锁的代码。

最后，lockInitProfile()函数被调用来初始化与locks（锁）相关的性能调试信息。

总之，lockInit()函数的作用是对一个互斥锁进行初始化，并优化其性能表现，同时提供了更好的调试和分析功能。它是Go语言运行时中非常重要的函数之一。



### getLockRank

getLockRank是一个用于获取锁等级的函数。在go的运行时中，每个锁都会被分配一个等级，通过getLockRank函数可以获取指定锁的等级。

该函数的具体作用是用于调试和监控锁的使用情况。通过获取锁的等级，可以分析程序中锁的使用情况和瓶颈，找到可能存在的死锁和竞争条件。

在实现上，该函数的主要作用是获取指定锁的等级，并根据等级返回相应的常量值。如果锁的等级未知，函数将返回一个默认值，表示未定义的等级。在编写并发程序时，可以利用该函数来监测锁的性能和使用情况，从而优化程序的性能和可靠性。



### lockWithRank

lockWithRank函数是在没有lock ranking模式下的锁实现。在没有lock ranking模式下，锁和其它同步原语的操作是以权重等级为依据，确保在一定条件下的执行顺序。

lockWithRank函数的作用是实现一个传统的互斥锁，以确保在共享资源的确切使用期间将其保持在单个协程或线程中。该函数采用了自旋锁和信号量结合的方式，通过循环自旋不断尝试获取锁，直到成功获取为止。当锁被释放时，通过信号量通知等待队列中的协程或线程可以尝试获取锁。

该函数的实现涉及以下步骤：

1. 初始化mutex结构体对象，其中包括互斥锁类型、锁的状态（是否被持有）以及等待队列、利用设置的值初始化自旋字
段的状态等。

2. 利用自旋锁的方式尝试获取锁，如果锁已经被其他协程或线程持有，则当前协程或线程将会进入自旋等待状态，尝试重新获取锁。

3. 在获取锁之后，将mutex的state变量标记为已持有状态，并返回表示成功获取锁的状态。

4. 在释放锁之前，需要首先将mutex的state变量标记为未持有状态，然后立即释放锁，并将等待队列中的一个协程或线程添加到锁的持有者队列中。

由于lockWithRank函数的实现方式非常基础，包含了自旋锁、权重等级等基本的同步原语，因此其运行效率相对较低，一般建议采用更高级别的锁机制。



### acquireLockRank

lockrank_off.go文件定义了一个在lockrank打开时会被调用的acquireLockRank函数。它的主要作用是获取当前goroutine的锁排名并将其保存在goroutine的上下文中，以便在goroutine陷入死锁时帮助排查问题。

具体来说，acquireLockRank函数会调用runtime.callers函数获取当前调用栈信息，并分析调用栈中获取锁的顺序，最终得出该goroutine的锁排名。锁排名是一个0~4294967295之间的数值，表示当前goroutine在竞争锁资源时的优先级，排名越小表示优先级越高。

获取锁排名后，acquireLockRank函数会将其保存在goroutine的上下文信息中，以便在死锁检测时使用。这样，在发生死锁时，可以通过查看每个goroutine的锁排名，确定哪些goroutine在获取锁时存在问题，从而更容易地排查和解决死锁问题。

总之，acquireLockRank函数是一个辅助函数，用于在lockrank打开时帮助排查死锁问题，提高程序稳定性和可靠性。



### unlockWithRank

unlockWithRank函数是在执行锁解除时，用来更新goroutine的锁排名（lock rank）的。锁排名是一个用来描述goroutine在锁中等待的相对顺序的值，它最初是从零开始的，每次当一个goroutine在获取锁时，它的排名将会被增加，并在锁被释放时减少。

unlockWithRank函数的作用是将当前的锁排名（lock rank）赋值给当前goroutine的lockwait字段，并将锁排名减1。这个操作是为了避免排名被“锁住”，并确保在下次获取锁时，该goroutine的排名不会太高，从而优先级过高导致其他goroutine无法获取锁。

总之，unlockWithRank函数的作用是保证锁的公平性，避免因为某个goroutine的高优先级导致其他goroutine无法获取锁。



### releaseLockRank

在go语言中，每个goroutine都会持有一个或多个锁，当一个goroutine持有一个锁时，其他goroutine就不能再获取这个锁了。锁的竞争是导致程序性能下降和死锁的主要原因之一。

为了帮助开发者诊断锁竞争问题，go语言提供了lock ranking功能。当开启lock ranking时，go会在线程的局部储存（TLS）中存储一个lock rank值，用于表示一个goroutine同一时间持有多少个锁。当一个goroutine抢占CPU时，go会检查当前的lock rank值是否合理，如果当前持有的锁的数量超过了预设的值，go就会输出警告信息并尝试自动打印堆栈信息。

在lockrank_off.go文件中，releaseLockRank（）函数的作用是关闭lock ranking。这个函数会将goroutine的lock rank值设置为一个非法值，使得当这个goroutine抢占CPU时，go不会输出任何警告信息。这个函数的作用是提高程序的性能，因为lock ranking功能会增加一些额外的开销，并且在一些场景下不是很有用，例如那些只持有一个锁的goroutine。

总而言之，releaseLockRank函数在lock ranking功能不必要时关闭这个功能，提高程序性能。



### lockWithRankMayAcquire

在Go语言中，lockWithRankMayAcquire这个函数是用于实现锁定机制的。它通过加锁和解锁来保护临界区代码的执行。

具体来说，lockWithRankMayAcquire函数的作用是尝试以给定的优先级获取锁，如果锁已经被持有，则等待直到可以获取锁为止。如果当前线程得到锁的优先级低于它想获取的优先级，则会将当前线程阻塞，直到其他线程释放锁或当前线程被唤醒。

此外，lockWithRankMayAcquire函数还可以对线程进行优先级排序，以确保优先级高的线程能够尽快获得锁，从而提高锁的性能和效率。

总之，lockWithRankMayAcquire函数在Go语言的并发编程中扮演着重要的角色，可以有效地保护共享资源，避免出现数据竞争和死锁等问题。



### assertLockHeld

在 Go 语言中，所有的锁都应该只能在持有锁的 goroutine 中被解锁。否则，会因为竞争条件而导致数据不一致或程序崩溃等问题。为了保证这一点，Go 在运行时提供了一些锁调试机制来检测每个锁的正确使用情况。assertLockHeld() 就是其中之一。

assertLockHeld() 函数的作用是检查当前 goroutine 是否持有某个锁。如果该锁不是在当前 goroutine 中被持有，assertLockHeld() 就会触发 panic，以便开发者及时发现潜在的错误。这个函数通常用于开发调试时，而不是正式运行时，因为它会损坏程序的性能。

assertLockHeld() 函数实现比较简单，它的核心就是调用 runtime 包中的 checkLockHeld() 函数来检查锁的持有情况。checkLockHeld() 默认只在 debug 编译模式下才会被编译进去，以避免对运行时的性能造成影响。



### assertRankHeld

在 Go 语言中，锁是一种用来控制并发访问的机制，以确保线程安全的数据访问。在实现锁时，必须要考虑到锁的权重和优先级，以避免死锁、饥饿等问题。lockrank_off.go 这个文件的作用就是定义了锁的优先级，同时提供了一些辅助函数，以方便在调试和优化时进行调用。

其中，assertRankHeld 函数的作用就是在检查锁的正确性时使用的。在实际的代码中，我们经常使用锁来控制线程的访问，如果在使用锁的过程中出现错误或者异常，就会导致线程出现死锁等问题。为了避免这种情况的发生，我们需要在程序运行期间尽可能地检查锁的状态和正确性。assertRankHeld 函数就扮演了这样的角色，它用来检查当前线程是否持有指定优先级的锁。如果当前线程没有持有指定优先级的锁，就会触发一个异常，以提示程序员有关问题的信息。

assertRankHeld 函数的实现非常简单，它只是通过调用 runtime 包中的 gLockRank 函数来获取当前线程持有的锁的优先级，然后将此值与函数参数中传递的优先级进行比较，以检查当前线程是否已持有指定的锁。如果不持有，则触发一个异常。这个函数的代码如下所示：

```
// 检查当前线程是否持有指定优先级的锁
// 如果没有，则触发一个异常
func assertRankHeld(rank lockRank) {
    if g := getg(); g.m.locksRank < rank || g.lockRank < rank {
        throw("lock rank inconsistency")
    }
}
```

总之，assertRankHeld 函数的作用是确保当前线程对指定优先级的锁进行排他访问，在调试和优化时可以大大增强代码健壮性和正确性。



### worldStopped

worldStopped函数是用于确认世界是否已经停止的一个辅助函数。在Go语言运行时中，世界是指所有的goroutine共同运行的上下文，在运行时中，可能会需要在世界停止的时候进行一些操作，例如垃圾回收。

在lockrank_off.go文件中，这个函数的作用是用于确认是否存在正在运行的非系统(goroutine没有被标记为系统)goroutine，如果不存在，则代表所有的非系统goroutine都已经停止，此时可以进行垃圾回收等操作。

具体实现方法是通过调用goexitSched函数，该函数会检查所有正在运行的goroutine的标记位，如果不存在非系统goroutine，则停止世界，否则继续让goroutine运行。这个函数在lockrank_off.go文件之外也有其他文件中使用。



### worldStarted

lockrank_off.go是Go语言中runtime包的一部分，该文件包含了lockrank模块的禁用实现。而worldStarted函数是lockrank模块的一个重要函数，其作用是在每个P启动M之后，在mstart()和acquirep()之间，它设置了一个P的全局状态，表示它已经启动并起动了任何相应的goroutine。其具体作用如下：

1. 确保每个P在世界启动后都有自己的goroutine运行。

2. 设置P状态，表示它已经开始运行，并已经“唤醒”了任何相应的goroutine。

3. 保证在P运行之前其他goroutine不会直接加锁，确保P的goroutine不会被卡住。

4. 通过确保所有P都已经启动并处于活动状态来协调所有调度器的问题。

总的来说，worldStarted函数在Go语言的运行时环境中起着非常重要的作用，确保了所有程序组件的同步和协调，是实现Go语言高效和鲁棒性的关键之一。



### assertWorldStopped

lockrank_off.go文件中的assertWorldStopped函数用于检验当前Goroutine所在的P状态是否为Psyscall，如果不是，则会抛出一个panic。这个函数的作用是在确保在临界区中暂停了全局抢占调度器的前提下，执行对临界区资源的加锁操作。

在Go语言中，当Goroutine执行临界区代码时，为了避免线程竞争和死锁等问题，需要使用一些锁机制来保护临界区资源。而为了保证锁的正确性和效率，需要对Goroutine执行的环境进行限制。在全局抢占调度器中，Goroutine会被随时中断，而这种中断可能会导致锁的状态不正确或者死锁等问题。因此，在执行临界区代码之前，我们需要先暂停全局抢占调度器，以保证临界区代码的正确性和效率。

该函数就是在临界区加锁操作中用来检验当前Goroutine所在的P状态的，只有当P的状态为Psyscall时，才可以执行加锁操作，否则会抛出一个panic。这个panic会中断Goroutine的执行，以提醒开发者需要遵守临界区加锁的规定，避免出现锁状态不正确的问题。

总之，assertWorldStopped函数是在临界区加锁之前执行的一个检验函数，用于保证当前Goroutine所在的P状态为Psyscall，从而保证临界区代码的正确性和效率。



### assertWorldStoppedOrLockHeld

assertWorldStoppedOrLockHeld函数的作用是检查当前goroutine是否处于一个已经停止或持有锁的状态。该函数通常在编写和调试锁相关的代码时使用，可以确保goroutine处于正确的状态，从而避免出现因竞争条件引起的并发问题。

具体来说，assertWorldStoppedOrLockHeld函数会检查四个条件，并对其进行断言：

1. 当前goroutine必须处于非投递状态，即没有处于任何M的运行队列中。

2. 当前goroutine必须持有排他锁。该锁可以是全局的锁（如glock），也可以是其他特定的锁，如gcBgMarkWorkerLock。

3. 如果全局停止标志（stopTheWorld标志）已经设置，则当前goroutine必须是停止状态。

4. 如果引导GC（全局或本地）正在进行，则当前goroutine必须属于工作线程。否则，该函数将引发Panic异常。

在满足这四个条件之后，该函数将退出，并且程序将继续执行。否则，如果有任何一个条件不满足，该函数将抛出异常，并且程序将停止。



