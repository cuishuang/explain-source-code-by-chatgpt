# File: base64.go

base64.go是一个Go语言包中的文件，它提供了一个对数据进行base64编码和解码的实现。

Base64是一种用于将二进制数据编码为可打印ASCII字符的方法，通常用于在网络上传输数据，或作为电子邮件中的附件。Base64编码将三个8位字节转换为四个6位字节，然后将这些6位值映射到可打印的ASCII字符集中。

在Go语言中，base64.go提供的API包括使用标准编码字符集的编码和解码方法，以及使用自定义字符集的编码和解码方法。此外，也提供了对IO流的编码和解码支持。

例如，对于一个字符串，要进行Base64编码，可以使用base64包中的StdEncoding.EncodeToString([]byte(string))方法；如果要对一个Base64字符串进行解码，可以使用base64包中的StdEncoding.DecodeString(string)方法。

总之，base64.go文件为Go语言开发者提供了一个方便且高效的Base64实现，使得在网络传输和数据存储中，Base64编码与解码变得容易。




---

### Var:

### StdEncoding

在 Go 语言中，base64 包含了三个“编码器”：StdEncoding、URLEncoding 和 RawStdEncoding。这些编码器可以将二进制数据编码成易于传输的 ASCII 字符序列。

StdEncoding 变量是最常用的编码器，它实现的是 RFC 4648 中的标准 Base64 编码规范。这个编码器将每 3 个字节的二进制数据编码成 4 个 ASCII 字符。

在 base64.go 文件中，StdEncoding 变量是一个 base64.Encoding 类型的变量，它包含了编码和解码所使用的字符集、填充字符等属性。

这个变量提供了 Encode、EncodeToString、Decode 和 DecodeString 等方法，用于对数据进行编码和解码。如果你想将一个字符串或者字节数组编码成 base64 字符串，可以使用 EncodeToString 方法。如果你想将一个 base64 字符串解码回原始的字节数组或字符串，可以使用 DecodeString 方法。

总之，StdEncoding 变量是 Go 语言中用于实现标准 Base64 编码规范的编码器，它提供了方便的方法和属性，可以帮助我们对数据进行编码和解码。



### URLEncoding

在Go语言的encoding包中，base64.go文件定义了Base64编解码算法相关的函数和类型。其中URLEncoding这个变量是一个定义在Base64编码表中的变体，它用于URL和文件名等需要对特定字符进行转义的场景，以避免产生歧义和冲突。

具体而言，URLEncoding将Base64编码表中的“+”和“/”字符替换成了“-”和“_”字符，同时去掉了末尾的“=”字符。这样可以保证Base64编码后的字符串不包含URL和文件名中可能会导致歧义的特殊字符，从而方便进行传输和存储。

在使用Base64编解码时，如果需要进行URL和文件名转义，就可以使用URLEncoding变量来指定编码方式，例如：

```go
data := []byte("hello world")
encoded := base64.URLEncoding.EncodeToString(data)
decoded, _ := base64.URLEncoding.DecodeString(encoded)
fmt.Println(string(decoded))
```

输出结果为：

```
hello world
```

可以看到，使用URLEncoding编码后的Base64字符串可以被正确地解码，并且也不包含任何需要进行URL和文件名转义的特殊字符。因此，在实际开发中，如果需要对Base64编码的数据进行传输和存储，就可以考虑使用URLEncoding这个变量来实现安全和方便的编码方式。



### RawStdEncoding

RawStdEncoding 是 encoding/base64 包中定义的一种编码方式。它是一种标准base64编码方式，但是在编码过程中不会添加额外的padding字符（=）。

具体来说，该变量是一个 base64.Encoding 类型的值，它定义了一个符号表，用于将字节数组转换为字符串。变量的定义如下：

var RawStdEncoding = StdEncoding.WithPadding(NoPadding)

可以看出，RawStdEncoding 是 StdEncoding 的一个实例，但是它采用了 NoPadding 的方式来进行编码，这也是它与 StdEncoding 的不同之处。

