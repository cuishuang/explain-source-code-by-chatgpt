# File: tools/gopls/internal/lsp/regtest/expectation.go

文件expectation.go定义了一些用于测试gopls（Go语言的LSP实现）的期望行为的数据结构和函数。

1. InitialWorkspaceLoad结构体表示初始工作区加载的期望行为。该结构体包含了一个LoadFiles字段，表示在初始加载期间需要加载的文件列表。

2. Verdict结构体表示测试结果的判断，包含一个Passed字段表示是否通过测试，并且可以附带一些额外的消息。

3. Expectation结构体表示对于特定操作的期望行为，包含了一个Name字段表示操作的名称（例如"workspace/didChangeConfiguration"），以及一个Verdict字段表示该操作的测试结果。

4. WorkStatus结构体表示对于工作区的状态的期望。包含一个Empty字段，表示工作区是否为空；一个OutstandingJobs字段，表示未完成的任务数量。

5. flatDiagnostic结构体表示一个扁平的诊断信息，包含了诊断的文件路径、行号、列号以及其他相关信息。

6. DiagnosticFilter结构体表示对诊断信息进行筛选和过滤的条件，比如根据文件路径或者诊断消息内容进行筛选。

7. 这个文件还定义了许多函数，用于构建和检查测试中的各种期望行为。

- String函数用于将期望行为的结果转换为字符串。
- OnceMet函数用于确定在一个切片中是否存在至少一个满足条件的元素。
- describeExpectations函数用于将期望行为的描述转换为字符串。
- Not函数用于对一个期望行为取反。
- AnyOf函数用于检查多个期望行为中是否有至少一个满足条件的。
- AllOf函数用于检查多个期望行为中是否全部满足条件。
- ReadDiagnostics函数用于从诊断的文件路径和内容中读取诊断信息。
- ReadAllDiagnostics函数用于从诊断的文件路径和内容列表中读取所有的诊断信息。
- NoOutstandingWork函数用于检查工作区是否没有未完成的任务。
- ShownDocument函数用于检查是否向用户显示了一个文档。
- NoShownMessage函数用于检查是否没有向用户显示消息。
- ShownMessage函数用于检查是否向用户显示了特定的消息。
- ShowMessageRequest函数用于检查是否发送了一个显示消息的请求。
- DoneDiagnosingChanges函数用于检查是否完成了对变更的诊断。
- AfterChange函数用于检查在对变更进行诊断之后是否触发了其他操作。
- DoneWithOpen函数用于检查是否完成了对打开操作的处理。
- StartedChange函数用于检查是否开始了对变更的处理。
- DoneWithChange函数用于检查是否完成了对变更的处理。
- DoneWithSave函数用于检查是否完成了对保存操作的处理。
- StartedChangeWatchedFiles函数用于检查是否开始了对变更监视的处理。
- DoneWithChangeWatchedFiles函数用于检查是否完成了对变更监视的处理。
- DoneWithClose函数用于检查是否完成了对关闭操作的处理。
- StartedWork函数用于检查是否开始了某种工作。
- CompletedWork函数用于检查是否完成了某种工作。
- CompletedProgress函数用于检查是否完成了某个进度通知。
- OutstandingWork函数用于检查是否有未完成的任务。
- NoErrorLogs函数用于检查是否没有错误日志。
- LogMatching函数用于检查是否有匹配特定条件的错误日志。
- NoLogMatching函数用于检查是否没有匹配特定条件的错误日志。
- FileWatchMatching函数用于检查是否匹配了特定条件的文件监视。
- NoFileWatchMatching函数用于检查是否没有匹配特定条件的文件监视。
- checkFileWatch函数用于检查是否匹配了特定条件的文件监视。
- jsonProperty函数用于从JSON中获取特定字段的值。
- Diagnostics函数用于检查是否存在特定的诊断信息。
- NoDiagnostics函数用于检查是否不存在特定的诊断信息。
- flattenDiagnostics函数用于将多个诊断信息展平到一个列表中。
- ForFile函数用于获取给定路径的诊断信息。
- FromSource函数用于从一个源字符串中获取诊断信息。
- AtRegexp函数用于从给定的字符串正则表达式处获取位置。
- AtPosition函数用于从给定的行列坐标处获取位置。
- WithMessage函数用于设置期望的诊断消息。
- WithSeverityTags函数用于设置期望的诊断严重性标签。

这些函数以及上述的数据结构与测试gopls期望行为相关的结果进行比较和验证，用于测试gopls的正确性和可靠性。

