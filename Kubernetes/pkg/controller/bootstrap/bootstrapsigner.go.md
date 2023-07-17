# File: pkg/controller/bootstrap/bootstrapsigner.go

pkg/controller/bootstrap/bootstrapsigner.go是Kubernetes集群启动过程中应用的控制器之一，负责为 API Server 的启动和运行提供签名证书和密钥。

在这个文件中，我们可以看到 SignerOptions 结构体的定义，它定义了用于生成和存储签名证书和密钥所需的所有选项。Signer 结构体定义了一个 Signer，它是一个签名证书生成器，用于生成签名证书并将其存储到 ConfigMap 中。

DefaultSignerOptions 函数返回一个默认的 SignerOptions 对象，NewSigner 函数用于创建一个新的签名器，Run 函数实现了控制器的实际逻辑，负责循环等待并调用子函数以处理 ConfigMap 和 Secret。

pokeConfigMapSync 函数用于当 ConfigMap 被更新时进行调用，serviceConfigMapQueue 用于服务一组 ConfigMap 和它们的资源版本列表，signConfigMap 函数用于生成签名证书并将其存储到 ConfigMap 中，updateConfigMap 函数用于将新的签名证书 ConfigMap 上传到 API Server 中，getConfigMap 用于检索一个 ConfigMap，listSecrets 用于列出所有已知的 Secret，getTokens 用于创建 Kubernetes 启动时需要使用的 API 访问 Token。

总的来说，pkg/controller/bootstrap/bootstrapsigner.go 文件是 Kubernetes 集群启动过程中必不可少的一部分，它提供了签名证书和密钥生成器，并支持签署证书以便 API Server 的安全启动。

