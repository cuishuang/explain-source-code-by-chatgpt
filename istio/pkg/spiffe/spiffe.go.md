# File: istio/pkg/spiffe/spiffe.go

在istio项目中，spiffe.go文件是实现SPIFFE (Secure Production Identity Framework For Everyone)相关操作的代码文件。

1. trustDomain：表示SPIFFE证书的信任域。
2. trustDomainMutex：用于保护trustDomain变量的互斥锁。
3. firstRetryBackOffTime：SPIFFE证书获取重试的初始退避时间。
4. spiffeLog：用于记录SPIFFE日志的logger对象。

以下是一些重要结构体和函数的详细介绍：

结构体：
1. Identity：表示SPIFFE身份信息的结构体，包含一个SPIFFE URI和证书的Trust Domain ID。
2. bundleDoc：表示以JSON格式存储的SPIFFE Bundle的结构体，包含了根证书和各种SPIFFE ID与证书的映射关系。
3. PeerCertVerifier：表示SPIFFE身份验证器的结构体，用于验证对等身份及其证书。

函数：
1. ParseIdentity：解析一个SPIFFE URI并返回Identity结构体。
2. String：将Identity结构体转换为字符串表示。
3. SetTrustDomain：设置信任域。
4. GetTrustDomain：获取信任域。
5. GenSpiffeURI：生成SPIFFE URI。
6. MustGenSpiffeURI：生成SPIFFE URI，如果有错误则抛出panic。
7. ExpandWithTrustDomains：扩展以逗号分隔的SPIFFE ID列表，使用信任域。
8. GetTrustDomainFromURISAN：从URI的SAN字段解析出信任域。
9. RetrieveSpiffeBundleRootCerts：从SPIFFE Bundle中提取根证书。
10. NewPeerCertVerifier：创建一个用于验证对等证书的验证器。
11. GetGeneralCertPool：获取全局证书池。
12. AddMapping：向Bundle中添加SPIFFE ID与证书的映射关系。
13. AddMappingFromPEM：解析PEM格式的证书链，并将其添加到Bundle中的映射关系中。
14. AddMappings：将一组SPIFFE ID与证书的映射关系添加到Bundle中。
15. VerifyPeerCert：使用PeerCertVerifier验证对等证书。

