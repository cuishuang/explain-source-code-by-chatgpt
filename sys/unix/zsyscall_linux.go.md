# File: /Users/fliter/go2023/sys/unix/zsyscall_linux.go

在Go语言的sys项目中，/Users/fliter/go2023/sys/unix/zsyscall_linux.go文件的作用是提供对Unix系统调用的封装和定义。该文件包含了大量的函数声明，每个函数对应一个Unix系统调用。

FanotifyInit: 创建一个与文件系统通知 (fanotify) 相关联的文件描述符。
fchmodat: 修改一个指定路径的文件的访问权限。
fchmodat2: 修改一个指定路径的文件的访问权限和文件创建方式。
ioctl: 对设备进行控制操作。
ioctlPtr: 对包含指针的设备进行控制操作。
Linkat: 创建硬链接。
openat: 通过路径名打开文件。
openat2: 打开指定路径的文件或目录。
pipe2: 创建一个管道。
ppoll: 等待描述符上的IO事件。
Readlinkat: 读取一个符号链接的目标路径。
Symlinkat: 创建一个符号链接。
Unlinkat: 删除指定路径的文件或目录。
utimensat: 修改指定路径的文件的读写时间。
Getcwd: 获取当前工作目录。
wait4: 等待子进程终止。
Waitid: 等待在进程组中有任一子进程终止。
KeyctlInt: 使用键ctl接口控制键的操作。
KeyctlBuffer: 使用键ctl接口控制键的操作，以缓冲区方式传递参数。
keyctlJoin: 加入一个session或者加入指定键的ring。
keyctlSearch: 搜索匹配指定标志和类型的键。
keyctlIOV: 使用键ctl接口控制键的操作，以IOV方式传递参数。
keyctlDH: 协商并输出一个Diffie-Hellman公钥和私钥。
keyctlRestrictKeyringByType: 限制一个keyring的访问权限。
keyctlRestrictKeyring: 限制一个keyring的访问权限。
ptrace: 控制一个进程的行为。
ptracePtr: 控制一个进程的行为，以指针方式传递参数。
reboot: 重启系统或者关闭系统。
mount: 挂载文件系统。
mountSetattr: 修改已挂载文件系统的属性。
Acct: 打开或关闭进程会计。
AddKey: 往内核键环中添加一个指定信息的新键。
Adjtimex: 控制内核的时钟调整。
Capget: 获取进程能力。
Capset: 设置进程能力。
Chdir: 改变当前工作目录。
Chroot: 改变根目录。
ClockAdjtime: 调整与指定时钟相关的时间偏差和频率的设置。
ClockGetres: 获取与指定时钟相关的分辨率。
ClockGettime: 获取与指定时钟相关的时间。
ClockNanosleep: 暂停指定时间。
Close: 关闭一个文件描述符。
CloseRange: 关闭一组文件描述符。
CopyFileRange: 在两个文件描述符之间复制数据。
DeleteModule: 卸载内核模块。
Dup: 复制文件描述符。
Dup3: 复制文件描述符和指定标志的文件控制属性。
EpollCreate1: 创建一个epoll文件描述符。
EpollCtl: 控制一个epoll文件描述符。
Eventfd: 创建一个event文件描述符。
Exit: 终止当前进程。
Fchdir: 改变当前工作目录到指定文件描述符所指向的路径。
Fchmod: 修改一个文件的访问权限。
Fchownat: 修改一个指定路径的文件的所有者和组。
Fdatasync: 强制写入文件描述符所指向的文件的数据。
Fgetxattr: 获取文件扩展属性的值。
FinitModule: 初始化内核模块。
Flistxattr: 列出文件的扩展属性。
Flock: 对文件进行加锁。
Fremovexattr: 移除文件的扩展属性。
Fsetxattr: 设置文件的扩展属性。
Fsync: 强制写入文件描述符所指向的文件的数据和元数据。
Fsmount: 操作文件系统信息。
Fsopen: 操作文件系统信息。
Fspick: 操作文件系统信息。
Getdents: 读取目录的内容。
Getpgid: 获取进程所在进程组的ID。
Getpid: 获取当前进程的ID。
Getppid: 获取当前进程的父进程的ID。
Getpriority: 获取进程的优先级。
Getrandom: 随机生成指定大小的字节。
Getrusage: 获取进程或者进程组的资源使用情况。
Getsid: 获取会话ID。
Gettid: 获取线程ID。
Getxattr: 获取文件或目录的扩展属性。
InitModule: 初始化内核模块。
InotifyAddWatch: 添加一个inotify监控。
InotifyInit1: 初始化一个inotify实例。
InotifyRmWatch: 移除一个inotify监控。
Kill: 给指定进程发送信号。
Klogctl: 控制内核日志。
Lgetxattr: 获取链接文件扩展属性的值。
Listxattr: 列出文件或目录的扩展属性。
Llistxattr: 列出链接文件的扩展属性。
Lremovexattr: 移除链接文件的扩展属性。
Lsetxattr: 设置链接文件的扩展属性。
MemfdCreate: 创建一个匿名文件。
Mkdirat: 创建目录。
Mknodat: 创建特殊文件。
MoveMount: 迁移已挂载文件系统。
Nanosleep: 暂停指定时间。
OpenTree: 打开一个文件树。
PerfEventOpen: 管理性能事件。
PivotRoot: 更改进程的根目录。
Prctl: 控制指定进程的行为。
pselect6: 轮询一组文件描述符的IO事件。
read: 从文件描述符读取数据。
Removexattr: 移除文件的扩展属性。
Renameat2: 重命名指定路径的文件或目录。
RequestKey: 请求一个公共或特定键类型的键值。
Setdomainname: 设置域名。
Sethostname: 设置主机名。
Setpgid: 设置指定进程的进程组ID。
Setsid: 设置新的会话ID。
Settimeofday: 设置系统的日期和时间。
Setns: 切换命名空间。
Setpriority: 设置进程的优先级。
Setxattr: 设置文件或目录的扩展属性。
signalfd: 创建一个文件描述符，用于接收进程接收到的信号。
Statx: 获取文件或目录的元数据。
Sync: 将缓冲区的数据写入存储设备。
Syncfs: 将缓冲区的数据写入存储设备。
Sysinfo: 获取系统信息。
TimerfdCreate: 创建一个定时器文件描述符。
TimerfdGettime: 获取定时器的当前超时时间。
TimerfdSettime: 设置定时器的超时时间。
Tgkill: 给指定线程组内的线程发送信号。
Times: 获取进程或线程的运行时间。
Umask: 设置新的文件创建权限掩码。
Uname: 获取系统的名称和版本信息。
Unmount: 卸载文件系统。
Unshare: 在进程间共享资源的基础上，完成系统重新初始化。
write: 向文件描述符写入数据。
exitThread: 终止当前线程。
readv: 从文件描述符中读取数据到多个缓冲区。
writev: 从多个缓冲区将数据写入文件描述符。
preadv: 从文件描述符中读取指定偏移量的数据到多个缓冲区。
pwritev: 从多个缓冲区将数据写入文件描述符的指定偏移量处。
preadv2: 从文件描述符中读取指定偏移量的数据到多个缓冲区，指定读取选项。
pwritev2: 从多个缓冲区将数据写入文件描述符的指定偏移处，指定写入选项。
munmap: 取消映射由mmap创建的内存区域。
mremap: 重新映射由mmap创建的内存区域，改变大小和位置。
Madvise: 告知内核有关给定内存区域的使用模式。
Mprotect: 修改给定内存区域的保护标志。
Mlock: 锁定给定内存区域中的页。
Mlockall: 锁定进程的整个虚拟地址空间中的页。
Msync: 同步文件到磁盘。
Munlock: 解锁给定内存区域中的页。
Munlockall: 解锁进程的整个虚拟地址空间中的页。
faccessat: 检查指定路径的文件的访问权限。
Faccessat2: 检查指定路径的文件的访问权限，指定选项。
nameToHandleAt: 获取指定路径的文件的持久化句柄。
openByHandleAt: 根据持久化句柄打开文件。
ProcessVMReadv: 将远程进程的虚拟内存区域的数据读取到本地缓冲区。
ProcessVMWritev: 将本地缓冲区中的数据写入远程进程的虚拟内存区域。
PidfdOpen: 获取指定进程的文件描述符。
PidfdGetfd: 获取指定的pidfd所引用的文件描述符。
PidfdSendSignal: 给指定进程发送信号，使用pidfd。
shmat: 将共享内存段附加到进程的地址空间。
shmctl: 控制共享内存段。
shmdt: 分离共享内存段。
shmget: 获取共享内存段的标识符。
getitimer: 获取指定的定时器的超时时间。
setitimer: 设置指定的定时器的超时时间。
rtSigprocmask: 改变进程的信号屏蔽字。
getresuid: 获取实际、有效和保存的用户ID。
getresgid: 获取实际、有效和保存的组ID。
schedSetattr: 设置指定进程的调度策略、优先级和相关选项。
schedGetattr: 获取指定进程的调度策略、优先级和相关选项。
Cachestat: 获取系统缓存状态。
