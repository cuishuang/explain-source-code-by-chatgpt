# File: chunked.go

chunked.go是Go语言中net包中的一个文件，它实现了HTTP的chunked编码和解码。

HTTP chunked编码是一种在传输过程中对HTTP消息体进行分块传输的方法，可以有效地处理消息体大小未知或非常大的情况。在chunked编码中，消息体被分成多个大小固定的块并且每个块都带有一个十六进制的长度值，表示该块的大小。最后一个块必须为0长度的块，表示消息体传输结束。

chunked.go文件中实现了chunkedReader和chunkedWriter两个struct，分别用于读取和写入chunked编码的HTTP消息体。当一个HTTP消息体需要被分块传输时，可以使用chunkedWriter将消息体进行编码并通过网络传输；而当接收方收到消息体时，可以使用chunkedReader将消息体进行解码并还原出原始的消息体数据。

除了chunkedReader和chunkedWriter之外，chunked.go文件中也实现了一些辅助函数和变量，包括对chunked编码进行解码的parseHex()函数和需要用到的常量和错误信息。

总之，chunked.go文件的作用就是实现了HTTP的chunked编码和解码，使得HTTP消息可以更加灵活和高效地进行传输。




---

### Var:

### ErrLineTooLong

ErrLineTooLong是一个net包中的错误变量，在chunked.go文件中的作用是表示HTTP报文中的行（例如请求行或响应头）超过了规定的最大长度。

HTTP协议规定，HTTP报文头的行的最大长度不能超过8192个字节，如果超过了这个限制，就会导致ErrLineTooLong错误。

在chunked编码中，每个chunk都会有一个类似于HTTP头的行结构，因此ErrLineTooLong错误也会在处理chunked编码的过程中出现。

如果在解析HTTP报文时出现ErrLineTooLong错误，通常会返回一个400 Bad Request的响应，表示请求语法错误。



### semi

在go/src/net/chunked.go文件中，semi变量是一个bool类型的变量，用于标记当前解析的chunk是否已经结束。

在HTTP协议中，一个chunk可以跨越多个TCP数据包，在解析chunk时需要使用semi变量来标记解析的进度。当semi为false时，表示当前chunk数据还未解析完整；当semi为true时，表示当前chunk数据已经解析完成。

具体来说，当解析到chunk的分块长度时，需要读取相应长度的数据，判断读取到的数据长度是否等于分块长度。如果相等，则将semi标记为true，表示当前chunk已经解析完整；否则继续读取后面的数据，直到解析完整个chunk。

总之，semi变量在解析chunked编码的HTTP响应体时，起到了标记当前chunk解析进度的作用。






---

### Structs:

### chunkedReader

chunkedReader结构体的作用是用于读取HTTP的chunked编码数据。

在HTTP协议中，chunked编码是一种数据传输方式，允许发送不定长度的消息体。消息体被分割成多个chunk，在每个chunk之前都会发送该chunk的长度。最后一个chunk的长度为0，表示消息体的结束。

chunkedReader结构体实现了io.Reader接口，用于读取chunked编码的数据。它内部维护了一个bytes.Buffer类型的缓冲区，用于存储从输入流中读取的数据。

当调用chunkedReader结构体的Read方法时，它会首先读取该chunk的长度，然后从输入流中读取对应长度的数据，并将数据写入缓冲区。直到读取到长度为0的chunk时，表示消息体已经读取完毕，Read方法才会返回io.EOF。

除了Read方法，还有一个ReadByte方法，用于从缓冲区中读取单个字节的数据。同时，该结构体还实现了io.ReadCloser接口，用于在读取完毕后关闭输入流。

总之，chunkedReader结构体的作用就是将输入流中的chunked编码数据解析成可供应用程序使用的消息体。



### chunkedWriter

chunkedWriter是一个实现了io.Writer接口的结构体，用于将数据进行分块传输。

在HTTP协议中，使用chunked编码时，发送的数据被分成若干段，每段数据的前面会加上这段数据的长度（16进制数字），最后以一个长度为0的chunk结束。接收方根据chunk的长度来分段接收数据，直到读取到长度为0的chunk为止。

chunkedWriter的作用就是将一个流式的数据分块，按照chunked编码的格式写入输出流中。它通过实现Write方法来处理传入的数据，并在每个数据块前写入该块的长度。同时，在写入完所有数据之后，它会写入长度为0的chunk，标志着该数据流的结束。

使用chunkedWriter可以防止在传输过程中出现延迟卡住的现象，提升数据传输的效率和稳定性。



### FlushAfterChunkWriter

FlushAfterChunkWriter是net包中chunked.go文件中的一个结构体，用于将写入的数据分块传输到网络连接中。

