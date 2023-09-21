# File: tools/gopls/internal/lsp/protocol/tsjson.go

在Golang的Tools项目中，`tsjson.go`文件的作用是提供与TypeScript JSON格式之间的转换功能。

`UnmarshalError`结构体用于表示解析JSON时可能发生的错误。它包含了错误的消息、位置以及相关联的错误。

`Error`函数用于创建一个新的UnmarshalError实例，并返回相应的错误消息。它接受一条错误消息、位置以及相关联的错误。

`MarshalJSON`是一个函数签名为`func(v interface{}) ([]byte, error)`的方法，它用于将给定的Go对象转换为对应的JSON数据。在`tsjson.go`中，它被用来序列化TypeScript JSON对象。

`UnmarshalJSON`是一个函数签名为`func(data []byte, v interface{}) error`的方法，它用于将JSON数据解析为Go对象的值。在`tsjson.go`中，它被用来反序列化TypeScript JSON对象。

总之，`tsjson.go`文件中的结构体和函数提供了用于处理TypeScript JSON格式的工具，包括将Go对象转换为JSON数据和将JSON数据解析为Go对象的功能。

