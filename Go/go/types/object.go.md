# File: object.go

object.go是Go编程语言的标准库中的一个文件，其作用是定义了代表所谓Go程序中的“对象”的结构体。

这个文件定义了三种对象类型的结构体：

1. Type：表示Go类型的对象，包括基本数据类型、结构体、函数和接口等，以及用户定义的类型。Type对象中包含类型的名称、基础类型、方法集等信息。

2. Value：表示Go程序中的值，包括各种类型的常量、变量等。Value对象中包含值的类型和具体的值内容，可以通过类型断言等方式将其转换为其他类型的值。

3. Package：表示Go程序中的包，包括标准库和用户定义的包。Package对象中包含包的名称、路径、依赖等信息。

除了定义对象类型的结构体外，该文件还定义了一些有关对象的操作函数，包括创建类型、值、包等对象，获取类型、值、包等对象的信息，以及比较、转换等操作。

总之，object.go文件是Go编程语言标准库中的一个重要文件，它定义了Go程序中的对象，并提供了许多与对象相关的操作函数，为Go程序的开发和实现提供了重要的基础支持。




---

### Structs:

### Object





### object





### color





### PkgName





### Const





### TypeName





### Var





### Func





### Label





### Builtin





### Nil





## Functions:

### isExported





### Id





### String





### colorFor





### Parent





### Pos





### Pkg





### Name





### Type





### Exported





### Id





### String





### order





### color





### scopePos





### setParent





### setType





### setOrder





### setColor





### setScopePos





### sameId





### less





### NewPkgName





### Imported





### NewConst





### Val





### isDependency





### NewTypeName





### _NewTypeNameLazy





### IsAlias





### NewVar





### NewParam





### NewField





### Anonymous





### Embedded





### IsField





### Origin





### isDependency





### NewFunc





### FullName





### Scope





### Origin





### hasPtrRecv





### isDependency





### NewLabel





### newBuiltin





### writeObject





### packagePrefix





### ObjectString





### String





### String





### String





### String





### String





### String





### String





### String





### writeFuncName





