# File: preempt_s390x.s

preempt_s390x.s是Go语言运行时中针对IBM System z（z/Architecture）平台的抢占式调度实现的汇编代码文件。

在抢占式调度中，操作系统可以在当前执行的 Goroutine 还没有完成时暂停该 Goroutine 的执行，在调度器的另一个 Goroutine 上执行一段时间，然后将其切换回原始 Goroutine，从而实现公平分配 CPU 时间，保证所有 Goroutine 获得公平的执行时间。

该文件中实现了针对z/Architecture的支持，包括：

1. 根据操作系统控制块（OSCB）的内容，恢复 Goroutine 的栈指针（sp）和程序计数器（pc），从而实现切换 Goroutine 的操作。

2. 根据被调度的 Goroutine 的状态，判断何时切换到下一个 Goroutine。

3. 在发生异常或中断时，恢复进程的状态并切换到内核的堆栈上处理异常或中断。

preempt_s390x.s 文件中实现了一些在 z/Architecture 平台上处理抢占式调度所必需的关键操作，但需要结合其他 Go 运行时组件一起工作，如调度器、任务队列、低级线程机制等。

