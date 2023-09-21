# File: grpc-go/internal/pretty/pretty.go

grpc-go/internal/pretty/pretty.go文件是grpc-go项目中的一个内部包，主要用于将gRPC消息转换为易于阅读和理解的格式。

该文件中的ToJSON函数接收一个proto.Message类型的消息，并将其转换成JSON格式的字符串。它使用了标准库中的json.MarshalIndent方法，并在其基础上进行了一些定制化的处理，以便更好地展示gRPC消息的结构。

而FormatJSON函数则在ToJSON的基础上进行了一些格式化处理，以便更好地展示嵌套的gRPC消息。它会对JSON字符串进行重新格式化，并在每个缩进级别前添加一个"↓"符号，以标识其嵌套层次。

这两个函数在调试和测试gRPC代码时非常有用。通过将gRPC消息转换成易于阅读和理解的JSON格式，开发人员可以更方便地查看和分析消息的内容、字段和嵌套关系，以便进行问题诊断和代码调试。

总结起来，grpc-go/internal/pretty/pretty.go文件提供了将gRPC消息转换成易读的JSON格式的功能，方便开发人员进行调试和分析。

