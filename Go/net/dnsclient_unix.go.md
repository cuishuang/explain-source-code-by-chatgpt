# File: dnsclient_unix.go

dnsclient_unix.go是Go语言标准库中net包下用于Unix系统的DNS客户端代码文件。它的主要作用是实现在Unix系统下对DNS服务器进行域名解析。具体来说，该文件中定义了一个名为dnsUnixClient的结构体，在该结构体中封装了DNS客户端连接通道的相关操作，以及读取并解析DNS服务器返回的响应的函数。 

在实现过程中，该文件使用了UDP协议进行DNS通信，通过DialUDP函数建立与DNS服务器之间的UDP连接，并使用DNS数据包来进行通信。客户端发送查询请求后，等待DNS服务器的回复，并对其进行解析，然后在一个超时时间内返回结果。 

此外，该文件中还实现了对DNS服务器地址解析的方法，用来解析给定的DNS服务器地址，并将其转换为网络字节序，在创建UDP连接时使用。 

总之，dnsclient_unix.go使得Go语言程序可以在Unix系统上通过使用UDP协议与DNS服务器进行通信，实现域名解析功能。




---

### Var:

### errLameReferral

在 Go 语言标准库中的 net 包中，dnsclient_unix.go 文件中的 errLameReferral 变量的作用是用于 DNS 的客户端查询时，当遇到 DNS 服务器返回的响应中包含了错误的引用时抛出异常。

具体而言，在 DNS 协议中，当 DNS 记录中的 DNS 服务器无法提供请求的主机名的 IP 地址时，DNS 服务器会返回指向其他 DNS 服务器的引用。当这个引用指向了错误或者无效的 DNS 服务器时，就会出现所谓的 Lame Referral 错误。此时，errLameReferral 变量就会被设置为一个错误信息，并被抛出给调用者处理。

在 Go 语言的 net 包中，当 errLameReferral 变量被设置时，它通常会在转发 DNS 请求时抛出异常，同时停止查询以避免陷入无限递归的情况。



### errCannotUnmarshalDNSMessage

在go/src/net中的dnsclient_unix.go文件中，errCannotUnmarshalDNSMessage变量的作用是表示在解析DNS消息时出现的错误。具体来说，这个变量用于表示当尝试将DNS消息从DNS响应解码为结构化的数据时出现问题时抛出的错误。

在DNS客户端中，当接收到一个DNS响应时，它需要将该响应转换成一个可读的结构化数据格式，以便能够进行后续的处理。如果无法将消息从DNS响应解析为结构化数据，则说明该响应包含不良数据或存在其他问题，这时就需要抛出一个错误，以保证程序的正确性和健壮性。

因此，errCannotUnmarshalDNSMessage变量的作用就在于表示此类解析错误，使得程序能够根据该错误做出相应的处理，从而提高程序的可靠性和稳定性。



### errCannotMarshalDNSMessage

在go/src/net中的dnsclient_unix.go文件中，errCannotMarshalDNSMessage是一个错误变量，用于在DNS查询过程中表示无法将DNS消息（message）进行编组（marshal）的错误。当向远程DNS服务器发送请求时，需要将DNS消息进行编组，以便可以在网络中进行传输。如果在此过程中发生编组错误，则会触发该错误变量，并将其作为错误信息返回给调用者。

在具体实现中，errorCannotMarshalDNSMessage定义为一个错误常量，其值为一个字符串，表示无法编组DNS消息。当在编组时遇到了错误，该常量被设置为错误信息，并返回给调用者。这有助于在发生错误时清晰地了解出现了什么问题，以便更好地诊断和解决问题。

总之，errCannotMarshalDNSMessage的作用是表示在DNS查询过程中无法将DNS消息进行编组的错误，并帮助调用者更好地了解遇到的问题。



### errServerMisbehaving

dnsclient_unix.go文件属于Go语言中与网络相关的代码模块，其中errServerMisbehaving是一个常量，用于表示一个DNS服务器的不良行为所导致的错误。

在DNS解析过程中，客户端向服务器发送请求来获取特定域名的解析结果。如果DNS服务器没有按照协议规范返回正确的响应，那么就表示该服务器的行为不当，可能导致后续的DNS解析出现问题。这种情况下，客户端需要通过一定的错误处理机制来处理这个问题。而errServerMisbehaving就是用于描述这种不良行为的错误类型，它告知客户端服务器的行为出了问题，需要进行错误处理，如重新请求或者直接放弃该服务器。

