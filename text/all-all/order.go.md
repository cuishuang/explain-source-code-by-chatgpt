# File: text/collate/build/order.go

在Go的text项目中，text/collate/build/order.go文件的作用是定义了排序规则的数据结构和相关操作。

logicalAnchor结构体表示排序时的逻辑锚点，用于连接不同的排序规则，每个排序规则都会有一个逻辑锚点。

entry结构体表示排序规则中的一个项，包含了该项在排序规则中的位置、编码信息和常见形式（如大小写转换）等。

sortedEntries是entry的切片，用于存储排序规则的所有项，按照位置从小到大排序。

ordering结构体表示排序规则的顺序，包含一个逻辑锚点和一组排序项。

String方法返回ordering的字符串表示。

skip方法返回跳过一定数量排序项后的ordering。

expansion方法返回将排序项从逻辑锚点位置扩展到指定长度后的ordering。

contraction方法返回将排序项从指定长度缩短到逻辑锚点位置后的ordering。

contractionStarter方法返回在逻辑锚点之前的排序项是否可以用作缩短的起点。

nextIndexed方法返回ordering中下一个已编码的排序项的位置。

remove方法从ordering中移除指定位置的排序项。

insertAfter方法在指定位置之后插入一个排序项。

insertBefore方法在指定位置之前插入一个排序项。

encodeBase方法返回编码后的指定排序项。

encode方法用于将排序规则编码为二进制格式。

entryLess方法比较两个排序项的顺序。

Len方法返回ordering中排序项的数量。

Swap方法交换ordering中两个排序项的位置。

Less方法比较ordering中两个排序项的顺序。

insert方法向ordering中插入一个排序项。

newEntry方法创建一个新的排序项。

find方法在sortedEntries中查找指定位置的排序项。

makeRootOrdering方法创建一个新的根排序规则。

patchForInsert方法更新ordering以适应插入操作。

clone方法复制一个ordering。

front方法返回ordering中第一个排序项。

sort方法对ordering中的排序项进行排序。

genColElems方法生成排序规则的元素列表。

以上是order.go文件中的主要结构体和函数，其主要功能是定义和操作排序规则的数据结构和相关操作，用于实现文本排序的功能。

