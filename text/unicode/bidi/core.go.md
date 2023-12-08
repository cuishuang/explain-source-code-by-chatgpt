# File: text/unicode/bidi/core.go

text/unicode/bidi/core.go文件是Go语言中text/unicode/bidi包的核心文件之一。该文件实现了双向文本算法（Bidirectional Text Algorithm，简称BTA），提供了处理混合方向文本的功能。

在该文件中，定义了多个结构体和函数来处理双向文本的解析、分析和处理。

1. 结构体：
- level：表示一个字符的嵌套等级。
- paragraph：表示一段文本中的一个段落，包含该段落的字符范围。
- directionalStatusStack：记录级联处理过程中的方向状态。
- isolatingRunSequence：表示从文本中提取的独立运行序列。

2. 函数：
- newParagraph：创建一个新的段落对象。
- Len：获取段落中字符的数量。
- run：从段落中获取指定索引位置的运行对象。
- determineMatchingIsolates：确定独立运行序列中每个隔离的末尾运行的索引。
- determineParagraphEmbeddingLevel：确定段落的嵌套等级。
- empty：检查段落是否为空。
- pop：弹出方向状态栈的顶部状态。
- depth：获取方向状态栈的深度。
- push：将方向状态推入栈中。
- lastEmbeddingLevel：获取最后一个嵌套等级。
- lastDirectionalOverrideStatus：获取最后一个方向覆盖状态。
- lastDirectionalIsolateStatus：获取最后一个方向隔离状态。
- determineExplicitEmbeddingLevels：确定显式嵌套等级。
- maxLevel：获取文本中最大的嵌套等级。
- isolatingRunSequence：从文本中提取独立运行序列。
- resolveWeakTypes：解析弱类型字符。
- resolveNeutralTypes：解析中立类型字符。
- setLevels：设置字符的嵌套等级。
- setTypes：设置字符的类型。
- resolveImplicitLevels：解析隐式嵌套等级。
- applyLevelsAndTypes：应用运行和字符的嵌套等级和类型。
- findRunLimit：在文本中找到运行的边界限制。
- assertOnly：验证字符只能有一个属性。
- determineLevelRuns：确定运行的嵌套等级。
- determineIsolatingRunSequences：确定独立运行序列。
- assignLevelsToCharactersRemovedByX9：为删除的字符分配嵌套等级。
- getLevels：获取字符的嵌套等级。
- getReordering：获取重新排序的索引数组。
- computeMultilineReordering：计算多行文本的重新排序。
- computeReordering：计算文本的重新排序。
- isWhitespace：检查字符是否为空白字符。
- isRemovedByX9：检查字符是否被X9规则删除。
- typeForLevel：根据嵌套等级获取字符的类型。
- validateTypes：验证字符的类型是否有效。
- validateParagraphEmbeddingLevel：验证段落嵌套等级是否有效。
- validateLineBreaks：验证换行符是否有效。
- validatePbTypes：验证段落嵌套等级和字符类型是否匹配。
- validatePbValues：验证段落嵌套等级和字符值是否匹配。

这些结构体和函数提供了双向文本处理所需的数据和算法支持，用于解决处理混合方向文本的问题，如嵌套等级的确定、字符类型的解析、字符重新排序等。

