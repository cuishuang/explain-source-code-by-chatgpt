# File: text/unicode/norm/normalize.go

在Go的text项目中，`text/unicode/norm/normalize.go`文件的作用是实现Unicode字符串的标准化。Unicode字符有多种表示方式，标准化就是将这些不同表示方式的字符统一化，以便进行更方便的比较和处理。

以下是`normalize.go`文件中一些结构体和函数的作用详细介绍：

**结构体：**

- `Form`: 这个结构体定义了字符串标准化的形式。有四个可用的形式，分别是`FormNFC`、`FormNFD`、`FormNFKC`和`FormNFKD`。每个形式代表不同种类的标准化规范。

**函数：**

- `Bytes`: 将unicode字符序列标准化后以字节的形式返回。
- `String`: 将unicode字符序列标准化后以字符串的形式返回。
- `IsNormal`: 检查给定的unicode字符序列是否已经标准化。
- `cmpNormalBytes`: 比较两个标准化的unicode字符序列的字节表示是否相等。
- `IsNormalString`: 检查给定的字符串是否已经标准化。
- `patchTail`: 修复字符串的末尾，使其包含标准化的检查点。
- `appendQuick`: 在字节切片的末尾追加标准化的unicode字符。
- `Append`: 在字节切片的末尾追加标准化的unicode字符序列。
- `doAppend`: 在字节切片的末尾追加标准化的unicode字符序列。
- `doAppendInner`: 在字节切片的末尾追加标准化的unicode字符序列。
- `AppendString`: 在字节切片的末尾追加标准化的字符串。
- `QuickSpan`: 寻找字节切片中下一个unicode字符的位置。
- `Span`: 寻找字节切片中下一个unicode字符的位置。
- `SpanString`: 寻找字符串中下一个unicode字符的位置。
- `quickSpan`: 寻找字节切片中下一个unicode字符的位置。
- `QuickSpanString`: 寻找字符串中下一个unicode字符的位置。
- `FirstBoundary`: 返回第一个在标准化unicode字符序列中的边界位置的索引。
- `firstBoundary`: 返回第一个在标准化unicode字符序列中的边界位置的索引。
- `FirstBoundaryInString`: 返回第一个在标准化unicode字符序列中的边界位置的索引。
- `NextBoundary`: 寻找下一个unicode字符的边界位置。
- `NextBoundaryInString`: 寻找下一个unicode字符的边界位置。
- `nextBoundary`: 寻找下一个unicode字符的边界位置。
- `LastBoundary`: 寻找最后一个unicode字符的边界位置。
- `lastBoundary`: 寻找最后一个unicode字符的边界位置。
- `decomposeSegment`: 将unicode字符序列进行分解。
- `lastRuneStart`: 返回最后一个unicode字符的起始位置。
- `decomposeToLastBoundary`: 将unicode字符序列进行分解，直到最后一个字符的边界位置。

