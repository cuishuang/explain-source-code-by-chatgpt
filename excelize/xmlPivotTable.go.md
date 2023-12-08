# File: excelize/xmlPivotTable.go

在Go生态excelize项目中，excelize/xmlPivotTable.go文件主要用于处理Excel中的数据透视表（PivotTable）功能。

xlsxPivotTableDefinition结构体表示Excel中数据透视表的定义，包括格式、源数据、名称等信息。

xlsxLocation结构体表示透视表的位置和大小。

xlsxPivotFields结构体表示透视表的字段。

xlsxPivotField结构体表示透视表字段的属性，包括名称、数据类型、汇总方式等。

xlsxItems结构体表示透视表字段的值。

xlsxItem结构体表示透视表字段的单个值。

xlsxAutoSortScope结构体表示透视表字段自动排序的范围。

xlsxRowFields结构体表示透视表行字段。

xlsxField结构体表示透视表字段的设置，如过滤、标签等。

xlsxRowItems结构体表示透视表行字段的值。

xlsxI结构体表示透视表中行的索引。

xlsxX结构体表示透视表中的位置。

xlsxColFields结构体表示透视表列字段。

xlsxColItems结构体表示透视表列字段的值。

xlsxPageFields结构体表示透视表页面字段。

xlsxPageField结构体表示透视表页面字段的设置，如过滤、标签等。

xlsxDataFields结构体表示透视表数据字段。

xlsxDataField结构体表示透视表数据字段的设置，如数据类型、汇总方式等。

xlsxConditionalFormats结构体表示透视表条件格式。

xlsxPivotTableStyleInfo结构体表示透视表样式信息。

以上结构体主要用于描述和存储Excel中数据透视表相关的信息，通过这些结构体可以对透视表进行创建、修改和删除等操作。

