# File: text/unicode/norm/forminfo.go

在Go的text项目中，text/unicode/norm/forminfo.go文件的作用是定义了Unicode标准的规范化信息，包括字符的规范化形式、规范化的细节和属性。

`formTable`是一个映射表，其键是Unicode字符的值，值是该字符的规范化信息。它根据Unicode代码点提供了规范化过程中所需的信息。

`Properties`结构体定义了一个Unicode字符的属性，例如大小写信息、类别、分解形式等。

`lookupFunc`是一个函数签名，用于查找特定Unicode字符的规范化信息。该函数将一个Unicode代码点作为输入，并返回该代码点的规范化信息。

`formInfo`结构体定义了一个规范化形式的信息，包括规范化类型、规范化映射函数、规范化名称等。

`qcInfo`结构体定义了一个字符序列的快速查找表，用于将多个字符的规范化映射结果组合成一个规范化字符串。

`BoundaryBefore`和`BoundaryAfter`用于表示字符的规范化边界。

`isYesC`和`isYesD`用于判断字符是否为规范化的Combining字符。

`combinesForward`和`combinesBackward`分别用于在规范化过程中判断字符是否可以与之前或之后的字符组合。

`hasDecomposition`用于判断字符是否有分解形式。

`isInert`用于判断字符是否是规范化不变的。

`multiSegment`用于判断字符是否需要多个段进行规范化。

`nLeadingNonStarters`和`nTrailingNonStarters`分别用于确定字符序列中起始和结束的非初始字符数量。

`Decomposition`是表示Unicode字符分解形式的切片。

`Size`是字符规范化结果的长度。

`CCC`、`LeadCCC`和`TrailCCC`分别表示字符的组合序列。

`buildRecompMap`函数用于构建字符重新组合的映射表。

`combine`函数用于将组合字符合并到前一个字符上。

`lookupInfoNFC`和`lookupInfoNFKC`函数用于查找NFC和NFKC规范化形式的信息。

`Properties`函数用于获取字符的属性。

`PropertiesString`函数用于将字符的属性转化为字符串。

`compInfo`函数用于在规范化过程中处理字符的组合情况。

这些函数和数据结构共同构成了Go的text项目中Unicode规范化功能的核心。它们负责处理字符的规范化形式、属性和组合关系，以确保文本的一致性和正确性。

