# File: html.go

go/src/html中的html.go文件是Go语言标准库中的一部分，它提供了HTML文本的解析功能。HTML是一种标记语言，它用于在Web页面中定义文本的结构、样式以及元数据信息等。

该文件包含了HTML文本分析器的实现，它可以将HTML文本解析为节点树。节点树是一种用于表示HTML文档结构的数据结构，其中每个节点具有一个标记名称和一组属性值。节点可以包含子节点，因此节点树可以用于表示HTML元素之间的嵌套关系和层次结构。

html.go文件还提供了一些有用的功能，例如，它支持HTML实体的解码，这是一种特殊的语法，用于表示在HTML文本中无法直接表示的字符。此外，它还可以将文本转换为HTML格式，例如，将字符串转换为带有合适标记的HTML文本。

总之，html.go文件提供了Go开发人员一个轻量级且易于使用的HTML解析器，使其能够处理HTML文本，并将其转换为易于操作的数据结构。它是Go语言Web开发的重要工具之一。




---

### Var:

### htmlReplacementTable

htmlReplacementTable是一个map类型的变量，它存储了HTML5中需要做转义处理的字符和对应的字符串。在HTML文档中，有些字符需要进行转义，如"<"要写成"&lt;"，">"要写成"&gt;"，否则会导致解析错误。

htmlReplacementTable的作用是在解析html文档时，将这些需要转义的字符替换成对应的字符串，保证文档能够正确解析。例如，当解析一个包含"<"字符的标签时，在解析过程中会从htmlReplacementTable中获取对应的替换字符串"&lt;"，并使用该字符串代替原有的字符。

总之，htmlReplacementTable在Go语言的html包中有着十分重要的作用，它是html包中用于HTML文档解析的一个重要的替换表。



### htmlNormReplacementTable

htmlNormReplacementTable 是一个 HTML 实体字符规范化替换表，包含了一系列常用的 HTML 实体字符及其对应的 Unicode 字符。

在 HTML 中，一些特殊字符需要使用实体字符的形式来进行表示，例如 < 符号需要使用 &lt; 表示。当解析 HTML 文档时，需要将这些实体字符转换为对应的 Unicode 字符，以便正确显示文本。

htmlNormReplacementTable 变量提供了一系列常用的 HTML 实体字符到 Unicode 字符的映射关系，可以方便的进行字符转换。这个映射关系也可以供其他程序使用，例如用于在 Web 开发中字符转义和解转义。

在实现 HTML 解析器时，htmlNormReplacementTable 可以用于将 HTML 文本中的实体字符转换为对应的 Unicode 字符，以便正确地解析和表示 HTML 文本。



### htmlNospaceReplacementTable

htmlNospaceReplacementTable这个变量是一个字符串映射表，用于在HTML文本中替换一些对应的实体字符（entity）。

HTML中有一些特殊字符（如空格、小于号、大于号等），在HTML文本中必须通过实体字符的形式来表示，否则会被浏览器解释成特殊含义而无法正确显示。例如，空格字符在HTML中需要使用“&nbsp;”实体字符来表示。

htmlNospaceReplacementTable变量中预定义了一些实体字符替换规则，例如将空格字符替换为“&nbsp;”实体字符。当HTML文本中存在这些字符时，会自动进行替换操作，从而保证HTML文本的正确性。

需要注意的是，该变量只在HTML解析器中使用，用户在编写HTML文本时，仍需手动使用对应的实体字符来表示特殊字符。



### htmlNospaceNormReplacementTable

变量htmlNospaceNormReplacementTable是一个包含HTML实体名称和对应实体字符的映射表。它的作用是在解析HTML文档时将实体引用替换为它们对应的字符。

在HTML中，有一些字符（如空格、小于号、大于号、引号等）不能直接作为文本内容输出，而是需要使用实体引用表示。例如，空格需要使用实体引用"&#32;"表示。当解析HTML文档时，需要将这些实体引用替换为对应的字符，以最终呈现正确的文本内容。

