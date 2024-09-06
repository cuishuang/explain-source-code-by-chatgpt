# File: testing.go

testing.go文件是Go语言标准库中testing包的核心文件之一，主要提供了以下功能：

1. 测试框架的核心功能。testing包提供了一套完整的测试框架，testing.go中定义了TestMain、Run和其他与测试执行相关的方法和结构体。

2. 测试用例的运行和输出。testing.go中定义了T和B类型，分别代表着普通测试和基准测试。它们提供了Fail、Error、Fatal、Log等方法，用于运行和输出测试结果。

3. 单元测试和压力测试的支持。testing.go中定义了Timer、Benchmark等结构体和方法，用于对代码进行单元测试和压力测试。

4. 辅助方法和工具类。testing.go中还提供了各种辅助方法，如Parallel、Short、Skip、MainStart等，用于方便测试的编写和执行。

总之，testing.go文件是Go语言测试框架的核心文件之一，提供了丰富的测试支持和功能，方便用户编写高质量的测试代码。

## Functions:

### CommentsDo

函数CommentsDo() 是在testing包中定义的函数，其作用是将注释中以 “Output:” 或 “Order:” 开头的行（以下简称为目标行）提取出来，并执行evalCode函数和applyOutput函数。

具体来说，CommentsDo()的主要任务是遍历给定的测试函数的注释，提取出目标行，并将其作为参数传递给evalCode和applyOutput函数。evalCode函数将目标行中以":"开头的字符串作为Go语言代码来执行，如果有任何错误，那么函数将返回一个错误信息。 applyOutput函数会将目标行中以输出形式显示的内容与标准输出进行比较，以检查测试是否通过。

从上述介绍可以看出，CommentsDo() 的主要作用是用于自动化测试中，检查代码是否能产生预期的输出，并记录结果以供后续测试。这可以帮助开发人员减少手工测试工作量，提高代码质量和开发效率。



### CommentMap

CommentMap函数的作用是读取Go测试文件中的注释，将其转换为一个键值对的映射关系（map），其中键是测试函数的名称，值是该测试函数的注释。

具体来说，CommentMap函数会从测试文件中解析出测试函数的名称和其注释，然后将它们存储在一个CommentMap的变量中。当测试运行时，测试框架会使用这个CommentMap来通过测试函数的名称查找其对应的注释。

注释应该包含在测试函数声明之前，格式如下：

func TestXxx(t *testing.T) {
  // 注释
  // ...
}

其中，Xxx是测试函数名称，注释部分可以是任何合法的Go代码，但通常是对测试函数的简短描述或说明。

通过使用CommentMap函数，测试框架可以在运行测试时自动加载测试注释，帮助开发者更好地了解各个测试函数的作用和测试目的，从而更好地进行测试开发和调试。



