# File: excelize/sheetview.go

在Go生态excelize项目中，excelize/sheetview.go文件的作用是处理Excel文件中的Sheet View（工作表视图）相关的操作。Sheet View是Excel中用于设置工作表的显示方式、缩放比例和分割窗口等属性的一种机制。

- `getSheetView`函数用于获取指定工作表的Sheet View对象。它接受一个工作表名和一个Sheet View索引作为参数，并返回对应的Sheet View对象。
- `setSheetView`函数用于设置指定工作表的Sheet View对象。它接受一个工作表名、一个Sheet View索引和要设置的Sheet View对象作为参数，并将指定的Sheet View对象设置到对应的工作表中。
- `SetSheetView`函数是对`setSheetView`函数的封装，用于设置指定工作表的Sheet View对象。它接受一个工作表名和一个要设置的Sheet View对象作为参数，并将指定的Sheet View对象设置到对应的工作表中。
- `GetSheetView`函数是对`getSheetView`函数的封装，用于获取指定工作表的Sheet View对象。它接受一个工作表名作为参数，并返回对应的Sheet View对象。

这些函数主要用于通过代码设置和获取Excel文件中工作表的显示方式、缩放比例和分割窗口等属性，提供了对Excel文件的更细致的控制和定制化操作的能力。

