# File: tsdb/tombstones/tombstones.go

在Prometheus项目中，tsdb/tombstones/tombstones.go文件是用来处理时间序列删除相关的逻辑。当一个时间序列被删除时，需要将其对应的样本数据标记为删除状态，以便在查询时正确处理。

在文件中，castagnoliTable 是一个64位的CRC-32校验表，用于计算时间序列标识符的哈希值，用于索引和查找数据。

Reader 结构体是用来读取tombstones文件的，其中包含了打开的文件句柄和相关的元数据信息。

Stone 结构体表示一个时间序列的删除记录，记录了该序列在每个时间区间内的删除时间戳范围。

MemTombstones 结构体表示内存中的所有时间序列删除记录，通过维护多个Stone对象来记录每个序列的删除范围。

Interval 结构体表示一个时间区间的起始和结束时间戳。

Intervals 结构体是一个有序的Interval切片，用于表示一系列间断的时间区间。

init 函数是用来初始化CRC32校验表。

newCRC32 函数创建一个新的CRC32校验对象。

WriteFile 函数将删除记录写入到文件中。

Encode 函数将Stone结构体编码为二进制数据。

Decode 函数将二进制数据解码为Stone结构体。

ReadTombstones 函数从文件中读取并解码删除记录。

NewMemTombstones 函数创建一个新的内存删除记录对象。

NewTestMemTombstones 函数用于在测试中创建一个内存删除记录对象。

Get 函数根据序列的标识符获取对应的删除记录。

DeleteTombstones 函数删除指定序列的删除记录。

TruncateBefore 函数删除指定时间戳之前的删除记录。

Iter 函数用于迭代指定时间区间内的删除记录。

Total 函数返回删除记录的数量。

AddInterval 函数将一个时间区间添加到Intervals结构体中。

Close 函数关闭文件读取器。

InBounds 函数检查一个时间戳是否在指定时间范围内。

IsSubrange 函数检查一个删除记录是否为另一个记录的子范围。

Add 函数将一个时间序列的删除记录添加到内存删除记录对象中。

以上是tsdb/tombstones/tombstones.go文件中的主要功能和相关结构体及函数的作用。通过这些功能，Prometheus实现了对时间序列的删除和查询规则的正确处理。

