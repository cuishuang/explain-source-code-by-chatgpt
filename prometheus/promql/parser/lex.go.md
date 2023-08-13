# File: promql/parser/lex.go

在Prometheus项目中，promql/parser/lex.go文件的作用是实现PromQL查询语言的词法分析器。该文件中定义了用于将查询字符串分解为一系列token的函数和数据结构。

以下是这些变量和结构体的作用：

- key：表示查询字符串中的关键字或标识符。
- ItemTypeStr：表示token的类型，如关键字、标识符、运算符等。
- Item：表示从查询字符串中提取的token。
- ItemType：表示token的类型。
- stateFn：表示状态函数，用于处理当前字符并决定下一步的操作。
- Pos：表示token在查询字符串中的位置。
- Lexer：表示词法分析器。

以下是这些函数的作用：

- String：将token的类型转换为可读的字符串表示。
- Pretty：将token的类型和值以易读的方式格式化。
- IsOperator：检查token是否为运算符。
- IsAggregator：检查token是否为聚合函数。
- IsAggregatorWithParam：检查token是否为带参数的聚合函数。
- IsKeyword：检查token是否为关键字。
- IsComparisonOperator：检查token是否为比较运算符。
- IsSetOperator：检查token是否为集合运算符。
- init：初始化词法分析器。
- desc：返回当前字符的描述。
- next：获取下一个字符并返回其类型。
- peek：返回下一个字符而不消耗它。
- backup：回退到上一个字符。
- emit：将当前token添加到token列表中。
- ignore：忽略当前字符。
- accept：接受指定类型的字符。
- acceptRun：连续接受指定类型的字符。
- errorf：生成词法分析错误。
- NextItem：获取下一个token。
- Lex：执行词法分析。
- lexStatements：词法分析语句。
- lexInsideBraces：词法分析大括号内的内容。
- lexValueSequence：词法分析值序列。
- lexEscape：词法分析转义字符。
- digitVal：获取数字字符的值。
- skipSpaces：跳过空白字符。
- lexString：词法分析字符串字面量。
- lexRawString：词法分析原始字符串。
- lexSpace：词法分析空白字符。
- lexLineComment：词法分析行注释。
- lexDuration：词法分析时间段。
- lexNumber：词法分析数字。
- lexNumberOrDuration：词法分析数字或时间段。
- acceptRemainingDuration：接受剩余的时间段。
- scanNumber：扫描数字。
- lexIdentifier：词法分析标识符。
- lexKeywordOrIdentifier：词法分析关键字或标识符。
- isSpace：检查字符是否为空白字符。
- isEndOfLine：检查字符是否为行尾。
- isAlphaNumeric：检查字符是否为字母或数字。
- isDigit：检查字符是否为数字。
- isAlpha：检查字符是否为字母。
- isLabel：检查字符是否为标签。

