# File: tools/gopls/internal/lsp/server_gen.go

`tools/gopls/internal/lsp/server_gen.go`文件是Gopls（Go语言的LSP服务器）中自动生成的文件，其作用是根据LSP规范定义的不同请求和通知方法生成相应的处理逻辑。

以下是对所列出的函数的作用的简要介绍：

- CodeAction: 提供代码操作（如重构、修复等）建议的列表。
- CodeLens: 提供代码镜头（CodeLens）的信息。
- ColorPresentation: 提供给定颜色值的呈现方式的信息。
- Completion: 提供代码补全的建议列表。
- Declaration: 提供指定位置的声明的位置信息。
- Definition: 提供指定位置的定义的位置信息。
- Diagnostic: 发送关于文档中的问题和错误的诊断信息。
- DiagnosticWorkspace: 发送关于工作区中的问题和错误的诊断信息。
- DidChange: 通知服务器指定文档已更改。
- DidChangeConfiguration: 通知服务器客户端的配置已更改。
- DidChangeNotebookDocument: 通知服务器指定笔记本文档已更改。
- DidChangeWatchedFiles: 通知服务器关于文件系统中的文件更改的信息。
- DidChangeWorkspaceFolders: 通知服务器工作区的文件夹列表已更改。
- DidClose: 通知服务器指定文档已关闭。
- DidCloseNotebookDocument: 通知服务器指定笔记本文档已关闭。
- DidCreateFiles: 通知服务器有文件被创建。
- DidDeleteFiles: 通知服务器有文件被删除。
- DidOpen: 通知服务器指定文档已打开。
- DidOpenNotebookDocument: 通知服务器指定笔记本文档已打开。
- DidRenameFiles: 通知服务器有文件被重命名。
- DidSave: 通知服务器指定文档已保存。
- DidSaveNotebookDocument: 通知服务器指定笔记本文档已保存。
- DocumentColor: 提供文档中颜色值的信息。
- DocumentHighlight: 提供指定位置的文档高亮信息。
- DocumentLink: 提供文档中链接的信息。
- DocumentSymbol: 提供文档中符号（函数、变量等）的信息。
- ExecuteCommand: 在服务器上执行命令。
- Exit: 发送LSP服务器退出信息。
- FoldingRange: 提供文档中折叠范围的信息。
- Formatting: 提供文档格式化的建议。
- Hover: 提供指定位置的悬停信息。
- Implementation: 提供指定位置的实现的位置信息。
- IncomingCalls: 提供指定位置的调用者的信息。
- Initialize: 客户端向服务器发送初始化请求。
- Initialized: 服务器向客户端发送初始化完成的通知。
- InlayHint: 提供文档中内嵌提示的信息。
- InlineCompletion: 提供在代码中内联的补全建议。
- InlineValue: 提供指定位置变量的内联值信息。
- LinkedEditingRange: 提供指定位置的链接编辑范围的信息。
- Moniker: 提供符号的唯一标识（Moniker）信息。
- NonstandardRequest: 处理非标准的请求。
- OnTypeFormatting: 提供在键入时格式化代码的建议。
- OutgoingCalls: 提供指定位置的被调用者的信息。
- PrepareCallHierarchy: 提供指定位置的调用层次结构的信息。
- PrepareRename: 准备对指定位置进行重命名。
- PrepareTypeHierarchy: 提供指定位置的类型层次结构的信息。
- Progress: 发送操作进度的信息。
- RangeFormatting: 提供指定范围内代码格式化的建议。
- RangesFormatting: 提供多个范围内代码格式化的建议。
- References: 提供指定位置的引用的位置信息。
- Rename: 对指定位置进行重命名。
- Resolve: 在提供的选项之间进行选择。
- ResolveCodeAction: 处理代码操作的解析。
- ResolveCodeLens: 处理代码镜头的解析。
- ResolveCompletionItem: 处理代码补全建议的解析。
- ResolveDocumentLink: 处理文档链接的解析。
- ResolveWorkspaceSymbol: 处理工作区符号的解析。
- SelectionRange: 提供指定位置的选项范围的信息。
- SemanticTokensFull: 提供完整的语义令牌信息。
- SemanticTokensFullDelta: 提供语义令牌的增量信息。
- SemanticTokensRange: 提供指定范围内的语义令牌信息。
- SetTrace: 设置LSP服务器的跟踪级别。
- Shutdown: 通知服务器关闭。
- SignatureHelp: 提供指定位置的函数、方法签名的信息。
- Subtypes: 提供指定类型的子类型的信息。
- Supertypes: 提供指定类型的父类型的信息。
- Symbol: 提供指定位置的符号信息。
- TypeDefinition: 提供指定位置的类型定义位置信息。
- WillCreateFiles: 通知服务器将要创建文件。
- WillDeleteFiles: 通知服务器将要删除文件。
- WillRenameFiles: 通知服务器将要重命名文件。
- WillSave: 通知服务器指定文档将要保存。
- WillSaveWaitUntil: 通知服务器指定文档将要保存，并返回保存结果。
- WorkDoneProgressCancel: 取消操作进度。

这些函数定义了LSP服务器的不同功能和处理能力，根据客户端的请求，服务器将根据这些函数提供相应的处理逻辑。

