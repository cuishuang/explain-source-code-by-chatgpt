# File: metrics/prometheus/collector.go

在go-ethereum项目中，`metrics/prometheus/collector.go`文件的作用是实现Prometheus指标收集器。该文件定义了用于收集和导出系统指标的结构体和方法。

`typeGaugeTpl`、`typeCounterTpl`、`typeSummaryTpl`、`keyValueTpl`和`keyQuantileTagValueTpl`这几个变量是用于指标模板的字符串模板。每个变量都对应一种类型的指标，并用于定义指标的格式和标签。

`collector`结构体是指标收集器的主体，用于保存和管理所有指标。它包含了一个字段`metrics`，是一个`map`类型，用于存储所有已注册的指标。

`newCollector`函数用于创建一个新的指标收集器。它会初始化`metrics`字段，并返回一个指针类型的`collector`结构体。

`addCounter`、`addCounterFloat64`、`addGauge`、`addGaugeFloat64`、`addHistogram`、`addMeter`、`addTimer`、`addResettingTimer`这些函数用于向指标收集器中添加不同类型的指标。这些方法会创建相应类型的指标，并将其存储到`metrics`字段中。

`writeGaugeCounter`、`writeSummaryCounter`、`writeSummaryPercentile`这些函数用于将指标的数值格式化为字符串，并写入到输出流中（通常是HTTP响应）。不同类型的指标有不同的格式化方式和输出逻辑。

`mutateKey`函数用于修改指标的键值（key）。它可以重命名指标或添加额外的标签。

这些函数和变量共同组成了一个完整的指标收集器，负责收集和导出系统中的各种指标。

