# File: runtime2_lockrank.go

runtime2_lockrank.go文件是Go语言sync（同步）包中的文件之一，其主要作用是为锁（mutex）提供了一个级别排序功能。

在并发编程中，锁是用于同步访问共享资源的一种机制。在代码中，锁可以控制对共享资源的访问，从而避免多个并发执行的线程访问同一资源导致的数据竞争问题。

然而，在使用锁的时候，可能会遇到多个锁相互依赖的情况。例如，如果某个线程已经持有了锁A，但需要等待另一个线程释放锁B后才能继续执行。这种情况可能会导致死锁问题。

为了解决这个问题，Go语言的sync包中提供了LockRank类型，可以通过对LockRank进行排序，避免出现循环等待的情况，从而避免死锁的发生。

该文件中定义了一个LockRank类型，用于实现锁的级别排序。LockRank可以与标准库中的互斥锁（Mutex）、读写锁（RWMutex）等配合使用。LockRank中提供了Lock和Unlock方法，用于获取或释放锁。在获取锁时，会根据LockRank的级别进行排序，从而避免出现死锁的情况。

总之，runtime2_lockrank.go 文件的作用是提供一个级别排序的锁实现，以避免在使用多个锁时出现死锁问题。




---

### Structs:

### notifyList

notifyList是一个用于等待通知的链表结构体。它被用于实现锁的等待和唤醒功能。每个等待锁的goroutine都会在notifyList上等待锁被释放，并被锁持有者通知唤醒。当锁释放时，锁持有者会把waiter从notifyList中取出，并将其唤醒。

notifyList结构体的定义如下：

```
type notifyList struct {
    wait   uint32   // 等待的 goroutine 数量
    notify uint32   // 通知的次数，每次 Lock 操作都会使得 notify 增加
    head   *waiter  // 链表头指针
    tail   *waiter  // 链表尾指针
}
```

它包含了四个成员变量：

- wait：等待锁的goroutine数量
- notify：已经通知goroutine的次数，每次锁被释放和唤醒waiter时都会增加
- head：链表的头指针，即第一个等待锁的waiter
- tail：链表的尾指针，即最后一个等待锁的waiter

waiter是notifyList中的链表元素：

```
type waiter struct {
    next *waiter // 下一个waiter指针
    prev *waiter // 前一个waiter指针
    wait *uint32 // 等待的次数
    key  uintptr // Unused by sync package, unlike some implementations.
}
```

waiter中也包含了四个成员变量：

- next：下一个waiter的指针
- prev：前一个waiter的指针
- wait：等待的次数
- key：保留字段，目前没有被使用

waiter中的wait成员变量表示当前waiter等待锁的次数，因为锁可以被释放和重新获取，所以waiter有可能会多次等待才能成功获取锁。

notifyList提供了以下四个方法：

- add：将waiter添加到链表尾部
- remove：将waiter从链表中移除
- notify：通知等待的goroutine
- notifyListLocked：获取锁并对notifyList进行操作

这些方法被用于实现锁的等待和唤醒功能。当一个goroutine需要等待锁时，它会调用add方法将自己添加到notifyList的链表尾部，并释放锁。当锁被释放时，锁持有者会调用notify方法通知等待的goroutine，其中notifyList表示了所有等待锁的goroutine。在notifyList的notify方法中会唤醒一个等待的goroutine，并将notify计数器加1。如果有多个等待的goroutine，它们会按FIFO的顺序被唤醒。唤醒的goroutine会通过调用Lock方法重新获取锁，它会在notifyList中等待，在获取锁之前会不断尝试等待。

notifyListLocked方法和notify方法类似，但是notifyListLocked是在获取锁的情况下对notifyList进行操作，而notify方法是在锁被释放的情况下对notifyList进行操作。在notifyListLocked方法中，notify计数器会被重置为0，这意味着唤醒操作一开始并没有被执行。notifyListLocked方法通常是在锁的实现中使用的，它会在锁获取失败时调用wait方法等待，并在锁被释放时调用notifyListLocked方法唤醒等待的goroutine。



