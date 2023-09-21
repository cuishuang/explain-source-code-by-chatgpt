# File: tools/cmd/benchcmp/benchcmp.go

在Golang的Tools项目中，tools/cmd/benchcmp/benchcmp.go文件的作用是生成和比较Go语言性能测试（benchmark）的报告。它可以接收两个benchmark输出文件并将其内容进行对比，显示出这两个测试之间的性能差异。

该文件中的`changedOnly`变量用于控制是否只显示具有性能变化的benchmark测试。当设置为true时，将仅显示具有变化的测试；当设置为false时，将显示所有测试。

`magSort`变量用于控制对测试结果进行排序的方式。如果设置为true，则按照逆序排序，也就是将结果从大到小排列；如果设置为false，则按照正序排序，结果从小到大排列。

`best`变量用于控制是否只显示最好/最坏的测试结果。当该变量设置为true时，benchcmp只会显示最好和最坏的测试结果；当设置为false时，benchcmp会显示所有测试结果。

`main`函数是benchcmp程序的入口点。它解析命令行参数、读取benchmark输出文件、比较两个文件的结果，并根据设置的选项输出结果。

`fatal`函数用于打印错误信息，并终止程序的运行。

`parseFile`函数用于解析一个benchmark输出文件，并返回Benchmark类型的slice。

`selectBest`函数用于根据设置的`best`选项，过滤出最好或最坏的测试结果。

`formatNs`函数用于将纳秒（ns）时间格式化为字符串，用于输出报告。

这些函数的组合和实现，使得`benchcmp`工具能够方便地进行benchmark结果的比较与分析，并输出易读的报告。