具体来说，FlushAfterChunkWriter会将每次写入流中的数据进行分块处理，然后将块数据写入底层网络连接中。分块传输是HTTP/1.1协议中的一种特性，它可以使得服务器在处理大数据时，在发送完整个响应之前，可以先将部分数据发送给客户端，从而可以更快地响应客户端的请求。

在使用FlushAfterChunkWriter时，通过传递一个io.Writer接口类型的对象给它，可以将数据写入这个对象中。每次写入的数据都会被分块传输到底层的网络连接中，从而可以更快地响应客户端的请求。

需要注意的是，如果使用FlushAfterChunkWriter的话，必须在完成数据发送后调用Close方法，以便将最后的空块传输到底层网络连接中。否则，服务器将无法判断响应已经结束，导致客户端一直等待，直到超时。

总之，FlushAfterChunkWriter是net包中chunked.go文件中的一个结构体，用于将写入的数据分块传输到网络连接中，从而让服务器可以更快地响应客户端的请求。



## Functions:

### NewChunkedReader

NewChunkedReader是一个函数，其作用是将一个io.Reader接口的实现转换为一个支持HTTP传输编码的分块传输格式的读取器。HTTP传输编码是一种在HTTP协议中用于分块传输的编码格式，它使用一组不定长度的块序列来传输数据，每个块都以块的大小和内容开始，以块大小为0的块作为结尾。

具体来说，NewChunkedReader函数接受一个io.Reader类型的参数r，用于向其输入数据。它返回一个*chunkedReader类型的值，该值实现了io.Reader接口和io.ReadCloser接口。该读取器可以从输入io.Reader中读取数据块，并将其作为一个完整的数据块返回，直到块大小为0，表示输入io.Reader已经完成。

此外，NewChunkedReader还对输入的数据进行检查，确保分块传输编码格式的正确性，防止不存在的块或块大小不合法的情况。如果检查失败，则返回一个错误。 

总的来说，NewChunkedReader函数使得我们可以将一个常规的io.Reader接口的实现转换为一个支持HTTP分块传输编码的读取器，方便在进行HTTP数据传输时使用。



### beginChunk

在这个文件中，beginChunk这个函数用于开始一个新的"chunk"，其中一个chunk是一个HTTP消息传输中的一个部分。它会使用格式为"{chunk-size}\r\n"（其中{chunk-size}是当前chunk的大小）的编码方式，告诉接收方正在发送多少个字节。然后，它会开始发送chunk的实际数据。在最后一个chunk后，必须发送一个0大小的chunk，以表示消息的结束。

具体而言，beginChunk函数会先将当前chunk的大小以16进制格式转换为字符串，并存储到一个暂存buf中。然后它会在buf中添加一个CRLF（"\r\n"），表示chunk大小的描述已经结束。最后，beginChunk函数会将buf中的数据写入到Conn中，并清空buf以供后续使用。这样就完成了一个新chunk的发送和描述，接收方就能解析这个chunk的大小并开始接收它的数据了。

总的来说，beginChunk函数是将当前chunk大小编码并发送到接收方，以便接收方能正确地解析和接收chunk的内容。



### chunkHeaderAvailable

chunkHeaderAvailable函数用于检查标头的可用性，以确保能够读取Chunked编码中的下一个块。如果当前缓冲区中的数据足够读取下一个块的标头，则返回true，否则返回false。

该函数首先检查当前缓冲区中的数据长度，如果数据长度小于块的标头大小，则返回false。

如果数据长度足够，该函数会在缓冲区中搜索下一个CRLF序列（"\r\n"），这是标头和块数据之间的分隔符。如果找到了CRLF序列，则返回true，否则返回false。

该函数的作用是确保能够正确读取Chunked编码中的每个块，以便能够按照正确的顺序获取所有的数据。



### Read

Read func在chunked.go文件中的作用是从HTTP响应中读取分块编码的数据，也就是将分块编码的数据流转换为普通的数据流。分块编码是一种HTTP传输协议用于数据编码的方法，它将大的数据流分成多个块，每个块包含数据长度和实际数据。这个编码方式可以在不知道内容长度的情况下发送数据。

Read func使用一个缓冲区并使用递归算法读取HTTP响应，直到读取到一个空的分块为止。每读取一个块，它将块的长度解析出来，然后从数据流中读取相应长度的数据并将其写入缓存中。如果还有剩余长度，则继续读取使用递归函数读取，直到读取到整个数据流的结尾。

在完成读取之后，Read func返回已读取的字节数和一个可能的错误。如果读取成功，则返回的错误为nil，否则返回一个错误。这个错误可能是由服务器发送的数据内容不符合分块编码规范所致。



### readChunkLine

在HTTP传输中，Chunked Transfer Encoding是一种用于传输数据的机制，该机制通过将原始数据分割成一个个称作“chunk”的数据块，并分别发送这些数据块，来实现数据的传输。

