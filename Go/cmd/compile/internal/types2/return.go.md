# File: return.go

go/src/cmd/return.go是Go语言标准库中的一个文件，其作用是实现函数返回值的语法。

具体来说，return.go中定义了return语句的语法和行为，包括：

1. return语句可以在函数中任何位置出现，用于提前退出函数并返回指定的值（或没有值）。

2. 如果函数声明了返回类型，return语句必须返回相应类型的值或零值。

3. 如果函数声明了多个返回值，return语句可以返回多个值（用逗号分隔）。

4. return语句可以用于defer语句中，使得当函数执行完后，最终结果仍然是期望的。

5. 如果函数中使用了named return values，return语句可以省略返回值，因为named return values已经将返回值声明过了。

6. 如果函数声明了panic类型的返回值，return语句在引发panic时自动触发。

除了return语句本身，return.go还定义了一些和返回值相关的辅助函数，如在没有声明返回类型的情况下推断返回类型、判断函数是否有返回值等。

总之，return.go是Go语言标准库中十分重要的一个文件，负责实现函数和返回值之间的关联和语法规则。

## Functions:

### isTerminating





### isTerminatingList





### isTerminatingSwitch





### hasBreak





### hasBreakList





### hasBreakCaseList





### hasBreakCommList





