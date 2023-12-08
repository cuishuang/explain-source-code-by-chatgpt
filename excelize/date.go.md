# File: excelize/date.go

在excelize这个项目中，excelize/date.go这个文件是负责处理Excel中日期相关功能的代码文件。下面我们来逐个介绍这些变量和函数的作用。

- daysInMonth: 定义了每个月的天数，是一个长度为12的切片，分别表示1月到12月的天数。
- excel1900Epoc: 定义了Excel中日期的起始时间（1900年1月1日）。
- excel1904Epoc: 定义了Excel中Mac版日期的起始时间（1904年1月1日）。
- excelMinTime1900: 定义了Excel日期中可以表示的最小时间（1900年1月0日）。
- excelBuggyPeriodStart: 定义了一个Excel在日期计算中存在问题的起始时间（1900年2月29日）。

下面是这些函数的作用：

- timeToExcelTime: 将时间类型转换为Excel中的日期格式。
- shiftJulianToNoon: 将时间从儒略日转换为儒略日的中午时间。
- fractionOfADay: 计算一个时间在一天中所占的比例。
- julianDateToGregorianTime: 将儒略日转换为公历时间。
- doTheFliegelAndVanFlandernAlgorithm: 根据日期计算天文参数。
- timeFromExcelTime: 将Excel中的日期格式转换为时间类型。
- ExcelDateToTime: 将Excel中的日期格式转换为时间类型，并指定Excel日期起始时间。
- isLeapYear: 判断一个年份是否为闰年。
- getDaysInMonth: 获取指定年份月份的天数。
- validateDate: 验证日期的合法性。
- formatYear: 对年份进行格式化处理。

这些函数主要用于日期和时间的转换、验证和计算等相关操作，可以方便地进行Excel文件中日期数据的处理。

