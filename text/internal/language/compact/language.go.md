# File: text/internal/language/language.go

在Go的text项目中，`text/internal/language/language.go`文件是一个内部包，定义了语言和地区标签的解析和处理函数。

`errPrivateUse`、`errInvalidArguments`、`errNoTLD`是在该文件中定义的错误变量。`errPrivateUse`表示使用了私有用途的标签，`errInvalidArguments`表示参数无效，`errNoTLD`表示没有顶级域名。

`Tag`是一个代表语言标签的结构体，包含基本子标签和可选的扩展标签和变体标签。

`Variant`是一个代表变体标签的结构体，记录变体的标识符。

`Make`函数用于根据指定的语言和地区信息创建一个Tag对象，并验证其合法性。

`Raw`函数将Tag表示为原始字符串。

`equalTags`函数比较两个Tag对象是否相等。

`IsRoot`函数判断一个Tag是否是根标签。

`IsPrivateUse`函数判断一个Tag是否是私有用途标签。

`RemakeString`函数根据给定的原始字符串重建一个Tag对象。

`genCoreBytes`函数生成一组核心字节。

`String`函数返回Tag对象的字符串表示。

`MarshalText`函数将Tag对象编码为文本格式。

`UnmarshalText`函数将文本格式的Tag解码为Tag对象。

`Variants`函数返回Tag对象的变体标签列表。

`VariantOrPrivateUseTags`函数返回Tag对象的变体标签和私有用途标签列表。

`HasString`函数检查一个Tag对象是否有字符串表示。

`Parent`函数返回Tag对象的父标签。

`ParseExtension`函数解析Tag对象中的扩展标签。

`HasVariants`函数检查Tag对象是否具有变体标签。

`HasExtensions`函数检查Tag对象是否具有扩展标签。

`Extension`函数返回指定键的扩展标签。

`Extensions`函数返回Tag对象的扩展标签列表。

`TypeForKey`函数返回指定键的语言标签。

`SetTypeForKey`函数为指定的键设置语言标签。

`findTypeForKey`函数根据键查找相应的语言标签。

`ParseBase`函数解析Tag对象中的基本子标签。

`ParseScript`函数解析Tag对象中的脚本标签。

`EncodeM49`函数将Tag对象编码为M49格式。

`ParseRegion`函数解析Tag对象中的地区标签。

`IsCountry`函数检查Tag对象是否表示国家。

`IsGroup`函数检查Tag对象是否表示地区群组。

`Contains`函数检查一个Tag对象是否包含另一个Tag对象。

`TLD`函数返回Tag对象的顶级域名标签。

`Canonicalize`函数对Tag对象进行规范化。

`ParseVariant`函数解析Tag对象中的变体标签。

