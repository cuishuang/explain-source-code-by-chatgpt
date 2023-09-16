# File: istio/security/pkg/pki/util/keycertbundle.go

在istio项目中，`istio/security/pkg/pki/util/keycertbundle.go` 这个文件是用于处理密钥和证书的工具函数。它定义了 `KeyCertBundle` 结构体及相关的函数，用于管理密钥和证书的加载、验证和处理。

下面是对 `KeyCertBundle` 结构体及其函数的详细介绍：

1. `KeyCertBundle` 结构体：表示密钥和证书的集合，包含以下字段：
   - `PrivateKey`：该密钥关联的私钥。
   - `CertChain`：该证书关联的证书链。
   - `RootCert`：根证书，用于验证证书链。
   - `CertOptions`：证书选项，控制证书的生成和加载。

2. 函数 `NewKeyCertBundleFromPem(pemBlock *pem.Block)`：从 PEM 格式的密钥和证书数据创建一个新的 `KeyCertBundle` 实例。

3. 函数 `NewVerifiedKeyCertBundleFromPem(pemBlocks []*pem.Block)`：从 PEM 格式的密钥和证书数据创建一个新的 `KeyCertBundle` 实例，并验证证书链及根证书。

4. 函数 `NewVerifiedKeyCertBundleFromFile(keyFile, certFile, rootCertFile string)`：从文件中加载密钥、证书和根证书，并创建一个新的 `KeyCertBundle` 实例，并验证证书链及根证书。

5. 函数 `NewKeyCertBundleWithRootCertFromFile(keyFile, certFile, rootCertFile string)`：从文件中加载密钥、证书和根证书，并创建一个新的 `KeyCertBundle` 实例，但不进行验证。

6. 函数 `GetAllPem()`：返回包含所有密钥和证书的 PEM 格式数据。

7. 函数 `GetAll()`：返回包含所有密钥和证书的字节数组数据。

8. 函数 `GetCertChainPem()`：返回证书链的 PEM 格式数据。

9. 函数 `GetRootCertPem()`：返回根证书的 PEM 格式数据。

10. 函数 `VerifyAndSetAll(keyPEM, certPEM, rootCertPEM []byte) error`：验证给定的密钥、证书和根证书，并设置到当前的 `KeyCertBundle` 实例。

11. 函数 `setAllFromPem(keyPEM, certPEM, rootCertPEM []byte) error`：从 PEM 格式数据中设置密钥、证书和根证书到当前的 `KeyCertBundle` 实例。

12. 函数 `CertOptions(certBytes []byte) (*x509.Certificate, error)`：从证书字节数据中解析出证书，并返回证书选项。

13. 函数 `UpdateVerifiedKeyCertBundleFromFile(keyFile, certFile, rootCertFile string) error`：从文件中加载密钥、证书和根证书，并更新当前的 `KeyCertBundle` 实例，并验证证书链及根证书。

14. 函数 `ExtractRootCertExpiryTimestamp(rootCert []byte) (time.Time, error)`：从根证书中提取过期时间戳。

15. 函数 `ExtractCACertExpiryTimestamp(caCert []byte) (time.Time, error)`：从 CA 证书中提取过期时间戳。

16. 函数 `TimeBeforeCertExpires(cert *x509.Certificate) time.Duration`：返回证书的有效期剩余时间。

17. 函数 `Verify(chains [][]*x509.Certificate, crls []pkix.CertificateList, trustedRoots []*x509.Certificate) error`：验证给定的证书链是否有效。

18. 函数 `extractCertExpiryTimestamp(certBytes []byte) (time.Time, error)`：从证书字节数据中提取过期时间戳。

19. 函数 `copyBytes(src []byte) []byte`：复制字节切片。

以上是 `keycertbundle.go` 文件中定义的结构体和函数的功能和作用的详细介绍。这些函数提供了一套用于加载、验证和处理密钥和证书的工具函数，可以用于实现与安全证书相关的功能。

