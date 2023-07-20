# File: core/rawdb/freezer_table.go

在go-ethereum项目中，core/rawdb/freezer_table.go文件的作用是实现了一个冻结表（freezer table），用于管理冻结数据的索引和文件。详细介绍如下：

errClosed：该变量表示操作在文件关闭状态下发生的错误。
errOutOfBounds：该变量表示操作在索引范围之外时发生的错误。
errNotSupported：该变量表示操作不受支持时发生的错误。

indexEntry：该结构体表示冻结表的索引项，包含了头部位置、长度和尾部位置。
freezerTable：该结构体表示冻结表，包含了文件路径、文件句柄、索引项和锁等信息。

unmarshalBinary：该函数用于将字节数据解码为索引项。
append：该函数用于向冻结表添加一个新的索引项。
bounds：该函数用于获取冻结表的头部和尾部索引项。
newFreezerTable：该函数用于创建一个新的冻结表。
newTable：该函数用于创建一个新的索引表。
repair：该函数用于修复由于意外中断或错误导致的冻结表问题。
preopen：该函数用于打开并预加载冻结表文件。
truncateHead：该函数用于删除冻结表的头部索引项及其对应的数据。
truncateTail：该函数用于删除冻结表的尾部索引项及其对应的数据。
Close：该函数用于关闭冻结表。
openFile：该函数用于打开冻结表文件。
releaseFile：该函数用于释放冻结表文件的句柄。
releaseFilesAfter：该函数用于释放冻结表文件句柄后的所有文件。
releaseFilesBefore：该函数用于释放冻结表文件句柄前的所有文件。
getIndices：该函数用于获取给定范围内的索引项列表。
Retrieve：该函数用于检索给定索引项的数据。
RetrieveItems：该函数用于检索给定索引项列表的数据。
retrieveItems：该函数用于通过索引项列表检索数据。
has：该函数用于检查给定索引项是否存在。
size：该函数用于获取冻结表的大小。
sizeNolock：该函数用于获取冻结表的大小（无锁版本）。
advanceHead：该函数用于前进头部索引指针。
Sync：该函数用于将冻结表的内容同步到持久化存储中。
dumpIndexStdout：该函数用于将冻结表的索引项打印到标准输出。
dumpIndexString：该函数用于将冻结表的索引项以字符串形式返回。
dumpIndex：该函数用于将冻结表的索引项写入到指定的文件中。

