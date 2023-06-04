# File: sema.go

sema.go这个文件是Go语言中实现信号量的关键文件，其中实现了两种类型的信号量：waitgroup和sema。

waitgroup实现了类似于Java中CountDownLatch的功能，即在等待一组并发任务完成后，再进行后续的操作。这个功能是通过实现一个计数器来实现的，该计数器初始值为0，每个并发任务完成后会通过Done方法将计数器减1，而等待组的主线程则会通过Wait方法一直等待计数器减为0。

sema是一个可重入的信号量，在实现Go语言的Mutex和Cond等同步机制时都用到了这个信号量。它实现了同步的基本方法——P（类似于Java中的acquire）和V（类似于Java中的release），P可用于申请资源，V则用于释放资源。sema的实现机制是通过一个channel来模拟信号量的操作，当一个goroutine请求信号量时将其放入channel中，而释放信号量时则从channel中取出一个goroutine来执行。这个时候如果channel中没有任何等待的goroutine，则P方法就会直接阻塞。

总之，sema.go这个文件的作用是提供了Go语言中实现信号量所需的核心算法和机制，它是Go语言中实现同步和并发的重要支撑。




---

### Var:

### semtable

semtable是一个用于存储抢占锁（preemptive locks）的哈希表（hash table），其中的键是一个32位整数，代表一个锁的goroutine ID，值是一个指向链表头部的指针，表示这个ID对应的锁在哈希表中的位置。

抢占锁是一种特殊类型的锁，它可以在任何时候被其他正在运行的goroutine获得。为了实现这样的锁，需要在正在持有锁的goroutine中插入抢占点（preemption point），以便其他goroutine有机会获得锁。semtable的作用就在于存储这些抢占点。

在Go语言中，抢占点是通过将锁的goroutine ID插入到semtable中来实现的。每当一个goroutine获取到抢占锁时，它会将自己的ID插入到semtable中，代表这个锁正在被它持有。当其他goroutine尝试获取同一个锁时，它们会向semtable中查找这个ID，如果能够找到，则表示这个锁正在被其他goroutine持有，并且会进入休眠状态，等待锁被释放。

当持有锁的goroutine退出临界区时，它会将自己的ID从semtable中删除，这样其他竞争者就有机会获取到锁。

总之，semtable是一个非常重要的数据结构，用于实现抢占锁的机制，保证Go语言中的并发访问操作的正确性。






---

### Structs:

### semaRoot

semaRoot是一个用于同步的结构体，用于控制一组goroutine的访问并保证其互斥。它的主要作用是维护一个信号量，用于限制并发访问的数量，防止竞态条件的出现。

该结构体中有两个字段：wg和sema，分别表示等待组和信号量。waiters表示正在等待的goroutine数量。

当一个goroutine想要访问某个共享资源时，它需要先获取semaRoot的信号量。如果semaRoot当前的waiters数量达到了限制值，则该goroutine将被阻塞，直到某个其他的goroutine释放了信号量。在获取了信号量之后，该goroutine可以安全地访问共享资源，在完成访问之后，它需要释放信号量。

通过维护这个信号量和等待组，semaRoot可以有效地防止多个goroutine同时访问某个共享资源，从而避免了竞态条件的出现，保证了程序的正确性和安全性。



### semTable

`semTable` 结构体是用来保存所有系统级信号量的，它会在程序运行时被初始化，用来协调并发访问共享资源的操作。信号量是一种同步机制，它允许多个线程在并行执行的同时，能够协调它们在共享资源上的行为，避免出现数据竞争和死锁等问题。

该结构体中包含一个列表 `semaArray`，它是由多个 `semaStruct` 构成的数组，每个 `semaStruct` 对应一个系统级信号量，用于协调一个共享资源的并发访问。每个 `semaStruct` 中包含一个协程队列 `gList` 和一个信号量值 `semaCount`，用来实现信号量的功能。

当一个协程需要访问共享资源时，它会首先尝试获取对应的 `semaStruct` 中的信号量。如果该信号量的值为正数，表示资源没有被占用，协程可以立即获取该资源并将信号量的值减去 1；如果该信号量的值为 0，表示资源正在被占用，协程需要加入到 `gList` 协程队列中等待信号量释放，并进入阻塞状态。当一个协程释放了该资源，它会将信号量的值加上 1，并从 `gList` 中唤醒等待的协程。

总之，`semTable` 结构体的作用就是为多个协程之间提供一种有效的同步机制，确保它们可以安全地并发访问共享资源，并且能够避免进程死锁和数据竞争等问题。



### semaProfileFlags

semaProfileFlags是一个位掩码，用于控制基于信号量的锁的轮询和阻塞的分析。它是在调试期间使用的一个工具，用于调优和分析信号量的使用情况。

该结构体包含以下几个字段：

- semaBlockProfile：如果设置，则会在信号量阻塞时记录调度堆栈信息。
- semaContentionProfile：如果设置，则会在信号量争用时记录调度堆栈信息。
- semaTrace：如果设置，则会记录每个信号量的操作信息，并将其聚合为跟踪文件。
- mutexProfileFraction：用于设置互斥量的阻塞分析的采样率。

其中，semaBlockProfile和semaContentionProfile用于分析信号量阻塞和争用的情况，而semaTrace用于记录每个信号量操作的详细信息。mutexProfileFraction用于设置互斥量阻塞分析的采样率，从而避免在高负载环境下引入过大的开销。

这些标志可以通过环境变量GOTRACE设置。例如：

```
export GOTRACE=threadmutexprofile=1
```

这会启用线程互斥量的阻塞分析，并将结果写入调试文件中。



### notifyList

