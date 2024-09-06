# File: base32.go

编码和解码base32的Go语言实现。

Base32是一种基于32个字符的编码方法，用于将二进制数据转换为可打印的ASCII字符。它的主要应用是在电子邮件、网站验证和其他情况下，需要将二进制数据转换为文本格式以便发送或存储。base32.go提供了一种方便的方法来编码和解码数据，以便在这些场景中使用。

base32.go文件中包括了以下几个函数：

- EncodeToString(data []byte) string

EncodeToString函数将二进制数据转换为Base32编码的字符串。它接受一个字节数组作为输入，返回一个字符串表示该字节数组的Base32编码。

- DecodeString(s string) ([]byte, error)

DecodeString函数将Base32编码的字符串转换为二进制数据。它接受一个字符串作为输入，返回一个字节数组表示该字符串的二进制数据。如果转换失败，它将返回一个错误。

- StdEncoding

StdEncoding是一个Base32编码器，它由RFC 4648定义。它使用ABCDEFGHIJKLMNOPQRSTUVWXYZ234567字符集进行编码，并使用=填充。它是最常用和最稳定的Base32编码方式。

- NewEncoder(enc *Encoding, w io.Writer) *Encoder

NewEncoder函数创建一个Base32编码器，它从给定的io.Writer接口中写入编码数据。它接受一个编码器和一个写入器作为参数，并返回一个编码器指针。

- NewDecoder(enc *Encoding, r io.Reader) *Decoder

NewDecoder函数创建一个Base32解码器，它从给定的io.Reader接口中读取编码数据。它接受一个编码器和一个读取器作为参数，并返回一个解码器指针。

这些函数提供了一种方便的方法来将数据进行Base32编码和解码。使用这些函数，可以将二进制数据转换为ASCII文本，或将Base32编码的数据转换回其原始的二进制格式。




---

### Var:

### StdEncoding

在 Go 语言中，base32.go 文件中的 StdEncoding 变量是一个 base32 编码器，其中包含了标准的 base32 编码字符集和编码/解码逻辑。它是一个常见的编码器，可以对任何数据进行 base32 编码和解码。

StdEncoding 变量是一个 Base32 对象，它具有以下属性：

1. 其字符集（base32 编码字符集）由 RFC 4648 定义，并使用“ABCDEFGHIJKLMNOPQRSTUVWXYZ234567”这 32 个字符表示。

2. 在 base32 编码中，“= ”为填充字符。

3. 默认的 padding 是使用“= ”填充，padding 的长度为 0 到 7 个字符。

4. 在编码和解码过程中，忽略所有的空格和换行符。

通过使用 StdEncoding，可以将数据编码为 base32 格式，并将其传输或存储在不支持二进制数据的环境中，例如电子邮件或 HTTP 请求。同样，使用该编码器可以将base32格式的数据还原为其原始二进制形式。

总之，StdEncoding 变量是 Go 语言中的一个标准 base32 编码器，可以用于对数据进行编码和解码，以及在不支持二进制数据的存储和传输环境中使用。



### HexEncoding

在go编程语言中，encoding/base32包中的HexEncoding变量是一个base32编码使用的编码方案。它是由32个字符串组成的字母表，每个字符代表5个位元（bit），并单独标记了十六进制数字中的每个字符（0到9和A到F）。这意味着每五个位元可以转换为一个字母，从而缩短了要编码的字符长度。

在HexEncoding变量内，每个字母的值都表示为相应的二进制值。例如，变量中的第一个字符“A”对应二进制值“00000”，变量中的最后一个字符“Z”对应二进制值“11111”。因此，将任何二进制数据编码为base32时，可以将其转换为相应的字母表中的字符。

在实际应用中，可以使用HexEncoding变量将字符数据编码为base32，以实现一定程度的数据保密性。此外，由于base32编码后的字符只包含字母，没有数字或符号，因此可以在某些情况下增强数据的可读性和可靠性。






---

### Structs:

### Encoding

在Go语言的encoding包中，base32.go文件中的Encoding结构体是一个处理base32编码的结构体。它嵌入了一个base编码结构体，并提供了base32特定的编码和解码方法。

Encoding结构体的主要作用是提供一种将二进制数据转换为可打印ASCII字符的方法。这种编码可以用于将数据存储在文本文件中、通过网络传输数据等场景。base32编码是一种将8位的二进制数据转换为5位的ASCII字符的编码方式，因此使用base32编码后的字符串长度比原始二进制数据长度长约60%。

