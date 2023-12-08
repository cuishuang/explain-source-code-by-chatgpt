# File: excelize/vmlDrawing.go

在Go生态中的excelize项目中，vmlDrawing.go文件的作用是处理Excel中的VML（Vector Markup Language）Drawing格式的相关操作。该文件包含了一系列结构体和处理方法，用于解析和生成VML Drawing数据。

下面逐个介绍这些结构体的作用：

1. vmlDrawing：表示VML Drawing的根对象，存储了页面上的所有Drawing元素。
2. xlsxShapeLayout：定义了形状的布局信息，如位置、大小等。
3. xlsxIDmap：存储了形状和XML元素ID之间的映射关系。
4. xlsxShape：表示一个形状，包括形状的类型、布局信息、边框、填充等。
5. xlsxShapeType：定义了形状的类型。
6. xlsxStroke：定义了形状的边框样式。
7. vPath：表示VML形状的路径。
8. vFill：表示VML形状的填充样式。
9. oFill：表示VML形状的透明度样式。
10. vShadow：表示VML形状的阴影样式。
11. vTextBox：表示VML形状的文本框。
12. xlsxDiv：定义了形状的分隔框。
13. vmlFont：定义了形状中的文本字体样式。
14. xClientData：存储了形状中的用户数据。
15. decodeVmlDrawing：解析一个Excel文件中的VML Drawing数据。
16. decodeShapeType：解析形状的类型。
17. decodeShape：解析形状的属性。
18. decodeShapeVal：解析形状的值。
19. decodeVMLFontU、decodeVMLFontI、decodeVMLFontB：分别解析形状中的字体样式的下划线、斜体、粗体属性。
20. decodeVMLFont：解析形状中的字体样式。
21. decodeVMLDiv：解析形状的分隔框。
22. decodeVMLTextBox：解析形状的文本框。
23. decodeVMLClientData：解析形状中的用户数据。
24. encodeShape：生成形状的XML字符串。
25. formCtrlPreset：定义了VML表单控件的预设值。
26. vmlOptions：定义了VML格式的导出选项。
27. FormControl：表示VML表单控件。

这些结构体和方法的集合组成了excelize中处理Excel VML Drawing格式的核心功能，通过解析和生成VML Drawing数据，可以实现对Excel文件中图形的读写操作。

