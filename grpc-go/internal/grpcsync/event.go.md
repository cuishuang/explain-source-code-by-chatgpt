# File: grpc-go/internal/grpcsync/event.go

在grpc-go项目中，grpc-go/internal/grpcsync/event.go文件定义了Event结构体和相关函数，用于实现一个简单的事件机制。

Event结构体是一个带有fire和done两个通道的同步事件信号量。它有三个主要的状态：初始状态、已触发状态和已完成状态。

- Fire函数用于触发一个事件。它会向fire通道发送一个信号，将事件的状态从初始状态变为已触发状态，唤醒任何等待事件触发的goroutine。
- Done函数用于等待事件的触发并完成。它会阻塞当前的goroutine，直到fire通道被触发，事件的状态变为已完成。一旦触发事件，Done函数会返回一个关闭的done通道，用于告知事件已完成。
- HasFired函数用于检查事件是否已经触发。它会检查事件的状态是否为已触发，如果是，则返回true，否则返回false。
- NewEvent函数用于创建一个新的事件。它返回一个初始状态的Event结构体，并初始化相关的通道。

这些函数和Event结构体的组合使得可以在不同的goroutine之间进行同步操作。通过等待事件被触发和完成，可以实现对共享资源进行同步访问或者进行并发控制。

