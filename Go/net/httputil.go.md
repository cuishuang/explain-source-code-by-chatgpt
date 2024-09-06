# File: httputil.go

httputil.go文件是Go语言标准库中net包中的一个子包，主要是提供HTTP相关的工具函数和结构体，如ReverseProxy、DumpRequest、DumpResponse等。具体作用如下：

1. ReverseProxy: 可以将传入的HTTP请求转发到另一个服务器上，它提供了代理功能，通过修改请求和响应的头、路径等来达到代理的效果。

2. DumpRequest: 可以将HTTP请求的信息转储成字符串形式，可以用来调试和分析请求。

3. DumpResponse: 可以将HTTP响应的信息转储成字符串形式，可以用来调试和分析响应。

4. NewFileTransport: 用于以文件系统为基础的Transport，可以通过该函数创建一个自定义的Transport，实现根据文件系统中资源路径返回响应数据的功能。

5. NewSingleHostReverseProxy: 它是一个非常常用的函数，可以创建一个反向代理，并且将传入的请求转换为目标url并将其发送到目标服务器。这个函数是ReverseProxy函数的一种快捷方式。

总之，httputil.go文件提供了HTTP相关的一些实用工具函数和结构体，使用这些函数和结构体可以简化HTTP开发过程中很多繁琐的任务。




---

### Var:

### ErrLineTooLong

ErrLineTooLong是一个错误类型变量，表示HTTP请求头或响应头的一行（例如请求行、响应行、请求头、响应头）长度超过了允许的最大长度。

在HTTP协议中，一行的最大长度限制为8192个字节。如果HTTP请求头或响应头中的一行长度超过了这个限制，则会抛出ErrLineTooLong错误。此错误通常用于处理HTTP请求或响应头时的数据验证和错误处理。

使用ErrLineTooLong类型错误变量可以提高程序的安全性和稳定性，因为它可以帮助程序及时发现并处理HTTP请求或响应头中的异常情况，最大限度地避免程序崩溃或被攻击。



## Functions:

### NewChunkedReader

NewChunkedReader函数的作用是创建一个读取HTTP传输编码为“chunked”的响应体的新读取器。“chunked”编码是一种HTTP传输编码，用于在HTTP协议的响应体中传递序列化数据。该编码将数据分成多个块，每个块的大小是在块的开头指定并以CRLF（回车符和换行符）结尾。

该函数返回的读取器可以解析“chunked”编码并提供有序的、原始的数据块。该函数还能够检测不正确的编码并在读取响应体时返回错误。

具体来说，NewChunkedReader函数接受一个io.Reader类型的参数，即HTTP响应体的原始数据读取器。它将返回一个新的*bufio.Reader类型的值，该值实现了io.Reader接口并提供了一些辅助解析“chunked”编码的方法。在读取器被使用前，需要确保读取器返回的最终结果是“chunked”编码的响应体。



### NewChunkedWriter

NewChunkedWriter函数是一个帮助函数，用于创建一个可以对数据进行分块编码的写入器（Chunked-Encoding）。Chunked-Encoding是HTTP/1.1规范中定义的一种传输编码方式，用于将数据划分为一系列大小不等的块，并且每个块都以一行十六进制数字的长度头开始。最后一个块是一个长度为0的块，标记着消息的结束。

NewChunkedWriter函数会返回一个ChunkedWriter类型的值，我们可以使用该值来将数据写入到一个io.Writer接口中，并使用Chunked-Encoding方式进行编码。对于发送大量数据的长时间接口，例如实时的网络流媒体，Chunked-Encoding能够很好地发送数据，并且可以避免在发送到底层TCP连接时，将所有数据都发送到底层缓冲区中时出现的问题。通过NewChunkedWriter创建Chunked-Encoding的写入器，可以让我们轻松地使用HTTP/1.1里面的chunked编码方式，从而更好地控制数据流的处理。



