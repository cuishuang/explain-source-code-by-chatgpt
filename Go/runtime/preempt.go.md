# File: preempt.go

preempt.go 是 Go 语言运行时的一个重要组件，它负责处理 Go 协程的调度。当一个 Go 协程在运行时长超过一定时间，或者发生阻塞时，调度器会将其停止并将处理器（内核线程）分配到其他的协程上，这个过程就是抢占。

preempt.go 中实现了一个 goroutine 执行时间片的计算器，它会在每个 goroutine 执行时向计数器增加一定的值（因为不同的操作系统时间片大小不同，所以该值在 Windows 和 Unix/Linux 平台上是不同的），当计数器超过限制时，就会触发抢占，让其他 goroutine 有机会运行。

此外，preempt.go 中还定义了一些宏，如 P 系统的最大数量、抢占的最小间隔等等。它还实现了 goroutine 的阻塞、唤醒和检查状态的功能，这些操作都是协调器构建在 CPU 和 OS 之上的一个重要组成部分。

总的来说，preempt.go 的作用是：实现了 Go 协程调度中的重要功能，包括时间片计算和抢占；定义了一些常量和宏来控制调度器的行为；实现了 goroutine 的阻塞和唤醒等协程状态管理操作。




---

### Var:

### asyncPreemptStack

在Go语言的运行时系统中，当一个goroutine正在执行时，它会一直运行直到当前的代码块执行完毕或者遇到了合适的阻塞操作。如果一个goroutine长时间占用了CPU，那么其它的goroutine就会被阻塞，甚至导致整个程序的响应性能下降。为了防止这种情况发生，Go语言的运行时系统实现了一种称为异步抢占的机制，可以让一个goroutine被另一个goroutine所抢占。异步抢占是Go语言实现并发和并行的重要机制之一。

在实现异步抢占的过程中，asyncPreemptStack这个变量起到了关键的作用。当一个goroutine执行时，它会在栈上记录哪些函数调用了其他的函数，这些函数调用的记录就叫做栈帧。当一个goroutine被抢占时，运行时系统需要记录当前goroutine的栈帧，以便在恢复时能够回到正确的位置。asyncPreemptStack就是用来记录这些栈帧的。

具体地说，asyncPreemptStack是一个指向栈帧的指针数组，其中每个指针都指向一个栈帧。当一个新的goroutine被创建时，它会从asyncPreemptStack中分配一个栈帧，并将该栈帧作为新goroutine的根栈帧。当一个goroutine被抢占时，它的栈帧会添加到asyncPreemptStack中，以便在恢复时可以回到正确的位置。当一个goroutine执行完毕时，其栈帧会从asyncPreemptStack中弹出，以便下一个goroutine使用。

总之，asyncPreemptStack起到了异步抢占机制中保存和恢复栈帧的作用，是Go语言实现并发和并行的重要组成部分。






---

### Structs:

### suspendGState

在Go语言的运行时系统中，goroutine是 Go 程序并发执行的基本单元。在运行时系统中，通过使用调度器来管理 Goroutine 的并行执行。由于调度器需要在多个 goroutine 之间切换，因此需要在不同的时间点暂停或恢复 goroutine 的执行状态。这个过程被称为抢占(preemption)。

在runtime/preempt.go中，suspendGState结构体是一个表示 goroutine 的执行状态的数据结构。它的作用是保存原本正在执行的 goroutine 的执行状态。当调度器需要切换到其他 goroutine 时，suspendGState被用于保存当前 goroutine 的执行状态（比如栈指针、寄存器状态等）和上下文信息，并在之后恢复到这个 goroutine执行时的状态。这个过程包含了一些处理器特定的机制，比如处理器状态和调度器等待的 goroutine，所以需要使用一些辅助函数来实现恢复和保存时间。

此外，此结构体还包含了一些状态标志，这些标志表示了调度器的处理状态，并控制了调度器的行为。例如：

