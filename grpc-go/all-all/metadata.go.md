# File: grpc-go/metadata/metadata.go

在grpc-go项目中，`grpc-go/metadata/metadata.go`文件的作用是定义了处理gRPC元数据的函数和结构。

`MD`是一个`map[string][]string`类型的别名，用于存储gRPC元数据中的键值对信息。

`mdIncomingKey`和`mdOutgoingKey`是`context`包中的私有类型，用于在上下文中传递接收和发送的元数据。

`rawMD`是`[]metadata.MD`类型的别名，用于存储原始的元数据。

以下是`metadata.MD`结构体的方法：

- `DecodeKeyValue`：解析键值对字符串，并返回键和值。

- `New`：创建一个空的`MD`对象。

- `Pairs`：返回一个包含所有键值对的切片。

- `Len`：返回键值对的数量。

- `Copy`：创建并返回一个`MD`的副本。

- `Get`：根据给定的键返回与之相关联的所有值。

- `Set`：设置给定键的值。

- `Append`：将给定值添加到指定键的现有值列表中。

- `Delete`：删除给定键的值。

- `Join`：将给定元数据和当前元数据结合为一个新的`MD`对象。

以下是`metadata`包中的函数：

- `NewIncomingContext`：创建一个新的上下文对象，并将给定元数据设置为接收到的元数据。

- `NewOutgoingContext`：创建一个新的上下文对象，并将给定元数据设置为要发送的元数据。

- `AppendToOutgoingContext`：向现有的发送元数据中追加给定的键值对。

- `FromIncomingContext`：从给定的上下文中获取接收到的元数据。

- `ValueFromIncomingContext`：从给定的上下文中获取接收到的元数据，以键和值的形式返回。

- `copyOf`：创建并返回给定`MD`对象的副本。

- `FromOutgoingContextRaw`：从给定的`rawMD`中获取发送的元数据。

- `FromOutgoingContext`：从给定的上下文中获取发送的元数据。

这些函数用于设置、处理和提取gRPC元数据的功能。

