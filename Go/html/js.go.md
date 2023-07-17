# File: js.go

js.go是Go语言实现HTML解析器的一部分，它主要用于处理HTML中的JavaScript代码。

具体来说，js.go文件中定义了一个jsScanner类型和相应的方法，用于扫描HTML文本中的JavaScript代码，并将其解析为Token（符号）序列。然后，这些Token序列可以被送到jsParser（另一个类型和相关的方法），进一步解析为AST（抽象语法树）。

通过解析AST，js.go可以帮助Go程序将HTML页面中的JavaScript代码转换为可执行的代码，从而实现对HTML中嵌入的JavaScript代码的处理和运行。

总之，js.go作为Go语言HTML解析器的一部分，是用于处理HTML中JavaScript代码的重要组成部分，它使得Go程序能够轻松地处理和运行HTML页面中的JavaScript代码。




---

### Var:

### regexpPrecederKeywords

在 Go 语言中的 html 包中，js.go 文件中定义了一个名为 regexpPrecederKeywords 的变量。这个变量是一个字符串切片，包含了一些 JavaScript 关键词，比如：`return`、`function`、`var`、`if` 等等。

这个变量的作用是在解析 JavaScript 代码时，判断当前位置是否应该考虑解析正则表达式。在 JavaScript 中，正则表达式可以使用 `/` 开头和结尾的字面量表示，也可以使用 `RegExp` 对象进行构造。但是，在一些情况下，对于 `/` 开头的表达式，可能会产生歧义。比如，以下代码：

```
var a = 10 / 2;
```

这个代码中，`/` 既可以是除法，也可以是正则表达式的开始。因此，我们需要一些方法来判断当前位置是否应该解析为正则表达式。

`regexpPrecederKeywords` 变量就是用来解决这个问题的。在解析 JavaScript 代码时，如果当前位置的前面有一个元素是 `regexpPrecederKeywords` 中的关键词之一，那么就认为当前位置应该被解析为正则表达式。如果不是，则认为当前位置是除法运算。

需要注意的是，这个变量只是用来简化代码的一个辅助变量，并不是一个正则表达式语法规则。JavaScript 中正则表达式的语法规则比较复杂，需要仔细理解。



### jsonMarshalType

在go/src/html中js.go这个文件中，jsonMarshalType变量用于标识要进行JSON编码的结构体的类型，它的类型是一个映射表，键是类型名称字符串，值是一个与该类型对应的JSON编码函数。通过这个映射表，我们可以方便地将Go语言中的结构体类型转换成JSON格式的数据。

为了方便与其他编程语言的交互，Go语言内置了JSON编解码函数，可以很方便地将一个Go语言对象（比如结构体）序列化成JSON格式的数据，或者将JSON格式的数据反序列化成Go语言对象。在实现这些函数时，需要了解每个Go语言类型与JSON数据类型的对应关系。例如，Go语言中的整型和浮点型对应着JSON的数值类型，字符串类型对应JSON的字符串类型，而Go语言中的结构体则需要根据其具体的定义方式进行转换。jsonMarshalType就是在定义结构体类型和JSON编码函数之间建立对应关系的一个映射表。在调用JSON编码函数时，首先会查找该结构体类型在jsonMarshalType中是否存在对应的JSON编码函数，如果存在则调用该函数进行编码，否则报错。

总之，jsonMarshalType变量在Go语言的JSON编解码过程中起着重要的作用，为结构体类型与JSON数据类型之间的转换提供了便利和支持。



### lowUnicodeReplacementTable

在 go/src/html/js.go 文件中，lowUnicodeReplacementTable 是一个用于编码 Unicode 字符的变量，它在 HTML/JS 中使用的 Unicode 编码方式中起着关键作用。 

在 HTML/JS 中，可以使用 Unicode 编码来表示非 ASCII 字符。 例如，字符“é”可以用 Unicode 编码 \u00E9 来表示。 对于 Unicode 码点小于 0x100 的字符，可以使用直接解码，但是對於大於0x100的字符则需要进行特殊处理。为了支持这样的编码，JS 包含了一些特殊的转义序列，并将 Unicode 编码的字符对应地转换为它们的 UTF-16 编码，这样在 HTML/JS 中使用时就可以正常显示出来。

