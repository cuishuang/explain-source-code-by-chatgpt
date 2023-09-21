# File: grpc-go/grpclog/glogger/glogger.go

在grpc-go项目中，glogger.go文件是grpc-go库自定义的日志记录器，提供了用于记录日志的函数和结构体。

glogger.go文件中定义了三个结构体：glogger，logEntry和severity，以及一些全局变量。

1. glogger结构体：这是grpc-go库的日志记录器，包含一个Logger接口。glogger结构体实现了grpclog.Logger接口，并提供了一组用于记录日志的函数。它封装了grpclog包的原始日志记录器，并在此基础上进行了一些自定义扩展。

2. logEntry结构体：这是日志条目的结构体，包含日志的级别、时间戳、文件名、行号等信息。

3. severity结构体：这是日志级别的枚举类型，定义了不同的级别：info、warning、error和fatal。

下面是glogger中的几个重要函数的作用：

1. init：这是初始化glogger的函数，它在glogger包被导入时自动执行，用于设置默认的日志级别和输出位置等配置。

2. Info、Infoln、Infof：这些函数用于记录信息级别的日志。Info函数用于记录一条信息日志，Infoln函数用于记录带换行符的信息日志，Infof函数用于记录格式化信息日志。

3. Warning、Warningln、Warningf：这些函数用于记录警告级别的日志，与信息级别的日志函数类似。

4. Error、Errorln、Errorf：这些函数用于记录错误级别的日志，与信息级别的日志函数类似。

5. Fatal、Fatalln、Fatalf：这些函数用于记录致命级别的日志，并调用os.Exit(1)终止程序。Fatal函数用于记录一条致命日志，Fatalln函数用于记录带换行符的致命日志，Fatalf函数用于记录格式化致命日志。

6. InfoDepth、WarningDepth、ErrorDepth、FatalDepth：这些函数用于记录带有调用栈深度信息的日志，可用于定位日志的具体调用位置。

V函数用于设置glogger的日志级别。

总之，glogger.go文件定义的glogger结构体和相关函数提供了grpc-go库自定义的日志记录器，用于记录不同级别的日志，并提供了一些扩展功能，如输出调用栈深度信息等。

