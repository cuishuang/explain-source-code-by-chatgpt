# File: typeset.go

go/src/cmd/typeset.go这个文件是Go语言编译器源码中的一个文件，负责实现标识符的类型推导和检查工作。该文件中定义了用于描述标识符类型的数据结构，以及完成类型检查的相关函数。

具体来说，该文件中定义的数据结构有：

1. Type：用于表示类型的结构体，包含类型名称、基础类型、元素类型等信息。

2. FuncType：用于表示函数类型的结构体，包含参数列表和返回值类型列表等信息。

3. Signature：用于表示函数签名的结构体，包含函数类型、参数和返回值列表等信息。

除了数据结构之外，该文件中还实现了一系列函数用于类型推导和检查。主要的函数包括：

1. ResolveType：用于将标识符的类型解析为Type结构体。

2. CheckExpr：对表达式进行类型推导和检查。

3. CheckStmt：对语句进行类型检查，包括赋值语句、返回语句、if语句、switch语句等。

4. CheckFunc：对函数进行类型检查，包括参数、返回值、函数体等。

总的来说，typeset.go文件是Go语言编译器中非常重要的一个文件，负责实现类型推导和检查，是Go语言强大类型系统的基础。




---

### Var:

### topTypeSet





### invalidTypeSet








---

### Structs:

### _TypeSet





### byUniqueMethodName





## Functions:

### IsEmpty





### IsAll





### IsMethodSet





### IsComparable





### NumMethods





### Method





### LookupMethod





### String





### hasTerms





### subsetOf





### is





### underIs





### computeInterfaceTypeSet





### intersectTermLists





### sortMethods





### assertSortedMethods





### Len





### Less





### Swap





### computeUnionTypeSet





