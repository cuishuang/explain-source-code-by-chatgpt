# File: union.go

union.go文件是Go语言标准库中的一个文件，主要实现了Go语言中的联合（union）数据类型，包括以下三个函数：

1. unsafe.Sizeof：用于获取联合体的大小，即占用的字节数。

2. unsafe.Offsetof：用于获取联合体中某个成员的偏移量（即相对于联合体起始地址的偏移量）。

3. unsafe.Alignof：用于获取联合体的对齐方式，即按几个字节对齐。

联合（union）是一种特殊的数据类型，与结构体类似，不同的是联合体中的成员共用同一段内存空间，只有其中一个成员有值。这种数据类型的特点是节省内存空间，但使用起来需要注意取值和赋值的正确性。

Go语言中提供了unsafe包来实现底层的内存操作，其中的Sizeof、Offsetof和Alignof函数可以用于联合体的操作。同时，由于unsafe包操作的是底层内存，因此使用时应特别注意安全性和正确性。

总之，union.go文件是Go语言标准库中实现联合体操作的重要文件，也展示了Go语言支持底层内存操作的特点。




---

### Structs:

### Union





### Term





## Functions:

### NewUnion





### Len





### Term





### Underlying





### String





### NewTerm





### Tilde





### Type





### String





### parseUnion





### parseTilde





### overlappingTerm





### flattenUnion





