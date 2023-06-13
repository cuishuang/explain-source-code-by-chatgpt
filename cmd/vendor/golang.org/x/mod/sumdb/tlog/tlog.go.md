# File: tlog.go

go/src/cmd/tlog.go是一个用于生成Go语言中的调试日志的工具。tlog是“trace log”的缩写，它能够在代码执行期间记录各种信息，例如函数调用、参数值、返回值等，并将这些信息输出到指定的日志文件中。这可以帮助开发人员更好地理解代码的执行过程，从而更容易地调试代码。

tlog的使用非常简单，只需要在代码中使用tlog.Trace函数进行标记即可。例如：

```go
func foo(x int) int {
    tlog.Trace("calling foo with argument x=%d", x)
    result := x * 2
    tlog.Trace("foo returned %d", result)
    return result
}
```

在代码执行过程中，tlog会自动输出相应的日志信息，例如：

```
[ 2017/09/18 20:20:01 ] calling foo with argument x=42
[ 2017/09/18 20:20:01 ] foo returned 84
```

tlog支持多种不同的日志输出格式，并且可以根据需要自定义输出格式。除了Trace函数之外，tlog还提供了一些其他的便利函数，例如Debug、Info、Warn和Error函数，用于记录不同级别的日志信息。

