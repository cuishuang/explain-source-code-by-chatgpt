# File: filter.go

filter.go文件位于Go语言标准库（go/src）的path包中，主要用于实现对Unix风格路径名的解析和构造。

在Unix系统中，文件路径名通常是由若干个名称（name）组成，这些名称被斜杠（/）分隔。例如，/usr/local/go/src/path/filter.go就是一个典型的Unix路径名。

filter.go文件中定义了以下重要的类型和函数：

类型：Path

Path类型表示一个路径名，它由多个名称组成，每个名称之间用斜杠分隔。Path类型定义了多个方法，包括：

- func (p Path) String() string：将Path类型转换成字符串类型，返回路径名。
- func (p Path) Base() string：返回路径名中的最后一个名称，即文件名。

函数：Clean、Join和Split

Clean函数用于规范化一个路径名，将其转换成标准格式。例如，/usr/local/go/../a/b/c变为/usr/local/a/b/c。

Join函数用于将多个名称拼接成一个完整的路径名，可以自动添加斜杠分隔符，也能够处理各种特殊情况和错误。

Split函数用于将一个路径名分解成多个名称，返回一个字符串切片。例如，/usr/local/go/src/path/filter.go被分解成/usr、local、go、src、path和filter.go。

总的来说，filter.go文件提供了基础性的路径名解析、构造和处理功能，方便用户在Unix系统上操作文件和目录。




---

### Var:

### separator








---

### Structs:

### Filter





### MergeMode





## Functions:

### exportFilter





### FileExports





### PackageExports





### filterIdentList





### fieldName





### filterFieldList





### filterCompositeLit





### filterExprList





### filterParamList





### filterType





### filterSpec





### filterSpecList





### FilterDecl





### filterDecl





### FilterFile





### filterFile





### FilterPackage





### filterPackage





### nameOf





### MergePackageFiles





