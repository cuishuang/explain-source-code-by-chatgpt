# File: tools/gopls/internal/lsp/protocol/mapper.go

tools/gopls/internal/lsp/protocol/mapper.go文件是Golang中gopls工具项目中的一个文件，其作用是提供一系列用于位置映射和转换的函数和结构体。

Mapper结构体用于在源代码和LSP中的位置之间进行映射。它包含了源代码和LSP之间的行列映射关系。MappedRange结构体表示源代码中的一个范围，并提供了用于获取LSP中位置和范围的函数。

下面是Mapper文件中的一些主要函数和结构体的作用：

- NewMapper: 创建一个新的Mapper实例。

- initLines: 初始化Mapper的行映射关系。它会根据源代码中的换行符来确定每行的起始和结束位置。

- SpanLocation: 将LSP中的Span（起始和结束的位置）转换为源代码中的行列位置。

- SpanRange: 将LSP中的Span转换为源代码中的范围。

- PointPosition: 将LSP中的Point（一个位置）转换为源代码中的行列位置。

- OffsetLocation: 将在源代码中的偏移量转换为行列位置。

- OffsetRange: 将在源代码中的偏移量范围转换为行列范围。

- OffsetSpan: 将在源代码中的偏移量转换为LSP中的Span。

- OffsetPosition: 将在源代码中的偏移量转换为LSP中的Point。

- lineCol16: 将一个16位整数解码为行列位置。

- lineCol8: 将一个8位整数解码为行列位置。

- line: 获取指定行的源代码内容。

- OffsetPoint: 将LSP中的Point转换为在源代码中的偏移量。

- OffsetMappedRange: 将在源代码中的偏移量范围转换为MappedRange。

- LocationSpan: 将LSP中的Location转换为Span。

- RangeSpan: 将LSP中的Range转换为Span。

- RangeOffsets: 将LSP中的Range转换为在源代码中的偏移量范围。

- PositionOffset: 将LSP中的行列位置转换为在源代码中的偏移量。

- PositionPoint: 将LSP中的行列位置转换为Point。

- PosPosition: 将源代码中的行列位置转换为LSP中的行列位置。

- PosLocation: 将源代码中的行列位置转换为LSP中的Location。

- PosRange: 将源代码中的范围转换为LSP中的Range。

- NodeRange: 将源代码中的AST节点转换为LSP中的Range。

- RangeLocation: 将LSP中的Range转换为Location。

- PosMappedRange: 将源代码中的范围转换为MappedRange。

- NodeMappedRange: 将源代码中的AST节点转换为MappedRange。

- Offsets: 获取源代码中的所有偏移量。

- URI: 对URI进行解析和格式化。

- Range、Location、Span、String：这些结构体表示位置、范围、字符串等，在位置映射和转换中用于数据的传输和处理。

- LocationTextDocumentPositionParams：一个结构体，表示一次请求中的位置参数。

总之，mapper.go文件中的这些函数和结构体提供了一种在源代码和LSP之间进行位置映射和转换的方式，方便了gopls工具处理源代码和LSP之间的数据交换。