Go中的chunked.go文件实现了HTTP Chunked Transfer Encoding的编解码器。readChunkLine函数是其中的一个方法，作用是从输入字节流中读取一个块的大小并返回之。

具体来说，readChunkLine函数的实现如下：

```
func readChunkLine(r *bufio.Reader) (size int, err error) {
	line, err := r.ReadSlice('\n')
	if err != nil {
		return 0, err
	}
	i := bytes.IndexByte(line, ';') 
	if i >= 0 {
		line = line[0:i] 
	}
	line = bytes.TrimSpace(line)
	size, err = strconv.ParseInt(string(line), 16, 64)
	return int(size), err
}

```

它的输入参数是一个bufio.Reader类型的指针r，表示从其中读取一个块的大小。该函数通过调用bufio.Reader的ReadSlice方法从输入流中读取一行字节，并返回包含该行字节的切片line。然后它查找线路中的“;”字符（如果有的话）并截断它，以便获得“chunk-size”。最后，它使用strconv.ParseInt函数将该值转换为int类型，并将其作为函数的返回值返回。



### trimTrailingWhitespace

在Chunked编码中，每个chunked的分块数据之间都需要使用CRLF（换行符）进行分隔。当传输的数据较大时，往往需要将数据切分成多个chunked块进行发送。在每个块之后，可能会存在一些空白字符（空格、制表符等），这些空白字符在Chunked编码中是被视为无效数据的。

trimTrailingWhitespace函数的作用就是剔除每个块末尾的空白字符，以保证Chunked编码的正确性。该函数会从块数据末尾向前遍历，直到遇到非空白字符时停止。然后，它会将末尾的空白字符与前面有效的数据字符分离出来，以便将其包含在下一个chunked块中。

例如，如果一个chunked块的末尾是"abc\t\r\n"，那么trimTrailingWhitespace函数会剔除制表符和换行符，将"abc"与前面的有效数据分离出来。然后，它就可以将"abc"作为一个完整的块发送，而将"\t\r\n"包含在下一个块中。

总之，trimTrailingWhitespace函数可以帮助保证Chunked编码的正确性和传输效率，同时优化网络通信的性能。



### isASCIISpace

isASCIISpace函数是用于判断一个ASCII字符是否为空格或制表符的。在HTTP协议中，由于每行头部信息都以CRLF(Carriage Return and Line Feed)结束，所以需要使用isASCIISpace函数来跳过一行头部信息的空格和制表符。

具体来说，isASCIISpace函数会检查一个字符是否为空格或制表符，如果是，则返回true，否则返回false。使用该函数可以将读取到的数据按行分割，并排除头部信息中的无用部分。



### removeChunkExtension

removeChunkExtension函数的作用是从一个字节数组中去掉chunks扩展名之后返回一个新的字节数组。在HTTP协议中，传输编码为chunked时，每个chunk的前面都会有一个chunk-size的字段指定了该chunk的长度。这个chunk-size字段的格式为16进制字符串，并且以"\r\n"结尾。从字节数组中去掉chunks扩展名之后，就可以直接解析chunk-size字段中的16进制字符串获取chunk的长度。

具体实现过程如下：

首先，removeChunkExtension函数会通过bytes.Index函数查找字节数组中"\r\n"的位置。如果没有找到，则直接返回原始字节数组。如果找到了，则使用bytes.TrimSuffix函数去掉该位置之前的所有字符，包括chunks扩展名和"\r\n"。

然后，removeChunkExtension函数再次调用bytes.Index函数查找字节数组中"\r\n"的位置。如果没有找到，则说明chunks扩展名被去掉后字节数组中只剩下chunk-size字段的16进制字符串。如果找到了，则使用bytes.TrimSuffix函数去掉该位置之后的所有字符，只保留chunk-size字段的16进制字符串。

最后，removeChunkExtension函数使用hex.DecodeString函数将chunk-size字段的16进制字符串转换为一个uint64类型的整数，即chunk的长度。

总结来说，removeChunkExtension函数的作用就是将字节数组中的chunks扩展名去掉，并返回一个新的字节数组和chunk的长度，方便后续处理。



### NewChunkedWriter

NewChunkedWriter函数是net包中的一个函数，用于构造一个写入器，该写入器可以将数据写入底层的io.Writer，并将数据分块（chunk）进行传输。数据分块（chunk）是指一个长度前缀，后跟一个数据块和一个CRLF终止符。

NewChunkedWriter函数的作用是封装底层的io.Writer，将数据写入该底层io.Writer时，根据HTTP分块编码格式将数据进行分块。使用分块传输编码（chunked-encoding）可以在不知道整个消息体长度的情况下，分块传输数据并减少内存使用。

