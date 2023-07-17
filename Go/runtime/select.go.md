# File: select.go

select.go这个文件是Go语言运行时（runtime）中的一个模块，主要负责实现Go语言中的select语句。Select语句是Go语言中一种用于处理多个通道并发操作的语句，可以让程序阻塞等待多个通道中的任意一个有可读取的数据或可写入的空间，从而实现对多个并发操作的统一控制。

在select.go文件中，主要包含了以下几个部分：

1. select结构体：用于定义select语句的结构体类型，包含了可读通道、可写通道、已执行状态等信息。

2. selectcase结构体：用于定义select语句中单个case子句的结构体类型，包含了通道、方向、发送/接收操作、已执行状态等信息。

3. selectgo函数：用于在新的goroutine中执行select语句。

4. selectnbrecv、selectnbrecv2、selectnbsend、selectnbsend2等函数：用于执行select语句中的接收/发送操作，并返回操作结果。

5. selectsetpc函数：用于将当前程序计数器（PC）设置为给定位置，从而跳转到对应的代码块中执行。

总的来说，select.go文件的作用是实现了Go语言中的select语句，利用多通道并发操作的特性实现了高效的并发编程。在使用select语句时，编译器会自动调用select.go中的代码来实现语句的执行。




---

### Var:

### chansendpc

变量chansendpc定义在select.go文件中，用于表示chansend函数的程序计数器（Program Counter，PC），它记录了正在执行chansend函数时的代码地址。在Golang中，程序计数器是一个寄存器，它保存了当前正在执行的代码地址。

在执行chansend函数时，如果发送操作无法立即完成，即通道已满，则该goroutine会被阻塞，等待接收者从通道中读取数据。这时，该goroutine会被放到通道的发送者队列中，等待接收者读取数据后唤醒它。同时，执行chansend函数的程序计数器也会被保存起来，后续该goroutine被唤醒时会从保存的PC处继续执行。

通过使用程序计数器，Golang可以实现goroutine的协作和调度，以实现高效的并发编程。



### chanrecvpc

在 Go 语言的 runtime 包中，select.go 文件中的 chanrecvpc 变量的作用是保存和处理一个 Go 程序中使用的通道的 receive 操作。chanrecvpc 变量是一个结构体类型的变量，其定义如下：

```go
type hchan struct {
    // ... 其他的字段
}

type sudog struct {
    // ... 其他的字段
}

type selectCtx struct {
    // ... 其他的字段
    recvOK    bool       // 接收操作是否成功
    elem      unsafe.Pointer // 接收到的元素
    selected  *sudog     // 已经被选中的 sudog
    order     uint32     // 选中的 sudog 在列表中的顺序
    //
    // Select 操作阻塞等待时会调用 onWait 方法，该方法会将 block 操作携带的 chanrecvpc 协程放入到 sudog 等待队列中
    //
    // 一旦 chanrecvpc 协程被唤醒，会调用 onWakeup 方法，它会将 block 操作携带的 chanrecvpc 协程从 sudog 等待队列中删除
    onWait    func(sg *sudog) // 暂存 chanrecvpc 协程的阻塞方法
    onWakeup  func(sg *sudog) // 查询 chanrecvpc 协程是否被唤醒的方法
    orderChan chan uint32 // 从 sudog 通知 chanrecvpc 排名信息的管道
}

// chanrecvpc 表示 "channel receive primitive with a PC"，即带有 PC 值的通道接收原语
type chanrecvpc struct {
    ptr *uintptr // chan 存储地址
    chan_ interface{} // 等待接收的通道
    recvx *uintptr // 接收操作的值存储地址
    recvk uintptr // 接收操作的值存储大小
    done *uint32 // 接收操作状态值
    ncases int // 管道列表长度，即当前 select 语句中管道的数量
    cstmt uint32 // 该 chanrecvpc 在 select 语句列表中的编号
    list  *selectCtx // select 语句中所有通道的列表
    // ... 其他的字段
}
```

