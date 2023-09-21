# File: grpc-go/credentials/local/local.go

在grpc-go项目中，`grpc-go/credentials/local/local.go`文件是用于提供本地测试凭证的功能。该功能允许您在本地环境中使用自签名证书进行TLS连接，而无需依赖第三方CA证书。

`info`结构体是用于存储证书的相关信息，包括证书和私钥的路径。

`localTC`结构体是一个实现了grpc的`TransportCredentials`接口的类型，它包装了证书相关的处理逻辑。

- `AuthType()`函数返回身份验证的类型，对于`localTC`来说，它总是返回`"tls"`。

- `Info()`函数返回证书的相关信息，如证书的路径。

- `getSecurityLevel()`函数返回TLS安全级别，默认是`tls.RequireAndVerifyClientCert`，即要求双向验证。

- `ClientHandshake()`用于从客户端发起的握手过程中，它会使用客户端的证书和私钥对连接进行加密。

- `ServerHandshake()`用于服务器端响应握手过程中，它会使用服务器的证书和私钥对连接进行加密。

- `NewCredentials()`函数用于创建新的`TransportCredentials`实例，该实例将使用给定的证书和私钥创建`localTC`。

- `Clone()`函数用于复制`TransportCredentials`实例。

- `OverrideServerName()`函数用于覆盖服务器名称。

这些函数共同实现了在本地环境中进行TLS连接的功能，方便了开发者进行本地测试和调试。

