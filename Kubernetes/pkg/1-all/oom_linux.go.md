# File: pkg/util/oom/oom_linux.go

在Kubernetes项目中，pkg/util/oom/oom_linux.go文件的作用是为Linux操作系统实现OOM（Out of Memory）调整器，该调整器用于设置进程组内进程的OOM分数。OOM分数是Linux内核中用于衡量进程对系统内存的需求以及应对内存不足的能力的一种指标。

接下来，我来详细介绍一下NewOOMAdjuster、getPids、applyOOMScoreAdj和applyOOMScoreAdjContainer这四个函数的作用：

1. NewOOMAdjuster: 该函数用于实例化一个OOM调整器对象。它会根据给定的进程组ID（PID）创建一个新的调整器对象，并将该进程组的进程的OOM分数保存在该对象的属性中。

2. getPids: 这个函数用于获取指定进程组（PID）的所有子进程的PID。它通过读取/proc目录下相应进程组ID的子目录来获取所有子进程的PID，并返回一个整数切片。

3. applyOOMScoreAdj: 该函数用于为给定的PID设置OOM分数。它通过读取/proc/[PID]/oom_score_adj文件，并将指定的OOM分数写入该文件来调整进程的OOM分数。OOM分数的范围为-1000到1000，其中-1000表示极高优先级，而1000表示极低优先级。

4. applyOOMScoreAdjContainer: 这个函数的作用是为一个Kubernetes容器的所有进程设置OOM分数。它通过获取容器的主进程的PID，并利用applyOOMScoreAdj函数来为容器内的所有进程设置相同的OOM分数。这样可以确保整个容器内的进程在内存不足时被操作系统精确地杀死。

以上是pkg/util/oom/oom_linux.go文件中的几个重要函数的作用简介。它们一起为Kubernetes项目提供了对Linux操作系统中OOM调整的支持，从而优化容器化应用的内存管理和系统的稳定性。