在 JS 中，当一个字符的 Unicode 码点大于 0xFFFF 时，不能直接表示为两个 16 位码元，因为 JS 使用16 位编码，并且不支持 surrogate pairs（代理项），而这些 Unicode 码点需要使用 surrogate pairs 进行表示。为了解决此问题，JS 引入了 low surrogate 和 high surrogate 的概念，通过这种方式将 Unicode 编码的字符分解为两个 16 位编码，然后通过这些编码来表示 Unicode 字符。

lowUnicodeReplacementTable 存储了 high surrogate 和 low surrogate 之间的 Unicode 编码范围。 在处理 Unicode 码点时，如果发现通过 surrogate pairs 表示的 Unicode 码点在 lowUnicodeReplacementTable 内，则返回替代字符 “\uFFFD” 来代替该字符，否则将该字符直接解码。

因此，lowUnicodeReplacementTable 在 go/src/html/js.go 文件中使用了一个映射表，来实现将解码无法表示的 Unicode 字符时替换为替代字符的功能，以保证在 HTML/JS 中使用此类字符时能够正常显示。



### jsStrReplacementTable

jsStrReplacementTable 变量是在 Go 语言的 html 包中的 js.go 文件中定义的一个映射表，用于替换 JavaScript 代码中的字面字符串常量。这个变量是一个类型为 map[string]string 的 map 类型，它定义了一对一的映射关系，其中键是需要被替换的字符串，值是替换后的字符串。

该变量的作用在于，在转换 HTML 文档中的 JavaScript 代码时，通过查找 JavaScript 代码中的字符串常量并使用键值对中对应的值进行替换，增强了 HTML 中 JavaScript 代码的嵌入能力，使得 HTML 内的 JavaScript 代码更加规范，易于维护。

例如，如果 JavaScript 代码中包含字符串 "Hello, world!"，它将被替换为 "$html_Hello_comma__world_$"。这个替换后的字符串可以被进一步处理，以避免 HTML 中的特殊字符，防止代码注入攻击等。具体来说，$ 符号表示了一个控制字符序列的起始位置，接下来的 html 表示了序列的标识，这使得我们可以使用 HTML 的转义字符来表示 JavaScript 字符串中的特殊字符。例如，"comma" 在这里表示逗号字符。使用这些特殊字符来表示原有的字符串常量，有助于提高代码的安全性和可维护性，避免了各种处理HTML中包含的JavaScript时的问题。



### jsStrNormReplacementTable

jsStrNormReplacementTable是一个字符串替换表，用于将HTML文档中的JavaScript代码中的转义字符替换为它们对应的Unicode字符。该变量包含一组键值对，其中键表示HTML中的转义字符，值表示对应的Unicode字符。例如，键“\t”表示制表符，值“\u0009”表示该字符的Unicode码点。当解析HTML文档时，如果遇到一个转义字符，就可以使用jsStrNormReplacementTable中的对应值来代替它。这样可以保证JavaScript代码中的字符串被正确地解析和执行。



### jsRegexpReplacementTable

在go/src/html/js.go文件中，jsRegexpReplacementTable是一个字符串替换表，用于在转义javascript代码中将某些特殊字符替换为它们的转义序列。

具体来说，JavaScript代码中有一些特殊字符，如单引号、双引号、反斜线、换行符等，如果直接写在字符串里，可能会导致解析错误。为了避免这种情况，需要对这些特殊字符进行转义，即将其替换为特定的转义序列，例如将单引号替换为 \ '，将双引号替换为 \ "，将反斜线替换为 \ \ 等等。

jsRegexpReplacementTable就是用于实现这种替换的字符串替换表。该表将一些特殊字符和它们的转义序列一一对应，当需要转义字符串时，可以使用这个表进行替换。例如，当需要将双引号替换为 \ "时，可以使用下面的代码：

str = strings.ReplaceAll(str, "\"", jsRegexpReplacementTable["\""])

这样就能把字符串中的双引号替换为 \ "了。

