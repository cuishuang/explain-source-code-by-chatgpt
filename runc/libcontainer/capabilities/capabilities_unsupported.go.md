# File: runc/libcontainer/capabilities/capabilities_unsupported.go

在runc项目中，文件`capabilities_unsupported.go`的作用是定义了一些 Linux 容器中不支持的能力（capabilities）列表。

在 Linux 中，能力是指操作系统对于进程或应用程序进行细粒度控制的一种机制。每个进程都被授予一组特定的能力，这些能力控制了进程能够执行的操作。通过调整进程的能力，可以限制其对系统资源的访问权限，提高系统的安全性。

在容器技术中，runc 作为一个基础的容器运行时工具，负责创建和管理 Linux 容器。为了提高容器的安全性，并限制容器中进程的特权级别，runc 对容器中的进程的能力进行了精确的控制。这就要求 runc 必须知道哪些能力是容器支持的，以对容器中的进程进行适当的能力限制。

而`capabilities_unsupported.go`文件则列出了一些 Linux 容器中不支持的能力，这些能力不能被容器中的进程使用。该文件的内容是一个包级别的变量`unsupportedCapabilities`，这个变量是一个能力字符串到布尔值的映射，表示哪些能力是不被支持的。每一个不被支持的能力都以布尔值`true`的方式指示。

通过指定不支持的能力列表，runc 可以在创建容器时对这些能力进行限制。当容器中的进程需要获得某个能力时，runc 将根据这个能力是否在不支持的列表中进行检查，并拒绝或允许相应的操作。

总之，`capabilities_unsupported.go`文件在 runc 项目中定义了容器中不支持的能力列表，用来控制容器中进程的权限和安全性。