使用 RawStdEncoding 进行编码和解码时，如果字节数组的长度不是 3 的倍数，则会被填充为 3 的倍数，但是不会添加任何 padding 字符。这种编码方式常用于 HTTP Cookie、URL 和文件名等场景。

需要注意的是，由于不添加 padding 字符，解码过程中需要确保输入的字节数组的长度是 4 的倍数，否则解码会失败。



### RawURLEncoding

RawURLEncoding是encoding/base64包中一个Base64编码器（encoder）的实例，它使用标准的Base64编码表（base64表，包含A到Z、a到z、0到9、+和/字符）来将数据进行编码，然后通过URL安全方式进行转义（+用-替代，/用_替代），以便于在URL上进行传递。

RawURLEncoding被设计为用于在url.Values、JSON等格式中对二进制数据进行编码。它不会将换行符插入到编码的字符串中，也不会强制使用64个字符一行的编码格式。

使用RawURLEncoding时需要注意的是，由于将字符+和/用于URL安全转义，因此使用URL之前必须再进行一次解码，以获得原始的Base64数据。






---

### Structs:

### Encoding

在Go语言中，base64编码是一种二进制数据到ASCII字符序列的编码方式。这种编码方式最常见的应用场景是在网络传输中。Go语言的encoding/base64包提供了base64编码和解码的功能。在该包中，Encoding结构体负责进行编码和解码操作。

Encoding结构体有以下主要作用：

1. 将原始数据进行编码：将原始二进制数据转换为ASCII字符序列。

2. 将编码后的数据进行解码：将编码后的ASCII字符序列还原为原始二进制数据。

3. 提供编码表：该结构体中定义了标准base64编码表和URL安全的base64编码表。

4. 设置编码参数：该结构体中定义了一些编码参数，例如是否要在编码后的字符序列中添加换行符等。

通过使用Encoding结构体，我们可以在Go语言中方便地进行base64编码和解码操作。在实现网络传输等功能时，使用该结构体能够有效地提高数据传输效率和数据安全性。



### encoder

在go语言的encoding包中，base64.go文件中的encoder结构体是用于对byte slice（二进制数据）进行Base64编码的结构体。具体介绍如下：

encoder结构体的定义如下：

```
type encoder struct {
    w            io.Writer
    strict       bool
    wrap         int
    wrapbuf      [4]byte
    wrapIndex    int
    padChar      rune
    encode       []byte
    encodeNoPad  []byte
    decodeMap    [256]byte
    decodeMapInit bool
    padding      [64]byte
}
```

其中各个字段的作用如下：

- w io.Writer：用于写入编码后的数据的接口。
- strict bool：表示是否使用strict编码规则，如果设置为true，会忽略不合法的字符。
- wrap int：如果wrap字段的值大于0，则在每隔wrap个字符的位置插入一个换行符。默认为76（RFC 2045）。
- wrapbuf [4]byte：用于缓存待写入的编码字符，最多4个，当缓存满时，会通过w接口一次性写入。
- wrapIndex int：跟wrap字段配合使用，表示当前缓存buffer中已缓存的字符数。
- padChar rune：base64编码的填充符，默认为'='。
- encode []byte：存储编码字符表，根据RFC 4648规范定义，由大小写字母、数字和“+”、“/”构成。
- encodeNoPad []byte：跟encode字段配合使用，在strict模式下，当编码后的字符中出现“=”填充符时，表示编码过程结束，后面的填充符部分会被忽略。
- decodeMap [256]byte：存储解码表，根据RFC 4648规范定义，该表将编码字符映射为字节（0-63）。
- decodeMapInit bool：表示解码表是否已初始化。
- padding [64]byte：存储填充符，只有在strict模式下使用。

encoder结构体的方法有：