在 Go 语言中，`notifyList` 是用来实现 Goroutine 和 Channel 之间的同步的数据结构。它是一个带有链表的互斥锁。

`notifyList` 结构体包含以下三个字段：

- waiters: 是一个链表，用来存储等待通知的 Goroutine。
- notify: 是一个标志位，表示是否有通知发生。
- lock：一个用来保护 waiters 字段和 notify 字段的互斥锁。

当一个 Goroutine 调用 Channel 的 `Receive` 方法时，如果 Channel 中没有数据，Go 会将该 Goroutine 加入到 `notifyList` 的等待链表中。在 Channel 写入数据时，会检查是否有等待的 Goroutine，如果有，则会通知它们，并从等待链表中删除。 (`notifyList` 负责通知，而非 Channel)

`notifyList` 的一个非常重要的特性是，它可以保证 Goroutine 按照加入等待列表的顺序逐一被唤醒，从而避免了因为唤醒顺序不确定而导致的“饥饿”等问题。

在 Go 的运行时系统中，`notifyList` 经常被用来实现对 Channel 的同步和等待操作，以此实现 Goroutine 之间的通讯和协作。



## Functions:

### rootFor

`rootFor`函数是Go运行时系统中用于锁定和解锁semagore的函数之一。它的作用是查找一个可用（或新的）semRoot对象，并返回它的指针。如果找到了可用的对象，则使用该对象，并将其locked状态设置为true。如果没有可用对象，则创建一个新的对象并将其locked状态设置为true，然后将其返回。

semRoot是一个包含信号量的结构体，用于控制并发访问。当多个Goroutine需要访问共享资源时，这些资源可以使用semaphore进行保护，以确保同一时间只有一个Goroutine可以访问资源。semaphore是信号量的一种实现方式，它具有计数器和等待队列，其中计数器用于跟踪当前可以访问共享资源的Goroutine数量，并控制等待队列中的Goroutine何时可以获取访问权限。

在Go运行时系统中，semaRoot是一个包含信号量计数器和等待队列的结构体。它允许Goroutine在需要访问共享资源时获取锁，并在完成后释放锁。在具体的实现中，Go运行时系统为每个semRoot对象创建一个等待队列，该队列包含在semRoot对象被锁定之前等待访问共享资源的Goroutine。

要使用semaphore进行锁定和解锁操作，必须先获得semRoot对象的指针。这就是rootFor函数的作用。它通过查找已锁定或可用的semRoot对象来返回semRoot对象的指针。如果找到了可用对象，则该数字将自增。

总之，rootFor函数是Go运行时系统中锁定并发访问时的关键函数之一。它根据当前并发访问的情况来查找或创建一个semRoot对象，并返回其指针，使得Goroutine可以对共享资源进行安全的并发访问。



### sync_runtime_Semacquire

这个函数的作用是获取一个信号量，用于同步并发访问共享资源。

在Go语言中，多个goroutine可以同时执行，共享同一进程的资源。如果这些goroutine同时访问一个共享的资源，就会发生竞争条件（race condition），导致程序出错或产生意料之外的结果。

为了避免竞争条件，需要在多个goroutine之间进行同步。在Go语言中，可以使用信号量（Semaphore）来实现同步。

这里的sync_runtime_Semacquire函数就是用于获取一个信号量的。具体来说，该函数会首先判断信号量计数器是否大于0，如果大于0，则将信号量计数器减1，并返回true，表示获取信号量成功；如果计数器为0，则将当前goroutine加入到等待队列，并进入阻塞状态，直到有其他goroutine释放了信号量后，再将该goroutine唤醒，并返回true。如果出现了错误，则返回false。

需要注意的是，该函数会使用一个spinlock来保护信号量，以避免锁竞争的问题。当spinlock被锁定时，该函数会忙等待直到spinlock释放锁，然后再获取信号量。

总之，sync_runtime_Semacquire函数是Go语言中实现同步的关键部分，用于保护共享资源的访问安全。



### poll_runtime_Semacquire

poll_runtime_Semacquire函数是Go语言运行时调度系统的一部分。它的作用是在没有可用的信号量的情况下，将Goroutine放入等待队列中并切换到其他Goroutine，以避免出现线程阻塞的情况。Semacquire函数会尝试在有信号量可用时立即获取锁，如果获取不到，则使用park函数挂起当前的Goroutine，并将其添加到等待队列中。在有新的信号量可用时，又会使用semaWakeup函数将挂起的Goroutine唤醒。

具体来说，poll_runtime_Semacquire函数的实现如下：

1. 初始化一个锁

2. 使用semacount获取当前信号量计数器的值

3. 如果当前信号量计数器不为0，则使用semacntdecrement将信号量计数器减一，并返回

4. 如果当前信号量计数器为0，则将当前Goroutine封装成一个sudog对象，并将其添加到等待队列中

5. 然后使用park函数将当前Goroutine挂起

6. 当有新的信号量可用时，由于信号量计数器已经减过了，因此可以使用semacountdecrement将计数器减一，并使用semaWakeup将等待队列中的Goroutine唤醒

7. 被唤醒的Goroutine重新执行，并使用semacntincrement将信号量计数器加一

需要注意的是，在等待队列中的Goroutine被唤醒时并不会立即获取锁，而是要再次尝试获取锁，如果获取不到，则又会被重新放入等待队列中。这个过程可能会重复很多次，因此poll_runtime_Semacquire实际上是一个自旋锁实现。

总的来说，poll_runtime_Semacquire函数的作用是实现了一种高效的Goroutine等待和唤醒机制，使得程序可以更好地利用CPU资源。



### sync_runtime_Semrelease

sync_runtime_Semrelease是在Go语言运行时中实现的一个信号量(semaphore)释放函数。信号量是用于控制并发访问的一种机制，可以用来限制同时访问某个资源的数量。

