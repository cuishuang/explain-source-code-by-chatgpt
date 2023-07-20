# File: metrics/runtimehistogram.go

文件runtimehistogram.go定义了一个运行时直方图的实现。

1. 结构体：
   - runtimeHistogram: 运行时直方图结构体，包含了直方图的最小值、最大值、桶的数量等信息。
   - runtimeHistogramSnapshot: 直方图快照结构体，用于保存直方图的快照，以便进行统计分析。
   - floatsAscendingKeepingIndex: 用于排序浮点数切片，并记录排序后的索引。
   - floatsByIndex: 带有索引的浮点数切片。

2. 函数：
   - getOrRegisterRuntimeHistogram: 获取或注册运行时直方图的函数。
   - newRuntimeHistogram: 创建一个新的运行时直方图实例。
   - update: 更新直方图的函数，将新的值添加到直方图中。
   - load: 从直方图快照加载数据。
   - Clear: 清空直方图。
   - Update: 更新直方图。
   - Sample: 从直方图中随机采样。
   - Snapshot: 创建直方图的快照。
   - Count: 统计直方图中值的数量。
   - Mean: 计算直方图中值的平均值。
   - StdDev: 计算直方图中值的标准差。
   - Variance: 计算直方图中值的方差。
   - Percentile: 计算给定百分位数下的值。
   - Percentiles: 计算给定百分位数列表下的值。
   - Max: 计算直方图中的最大值。
   - Min: 计算直方图中的最小值。
   - Sum: 计算直方图中所有值的总和。
   - mean: 计算切片中所有浮点数的平均值。
   - midpoint: 计算切片中浮点数的中点。
   - computePercentiles: 计算给定百分位数下的值。
   - Len: 获取切片的长度。
   - Less: 判断两个切片中的浮点数是否满足从小到大的排序关系。
   - Swap: 交换两个切片中的浮点数。

