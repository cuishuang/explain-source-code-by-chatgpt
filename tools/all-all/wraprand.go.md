# File: tools/cmd/signature-fuzzer/internal/fuzz-generator/wraprand.go

在Golang的Tools项目中，`tools/cmd/signature-fuzzer/internal/fuzz-generator/wraprand.go`文件的作用是实现一个包装了`math/rand`的随机数生成器。

`NewWrapRand`函数用于创建一个新的`WrapRand`结构体实例，该结构体继承了`math/rand.Rand`的所有方法，并添加了一些新方法。

`captureCall`函数用于捕获调用`math/rand`的函数及其参数，并将其记录在`WrapRand`结构体的`calls`字段中。

`Intn`函数用于代理调用`math/rand.Intn`函数，并通过`captureCall`记录这次调用。

`Float32`函数用于代理调用`math/rand.Float32`函数，并通过`captureCall`记录这次调用。

`NormFloat64`函数用于代理调用`math/rand.NormFloat64`函数，并通过`captureCall`记录这次调用。

`emitCalls`函数用于将所有被记录的调用信息生成一个字符串，并将字符串返回。

`Equal`函数用于比较两个`WrapRand`结构体是否相等，即比较它们所记录的调用信息是否完全一致。

`Check`函数用于检查给定的`WrapRand`结构体是否应该触发panic。

`Checkpoint`函数用于在`WrapRand`结构体的`calls`字段中添加一个检查点，用于验证后续的调用顺序是否正确。

通过这些方法，`WrapRand`结构体可以记录所有调用情况，并方便对随机数生成的过程进行分析和测试。

