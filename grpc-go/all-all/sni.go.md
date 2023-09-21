# File: grpc-go/security/advancedtls/sni.go

文件 `sni.go` 是 gRPC-Go 安全包中的一个源文件，位于 `grpc-go/security/advancedtls/` 目录下。

该文件的主要作用是提供选择证书的逻辑，基于 SNI 扩展（Server Name Indication）来选择不同的证书进行TLS握手。

在TLS握手过程中，客户端会发送一个SNI字段，该字段指定了客户端希望访问的服务器的域名。服务器可以使用这个信息来选择相应的证书进行握手，以确保安全连接的建立。

下面是 `sni.go` 文件中 `buildGetCertificates` 函数的详细介绍：

1. `buildGetCertificates`

   函数原型： `func buildGetCertificates(certMaps []*credentials.CertKeyMap) func(*tls.ClientHelloInfo) (*tls.Certificate, error)`

   - 函数参数：`certMaps []*credentials.CertKeyMap`，证书和私钥的映射关系列表
   - 函数返回值：返回一个函数类型，该函数接收 `*tls.ClientHelloInfo` 参数，并返回一个 `*tls.Certificate` 类型的指针和一个 `error` 类型的错误
   
   该函数用于构建一个函数，这个函数将根据传入的 `certMaps` 参数选择相应的证书进行SNI选择，并返回用于TLS握手的证书及可能的错误。

   `certMaps` 参数是一个包含多个 `CertKeyMap` 对象的切片。`CertKeyMap` 对象是一个结构体，用于表示证书和对应私钥的映射关系。

2. `getCertificate`

   函数原型：`func getCertificate(certMaps []*credentials.CertKeyMap, sniName string) (*tls.Certificate, error)`

   - 函数参数：`certMaps []*credentials.CertKeyMap`，证书和私钥的映射关系列表；`sniName string`，SNI字段中指定的域名
   - 函数返回值：返回一个 `*tls.Certificate` 类型的指针和一个 `error` 类型的错误
   
   该函数用于根据SNI字段中指定的域名从 `certMaps` 参数中选择相应的证书，并返回用于TLS握手的证书及可能的错误。

   在函数内部，会遍历 `certMaps` 切片，对比其中的每个 `CertKeyMap` 对象的域名与 `sniName` 是否相等。如果相等，说明找到了匹配的证书，将其转换为 `*tls.Certificate` 类型，并返回。

   如果遍历完整个切片后仍未找到匹配的证书，函数将返回错误。

3. `getDefaultCertificate`

   函数原型：`func getDefaultCertificate(certMaps []*credentials.CertKeyMap) (*tls.Certificate, error)`

   - 函数参数：`certMaps []*credentials.CertKeyMap`，证书和私钥的映射关系列表
   - 函数返回值：返回一个 `*tls.Certificate` 类型的指针和一个 `error` 类型的错误
   
   该函数用于从 `certMaps` 参数中选择默认的证书，如果没有找到默认证书，则返回错误。

   默认证书是指在 `CertKeyMap` 对象中没有域名的证书。函数会遍历 `certMaps` 切片，在找到第一个没有域名的 `CertKeyMap` 对象时，将其转换为 `*tls.Certificate` 类型，并返回。

   如果遍历完整个切片后仍未找到默认证书，函数将返回错误。

