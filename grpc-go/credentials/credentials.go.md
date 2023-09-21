# File: grpc-go/credentials/credentials.go

`grpc-go/credentials/credentials.go`文件是gRPC-go库中负责处理认证和安全相关功能的核心文件。它定义了一些接口，结构体和函数，用于处理与认证和安全相关的操作。

以下是对每个变量和结构体的详细介绍：

1. `ErrConnDispatched`是一个错误变量，当连接已经分发给客户端进行使用时会被返回，表示连接已被使用。

以下是对每个结构体的详细介绍：

1. `PerRPCCredentials`是一个接口，定义了在每个RPC调用中使用的认证凭据的方法。
2. `SecurityLevel`是一个接口，定义了安全等级的方法。
3. `CommonAuthInfo`是一个接口，定义了共同的认证信息的方法。
4. `ProtocolInfo`是一个接口，定义了特定协议的认证信息和安全等级的方法。
5. `AuthInfo`是一个接口，定义了认证信息的方法。
6. `TransportCredentials`是一个接口，定义了传输层凭据的方法。
7. `Bundle`是一个接口，定义了证书捆绑的方法。
8. `RequestInfo`是一个接口，定义了请求信息的方法。
9. `ClientHandshakeInfo`是一个接口，定义了客户端握手信息的方法。
10. `ChannelzSecurityInfo`是一个接口，定义了通道安全信息的方法。
11. `ChannelzSecurityValue`是一个接口，定义了通道安全值的方法。
12. `OtherChannelzSecurityValue`是一个接口，定义了其他通道安全值的方法。

以下是对每个函数的详细介绍：

1. `String`函数返回一个字符串，表示该认证凭据的类型。
2. `GetCommonAuthInfo`函数返回一个共同的认证信息。
3. `RequestInfoFromContext`函数从上下文中获取请求信息。
4. `ClientHandshakeInfoFromContext`函数从上下文中获取客户端握手信息。
5. `CheckSecurityLevel`函数检查给定的安全等级是否与所需的安全等级相匹配。

总体来说，`grpc-go/credentials/credentials.go`文件提供了认证和安全相关的功能和接口，供gRPC-go库实现和扩展使用。

