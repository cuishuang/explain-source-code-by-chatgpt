# File: tools/gopls/internal/lsp/testdata/builtins/builtin_go117.go

文件 `builtin_go117.go` 的作用是提供 Golang 1.17 中内置的函数和类型的示例代码，用于 `gopls/internal/lsp/testdata` 目录下的测试用途。

以下是该文件中的函数及其作用的介绍：

1. `func Name() string`：
   - 功能：返回当前文件的名称。
   - 作用：用于测试函数名称的提取和返回是否正确。

2. `func funcWithArgsAndReturn() (bool, int, string)`：
   - 功能：定义了一个有参数和返回值的函数。
   - 作用：测试函数参数和返回值的提取和匹配是否正确。

3. `func go117SpecificFunc() int`：
   - 功能：定义了一个只在 Golang 1.17 中新增的函数。
   - 作用：测试对于特定版本新增的函数的支持是否正确。

4. `func unusedFunc()`：
   - 功能：定义了一个未被使用的函数。
   - 作用：测试未使用的函数是否会触发警告或错误信息。

5. `func unusedType()`：
   - 功能：定义了一个未被使用的类型。
   - 作用：测试未使用的类型是否会触发警告或错误信息。

总体而言，`builtin_go117.go` 文件提供了一组用于测试 Golang 1.17 内置函数和类型支持的示例代码，通过这些代码可以验证 `gopls` 工具是否正确解析和处理这些特性。