Encoding结构体包含以下字段：

- encode：一个长度为32的数组，包含所有32个可能的base32字符
- decode：一个长度为256的数组，包含所有256个可能的字节的解码值

该结构体定义了以下方法：

- EncodeToString(src []byte) string：将src中的二进制数据编码为base32字符串。
- Encode(dst, src []byte) int：将src中的二进制数据编码为base32并将结果写入dst。返回写入的字节数。
- DecodeString(s string) ([]byte, error)：将base32字符串s解码为原始二进制数据。
- Decode(dst, src []byte) (int, error)：将src中的base32编码数据解码为原始二进制数据并写入dst。返回写入的字节数和可能的错误。

总之，Encoding结构体是一个方便使用base32编码和解码的封装，它使得将二进制数据编码为可打印字符和从编码后的字符串解码回原始二进制数据变得非常简单。



### encoder

encoder这个结构体是用于base32编码的，它定义了一些用于编码的参数和函数。具体来说，encoder结构体包含以下字段：

- enc：一个[32]byte数组，用于将5个输入字节转换成8个输出字符的映射表。
- paddingChar：用于填充输入字节的字符。通常为'='。
- encodeTable：一个[256]byte数组，用于将ASCII字符转换成base32字符的映射表。
- buf：一个[]byte类型的缓冲区，用于暂存未完成的编码结果。
- n：buf中已有的字节数。

encoder结构体中定义了以下方法：

- Encode：将指定字节切片进行base32编码。
- EncodeToString：将指定字节切片进行base32编码，并返回编码结果的字符串表示。
- Write：将指定字节切片写入编码器，并返回写入字节数。
- Close：关闭编码器，并返回任何未完成的编码结果。

在实际使用中，encoder结构体可以创建多个实例，并在每个实例中设置不同的paddingChar值或enc映射表，以适应各种需求。



### CorruptInputError

在encoding/base32.go文件中，CorruptInputError是一个结构体，它表示在解码base32输入时出现了错误。

当解码base32时，输入数据不一定总是有效或完整的，它可能包含错误，比如输入数据包含不合法的字符、长度不合法、或者是缺失数据。这时候就需要使用CorruptInputError来表示解码错误。

CorruptInputError结构体包含一个错误信息，用于指示输入数据中有什么错误。它通常与ErrLength和ErrFormat一起使用，一旦发现解码错误，就会返回一个CorruptInputError类型的错误，以便调用者可以识别和处理错误。

总之，CorruptInputError的作用是表示base32解码输入时的错误，并提供错误信息以方便调用者处理错误。



### decoder

decoder结构体是用于将base32编码的数据解码成原始数据的结构体。它包含一个解码表和一个缓冲区，用于存储解码过程中的中间值。decoder结构体实现了io.Reader接口，因此可以通过Read方法一次性读取所有数据并解码。

具体来说，decoder结构体中的解码表包含了对应每个base32字符的十进制值，以便于将其转换成二进制形式进行解码。在解码过程中，decoder结构体将输入的base32字符序列逐个解析，将每个字符转换成其对应的十进制值，并将其转换成二进制形式。然后，decoder结构体将这些二进制位按照一定规则重新组合，得到解码后的原始数据。在这个过程中，decoder结构体会利用缓冲区中的空间存储中间结果，以便于在后续的解码过程中能够高效地访问这些中间结果。

总之，decoder结构体是用于将base32编码的数据解码成原始数据的核心结构体，它包含了解码表和缓冲区等关键组件，能够高效地实现base32数据的解码。



### newlineFilteringReader

newlineFilteringReader是一个结构体，实现了io.Reader接口，用于过滤换行符。

在base32编码中，每5个字节会被编码成8个字符，如果源数据不是5的倍数，则需要在末尾补0。所以在编码过程中，将源数据分割成5个字节一组进行编码。但有些情况下，源数据在存储或传输过程中会自动加入换行符，这样会破坏数据的完整性。所以newlineFilteringReader就是用于过滤这些换行符的。

具体来说，newlineFilteringReader实现了io.Reader接口的Read方法，它会从输入流中读取数据并过滤掉换行符后返回数据。如果输入流中没有更多数据可读，则返回EOF。在base32编码过程中，数据源就是通过newlineFilteringReader来读取的，以保证数据的完整性。



## Functions:

### NewEncoding

NewEncoding函数是base32包中一个非常重要的函数，其作用是创建一个新的base32编码器，以便在编解码中使用。

