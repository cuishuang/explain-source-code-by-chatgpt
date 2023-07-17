# File: symbolizer.go

symbolizer.go是一个Go语言程序，它的作用是将二进制文件中的地址映射成其对应的源代码位置。具体地说，它是一个Symbolizer，即一个符号化器。

在软件开发中，人们通常会使用调试器来调试程序，以查找错误或进行性能分析。调试器需要知道程序中变量、函数等符号地址和源代码位置之间的关系，以便在Debug时提供有足够有效的信息。

然而，程序经过编译后，符号地址经常和源代码位置不一致。这时，Symbolizer就显得尤为重要了。Symbolizer能够读取一个二进制文件中的调试信息（Debug Info），解析并将其中的地址映射到其对应的源代码位置上。它能够解决符号地址和源代码位置不同的问题，帮助调试器更加方便地进行程序调试和分析。

在symbolizer.go文件中，主要实现了ElfSymbolizer和MachOSymbolizer两个Symbolizer。

ElfSymbolizer是针对Linux、BSD等系统的二进制文件的调试信息。它能够读取文件中的Elf符号表（Symbol Table）和Debug Info，并将地址映射到源代码位置上。 

MachOSymbolizer是针对macOS系统二进制文件的调试信息。它能够读取文件中的Macho符号表和Debug Info，同样将地址映射到源代码位置上。 

综上，symbolizer.go的主要作用是实现符号化器，将程序中的符号地址映射成对应的源代码位置，为调试器提供有足够有效的信息，进而方便程序调试和分析。

