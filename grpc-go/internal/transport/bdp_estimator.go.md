# File: grpc-go/internal/transport/bdp_estimator.go

在grpc-go项目中，`bdp_estimator.go`文件的作用是实现基于平滑加权移动平均算法的带宽延迟积（Bandwidth Delay Product，BDP）估算器，用于估算客户端与服务器之间的带宽和延迟。

`bdpPing`是一个用于存储BDP估算器统计信息的结构体。它有以下几个变量：
- `rtt`：表示BDP估算器的往返时延（Round Trip Time, RTT）。
- `timestamp`：表示时间戳，在计算BDP时使用。
- `ackedPings`：表示接收到的已确认ping的数量。
- `ackedPingSize`：表示已确认ping的大小。
- `outstandingPings`：表示未确认的ping的数量。
- `outstandingPingSize`：表示未确认ping的大小。
- `smoothedBDP`：表示平滑后的BDP值。
- `inWindow`：表示是否处于可用的时间窗口内。

`bdpEstimator`是BDP估算器的主要结构体，它有以下几个变量：
- `mu`：用于保护结构体的并发访问。
- `r`：表示平滑系数，即加权移动平均算法中的权重因子。
- `windowSize`：表示可用的时间窗口大小。
- `bdpPing`：表示BDP估算器的统计信息。
- `bdpCrossed`：表示是否已经跨越BDP边界，即是否足够算出一个有效的BDP值。

`timesnap`函数用于获取当前时间的时间戳，以纳秒为单位。

`add`函数用于更新BDP估算器的统计信息。它根据接收到的ping的确认状态、大小和时间戳进行更新。如果确认成功，它将更新已确认ping的数量和大小；否则，它将更新未确认的ping的数量和大小。

`calculate`函数用于计算BDP估算器的当前平滑BDP值。它利用已确认的ping和时间窗口大小来计算平滑BDP值。如果BDP边界被越过，它将返回true；否则，返回false。

