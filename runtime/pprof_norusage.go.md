# File: pprof_norusage.go

pprof_norusage.go是Go语言标准库中的一个文件，其作用是提供一组函数，用于采集运行时系统的性能数据，并生成性能分析报表。具体而言，pprof_norusage.go实现了以下功能：

1. 收集CPU使用率数据：通过调用runtime/pprof包中的CPUProfile函数，采集程序的CPU使用率数据，并将其写入指定的文件中。采集数据的时间间隔可以通过SetCPUProfileRate函数设置。

2. 收集内存使用率数据：通过调用runtime/pprof包中的WriteHeapProfile函数，采集程序的内存使用率数据，并将其写入指定的文件中。该函数还会将goroutine的堆栈信息一并写入文件，以供进一步分析。

3. 收集阻塞事件数据：通过调用runtime/pprof包中的BlockProfile函数，采集程序的阻塞事件数据，并将其写入指定的文件中。采集的数据包括发生阻塞的goroutine数量、阻塞时间、等待时间等。

4. 收集互斥锁竞争数据：通过调用runtime/pprof包中的MutexProfile函数，采集程序中互斥锁的竞争情况，并将其写入指定的文件中。采集的数据包括互斥锁的名称、竞争次数、持有该锁的goroutine等信息。

5. 读取CPU使用率数据：通过调用runtime/pprof包中的LookupCPUProfile函数，读取指定文件中的CPU使用率数据，并返回相应的分析报表。报表包括调用栈信息、每个函数占用CPU时间的占比等。

6. 读取内存使用率数据： 通过调用runtime/pprof包中的LookupHeapProfile函数，读取指定文件中的内存使用率数据，并返回相应的分析报表。报表包括每个堆对象的大小、分布情况等。

7. 读取阻塞事件数据： 通过调用runtime/pprof包中的LookupBlockProfile函数，读取指定文件中的阻塞事件数据，并返回相应的分析报表。报表包括阻塞事件的数量、所占CPU时间的比例、主要发生在哪些goroutine中等。

8. 读取互斥锁竞争数据： 通过调用runtime/pprof包中的LookupMutexProfile函数，读取指定文件中的互斥锁竞争数据，并返回相应的分析报表。报表包括互斥锁名称、竞争次数、持有该锁的goroutine等信息。

总的来说，pprof_norusage.go提供了基于函数调用跟踪的性能分析工具，方便开发人员对Go程序的性能问题进行诊断和优化。

## Functions:

### addMaxRSS

在Go语言的运行时中，pprof_norusage.go这个文件中的addMaxRSS函数主要用于将操作系统进程的最大常驻集大小（即进程使用的最大物理内存量）添加到pprof文件中。这个函数会在应用程序结束时被调用，用于生成一个pprof文件，以便开发人员可以对应用程序进行性能分析和调优。

具体来说，这个函数会通过调用操作系统函数getrusage，获取当前进程的系统资源使用情况（如CPU时间、内存使用等）。然后，它会将常驻集大小转换为字节数，然后将其添加到pprof文件中。这可以帮助开发人员识别应用程序中可能存在的内存泄漏或性能瓶颈，从而进行优化。

需要注意的是，addMaxRSS函数只会在应用程序结束时被调用，因此它不能提供实时的系统资源使用情况。如果要实时监控系统资源使用情况，请使用其他工具或库（如Prometheus等）。



