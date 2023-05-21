# File: lock_sema.go

lock_sema.go文件是Go语言运行时包中的一个源文件，其作用是实现了基于信号量的锁机制。

在Go的并发操作中，不同的goroutine之间需要协作完成任务，而在这个过程中可能会存在共享内存的并发访问问题。为了解决这种并发访问问题，Go语言提供了一种轻量级线程模型goroutine，以及对应的同步机制，如Mutex、RWMutex等。

而lock_sema.go中实现的信号量锁机制，是为了在一些特殊情况下提供更高效、更灵活的同步机制。相比于Mutex等锁机制，信号量锁的实现更为底层，更为轻量级，因此可以更快地获取和释放锁，适用于一些高并发场景。

具体实现上，lock_sema.go中实现了三个结构体，分别是semacquire、semrelease、and semasleep。其中semacquire用于获取锁，semrelease用于释放锁，semasleep用于等待锁被释放。通过这三个结构体的组合使用，可以实现信号量锁机制的高效同步。

在实际应用中，lock_sema.go中信号量锁机制的使用并不广泛，大部分情况下还是使用Mutex等锁机制。但在一些特殊的高并发场景下，信号量锁机制可以提供更为高效、精细的同步控制，因此在Go语言的运行时库中得以保留和实现。

## Functions:

### lock

lock_sema.go中的lock函数是用来抢占一个信号量低级锁的。一个低级锁是一个定义了互斥访问共享资源的机制。当一个goroutine（协程）需要访问共享资源时，它可以尝试获取低级锁，如果该锁没有被其他goroutine持有，则该goroutine获取锁并继续执行相应的操作。如果该锁已经被占用，则该goroutine将被阻塞，等待低级锁的释放。

在lock_sema.go中，lock函数使用了一个等待队列来管理等待该锁的goroutine。如果该锁已经被其他goroutine持有，则当前goroutine将会被添加到该等待队列中，并使用semaBlock函数将goroutine挂起。当锁释放时，锁的拥有者将调用unlock函数，该函数将唤醒一个等待该锁的goroutine，并将其从等待队列中移除，并在该goroutine从semaBlock函数中恢复时允许其获取锁。

除了等待队列和阻塞函数之外，lock函数还定义了一些锁状态，例如mutexLocked和mutexWoken。这些状态用于管理和同步锁状态的变化。

总之，lock函数是锁的重要组成部分。它负责协调和同步goroutine对共享资源的访问，确保每个goroutine都可以安全地访问资源，并保证资源的正确性和一致性。



### lock2

lock_sema.go中的lock2函数实现了对于一个互斥锁（mutex）的加锁功能。这个锁采用了一种基于信号量（semaphore）的实现方式。

具体来说，lock2函数会先试图以“快速路径”（fast path）的方式获取锁，即检查锁是否处于空闲状态，如果是，则立即获取并返回；否则，进入“慢路径”（slow path）。

在slow path中，lock2函数会为当前goroutine创建一个休眠信号量（sema），将它加入到mutex的等待队列中，并调用park函数将当前goroutine挂起。在park函数挂起时，该goroutine会进入休眠状态，等待锁释放并唤醒它。当锁释放时，唤醒信号量会被传递给等待队列中的一个休眠信号量，并由该信号量唤醒相应的goroutine继续执行；其它等待的goroutine会继续休眠。这样就实现了锁的等待和唤醒机制。

需要注意的是，lock2函数中还会设置一个isHandoff标记，用于指示当前唤醒等待的goroutine是否需要立即获取锁。如果isHandoff为true，则表示当前等待的goroutine必须要立即获取锁，否则它将会重新进入到等待队列中。如果isHandoff为false，则表示当前等待的goroutine可以选择将锁设置为“放手模式”，即使在当前goroutine等待期间，其它goroutine获取了锁也可以立即放弃锁而不必等待。

