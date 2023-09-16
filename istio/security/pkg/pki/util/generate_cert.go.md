# File: istio/security/pkg/pki/util/generate_cert.go

generate_cert.go文件的作用是用于生成证书，并提供了一些相关的功能函数和结构体。
1. SupportedECSignatureAlgorithms：这个结构体定义了一些支持的ECDSA签名算法。
2. SupportedEllipticCurves：这个结构体定义了一些支持的椭圆曲线。
3. CertOptions：这个结构体包含了生成证书所需的各种选项，如公钥、私钥、有效期等。

接下来是一些功能函数的介绍：
1. GenCertKeyFromOptions：根据给定的证书选项生成公钥和私钥。
2. genCert：根据给定的证书选项和私钥生成证书。
3. publicKey：从给定的私钥生成公钥。
4. GenRootCertFromExistingKey：根据给定的私钥生成根证书。
5. GetCertOptionsFromExistingCert：从给定的证书中提取证书选项。
6. MergeCertOptions：合并多个证书选项，生成新的证书选项。
7. GenCertFromCSR：根据给定的证书签名请求（CSR）生成证书。
8. LoadSignerCredsFromFiles：从文件中加载签名者的凭证。
9. genCertTemplateFromCSR：根据给定的CSR生成证书模板。
10. genCertTemplateFromOptions：根据给定的证书选项生成证书模板。
11. genSerialNum：生成一个序列号。
12. encodePem：将数据编码为PEM格式。

这些函数和结构体的组合使用，可以实现根据不同的选项来生成证书，包括生成自签名证书、根证书、从CSR生成证书等。同时提供了一些辅助函数，如从文件加载签名者凭证、编码数据为PEM格式等。

