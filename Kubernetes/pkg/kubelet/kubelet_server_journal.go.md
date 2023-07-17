# File: pkg/kubelet/kubelet_server_journal.go

pkg/kubelet/kubelet_server_journal.go文件的作用是实现kubelet server的journal日志功能。该文件包含了与journal相关的操作函数和结构体。

- journal是一个字符串常量，指定了journal的位置。
- reServiceNameUnsafeCharacters是一个正则表达式，用于过滤不安全的字符。
- journalServer是一个结构体，包含了journal server的配置信息。
- nodeLogQuery是一个结构体，表示获取节点日志的查询。
- options是一个结构体，包含了journal server的选项。
- readerCtx是一个结构体，表示journal读取上下文。

以下是各个函数的作用：

- ServeHTTP函数用于处理HTTP请求，接受了journalServer作为参数。
- newNodeLogQuery函数用于创建一个新的节点日志查询。
- validateServices函数用于验证服务。
- validate函数用于验证journal server配置。
- Copy函数用于复制journal。
- copyForBoot函数用于为引导过程复制journal。
- splitNativeVsFileLoggers函数用于分割本地日志和文件日志。
- copyServiceLogs函数用于复制服务日志。
- copyFileLogs函数用于复制文件日志。
- heuristicsCopyFileLogs函数用于通过启发式方法复制文件日志。
- Read函数用于读取journal日志。
- newReaderCtx函数用于创建一个新的journal读取上下文。
- heuristicsCopyFileLog函数用于通过启发式方法复制文件日志。
- safeServiceName函数用于过滤不安全的服务名称。

