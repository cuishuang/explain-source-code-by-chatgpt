# File: plugin/pkg/auth/authorizer/rbac/subject_locator.go

在Kubernetes项目中，`plugin/pkg/auth/authorizer/rbac/subject_locator.go`文件的作用是定位主体（Subject）信息，用于授权验证。

- `_`：这个变量是一个占位符，表示忽略不需要使用的值。在这个文件中，这些占位符被用于忽略不需要的返回值。

- `RoleToRuleMapper`：这个结构体用于将角色（Role）映射到规则（Rule）。它负责解析和映射角色定义的规则，并将其关联到相应的角色名称。

- `SubjectLocator`：这个结构体用于定位主体信息。它包含一些函数和字段，用于根据提供的主体名称和类型查找相应的主体。

- `SubjectAccessEvaluator`：这个结构体用于评估主体的访问权限。它根据规则评估是否允许主体执行某个操作。

- `NewSubjectAccessEvaluator`：这个函数创建一个新的主体访问评估器。它接受一些参数，如主体定位器和角色到规则的映射，创建相应的评估器实例。

- `AllowedSubjects`：这个函数根据指定的资源和操作，返回允许执行该操作的主体列表。它基于主体访问评估器的结果生成允许的主体列表，并排除不符合条件的主体。这个函数通常用于授权验证过程中确定能够执行操作的主体。

