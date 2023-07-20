# File: accounts/keystore/file_cache.go

在go-ethereum项目中，accounts/keystore/file_cache.go文件是keystore文件缓存功能的实现。该文件定义了fileCache结构体和相关方法，用于管理以文件形式存储的密钥库文件。

fileCache结构体是keystore文件缓存的核心数据结构，它有以下几个作用：
- 缓存keystore文件，减少文件操作的次数，提高性能。
- 记录最近访问的keystore文件列表，并根据使用频率进行排列，以实现最近最少使用（LRU）缓存淘汰策略。
- 同时维护一个keyFileMap字典，存储keystore文件的路径和对应的file结构体。

scan函数是fileCache的方法函数，在指定的目录中查找所有的keystore文件，并将其添加到缓存中。它通过遍历目录结构，判断文件是否符合keystore文件的命名规则（以"UTC--"开头），并读取文件头部信息，确认文件格式正确后，生成一个file结构体，并将其添加到缓存中。

nonKeyFile函数是fileCache的方法函数，判断给定的文件是否是keystore文件。它通过检查文件的扩展名是否符合规范（.json）以及文件名是否以"UTC--"开头来确定。该函数用于过滤目录下除了keystore文件之外的其他文件。

