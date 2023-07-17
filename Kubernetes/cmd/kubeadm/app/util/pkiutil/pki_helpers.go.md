# File: cmd/kubeadm/app/util/pkiutil/pki_helpers.go

在kubernetes项目中，`cmd/kubeadm/app/util/pkiutil/pki_helpers.go`文件是PKI（Public Key Infrastructure）辅助函数的实现。这些函数用于生成、加载、验证和操作与认证和加密相关的密钥和证书。

以下是`NewPrivateKey`这几个变量的作用：
- `rsaPrivateKey`：RSA算法生成的私钥
- `ecPrivateKey`： elliptic curve (EC) 算法生成的私钥
- `ed25519PrivateKey`：Ed25519算法生成的私钥

以下是`CertConfig`这几个结构体的作用：
- `CertificateConfig`：配置用于创建自签名证书的参数，如私钥、公钥、主题等。
- `CSRConfig`：配置用于创建CSR（Certificate Signing Request）的参数，包括私钥、公钥和主题。
- `CertKeyPair`：封装了证书和私钥，用于存储和加载。

以下是这些函数的作用：

- `NewCertificateAuthority`：生成新的证书签发机构（CA）。
- `NewIntermediateCertificateAuthority`：生成新的中间证书签发机构（intermediate CA）。
- `NewCertAndKey`：生成新的证书和私钥对。
- `NewCSRAndKey`：生成新的CSR和私钥对。
- `HasServerAuth`：检查给定的证书是否具有服务器身份验证。
- `WriteCertAndKey`：将证书和私钥写入磁盘。
- `WriteCert`：将证书写入磁盘。
- `WriteCertBundle`：将证书捆绑写入磁盘。
- `WriteKey`：将私钥写入磁盘。
- `WriteCSR`：将CSR写入磁盘。
- `WritePublicKey`：将公钥写入磁盘。
- `CertOrKeyExist`：检查给定路径上是否存在证书或私钥文件。
- `CSROrKeyExist`：检查给定路径上是否存在CSR或私钥文件。
- `TryLoadCertAndKeyFromDisk`：尝试从磁盘加载证书和私钥。
- `TryLoadCertFromDisk`：尝试从磁盘加载证书。
- `TryLoadCertChainFromDisk`：尝试从磁盘加载证书链。
- `TryLoadKeyFromDisk`：尝试从磁盘加载私钥。
- `TryLoadCSRAndKeyFromDisk`：尝试从磁盘加载CSR和私钥。
- `TryLoadPrivatePublicKeyFromDisk`：尝试从磁盘加载私钥和公钥。
- `TryLoadCSRFromDisk`：尝试从磁盘加载CSR。
- `PathsForCertAndKey`：根据给定的目录路径和文件名生成证书和私钥的路径。
- `pathForCert`：根据给定的目录路径和文件名生成证书路径。
- `pathForKey`：根据给定的目录路径和文件名生成私钥路径。
- `pathForPublicKey`：根据给定的目录路径和文件名生成公钥路径。
- `pathForCSR`：根据给定的目录路径和文件名生成CSR路径。
- `GetAPIServerAltNames`：获取用于API服务器的备用名（alternate name）列表。
- `GetEtcdAltNames`：获取用于Etcd的备用名列表。
- `GetEtcdPeerAltNames`：获取用于Etcd peer的备用名列表。
- `getAltNames`：从给定的主机名和备用名中生成备用名列表。
- `appendSANsToAltNames`：将备用名列表添加到给定的主机名和备用名中。
- `EncodeCSRPEM`：编码CSR为PEM格式。
- `parseCSRPEM`：解析PEM编码的CSR。
- `CertificateRequestFromFile`：从文件加载证书请求。
- `NewCSR`：生成新的CSR。
- `EncodeCertPEM`：编码证书为PEM格式。
- `EncodeCertBundlePEM`：编码捆绑的证书为PEM格式。
- `EncodePublicKeyPEM`：编码公钥为PEM格式。
- `GeneratePrivateKey`：生成新的私钥。
- `NewSignedCert`：根据给定的配置和证书签发机构生成新的签名证书。
- `RemoveDuplicateAltNames`：从备用名列表中去除重复的备用名。
- `ValidateCertPeriod`：验证证书的有效期。
- `VerifyCertChain`：验证整个证书链的有效性。