此函数用于释放信号量，唤醒其他等待信号量的goroutine。具体来说，它会检查等待队列中是否有正在等待信号量的goroutine，如果有，则唤醒最先等待的goroutine并从等待队列中移除，否则，将信号量的计数值加1。

在Go语言中，sync.Mutex和sync.WaitGroup等标准库包中都是基于信号量机制实现的，因此，这个函数也会被相关的标准库包使用。例如，当调用sync.WaitGroup.Wait()方法时，如果计数器的值为0，该方法会立即返回；否则，它会等待计数器归零，而计数器归零的关键在于调用了Semrelease()函数。

总之，sync_runtime_Semrelease()函数是在Go语言运行时中用于实现信号量机制的重要函数，用于控制并发访问和同步等场景。



### sync_runtime_SemacquireMutex

sync_runtime_SemacquireMutex是Go语言中实现的一种同步机制，用于在协程之间传递访问资源的所有权。

具体来说，它的作用是获取一个互斥锁，如果当前该锁被其他协程所持有，则会将当前协程阻塞在该位置等待锁的释放。当锁被释放后，阻塞在该位置的协程会唤醒并继续执行。

在实现上，sync_runtime_SemacquireMutex的具体实现方式可以依赖于不同的操作系统平台，但是其核心思想均是通过操作系统提供的原子操作来实现并发访问资源的互斥。

需要注意的是，由于该同步机制会导致协程的阻塞，因此在使用时需要谨慎并避免出现死锁等问题。



### sync_runtime_SemacquireRWMutexR

func sync_runtime_SemacquireRWMutexR(mutex *sync.RWMutex, read bool, blockprof *uint32) {
	// Grab the readers lock
	if read {
		mutex.RLock()
		return
	}

	// Acquire the writers semaphore, spinning until successful
	if blockprof != nil {
		block((*int32)(blockprof), func() { mutex.Lock() })
	} else {
		mutex.Lock()
	}
}

这个函数的作用是实现读写锁RWMutex的加锁操作，具体代码实现如下：

如果是读取锁，则直接使用RWMutex的RLock()方法获取读锁；如果是写锁，则使用Mutex的Lock()方法获取写锁。

在获取写锁时，如果blockprof参数不为nil，则使用block()进行阻塞等待，同时记录阻塞的时间。如果blockprof为nil，则直接获取写锁。

总的来说，这个函数是用来实现读写锁RWMutex的加锁操作，保证多线程下对共享数据的访问安全。



### sync_runtime_SemacquireRWMutex

sync_runtime_SemacquireRWMutex是Go语言中读写锁(RWMutex)的内部实现函数之一，它的作用是获取RWMutex的读锁或写锁，用于控制对共享资源的访问。具体来说，如果当前没有任何协程持有锁，它会立即获取锁并返回true；否则，它会将当前协程阻塞在一个等待队列中，直到有其他协程释放了锁。

在具体代码实现中，sync_runtime_SemacquireRWMutex会调用runtime_Semacquire函数实现锁的获取。读写锁在Go语言标准库中是通过两个计数器实现的，一个计数器记录当前持有写锁的协程数量(最多只能为1)，另一个计数器记录当前持有读锁的协程数量。当某个协程请求获取锁时，函数会根据锁的类型(读锁/写锁)和当前锁的状态(空闲/被占用)来决定是否能够获取锁。如果当前有写锁被持有，请求读锁的协程会被加入等待队列；如果当前有任何锁被持有，请求写锁的协程会被加入等待队列。

总之，sync_runtime_SemacquireRWMutex是一个非常重要的函数，Go语言中的RWMutex的正确性和性能都依赖于它的正确实现。



### poll_runtime_Semrelease

poll_runtime_Semrelease函数是runtime包中用于在内部信号量(semaphore)计数器被释放时将一个或多个goroutine放到可运行状态的函数。

该函数是内部函数，只在runtime包中使用。它与semacquire函数一起使用来提供一种协调goroutine的机制。semacquire函数用于等待并获取内部信号量，而poll_runtime_Semrelease函数用于在信号量被释放时将等待在该信号量上的goroutine放置在可运行状态。

该函数的作用是遍历所有对该信号量进行等待的goroutine，并将它们添加到全局运行队列中。它还会更新关于信号量计数的状态。

poll_runtime_Semrelease函数的主要逻辑如下：

1.从等待信号量的队列中取出一个或多个goroutine。

2.将goroutine添加到全局运行队列中。

3.更新内部信号量的计数器状态。

当某个goroutine释放信号量时，该函数会遍历所有等待该信号量的goroutine，并将它们添加到全局运行队列中，使它们可以运行。

poll_runtime_Semrelease函数是内部实现的一部分，普通的go开发者通常不需要直接使用它。



### readyWithTime

readyWithTime是Go语言运行时系统中用于实现信号量调度的核心函数之一。

在Go语言中，信号量通常用于控制并发访问共享资源的情况。一个信号量包含一个计数器和一组等待队列，计数器用于记录可用的资源数量，等待队列用于存储等待该资源的协程列表。当一个协程需要访问共享资源时，如果此时信号量计数器为0，则该协程将被阻塞并放入等待队列中，直到计数器大于0时才能被唤醒继续执行。

readyWithTime函数的作用就是将等待队列中的协程唤醒，并将其放入调度器的就绪队列中等待调度。它接收一个等待队列的指针作为参数，遍历队列中的所有协程，并将它们分别加入到调度器的就绪队列中。如果协程需要设置超时时间，则readyWithTime还会根据超时时间设置一个定时器，并将定时器加入到调度器的计时器堆中，以便在超时时间到达时唤醒协程。

