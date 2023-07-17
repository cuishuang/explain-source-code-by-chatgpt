# File: operand.go

operand.go文件定义了表达式的操作数结构，包括常量、变量、函数和表达式等。这个文件还提供了一些有用的帮助函数，用于解析和校验操作数。

具体来说，文件中定义了以下几个struct：

- Operand：表示表达式中的一个操作数，可能是一个常量、一个变量、一个函数或者一个表达式。该struct中定义了一个kind字段来表示操作数的类型，以及value字段用来存储操作数的值。
- OperandMode：表示表达式的操作数模式。可能的值包括值、寄存器、内存和分支等。
- OperandStackView：表示操作数栈的视图，即一个stack of OperandMode。该struct提供了一些有用的帮助函数，例如指定操作数模式以及从操作数栈中弹出操作数等。
- Obj：表示一个go程序的object。每个Obj都有一个类型、一个标识符以及一些旁边的信息。其实在go中，一切都是对象，也就是说每个变量、每个函数都是一个对象。




---

### Var:

### operandModeString





### kind2tok








---

### Structs:

### operandMode





### operand





## Functions:

### Pos





### operandString





### String





### setConst





### isNil





### assignableTo