该函数接受一个*Encoding类型的参数，如果该参数为nil，则默认使用RFC 4648（标准base32）编码器进行编解码。如果你想使用一个不同的base32编码，可以使用该函数自定义。

在实现自定义编码器时，最重要的是编码表（EncodingTable），它指定了使用哪些字符代表每个数值。在RFC 4648中，使用的字符集为ABCDEFGHIJKLMNOPQRSTUVWXYZ234567。你可以通过调用WithEncodingTable来指定自己的编码表，并使用NewEncoding生成一个自定义的编码器。 

例如，如果你想使用RFC 3548的Base32编码方案，其字符集为ABCDEFGHIJKLMNOPQRSTUVWXYZ234567，你可以通过以下代码进行设置：

```
import "encoding/base32"

// Create a custom encoding using RFC 3548 specification
customEncoding := base32.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567").WithPadding(base32.NoPadding)
```

在上面的代码中，NewEncoding创建了一个新的编码器，使用自定义字符集"ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"，并通过WithPadding函数指定了无填充编码方案，这样就可以使用自定义的base32编码器进行编解码操作了。

总之，NewEncoding函数提供了一个非常方便的接口来创建和自定义base32编码器。它是在编解码数据时非常实用的工具函数。



### WithPadding

WithPadding在base32编码中设置填充字符。当以base32编码的数据长度不足8等于时，通常需要在数据末尾填充一些字符，以保证编码后的数据长度为8的倍数。填充字符可以是任何字符，但大多数情况下使用'='符号作为填充字符。

WithPadding函数返回的是指定填充字符的新的base32编码程序。可以使用此函数的返回值进行编码操作。

函数声明如下：

```
func (enc *Encoding) WithPadding(padding rune) *Encoding
```

其中，padding是填充字符的rune值。

例如，如果要使用'.'作为填充字符，则可以这样调用WithPadding函数：

```
enc := base32.WithPadding('.')
```

这将返回一个新的base32编码程序，使用'.'作为填充字符。之后就可以使用此编码程序进行base32编码操作。



### Encode

Encode函数是用于将字节数组转换为base32编码的字符串。这个函数主要用于数据传输和存储，因为base32编码可以更好地保证数据的可读性和传输正确性。

该函数的参数是一个字节数组，代表要进行编码的数据。返回值是一个字符串，代表base32编码后的结果。

在具体实现过程中，该函数会将输入字节数组按照5个字节一组进行处理。每组5个字节转换成8组5个bit的编码，然后根据base32字符表，将编码转换成32进制下的字符表示形式。

这个函数是编码过程中的关键函数之一，通过该函数的调用，可以快速将字节数组进行base32编码，并且方便地进行数据处理和存储。



### EncodeToString

EncodeToString是一个在Go语言中将字节数组转换为base32编码字符串的函数。

在计算机中，数据传输和存储需要使用二进制数据。但是，二进制数据并不总是直观或易于记忆的。因此，我们需要将二进制数据转换为可读、易于存储和传输的字符串格式。在这种情况下，编码和解码功能是至关重要的。

base32编码是一种用于在二进制数据和可读字符串之间转换的编码格式。它将8位二进制数据转换为5位的base32字符。在转换过程中，编码算法可以保证输出字符串只包含ASCII字符集中的可打印字符，它可以安全地在电子邮件、DNS、JSON等地方使用。

在base32.go文件中，EncodeToString函数接收一个字节数组并返回一个base32编码的字符串。在该函数内部，它首先创建一个新的base32编码器，然后使用该编码器将字节数组编码为base32字符串。 最后，它返回转换后的base32字符串。

这个函数在很多场合中非常有用。例如，当需要将二进制数据存储在数据库或发送到其他服务器时，就可以使用EncodeToString将二进制数据转换为base32编码的字符串，然后存储或发送base32字符串。在需要将其还原为原始数据时，可以使用对应的DecodeString函数来解码。



### Write

Write函数是实现io.Writer接口的方法，它是在base32编码过程中的一个核心函数。基本作用是将输入的字节切片中的字节，逐个转化为base32编码后的字符，并将结果输出到缓存中。该函数接受一个字节切片b作为输入，返回两个值：n表示输出的字符数，err表示是否发生了错误。

具体实现过程如下：

1. 循环遍历输入的字节切片b，先将当前字节和前面未处理完的部分组成一个64位的代码（code），并获取当前代码的有效位数（bits）。

2. 检查当前代码是否满足base32编码的字符数量要求（5或8位），如果不满足，则继续读取下一个字节，并重新组成代码。

