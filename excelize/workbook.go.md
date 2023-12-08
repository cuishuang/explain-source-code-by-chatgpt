# File: excelize/workbook.go

在Go生态excelize项目中，文件workbook.go是用于处理Excel工作簿的文件。它包含了一系列用于操作工作簿的函数。

1. SetWorkbookProps: 此函数用于设置工作簿的属性，例如标题、作者、主题等。

2. GetWorkbookProps: 此函数用于获取工作簿的属性，返回一个包含所有属性的结构体。

3. ProtectWorkbook: 此函数用于保护工作簿的内容，可以设置密码以防止未经授权的修改。

4. UnprotectWorkbook: 此函数用于取消保护工作簿的内容，即去除密码保护。

5. setWorkbook: 此函数用于设置工作簿的路径。

6. getWorkbookPath: 此函数用于获取工作簿的路径。

7. getWorkbookRelsPath: 此函数用于获取工作簿的关系路径。

8. deleteWorkbookRels: 此函数用于删除工作簿的关系。

9. workbookReader: 此函数用于读取工作簿的内容。

10. workBookWriter: 此函数用于写入工作簿的内容。

11. setContentTypePartRelsExtensions: 此函数用于设置内容类型部分的关系扩展。

12. setContentTypePartImageExtensions: 此函数用于设置内容类型部分图像扩展。

13. setContentTypePartVMLExtensions: 此函数用于设置内容类型部分VML扩展。

14. addContentTypePart: 此函数用于添加内容类型部分。

15. removeContentTypesPart: 此函数用于删除内容类型部分。

这些函数提供了对Excel工作簿的各种操作，包括设置属性、保护工作簿、读取和写入内容等。可以根据需求使用这些函数来操作Excel工作簿。