总的来说，lock2函数实现了互斥锁的等待和唤醒机制，支持快速路径和慢路径两种获取锁的方式，是Go语言运行时中重要的锁实现之一。



### unlock

在Go语言中，锁是一种用于协调多个goroutine之间并发访问共享资源的机制。在runtime包中的lock_sema.go文件中，实现了基于信号量机制的锁，其中的unlock函数用于释放一个锁。

具体来说，当一个goroutine获得了一个锁时，其他goroutine就无法获得相同的锁，直到该goroutine将锁释放。unlock函数的作用就是将此锁的占用状态标记为可用，并通知等待队列中的其他goroutine有锁可用。具体实现方法如下：

1. 首先，将锁状态设置为可用状态。在lock_sema.go文件中，锁的状态被定义为一个semaphore类型的指针，使用atomic.StorePointer函数将其设置为nil。

2. 然后，唤醒等待队列中的其他goroutine。当一个goroutine在等待队列中等待一个锁时，它会在Lock函数中相应的信号量上递减，随着unlock函数的调用，这个信号量会递增回释放信号量，从而唤醒等待队列中的其他goroutine。

总之，unlock函数用于释放一个锁，并唤醒等待队列中的其他goroutine，使它们有机会竞争该锁。



### unlock2

unlock2函数作用是将信号量锁的计数值减1并唤醒等待该信号量锁的协程。

在Go语言中，mutex和sync.Cond都是基于信号量实现的，具体来说，就是使用了内置的信号量锁。在lock_sema.go文件中，通过一个lockWithRank和unlock2函数来实现信号量锁的加锁和解锁操作。

在unlock2函数中，首先通过atomic.AddInt32函数对计数器lock.sema等于waiters的值减1，表示有一个协程离开了信号量锁的等待队列，这样就可以唤醒一个处于等待队列中的协程。接下来再通过系统调用futex将等待在该信号量锁上的协程从等待状态转换为就绪状态，这样一来，被唤醒的协程就可以立即获取到该信号量锁。

总的来说，unlock2函数的作用就是将信号量锁的计数值减1并唤醒等待该信号量锁的协程，这样就可以保证在多个协程同时抢占同一个锁的时候，只有一个协程能够获得锁，并且其他协程会被等待唤醒，保证了协程之间的同步和互斥性。



### noteclear

在Go语言的runtime包的lock_sema.go文件中，noteclear是一个函数，用于释放和清除一个已经被使用过的信号量（semaphore）中的资源。

在Go语言中，信号量是一种同步原语，通过信号量可以实现线程之间的互斥和同步。当一个线程需要进入临界区时，它会申请一个信号量，如果信号量的值为0，说明此时有别的线程正在使用这段临界区，申请锁的线程必须等待。当进入临界区的线程执行完毕并退出临界区时，它会释放占用的信号量，这样其他线程就可以进入临界区了。

在Go语言中，noteclear函数的作用就是在一个信号量被释放之后，从信号量中清除所有的相关信息，以保证信号量可以重用。noteclear函数会将信号量的等待队列和通知队列中的所有元素都清除掉，并将信号量的状态重置为未锁定状态，从而使得其他线程可以再次申请该信号量。通过释放信号量并调用noteclear函数，可以避免信号量的资源浪费和性能问题。



### notewakeup

notewakeup是一个用于唤醒等待在WaitNote上的goroutine的函数。

在Go的SDK中，goroutine是调度的基本单位，它往往是轻量级的线程，一个Go程序可以同时运行多个goroutine，它们之间是并发执行的。在运行时，goroutine会不断地从全局任务队列中取出任务进行执行。但是，有些时候，任务队列是空的，这时goroutine需要进入等待状态，等待其他goroutine的任务进入队列。这样就需要用到WaitGroup或Cond或Mutex等同步机制。其中，notewakeup主要用于Cond中，实现其Wait方法的功能。

