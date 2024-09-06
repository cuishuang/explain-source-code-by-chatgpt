# File: recorder.go

recorder.go这个文件的作用是记录网络请求和响应的信息，包括请求时间、响应时间、请求信息、响应信息等。这个文件中的Recorder结构体实现了http.RoundTripper接口，可以被用作http.Client中的Transport。

当使用http.Client发送网络请求时，请求会被送到http.RoundTripper接口的实现中处理，Recorder就是其中一种可能的实现。在请求被处理时，Recorder会记录请求的信息，同时调用下一个RoundTripper进行后续处理。当响应返回时，Recorder会记录响应的信息，并将响应传递给调用者。

通过使用Recorder，我们可以方便地记录网络请求和响应的信息，用于调试、性能优化等目的。在测试时，我们也可以使用Recorder来模拟网络请求和响应，以便于编写测试用例。




---

### Structs:

### ResponseRecorder

在 Go 语言的 net 包中，recorder.go 文件中的 ResponseRecorder 结构体是一个类似于 http.ResponseWriter 的记录器。

ResponseRecorder 可以用于记录处理请求时的响应，它保存了请求响应的状态码、响应头和响应数据等信息。另外，ResponseRecorder 还提供了一些方法可以让开发者方便地获取响应值。

在测试中，我们可以使用 ResponseRecorder 来代替真实的响应，用于模拟响应过程并方便地检查响应结果。

ResponseRecorder 结构体的主要字段如下：

```go
type ResponseRecorder struct {
    Code       int           // 响应状态码
    HeadersMap map[string][]string  // 响应头
    Body       *bytes.Buffer // 响应体
    Flushed    bool          // 响应是否被刷新过
}
```

ResponseRecorder 结构体的主要方法如下：

1. Header()：返回 ResponseRecorder 中的响应头。
2. Write()：将响应数据写入 ResponseRecorder 中的响应体。
3. WriteHeader()：写入响应状态码。

通过 ResponseRecorder 我们可以轻松地模拟 HTTP 请求并获取响应结果，这对于测试和调试是一个非常方便的工具。



## Functions:

### NewRecorder

NewRecorder函数是一个工厂函数，用于创建一个Record结构体的实例，并将与传入的 ResponseWriter 和 Request 相关的数据存储在Record结构体中。Record结构体是一个实现了net/http/httptest包中的ResponseRecorder接口的结构体。它保存了响应的状态、头部和内容，并且它的Write()方法存储了发送给客户端的数据。

NewRecorder函数的定义如下：

```
func NewRecorder() *Record {
    return &Record{
        HeaderMap: make(http.Header),
    }
}
```

它不需要传入参数，直接返回一个指向Record结构体的指针。 Record结构体包含几个成员变量，如下所示：

```
type Record struct {
    Writer      http.ResponseWriter
    Code        int
    HeaderMap   http.Header
    Body        bytes.Buffer
    Flushed     bool
    ContentLen  int
    Request     *http.Request
}
```

- Writer：用于存储内部使用的ResponseWriter。
- Code：指示响应的 HTTP 状态码。
- HeaderMap：用于存储响应头信息的map。
- Body：用于存储响应主体的缓冲区。
- Flushed：用于标识响应是否已经被发送到客户端。
- ContentLen：指示响应主体的长度。
- Request：指向客户端发送请求时传入的 Request 指针。

NewRecorder函数创建一个Record结构体的实例，并将其初始化为默认值。在使用NewRecorder函数时，我们往往会将它作为参数传递到http.HandlerFunc或http.Handler的实现中，以记录对请求的处理结果，并返回记录的响应内容供测试使用。



### Header

Header这个func的作用是返回Record类型值的headers。其中，Header函数的签名如下：

```
func (r *Record) Header() Header
```

其实现方式是通过从Record类型指针获取一个unwrap专用请求，然后通过unwrap请求的Header方法返回可以复制的headers。

具体来说，Header这个func返回的是headers类型的值。headers类型是一个map[string][]string类型的值，保存了Record中的所有headers。在HTTP传输过程中，header用于向接收方提供请求的元数据信息，如Content-Type、Content-Length、Accept等。返回headers方便对请求的headers进行处理和操作。



### writeHeader

