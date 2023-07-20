# File: les/downloader/queue.go

在go-ethereum项目中，les/downloader/queue.go文件的作用是管理下载队列，它负责协调和控制从网络中下载区块、头部和收据。

以下是对每个变量的详细解释：

1. blockCacheMaxItems：缓存中允许的最大区块数量。
2. blockCacheInitialItems：缓存中初始的区块数量。
3. blockCacheMemory：内存中缓存区块数据的大小。
4. blockCacheSizeWeight：区块缓存占用内存的比重。
5. errNoFetchesPending：表示没有待处理的下载请求的错误信息。
6. errStaleDelivery：表示传递的下载结果已过期的错误信息。

以下是对每个结构体的详细解释：

1. fetchRequest：表示一个下载请求，包含要下载的区块头、区块体或收据的哈希值以及该请求的回调函数。
2. fetchResult：表示一个下载请求的结果，包含下载到的区块头、区块体或收据的数据以及错误信息。
3. queue：表示一个下载队列，包含保存待处理的下载请求和已经下载和待传递的区块头、区块体和收据。

以下是每个函数的详细解释：

1. newFetchResult：创建一个新的下载结果对象。
2. SetBodyDone：标记下载请求的区块体已完成下载。
3. AllDone：判断是否所有的下载任务都已完成。
4. SetReceiptsDone：标记下载请求的收据已完成下载。
5. Done：处理下载请求的完成状态。
6. newQueue：创建一个新的下载队列。
7. Reset：重置下载队列。
8. Close：关闭下载队列。
9. PendingHeaders：获取待处理的区块头的数量。
10. PendingBlocks：获取待处理的区块体的数量。
11. PendingReceipts：获取待处理的收据的数量。
12. InFlightHeaders：获取正在下载中的区块头的数量。
13. InFlightBlocks：获取正在下载中的区块体的数量。
14. InFlightReceipts：获取正在下载中的收据的数量。
15. Idle：判断下载队列是否处于空闲状态。
16. ScheduleSkeleton：安排下载区块头的任务。
17. RetrieveHeaders：获取需要下载的区块头的哈希值。
18. Schedule：安排下载区块体或收据的任务。
19. Results：获取下载的结果。
20. Stats：获取下载队列的统计信息。
21. stats：更新下载队列的统计信息。
22. ReserveHeaders：预留一定数量的区块头的空间。
23. ReserveBodies：预留一定数量的区块体的空间。
24. ReserveReceipts：预留一定数量的收据的空间。
25. reserveHeaders：获取预留的区块头的空间。
26. CancelHeaders：取消下载区块头的任务。
27. CancelBodies：取消下载区块体的任务。
28. CancelReceipts：取消下载收据的任务。
29. cancel：取消指定哈希值的下载任务。
30. Revoke：撤销已经完成的下载任务。
31. ExpireHeaders：删除已过期的区块头。
32. ExpireBodies：删除已过期的区块体。
33. ExpireReceipts：删除已过期的收据。
34. expire：删除已过期的下载任务。
35. DeliverHeaders：传递已下载的区块头。
36. DeliverBodies：传递已下载的区块体。
37. DeliverReceipts：传递已下载的收据。
38. deliver：传递下载结果的数据。
39. Prepare：准备下载队列，包括建立缓存和初始化变量等操作。

