# File: runc/libcontainer/cgroups/manager/new.go

在runc项目中，`runc/libcontainer/cgroups/manager/new.go`文件的作用是为cgroup的管理器创建新的实例。

`New`函数是创建Cgroup管理器的函数，它通过接收一个配置对象和一个父路径参数，返回一个新的Cgroup管理器实例。该实例包含了所需的信息和方法来管理和操作cgroup资源。

`NewWithPaths`函数是一个类似于`New`函数的变体，它接收一个配置对象、父路径参数和内存、CPU路径参数，用于指定保存在配置中的内存和CPU组相应的cgroup路径。这使得可以在更高级别的系统上进行更精细的控制。

`getUnifiedPath`函数根据指定的子系统名称和父路径，返回一个cgroup的绝对路径。这个绝对路径是在使用默认的cgroup组织结构时通过将子系统名称与父路径拼接而成的。

通过这几个函数的组合，可以创建一个实例化的Cgroup管理器，并在文件系统中创建相应的cgroup目录来管理资源。这个cgroup目录提供了一种将进程限制到指定资源的机制。

