# File: tools/cmd/guru/what.go

在Golang的Tools项目中，`tools/cmd/guru/what.go`文件的作用是实现了`guru what`命令的功能，用于查询Go代码中的符号信息。

该文件中定义了多个结构体和函数，下面分别介绍它们的作用：

1. 结构体`whatResult`：用于存储符号信息查询的结果。它包含了以下字段：
   - `what`：表示查询的符号类型，例如"callers"表示调用该函数的地方，"definition"表示符号的定义，"describe"表示符号的描述等。
   - `guessImportPath`：表示符号所属的包的导入路径。
   - `segments`：表示符号所属的代码文件和行号信息。
   - `prefixLen`：表示在查询结果中给出的前缀长度。
   - `PrintPlain`：表示结果是否以普通文本形式打印。
   - `JSON`：表示结果是否以JSON格式输出。

2. 函数`what`：根据查询的符号和所在的文件位置，向用户打印该符号的信息。它接受以下参数：
   - `conf`：用于设置查询相关的配置。
   - `pos`：表示符号所在的代码位置。
   - `arg`：表示查询的符号。

3. 函数`guessImportPath`：根据给定的代码位置，猜测符号所属的包的导入路径。它接受以下参数：
   - `arg`：表示查询的符号。
   - `pos`：表示符号所在的代码位置。

4. 函数`segments`：获取符号所属的代码文件和行号信息。它接受以下参数：
   - `arg`：表示查询的符号。
   - `pos`：表示符号所在的代码位置。

5. 函数`prefixLen`：返回查询结果中给出的前缀长度。它接受以下参数：
   - `arg`：表示查询的符号。

6. 函数`PrintPlain`：将查询结果以普通文本的形式打印输出。它接受以下参数：
   - `fr *referrers`：表示反向引用信息。
   - `cl *ChangeList`：表示变更列表。
   - `pkg *Package`：表示符号所属的包信息。

7. 函数`JSON`：将查询结果以JSON格式输出。它接受以下参数：
   - `fr *referrers`：表示反向引用信息。
   - `cl *ChangeList`：表示变更列表。

通过使用上述函数和结构体，`tools/cmd/guru/what.go`文件实现了`guru what`命令的符号信息查询功能。