在net/recorder.go文件中，writeHeader函数是一个用于写HTTP响应头的函数。该函数的作用是将status、header和w（writer）写入到记录器中，以便后续分析和调试。

具体来说，writeHeader函数的主要功能如下：

1. 通过调用http.Response对象的Write函数，将http响应头写入到记录器的缓冲区中。

2. 设置记录器的status属性为HTTP响应状态。

3. 如果header参数不为nil，则将header写入到记录器的Header属性中，否则将默认的http头写入到记录器中。

通过调用writeHeader函数，我们可以捕获和记录所有请求和响应的数据，这对于调试和分析应用程序的性能和行为非常有用。同时，writeHeader函数还为我们提供了一个简单而强大的HTTP请求/响应记录器，使我们能够更好地理解和优化Web应用程序。



### Write

Write函数是一个实现了io.Writer接口的方法，它的作用是将数据写入到recorder缓冲区中。recorder是一个记录器，它会记录下写入到它缓冲区的所有数据。它通常被用于测试的场景中，比如用来捕捉HTTP请求和响应的流量。

具体来说，Write函数接收一个字节切片作为数据，然后将这个数据写入到recorder的缓冲区中。写入完成后，函数会返回写入的字节数和nil作为error。如果写入过程中出现错误，比如缓冲区已满，函数会返回已写入的字节数和相应的error信息。

在测试中，我们可以将recorder作为HTTP请求的目标，然后对发出的请求进行记录。我们还可以使用recorder来测试我们的代码是否正确的解析了HTTP响应。它还可以用来模拟网络连接，使得我们的测试更加稳定和可靠。



### WriteString

WriteString函数是net/recorder包中的一个方法，其作用是将字符串写入缓存器中。该函数的定义如下：

```go
func (r *Recorder) WriteString(s string) (int, error) {
    return r.Write([]byte(s))
}
```

该函数接收一个字符串参数s，该参数将被转换为字节数组，并使用Write方法将其写入缓冲区。然后返回写入的字节数以及是否发生了错误。

更具体地说，WriteString函数是Record类型的一个方法，Record用于记录传入和传出的流量。WriteString函数是向Record的输出缓冲区中写入字符串的快捷方式。WriteString操作可以用于记录服务器响应、读取网页内容等。

举一个具体的例子，假设我们希望在客户端和服务器通信时记录数据流，我们可以使用recorder包中的Recorder类型，创建一个记录器，然后使用WriteString方法将字符串写入缓存区：

```go
import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "net/http/httputil"
)

func main() {
    recorder := httptest.NewRecorder()
    handler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, World!")
    }
    req := httptest.NewRequest(http.MethodGet, "/do/something", nil)
    handler(recorder, req)
    dump, _ := httputil.DumpRequestOut(req, true)
    fmt.Println(string(dump))
    fmt.Println("Response:", recorder.Body.String())
}
```

在这个例子中，我们创建了一个记录器recorder和一个处理器handler，并调用handler来处理请求。然后使用WriteString方法将字符串“Hello,World!”写入缓存区。最后，我们使用DumpRequestOut将请求输出为可读的格式，并打印响应主体的内容。



### checkWriteHeaderCode

checkWriteHeaderCode是go/src/net/recorder.go文件中的一个函数，主要用于验证写入的header的合法性。

在HTTP通信中，请求和响应都需要发送一个header，这个header包含了关于请求或响应的一些元数据信息，如Content-Type、Content-Length等。其中，HTTP头是由键值对组成的，键和值之间用冒号分隔，并以一个空白行作为结束。

checkWriteHeaderCode函数的作用是跟踪写入Header的代码，并检查写入的Header键是否符合HTTP协议中定义的规范。

该函数实际上是一个钩子函数，在Record类型的Writer中的Write函数中被调用。

函数的参数包括了已写入Header的键，“夹具”中的状态以及状态的改变量。它会根据参数计算出新状态代码，并判断状态代码是否符合HTTP协议的规范。如果不是，则会panic。

下面是checkWriteHeaderCode函数的代码：

```
func checkWriteHeaderCode(h int, f *recordFrame, status int) {
    if !isHeadAuthorizedForStatus(h, status) {
        panic(fmt.Sprintf("invalid header write of %q for status code %v",
            httpResponseHeader[h], status))
    }
    f.WriteHeader(h)
}
```

