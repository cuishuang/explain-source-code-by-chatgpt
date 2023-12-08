# File: excelize/xmlCore.go

在Go生态Excelize这个项目中，excelize/xmlCore.go文件的主要作用是处理Excel文件的元数据，包括DocProperties、decodeDcTerms、decodeCoreProperties、xlsxDcTerms和xlsxCoreProperties这几个结构体的定义和相关函数的实现。

1. `DocProperties` 结构体用于定义Excel文件的文档属性，如标题、主题、作者等。

2. `decodeDcTerms` 函数用于解析Excel文件的 Dublin Core（DC）元数据，它是一种描述文档的标准，通过 XML 存储并嵌入在 Excel 文件中。

3. `decodeCoreProperties` 函数是用于解析Excel文件的核心属性，包括创建日期、修改日期、文档分类等信息。

4. `xlsxDcTerms` 结构体用于定义 Excel 的 Dublin Core（DC）元数据对应的 XML 格式。

5. `xlsxCoreProperties` 结构体定义了 Excel 的核心属性对应的 XML 格式。

这些结构体和函数的作用是将 Excel 文件中的元数据解析出来并进行处理，使得用户可以获取和修改 Excel 文件的文档属性和核心属性。同时，也可以使用这些结构体和函数来创建新的 Excel 文件并设置元数据。这些元数据信息对于文件的描述、分类、版本控制等有着重要的作用。

