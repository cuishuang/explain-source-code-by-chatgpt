# File: text/currency/format.go

在Go的text项目中，text/currency/format.go文件用于实现货币格式化和解析的功能。

该文件中包含了一些常量变量和结构体，以及一些函数。

变量解释：
- dummy: 用于标记占位符的dummy索引。
- defaultFormat: 表示默认的货币格式。
- NarrowSymbol, Symbol, ISO: 用于存储不同货币格式的标识符。
- optISO, optSymbol, optNarrow: 用于存储可选的格式标识符。

结构体解释：
- Amount：表示一个金额的结构体，包括整数和小数部分。
- formattedValue: 用于存储机写格式的货币值。
- Formatter: 代表一种格式化货币的机写式的货币格式。
- options: 表示格式化货币的选项。
- symbolIndex: 存储不同货币符号的索引。

函数解释：
- Currency：根据给定的符号或3个字母的ISO代码返回货币。同时会进行一些标准化处理，如大小写转换等。
- Format：根据给定的货币、金额和选项，返回格式化后的机写式货币值。
- adjust：对金额进行四舍五入和溢出处理。
- Default：根据给定的货币返回其默认的机写式货币符号。
- Kind：返回货币的格式类型，如ISO、Symbol等。
- format：根据给定的符号返回其对应的格式。
- formISO, formSymbol, formNarrow：根据给定的符号返回对应的格式标识符。
- lookupISO, lookupSymbol, lookupNarrow：根据给定的格式标识符返回对应的符号。
- lookup：根据给定的符号返回对应的格式标识符。

总体而言，text/currency/format.go文件中的变量、结构体和函数实现了货币的格式化和解析功能，主要包括货币格式、金额格式化的机写式和货币符号的处理等。这些功能在国际化和本地化场景下非常有用，可以帮助开发者处理货币的展示和解析。

