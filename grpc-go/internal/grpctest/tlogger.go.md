# File: grpc-go/internal/grpctest/tlogger.go

grpc-go/internal/grpctest/tlogger.go文件的作用是定义了一个用于测试的日志记录器（logger）。

TLogger中的变量：

- logType：用于标识日志类型的枚举类型，包括InfoLog、WarningLog、ErrorLog和FatalLog。
- tLogger：保存日志的结构体，包含了日志级别、计数器和记录的日志消息。

logType枚举类型：定义了不同的日志类型。

tLogger结构体：用于保存日志的结构体，包含了以下字段：

- level：日志级别，可以是Info，Warning，Error和Fatal。
- counts：用于计数各种日志类型的计数器。
- logs：保存实际的日志消息。

TLogger提供以下方法：

- String：将TLogger的内容格式化为字符串表示。
- init：初始化一个新的TLogger实例。
- getCallingPrefix：获取调用者的信息，用于显示在日志中。
- log：将日志消息添加到日志记录器中。
- Update：根据给定的更新函数更新日志记录器。
- ExpectError：检查是否有错误发生，如果有，记录错误并返回错误消息。
- ExpectErrorN：检查指定数量的错误是否发生，如果超出数量，记录错误并返回错误消息。
- EndTest：在测试结束时调用，检查是否有未记录的日志。
- expected：获取记录器中预期的日志条目数。
- Info：记录一条Info级别的日志。
- Infoln：记录一条带换行符的Info级别的日志。
- Infof：格式化并记录一条Info级别的日志。
- InfoDepth：记录带有调用者信息和Info级别的日志。
- Warning：记录一条Warning级别的日志。
- Warningln：记录一条带换行符的Warning级别的日志。
- Warningf：格式化并记录一条Warning级别的日志。
- WarningDepth：记录带有调用者信息和Warning级别的日志。
- Error：记录一条Error级别的日志，并返回错误消息。
- Errorln：记录一条带换行符的Error级别的日志，并返回错误消息。
- Errorf：格式化并记录一条Error级别的日志，并返回错误消息。
- ErrorDepth：记录带有调用者信息和Error级别的日志，并返回错误消息。
- Fatal：记录一条Fatal级别的日志，并终止测试。
- Fatalln：记录一条带换行符的Fatal级别的日志，并终止测试。
- Fatalf：格式化并记录一条Fatal级别的日志，并终止测试。
- FatalDepth：记录带有调用者信息和Fatal级别的日志，并终止测试。
- V：根据给定的日志级别记录日志。

