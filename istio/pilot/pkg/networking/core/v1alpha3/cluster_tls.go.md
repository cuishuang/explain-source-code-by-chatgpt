# File: istio/pilot/pkg/networking/core/v1alpha3/cluster_tls.go

在Istio项目中，`pkg/networking/core/v1alpha3/cluster_tls.go`文件负责定义与集群间TLS通信相关的功能。它定义了用于构建TLS配置的类型和方法，以及控制TLS通信的函数。

现在来详细介绍每个变量和函数的作用：

- `istioMtlsTransportSocketMatch`：用于匹配`Istio Mutual TLS`传输套接字规则的变量。
- `internalUpstreamSocket`：用于表示用于内部上游通信的传输套接字。
- `hboneTransportSocket`：用于表示`HBone`传输套接字。
- `hboneOrPlaintextSocket`：用于表示同时支持`HBone`和`Plaintext`传输套接字。

接下来是函数的作用：

- `applyUpstreamTLSSettings`：将上游TLS设置应用于上游集群。
- `buildUpstreamClusterTLSContext`：构建上游集群的TLS上下文。
- `applyTLSDefaults`：应用TLS默认设置。
- `setAutoSniAndAutoSanValidation`：设置自动SNI和自动SAN验证。
- `applyHBONETransportSocketMatches`：应用HBone传输套接字匹配规则。
- `defaultUpstreamCommonTLSContext`：构建默认的上游常规TLS上下文。
- `defaultTransportSocketMatch`：返回默认的传输套接字匹配规则。
- `buildUpstreamTLSSettings`：构建上游TLS设置。
- `hasMetadataCerts`：检查是否存在元数据证书。
- `buildMutualTLS`：构建互联网TLS配置。
- `buildIstioMutualTLS`：构建Istio Mutual TLS配置。

这些变量和函数是Istio中构建和应用TLS配置的关键组件，通过它们可以确保集群间的通信是安全的，并支持诸如自动验证和传输套接字匹配等功能。