总之，errServerMisbehaving的作用是在DNS解析过程中捕捉DNS服务器的不当行为，并为客户端提供相应的错误处理机制，以保障DNS解析的正确性和安全性。



### errInvalidDNSResponse

在go语言的net包中，dnsclient_unix.go文件实现了Unix平台下的DNS解析功能。在该文件中，errInvalidDNSResponse是一个变量，定义如下：

var errInvalidDNSResponse = &DNSError{Err: "invalid DNS response"}

该变量的作用是当DNS响应无效时返回错误信息。DNS响应包括两个部分：头部和资源记录。如果响应包的头部不完整或者资源记录不完整或者不符合协议，则认为该响应是无效的。此时就可以通过errInvalidDNSResponse变量返回错误信息，提示用户DNS响应无效。

通过这个变量，程序可以在解析DNS响应时对错误情况进行处理，提高了程序的容错性和可靠性。



### errNoAnswerFromDNSServer

errNoAnswerFromDNSServer这个变量是一个自定义错误类型，用于表示从DNS服务器未收到响应的情况。

在发起DNS查询请求时，客户端程序需要等待DNS服务器的响应。如果DNS服务器在一定时间内未响应，则客户端程序将会认为查询失败，此时就会抛出errNoAnswerFromDNSServer这个错误，以告知上层调用者出现了无响应的情况。

该变量主要的作用是提高程序的稳定性，在网络环境不好或DNS服务器出现异常时能够及时发现问题，避免程序长时间等待或一直卡住不动的情况出现。同时，也可以方便地判断DNS查询结果是否有效，从而采取相应的处理措施。



### errServerTemporarilyMisbehaving

在go/src/net中的dnsclient_unix.go文件中，errServerTemporarilyMisbehaving变量的作用是表示DNS服务器临时发生故障，无法提供DNS服务。它是一个错误类型的变量，表示DNS解析时遇到的错误之一。

当DNS客户端在解析主机名时遇到errServerTemporarilyMisbehaving错误，通常会尝试重新发送DNS查询请求。如果DNS服务器仍然无法响应，则DNS解析请求可能会被传输到其他DNS服务器，或者返回给客户端一个错误报告。

这个变量的作用是为了帮助开发者在开发过程中更好地处理DNS解析错误情况，提高程序的稳定性和可靠性。例如，在代码中可以使用这个变量来编写错误处理逻辑，以便在遇到故障时能够更快地恢复DNS服务。



### resolvConf

在Go语言中，resolvConf变量用于存储操作系统中的resolv.conf文件中的DNS服务器配置信息。这个文件通常用来指定域名服务器（DNS服务器），以便操作系统能够正确地解析域名并查找对应的IP地址。

在Unix系统中，DNS客户端通常使用resolv.conf文件来查找DNS服务器的地址。在Go语言中，dnsclient_unix.go文件定义了一个名为resolvConf的变量，用于读取和解析resolv.conf文件，并提供了一个函数，可以使用这些信息来构建DNS请求。

这个变量类型是resolvConf结构体，包含了多个域名服务器信息，例如nameserver（域名服务器IP地址）、search（域名搜索路径）和options（器选项）等。

通过使用resolvConf变量，Go语言中的DNS客户端能够快速地查找并使用配置在操作系统中的DNS服务器。这避免了每次DNS查询时重新读取resolv.conf文件的耗时操作，从而提高了DNS解析的效率和速度。



### lookupOrderName

在net包中，dnsclient_unix.go文件中的lookupOrderName变量定义了在对主机名进行DNS解析时，使用的默认DNS服务器查找顺序。

lookupOrderName是一个字符串数组，其中包含以下内容的任意组合：

- "dns"：将在/ etc / resolv.conf中指定的DNS服务器中查找主机名。
- "files"：将在/ etc / hosts文件中查找主机名，以便通过IP地址获取主机名。
- "mdns"：使用mDNS查找主机名，这是一种局域网广播协议，用于在本地网络上发现设备。

从上述内容中可以看出，lookupOrderName影响主机名的解析方式，可以在查找DNS解析器之前，先查找本地主机名映射文件，也可以使用mDNS直接在网络中查找设备。lookupOrderName的默认值为{"dns"}，表示在/ etc / resolv.conf中指定的DNS服务器中查找主机名。

