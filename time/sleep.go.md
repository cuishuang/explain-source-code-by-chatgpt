# File: sleep.go

sleep.go文件是Go语言标准库中time包的一个源文件，主要实现了time包中的Sleep函数。

Sleep函数的作用是让当前goroutine暂停一段时间，指定的时间可以是纳秒、微秒、毫秒、秒、分、小时等。在暂停期间，当前goroutine会被阻塞，但其他goroutine仍可进行操作。当暂停时间过去后，当前goroutine会被唤醒重新执行。

具体实现上，Sleep函数调用了runtime包中的gosched函数，该函数会将当前goroutine加入到调度器的可运行队列中，并让调度器立即切换到其他可运行的goroutine。在指定的时间到达后，调度器会将该goroutine从可运行队列中取出并重新执行。

Sleep函数的实现考虑到了系统的性能和精度的平衡。当指定的时间较长时，Sleep函数会使用系统提供的精度较低的计时器，在计时器到期前可能会出现一定误差。当指定的时间较短时，Sleep函数会使用系统提供的精度较高的计时器，但这也会占用一定的系统资源。

总之，Sleep函数是Go语言中用于暂停当前goroutine的重要函数，提供了不同精度和开销的时间单位，同时考虑到了系统的性能和精度的平衡。




---

### Structs:

### runtimeTimer

在Go语言的标准库中，time包提供了对时间的处理，包括计时器和时钟等。其中，sleep.go文件提供了睡眠功能的实现，用于让当前的Goroutine进入一个睡眠状态。在该文件中定义了一个runtimeTimer结构体，具体作用如下：

1. 等待时间的计时器

runtimeTimer结构体用于计算等待时间，并在一定时间后向Goroutine发送一个当定的信号。通过runtimeTimer结构体，我们可以设定等待多长时间，然后在该时间结束后执行某个操作。

2. 用于执行操作的回调函数

runtimeTimer结构体中包含了一个f函数字段，我们可以在这里指定一个操作，当定时器到期时自动执行该操作。

3. Timer链表节点

将所有的计时器信息保存到一个链表中，通过“next”字段将各个计时器节点连接起来。这些节点按照等待时间的递增顺序排列，可以有效提高计时器触发的效率。

总之，runtimeTimer结构体是Go内部实现计时器的基础设施，它承担了计时和操作执行的功能，可以非常方便地实现各种时间控制的功能。



### Timer

Timer结构体在time包中是一个计时器，用于在固定时间后触发一个事件。Timer结构体主要用于以下场景：

1. 延迟执行

使用Timer可以在指定时间后执行一段代码，例如：

```
timer := time.NewTimer(time.Second)
<-timer.C
fmt.Println("1 second has passed")
```

上面的代码会等待1秒钟后输出"1 second has passed"。

2. 定时执行

使用Timer还可以定时执行一段代码，例如：

```
func doSomething() {
    fmt.Println("do something")
}

timer := time.NewTimer(time.Second)
for {
    <-timer.C
    doSomething()
    timer.Reset(time.Second)
}
```

上面的代码每隔1秒钟就会执行一次doSomething函数。

3. 超时控制

使用Timer还可以控制一段代码的超时时间，例如：

```
func doSomething() error {
    // do something that may take some time
}

func main() {
    timeout := time.NewTimer(time.Second * 10)
    select {
    case <-timeout.C:
        fmt.Println("time is up")
    case err := <-doSomething():
        if err != nil {
            fmt.Println("do something error")
        }
    }
}
```

上面的代码通过Timer实现了对doSomething函数执行的超时控制，如果doSomething函数在10秒钟内没有执行完毕，就会输出"time is up"。



## Functions:

### Sleep

Sleep是time包中的一个函数，它的作用是让当前的Goroutine（协程）休眠一段时间，也就是暂停执行一段时间。具体来说，Sleep函数会暂停当前协程的执行，等待指定的时间过去后，当前协程重新被唤醒继续执行。

Sleep函数的定义如下：

```
func Sleep(d Duration)
```

