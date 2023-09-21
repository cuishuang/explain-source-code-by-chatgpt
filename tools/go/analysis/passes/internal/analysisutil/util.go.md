# File: tools/go/callgraph/util.go

在Golang的Tools项目中，tools/go/callgraph/util.go文件的作用是提供对调用图(call graph)的一些常见操作和功能。该文件中包含了一些函数，下面是对这些函数的详细介绍：

1. CalleesOf(node *graph.Node) (callees []*graph.Node):
   这个函数用于获取给定节点(node)的所有被调用者(callees)。它基于调用图中的边缘关系，遍历图结构并返回与给定节点相关的所有被调用者节点。

2. GraphVisitEdges(graph *graph.Graph, visit graph.Visitor):
   这个函数用于遍历给定的调用图，并在遍历每条边时调用visit函数。传入的visit函数用于处理每条边的操作。通过这个函数，可以从调用图中获取节点的邻居节点。

3. PathSearch(src, tgt *graph.Node, filter func(*graph.Node) bool) []*graph.Node:
   这个函数用于在调用图中查找从源节点(src)到目标节点(tgt)的路径。可以通过filter函数来过滤不符合要求的节点。它使用深度优先搜索算法来查找路径。

4. DeleteSyntheticNodes(graph *graph.Graph):
   这个函数用于删除调用图中的合成节点(synthetic nodes)。合成节点是在分析调用图时自动生成的临时节点，不代表实际的函数。

5. isInit(node *graph.Node) bool:
   这个函数用于检查给定节点是否表示初始化函数。如果是则返回true，否则返回false。它根据节点的名称来判断是否是初始化函数。

6. DeleteNode(graph *graph.Graph, node *graph.Node):
   这个函数用于删除调用图中的给定节点及其相应的边缘。它会从调用图中删除给定的节点，并删除与该节点相关的所有边缘。

7. deleteIns(node *graph.Node):
   这个函数用于删除调用图中给定节点的所有入边。它会删除与该节点的所有入边，并从相应的出边的源节点的邻居列表中删除该节点。

8. deleteOuts(node *graph.Node):
   这个函数用于删除调用图中给定节点的所有出边。它会删除与该节点的所有出边，并从相应的入边的目标节点的邻居列表中删除该节点。

9. removeOutEdge(source, target *graph.Node):
   这个函数用于删除调用图中给定边缘的出边。它会从源节点的出边列表中删除目标节点，并从目标节点的入边列表中删除源节点。

10. removeInEdge(source, target *graph.Node):
    这个函数用于删除调用图中给定边缘的入边。它会从源节点的入边列表中删除目标节点，并从目标节点的出边列表中删除源节点。

这些函数提供了对调用图的一些常见操作和功能，可以用于分析和处理Golang代码的调用关系。

