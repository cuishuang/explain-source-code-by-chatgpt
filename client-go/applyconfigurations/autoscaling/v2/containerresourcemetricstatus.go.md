# File: client-go/applyconfigurations/autoscaling/v2beta1/containerresourcemetricstatus.go

文件`containerresourcemetricstatus.go`是`client-go`项目中`autoscaling/v2beta1`包下的文件，它提供了与自动伸缩相关的容器资源度量状态的配置。

作用概述：
该文件定义了用于配置容器资源度量状态的结构体和方法，通过这些结构体和方法可以设置容器的名称、当前平均利用率、当前平均值以及特定容器的度量指标。

作用详细介绍：
1. `ContainerResourceMetricStatusApplyConfiguration`结构体：该结构体用于配置容器资源度量状态的应用配置，通过设置该配置可以创建或更新自动伸缩相关的容器资源度量状态。

2. `ContainerResourceMetricStatus`结构体：该结构体表示容器资源度量状态，包括容器名称、当前平均利用率和当前平均值。

3. `WithName`方法：该方法用于设置容器的名称。可以使用`WithName(name string)`来指定容器的名称。

4. `WithCurrentAverageUtilization`方法：该方法用于设置当前平均利用率。可以使用`WithCurrentAverageUtilization(utilization int32)`来指定当前平均利用率的数值。

5. `WithCurrentAverageValue`方法：该方法用于设置当前平均值。可以使用`WithCurrentAverageValue(value *resource.Quantity)`来指定当前平均值。

6. `WithContainer`方法：该方法用于设置特定容器的度量指标。可以使用`WithContainer(container string)`来指定特定容器的度量指标。

这些方法和结构体的组合使用可以对容器资源度量状态进行灵活配置，以满足不同的自动伸缩需求。

