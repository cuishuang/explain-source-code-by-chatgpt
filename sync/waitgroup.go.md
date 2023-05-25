# File: waitgroup.go

waitgroup.go 是 Go 语言标准库中 sync 包中的一个文件，主要定义了 sync.WaitGroup 类型，该类型被用来等待一组 goroutine 的执行结束。

在 Go 语言中，每个 goroutine 都是独立的运行线程，因此在程序中可能存在多个 goroutine 在同时执行。在某些情况下，我们需要等待某些 goroutine 执行完毕后再执行下一步操作，这时候就需要用到 sync.WaitGroup 类型来实现。具体实现的方式是在每个 goroutine 开始执行前调用 sync.WaitGroup 的 Add 方法，在 goroutine 中执行完任务后调用 Done 方法，最后在主线程中等待所有 goroutine 完成调用 Wait 方法。

sync.WaitGroup 中的方法包括：

1. Add(delta: int)：将 WaitGroup 的计数器增加 delta，表示要等待的 goroutine 数量增加了 delta。

2. Done()：将 WaitGroup 的计数器减1，表示一个 goroutine 执行完任务。

3. Wait()：阻塞当前 goroutine，直到 WaitGroup 的计数器减为零，即所有 goroutine 都执行完毕。

下面是一个示例代码：

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("worker %d starting\n", id)
    // 模拟一些业务逻辑
    for i := 0; i < 5; i++ {
        fmt.Printf("worker %d working...\n", id)
    }
    fmt.Printf("worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }
    fmt.Println("main thread waiting...")
    wg.Wait()
    fmt.Println("main thread done")
}
```

以上代码启动了 3 个 goroutine 来执行 worker 函数，每个 goroutine 的任务都是执行一些业务逻辑。在每个 goroutine 执行开始时调用 wg.Add(1) 方法增加 WaitGroup 的计数器，表示需要等待的 goroutine 数量增加了 1。在每个 goroutine 执行结束时调用 wg.Done() 方法将计数器减 1，表示一个 goroutine 的任务执行完毕。最后在主线程中调用 wg.Wait() 方法，等待所有 goroutine 执行完毕。

在这里，我们可以看到 sync.WaitGroup 的作用：它能够控制 goroutine 执行的同步，保证一个 goroutine 只有在其它相关的 goroutine 执行完毕时才能结束。它可以有效的避免 goroutine 之间的竞争和数据竞态，从而提高程序的稳定性和可读性。




---

### Structs:

### WaitGroup

WaitGroup是一个同步工具，用于协调多个goroutine的执行顺序。它的作用是协调多个goroutine之间的执行流程，当一个goroutine需要等待其他goroutine执行完毕之后才能继续执行时，可以使用WaitGroup来完成这个功能。

WaitGroup结构体包含三个方法：

1. Add(delta int)：向WaitGroup中添加或减少计数器的值。delta为正则增加计数器的值，为负则减少计数器的值。

2. Done()：减少WaitGroup中计数器的值，相当于Add(-1)。

3. Wait()：等待所有的goroutine执行完毕，即等待计数器的值变为0。

当WaitGroup中的计数器值为0时，Wait方法将立即返回。如果计数器的值为负数，Wait方法将会panic。

例如，在一个程序中，有两个goroutine需要执行，但是第一个goroutine需要等待第二个goroutine执行完毕之后才能继续执行。可以使用WaitGroup来完成这个功能：

```
var wg sync.WaitGroup
wg.Add(1)
go func() {
    // 执行第二个goroutine的任务
    wg.Done()
}()
// 执行第一个goroutine的任务
wg.Wait()
```

如上代码中，第一个goroutine先执行添加一次计数器，然后进入等待状态。第二个goroutine执行完成任务之后，调用Done方法将计数器减一。此时，计数器的值为0，第一个goroutine解除等待状态，继续执行下去。



## Functions:

### Add

在 Go 语言中，sync.WaitGroup 是一种用于等待一组 goroutine 完成执行的工具。我们可以调用 sync.WaitGroup.Add() 方法来将等待组中待执行的 goroutine 数量加一，再调用 sync.WaitGroup.Done() 方法将完成执行的 goroutine 数量减一。

sync.WaitGroup.Add() 方法有一个参数，该参数表示需要等待执行的 goroutine 的数量。可以在主函数中使用 sync.WaitGroup 来等待所有 goroutine 执行完成。

在 sync.WaitGroup 中，Add() 方法的作用是向等待组中添加要执行的 goroutine 的数量。当我们需要等待多个 goroutine 执行完成后再执行其他操作时，可以使用 Add() 方法来增加等待组的计数器。

具体来说，这个方法会将计数器的值加上传递进来的数量，这个值必须是正整数。如果当前计数器的值小于等于 0，则调用 Add() 方法执行会造成 panic 异常。

当所有被加入等待组的 goroutine 都执行完毕后，等待组的计数器会变为 0，此时等待组会停止等待，主函数可以继续执行后续的操作。

总结来说，sync.WaitGroup.Add() 方法的作用是将等待组中待执行的 goroutine 数量加一。它是 sync.WaitGroup 中最重要的方法之一，常用于等待多个 goroutine 执行完成后再执行其他操作。



### Done

sync.WaitGroup是一种特殊的锁同步机制，它允许多个goroutine同时等待一个或多个任务的完成。这个机制实现了对并发代码的控制，避免了由于goroutine的执行顺序不确定而产生的不可预测的结果。

WaitGroup内部有一个计数器，它记录还剩余多少个任务没有完成。每个goroutine执行完一个任务时，需要调用WaitGroup的Done方法来通知计数器减少一个任务。当计数器减少到0时，表示所有任务都已完成，则所有等待的goroutine被唤醒。

Done方法的作用就是通知WaitGroup已经完成了一个任务。每次调用Done方法，内部计数器都会减1。如果计数器减为0，则表示所有任务都已完成。这时，所有阻塞在Wait方法上的goroutine都会被唤醒，并继续执行后续的代码。

Done方法的语法为：

``` go
func (wg *WaitGroup) Done()
```

参数wg是一个指向WaitGroup对象的指针。可以在多个goroutine中同时调用Done方法，每次调用都会使计数器减1。但是需要注意的是，计数器不能小于0，所以Done方法不可重复调用。

Done方法的应用非常广泛。在多线程编程中，经常需要等待一组任务完成后才能继续执行下面的代码，这时就可以使用WaitGroup来解决问题。



### Wait

在Go语言的sync包中，WaitGroup是用来等待若干个goroutine执行完毕的同步工具。WaitGroup提供了三个方法：Add、Done和Wait，其中最重要的是Wait方法。下面详细介绍Wait方法的作用。

Wait方法的作用是阻塞当前的goroutine，直到WaitGroup所管理的goroutine都执行完毕。如果WaitGroup的计数器不为零，Wait方法会一直阻塞当前goroutine，直到计数器归零。

当我们使用WaitGroup时，可以将Add方法设置的计数器的值设置为需要等待的goroutine的数量。然后，在每个goroutine的最后调用Done方法，即将计数器减一。当所有的goroutine执行完毕后，计数器变为零，即可调用Wait方法，来挂起当前的goroutine，等待所有goroutine都执行完毕。

例如，下面的代码中创建了两个goroutine，它们会执行一个较为耗时的任务。我们使用WaitGroup来等待这两个goroutine执行完毕，并输出结果。

```Go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		time.Sleep(time.Second)
		fmt.Println("goroutine 1 done.")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine 2 done.")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("all goroutines done.")
}
```

输出结果：

```
goroutine 1 done.
goroutine 2 done.
all goroutines done.
```

上面的输出结果显示，两个goroutine都执行完毕后，才会继续执行main函数中的代码，即输出"all goroutines done."。在使用WaitGroup时，需要注意，Add方法和Done方法需要在不同的goroutine中调用，否则会造成死锁。同时，需要确保Add方法的计数器的值与所管理的goroutine数量相等，否则程序会出现错误。



