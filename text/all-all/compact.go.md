# File: text/internal/language/compact.go

text/internal/language/compact.go文件是Go的text项目中的一个内部语言包文件，它提供了一种压缩的方式来存储和检索语言数据，以实现更高效的文本处理。

CompactCoreInfo结构体是压缩语言数据的核心信息，它包含了语言数据的元信息，如语言ID、版本、最小代码点和最大代码点。这些信息用来描述语言数据的范围和特征。

GetCompactCore函数用于获取给定语言的压缩核心数据信息。它接收一个语言ID作为输入参数，并返回一个CompactCoreInfo结构体，其中包含了该语言的压缩核心信息。

Tag函数用于获取给定代码点的语言数据标签。它接收一个代码点和语言ID作为输入参数，并返回一个表示该代码点在给定语言中的标签。这个标签可以表达出字符的语言特征，如大小写、字母、数字等。

总之，text/internal/language/compact.go文件中的CompactCoreInfo结构体和相关函数主要用于压缩语言数据的存储和检索，通过有效地存储语言信息和字符标签，实现了高效的文本处理。

