# File: istio/security/pkg/stsservice/tokenmanager/google/mock/mockserver.go

mockserver.go是一个模拟服务器，用于在Google Cloud中模拟Istio服务的行为。它用于在测试或开发过程中提供一个虚拟的Google Cloud环境，以完成与Google OIDC服务器的交互。

具体变量的作用如下：
- FakeFederatedToken：模拟的联邦令牌值。
- FakeAccessToken：模拟的访问令牌值。
- FakeTrustDomain：模拟的信任域。
- FakeSubjectToken：模拟的主体令牌。
- FakeProjectNum：模拟的项目编号。
- FakeGKEClusterURL：模拟的GKE集群URL。
- FakeExpiresInSeconds：模拟的令牌过期时间（以秒为单位）。

具体结构体的作用如下：
- federatedTokenRequest：联邦令牌请求的结构体。
- federatedTokenResponse：联邦令牌响应的结构体。
- Duration：持续时间的结构体。
- accessTokenRequest：访问令牌请求的结构体。
- accessTokenResponse：访问令牌响应的结构体。
- AuthorizationServer：授权服务器的结构体。
- Config：配置信息的结构体。

具体函数的作用如下：
- StartNewServer：启动一个新的模拟服务器。
- SetGenFedTokenError：设置生成联邦令牌时的错误。
- BlockFederatedTokenRequest：阻塞联邦令牌请求。
- BlockAccessTokenRequest：阻塞访问令牌请求。
- SetGenAcsTokenError：设置生成访问令牌时的错误。
- SetTokenLifeTime：设置令牌的生命周期。
- SetAccessToken：设置访问令牌值。
- EnableDynamicAccessToken：启用动态访问令牌。
- NumGetAccessTokenCalls：获取访问令牌的调用次数。
- NumGetFederatedTokenCalls：获取联邦令牌的调用次数。
- Start：启动模拟服务器。
- Stop：停止模拟服务器。
- getFederatedToken：获取联邦令牌。
- getAccessToken：获取访问令牌。

这些函数和变量的目的是通过模拟Google OIDC服务器的行为来帮助测试和开发Istio项目中与安全相关的功能。

