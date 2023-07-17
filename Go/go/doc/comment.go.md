# File: comment.go

comment.go 是 Go 语言标准库中注释相关的代码文件，主要定义了操作注释的函数和结构体。

具体来说，comment.go 定义了以下几个重要的类型和函数：

- Comment：表示源代码中的注释内容（除了文档注释）。Comment 结构体有一个 Pos 字段表示注释在源代码中的位置，以及一个 Text 字段表示注释的内容。
- PackageComment：表示 package 开头的注释内容。PackageComment 结构体包含一个 Comment 字段表示 package 注释的内容。
- CommentMap：表示一个文件中所有注释的映射。CommentMap 结构体是一个 map 类型，键为注释在源代码中的位置，值为一个 Comment 类型的结构体，表示该位置上的注释内容。
- ExtractComment：用于提取源代码中的注释信息，返回一个 CommentMap 类型的值。该函数会忽略掉文档注释。
- ExtractPackageComment：用于提取 package 注释信息，返回一个 PackageComment 类型的值。
- FindAdjacentComments：用于查找一个位置上、或者一个位置周围的注释内容，返回一个 Comment 类型的切片。

这些类型和函数提供了一些基本的操作注释的方法，可用于实现自动化文档生成工具和注释检查工具等功能。

## Functions:

### ToHTML





### ToText





