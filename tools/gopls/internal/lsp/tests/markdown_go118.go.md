# File: tools/gopls/internal/lsp/tests/markdown_go118.go

文件 `markdown_go118.go` 是 `gopls` 工具的一个内部测试文件，用于测试 Markdown 相关的功能。

该文件中包含了两个函数：

1. `DiffMarkdown`：该函数用于比较两个 Markdown 文件的内容差异，并生成一个可读性更高的 Markdown 格式的差异报告。

2. `normalizeMarkdown`：该函数用于规范化 Markdown 文本，使其符合一些规范和风格要求。它会修复一些常见的 Markdown 错误或问题，并对文本进行格式化和标准化。

这些函数的具体作用如下：

1. `DiffMarkdown` 函数通过比较两个 Markdown 文件的内容，找出它们之间的差异并生成差异报告。这个函数主要用于修改或重构 Markdown 文件时，帮助开发者理解文件内容的变化。它可以显示出两个 Markdown 文件之间新增、删除、修改的内容，并以易于阅读的方式展示出来，帮助开发者更好地理解和分析文件的变化。

2. `normalizeMarkdown` 函数用于规范化 Markdown 文本。Markdown 文本的格式可能因为编辑器设置不同、人为错误或其他原因而不一致。这个函数会修复一些常见的 Markdown 错误，并对文本进行格式化和标准化，使其更符合一些约定和规范。它会修复缩进、链接、标题、代码块等方面的问题，并对文本进行一些清理操作，使其在处理和显示时更加统一和易读。

这些函数在 `gopls` 工具的开发中起到了测试和辅助的作用，帮助开发者在进行 Markdown 处理时能更好地理解和调试代码，并确保 Markdown 文本在处理和展示时具有一致的格式和风格。

