# File: entity.go

entity.go是Golang标准库中html包的一部分。它主要用于处理HTML实体，即HTML中特殊字符的编码方式。

HTML实体是一些特殊字符，如小于号（<），大于号（>），引号（"），单引号（'）等字符在HTML中的替代编码方式，以避免这些字符与HTML标记混淆，导致页面布局混乱或安全漏洞。

entity.go中的代码提供了一组函数，用于将HTML实体编码或解码。它们分别是：

- EscapeString：将字符串中的特殊字符转换为该字符的HTML实体编码。例如，小于号转换为&lt;。
- UnescapeString：将字符串中的HTML实体解码为原来的字符。例如，&lt;解码为小于号。
- Escape：将字符c的HTML实体编码返回。例如，字符<返回&lt;。
- Unescape：将HTML实体s解码为原来的字符并返回。例如，字符串&lt;解码为小于号。

这些函数使得在HTML中处理特殊字符变得容易，从而保证HTML页面的正常显示和安全性。




---

### Var:

### entity

entity变量是一个映射，它将常用的HTML实体名称映射到它们对应的Unicode字符。在HTML中，有一些特殊字符，如小于号 < 和大于号 >，它们需要使用实体名称或实体编号来表示，以便正确的被渲染。entity变量提供了一种方便的方式来查找这些特殊字符实体名称对应的Unicode字符。例如，entity["lt"]将返回Unicode字符“<”。

在html中entity.go文件中，通过在entity变量中存储实体名称和对应的Unicode字符，可以在分析和渲染HTML时方便地引用这些字符。entity变量是一个map类型，其中key是实体名称，value是对应的Unicode字符。当解析HTML字符串时，可以通过使用entity变量来处理这些特殊字符，以便正确的渲染HTML页面。

例如，当在HTML页面中使用“<”时，应该使用实体名称“&lt;”来替代，“&lt;”将被视为一个特殊字符并被正确解析。entity变量包含了很多这样的映射，以便正确处理HTML中所有的特殊字符。



### entity2

entity2变量是HTML中预定义实体字符的映射表，用于在HTML文本中转义特殊字符，如“<”、“>”、“&”、“"”和“'”。

entity2变量是map[string]string类型，其中键是实体字符的名称，值是它们对应的Unicode编码。例如，键“lt”对应的值是“＜”。

通过使用entity2变量，可以将HTML文本中的预定义字符转换为它们的Unicode编码。这样，可以避免这些字符对HTML的解析造成干扰。

在编写HTML解析器时，可以使用entity2变量来处理HTML文本。例如，可以将下面的HTML片段：

    <p>This is a "test" & <i>example</i></p>

使用entity2变量进行转义，得到如下结果：

    &lt;p&gt;This is a &quot;test&quot; &amp; &lt;i&gt;example&lt;/i&gt;&lt;/p&gt;

这样，在解析HTML时，就可以正确处理特殊字符，而不会影响HTML的结构和语义。



### populateMapsOnce

`populateMapsOnce`是一个常量，表示HTML包中实体映射的初始化已经完成。它的作用是防止重复初始化实体映射，从而提高代码性能。在HTML包中，实体映射是一种将常见HTML实体名称映射到它们对应的Unicode代码点的方法。如果未初始化实体映射，解析HTML文档中包含的实体名称将变得非常困难。

在`entity.go`文件中，`populateMapsOnce`常量定义如下：

```
var populateMapsOnce sync.Once
```

在实体映射初始化时，`populateMapsOnce`被用于确保实体映射仅被初始化一次。`sync.Once`是一个同步原语，可确保代码只执行一次，即使多个goroutine同时调用该方法。一旦实体映射完成初始化，将不再需要对其进行初始化，因为实体映射已以映射表的形式存储在内存中。通过使用`populateMapsOnce`常量，无需重复执行实体映射的初始化代码，提高了代码性能。



## Functions:

### populateMaps

在 HTML 中，有一些特殊的字符，例如小于号（<）、大于号（>）和和号（&）等等，这些特殊字符不能直接在 HTML 中使用，必须使用字符实体来进行替换。

populateMaps 这个函数的作用是将字符实体的名称和对应的码点（Unicode 编码）存储到两个 map 中：entityNameToCodePoint 和 codePointToEntityName。这样，在解析 HTML 文件时，如果遇到字符实体，就可以用这两个 map 进行编码和解码。

具体来说，假设遇到了 `&lt;`，就可以通过 entityNameToCodePoint 中的 `lt` 找到对应的码点 `U+003C`，然后将其转换成 `<`。

这个函数非常重要，因为在 HTML 的解析器中，需要频繁地进行字符实体的编码和解码，而这个函数确保了这些操作可以在 O(1) 的时间内完成，从而提高了解析效率。



