# File: text/unicode/bidi/bracket.go

在Go的text项目中，text/unicode/bidi/bracket.go这个文件的作用是处理Unicode文本中的括号匹配和方向控制。

1. bracketType：表示括号的类型，包括开括号、闭括号和中立括号。
2. bracketPair：表示一个括号对，包含了括号的起始位置、结束位置和括号类型。
3. bracketPairs：表示多个括号对的集合，提供了操作括号对的方法。
4. bracketPairer：括号对提供器，用于在文本中定位出现的括号对。

下面是这些结构体的具体功能：

- String() string：将bracketPair转换为字符串形式。
- Len() int：返回bracketPairer中括号对的数量。
- Swap(i, j int)：交换bracketPairer集合中的两个括号对的位置。
- Less(i, j int) bool：比较给定位置的两个括号对是否满足less关系。

下面是这些函数的具体功能：

- resolvePairedBrackets(bp *bracketPairs)：处理括号对之间的嵌套关系，将嵌套的括号对解析出来。
- matchOpener(opener rune) rune：给定一个开括号，返回对应的闭括号。如果不存在对应的闭括号，则返回开括号本身。
- locateBrackets(s string, ds *unicode.RangeTable) *bracketPairs：在文本中定位出现的括号对。ds参数是用于处理特定类别括号的字符范围。
- getStrongTypeN0(c rune) bool：判断给定字符是否为双向文本中的Strong类别字符。
- classBeforePair(s string, idx int) int：返回括号对之前非空字符的字符类别。
- assignBracketType(s string, bp *bracketPairs)：为括号对设置括号类型，即开括号、闭括号或中立括号。
- setBracketsToType(s string, ds1, ds2 *unicode.RangeTable, bp *bracketPairs)：为括号对设置括号类型，ds1和ds2用于处理特定类别括号的字符范围。
- resolveBrackets(s string) *bracketPairs：解析文本中的括号，并返回括号对集合。

