# File: istio/pilot/pkg/security/authz/builder/logger.go

在Istio项目中，`logger.go`文件的作用是构建授权日志，并提供将日志输出到不同目标的功能。

`authzLog`是一个提供日志记录的全局变量，用于记录授权相关的日志信息。它包含了三个字段：
- `debug`：一个布尔值，表示是否启用调试日志级别。
- `out`：一个`io.Writer`接口，表示日志输出目标。
- `errOut`：一个`io.Writer`接口，表示错误日志输出目标。

`AuthzLogger`结构体是授权日志记录器，提供了四个方法：
- `AppendDebugf(format string, args ...interface{})`：将格式化的调试信息添加到日志。
- `AppendError(err error)`：将错误信息添加到日志。
- `Report() string`：收集并以字符串形式返回当前日志的内容。
- `SetOutput(output io.Writer)`：设置日志输出目标。

`AppendDebugf`方法用于将格式化的调试信息追加到日志中，以方便调试和排查问题。

`AppendError`方法用于将错误信息追加到日志中，以记录授权过程中出现的错误。

`Report`方法用于收集当前日志的内容，并以字符串形式返回，用于输出或记录日志。

这些方法结合使用，可以将授权相关的调试和错误信息记录到日志中，并根据需要将日志输出到不同的目标，例如控制台、文件或网络等。这样可以帮助开发者跟踪授权过程中的问题，并对其进行调试和优化。

