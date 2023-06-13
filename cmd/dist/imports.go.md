# File: imports.go

在Go语言中，通过import关键字来引入其他包中的代码，但是在导入时需要遵循一定的规则和流程，不同的导入路径有不同的处理方式。在Go标准库中，有一个名为imports.go的文件，其作用是负责处理导入路径，解析和检查import语句，根据不同情况生成不同的导入信息。 

imports.go文件中包含了多个函数和结构体，其中最重要的是build和Import函数。build函数是整个导入过程的核心，其原型如下：

func build(ctxt *Context, path_ string, mode buildMode, stk *importStack) (*Package, error)

该函数接受4个参数，一个是Context类型的编译环境，一个是导入路径，一个是构建模式，还有一个是importStack类型的堆栈。build函数通过上述参数判断导入路径类型、检查该路径下的代码是否具有可导出性、递归导入所有依赖包，并在返回时检查是否出现了循环依赖等问题。最终，build函数返回一个Package类型的数据结构，其中包含导入的代码的所有信息（比如文件路径、代码位置等）。

而Import函数则是对build函数的一个封装，他只接受一个路径参数，并返回一个*Package类型的指针。下面是Import函数的原型：

func Import(path string) (*Package, error)

除了路径参数之外，Import函数内部还需要构建出一个默认的buildContext对象，来提供给build函数调用。在具体实现中，build和Import函数并不是分别实现了整个导入过程，而是通过一些辅助函数实现了各自的特定功能，并相互协作完成了整个导入流程。

总之，imports.go文件的作用是实现了Go语言中的导入规则和流程，以方便用户在自己的代码中引用其他包中的代码，同时还提供了一些辅助函数和数据结构，在导入过程中确保了各种异常情况的处理（如循环依赖等），使得我们可以愉快的使用import语句。




---

### Var:

### errSyntax





### errNUL








---

### Structs:

### importReader





## Functions:

### isIdent





### syntaxError





### readByte





### peekByte





### nextByte





### readKeyword





### readIdent





### readString





### readImport





### readComments





### readimports





### resolveVendor





