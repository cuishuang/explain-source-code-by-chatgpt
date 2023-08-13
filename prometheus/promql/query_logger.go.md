# File: promql/query_logger.go

promql/query_logger.go文件是Prometheus项目中的一个组件，用于记录和跟踪PromQL查询的执行情况。该文件中定义了ActiveQueryTracker和Entry两个结构体，以及一些相关的函数。

1. ActiveQueryTracker结构体：用于跟踪和管理活动的查询，包括查询的开始时间、状态以及相关的元数据等。

2. Entry结构体：代表一个查询日志条目，包括查询的ID、类型（Instant/Range）、开始时间、持续时间、实例（instance）等信息。

3. parseBrokenJSON函数：用于解析损坏的JSON数据。

4. logUnfinishedQueries函数：记录未完成的查询日志。

5. getMMapedFile函数：获取内存映射文件。

6. NewActiveQueryTracker函数：创建一个新的ActiveQueryTracker实例。

7. trimStringByBytes函数：按照字节截取字符串。

8. _newJSONEntry函数：创建一个新的JSON格式的日志条目。

9. newJSONEntry函数：创建一个新的JSON格式的日志条目，并将其写入文件。

10. generateIndices函数：生成索引。

11. GetMaxConcurrent函数：获取最大并发查询数。

12. Delete函数：从ActiveQueryTracker中删除指定的查询。

13. Insert函数：向ActiveQueryTracker中插入一个新的查询。

这些函数和结构体的作用是为了实现Prometheus的查询跟踪和日志记录功能。通过跟踪和记录查询的执行情况，可以对系统性能进行监控和分析，并及时发现和解决潜在的性能问题。

