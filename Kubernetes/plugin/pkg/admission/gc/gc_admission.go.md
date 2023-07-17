# File: plugin/pkg/admission/gc/gc_admission.go

在 Kubernetes 项目中，`plugin/pkg/admission/gc/gc_admission.go` 文件的作用是实现资源垃圾回收的逻辑。它是一个 Kubernetes 的 Admission Controller 插件，用于在资源的创建、更新和删除之前进行验证和修改。

以下是关于提到的各个变量和函数的详细介绍：

- `_`：在 Go 语言中，下划线 `_` 被用作一个空标识符，表示一个变量被声明但没有被使用。在这个文件中，`_` 变量是用于表示不使用的某些返回值或参数。
- `gcPermissionsEnforcement`：这是一个全局布尔变量，用于配置是否启用资源垃圾回收的权限强制控制。
- `whiteListItem`：这是一个结构体，用于存储资源垃圾回收的白名单项。它包含资源的 API 组、版本和资源类型等信息。
- `Register`：这个函数用于注册垃圾回收的 Admission Controller 插件。
- `isWhiteListed`：这个函数检查指定的资源是否在白名单中，确定是否进行垃圾回收。
- `Validate`：这个函数用于验证资源对象是否满足垃圾回收的要求。它检查资源对象中是否存在与垃圾回收相关的注释和标签，并进行相应的处理。
- `isChangingOwnerReference`：这个函数检查资源对象是否正在修改其拥有者引用，用于确定是否进行垃圾回收。
- `finalizeAnythingRecord`：这个函数用于标记特定资源对象的垃圾回收状态。
- `ownerRefToDeleteAttributeRecords`：这个函数用于删除指定资源对象的垃圾回收记录。
- `blockingOwnerRefs`：这个函数用于查找和返回拥有者引用的资源对象。
- `indexByUID`：这是一个用于索引资源对象的函数，以便更高效地进行访问和检查。
- `newBlockingOwnerDeletionRefs`：这个函数用于创建拥有者引用的删除记录。
- `SetAuthorizer`：这个函数用于设置垃圾回收的授权器，用于验证用户对资源的操作权限。
- `SetRESTMapper`：这个函数用于设置 REST 映射器，以便将资源对象映射到相应的 API 组、版本和资源类型。
- `ValidateInitialization`：这个函数用于验证垃圾回收插件的初始化，并确保所需的依赖关系和配置正确设置。

这些变量和函数共同协作，实现了 Kubernetes 中资源对象的垃圾回收功能，确保系统中的资源得到有效管理和释放。

