# File: plugin/pkg/auth/authorizer/node/graph.go

在kubernetes项目中，plugin/pkg/auth/authorizer/node/graph.go文件是用来实现用于节点授权的图形数据结构和相关操作的。它充当了一种权限验证方式，用于确定节点之间的连接和关联关系。以下是关于文件中相关部分的详细介绍：

- vertexTypes是一个枚举类型，定义了节点的不同类型，例如Pod、PV、VolumeAttachment等，用来标识不同类型的节点。

- namedVertex结构体表示一个带有名称的节点，用于表示具体的资源对象。它包含一个ID字段，用来唯一标识该节点，以及一个String字段，用来表示节点名称。

- destinationEdge结构体用于表示节点之间的连接关系，包括从一个节点到另一个节点的权重。

- Graph结构体表示整个图形，包含了与节点和边相关的信息。它包括了多个索引和映射，用于快速查找和操作节点和边。

- namespaceVertexMapping结构体用于映射命名空间到具体的节点，方便查找某个命名空间下的节点。

- nameVertexMapping结构体用于映射名称到具体的节点，方便根据名称查找节点。

- vertexType是一个映射表，用于将节点类型映射到对应的枚举类型。

- newNamedVertex函数用于创建一个新的带有名称的节点。

- ID函数用于获取节点的唯一标识。

- String函数用于获取节点的名称。

- newDestinationEdge函数用于创建一个从源节点到目标节点的连接。

- From函数用于获取边的源节点。

- To函数用于获取边的目标节点。

- Weight函数用于获取边的权重。

- DestinationID函数用于获取边的目标节点的唯一标识。

- NewGraph函数用于创建一个新的图形对象。

- getOrCreateVertex_locked函数用于获取或创建一个节点，并在操作过程中加锁。

- getVertex_rlocked函数用于获取一个节点的只读副本，不会加锁。

- createVertex_locked函数用于创建一个新的节点，并在操作过程中加锁。

- deleteVertex_locked函数用于删除一个节点，并在操作过程中加锁。

- deleteEdges_locked函数用于删除与节点相关的边，并在操作过程中加锁。

- removeEdgeFromDestinationIndex_locked函数用于从目标节点的索引中删除一条边，并在操作过程中加锁。

- addEdgeToDestinationIndex_locked函数用于向目标节点的索引中添加一条边，并在操作过程中加锁。

- removeVertex_locked函数用于从图形中删除一个节点，并在操作过程中加锁。

- recomputeDestinationIndex_locked函数用于重新计算目标节点的索引，并在操作过程中加锁。

- AddPod函数用于向图形中添加一个Pod节点。

- DeletePod函数用于从图形中删除一个Pod节点。

- AddPV函数用于向图形中添加一个PV节点。

- DeletePV函数用于从图形中删除一个PV节点。

- AddVolumeAttachment函数用于向图形中添加一个VolumeAttachment节点。

- DeleteVolumeAttachment函数用于从图形中删除一个VolumeAttachment节点。

这些函数都是用于在图形中进行节点和边的操作，从而实现节点授权的功能。

