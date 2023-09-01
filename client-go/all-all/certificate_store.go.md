# File: client-go/util/certificate/certificate_store.go

在K8s组织下的client-go项目中，certificate_store.go文件位于client-go/util/certificate目录下，它的作用是提供一种机制来处理证书相关操作，例如加载、存储和更新证书。

该文件中定义了两个结构体：fileStore和FileStore。

fileStore结构体是certificate_store.go文件的私有结构体，用于存储证书相关的信息，如证书文件路径、私钥文件路径、证书过期时间等。

FileStore结构体是公共的接口，定义了一组对证书存储进行操作的方法，包括加载证书、更新证书、判断证书文件是否存在等。

下面是一些重要的方法和函数的介绍：

- NewFileStore：该函数用于创建一个新的fileStore实例，并返回一个实现了FileStore接口的对象。

- CurrentPath：该方法返回当前证书的文件路径。

- recover：该方法用于从panic中恢复，并返回一个错误信息。

- Current：该方法返回当前证书的X509证书和私钥。

- loadFile：该方法用于加载证书文件，返回加载的证书内容。

- Update：该方法用于更新证书文件，将新证书内容写入文件中。

- updateSymlink：该方法用于更新符号链接文件。

- filename：该方法返回证书文件的名称。

- loadX509KeyPair：该方法用于加载X509证书和私钥对。

- fileExists：该方法用于判断证书文件是否存在。

总体而言，certificate_store.go文件提供了一套用于管理和操作证书的工具，可以方便地进行证书的加载、存储和更新等操作。

