# File: tools/present/html.go

在Golang的Tools项目中，文件`html.go`位于`tools/present`目录下，主要负责处理与HTML相关的逻辑。

该文件中定义了一些HTML相关的结构体和函数，包括：

1. `HTML`结构体：表示生成的HTML页面的内容，包含标题、CSS和页面主体。
2. `HTMLSlide`结构体：表示一个幻灯片页面的内容，包含标题、索引、片段和注释。
3. `HTMLFooter`结构体：表示页面脚部的内容，包含版权信息和链接。
4. `init`函数：为HTML相关的结构体初始化默认值，如设置CSS样式、尾部链接等。
5. `parseHTML`函数：解析输入的Markdown文件，并将其转换为HTML格式。
6. `PresentCmd`函数：用于执行`present`命令，生成HTML页面文件。
7. `TemplateName`函数：根据文件名和目录生成模板名称，用于加载HTML模板。

这些结构体和函数的作用如下：

1. `HTML`结构体：存储生成的HTML页面的内容，包括页面标题、CSS样式和页面主体。
2. `HTMLSlide`结构体：存储一个幻灯片页的内容，包括标题、索引、片段和注释。
3. `HTMLFooter`结构体：存储页面脚部的内容，包括版权信息和链接。
4. `init`函数：初始化HTML相关结构体的默认值，例如设置CSS样式和尾部链接。
5. `parseHTML`函数：将输入的Markdown文件解析为HTML格式，并生成相应的幻灯片页面。
6. `PresentCmd`函数：执行`present`命令，将解析后的幻灯片页面生成为HTML页面文件。
7. `TemplateName`函数：根据文件名和目录生成用于加载HTML模板的模板名称。

这些结构体和函数的主要目的是解析Markdown文件并生成对应的HTML页面，其中包括幻灯片页、页面标题、脚部等内容。