- Write([]byte) (int, error)：实现io.Write接口，将需要编码的数据写入缓存buffer中。
- Close() error：关闭encoder，将缓存buffer中的有效数据进行Base64编码，并写入w接口中。
- padLength(dataLen int) int：根据数据长度计算应该添加的填充字节数。
- padString(s []byte) []byte：在编码后的字符串中添加填充符。
- encode(dst, src []byte, stdEncoding bool) int：对数据进行Base64编码（标准版或URL安全版）。

总的来说，encoder结构体是Base64编码的核心数据结构，实现了对数据进行Base64编码的核心逻辑。



### CorruptInputError

在Go语言中，base64编码是一种将二进制数据转换为可视化字符串的方法。Go语言内置的encoding/base64包提供了对base64编码和解码的支持。

在base64编解码的过程中，可能会发生输入数据损坏的情况。例如，在解码过程中输入的字符串包含非法字符、长度不是4的倍数等等。为了表示这种错误，encoding/base64包中定义了CorruptInputError结构体。

当编解码过程中发生输入数据损坏的情况时，函数会返回一个CorruptInputError结构体的实例，告诉用户输入数据存在问题。通过检查CorruptInputError结构体中的错误信息，用户可以了解到具体的错误原因并采取相应的处理措施。

例如，当我们在使用encoding/base64包中的DecodeString函数解码一个字符串时，如果输入的字符串不是合法的base64编码格式，DecodeString函数会返回一个CorruptInputError实例，告诉我们输入数据存在问题。我们可以根据这个错误信息进行相应的处理，例如输出错误信息给用户，或者跳过这个输入数据继续处理下一个数据。



### decoder

在base64编码解码中，decoder结构体的作用是实现将base64编码转换回原始数据的功能。它包含一个io.Reader类型的r字段，用于从base64编码字符串中读取数据。同时，它还包含一个byte类型的buf字段，用于临时存储读取的数据。

decoder结构体的主要方法是Read方法，用于读取并解码base64编码的数据。它从r字段中读取base64编码数据，将其解码成原始数据，并将解码后的数据存储在buf字段中。Read方法接受一个字节切片作为参数，用于存储解码后的数据。如果buf字段中已经存在解码后的数据，则将buf字段中的数据复制到字节切片中，并将buf字段清空。如果r字段中没有数据可以读取，则返回io.EOF。

decoder结构体还实现了一个Close方法，用于关闭decoder并释放资源。

总之，decoder结构体的作用是将base64编码字符串转换回原始数据，并提供了Read和Close等方法来实现这一功能。



### newlineFilteringReader

在encoding/base64包中，newlineFilteringReader是一个结构体，它实现了io.Reader接口，并利用该接口将输入流的数据进行过滤和适当的处理，以将输入数据的换行符进行过滤和压缩，从而生成 base64 格式的编码数据。

newlineFilteringReader结构体的作用是读取输入流中的数据，并将其转换为base64编码，同时过滤和压缩输入数据中的换行符。这有助于减小编码数据的体积，并使其更加紧凑。它还可以防止邮件和HTTP等协议中可能会对数据进行换行的情况，从而保证数据传输的完整性和正确性。

该结构体可以在base64编码中起到非常重要的作用。当对数据进行base64编码时，使用newlineFilteringReader可以实现更好的数据传输效果，并且可以使编码长度更短。它的使用可以有效地提高编码效率和减少数据传输量。



## Functions:

### NewEncoding

NewEncoding函数的作用是创建并返回一个新的Base64编码方案。

具体来说，它将 base64 编码字符集中使用的 64 个字符（包括大写字母、小写字母、数字、加号和斜线等）转化为一个 64 位的自定义字符集，以便利用其进行自定义的编码操作。

该函数使用参数 alphabet 定义自定义字符集。如果未提供 alphabet 参数，则使用一组默认的编码字符集。

NewEncoding 函数返回的是一个指向新编码方案的指针，可以使用该指针中的方法进行编码和解码操作，如 EncodeToString 和 DecodeString 等。同时，新的编码方案可以使用其他基础编码方案进行转义和解码等操作。



