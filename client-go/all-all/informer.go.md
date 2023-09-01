# File: client-go/dynamic/dynamicinformer/informer.go

在Kubernetes官方的client-go库中，client-go/dynamic/dynamicinformer/informer.go文件实现了动态资源的Informer接口。

而Informer的目标是在Kubernetes集群中跟踪特定资源的更改。它通过观察Kubernetes API server上特定资源的watch事件流来实现这一点，并将这些事件通知给与该资源相关联的处理函数。因此，通过使用Informer，可以方便地获取、监听和处理Kubernetes集群中资源的更改。

_interfaces定义了Informer接口：

```go
type Informer interface {
    factory.Getter
    GetStore() Store
    AddEventHandler(handler ResourceEventHandler)
    AddEventHandlerWithResyncPeriod(handler ResourceEventHandler, resyncPeriod time.Duration)
    Run(stopCh <-chan struct{})
    HasSynced() bool
    LastSyncResourceVersion() string
}
```

其中的下划线_变量是一个无类型的占位符，通常用于表示不需要使用该变量。

dynamicSharedInformerFactory是一个用于创建和管理dynamicInformer对象的工厂。它包含一组资源的配置和选项，并提供用于创建和管理dynamicInformer的函数。

dynamicInformer是dynamicSharedInformerFactory创建的实际Informer对象。它实现了Informer接口，并封装了底层Watcher，用于监听资源的事件。dynamicInformer还维护了一个本地缓存来存储资源的最新状态，并将这些状态通知给注册的处理函数。

以下是informer.go中的一些重要函数的简要介绍：

- NewDynamicSharedInformerFactory：创建一个新的dynamicSharedInformerFactory对象。
- NewFilteredDynamicSharedInformerFactory：创建一个新的dynamicSharedInformerFactory对象，该工厂会根据给定的谓词函数对资源进行筛选。
- ForResource：为给定的资源创建一个dynamicSharedInformerFactory对象。
- Start：启动dynamicSharedInformerFactory对象中所有的informer。
- WaitForCacheSync：等待dynamicSharedInformerFactory对象中所有informer的缓存同步完成。
- Shutdown：关闭dynamicSharedInformerFactory对象中所有的informer。
- NewFilteredDynamicInformer：创建一个新的dynamicInformer对象，该informer会根据给定的谓词函数对资源进行筛选。
- Informer：返回给定dynamicInformer对象的Informer接口。
- Lister：返回一个资源的Lister接口，用于从informer的缓存中获取指定资源的列表。

通过使用这些函数，可以方便地创建和管理dynamicInformer对象，并在应用程序中跟踪和处理与Kubernetes集群中资源相关的事件。

