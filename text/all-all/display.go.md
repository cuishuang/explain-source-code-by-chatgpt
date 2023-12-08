# File: text/language/display/display.go

文件`display.go`是`text/language/display`包的一部分，该包提供了一些用于操作和获取语言、地区、脚本等显示相关信息的功能。

以下是对其中变量和结构体的详细介绍：

Variables（变量）:
1. Supported：存储已支持的显示格式列表。
2. Values：一个map，存储支持的值及其相应的显示格式。
3. matcher：一个用于匹配的函数。
4. Self：表示自定义内容。
5. self：表示默认的自定义值。

Structs（结构体）:
1. Formatter：显示格式的格式化器，用于将某个显示值格式化为符合预定格式的字符串。
2. Namer：用于获取某个值的名称。
3. languageNamer：用于获取某个语言的名称。
4. scriptNamer：用于获取某个脚本的名称。
5. regionNamer：用于获取某个地区的名称。
6. tagNamer：用于获取某个标签的名称。
7. Dictionary：在字典中查找对应的对象。
8. dictTags：一组标签和其对应的对象集合。
9. dictLanguages：一组语言和其对应的对象集合。
10. dictScripts：一组脚本和其对应的对象集合。
11. dictRegions：一组地区和其对应的对象集合。
12. SelfNamer：将自定义值（self）作为名称的标记。

Functions（函数）:
1. Format：将一个值格式化为指定的显示格式，并返回字符串。
2. Language：获取一个语言的显示信息。
3. Region：获取一个地区的显示信息。
4. Script：获取一个脚本的显示信息。
5. Tag：获取一个标签的显示信息。
6. init：初始化函数，在使用包前会被调用，用于初始化一些变量。
7. Languages：返回所有已知的语言。
8. langFunc：用于获取给定语言的显示信息。
9. name：获取一个值的名称。
10. Name：获取一个值的名称，并根据显示格式进行适当的格式化。
11. nonEmptyIndex：找到数组中第一个非空元素的索引。
12. Scripts：返回所有已知的脚本。
13. scriptFunc：用于获取给定脚本的显示信息。
14. Regions：返回所有已知的地区。
15. regionFunc：用于获取给定地区的显示信息。
16. Tags：返回所有已知的标签。
17. tagFunc：用于获取给定标签的显示信息。
18. lookup：在字典中查找给定值对应的对象。
19. getScript：获取给定标签的脚本信息。

总之，`display.go`文件定义了一系列变量、结构体和函数，用于处理和获取语言、地区、脚本等显示相关的信息，包括格式化显示值、获取名称以及查找相关信息等操作。

