# File: named.go

named.go文件是Go语言标准库中的一部分，它定义了一个名字解析器（resolver）的接口，以及用于创建和配置名字解析器的函数。

名字解析器的作用是将主机名（host name）和服务名（service name）转换为网络地址。在网络编程中，通常需要通过主机名或服务名来定位网络上的服务或计算机，而名字解析器就起到了这个作用。

named.go文件中的接口定义了名字解析器的方法，其中包括解析主机名和服务名、返回网络地址等。标准库中的实现包括DNS解析器和本地主机文件解析器。

此外，named.go文件中还定义了一些函数，用于创建和配置名字解析器。其中包括DefaultResolver函数，用于获取默认的名字解析器，以及LookupAddr、LookupPort等函数，用于解析主机名和服务名。

总之，named.go文件是标准库中重要的网络编程工具，提供了名字解析器的抽象和实现，方便开发者进行网络编程。




---

### Structs:

### Named





### instance





### namedState





## Functions:

### NewNamed





### resolve





### state





### setState





### newNamed





### newNamedInstance





### cleanup





### Obj





### Origin





### TypeParams





### SetTypeParams





### TypeArgs





### NumMethods





### Method





### expandMethod





### SetUnderlying





### AddMethod





### Underlying





### String





### under





### setUnderlying





### lookupMethod





### context





### expandUnderlying





### safeUnderlying





