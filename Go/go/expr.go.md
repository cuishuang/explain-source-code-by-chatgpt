# File: expr.go

expr.go文件的作用是解析和评估Go语言的表达式。

Go语言是一种静态类型语言，表达式的类型和值在编译时就确定下来了。表达式可以用于进行计算、比较、逻辑运算等，是编写Go程序必不可少的语言元素。

expr.go文件包含了处理表达式相关的代码，包括词法分析、语法分析、类型检查和求值。在表达式求值中，expr.go文件会遵循Go语言的类型转换规则和运算优先级，确保表达式求值的正确性。

除了表达式解析和求值，expr.go文件还包含了一些常用的内置函数，如len、cap、make、append等。这些内置函数在表达式求值中经常会被使用到。

总之，expr.go文件是Go语言编译器和运行时环境中非常重要的一部分，它使得Go语言能够灵活地处理各种表达式，让程序员能够更加方便地编写高质量的Go程序。




---

### Var:

### errNotConstraint





### errComplex








---

### Structs:

### Expr





### TagExpr





### NotExpr





### AndExpr





### OrExpr





### SyntaxError





### exprParser





## Functions:

### isExpr





### Eval





### String





### tag





### isExpr





### Eval





### String





### not





### isExpr





### Eval





### String





### andArg





### and





### isExpr





### Eval





### String





### orArg





### or





### Error





### Parse





### IsGoBuild





### splitGoBuild





### parseExpr





### or





### and





### not





### atom





### lex





### IsPlusBuild





### splitPlusBuild





### parsePlusBuildExpr





### isValidTag





### PlusBuildLines





### pushNot





### appendSplitAnd





### appendSplitOr