需要注意的是，在/ etc / resolv.conf文件中指定的DNS服务器可能已被编辑或删除，所以在实际使用时，需要根据情况进行调整。例如，在寻找内部网络主机时，可能需要添加额外的DNS服务器或在files中添加本地映射。

总之，lookupOrderName变量的作用是定义主机名解析顺序，并且它们的组合方式可以根据应用程序的需求进行自定义。






---

### Structs:

### resolverConfig

resolverConfig结构体是用于表示系统DNS解析器的配置信息，它包含了几个字段，分别是：

- nameservers：一个字符串数组，表示系统配置的DNS服务器地址。
- search：一个字符串数组，表示系统配置的DNS搜索域。
- ndots：一个整型值，表示DNS域名中点号分隔符的数量。如果域名中点号的数量少于这个值，则系统会在域名后面自动添加搜索域。
- timeout：一个时间间隔值，表示DNS查询的超时时间。
- attempts：一个整型值，表示DNS查询的最大重试次数。

resolverConfig结构体的作用是提供系统DNS解析器的配置信息，这些配置信息可以被net包中的DNS客户端使用，以便向指定的DNS服务器进行域名解析。在DNS客户端中，会首先尝试使用系统配置的DNS服务器进行解析，如果无法解析则尝试使用其他的DNS服务器进行解析。通过使用resolverConfig结构体，可以灵活地配置DNS客户端的行为，从而在网络连接出现故障或DNS服务器无法正常工作时，仍能提供可靠的域名解析服务。



### hostLookupOrder

在Go语言标准库中执行DNS解析时，DNS服务器的搜索顺序由hostLookupOrder结构体决定。 

hostLookupOrder结构体是一个数组，用于指定DNS服务的搜索顺序。该数组的元素分别是“首选DNS服务器”、“备用DNS服务器”、“网络接口的本地解析器”。

首选DNS服务器通常设置为本地ISP或其他“知名”的DNS服务器。备用DNS服务器则可以是另外一台服务器，或者您可以指定Google或其他公共DNS服务器。网络接口的本地解析器是指运行在操作系统上的本地DNS解析器。

DNS客户端会按照数组元素的顺序依次向这些服务器发送DNS查询请求，直到查询成功为止。如果所有服务器都无法返回结果，DNS解析将失败并返回错误信息。



## Functions:

### newRequest

在Go语言中，dnsclient_unix.go文件中的newRequest函数主要用于创建一个DNS查询请求，并返回一个DNS请求结构体。在DNS协议中，查询请求以及查询响应都需要遵循DNS消息格式。

在具体实现中，newRequest函数会根据查询类型和查询域名构造一个DNS请求消息，其中包含了查询类型、查询类、查询ID、DNS服务器地址等信息。在构造完请求消息之后，函数会将消息打包成一个字节数组，以便于在网络上传输。最终，该函数会返回一个DNS请求结构体，该结构体包含了DNS请求的相关信息。

通过newRequest函数，我们可以方便地创建一个DNS查询请求，并将其发送到DNS服务器，从而获取目标域名对应的IP地址等信息。



### checkResponse

checkResponse函数在net包中的dnsClient类型中被调用，主要用于验证从DNS服务器返回的DNS消息是否有效。它接收一个包含返回的DNS消息的字节切片，并且返回一个string类型的错误信息，如果返回的DNS消息没有错误，则返回nil。

该函数中主要验证以下方面：

1. 验证返回的DNS消息是否是一个合法的DNS消息。如果不是，返回错误信息"invalid DNS message"。
2. 验证DNS消息的信息是否与请求匹配。如果收到错误的消息类型或ID，则返回错误信息"mismatched DNS message"。
3. 验证DNS消息是否包含错误码。如果包含错误码，返回错误信息"DNS server error"。
4. 检查DNS消息中的回答数量。如果不是一个，则返回错误信息"DNS response format error, answer count is not 1".

总之，checkResponse函数的作用是确保返回的DNS消息是有效的，并且符合请求的要求。它确保在进行后续操作之前对DNS服务器返回的消息进行了有效的验证。



### dnsPacketRoundTrip

dnsPacketRoundTrip是一个函数，它在net/dnsclient_unix.go文件中定义。它的作用是向指定的DNS服务器发送DNS请求，并等待响应。