在 Go 语言中，通道是一种基本的数据类型，它可以实现不同的 goroutine 之间的同步和通信。chanrecvpc 变量的作用是在实现运行时的 Select 操作过程中，处理当前 select 语句中使用的通道接收操作（receive operation）的相关状态和内容。Select 操作可以同时等待多个通道的数据到来，并根据到来的数据进行不同的处理，表示一个 select 代码块的chanrecvpc用于处理一个通道接收操作的状态信息。同时，每个 chanrecvpc 对应一个 selectCtx 变量，该联合体变量用于保存 select 语句的相关状态，包括接收状态的输入信息和结果状态的输出信息。当这些输入或输出状态发生变化时，selectCtx 将用于通知其他相关联的术语和变量，来执行正确的操作。因此，chanrecvpc 变量是实现 Select 操作的必要部分，因为它可以协调通道数据的接收以及 goroutine 和 channel 之间的交互。






---

### Structs:

### scase

scase 结构体表示一个 select 语句中的一个 case 子句，它记录了 case 子句中的操作、通道以及通道操作的结果等。

scase 结构体的定义如下：

```
type scase struct {
    c    *hchan             // Channel to use (may be nil)
    recv reflect.Value      // Argument to send/receive
    kind reflect.SelectCase // select{Recv,Send,Default}
    pc   uintptr            // return pc
}
```

- c：表示要进行 channel 操作的通道，如果是 default 分支则为 nil。

- recv：表示要发送或接收的值。如果是发送操作，则该字段必须是可寻址的，否则会抛出 panic。如果是接收操作，则该字段必须是指向储存接收值的地址的指针，否则同样会抛出 panic。

- kind：表示这个 case 分支所对应的类型，可选值有 reflect.SelectRecv（表示接收操作）、reflect.SelectSend（表示发送操作）和 reflect.SelectDefault（表示默认分支）。

- pc：表示分支语句的返回地址，用于跳转。如果是 default 分支，则该字段为 0。

在 select 语句的执行过程中，系统会为每一个 case 分支创建一个 scase 对象，并将分支所对应的数据赋值给它对应的 scase 对象。同时，这些 scase 对象将按照在 select 语句中出现的顺序存储在 sel 中的 scase 数组中。之后，系统通过执行一系列轮询和条件判断来确定应该选择哪个分支。

总之，scase 结构体记录了 select 语句中的 case 子句所需要的全部信息，为 select 语句的运行提供了必要的支持。



### runtimeSelect

在Go语言中，select 语句用于监听多个通道的读写任务，它可以等待多个通道中的第一个 I/O 操作完成并执行相应的操作。而 runtimeSelect 结构体就是用来支持 select 语句的。

runtimeSelect 结构体定义在 select.go 文件中，其中包含了一些字段和方法。下面是 runtimeSelect 结构体的定义：

```go
type runtimeSelect struct {
    tcase       unsafe.Pointer // []scase
    pollorder   unsafe.Pointer // []uint16 用于轮询scase的顺序
    lockorder   *uint16        // 多goroutine select时的锁和顺序，用于避免饥饿
    ncase       uint16
    pollclosed  uint32
    waitclosed  uint32
    startTime   int64
    context     *sudog
    options     uint32
}
```

runtimeSelect 结构体的几个主要字段含义如下：

- tcase：一个 scase 类型的切片，每个 scase 类型表示一个通道的读写操作；
- pollorder：一个 uint16 类型的切片，用于存储 tcase 中通道的索引；
- lockorder：一个指向 uint16 类型的指针，表示多个 goroutine 使用相同的 select 时的锁和顺序；
- ncase：表示 tcase 中 scase 的数量；
- pollclosed 和 waitclosed：表示 select 是否关闭；
- context：一个指向 sudog 类型的指针，用于唤醒此 select 监听的 goroutine；
- options：表示 select 的一些选项，例如是否设置了超时和抢占等。

