# File: tsdb/record/record.go

tsdb/record/record.go文件是Prometheus项目中的一个文件，用于定义和处理时间序列数据的记录。

文件中定义了一些常量，如ErrNotFound，它用来表示当在操作记录时找不到特定项时的错误。

另外还定义了一些结构体：

- Type结构体是一个枚举类型，用于表示记录的类型，包括Metric，Sample，Metadata等。

- MetricType结构体表示时间序列数据的指标类型，如Counter，Gauge，Histogram等。

- RefSeries结构体表示引用时间序列数据的引用。

- RefSample结构体表示引用时间序列数据样本的引用。

- RefMetadata结构体表示引用时间序列数据元数据的引用。

- RefExemplar结构体表示引用时间序列数据样本示例的引用。

- RefHistogramSample结构体表示引用直方图样本的引用。

- RefFloatHistogramSample结构体表示引用浮点数直方图样本的引用。

- RefMmapMarker结构体表示引用内存映射标记的引用。

- Decoder结构体用于解码记录数据。

- Encoder结构体用于编码记录数据。

此外，文件中还定义了一些函数：

- String函数用于将记录的类型转换为字符串表示。

- GetMetricType函数用于获取记录中的指标类型。

- ToTextparseMetricType函数用于将字符串表示的指标类型转换为MetricType类型。

- Type函数用于获取记录的类型。

- Series函数用于获取时间序列数据的引用。

- Metadata函数用于获取时间序列数据元数据的引用。

- DecodeLabels函数用于解码记录中的标签。

- Samples函数用于获取时间序列数据的样本。

- Tombstones函数用于获取时间序列数据的删除标记。

- Exemplars函数用于获取时间序列数据的示例。

- ExemplarsFromBuffer函数用于从字节缓冲中解码示例。

- MmapMarkers函数用于获取内存映射标记。

- HistogramSamples函数用于获取直方图样本。

- FloatHistogramSamples函数用于获取浮点数直方图样本。

- EncodeLabels函数用于编码标签。

- EncodeExemplarsIntoBuffer函数用于将示例编码到字节缓冲中。

这些函数用于处理和操作记录数据的不同部分，如获取和解码时间序列数据的属性和子元素等。