### WithPadding

WithPadding是base64.go文件中的一个函数，其作用是返回一个新的 base64 编码器，该编码器使用指定的填充字符代替默认的填充字符。

在标准的base64编码中，用来填充编码后字符串末尾的字符是“=”，但在某些情况下，可能需要使用其他字符来进行填充。WithPadding函数允许用户传入参数指定所需的填充字符。

函数定义如下：

func (enc *Encoding) WithPadding(padding rune) *Encoding{}

参数padding是填充字符。该函数返回一个新的编码器，该编码器与enc编码器相同，但使用padding作为填充字符。

使用示例：

例如，如果用户想要用“#”号代替“=”号作为填充字符，则可以使用以下代码：

enc := base64.StdEncoding.WithPadding('#')

这将创建一个新的编码器，该编码器使用“#”号作为填充字符。

总之，WithPadding函数使用户能够控制base64编码器的填充字符，以便根据具体应用场景进行灵活配置。



### Strict

在 Go 语言的 encoding/base64 包中，Strict() 是一个函数，用于创建一个 Base64 编码器（Encoder）对象，其编码和解码过程中将严格检查输入数据的有效性。

具体来说，Strict() 函数创建的编码器对象会在编码和解码输入数据时，进行以下严格检查：

- 输入数据长度必须是 4 的倍数
- 输入数据必须由标准 Base64 字符集中的字符组成，否则将返回一个错误
- 编码后的输出数据必须是没有空格和换行符的一行字符流，否则将返回一个错误

Strict() 函数的作用是确保 Base64 编码和解码的输入数据的准确性和完整性，以避免在处理敏感数据时出现潜在的安全问题。

在处理非标准 Base64 数据（例如 URL 安全的 Base64 数据）时，可以使用 URLSafeEncoder() 函数创建一个不严格检查输入数据的编码器对象。但是，在安全性要求较高的应用中，仍建议使用 Strict() 函数创建严格检查输入数据的编码器对象。



### Encode

Encode函数是用于将一个字节数组编码成Base64格式的字符串。它将每3个字节（24个位）转换为4个Base64字符（每个字符6位），并添加填充字符(如果需要)使得字符串长度是4的倍数。

具体的实现过程如下：

1. 按照Base64规则将字节数组划分为长度为3个字节的块（如果字节数组长度不是3的倍数，则需要使用0填充到3的倍数），每个块将被转换为4个Base64字符。

2. 对于已经转换成Base64字符的块，它们的ASCII码值加上'+'、'/'、'A'，即可获得对应的Base64字符。

3. 为了避免出现Base64字符长度不足4的情况，需要在字节数组末尾添加填充字符'='，具体的填充规则如下：

- 当字节数组长度为1时，需要添加两个填充字符'=='；
- 当字节数组长度为2时，需要添加一个填充字符'='。

Encode函数的主要作用是为了在网络传输和数据存储中使用，因为Base64编码后的字符串只包含可打印ASCII字符，因此可以直接使用文本方式传输，而不需要考虑字节序或编码问题。此外，Encode函数也可以用于加密和数据签名等应用。



### EncodeToString

EncodeToString是encoding/base64包中一个函数，其作用是将一个byte类型的slice转换成字符串类型的base64编码。所谓的base64编码是一种将二进制数据按照固定格式转化成字符串的方法，常用于在网络传输和数据存储中使用。

具体来说，EncodeToString的实现过程如下：

1. 首先新建一个bytes.Buffer类型的变量b，用于存储base64编码后的结果。

2. 然后在一个for循环内，每次从输入的byte slice中取出3个字节，将其转换成4个base64字符，并逐个写入b中。如果输入的byte slice不足3个字节，需要进行padding操作。

3. 最后将bytes.Buffer类型的变量b转换成string类型并返回。

在实际使用中，EncodeToString函数常用于将一些二进制数据转换为可读的字符串形式，如在HTTP传输中对图片和文件的base64编码，以便于数据传输和存储。



