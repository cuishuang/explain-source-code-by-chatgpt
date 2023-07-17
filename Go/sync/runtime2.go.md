# File: runtime2.go

在 Go 语言的 sync 包中，runtime2.go 文件的作用是提供了低级别的锁原语和底层锁实现。这些原语可以用来编写高性能的锁和同步结构。

这个文件中包含了一些平台相关的代码，比如 Linux、Windows、Darwin 等系统。它们实现了各种锁算法，如互斥锁、读写锁、信号量等。

其中，runtime2.go 文件中的主要函数包括：

- sema_init 和 sema_wakeup：实现了一种信号量结构，可以用于实现线程同步和互斥。
- mutex_init 和 mutex_unlock：实现了互斥锁。
- rwmutex_init 和 rwmutex_unlock：实现了读写锁。

这些函数在 Go 的底层运行时中被广泛使用，以提供线程安全和并发性。

需要注意的是，这些函数仅适用于底层编程和高性能应用中。在普通的 Go 开发中，应该优先使用标准库提供的高级别锁和同步结构，比如 sync.Mutex、sync.RWMutex、sync.WaitGroup、sync.Once 等。




---

### Structs:

### notifyList

notifyList是一个用于通知goroutine的结构体。

在go语言中，实现通知goroutine的经典方式是使用channel。但是，在某些情况下，使用channel可能会出现一些问题，比如当有大量的goroutine等待消息时，通知所有等待的goroutine可能会导致性能问题。

notifyList的设计就是为了解决这些性能问题。它使用了一种高效的设计，即在通知列表中记录goroutine，而不是使用channel，这样可以避免性能问题。

notifyList结构体中主要包含两个字段：

1. wait 表示等待通知的goroutine列表。

2. notify 表示通知事件已发生，是否有goroutine正在等待事件，以及在等待期间是否有新的goroutine加入等待列表。

当某个goroutine需要等待事件发生时，它可以调用notifyList的add方法将自己加入等待列表，并在之后等待通知。当事件发生时，调用notifyList的notify方法就可以将所有等待列表中的goroutine都唤醒。

在调用notifyList的notify方法时，会检查等待列表是否为空。如果为空，则不做任何操作。如果等待列表不为空，则遍历等待列表中的goroutine，并检查它们是否已经被唤醒过。如果已经唤醒过，则从等待列表中删除对应的goroutine，否则将对应的goroutine加入唤醒列表中。

使用notifyList可以避免使用channel时可能出现的性能问题，特别是在有大量的等待goroutine时。因此，notifyList是go语言中一个非常重要的通知机制。



