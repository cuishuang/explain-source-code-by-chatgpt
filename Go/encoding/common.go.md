# File: common.go

common.go文件位于Go语言标准库中encoding包中，是一个公共文件，提供了一些编码和解码所需的通用方法和变量。它主要负责以下几个功能：

1. 定义并实现了一些常见的编解码需要的函数。如IsSpace(r rune) bool、readFull(r io.Reader, buf []byte) (n int, err error)、writeFull(w io.Writer, buf []byte) (n int, err error)等。

2. 定义了一些常量。如UTF8、UTF16LE、UTF16BE、UTF32LE、UTF32BE等常量是用于标识Unicode字符集编码方式的常量。

3. 实现了一些接口。如编码解码器的接口，包括Decoder、Encoder、Marshaler和Unmarshaller接口。

4. 提供了一些针对字符编码的工具方法。如NewDecoder(charset string, r io.Reader) io.Reader、NewEncoder(charset string, w io.Writer) io.Writer等。

总的来说，common.go文件的功能是为encoding包中的其他文件提供通用的函数和变量，方便对不同的字符编码进行编解码操作。它是整个encoding包的基础，为其他相关文件的运行提供了支持。




---

### Structs:

### tagAndLength

tagAndLength是一个结构体类型，用于表示一组数据的标记和长度。具体来说，这个结构体用于实现ASN.1编码中的TLV（Type-Length-Value）格式，即类型、长度和值的编码方式。

该结构体包含两个字段：tag和length，分别表示数据的标记和长度。tag字段通常用来表示数据的类型，它在ASN.1编码中是一个数值，用于唯一标识不同类型的数据。length字段表示数据的长度，它在ASN.1编码中可以是固定长度或者可变长度，具体取决于数据的类型和编码规则。

在编解码过程中，tagAndLength结构体通常与其他数据结构（如byte数组或者其他结构体）组合使用，用于将数据按照TLV格式进行编码或解码。例如，在ASN.1编码中，一个完整的TLV数据由一个tagAndLength结构体和一个值（即数据内容）组成。编码时，需要将tag、length和值按照一定的规则转换成一个二进制流，解码时，则需要对二进制流进行逆向操作，即将二进制流解析成tagAndLength结构体和数据值。

总之，tagAndLength结构体在ASN.1编码中扮演着重要的角色，用于确保数据的正确编解码。



### fieldParameters

在go的encoding包中，common.go文件中的fieldParameters结构体用于表示一个编码字段的参数，它包含以下字段：

- tag: 字段的tag字符串，用于标识字段在编码和解码时的处理方式。
- index: 字段在结构体中的索引。
- omitEmpty: 是否忽略空值字段。
- quoted: 是否对字段进行引用包裹。

这些参数会在编码和解码结构体时起到非常重要的作用，例如通过tag可以设置编码和解码时使用的字段名称或者值的格式化方式。omitEmpty则可以控制是否忽略掉值为空的字段，从而减少编码数据的大小。而quoted则可以在编码时对某些特殊字符进行替换，避免在传输过程中出现问题。

总之，fieldParameters结构体可以帮助我们更加灵活地控制编码和解码的流程，从而使得数据的传输更加高效和可靠。



## Functions:

### parseFieldParameters

parseFieldParameters是一个函数，用于解析给定字符串中的字段参数。 在encoding / json中，该函数用于从标记名称中解析字段标记，例如“foo，omitempty”。 

具体而言，它接收一个字符串参数，该字符串表示需要解析的标记名称。 它首先使用strings.Split函数将标记名称拆分为逗号分隔的字符串切片，然后对切片中的每个项目进行处理。 对于每个项目，它检查是否存在适当标识符（例如“omitempty”），并相应地设置FieldOptions结构体中的字段标志。 

最后，函数返回一个FieldOptions结构体，它包含有关该字段的所有标志。 如果标识符无效，则函数将返回一个err，表示该标记名称无效。 

在encoding / json中，FieldOptions结构体被用于表示字段选项（例如是否应忽略字段中的空值），并用于生成JSON格式的数据。 通过使用parseFieldParameters函数，字符串中的标记可以被解析成FieldOptions结构体，进而控制生成的JSON数据。



### getUniversalType

getUniversalType函数的作用是获取一个值对应的通用类型。

在编码中，不同的数据类型需要使用不同的编码方式进行表示。因此需要先将一个值转换为通用类型，然后再根据不同的具体数据类型进行编码。

getUniversalType函数会将一个值分配到以下几种通用类型之一：

- Int
- Uint
- Float
- Bytes
- String
- Bool
- Nil

具体分配方式如下：

- 如果值是一个 int 类型或者对 int 类型可以进行正确的类型断言，则分配到 Int 类型。
- 如果值是一个 uint 类型或者对 uint 类型可以进行正确的类型断言，则分配到 Uint 类型。
- 如果值是一个 float32 或者 float64 类型或者对这些类型可以进行正确的类型断言，则分配到 Float 类型。
- 如果值是一个 []byte 或者一个字节数组类型，则分配到 Bytes 类型。
- 如果值是一个 string 类型，则分配到 String 类型。
- 如果值是一个 bool 类型或者对 bool 类型可以进行正确的类型断言，则分配到 Bool 类型。
- 如果值是一个 nil，则分配到 Nil 类型。

在编码时，可以根据值的通用类型选择相应的编码方式。例如，对于 Int 类型的值可以使用 varint 编码方式进行表示。



