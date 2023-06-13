# File: builtin_test.go

builtin_test.go文件是Go语言的内置包testing中的文件之一。该文件中定义了一些内置的测试函数，用于测试Go语言中的一些内置函数和方法的行为和正确性。

该文件中定义的测试函数包括测试append、copy、len、cap、make、new等内置函数以及对应类型的方法，如slice类型上的append方法、map类型上的delete方法等。

测试函数通过使用Go语言的testing包中的测试框架来实现测试。测试框架包括TestMain、TestCase和TestSuite等函数和类型，用于组织和运行测试。测试函数会通过输入一些测试数据，调用被测试函数或方法，然后使用断言函数来验证测试结果是否为期望值。

由于Go语言的内置函数和方法是使用Go语言本身实现的，因此在对它们进行测试时需要保证被测试函数和方法的行为正确性，以保证Go语言程序的正确性和稳定性。因此，builtin_test.go文件在Go语言编程中扮演着重要的角色。

## Functions:

### TestBuiltin





