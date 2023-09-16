# File: istio/pkg/kube/inject/inject.go

在Istio项目中，`inject.go`文件的作用是对Kubernetes资源进行注入Istio sidecar代理。它是Istio注入功能的核心实现。

以下是相关变量和结构体的详细介绍：

- `KnownImageTypes`：这是一个包含已知镜像类型的映射表，用于确定需要注入的容器类型，如`sidecar`、`init`或`proxy_init`。

- `InjectionPolicy`：这个结构体用于表示注入策略，包括是否强制注入、注入的命名空间和是否注入初始化容器。

- `SidecarTemplateData`：这个结构体用于储存模板数据，用于生成sidecar代理的注入配置。

- `Template`：这个结构体表示用于注入的模板，包含要注入的容器、卷和环境变量。

- `Injector`：这个结构体封装了注入逻辑，包括从标准输入读取资源、解析注入配置、生成注入后的资源、输出到标准输出等。

- `Config`：这个结构体用于解析和存储注入配置。

- `SidecarInjectionStatus`：这个结构体用于表示注入状态，记录注入是否成功以及失败的原因。

以下是相关函数的功能介绍：

- `UnmarshalConfig`：将注入配置反序列化为`Config`结构体。

- `injectRequired`：检查给定的资源是否需要进行注入。

- `ProxyImage`：根据给定的代理版本和代理镜像类型返回完整的代理镜像地址。

- `imageURL`：从给定的Docker镜像名称解析镜像仓库URL。

- `updateImageTypeIfPresent`：更新给定镜像名称的镜像类型为已知类型。

- `RunTemplate`：执行模板，生成注入后的资源。

- `knownTemplates`：返回已知模板的映射表。

- `selectTemplates`：选择适用于给定资源的模板。

- `resolveAliases`：解析注入配置中的别名。

- `stripPod`：移除Pod中的注入标记。

- `injectionStatus`：返回注入状态。

- `parseDryTemplate`：解析用于干运行注入的模板。

- `runTemplate`：运行模板，生成注入后的资源。

- `IntoResourceFile`：将注入后的资源写入文件。

- `FromRawToObject`：将原始的JSON表示转换为Kubernetes对象。

- `IntoObject`：将注入后的资源写入Kubernetes对象。

- `applyJSONPatchToPod`：应用JSON Patch到Pod对象。

- `potentialPodName`：返回Pod对象的名称。

- `overwriteClusterInfo`：覆盖集群信息。

- `updateClusterEnvs`：更新环境变量中的集群信息。

- `GetProxyIDs`：获取代理的ID列表。

这些函数的功能在注入过程中起着关键作用，包括解析配置、选择模板、生成注入后的资源以及更新相关对象。

