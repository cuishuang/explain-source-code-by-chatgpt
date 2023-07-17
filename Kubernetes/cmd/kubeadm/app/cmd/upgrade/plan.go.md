# File: cmd/kubeadm/app/cmd/upgrade/plan.go

在kubeadm项目中，cmd/kubeadm/app/cmd/upgrade/plan.go文件的作用是实现"kubeadm upgrade plan"命令，用于生成当前集群的升级计划。它基于当前集群的状态和配置信息，分析出需要升级的组件、版本号及其它相关信息，并以可读的方式打印出来。

以下是对每个相关的结构体和函数的详细介绍：

结构体：
1. planFlags：存储"kubeadm upgrade plan"命令的命令行标志参数。
2. upgradePlanPrintFlags：用于打印升级计划的命令行标志参数，控制打印方式的格式。
3. upgradePlanJSONYamlPrintFlags：用于以JSON或YAML格式打印升级计划的命令行标志参数。
4. upgradePlanJSONYAMLPrinter：将升级计划打印为JSON或YAML格式。
5. upgradePlanTextPrinter：将升级计划以可读的文本格式打印。
6. upgradePlanTextPrintFlags：用于以文本格式打印升级计划的命令行标志参数。

函数：
1. newCmdPlan：创建一个"kubeadm upgrade plan"命令的cobra.Command对象。
2. newComponentUpgradePlan：创建一个新的ComponentUpgradePlan对象，记录一个组件的升级信息。
3. newUpgradePlanPrintFlags：创建一个UpgradePlanPrintFlags对象，用于记录打印升级计划时的标志参数。
4. AllowedFormats：返回支持的打印格式列表。
5. AddFlags：为"kubeadm upgrade plan"命令的cobra.Command对象添加命令行标志参数。
6. ToPrinter：根据UpgradePlanPrintFlags指定的打印格式，返回相应的打印器。
7. newUpgradePlanJSONYAMLPrinter：创建一个新的UpgradePlanJSONYAMLPrinter对象，用于将升级计划打印为JSON或YAML格式。
8. PrintObj：将对象以JSON或YAML格式打印到指定的Writer中。
9. Flush：刷新升级计划的打印器。
10. Close：关闭升级计划的打印器。
11. runPlan：执行"kubeadm upgrade plan"命令的实际操作。
12. appendDNSComponent：将DNS组件的升级信息添加到升级计划中。
13. genUpgradePlan：生成当前集群的升级计划，包括各组件的版本信息。
14. getComponentConfigVersionStates：获取集群中的各组件的配置和版本信息。
15. printUpgradePlan：打印升级计划。
16. sortedSliceFromStringIntMap：将以字符串为键、整数为值的映射按键排序并返回切片。
17. strOrDash：如果字符串非空，则返回字符串，否则返回"-"。
18. yesOrNo：如果布尔值为true，则返回"Yes"，否则返回"No"。
19. printLineSeparator：打印分隔线。
20. printComponentConfigVersionStates：打印组件的配置和版本信息。

