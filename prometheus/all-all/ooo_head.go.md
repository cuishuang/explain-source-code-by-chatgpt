# File: tsdb/ooo_head.go

在Prometheus项目中，tsdb/ooo_head.go文件的作用是定义了OOOHead类型，该类型表示一个存储参与OOM实现的块的元信息。

首先，_这几个变量是占位符，代表暂不需要使用的变量，通常用于忽略某个值。

接下来，OOOChunk结构体表示一个块的元信息，在该结构体中，有以下字段：
- Ref records：该块的引用计数。
- MinTime records：该块所包含样本数据的最小时间戳。
- MaxTime records：该块所包含样本数据的最大时间戳。
- FirstVal records：该块中的第一个样本点值。
- LastVal records：该块中的最后一个样本点值。

OOORangeHead结构体表示一个范围的元信息，在该结构体中，有以下字段：
- FileName records：该范围所对应的文件名。
- MinTime records：该范围所包含样本数据的最小时间戳。
- MaxTime records：该范围所包含样本数据的最大时间戳。
- IndexOffset records：该范围所在的文件中索引的偏移。

NewOOOChunk函数用于创建一个所包含样本数据的新的块。

Insert函数用于向块中插入样本点。

NumSamples函数用于返回块中的样本点数量。

ToXOR函数用于计算两个块的异或。

ToXORBetweenTimestamps函数用于计算两个时间戳之间的块的异或。

NewOOORangeHead函数用于创建一个表示范围的新的元信息。

Index函数用于返回范围元信息中的索引值。

Chunks函数用于返回范围元信息中的块列表。

Tombstones函数用于返回范围元信息中的结果列表。

Meta函数用于返回范围元信息中的元信息数据。

Size函数用于返回范围元信息中的大小。

String函数用于返回范围元信息的字符串表示。

MinTime函数用于返回范围元信息中的最小时间戳。

MaxTime函数用于返回范围元信息中的最大时间戳。

