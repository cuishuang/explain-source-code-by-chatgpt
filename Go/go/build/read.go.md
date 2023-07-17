# File: read.go

read.go文件的作用是提供读取操作的工具函数。该文件包含多个函数，其中最重要的是ReadFull和ReadAtLeast。这些函数允许用户读取指定长度的数据，避免了读取不足的问题。

具体来说，ReadFull函数用于读取指定长度的数据并将其存储在提供的缓冲区中。如果读取的数据长度不足，则返回错误。而ReadAtLeast函数则用于读取至少指定长度的数据，如果读取的数据长度不足，则会一直阻塞直到读取指定长度为止。

此外，read.go文件还包含了其他一些与读取操作相关的函数，例如ReadByte、ReadRune、NewReader等。这些函数简化了读取操作，并提供了更高级别的读取操作，例如读取单个字节或Unicode字符等。

总的来说，read.go文件是Go标准库中非常重要的一个文件，提供了读取数据的工具函数，使得从IO读取数据变得更加方便，安全和可靠。




---

### Var:

### bom





### errSyntax





### errNUL





### goEmbed








---

### Structs:

### importReader





## Functions:

### newImportReader





### isIdent





### syntaxError





### readByte





### readByteNoBuf





### peekByte





### nextByte





### findEmbed





### readKeyword





### readIdent





### readString





### readImport





### readComments





### readGoInfo





### parseGoEmbed





