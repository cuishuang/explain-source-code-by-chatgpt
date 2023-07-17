# File: library.go

library.go是Go语言标准库的一个文件，用于管理和加载Go语言库的代码。

具体来说，library.go定义了两个重要的数据结构：

1. Package：表示一个Go语言包，包含了包名、导入路径、所在目录等信息。
2. ImportStack：用于加载Go语言库时跟踪已经被导入的包。

同时，library.go还定义了几个重要的函数：

1. importPackage：用于加载一个Go语言包，并将其加入到ImportStack中。
2. importFromFiles：用于从本地文件中加载Go语言库。
3. importPath：用于根据包名和导入路径导入一个Go语言包。
4. loadImports：用于加载所有被导入的Go语言包，并按照依赖关系排序。

这些函数共同实现了Go语言库的加载和管理，使得开发者可以轻松地使用和扩展标准库。