3. 如果当前代码满足编码要求，则通过实现的encode函数将当前代码转化为base32编码后的字符，并将结果输出到缓存中。

4. 重复执行以上步骤，直至整个字节切片b被处理完毕。

因此，Write函数是实现base32编码的一个必要函数，同时也是基于io.Writer接口实现的常见方式。它将一组字节转换为可读的文本字符串，并存储在缓冲区中。如果缓存区中的空间不够存储编码后的字符，则会返回Write函数错误。



### Close

在Go语言的encoding/base32包中，Close函数被定义为无操作函数。它仅是为了满足io.Closer接口的要求而存在。

io.Closer接口表示一个可以关闭的资源。在Go中，我们经常需要处理文件、网络连接和其他情况下，需要在使用完后将资源释放（关闭）以释放资源。

但是，在编码和解码操作中，数据是直接从内存读取和写入，而不是使用资源对象。因此，Close函数没有什么实际用途。

在base32.go文件中，Close函数的定义如下：

```
func (e *Encoding) Close() error {
	return nil
}
```

我们可以看到，它只是简单地返回一个nil值，因为在进行编码和解码时并不需要对资源进行关闭操作。



### NewEncoder

NewEncoder函数在encoding/base32包中定义，其作用是创建一个新的base32编码器。当需要将二进制数据编码为base32格式时，可以使用此函数创建一个编码器，并使用其Encode方法进行编码。

具体来说，NewEncoder函数接收一个io.Writer类型的参数，用于指定编码后数据的输出位置。当调用编码器的Encode方法时，编码后的数据将被写入该io.Writer中。例如，可以使用如下代码创建一个base32编码器，将编码后的数据写入标准输出：

```
package main

import (
    "encoding/base32"
    "os"
)

func main() {
    encoder := base32.NewEncoder(base32.StdEncoding, os.Stdout)
    encoder.Write([]byte("hello world"))
    encoder.Close()
}
```

上述代码中，NewEncoder函数的第一个参数是base32.StdEncoding，表示使用标准的base32编码方式进行编码。第二个参数是os.Stdout，表示将编码后的数据写入标准输出。后续使用Write方法将数据写入编码器，并使用Close方法确保所有数据都已被写入。执行上述代码将会输出如下结果：

```
NBSWY3DPEB3W64TMMQQQ====
```

需要注意的是，NewEncoder函数返回的编码器是一个io.WriteCloser接口，即它还具有关闭方法。当编码完成后，应该调用Close方法以确保所有数据都已经写入。如果不调用Close方法，可能会导致数据丢失或不完整。



### EncodedLen

EncodedLen函数是一个用于计算基于指定编码方式的数据长度的函数。在base32编码中，它计算了未经填充的数据块的编码长度。

函数的具体作用如下：

1. 接受参数n，它表示要编码的字节数。

2. 使用算法计算出基于指定编码方式的未经填充的数据块的编码长度。

3. 返回结果，也就是未经填充的数据块的编码长度。

在base32编码中，每个数据块使用5个位来表示，所以每个数据块可以表示32个不同的值（2^5）。如果要编码n个字节，它将转换为n×8个位。根据base32编码规则，每个数据块实际上只使用其中的5个位。我们可以使用下面的公式计算编码块的数量：

numBlocks = (n * 8 + 4) / 5

由于每个块未经填充的数据长度是5个位，因此对编码块的数量进行舍入，可以得出未经填充的数据块总长度：

unpaddedLen = numBlocks * 5

但是，在修补填充字节的情况下，每个编码块可能需要填充。因此，我们可以使用以下公式计算填充的数据长度：

paddingLen = (8 - (n * 8) % 5) % 5

最后，我们可以将无填充长度和填充长度相加，以得出最终的编码长度：

finalLen = unpaddedLen + paddingLen

EncodedLen函数就是使用这些公式来计算基于指定编码方式的数据长度的。



### Error

在encoding/base32.go文件中，Error函数用于返回一个字符串，该字符串包含了错误信息，用于解释为什么一个操作失败或者非法输入。这个函数的具体作用如下：

1. 检测编码或解码等操作是否出错
在编码或解码过程中，如果出现了非法字符或者数据不完整等错误，就会通过调用Error函数来返回错误信息，使得编码或解码的操作失败。

2. 提供错误信息给用户
当程序运行出错时，可以将Error函数返回的错误信息展示给用户，以便于他们了解问题所在并采取相应的措施。

