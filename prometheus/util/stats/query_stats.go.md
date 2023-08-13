# File: util/stats/query_stats.go

在Prometheus项目中，util/stats/query_stats.go文件是用于记录和统计查询的性能指标和统计数据的工具。

以下是这些结构体的作用：

1. QueryTiming：记录查询的起始时间和消耗的时间。
2. stepStat：用于统计每个步骤（step）的指标。
3. queryTimings：保存查询的历史时间信息。
4. querySamples：保存查询的样本数据。
5. BuiltinStats：内置的统计数据，用于跟踪查询统计信息的历史记录。
6. QueryStats：用于记录查询的统计信息。
7. SpanTimer：计时器，用于跟踪子查询的耗时。
8. Statistics：保存查询的统计数据。
9. QueryTimers：保存查询的耗时记录。
10. TotalSamplesPerStep：记录每个步骤的样本数量。
11. QuerySamples：保存查询的样本数据。
12. Stats：查询的统计信息。

以下是这些函数的作用：

1. String：将结构体转换为可打印的字符串。
2. SpanOperation：用于创建一个新的SpanTimer。
3. MarshalJSON：将结构体序列化为JSON格式。
4. Builtin：获取内置的统计信息。
5. NewQueryStats：创建一个新的QueryStats结构体。
6. TotalSamplesPerStepMap：获取记录每个步骤样本数量的映射。
7. totalSamplesPerStepPoints：获取每个步骤样本数量的点。
8. NewSpanTimer：创建一个新的SpanTimer。
9. Finish：计算计时器的总时间和最大时间。
10. InitStepTracking：初始化步骤跟踪。
11. IncrementSamplesAtStep：在指定步骤增加样本数量。
12. IncrementSamplesAtTimestamp：在指定时间戳增加样本数量。
13. UpdatePeak：更新峰值。
14. UpdatePeakFromSubquery：从子查询更新峰值。
15. NewQueryTimers：创建一个新的QueryTimers结构体。
16. NewQuerySamples：创建一个新的QuerySamples结构体。
17. NewChild：创建一个新的SpanTimer子计时器。
18. GetSpanTimer：获取计时器的时间。

