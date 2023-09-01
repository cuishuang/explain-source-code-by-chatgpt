# File: client-go/util/jsonpath/jsonpath.go

client-go/util/jsonpath/jsonpath.go文件的作用是实现了对JSON数据进行筛选和提取的功能，具体通过使用JSONPath语法来实现。JSONPath是一种在JSON数据中进行路径导航和筛选的语言，类似于XPath。

jsonpath.go文件中定义了几个重要的结构体，它们分别是：

1. JSONPath：表示一个JSONPath表达式，包含了解析后的路径和筛选条件。
2. JSONPathOptions：表示JSONPath执行的选项，包含了一些可选配置参数。
3. JSONPathOutputOptions：表示输出JSON结果时的选项，包含了一些可选参数。

下面是这些结构体的具体作用和用法：

1. New：构造一个JSONPath对象，参数为JSONPath表达式字符串。
2. AllowMissingKeys：设置JSONPath执行时是否允许缺失的字段，默认为false，即如果字段不存在将会报错。
3. Parse：解析JSONPath表达式字符串，返回一个JSONPath对象。
4. Execute：执行JSONPath表达式，参数为待筛选的JSON数据和JSONPath选项，返回匹配的结果集。
5. FindResults：从JSON数据的某个节点开始，根据JSONPath表达式查找满足条件的结果集。
6. EnableJSONOutput：设置JSONPath执行时是否允许输出JSON结果，默认为false，即输出字符串形式的结果。
7. PrintResults：将JSONPath执行的结果以合适的方式输出，可根据JSONPathOutputOptions进行格式化配置。
8. walk：递归遍历JSON数据中的字段，并根据解析后的JSONPath表达式进行路径导航和筛选操作。
9. evalInt：根据给定的字段路径，从JSON数据中获取并返回整型值。
10. evalFloat：根据给定的字段路径，从JSON数据中获取并返回浮点型值。
11. evalBool：根据给定的字段路径，从JSON数据中获取并返回布尔型值。
12. evalList：根据给定的字段路径，从JSON数据中获取并返回列表值。
13. evalIdentifier：根据给定的字段路径，从JSON数据中获取并返回字段的值。
14. evalArray：根据给定的字段路径，从JSON数据中获取并返回数组。
15. evalUnion：根据给定的字段路径，从JSON数据中获取并返回所有满足条件的字段值的结果集。
16. findFieldInValue：在JSON数据的某个字段中查找指定的字段名。
17. evalField：根据给定的字段路径，从JSON数据的某个节点中获取并返回字段的值。
18. evalWildcard：根据给定的字段路径，从JSON数据中获取并返回通配符（*）匹配的字段值。
19. evalRecursive：根据给定的字段路径，从JSON数据中递归获取并返回字段的值。
20. evalFilter：根据指定的条件对结果集进行过滤操作。
21. evalToText：将结果集转化为文本形式。

