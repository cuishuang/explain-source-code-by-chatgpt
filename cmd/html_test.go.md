# File: html_test.go

html_test.go是Go语言内置命令`go test`命令的测试文件，其作用是测试html包中的函数和方法是否正常运作。

在该文件中，包含了一系列针对html包的测试用例，验证了在不同的情况下html包中的函数和方法是否符合预期的输出结果。这些测试用例包括：

- TestEscapeString：测试EscapeString函数是否能够正确地将特殊字符进行转义，避免XSS漏洞。
- TestUnescapeString：测试UnescapeString函数是否能够正确地将特殊字符进行反转义。
- TestEscapeReader：测试escapeReader类型是否正确地将特殊字符进行转义，避免XSS漏洞。
- TestPretty：测试Pretty函数是否能够正确地将HTML格式化输出，便于debug和阅读。
- TestCompact：测试Compact函数是否能够正确地将HTML紧凑输出，节省空间。
- TestStripTags：测试StripTags函数是否能够正确地将HTML标记过滤掉，只保留纯文本内容。

通过这些测试用例，我们可以保证html包中的函数和方法在不同场景下都能够正确地工作，确保应用程序的安全和性能。

