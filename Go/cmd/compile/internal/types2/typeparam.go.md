# File: typeparam.go

参数化类型（generics）是一种通用编程构造，它允许形式上相似但类型不同的代码被重用。 Go 语言一直以来都没有支持参数化类型，但是在 go1.18 版本中，引入了一个新的实验性特性: type parameter。

文件 go/src/cmd/typeparam/typeparam.go 就是 type parameter 特性的实现。它为 Go 语言引入了参数化类型的能力。在使用 type parameter 前，需要在文件开头加上一个 //go:build go1.18 的编译指令。

在这个文件中，主要实现了以下几个函数：

- typeparams.Check：对类型参数进行语义检查，在函数或方法的参数列表以及泛型声明部分中出现定义的参数名，检查是否定义过或者重复，是否可以比较，是否是顶级命名，类型参数和结果类型中是否兼容，等一系列语义检查工作。
- typeparams.Instantiate： 对模板中的类型参数进行实例化，将函数和方法声明中引用的类型参数替换为实际的类型。此外，还需要进行类型checking，以保证类型的一致性和正确性。
- typeparams.ResolveStub：将模板中的 stub 类型替换为具体的实现，并生成新的 AST。因为类型参数化的代码是在编译期实现的，所以编译器必须了解所有的类型，才能正确地进行代码生成。因此，在这个阶段，需要对 stub 进行解析和替换。

这些功能的实现让 Go 语言的类型更加灵活和泛化，避免了大量的重复代码。不过，在使用过程中，需要注意参数类型的一致性和约束条件的准确性，以避免出现错误。




---

### Var:

### lastID








---

### Structs:

### TypeParam





## Functions:

### nextID





### NewTypeParam





### newTypeParam





### Obj





### Index





### Constraint





### SetConstraint





### Underlying





### String





### cleanup





### iface





### is





### underIs





