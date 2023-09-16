# File: istio/pkg/wasm/imagefetcher.go

文件imagefetcher.go的作用是定义了用于提取和解析容器镜像和WASM二进制文件的功能。

- `ImageFetcherOption`结构体用于存储`ImageFetcher`的配置选项。
- `ImageFetcher`结构体是一个用于提取和解析容器镜像和WASM二进制文件的工具，它使用`docker/docker`和`github.com/opencontainers/image-spec/specs-go/v1`包来处理镜像。
- `wasmKeyChain`结构体用于表示WASM二进制文件的密钥链。
- `useDefaultKeyChain`函数用于判断是否使用默认的密钥链。
- `String`方法用于将`ImageFetcherOption`结构体转换为字符串表示。
- `NewImageFetcher`函数用于创建一个新的`ImageFetcher`实例。
- `PrepareFetch`函数用于准备提取容器镜像或WASM文件。
- `extractDockerImage`函数用于提取并解析Docker镜像。
- `extractOCIStandardImage`函数用于提取并解析OCI标准的镜像。
- `extractWasmPluginBinary`函数用于提取并解析WASM插件二进制文件。
- `extractOCIArtifactImage`函数用于提取并解析OCI标准的Artifact镜像。
- `Resolve`函数用于解析容器镜像或WASM文件并返回解析结果。

总之，`imagefetcher.go`文件定义了一系列与容器镜像和WASM文件提取、解析相关的结构体和函数，提供了在istio项目中操作容器镜像和WASM文件的基本功能。

