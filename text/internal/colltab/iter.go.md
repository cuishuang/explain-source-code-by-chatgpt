# File: text/internal/colltab/iter.go

在Go的text项目中，text/internal/colltab/iter.go文件的作用是提供了Unicode规范化和字典排序的功能。

该文件中定义了一些结构体和函数，下面对它们进行介绍：

1. 结构体Iter：表示一个文本迭代器，用于迭代处理文本。
   - Index 字段：表示当前处理的文本索引位置。
   - Norm 字段：表示是否要进行Unicode规范化。
   - NormIter 字段：表示Unicode规范化的迭代器。
   - Text 字段：表示要处理的文本。
   - Next 字段：表示要处理的下一个字符。
   - NormNext 字段：表示进行规范化后的下一个字符。

2. 结构体ForwardIter：Iter的子结构体，用于向前处理文本。

3. 结构体BackIter：Iter的子结构体，用于向后处理文本。

以下是函数列表及其作用描述：

- Reset：将迭代器重置为初始状态，以便重新进行迭代。
- Len：返回要处理的文本的长度。
- Discard：将当前迭代器指向下一个字符，跳过当前字符。
- End：判断当前迭代器是否已经处理完所有文本。
- SetInput：设置迭代器要处理的文本。
- SetInputString：设置迭代器要处理的文本（使用字符串形式）。
- done：判断当前字符是否已经处理完，即是否到达文本末尾。
- appendNext：将当前字符追加到指定的字节切片中。
- Next：获取并返回下一个字符，同时将迭代器指向下一个字符。
- nextNoNorm：获取下一个字符，但不进行规范化处理。
- doNorm：进行Unicode规范化处理，将NormNext设置为规范化后的字符。

这些函数及结构体的目的是为了支持文本排序和Unicode规范化。它们提供了对文本的迭代和处理的功能，使得开发者能够在文本处理中更加灵活和高效。

