# File: tools/cmd/guru/guru.go

在Golang Tools项目中，tools/cmd/guru/guru.go是Guru工具的主要源代码文件。Guru是Golang静态代码分析工具，用于提供代码语义的查询功能，以便开发人员可以获得更好的理解和导航他们的代码。

该文件中定义了一些重要的结构体和函数，以下是它们的作用：

1. printfFunc结构体：表示要查询的printf样式函数的信息，包含参数名称、格式和位置等信息。
2. QueryResult结构体：表示代码查询结果的信息，包括查询的位置、结果描述、相关联的对象和全文等。
3. queryPos结构体：表示要查询的代码位置的信息，包括文件名、所在行号等。
4. Query结构体：表示一个代码查询的信息，包括查询位置和查询标记。

关于函数的作用，请参考下面的说明：

1. typeString函数：将给定类型转换为字符串表示形式。
2. objectString函数：将给定对象转换为字符串表示形式。
3. Run函数：Guru工具的主要入口函数，用于解析命令行参数并执行相应查询。
4. importQueryPackage函数：导入查询位置所在的包，并返回其包名。
5. pkgContainsFile函数：判断给定包是否包含指定的文件。
6. parseQueryPos函数：解析查询位置的字符串表示形式，返回queryPos结构体。
7. loadWithSoftErrors函数：加载包并允许包含可恢复的错误的文件。
8. containsHardErrors函数：判断包是否包含无法恢复的错误。
9. allowErrors函数：判断是否允许报告错误。
10. unparen函数：从表达式中删除括号。
11. deref函数：取指针类型的基础类型。
12. fprintf函数：将格式化的文本写入标准输出。
13. toJSON函数：将给定数据结构转换为JSON格式，并写入标准输出。

这些结构体和函数的组合是Guru工具的核心，它们提供了代码查询和分析的功能，帮助开发人员更好地理解和导航Golang代码。

