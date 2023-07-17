# File: oncefunc.go

sync/oncefunc.go是sync包中实现同步机制的文件之一，主要提供了一种"只执行一次"的机制，可以在多个goroutine中仅执行一次指定的函数或方法。

oncefunc.go文件中包含了一个名为"once"的结构体，该结构体提供了三个方法：Do、done和isDone，分别用于执行一次指定的方法或函数、标记该方法或函数已经完成执行，以及判断该方法或函数是否已经执行完毕。

其中Do方法是该结构体的主要方法，用于执行函数或方法，确保其仅会被执行一次。当多个goroutine尝试执行Do方法时，只有其中一个能够成功执行，其他的goroutine会等待直到该方法执行完成。

该功能主要用于避免并发场景下重复执行某些代码的问题，可以保证某些初始化工作只会被执行一次，或者避免重复创建某些对象。

总的来说，sync/oncefunc.go的作用是提供了一种在多个goroutine中只执行一次某个函数或方法的机制，帮助我们避免并发问题。

## Functions:

### OnceFunc

OnceFunc是sync包中的一个类型，它用于确保某个函数只被执行一次。它通常被用于初始化操作，例如初始化全局变量。

OnceFunc的特点是可以并发安全地调用某个函数，并且只有第一次调用时会执行这个函数。一旦函数执行了一次，之后的并发调用都不会再次执行该函数。

具体来说，OnceFunc包含一个Done通道和一个Mutex锁。当OnceFunc第一次调用时，它会获取锁并且执行传入的函数。然后它会关闭Done通道，以此通知所有其他并发调用者这个函数已经执行完毕。其他并发调用者在调用Do方法时会发现Done通道已经关闭，从而不会再次执行传入的函数。

除了Do方法之外，OnceFunc还包含另外两个方法：Done和Reset。Done方法返回一个只读的Done通道，用于在其他goroutine中等待函数执行完毕。Reset方法用于清空Done通道和锁，从而允许再次执行传入的函数。

总之，OnceFunc提供了一种简单、安全且高效的方式来确保某个函数只被执行一次，并且合理地处理了并发执行的情况。



### OnceValue

OnceValue这个func是sync包中的一个结构体，用于保证一个函数只会被执行一次。

具体来说，OnceValue内部维护了一个done标志位，表示相关函数是否已经被执行过了。当调用Do()方法时，OnceValue首先会检查done标志位，如果已经被置为true，则直接返回；否则，调用相关函数进行执行，并将done标志位置为true，以保证下次调用Do()时可以直接返回。

OnceValue可用于实现某些只需要执行一次的操作，例如全局变量初始化、单例模式等，可以有效地避免竞态条件和重复执行的问题。



### OnceValues

OnceValues是sync包中Once类型的一个方法，用于返回Once类型的状态和值。

在并发编程中，有时我们只需要执行某个函数一次，而不需要重复执行。这时可以使用sync包中的Once类型来保证只执行一次。

OnceValues方法返回Once类型内部状态和值。如果Once已经执行过，将会返回一个特殊的空值nil，否则返回nil以外的值。可以用来检查Once是否执行过，以及执行的结果。

简单的示例代码如下：

```
package main

import (
    "fmt"
    "sync"
)

func setup() {
    fmt.Println("do setup")
}

var (
    once     sync.Once
    onceArgs string
)

func main() {
    // 执行一次，输出 "do setup"
    once.Do(setup)

    // 再次执行，不输出任何内容
    once.Do(setup)

    // 查看Once的状态和值
    state, val := once.Values()
    fmt.Println(state, val)

    // 执行另一个函数，修改onceArgs
    once.Do(func() { onceArgs = "hello" })

    // 查看Once的状态和值
    state, val = once.Values()
    fmt.Println(state, val)
}
```

输出结果为：

```
do setup
<nil> <nil>
<nil> <nil>
```

可以看到，在第一次执行setup后，Once的状态为nil，值也为nil。再次执行setup时，由于Once已经执行过，不会再输出任何内容。使用Values方法可以查看Once的状态和值，可以看到此时状态仍然为nil，值也为nil。

在执行另一个函数后，Once的状态已经发生了变化，可以看到状态不再是nil，值也不再是nil，而是一个字符串"hello"。



