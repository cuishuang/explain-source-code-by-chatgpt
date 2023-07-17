# File: pkg.go

pkg.go 是 Go 语言标准库中 cmd 目录下的一个文件，其主要作用是导出 Go 命令中的各个子命令，让其可以在命令行界面中被执行。

具体来说，pkg.go 文件定义了一个名为 main 的 package，其中包含了 Go 命令的主函数和各个子命令的导出函数。通过导出函数，pkg.go 文件将 Go 命令中的子命令全部暴露给了外界，使得这些命令可以在命令行中被执行。

在 pkg.go 文件中，还定义了一些全局变量和常量，用于设置 Go 命令的默认行为和输出格式等。此外，pkg.go 文件中也提供了一些函数，用于处理命令行参数、输出帮助信息等。

总之，pkg.go 文件是 Go 语言标准库中的一个重要文件，它承担了 Go 命令中的核心功能。了解 pkg.go 文件的作用，有助于深入理解 Go 命令的实现原理，并能帮助开发者自定义和扩展 Go 命令。




---

### Var:

### pkgMap





### nopkg





### internedStringsmu





### internedStrings








---

### Structs:

### Pkg





### byPath





## Functions:

### NewPkg





### ImportedPkgList





### Len





### Less





### Swap





### Lookup





### LookupOK





### LookupBytes





### LookupNum





### InternString





### CleanroomDo





