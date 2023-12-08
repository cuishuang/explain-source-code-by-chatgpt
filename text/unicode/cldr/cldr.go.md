# File: text/unicode/cldr/cldr.go

在Go的text项目中，text/unicode/cldr/cldr.go文件的作用是为Unicode CLDR（Common Locale Data Repository）提供解析和检索功能。

Unicode CLDR是一个用于提供众多语言和地区相关数据的项目，这些数据包括日期格式、货币符号、数字格式、时区等，旨在帮助开发人员实现国际化和本地化。

在cldr.go文件中，有三个重要的变量：DraftBCP47、DraftCurrency和DraftTime，它们分别表示BCP47、货币和时间的"草案"（draft）级别。这些级别包括正式（Official）、草案（Contributed）和废弃（Deprecated）等。

CLDR结构体表示整个CLDR，它包含了各种语言和地区的相关数据。Draft结构体表示CLDR中的草案。makeCLDR函数用于从CLDR XML文件创建CLDR结构体。BCP47函数用于解析BCP47规范的标签。ParseDraft函数用于解析草案级别字符串，转换为Draft结构体。String函数用于将Draft结构体转换为字符串。SetDraftLevel函数用于设置草案级别。RawLDML函数将LDML XML文件转换为rawLDML结构体。LDML函数将rawLDML结构体转换为LDML结构体。Supplemental函数用于获取补充信息的Supplemental结构体。Locales函数用于获取CLDR中定义的所有语言和地区。Get函数用于根据给定的语言和地区标签，获取相应的数据。

总而言之，cldr.go文件提供了解析和检索Unicode CLDR数据的功能，并包含各种函数和结构体来处理不同类型的数据。

