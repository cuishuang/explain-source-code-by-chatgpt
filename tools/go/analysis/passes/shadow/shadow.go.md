# File: tools/go/analysis/passes/shadow/shadow.go

在Golang的Tools项目中，tools/go/analysis/passes/shadow/shadow.go是一个用于检测变量和导入阴影（shadowing）的分析器。阴影指的是在同一个作用域内使用了相同的名称，导致一个实体（变量或导入）覆盖了另一个。

该文件中的doc变量是一个包级别的字符串文档，用于提供有关该分析器的信息。Analyzer变量是一个提供分析器的注册实例，用于将该分析器添加到Golang Tools中。strict变量是一个布尔值，用于指定是否启用严格模式，即仅在导入阴影时报告错误。

关于span，它们是该分析器中使用的结构体，表示代码中的位置范围。有三个不同的span结构体：span, spanStack和shadowStack。

- span结构体表示代码中的一个位置范围，包含文件名、行号和列号等信息。

- spanStack结构体是span结构体的堆栈，用于表示不同的位置范围的层次结构。

- shadowStack结构体是一个记录所有阴影变量的堆栈，用于在分析过程中跟踪阴影变量的定义和使用。

在该文件中，还定义了一些函数：

- init函数用于初始化该分析器，设置其名称、文档和其他属性。

- run函数是实际执行分析的函数，它接收AST（抽象语法树）和通过分析代码报告阴影变量。

- contains函数用于检查一个span是否包含另一个span，即用于判断两个位置范围之间的嵌套关系。

- growSpan函数用于扩展给定的span，将其范围延伸到包含另一个span。

- checkShadowAssignment函数用于检查阴影变量的赋值语句。

- idiomaticShortRedecl函数用于检查阴影变量的短声明构造。

- idiomaticRedecl函数用于检查阴影变量的声明。

- checkShadowDecl函数用于检查阴影变量的声明语句。

- checkShadowing函数是一个递归函数，用于检查所有的阴影变量。

这些函数共同工作，以检测和报告代码中出现的阴影变量，并根据配置的strict标记确定是否报告导入阴影。

