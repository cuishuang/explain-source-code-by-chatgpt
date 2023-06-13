# File: sym.go

sym.go这个文件是Go语言编译器的一部分，主要用于管理程序中的符号（symbols）。

在编程语言中，符号通常指代变量名、函数名等标识符，Symbol Table是指存储这些标识符的数据结构。在Go语言编译过程中，编译器会把程序中的符号都存储到Symbol Table中，以便在后续的编译、链接过程中能够快速地查找和处理这些标识符。

sym.go文件中定义了一系列结构体和函数，用于表示和管理符号表中的各种信息，例如符号的类型、大小、对齐等。

具体来说，sym.go文件中的一些重要函数包括：

- Addrel：用于将一个符号与一个重定位（relocation）信息关联起来，在链接时使用。
- Lookup：在符号表中查找一个给定名称的符号，并返回其引用。如果该符号不存在，则创建一个新的符号。
- Addstring：向符号表中添加一个字符串符号。
- Sbrk：向符号表中分配一块额外的内存空间，以存储一些特殊符号。

另外，sym.go文件还定义了一些常量，如符号表中符号的大小、类型等，这些常量在编译过程中用于确定和处理不同种类的符号。




---

### Structs:

### Sym





## Functions:

### OnExportList





### Uniq





### Siggen





### Asm





### Func





### SetOnExportList





### SetUniq





### SetSiggen





### SetAsm





### SetFunc





### IsBlank





### Linksym





### LinksymABI





### Less





### IsExported





