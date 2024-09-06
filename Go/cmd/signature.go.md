# File: signature.go

signature.go是go语言中的一个文件，其作用是提供函数签名的类型。函数签名由函数的参数及返回值类型组成。

在go语言中，函数也是一种类型，可以将函数作为参数传递给另一个函数，或者将函数作为返回值返回给调用方。在这些场景下，函数的签名起着很重要的作用。

signature.go文件定义了三种类型：Type, FuncType, and Signature。其中，Type类型代表任意数据类型，FuncType类型代表函数类型，Signature代表函数签名类型。

Signature类型定义了函数的参数类型和返回值类型，以及函数是否是变参函数。在go语言中，变参函数的参数个数可以是任意个。

下面是Signature类型的定义：

```go
type Signature struct {
    Rparams     *Tuple // 返回参数列表
    Rresults    *Tuple // 返回值列表
    Variadic    bool   // 是否是变参函数
}
```

Signature类型作为函数类型的一部分来使用。例如：

```go
type FuncType struct {
    ...
    Params *Tuple
    Results *Tuple
}
```

可以看到，FuncType类型包含了Signature类型。这是因为函数类型必须知道函数的参数及返回值类型，Signature类型提供了这些信息。

总之，signature.go文件定义了函数签名的类型，为函数的类型转换、参数传递、返回值等操作提供了基础支持。




---

### Structs:

### Signature





## Functions:

### NewSignatureType





### Recv





### TypeParams





### SetTypeParams





### RecvTypeParams





### Params





### Results





### Variadic





### Underlying





### String





### funcType





### collectParams