### Write

在编码过程中，需要将字节序列转换为可打印字符的形式。 Write 函数是将编码后的数据写入缓冲区的函数，作用是将编码后的数据写入到指定的 byte slice 中，并返回写入的字节数和一个 error。该函数的实现主要是将输入的字节切片进行编码后写入传入的缓冲区中。

具体来说，其实现过程如下：

1. 首先，通过调用 encodeBuffer 方法，将需要编码的数据传入，此方法会返回一个编码后的字节数组，并记录编码过程中处理的字节个数。

2. 接着，将 encodeBuffer 方法返回的结果写入到缓冲区 slice 中，并更新 slice 的偏移量。

3. 最后，返回写入缓冲区的字节数和一个可能的错误。

通过对 Write 方法的调用，我们可以对需要编码的数据进行 base64 编码，并将编码后的结果写入指定的缓冲区中。这样，我们就可以方便地将数据进行传输或存储，并在需要的时候进行解码还原出原始数据。



### Close

在encoding/base64.go文件中，Close函数没有实际的实现代码，但是它是一个io.Closer接口。

io.Closer接口定义了一个Close()函数，调用这个函数可以释放资源，结束读写操作。Close()函数通常用于释放文件句柄、数据库连接、套接字等资源。在encoding/base64.go文件中，Close()函数用于关闭base64编码器的缓冲区。缓冲区在编码完成之后需要释放，以便给其他的编码器使用。

在Go语言中，io.Closer接口是一个非常常见的接口，几乎所有实现了io.ReadCloser、io.WriteCloser、io.ReaderAt、io.WriterAt等接口的类型都会实现Close()函数。Close()函数的作用是确保资源被正确释放，避免内存泄漏和资源浪费。



### NewEncoder

NewEncoder函数的作用是返回一个新的base64编码器，这个编码器可以将一段数据编码成base64格式。它接受一个io.Writer类型的参数，用于输出编码数据。当使用编码器进行编码时，数据会被分成3字节的块，每个块将被转换成4个base64字符。如果数据长度不是3的倍数，则会自动补充零字节直到长度为3的倍数。

调用NewEncoder函数后得到一个编码器后，可以使用编码器的Write方法将数据传递给编码器进行编码。当所有数据都编码完成后，需要调用编码器的Close方法来完成编码操作，这将会确保所有数据都已经被编码并输出。



### EncodedLen

EncodedLen函数返回给定长度的输入的base64编码字符串的长度，同时还会包括任何填充字符的数量。base64编码将3个字节的数据转换为4个字节的编码字符。如果不足3个字节，那么将使用填充字符“=”。因此，使用EncodedLen函数可以计算出输入数据的base64编码所需的最小字符串长度，以便在进行base64编码时为编码后的字符串预分配内存。EncodedLen函数的签名为：

```go
func EncodedLen(n int) int
```

其中n是输入数据的长度。EncodedLen可以通过以下公式计算：
```
encLen := (n + 2) / 3 * 4
```
这个公式中，n + 2是为了捕获剩余的1或2字节而做出的调整；然后将其除以3来计算需要进行多少次3字节编码，乘以4得到最终编码长度。EncdedLen函数主要用于预分配输出缓冲区大小，以避免在编码期间再进行内存分配。该函数对于需要大量的base64编码操作，如在网络上传输数据或将数据存储在数据库中，非常有用。



### Error

在go/src/encoding/base64.go中，Error函数是一个非导出的函数，其作用是返回一个错误字符串。

具体来说，当base64解码过程中出现错误，例如输入字符串长度不是4的倍数，或者解码出来的字节长度不是3的倍数等，就会调用Error函数返回一个错误字符串。该错误可以在调用base64.Decode函数时被捕获并处理。这为使用base64编码和解码提供了一种安全的方式，允许开发者检测并处理错误情况。

下面是Error函数的代码实现：

