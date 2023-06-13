# File: binutils.go

go/src/cmd/binutils.go这个文件是Go语言中内置的命令行工具"go tool objdump"和"go tool nm"的实现代码。这两个工具用于分析和查看Go语言编译后的二进制文件。

具体来说，binutils.go文件中定义了两个类型: nmCmd和objdumpCmd。这两个类型都实现了Cmd接口，包括Name方法、Usage方法、Short方法和Long方法，用于定义命令行工具的基本信息和使用方式。

nmCmd实现了Name方法、Usage方法和Short方法，用于定义"go tool nm"命令的基本信息和使用方式。同时，nmCmd还实现了Run方法，用于解析命令行参数，加载目标二进制文件，分析并打印出其中的符号信息。

objdumpCmd实现了Name方法、Usage方法和Short方法，用于定义"go tool objdump"命令的基本信息和使用方式。同时，objdumpCmd还实现了Run方法，用于解析命令行参数，加载目标二进制文件，分析并打印出其中的汇编代码和调试信息。

总的来说，binutils.go文件实现了Go语言中内置的两个二进制文件分析工具， 方便用户了解编译后的二进制文件的符号信息、汇编代码和调试信息。

