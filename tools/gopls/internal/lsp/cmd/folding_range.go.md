# File: tools/gopls/internal/lsp/cmd/folding_range.go

在Golang的Tools项目中，`folding_range.go`文件位于`tools/gopls/internal/lsp/cmd/`目录下。该文件的作用是实现了LSP（Language Server Protocol）的折叠范围相关的命令。

`foldingRanges`是`foldingRangeSource`接口的实现。该接口用于从源代码中提取可折叠范围。

以下是`foldingRanges`中的几个结构体及其作用：

1. `Name`：用于存储折叠范围的名称。
2. `Parent`：指示该折叠范围的上级范围。
3. `Usage`：用途描述。
4. `ShortHelp`：提供简短的帮助信息。
5. `DetailedHelp`：提供详细的帮助信息。
6. `Run`：执行具体的折叠范围操作。

`Run`函数中的具体逻辑会遍历文件中的代码，并根据特定的规则标记代码片段，以便将其折叠显示。这些折叠范围可以是函数、条件语句、循环等代码块。

`foldingRanges`结构体通过提供折叠范围的名称、父级范围、使用说明等信息，可以帮助编辑器在显示代码时进行折叠。

通过对`folding_range.go`文件中的这些结构体和函数的使用和调用，编辑器可以实现对代码的折叠显示功能，提供更好的代码阅读和编辑体验。

