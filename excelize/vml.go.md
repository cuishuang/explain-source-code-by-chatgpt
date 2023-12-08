# File: excelize/vml.go

在Go生态excelize项目中，excelize/vml.go文件的作用是处理Excel中的VML（Vector Markup Language）绘图对象，该文件包含一些函数和结构体用于操作和处理Excel中的VML对象。

formCtrlPresets变量是一个map，用于存储预定义的控件类型和其属性值。

FormControlType是一个枚举类型，定义了不同类型的控件，如按钮、复选框、下拉框等。

GetComments函数用于获取指定工作表中的所有批注。

GetSheetComments函数用于获取指定工作表上的所有批注的ID和位置。

AddComment函数用于在指定位置添加批注。

DeleteComment函数用于删除指定位置的批注。

addComment函数用于添加一条批注。

countComments函数用于计算工作表上的批注数量。

commentsReader和commentsWriter用于读取和写入批注信息。

AddFormControl函数用于向指定位置添加控件。

DeleteFormControl函数用于删除指定位置的控件。

countVMLDrawing函数用于计算工作表上的VML绘图对象个数。

decodeVMLDrawingReader和vmlDrawingWriter用于读取和写入VML绘图对象。

addVMLObject函数用于向工作表添加VML对象。

prepareFormCtrlOptions函数用于处理控件的属性值。

formCtrlText函数用于处理控件的显示文本。

addFormCtrl、addFormCtrlShape和addDrawingVML函数用于添加控件到工作表上。

GetFormControls函数用于获取工作表上的所有控件。

extractFormControl函数用于提取控件的属性值。

extractAnchorCell函数用于提取VML对象的位置信息。

extractVMLFont函数用于提取VML字体对象的属性值。

上述函数和变量的作用是实现对Excel中VML对象的读取、写入、处理和操作，例如添加、删除、获取批注以及添加、删除、获取控件等操作。

