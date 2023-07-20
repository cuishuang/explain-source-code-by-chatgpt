# File: miner/worker.go

miner/worker.go文件的作用是实现了矿工（miner）的工作逻辑，即处理和挖掘新的区块的相关操作。

1. errBlockInterruptedByNewHead：当新的区块头到达时，挖掘过程中的区块被中断的错误变量。
2. errBlockInterruptedByRecommit：当需要重新计算工作证明时，挖掘过程中的区块被中断的错误变量。
3. errBlockInterruptedByTimeout：当挖掘过程中超时时，挖掘过程中的区块被中断的错误变量。

结构体：
1. environment：保存矿工工作环境相关的信息，如块链信息、奖励等。
2. task：保存矿工任务相关的信息，如当前正在处理的区块头、挖矿难度等。
3. newWorkReq：用于接收新的挖矿工作请求的结构体。
4. newPayloadResult：用于传递新的工作负载请求结果的结构体。
5. getWorkReq：用于接收挖矿任务请求的结构体。
6. intervalAdjust：用于调整矿工挖矿间隔的结构体。
7. worker：保存矿工的基本数据结构，包括通信通道等。
8. generateParams：运行矿工的生成参数的结构体。

函数：
1. copy：用于复制新的工作结构体。
2. discard：用于丢弃当前的工作结构体。
3. newWorker：用于创建新的矿工实例。
4. setEtherbase：设置矿工挖矿的地址。
5. etherbase：获取矿工挖矿的地址。
6. setGasCeil：设置矿工挖矿的燃气上限。
7. setExtra：设置矿工挖矿的额外数据。
8. setRecommitInterval：设置重新计算工作证明的时间间隔。
9. pending：返回尚未挖掘的交易数。
10. pendingBlock：返回要挖掘的区块。
11. pendingBlockAndReceipts：返回要挖掘的区块以及相关的收据。
12. start：启动矿工。
13. stop：停止矿工。
14. isRunning：检查矿工是否正在运行。
15. close：关闭矿工。
16. recalcRecommit：重新计算工作证明。
17. newWorkLoop：处理新的工作请求的循环。
18. mainLoop：矿工的主循环，负责处理新的工作请求和发送结果。
19. taskLoop：处理挖矿任务的循环。
20. resultLoop：处理发送结果的循环。
21. makeEnv：创建矿工的工作环境。
22. updateSnapshot：更新矿工的工作快照。
23. commitTransaction：提交挖掘的交易。
24. commitTransactions：提交挖掘的所有交易。
25. prepareWork：准备挖掘的工作。
26. fillTransactions：填充要挖掘的交易。
27. generateWork：生成新的挖掘工作。
28. commitWork：提交挖掘的工作。
29. commit：提交新的挖掘工作。
30. getSealingBlock：获取要封闭的块。
31. isTTDReached：检查时间轴触发深度是否达到。
32. copyReceipts：复制交易收据。
33. totalFees：计算交易的总费用。
34. signalToErr：将信号转换为错误。

