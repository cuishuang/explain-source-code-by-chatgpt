# File: tools/godoc/format.go

在Golang的Tools项目中，tools/godoc/format.go文件的作用是格式化Go源代码的文档注释和示例代码。

selRx是一个正则表达式，用于匹配注释和示例代码的开始和结束标记。

startTags和endTag是用来构建注释和示例代码的结构体的开始和结束标签。

Segment结构体表示一个文本段落，其中包含原始文本数据和段落所处的位置信息。

Selection结构体表示一个用户选择的文本片段，包括起始行和列以及结束行和列。

LinkWriter结构体实现了Write方法，用于在生成的HTML文档中添加链接。

SegmentWriter结构体用于将文本段落和选择结构体的内容写入HTML文档。

merger结构体用于将注释和示例代码段合并为一个段落。

isEmpty函数用于判断一个文本段落是否为空。

FormatSelections函数用于将格式化后的注释和示例代码段写入HTML文档。

newMerger函数用于创建一个新的合并器。

next函数用于查找下一个注释或示例代码段。

lineSelection函数用于创建一个选择结构体，表示单行注释或示例代码。

tokenSelection函数用于创建一个选择结构体，表示一个范围内的注释或示例代码。

makeSelection函数用于创建一个选择结构体，根据给定的行和列信息。

regexpSelection函数用于根据正则表达式创建一个选择结构体。

RangeSelection函数用于创建一个选择结构体，表示一段范围内的注释或示例代码。

selectionTag函数用于根据给定的选择结构体创建开始或结束标签。

FormatText函数用于格式化给定的文本，将空格和换行符转换为HTML实体。