3. 记录日志
当出现错误时，Error函数返回的错误信息可用于记录日志，以便于在事后查看问题、定位错误并解决错误。

总之，Error函数是一个通用的、用于解释错误信息的函数，它可以提高代码的可读性、可维护性和可靠性，是 Go 语言中的一个常用函数。



### decode

decode函数的作用是将给定的base32编码的字符串解码成原始数据。

具体来说，decode函数会将base32编码的字符串逐个字符地解码成5个二进制位，然后根据这些二进制位重组出原始数据。在解码过程中，还会忽略空格、回车等无效字符。

该函数的输入参数为一个字符串，表示需要解码的base32编码字符串。输出为解码后的原始数据，类型为[]byte。

在实际应用中，decode函数通常用于处理通过base32编码的数据，以便将其转换成可读的原始数据，例如将电子邮件地址中使用的base32编码字符串解码成原始的邮件地址。



### Decode

Base32是一种用于将二进制数据编码成文本字符串的方法。Base32编码使用32个不同的字符表示5个比特位的值。base32.go文件中的Decode函数就是将Base32编码的文本字符串解码成二进制数据。

具体来说，Decode函数接收一个Base32编码的文本字符串作为输入，然后将其中的每个字符转换成相应的5个比特位的值，最终将这些值组合成一个二进制数据数组返回。例如，输入字符串"KRYRQ===="会被解码为二进制数据{0x9f, 0x46, 0x23}。

在Go标准库中，通常使用Base64编码二进制数据。但是在某些情况下，由于Base64使用"+"和"/"这两个字符作为编码字符，因此可能会在URL参数等特殊情况下出现问题。因此，有时候需要使用Base32编码。Decode函数则是用来解码Base32编码的字符串，方便在程序中进行处理。



### DecodeString

DecodeString是从base32编码字符串中解码数据的函数。它将一个base32编码的字符串作为参数，并返回解码后的原始字节数组。

函数签名如下：

func DecodeString(s string) ([]byte, error)

其中，s是输入的base32编码字符串，返回值是解码后的字节数组和一个error类型的错误信息。

具体而言，DecodeString会执行以下步骤：

1. 检查输入字符串的长度是否满足要求。根据RFC 3548规范，base32编码字符串的长度必须是8的倍数。

2. 创建一个新的字节数组，用于存储解码结果。

3. 对输入字符串进行解码。DecodeString会逐个读取输入字符串的字符，并将其转换为对应的二进制值。然后，它会将这些二进制值组合起来，生成一个完整的字节，并将其保存到新的字节数组中。如果输入字符串的长度不是8的倍数，则DecodeString会在最后一个字节中填充一些0位，以保证字节数组的长度正确。

4. 返回解码后的字节数组和一个可能的错误信息。如果解码过程中发生错误，DecodeString会返回一个非nil的error类型值。

总之，DecodeString函数是将base32编码字符串转换为原始数据的重要工具。在许多应用程序中，它被用于解码经过base32编码的数据，例如在网络传输和数据存储中。



### readEncodedData

readEncodedData这个函数的作用是解码输入的base32数据并将其写入给定的字节数组中。

该函数的输入参数有Reader接口类型的r，表示输入的base32数据流；output []byte类型的buf，表示写入解码后数据的字节数组；alphabet []byte类型的encoding alphabet，表示使用的base32编码表。

函数的具体原理是从输入的数据流中读取数据，然后根据编码表将每个字符解码成5位二进制数字，并将连续的8个二进制数字组合成一个字节，写入输出字节数组中。

函数的返回值为写入buf中的字节数，同时会返回一个error类型的值，表明解码过程中是否出现了错误。

总之，readEncodedData函数是实现base32解码的核心函数，它能够将输入的base32数据流解码成普通的二进制数据，并保存到指定的字节数组中，从而方便后续的数据处理和使用。



### Read

在go语言的encoding包中，base32.go文件中的Read函数实现了io.Reader接口，用于从base32编码的输入流中读取数据。

具体来说，该函数的作用是从一个byte数组的指定位置开始读取指定长度的数据，并将其解码为base32格式，然后将解码后的数据存储到一个指定的byte数组中。该函数返回读取的字节数和一个error。如果读取的数据长度不足或者解码失败，会报错并返回相应的错误信息。

该函数的定义如下：

```
func (dec *Decoder) Read(p []byte) (n int, err error) {
    // ...
}
```

其中，`dec`是一个Decoder类型的结构体，表示一个base32解码器。`p`是一个byte数组，表示一个读取缓冲区，这个函数会将解码后的数据存储到这个数组中。

