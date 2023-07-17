# File: support.go

该文件是Go语言的示例程序，用于展示如何在Go程序中使用C语言的库函数。该文件定义了一个名为"Average"的函数，该函数通过调用C语言的库函数来计算一组数字的平均值。

该文件包含了以下几个函数：

1. average - 该函数接收一个整数数组作为参数，并通过调用C语言的库函数来计算这些数字的平均值。

2. main - 该函数是程序的主入口点，该函数调用average函数，并打印出计算结果。

3. cgoHelloWorld - 该函数是使用CGO调用C函数的示例。

4. sum - 该函数计算一个整数数组的总和，并返回结果。

在该文件中，我们使用了CGO技术，这是一个特殊的工具，用于将Go代码与C代码相互操作。它允许我们在Go程序中调用C函数，同时还允许C函数调用Go函数。使用CGO技术，我们可以利用C语言库中的功能来扩展Go语言的功能。

总之，support.go文件提供了一个简单的示例，展示了如何在Go程序中使用C语言的库函数，并且介绍了CGO的基本概念和用法。




---

### Var:

### fakeLines





### fakeLinesOnce





### predeclared








---

### Structs:

### fakeFileSet





### fileInfo





### anyType





### derivedInfo





### typeInfo





## Functions:

### assert





### errorf





### pos





### setLines





### chanDir





### Underlying





### String





### splitVargenSuffix