具体来说，readyWithTime函数的主要步骤包括：

1. 遍历等待队列中的所有协程，并将它们依次从队列中取出。

2. 如果协程设置了超时时间，则创建一个定时器，并将协程和定时器关联起来，以便在超时时间到达时唤醒协程。

3. 将协程加入调度器的就绪队列中，等待被调度执行。

4. 重复执行步骤1-3，直到等待队列为空为止。

由于信号量调度在Go语言的并发编程中极为常见，在Go语言运行时系统中，readyWithTime函数被广泛使用，例如在channel、mutex和cond等类型中都有涉及。通过调用readyWithTime函数，运行时系统可以很好地实现协程的调度和等待，保证程序能够正确地访问共享资源并避免竞态条件和死锁等问题。



### semacquire

semacquire函数是Golang运行时中的一个同步原语，用于实现对共享资源的并发访问控制。

在多线程环境中，多个线程需要访问同一个资源，为了避免竞争条件和其他并发问题，需要限制只有一个线程可以访问该资源。semacquire函数就是用来获取访问资源的锁并阻塞等待锁释放的函数。

semacquire函数的主要作用是通过获取一个信号量来实现同步。信号量是一个整数，表示可用的资源数量。当信号量为0时，线程会被阻塞等待资源可用。

具体来说，semacquire函数会首先检查当前是否已经获取到了锁。如果已经获取到了锁，就会直接返回；否则，会先增加锁的等待计数器，然后尝试获取锁。

如果获取锁失败，说明当前资源正在被其他线程访问，semacquire函数会将线程加入等待队列，并执行调度器从队列中选择一个新的线程来运行。当其他线程释放了锁，semacquire函数会重新尝试获取锁，并顺利返回。

总的来说，semacquire函数是Golang运行时中非常重要的一个同步机制，用于保证并发访问共享资源的正确性和可靠性。



### semacquire1

semacquire1这个func的作用是在并发编程中控制上锁和解锁共享资源的同步方法。

在Go语言中，通过加上sync.Mutex类型的锁来实现并发控制，但是在多线程的情况下，无法保证每个线程都有机会解锁。这时候就需要一个类似于信号量（Semaphore）的机制，保证同一时间只有一个线程的锁能够生效，其他线程只能等待。

semacquire1函数就是一个实现这个机制的函数。它通过判断锁是否处于被占用状态，如果是，则将当前协程加入该锁的等待队列中。直到锁被释放为止。

在实现上，semacquire1函数使用了一些汇编代码，来保证并发安全性和性能，同时也利用管道进行状态传递。用法如下：

```
func (s *Sema) Acquire() {
    // ...
    semacquire1(&s.sema)
    // ...
}

func (s *Sema) Release() {
    // ...
    semrelease(&s.sema)
    // ...
}
```

其中，Acquire是加锁方法，Release是解锁方法。

总之，semacquire1函数是Go语言中实现并发控制机制的核心函数之一。



### semrelease

semrelease函数是Go调度器中的一个关键函数，它的作用是释放信号量。

在Go运行时中，信号量用于控制系统资源的访问和共享。例如，当一个goroutine需要使用一个共享的系统资源（例如内存、文件等）时，它需要获取信号量来保证其他goroutine不会同时访问这个资源。当这个goroutine没有再使用这个资源时，就需要释放这个信号量让其他goroutine使用它。

semrelease函数的功能就是实现信号量的释放。它会将信号量的计数器加1，如果有在等待信号量的goroutine，则会唤醒其中的一个goroutine。如果没有等待的goroutine，那么这个信号量就可以被其他的goroutine再次获取，并且在获取的时候不需要进入等待状态。

需要注意的是，semrelease函数并不是一个普通的函数，它是在Go调度器内部被调用的。因此，它的具体实现可能会因不同的系统环境、硬件架构等因素而有所不同。在不同的系统环境下，Go调度器也会采用不同的信号量实现方式，以保证其在不同平台上的正确性和高效性。



### semrelease1

semrelease1是一个函数，它用于释放锁定的信号量。

在Go语言中，信号量是一个用于同步和互斥访问的计数器。当一个进程或线程需要访问一个共享资源时，它必须获取该资源的信号量锁，以确保其他进程或线程不能同时访问该资源。

在semrelease1函数中，它首先会从当前goroutine获取锁定的信号量，并检查是否存在多余的信号量，如果有多余的信号量，则表示当前goroutine没有完全释放锁定的资源，并增加可用信号量的计数。

接下来，如果存在等待该锁的goroutine，则会将其中一个goroutine唤醒并将锁定的信号量传递给它。如果没有等待的goroutine，则将该信号量标记为可用并从等待列表中移除。

最后，semrelease1函数会解锁操作系统中的信号量，并返回。

因此，semrelease1函数的作用是确保对共享资源的互斥访问，并协调所有等待该资源的goroutine，以避免竞争条件和死锁的发生。



### cansemacquire

在 Go 语言的 runtime 包中的 sema.go 文件中，cansemacquire 函数用于判断是否可以获取信号量。

在 Go 中，信号量是用于协调并发访问共享资源的机制。它可以控制同时访问共享资源的协程数量。在 runtime 包中，cansemacquire 函数被用于协程调度和同步过程中，具体作用如下：

1. 限制并发访问：cansemacquire 函数用于限制并发访问共享资源的数量。它通过检查当前可用的信号量数量来确定是否允许协程获取信号量。如果可用信号量数量大于零，则允许获取；否则，需要等待其他协程释放信号量后才能获取。

2. 阻塞和唤醒：如果 cansemacquire 函数判断当前无法获取信号量（即可用信号量数量为零），则会将当前协程阻塞，使其进入等待状态，直到其他协程释放信号量并唤醒等待的协程。

