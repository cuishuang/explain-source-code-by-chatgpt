# File: grpc-go/security/authorization/engine/engine.go

在grpc-go项目中，`grpc-go/security/authorization/engine/engine.go`文件的作用是实现gRPC的授权引擎，用于对接收到的请求进行授权决策。

以下是对各个变量和结构体的作用的详细介绍：

1. `logger`：用于记录日志的接口类型。
2. `stringAttributeMap`：用于存储字符串类型的属性键值对的映射。
3. `intAttributeMap`：用于存储整数类型的属性键值对的映射。

以下是对各个结构体的作用的详细介绍：

1. `activationImpl`：授权引擎的激活实现。
2. `AuthorizationArgs`：授权决策的参数。
3. `Decision`：授权策略的决策结果。
4. `AuthorizationDecision`：授权决策结果的结构体。
5. `policyEngine`：授权引擎的策略实现。
6. `AuthorizationEngine`：授权引擎的主要结构体。

以下是对各个函数的作用的详细介绍：

1. `ResolveName`：根据指定的名称解析对应的属性值。
2. `Parent`：获取授权请求的上级请求。
3. `newActivation`：创建一个授权激活。
4. `getRequestURLPath`：获取授权请求的URL路径。
5. `getRequestHost`：获取授权请求的主机。
6. `getRequestMethod`：获取授权请求的方法。
7. `getRequestHeaders`：获取授权请求的头部信息。
8. `getSourceAddress`：获取授权请求的源地址。
9. `getSourcePort`：获取授权请求的源端口。
10. `getDestinationAddress`：获取授权请求的目标地址。
11. `getDestinationPort`：获取授权请求的目标端口。
12. `getURISanPeerCertificate`：获取授权请求的URI SAN证书。
13. `getSourcePrincipal`：获取授权请求的源主体。
14. `String`：将授权结果转换为字符串。
15. `exprToParsedExpr`：将表达式转换为解析的表达式。
16. `exprToProgram`：将表达式转换为程序。
17. `newPolicyEngine`：创建一个新的策略引擎。
18. `getDecision`：获取授权决策结果。
19. `evaluate`：对授权策略进行评估。
20. `NewAuthorizationEngine`：创建一个新的授权引擎。
21. `Evaluate`：对授权请求进行评估。

以上是对`grpc-go/security/authorization/engine/engine.go`文件中各个变量和函数的作用的介绍。