可以看到，该函数通过调用isHeadAuthorizedForStatus函数来判断当前状态下，写入Header的键是否合法。如果不合法，则会抛出一个panic异常。

总的来说，checkWriteHeaderCode函数是一个用于检查HTTP协议合法性的辅助函数。确保写入的Header符合HTTP协议的格式和规范。



### WriteHeader

在 Go 的 net 包中，recorder.go 文件中的 WriteHeader 函数被用来向 HTTP 响应中写入状态码和头部信息。具体来说，它的作用如下：

1. 在 HTTP 响应中设置状态码：WriteHeader 可以通过传入一个整数类型的参数来设置 HTTP 响应的状态码。

2. 在 HTTP 响应中设置头部信息：WriteHeader 可以通过传入一个 http.Header 类型的参数来设置 HTTP 响应的头部信息。

3. 记录 HTTP 响应的状态码：WriteHeader 还可以记录 HTTP 响应的状态码，以便稍后在代码中进行访问和分析。

例如，当我们使用 net 包中的 ResponseRecorder 对象来捕获 HTTP 响应时，我们可以使用 WriteHeader 函数来写入状态码和头部信息，然后使用该对象的 Result() 方法来访问、分析或处理该响应。具体可以参考以下示例代码：

```
import (
    "net/http"
    "net/http/httptest"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
}

func TestHandleRequest(t *testing.T) {
    req, err := http.NewRequest("GET", "/test", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handleRequest)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := `application/json`
    if rr.Header().Get("Content-Type") != expected {
        t.Errorf("handler returned wrong content type: got %v want %v",
            rr.Header().Get("Content-Type"), expected)
    }
}
```

在此示例中，我们使用 ResponseRecorder 对象（rr）来捕获处理请求的 HTTP 响应。在 handleRequest 函数中，我们使用 WriteHeader 函数来设置 HTTP 响应的状态码和头部信息。最后，我们使用 ResponseRecorder 对象的 Code() 和 Header().Get() 方法来访问响应的状态码和头部信息，以进行测试和验证。



### Flush

recorder.go文件中的Flush函数是用于将当前缓存中的数据写入到底层的io.Writer中。具体来说，它的作用是：

1. 将当前缓存中的数据写入到io.Writer中。即使当前缓存没有达到缓存边界也可以使用Flush方法，因为它可以确保当前缓存中的所有数据都被写入底层Writer中。

2. 如果Flush过程中出现错误，它会将这个错误返回给调用者。这个错误可以是任何类型的错误，例如底层Writer返回的错误或者其他类型的错误，所以在使用Flush方法时应该检查返回值以确定是否有错误发生。

3. Flush方法是可以并发安全的，因为在调用Flush方法时底层Writer是不会被并发访问的。

总之，Flush方法的作用是将当前缓存中的数据写入底层io.Writer中，并且它是并发安全的，可以在任何时候调用。



### Result

在go/src/net/recorder.go中，Result（）函数的作用是返回HTTP响应的记录。

具体来说，Result（）函数是Recorder结构体的一个方法，该结构体保存了对http.Handler接口调用的记录。在记录HTTP响应之后，可以通过调用Result函数来获取响应。Result()函数返回两个值，分别是*http.Response和error。Response是HTTP响应，error是可能发生的错误。

在测试代码中，可以使用Recorder来模拟HTTP请求和响应，对应用程序进行测试。在测试期间，通过记录HTTP响应，我们可以检查是否按预期为每个请求设置了响应。

总之，Result()函数是在HTTP测试时捕获记录的机制，用于获取测试响应。



### parseContentLength

parseContentLength是一个用于解析HTTP响应中Content-Length头部的函数。它的作用是从HTTP响应头部中获取Content-Length值，并将其转换为int类型。

具体来说，parseContentLength函数接受一个字符串类型的Content-Length头部值作为参数，并使用strconv.Atoi函数将该值转换为int类型。如果转换成功，函数返回一个int类型的值；否则返回-1。

在recorder.go这个文件中，parseContentLength函数经常被用于解析HTTP响应体的长度。这个函数可以帮助我们更好地理解HTTP响应头部中Content-Length的含义，并且帮助我们更好地处理HTTP响应体的内容。



