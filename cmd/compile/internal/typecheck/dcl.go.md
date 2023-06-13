# File: dcl.go

dcl.go是Go编译器中的一个文件，全称为declaration。它的作用是解析Go程序中的声明语句，包括变量、函数、常量、结构体等的声明，并将其存储到符号表中供后续的语法分析和代码生成使用。

具体来说，dcl.go会预处理Go程序中的声明语句，包括变量和函数的声明和定义。它会对每个声明进行语法分析，确定其类型、作用域等信息，并生成相应的符号表项。

在处理变量声明时，dcl.go会检查变量的生命周期和作用域，并解析变量初始化表达式。它还支持多个变量同时声明的语法，如var a, b, c int。

在处理函数声明时，dcl.go会检查函数的参数和返回值，并将函数定义的符号表项存储到符号表中。它还支持嵌套函数的声明，该语法允许将一个函数作为另一个函数的局部变量或返回值。

总之，dcl.go是Go编译器中非常重要的一个文件，它完成了语法分析的重要部分，解析程序中的各种声明，然后将其存储在符号表中。这些信息保证了程序在后面的处理阶段能够正确地被编译和执行。




---

### Var:

### DeclContext





### funcStack





### autotmpnamesmu





### autotmpnames








---

### Structs:

### funcStackEnt





## Functions:

### DeclFunc





### Declare





### Export





### StartFuncBody





### FinishFuncBody





### CheckFuncStack





### autoexport





### checkdupfields





### checkembeddedtype





### declareParams





### declareParam





### Temp





### TempAt





### autotmpname





### NewMethodType