- runnable：表示当前 goroutine 的执行状态是否可运行。
- scan：表示当前 goroutine 执行状态和工作窃取时的扫描状态。
- preemptionRequested：表示是否已经授权该 goroutine 抢占调度器。



## Functions:

### suspendG

`suspendG` 函数是 Go 语言运行时中的一个重要函数，它用于把一个 Goroutine 从运行状态转为挂起状态。下面是它的定义：

```go
func suspendG() bool {
    gp := getg()
    if !gp.isSystemGoroutine(true) || readgstatus(gp)&^_Gscan != _Grunning {
        return false
    }
    casgstatus(gp, _Grunning, _Gwaiting)
    park_m(gp)
    return true
}
```

该函数的执行过程如下：

1. 获取当前 Goroutine (`gp = getg()`)。
2. 检查当前 Goroutine 是否是系统 Goroutine 以及是否处于运行状态（`gp.isSystemGoroutine(true)` 和 `readgstatus(gp)&^_Gscan != _Grunning`）。
3. 如果不满足上述条件，函数直接返回 `false`，即不进行挂起操作。
4. 如果满足上述条件，将 Goroutine 的状态从运行状态改为等待状态（`casgstatus(gp, _Grunning, _Gwaiting)`）。
5. 而后，调用 `park_m(gp)` 函数，使得当前 Goroutine 进入休眠状态，直到被唤醒。

通过调用 `suspendG()` 函数，应用程序在协作式调度时可以主动将当前 Goroutine 从运行状态转为阻塞状态，从而让其他 Goroutine 有机会获得执行权。这样做的好处是可以减少 Goroutine 之间的竞争，提高程序的性能和可靠性。

需要注意的是，由于 `suspendG()` 函数会将当前 Goroutine 置为等待状态，因此它需要在其他 Goroutine 中被调用，才能起到转移执行权的作用。



### resumeG

resumeG函数是Go语言运行时系统(runtime)中实现调度器的重要函数之一。在Go语言的并发编程中，调度器(scheduler)负责管理和调度被创建的goroutine，以便让它们在多个CPU核心之间高效地共享资源。调度器的工作是在Go语言的运行时系统中完成的，而resumeG函数是其中一个实现调度器功能的重要函数。

resumeG函数的主要作用是将被暂停(pause状态)的goroutine解除暂停(paused状态)，并将其重新放回到调度器的调度队列中，以便被调度器再次调度和执行。当goroutine被暂停时，调度器会将其信息保存在一个数据结构中，以便在以后调度时可以快速恢复其执行状态。调度器使用resumeG函数来实现恢复被暂停的goroutine的执行状态，并将其加入到可以执行的goroutine队列中。

具体地说，resumeG函数会读取被暂停的goroutine的执行状态信息，并使用运行时系统中的线程(thread)来恢复该goroutine的执行状态。一旦该goroutine的执行状态被恢复，resumeG函数会将该goroutine加入到可以执行的goroutine队列中，并在下一次调度时选中该goroutine来执行。这样，被暂停的goroutine就可以重新执行，并继续完成其对应的任务。

总的来说，resumeG函数在Go语言的并发编程中具有非常重要的作用，它是实现调度器等核心功能的关键之一。它用于恢复被暂停的goroutine的执行状态，并将其重新放回调度器的调度队列中，以便被再次调度和执行。



### canPreemptM

canPreemptM是一个函数，用于判断当前M（machine）是否可以被抢占（preempt），即是否可以在当前运行的代码中间被停止，把控制权转交给其他M执行。这个函数的主要作用是在协程（goroutine）的调度中，通过调用它来判断当前M是否被允许被抢占，以便进行协程的调度。

在Go语言中，每个可执行的goroutine都会被分配到一个M上进行执行。可以通过M的数量来控制并行度，从而提高程序的性能。当一个M正在执行一个goroutine时，如果有其他goroutine也需要执行，那么调度器就会决定是否将当前的M抢占，让其他M来运行这个goroutine。canPreemptM函数的主要作用就是判断当前M是否可以被抢占。