3. 协程调度：cansemacquire 函数在协程调度过程中发挥重要作用。当一个协程需要获取信号量时，如果无法立即获取，它会暂时释放处理器，让其他协程有机会执行。这种协程的暂停和恢复是通过 cansemacquire 函数来实现的。

总结来说，cansemacquire 函数在 Go 语言的运行时系统中负责控制协程对信号量的获取和释放，通过它实现了并发访问的同步和调度机制，以保证共享资源的正确性和性能。



### queue

在 Go 语言的 runtime/sema.go 文件中，queue 函数用于将 goroutine（轻量级线程）添加到信号量等待队列中。下面是该函数的源代码：

```go
// queue adds the calling goroutine to the semaphore queue.
// It reports whether the goroutine must block.
// It must not be called with g.m.p == 0; use acquire instead.
//
//go:linkname queue sync.runtime_SemacquireMutex
func queue(sem *uint32, lifo bool) bool {
	if *sem > 0 {
		throw("queue: *sem > 0")
	}
	if getg().m.locks != 0 {
		throw("queue: holding locks")
	}
	mp := acquirem() // 获取当前执行的 goroutine 所在的 M（机器线程）
	semrelease1(*sem, false)
	// Add ourselves to nwait to disable "easy" wake-up.
	old := *sem
	// 由于这里是一些原子操作，具体流程可能会根据 Go 版本和编译器实现略有不同
	new := old + 1
	if lifo {
		// Add to head of queue
		new = old + 2
	}
	for {
		if atomic.Cas(sem, old, new) {
			break
		}
		old = *sem
		if lifo {
			// Add to head of queue
			new = old + 2
		}
		if old&1 != 0 {
			// sem is actively being acquired by a
			// different goroutine, so queue ourselves
			// so that semrelease1 can wake us up.
			 // 将当前 goroutine 添加到 semaRoot 的等待队列中
			lock(&semaRoot)
			gp := getg()
			gp.schedlink = (*sudog)(unsafe.Pointer(semaRoot.waiting))
			semaRoot.waiting = uintptr(unsafe.Pointer(gp))
			unlock(&semaRoot)
			goparkunlock(&mp.lock, waitReasonSema, traceEvGoBlockSema, 3) // 将当前 goroutine 阻塞
			lock(&semaRoot)
			gp = getg()
			// remove from semaRoot.waiting - it's not
			// running so it won't get lost.
			for pp := &semaRoot.waiting; ; pp = &(*sudog)(unsafe.Pointer(pp)).schedlink {
				sgp := (*sudog)(unsafe.Pointer(*pp))
				if sgp == gp {
					*pp = uintptr(unsafe.Pointer(sgp.schedlink))
					break
				}
			}
			unlock(&semaRoot)
			return false
		}
		new = old + 1
	}
	// 能够到达这里，说明当前 goroutine 成功获取了 sema
	// 执行下面的代码之前，应该先调用 semacquire 记录当前状态
	mp.mcache.sema++
	releasem(mp) // 释放当前 goroutine 所在的 M
	return true
}
```
queue 函数的主要作用是将当前的 goroutine 添加到信号量（semaphore）的等待队列中，并判断当前 goroutine 是否需要阻塞。下面是该函数的详细解释：

1. 首先，函数会检查传入的信号量 sem 的值是否大于 0。如果大于 0，则表示存在问题，因为只有在信号量为 0 时，才需要将 goroutine 添加到等待队列中。

2. 接下来，函数会检查当前 goroutine 是否持有其他锁。如果当前 goroutine 持有锁，则会报错，因为不允许在持有锁的情况下调用 queue 函数。

3. 然后，函数会调用 acquirem 函数获取当前 goroutine 所在的 M（机器线程），并调用 semrelease1 函数对信号量进行释放操作。semrelease1 函数用于释放信号量，并唤醒等待队列中的其他 goroutine。

4. 在将当前 goroutine 添加到等待队列之前，函数会先将当前信号量的值保存在 old 变量中，并计算出新的信号量值 new。如果需要使用 LIFO（后进先出）方式管理等待队列，则 new 的计算会有所不同。

5. 接下来，函数会使用 atomic.Cas 函数尝试原子地将 new 的值赋给信号量 sem。如果赋值成功，表示当前 goroutine 成功将自己添加到了等待队列，并且信号量的值已更新。此时函数会继续执行后续的代码。

6. 如果 atomic.Cas 失败，表示在当前 goroutine 尝试将自己添加到等待队列的过程中，有其他 goroutine 正在活动地获取信号量。这时，当前 goroutine 需要将自己添加到信号量的等待队列中，并阻塞自己。

7. 在添加到等待队列之前，函数会先获取 semaRoot 的锁，然后将当前 goroutine 添加到等待队列的头部。之后，函数会使用 goparkunlock 函数将当前 goroutine 阻塞。阻塞期间，该 goroutine 将不会占用 M，这意味着 M 可以用于执行其他 goroutine。

8. 当其他 goroutine 在释放信号量时，会尝试唤醒等待队列中的 goroutine。被唤醒的 goroutine 会重新获取 semaRoot 的锁，并从等待队列中移除自己。

9. 最后，函数会对获取到信号量的 goroutine 进行一些处理，例如增加其关联的 M 的 sema 值，并释放当前 goroutine 所在的 M。

总结来说，queue 函数的作用是将当前 goroutine 添加到信号量的等待队列中，并根据情况决定是否需要阻塞当前 goroutine。该函数是 Go 运行时中信号量实现的核心部分，用于实现协程之间的同步和互斥操作。



### dequeue