除了上述字段之外，runtimeSelect 结构体也提供了一些方法，用于处理 select 语句的执行过程，例如 runtimeSelect.selectcase 和 runtimeSelect.block 等。

在 Go 的运行时中，当执行到 select 语句时，会创建一个 runtimeSelect 类型的结构体，并根据 select 中的 case 条件初始化 tcase 和 pollorder 等字段。接着，runtimeSelect.selectcase 方法会遍历 tcase 中所有的 scase，检查其中的操作是否可以执行。如果有操作可以执行，就会执行这个操作并返回结果。如果所有操作都无法执行，则进入 runtimeSelect.block 方法等待某个操作可以执行。当有操作可以执行时，runtimeSelect.selectcase 方法就会处理相应的操作并返回。这样，就完成了 select 语句的执行。



### selectDir

在 Go 语言中，select 语句用于从多个 channel 中接收数据，它会一直阻塞，直到其中一个 channel 可以进行数据交互为止。为了实现这一功能，Go 在 runtime 包中定义了 selectDir 结构体。

selectDir 结构体表示一个 channel 的数据传输方向，它有以下几个字段：

- casep：指向 case 结构体的指针，表示 select 语句中的某个 case 语句。
- dir：枚举类型，表示 channel 的传输方向，有发送(send)、接收(recv)和默认(default)三种。
- index：对于发送操作，表示要发送的数据在 case 中的位置；对于接收操作，表示要接收数据的变量在 case 中的位置；对于默认操作，表示该 case 在所有 case 中的位置。

selectDir 结构体的作用主要体现在 selectnbrecv 和 selectnbsend 函数中。这两个函数会在 selectDir 中表示的 channel 上进行非阻塞的接收或发送操作。

总的来说，selectDir 结构体是 select 语句实现中的一个重要组成部分，用于表示 channel 的数据传输方向，并支持非阻塞的数据传输操作。



## Functions:

### selectsetpc

selectsetpc()函数是Go语言中runtime包中select.go文件中的一个函数，它的作用是设置goroutine的PC值和SP值，以此切换到选择的case执行体中。

在Go语言中，select语句用于在多个通道之间进行选择。当一个或多个channel中出现数据时，select语句会选出其中一个case进行处理，而其他的case将被忽略。

在实现select语句时，Go语言使用了一种特殊的方式来实现并发的选择。它首先会将所有case包装为一个select-case结构体，并存放到一个数组中，然后调用selectgo()函数执行。

在selectgo()函数中，会使用selectsetpc()函数将下一个需要执行的case的PC值和SP值保存到当前的运行栈中，然后使用Go语言的栈切换机制，将当前的goroutine挂起，切换到下一个goroutine执行。

当下一个goroutine执行完毕后，会将之前保存的PC值和SP值恢复，继续执行selectgo()函数，直到所有的case都被处理完毕或者等待时间超时。

总之，selectsetpc()函数的作用是在Go语言的select实现中，对goroutine的PC值和SP值进行设置，以此实现并发的选择和切换。



### sellock

在Go语言中，`select`语句用于在多个通道之间进行选择操作。当有多个通道都处于可操作状态时，`select`语句会随机选择一个通道进行操作，从而实现非阻塞的并发编程。

`sellock`函数是`select`语句中的一个关键函数，它用于对所有的通道进行加锁操作。`sellock`函数会先获取当前`g`（goroutine）的`selectdone`指针，然后遍历所有的通道，逐个对其进行加锁操作。当某个通道被选中后，`sellock`函数会对其进行解锁操作，然后返回该通道的索引。

`sellock`函数的实现源码如下：

