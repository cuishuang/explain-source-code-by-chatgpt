# File: istio/pilot/pkg/security/authn/v1beta1/policy_applier.go

在Istio项目中，`istio/pilot/pkg/security/authn/v1beta1/policy_applier.go`文件负责实现身份认证策略的应用。它提供了一些函数和结构体，用于将PeerAuthentication策略应用到envoy代理中。

`authnLog`变量是用于记录身份认证策略应用过程中的日志信息。

`v1beta1PolicyApplier`是一个结构体，包含用于应用PeerAuthentication策略的函数。

`MergedPeerAuthentication`是一个结构体，用于表示合并后的PeerAuthentication策略配置。这个结构体由多个PeerAuthentication策略进行合并而成。

`NewPolicyApplier`是一个函数，用于创建一个新的`v1beta1PolicyApplier`对象。

`JwtFilter`是一个结构体，表示Envoy中的Jwt Filter配置。

`defaultAuthnFilter`是一个函数，用于创建默认的AuthN Filter配置。

`setAuthnFilterForRequestAuthn`是一个函数，用于根据请求级别的AuthN配置设置AuthN Filter。

`AuthNFilter`是一个结构体，表示AuthN Filter配置。

`InboundMTLSSettings`是一个结构体，表示入站请求的MTLS配置。

`convertToEnvoyJwtConfig`是一个函数，用于将Istio的Jwt Filter配置转换为Envoy的Jwt Filter配置。

`PortLevelSetting`是一个结构体，表示端口级别的配置。

`GetMutualTLSModeForPort`是一个函数，用于获取端口级别的MTLS模式配置。

`ComposePeerAuthentication`是一个函数，用于合并多个PeerAuthentication策略配置。

`isMtlsModeUnset`是一个函数，用于判断MTLS模式是否未被设置。

这些函数和结构体相互配合，实现了将PeerAuthentication策略应用到envoy代理中的功能。

