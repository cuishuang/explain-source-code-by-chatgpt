# File: mkbuiltin.go

mkbuiltin.go文件是Go语言中的一个工具，用于生成内置函数表和类型表。Go语言编译器会使用这些表来检查和优化代码。

该工具会扫描Go语言标准库中的所有文件，并将其中的内置函数和类型信息提取出来，然后生成一个名为“builtin”的Go包。在编译器中，当遇到使用内置函数或类型的代码时，编译器会查找这个包中相应的函数或类型。

此外，该工具还会检查内置函数和类型的元数据，并将其保存在Go语言的文档中。这些文档可以通过godoc命令查看，在Go语言的文档中心也可以查看。

总之，mkbuiltin.go文件是Go语言编译器的重要组成部分，它可以提高代码的性能和可读性，也可以方便用户查找和使用内置函数和类型。




---

### Var:

### stdout





### nofmt








---

### Structs:

### typeInterner





## Functions:

### main





### mkbuiltin





### intern





### subtype





### mktype





### fields





### intconst





