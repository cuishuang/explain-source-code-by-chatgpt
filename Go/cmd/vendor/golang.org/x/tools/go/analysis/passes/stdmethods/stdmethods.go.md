# File: stdmethods.go

stdmethods.go是Go语言中标准库的一部分，它定义了一些基础类型的方法。

具体来说，stdmethods.go中定义了以下类型的方法：

- bool类型：它只有一个方法——Not()，用于将布尔值取反。

- int、int8、int16、int32、uint、uint8、uint16、uint32、uintptr和float32、float64、complex64、complex128类型：它们都实现了fmt包中相应类型的Stringer接口，该接口用于将值格式化为字符串。此外，还定义了一些比较大小的方法（如IntAbs()、Float32bits()等）。

- string类型：它实现了fmt包中的Stringer接口，该接口用于将字符串格式化为字符串。此外，还定义了一些字符串操作的方法（如ToUpper()、ToLower()等）。

- slice类型：它实现了sort包中的接口，用于排序。此外，还定义了一些操作slice的方法（如Copy()、Append()等）。

- map类型：它实现了runtime包中的hash方法，用于计算哈希值。此外，还定义了一些操作map的方法（如Delete()、Load()等）。

- error类型：它只有一个方法——Error()，用于返回错误信息的字符串表示形式。

stdmethods.go的作用是使这些基础类型的方法更加易于使用和扩展。用户在使用这些类型时，可以方便地调用相关的方法，而不用自己手动实现。同时，用户在自定义类型时，也可以参考这些基础类型的方法来定义自己的方法。

