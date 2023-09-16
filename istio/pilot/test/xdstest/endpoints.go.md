# File: istio/pilot/test/xdstest/endpoints.go

文件istio/pilot/test/xdstest/endpoints.go是Istio Pilot的测试文件，用于测试xDS协议中的Endpoint。

LbEpInfo是一个结构体，用于存储Endpoint的相关信息，包括Cluster，在线状态(isOnline)，地址列表(Addrs)，以及相应的锁(Mutex)。LocLbEpInfo是LbEpInfo的本地版本，用于保存本地Endpoints信息。

以下是每个函数的详细作用：
1. GetAddrs(cluster string) ([]string, bool): 该函数用于获取特定集群(cluster)的地址列表(Addrs)，如果Cluster不存在，则会返回一个空的地址列表和false。

2. CompareEndpointsOrFail(cluster string, newEps []*xds.Endpoint): 该函数用于比较给定集群(cluster)的新Endpoints(newEps)与之前的Endpoints是否相同。如果不相同，会触发测试失败。

3. CompareEndpoints(cluster string, newEps []*xds.Endpoint) error: 该函数与CompareEndpointsOrFail类似，但是不会触发测试失败，而是返回一个错误。

4. getLbEndpointAddrs(cluster string) ([]string, bool): 该函数用于获取特定集群(cluster)的负载均衡Endpoint的地址列表(Addrs)，如果Cluster不存在或者没有负载均衡Endpoint，则会返回一个空的地址列表和false。