在 Go 语言的 runtime 包中，`sema.go` 文件中的 `dequeue` 函数是用于从一个 goroutine 队列中移除并返回队列中的第一个 goroutine。

`dequeue` 函数主要用于实现 Go 语言的调度器中的信号量（sema）机制。信号量机制用于控制并发执行的 goroutine 数量，它可以限制同时执行的 goroutine 数目，以防止资源竞争和过度的并发。

下面是 `dequeue` 函数的主要作用和行为的详细解释：

1. 队列检查：函数首先检查传入的队列是否为空。如果队列为空，表示没有可用的 goroutine，函数将返回 `nil`。
2. 队首元素移除：如果队列不为空，函数将从队列中移除第一个元素（即第一个 goroutine）。
3. 队首元素状态检查：在移除队首元素之后，函数会检查该 goroutine 的状态。这是因为在并发环境下，队列中的 goroutine 可能已经被其他部分取消或者从队列中移除。
4. 队首元素的下一个状态：如果队首元素的状态是 `goSched`，表示该 goroutine 准备就绪并可以执行，函数将返回该 goroutine。
5. 队首元素的下一个状态：如果队首元素的状态是 `goNil`，表示该 goroutine已经取消或者从队列中移除，函数将继续从队列中寻找下一个可用的 goroutine。
6. 队首元素的下一个状态：如果队首元素的状态是其他状态（例如 `goSleep`），表示该 goroutine 正在等待某个条件满足，函数将继续从队列中寻找下一个可用的 goroutine。

总之，`dequeue` 函数在 Go 语言的 runtime 调度器中起到了从队列中移除并返回队列中的第一个 goroutine 的作用。它是调度器中的关键部分，用于决定哪个 goroutine 将被执行，并确保并发的合理调度。



### rotateLeft

在 Go 语言的 runtime 包中的 sema.go 文件中，rotateLeft 函数用于将无符号整数（uint）按位循环左移（rotate left）指定的位数。让我们更详细地解释一下它的作用和实现原理。

函数签名如下：
```go
func rotateLeft(x uint, k uint) uint
```

rotateLeft 函数将无符号整数 x 按位循环左移 k 位，并返回结果。具体来说，它将 x 的二进制表示向左循环移动 k 位。循环移动意味着超过整数长度的位将从右侧重新进入左侧，保持整数的位数不变。

以下是 rotateLeft 函数的实现原理：

1. 首先，检查 k 是否等于 0。如果是，表示无需移动，直接返回原始值 x。
2. 接下来，获取 x 的位数（bitSize），通常为 32 或 64 位，取决于运行时环境。
3. 通过位运算，将 k 的值限制在位数范围内（k %= bitSize）。这是为了避免移位操作超出整数长度的情况。
4. 使用位运算左移操作符 (<<) 将 x 向左移动 k 位。这会将位从左侧移出并进入右侧。但是，这仍然可能导致超出位数范围的问题。
5. 使用位运算或操作符（|）将超出位数范围的位从右侧重新进入左侧。这是通过将右侧的位与左侧位进行逻辑或运算实现的。
6. 返回结果。

rotateLeft 函数的目的是实现一种高效的按位循环左移操作，而不需要借助其他辅助变量或循环。这在一些并发和同步算法中可能会被使用，例如在信号量（semaphore）的实现中。

以下是 rotateLeft 函数的示例代码：
```go
func rotateLeft(x uint, k uint) uint {
	if k == 0 {
		return x
	}
	bitSize := uint(unsafe.Sizeof(x)) * 8
	k %= bitSize
	return (x << k) | (x >> (bitSize - k))
}
```

请注意，此函数的实现依赖于具体的运行时环境，其中整数的位数可能是 32 位或 64 位。此外，rotateLeft 函数在 runtime 包的内部使用，并不直接暴露给一般的 Go 语言开发者使用。



### rotateRight

在 Go 语言的 runtime 包中，sema.go 文件中的 rotateRight 函数用于实现循环右移操作。这个函数的作用是将一个无符号整数（uint）的位向右循环移动指定的位数。

下面是 rotateRight 函数的源代码（截取自 Go 1.17 版本的 runtime/sema.go）：

```go
// rotateRight returns x rotated right by k bits.
func rotateRight(x uint, k unsigned) uint {
    return (x >> k) | (x << (wordSize - k))
}
```

这个函数的实现使用了位操作符（bitwise operators）来完成循环右移操作。具体解释如下：

1. 函数接收两个参数：x 为要进行循环右移的无符号整数，k 为循环移动的位数。
2. 首先，通过右移操作 `(x >> k)`，将 x 的二进制表示向右移动 k 位。这会使原本位于低位的 k 位移到高位，被丢弃的高位空出来。
3. 然后，通过左移操作 `(x << (wordSize - k))`，将 x 的二进制表示向左移动 wordSize-k 位。这会将被丢弃的高位移动到低位，填补右移操作中空出来的高位。
4. 最后，使用位或操作符 `(x >> k) | (x << (wordSize - k))`，将两个移位操作的结果合并，得到最终的循环右移结果。

需要注意的是，rotateRight 函数中的 `wordSize` 是一个常量，表示机器字（machine word）的位数。它的值根据不同的操作系统和架构而变化，例如在 64 位系统上为 64。

循环右移操作在某些算法和数据结构中非常有用，例如循环队列、哈希函数等。通过将位从右边循环移动到左边，可以实现循环的效果，使得数据在某个固定大小的空间中循环使用。rotateRight 函数提供了一种方便且高效的方式来执行这种操作。



### less


在 Go 语言的 `runtime/sema.go` 文件中，`less` 函数是用于计算信号量中的等待者优先级的函数。下面我将详细解释该函数的作用和实现。

