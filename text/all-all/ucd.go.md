# File: text/internal/ucd/ucd.go

文件"ucd.go"的作用是实现了对Unicode字符数据库（UCD）的解析和处理。

变量解释：
- KeepRanges：用于存储表示要保留的Unicode字符范围。
- errIncorrectLegacyRange：表示无效的Unicode字符范围错误。
- reRange：正则表达式模式，用于匹配和提取Unicode字符范围。
- bools：用于存储布尔值的切片。
- errUndefinedEnum：表示未定义的枚举错误。

结构体解释：
- Option：用于配置解析器的选项。
- Parser：解析器结构体，用于解析Unicode字符范围和属性文件。
 
函数解释：
- Parse：接受UCD解析器和一个字符串作为输入，解析字符串中的Unicode字符范围。
- keepRanges：将解析的Unicode字符范围添加到给定的切片中。
- Part：辅助函数，用于从字符串中提取指定位置的部分。
- CommentHandler：用于处理注释的回调函数。
- setError：用于设置解析器错误。
- getField：辅助函数，用于从字符串中提取指定字段。
- Err：返回解析器的错误。
- New：创建一个新的UCD解析器。
- Next：将解析器的位置移到下一个目标（字符范围或属性）。
- parseRune：解析给定字符串中的一个Unicode字符。
- Rune：用于返回当前解析位置的Unicode字符。
- Runes：用于返回字符范围的切片。
- Range：返回当前解析位置的字符范围。
- getRange：用于解析给定字符串中的字符范围。
- Bool：用于解析给定字符串中的布尔值。
- Int：用于解析给定字符串中的整数。
- Uint：用于解析给定字符串中的无符号整数。
- Float：用于解析给定字符串中的浮点数。
- String：用于解析给定字符串中的字符串。
- Strings：用于解析给定字符串中的字符串切片。
- Comment：解析注释并返回相应的值。
- Enum：用于解析指定枚举类型的字符串。

