# File: tools/cmd/guru/definition.go

在Golang的Tools项目中，位于tools/cmd/guru/definition.go文件是guru工具中的一部分，它实现了查找Go代码中标识符的定义的功能。

该文件中定义了几个结构体以及相关的函数，以下是对它们的详细介绍：

1. definitionResult 结构体：表示一个标识符的定义结果。它包含了标识符的名称、所在的包、位置信息以及其他相关的信息。

2. definition 结构体：表示一个标识符的定义。它继承了definitionResult，并添加了定义的类型信息。这个结构体是guru定义功能的核心。

3. packageForQualIdent 结构体：表示一个限定符标识符（如"pkg.Ident"）所对应的包的信息。包含了包的路径、名称等信息。

4. findPackageMember() 函数：根据限定符标识符查找其所属的包。它会解析限定符的包名，并返回packageForQualIdent结构体。

5. PrintPlain() 函数：以纯文本的形式打印标识符的定义信息。它接收一个*definitionResult参数，并将定义信息打印到标准输出。

6. JSON() 函数：以JSON格式的形式返回标识符的定义信息。它接收一个*definitionResult参数，并将定义信息以JSON字符串的形式返回。

这些函数提供了定义查询相关的功能，可以根据标识符的名称、位置等信息，查找并返回其定义的详细信息。同时，通过PrintPlain()函数和JSON()函数，可以以不同的格式输出定义信息，方便使用者进行阅读和处理。

