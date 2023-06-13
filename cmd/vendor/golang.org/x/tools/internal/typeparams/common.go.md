# File: common.go

common.go是Go语言标准库中cmd目录下的一个文件，主要是定义了一些常用的变量和函数。它的作用是为了方便cmd目录下的各个命令包引用一些常用的变量和函数。

具体来说，common.go文件中定义的一些常用变量包括：

- version：当前的Go语言版本号；
- GOARCH、GOOS：当前系统的操作系统和CPU架构；
- (
  默认的环境变量名字（例如$GOOS和$GOARCH）和对应的值；
- (
  sBuildID：程序的编译ID；
- goBootstrap：标记程序是否是在Go语言自举时使用的，如果是，则为true。

此外，common.go文件中还定义了一些常用的函数，主要包括：

- fatalf(format string, args ...interface{})：输出错误信息并退出程序；
- mkTempDir(dir, prefix string) string：在指定的目录下创建一个临时目录，并返回其路径；
- maybePrintProfile()：根据环境变量的设置选择性地输出性能数据；
- setExitStatus(code int)：设置程序的退出状态码。

这些变量和函数都是在Go语言标准库中不同的命令包中频繁使用的，通过将它们放在common.go文件中，减少了代码冗余，提高了代码复用性和整体代码的可维护性。

