# File: grpc-go/security/advancedtls/crl.go

在grpc-go项目中，`grpc-go/security/advancedtls/crl.go`文件的作用是提供了用于验证TLS证书撤销的功能。

下面是关于各个变量和结构体的详细说明：

1. `grpclogLogger`：这个变量是用于记录日志的logger。
2. `oidDeltaCRLIndicator`：这个变量表示CRL的OID（Object Identifier），用于指示是否有增量CRL。
3. `oidIssuingDistributionPoint`：这个变量表示CRL的OID，用于指示颁发CRL的分发点。
4. `oidCertificateIssuer`：这个变量表示CRL的OID，用于指示颁发CRL的证书颁发者。
5. `oidAuthorityKeyIdentifier`：这个变量表示CRL的OID，用于指示颁发CRL的证书的密钥标识。
6. `crlPemPrefix`：这个变量表示CRL文件的PEM编码前缀。

接下来是关于各个结构体的详细说明：

1. `Cache`：这个结构体用于缓存CRL并提供操作缓存的方法。
2. `RevocationConfig`：这个结构体用于存储撤销验证的配置信息，包括是否启用CRL验证和是否启用OCSP验证等。
3. `RevocationStatus`：这个结构体用于存储撤销验证的状态信息，包括是否撤销和撤销原因等。
4. `certificateListExt`：这个结构体用于表示X.509证书的扩展信息。
5. `authKeyID`：这个结构体用于表示CRL的颁发者的密钥标识。
6. `issuingDistributionPoint`：这个结构体用于表示CRL的颁发者的分发点。

下面是关于各个函数的详细说明：

1. `String`：这个函数用于将结构体转换为字符串。
2. `x509NameHash`：这个函数用于计算X.509证书名称的哈希值。
3. `CheckRevocation`：这个函数用于检查证书是否被撤销。
4. `CheckChainRevocation`：这个函数用于检查证书链是否有任何一个证书被撤销。
5. `checkChain`：这个函数用于检查证书链中每个证书是否被撤销。
6. `cachedCrl`：这个函数用于获取缓存的CRL。
7. `fetchIssuerCRL`：这个函数用于获取颁发者CRL。
8. `checkCert`：这个函数用于检查证书是否被撤销。
9. `checkCertRevocation`：这个函数用于检查证书是否被撤销，如果是，则返回撤销原因。
10. `parseCertIssuerExt`：这个函数用于解析证书的颁发者扩展信息。
11. `parseCRLExtensions`：这个函数用于解析CRL的扩展信息。
12. `fetchCRL`：这个函数用于获取CRL。
13. `verifyCRL`：这个函数用于验证CRL的签名。
14. `crlPemToDer`：这个函数用于将CRL的PEM编码转换为DER编码。
15. `extractCRLIssuer`：这个函数用于从CRL中提取颁发者信息。
16. `hasExpired`：这个函数用于检查CRL是否已过期。
17. `parseRevocationList`：这个函数用于解析CRL的撤销列表。

以上是对`grpc-go/security/advancedtls/crl.go`文件中变量，结构体和函数的详细介绍。

