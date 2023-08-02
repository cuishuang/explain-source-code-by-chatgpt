# File: runc/libcontainer/devices/device_unix.go

在runc项目中，`runc/libcontainer/devices/device_unix.go`文件的作用是实现了与设备相关的函数和结构。下面是对其中相关变量和函数的详细介绍：

1. `ErrNotADevice`变量：表示在操作设备时出现的错误类型，表示该文件不是设备。

2. `unixLstat`变量：用于获取文件的信息，类似于`stat`函数，但可以处理符号链接，返回文件的元数据。

3. `osReadDir`变量：用于读取目录中的文件和文件夹。

4. `mkDev(major, minor uint32) uint64`函数：根据给定的主设备号和次设备号创建设备号。它将主设备号左移20位并与次设备号进行或运算，返回一个64位的设备号。

5. `DeviceFromPath(path string) (Device, error)`函数：根据给定的路径返回与之对应的设备实例和错误类型。它通过使用`unixLstat`函数获取路径的文件信息，并使用`mkDev`函数创建设备号。

6. `HostDevices() ([]Device, error)`函数：返回运行runc的主机上的设备列表。它通过读取`/dev`目录下的设备文件来获取设备列表，并使用`DeviceFromPath`函数创建设备实例。

7. `GetDevices()`函数：返回支持的设备列表，即容器可以使用的设备列表。它获取主机设备列表并从中过滤出容器可使用的设备。

