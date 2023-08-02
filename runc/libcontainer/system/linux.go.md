# File: runc/libcontainer/system/linux.go

在runc项目中，runc/libcontainer/system/linux.go文件负责实现与Linux系统相关的低级操作。主要涉及进程管理、资源控制和系统调用等功能。

ParentDeathSignal这几个结构体用于定义进程的父进程死亡时的信号处理。

- Restore用于在进程恢复后重新设置父进程死亡时的信号处理。
- Set用于设置进程的父进程死亡时的信号处理。
- Execv用于替代当前进程执行新的程序，并且传入替换程序的路径和参数。
- Exec用于在当前进程上下文中执行新的程序，可以用来替代当前正在执行的程序。
- SetParentDeathSignal用于设置进程的父进程死亡时的信号处理。
- GetParentDeathSignal用于获取进程的父进程死亡时的信号处理。
- SetKeepCaps用于设置是否保持进程的特权权限。
- ClearKeepCaps用于清除进程的特权权限。
- Setctty用于设置控制终端。
- SetSubreaper用于设置进程为subreaper，即可以接收孤儿进程的状态改变通知。
- GetSubreaper用于获取进程是否是subreaper。

这些函数提供了对进程的控制和配置功能，例如设置信号处理、执行新程序、设置特权权限等。