首先，为了更好地理解 `less` 函数的作用，让我们回顾一下在 Go 语言中的信号量（Semaphore）的概念。信号量是一种用于协调共享资源访问的同步原语。它维护了一个计数器，该计数器表示可用资源的数量。当一个协程需要访问资源时，它会尝试获取信号量。如果计数器大于零，协程可以获取资源并将计数器减一；否则，协程将被阻塞，直到有可用的资源。

在 Go 语言的 `runtime` 包中，`sema.go` 文件实现了一个非常基础的信号量类型 `sema`，它用于调度协程的执行。`sema` 结构体中有一个 `int64` 类型的计数器字段 `n`，表示可用资源的数量。另外，还有一个保存等待者的双向链表 `waiters`，其中每个等待者都包含一个指向 `g`（goroutine）的指针和一个 `int64` 类型的 `deadline` 字段。

现在我们来看一下 `less` 函数的实现。`less` 函数的定义如下：

```go
func less(a, b waitlinkptr) bool {
    // ...
}
```

该函数接收两个 `waitlinkptr` 类型的参数 `a` 和 `b`，它们都是等待者链表中的节点。`waitlinkptr` 实际上是一个指向等待者节点的指针。

`less` 函数的作用是比较两个等待者的优先级，以确定它们在链表中的顺序。具体来说，它通过比较 `a` 和 `b` 的 `deadline` 字段的值来判断哪个等待者的优先级更高。较早到期的等待者被认为是优先级更高的，因此应该在链表中处于较前的位置。

函数的实现相对简单，它比较了 `a` 和 `b` 的 `deadline` 字段，并返回一个布尔值，指示是否 `a` 的优先级较高。比较操作符 `==` 和 `<` 在这里起到了关键的作用。

总结起来，`less` 函数在 `runtime/sema.go` 文件中用于确定等待者在信号量链表中的优先级顺序。通过比较等待者的 `deadline` 字段，它决定哪个等待者的优先级更高。这对于协程的调度和资源的分配非常重要。


### notifyListAdd


`notifyListAdd` 是 Go 语言中 `sync` 包中的一个函数，位于 `sema.go` 文件中。该函数的作用是向一个通知列表（notify list）中添加一个新的通知项（notification entry）。

在 Go 语言的运行时（runtime）中，`notifyList` 是用于实现同步原语的数据结构之一。它通常用于实现类似于条件变量（condition variable）的功能。

通知列表（notify list）是一个链表，用于保存等待某个特定事件发生的一组 goroutine。每个通知项（notification entry）表示一个等待事件的 goroutine，并包含了与该 goroutine 相关的一些信息。

`notifyListAdd` 函数的具体作用如下：

1. 创建一个新的通知项（notification entry），其中包含了等待事件的 goroutine 和与该 goroutine 相关的一些信息。
2. 将新的通知项添加到通知列表的末尾，成为最新的等待项。

通常，在并发编程中，当一个或多个 goroutine 需要等待某个条件满足时，它们可以将自己添加到一个通知列表中。当条件满足时，可以使用通知列表来唤醒这些等待的 goroutine，以便它们可以继续执行。

`notifyListAdd` 函数是 `sync` 包中用于管理通知列表的内部函数，它在运行时中起到了管理和组织等待事件 goroutine 的作用，确保它们能够正确地被唤醒和处理。但是，请注意，由于 `notifyList` 是运行时的内部数据结构，因此该函数的具体实现和细节可能会因版本和平台而有所不同。


### notifyListWait


在 Go 语言的运行时（runtime）包中，`sema.go` 文件包含了用于实现信号量（Semaphore）的相关代码。其中的 `notifyListWait` 函数用于等待通知列表（notify list）中的通知。

通知列表是一种用于在并发环境中等待通知的数据结构。它通常与条件变量（condition variable）一起使用，用于协调不同的 goroutine（Go 协程）之间的操作。

具体来说，`notifyListWait` 函数的作用是将当前 goroutine 添加到通知列表中，并等待被通知。下面是 `notifyListWait` 函数的详细解释：

1. 首先，它会创建一个 `notifyList` 类型的变量 `l`，该类型是一个双向链表，用于存储等待通知的 goroutine。
2. 然后，它会将当前的 goroutine 封装在一个 `waitNode` 结构体中，该结构体包含了当前 goroutine 的相关信息。
3. 接下来，它会将 `waitNode` 添加到 `notifyList` 中，通过调用 `l.add(&wait)` 实现。这样，当前 goroutine 就被添加到了等待通知的列表中。
4. 接下来，当前 goroutine 会进入休眠状态，等待被通知。它会调用 `runtime.goparkunlock` 函数，将当前 goroutine 进行休眠并解锁当前的互斥锁（如果存在）。
5. 当其他 goroutine 调用通知列表的通知函数（`notifyListNotify` 或 `notifyListNotifyAll`）时，会遍历通知列表中的所有等待 goroutine，并唤醒它们。
6. 被唤醒的 goroutine 会从休眠状态返回，并继续执行之前被休眠的代码。

总结来说，`notifyListWait` 函数的作用是将当前 goroutine 添加到通知列表中，并进入休眠状态，等待被其他 goroutine 通知唤醒。它在协调并发操作和实现同步等待的场景中起着重要的作用。


### notifyListNotifyAll


在 Go 语言的 runtime 包中，`sema.go` 文件中的 `notifyListNotifyAll` 函数用于通知一个等待列表中的所有等待者（waiter）。让我们详细解释一下该函数的作用。

在 Go 的并发编程中，经常需要使用等待组（wait group）或条件变量（condition variable）来同步协程的执行。`notifyListNotifyAll` 函数实现了一种条件变量的功能，用于通知所有正在等待的协程。

