# File: http.go

http.go是net包中的一个子包，它的作用是提供HTTP协议的实现。HTTP是一种应用层协议，它允许客户端和服务器之间进行交互，并传输数据。

http.go实现了请求和响应的解析和生成，包括从URL中解析出请求的目标主机和路径，对请求进行方法验证、请求头解析、请求体解析、响应头的生成和响应体的生成，并提供了一些HTTP协议中常用的功能，如cookie、gzip压缩等。

在http.go中，定义了多个类型，其中包括：

- Client：HTTP客户端。
- Server：HTTP服务器。
- Request：HTTP请求实体。
- Response：HTTP响应实体。
- Header：HTTP请求和响应头。

除此之外，http.go还提供了一些常用的函数，如Get、Post、PostForm、HandleFunc等，一些处理请求和响应的函数，如Serve、ServeHTTP等。

总的来说，http.go是net包中HTTP协议的具体实现，它提供了HTTP客户端和HTTP服务器的实现，以及与HTTP协议有关显著的函数和类型，这些都为使用HTTP协议的开发者提供了很大的方便。




---

### Var:

### aLongTimeAgo

在go/src/net中http.go文件中，aLongTimeAgo这个变量是一个time.Time类型的变量，用于表示一个很久以前的时间。

具体来说，aLongTimeAgo这个变量的初值被设置为“Mon, 01 Jan 1950 00:00:00 GMT”，也就是“1950年1月1日 00:00:00 格林威治标准时间”，也是互联网第一个标准时间戳。

在http.go文件中，aLongTimeAgo这个变量主要用于HTTP协议中的时间处理。例如，在HTTP响应头中，HTTP协议要求包含“Date”字段，表示响应生成时间。如果响应头中没有Date字段，就需要使用aLongTimeAgo变量作为默认的时间戳。

另外，在HTTP请求处理中，当客户端发起请求时，如果请求头中没有“If-Modified-Since”字段，服务器也会使用aLongTimeAgo变量作为默认的修改时间。

因此，可以说aLongTimeAgo这个变量是HTTP协议中非常重要的一个变量，在时间处理方面起到了重要的作用。



### omitBundledHTTP2

在Go语言的net包中的http.go文件中，omitBundledHTTP2是一个布尔型变量，其默认值为false。如果设置为true，则会禁用内置的HTTP/2支持，强制使用外部的HTTP/2库来处理HTTP/2流量。

HTTP/2是一种新的协议，旨在改进Web应用程序的性能和安全性。作为Go语言的标准库，net/http包提供了HTTP服务的核心部分，而HTTP/2则是其中的一种特性。

然而，内置的HTTP/2支持可能会引入一些不兼容性问题，而且在某些情况下它可能无法正常工作。因此，如果遇到了HTTP/2相关的问题，可以将omitBundledHTTP2设置为true，强制使用外部的HTTP/2库来处理流量。

需要注意的是，如果使用外部的HTTP/2库，则需要手动导入和初始化该库，以及提供相应的配置。同时，该选项可能会导致性能上的损失，因为外部库可能会采用不同的优化策略或实现方式。



### NoBody

NoBody 是一个特殊的 ReadCloser 类型的变量，在 HTTP 请求和响应中用于表示空的消息体。

HTTP 协议中有时需要发送一个请求或响应的头部和状态信息，但不需要传递任何消息的正文部分。这种情况下，可以使用 NoBody 变量来表示一个空的正文部分。使用 NoBody 变量可以避免显式地创建一个空的正文部分，减少内存分配和复制量。

例如，在 HTTP GET 请求中，消息体应该是空的。这时可以设置 Request.Body 为 NoBody 变量。在 HTTP 响应中，如果响应状态码为 204 No Content 或 304 Not Modified，则响应正文部分也应该是空的。这时可以设置 Response.Body 为 NoBody 变量。

