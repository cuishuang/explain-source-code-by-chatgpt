# File: text/unicode/rangetable/rangetable.go

在Go的text包中，rangetable.go文件主要用于定义和操作Unicode字符范围表。

rangetable.go中定义了以下几个结构体：byRune、byRange16和byRange32。其中byRune结构体用于存储字符范围的元数据和操作方法，byRange16和byRange32结构体则是byRune的衍生结构体，用于处理对应的字符范围。

byRune结构体有以下作用：
- 存储字符范围的元数据，包括字符范围的起始和结束代码点、字符范围在rangetable中的索引以及相关的字符分类信息。
- 提供Len()方法，用于返回字符范围的长度。
- 提供Swap()方法，用于交换两个字符范围的位置。
- 提供Less()方法，用于判断当前字符范围是否小于另一个字符范围。
- 提供Visit()方法，用于访问字符范围中的每个字符。
- 提供Assigned()方法，用于判断某个字符是否属于当前字符范围。

byRange16和byRange32结构体的作用类似于byRune，不同之处在于它们只处理16位和32位Unicode字符范围。

New()函数用于创建一个新的Unicode字符范围表。它通过读取Unicode数据文件（如UnicodeData.txt和EastAsianWidth.txt）来生成字符范围表的数据，并将其填充到rangetable中。

Len()函数用于返回Unicode字符范围表中的字符范围数量。

Swap()函数用于交换两个字符范围在Unicode字符范围表中的位置。

Less()函数用于判断一个字符范围是否小于另一个字符范围。

Visit()函数用于遍历一个字符范围，并对范围内的每个字符执行指定的操作。

Assigned()函数用于判断某个字符是否属于Unicode字符范围表中的一个字符范围。

