# File: excelize/lib.go

在Go生态excelize项目中，excelize/lib.go文件是excelize库的主要实现文件之一。它包含了一些关键的函数和结构体，用于处理Excel文件的读取和写入。

在该文件中，bstrExp和bstrEscapeExp是用于匹配和转义XML字符串中的特殊字符的正则表达式。这些正则表达式用于确保在将XML字符串写入Excel文件时，特殊字符不会破坏文件结构或格式。

Stack是excelize库中使用的栈结构，用于在处理过程中存储和管理数据。具体来说，Push函数用于将元素推入栈中，Pop函数用于弹出栈顶的元素，Peek函数用于获取栈顶的元素，Len函数用于获取栈中元素的数量，Empty函数用于检查栈是否为空。

其他一些函数的功能如下：
- ReadZipReader用于读取ZIP文件中的内容，并返回读取结果。
- unzipToTemp用于将ZIP文件解压到临时文件夹。
- readXML用于读取XML文件并返回其内容。
- readBytes用于从文件中读取字节。
- readTemp用于读取临时文件。
- saveFileList用于保存文件列表。
- readFile用于读取文件。
- SplitCellName用于将单元格名称拆分为坐标。
- JoinCellName用于将坐标拼接为单元格名称。
- ColumnNameToNumber用于将列名称转换为数字。
- ColumnNumberToName用于将列数字转换为名称。
- CellNameToCoordinates用于将单元格名称转换为坐标。
- CoordinatesToCellName用于将坐标转换为单元格名称。
- rangeRefToCoordinates用于将范围引用转换为坐标。
- cellRefsToCoordinates用于将单元格引用列表转换为坐标列表。
- sortCoordinates用于对坐标列表进行排序。
- coordinatesToRangeRef用于将坐标转换为范围引用。
- getDefinedNameRefTo用于获取定义名称的引用范围。
- flatSqref用于将sqref字符串分割为坐标列表。
- inCoordinates用于检查坐标是否在坐标列表中。
- inStrSlice用于检查字符串是否在字符串切片中。
- inFloat64Slice用于检查Float64是否在Float64切片中。
- boolPtr, intPtr, uintPtr, float64Ptr, stringPtr用于将常用类型转换为指针类型。
- Value是表示Excel单元格值的结构体。
- MarshalXML和UnmarshalXML用于XML序列化和反序列化。
- namespaceStrictToTransitional用于将严格模式的命名空间转换为过渡模式。
- bytesReplace用于替换字节中的内容。
- genSheetPasswd用于生成工作表密码。
- getRootElement用于获取XML文件的根元素。
- genXMLNamespace用于生成XML命名空间。
- replaceNameSpaceBytes用于替换字节中的命名空间。
- addNameSpaces用于添加命名空间。
- setIgnorableNameSpace用于设置可忽略的命名空间。
- addSheetNameSpace用于添加工作表的命名空间。
- isNumeric用于检查字符串是否为数字。
- bstrUnmarshal用于将二进制字符串反序列化为uint16列表。
- bstrMarshal用于将uint16列表序列化为二进制字符串。
- newRat用于创建一个新的有理数。
- continuedFraction用于计算连分数。
- NewStack用于创建新的栈。
- Push用于将元素推入栈中。
- Pop用于弹出栈顶的元素。
- Peek用于获取栈顶的元素。
- Len用于获取栈中元素的数量。
- Empty用于检查栈是否为空。

以上这些函数和变量提供了excelize库在处理Excel文件时所需的一些基本功能和数据结构支持。

