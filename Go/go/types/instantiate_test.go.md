# File: instantiate_test.go

instantiate_test.go文件是Go语言的标准库中的一个测试文件，主要目的是测试在Go语言中如何实例化变量（即创建变量并赋值）。测试包括使用各种数据类型、使用字面量、使用函数等方式进行实例化。

该文件主要包含了对于以下函数的测试：

1. new()函数：该函数用于分配内存空间，返回一个指向已分配空间的指针。

2. make()函数：该函数用于分配复杂数据类型的空间，比如数组、切片、映射和通道等。

3. 字面量：在Go语言中，可以直接使用字面量来实例化变量，例如使用“var a int = 10”这样的语句来创建一个整型变量。

4. 函数：可以通过函数来创建并初始化变量，并返回变量的值。

总之，该文件是为了确保创建变量的各种方式在Go语言中能够成功地进行，防止程序出现问题，保证代码的正确性。

## Functions:

### TestInstantiateEquality





### TestInstantiateNonEquality





### TestMethodInstantiation





### TestImmutableSignatures





### stripAnnotations





