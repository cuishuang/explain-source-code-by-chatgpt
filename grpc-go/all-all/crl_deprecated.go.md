# File: grpc-go/security/advancedtls/crl_deprecated.go

在grpc-go项目中，grpc-go/security/advancedtls/crl_deprecated.go文件的作用是提供了一个基于证书吊销列表（Certificate Revocation List，CRL）的证书验证功能。以下是对于所提到的变量、结构体和函数的详细介绍。

1. grpclogLogger: 这是用于日志记录的grpclog.Logger。

2. oidDeltaCRLIndicator: 表示CRL扩展的OID（Object Identifier），“delta CRL indicator”。

3. oidIssuingDistributionPoint: 表示CRL扩展的OID，查找颁发者分发点的扩展字段。

4. oidCertificateIssuer: 表示CRL扩展的OID，用于指定证书颁发者。

5. oidAuthorityKeyIdentifier: 表示CRL扩展的OID，用于指定证书的颁发者关键字。

6. crlPemPrefix: CRL文件的PEM（Privacy Enhanced Mail）格式前缀。

7. Cache: 用于缓存已验证过的CRL的结构体。

8. RevocationConfig: 用于配置吊销检查的结构体，包含CRL地址和是否启用吊销检查。

9. RevocationStatus: 用于表示证书吊销状态的结构体，包含了证书的序列号信息以及吊销状态。

10. certificateListExt: 用于存储CRL扩展字段的结构体，包含deltaCRLIndicator、issuingDistributionPoint和certificateIssuer。

11. authKeyID: 用于存储证书颁发者关键字的结构体，包含authorityKeyIdentifier。

12. issuingDistributionPoint: 用于存储颁发者分发点的结构体。

以下是一些函数的详细介绍：

- String: 用于将CRL扩展OID转换为字符串表示。

- x509NameHash: 计算X.509证书主体名称的哈希值。

- CheckRevocation: 检查证书是否被吊销。

- CheckChainRevocation: 检查证书链中的证书是否被吊销。

- checkChain: 检查证书链中的证书是否被吊销，具体调用了CheckChainRevocation。

- cachedCrl: 获取缓存的CRL。

- fetchIssuerCRL: 获取颁发者的CRL。

- checkCert: 检查证书是否被吊销，具体调用了CheckRevocation。

- checkCertRevocation: 检查证书是否被吊销的内部实现。

- parseCertIssuerExt: 解析证书颁发者扩展。

- parseCRLExtensions: 解析CRL扩展字段。

- fetchCRL: 获取CRL。

- verifyCRL: 验证CRL的合法性。

- extractCRLIssuer: 提取CRL的颁发者信息。

这些函数实现了CRL的下载、解析和验证，以及证书是否被吊销的检查等功能。