具体来说，该函数会构造一个DNS请求数据包（包括请求头和请求体），并将其发送到指定的DNS服务器。然后，它会等待服务器的响应，并返回响应数据（包括响应头和响应体）。

如果在超时时间内没有收到响应，该函数会返回一个超时错误。

dnsPacketRoundTrip函数的输入参数包括：

- DNS服务器地址
- 请求数据包（包括请求头和请求体）
- 超时时间

函数的返回值包括：

- 响应数据包（包括响应头和响应体）
- 错误（如果有的话）

该函数是网络编程中常用的函数之一，它在使用Go语言进行网络开发时非常有用。在实际应用中，可以使用该函数来查询域名、进行DNS解析等操作。



### dnsStreamRoundTrip

dnsStreamRoundTrip是一个在Unix系统下用于DNS解析的函数。它的作用是向指定的DNS服务器发送DNS查询请求，接收DNS服务器返回的数据，并将返回的数据解析成一个DNS消息。

具体来说，该函数首先使用net.Dial()函数连接到指定的DNS服务器，并向该服务器发送DNS查询请求。然后，它从连接中读取并解析DNS服务器返回的消息，直到返回的消息中包含所有查询请求的响应，或者超时时间到达。

在解析DNS消息时，该函数使用了net/dns包中的DNS消息解析函数，将返回的消息解析为一个dns.Msg对象。最后，该函数将解析出的DNS响应消息作为函数的返回值返回。

总之，dnsStreamRoundTrip函数是一个在Unix系统下实现DNS查询功能的函数，它使用了网络连接和DNS消息解析等技术，能够将DNS查询请求发送到指定的DNS服务器，并解析出服务器返回的DNS响应消息，从而完成DNS查询的功能。



### exchange

在go语言的net包中，dnsclient_unix.go文件中的exchange函数是用于向DNS服务器发送DNS请求和接收DNS响应的功能函数。当我们在客户端应用程序中需要发起DNS请求时，通过调用exchange函数向DNS服务器发送请求，然后等待DNS服务器的响应。

具体来说，exchange函数的作用是：

1. 创建DNS请求并向DNS服务器发送请求

在exchange函数中，先通过DNS请求类型和域名构建出DNS请求报文，并使用操作系统调用向DNS服务器发送请求报文。这里会根据配置的DNS服务器，向本地的DNS缓存服务器或者向公共的DNS服务器发送请求。

2. 接收DNS服务器的响应报文

等待DNS服务器返回的响应报文，并将响应报文的内容读取出来，根据协议规定的格式进行解析。在收到响应报文后，需要对响应报文的格式进行校验，确保响应报文的合法性。

3. 解析响应报文，获取IP地址

最后，从响应报文中获取到目标域名对应的IP地址，并将其返回给客户端应用程序，完成DNS查询的任务。

总之，exchange函数是一个非常关键的函数，在DNS查询过程中承担了发送请求、接收响应和解析响应等重要任务。通过exchange函数的执行，我们可以实现在应用程序中查询目标域名对应的IP地址，并将其用于后续的网络通信任务中。



### checkHeader

在go/src/net中dnsclient_unix.go文件中的checkHeader函数是一个用于检查DNS响应报文头的函数。

它的作用是检验DNS响应报文头部是否符合规定的格式，包括检查报文ID、标志字段、回答数量、权威数量和附加数量。如果检测失败，则会返回错误信息。

具体地，checkHeader函数会首先获取DNS响应报文头的各个字段的值，然后依次检查这些值是否符合规定的范围和格式。其中，报文ID必须与请求报文ID相同，标志字段中的QR位必须为1表示响应报文，回答数量、权威数量和附加数量必须均大于等于0，且不超过65535。

如果检测失败，则函数会返回一个错误码来标识错误的原因，例如：报文ID不匹配、非响应报文或者数量错误等。否则，函数会返回nil，表示检测通过，可以继续进行后续的解析和处理。

总的来说，checkHeader函数在DNS客户端中起到了确保DNS响应报文符合规定格式的重要作用，能够有效避免网络攻击和数据损坏等问题，保障了DNS服务的可靠性和安全性。



### skipToAnswer

skipToAnswer函数用于跳过DNS响应消息中的非Answer部分。它接受一个DNS响应消息的字节数组和一个指针，指针指向响应消息中要跳过的部分。

