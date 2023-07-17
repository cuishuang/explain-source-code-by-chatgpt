# File: pkg/controller/certificates/approver/sarapprove.go

pkg/controller/certificates/approver/sarapprove.go是Kubernetes项目中负责证书认证的控制器，用于审批证书签发请求。

csrRecognizer结构体是用于识别证书签发请求的认证请求对象，sarApprover结构体是审批对象。该控制器通过识别签署请求并取消请求，以确保不被授权或需要的更新时为集群颁发证书。

该控制器的主要流程如下：

1. 从kube-apiserver中获取证书认证请求对象。
2. 根据识别器recognizers的预定义规则，判断请求是否应被此控制器处理。
3. 根据请求中提供的用户信息，用函数authorize判断是否允许根据请求进行证书签发。
4. 用函数appendApprovalCondition来追加个签发请求的审批条件。
5. 通过使用针对标记certificates.k8s.io下的CertificateSigningRequest类型的对象的标准控制器来创建请求。
6. 如果签名请求失败，那么控制器将用相同的方式撤消这个请求。

NewCSRApprovingController函数用于创建新控制器的一个实例，并返回一个用于启动和停止控制器的Stopable接口类型的对象。

recognizers、handle、authorize、appendApprovalCondition、isNodeClientCert、isSelfNodeClientCert、usagesToSet都是该控制器用到的实用函数。下面是各个函数的作用：

-recognizers函数是由csrRecognizer和sarApprover结构体组成的数组，用于确定哪些请求应该被该控制器处理。

-handle函数用于调用实际的处理逻辑，这个处理逻辑是定义在控制器的其他函数中。

-authorize函数接受一个user.Info对象，用于判断是否允许这个用户根据请求签发证书。

-appendApprovalCondition函数用于根据签名请求中提供的信息追加对请求的附加条件。

-isNodeClientCert函数用于检查请求的用途是否允许使用“节点客户端证书”的角色。

-isSelfNodeClientCert函数用于检查请求中提供的证书特征是否标识它为“自我签名”的节点客户端证书。

-usagesToSet函数用于确定可以在颁发的证书上设置的用途列表，这些用途通常包括“服务器身份验证”和“客户端身份验证”。

