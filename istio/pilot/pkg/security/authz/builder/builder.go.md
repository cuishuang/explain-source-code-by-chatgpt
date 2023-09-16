# File: istio/pilot/pkg/security/authz/builder/builder.go

在Istio项目中，istio/pilot/pkg/security/authz/builder/builder.go文件的作用是构建和管理Istio的授权策略。该文件定义了一组API来帮助开发人员生成和配置与授权相关的配置项。

在该文件中，rbacPolicyMatchNever变量是一个仅包含字符串"match-never"的常量，它用于定义RBAC策略中的一个特殊值，表示永远不匹配。

Option结构体是用于配置授权策略的选项，它包含一些字段，例如是否启用RBAC、是否启用委派等。

Builder结构体是用于构建授权规则的主要结构。它有一个字段用于配置选项，还有其他私有字段用于存储配置中间结果，如构建的HTTP规则和TCP规则等。

builtConfigs结构体保存了已经构建的授权策略的配置。它有两个字段，一个用于HTTP规则，一个用于TCP规则。

New函数是用于创建一个新的Builder实例。

BuildHTTP函数是用于根据给定的配置生成HTTP授权规则。

BuildTCP函数是用于根据给定的配置生成TCP授权规则。

isDryRun函数用于判断当前是否为干扰模式（dry run），干扰模式下策略修改不会被应用。

shadowRuleStatPrefix函数用于返回一个带有阴影规则统计前缀的字符串。

build函数用于构建授权规则，并将结果保存到builtConfigs中。

buildHTTP函数用于构建HTTP授权规则，并将结果保存到builtConfigs中。

buildTCP函数用于构建TCP授权规则，并将结果保存到builtConfigs中。

policyName函数用于生成一个唯一的策略名称。

这些函数在builder.go文件中的实现，通过组合不同的配置选项和参数，来生成和管理与授权策略相关的配置，以提供授权功能。


