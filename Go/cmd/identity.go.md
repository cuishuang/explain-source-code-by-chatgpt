# File: identity.go

identity.go是Go标准库中cmd目录下的一个文件，它的作用是生成和打印出给定文件的基本身份和属性信息。在命令行中运行“go tool dist list”命令时，它将被调用。

具体而言，identity.go主要实现了以下几个功能：

1. 获取文件的基本信息：包括文件名、文件大小、文件最后修改时间等信息。

2. 获取文件的SHA256哈希值：identity.go会对文件的内容进行哈希处理，从而生成一个独一无二的SHA256哈希值。

3. 打印文件的属性信息：最后，identity.go会将上述信息打印出来，包括文件名、文件大小、SHA256哈希值等。

总体上来说，identity.go的作用是辅助开发者在识别和确认不同文件之间的差异，确保文件的完整性和安全性。




---

### Structs:

### typePair





## Functions:

### Identical





### IdenticalIgnoreTags





### IdenticalStrict





### identical





