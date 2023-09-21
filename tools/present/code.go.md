# File: tools/present/code.go

在Golang的Tools项目中，tools/present/code.go文件的作用是实现对代码的解析和处理，用于在present工具中展示代码。

具体变量的作用如下：

- PlayEnabled：用于确定是否启用代码运行的功能。
- NotesEnabled：用于确定是否启用代码注释的功能。
- highlightRE：用于匹配代码中需要高亮的部分。
- hlCommentRE：用于匹配代码中的注释部分。
- codeRE：用于匹配代码块。
- leadingSpaceRE：用于匹配代码行前的空格。
- codeTemplate：用于存储代码展示的模板。

具体结构体的作用如下：

- Code：用于表示代码块的结构体，包含代码内容和代码类型等信息。
- codeTemplateData：用于表示代码模板的数据结构，存储了代码块的相关信息。
- codeLine：用于表示代码行的结构体，包含代码内容和代码行号等信息。

具体函数的作用如下：

- init：初始化代码模板。
- PresentCmd：用于解析并展示代码，接受代码行和注释行，并根据定义的模板渲染展示。
- TemplateName：获取代码展示所使用的模板名。
- parseCode：解析代码块，将代码分割成行，并提取其中的注释和代码。
- formatLines：格式化代码行，用于去除行首空格等。
- rawCode：获取原始的代码内容。
- codeLines：获取代码的行数。
- parseArgs：解析输入参数，将命令行参数转换为代码块。

