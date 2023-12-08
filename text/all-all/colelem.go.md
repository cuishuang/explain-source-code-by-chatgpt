# File: text/collate/build/colelem.go

在Go的text项目中，text/collate/build/colelem.go文件的作用是定义了用于排序和比较Unicode字符串的相关数据结构和函数。它是构建基元元素（collation element）的核心文件。基元元素是排序算法的基础，通过比较它们可以确定字符串的排序顺序。

文件中包含了一些与基元元素相关的结构体:
1. rawCE：表示原始的基元元素。包含了权重（weight）和值（value）等信息。
2. Primary：rawCE的一种特殊类型，表示主权重，并有拓展值（expansion）。
3. secondary：rawCE的一种特殊类型，表示次级权重。
4. tertiary：rawCE的一种特殊类型，表示三级权重。
5. quaternary：rawCE的一种特殊类型，表示四级权重。

以下是一些在该文件中定义的函数及其作用：
1. makeRawCE：根据给定的权重和值创建一个rawCE。
2. makeCE：根据给定的权重和值创建一个基元元素。
3. makeContractIndex：创建一个用于记录字符串的合并位置的索引。
4. makeExpandIndex：创建一个用于记录字符串的拓展位置的索引。
5. makeExpansionHeader：创建一个用于记录拓展值长度的头部。
6. makeDecompose：基于Unicode分解规则创建一个映射表。
7. implicitPrimary：返回给定基元元素的隐式主权重。
8. convertLargeWeights：将给定的权重转换为大权重表示形式。
9. nextWeight：根据给定的权重计算下一个权重。
10. nextVal：根据给定的rawCE计算下一个拓展值。
11. compareWeights：比较两个权重的大小。
12. equalCE：比较两个基元元素是否相等。
13. equalCEArrays：比较两个基元元素数组是否相等。

这些函数和结构体一起提供了在排序和比较Unicode字符串时所需的低级操作和工具函数。