canPreemptM函数的实现如下：

```
// canPreemptM reports whether mp is in a state that it can be preempted.
// It returns true if mp is running in a non-invasive or preemptable state.
// mp must be locked.
//go:systemstack
func canPreemptM(mp *m) bool {
    if mp.preemptoff != "" || mp.mallocing != 0 || mp.spinning || mp.fastrand()&0xff == 0 {
        return false
    }
    gp := mp.curg
    if gp == nil || gp.startedSyscall() || !gp.atomicstatus.inStatus(gpatomicstatusrunnable) {
        return true
    }
    return false
}
```

可以看到，canPreemptM函数首先判断了一些状态，包括当前M的抢占状态（mp.preemptoff）、内存分配状态（mp.mallocing）、自旋状态（mp.spinning）以及随机数等。如果这些状态满足一定条件，则表示当前M不能被抢占。

接下来，canPreemptM函数判断当前M中正在运行的goroutine的状态。如果该goroutine正在执行系统调用（gp.startedSyscall()）或者处于可以运行的状态（gp.atomicstatus.inStatus(gpatomicstatusrunnable)），则说明当前M可以被抢占。

最后，如果当前M可以被抢占，则返回true，否则返回false。

总的来说，canPreemptM函数的作用就是在协程调度时，判断当前M是否可以被抢占，以便进行调度，提高程序的并行度和性能。



### asyncPreempt

asyncPreempt函数是Golang运行时系统中的一部分，其主要作用是异步抢占当前正在执行的goroutine。

在Golang中，goroutine是一种轻量级的线程，可以实现高并发和异步编程，但是如果goroutine的执行时间过长，可能会影响程序的性能和响应速度。为了解决这个问题，Golang运行时系统会定期检查当前正在执行的goroutine的执行时间是否超过了一定的阈值，如果超过了阈值，就会抢占当前的goroutine，并切换到其他的goroutine。

asyncPreempt函数就是实现这个抢占过程的函数，它会在当前goroutine执行的一些关键点上检查是否需要抢占，并在需要时抢占当前的goroutine。具体的实现方式是通过设置goroutine的状态为“抢占中”，然后将goroutine加入到运行时系统的抢占队列中，等待下一次调度器的调度。

asyncPreempt函数的作用在于使得Golang运行时系统具有了一定的抢占能力，可以避免某个goroutine执行时间过长而影响程序的性能和响应速度。与之相关的还有其他的函数，比如preempt和preemptM，它们也都是运行时系统中的抢占相关函数，但是它们的作用略微不同，preempt函数用于在某个goroutine执行的临界点上抢占当前的goroutine，而preemptM函数用于在某个线程执行的临界点上抢占当前的线程。



### asyncPreempt2

asyncPreempt2是Go语言中用于协程抢占的函数之一，它的作用是提供异步抢占能力，通过将G标记为异步抢占标记，可以实现异步抢占的目的。

在Go语言中，协程的调度是基于时间片轮转的，如果一个协程执行的时间超过了时间片，它就会被挂起，等待调度器重新分配时间片。但是，如果一个协程长时间阻塞在某个系统调用或者其他同步操作中，它会占用一个P（处理器），而其他协程却无法运行。为了解决这个问题，Go语言提供了异步抢占机制，即在协程执行时间过长时，强制将其挂起，以便其他协程有机会运行。

asyncPreempt2函数就是用于实现异步抢占的关键函数之一。当调度器检测到一个协程执行时间过长时，会调用asyncPreempt2函数，向其传递当前协程所在的上下文信息，并标记该协程为异步抢占标记。在函数内部，asyncPreempt2会根据当前协程所在的位置，找到相应的m（执行线程），并调用m的asyncPreempt方法，将异步抢占标记设置到m对应的G上，从而实现异步抢占的目的。

总之，asyncPreempt2函数在Go语言中扮演着实现协程异步抢占的重要角色，通过它的实现，Go语言可以更加高效和灵活地进行协程调度和管理。



