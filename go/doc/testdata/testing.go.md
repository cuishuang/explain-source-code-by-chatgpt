# File: testing.go

testing.go是Go语言中testing包中的一个主要源文件，它主要用于定义一些供测试使用的辅助函数和结构体。

具体来说，testing.go文件中定义了如下几个重要的类型和函数：

1. testing.T结构体：表示一个测试用例，包含许多测试相关的方法，如Error、Fail、FailNow等。每个测试函数都会创建一个测试用例对象，用于检查函数的正确性。

2. testing.M结构体：表示一个单元测试集合，包含多个测试用例。使用该结构体可以方便地将多个测试函数组织在一起。

3. testing.B结构体：使用该结构体可以进行性能测试。与testing.T结构体类似，testing.B结构体也有一系列相关的方法来进行测试统计和输出。

4. testing.C结构体：表示一个补充测试（Benchmark）函数。Benchmark函数用于测试某个功能的性能，通过该结构体可以获取测试的输入数据和统计测试结果。

5. testing.Short()函数：该函数用于判断当前测试是否为短时间测试。短时间测试通常用于快速检查函数的常见问题，因此某些高质量测试可能需要长时间运行，这时可以通过该函数来判断当前测试是否为短时间测试。如果是短时间测试，则可以跳过一些本应进行的测试来节省时间。

总之，testing.go是Go语言中testing包的核心文件，它定义了一系列用于测试的结构体和函数，能够方便地进行单元测试和性能测试。




---

### Var:

### short





### chatty





### match





### memProfile





### memProfileRate





### cpuProfile





### timeout





### cpuListStr





### parallel





### cpuList





### timer








---

### Structs:

### common





### T





### InternalTest





## Functions:

### Short





### decorate





### Fail





### Failed





### FailNow





### log





### Log





### Logf





### Error





### Errorf





### Fatal





### Fatalf





### Parallel





### tRunner





### Main





### report





### RunTests





### before





### after





### startAlarm





### stopAlarm





### alarm





### parseCpuList





