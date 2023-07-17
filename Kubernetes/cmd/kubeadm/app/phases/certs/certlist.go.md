# File: cmd/kubeadm/app/phases/certs/certlist.go

cmd/kubeadm/app/phases/certs/certlist.go文件的作用是定义了一些关于证书操作的数据结构和函数。

- configMutatorsFunc是一个函数类型，表示配置的变更函数。
- KubeadmCert是一个表示kubeadm证书的结构体，包含了证书的名称、签发者、使用者等信息。
- CertificateTree是一个表示证书树的结构体，用于存储证书的层次结构。
- CertificateMap是一个表示证书映射的结构体，用于存储证书的名称和对应的证书对象。
- Certificates是一个表示证书列表的结构体，包含了多个证书的信息。

以下是几个关键函数和结构体的作用：

- GetConfig函数从kubeconfig配置文件中获取集群和认证的信息。
- CreateFromCA函数使用给定的CA证书和秘钥创建kubeadm证书。
- CreateAsCA函数使用给定的CA证书和秘钥创建一个CA证书。
- CreateTree函数创建一个证书树，根据已有的证书和相关配置。
- CertTree函数根据已有的证书和配置，创建一个证书树并返回。
- AsMap函数将证书树转换为证书映射。
- GetDefaultCertList函数返回默认的kubeadm证书列表。
- GetCertsWithoutEtcd函数返回不包含etcd证书的kubeadm证书列表。
- KubeadmCertRootCA函数返回Root CA证书相关的kubeadm证书。
- KubeadmCertAPIServer函数返回APIServer证书相关的kubeadm证书。
- KubeadmCertKubeletClient函数返回Kubelet Client证书相关的kubeadm证书。
- KubeadmCertFrontProxyCA函数返回FrontProxy CA证书相关的kubeadm证书。
- KubeadmCertFrontProxyClient函数返回FrontProxy Client证书相关的kubeadm证书。
- KubeadmCertEtcdCA函数返回Etcd CA证书相关的kubeadm证书。
- KubeadmCertEtcdServer函数返回Etcd Server证书相关的kubeadm证书。
- KubeadmCertEtcdPeer函数返回Etcd Peer证书相关的kubeadm证书。
- KubeadmCertEtcdHealthcheck函数返回Etcd Healthcheck证书相关的kubeadm证书。
- KubeadmCertEtcdAPIClient函数返回Etcd API Client证书相关的kubeadm证书。
- makeAltNamesMutator函数返回一个配置变更函数，用于添加备用名称到证书的SAN字段。
- setCommonNameToNodeName函数返回一个配置变更函数，用于将节点的名称设置为证书的CN字段。
- leafCertificates函数返回一个包含所有叶子证书的列表。
- createKeyAndCSR函数创建私钥和CSR(certificate signing request)文件。
- CreateDefaultKeysAndCSRFiles函数创建默认的私钥和CSR文件。

