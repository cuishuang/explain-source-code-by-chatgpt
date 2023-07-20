# File: p2p/netutil/net.go

在go-ethereum项目中，p2p/netutil/net.go文件是一个网络工具包，提供了一些用于处理网络相关操作的函数和结构体。

lan4,errInvalid,errUnspecified,errSpecial,errLoopback,errLAN这几个变量用于定义不同类型的网络，例如`lan4`表示IPv4局域网，`errInvalid`表示无效的网络类型，`errUnspecified`表示未指定网络类型，等等。

Netlist结构体表示一个网络列表，用于存储网络类型。DistinctNetSet是一个存储唯一网络类型的集合。

init函数主要用于初始化网络类型列表，将预定义的网络类型加入到Netlist中。

ParseNetlist函数用于解析字符串形式的网络列表，并返回Netlist。

MarshalTOML和UnmarshalTOML函数用于将网络列表转换为TOML格式或从TOML格式转换为Netlist。

Add函数用于将网络类型加入到Netlist中。

Contains函数用于判断一个网络列表是否包含指定的网络类型。

IsLAN函数用于判断一个网络类型是否为局域网。

IsSpecialNetwork函数用于判断一个网络类型是否为特殊网络。

CheckRelayIP函数用于检查中继IP地址是否有效。

SameNet函数和sameNet方法用于判断两个网络类型是否相同。

Remove函数用于从网络列表中移除指定的网络类型。

Len函数用于获取网络类型列表的长度。

key方法用于生成一个网络类型的唯一标识符。

String方法用于将网络类型转换为字符串表示。