htmlNospaceNormReplacementTable变量中包含了一些常见的实体引用的映射表，它可以确保HTML文档中的实体引用被正确地解析和替换为对应的字符。这些字符通常是空格、制表符、换行符等标点和空白字符，用于控制文本的格式。



## Functions:

### htmlNospaceEscaper

htmlNospaceEscaper是一个函数，用于转义HTML文本，但不对空格进行转义。

在HTML中，一些特殊字符（如<、>、&等）需要进行转义，否则会导致浏览器将它们解释为HTML标签或命令。htmlNospaceEscaper函数可以将这些特殊字符转义，并将它们替换为相应的HTML实体。

与其他转义函数不同的是，htmlNospaceEscaper不转义空格。这是因为HTML中的空格通常用于实现格式化，而不是传递有意义的信息。如果对空格进行转义，可能会破坏HTML的格式化效果。

例如，将字符串"<p>Hello, World!</p>"传递给htmlNospaceEscaper函数，它将返回"&lt;p&gt;Hello, World!&lt;/p&gt;"。换句话说，函数将尖括号和引号转义为HTML实体，但保留空格。

总之，htmlNospaceEscaper函数是一个用于转义HTML文本的工具，它可以在保留空格的前提下，将HTML中的特殊字符转义为HTML实体，以确保HTML文本得到正确地解释和显示。



### attrEscaper

attrEscaper() 函数是用于转义HTML属性的字符编码的方法。在HTML文档中，属性值必须使用引号括起来，双引号 (" ") 或单引号 (' ')。如果属性值中包含引号，则必须使用相反类型的引号。此外，所有属性值中的字符必须使用字符编码进行转义，以便它们能够正确显示。这是因为某些字符有特殊的意义，例如小于号 (<) 和大于号 (>) 可能被解释为HTML标记。