其中，d是一个Duration类型的参数，表示要休眠的时间段。Duration类型是一个表示时间段的基本类型，单位可以是纳秒、微秒、毫秒、秒等等，可以通过time包提供的常量进行表示。

Sleep函数在底层是通过调用操作系统提供的睡眠函数实现的，实现方式会依赖于具体的操作系统实现。在Linux系统中，Sleep函数会使用nanosleep系统调用；在Windows系统中，Sleep函数会使用SleepEx函数或者WaitForSingleObjectEx函数。无论是哪种实现方式，Sleep函数都是一种非常高效和精确的暂停协程的方式，可以用于实现各种需要暂停一段时间的场景。



### when

当我们调用time包中的sleep函数时，其实最终会调用的是when函数。when函数的作用是计算出当前时间和目标时间之间的时间差，然后调用计时器（timer）的sleep函数，使当前goroutine暂停指定的时间。 

具体来说，when函数会首先计算出当前时间和目标时间之间的纳秒数，然后调用time包中的sleepNanoseconds函数进行实际的休眠。在进行实际休眠之前，如果当前的计时器状态不是正在运行中，那么when函数会先调用startTimer函数来启动计时器。

在启动计时器之前，when函数还会检查当前时间是否小于目标时间。如果小于目标时间，则直接返回；否则将计时器状态置为正在运行中，并且启动计时器。

总之，when函数的作用是通过计算当前时间和目标时间之间的时间差，来使当前goroutine暂停指定的时间。它是time包中sleep函数的核心实现。



### startTimer

startTimer是time包中的一个函数，用于启动一个定时器。

具体来说，startTimer的作用是：

1. 创建一个timer对象，即一个触发信号通道和一个计时器。timer对象表示一个计时器，可以在指定时间后向其触发信号通道发送信号。

2. 将timer对象加入到全局计时器堆中，时间戳为当前时间加上指定时间。

3. 用一个goroutine循环等待计时器的信号，一旦收到信号就会执行相应的处理。

startTimer函数的代码如下：

```go
func startTimer(t *timer) {
    when := t.when.UnixNano()
    runtimeNano := runtimeNano()
    if when < runtimeNano {
        when = runtimeNano + 1
    }
    t.when = time.Unix(0, when)
    t.f = func() {
        t.when = zeroTime
        t.c.sendTime(runtimenano())
        lock(&timersLock)
        heapRemove(&timers, t.i)
        unlock(&timersLock)
    }

    lock(&timersLock)
    heapPush(&timers, t)
    unlock(&timersLock)

    if timers[0] == t {
        if t.when.Sub(now()) <= 0 {
            go t.f() // non-blocking send
        } else {
            startTimerThreadLocked()
        }
    }
}
```

在具体实现中，startTimer函数首先计算出timer应该触发的时间戳when(单位为纳秒)，然后将timer对象加入到全局计时器堆中。如果timer触发时间早于当前时间，则将触发时间设置为当前时间加一纳秒，以确保计时器能够正常触发。

最后一步是检查计时器堆顶元素是否为当前timer对象。如果是，就判断计时器是否已经过期，如果已经过期就直接执行计时器函数；如果未过期则启动一个新的goroutine等待计时器过期。这个过程是在startTimerThreadLocked函数中完成的。



### stopTimer

在 Go 语言的 time 包中，stopTimer() 函数是用于停止计时器的。它的具体作用是停止 Timer 结构体类型的计时器对象。一旦计时器被停止，它就不能被重新启动，而必须使用一个新的计时器对象。

在 sleep.go 文件中，stopTimer() 函数被用来停止计时器。在执行休眠操作时，为了保证休眠的精度和可靠性，使用了一个计时器对象来进行计时和触发信号。当休眠操作需要被中断时，stopTimer() 函数被调用来停止计时器，以避免触发不必要的信号。

此外，stopTimer() 函数还会将计时器结构体对象的状态设置为“已停止”，以便其他程序可以检测到计时器已经停止，从而避免在停止计时器之后再次使用它。



### resetTimer

resetTimer函数是time包中的一个私有函数，主要用于重新设置定时器的计时器，并且返回该定时器剩余的时间。该函数的代码如下：

