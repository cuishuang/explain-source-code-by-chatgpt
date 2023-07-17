# File: cmd/kubeadm/app/phases/certs/certs.go

文件`cmd/kubeadm/app/phases/certs/certs.go`是Kubernetes项目中的一个文件，用于实现证书相关的功能。下面将详细介绍该文件中的变量和函数。

**certPeriodValidationMutex:** 这是一个互斥量，用于确保证书有效期检查的并发安全性。

**certPeriodValidation:** 该变量是一个布尔值，用于确定是否应该进行证书有效期检查。如果为true，则会在创建证书时检查证书的有效期。

**certKeyLocation:** 这是一个结构体，用于存储证书和密钥的文件路径。

**CreatePKIAssets:** 这个函数用于创建PKI（Public Key Infrastructure）相关的证书和密钥文件。

**CreateServiceAccountKeyAndPublicKeyFiles:** 这个函数用于创建服务账号的密钥和公钥文件。

**CreateCACertAndKeyFiles:** 这个函数用于创建CA（Certificate Authority）的证书和密钥文件。

**NewCSR:** 这个函数用于创建一个新的CSR（Certificate Signing Request）。

**CreateCSR:** 这个函数用于创建CSR，并使用给定的证书颁发机构进行签名。

**CreateCertAndKeyFilesWithCA:** 这个函数用给定的CA签名颁发机构，创建一对新的证书和密钥文件。

**LoadCertificateAuthority:** 这个函数用于加载证书颁发机构的证书和密钥。

**writeCertificateAuthorityFilesIfNotExist:** 这个函数用于如果证书颁发机构的证书和密钥文件不存在，则创建并写入相关文件。

**writeCertificateFilesIfNotExist:** 这个函数用于如果证书文件不存在，则创建并写入相关文件。

**writeCSRFilesIfNotExist:** 这个函数用于如果CSR文件不存在，则创建并写入相关文件。

**SharedCertificateExists:** 这个函数用于检查是否存在共享的证书。

**UsingExternalCA, UsingExternalFrontProxyCA, UsingExternalEtcdCA:** 这些函数用于检查是否使用了外部的证书颁发机构。

**validateCACert, validateCACertAndKey, validateSignedCert, validateSignedCertWithCA, validatePrivatePublicKey, validateCertificateWithConfig, CheckCertificatePeriodValidity:** 这些函数用于验证证书和密钥的有效性、检查证书的有效期等证书相关的验证操作。

