# File: runc/libcontainer/cgroups/systemd/common.go

在runc项目中，runc/libcontainer/cgroups/systemd/common.go文件的作用是提供了与systemd相关的函数和变量，用于管理和操作Linux系统中的服务单元。

具体来说，该文件包含了以下几个变量和函数：

1. versionOnce：用于确保version在整个程序运行过程中只被初始化一次的sync.Once对象。

2. version：用于存储已解析的systemd版本号。

3. isRunningSystemdOnce：用于确保isRunningSystemd在整个程序运行过程中只被检查一次的sync.Once对象。

4. isRunningSystemd：用于标识当前系统是否正在运行systemd。

5. GenerateDeviceProps：用于根据指定的配置生成设备属性列表。

接下来，我们分别介绍每个变量和函数的具体作用：

- IsRunningSystemd：用于检查系统中是否正在运行systemd。

- ExpandSlice：用于将单位分段的字符串扩展成完整的切片。

- newProp：用于创建一个新的Property对象，表示systemd服务单元的属性。

- getUnitName：获取给定路径（例如cgroup路径）的unit名称。

- getUnitType：获取指定路径的unit类型。

- isDbusError：检查给定的DBUS错误是否表示unit未找到。

- isUnitExists：检查指定的unit是否存在。

- startUnit：启动指定的unit。

- stopUnit：停止指定的unit。

- resetFailedUnit：重置指定的unit的失败次数。

- getUnitTypeProperty：获取指定unit类型的属性。

- setUnitProperties：设置指定unit的属性。

- getManagerProperty：获取指定manager属性的值。

- systemdVersion：获取当前系统的systemd版本。

- systemdVersionAtoi：将systemd版本号解析为整数。

- addCpuQuota：添加cpu配额限制。

- addCpuset：添加cpuset限制。

- generateDeviceProperties：根据指定的配置生成设备属性列表。

总之，runc/libcontainer/cgroups/systemd/common.go文件在runc项目中是与systemd相关的功能的集合，提供了对systemd服务单元的管理和操作的函数和变量。

