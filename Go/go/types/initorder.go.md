# File: initorder.go

initorder.go是一个Go语言的例子程序，用于演示Go语言中包(Package)的初始化顺序。

在Go语言中，包的初始化顺序是很重要的，因为它可以影响到程序的正确性。Go语言的包(Package)可以通过init()函数来进行初始化操作，这些初始化操作将在程序运行前执行。

initorder.go文件中包含了三个包：package main、package a和package b。这三个包中都定义了init()函数，用于演示Go语言的包初始化顺序。

在initorder.go文件中，main包的init()函数首先被执行，然后是package a中的init()函数，最后是package b中的init()函数。这是因为init()函数的执行顺序是按照包导入的顺序来决定的。

通过这个例子程序，我们可以学习到Go语言包初始化的规则，进而提高程序的正确性和质量。




---

### Structs:

### dependency





### graphNode





### nodeSet





### nodeQueue





## Functions:

### initOrder





### findPath





### reportCycle





### cost





### add





### dependencyGraph





### Len





### Swap





### Less





### Push





### Pop





