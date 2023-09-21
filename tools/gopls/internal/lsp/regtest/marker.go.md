# File: tools/gopls/internal/lsp/regtest/marker.go

在Golang的Tools项目中，tools/gopls/internal/lsp/regtest/marker.go这个文件的作用是定义了一些用于LSP（Language Server Protocol）测试的辅助函数和结构体。

- update：一个全局变量，用于标识是否更新测试结果。
- markerFuncs：一个包含各种类型标记函数的映射表。用于将特定类型的标记字符串转换为处理函数。
- goldenType：一个用于表示golden文件类型的枚举类型。包括GoldenMarker，GoldenPositionMarker和GoldenRangeMarker。
- locationType：一个用于表示location类型的枚举类型。包括LocationLine，LocationLineStart和LocationRange。
- markerType：一个用于表示marker类型的枚举类型。包括MarkerPoint和MarkerRange。
- regexpType：一个用于表示正则表达式的类型。用于在字符串中查找匹配的内容。
- wantErrorType：一个用于表示期望的错误类型的枚举类型。包括WantError，WantNoError和WantErrSnippet。

以下是一些结构体及其作用：

- marker：一个结构体，用于表示一个标记。包含文件路径、位置和标记类型等信息。
- markerTest：一个结构体，用于表示一个标记测试。包含了要测试的标记和期望的结果等信息。
- stringListValue：一个结构体，用于表示字符列表类型的值。包含一个字符串切片和是否需要换行输出的标志。
- Golden：一个结构体，用于表示测试中的golden文件。包含一个文件路径和文件内容。
- markerFunc：一个函数类型，用于处理标记字符串并返回一个marker对象。
- markerTestRun：一个函数类型，用于运行标记测试并返回测试结果。
- converter：一个函数类型，用于将特定类型的标记字符串转换为对应的数据结构。
- wantError：一个结构体，用于表示期望的错误结果，并包含一个处理错误的函数。

以下是一些函数及其作用：

- RunMarkerTests：运行一组标记测试。
- server：表示LSP服务器的接口对象。
- errorf：用于输出错误消息的辅助函数。
- execute：发送LSP请求并返回响应的辅助函数。
- flagSet：表示命令行标志的集合。
- Set：从命令行标志设置字符串的辅助函数。
- String：返回命令行标志设置字符串的辅助函数。
- getGolden：返回golden文件的内容。
- Get：获取路径对应的文件内容，并返回一个Golden对象。
- loadMarkerTests：加载标记测试的辅助函数。
- loadMarkerTest：加载单个标记测试的辅助函数。
- formatTest：格式化测试结果。
- newEnv：创建一个包含环境变量的副本。
- sprintf：格式化字符串的辅助函数。
- uri：表示LSP URI的类型。
- path：表示文件路径的类型。
- fmtPos：格式化位置的辅助函数。
- fmtLoc：格式化位置信息的辅助函数。
- fmtLocDetails：格式化位置信息和更多细节的辅助函数。
- makeMarkerFunc：创建一个用于处理标记字符串的处理函数。
- makeConverter：根据给定的标记类型创建一个转换函数。
- locationConverter：将位置标记字符串转换为位置信息对象的转换函数。
- findRegexpInLine：在给定的字符串中查找正则表达式匹配的辅助函数。
- linePreceding：返回指定行的前一行的内容的辅助函数。
- wantErrorConverter：将错误标记字符串转换为错误结果对象的转换函数。
- check：检查给定的标记是否符合期望的结果的辅助函数。
- goldenConverter：将golden标记字符串转换为golden对象的转换函数。
- checkChangedFiles：检查修改后的文件是否符合golden规范的辅助函数。
- completeMarker：表示代码补全的标记处理函数。
- acceptCompletionMarker：表示接受代码补全建议的标记处理函数。
- defMarker：表示定义的标记处理函数。
- foldingRangeMarker：表示折叠范围的标记处理函数。
- formatMarker：表示格式化的标记处理函数。
- highlightMarker：表示高亮显示的标记处理函数。
- hoverMarker：表示鼠标悬停的标记处理函数。
- locMarker：表示位置的标记处理函数。
- diagMarker：表示诊断消息的标记处理函数。
- removeDiagnostic：表示移除诊断消息的标记处理函数。
- renameMarker：表示重命名的标记处理函数。
- renameErrMarker：表示重命名错误的标记处理函数。
- rename：执行重命名操作的辅助函数。
- applyDocumentChanges：执行文档更改的辅助函数。
- codeActionMarker：表示代码操作的标记处理函数。
- codeActionErrMarker：表示代码操作错误的标记处理函数。
- codeAction：执行代码操作的辅助函数。
- codeLensesMarker：表示代码镜头的标记处理函数。
- collectExtraNotes：收集额外信息的辅助函数。
- suggestedfixMarker：表示推荐修复的标记处理函数。
- codeAction：执行代码操作的辅助函数。
- refsMarker：表示引用的标记处理函数。
- implementationMarker：表示实现的标记处理函数。
- symbolMarker：表示符号的标记处理函数。
- compareLocations：比较位置信息的辅助函数。
- workspaceSymbolMarker：表示工作区符号的标记处理函数。

以上是marker.go文件中定义的主要函数、变量和结构体，用于支持LSP测试。