NoBody 也实现了 io.Reader 和 io.Closer 接口，使得它可以像其他消息正文一样被处理。当调用 Read 时，将返回一个空的字节数组，表示读取完所有数据。当调用 Close 时，将不执行任何操作。



### _

_ 这个变量在 Go 语言中是一个特殊的变量名，在不需要使用这个变量的情况下，可以作为占位符使用。在 net/http 包中，_ 这个变量被使用在两个地方。

1. 在 http 协议的定义中，使用 _ 变量作为占位符传递函数参数

在 net/http/http.go 文件中，定义了一个 http 协议的结构体。在这个结构体中的函数定义中，会使用 _ 变量作为占位符传递参数。

例如，在 ServerHTTP 函数中，我们可以看到代码：

```go
func (hs *httpServer) ServeHTTP(w ResponseWriter, r *Request) {
    handler := hs.Handler
    if handler == nil {
        handler = NotFoundHandler()
    }
    if r.Method == "CONNECT" {
        handler = tunnelHandler{}
    }
    handler.ServeHTTP(w, r) // use unexported ServeHTTP to avoid
}
```

在上述代码中，ServeHTTP 函数的参数列表中，使用 _ 变量作为占位符传递参数。这意味着，我们可以传递任何类型的参数，但是这些参数不会被使用到。

2. 在 http 包的初始化函数中，_ 变量用于导入子包

在 net/http 包中，有一个未导出的 init 函数，用于初始化包内变量和函数。在这个函数中，会使用 _ 变量用于导入 net/http 子包和其他一些子包。

例如，在 http.go 文件中，有一个 init 函数：

```go
func init() {
    // ......
    _ = ProtocolError{} // so that code that uses http.ProtocolError still works
    _ = http.DefaultClient // so that the Transport type is in the symbol table
    _ = mimeGlobal // eagerly parse mime types
    _ = goUnusedProtection__ // http2error.go, to silence import unused
    // ......
}
```

在上述代码中，我们可以看到，某些子包被导入并赋值给 _ 变量，这些子包可以使用包路径调用。使用 _ 变量的原因是，如果这些子包不被使用到，可以防止 Go 编译器报出未使用导入错误。



### _

在 `go/src/net/http/http.go` 中有一个“_”变量，它的作用是导入一个包但不使用它。这通常用于在导入包时执行其初始化功能。

例如，在 `http.go` 文件中，通过导入 `time` 包的 `_` 标识符来调用其 `init()` 函数，这个函数会初始化 Time 字符串。而这个标识符的放置位置不确定，可以在任何语句都可以导入这个包，只要这个标识符不会被引用到。

值得注意的是，`_` 标识符在Go语言中只是一个特殊的标识符，并没有任何实际的意义，因此还可以用其他的标识符代替。但是，由于 `_` 是一个惯用标识符，因此它在Go语言中经常用于这种情况下。






---

### Structs:

### incomparable

在go/src/net/http.go中，incomparable结构体是一个未导出的辅助结构体，它用于在HTTP请求中跟踪请求体的长度和类型等信息。

具体来说，incomparable结构体包含了以下几个字段:

- length int64：表示请求体的长度，以字节为单位。如果请求体长度未知或不可计算，则值为-1。
- fragmented bool：表示请求体是否分段传输。如果请求体被分成多个部分分别发送，则值为true；否则为false。
- mediaType string：表示请求体的MIME类型，例如"text/plain"或"application/json"等。如果不知道请求体的MIME类型，则值为""。

在处理HTTP请求时，我们通常需要使用这些信息来读取并解析请求体数据。通过使用incomparable结构体，我们可以更方便地获取和管理这些信息，从而保证请求体传输的安全、完整和准确。



### contextKey

在Go语言中，http请求中的上下文（Context）扮演了非常重要的角色。而contextKey结构体则是用于定义Context中key的类型的结构体。

具体来说，Context是Go语言中的一个非常轻量级的线程安全的上下文容器，用于在不同的Goroutine之间传递数据和请求的上下文信息。而contextKey则定义了Context中的键的类型，并且可以用于获取或者设置某个上下文信息。

