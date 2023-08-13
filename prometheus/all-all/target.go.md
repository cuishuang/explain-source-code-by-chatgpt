# File: scrape/target.go

在Prometheus项目中，scrape/target.go文件的作用是定义了Target结构体以及与之相关的函数和方法，用于管理和处理被抓取的目标（target）的信息。

- errSampleLimit和errBucketLimit是两个错误变量，分别表示样本限制错误和桶限制错误，用于在抓取目标时处理相关错误。

- TargetHealth是一个表示目标健康状态的枚举类型。
- Target是代表被抓取的目标的结构体，包含了目标的URL、健康状态、最近一次抓取的时间和持续时间等信息。
- MetricMetadataStore是一个接口，用于存储指标元数据。
- MetricMetadata是指标元数据的结构体，包含了指标的名称、帮助信息、标签等。
- Targets是一个Target类型的切片，表示多个目标。
- limitAppender、timeLimitAppender和bucketLimitAppender是用于将限制应用于被抓取的指标样本集合的函数，用于限制样本数量、时间和桶的数量。
- NewTarget是创建一个新的Target对象的函数。
- String是Target结构体的String方法，用于将Target对象转换为字符串表示。
- MetadataList是MetricMetadataStore接口的一个方法，用于获取所有指标的元数据列表。
- MetadataSize和MetadataLength是MetricMetadataStore接口的两个方法，分别用于获取指标元数据的大小和长度。
- Metadata是MetricMetadataStore接口的一个方法，用于根据指标名称获取指标的元数据。
- SetMetadataStore是给Target结构体添加元数据存储功能的方法。
- hash、offset、Labels、LabelsRange、DiscoveredLabels和SetDiscoveredLabels是用于处理和操作标签的函数和方法。
- URL、Report、LastError、LastScrape、LastScrapeDuration、Health、intervalAndTimeout、GetValue、Len、Less、Swap、Append、AppendHistogram、PopulateLabels和TargetsFromGroup是用于操作和管理Target结构体的函数和方法，包括获取URL、报告、最近抓取的错误和时间、健康状态、获取值、长度、排序等。