### init

init函数是Go语言中一个特殊的函数，用于在Go程序启动时自动执行，它没有参数和返回值。在预编译阶段，编译器会根据程序中所有的init函数生成一个init程序，并将其链接到已经编译好的程序中。这样，在Go程序启动时，所有的init函数就会被自动执行。

在go/src/runtime/preempt.go中，init函数的作用如下：

1. 初始化preemptTimer变量

init函数首先会初始化preemptTimer变量，该变量用于定时触发协程抢占操作。preemptTimer是runtime包中的全局变量，其类型为*time.Timer。在Go程序启动时，会启动一个goroutine来定时触发preemptTimer。

2. 注册钩子函数

init函数还会注册一个钩子函数（hook），该钩子函数会在goroutine被抢占之前执行。该函数的作用是记录goroutine的状态，例如它的堆栈信息和调度器状态等。这些信息对于Go语言的debugger和profiler有很大帮助。

3. 注册抢占函数

init函数还会注册一个抢占函数（preempt），该函数会在preemptTimer时间到达时被调用，从而触发协程抢占操作。该函数的作用是将当前执行的goroutine阻塞，并强制调度器调度另一个可执行的goroutine来执行。

总之，preempt.go文件中的init函数主要用于初始化一些全局变量、注册钩子函数和抢占函数等，它们都是Go程序调度器实现的一部分。这些函数的工作机制复杂，但它们都是为了实现Go语言的并发编程模型而存在的。



### wantAsyncPreempt

preempt.go文件中的wantAsyncPreempt函数用于设置当前线程是否希望异步抢占。异步抢占是一种机制，允许操作系统通过中断或信号强制地挂起当前线程，以便其他线程能够访问共享资源或执行临界区域的代码。

具体来说，wantAsyncPreempt函数检查当前线程是否可以安全地进行异步抢占，即当线程正在运行可抢占函数时，且当前线程没有被阻塞，也没有处于临界区域，同时该线程还没有请求抢占或等待抢占时，这个函数会返回true。如果wantAsyncPreempt返回true，则表示当前线程希望在适当的时候进行异步抢占。

该函数使用的一些关键变量和函数包括：

- gp：当前线程的goroutine结构。
- g_preempt：表示当前线程是否在等待抢占。
- getg()：返回当前线程的goroutine结构。
- canPreemptM()：检查当前M是否可抢占。
- mheap_.lock：堆锁，同一时刻只能有一个M进行堆操作。

在进行异步抢占时，操作系统会中断当前线程的执行，并将控制转移到其他线程。然后，当当前线程再次被分配CPU时间时，它将从中断的位置开始执行。因此，异步抢占也会产生一定的性能开销。

在Go语言中，wantAsyncPreempt函数主要用于在调度器中进行线程的抢占操作。通过适当地设置wantAsyncPreempt，调度器可以在需要时在不打断线程的进程中进行抢占操作，从而提高程序的性能和响应速度。



### isAsyncSafePoint

isAsyncSafePoint函数的作用是判断当前线程是否处于一个异步安全点。

异步安全点（Async Safe Point）指的是程序执行时，能够安全地停止当前线程并进行一些操作，而不会导致数据结构不一致或者其他不安全的情况。

在Go语言中，异步安全点的使用场景主要包括垃圾回收、栈扩容等。当线程处于这些场景时，需要保证其处于异步安全点，以确保程序的正确性。

isAsyncSafePoint函数的具体实现主要包括以下几个步骤：

1. 判断当前线程的状态是否处于阻塞状态。如果是，说明线程正处于一个异步安全点。

2. 判断当前线程是否需要参与垃圾回收。如果是，说明线程正处于一个异步安全点。

3. 判断当前线程是否处于栈扩容状态。如果是，说明线程正处于一个异步安全点。

通过这些判断，isAsyncSafePoint函数能够准确地判断当前线程是否处于一个异步安全点，从而保证程序的正确性和安全性。



