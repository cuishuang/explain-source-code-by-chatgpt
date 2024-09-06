# File: size.go

size.go是Go语言编译器的一个重要组件，用于生成目标平台上每个符号的大小和对齐方式的报告。

在Go语言中，编译器生成的每个二进制程序中，都会包含很多符号，例如变量、函数、结构体等。这些符号在编译过程中都会被赋予一定的大小，而且这个大小和对齐方式都会对程序的运行效率产生一定的影响。

size.go会在编译器生成目标文件的过程中，对每个符号的大小进行统计，并生成相应的报告，以帮助开发者优化程序的运行效率。

具体来说，size.go会遍历编译器生成的目标文件，并使用目标平台上的ABI（应用二进制接口）规范，计算每个符号的大小和对齐方式。然后，它会生成一个报告，其中包含了每个符号的名称、大小、对齐方式等信息，供开发者参考。

通过分析size.go生成的报告，开发者可以了解到每个符号的大小和对齐方式，在程序设计和优化过程中，可以采取相应的措施，以减少程序的内存占用和提高程序的运行效率。

总之，size.go是Go语言编译器中一个重要的组成部分，它负责生成目标文件中每个符号的大小和对齐方式的报告，帮助开发者优化程序的运行效率。




---

### Var:

### PtrSize





### RegSize





### SlicePtrOffset





### SliceLenOffset





### SliceCapOffset





### SliceSize





### StringSize





### SkipSizeForTracing





### MaxWidth





### CalcSizeDisabled





### defercalc





### deferredTypeStack





## Functions:

### typePos





### RoundUp





### expandiface





### calcStructOffset





### isAtomicStdPkg





### CalcSize





### CalcStructSize





### RecalcSize





### widthCalculated





### CheckSize





### DeferCheckSize





### ResumeCheckSize





### PtrDataSize





