# File: pkg/kubelet/images/image_gc_manager.go

pkg/kubelet/images/image_gc_manager.go文件的作用是实现了镜像的垃圾回收（Garbage Collection）功能，用于清理不再需要的镜像。

- StatsProvider是一个接口，用于获取容器运行时的统计信息。
- ImageGCManager是一个结构体，负责管理垃圾回收的相关逻辑。
- ImageGCPolicy是一个接口，用于指定垃圾回收的策略。
- realImageGCManager是ImageGCManager的实现结构体。
- imageCache是一个缓存镜像信息的结构体。
- imageRecord是一个记录镜像信息的结构体。
- evictionInfo是一个记录镜像删除信息的结构体。
- byLastUsedAndDetected是一个排序器，按照最近使用和检测时间进行排序。

下面是各个函数的作用说明：
- set用于设置镜像的使用状态。
- get用于获取镜像的使用状态。
- NewImageGCManager是一个构造函数，用于创建ImageGCManager实例。
- Start用于启动垃圾回收任务。
- GetImageList用于获取镜像列表。
- detectImages用于检测哪些镜像需要被清理。
- GarbageCollect用于执行垃圾回收操作，清理不再使用的镜像。
- DeleteUnusedImages用于删除不再使用的镜像。
- freeSpace用于释放空间。
- Len用于获取镜像列表的长度。
- Swap用于交换镜像列表中的两个元素。
- Less用于比较两个镜像的使用时间和检测时间。
- isImageUsed用于判断镜像是否在使用中。

