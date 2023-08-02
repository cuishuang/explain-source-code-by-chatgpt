# File: runc/libcontainer/specconv/spec_linux.go

在runc项目中，runc/libcontainer/specconv/spec_linux.go文件的作用是将OCI规范（Open Container Initiative Specification）的容器配置转换为Linux平台上的配置项。

下面是对每个变量和函数的详细介绍：

变量：
1. initMapsOnce：一个sync.Once类型的变量，用于确保initMaps()函数只被执行一次。
2. namespaceMapping：一个映射表，用于将命名空间类型与对应的标识符进行映射。
3. mountPropagationMapping：一个映射表，用于将挂载传播选项与对应的标识符进行映射。
4. recAttrFlags：一个映射表，用于将文件记录标签属性与对应的标识符进行映射。
5. mountFlags：一个映射表，用于将挂载选项与对应的标识符进行映射。
6. AllowedDevices：一个表示允许的设备列表，用于控制容器中哪些设备可以被访问。

结构体：
1. CreateOpts：用于存储创建容器的选项的结构体，包括容器的命令、环境变量、工作目录等信息。

函数：
1. initMaps：初始化namespaceMapping、mountPropagationMapping、recAttrFlags和mountFlags变量，将Linux平台上的标识符与对应的值进行映射。
2. KnownNamespaces：获取已知的命名空间类型的列表。
3. KnownMountOptions：获取已知的挂载选项的列表。
4. getwd：获取当前的工作目录。
5. CreateLibcontainerConfig：将OCI规范的容器配置转换为libcontainer库使用的配置对象。
6. toConfigIDMap：将OCI规范的配置映射转换为libcontainer库使用的ConfigIDMap对象。
7. createLibcontainerMount：将OCI规范的挂载配置转换为libcontainer库使用的Mount对象。
8. checkPropertyName：检查给定的属性名是否是有效的属性。
9. convertSecToUSec：将秒数转换为微秒数。
10. initSystemdProps：初始化systemd属性。
11. CreateCgroupConfig：将OCI规范的cgroup配置转换为libcontainer库使用的CgroupConfig对象。
12. stringToCgroupDeviceRune：将设备字符表示的字符串转换为设备字符表示的rune类型。
13. stringToDeviceRune：将设备字符表示的字符串转换为设备字符表示的rune类型。
14. createDevices：将OCI规范的设备配置转换为libcontainer库使用的Cgroup设备配置。
15. setupUserNamespace：设置用户命名空间。
16. parseMountOptions：解析挂载选项字符串。
17. SetupSeccomp：设置seccomp配置。
18. createHooks：将OCI规范的钩子配置转换为libcontainer库使用的Hook配置。
19. createCommandHook：创建用于执行命令的钩子。

