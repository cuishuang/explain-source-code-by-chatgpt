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





### queue





### dequeue





### rotateLeft





### rotateRight





### less





### notifyListAdd





### notifyListWait





### notifyListNotifyAll





### notifyListNotifyOne





### notifyListCheck





### sync_nanotime





