# File: cmd/yamlfmt/yamlfmt.go

在Kubernetes项目中，cmd/yamlfmt/yamlfmt.go文件的作用是实现一个用于格式化和验证YAML文件的命令行工具。该工具用于自动调整和标准化Kubernetes配置文件中的缩进、空格、换行等格式，以提高文件的可读性和可维护性。

以下是对于main、fetchYaml和streamYaml这几个函数的详细介绍：

1. main函数：作为程序入口点，负责处理命令行参数和调用其他函数来执行相应的操作。它使用flag包来解析命令行标志，并根据标志的值选择执行格式化、验证或打印版本等操作。

2. fetchYaml函数：该函数用于从命令行参数传递的URL或文件路径中读取YAML配置文件。它根据参数的前缀（http://、https://等）来确定文件的来源类型，并相应地选择使用不同的方法来获取文件内容。

3. streamYaml函数：该函数负责处理并格式化YAML文件。它使用了yaml.v3包提供的功能来解析和序列化YAML数据。streamYaml函数首先将输入的YAML文件解析为k8s.io/apimachinery中的Unstructured对象，然后通过设置缩进、空格等选项，将其重新序列化为格式化的YAML格式，并将结果通过标准输出打印出来。

这些函数共同协作，使得yamlfmt工具能够读取、格式化和输出YAML文件，以帮助开发人员和操作人员提高工作效率和代码质量。