跳过的部分包括：

1. Header部分：DNS响应消息的开头包含一个12字节的标头，标头携带了关于响应消息的基本信息。

2. Question部分：DNS响应消息中Question部分包含一个或多个DNS问题字段，每个DNS问题字段都可以包含一个域名和与此域名相关的查询类型和查询类。

skipToAnswer函数通过将指针设置为Header和Question部分的长度来跳过这些部分，因为这些部分不包含Answer。在跳过Header和Question后，指针指向响应消息中的Answer部分，这就是skipToAnswer函数的主要作用。

在Net包中，skipToAnswer函数是用于Unix系统下进行DNS解析的，通过调用该函数将指针移动到DNS响应消息中的Answer部分，然后对Answer部分进行解析，提取所需要的信息。



### tryOneName

tryOneName是net包中dnsclient_unix.go文件中的一个函数，作用是尝试使用DNS服务器解析一个主机名，并返回可用的IP地址。

具体来说，tryOneName函数接收三个参数：主机名(hostname)、域名(dns)以及协议类型(proto)。它将主机名和域名组合成一个完整的带点号的域名，并使用系统库中默认的DNS服务器解析这个域名。如果解析成功，它将返回所有可用的IP地址，并按照优先级从高到低的顺序排列。

如果DNS服务器无法解析这个域名，或者解析出来的IP地址不可用，则会继续尝试使用下一个DNS服务器。如果所有的DNS服务器都无法解析这个域名，则tryOneName函数将返回一个错误。

在整个DNS解析过程中，tryOneName函数会依次尝试使用多个DNS服务器，以提高解析成功率。这些DNS服务器可以是本地系统上配置的，默认的DNS服务器，也可以是用户指定的备用DNS服务器。



### getSystemDNSConfig

`getSystemDNSConfig`函数的作用是获取当前操作系统的DNS配置信息。

具体来说，该函数会首先检查操作系统是否提供了获取DNS配置信息的API（如Linux上的`resolvconf`命令），如果有，则使用该API获取DNS配置信息。如果没有，则尝试读取本地的`/etc/resolv.conf`文件来获取DNS配置信息。

该函数返回的是一个`net.ResolverConfig`类型的结构体，其中包含了DNS服务器的地址、DNS搜索域、DNS超时设置等信息。这些信息对于网络编程非常重要，因为它们决定了程序如何解析域名并与远程服务器进行通信。

在实际使用中，程序员可以通过调用`net.LookupIP`等函数来使用这些DNS配置信息进行域名解析，或者手动配置自己的DNS解析器来修改默认的DNS配置信息。



### init

文件路径：go/src/net/dnsclient_unix.go

init函数是一个特殊的函数，它在包被加载时自动执行。dnsclient_unix.go文件中的init函数主要是初始化了一些变量和数据结构，以供后续的DNS查询使用。

具体来说，这个init函数主要完成了以下几个工作：

1. 初始化了一个名为dnsConfig的私有变量，它是一个dnsConfigStruct类型的数据结构，用于保存本地DNS服务器的相关信息（如IP地址、端口号等）以及DNS服务器的列表。

2. 通过调用系统的res_init函数初始化了系统的默认DNS服务器列表，并将其保存到dnsConfig这个变量中。同时，它还根据操作系统类型和版本，更新了dnsConfig中一些特定的字段。

3. 通过调用getHosts函数，获取本地/etc/hosts文件中的主机名和IP地址的映射关系，并将其保存到dnsConfig的hostsMap字段中。

4. 将dnsConfig这个变量设置为默认的DNS配置。

总之，这个init函数完成了在Unix系统上使用net包进行DNS查询之前，必要的一些初始化工作，以保证DNS查询的正常进行。



### tryUpdate

tryUpdate是一个函数，它是用于Unix平台的DNS客户端实现。它的作用是在没有收到响应的情况下尝试重新查询DNS服务器。

在DNS查询期间，如果DNS服务器无响应或者网络故障，tryUpdate将会调用waitRetry函数，然后等待retryDuration（默认设置是2秒）的时间。在此期间，它会检查网络连接并更新DNS服务器地址，以便下次查询时使用新的地址。

如果在重试期间仍然没有收到响应，tryUpdate会根据重试次数逐渐增加等待时间（默认的等待时间上限是64秒），并在超过等待时间上限时返回错误。