例如，在http请求处理函数中，可以通过Context来传递请求的上下文信息，比如说请求的Method、Headers、URL等信息。而对于contextKey结构体，通常是定义成某个常量，作为Context中键的标识，方便在不同的地方使用。

例如，Go语言的标准库中就定义了很多contextKey结构体，比如说：

- MethodCtxKey：表示请求的Method信息
- URLCtxKey：表示请求的URL信息
- HeaderCtxKey：表示请求的Headers信息

这些结构体本身没有太多的逻辑，仅仅是用于定义Context中的键的类型，并且让编程者可以使用它来获取或者设置某个上下文信息。同时，这些contextKey结构体的值（通常是常量）也可以在不同的地方使用，以保证代码的可读性和一致性。



### noBody

结构体noBody定义在net/http包下的http.go文件中，它实现了io.Reader和io.Closer接口。该结构体的作用是用于请求中不带请求体时的空请求体，用来代替nil。在HTTP协议中，有些方法如GET、DELETE等是没有请求体的，此时使用noBody结构体来代替nil，避免出现错误。

noBody结构体的定义如下：

type noBody struct{}

func (_ noBody) Read([]byte) (int, error) {
    return 0, io.EOF
}

func (_ noBody) Close() error {
    return nil
}

在noBody结构体的Read方法中，每次读取请求体时，返回0和io.EOF，表示没有请求体可读。在Close方法中，返回nil，表示空的请求体已经关闭。这种空的请求体通常是一个代表没有内容的空字符，有时也称为"payload"。

使用noBody结构体有助于区分有body和没有body的请求，并提高HTTP请求的可读性和可维护性。



### PushOptions

PushOptions是一个结构体，用于在处理HTTP/2 Push请求时指定一些选项。 当客户端使用HTTP/2协议与服务器进行通信时，服务器可以提前向客户端推送一些与请求相关的资源，例如css，js，图片等。 这样可以加快页面加载速度，提高用户体验。

PushOptions结构体具有以下字段：

- Method string：推送资源的HTTP方法，默认为GET
- Header http.Header：推送资源的HTTP标头
- Pusher *Pusher：推送HTTP/2流
- 推送范围Range：推送资源的字节范围，默认为整个文档
- 紧急程度Priority：推送资源的优先顺序

使用PushOptions结构体，可以指定推送资源的HTTP方法，标头，范围等信息，并控制推送资源的优先级。 这可以帮助提高页面速度和性能，从而改善用户体验。



### Pusher

Pusher结构体是一个接口，它用于在HTTP请求处理过程中推送服务器端资源给客户端。 

在HTTP请求响应过程中，Pusher结构体提供了一个方法Push，可以将资源封装并发送到客户端，客户端会即时接收到该资源，这种机制促使了Web应用程序更快、更高效地提供响应。 

Pusher结构体是在Server.ServeHTTP方法中被调用。如果HTTP服务器检测到请求是否可以推送服务器端资源，它将返回一个类型为http.Pusher的请求，然后调用Pusher结构体的Push方法以将资源推送给客户端。 

需要注意的是，Pusher结构体是一个很灵活的接口，它适用于多种不同类型的资源，但是HTTP/2服务器可以选择忽略Pusher结构体推送资源的请求，所以客户端不应该依赖于服务器端推送资源。



## Functions:

### String

http.go文件中的String()函数是用于将HTTP请求或响应的数据转换为字符串的函数。它是在http.Request和http.Response结构体上定义的。

该函数返回HTTP请求或响应的字符串表示形式，其中包括HTTP方法、URL，协议版本和HTTP头信息。例如，在HTTP请求中，字符串可能如下所示：

```
GET /index.html HTTP/1.1
Host: example.com
User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
```

这个字符串显示了HTTP请求使用的方法（GET）、请求的URL（/index.html），HTTP协议版本（HTTP/1.1）和HTTP头信息（Host、User Agent等）。

