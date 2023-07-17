# File: check_test.go

check_test.go是Go语言中cmd包中的一个测试文件，主要用于测试Go语言源码中的代码检查命令。该文件中包含了多个测试函数，用于测试源码中代码检查的各种功能、选项和参数的正确性。

具体来说，check_test.go测试了以下命令和参数：

1. go vet：这个命令用于静态检查Go语言源码中的常见错误。check_test.go中的测试函数测试了go vet命令在不同情况下的输出是否正确。

2. go fmt：这个命令用于格式化Go语言源码。check_test.go中的测试函数测试了go fmt命令在不同情况下的输出是否正确。

3. go fix：这个命令用于将旧版本的Go语言代码更新为新版本的代码风格。check_test.go中的测试函数测试了go fix命令在不同情况下的输出是否正确。

4. goimports：这个命令用于自动添加、删除和排序Go语言源码中的import语句。check_test.go中的测试函数测试了goimports命令在不同情况下的输出是否正确。

总之，check_test.go是Go语言源代码检查命令的测试文件，通过测试不同命令和参数的输出是否正确，保证了这些命令的正确性和可靠性。




---

### Var:

### haltOnError





### verifyErrors





## Functions:

### parseFiles





### unpackError





### absDiff





### parseFlags





### testFiles





### TestManual





### TestCheck





### TestSpec





### TestExamples





### TestFixedbugs





### TestLocal





### testDirFiles





### testDir





