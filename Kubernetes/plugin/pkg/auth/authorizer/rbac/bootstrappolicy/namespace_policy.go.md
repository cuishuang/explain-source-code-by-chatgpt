# File: plugin/pkg/auth/authorizer/rbac/bootstrappolicy/namespace_policy.go

在 Kubernetes 项目中，`plugin/pkg/auth/authorizer/rbac/bootstrappolicy/namespace_policy.go` 文件的作用是初始化和创建命名空间（Namespace）级别的 Role（角色）和 RoleBinding（角色绑定）。

下面是关于这个文件中各个变量和函数的详细介绍：

- `namespaceRoles`：这是一个存储命名空间角色（Role）的切片（slice），用于存储 Role 对象的定义。通过执行 `AddNamespaceRole` 函数，可以向该切片添加新的 Role。

- `namespaceRoleBindings`：这也是一个存储命名空间角色绑定（RoleBinding）的切片，用于存储 RoleBinding 对象的定义。通过执行 `AddNamespaceRoleBinding` 函数，可以向该切片添加新的 RoleBinding。

- `AddNamespaceRole` 函数：用于向 `namespaceRoles` 中添加一个命名空间角色。该函数接收角色的名称和权限规则，并创建一个 Role 对象，然后将其添加到 `namespaceRoles` 中。

- `AddNamespaceRoleBinding` 函数：用于向 `namespaceRoleBindings` 中添加一个命名空间角色绑定。该函数接收绑定名称、绑定的角色名称和绑定的命名空间，并创建一个 RoleBinding 对象，然后将其添加到 `namespaceRoleBindings` 中。

- `init` 函数：在初始化时被调用，用于创建和填充命名空间角色和角色绑定。该函数首先会调用 `AddNamespaceRole` 和 `AddNamespaceRoleBinding` 函数，添加一些默认的命名空间角色和角色绑定。

- `NamespaceRoles` 函数：返回 `namespaceRoles` 切片，用于获取已定义的命名空间角色。

- `NamespaceRoleBindings` 函数：返回 `namespaceRoleBindings` 切片，用于获取已定义的命名空间角色绑定。

通过这些函数和变量，`namespace_policy.go` 文件提供了创建和管理命名空间角色和角色绑定的能力。这些角色和角色绑定可以用于控制和限制某些用户或服务对命名空间中资源的访问和操作权限。

