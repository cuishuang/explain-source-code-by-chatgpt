# File: tools/cmd/guru/implements.go

在Golang的Tools项目中，`tools/cmd/guru/implements.go`文件的作用是实现了`guru-implements`命令的主要逻辑。

`implementsResult`结构体表示了实现关系的结果，包含了实现关系的源类型和目标类型之间的关联信息。

`typesByString`结构体是一个用于缓存类型的映射表，它以类型的字符串形式作为键，类型本身作为值。

- `implements`函数用于检查给定类型的实现关系，并返回实现关系的结果。
- `PrintPlain`函数用于以纯文本的方式打印实现关系的结果。
- `JSON`函数用于以JSON格式打印实现关系的结果。
- `makeImplementsTypes`函数用于创建`implementsResult`结构体的实例。
- `makeImplementsType`函数用于创建`implementsResult`结构体的实例，其中目标类型为接口类型。
- `typeKind`函数用于获取给定类型的类型信息。
- `isInterface`函数用于判断给定类型是否是接口类型。
- `Len`、`Less`、`Swap`函数分别是用于排序实现关系结果的切片的方法，实现了`sort.Interface`接口。

这些函数共同实现了对于给定类型的实现关系的检查和打印功能。其中，`implements`函数是核心函数，其余函数则辅助实现了结果的打印和处理。

