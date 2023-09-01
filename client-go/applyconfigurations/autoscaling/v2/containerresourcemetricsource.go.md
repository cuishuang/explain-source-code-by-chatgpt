# File: client-go/applyconfigurations/autoscaling/v2beta1/containerresourcemetricsource.go

在`client-go`项目中，`containerresourcemetricsource.go`文件定义了`v2beta1`版本中的容器资源度量源（Container Resource Metric Source）的配置对象和操作函数。

`ContainerResourceMetricSourceApplyConfiguration`是一个配置应用的对象，用于将特定的配置应用于`ContainerResourceMetricSource`对象。它有以下几个结构体：

1. `ContainerResourceMetricSource`：这个结构体定义了容器资源度量源的配置。它包含以下字段：
   - `Name`：度量指标的名称。
   - `TargetAverageUtilization`：目标容器平均利用率的百分比。
   - `TargetAverageValue`：目标容器平均值。

2. `WithName`：这个函数用于设置`ContainerResourceMetricSource`对象的`Name`字段。

3. `WithTargetAverageUtilization`：这个函数用于设置`ContainerResourceMetricSource`对象的`TargetAverageUtilization`字段。

4. `WithTargetAverageValue`：这个函数用于设置`ContainerResourceMetricSource`对象的`TargetAverageValue`字段。

5. `WithContainer`：这个函数用于设置`ContainerResourceMetricSource`对象的`Container`字段，即容器的名称。

这些函数可以对`ContainerResourceMetricSource`对象进行配置和设置，以便在使用容器资源度量源时指定度量指标的名称、目标平均利用率或目标平均值，并关联特定容器。

