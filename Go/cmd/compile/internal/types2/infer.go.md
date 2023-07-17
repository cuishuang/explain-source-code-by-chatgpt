# File: infer.go

infer.go是一个用于Go源代码推理的工具，它可以识别和解析Go源代码，并生成各种类型的输出，例如：

- 依赖关系：显示一个Go程序中的所有依赖项和它们之间的关系
- 调用图：显示Go程序中的函数调用结构
- 接口实现：显示Go中接口和结构体之间的实现关系
- 控制流：显示Go程序的控制流结构

infer.go的基本原理是解析Go AST（抽象语法树）并使用类型信息来推导程序的行为。它还利用Go的工具链和内部信息来处理import路径和GOPATH（Go项目的路径）等细节。它可以作为命令行工具使用，也可以作为库引用到其他Go程序中。infer.go是Go语言社区中流行的一种代码推理工具，它被广泛应用于代码分析、重构和自动化测试等领域。




---

### Structs:

### tpWalker





### cycleFinder





## Functions:

### infer





### renameTParams





### typeParamsString





### isParameterized





### isParameterized





### varList





### coreTerm





### killCycles





### typ





### varList





### tparamIndex