在实际开发中，String()函数通常用于调试和日志记录。例如，在调试过程中，我们可以使用String()函数来输出HTTP的请求和响应内容，以便于确定问题所在。

总之，String()函数是http.go文件中重要的一个功能，它提供了一个方便的方式来查看HTTP请求和响应的内容以及头信息，提高了HTTP协议的可读性和可维护性。



### hasPort

在Go语言的net包中，http.go文件中的hasPort函数用于检查一个网络地址字符串是否包含端口号。如果地址字符串中包含“:”字符，则认为该字符串具有端口号。

函数定义如下：

```go
func hasPort(s string) bool {
    return strings.LastIndexByte(s, ':') > strings.LastIndexByte(s, ']')
}
```

该函数使用了Go字符串库中的LastIndexByte函数，通过查找字符串中指定字节的最后一个出现位置来确定是否包含端口号。如果“:”字符的最后一个出现位置在“]”字符的最后一个出现位置之后，则认为该地址字符串不包含端口号。

在HTTP请求中，当地址字符串中不含有端口号时，默认使用80端口（HTTP协议默认端口）或443端口（HTTPS协议默认端口）。如果包含端口号，则使用指定的端口号进行访问。

因此，hasPort函数是一个非常重要的工具函数，用于帮助HTTP客户端和服务器分析和处理网络地址字符串，以确保能够正确地使用指定的端口号进行通信。



### removeEmptyPort

removeEmptyPort是http.go文件中的一个函数，主要作用是从给定的地址中移除空端口号。它通常在HTTP客户端和服务器中用于处理地址，确保一个有效的端口号。

HTTP地址通常具有形式http://host:port/path，其中端口号是可选的。如果地址中没有指定端口号，则使用默认端口号，对于HTTP协议是80。

removeEmptyPort函数的入参是一个字符串类型的地址，它首先通过调用net.SplitHostPort函数将地址分割为主机名和端口号。如果主机名中包含IPv6地址，那么SplitHostPort函数还会去除方括号。接下来，如果端口号为空或为"0"，则会将其移除，否则保留端口号。

例如，如果输入的地址是 "http://www.example.org:8080/path"，那么removeEmptyPort函数会返回 "http://www.example.org:8080/path"，因为端口号不为空。而如果输入的地址是 "http://www.example.org/path"，那么removeEmptyPort函数会返回 "http://www.example.org/path"，因为使用默认端口号80。



### isNotToken

isNotToken是一个辅助函数，用于检查HTTP头部字段是否符合标准的Token格式。

在HTTP协议中，Token是由一系列ASCII字符构成的非空字符串，以空格分隔。例如，HTTP头部字段"User-Agent"的值可以是"Mozilla/5.0"，它符合Token格式。而"User-Agent: Mozilla/5.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; MS-RTC EA 2"则不符合Token格式，因为其中包含非Token字符";"和空格。

isNotToken函数会遍历头部值的每一个字符，如果发现不属于ASCII字符集或者是非Token字符，就返回true，表示头部值不符合Token格式。否则返回false，表示头部值符合Token格式。

在http.go中，isNotToken函数主要用于解析HTTP头部字段。在解析器中，当解析到头部字段时，会调用isNotToken函数检查头部值是否符合Token格式。如果不符合，解析器会抛出错误，表示HTTP消息格式错误。如果符合，解析器会将头部字段名和头部值存储到HTTP请求/响应结构体中，继续解析下一个头部字段。



### stringContainsCTLByte

stringContainsCTLByte函数用于检查字符串中是否包含控制字符（Control characters）或无效字符（Invalid characters）。在HTTP协议中，控制字符和无效字符会引起一些安全问题，比如HTTP协议注入攻击，因此需要对它们进行检查和过滤，以防攻击。

