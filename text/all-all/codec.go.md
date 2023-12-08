# File: text/internal/catmsg/codec.go

text/internal/catmsg/codec.go 文件是 Go 的 text 项目中的一个文件，主要用来处理消息编码和解码的逻辑。下面详细介绍各个变量和结构体的作用，以及各个函数的功能：

- errUnknownHandler: 这是一个错误变量，表示未知的处理程序。在编码或解码消息时，如果找不到相应的处理程序，就会返回这个错误。

结构体：

- Renderer: 这个结构体用于渲染文本模板，并将模板渲染的结果写入 io.Writer 流中。

- Dictionary: 这是一个字典结构体，用于存储消息和对应的处理程序。可以通过它来注册和查找处理程序。

- Encoder: 这个结构体用于将消息编码为二进制格式，并将编码的结果写入 io.Writer 流中。

- keyVal: 这个结构体用于表示键值对。

- Decoder: 这个结构体用于将二进制格式的消息解码为文本格式，并返回解码后的消息。

函数：

- Language: 这个函数用于设置消息的语言。

- setError: 这个函数用于设置编码或解码过程中发生的错误。

- EncodeUint: 这个函数用于将无符号整数编码为二进制格式。

- EncodeString: 这个函数用于将字符串编码为二进制格式。

- EncodeMessageType: 这个函数用于将消息类型编码为二进制格式。

- EncodeMessage: 这个函数用于将消息编码为二进制格式，并将编码的结果写入 io.Writer 流中。

- checkInBody: 这个函数用于检查消息体中是否包含指定的变量。

- stripPrefix: 这个函数用于从消息体中删除指定的前缀。

- flushTo: 这个函数用于将消息中未编码的数据刷新到 io.Writer 流中。

- addVar: 这个函数用于向字典中添加变量并设置对应的处理程序。

- EncodeSubstitution: 这个函数用于将文本模板中的插值表达式编码为二进制格式。

- NewDecoder: 这个函数用于创建一个新的解码器。

- Done: 这个函数用于表示解码过程已完成。

- Render: 这个函数用于渲染文本模板，并将结果写入 io.Writer 流中。

- Arg: 这个函数用于获取指定位置的参数。

- DecodeUint: 这个函数用于将二进制格式的无符号整数解码为文本格式。

- DecodeString: 这个函数用于将二进制格式的字符串解码为文本格式。

- SkipMessage: 这个函数用于跳过指定消息的解码。

- Execute: 这个函数用于执行文本模板。

- execute: 这个函数用于执行消息处理程序。

- executeMessageFromData: 这个函数用于从二进制数据中执行消息处理程序。

- executeMessage: 这个函数用于执行消息处理程序。

- ExecuteMessage: 这个函数用于执行消息处理程序，并将结果写入 io.Writer 流中。

- ExecuteSubstitution: 这个函数用于执行插值表达式，并将结果写入 io.Writer 流中。