func resetTimer(now int64, d Duration) int64 {
    if d <= 0 {
        return now
    }
    t := now + int64(d)
    t -= t % int64(d)
    return t
}

resetTimer函数接受两个参数：当前时间戳和要重新设置的计时器的时间间隔。如果时间间隔小于等于0，则直接返回当前时间戳并停止计时器；否则，将当前时间戳和时间间隔相加得到计时器结束时间戳t，然后将t舍去余数，即将t舍去到最近的时间间隔的整数倍。比如，如果时间间隔是1秒，当前时间戳是1595260410，那么t就是1595260411，因为1秒后是1595260411秒，并将t舍去余数后得到1595260410，即最近的时间间隔的整数倍。最后，resetTimer函数返回t作为重新设置后的计时器结束时间戳，并且返回剩余时间d，即t-now。

resetTimer函数的作用是重新设置定时器的计时器并返回剩余时间，比如在time包中的sleep函数中，如果需要在sleep期间进行一些操作，比如定时打印日志，就需要重新设置计时器的计时器，以便在sleep期间能够及时执行操作。resetTimer函数可以方便地计算出新的计时器结束时间戳，并且返回剩余时间，使得调用方能够及时处理操作。



### modTimer

modTimer这个函数的作用是修改已经存在的计时器的触发时间。在goroutine中调用time.Sleep()时，会创建一个计时器并添加到全局的计时器堆中，等待到达设定的时间后便会唤醒该goroutine。

在一些情况下，我们可能需要修改已经存在的计时器的触发时间，比如当某个事件提前或延迟发生时，我们需要重新设置该事件的触发时间，这时就可以通过调用modTimer函数来实现。

该函数接收两个参数，第一个参数是计时器的指针，第二个参数是新的触发时间。函数会先将指针对应的计时器从堆中删除，然后根据新的触发时间重新计算计时器的超时时间并添加到堆中，以便触发唤醒操作。同时也需要保证堆中计时器的有序性。



### Stop

Stop函数用于停止一个正在运行的Timer（计时器）并且防止它在未来触发。如果Timer已经被停止或者过期，Stop方法返回false。如果Timer是nil，Stop方法不会产生错误。

当调用Stop方法时，Timer会立即停止，Timer的回调函数不会被执行。如果以后再次调用Timer的Reset方法，将不会影响它之前的触发时间，只会根据当前的时间再次计算触发时间。

Stop方法在多线程环境中可以安全地使用，包括同时调用Reset和Stop方法。此外，在Timer已经过期但回调函数还没有被执行时调用Stop方法不会阻止回调函数的执行。

总之，Stop方法用于停止正在进行的Timer并取消它的回调函数，防止它在未来触发。



### NewTimer

NewTimer函数是用于创建一个新的Timer对象的函数。Timer表示一个定时器，它会在指定的时间后向其管道C发送一个时间值。

函数定义如下：

```go
func NewTimer(d Duration) *Timer
```

它接收一个Duration类型的参数d，表示定时器的时长。返回一个指向新创建的Timer对象的指针。

例如，如果想要创建一个5秒的定时器，可以像下面这样使用NewTimer函数：

```go
t := time.NewTimer(5 * time.Second)
```

创建了一个新的Timer对象后，可以使用它的Reset方法来重置定时器的时长，也可以使用它的Stop方法来停止定时器。

另外，如果希望立即触发定时器，可以向其管道C发送当前时间值，例如：

```go
func main() {
	t := time.NewTimer(5 * time.Second)
	now := time.Now()
	t.Reset(1 * time.Second) // 重置定时器的时长为1秒
	<-t.C // 等待定时器触发，打印当前时间和触发时间
	fmt.Println("current time:", now)
	fmt.Println("timer fired at:", time.Now())
}
```



### Reset

Reset函数在调用时会重置计时器的时间，并开始一个新的计时。具体而言，它会清空计时器的时间，停止之前的计时（如果有的话），并重新开始一个新的计时。具体实现过程如下：

