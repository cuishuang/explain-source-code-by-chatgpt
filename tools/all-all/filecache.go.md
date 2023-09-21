# File: tools/gopls/internal/lsp/filecache/filecache.go

文件filecache.go的作用是提供一个文件缓存，用于在gopls工具中管理文件的读取、写入和缓存操作。

下面是对各个变量和函数的详细介绍：

变量：
1. memCache: 内存缓存，用于存储已加载的文件内容。
2. ErrNotFound: 表示文件未找到的错误。
3. active: 一个计数器，用于记录当前有多少个请求正在处理中。
4. iolimit: 一个计数器，用于限制并发I/O操作的数量。
5. budget: 一个计数器，用于控制文件加载的预算，避免过多的文件加载操作。
6. cacheDirOnce: 一个sync.Once对象，用于确保cacheDir和cacheDirErr只被初始化一次。
7. cacheDir: 缓存目录的路径。
8. cacheDirErr: 获取缓存目录路径时发生的错误。

结构体：
1. memKey: 用于为文件缓存生成唯一的键值。包含了文件名和文件目录的哈希值。

函数：
1. Start: 启动文件缓存机制，设置缓存目录和缓存大小。
2. Get: 从缓存中获取给定文件的内容，如果不存在则从磁盘加载。
3. Set: 设置给定文件的内容到缓存中。
4. writeFileNoTrunc: 将数据写入文件，会创建新文件或覆盖现有文件。
5. SetBudget: 设置缓存的预算大小。
6. filename: 根据给定的键值生成文件名，用于存储缓存文件。
7. getCacheDir: 获取缓存目录的路径。
8. hashExecutable: 计算可执行文件的哈希值。
9. gc: 执行缓存垃圾回收，清理过期的缓存文件。
10. init: 初始化文件缓存相关变量。
11. BugReports: 提供一个帮助信息，用于报告文件缓存相关的问题或错误。

