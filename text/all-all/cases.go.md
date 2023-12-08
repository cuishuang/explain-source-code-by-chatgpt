# File: text/cases/cases.go

text/cases/cases.go文件是Go的text项目中的一个文件，其中定义了一些变量、结构体和函数，用于处理文本的大小写转换和标题处理。

NoLower变量表示不转换小写的选项，Compact变量表示压缩的选项。

Caser结构体表示一个文本的大小写转换器，Option结构体定义了转换的选项，options结构体用于承载多个Option选项。

Bytes函数将一个字符串按照指定的选项进行大小写转换，并返回转换后的结果。

String函数类似于Bytes函数，但是输入和输出都是字符串类型。

Reset方法用于重置大小写转换器的状态。

Transform方法根据给定的大小写转换器和选项对一个文本进行转换，并返回转换后的结果。

Span函数用于计算一个字符串的边界。

Upper方法将一个字符串转换为大写。

Lower方法将一个字符串转换为小写。

Title方法将一个字符串转换为标题格式。

Fold方法将一个字符串转换为大小写折叠形式。

getOpts函数从一个选项列表中获取相应的选项。

noLower函数检查选项列表中是否包含NoLower选项。

compact函数检查选项列表中是否包含Compact选项。

HandleFinalSigma函数用于处理希腊字母的最后一个字符。

ignoreFinalSigma函数检查选项列表中是否包含忽略最后一个希腊字母的选项。

handleFinalSigma函数根据选项列表对最后一个希腊字母进行处理。