总之，tryUpdate函数通过监测网络故障和重新查询DNS服务器，以确保DNS查询功能在Unix平台上的正确运行。



### tryAcquireSema

在Go语言的net包中，dnsclient_unix.go文件中的tryAcquireSema函数是用于限制DNS查询并发执行的。在进行DNS查询时我们有时需要进行并发查询以提高查询效率，但如果同时进行过多的并发查询，会导致DNS服务器的负载过高，造成查询超时或失败等问题。因此，在进行DNS查询时我们需要对并发进行适当的限制。 

tryAcquireSema函数就提供了这种限制机制。它是基于信号量（semaphore）实现的锁机制，用于控制DNS查询并发量。当一个goroutine想要进行DNS查询时，它会先调用tryAcquireSema函数，如果当前的并发查询量小于maxActive，它就可以直接进行查询；否则，tryAcquireSema函数将会使当前goroutine等待，直到有其他goroutine完成了查询并释放了信号量，这时它才能开始自己的查询。 

tryAcquireSema函数的具体实现机制是：它会先尝试使用CAS原子操作来获取信号量，如果获取成功就返回true；否则它就会新建一个channel，并将该channel存储到dnsSemaChan数组中，并将这个channel返回。其他查询goroutine在调用tryAcquireSema函数时，会等待这个channel上的事件，当之前被占用的信号量被释放时，会发送一个事件通知到这个channel上，从而唤醒等待的goroutine。 

通过tryAcquireSema函数的限制机制，我们可以非常方便地控制DNS查询并发量，避免过多的并发查询造成DNS服务器的负载过高。



### releaseSema

releaseSema函数是一个信号量释放函数，它的主要作用是释放DNS请求完成的信号量，以便通知等待DNS请求结果的goroutine，DNS解析已经完成。当DNS请求完成时，releaseSema函数会将dnsPorts[d.Port]中的信号量加1，以便其他goroutine可以继续操作。这样就可以保证在DNS请求完成之前，其他goroutine会被阻塞，以避免在DNS请求结果未完成时发生竞争条件。通过释放信号量，可以确保在DNS请求完成后通知等待结果的goroutine，以便它们可以继续执行。

在DNS解析期间，等待DNS请求结果的goroutine会在DNS查询结果返回之前被阻塞。一旦DNS查询结果返回，DNS解析器将使用releaseSema函数释放信号量，以通知等待结果的goroutine，DNS查询已完成。这样，等待DNS结果的goroutine就可以继续执行，以处理DNS查询结果并返回结果给调用方。 

总的来说，releaseSema函数的主要作用是释放DNS结果完成的信号量，以通知等待DNS查询结果的goroutine，DNS查询已经完成。它保证了在DNS查询结果未返回之前，不会发生竞争条件，从而确保了并发安全性和正确性。



### lookup

lookup func是用于在Unix系统中使用DNS解析域名的函数。它使用系统的getaddrinfo函数实现，该函数将解析域名并返回与之对应的IP地址。

具体功能如下：

1. 接收一个主机名（例如"www.google.com"）和一个网络类型（例如"tcp"或"udp"）作为参数。

2. 首先将主机名转换为C语言的字符串，然后将网络类型转换为C语言的常量（例如"SOCK_STREAM"或"SOCK_DGRAM"）。

3. 调用getaddrinfo函数并将它们传递给它，它将返回一个地址信息列表，其中每个条目都包含一个IP地址、端口号和其他信息。

4. 函数将地址信息列表转换为Go语言的形式，并返回它们，以便调用者可以使用它们。

总之，lookup函数的作用是将一个主机名解析为与之对应的IP地址，并返回IP地址列表供调用者使用。它是网络编程中非常常用的一个函数，用于实现套接字编程等场景。



### avoidDNS

在 Go 语言中，dnsclient_unix.go 文件中的 avoidDNS 函数作为 DNS 查询的主要实现，用于在解析 IP 地址时避免使用指定的 DNS 服务器。该函数主要用于解决一些操作系统上存在的 DNS 规避机制（系统将某些 DNS 服务器列入规避清单，以避免 DNS 循环依赖或集中攻击等问题）。

