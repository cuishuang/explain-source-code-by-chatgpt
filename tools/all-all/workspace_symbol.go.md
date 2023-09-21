# File: tools/gopls/internal/lsp/cmd/workspace_symbol.go

文件 `workspace_symbol.go` 的作用是定义和实现 LSP（Language Server Protocol）中的 workspace symbol 相关的功能。

`workspaceSymbol` 这几个结构体的作用如下：
- `Name`：表示一个符号的名称。
- `Parent`：表示一个符号的父级符号。
- `Usage`：表示一个符号的用法。
- `ShortHelp`：表示一个符号的简短帮助信息。
- `DetailedHelp`：表示一个符号的详细帮助信息。
- `Run`：表示执行一个操作。

这些结构体用于存储和表示一个符号在 workspace 中的信息。

`Name` 代表符号的名称，可以是函数名、变量名等。`Parent` 表示符号的父级符号，例如在一个函数内部，`Parent` 可能是函数本身。`Usage` 描述了符号的用法，可以帮助开发者理解如何正确使用该符号。

`ShortHelp` 是符号的简短帮助信息，通常包含了符号的摘要描述。`DetailedHelp` 则提供了更详细的帮助信息，包括符号的详细描述、示例代码等。这些信息可以帮助开发者更好地了解符号的用途和功能。

`Run` 函数用于执行一个操作，例如在命令行中执行一个特定的命令。这个函数通常与符号相关联，以便开发者可以直接在命令行中执行某些与符号相关的操作，比如查看文档、跳转到定义等。

总的来说，`workspace_symbol.go` 文件定义了 workspace symbol 功能的相关结构体和函数，用于提供符号的查询、显示和操作功能。

