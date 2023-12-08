# File: excelize/errors.go

在Go生态excelize项目中，errors.go文件定义了一些错误常量、结构体和函数，用于处理在处理Excel文件过程中可能出现的各种错误。

这些错误常量包括：

1. ErrAddVBAProject：添加VBA项目错误
2. ErrAttrValBool：属性值布尔类型错误
3. ErrCellCharsLength：单元格字符长度错误
4. ErrCellStyles：单元格样式错误
5. ErrColumnNumber：列数错误
6. ErrColumnWidth：列宽错误
7. ErrCoordinates：坐标错误
8. ErrCustomNumFmt：自定义数字格式错误
9. ErrDataValidationFormulaLength：数据验证公式长度错误
10. ErrDataValidationRange：数据验证范围错误
11. ErrDefinedNameDuplicate：定义名称重复错误
12. ErrDefinedNameScope：定义名称范围错误
13. ErrExistsSheet：工作表已存在错误
14. ErrExistsTableName：表名已存在错误
15. ErrFontLength：字体长度错误
16. ErrFontSize：字体大小错误
17. ErrFormControlValue：表单控件值错误
18. ErrGroupSheets：分组工作表错误
19. ErrImgExt：图片扩展名错误
20. ErrInvalidFormula：无效的公式错误
21. ErrMaxFilePathLength：文件路径长度超出限制错误
22. ErrMaxRowHeight：行高超出限制错误
23. ErrMaxRows：行数超出限制错误
24. ErrNameLength：名称长度错误
25. ErrOptionsUnzipSizeLimit：选项解压缩大小限制错误
26. ErrOutlineLevel：大纲级别错误
27. ErrParameterInvalid：无效参数错误
28. ErrParameterRequired：缺少必需参数错误
29. ErrPasswordLengthInvalid：密码长度无效错误
30. ErrSave：保存错误
31. ErrSheetIdx：工作表索引错误
32. ErrSheetNameBlank：工作表名称为空错误
33. ErrSheetNameInvalid：无效的工作表名称错误
34. ErrSheetNameLength：工作表名称长度错误
35. ErrSheetNameSingleQuote：工作表名称含有单引号错误
36. ErrSparkline：活动中的图表错误
37. ErrSparklineLocation：图表位置错误
38. ErrSparklineRange：图表范围错误
39. ErrSparklineStyle：图表样式错误
40. ErrSparklineType：图表类型错误
41. ErrStreamSetColWidth：设置列宽错误
42. ErrStreamSetPanes：设置窗格错误
43. ErrTotalSheetHyperlinks：工作表超链接总数错误
44. ErrUnknownEncryptMechanism：未知的加密机制错误
45. ErrUnprotectSheet：解除保护工作表错误
46. ErrUnprotectSheetPassword：解除保护工作表密码错误
47. ErrUnprotectWorkbook：解除保护工作簿错误
48. ErrUnprotectWorkbookPassword：解除保护工作簿密码错误
49. ErrUnsupportedEncryptMechanism：不支持的加密机制错误
50. ErrUnsupportedHashAlgorithm：不支持的哈希算法错误
51. ErrUnsupportedNumberFormat：不支持的数字格式错误
52. ErrWorkbookFileFormat：工作簿文件格式错误
53. ErrWorkbookPassword：工作簿密码错误

这些错误常量对应的结构体用于在发生错误时返回错误信息，并包含了错误信息的相关属性和方法。

这些函数用于创建和处理不同类型的错误：

1. Error：创建普通错误
2. newCellNameToCoordinatesError：创建单元格名称转坐标错误
3. newCoordinatesToCellNameError：创建坐标转单元格名称错误
4. newFieldLengthError：创建字段长度错误
5. newInvalidAutoFilterColumnError：创建无效的自动筛选列错误
6. newInvalidAutoFilterExpError：创建无效的自动筛选表达式错误
7. newInvalidAutoFilterOperatorError：创建无效的自动筛选操作符错误
8. newInvalidCellNameError：创建无效的单元格名称错误
9. newInvalidColumnNameError：创建无效的列名错误
10. newInvalidExcelDateError：创建无效的Excel日期错误
11. newInvalidLinkTypeError：创建无效的链接类型错误
12. newInvalidNameError：创建无效的名称错误
13. newInvalidRowNumberError：创建无效的行号错误
14. newInvalidSlicerNameError：创建无效的切片器名称错误
15. newInvalidStyleID：创建无效的样式ID错误
16. newNoExistTableError：创建不存在的表格错误
17. newNotWorksheetError：创建非工作表错误
18. newPivotTableDataRangeError：创建数据透视表数据范围错误
19. newPivotTableRangeError：创建数据透视表范围错误
20. newStreamSetRowError：创建设置行错误
21. newUnknownFilterTokenError：创建未知的筛选器令牌错误
22. newUnsupportedChartType：创建不支持的图表类型错误
23. newUnzipSizeLimitError：创建解压缩大小限制错误
24. newViewIdxError：创建视图索引错误

这些函数用于创建特定类型的错误对象，并返回错误信息。