```
func (e *CorruptInputError) Error() string {
    if e.isBinary {
        return "illegal base64 data at input byte " + strconv.FormatInt(int64(e.offset), 10)
    }
    return "illegal base64 URL encoding at input byte " + strconv.FormatInt(int64(e.offset), 10)
}
```

其中，CorruptInputError是一个自定义错误类型，用于表示输入数据不合法的错误。如果输入数据是二进制数据，则返回“illegal base64 data”错误信息，否则返回“illegal base64 URL encoding”错误信息。在错误信息中还会包含错误的字节偏移量，方便开发者定位错误。



### decodeQuantum

在Go语言中的encoding/base64包中，decodeQuantum方法用于将一个64位base64编码字符串转换为一个3-byte的字节数组。

该方法的作用是将base64编码字符串的每一个字符转换成对应的6-bit二进制数值，然后将4个这样的6-bit二进制数值组成一个24-bit的二进制数，再将这个24-bit的二进制数拆分为3个8-bit的字节表示。

具体来说，该方法首先通过查找字符表中给定字符的位置来确定每个字符所代表的6-bit二进制数值，然后将这些6-bit二进制数值进行组合，并使用位运算来生成3个8-bit的字节表示。

该方法可以单独使用，也可以与其他base64解码方法一起使用，以便在解码base64编码字符串时将多个字符解码为有效的字节值。



### DecodeString

DecodeString函数用于将base64编码的字符串解码为原始二进制数据。它采用base64编码规则将给定的字符串转换回原始数据。

函数签名为：

func DecodeString(s string) ([]byte, error)

DecodeString函数接收一个字符串参数，该字符串是base64编码的数据字符串。如果解码成功，则返回解码后的原始二进制数据切片以及一个nil错误值。如果解码失败，则返回值为nil切片和一个非nil的错误值。

下面是DecodeString函数的实现原理：

首先，函数将base64编码的字符串切片转换为字节数组。

然后，函数采用base64规则从字节数组中提取出二进制数据，并存储在新的字节数组中。

最后，函数返回新的字节数组以及是否遇到错误。

DecodeString函数是Go语言标准库中一个非常重要的函数，因为它是用于处理base64编码和解码数据的常见方法之一。在网络通信、数据传输和文件储存等场景中，通常需要使用base64编码来确保数据的安全性和可读性。DecodeString函数提供了将base64编码数据解码为原始数据的快捷方式。



### Read

在go语言的encoding包中，base64.go文件定义了一系列用于base64编码和解码的函数和数据结构。其中，Read函数是一个实现了io.Reader接口的函数。

Read函数的作用是从一个byte slice中读取数据，并将其解码成base64编码格式。具体来说，Read函数接收一个byte slice作为参数buf，用于存储从输入流中读取的数据，同时也要求调用者提供一个表示输入字节流的io.Reader接口实例r。Read函数将从r中读取数据，并将其解码成base64编码格式，然后将解码后的数据存储到buf中，并返回读取的字节数以及可能出现的任何错误。

例如，可以使用以下代码段来读取base64编码的字符串：

```
package main

import (
  "encoding/base64"
  "fmt"
  "strings"
)

func main() {
  str := "SGVsbG8gV29ybGQgRG9l"
  data := make([]byte, base64.StdEncoding.DecodedLen(len(str)))
  n, err := base64.StdEncoding.Decode(data, []byte(str))
  if err != nil {
    panic(err)
  }
  fmt.Printf("%s", data[:n])
}
```

值得注意的是，Read函数只是base64解码算法的一部分，通常不单独使用。用户通常会使用base64包中提供的其他函数，如Encode，Decode，EncodeToString和DecodeString等函数，这些函数会自动调用Read函数来实现base64编码和解码。



### Decode

Decode函数是将base64编码的字符串解码为原始数据的一个函数。

