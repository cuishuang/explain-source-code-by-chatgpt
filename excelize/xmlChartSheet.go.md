# File: excelize/xmlChartSheet.go

在excelize这个项目中，excelize/xmlChartSheet.go文件的作用是实现Chartsheet（图表工作表）相关的XML解析和生成功能。

xlsxChartsheet结构体表示一个图表工作表的定义，包含了工作表的属性和视图信息等。它具有以下字段：
- Name：图表工作表的名称。
- SheetViews：图表工作表的视图信息，包含了显示样式和窗口状态等。
- SheetFormatPr：图表工作表的格式属性，比如行高、列宽等。
- SheetProtection：图表工作表的保护属性，比如是否允许编辑等。
- CustomSheetViews：自定义的视图信息，比如工作表的缩放比例等。

xlsxChartsheetPr结构体表示图表工作表的属性，包含了图表工作表的属性信息。它具有以下字段：
- CodeName：图表工作表的VBA编程名称。
- FilterMode：图表工作表是否处于筛选模式。
- SyncHorizontal：图表工作表是否水平同步滚动。
- SyncVertical：图表工作表是否垂直同步滚动。

xlsxChartsheetViews结构体表示图表工作表的视图信息，包含了图表工作表的显示样式和窗口状态等。它具有以下字段：
- SheetView：具体的图表工作表视图信息。

xlsxChartsheetView结构体表示图表工作表的具体视图信息，包含了图表的显示样式和窗口状态等。它具有以下字段：
- WindowProtection：图表工作表窗口的保护属性，比如是否允许拖动等。
- ZoomScale: 缩放比例。
- WorkbookViewID：工作簿视图ID。

xlsxChartsheetProtection结构体表示图表工作表的保护属性，包含了是否允许编辑等信息。

xlsxCustomChartsheetViews结构体表示自定义的图表工作表的视图信息，比如工作表的缩放比例等。

xlsxCustomChartsheetView结构体表示自定义的图表工作表的具体视图信息，包含了工作表的缩放比例等。

这些结构体通过在xmlChartSheet.go文件中的解析和生成函数，实现了对图表工作表的XML解析和生成功能。具体而言，可以通过这些结构体和相关函数实现对图表工作表的创建、编辑和删除等操作。

