# File: tools/gopls/internal/lsp/protocol/enums.go

文件enums.go的作用是定义与LSP（Language Server Protocol）相关的枚举类型和相关函数。

- namesTextDocumentSyncKind: 一个包含LSP文本文档同步类型的名称映射。
- namesMessageType: 一个包含LSP消息类型的名称映射。
- namesFileChangeType: 一个包含LSP文件变化类型的名称映射。
- namesWatchKind: 一个包含LSP文件监视类型的名称映射。
- namesCompletionTriggerKind: 一个包含LSP代码补全触发类型的名称映射。
- namesDiagnosticSeverity: 一个包含LSP诊断严重性类型的名称映射。
- namesDiagnosticTag: 一个包含LSP诊断标签类型的名称映射。
- namesCompletionItemKind: 一个包含LSP代码补全项类型的名称映射。
- namesInsertTextFormat: 一个包含LSP插入文本格式类型的名称映射。
- namesDocumentHighlightKind: 一个包含LSP文档高亮类型的名称映射。
- namesSymbolKind: 一个包含LSP符号类型的名称映射。
- namesTextDocumentSaveReason: 一个包含LSP文本文档保存原因类型的名称映射。

- init: 初始化上述所有的名称映射变量。
- formatEnum: 格式化指定的枚举类型并返回相应的字符串表示。
- parseEnum: 解析指定的字符串表示，并返回相应的枚举类型值。
- Format: 是一个通用的格式化函数，根据给定枚举类型和值返回其对应的字符串表示。
- ParseTextDocumentSyncKind: 解析字符串表示的文本文档同步类型并返回对应的值。
- ParseMessageType: 解析字符串表示的消息类型并返回对应的值。
- ParseFileChangeType: 解析字符串表示的文件变化类型并返回对应的值。
- ParseWatchKind: 解析字符串表示的文件监视类型并返回对应的值。
- ParseCompletionTriggerKind: 解析字符串表示的代码补全触发类型并返回对应的值。
- ParseDiagnosticSeverity: 解析字符串表示的诊断严重性类型并返回对应的值。
- ParseDiagnosticTag: 解析字符串表示的诊断标签类型并返回对应的值。
- ParseCompletionItemKind: 解析字符串表示的代码补全项类型并返回对应的值。
- ParseInsertTextFormat: 解析字符串表示的插入文本格式类型并返回对应的值。
- ParseDocumentHighlightKind: 解析字符串表示的文档高亮类型并返回对应的值。
- ParseSymbolKind: 解析字符串表示的符号类型并返回对应的值。
- ParseTextDocumentSaveReason: 解析字符串表示的文本文档保存原因类型并返回对应的值。

这些功能函数通过将字符串与名称映射中的对应关系进行匹配、查找、转换等操作，可以方便地进行LSP枚举类型的解析和格式化。