总之，jsRegexpReplacementTable是一个用于转义javascript代码中特殊字符的字符串替换表，它能够帮助开发人员更方便地进行字符串转义操作，确保javascript代码的正确性和安全性。



## Functions:

### nextJSCtx

在 go/src/html/js.go 文件中，nextJSCtx 函数用于提取 HTML 中的 JavaScript 代码。

该函数首先会查找 HTML 中下一个 script 标记，并返回其内容。如果 script 标记带有 type 属性，且不是 "text/javascript" 或 "application/javascript"，则会跳过该标记，查找下一个 script 标记。

在提取完 JavaScript 代码后，该函数还会检查代码是否符合 ECMA-262 标准（即 JavaScript 标准）。如果不符合标准，则会返回一个 error。

需要注意的是，如果 HTML 中没有 script 标记，该函数会返回 EOF 错误。此外，调用者需要保证该函数在正确的文本上下文中调用。如果调用者在查找 script 标记之外的文本上下文中调用该函数，可能会导致错误的结果。



### indirectToJSONMarshaler

indirectToJSONMarshaler是html/js包中的函数之一，用于将一个值解码为JSON对象并返回字节切片，具体作用如下：

在JavaScript中，有一些内置对象实现了MarshalJSON()方法来将其编码为JSON格式，例如Date对象和Map对象等。当我们需要将这些对象编码为JSON并在Go代码中处理时，通常无法直接调用MarshalJSON()方法。这时，我们可以使用indirectToJSONMarshaler函数来解决这个问题。

indirectToJSONMarshaler函数接收一个值作为参数，并检查该值是否实现了MarshalJSON()方法。如果实现了，则调用该方法将值转换为JSON格式的字节切片。如果未实现该方法，则将值转换为其默认的JSON格式。

此外，该函数还支持嵌套调用，即在对象中包含实现了MarshalJSON()方法的对象时，可以递归地调用该方法来编码所有子对象。

总之，indirectToJSONMarshaler函数是一个方便的工具，用于在Go和JavaScript之间编解码JSON对象时转换内置对象的格式。



### jsValEscaper

jsValEscaper 函数的作用是将HTML文档中的JavaScript值进行转义，以便在JavaScript中使用。转义的过程包括对特殊字符如双引号、单引号和反斜杠进行转义，以及对Unicode字符进行编码。

具体来说，该函数会遍历字符串中的每个字符，如果遇到需要转义的字符则进行相应的转义操作，最终返回一个包含转义后的字符串的副本。

例如，输入字符串 `"<script>alert('hello');<\/script>"`，会被转义为 `"\\u003cscript\\u003ealert(\\u0027hello\\u0027);\\u003c\\/script\\u003e"`。这样转义后，值可以在JavaScript中使用，而不会被解释成HTML标签或其他非法字符。



### jsStrEscaper

jsStrEscaper这个函数主要用于在HTML文档中对JavaScript字符串进行转义操作，以防止XSS攻击。XSS（Cross-Site Scripting，跨站脚本攻击）是一种Web安全漏洞，攻击者通过向页面注入恶意代码来窃取用户的敏感信息或者欺骗用户执行某些特定的操作。

在JavaScript中，字符串值需要用引号将其包围起来。如果一个字符串值本身含有引号，则需要进行转义。例如，有如下字符串：

```javascript
var name = "Tom \"the cat\"";
```

字符串中的双引号需要转义，否则会导致语法错误。这个函数就是用来对这种字符进行转义，将其转义成HTML实体，使其在HTML中正常显示，并且不会被当做JavaScript代码解析。

在jsStrEscaper函数中，使用了一个map来保存需要转义的字符，通过遍历字符串，找到需要转义的字符，然后使用map中定义好的映射进行转义。例如，双引号需要转义成“\&quot;”，单引号需要转义成“\&apos;”。

这个函数的具体实现可以参考go/src/html/js.go文件中的代码。



### jsRegexpEscaper

jsRegexpEscaper这个函数是用于转义JavaScript中正则表达式的特殊字符的。在JavaScript中，有一些字符是有特殊意义的，例如星号、加号、问号、括号等，如果不进行转义，这些字符会被解释为正则表达式的控制字符，导致意料之外的匹配结果。

