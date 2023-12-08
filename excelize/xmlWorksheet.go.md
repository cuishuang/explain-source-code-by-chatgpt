# File: excelize/xmlWorksheet.go

在Go生态excelize项目中，excelize/xmlWorksheet.go文件是用于处理Excel文档中的工作表（worksheet）的。它定义了处理工作表所需的结构体和方法。

这个文件中包含以下结构体和方法：

1. xlsxWorksheet：表示Excel文档中的一个工作表。它包含了工作表的名称、索引、列的宽度、行的高度等信息。

2. xlsxDrawing和xlsxLegacyDrawing：表示工作表中的绘图。它包含了工作表中的图形、图片等绘图对象的信息。

3. xlsxHeaderFooter：表示工作表的页眉页脚。它包含了工作表的页眉页脚的内容和样式。

4. xlsxDrawingHF：表示工作表的页眉页脚中的绘图。它包含了工作表的页眉页脚中的图形、图片等绘图对象的信息。

5. xlsxPageSetUp：表示工作表的页面设置。它包含了工作表的打印设置、纸张大小、打印区域等页面设置的信息。

6. xlsxPrintOptions：表示工作表的打印选项。它包含了工作表的打印选项的信息。

7. xlsxPageMargins：表示工作表的页边距。它包含了工作表的上下左右的页边距的信息。

8. xlsxSheetFormatPr：表示工作表的格式设置。它包含了工作表的默认行高、列宽等格式设置的信息。

9. xlsxSheetViews和xlsxSheetView：表示工作表的视图。它包含了工作表的视图类型、网格线的显示等视图的信息。

10. xlsxSelection和xlsxPane：表示工作表的选择区域和窗口坐标。它包含了工作表的当前选择区域和窗口坐标的信息。

11. xlsxSheetPr：表示工作表的属性。它包含了工作表的隐藏、活动、展开等属性的信息。

12. xlsxOutlinePr：表示工作表的大纲设定属性。它包含了工作表的大纲设定相关的属性的信息。

13. xlsxSheetData和xlsxRow：表示工作表的数据。它包含了工作表中的数据行、单元格和值的信息。

14. xlsxSortState：表示工作表的排序设置。它包含了工作表的排序方式、排序列等排序设置的信息。

15. xlsxCustomSheetViews：表示工作表的自定义视图。它包含了工作表的自定义视图的信息。

16. xlsxBrk、xlsxRowBreaks、xlsxColBreaks和xlsxBreaks：表示工作表的分页符。它包含了工作表的行分页符、列分页符、分页符的信息。

17. xlsxCustomSheetView：表示工作表的自定义视图设置。它包含了工作表的自定义视图设置的信息。

18. xlsxMergeCell和xlsxMergeCells：表示工作表的合并单元格。它包含了工作表的合并单元格的信息。

19. xlsxDataValidations和xlsxDataValidation：表示工作表的数据验证。它包含了工作表的数据验证的信息。

20. xlsxC和xlsxF：表示工作表的列和行的选项。它包含了工作表的列和行的选项的信息。

21. xlsxSheetProtection：表示工作表的保护设置。它包含了工作表的保护设置的信息。

22. xlsxPhoneticPr：表示工作表的拼音属性。它包含了工作表的拼音属性的信息。

23. xlsxConditionalFormatting、xlsxCfRule、xlsxColorScale、xlsxDataBar、xlsxCfvo和xlsxIconSet：表示工作表的条件格式设置。它包含了工作表的条件格式的信息。

24. xlsxHyperlinks和xlsxHyperlink：表示工作表的超链接。它包含了工作表的超链接的信息。

25. xlsxTableParts和xlsxTablePart：表示工作表的表格式。它包含了工作表的表格式的信息。

26. xlsxPicture和xlsxLegacyDrawing：表示工作表的图片和图形。它包含了工作表的图片和图形的信息。

27. decodeX14SparklineGroups、decodeX14ConditionalFormattingExt、decodeX14ConditionalFormattings、decodeX14ConditionalFormatting和decodeX14CfRule：表示解码工作表中的扩展元素。

这些结构体和方法组成了excelize/xmlWorksheet.go文件，用于处理Excel文档中工作表的各种信息和属性。通过这些结构体和方法，可以读取、写入、修改和删除Excel工作表中的数据、样式、格式等内容。

