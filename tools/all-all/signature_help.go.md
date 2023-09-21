# File: tools/gopls/internal/lsp/source/signature_help.go

文件 signature_help.go 是 gopls 工具项目中的一部分，它提供了关于函数参数提示的功能。

SignatureHelp 是一个结构体，包含了函数的签名信息，以及光标所在位置的参数提示。

builtinSignature 是一个类型，存储了内置函数的签名信息。

activeParameter 是一个函数，用于确定当前光标所在的参数位置。

stringToSigInfoDocumentation 是一个函数，用于将函数签名信息和文档字符串转换为 SignatureInformation 和 Documentation 数据类型。

具体来说，SignatureHelp 结构体中的 Signature 字段是一个切片，每个元素包含了一个函数的参数和返回值的签名信息。Documentation 字段则是当前函数的文档字符串。光标所在的参数位置会通过 activeParameter 函数得到。

builtinSignature 存储了一些内置函数的签名信息，这些内置函数可以在代码补全和参数提示中使用。

activeParameter 函数会根据光标的位置，判断当前正在输入的参数是哪个。这个函数会遍历当前参数的前后字符，并根据 Go 语言的语法规则判断参数的位置。

stringToSigInfoDocumentation 函数会将函数的签名信息和文档字符串转换为 SignatureInformation 和 Documentation 类型。SignatureInformation 包含了函数的参数信息以及相应的文档字符串，Documentation 则表示函数的详细文档。

总之，signature_help.go 这个文件提供了函数参数提示的功能，能够根据光标的位置和函数的签名信息，在编辑器中显示函数参数的提示信息。

