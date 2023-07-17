# File: cmd/kubeadm/app/util/config/common.go

在Kubernetes项目中，`cmd/kubeadm/app/util/config/common.go`这个文件是 `kubeadm` 工具中的一个通用配置文件处理库。它包含了一些用于处理和验证 `kubeadm` 配置的函数。

1. `MarshalKubeadmConfigObject`：将 `kubeadm` 配置对象序列化为YAML格式的字节流。
2. `validateSupportedVersion`：验证指定的版本是否为支持的 Kubernetes 版本。
3. `NormalizeKubernetesVersion`：规范化 Kubernetes 版本字符串，确保它符合规范的格式。
4. `LowercaseSANs`：将 Subject Alternative Names（SANs）中的所有域名转换为小写格式。
5. `VerifyAPIServerBindAddress`：验证 `apiserver` 绑定地址是否合法。
6. `ChooseAPIServerBindAddress`：根据不同情况选择合适的 `apiserver` 绑定地址。
7. `validateKnownGVKs`：验证配置文件中的已知 GroupVersionKinds (GVKs) 是否是正确的。
8. `MigrateOldConfig`：支持旧版本配置文件的迁移，将旧版本的配置文件转换为当前版本的配置文件格式。
9. `ValidateConfig`：验证给定的 `kubeadm` 配置是否有效。
10. `isKubeadmPrereleaseVersion`：检查给定的版本是否为 `kubeadm` 的预发布版本。

这些函数在 `kubeadm` 工具的配置文件处理过程中起到了重要的作用，用于序列化、验证、规范化配置参数，并且还包含了一些旧版本配置的迁移功能。

