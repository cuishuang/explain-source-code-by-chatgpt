# File: excelize/sheet.go

在Go生态excelize项目中，excelize/sheet.go文件是用于处理Excel中的表格(sheet)的功能。

1. NewSheet：创建新的表格(sheet)。
2. contentTypesReader：读取内容类型(content type)。
3. contentTypesWriter：写入内容类型(content type)。
4. getWorksheetPath：获取工作表路径。
5. mergeExpandedCols：合并扩展列。
6. workSheetWriter：写入工作表。
7. trimRow：修剪行。
8. trimCell：修剪单元格。
9. setContentTypes：设置内容类型。
10. setSheet：设置表格。
11. relsWriter：写入关系。
12. replaceRelationshipsBytes：替换关系 bytes。
13. SetActiveSheet：设置活动表格。
14. GetActiveSheetIndex：获取活动表格索引。
15. getActiveSheetID：获取活动表格ID。
16. SetSheetName：设置表格名称。
17. GetSheetName：获取表格名称。
18. getSheetID：获取表格ID。
19. GetSheetIndex：获取表格索引。
20. GetSheetMap：获取表格映射。
21. GetSheetList：获取表格列表。
22. getSheetMap：获取表格映射。
23. getSheetXMLPath：获取表格XML路径。
24. SetSheetBackground：设置表格背景。
25. SetSheetBackgroundFromBytes：从字节设置表格背景。
26. setSheetBackground：设置表格背景。
27. DeleteSheet：删除表格。
28. deleteAndAdjustDefinedNames：删除并调整定义名称。
29. deleteSheetFromWorkbookRels：从工作簿关系删除表格。
30. deleteSheetRelationships：删除表格的关系。
31. getSheetRelationshipsTargetByID：通过ID获取表格关系目标。
32. CopySheet：复制表格。
33. copySheet：复制表格。
34. getSheetState：获取表格状态。
35. SetSheetVisible：设置表格可见性。
36. setPanes：设置窗格。
37. SetPanes：设置窗格。
38. getPanes：获取窗格。
39. GetPanes：获取窗格。
40. GetSheetVisible：获取表格可见性。
41. SearchSheet：搜索表格。
42. searchSheet：搜索表格。
43. attrValToInt：属性值转为整数。
44. attrValToFloat：属性值转为浮点数。
45. attrValToBool：属性值转为布尔值。
46. SetHeaderFooter：设置页眉页脚。
47. GetHeaderFooter：获取页眉页脚。
48. ProtectSheet：保护表格。
49. UnprotectSheet：解除表格保护。
50. checkSheetName：检查表格名称。
51. SetPageLayout：设置页面布局。
52. newPageSetUp：新建页面设置。
53. setPageSetUp：设置页面设置。
54. GetPageLayout：获取页面布局。
55. SetDefinedName：设置定义名称。
56. DeleteDefinedName：删除定义名称。
57. GetDefinedName：获取定义名称。
58. GroupSheets：将表格分组。
59. UngroupSheets：取消分组表格。
60. InsertPageBreak：插入分页符。
61. insertPageBreak：插入分页符。
62. RemovePageBreak：移除分页符。
63. relsReader：读取关系。
64. prepareSheetXML：准备表格XML。
65. fillColumns：填充列。
66. makeContiguousColumns：使列连续。
67. SetSheetDimension：设置表格维度。
68. GetSheetDimension：获取表格维度。

这些函数都提供了丰富的功能，帮助用户在处理Excel表格时进行各种操作，包括创建、读取、写入、删除、复制、设置样式、修改属性等。

