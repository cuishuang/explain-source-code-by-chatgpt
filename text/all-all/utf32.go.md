# File: text/encoding/unicode/utf32/utf32.go

在Go的text项目中，`text/encoding/unicode/utf32/utf32.go` 文件的作用是实现对UTF-32编码的处理。

- `All` 是一个`byte`类型的切片，包含了所有UTF-32编码的字符。
- `ErrMissingBOM` 是一个错误，表示缺少字节顺序标记（BOM）。
- `mibValue` 是一个`int`类型的常量，表示UTF-32编码的 MIB 值（在编码标准中标识编码的唯一数字）。
- `BOMPolicy` 是一个表示字节顺序标记的策略的类型。
- `Endianness` 是一个表示字节序（大小端）的类型。
- `config` 是一个配置结构体，用于设置编码器和解码器的字节顺序和错误策略。
- `utf32Encoding` 是一个实现了`encoding.Encoding`接口的结构体，用于UTF-32编码的编码器。
- `utf32Decoder` 是一个实现了`encoding.Decoder`接口的结构体，用于将UTF-32编码的字节流解码为Unicode字符。
- `utf32Encoder` 是一个实现了`encoding.Encoder`接口的结构体，用于将Unicode字符编码为UTF-32字节流。
- `UTF32` 是一个表示UTF-32编码的类型。
- `NewDecoder` 是一个创建解码器的函数，用于将UTF-32编码的字节流解码为Unicode字符。
- `NewEncoder` 是一个创建编码器的函数，用于将Unicode字符编码为UTF-32字节流。
- `ID` 是一个返回编码的MIB值的方法。
- `String` 是一个返回编码的名称的方法。
- `Reset` 是一个重置解码器的方法。
- `Transform` 是一个将给定的字节切片转换为Unicode字符的方法。