在函数的实现中，首先会判断输入的缓冲区是否为空，如果是则返回0字节读取，不进行任何处理。接着，该函数会先计算出所需的字节数，然后调用Decoder的decode方法进行解码。在decode方法中，该函数会先将输入缓冲区标准化为不含填充字符的长度，再将其解码为UTF-8格式，最后将解码后的结果复制到输出缓冲区中。

在处理完所有的输入后，该函数会返回实际读取的字节数和一个可能的错误。如果解码过程中出现了错误，那么该函数会返回相应的错误信息。



### stripNewlines

stripNewlines函数的作用是从字符串中移除所有的换行符(\n)和回车符(\r)。这个函数主要用于将base32编码中可能出现的换行符和回车符去除，以便使编码结果更加紧凑。

stripNewlines函数的实现很简单，它使用了strings.Replace函数将"\n"和"\r"替换为空字符串""。具体实现如下：

```
func stripNewlines(str string) string {
    return strings.Replace(strings.Replace(str, "\n", "", -1), "\r", "", -1)
}
```

其中，函数的输入参数为一个字符串(str)，返回值为一个与str相同但不包含换行符和回车符的字符串。函数使用了内置的strings包中的Replace函数，将"\n"和"\r"分别替换为空字符串""，并使用-1作为替换的最大次数，这样就可以将所有的换行符和回车符都去除了。



### NewDecoder

在encoding/base32.go文件中，NewDecoder函数用于创建一个新的base32解码器。解码器可以将base32编码的数据转换成原始数据。

该函数的返回值是一个*Decoder类型的指针，具体定义如下：

func NewDecoder(enc *Encoding, r io.Reader) io.Reader

这个函数有两个参数：

1. enc：一个*Encoding类型的指针，指定使用的base32编码格式。在Go语言中提供了4种base32编码格式，分别是：StdEncoding、HexEncoding、ExtendedHexEncoding和NewEncoding。其中，StdEncoding是标准的Base32编码，HexEncoding是Base32编码的16进制版本，ExtendedHexEncoding是HexEncoding的扩展版本，NewEncoding则是一个可定制的Base32编码格式。

2. r：一个io.Reader类型的参数，这个参数指定要解码的数据源。

NewDecoder函数返回的是一个io.Reader类型的解码器，可以通过调用Read方法，将解码后的数据读取出来。

举个例子，假设要将一个base32编码的字符串解码成原始的字节数组，代码如下所示：

```go
package main

import (
    "encoding/base32"
    "fmt"
    "strings"
)

func main() {
    input := "MZXW6YTB"
    reader := strings.NewReader(input)
    decoder := base32.NewDecoder(base32.StdEncoding, reader)
    output := make([]byte, 10)
    n, err := decoder.Read(output)
    if err != nil {
        fmt.Println("error:", err)
    }

    fmt.Println(output[:n])
}
```

在这个例子中，我们首先声明了一个输入字符串input，它是一个base32编码的字符串。接下来，我们使用strings.NewReader函数将其转换成io.Reader类型的数据源。

然后，我们调用NewDecoder函数创建了一个新的base32解码器，使用了标准的Base32编码。接着，我们声明了一个长度为10的字节数组output，用于存储解码后的数据。

最后，通过调用decoder的Read方法，将解码后的数据读取到output中，并且打印输出。

此时，输出的是原始数据的字节数组，即：[104 101 108 108 111]。



### DecodedLen

DecodedLen函数是计算使用Base32编码的字节数组解码后的实际字节数的函数。在使用Base32编码和解码时，同一段内容经过编码后会变成大于原来的字节数，因为Base32编码后每个字节被扩展到了5位。因此，在实际使用中，需要计算出解码后的字节数，以便在解码时分配足够的字节数组。

DecodedLen函数的作用是根据Base32编码的字节数组的长度n，计算出解码后的实际字节数。具体计算方法为：解码后的字节数 = （n + 7）/ 8 * 5。其中，n + 7表示为了保证能够分配足够的字节数组而添加的7个空间，/ 8表示将结果除以8向下取整，* 5表示每个字节被扩展到了5位。

例如，如果Base32编码的字节数组长度为10，则解码后的字节数为8，因为（10+7）/ 8 * 5 = 17.5，向下取整为8。因此，当我们要解码使用Base32编码的字节数组时，可以使用DecodedLen函数计算出解码后的实际字节数，并分配足够的字节数组。