该函数的定义如下：

```go
func NewChunkedWriter(w io.Writer) io.WriteCloser
```

其中，参数w是底层的io.Writer，返回值是一个实现了io.WriteCloser接口的对象。使用该对象进行写操作时，会将数据按照HTTP分块编码进行发送，并在写入完成后关闭底层的io.Writer。

例如，可以使用以下代码创建一个使用chunked encoding的写入器：

```go
package main

import (
    "net/http"
    "net/http/httputil"
    "os"
)

func main() {
    req, _ := http.NewRequest("POST", "http://example.com", nil)
    req.TransferEncoding = []string{"chunked"}

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        println(err.Error())
        return
    }
    defer resp.Body.Close()

    dump, err := httputil.DumpResponse(resp, true)
    if err != nil {
        println(err.Error())
        return
    }
    os.Stdout.Write(dump)
}
```

在上述代码中创建了一个POST请求，并设置TransferEncoding为chunked。这样，在发送请求时就会使用chunked-encoding将请求的消息体作为分块进行传输。因为没有设置消息体，所以在请求中没有数据块被分块传输。请求返回后，使用httputil包中的DumpResponse函数将响应数据打印到控制台上。

需要注意的是，如果底层的io.Writer不支持chunked encoding，则会出现错误。此时需要在构造NewChunkedWriter的时候使用另外一种底层io.Writer来进行替换。



### Write

chunked.go中的Write函数定义如下：

```go
func (cw *chunkedWriter) Write(p []byte) (n int, err error)
```

这个函数的作用是将字节切片p写入到chunkedWriter中，并使用chunked编码格式将数据分块。

chunked编码是一种HTTP传输编码，用于在HTTP发送消息主体时动态生成消息长度。在chunked编码中，消息主体被分割成一系列使用长度前缀标记的块。每个块包含一个描述其长度的十六进制数的标头，然后是块本身，后跟一个空行作为块结束符。

在Write函数中，将要写入的字节切片p被分成多个块，每个块的最大长度为chunkSize。然后，在块开始之前写入十六进制数值的块长度。最后，在写入完整个消息主体后，写入一个零长度的块，表示消息主体的结束。

Write函数返回已写入的字节数和任何错误。如果写入的字节数不是块长度的倍数，则返回错误io.ErrShortWrite。如果chunkedWriter已被关闭，则Write函数将返回错误io.ErrClosedPipe。



### Close

chunked.go文件中的Close函数是为了关闭chunked编码的流而设计的。HTTP协议使用chunked编码来传输可变长度的数据。在chunked编码中，响应体被分为多个块，每个块都有一个指定的大小。在传输数据时，响应体的大小不需要预先知道，块的大小可以在发送时动态确定。

Close函数的主要作用是在响应体的最后添加一个空块，并在传输完成后关闭连接。这样可以确保所有的数据都已经传输完毕，并且通知客户端已经完成传输。

函数中的代码逻辑如下：

1. 如果当前块的大小大于0，则添加一个空的块以结束流。
2. 如果当前chunkedReader的parent是一个io.Closer，则关闭这个parent。
3. 将chunkedReader的父对象设置为nil，以确保在关闭多次调用后不会发生死锁。
4. 返回一个nil的error，表明关闭操作已经完成。

总的来说，Close函数是为了确保chunked编码流的正确关闭，而不会导致数据丢失或死锁。它是http协议中实现数据传输的重要组成部分。



### parseHexUint

在HTTP协议中，数据传输是通过一个称为“传输编码”的机制进行的。其中，Chunked编码是一种流式传输的机制，用于通过在每个块的结束处添加一个带有块的长度的十六进制字串来对数据进行分块。从而，接收方能够知道每个块的长度并将其作为单独的数据块读取。

在net包中的chunked.go文件中，parseHexUint函数的作用是解析16进制数，将其转换为十进制数。chunked.go文件是HTTP响应分块解码器的实现，该函数用于解析传输编码中用于描述数据块长度的十六进制字串。

该函数的具体实现如下：

```
func parseHexUint(text string) (n uint64, err error) {
    for _, c := range text {
        switch {
        case '0' <= c && c <= '9':
            n = n*16 + uint64(c-'0')
        case 'a' <= c && c <= 'f':
            n = n*16 + uint64(c-'a'+10)
        case 'A' <= c && c <= 'F':
            n = n*16 + uint64(c-'A'+10)
        default:
            return 0, &strconv.NumError{Func: "ParseUint", Num: text, Err: strconv.ErrSyntax}
        }
    }
    return n, nil
}
```

解析过程中，函数遍历传入的字串中的每个字符，根据其范围减去相应的ascii码，实现了将16进制数转为10进制数。如果遇到非法字符，则返回相应的错误信息。