具体来说，该函数首先会在系统上查找可用的 DNS 服务器列表，并对这些服务器进行排除和筛选，清除不可用的 DNS 服务器或被系统规避的 DNS 服务器。然后，该函数会将可用的 DNS 服务器与当前设置的 DNS 服务器进行对比，如果存在匹配的 DNS 服务器，则将其从可用列表中移除，以避免查询使用该 DNS 服务器出现问题。

简而言之，avoidDNS 函数的作用是在 Go 语言程序中通过操作系统获取可用 DNS 服务器列表，并解决可能存在的 DNS 规避问题，确保 DNS 查询结果正确且可靠。



### nameList

在go/src/net中的dnsclient_unix.go文件中，nameList这个函数的作用是解析DNS服务器Hosts文件的条目并返回一个字符串切片。在Unix操作系统上，DNS服务器配置通常存储在一个名为“/etc/hosts”的文件中，其中包含一系列IP地址和对应的主机名。

nameList函数的实现会打开这个文件并读取其中的所有行，然后将以空格或制表符分隔的第一个字段解释为IP地址，其余字段解释为主机名或别名。在返回的字符串切片中，每个IP地址和主机名的对应关系都以字符串“IP地址 主机名”表示。如果行中出现了注释字符（“#”），则后面的所有字符都被视为注释并被忽略。

在网络编程中，解析Hosts文件中的条目通常是在查询DNS服务器之前进行的，这样可以通过本地映射将某些常用的主机名映射到已知的IP地址，从而提高网络访问的效率和可靠性。



### String

在go/src/net/dnsclient_unix.go文件中，String函数的作用是返回一个字符串，该字符串以“dnsclient_unix”开头，并包含UNIX域套接字的路径信息。

具体来说，String函数返回一个包含UnixDialer结构体中socket的路径信息的字符串。UnixDialer结构体表示一个UNIX域套接字的拨号器，在DNS查询期间使用该拨号器来与名字服务器通信。String函数返回的字符串用于在调试和日志记录中提供上下文信息，有助于快速定位问题。

例如，如果我们使用UnixDialer来连接一个UNIX域套接字，代码如下：

```
d := net.Dialer{Timeout: 5 * time.Second, Deadline: time.Now().Add(10 * time.Second)}
u := net.UnixDialer{Timeout: 5 * time.Second, Deadline: time.Now().Add(10 * time.Second), Network: "unixgram", Path: "/var/run/dns.sock"}
conn, err := d.DialContext(ctx, u.Network, u.Path)
if err != nil {
    log.Fatal("Failed to dial UNIX domain socket:", err)
}
```

如果在连接过程中出现错误，我们可以打印UnixDialer结构体的字符串表示来获取更多上下文信息：

```
log.Println("UnixDialer:", u.String())
```

这样我们就可以知道哪个路径导致了连接错误，从而更容易地解决问题。



### goLookupHostOrder

goLookupHostOrder函数是net包中用于解析域名的一个函数，主要作用是根据操作系统的默认解析顺序以及用户配置的选项来确定解析的优先顺序。

具体来说，goLookupHostOrder首先会通过获取操作系统中的默认DNS服务器地址以及搜索域列表来确定默认的解析顺序。如果用户在环境变量中或者其他方式中指定了特定的DNS服务器或者搜索域列表，这些选项也会被加入到解析顺序中。

然后，该函数会依次查询每一个指定的域名服务器，直到找到一个能够解析出IP地址的服务器。如果所有的服务器都无法解析该域名，则返回错误信息。

在进行域名解析时，该函数还支持IPv6和IPv4两种地址族，并会自动选择合适的地址族进行解析。

总的来说，goLookupHostOrder函数提供了一个方便且灵活的域名解析接口，能够满足不同用户的需求。



### goLookupIPFiles

goLookupIPFiles是一个在Unix系统中实现的函数，它的作用是查找系统中的DNS解析配置文件并返回其中的IP地址列表。在Unix系统中，DNS解析配置文件通常是/etc/resolv.conf。该函数首先会通过getResolverConfig函数获取到系统中的DNS解析配置，然后解析出其中的DNS服务器地址列表和搜索域列表。接下来，它会通过resolveHostnames函数将需要解析的主机名解析为IP地址，并返回解析结果。

具体来说，goLookupIPFiles函数会执行以下步骤：

1. 调用getResolverConfig函数获取到系统中DNS解析配置。

2. 根据配置文件获取到DNS服务器地址列表和搜索域列表。

3. 调用resolveHostnames函数，将需要解析的主机名解析为IP地址。

