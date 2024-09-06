# File: exp_fieldtrack_on.go

go/src/internal/exp_fieldtrack_on.go 文件是Go语言中的一个内部实现文件，主要作用是用于开启或关闭结构体域跟踪功能。

结构体域跟踪是一种针对具有内部转换或嵌套结构的Go语言结构的精确定位跟踪技术。它在编译期跟踪结构体中的字段，并为程序生成辅助信息。该技术可以大大简化程序员的工作，从而提高代码的可维护性和可读性。

exp_fieldtrack_on.go 文件中定义了一个名为fieldTrackEnabled的全局变量，该变量默认值为假。程序员可以在编译期启用该功能，通过将该变量值设为真来开启结构体域跟踪功能。

启用这个功能的方式如下：

go build -gcflags="-l -B" -tags fieldtrack

其中，-tags选项用来添加一个fieldtrack标签，告诉编译器开启结构体域跟踪功能。具体的实现细节和使用方式可以参考官方文档：https://golang.org/pkg/go/internal/fieldtrack/

