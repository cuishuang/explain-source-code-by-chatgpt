# File: grpc-go/credentials/sts/sts.go

在grpc-go项目中，`grpc-go/credentials/sts/sts.go`文件是实现使用AWS Security Token Service (STS)进行认证的功能。STS是一种Web服务，允许您为AWS资源生成临时安全凭据。

下面是对每个变量和结构体的作用的详细介绍：

1. `loadSystemCertPool`：该变量用于加载系统的证书池，并返回一个`*x509.CertPool`对象，其中包含系统证书池中的可用证书。

2. `makeHTTPDoer`：该变量是用于创建HTTP客户端，用于发送HTTP请求。

3. `readSubjectTokenFrom`：该变量用于从给定的`*metadata.MD`中读取并返回`sts.SubjectTokenType`的值。这个值用作请求参数中的"subject_token_type"。 

4. `readActorTokenFrom`：该变量用于从给定的`*metadata.MD`中读取并返回`sts.ActorTokenType`的值。这个值用作请求参数中的"actor_token_type"。

5. `logger`：该变量用于在调试和记录日志信息时，输出相关信息。

以下是每个结构体的作用：

1. `Options`：这个结构体包含用于STS凭据认证的配置选项，包括Endpoint和STS客户端的自定义配置。

2. `callCreds`：这个结构体是`credentials.PerRPCCredentials`接口的实现，它持有STS凭据，并在请求携带凭据时使用。

3. `httpDoer`：这个结构体是一个HTTP客户端，它用于发送HTTP请求。

4. `requestParameters`：这个结构体包含用于构造请求的参数，包括主题令牌（SubjectToken）和令牌类型（TokenType）。

5. `responseParameters`：这个结构体包含从STS响应中提取的参数，包括Access Key、Secret Access Key、Session Token和Expiration。

6. `tokenInfo`：这个结构体包含从访问令牌（Access Token）或AWS STS响应中提取的令牌信息，如过期时间。

下面是每个函数的作用的详细介绍：

1. `String`：这个函数将`Options`结构体转换为人类可读的字符串，并返回结果。

2. `NewCredentials`：这个函数根据给定的`Options`创建一个新的STS凭据。

3. `GetRequestMetadata`：这个函数返回一个`map[string]string`类型的元数据，用于将凭据信息添加到gRPC请求的元数据中。

4. `RequireTransportSecurity`：这个函数返回一个布尔值，指示是否要求传输安全。

5. `makeHTTPClient`：这个函数创建一个HTTP客户端，用于发送HTTP请求。

6. `validateOptions`：这个函数用于验证Options结构体中的配置选项是否有效。

7. `cachedMetadata`：这个函数返回一个`transport.CachedMetadata`对象，用于缓存元数据。

8. `constructRequest`：这个函数根据给定的请求参数构造一个AWS STS请求。

9. `sendRequest`：这个函数使用HTTP客户端发送请求，并返回AWS STS的响应。

10. `tokenInfoFromResponse`：这个函数从AWS STS响应中解析并返回令牌信息。

这些函数和结构体的组合构成了使用AWS Security Token Service (STS)进行认证的完整实现。