```go
func sellock(sc *sudog, locked *bool) (selIdx uint16) {
    // ...
    // Get the g.selectDone (if not already cached).
    sdc := g.selectDone
    if sdc == nil {
        sdc = new(selectDone)
        g.selectDone = sdc
    }
    // ...
    // Lock all channels involved.
    sc.locked = locked
    sc.next = sdc.waitq
    sdc.waitq = sc
    lock(&sc.elem.getPtr().qlock)

    for i := uint16(0); i < sc.sel.n; i++ {
        if sc.sel.rcases[i].kind == recv {
            lock(&sc.sel.rcases[i].ch.recvq.lock)
            sc.sel.rcases[i].ch.recvq.waiting++
        } else {
            lock(&sc.sel.rcases[i].ch.sendq.lock)
            sc.sel.rcases[i].ch.sendq.waiting++
        }
    }
    // ...
    return selected
}
```

具体来说，`sellock`函数实现的核心逻辑包括：

1. 获取当前`g`的`selectdone`指针，如果指针为`nil`则新建一个`selectDone`对象。
2. 将`sudog`对象的`locked`字段设置为`locked`。
3. 将`sudog`对象加入到`selectdone`的等待队列中。
4. 对所有的通道进行加锁操作，统计等待通道个数。
5. 返回被选中的通道的索引。



### selunlock

在select的实现中，selunlock()是用来解锁互斥锁m锁的。在执行select操作期间，多个goroutine可能被阻塞在不同的通道上，当其中一个通道准备好后，就需要唤醒对应的goroutine进行处理。

selunlock()的主要作用是用来解锁互斥锁，允许在通道上等待的goroutine可以重新竞争互斥锁，以便抢占其他通道中的数据。

具体地说，当某个goroutine被阻塞在某个通道上调用selunlock()后，会解除阻塞并将goroutine从g等待队列中删除。同时，该通道下的其他等待goroutine也会被唤醒并竞争互斥锁。如果没有其他等待goroutine，则互斥锁就会被释放并返回。

总之，selunlock()函数可以协调多个goroutine在select操作期间对通道的竞争，解除阻塞并唤醒需要处理的goroutine。



### selparkcommit

selparkcommit函数是用于在select操作中进行park的函数，并且在park时重新提交M（goroutine所在的线程）以回到syscall。select操作是一种用于在多个channel上等待数据的机制，而selparkcommit函数的作用就是将当前的M挂起，在等待期间由其他线程负责M的工作，当有数据可用时唤醒M并重新提交M以继续执行后续的操作。

具体来说，selparkcommit函数首先获取当前M的g状态（也就是goroutine的状态），如果当前g的状态为Grunning（正在运行），则将其状态修改为Gwaiting（等待中）以便于在等待期间不被调度器选中，然后再调用park函数将当前M挂起。挂起之后，当前M将进入waiting状态，等待其他线程的唤醒。

当其他线程有数据可用时，它会通过localWake（本地唤醒）或globalWake（全局唤醒）函数唤醒当前M，此时selparkcommit函数会重新将g的状态修改为Grunnable（可运行）并将其提交以继续执行。如果在等待期间select操作被取消或超时，selparkcommit函数将直接从park中返回，当前M将不会进入waiting状态，而是立即继续执行后续的操作。

总的来说，selparkcommit函数的作用是在select操作中进行park，挂起当前M并等待其他线程进行唤醒，以便于在等待期间由其他线程负责M的工作，从而提高了系统的并发性能。



### block

在 Go 语言中，select 语句用于从一组通道中检索一个可以进行通信的操作。没有任何操作可以进行时，select 语句将会阻塞，直到至少有一个通道可以进行通信。

在 runtime 包的 select.go 文件中，block 函数用于将 goroutine 阻塞在 select 语句中，直到至少有一个通道可以进行通信。

具体来说，block 函数会创建一个 SudoG 结构体，将其加入到 select 语句的阻塞列表中，并将该 goroutine 挂起。当至少有一个通道可以进行通信时，会从阻塞列表中移除 SudoG 并唤醒对应的 goroutine，使其继续执行。

此外，block 函数还会记录当前 goroutine 的 m 和 g，这些信息可以在唤醒时恢复，使得恢复后的 goroutine 可以继续执行之前的操作。

