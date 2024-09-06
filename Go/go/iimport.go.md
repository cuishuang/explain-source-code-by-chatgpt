# File: iimport.go

iimport.go 是 Go 语言编译器（go）的一个重要文件，主要用于解析和处理 Go 程序中的导入语句，包括解析、分析和处理程序中使用的所有依赖库。

具体来说，iimport.go 通过以下方式实现：

1. 解析 import 语句：iimport.go 会对 Go 语言程序中的 Import 语法进行解析，将其转换为相应的导入组件并生成相关的AST（抽象语法树）。

2. 分析导入组件：iimport.go 会对所有导入文件进行分析并加载，找出所有需要使用的包及其依赖关系，并生成包的信息描述符（pkg）。

3. 处理依赖库：iimport.go 会递归地处理所有依赖库，并针对每个包生成一个唯一的版本号，并将其与全局包路径建立映射关系，以便在编译时正确地查找并使用不同版本的包。

4. 链接依赖库：iimport.go 会将程序中所有的依赖库编译为目标文件，并将其链接到主程序文件中以创建完整的可执行程序。

总之，iimport.go 的主要作用是管理和处理 Go 语言程序中的所有依赖库，并将其转换为可执行的二进制文件。这个过程是 Go 语言编译器的重要组成部分，保证了 Go 语言程序的可靠性、可移植性和可扩展性。




---

### Structs:

### intReader





### ident





### itag





### setConstraintArgs





### iimporter





### importReader





## Functions:

### int64





### uint64





### iImportData





### doDecl





### stringAt





### pkgAt





### typAt





### canReuse





### obj





### declare





### value





### intSize





### mpint





### mpfloat





### ident





### qualifiedIdent





### pos





### posv0





### posv1





### typ





### isInterface





### pkg





### string





### doType





### kind





### signature





### tparamList





### paramList





### param





### bool





### int64





### uint64





### byte





### baseType





### tparamName





