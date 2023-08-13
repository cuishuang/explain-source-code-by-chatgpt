# File: storage/interface.go

storage/interface.go文件是Prometheus项目中定义存储接口的文件。它定义了与时间序列数据和指标元数据交互的方法和结构体。

以下是每个变量的作用：

- `ErrNotFound`：表示请求找不到数据的错误。
- `ErrOutOfOrderSample`：表示样本的时间戳顺序错误的错误。
- `ErrOutOfBounds`：表示请求的时间范围超出了存储范围的错误。
- `ErrTooOldSample`：表示样本时间戳太旧以至于存储已经将其丢弃的错误。
- `ErrDuplicateSampleForTimestamp`：表示给定时间戳下有重复样本的错误。
- `ErrOutOfOrderExemplar`：表示示例的时间戳顺序错误的错误。
- `ErrDuplicateExemplar`：表示给定时间戳下有重复示例的错误。
- `ErrExemplarLabelLength`：表示示例标签长度无效的错误。
- `ErrExemplarsDisabled`：表示存储不支持示例的错误。
- `ErrNativeHistogramsDisabled`：表示存储不支持直方图的错误。
- `ErrHistogramCountNotBigEnough`：表示直方图计数太小的错误。
- `ErrHistogramNegativeBucketCount`：表示直方图桶计数为负数的错误。
- `ErrHistogramSpanNegativeOffset`：表示直方图范围负偏移量的错误。
- `ErrHistogramSpansBucketsMismatch`：表示直方图范围和桶计数不匹配的错误。
- `emptySeriesSet`：表示一个空的SeriesSet。
- `emptyChunkSeriesSet`：表示一个空的ChunkSeriesSet。

以下是每个结构体的作用：

- `SeriesRef`：表示一个时间序列的引用。
- `Appendable`：表示可以添加样本或示例的接口。
- `SampleAndChunkQueryable`：表示可以查询样本和分块的接口。
- `Storage`：表示存储引擎的接口。
- `ExemplarStorage`：表示示例存储引擎的接口。
- `Queryable`：表示可查询存储的接口。
- `MockQueryable`：表示用于测试目的的查询接口。
- `Querier`：表示查询器的接口。
- `MockQuerier`：表示用于测试目的的查询器接口。
- `ChunkQueryable`：表示可查询分块的接口。
- `ChunkQuerier`：表示分块查询器的接口。
- `LabelQuerier`：表示标签查询器的接口。
- `ExemplarQueryable`：表示可查询示例的接口。
- `ExemplarQuerier`：表示示例查询器的接口。
- `SelectHints`：表示查询的选项。
- `QueryableFunc`：表示一个可查询存储的函数。
- `Appender`：表示可以追加样本的接口。
- `GetRef`：表示获取时间序列引用的方法。
- `ExemplarAppender`：表示可以追加示例的接口。
- `HistogramAppender`：表示可以追加直方图的接口。
- `MetadataUpdater`：表示更新元数据的接口。
- `SeriesSet`：表示一组时间序列的接口。
- `testSeriesSet`：表示用于测试目的的时间序列集合。
- `errSeriesSet`：表示返回错误的时间序列集合。
- `errChunkSeriesSet`：表示返回错误的分块时间序列集合。
- `Series`：表示一个时间序列的结构体。
- `mockSeries`：表示用于测试目的的时间序列。
- `ChunkSeriesSet`：表示一组分块时间序列的接口。
- `ChunkSeries`：表示一个分块时间序列的结构体。
- `Labels`：表示一组标签的结构体。
- `SampleIterable`：表示可迭代样本的接口。
- `ChunkIterable`：表示可迭代分块的接口。
- `Warnings`：表示一组警告的结构体。

以下是每个函数的作用：

- `Querier`：返回一个新的查询接口。
- `LabelValues`：返回指定标签的可能值。
- `LabelNames`：返回所有存在的标签名称。
- `Close`：关闭存储引擎。
- `Select`：执行给定查询选项的查询并返回结果迭代器。
- `EmptySeriesSet`：返回一个空的时间序列集合。
- `Next`：返回下一个时间序列。
- `At`：返回给定时间戳的样本和示例迭代器。
- `Err`：返回迭代过程中的错误。
- `Warnings`：返回迭代过程中的警告。
- `TestSeriesSet`：返回用于测试目的的时间序列集合。
- `ErrSeriesSet`：返回带有错误的时间序列集合。
- `EmptyChunkSeriesSet`：返回一个空的分块时间序列集合。
- `ErrChunkSeriesSet`：返回带有错误的分块时间序列集合。
- `Labels`：返回时间序列的标签。
- `Iterator`：返回一个迭代器接口。
- `MockSeries`：返回用于测试目的的时间序列结构体。

