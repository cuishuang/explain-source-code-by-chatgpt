# File: plugin/pkg/auth/authenticator/token/bootstrap/bootstrap.go

在 Kubernetes 项目中，`plugin/pkg/auth/authenticator/token/bootstrap/bootstrap.go` 文件的作用是实现了 TokenAuthenticator 接口的默认鉴权方法。下面对其中的结构体和函数进行详细介绍：

1. `TokenAuthenticator` 结构体：它是一个接口类型，定义了对 token 进行鉴权的方法。包括 `AuthenticateToken` 和 `Wrap`。

2. `wrappedTokenAuthenticator` 结构体：它是 `TokenAuthenticator` 接口的实现，包装了一个 `TokenAuthenticator` 实例和一个 `ErrorHandler`。当使用 `AuthenticateToken` 方法进行鉴权时，会先调用 `wrappedTokenAuthenticator` 的 `Resolve` 方法获取到实际的 `TokenAuthenticator` 实例，并调用其 `AuthenticateToken` 方法进行鉴权。

3. `NewTokenAuthenticator` 函数：它是一个构造函数，用于创建一个 `TokenAuthenticator` 实例。该函数接受一个参数 `tokenAllower`，是一个函数类型，用于判断 token 是否允许通过身份验证。`NewTokenAuthenticator` 函数返回一个 `TokenAuthenticator` 接口的实例。

4. `tokenErrorf` 函数：它是一个辅助函数，用于格式化错误信息，并返回一个 `AuthenticationError` 实例。

5. `AuthenticateToken` 函数：它是 `TokenAuthenticator` 接口的方法，用于对输入的 token 进行鉴权。该函数先调用 `tokenAllower` 判断 token 是否允许通过身份验证，如果允许则返回对应的 `UserInfo` 实例，否则返回一个 `AuthenticationError`。

总结起来，`bootstrap.go` 文件中的代码实现了对 token 的鉴权操作。通过 `NewTokenAuthenticator` 函数创建一个 `TokenAuthenticator` 实例，并使用 `AuthenticateToken` 方法对输入的 token 进行鉴权，根据鉴权结果返回对应的 `UserInfo` 或 `AuthenticationError` 实例。同时，在鉴权过程中，可以通过 `tokenAllower` 函数自定义判断 token 是否允许通过身份验证的逻辑。

