# File: excelize/cell.go

在Go生态项目excelize中，cell.go文件是用于处理Excel单元格相关的操作。它定义了一些变量和函数来处理单元格类型、公式、超链接等。

下面是对cell.go文件中相关变量和结构体作用的详细介绍：

- cellTypes变量：用于定义了单元格的类型常量，包括字符串、数字、公式、时间等。

- CellType结构体：表示单元格的类型，包含了单元格的值和格式。

- FormulaOpts结构体：表示公式的选项，包括公式的计算缓存和计算模式等。

- HyperlinkOpts结构体：表示超链接的选项，包括链接的类型和链接的地址等。

下面是对cell.go文件中相关函数的作用的详细介绍：

- GetCellValue：获取单元格的值。

- GetCellType：获取单元格的类型。

- SetCellValue：设置单元格的值。

- String：将一个接口类型转换为字符串类型。

- hasValue：检查单元格是否有值。

- removeFormula：移除单元格中的公式。

- setCellIntFunc：设置整数类型值的单元格函数。

- setCellTimeFunc：设置时间类型值的单元格函数。

- setCellTime：设置时间类型值的单元格。

- setCellDuration：设置时长类型值的单元格。

- SetCellInt：设置整数类型值的单元格。

- setCellInt：设置整数类型值的单元格。

- SetCellUint：设置无符号整数类型值的单元格。

- setCellUint：设置无符号整数类型值的单元格。

- SetCellBool：设置布尔类型值的单元格。

- setCellBool：设置布尔类型值的单元格。

- SetCellFloat：设置浮点类型值的单元格。

- setCellFloat：设置浮点类型值的单元格。

- SetCellStr：设置字符串类型值的单元格。

- setCellString：设置字符串类型值的单元格。

- sharedStringsLoader：加载共享字符串。

- setSharedString：设置共享字符串。

- trimCellValue：修剪单元格的值。

- setCellValue：设置单元格的值。

- setInlineStr：设置内联字符串。

- setStr：设置字符串。

- getCellBool：获取布尔类型值的单元格。

- setCellDefault：设置默认类型的单元格。

- getCellDate：获取日期类型值的单元格。

- getValueFrom：从单元格中获取的值。

- SetCellDefault：设置默认类型的单元格。

- GetCellFormula：获取单元格的公式。

- SetCellFormula：设置单元格的公式。

- setSharedFormula：设置共享公式。

- countSharedFormula：计算共享公式的个数。

- GetCellHyperLink：获取单元格的超链接。

- SetCellHyperLink：设置单元格的超链接。

- getCellRichText：获取富文本类型值的单元格。

- GetCellRichText：获取富文本类型值的单元格。

- newRpr：创建富文本的rpr标签。

- newFont：创建字体标签。

- setRichText：设置富文本。

- SetCellRichText：设置富文本类型值的单元格。

- SetSheetRow：设置工作表的行。

- SetSheetCol：设置工作表的列。

- setSheetCells：设置工作表的单元格。

- prepareCell：准备单元格。

- getCellStringFunc：获取字符串类型值的单元格函数。

- formattedValue：格式化值。

- getCustomNumFmtCode：获取自定义数字格式代码。

- prepareCellStyle：准备单元格的样式。

- mergeCellsParser：合并单元格解析器。

- checkCellInRangeRef：检查单元格是否在指定范围内引用。

- cellInRange：判断单元格是否在指定范围内。

- isOverlap：判断两个合并单元格是否有重叠。

- parseSharedFormula：解析共享公式。

- getSharedFormula：获取共享公式。

- shiftCell：移动单元格。