jsRegexpEscaper函数接收一个字符串参数，它将该字符串中的正则表达式特殊字符进行转义，返回一个转义后的字符串。实现方式是使用正则表达式匹配需要转义的字符，并在其前面添加反斜杠进行转义。

这个函数主要用于HTML模板中的JavaScript代码的转义，以避免XSS（跨站脚本攻击）和其他安全问题。当用户输入一些恶意代码时，这些代码可能会被浏览器解释并执行，给网站带来安全威胁。jsRegexpEscaper函数可以将用户输入的代码进行转义，这样就能够消除代码注入的安全问题。



### replace

在go/src/html/js.go文件中，replace函数用于替换HTML文本中的特定字符串。

具体来说，JSValue类型的replace函数通过接收两个参数，一个正则表达式和一个字符串，用来匹配文本中符合正则表达式的字符串并将其替换为指定的字符串。该功能等效于JavaScript字符串对象的replace()方法。

在HTML文档中，有时需要使用JavaScript代码来操纵文本内容。由于JavaScript代码和HTML标签使用相同的符号，因此在处理JS代码和HTML标记时可能会出现冲突。replace函数可以用于解决这个问题，将HTML中的特殊字符（如<、>、&等）替换为对应的转义字符，以便在文档中正常显示。

例如，下面的代码使用replace函数将HTML字符串中的特殊字符替换为转义字符：

```
s := "<script>alert('Hello World!');</script>"
escaped := js.Global().Get("encodeURIComponent").Invoke(s).String()
document.Call("write", "Escaped: "+escaped+"<br>")
fixed := js.Global().Get("decodeURIComponent").Invoke(escaped).String()
document.Call("write", "Fixed: "+fixed+"<br>")
```

在上面的代码中，首先将包含JavaScript代码的HTML字符串定义为s变量。然后，使用encodeURIComponent函数对s进行编码，并将结果赋值给escaped变量。使用replace函数处理s的主要目的是确保HTML中的所有特殊字符都被正确地编码。最后，使用decodeURIComponent函数对escaped进行解码，并将结果赋给fixed变量。通过这样的操作，我们可以将包含JS代码的HTML字符串正确地渲染为Web页面。



### isJSIdentPart

isJSIdentPart函数的作用是检查给定的代码字符是否是JavaScript标识符的一部分，如果是，则返回true，否则返回false。由于JavaScript标识符可以包含字符、数字或下划线，因此该函数检查了字符串是否合法。

isJSIdentPart函数还考虑了Unicode字符集和ECMAScript标准，这意味着它可以正确地处理所有的JavaScript代码字符，包括Unicode字符、转义序列等。它还考虑了一些JavaScript特定的字符，如$、@、#等符号。

isJSIdentPart函数可用于HTML中解析JavaScript代码时的标识符识别。在HTML代码中，JavaScript代码可能出现在标签的属性中，如onclick或onload。在这种情况下，HTML解析器将JavaScript代码作为文本处理，并传递给JavaScript解析器。解析器需要检查给定的文本是否包含有效的标识符，以便正确识别其含义。

总之，isJSIdentPart函数是一个非常重要的JavaScript字符检查函数，它用于确定JavaScript标识符是否合法。它在HTML解析器和JavaScript解析器之间提供了有效的桥梁，确保JavaScript代码的正确解析和执行。



### isJSType

isJSType是html中js.go文件中的一个函数，它的作用是判断一个字符串是否为JavaScript的数据类型。

该函数的具体实现如下：

```
// isJSType reports whether typ represents a JavaScript type.
func isJSType(typ string) bool {
    switch typ {
    case "bool", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64",
        "float32", "float64", "complex64", "complex128", "string":
        return true
    default:
        return false
    }
}
```

该函数首先使用一个switch语句检查参数typ，如果符合JavaScript支持的数据类型，则返回true；否则返回false。

这个函数的作用是方便判断参数类型是否为JavaScript支持的类型。在html中，有一些需要判断参数类型的地方，比如在解析script标签内容时需要知道其所属的数据类型，该函数可以提供较为简单的方式来进行判断，并且保证了判断结果的准确性。



