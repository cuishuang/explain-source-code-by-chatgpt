# File: p2p/discover/table.go

在go-ethereum的p2p/discover包中，table.go的作用是实现节点发现的网络表。它负责维护已发现的节点列表，并提供节点查找、更新和删除的功能。

以下是p2p/discover/table.go文件中关键结构体的作用：

1. Table：该结构体表示网络表，用于存储已发现的节点。它由一组bucket组成，每个bucket存储具有相同距离范围的节点。

2. Transport：该结构体定义了网络传输的接口。它负责发送和接收节点发现消息。

3. Bucket：该结构体用于存储一组节点。每个bucket根据节点的距离排序，并有一个最大容量限制。

4. nodesByDistance：该结构体用于按节点ID的距离进行排序的辅助结构。

以下是p2p/discover/table.go文件中的关键函数的作用：

1. newTable：创建一个新的Table实例。

2. newMeteredTable：创建一个带有流量计的新Table实例。

3. Nodes：获取表中的所有节点。

4. self：获取本地节点的节点ID。

5. seedRand：随机数生成器，用于从种子节点中选择用于初始发现的节点。

6. getNode：根据节点ID选择并返回一个节点。

7. close：关闭Table，停止刷新和验证。

8. setFallbackNodes：设置备用种子节点。

9. isInitDone：检查Table是否初始化完成。

10. refresh：重新发现网络中的其他节点。

11. loop：定期刷新节点表。

12. doRefresh：进行实际的刷新操作。

13. loadSeedNodes：加载种子节点。

14. doRevalidate：执行节点验证。

15. nodeToRevalidate：选择需要验证的下一个节点。

16. nextRevalidateTime：获取下一次需要进行验证的时间。

17. nextRefreshTime：获取下一次需要进行刷新的时间。

18. copyLiveNodes：复制可用节点。

19. findnodeByID：通过节点ID查找节点。

20. len：获取表中节点的数量。

21. bucketLen：获取指定bucket的节点数量。

22. bucket：获取指定索引的bucket。

23. bucketAtDistance：根据距离获取bucket。

24. addSeenNode：将已看到的节点添加到表中。

25. addVerifiedNode：将已验证的节点添加到表中。

26. delete：从表中删除一个节点。

27. addIP：将IP地址添加到节点。

28. removeIP：从节点中删除IP地址。

29. addReplacement：添加一个替代节点。

30. replace：替换表中的一个节点。

31. bumpInBucket：将一个节点在其bucket中上移。

32. deleteInBucket：从bucket中删除一个节点。

33. contains：检查表中是否包含指定的节点。

34. pushNode：将节点添加到表中。

35. deleteNode：从表中删除一个节点。

36. push：向网络中的节点发送Pong消息。