在 attrEscaper() 函数中，除了将所有引号替换为正确的对应引号之外，还分情况处理了危险字符，例如小于号 (<)，大于号 (>)，单引号 (')，双引号 (")，和符号 (&)。这些字符必须正确转义以防止注入攻击或破坏文档格式。

因此，attrEscaper() 在HTML编写过程中非常有用，可以让你以安全的方式添加HTML标记，并排除任何潜在的漏洞或格式化问题。



### rcdataEscaper

rcdataEscaper是一个函数，其作用是将文本中的特殊字符进行转义，以便在HTML文档中正确地显示。具体来说，它将“&”符号替换为“&amp;”、“<”符号替换为“&lt;”、“>”符号替换为“&gt;”、“"”符号替换为“&quot;”、“'”符号替换为“&#39;”等。

在HTML文档中，部分文本被称为可重定向数据（RCDATA），如<script>和<style>标签内的内容等。这些标签内的文本需要被解释器特殊处理以避免解释器错误解释为HTML标签。因此，在HTML解析器中，需要将这些特殊字符进行转义，以避免这些字符被错误解释。

rcdataEscaper是一个重要的函数，因为它确保了HTML文档中的RCDATA部分中的文本正确地解释为纯文本，从而使文档能够正确地显示。



### htmlEscaper

htmlEscaper是一个函数，用于将输入字符串转义为HTML形式。

它的作用是防止在HTML文档中出现非法字符导致文档结构错误或安全问题。例如，如果未将字符串中的<符号转义，则可能在HTML文档中创建一个标签，而不是想要的字符串一部分。htmlEscaper的任务是将这些字符转义为HTML实体。

例如，下面是一个字符串：

```
This is a <strong>bold</strong> statement.
```

如果该字符串被用于生成HTML文档，那么<strong>将被解释为HTML标签。为了避免这种情况，htmlEscaper被用于将其转换为：

```
This is a &lt;strong&gt;bold&lt;/strong&gt; statement.
```

这样，它将被渲染为预期的字符串，而不是HTML标签。

htmlEscaper是将输入字符串中的字符检查并转义为HTML实体。它将以下字符转义为相应的HTML实体：

```
<   -->   &lt;
>   -->   &gt;
&   -->   &amp;
'   -->   &#39;
"   -->   &quot;
```

这些实体被处理为原始的HTML字符，而不是它们在HTML中的字面意义。htmlEscaper函数确保生成的HTML文档是有效的，并能正确显示所有字符。



### htmlReplacer

htmlReplacer是一个函数，它的作用是在HTML文本中替换一些特殊字符，使得这些字符能够正确地显示在浏览器中。

具体来说，htmlReplacer会依次查找HTML文本中的特殊字符，如"&"、"<"、">"、"'"和"""，并将其替换成它们在HTML中对应的转义字符，如"&amp;"、"&lt;"、"&gt;"、"&apos;"和"&quot;"。这样，浏览器在解析HTML文本时就能够正确地显示这些特殊字符，而不会将它们误认为是HTML标记。

htmlReplacer的实现基于Go语言的字符串替换工具strings包中的Replace函数，不过在这里它使用了一个名为htmlReplacerTable的map，它将每个特殊字符映射到其对应的转义字符。这个方法具有高效、简便的特点，可以让Go语言轻松地处理HTML文本中的特殊字符。



### stripTags

stripTags函数是将字符串中的HTML标签删除并返回结果的函数。它接受一个字符串作为参数，并使用正则表达式删除所有HTML标签，然后返回处理后的结果。

该函数的作用是将HTML文本转换为纯文本，以便更容易地对文本进行处理和显示。在处理web应用程序中的用户输入时，该函数非常有用，因为它可以防止恶意用户利用HTML标签来执行攻击，例如跨站脚本攻击（XSS）。

stripTags函数的实现方式是使用正则表达式匹配所有HTML标签，并将它们删除。它还使用一些特殊的正则表达式来处理不同类型的标签和属性。

例如，该函数可以处理以下类型的标签：

- div、span、p等块状元素标签
- a、img等内联元素标签
- b、strong等文本格式化标记

另外，该函数还可以处理一些特殊的标记和属性，例如注释、JavaScript脚本、CSS样式等。



### htmlNameFilter

htmlNameFilter是一个函数，它的作用是过滤HTML标签中的属性名（attribute name）。在解析HTML文档过程中，出现的标签可能带有多个属性，例如，<a href="https://www.example.com" target="_blank">，其中href和target就是标签的属性名。从安全角度考虑，为了防止一些恶意攻击，通常需要对属性名进行过滤。htmlNameFilter就是用来做这个事情的。

函数内部实现是通过正则表达式进行匹配。如果属性名符合规范，则直接返回原始字符串；否则，将属性名替换为&zwnj;，这是零宽非连接符，可以避免在浏览器中出现解析错误或被误解为其它的字符。

总的来说，htmlNameFilter这个函数是用来保证代码安全性的一个重要环节，它保证了从HTML文档中读取属性名是安全的。



### commentEscaper

在html.go文件中，commentEscaper是一个功能强大的函数，被用于HTML文档中注释的转义。

HTML中的注释采用“<!--”和“-->”标记。通常情况下，如果用户在注释中输入“-->”字符，代码将会认为注释结束，从而产生错误。为了避免这种情况，函数commentEscaper会将注释中的“--”序列替换成“- -”（中间有空格），从而避免“-->”这个字符出现在注释中。

具体来说，该函数检查注释的文本节点内容，对于以“-”作为开头的文本内容，在其前插入一个空格。对于以“-”作为结尾的文本节点内容，在其后插入一个空格。这样，注释内容中的“--”字符序列就会转换成“- -”，从而避免后续代码的解析错误。

总结一下，commentEscaper函数的主要功能是确保HTML文档中的注释不会包含“-->”字符序列，从而避免产生语法错误。



