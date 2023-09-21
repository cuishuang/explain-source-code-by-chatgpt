# File: grpc-go/balancer/weightedtarget/weightedtarget_config.go

在grpc-go中，weightedtarget_config.go文件定义了WeightedTargetBalancerConfig和weightedTargetBuilder结构体，用于负载均衡策略的配置和构建。

WeightedTargetBalancerConfig是负载均衡策略的配置结构体，它包含一个WeightedTarget列表，用于表示每个目标服务器的权重和地址。

type WeightedTargetBalancerConfig struct {
    Targets []*Target
}

type Target struct {
    Address string
    Weight  int
    Metadata map[string]interface{}
}

其中，WeightedTargetBalancerConfig.Targets列表表示了一组目标服务器，每个目标服务器使用Target结构体表示，其中Address表示服务器地址，Weight表示服务器的权重，Metadata是可选的元数据。

weightedTargetBuilder是负载均衡器的构建器结构体，实现了balancer.Builder接口，用于根据配置构建负载均衡器。

type weightedTargetBuilder struct {}

func (w *weightedTargetBuilder) Build(readySCs map[resolver.Address]grpc.ClientConn) balancer.Balancer {
    return &weightedTargetBalancer{}
}

func (w *weightedTargetBuilder) Scheme() string {
    return weightedtargetName
}

parseConfig函数用于解析负载均衡策略的配置，并将其转换为WeightedTargetBalancerConfig结构体。具体来说，parseConfig将传入的配置转换为map[string]interface{}，获取其中的"targets"字段，根据字段的值解析出WeightedTargetBalancerConfig.Targets列表。

func parseConfig(c *Config) (*WeightedTargetBalancerConfig, error) {
    var wtc WeightedTargetBalancerConfig
    if targets, ok := c.RawConfig["targets"]; ok {
        targetsList, ok := targets.([]interface{})
        if !ok {
            return nil, fmt.Errorf("bad formatted targets (wanted []interface{} got %T)", targets)
        }
        for _, submap := range targetsList {
            m, ok := submap.(map[string]interface{})
            if !ok {
                return nil, fmt.Errorf("target is not a map[string]interface{}: %v", submap)
            }
            if addr, ok := m["addr"].(string); ok {
                wtc.Targets = append(wtc.Targets, &Target{
                    Address: addr,
                    Weight:  int(m["weight"].(float64)), // assume it exists
                    Metadata: m["metadata"].(map[string]interface{}), // assume it exists
                })
            } else {
                err := fmt.Errorf("target doesn't contain addr field: %v", submap)
                return nil, err
            }
        }
    } else {
        return nil, fmt.Errorf("no targets found in weighted_target balancing config")
    }
    return &wtc, nil
}

函数首先获取配置中的"targets"字段，然后依次遍历targetsList，在每个target map中获取addr、weight和metadata字段，并根据字段的值构建Target对象，最后将Target对象添加到WeightedTargetBalancerConfig.Targets列表中。

这些结构体和函数的作用是为了实现自定义的权重目标负载均衡策略，解析配置并构建负载均衡器。

