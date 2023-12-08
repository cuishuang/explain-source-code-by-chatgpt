# File: text/internal/export/idna/tables13.0.0.go

text/internal/export/idna/tables13.0.0.go文件是Go的text项目中的一个文件，其作用是存储用于国际化域名（IDNA）转换的数据表。

1. mappings: 这是一个映射表，用于将Unicode字符转换成IDNA值。每个Unicode字符对应一个IDNA值。
2. mappingIndex: 这是一个索引表，用于根据输入的字符查找对应的映射值在mappings中的位置。
3. xorData: 这是一个异或表，用于将IDNA值与其他值进行异或操作以生成最终的转换结果。
4. idnaValues: 这是一个IDNA值的列表，存储了所有可能的IDNA值。
5. idnaIndex: 这是一个索引表，用于根据IDNA值查找其在idnaValues中的位置。
6. idnaSparseOffset: 这是一个偏移表，用于查找相应IDNA值在idnaSparseValues中的位置。
7. idnaSparseValues: 这是一个稀疏数组，存储IDNA值与其相应属性的映射关系。

idnaTrie是一个数据结构，用于高效地进行IDNA转换。它有以下几个结构体作为成员变量：
1. root: 表示idnaTrie的根节点。
2. firstTable: 存储了第一级查找表的指针，用于快速路由到正确的子表。
3. tables: 存储了所有的子表，用于进行具体的字符查找。

上述文件中的以下函数用于支持IDNA转换：
1. lookup: 根据输入字符在idnaTrie中查找相应的IDNA值并返回。
2. lookupUnsafe: 与lookup函数类似，但是在查找字符时不进行边界检查，提高了查找速度。
3. lookupString: 根据输入字符串在idnaTrie中查找相应的IDNA值并返回。
4. lookupStringUnsafe: 与lookupString函数类似，但是在查找字符时不进行边界检查，提高了查找速度。
5. newIdnaTrie: 创建一个新的idnaTrie数据结构。
6. lookupValue: 根据IDNA值在idnaSparseValues中查找相应的属性值并返回。

