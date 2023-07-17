# File: mcheckmark.go

mcheckmark.go这个文件是Go语言运行时中的一个组成部分，其主要作用是在协程调度时控制M的工作以及异常回复。

Go语言是基于协程的并发模型，其使用调度器来控制协程的执行。在运行时，每个协程都会分配一个M，用于执行协程中的函数。mcheckmark.go中的代码实现了M的标记检查点机制，用于控制M的调度，并处理异常情况。

具体来说，mcheckmark.go中的代码实现了以下功能：

1. 在协程调度时，检查M是否需要进行退出操作，如需要则执行退出操作；

2. 标记M的状态，如设置M的运行状态；

3. 处理M的异常情况，如崩溃和死锁等，以保证程序的正确执行。

总之，mcheckmark.go是Go运行时中非常重要的一部分，它实现了M的管理和调度，保证了Go语言程序的正确执行。




---

### Var:

### useCheckmark

在Go语言的并发模型中，每个线程（goroutine）都是由M（Machine）所托管。M管理着与goroutine相关的一些状态，例如它所属的P（Processor）和当前运行的抢占点（preemption point）。当goroutine处于阻塞状态时，M需要尝试去帮助其它goroutine来运行，以提高整个程序的并发性能。

MCheckmark机制就是为了实现这一点而诞生的。当MCheckmark被启用时，会在一定的时间段内随机触发特定的抢占点给当前线程，在此处中断线程，发现有新的未被抢占的goroutine然后停止当前做的工作，到其它线程中找一个空闲的并协助其调度。

useCheckmark就是用来控制MCheckmark这一机制是否开启的开关。当useCheckmark为1时，MCheckmark开启，否则关闭。这个变量最开始被设置为1（开启状态），但在某些情况下会被设置为0（关闭状态），例如在垃圾回收期间禁用MCheckmark，以避免在GC期间发生抢占点，从而导致GC的不稳定性。

在默认情况下，useCheckmark被设置为1，以允许Go语言的运行时系统动态地调整M的状态，以优化程序的并发性能。






---

### Structs:

### checkmarksMap

checkmarksMap结构体在mcheckmark.go文件中用于记录哪些Goroutine需要在执行期间执行垃圾回收标记。这个结构体是一个记录Goroutine ID的map，它可以用于在Goroutine启动和退出时维护需要执行垃圾回收标记的信息。当一个Goroutine启动时，它会添加到checkmarksMap中，当它退出时，它会从checkmarksMap中删除。在执行期间，垃圾回收器会定期扫描checkmarksMap来检查哪些Goroutine需要标记，并执行必要的操作。

通过记录需要执行垃圾回收标记的Goroutine信息，checkmarksMap结构体使得垃圾回收器能够在不影响正在执行的Goroutine的情况下执行标记，从而提高了程序的性能和稳定性。



## Functions:

### startCheckmarks

startCheckmarks这个函数的主要作用是启动goroutine调度器中的插入抢占式调度的标记（checkmark），以确保Goroutine调度器可以快速切换执行的Goroutine。

在Goroutine调度器中，当一个Goroutine调用了一些阻塞且会引起长时间等待的函数（如IO操作），Goroutine调度器就会挂起当前的Goroutine，并且开始寻找可以被执行的Goroutine。

但是，Goroutine调度器在寻找可执行的Goroutine时，可能会因为某些原因而卡住，比如说某些任务过多，或者零散的任务太多，导致调度器很难找到一组可以执行的任务。

为了解决这个问题，Goroutine调度器引入了checkmarks的概念。checkmark是一个标记，表示在特定的时间点，调度器需要立即挂起当前正在执行的Goroutine，并强制切换到另一个可执行的Goroutine。checkmark机制可以使调度器更加高效，无论在何时都可以快速切换Goroutine。

startCheckmarks函数的作用就是启用checkmark机制。当startCheckmarks函数被调用时，Goroutine调度器会开始为所有活动的Goroutine创建checkmarks。这些checkmarks会被插入到调度器的任务队列中，以确保可以轮流执行它们。

总之，startCheckmarks函数是Goroutine调度器中一个非常重要的函数，它可以使调度器更高效和可靠，从而更好地管理和调度Goroutine。



### endCheckmarks

endCheckmarks函数是运行时系统中的一个私有函数，其主要作用是停止异步抢占式调度器的检查点标记，将异步的抢占式调度器重新恢复到非检查点状态。在调用endCheckmarks函数之前，调度器会暂停当前的goroutine并安排运行另一个goroutine。当调度器再次调用当前goroutine时，endCheckmarks函数将被调用以恢复goroutine的状态。

具体来说，endCheckmarks函数的功能如下：

1. 将当前的代码执行状态清零，以便重新启动抢占式调度器。

2. 如果当前的调度状态存储了一个检查点，该检查点将被撤销并解锁。

3. 将当前goroutine的M的gcing标志复位为false，以确保GC期间不会再次抢占该goroutine。

4. 恢复与检查点相关的计数器和指针，包括检查点堆栈指针、时间戳计数器和检查点标记计数器。

总之，endCheckmarks函数的作用是恢复运行时系统的状态，以便继续执行当前的goroutine。



### setCheckmark

setCheckmark是runtime中的一个函数，它的主要作用是将标记位设置为1，以便在后续的垃圾回收中收集扫描根对象并标记其他对象。

更具体地说，当Go程序需要启动垃圾收集时，它会调用setCheckmark来确保收集器在扫描对象时不会遗漏任何根对象。在执行垃圾回收时，收集器将扫描所有当前正在运行的goroutine栈中的变量，并标记它们作为根对象。如果setCheckmark没有被调用，这些根对象就有可能被遗漏，导致内存泄漏和程序行为异常。

另外，setCheckmark还会设置一个全局计数器，以便于跟踪当前已经设置的标记位数量。该计数器对于垃圾回收算法的优化非常重要，因为它可以帮助收集器更准确地估计未标记对象的数量，从而优化内存回收和程序性能。

综上所述，setCheckmark是一个重要的函数，它确保了程序在进行垃圾回收时能够正确处理所有必要的引用。



