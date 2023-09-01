# File: client-go/transport/round_trippers.go

在client-go项目中，`round_trippers.go`文件定义了一系列的RoundTripper，它们被用于向Kubernetes集群发送HTTP请求。下面逐一介绍这些变量和函数的作用：

变量：
- `_`: 表示一个占位符，用于忽略某个值。
- `knownAuthTypes`: 一个包含已知身份验证类型的切片，这些类型是用于设置HTTP请求头中的`Authorization`字段的。现有的类型包括`BearerToken`和`BasicAuth`。
- `legalHeaderKeyBytes`: 合法的HTTP请求头名称的字节表示，用于检查请求头是否有效。

结构体：
- `authProxyRoundTripper`: 一个RoundTripper，负责将请求代理到认证代理服务器，以进行身份验证。
- `userAgentRoundTripper`: 一个RoundTripper，用于设置请求的用户代理(User-Agent)头。
- `basicAuthRoundTripper`: 一个RoundTripper，用于在HTTP请求头中添加基本身份验证信息。
- `impersonatingRoundTripper`: 一个RoundTripper，用于在请求中模拟另一个用户的身份。
- `bearerAuthRoundTripper`: 一个RoundTripper，用于在HTTP请求头中添加Bearer令牌身份验证信息。
- `requestInfo`: 一个结构体，包含有关HTTP请求的信息，如方法、URL和请求主体。
- `debuggingRoundTripper`: 一个用于调试的RoundTripper，它会将HTTP请求和响应的详细信息记录下来。
- `DebugLevel`: 调试级别，用于控制调试日志的详细程度。

函数：
- `HTTPWrappersForConfig`: 根据给定的Kubernetes配置，返回一系列HTTP请求包装器。
- `DebugWrappers`: 根据调试级别返回一系列HTTP请求包装器。
- `NewAuthProxyRoundTripper`: 创建一个认证代理RoundTripper。
- `RoundTrip`: 执行一个HTTP请求，并返回响应。
- `SetAuthProxyHeaders`: 在请求中设置认证代理头。
- `CancelRequest`: 取消正在进行的HTTP请求。
- `WrappedRoundTripper`: 包装给定的RoundTripper并返回。
- `NewUserAgentRoundTripper`: 创建一个用户代理RoundTripper。
- `NewBasicAuthRoundTripper`: 创建一个基本身份验证RoundTripper。
- `NewImpersonatingRoundTripper`: 创建一个模拟身份验证RoundTripper。
- `NewBearerAuthRoundTripper` 和 `NewBearerAuthWithRefreshRoundTripper`: 分别创建Bearer令牌身份验证RoundTripper。
- `newRequestInfo`: 创建一个请求信息结构体。
- `complete`: 根据请求信息和基础URL生成完整URL。
- `toCurl`: 将HTTP请求信息转换为curl命令。
- `NewDebuggingRoundTripper`: 创建一个调试RoundTripper。
- `maskValue`: 用“***”替换敏感值。
- `legalHeaderByte`: 检查单个字节是否是合法的HTTP请求头名称字节。
- `shouldEscape`: 检查字符是否应该进行URL编码。
- `headerKeyEscape`: 对HTTP请求头名称进行URL编码，确保有效性。

