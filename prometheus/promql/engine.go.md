# File: promql/engine.go

在Prometheus项目中，promql/engine.go文件负责实现Prometheus查询语言（PromQL）的查询引擎。它将用户提交的查询语句解析、验证和执行，并返回查询结果。

在该文件中，ErrValidationAtModifierDisabled、ErrValidationNegativeOffsetDisabled、fPointPool、hPointPool这几个变量的作用如下：
- ErrValidationAtModifierDisabled和ErrValidationNegativeOffsetDisabled是用于在查询中禁用At修饰符和负偏移量。当用户的查询中包含这些被禁用的修饰符时，引擎会返回相应的错误。
- fPointPool和hPointPool是用于复用内存的对象池。它们用于减少内存分配的开销，提高执行效率。

以下是engineMetrics, ErrQueryTimeout, QueryLogger, Query, QueryOpts, query, QueryOrigin, QueryTracker, EngineOpts, Engine, errWithWarnings, evaluator, EvalSeriesHelper, EvalNodeHelper, groupedAggregation这几个结构体的功能概述：

- engineMetrics: 包含引擎的度量信息，用于统计查询执行的指标。
- ErrQueryTimeout: 表示查询超时时返回的错误。
- QueryLogger: 用于记录查询的日志。
- Query: 表示一个完整的查询，包括查询语句和查询参数。
- QueryOpts: 包含查询的各种选项，如超时时间、报警规则等。
- query: 表示正在执行的查询。
- QueryOrigin: 表示查询的来源，包含了查询的用户、地址等信息。
- QueryTracker: 用于跟踪查询的状态，比如查询的开始时间、结束时间等。
- EngineOpts: 包含引擎的各种选项，如记录日志、缓存配置等。
- Engine: 表示查询引擎，用于执行和管理查询。
- errWithWarnings: 表示带有警告的错误。
- evaluator: 查询的评估器，用于执行实际的查询操作。
- EvalSeriesHelper: 辅助数据结构，用于帮助评估系列选择器表达式。
- EvalNodeHelper: 辅助数据结构，用于帮助评估节点表达式。
- groupedAggregation: 表示分组聚合的辅助数据结构。

以下是一些常用的函数及其作用：

- convertibleToInt64: 将通用类型转换为int64类型。
- Error: 生成一个错误。
- Statement: 表示一个查询语句。
- String: 将表达式转换为字符串。
- Stats: 表示查询的统计信息。
- Cancel: 取消执行中的查询。
- Close: 关闭引擎。
- Exec: 执行查询。
- contextDone: 检查查询上下文是否已完成。
- contextErr: 获取查询上下文中的错误。
- NewEngine: 创建一个新的查询引擎。
- SetQueryLogger: 设置查询日志记录器。
- NewInstantQuery: 创建一个新的即时查询。
- NewRangeQuery: 创建一个新的范围查询。
- newQuery: 创建一个新的查询。
- validateOpts: 验证查询选项。
- newTestQuery: 创建用于测试的查询。
- exec: 执行查询操作。
- queueActive: 将查询添加到活动查询队列中。
- timeMilliseconds: 将时间转换为毫秒。
- durationMilliseconds: 计算时间间隔的毫秒数。
- execEvalStmt: 执行评估的语句。
- subqueryTimes: 获取子查询的时间范围。
- findMinMaxTime: 查找指定向量的最大和最小时间戳。
- getTimeRangesForSelector: 获取选择器的时间范围。
- getLastSubqueryInterval: 获取最后一个子查询的时间间隔。
- populateSeries: 填充原始系列数据。
- extractFuncFromPath: 从路径中提取函数信息。
- extractGroupsFromPath: 从选择器路径中提取分组信息。
- checkAndExpandSeriesSet: 检查并扩展系列集。
- expandSeriesSet: 扩展系列集。
- errorf: 格式化输出错误信息。
- error: 生成一个错误。
- recover: 恢复从panic中恢复执行。
- Eval: 执行一个查询。
- resetBuilder: 重置查询构建器。
- DropMetricName: 移除指定向量中的度量名称。
- rangeEval: 范围查询的评估。
- evalSubquery: 执行子查询。
- eval: 评估查询语句。
- vectorSelector: 向量选择器。
- vectorSelectorSingle: 向量选择器（单个）。
- getFPointSlice、putFPointSlice、getHPointSlice、putHPointSlice: 获取和释放具有特定类型的切片。
- matrixSelector: 矩阵选择器。
- matrixIterSlice: 矩阵选择器的迭代器切片。
- VectorAnd、VectorOr、VectorUnless: 向量操作函数。
- VectorBinop: 向量二元操作函数。
- signatureFunc: 函数签名。
- resultMetric: 一个度量补丁。
- VectorscalarBinop: 向量与标量的二元操作函数。
- dropMetricName: 移除向量中的度量名称。
- scalarBinop: 标量二元操作函数。
- vectorElemBinop: 向量元素二元操作函数。
- aggregation: 聚合函数。
- generateGroupingKey: 生成分组键。
- btos: 将布尔值转为字符串。
- shouldDropMetricName: 判断是否需要移除度量名称。
- NewOriginContext: 创建一个新的查询原始上下文。
- formatDate: 格式化日期。
- unwrapParenExpr: 解析括号表达式。
- unwrapStepInvariantExpr: 解析步骤不变表达式。
- PreprocessExpr: 预处理表达式。
- preprocessExprHelper: 辅助函数，用于预处理表达式。
- newStepInvariantExpr: 创建步骤不变的表达式。
- setOffsetForAtModifier: 为At修饰符设置偏移量。
- makeInt64Pointer：创建一个int64类型的指针。

