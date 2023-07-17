# File: scanner.go

scanner.go 是 Go 语言中 scanner 包的核心文件之一，其作用是提供一个用于读取文本数据流的工具，以将其转换为基本数据类型或标记（token）序列。它被广泛用于代码解析、文本分析、文件读取等场景中，可大大简化处理文本数据时的代码实现。

在具体实现中，scanner.go 主要定义了 Scanner 结构体和相关的方法，其中包括：

- Init 方法：用于初始化 Scanner 实例，其中可以设置 Scanner 的输入源、分隔符等参数。
- Scan 方法：用于从输入源中读取下一个 token，并返回 token 的类型和值。可以通过循环调用 Scan 方法，逐一处理输入源中的所有 token。
- Buffer 方法：用于在输入源的当前位置缓存一定长度的文本数据，以提高后续 Scan 操作的效率。
- Text 方法：用于获取当前 token 的文本表示。
- Split 方法：用于设置默认的分隔符，并返回一个函数用于自定义分隔符。
- TokenString 方法：用于将 token 的类型转换为对应的字符串表示，方便输出和调试。

总之，scanner.go 是 Go 语言中实现基于文本数据流的处理工作的重要组成部分，为用户提供了高效、简便的文本读取和分析工具。




---

### Var:

### prefix








---

### Structs:

### ErrorHandler





### Scanner





### Mode





## Functions:

### next





### peek





### Init





### error





### errorf





### scanComment





### updateLineInfo





### trailingDigits





### isLetter





### isDigit





### scanIdentifier





### digitVal





### lower





### isDecimal





### isHex





### digits





### scanNumber





### litname





### invalidSep





### scanEscape





### scanRune





### scanString





### stripCR





### scanRawString





### skipWhitespace





### switch2





### switch3





### switch4





### Scan





