# File: tools/gopls/doc/generate.go

在Golang的Tools项目中，tools/gopls/doc/generate.go文件的作用是用于生成Gopls的文档。它通过解析Gopls的API定义文件，并生成对应的Markdown格式的文档。

optionsGroup结构体是用来表示一组选项的集合。它包括了组名、组的简要描述、选项列表等信息，用于在文档中显示每个选项的详细说明。

下面是各个函数的作用解释：

- main函数是程序的入口，负责解析命令行参数，并调用doMain函数进行文档生成。
- doMain函数是主要的文档生成函数，它负责加载Gopls的API定义文件，并调用各个子函数生成文档的不同部分。
- pkgDir函数用于获取Gopls包的位置。
- loadAPI函数用于加载Gopls的API定义文件，并返回对应的数据结构。
- loadOptions函数通过解析API定义文件获取所有的选项，并构建对应的optionsGroup结构体。
- loadEnums函数通过解析API定义文件获取所有的枚举类型，并返回一个映射表。
- collectEnumKeys函数用于收集枚举类型的键名。
- formatDefaultFromEnumBoolMap函数用于格式化带有枚举类型的布尔选项的默认值。
- formatDefault函数用于格式化选项的默认值。
- valueDoc函数用于生成选项的值文档。
- loadCommands函数用于解析API定义文件并获取所有的命令。
- argsDoc函数用于生成命令的参数文档。
- typeDoc函数用于生成类型的文档。
- structDoc函数用于生成结构体的文档。
- loadLenses函数用于解析API定义文件并获取所有的代码镜头（lenses）。
- loadAnalyzers函数用于解析API定义文件并获取所有的代码分析器。
- loadHints函数用于解析API定义文件并获取所有的代码提示（hints）。
- lowerFirst函数用于将字符串的首字母转换为小写。
- upperFirst函数用于将字符串的首字母转换为大写。
- fileForPos函数用于根据位置获取对应的文件。
- rewriteFile函数用于重写文件。
- rewriteAPI函数用于重写API文档。
- rewriteSettings函数用于重写设置文档。
- collectGroups函数用于收集选项的组信息。
- hardcodedEnumKeys函数用于获取硬编码的枚举示例的键名。
- writeBullet函数用于生成Markdown格式的项目符号。
- writeTitle函数用于生成Markdown格式的标题。
- capitalize函数用于将字符串的首字母大写。
- strMultiply函数用于生成重复多次的字符串。
- rewriteCommands函数用于重写命令文档。
- rewriteAnalyzers函数用于重写代码分析器文档。
- rewriteInlayHints函数用于重写代码提示文档。
- replaceSection函数用于替换Markdown文档中的指定部分。

