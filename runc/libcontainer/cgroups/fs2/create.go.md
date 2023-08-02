# File: runc/libcontainer/cgroups/fs2/create.go

在runc项目中，runc/libcontainer/cgroups/fs2/create.go文件的作用是用于创建cgroup来管理容器的资源限制和控制。下面会逐一介绍所提到的这些函数的作用。

1. supportedControllers函数用于获取系统上支持的所有cgroup控制器。在linux系统上，cgroup是以层次结构进行组织，支持不同种类的控制器用于限制不同的资源。

2. needAnyControllers函数用于检查是否至少一个被请求的控制器被当前系统所支持。当需要创建cgroup并限制资源时，需要确保系统上至少支持一个请求的控制器。

3. containsDomainController函数用于检查控制器列表是否包含容器的域控制器。在创建cgroup时，需要确保域控制器（比如cpu、memory等）存在于控制器列表中。

4. CreateCgroupPath函数用于创建cgroup的hierarchy path。该路径由容器的父目录路径加上要创建的cgroup名字组成。在创建路径的过程中，按需创建缺失的目录并设置正确的权限。

这些函数的作用在于在创建cgroup之前，对控制器的支持情况进行检查，保证所需要的控制器列表和路径的正确性。

