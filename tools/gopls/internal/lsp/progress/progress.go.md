# File: tools/gopls/internal/lsp/progress/progress.go

在Golang的Tools项目中，`tools/gopls/internal/lsp/progress/progress.go` 文件的作用是提供了一套用于跟踪和报告进度的工具。

以下是对每个结构体的功能的详细介绍：

1. `Tracker`: 是一个用于跟踪进度的结构体。它可以创建一个进度报告，以便在某个任务的过程中向客户端发送进度信息，例如任务的完成百分比，当前执行的操作等。

2. `WorkDone`: 代表一个正在进行中的工作，例如编译、分析等。它提供了一个方法来获取工作的完成百分比。

3. `EventWriter`: 是一个用于向客户端发送进度事件的结构体。它提供了几个方法来报告不同类型的进度事件，如开始、进度更新、完成等。

4. `WorkDoneWriter`: 是一个用于报告工作进度的结构体。它提供了一些方法来报告工作的进度更新，如开始、更新、完成和取消。

以下是对每个函数的功能的详细介绍：

1. `NewTracker`: 创建一个新的进度跟踪器，并返回一个 `Tracker` 结构体的实例。

2. `SetSupportsWorkDoneProgress`: 设置进度跟踪器是否支持工作完成进度报告。

3. `SupportsWorkDoneProgress`: 检查进度跟踪器是否支持工作完成进度报告。

4. `Start`: 开始一个新的工作，并返回一个工作标识符 `Token`。

5. `Cancel`: 取消指定标识符的工作。

6. `Token`: 返回标识符 `Token` 的字符串表示。

7. `doCancel`: 执行取消操作，将取消信号发送给 `Tracker`。

8. `Report`: 向进度 `Tracker` 发送进度更新。

9. `End`: 结束指定标识符的工作。

10. `NewEventWriter`: 创建一个新的事件写入器，并返回一个 `EventWriter` 结构体的实例。

11. `Write`: 向客户端发送进度事件。

12. `NewWorkDoneWriter`: 创建一个新的工作进度写入器，并返回一个 `WorkDoneWriter` 结构体的实例。

以上是对于 `tools/gopls/internal/lsp/progress/progress.go` 文件中的结构体和函数的详细介绍。这些工具提供了一套方便的方法来跟踪和报告进度，以便在进行长时间或需要反馈给用户的任务时使用。

