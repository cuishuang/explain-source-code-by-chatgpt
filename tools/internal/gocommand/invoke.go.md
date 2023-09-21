# File: tools/internal/gocommand/invoke.go

文件 `invoke.go` 的作用是定义了调用Go命令的函数和相关的数据结构。

- `modConcurrencyError` 是一个全局变量，用于控制在运行 `go mod` 命令时是否进行并发更新。
- `verb` 是一个全局变量，用于存储 golang 命令的当前运行动作。
- `DebugHangingGoCommands` 是一个全局变量，用于控制是否打印输出调试信息。

`Runner` 结构体表示一个运行命令行操作的实例，它包含了执行系统命令所需的参数和环境信息。
`Invocation` 结构体表示一个执行的命令和它的参数。

以下是该文件中的主要函数介绍：

- `initialize()` 函数用于初始化一些全局变量。
- `invLabels()` 函数返回一个用于显示命令行输出时的标签。
- `Run()` 函数是执行一个命令行操作的入口函数。
- `RunPiped()` 函数用于运行一个管道命令行操作。
- `RunRaw()` 函数用于运行一个不带任何处理的原始命令行操作。
- `runConcurrent()` 函数用于并发运行多个命令行操作。
- `runPiped()` 函数用于运行一个管道命令行操作。
- `runWithFriendlyError()` 函数用于以友好的方式运行一个命令行操作，并处理可能出现的错误。
- `run()` 函数用于运行一个命令行操作，并处理可能出现的错误。
- `runCmdContext()` 函数用于运行一个带有上下文环境的命令行操作。
- `HandleHangingGoCommand()` 函数用于处理因 Go 命令卡住而无法退出的情况，以及显示调试信息。
- `cmdDebugStr()` 函数返回一个包含命令行的调试信息的字符串。

