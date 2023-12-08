# File: excelize/merge.go

在Go生态excelize项目中，excelize/merge.go文件的作用是处理Excel中单元格的合并和拆分。

MergeCell这几个结构体的作用如下：

1. Rect：表示一个矩形区域，包含起始单元格和结束单元格的坐标；

2. MergeCell：表示一个合并的单元格，包含合并单元格的坐标和样式；

3. UnmergeCell：表示一个要拆分的单元格，包含要拆分单元格的坐标。

这些结构体用于描述Excel中的合并和拆分操作。

以下是这些函数的作用：

1. GetMergeCells：获取所有的合并单元格，返回一个MergeCell结构体的切片；

2. overlapRange：判断两个矩形区域是否有重叠，如果有重叠，则返回重叠部分的矩形区域；

3. flatMergedCells：将所有的合并单元格扁平化为一个矩形区域的切片；

4. mergeOverlapCells：合并所有相邻的重叠单元格；

5. mergeCell：合并指定的单元格；

6. GetCellValue：获取指定单元格的值；

7. GetStartAxis：获取合并单元格的起始坐标；

8. GetEndAxis：获取合并单元格的结束坐标。

这些函数用于处理合并和拆分操作，从而实现对Excel中单元格的合并和拆分操作。