总之，block 函数的作用是将 goroutine 阻塞在 select 语句的阻塞列表中，直到至少有一个通道可以进行通信，并在恰当的时候唤醒对应的 goroutine。



### selectgo

selectgo函数是runtime包中select语句的实现。它用于在多个通道之间选择交互某些操作。在Golang中，select语句是用于同时等待多个通道操作的高级控制结构。

selectgo函数的主要作用是：

1.管理Goroutine的选择逻辑，使得Goroutine可以在多个通道中进行交互。
2.用于实现select语句，因此被调用时它会等待所有与该语句关联的通道都有可用的数据或者发生读写事件。
3.一旦某个通道有可用的数据或者发生读写事件，selectgo函数会选择这个通道并执行相应的操作。如果多个通道同时有可用的数据或者多个通道同时发生读写事件，则随机选择一个通道并执行相应操作。

总的来说，selectgo函数是Golang中非常重要的一个函数，它扮演了多通道协作的枢纽角色，使得Goroutine可以更加高效地处理并发问题。



### sortkey

在Go语言中，select语句用于处理多个通道的数据。在select.go文件中，sortkey（）函数用于对通道集合进行排序，以确定哪个通道可以被选择。该函数使用的策略是将通道集合按照通道ID的顺序进行排序，从而使较小的ID优先被选择。

具体地说，sortkey（）函数接收一个通道集合作为参数，并返回一个排序后的通道集合和一个值映射表。其中值映射表用于将排序后的通道索引映射回原始的通道索引，从而在后续处理中使用。该函数的实现使用了排序算法中的快速排序（quicksort）。

总之，sortkey（）函数在select语句的实现中发挥了非常重要的作用，它通过排序通道集合确定了哪个通道可以被选择，从而保证了程序的正确性和稳定性。



### reflect_rselect

在Go语言中，select语句用于同时等待多个通道的操作。在runtime包的select.go文件中，reflect_rselect函数是用于执行select语句的实现。当select语句中的case子句中有一个或多个通道时，reflect_rselect函数会调用reflect.chanrecv函数对通道进行接收操作，并返回第一个准备好的通道的操作结果。

reflect_rselect函数的具体作用包括以下几点：

1. 调用reflect.chanrecv函数执行通道的接收操作。

2. 根据通道的操作结果执行相应的case子句。

3. 如果所有的case子句都没有准备好，则等待其中一个准备好再执行相应的case子句。

4. 如果所有的case子句都已经准备好，则随机选择一个case子句执行相应的操作。

总体来说，reflect_rselect函数是select语句的核心实现，它实现了等待多个通道的操作，对准备好的通道进行操作，并根据结果执行相应的case子句。



### dequeueSudoG

在 Go 语言中，`select` 语句可以同时等待多个 `channel` 的操作。当其中一个 `channel` 准备就绪，就会执行对应的 `case` 分支。`dequeueSudoG` 是 `select` 的一个实现方法，作用是从某个 `channel` 中取出可执行的 `goroutine`（协程）。

`dequeueSudoG` 会从一个 `waitq` 的队列里取出一个 `goroutine`，然后将其加入到 `running` 的队列里。具体流程为：

1. 从队列头取出一个 `sudog`，代表一个等待着可以执行的 `goroutine`。
2. 将 `sudog` 的状态设置为执行状态。
3. 将 `sudog` 添加到 `running` 的队列中。
4. 返回 `sudog`。

`dequeueSudoG` 中涉及的 `waitq` 和 `running` 队列是 `runtime` 包中用来管理协程的两个重要队列。其中，`waitq` 存放等待 `channel` 的 `goroutine`，`running` 存放正在运行的 `goroutine`。

总的来说，`dequeueSudoG` 的作用是从等待 `channel` 的队列中取出一个可以立即执行的协程，并将其加入到运行队列中，以确保并发执行的正确性和高效性。



