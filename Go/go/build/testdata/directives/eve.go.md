# File: eve.go

eve.go是Go语言官方网站Go tour中的一个示例，用来展示如何使用goroutine和channel来实现事件循环（event loop）。

事件循环是一种常用的编程模型，通常用于GUI编程、网络编程和控制流。它通过不断地等待事件的发生，然后调用相应的处理函数来处理事件。在多线程编程中，事件循环也可以用来避免线程阻塞和CPU占用过高的问题。

在eve.go中，使用两个goroutine来模拟事件的发生和处理。一个goroutine不断地生成随机数，并通过channel发送给另一个goroutine进行处理；另一个goroutine则接收随机数，并进行奇偶数的判断，然后输出相应的信息。

具体实现中，使用了select语句来等待事件的发生，并通过默认分支来避免阻塞。同时，也使用了关闭channel来告诉goroutine停止运行。

总之，eve.go展示了Go语言并发编程中重要的概念和技术，包括goroutine、channel、select等，对于学习和理解Go语言并发编程非常有帮助。

