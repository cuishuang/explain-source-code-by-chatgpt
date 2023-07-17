# File: dnsconfig_unix.go

dnsconfig_unix.go是Go语言中net包下的一个文件，主要提供了Unix平台下获取DNS配置相关的功能。

在Unix平台下，DNS配置信息一般存储在resolv.conf文件中。dnsconfig_unix.go 中定义了一个`dnsReadConfigFile`函数，用于从resolv.conf文件中读取DNS配置信息。

具体步骤如下：

1. 首先，根据文件路径“/etc/resolv.conf”打开resolv.conf文件，并检查是否存在。如果文件不存在，直接返回一个nil值。
2. 然后，从resolv.conf文件中逐行读取内容，并对每一行的内容进行解析。由于resolv.conf文件的格式比较灵活，可以通过注释、空白行等对内容进行分割，因此需要对每一行的内容进行解析和处理。
3. 对于每一行的内容，首先需要检查其是否是合法的DNS配置信息。如果不是，需要跳过该行。如果是，则需要解析出DNS服务器的IP地址，并将其添加到返回结果中。
4. 最后，将解析出来的DNS配置信息按照一定的顺序排序，并返回这些信息。

总的来说，dnsconfig_unix.go 文件提供的功能主要是能够帮助开发者在Unix平台下快速获取DNS配置信息，方便网络程序的开发和维护。

## Functions:

### dnsReadConfig

dnsReadConfig是一个在Unix系统中读取DNS配置的函数。它首先会尝试从resolv.conf文件中读取DNS服务器的地址和搜索域名，如果读取失败，则会使用默认值。

这个函数返回一个结构体类型的DNS配置信息，包括了DNS服务器地址、搜索域名、是否在IPv6中使用RFC 3484规则等内容。这个配置信息可以在网络连接的时候使用，来决定使用哪个DNS服务器和搜索哪些域名。

在Unix系统中，DNS配置通常存储在resolv.conf文件中，并且在每次网络连接时被更新。dnsReadConfig函数可以从这个文件中读取配置信息，并且可以通过调用res_init函数将这些配置信息应用到当前进程中。



### dnsDefaultSearch

dnsDefaultSearch函数是一个内部函数，用于检查系统配置文件（如/ etc / resolv.conf）中是否存在“search”选项并返回该选项的值列表。如果未找到“search”选项，则会返回一个空的字符串切片。

该函数的作用是在UNIX系统上处理默认的DNS搜索域的配置。DNS搜索域是一组域名，在域名解析时，如果查询的域名不包含在这组域名中，则DNS解析器会自动在域名前添加这组域名的一个或多个成员，以扩展查询范围。 

例如，假设系统配置文件中存在以下内容：

```
search example.com subdomain.example.com
```

则当解析“example”时，DNS解析器将自动尝试以下查询：

```
example.com
example.subdomain.example.com 
```

如果没有找到“search”选项，则将使用默认DNS搜索域（通常是本地域名）进行查询。

在实际使用中，可以使用此函数设置自定义DNS搜索域，从而使应用程序能够更准确地解析域名。



### hasPrefix

在go/src/net中dnsconfig_unix.go文件中，hasPrefix函数的作用是判断一个字符串是否以另一个字符串开头。

函数定义如下：

```go
func hasPrefix(s, prefix []byte) bool {
    return len(s) >= len(prefix) && bytes.Equal(s[:len(prefix)], prefix)
}
```

该函数接收两个参数，s和prefix，均为字节切片（[]byte）。函数首先比较s的长度是否大于等于prefix的长度，如果是，再用bytes.Equal()函数判断s的前缀是否与prefix相同，如果相同，返回true，否则返回false。

该函数主要用于解析DNS服务配置文件中的参数。在DNS服务配置文件中，各个参数之间以空格或tab键分隔，因此需要使用hasPrefix函数来判断参数的开头是否正确。例如，以下代码用于解析DNS服务配置文件中的nameserver参数，并判断其开头是否为“nameserver”：

```go
func parse(resolvConf []byte) (*Config, error) {
    config := new(Config)
    for len(resolvConf) > 0 {
        line, rem := parseLine(resolvConf)
        if len(line) == 0 {
            resolvConf = rem
            continue
        }
        switch {
        case hasPrefix(line, []byte("nameserver")):
            servers := splitWords(line)
            if len(servers) < 2 {
                return nil, errors.New("bad nameserver line")
            }
            config.Servers = append(config.Servers, servers[1])
        case hasPrefix(line, []byte("domain")):
            if config.Search == nil {
                config.Search = new(SearchList)
            }
            *config.Search = append(*config.Search, string(splitWords(line)[1]))
        case hasPrefix(line, []byte("search")):
            if config.Search == nil {
                config.Search = new(SearchList)
            }
            *config.Search = splitWords(line)[1:]
        case hasPrefix(line, []byte("options")):
            // ignored for now
        }
        resolvConf = rem
    }
    return config, nil
}
```

代码使用hasPrefix函数判断line是否以“nameserver”开头，如果是，则将第二个参数解析出来并添加到Config结构体中的Servers字段中。如果line以其他关键字开头，则按照相应的方式解析。



### ensureRooted

ensureRooted函数的作用是确保DNS服务器的根节点以“.”结尾。它是在DNS resolver在解析主机名时使用的。

在Unix系统中，如果DNS服务器的根节点没有以“.”结尾，则可能会发生意外的行为。例如，如果没有以“.”结尾，则可能会将其解释为相对于本地域的名称，而不是绝对名称。这可能会导致不必要的延迟和错误的解析结果。

确保DNS服务器的根节点以“.”结尾可以防止这种情况的发生。如果DNS服务器的根节点不是以“.”结尾，则ensureRooted函数将在字符串的末尾添加“.”，以确保其是绝对名称。

总之，ensureRooted函数在Unix系统中确保DNS服务器的根节点以“.”结尾，以避免DNS解析出现问题。



