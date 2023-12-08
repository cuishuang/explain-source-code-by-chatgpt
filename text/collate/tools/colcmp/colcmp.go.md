# File: text/collate/tools/colcmp/colcmp.go

文件colcmp.go的作用是实现了一个用于比较排序结果的工具。该工具用于验证text/collate包中的排序算法的正确性。

以下是这些变量和结构体的详细介绍：

变量：
- doNorm：一个bool类型的变量，表示是否要对输入进行Unicode规范化。
- cases：一个包含不同大小写变换的字符串切片，用于在比较时考虑不同大小写的情况。
- verbose：一个bool类型的变量，表示是否输出详细的比较信息。
- debug：一个bool类型的变量，表示是否输出调试信息。
- locales：一个包含要测试的不同locale的字符串切片。
- col：一个字符串，表示要测试的collator类型。
- gold：一个包含用于验证排序结果的golden文件的路径字符串切片。
- usecmp：一个字符串，表示要使用的比较函数类型。
- cpuprofile：一个字符串，表示保存CPU profile的文件路径。
- exclude：一个包含要排除的测试的字符串切片。
- limit：一个整数，表示要限制的测试数量。
- lastLen：一个整数，表示上一次运行的输入长度。
- commands：一个字符串切片，表示要运行的命令。
- cmdSort：一个bool类型的变量，表示是否运行sort命令。
- cmdBench：一个bool类型的变量，表示是否运行bench命令。
- cmdRegress：一个bool类型的变量，表示是否运行regress命令。

结构体：
- Test：用于保存测试的结果的结构体。包含排序前和排序后的输入字符串切片。
- testCompare：用于保存比较结果的结构体。包含一个测试和实际的比较结果。
- testRestore：用于保存还原结果的结构体。包含一个测试和还原后的结果。
- Context：用于保存上下文信息的结构体。包含比较函数和unicode规范化函数。
- Command：用于保存命令信息的结构体。包含命令名称和帮助信息。

函数：
- failOnError：用于输出错误信息并返回错误状态码的函数。
- clear：用于清除缓冲区的函数。
- SetStatus：用于设置状态的函数。
- Start：用于开始计时的函数。
- Stop：用于停止计时并输出结果的函数。
- generateKeys：用于生成key的函数。
- Sort：对字符串切片进行排序的函数。
- Swap：用于交换切片中元素位置的函数。
- Less：用于比较两个元素大小的函数。
- Len：用于获取切片长度的函数。
- GenerateInput：用于生成输入的函数。
- Printf：用于输出格式化字符串的函数。
- Print：用于输出字符串的函数。
- assertBuf：用于断言缓冲区中的内容是否与预期相同的函数。
- flush：用于清空缓冲区的函数。
- parseTests：用于解析测试文件的函数。
- Test：对输入字符串切片进行排序并返回结果的函数。
- parseInput：用于解析输入字符串的函数。
- Name：用于获取测试名称的函数。
- runSort：运行sort命令的函数。
- runBench：运行bench命令的函数。
- keyStr：用于将排序的结果转换为字符串的函数。
- runRegress：运行regress命令的函数。
- runHelp：运行help命令的函数。
- main：工具的入口函数。根据命令行参数执行不同的命令。

