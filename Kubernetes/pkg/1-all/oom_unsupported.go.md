# File: pkg/util/oom/oom_unsupported.go

在Kubernetes项目中，`pkg/util/oom/oom_unsupported.go`文件的作用是为不受支持的操作系统提供OOM（Out Of Memory）调整器的实现。由于某些操作系统的内核不支持OOM调整，因此这个文件为这些操作系统提供了兼容性支持。

这个文件中定义了如下几个要点：

1. `unsupportedErr`：这是一个常量，表示不支持OOM调整的错误信息。

2. `NewOOMAdjuster`函数：这个函数用于在不支持OOM调整的操作系统上创建OOM调整器。它返回一个实现了`OOMAdjuster`接口的结构体，该结构体的方法将不执行任何实际的OOM调整操作，而是返回一个错误，表示不支持此操作。

3. `unsupportedApplyOOMScoreAdj`函数：这个函数用于在不支持OOM调整的操作系统上应用OOM Score Adjustment。它返回一个错误，表示不支持此操作。

4. `unsupportedApplyOOMScoreAdjContainer`函数：这个函数用于在不支持OOM调整的操作系统上为容器应用OOM Score Adjustment。它返回一个错误，表示不支持此操作。

这些函数和变量的作用是为了在不支持OOM调整的操作系统上提供一个应对方案，以兼容这些操作系统的限制。当在不支持OOM调整的操作系统上调用OOM调整相关的功能时，这些函数将返回错误，告知调用者操作系统的限制。

需要注意的是，这些函数和变量只在不支持OOM调整的操作系统上起作用，而在支持OOM调整的操作系统上将使用其他实现。由于Kubernetes支持多种操作系统和平台，这种兼容性设计可以确保在各种环境中稳定地运行。

