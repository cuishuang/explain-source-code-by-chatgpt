# File: tools/gopls/internal/lsp/protocol/tsserver.go

tools/gopls/internal/lsp/protocol/tsserver.go文件是Golang的tools/gopls项目中实现了与TypeScript服务器通信的协议。

该文件定义了TSServer结构体和一系列与TypeScript服务器通信的方法。下面是对文件中一些重要结构体和方法的介绍：

1. Server结构体：表示TSServer服务器。它包含了与TypeScript服务器通信的连接以及发送和接收消息的方法。

2. serverDispatch方法：根据收到的消息类型，调用对应的方法进行处理。它是服务器接收到客户端请求的入口函数。

3. Progress方法：处理服务器发送来的进度通知。它用于向客户端报告长时间运行的操作的进度。

4. SetTrace方法：设置跟踪级别。它用于通知TypeScript服务器要在哪个级别上输出调试信息。

5. IncomingCalls和OutgoingCalls方法：处理对内部函数的调用和对外部函数的调用。它们是为了实现函数调用相关的功能。

6. ResolveCodeAction、ResolveCodeLens、ResolveCompletionItem、ResolveDocumentLink等方法：处理对代码操作的解析请求。它们将解析并返回与代码操作相关的详细信息。

7. Initialize、Initialized、Resolve、DidChangeNotebookDocument、DidCloseNotebookDocument等方法：处理与项目初始化、修改、关闭等相关的请求。

8. Shutdown方法：处理服务器关闭请求。它用于关闭与TypeScript服务器的连接并清理资源。

9. Diagnostic、DocumentSymbol、Hover、Completion、References、CodeLens等方法：处理与代码编辑和分析相关的请求。

10. Symbol、ExecuteCommand、WorkspaceSymbol、Rename、DidChangeConfiguration、DidChangeWatchedFiles等方法：处理与工作区和符号相关的请求。

11. NonstandardRequest方法：处理非标准请求。它用于处理一些自定义的请求，对于标准请求无法满足的情况。

这些方法在Golang的tools/gopls项目中与TypeScript服务器进行通信，并处理各种客户端的请求。通过这些方法，可以实现代码编辑、分析、跳转等功能。