首先，Reset函数会执行一个atomic操作，即将计时器的状态标记为resetting。这个标记的作用是防止在重置过程中出现竞态条件。

接着，如果计时器已经在运行中，则调用stopTimer函数停止计时器。stopTimer函数会停止计时器并将其状态标记为负数。

然后，Reset函数会清空计时器的时间和duration，以确保时间从0开始计时。

最后，如果计时器之前是处于运行状态，则调用startTimer函数重新启动计时。startTimer函数会重新设置计时器的时间和duration，并将其状态标记为正数。

总之，Reset函数的作用是清空计时器的时间并重新开始计时，以便使用它来测量一个新的时间段。



### sendTime

sendTime函数是一个内部函数，用于向当前的时间管道中发送当前时间值。

在Go的time包中，时间管道是一个用于传递时间值的通道。当我们使用time.Sleep()函数时，实际上是将当前goroutine阻塞，并发送一个时间值到时间管道中。等到时间到达后，时间管道会自动发送一个通知，通知阻塞的goroutine继续执行。

sendTime函数的作用是在time.Sleep()函数中发送时间值到时间管道中，以通知阻塞的goroutine继续执行。具体代码如下：

```go
func sendTime(c chan Time, t Time) {
    select {
    case c <- t:
    default:
    }
}
```

在这个函数中，首先使用了一个select语句来向时间管道中发送当前时间值。如果时间管道已经满了，select语句就会执行default分支，直接忽略当前的时间值。这样做的原因是为了避免阻塞的goroutine一直等待时间管道空出位置，从而影响程序的性能。

总之，sendTime函数是time包中的一个内部函数，用于向管道中发送时间值，以通知阻塞的goroutine继续执行。



### After

After函数的作用是返回一个通道chan，当该通道可读时，表示经过指定时间后已经到达了指定的时间点。具体来说，After函数等待指定的时间后，向该通道写入当前时间，并立即返回该通道，从而允许程序在未到达指定时间点时继续执行其他任务。

例如，以下代码将在10秒后打印一条消息：

```
select {
case <-time.After(10 * time.Second):
    fmt.Println("Time's up!")
}
```

在需要等待一段时间后执行某个操作时，After函数非常有用。它可以用来实现超时控制、定时任务等。

需要注意的是，由于After函数返回的通道可以被多次读取，因此需要在使用该函数时确保只读取一次通道。否则可能会导致结果不正确或死锁等问题。例如，以下代码使用了一个与After函数返回的通道相同的变量，将导致死锁：

```
ch := time.After(10 * time.Second)
select {
case <-ch:
    // 这里会一直阻塞
}
```



### AfterFunc

AfterFunc是time包中的一个函数，它的作用是在指定的时间后执行一个函数。

它的函数签名为：
```
func AfterFunc(d Duration, f func()) *Timer
```

其中，d表示时间间隔，f表示需要执行的函数。函数f会在d时间之后执行。

返回值是一个Timer类型的指针，它可以用来停止等待时间的执行。

这个函数通常用于在指定时间后触发一些操作，例如超时处理、定时任务等等。它支持任意的函数作为参数，因此可以实现非常灵活的定时任务。

需要注意的是，这个函数不是精准的定时器，它的精度受到系统时间调度的影响。如果需要精确的定时器，可以使用Ticker类型来实现。



### goFunc

在go/src/time中的sleep.go文件中，goFunc函数是一个用于启动goroutine的函数。它的作用是创建一个新的goroutine，将指定的函数f作为其参数，并在goroutine中执行该函数。

具体来说，goFunc函数会创建一个新的goroutine，并调用runtime.newproc函数，将指定的函数f和它的参数一起传递给newproc函数。newproc函数会创建一个新的goroutine，并将该函数及其参数绑定到该goroutine中，最终该goroutine会执行这个函数。

在time.Sleep函数中，goFunc函数的作用是启动一个新的goroutine来执行timerSleep函数，该函数是一个阻止调用goroutine并作为时间计时器的核心部分。

总之，goFunc函数是用于启动goroutine的函数，在time.Sleep函数中，它被用作启动新的goroutine执行timerSleep函数的工具。



