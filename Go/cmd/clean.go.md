# File: clean.go

在 Go 语言中，clean.go 文件是编译命令的一部分。其作用是在构建和安装 Go 语言程序时执行清理操作。这个文件中定义了一个函数 CleanCmd，它接收命令行参数并执行清理操作。

清理操作的主要目的是删除之前构建过程中生成的目标文件（可执行程序、库文件、中间文件、缓存文件等）。这些文件可能占用大量存储空间并影响性能，因此清理操作非常重要。

在执行清理操作时，clean.go 使用了一些命令行选项和环境变量来控制其行为。其中一些选项包括：

- -i：移除安装的目标文件。
- -r：递归地清理目标文件。
- -n：不实际执行清理操作，而是显示将要删除的文件列表。
- -x：显示每个删除操作的详细输出。

通过这些选项和其他环境变量，clean.go 提供了一些灵活的配置选项，使得它可以适用于不同类型的构建和安装场景。

总之，clean.go 文件的作用是帮助开发者在构建 Go 语言程序时执行清理操作，以确保生成的目标文件占用空间最少、执行效率最高。




---

### Var:

### CmdClean





### cleanI





### cleanR





### cleanCache





### cleanFuzzcache





### cleanModcache





### cleanTestcache





### cleaned





### cleanDir





### cleanFile





### cleanExt





## Functions:

### init





### runClean





### clean





### removeFile





