# File: les/downloader/resultstore.go

在go-ethereum项目中，les/downloader/resultstore.go文件的作用是定义并实现了结果存储相关的逻辑。该文件中的结构体和函数用于管理和存储下载任务的结果。

结构体：
1. resultStore: 表示结果存储对象，用于存储下载任务结果的容器。
2. fetchResult: 用于封装单个下载任务的结果及相关信息，包括区块号、数据和错误等。
3. completedItem: 表示已完成的结果项，包括index和标识下载任务是否完成的completed字段。

函数：
1. newResultStore: 创建并返回一个新的结果存储对象resultStore。
2. SetThrottleThreshold: 设置下载速率控制阈值，限制每秒下载的区块数量。
3. AddFetch: 向结果存储对象中添加一个下载任务，并返回任务在结果存储对象中的唯一标识。
4. GetDeliverySlot: 获取可用的传递时隙（用于限制同时传递的区块数量）。
5. getFetchResult: 获取指定下载任务结果的引用（fetchResult）。
6. HasCompletedItems: 检查是否有已完成的下载任务。
7. countCompleted: 统计已完成的下载任务数量。
8. GetCompleted: 获取所有已完成的下载任务结果。
9. Prepare: 通过检查已完成的下载任务，准备进行传递和读取结果。

简要介绍了resultStore文件中的结构体和函数。这些结构体和函数的作用是为了有效管理和存储下载任务的结果，并提供一些基本的操作和查询功能。

