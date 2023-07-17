# File: builtins.go

builtins.go这个文件是Go语言编译器中的内置函数实现，它定义了所有的内置函数及其实现方式。内置函数是Go语言运行时环境中预先定义的函数，不需要引入包就可以直接使用。这些函数包括类型转换、切片、map、make、new、len、cap、append、copy、delete等等。

在builtins.go中，每个内置函数都有一个名字和对应的实现代码。这些内置函数都被编译进了Go程序中，在程序运行时直接调用。

除了内置函数的实现，builtins.go还提供了一些辅助函数，用于Go语言编译器的内部实现。这些辅助函数包括了各种类型转换和类型检测函数，例如类型断言、类型比较以及指针的一些操作等等。

总的来说，builtins.go是Go语言编译器中非常重要的文件，它定义了Go语言运行时环境中预定义的内置函数，这些函数能够满足程序开发中的基本需求，在提高程序开发效率和程序性能方面发挥了重要作用。

## Functions:

### builtin





### hasVarSize





### applyTypeFunc





### makeSig





### arrayPtrDeref





### unparen





