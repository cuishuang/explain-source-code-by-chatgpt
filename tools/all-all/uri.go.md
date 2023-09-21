# File: tools/gopls/internal/span/uri.go

tools/gopls/internal/span/uri.go文件的主要作用是处理URI与文件路径之间的转换以及相关的辅助函数。

URI这几个结构体的作用如下：

1. `URI` 结构体用于表示一个标准的URI字符串，一般是一个绝对路径。

2. `URIError` 结构体用于表示URI相关的错误信息。

3. `URIRaw` 结构体用于表示原始的URI字符串，可以包含相对路径。

这些结构体提供了不同层次上处理URI的方式，以满足不同的需求。

以下是一些重要的函数及其作用：

1. `IsFile(uri URI) bool`：判断给定的URI是否是一个文件URI，即判断URI的scheme是否为"file"。

2. `Filename(uri URI) string`：从给定的URI中提取文件路径部分并返回字符串形式。

3. `filename(uri URI) string`：从给定的URI中提取文件路径中的文件名部分并返回字符串形式。

4. `URIFromURI(s string) (URI, error)`：根据给定的URI字符串创建一个URI对象，并返回。

5. `SameExistingFile(a, b URI) bool`：判断两个URI表示的文件是否相同。

6. `URIFromPath(path string) (URI, error)`：根据给定的文件路径创建一个URI对象，并返回。

7. `isWindowsDrivePath(path string) bool`：判断给定的文件路径是否是Windows驱动器路径。

8. `isWindowsDriveURIPath(uriPath string) bool`：判断给定的URI路径是否是Windows驱动器URI路径。

这些函数提供了方便的方法来处理URI和文件路径之间的转换，以及检查URI是否符合特定的规则。它们在 `span` 包中被广泛应用，用于处理代码分析和转换过程中的文件路径和URI操作。

