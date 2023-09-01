# File: client-go/applyconfigurations/policy/v1beta1/poddisruptionbudgetspec.go

在client-go的applyconfigurations/policy/v1beta1/poddisruptionbudgetspec.go文件中，定义了PodDisruptionBudgetSpecApplyConfiguration结构体和一些用于配置PodDisruptionBudgetSpec的函数。

PodDisruptionBudgetSpec用于配置PodDisruptionBudget的规范，可以指定最小可用的Pod数量、选择器以选择要受限制的Pod、最大不可用的Pod数量以及当前不健康的Pod的驱逐策略。

PodDisruptionBudgetSpecApplyConfiguration结构体是一个用于应用配置的中间结构体，它允许用户对PodDisruptionBudgetSpec进行配置，并在应用时将其转换为底层的PodDisruptionBudgetSpec结构体。

以下是几个主要函数的作用：

1. WithMinAvailable：指定PodDisruptionBudget需要保证的最小可用的Pod数量。可以按百分比（如"50%"）或绝对值（如"3"）来定义。默认为nil，表示没有最小可用数量要求。

2. WithSelector：指定选择器，用于选择要受限制的Pod。可以使用标签选择器（如"app=nginx"）或字段选择器（如"metadata.name=nginx-pod"）进行定义。默认为nil，表示选择所有的Pod。

3. WithMaxUnavailable：指定PodDisruptionBudget允许的最大不可用的Pod数量。可以按百分比（如"50%"）或绝对值（如"3"）来定义。默认为nil，表示没有最大不可用数量限制。

4. WithUnhealthyPodEvictionPolicy：指定当前不健康的Pod的驱逐策略。有两个可选值：EvictionPolicyDelete（删除不健康的Pod）和EvictionPolicyWait（等待不健康的Pod变得可用）。默认为nil，表示使用默认的驱逐策略。

这些函数用于配置PodDisruptionBudgetSpec，可以根据具体需求设置所需的限制和策略，以确保应用程序在进行节点维护或故障转移时保持可用性。