该函数检查字符串中每个字符是否在0到31的ASCII码范围内或128到255的ASCII码范围内，如果包含则返回true，否则返回false。在HTTP协议中，通常使用该函数来检查URL的合法性（是否包含不允许的字符），以及HTTP头部中的字段名称和字段值是否包含不允许的字符。如果包含不允许的字符，就需要进行过滤或拒绝请求。

例如，检查URL是否包含不允许的字符的代码片段：

```
func checkURL(url string) error {
    if stringContainsCTLByte([]byte(url)) {
        return errors.New("URL contains control characters")
    }
    // other checks
    return nil
}
```

该函数接受一个字节数组作为参数，而不是字符串，这是因为在Go语言中，字符串是不可变的，而字节数组可以进行修改，更方便进行检查和过滤。



### hexEscapeNonASCII

hexEscapeNonASCII是一个在net/http包中定义的函数，其作用是将非ASCII字符转换为对应的hex编码。在http传输中，只有ASCII字符是安全的，因此需要将非ASCII字符进行编码，以便在http传输中传输。

具体来说，hexEscapeNonASCII函数将字符串中的每个非ASCII字符替换为对应的hex编码，即用"%xx"的形式表示。例如，字符'冯'在hex编码下表示为E5 86 AF，因此hexEscapeNonASCII函数会将其替换为"%E5%86%AF"，这样就可以在http传输中传递了。

此外，hexEscapeNonASCII函数还会对空格、加号、等号等特殊字符进行编码，以确保它们在http传输中不会被误解。例如，在http传输中，空格应该被转换为"%20"，加号应该被转换为"%2B"。

总之，hexEscapeNonASCII函数的作用是将一个字符串进行编码，使其可以在http传输中安全地传递。



### Read

http.go 文件中的 Read() 函数是实现 io.Reader 接口的一个方法，主要负责读取连接中的数据并将其填充到给定的字节切片中。

具体来说，当 Read() 函数被调用时，它首先会检查是否存在尚未读取的缓冲区数据，如果缓冲区存在数据，就会直接从缓冲区读取并返回数据。否则，Read() 函数会调用 underlying connection 上的读取方法（例如 net.Conn.Read()）来读取更多的数据，并将其存储到为 Reads 调用分配的临时缓冲区中。最后，Read() 函数将从临时缓冲区填充足够的字节数到给定的字节切片中，并返回写入数据的字节数。

这个方法的作用非常重要，它使得 HTTP 服务器能够从客户端读取请求数据，并将其转换为 HTTP 请求消息。而 HTTP 客户端也可以使用这个方法读取来自服务器的 HTTP 响应。因此，Read() 函数是 HTTP 协议实现的重要组成部分。



### Close

http.go文件中的Close函数主要用于关闭HTTP连接。

具体来说，Close函数会关闭当前HTTP连接并将其相关的资源释放。在HTTP请求结束后，应该及时调用Close函数，以避免一些资源占用过长时间而导致的问题。

同时，Close函数还可以允许用户指定一个返回error类型的回调函数，以在关闭连接出现问题时进行处理。

总之，Close函数是一个非常重要的函数，它可以保证HTTP连接的安全关闭和资源释放，避免因此带来的一些不必要的问题。



### WriteTo

WriteTo是net/http包中的一个函数，用于将HTTP响应内容写入到一个io.Writer接口中。具体来说，它会将HTTP响应的Header和Body部分分别写入到给定的io.Writer接口中。

例如，假设我们有一个HTTP响应resp和一个文件f，我们可以使用WriteTo函数将响应写入到文件中：

```
resp.WriteTo(f)
```

该函数的签名如下：

```
func (resp *Response) WriteTo(w io.Writer) (n int64, err error)
```

其中，resp表示要写入的HTTP响应，w表示要写入的io.Writer接口。函数会返回写入的字节数以及可能出现的错误。

WriteTo函数的作用在于，将HTTP响应写入到一个io.Writer接口中，这个接口可以是文件、内存缓冲区等等。这样我们可以方便地将HTTP响应保存到本地文件或者处理其它相关操作。



