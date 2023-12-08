# File: text/unicode/norm/triegen.go

在Go的text项目中，`text/unicode/norm/triegen.go`这个文件的作用是用于生成Unicode标准化的字典数据结构。

`normCompacter`是一个辅助结构体，用于压缩Unicode标准化的字典数据。它包含以下字段：
- `mostFrequentStride`：表示字符在最常用的范围内所占的空间大小。
- `countSparseEntries`：标识稀疏条目数。
- `Size`：标识标准化字典的尺寸。
- `Store`：表示Unicode字符的编码范围。
- `Handler`：标准化处理程序。
- `Print`：用于打印标准化处理过程的调试信息。

以下是`triegen.go`中的一些功能函数的解释：
- `newNormCompacter`：创建一个新的`normCompacter`实例。
- `calcSizes`：计算标准化字典的各个部分的尺寸。
- `convertCases`：将Unicode字符的大小写转换规则转换为字典的编码形式。
- `generateData`：生成标准化字典的数据。
- `generateProperties`：生成标准化字典的属性数据。
- `generateDecomp`：生成标准化字典的分解数据。
- `generateCompositions`：生成标准化字典的组合数据。
- `generateCanonicals`：生成标准化字典的规范数据。
- `readCasefold`：读取Unicode字符的大小写转换规则。
- `readDecomposition`：读取Unicode字符的分解数据。
- `readComposition`：读取Unicode字符的组合数据。
- `writeTrie`：将生成的标准化字典写入输出流。

