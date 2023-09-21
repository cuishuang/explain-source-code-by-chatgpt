# File: tools/gopls/internal/lsp/glob/glob.go

文件`glob.go`的作用是提供了用于处理通配符匹配的函数和结构体。

- `errBadRange`和`errInvalidUTF8`是用于表示错误的变量。`errBadRange`表示在范围表达式中给定了无效的范围，而`errInvalidUTF8`表示在解析中出现了无效的UTF-8序列。
  
- `Glob`结构体表示一个通配符表达式，包含了要匹配的模式和一些其他的信息。
  
- `element`结构体表示一个元素，可以是一个通配符字符、一个范围或一个普通字符。

- `slash`结构体表示了斜杠字符的信息，它有两个字段表示是否匹配零个或多个斜杠。
  
- `Parse`函数用于解析通配符模式，将其转换为`Glob`结构体。

- `parse`函数用于解析通配符模式的剩余部分并返回一个元素。

- `readRangeRune`函数用于解析范围中的一个字符并返回正确的Unicode字符。

- `parseLiteral`函数用于解析一个字面量字符。

- `String`函数用于将一个元素或范围表达式转换为相应的字符串。

- `Match`函数用于判断给定字符串是否与通配符模式匹配。

- `match`函数用于实际执行匹配逻辑的内部函数。

- `split`函数用于将通配符模式按斜杠拆分为多个子模式。

