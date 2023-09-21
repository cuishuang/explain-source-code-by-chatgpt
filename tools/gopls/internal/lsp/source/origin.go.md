# File: tools/gopls/internal/lsp/source/origin.go

在Golang的Tools项目中，`origin.go`文件位于`tools/gopls/internal/lsp/source`目录下，它的作用是处理源代码位置信息的原始来源。

在代码编辑器中，当用户进行鼠标悬停、跳转到定义等操作时，需要获取相应代码位置的原始信息，以便提供准确的结果。`origin.go`文件定义了与原始代码位置相关的结构体和函数，用于解析和处理这些信息。

下面是对`origin.go`文件中关键函数`containsOrigin`的作用进行详细介绍：

1. `containsOrigin`函数：

```go
func containsOrigin(ancestors []compareFuncData, pos token.Position) bool
```

该函数用于检查给定的代码位置`pos`是否包含在`ancestors`中的某个位置区间中。`ancestors`是一个由`compareFuncData`结构体组成的切片，每个`compareFuncData`结构体包含了一个位置区间（`start`和`end`表示）。

`containsOrigin`函数的运行逻辑如下：

- 首先，通过二分查找找到`ancestors`中包含`pos`的位置区间。
- 如果找到了包含`pos`的位置区间，则返回`true`，表示`pos`在该区间内。
- 如果没有找到包含`pos`的位置区间，则返回`false`，表示`pos`不在任何区间内。

该函数的作用是帮助确定某个代码位置`pos`是否位于给定的位置区间内，以此判断代码位置是否与给定的位置信息相关联。

总结：`tools/gopls/internal/lsp/source/origin.go`文件的主要作用是处理和解析源代码位置的原始信息。`containsOrigin`函数用于检查给定的代码位置是否位于指定区间内。通过这些功能实现，可以在代码编辑器中提供准确的代码导航、查找定义等功能。

