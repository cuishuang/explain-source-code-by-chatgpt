# File: tools/gopls/internal/astutil/purge.go

文件`purge.go`位于`Golang`项目`gopls`的`internal/astutil`包中，其目的是提供一些辅助函数来删除Go语言源代码中的函数体。下面我们将详细介绍该文件的作用和`PurgeFuncBodies`函数的功能。

`purge.go`中的几个函数都是围绕函数体的删除操作展开，其中主要涉及的是删除Go语言源代码文件中的函数体，以便进行一些代码检查、代码生成或修复等任务。以下是`purge.go`中的一些重要函数：

1. `PurgeFuncBodies`函数签名：`func PurgeFuncBodies(f *ast.File, p token.Position) []*ast.FuncDecl`，该函数的作用是在指定的`ast.File`中删除所有函数的函数体，并返回所有被删除函数的`ast.FuncDecl`列表。该函数接受一个`ast.File`类型的参数`f`，以及一个`token.Position`类型的参数`p`，用于标识函数操作的位置。函数通过遍历`ast.File`的语法树来找到所有的函数声明，然后将每个函数的函数体设置为空代码块，并将被删除函数的`ast.FuncDecl`添加到返回的列表中。

2. `PurgeAllFuncBodies`函数签名：`func PurgeAllFuncBodies(fset *token.FileSet, filenames []string) ([]*ast.File, error)`，该函数的作用是在指定的文件集合中删除所有函数的函数体，并返回更新后的`ast.File`列表。该函数接受一个`token.FileSet`类型的参数`fset`，用于管理文件集合，以及一个字符串类型的数组`filenames`，指定要进行操作的文件。通过`token.FileSet`和`filenames`，函数将文件逐个打开并分析，然后调用`PurgeFuncBodies`函数删除每个文件中的所有函数体。最后，函数将更新后的`ast.File`添加到返回的列表中。

3. `WritePurgedFiles`函数签名：`func WritePurgedFiles(fset *token.FileSet, filenames []string, output string, mode os.FileMode) ([]string, error)`，该函数的作用是将经过函数体清除的源代码写入到指定的输出文件中，并返回写入成功的文件列表。该函数接受一个`token.FileSet`类型的参数`fset`，用于管理文件集合，一个字符串类型的数组`filenames`，指定要进行操作的文件，一个字符串类型的参数`output`，指定输出文件的路径，以及一个`os.FileMode`类型的参数`mode`，指定输出文件的权限。函数会遍历`filenames`中的文件，并将每个文件的代码进行函数体清除后写入到输出文件中。最终，函数返回成功写入的文件列表。

总结来说，`purge.go`文件中的函数用于删除Go语言源代码中的函数体，从而用于诸如代码检查、代码生成或修复等任务。`PurgeFuncBodies`函数可以单独用于删除指定文件中的函数体，而`PurgeAllFuncBodies`函数则可以用于删除多个文件的函数体，并返回更新后的`ast.File`列表。`WritePurgedFiles`函数则进一步将经过函数体清除的源代码写入到指定的输出文件中。

