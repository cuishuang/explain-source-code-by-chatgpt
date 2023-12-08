# File: excelize/sheetpr.go

在Go生态excelize项目中，excelize/sheetpr.go文件的作用是处理Excel工作表的页面属性。

- SetPageMargins函数用于设置工作表的页面边距。可以设置上、下、左、右四个边距的数值。

- GetPageMargins函数用于获取工作表的页面边距。

- prepareSheetPr函数是一个私有函数，用于准备工作表的页面属性。它会创建一个新的SheetPr元素，并将其添加到Sheet元素中。

- setSheetOutlineProps函数用于设置工作表的大纲属性。例如，可以设置工作表是否显示大纲符号、大纲符号的折叠状态和级别等。

- setSheetProps函数用于设置工作表的属性。可以设置工作表的名称、显示网格线、默认列宽、默认行高等。

- SetSheetProps函数是对setSheetProps函数的封装，用于将指定的工作表属性应用到工作表中。

- GetSheetProps函数用于获取工作表的属性。返回的是一个SheetPr结构体，包含了工作表的各种属性信息。

这些函数可以通过调用相关的API来操作Excel工作表的页面属性，如设置页面边距、设置大纲属性等。你可以根据自己的需求使用这些函数来定制和控制Excel工作表的外观和布局。

