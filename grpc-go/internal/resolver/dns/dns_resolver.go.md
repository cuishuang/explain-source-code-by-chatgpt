# File: grpc-go/internal/resolver/dns/dns_resolver.go

对于grpc-go项目中的dns_resolver.go文件，它的作用是为gRPC提供一种解析DNS名称的功能。详细介绍如下：

1. 变量解释：
- EnableSRVLookups：用于启用或禁用SRV查询的标志。当启用时，该解析器将尝试使用SRV记录进行负载平衡。
- logger：用于记录日志的接口实例。
- newTimer：用于创建计时器的方法。
- newTimerDNSResRate：用于每隔一段时间计算DNS解析速度的方法。
- errMissingAddr：当解析的地址为空时返回的错误。
- errEndsWithColon：当解析的地址以冒号结尾时返回的错误。
- defaultResolver：默认的DNS解析器。
- minDNSResRate：最小的DNS解析速率，用于控制频繁执行解析操作的频率。
- addressDialer：用于建立网络连接的方法。
- newNetResolver：用于创建网络解析器的方法。

2. 结构体解释：
- dnsBuilder：实现了Builder接口，用于构建解析器。
- netResolver：实现了Resolver接口，用于解析网络地址。
- deadResolver：实现了Resolver接口，用于解析无效的网络地址，用于保留旧的解析结果。
- dnsResolver：实现了Resolver接口，用于解析DNS名称。
- rawChoice：表示选择的解析结果。

3. 函数解释：
- init：初始化方法，在包加载时调用。
- NewBuilder：创建并返回dnsBuilder实例。
- Build：创建一个新的解析器。
- Scheme：返回解析器的Scheme名称。
- ResolveNow：立即重新解析当前已知地址。
- Close：关闭解析器。
- watcher：解析器的观察对象。
- lookupSRV：查询SRV记录。
- handleDNSError：处理DNS解析错误。
- lookupTXT：查询TXT记录。
- lookupHost：查询主机记录。
- lookup：查询DNS记录。
- formatIP：格式化IPv4或IPv6地址。
- parseTarget：解析目标地址。
- containsString：检查字符串切片中是否包含指定字符串。
- chosenByPercentage：根据百分比设置选项。
- canaryingSC：支持金丝雀部署的服务配置。

