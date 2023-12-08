# File: text/cmd/gotext/rewrite.go

在Go的text项目中，text/cmd/gotext/rewrite.go文件的作用是实现重写规则的解析和应用。

该文件中定义了一个名为cmdRewrite的结构体，该结构体用于表示重写规则。cmdRewrite结构体的字段包括Pattern（匹配模式）、Source（源字符串）、Dest（目标字符串）、PreserveCase（是否保留原始字符串的大小写）、Regex（是否使用正则表达式匹配）等。cmdRewrite结构体的定义如下：

```
type cmdRewrite struct {
    Pattern       string
    Source        string
    Dest          string
    PreserveCase  bool
    Regex         bool
    compiledRegex *regexp.Regexp
}
```

其中，Pattern字段表示匹配模式，可以是普通字符串或正则表达式。Source字段表示待替换的字符串，Dest字段表示替换后的字符串。PreserveCase字段决定是否保留原始字符串的大小写，Regex字段表示是否使用正则表达式进行匹配。compiledRegex字段存储已编译的正则表达式。

该文件中还定义了以下函数：

- initRewrite：用于初始化重写规则。该函数会解析重写规则文件，根据文件中的规则创建cmdRewrite结构体的实例，并将这些实例存储在全局变量rewriteRules中。重写规则文件的默认路径为~/.gotext_rewrite。
- runRewrite：用于应用重写规则。该函数会遍历rewriteRules中的每个重写规则实例，并对给定的字符串应用这些规则。如果匹配到了某个规则的Pattern，就会执行对应的替换规则，将字符串中的Source替换为Dest。如果PreserveCase字段为真，则保持原始字符串的大小写。

总之，rewrite.go文件实现了重写规则的解析和应用功能，允许用户定义自定义的替换规则，并在需要时对字符串应用这些规则。

