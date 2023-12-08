# File: /Users/fliter/go2023/sys/windows/dll_windows.go

在Go语言的sys项目中，/Users/fliter/go2023/sys/windows/dll_windows.go文件的作用是用于对Windows操作系统的动态链接库（DLL）进行加载和函数调用的支持。
下面是对相关变量和结构体的详细说明：

canDoSearchSystem32Once
这是一个sync.Once类型的变量，用于确保canDoSearchSystem32变量只被初始化一次。它的作用在于判断是否允许搜索系统目录进行DLL加载。

DLLError
这是一个常见的错误类型，用于表示DLL相关的错误。

DLL
这是一个结构体类型，用于表示被加载的DLL的句柄。它包含一个uintptr类型的字段，表示DLL的句柄值。

Proc
这是一个结构体类型，用于表示DLL中的函数。它包含一个uintptr类型的字段，表示函数的地址。

LazyDLL
这是一个结构体类型，用于延迟加载DLL。它包含一个字符串类型的字段，表示DLL的名称。

LazyProc
这是一个结构体类型，用于延迟加载DLL中的函数。它包含一个字符串类型的字段，表示函数的名称。

errString
这是一个结构体类型，用于表示错误信息。它实现了error接口，用于返回解析错误信息的字符串。

syscall_loadlibrary
这个函数用于加载指定的DLL，并返回DLL的句柄。

syscall_getprocaddress
这个函数用于从指定的DLL中获取指定名称的函数地址。

Error
这是一个类型，用于表示DLL相关的错误。

Unwrap
这个方法用于返回错误的底层原因。

LoadDLL
这个函数用于加载指定名称的DLL文件，并返回DLL结构体。

MustLoadDLL
这个函数用于加载指定名称的DLL文件，如果加载失败则 panic。

FindProc
这个函数用于查找DLL中指定名称的函数，并返回Proc结构体。

MustFindProc
这个函数用于查找DLL中指定名称的函数，如果查找失败则 panic。

FindProcByOrdinal
这个函数用于通过序数值（Ordinal）查找DLL中的函数，并返回Proc结构体。

MustFindProcByOrdinal
这个函数用于通过序数值（Ordinal）查找DLL中的函数，如果查找失败则 panic。

Release
这个方法用于释放DLL结构体。

Addr
这个方法用于从Proc中获取函数的地址。

Call
这个方法用于调用Proc中的函数。

Load
这个方法用于加载指定名称的DLL文件，并返回DLL结构体。

mustLoad
这个方法用于加载指定名称的DLL文件，如果加载失败则 panic。

Handle
这个方法用于获取DLL的句柄值。

NewProc
这个方法用于创建Proc结构体。

NewLazyDLL
这个函数用于创建LazyDLL结构体。

NewLazySystemDLL
这个函数用于创建LazyDLL结构体，指定了系统目录进行搜索。

Find
这个方法用于在DLL中查找指定名称的函数，并返回Proc结构体。

mustFind
这个方法用于在DLL中查找指定名称的函数，如果查找失败则 panic。

initCanDoSearchSystem32
这个函数用于初始化canDoSearchSystem32变量，判断是否允许搜索系统目录进行DLL加载。

canDoSearchSystem32
这个变量用于判断是否允许搜索系统目录进行DLL加载。

isBaseName
这个函数用于判断给定的文件名是否是基本名称。

loadLibraryEx
这个函数用于加载指定名称的DLL文件，并返回DLL的句柄。

