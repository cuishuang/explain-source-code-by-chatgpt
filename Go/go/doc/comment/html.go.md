# File: html.go

html.go文件是Go标准库中的一个package，主要用于HTML模板和HTTP服务。

具体作用如下：

1. ParseInLocation：用于解析具有指定时区的时间字符串。

2. EscapeString和UnescapeString：用于处理HTML字符串中的特殊字符，避免XSS攻击。

3. HTMLEscape和HTMLUnescape：用于将HTML转义为实体和将实体转回原始HTML。

4. NewTokenizer：用于将HTML文本分解为各种令牌，包括标签、文本和属性。

5. NewEncoder和NewDecoder：用于转换HTML编码（例如从UTF-8到ISO-8859-1）。

此外，在http包中，HTML相关的函数也通常会调用此文件中的函数来进行处理。

总之，html.go文件提供了强大的HTML处理工具，可以方便地解析和生成HTML文本，以及在HTTP服务中提供对HTML文本的各种操作。




---

### Structs:

### htmlPrinter





## Functions:

### HTML





### block





### inc





### text





### escape





