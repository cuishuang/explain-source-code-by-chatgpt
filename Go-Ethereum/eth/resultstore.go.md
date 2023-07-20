# File: eth/downloader/resultstore.go

在go-ethereum项目中，eth/downloader/resultstore.go文件的作用是实现下载器结果的存储和管理。

该文件中定义了几个结构体，分别是resultStore、resultRequest和resultSet。

- resultStore结构体用于存储整个下载过程中的结果。它包含了一个结果请求队列和一个结果集合。结果请求队列按照优先级和时间戳排序，结果集合用于存储已经完成的结果。

- resultRequest结构体表示一个结果请求，它包含了请求的哈希值、权重、优先级和时间戳等信息。下载器根据这些信息来优先处理请求。

- resultSet结构体用于存储一个结果集合。

下面是结果存储结构体（resultStore）的几个方法的介绍：

- newResultStore：创建一个新的结果存储结构体。

- SetThrottleThreshold：设置下载速率的阈值。

- AddFetch：添加一个结果请求到结果请求队列中。

- GetDeliverySlot：获取一个结果请求可以交付的位置。

- getFetchResult：根据哈希值获取结果请求的结果。

- HasCompletedItems：检查是否有已经完成的结果请求。

- countCompleted：统计已经完成的结果请求数量。

- GetCompleted：获取已经完成的结果请求集合。

- Prepare：准备下载器结果存储结构体，启动处理结果的goroutine。

这些方法的具体作用如下：

- newResultStore方法用于创建一个新的结果存储结构体。

- SetThrottleThreshold方法用于设置下载速率的阈值，以控制下载器的速率。

- AddFetch方法用于将一个结果请求添加到结果请求队列中，按照优先级和时间戳排序。

- GetDeliverySlot方法用于获取一个结果请求可以交付的位置，根据下载速率和阈值来计算。

- getFetchResult方法用于根据哈希值获取结果请求的结果，如果结果还没有准备好，它将会等待。

- HasCompletedItems方法用于检查是否有已经完成的结果请求。

- countCompleted方法用于统计已经完成的结果请求数量。

- GetCompleted方法用于获取已经完成的结果请求集合。

- Prepare方法用于准备下载器结果存储结构体，启动处理结果的goroutine。它将会处理结果请求队列中的请求，并将已经完成的结果添加到结果集合中。

总结来说，resultStore.go文件中的resultStore结构体和相关方法用于管理和存储下载器的结果请求和结果集合，并提供相应的查询和操作接口。