4. 返回解析结果。

需要注意的是，goLookupIPFiles函数只会查询系统中的DNS解析配置文件，如果需要解析的主机名没有被配置，则无法解析出IP地址。同时，该函数只能在Unix系统中使用，因为Windows系统中并没有类似的DNS解析配置文件。



### goLookupIP

goLookupIP是一个函数，在net/dnsclient_unix.go中定义，它的作用是执行DNS查找操作，将域名转换为IP地址。具体而言，它会检查DNS域名解析器的缓存，如果缓存中已经存在该域名对应的结果，则直接返回；否则，它将向DNS服务器发送DNS请求，获取域名解析结果。

goLookupIP函数的输入参数包括域名（host）以及一些可选的参数，例如查询类型（queryType）和DNS服务器地址（dnsServer）。它返回一个IP地址列表和一个error对象。IP地址列表中包含了该域名解析的所有结果，可能是IPv4地址、IPv6地址、或者两者都包含。

在实现过程中，如果注入了DNS服务器地址，则会通过调用dnsExchange函数来发送DNS请求；否则会通过使用默认的DNS解析器来进行解析。默认情况下，系统会自动配置一个默认的DNS解析器，通常是用户的网络配置中提供的一个DNS服务器地址。

总之，goLookupIP函数提供了一个简单而高效地方式来执行DNS查找，进而实现从域名到IP地址的转换。它是net包中对于DNS查找的核心实现，可以帮助应用程序快速地获取所需的网络地址信息。



### goLookupIPCNAMEOrder

goLookupIPCNAMEOrder这个函数是在net包中实现的，主要作用是确定DNS查询顺序。它被用于解析URL中的域名，从而确定这个域名所对应的IP地址。

在函数内部，首先根据OS系统的不同，进行操作系统层面的查询。然后根据查询得到的结果，决定是否需要进行下一步操作（比如查询DNS服务器等）。最后，如果在所有查询操作中都没有找到IP地址，那么会返回一个错误信息。

具体来说，这个函数先尝试使用本地系统的hosts文件查询目标主机名字对应的IP地址，如果查询不到，则将查询请求发送给DNS服务器。如果DNS服务器返回了相应的IP地址，则直接使用该IP地址返回；否则，函数将会对DNS服务器返回的IP地址列表进行排序，并按照优先级顺序进行查询，直到找到一个可用的IP地址或者所有的查询操作都失败了。

总之，goLookupIPCNAMEOrder这个函数的作用就是为了确定访问目标主机所需的IP地址，将查询IP地址的操作封装起来，使得用户可以方便地进行IP地址的查询和使用。



### goLookupCNAME

goLookupCNAME函数是用来进行DNS解析中的CNAME记录查询的。在DNS中，CNAME记录是用来指向另一个域名的别名记录。因此，当我们向DNS服务器查询一个域名的IP地址时，我们需要查询该域名的所有CNAME记录，直到最终获得对应的IP地址。goLookupCNAME函数就是用来进行这个过程的。

具体来说，goLookupCNAME函数会首先向本地的DNS服务器进行查询（也可以通过resolv.conf文件指定其他DNS服务器）。如果本地DNS服务器无法解析该域名的CNAME记录，则该函数会向Internet上的根DNS服务器发出递归查询。递归查询就是在每次发出查询请求时都向上一级DNS服务器询问，直到最终获得所需的DNS记录。

最终，goLookupCNAME函数会返回该域名的最终目标地址（即A记录对应的IP地址）。如果该域名没有CNAME记录，则直接返回A记录的IP地址。

综上所述，goLookupCNAME函数在DNS解析中扮演了非常关键的角色，它能够帮助我们快速、准确地获取到域名对应的IP地址，从而保证了网络通信的顺利进行。



### goLookupPTR

函数名称：goLookupPTR

函数功能：完成PTR查询，返回反向域名

函数实现：

- 根据IP地址获取划分的网络类型并返回其类型标识符；
- 检查IP是否为IPv4或IPv6类型，如果不是则返回错误信息；
- 组成DNS查询消息的数据包，并向DNS服务器发送查询请求；
- 读取DNS响应数据，并解析响应以获取反向域名；
- 返回反向域名。

该函数的参数为IP地址(ip)和一个回调函数(fn)，回调函数用于处理反向域名的查询结果。