`notifyListNotifyAll` 函数的主要作用是唤醒所有在等待列表中的等待者。等待列表是一个链表，其中包含等待该条件变量的协程。当某个条件满足时，通过调用 `notifyListNotifyAll` 函数，可以同时唤醒所有在等待列表中的协程，使它们继续执行。

以下是 `notifyListNotifyAll` 函数的伪代码表示：

```go
func notifyListNotifyAll(l *notifyList) {
    for l.wait != 0 {
        // 唤醒一个等待者
        notifyListNotifyOne(l)
    }
}
```

该函数通过循环调用 `notifyListNotifyOne` 函数来逐个唤醒等待列表中的等待者。在每次循环中，会从等待列表中选择一个等待者进行唤醒。这样，所有在等待列表中的等待者都会被唤醒，并可以继续执行它们的任务。

需要注意的是，该函数在内部调用了 `notifyListNotifyOne` 函数，它是 `sema.go` 文件中的另一个函数，用于唤醒单个等待者。

总结起来，`notifyListNotifyAll` 函数用于在 Go 语言的并发编程中实现条件变量的功能，通过唤醒等待列表中的所有等待者来通知它们某个条件已满足，从而使它们可以继续执行。


### notifyListNotifyOne

`sema.go` 文件中的 `notifyListNotifyOne` 函数是 Go 语言中用于通知等待中的 goroutine 的函数之一。它是在 Go 语言的调度器中用于同步和调度 goroutine 的一部分。

在 Go 语言中，`notifyListNotifyOne` 函数用于唤醒等待队列中的一个 goroutine。它的作用是通过通知等待队列中的一个 goroutine，告诉它可以继续执行。该函数的详细解释如下：

1. `notifyListNotifyOne` 函数首先会检查等待队列是否为空。如果等待队列为空，则没有需要唤醒的 goroutine，函数将直接返回。

2. 如果等待队列不为空，函数会选择一个等待中的 goroutine（通常是队列中的第一个）并从等待队列中移除它。

3. 接下来，函数会将该 goroutine 的状态设置为可运行，并将其放入可运行队列中，以便调度器可以在适当的时候执行该 goroutine。

4. 通过上述步骤，`notifyListNotifyOne` 函数成功地唤醒了一个等待中的 goroutine，并使其可以继续执行。

这个函数通常与其他的同步原语（如锁、条件变量等）一起使用，以实现对共享资源的安全访问和协调。当某个条件满足时，等待中的 goroutine 会通过调用 `notifyListNotifyOne` 函数来被唤醒，从而继续执行后续的操作。

需要注意的是，上述解释是基于 Go 语言运行时源代码中 `sema.go` 文件的通用逻辑。具体的使用和含义可能会因为不同的上下文而有所不同。因此，具体的使用方式和含义还需要根据您所查看的具体代码和文档来确定。



### notifyListCheck

在 Go 语言的运行时（runtime）包中，`sema.go` 文件定义了 Go 语言中的信号量实现。其中，`notifyListCheck` 函数用于检查并更新一个通知列表（`notifyList`）中等待的 Goroutine（Go 协程）的状态。

下面是 `notifyListCheck` 函数的详细解释：

1. `notifyListCheck` 函数的主要目的是扫描通知列表中等待的 Goroutine，检查它们是否可以继续执行。
2. 该函数被用于实现 Go 语言中的锁、条件变量等同步原语。它用于确保 Goroutine 在等待特定事件时能够正确地被唤醒和执行。
3. 当某个事件发生并满足特定条件时，`notifyListCheck` 函数会遍历通知列表，并将满足条件的 Goroutine 的状态更新为可执行状态，以便它们能够继续执行。
4. 在遍历通知列表期间，函数会检查每个 Goroutine 的状态并根据条件进行更新。具体的更新包括将 Goroutine 的状态从等待状态（waiting）更新为可运行状态（runnable）。
5. 更新状态后，`notifyListCheck` 函数可能还会执行一些额外的操作，例如将 Goroutine 添加到调度器的运行队列中，以确保它们能够被调度并执行。
6. 通常情况下，`notifyListCheck` 函数在某个事件发生后由调度器或相关同步原语调用，以唤醒等待的 Goroutine。

总而言之，`notifyListCheck` 函数在 Go 语言的运行时中扮演着关键的角色，负责管理 Goroutine 的状态，并确保它们能够正确地被唤醒和执行。它是实现同步原语的重要组成部分，用于保证并发程序的正确性和可靠性。



### sync_nanotime


在Go语言的runtime包中的sema.go文件中，sync_nanotime函数是用于获取当前的纳秒级时间戳的函数。这个函数使用了与操作系统相关的底层调用来获取高精度的时间戳。

sync_nanotime函数的作用是提供一个相对可靠的时间源，用于在同步原语（synchronization primitives）中进行时间相关的操作，如超时控制、定时器等。在并发编程中，时间是一种重要的资源，它被用于实现各种同步和调度机制。

具体来说，sync_nanotime函数返回一个int64类型的纳秒级时间戳，表示从某个参考时间点到当前时刻所经过的纳秒数。这个时间戳通常用于计算时间间隔、判断是否超时以及限制某些操作的执行时间。

由于不同的操作系统可能提供不同的时间获取方法，并且不同的架构可能有不同的时间计数器，sync_nanotime函数在实现上会根据操作系统和硬件架构选择最适合的方式来获取时间戳。这确保了在不同的平台上，程序都可以获取到高精度的时间戳。

总结起来，sync_nanotime函数在Go语言的runtime中提供了一个可靠的方法来获取纳秒级时间戳，为同步原语提供了时间相关的操作支持，从而在并发编程中实现超时控制、定时器等功能。


