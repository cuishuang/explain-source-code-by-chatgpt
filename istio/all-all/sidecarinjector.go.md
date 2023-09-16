# File: istio/pilot/pkg/bootstrap/sidecarinjector.go

在Istio项目中，`sidecarinjector.go`文件的作用是启动和配置Sidecar注入器。Sidecar注入器是Istio的一个组件，用于自动注入与应用程序一起运行的Sidecar代理。

该文件中的`injectionEnabled`变量用于判断是否启用了Sidecar注入功能。它是一个布尔类型的变量，如果为true，则表示启用了Sidecar注入，否则表示禁用了Sidecar注入。

`initSidecarInjector`函数的作用是初始化Sidecar注入器。它首先获取注入配置的命名空间和ConfigMap名称，然后将ConfigMap名称设置为启动参数`--kube-injector-configmap`的值。接下来，它关联一个Kubernetes事件处理器，用于监听ConfigMap的变化。最后，它启动Sidecar注入器的主循环。

`getInjectorConfigMapName`函数的作用是获取且返回Sidecar注入器的ConfigMap名称。它首先检查环境变量`INJECTOR_CONFIG_NAME`是否已设置，如果设置了，则返回环境变量的值。否则，它使用默认的ConfigMap名称`istio-inject`。

总的来说，`sidecarinjector.go`文件中的这些变量和函数负责启动和配置Sidecar注入器，并提供相关的配置信息。

