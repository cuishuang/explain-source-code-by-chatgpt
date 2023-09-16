# File: istio/istioctl/pkg/injector/injector-list.go

在Istio项目中，`istio/istioctl/pkg/injector/injector-list.go`文件的作用是实现了与Istio注入相关的命令行工具。下面将详细介绍其中的结构体和函数。

1. `revisionCount`是一个用于统计pod版本数量的结构体。它有两个字段：`revision`表示版本名称，`count`表示该版本pod的数量。

2. `Cmd`是一个注入相关命令的父命令，包含了`injectorListCommand`、`filterSystemNamespaces`、`getNamespaces`、`printNS`、`printHooks`、`getInjector`、`getInjectedRevision`、`getMatchingNamespaces`、`getPods`、`getInjectedImages`、`podCountByRevision`、`extractRevisionFromPod`、`injectionDisabled`和`renderCounts`等函数。

3. `injectorListCommand`函数是一个子命令，用于从集群中获取与注入相关的信息。

4. `filterSystemNamespaces`函数用于过滤掉系统级的命名空间，只返回用户定义的命名空间。

5. `getNamespaces`函数用于获取所有命名空间的列表。

6. `printNS`函数用于打印命名空间的信息。

7. `printHooks`函数用于打印注入的钩子信息，包括`pre`、`post`和`final`三种类型的钩子。

8. `getInjector`函数用于获取注入器的信息，包括注入器的名称和配置。

9. `getInjectedRevision`函数用于获取pod已注入的版本。

10. `getMatchingNamespaces`函数用于获取与给定标签匹配的命名空间。

11. `getPods`函数用于获取指定命名空间中的所有pod。

12. `getInjectedImages`函数用于获取pod中已注入的镜像列表。

13. `podCountByRevision`函数用于统计每个版本的pod数量。

14. `extractRevisionFromPod`函数用于从pod注释中提取版本信息。

15. `injectionDisabled`函数用于检查是否禁用了注入。

16. `renderCounts`函数用于将pod版本数量信息以表格形式打印出来。

这些函数通过调用Kubernetes API和Istio API，获取到与注入相关的信息，并进行处理和展示。

