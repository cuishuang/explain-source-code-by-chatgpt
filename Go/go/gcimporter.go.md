# File: gcimporter.go

gcimporter.go这个文件的作用是解析Go语言的编译后二进制包文件，并将其转换成抽象语法树（AST），以便在编译Go程序时使用。

具体来说，当我们使用Go语言中的import语句导入一个二进制包文件时，这个包文件会被传递到gcimporter.go中进行解析。gcimporter.go会读取这个二进制包文件中的信息，并将其转换成抽象语法树，然后将这些信息传递给Go编译器，用于生成最终的可执行程序。

此外，gcimporter.go还提供了一些与导入、解析和编译二进制包相关的功能，比如处理导入路径、解析标识符、获取程序包信息等。

总之，gcimporter.go是Go语言编译器的一个重要组成部分，负责将导入的二进制包文件解析为Go程序可以识别的抽象语法树，是Go语言实现模块化编程的重要基础。




---

### Var:

### exportMap





### pkgExts








---

### Structs:

### byPath





## Functions:

### lookupGorootExport





### FindPkg





### Import





### Len





### Swap





### Less





