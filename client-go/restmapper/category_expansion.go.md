# File: client-go/restmapper/category_expansion.go

在Kubernetes（K8s）组织下的client-go项目中，client-go/restmapper/category_expansion.go文件是一个用于资源映射的辅助工具。该文件中定义了一些结构体和函数，用于扩展和管理Kubernetes API中的资源种类。

1. CategoryExpander结构体：
   - CategoryExpander是一个接口，定义了一种方法`Expand(apiVersion, resourceName string) []schema.GroupVersionResource`。
   - 其目的是根据给定的apiVersion和resourceName，扩展成一组GroupVersionResource（Kubernetes资源的组、版本和资源类型）。
   - 该接口的实现将特定的apiVersion和resourceName映射为一组GroupVersionResource集合。

2. SimpleCategoryExpander结构体：
   - SimpleCategoryExpander是CategoryExpander的一个简单实现。
   - 它通过将给定的apiVersion和resourceName解析为GroupVersionResource来扩展资源种类。

3. discoveryCategoryExpander结构体：
   - discoveryCategoryExpander是CategoryExpander的另一个实现，它使用DiscoveryClient来发现Kubernetes API服务器上的资源种类。
   - 它通过调用Kubernetes API的Server PreferredResources接口获取到的资源列表来扩展资源种类。
   - 首次调用接口时，会使用discovery.Expander来解析资源并缓存其结果，以提高性能。

4. UnionCategoryExpander结构体：
   - UnionCategoryExpander是CategoryExpander的一个包装器实现。
   - 它将多个CategoryExpander组合在一起，并将它们的结果合并为一个GroupVersionResource集合，以实现资源类型扩展的联合效果。

5. Expand函数：
   - Expand是一个实用函数，用于扩展给定的apiVersion和resourceName为一组GroupVersionResource。
   - 它根据传入的CategoryExpander实例来选择合适的实现并返回扩展结果。

6. NewDiscoveryCategoryExpander函数：
   - NewDiscoveryCategoryExpander是一个构造函数，用于创建一个新的discoveryCategoryExpander实例。

这些结构体和函数提供了一种用于资源映射和扩展的方法，使开发者可以更方便地处理Kubernetes API中的不同资源种类。可以根据特定的需求选择和配置适当的CategoryExpander实现，并使用Expand函数来获取扩展结果。

