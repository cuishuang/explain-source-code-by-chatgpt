# File: discovery/kubernetes/endpoints.go

discovery/kubernetes/endpoints.go文件是Prometheus项目中的一个文件，主要用于从Kubernetes API获取Endpoints数据，并将其转换成Prometheus的Target格式。

该文件中的epAddCount，epUpdateCount和epDeleteCount变量分别用于计数Endpoints的增加、更新和删除操作。

Endpoints结构体用于存储一个Endpoints资源的所有信息，包括名称、命名空间、子集等。

- NewEndpoints函数用于创建一个新的Endpoints资源。

- enqueueNode函数负责将某个节点标记为需要进行处理，以获取其上的Endpoints资源。

- enqueue函数用于将Endpoints资源添加到任务队列中，以进行后续处理。

- Run函数是Endpoints的主要执行函数，用于从Kubernetes API获取Endpoints的变化并进行处理。

- process函数是Run函数的子函数，用于处理具体的Endpoints操作，如增加、更新和删除。

- convertToEndpoints函数用于将Kubernetes API返回的Endpoints数据转换为Prometheus的Target格式。

- endpointsSource函数用于从Kubernetes API获取Endpoints资源的源。

- endpointsSourceFromNamespaceAndName函数根据给定的命名空间和名称获取Endpoints资源的源。

- buildEndpoints函数用于构建Endpoints资源。

- resolvePodRef函数用于解析Pod的引用，将其转换为Endpoints的Target格式。

- addServiceLabels函数用于向Endpoints添加服务标签。

- addNodeLabels函数用于向Endpoints添加节点标签。

总的来说，discovery/kubernetes/endpoints.go文件的作用是与Kubernetes API交互，从中获取Endpoints资源并进行转换和处理，以符合Prometheus的Target格式。