具体来说，它会将输入的字符串首先进行base64解码，然后将解码后的字节流写入一个输出缓冲区，并返回缓冲区中写入的字节数。如果输入的字符串不是有效的base64编码，或者输出缓冲区的大小不足以容纳解码后的数据，函数将返回错误。

Decode函数在许多应用程序和协议中都被广泛使用，例如在HTTP头中传输二进制数据，或将邮件附件转换为纯文本格式。它是Go语言标准库中基本的编码/解码函数之一，提供了方便而高效的方法来处理base64编码数据。



### assemble32

在go/src/encoding/base64.go文件中，assemble32是一个函数，用于将4个相邻的6位字节合并成一个32位字节。

具体来说，这个函数接受一个长度为4的切片作为输入，将其转换为一个32位无符号整数。具体方法是将输入切片中的4个值（每个值表示6位二进制数）拼接成一个24位的二进制数，然后将这个24位的二进制数右移18位，再通过位或运算将其合并到一个32位的无符号整数中，返回这个整数。

在base64编码中，每3个8位字节会被转换成4个6位字节，而assemble32函数则将这4个6位字节重新组装成一个32位的无符号整数，方便进行后续的处理。



### assemble64

func assemble64(dst, src []byte, table *Encoding) (int, error) 是encoding/base64包中的一个函数，它的作用是将base64字符串解码为原始的字节数组，并将结果存储在目标字节数组中。

这个函数有三个参数：目标字节数组（dst）、源字节数组（src）、以及一个base64编码表（table）。目标字节数组用于存储解码后的结果，而源字节数组则包含需要解码的base64字符串。编码表则为base64加密、解密提供字符映射。

在函数内部，首先会对源字节数组进行切片，然后遍历每个字符，根据编码表找到对应的数值。如果该字符不在编码表中，函数会返回一个解码错误。如果字符是填充字符（‘=’字符），则表示该字节串已经结束，函数会返回已解码的字节数。

根据base64编码的规则，每4个字符对应3个字节。因此在函数的循环中，每次都先判断是否有足够的输入字符，如果不足4个字符，函数就返回一个解码错误。然后在循环体内将4个输入字符解码为3个输出字节。

最后，函数返回已解码的字节数以及可能的错误。如果没有错误，则返回的字节数应该等于目标字节数组的长度。



### NewDecoder

NewDecoder函数是用于创建一个新的base64解码器的函数，具体作用如下：

1.接收一个io.Reader类型的参数，用于从该读取数据并解码。

2.返回一个base64编码解码器指针。

3.可以使用创建的base64编码解码器指针的Decode方法来解码编码的数据。

4.解码器在读取数据时，遇到无效的base64字符时会返回一个错误，以指示解码失败。

5.解码器会忽略输入数据中的空格和回车等空白字符。

总之，NewDecoder函数是为了简化base64编写和解析过程，让它们更加简单、易用。通过使用NewDecoder函数，我们可以轻松、快速地解码任何使用base64编码的数据。



### DecodedLen

DecodedLen 函数是在编码的字符长度已知的情况下计算将使用 Base64 编码的数据解码后的长度。Base64 编码将会把每 3 个字节的数据分成 4 段，在每段后面都会加上 1 到 2 个填充字符。因此，解码后的长度将是编码字符数除以 4 的结果，而如果有填充字符，则需要减去相应的数量。

该函数的具体实现方式是通过计算字符长度与 4 取模的结果，得出填充字符的个数，再将编码字符数减去填充字符和基础长度（由编码数据的每 3 个字节分成 4 段得到，对于每个没被分成 4 段的字节会被填充为 0），最终得到解码后的长度。

例如，如果编码字符数为10，那么解码后的长度将是 7，因为将 10 除以 4 就是 2，不过因为总长度不够，还需要加上 2 个填充字符才能凑成 4 的倍数，因此需要再减去 2。

DecodedLen 函数在使用 Base64 进行数据解码之前，可以帮助开发者先计算出需要分配的存储空间，从而在内存空间和处理时间上达到更好的平衡。



