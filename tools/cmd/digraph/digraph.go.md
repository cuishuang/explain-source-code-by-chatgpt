# File: tools/cmd/digraph/digraph.go

在Golang的tools/cmd/digraph/digraph.go文件中，定义了一个用于处理图形数据的工具。

doc变量是一个字符串，用于存储程序的文档信息。
stdin变量是一个io.Reader类型，表示标准输入流。
stdout变量是一个io.Writer类型，表示标准输出流。

nodelist结构体是一个自定义类型，用于存储节点列表，并提供了一些操作方法。
nodeset结构体是一个自定义类型，用于存储节点集，并提供了一些操作方法。
graph结构体是一个自定义类型，用于存储图，包含一个节点列表和一个邻接矩阵。

usage函数用于显示程序的使用方法。
main函数是程序的入口，负责解析命令行参数、读取输入数据、创建图对象等。
println函数用于输出字符串到标准输出流。
sort函数用于对节点列表进行排序。
addAll函数用于向节点集合中添加一组节点。
addNode函数用于向节点列表中添加一个节点。
addEdges函数用于向图中添加一组边。
nodelist函数用于创建一个节点列表。
reachableFrom函数用于查找从指定节点可达的节点列表。
transpose函数用于将图的邻接矩阵进行转置。
sccs函数用于查找图的强连通分量。
allpaths函数用于查找从起始节点到目标节点的所有路径。
somepath函数用于查找从起始节点到目标节点的一条路径。
toDot函数用于将图转换为DOT格式字符串。
parse函数用于解析输入数据，构建图对象。
digraph函数是一个包装函数，用于创建并配置一个图对象。
split函数用于将字符串通过空格分割为多个子字符串。
quotedLength函数返回去除引号后的字符串长度。