Cond（条件变量）是一种用于goroutine之间的同步的机制，Wait方法会使调用它的goroutine进入等待状态，等待一个唤醒信号。当其他goroutine的行为改变了等待条件时，唤醒信号会被发出，此时Wait方法会将进入等待状态的goroutine唤醒。notewakeup函数被用于在某个goroutine满足某个条件时唤醒其他等待的goroutine。

notewakeup函数的具体实现是通过向被唤醒的goroutine的task中加入一个suspend点，使得它下一次执行时会被挂起，继而从等待列表中移除。再通过调度器的goexit方法执行defer函数，最后再一次调用调度器的schedule方法，重新对该goroutine进行调度。这样，被唤醒的goroutine就可以回到任务队列继续执行。



### notesleep

lock_sema.go这个文件是Go语言的运行时系统（runtime）中的一部分，它包含了锁实现的代码。notesleep函数是其中一个辅助函数，用于等待一个信号量，直到被唤醒。

信号量是一个非负整数，可以用于多线程同步和互斥。当信号量的值为0时，所有等待这个信号量的线程都会被阻塞。当有其他线程释放这个信号量时，等待线程中的一个会被唤醒，继续执行。

notesleep函数被用于实现带有timeout的锁等待，如果锁不可用，调用notesleep函数会等待一定的时间。具体实现方式是先尝试获取锁，如果获取成功直接返回；获取失败则调用notesleep函数，等待一段时间后再次尝试获取锁。如果超时依然无法获得锁，则返回失败。

在实现锁等待的过程中，notesleep函数使用了Go语言中的select语句，结合channel来实现等待和唤醒线程的操作。这种实现方式比传统的信号量方式更加灵活和高效。



### notetsleep_internal

notetsleep_internal函数是Go语言中的一个阻塞等待机制，它是Go语言实现的一种基于信号量的等待机制，在很多地方都被广泛使用。notetsleep_internal函数用于等待一个事件的发生，当等待的事件发生时，该函数将返回，否则将一直阻塞。

notetsleep_internal函数主要作用是让goroutine等待事件的发生，并将该goroutine的状态设置为睡眠状态。在等待事件的过程中，该goroutine不会消耗CPU资源，因此可以有效节省系统资源。当事件发生时，notetsleep_internal函数会被唤醒，并将该goroutine的状态设置为就绪状态。

notetsleep_internal函数的实现是基于一个计数器和两个信号量的机制。当计数器为0时，表示事件没有发生，此时goroutine会被阻塞。当事件发生时，计数器加1，同时唤醒等待该事件的所有goroutine。当goroutine被唤醒时，它会将计数器减1，然后继续执行。

notetsleep_internal函数的实现在go/src/runtime/lock_sema.go文件中，是Go语言实现中比较核心的一个部分。通过使用notetsleep_internal函数，可以实现很多基于等待机制的并发控制，如互斥锁、读写锁等。



### notetsleep

函数名称：notetsleep

作用：notetsleep 函数主要是用于等待 goroutine 结构体中的通知或休眠信号。它会一直等待，直到收到信号，然后返回。如果等待过程中出现错误，则会将错误信息返回。

函数参数：

- note *note：指向 goroutine 结构体的指针。
- ns int64：等待时间，以纳秒为单位。

函数返回值：

- bool：如果函数调用成功，则返回 true，否则返回 false。
- error：如果出现错误，则返回错误信息。

函数实现原理：notetsleep 函数主要是通过调用 notetsleep_internal 函数实现的。notetsleep_internal 函数会注册在 goroutine 结构体的等待队列中，然后处于循环状态来等待通知或休眠信号。如果收到信号，则会从等待队列中删除自己，并返回 true。如果出现错误，则会返回错误信息，否则会一直等待，直到等待时间过期。

