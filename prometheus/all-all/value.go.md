# File: promql/value.go

在Prometheus项目中，promql/value.go文件是用于定义PromQL表达式中数据类型的文件。

以下是这些结构体的作用：

- String：表示一个字符串值。
- Scalar：表示一个标量值，即单个时间序列的值。
- Series：表示一个时间序列。
- FPoint：表示一个浮点数时间序列的样本。
- HPoint：表示一个直方图时间序列的样本。
- Sample：表示一个样本，包括时间戳和值。
- Vector：表示一个向量，包含多个时间序列。
- Matrix：表示一个矩阵，包含多个向量。
- Result：表示一个查询结果，包含标签和值。
- StorageSeries：表示一个持久化的时间序列。
- storageSeriesIterator：表示一个用于迭代持久化时间序列的迭代器。

以下是这些函数的作用：

- Type：返回给定值的类型。
- String：返回给定值的字符串表示。
- MarshalJSON：将给定值转换为JSON字符串。
- ContainsSameLabelset：检查两个Result是否具有相同的标签集。
- TotalSamples：返回给定矩阵中的总样本数。
- Len：返回给定向量的长度。
- Less：比较两个向量的元素是否小于。
- Swap：交换两个向量的位置。
- Vector：返回给定向量中指定索引处的值。
- Matrix：返回给定矩阵中指定索引处的向量。
- Scalar：返回给定标量中的值。
- NewStorageSeries：创建一个新的持久化时间序列。
- Labels：返回给定持久化时间序列的标签集。
- Iterator：返回给定持久化时间序列的迭代器。
- newStorageSeriesIterator：创建一个新的持久化时间序列的迭代器。
- reset：重置迭代器的状态。
- Seek：将迭代器指向指定时间戳之前的最后一个样本。
- At：返回迭代器当前位置的样本。
- AtHistogram：返回迭代器当前位置的直方图样本。
- AtFloatHistogram：返回迭代器当前位置的浮点直方图样本。
- AtT：返回迭代器当前位置的时间戳。
- Next：将迭代器移动到下一个样本。
- Err：返回迭代器的错误。

