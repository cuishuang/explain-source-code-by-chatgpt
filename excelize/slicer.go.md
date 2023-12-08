# File: excelize/slicer.go

在Go生态excelize这个项目中，excelize/slicer.go文件的作用是处理Excel的切片器（Slicer）相关的操作。切片器是Excel中的一种数据筛选工具，它允许用户以交互方式筛选数据。

SlicerOptions这几个结构体分别有以下作用：
- SlicerOptions：切片器选项，存储切片器的属性，如样式、位置、源数据等。
- SlicerStyleOptions：切片器样式选项，存储切片器的样式属性，如背景色、边框等。
- SlicerColumnOptions：切片器列选项，存储切片器列的属性，如字段、数据源等。

下面是各个函数的作用：
- AddSlicer：向指定的工作表中添加切片器，并返回切片器在工作表中的索引。
- parseSlicerOptions：解析并返回切片器的选项值。
- countSlicers：计算工作表中的切片器数量，返回结果为切片器数量。
- countSlicerCache：计算切片器缓存数量，返回结果为缓存数量。
- getSlicerSource：获取切片器的源数据。
- addSheetSlicer：将切片器添加到指定的工作表上。
- addSheetTableSlicer：将切片器添加到指定的表格上。
- addSlicer：向切片器列表中添加切片器。
- genSlicerName：生成切片器的名称。
- genSlicerCacheName：生成切片器缓存的名称。
- setSlicerCache：设置切片器缓存。
- slicerReader：读取切片器数据。
- timelineReader：读取时间线数据。
- addSlicerCache：向切片器缓存列表中添加切片器缓存。
- addPivotCacheSlicer：向透视表缓存添加切片器。
- addDrawingSlicer：向绘图中添加切片器。
- addWorkbookSlicerCache：向工作簿的切片器缓存列表中添加切片器缓存。