notetsleep 函数的作用是在等待 goroutine 结构体中的通知或休眠信号时候使用的。它和其他的等待函数不同的是，它会注册到 goroutine 的等待队列中，并一直等待，直到收到信号。在等待过程中，它会不停地轮询 goroutine 等待队列，以便及时收到信号。如果等待时间过期，则函数会立刻返回，并将返回值设置为 false。

notetsleep 函数通常用于阻塞当前的 goroutine，直到接收到其他 goroutine 中的通知。它可以和其他的同步原语一起使用，比如 WaitGroup、Mutex 和 Cond 等，以实现多个 goroutine 的同步操作。



### notetsleepg

notetsleepg是golang中实现的一种基于信号量的非阻塞睡眠机制，用于实现goroutine的调度。在golang中，每个goroutine都有一个G struct，在这个struct中存储着goroutine相关的信息。

在goroutine的调度过程中，有一种常见的情况就是goroutine需要等待某个事件的产生或者某个条件的满足才能继续向下执行，这时就需要用到notetsleepg函数。notetsleepg函数主要有以下几个作用：

1. 等待事件的产生或条件的满足：当一个goroutine需要等待某个事件的产生或者某个条件的满足时，可以调用notetsleepg函数进行等待。这个函数使用信号量来实现非阻塞的等待，避免了采用传统的阻塞等待方式引起的性能问题。

2. 释放CPU资源：在等待期间，notetsleepg函数会释放CPU资源，避免了CPU资源的浪费。

3. 响应事件：当事件发生或条件满足时，notetsleepg函数会响应当前goroutine的调度，使得该goroutine可以继续执行。

4. 支持多个goroutine的唤醒：notetsleepg函数支持多个goroutine的唤醒，即当事件发生或条件满足时，可以唤醒等待该事件或条件的多个goroutine。

总之，notetsleepg函数将信号量机制应用于goroutine的调度中，提高了goroutine的响应速度和系统的性能。



### beforeIdle

在Go语言运行时的lock_sema.go文件中，beforeIdle这个函数是用来检查goroutine是否被阻塞并且没有可以运行的任务。如果是这种情况，beforeIdle函数将会设置一个标志，通知调度器可以进行操作系统级别的休眠（进入睡眠状态）。通过这种方式，可以防止占用CPU资源，并允许其他goroutine执行。

在beforeIdle函数中，会先获取当前被阻塞的goroutine的P（处理器），将其放入一个临时的切片中，并且调用readyForSyscall函数。这个函数将设置一个标志，表示当前协程可以让出处理器，并且准备进入系统调用（例如Unix的select，poll或epoll）。此外，beforeIdle还会检查是否存在待处理的系统调用，如果存在，则会中止休眠并执行系统调用。

通过beforeIdle函数，可以确保在没有活动任务时，Go运行时会正确地进入休眠状态，而不会占用CPU资源。同时，beforeIdle也是Go语言调度器实现中重要的一环，有助于保证系统的稳定性和性能。



### checkTimeouts

checkTimeouts函数的主要作用是检查所有的锁等待队列，看看是否有超时的锁等待者。如果存在超时的锁等待者，它会将其从等待队列中移除，并将其标记为超时。

具体来说，checkTimeouts函数会遍历所有P的local runqueue和global runqueue中的锁等待队列。对于每个队列中的锁等待者，它会计算出等待时间是否超过了其指定的超时时间。如果等待时间超过了超时时间，它就会将该锁等待者从队列中移除，并将其状态设置为超时。

checkTimeouts函数还会根据锁等待者的类型，执行不同的超时处理逻辑。对于Mutex等进程内同步锁，它会将超时的锁等待者放入死锁检测队列，以便后续的死锁检测。对于Semaphore等信号量类型的锁，它会直接将超时的锁等待者从等待队列中移除，并且将其状态设置为超时。

总的来说，checkTimeouts函数的作用是避免锁等待队列中的锁等待者因为死锁或者超时而一直占用资源，导致程序性能下降或者崩溃。



