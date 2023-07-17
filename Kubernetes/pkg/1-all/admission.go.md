# File: pkg/kubeapiserver/options/admission.go

在kubernetes项目中，pkg/kubeapiserver/options/admission.go这个文件的作用是为kube-apiserver提供命令行选项和配置验证，以及为插件提供注册和管理。

AdmissionOptions中定义了一组选项，用于配置和控制Admission插件。这些选项包括：

- Plugins：定义了启用的Admission插件列表。
- PluginConfigFile：指定Admission插件的配置文件路径。
- Enabled：指示是否启用Admission插件。
- Initializers：定义了在请求流入之前要应用的初始化器。

接下来，让我们逐个介绍每个相关的函数：

1. NewAdmissionOptions: 该函数用于创建一个新的Admission插件选项配置，并返回AdmissionOptions结构体的实例。

2. AddFlags: 该函数用于将Admission插件选项的命令行标志添加到给定的FlagSet中。这样，用户在运行kube-apiserver时就可以使用这些标志进行配置。

3. Validate: 该函数用于验证Admission插件选项的有效性。它检查选项是否符合一些基本的约束，例如插件列表中是否包含有效的插件名称。

4. ApplyTo: 该函数用于将Admission插件的配置应用到给定的插件集。它遍历所有插件，将配置应用到每个插件。

5. computePluginNames: 该函数用于从插件列表中解析插件名称。它先从插件列表中获取插件名称，然后根据指定的策略（例如ALLOW和DENY）来解析插件名称。

总之，pkg/kubeapiserver/options/admission.go文件定义了Admission插件选项的配置和管理，提供了用于创建、添加标志、验证配置、应用配置和解析插件名称等一系列函数。这些函数是为了方便用户对Admission插件进行配置和管理。

