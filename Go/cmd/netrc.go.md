# File: netrc.go

netrc.go是Go语言标准库中的一个文件，它实现了从 .netrc 文件中获取认证信息的功能。

Netrc（“网络文件”，即.netrc）是一个用于保存FTP客户端认证信息的文件。它通常是在用户家目录下的隐藏文件，保存了登录FTP服务器所需要的用户名、密码、主机名等信息。当用户使用FTP客户端访问FTP服务器时，FTP客户端会自动读取.netrc文件中的信息，然后将这些信息用于认证登录FTP服务器。

netrc.go文件提供了从.netrc文件中获取认证信息的接口。它能够解析.netrc文件中的信息，并提供了函数接口可以获取特定主机名下的用户名和密码信息。

该文件通过下面几个函数提供了获取.netrc信息的功能：

- func netrc() (map[string]*userData, error)：解析标准的.netrc文件，并返回主机名到认证用户信息的映射。
- func (rc *Netrc) FindMachine(name string) *Machine：根据主机名查找认证信息。
- func (rc *Netrc) Password(machine string) (string, error)：根据主机名查找密码。
- func (rc *Netrc) Login(machine string) (string, string, error)：根据主机名查找用户名和密码。

在实际使用中，我们可以直接使用标准库中的net包的Dial、DialTimeout等函数连接FTP服务器，net包会自动读取用户的.netrc文件中的信息进行连接认证。我们只需要提供FTP服务器的主机名和端口号即可。这极大地简化了FTP客户端的使用。




---

### Var:

### netrcOnce





### netrc





### netrcErr








---

### Structs:

### netrcLine





## Functions:

### parseNetrc





### netrcPath





### readNetrc





