# File: promql/parser/prettier.go

在 Prometheus 项目中，promql/parser/prettier.go 文件的作用是提供一个用于美化 PromQL 查询表达式的工具。它处理了对查询表达式进行格式化，使其更易读和整洁。

文件中的 maxCharactersPerLine 变量定义了一行代码的最大字符数。当进行代码美化时，如果某一行的字符数超过了这个值，就会进行换行操作。

下面是对文件中几个重要函数的介绍：

1. Prettify(query string, indent string): 这个函数是整个工具的入口。它接收一个待美化的查询表达式字符串和缩进字符串作为参数，对查询表达式进行解析和格式化，并返回美化后的表达式字符串。

2. Pretty(node Node, indent string): 这个函数是 Prettify 函数的辅助函数，用于递归地遍历查询表达式的语法树。它接收一个语法树节点和缩进字符串作为参数，根据节点的类型和内容对参数进行格式化，并返回处理后的字符串。

3. getCommonPrefixIndent(lines []string): 这个函数用于获取多行代码中的公共缩进字符串。它接收一个字符串数组作为参数，遍历数组中的所有行，找出它们的最长公共前缀，并返回公共前缀作为缩进字符串。

4. needsSplit(line string): 这个函数用于判断一行代码是否需要进行换行。它接收一行代码作为参数，判断该行代码的长度是否超过了预定义的 maxCharactersPerLine 值，若超过则返回 true，否则返回 false。

5. indent(str string, indent string): 这个函数用于给一段字符串添加缩进。它接收一个字符串和一个缩进字符串作为参数，将缩进字符串添加到输入字符串的每一行之前，并返回添加缩进后的字符串。

这些函数共同协作，实现了对查询表达式的解析和格式化，使其在可读性和整洁性方面得到了优化。

