# File: under.go

under.go是一个命令行工具，用于查找Go源代码中未导出的函数、变量和常量，包括私有方法和字段，在Go语言中称为“未导出符号”。

未导出符号在Go语言中是一种信息隐藏机制，它们只能在同一个包中被访问，不能被其他包引用。这种隐藏机制可以使代码更好地维护和测试，同时也可以防止不经意地泄漏实现细节，提高代码的安全性。

但是，在进行代码重构或调试等工作时，有时也需要查找未导出符号，这时可以使用under工具快速定位它们所在的文件和行号。

under工具的使用方法如下：

```shell
usage: under [-strict] [-tests] <patterns>

The under tool prints the file name and line number for each instance of a
private identifier's use within a specified set of Go packages.

  -strict
      Only match exact identifier names.

  -tests
      Include test files.

Each package is specified as an import path. Example:

  under -tests io/ioutil bufio bytes.Reader.Read
```

```
使用方式: under [-strict] [-tests] <patterns>

under工具可以打印指定一组Go包中的私有标识符的使用实例所在的文件名和行号。

  -strict
      只匹配完全的标识符名称。

  -tests
      包括测试文件。

每个包都是通过导入路径来指定的。例如：

  under -tests io/ioutil bufio bytes.Reader.Read
```

为了使用under工具，您需要首先安装Go编程语言的环境，并将其添加到您的操作系统的PATH环境变量中。然后，您可以在命令行中使用“go get”命令安装under工具：

```shell
go get golang.org/x/tools/cmd/under
```

安装完成后，您可以使用“under”命令搜索未导出的符号，命令格式如下：

```shell
under [-strict] [-tests] <patterns>
```

其中，patterns是要搜索的未导出符号名称列表，可以使用通配符匹配多个模式。例如，“bufio bytes.Reader.Read”表示在“bufio”包中查找“bytes.Reader.Read”方法。

如果不指定“-tests”选项，则under将只搜索默认的Go包。如果指定了“-tests”选项，则还会搜索测试文件。

如果指定“-strict”选项，则搜索将匹配完全相同的名称，否则将忽略名称的大小写和标点符号。

搜索结果将按照文件名称和行号的顺序显示。如果找不到匹配模式的符号，则不会产生输出。

## Functions:

### under





### coreType





### coreString





### match





