# File: whole_function.go

whole_function.go是Go语言源代码的一部分，主要用于标准库中的testing包，用于比较两个函数是否等效，也就是相同的输入不同的输出是否都符合预期。

该文件中的主要函数是WholeFunctionEqual，它接受两个参数，分别为待比较的函数f和g，返回一个布尔值，表示两个函数是否等效。该函数的实现首先调用reflect.TypeOf()函数来比较f和g的函数签名是否相同，然后使用reflect.ValueOf()函数来调用f和g，并比较它们的返回值是否相同。

除了WholeFunctionEqual函数外，该文件还定义了一些辅助函数，比如toThrowPanic()函数用于判断一个函数是否会抛出异常。在测试中，这些函数可以被用来验证使用相同输入的两个函数是否产生相同的输出，从而确保函数的正确性。

总之，whole_function.go文件的作用是提供一个用于测试Function是否等效的函数实现，使得开发人员可以使用该函数来确保程序的正确性。

## Functions:

### Foo





### Example





