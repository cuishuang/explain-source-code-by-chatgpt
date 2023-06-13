# File: head.go

head.go是一个命令行工具，用于显示文件的头几行。该工具的作用是快速查看文件的开头内容，也可用于筛选文件的前几行，并将其作为另一个文件的输入。head命令在UNIX和类UNIX系统中非常常用。

该工具的使用方法为： 

```
Usage of head:
  -c bytes
      Print the first bytes of each file
  -n lines
      Print the first lines of each file (default 10)
```

参数-c可以用于输出文件的前几个字节，参数-n用于输出文件的前几行，两者均可配合文件名使用。例如：

```
$ head -n 5 file.txt
```

上述命令会输出文件file.txt的前5行内容。

head.go的实现原理是基于Go语言的标准库中的bufio和os包实现的。它首先打开文件，然后使用bufio.Scanner读取文件的每一行，并输出到标准输出流中，直到输出的行数达到指定的行数或读取完整个文件为止。如果设置了-c参数，则它会按照字节的方式读取文件。

总的来说，head.go是一款简单实用的工具，可以在控制台中快速查看文件的开头内容，对于快速定位文件的某些信息非常有帮助。

