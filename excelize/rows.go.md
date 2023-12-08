# File: excelize/rows.go

在Go生态excelize这个项目中，excelize/rows.go文件包含了与Excel表格行相关的功能和结构体。

1. duplicateHelperFunc：这几个变量是用于辅助生成行数据副本时的复制函数。通过使用这些函数，可以复制行数据中的样式、格式、合并单元格等属性。

2. Rows：Rows结构体表示一个Excel表格的行。它包含了当前行的索引、行数据以及与该行相关的一些操作方法。

3. rowXMLIterator：rowXMLIterator结构体用于迭代解析XML文档的行数据。它可以按照一定的规则从XML中读取出行数据，并返回给调用方。

4. GetRows：GetRows函数用于获取指定范围内的行数据。通过设置起始行和结束行的索引，可以获取这些行的全部数据。

5. Next：Next函数用于将迭代器移动到下一行的位置，并返回该行的数据。

6. GetRowOpts：GetRowOpts函数用于获取指定行的选项。选项包括行高、是否可见、大纲级别等。

7. Error：Error函数用于获取行操作过程中的错误信息。

8. Close：Close函数用于关闭行操作，释放资源。

9. Columns：Columns函数用于返回指定行的列数。

10. extractRowOpts：extractRowOpts函数用于提取行的选项数据。

11. appendSpace：appendSpace函数用于在行数据的切片中追加空白数据。

12. rowXMLHandler：rowXMLHandler函数用于处理行的XML数据。

13. getFromStringItem：getFromStringItem函数用于从字符串中获取XML数据项。

14. xmlDecoder：xmlDecoder函数用于解码XML数据。

15. SetRowHeight、getRowHeight、GetRowHeight：这几个函数用于设置、获取行高。

16. sharedStringsReader：sharedStringsReader函数用于读取共享字符串。

17. SetRowVisible、GetRowVisible：这两个函数用于设置、获取行的可见性。

18. SetRowOutlineLevel、GetRowOutlineLevel：这两个函数用于设置、获取行的大纲级别。

19. RemoveRow：RemoveRow函数用于删除指定的行。

20. InsertRows：InsertRows函数用于在指定位置插入一定数量的行。

21. DuplicateRow、DuplicateRowTo：这两个函数用于复制行数据到指定位置。

22. duplicateConditionalFormat、duplicateDataValidations、duplicateMergeCells：这几个函数用于复制行数据时处理行数据中的条件格式、数据验证和合并单元格。

23. checkRow、hasAttr：这两个函数用于检查行数据是否符合格式和判断是否存在某个属性。

24. SetRowStyle：SetRowStyle函数用于设置行的样式。

25. convertRowHeightToPixels：convertRowHeightToPixels函数用于将行高转换为像素单位。

