# File: order.go

order.go是Go语言编译器(cmd/compile)中的一个文件，其主要作用是实现Go语言中的代码链接和初始化顺序，在编译过程中为每个包生成的对象文件中添加初始化和链接信息，确保依赖关系正确地被解析。

具体来说，order.go实现了以下几个功能：

1. 对函数和变量进行排序。根据符号引用关系，将每个包中的函数和变量按照依赖关系进行排序。排序算法使用了拓扑排序，将依赖性从高到低进行排序。

2. 为函数和变量生成初始化代码。对于有需要进行初始化的变量或函数，order.go会为其生成一段代码片段，在程序运行时确保其被正确初始化。

3. 为包生成链接信息。在Go语言中，每个包可能会依赖其他包中的函数或变量，因此order.go会为每个包生成一个init函数，并将这个函数的地址和其他需要链接的包的地址记录在init_array中，确保在程序启动时，所有需要链接的包都已正确链接。

总的来说，order.go文件起到连接不同包中代码的作用，保证包的依赖关系和初始化顺序正确，确保程序可以正确运行。




---

### Structs:

### orderState





### ordermarker





## Functions:

### order





### append





### newTemp





### copyExpr





### copyExprClear





### copyExpr1





### cheapExpr





### safeExpr





### addrTemp





### mapKeyTemp





### mapKeyReplaceStrConv





### markTemp





### popTemp





### stmtList





### orderMakeSliceCopy





### edge





### orderBlock





### exprInPlace





### orderStmtInPlace





### init





### call





### mapAssign





### safeMapRHS





### stmt





### hasDefaultCase





### exprList





### exprListInPlace





### exprNoLHS





### expr





### expr1





### as2func





### as2ok





### isFuncPCIntrinsic





### isIfaceOfFunc





