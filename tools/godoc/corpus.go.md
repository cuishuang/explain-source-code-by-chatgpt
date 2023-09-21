# File: tools/godoc/corpus.go

在Golang的Tools项目中，`corpus.go`文件是`godoc`工具的一部分，它负责管理文档的语料库。以下是对该文件中各功能的详细介绍：

1. Corpus结构体：`Corpus`结构体表示整个语料库。它包括以下字段：
   - `tree`：表示语料库的文件系统树。
   - `path`：表示语料库的根路径。
   - `index`：表示语料库的索引。

2. NewCorpus函数：`NewCorpus`函数用于创建一个新的语料库。它接受一个路径参数，并返回一个初始化后的`Corpus`结构体。

3. CurrentIndex函数：`CurrentIndex`函数返回语料库的当前索引。如果语料库中没有索引，它会尝试创建一个新的索引。

4. FSModifiedTime函数：`FSModifiedTime`函数返回语料库中最新的文件更改时间。

5. Init函数：`Init`函数用于初始化语料库。它会创建一个新的`Corpus`结构体，并将语料库的根路径设置为指定的路径。

6. initFSTree函数：`initFSTree`函数用于初始化语料库的文件系统树。它会遍历语料库的文件和目录，并将其映射到`Corpus`结构体的`tree`字段上。

这些功能一起工作，通过读取和解析代码包来构建一个完整的文档语料库。该语料库可被`godoc`工具使用来生成代码文档、搜索、浏览和查看Go程序的详细信息。

