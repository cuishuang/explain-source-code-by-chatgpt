# File: dirs.go

go/src/cmd/dirs.go文件的作用是实现命令`go env GOROOT GOTOOLDIR GOBIN`所需的功能。具体来说，该文件定义了以下几个函数：

1. `goroot`函数用于返回当前Go语言的根目录路径。
2. `goToolDir`函数用于返回当前Go语言的工具目录路径。
3. `goBin`函数用于返回当前Go语言的二进制目录路径。

这些函数的实现都依赖于操作系统环境变量或相关配置文件的设置。因此，在调用这些函数之前，会先调用其他函数如`findExecDir`，来尝试获取可执行文件所在的目录，从而推导出上述路径。

这些函数的主要作用是为使用Go语言的开发者提供一种方便且可靠的方式，来获取当前Go语言的根目录、工具目录和二进制目录等重要路径信息。这些信息对于开发者调试和编译Go语言项目非常重要，因此Go语言提供了这些内置函数来方便开发者使用。




---

### Var:

### dirs





### testGOPATH





### codeRootsCache





### usingModules





### modFlagRegexp








---

### Structs:

### Dir





### Dirs





### moduleJSON





## Functions:

### dirsInit





### goCmd





### Reset





### Next





### walk





### bfsWalkRoot





### codeRoots





### findCodeRoots





### vendorEnabled





### getMainModuleAnd114





