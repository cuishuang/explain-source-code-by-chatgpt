# File: promql/parser/ast.go

在Prometheus项目中，promql/parser/ast.go文件定义了PromQL查询语言的抽象语法树（Abstract Syntax Tree, AST）。该文件包含了一系列结构体和函数，用于表示和操作PromQL语句的语法结构。

- Node：表示AST中的节点，是所有结构体的基础类型。
- Statement：表示一个完整的查询语句，包含多个表达式。
- EvalStmt：表示一个待计算的查询语句。
- Expr：表示一个表达式，可以是单独的表达式或者多个表达式的组合。
- Expressions：表示一个表达式列表，用于存储多个表达式。
- AggregateExpr：表示一个聚合表达式，用于对多个时间序列进行聚合操作。
- BinaryExpr：表示一个二元表达式，包含左右两个操作数和操作符。
- Call：表示一个函数调用表达式。
- MatrixSelector：表示一个矩阵选择器，用于选择多个时间序列。
- SubqueryExpr：表示一个子查询表达式，用于在查询中执行子查询操作。
- NumberLiteral：表示一个数值字面量。
- ParenExpr：表示一个括号表达式，用于控制运算顺序。
- StringLiteral：表示一个字符串字面量。
- UnaryExpr：表示一个一元表达式，包含一个操作数和操作符。
- StepInvariantExpr：表示一个时刻不变表达式，用于在查询中标记时刻不变条件。
- VectorSelector：表示一个向量选择器，用于选择一个或多个时间序列。
- TestStmt：表示一个测试语句，用于测试查询语句。
- VectorMatchCardinality：表示向量匹配的基数。
- VectorMatching：表示向量匹配方式的枚举类型。
- Visitor：用于访问和操作AST中的节点。
- Inspector：用于检查AST中的节点。
- PositionRange：表示节点在源代码中的位置范围。

以下是一些关键函数的作用：

- PromQLStmt：将AST节点转换为PromQL语句字符串。
- String：将AST节点转换为字符串。
- PositionRange：标识节点在源代码中的位置范围。
- Pretty：将AST节点以漂亮的形式打印出来，方便人类阅读。
- Type：获取AST节点的类型。
- PromQLExpr：将AST节点转换为PromQL查询表达式字符串。
- Walk：遍历AST节点，并调用指定的函数进行处理。
- ExtractSelectors：提取AST节点中的选择器。
- Visit：访问AST节点，并调用指定的函数进行处理。
- Inspect：检查AST节点，并调用指定的函数进行处理。
- Children：获取AST节点的子节点。
- mergeRanges：合并多个位置范围。

