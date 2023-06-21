# File: exec_linux.go

exec_linux.go 文件是 Go 语言标准库中 syscall 包中的一个实现文件，主要用于在 Linux 操作系统上执行系统调用中的 exec 系列函数。

exec 系列函数是 Linux 操作系统下的系统调用函数，用于在当前进程中加载并运行一个新的程序。在 Go 语言中，通过 syscall 包中的 exec 函数可以直接调用 exec 系列函数完成这个操作。

exec_linux.go 文件中定义了 exec 系列函数在 Linux 操作系统下的具体实现，主要涉及到了以下几个方面：

1. 定义了一个名为 "RawSyscall" 的函数，该函数用来直接调用 Linux 操作系统内核提供的系统调用，包括 execve、execveat 等。

2. 定义了一个名为 "Exec" 的函数，该函数是 Go 语言中的一个包装函数，用来调用 RawSyscall 函数实现调用 execve 或 execveat 系列函数。该函数的具体实现方式则是通过构造系统调用参数和系统调用号，然后调用 RawSyscall 函数完成操作。

3. 定义了其他一些相关的函数和变量，包括 "ExecErr"、"Executable"、"ErrExecutableNotFound" 等。

总之，exec_linux.go 文件是 syscall 包的一个实现文件，主要用于在 Linux 操作系统下完成 exec 系列函数的具体实现。




---

### Var:

### none

在go/src/syscall中exec_linux.go这个文件中，定义了一个名为none的变量，它的作用是在进程执行execve时，如果指定了文件名以及参数，则使用none变量覆盖掉最后一个参数，以保证执行的新进程不会继承已有的文件描述符。

在Linux系统中，execve系统调用是用来执行新进程的，它的语法格式如下：

```
int execve(const char *filename, char * const argv[], char *const envp[]);
```

其中，filename表示要执行的进程的文件名，argv表示传递给新进程的参数列表，envp表示传递给新进程的环境变量。在调用execve之前，需要先执行fork系统调用，创建一个新的进程，然后使用execve来替换新进程。

但是在执行execve时，有时候需要覆盖掉最后一个参数，以确保执行的新进程不会继承已有的文件描述符。这时就可以使用none变量，将最后一个参数设置为nil，表示不继承任何文件描述符。例如：

```
syscall.Exec("/usr/bin/env", []string{"env", "-i"}, []string{"PATH=/bin:/usr/bin"}, []syscall.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}}) // 这里的最后一个参数就不能是nil或空数组，需要使用none变量来覆盖掉
```

none变量的定义如下：

```
var none = []string{""}
```

它是一个空字符串的数组，表示不继承任何文件描述符。使用none变量来覆盖掉最后一个参数，可以确保执行的新进程不会继承已有的文件描述符，从而保证进程间的独立性和互不干扰。



### slash

在 `exec_linux.go` 文件中，`slash` 这个变量是一个 `byte` 类型，它的值是反斜杠字符（`\`）的 ASCII 码值。它的作用是用于在非 Windows 系统中构建可执行文件路径。

在类 Unix 系统中，文件路径使用正斜杠（`/`）作为分隔符，而 Windows 系统中使用反斜杠（`\`）作为分隔符。当程序需要在不同操作系统上运行时，需要根据不同的操作系统来构建文件路径，因此使用 `slash` 变量可以快捷地区分不同操作系统使用的路径分隔符，简化代码实现。

`slash` 变量的赋值语句如下：

```go
var slash = byte('/')
if runtime.GOOS == "windows" {
   slash = '\\'
}
```

其中，`runtime.GOOS` 是一个常量字符串，表示当前程序所运行的操作系统。在 Windows 系统中，将 `slash` 变量赋值为反斜杠的 ASCII 码值，否则将其赋值为正斜杠的 ASCII 码值。

当需要在程序中构建文件路径时，可以使用 `slash` 变量来代替具体的路径分隔符，以实现跨平台兼容性。例如：

```go
path := "/path/to/file" + string(slash) + "name.txt"
```

在 Windows 系统中，`slash` 变量的值为反斜杠，因此上述代码将会构建出如下的文件路径：

```
\path\to\file\name.txt
```

而在类 Unix 系统中，`slash` 变量的值为正斜杠，因此上述代码将会构建出如下的文件路径：

```
/path/to/file/name.txt
```






---

### Structs:

### SysProcIDMap

SysProcIDMap是一个结构体，用于在Linux上执行进程时，通过设置进程命名空间映射关系，将进程在其中运行的进程ID（PID）与主机上的PID分离。这种分离意味着一个进程在不同的命名空间中运行时，其PID会在每个命名空间中有不同的值。这是Linux上命名空间的核心概念之一。

该结构体包含以下字段：

- ContainerID：用于标识与进程命名空间相关联的容器ID。
- HostID：用于标识在进程外部运行的PID。
- Size：用于定义映射关系的大小。

通过使用这些字段，可以通过设置SysProcIDMap结构体来设置命名空间映射关系，使进程在其中运行时其PID与主机上的PID分离，从而增加了进程的安全性以及进程管理的灵活性。

在Linux容器技术中，使用SysProcIDMap结构体可以有效地实现多个容器之间的PID隔离，使得它们可以在同一主机上运行，而不会产生PID冲突。这为容器化应用程序的开发和管理提供了很大的便利。



### SysProcAttr

SysProcAttr是一个结构体，用于表示系统进程的属性，它包含以下字段：

- Chroot: 字符串类型，表示要修改为的默认根目录。如果为“”，则表示不进行修改。
- Credential: *syscall.Credential类型，用于设定进程所属的用户和组权限。如果为nil，则表示从父进程继承。
- Ptrace: bool类型，表示是否允许调试这个进程。如果为true，则表示允许。
- Setsid: bool类型，表示是否将进程设置为新的会话组和进程组的领头进程。如果为true，则表示设置。
- Setpgid: bool类型，表示是否将进程设置为新的进程组。如果为true，则表示设置。
- Noctty: bool类型，表示进程是否从应用的控制终端独立。如果为true，则表示独立。

这些属性可以用于执行新的系统进程时进行属性的设置。比如：

```go
attr := &syscall.SysProcAttr{
    Chroot:     "/var/run/myroot",
    Credential: &syscall.Credential{Uid: uint32(1), Gid: uint32(1)},
}

// 设置在 chroot 之后并以指定的用户及组权限启动一个进程
cmd := exec.Command("/bin/sh")
cmd.SysProcAttr = attr
```

其中attr指定了在调用/bin/sh时，要将其Chroot到/var/run/myroot目录下，并指定用户及组权限为1。在调用exec.Command()时，将attr赋值给cmd.SysProcAttr，表示启动的进程需要这些属性。



### capHeader

capHeader结构体定义了Linux内核中的能力集头部，包括了能力集的版本号、持有该能力集的进程ID、以及能力集的数量等信息。一个进程在运行时可以拥有多个能力集，每个能力集中定义了一系列能力，这些能力定义了进程能够执行的特权操作。

在执行Linux系统调用中的execve()函数时，会使用capHeader结构体来传递进程要继承的能力集。这些能力集可以用来控制进程是否具备访问某些特权资源的权限。在安全敏感的应用程序中，使用这些能力集可以降低系统的安全风险，并防止不信任的代码执行潜在的攻击操作。在这种情况下，capHeader结构体充当了控制程序权限的重要载体。



### capData

在Go语言的syscall包中，exec_linux.go文件中定义了capData结构体，主要作用是封装Linux内核的capabilities数据结构，用于描述进程的能力。

capabilities是Linux内核中的一个机制，用于控制进程对系统资源的访问权限。每个进程有一组capabilities，描述了进程当前所具备的访问权限。Linux内核中有很多不同的capabilities，例如CAP_NET_ADMIN可以控制网络配置、CAP_SYS_ADMIN可以控制系统管理等等。

capData结构体是syscall包中用于封装capabilities的数据结构，包含3个字段：

- effective：表示进程的有效capabilities，即当前具备的访问权限；
- permitted：表示进程的预设capabilities，即进程启动时被赋予的初始访问权限；
- inheritable：表示进程的可继承capabilities，即进程fork时可被子进程继承的访问权限。

在创建进程时，可以使用syscall.Exec函数指定启动的进程的capabilities，以修改进程的访问权限。capData结构体就是用来描述这些capabilities的数据结构。



### caps

在 Linux 中，capabilties（简称 caps）表示了一个进程的权限集合，包括一些可以打开一些受限制文件、设置一些特殊权限、连接到一些网络端口等操作的权限。每个进程都有自己的 caps，但并不是所有进程都能自由地修改它们。

caps 是一个结构体，它定义了进程的十六个 capabilities。在 exec_linux.go 文件中，这个结构体是用来设置新建进程的 capability 集合的。具体来说，它包含下列十六个成员：

```golang
type caps struct {
    CapList [2]uint32
    Len     uint32
}
```

- CapList：一个长度为 2 的 uint32 数组，用来表示进程 capability 的列表。每个元素表示 32 个 capabilities，即一共包含 64 个 capabilities。
- Len：进程的 capability 数量。

当创建进程时，旧的 caps 将被新的 caps 覆盖，这意味着新进程将继承新的 caps。如果 caps 为空，那么进程将被赋予最小的权限集合（即仅拥有基本 capabilities）。如果 caps 中包含所有 capabilities，那么进程将拥有全部权限。



### cloneArgs

在 go/src/syscall/exec_linux.go 文件中，cloneArgs 结构体是用于传递 Linux 系统调用 Clone() 的参数的。

Clone() 是 Linux 中的一个系统调用，它创建一个新的进程或线程，并在子进程中执行指定的函数。CloneArgs 结构体包含了创建进程所需的所有信息，例如子进程的栈大小、标志、命名空间等。这些参数都在创建新进程时被使用。

exec_linux.go 文件中的 cloneArgs 结构体定义了如下字段：

type cloneArgs struct {
	flags uintptr
	stack uintptr
	pid   *uintptr
	tls   uintptr
	ctid  *int32
}

- flags：表示创建新进程时的标志，类型为 uintptr。
- stack：表示子进程使用的栈大小，类型为 uintptr。
- pid：表示新进程的进程 ID，类型为 *uintptr。
- tls：表示用于存储线程本地存储的内存块地址，类型为 uintptr。
- ctid：表示用于存储子进程 ID 的缓冲区地址，类型为 *int32。

除了 cloneArgs 结构体外， exec_linux.go 文件还包含了其他一些与进程创建和运行相关的函数和结构体，例如 syscall.SysProcAttr 结构体、syscall.ForkExec() 函数等。这些函数和结构体提供了更高级别的接口，使得用户能够更方便地创建和管理进程。



## Functions:

### runtime_BeforeFork

在exec_linux.go中，runtime_BeforeFork函数是一个前置函数，其作用是在调用fork函数之前，预处理一些信息，防止在子进程中无效复制父进程中的文件句柄，而导致资源浪费或者不稳定情况的发生。

具体来说，runtime_BeforeFork做了以下几件事情：

1. 关闭当前进程中的所有pthread（POSIX 线程），避免子进程中因为不必要的线程创建导致资源浪费和性能问题。

2. 使用sysconf获取pipe缓存的大小，并将该值与某些常数比较得出一个缓存大小阈值，以便在exec调用后，当文件描述符被打开/关闭时，可以避免大量kernel memory被不必要的缓存占用。

3. 确定进程的cloned flag。当fork系统调用被调用时，一个新的进程将被创建。如果请求调用方的进程更改某些属性来影响子进程的初始化过程，则这些属性必须通过fork的cloned flag进行控制。在这里，进程更改了cloned flag，这样会影响子进程的复制行为。

总的来说，runtime_BeforeFork的作用就是为了保证在子进程中有一个干净的环境，以避免资源浪费和其他可能导致程序运行不稳定的问题。



### runtime_AfterFork

在go/src/syscall/exec_linux.go文件中，runtime_AfterFork函数是用来确保进程在 fork 之后可以正确的继续执行。该函数的作用是重置一些基于glibc的内部状态，因为这些状态可能与新的进程环境不兼容。

具体来说，该函数会重置以下状态：

1. 清除已注册的fork处理程序，并重新注册它们，以保证这些程序在新进程中执行。

2. 重置glibc的内存状态，包括malloc状态以及malloc_hook和free_hook的值。

3. 重置glibc的文件状态，包括STDIN、STDOUT和STDERR的值以及其他环境变量。

综上所述，runtime_AfterFork函数主要用于在进程 fork 之后重置一些glibc的内部状态，以确保进程可以继续正常工作。



### runtime_AfterForkInChild

func runtime_AfterForkInChild() {
	// Disable profiling, as we might confuse the profiler if its
	// signal handler runs concurrently with the fork.
	disableProfiler()
	// Send a message on the pipe to signal that we're done,
	// and the parent can proceed.
	sendAfterForkMessage()
}

这个函数在一个进程调用fork（）创建子进程之后，该子进程会继承来自父进程的大部分资源。此函数的作用就是在子进程Fork成功后创建一个管道（pipe），用于在子进程和父进程间进行通信。在子进程中，有些操作需要在子进程自己完成，在父进程中不需要或不允许执行，例如关闭监听的socket和清空锁的等等。

在这个函数中，我们还禁用了CPU分析器，并通过管道向父进程发送了一个消息，告诉它我们已经完成了在子进程中必须完成的操作，并准备接受新任务。



### forkAndExecInChild

forkAndExecInChild是一个函数，用于在新的子进程中执行指定的命令。它主要包括以下几个步骤：

1. 首先，使用fork函数创建一个新的子进程。

2. 然后，在子进程中使用execve函数执行指定的命令。execve函数会将当前进程替换为新的进程，从而完成命令的执行过程。

3. 在子进程中，如果execve函数调用失败，则会返回一个错误信息。此时，子进程会通过exit函数退出，并向父进程发送一个SIGCHLD信号，通知父进程子进程已经退出。

4. 在父进程中，通过调用wait4函数等待子进程的退出状态。如果子进程成功退出，父进程将会收到SIGCHLD信号，并且可以通过wait4函数获取子进程的退出状态。

总的来说，forkAndExecInChild函数的作用是创建一个新的子进程，并在子进程中执行指定的命令。它通常会被用来创建新的进程，并在新的进程中执行特定的任务，比如执行系统命令或者启动新的服务进程。



### capToIndex

capToIndex函数的作用是将Linux内核的安全能力（也称为“功能”）映射到在syscall/exec_linux_amd64.go中定义的安全能力索引。

在Linux内核中，每个安全能力均由一个位（bit）表示，该位在每个进程的有效能力集（effective capability set）中设置或清除。capToIndex函数的任务是在能力值和索引之间进行转换，不仅用于最终调用execve系统调用时设置进程的有效能力集，还用于其他操作，例如检查进程的能力集是否满足安全要求等。

此函数的实现并不复杂，它使用位操作和一个简单的映射表将安全能力值转换为相应的索引。由于其高效性和可重用性，该函数被广泛用于系统调用库的多个部分以及其他Go程序中。



### capToMask

capToMask函数的作用是将一个能力集合的字符串表示转换为对应的整型权限掩码。在Linux中，每个进程都有一组能力（capabilities），它控制了进程的各种系统特权。这些权限可以通过setcap和getcap命令设置和查询。

capToMask函数接受一个字符串参数，该字符串是一个用“+”和“-”表示添加和删除权限的列表。例如，capToMask("cap_setuid,+ep")表示将cap_setuid权限设置为开启，并且将EPERM权限添加到现有权限集中。

函数内部首先创建一个uint64类型的掩码变量mask，然后依次遍历字符串中的每个权限。如果权限的第一个字符是“+”，则将对应位设置为1，如果是“-”则将对应位清空为0。最后返回掩码变量mask。

这个函数主要是为了方便程序员对Linux的进程权限进行管理，将权限字符串转换为对应的掩码，然后可以通过setuid等系统调用进行设置。



### forkAndExecInChild1

在syscall包中，exec_linux.go文件是用于实现Linux系统调用相关的函数的。其中，forkAndExecInChild1函数的作用是在子进程中执行命令。

具体来说，该函数被调用时会先通过调用clone函数创建一个新的进程，然后在新的进程中执行命令。该函数实现的是一个纯粹的子进程，在该子进程中会进行资源清理、环境变量、参数等的处理，最终调用execve系统调用执行命令。

在执行命令的过程中，该函数还会根据用户传入的参数，设置一些环境变量、文件描述符、信号等信息，以确保命令可以正确的执行。最后，该函数会返回一个错误值，用于判断命令是否执行成功。



### forkExecPipe

forkExecPipe函数在syscall package中的exec_linux.go文件中实现，其作用是创建一个管道(pipe)，用于父进程和子进程之间的通信。

在forkExecPipe函数中，首先使用syscall.Pipe函数创建两个文件描述符(fd)，一个用于读取(pipe[0])，一个用于写入(pipe[1])。接下来，通过syscall.RawSyscall6调用进行进程复制（fork），创建子进程，如果创建成功则执行以下操作：

1. 将新进程的文件描述符表[0]设置为pipe[1], 并读文件描述符表[1]。
2. 如果该进程不是组长进程，则该进程成为新的会话组长。
3. 关闭父进程的pipe[1]和子进程的pipe[0], 将父进程的pipe[0]赋值给stdin, 将子进程的pipe[1]赋值给stdout和stderr, 最后返回子进程PID。

这样，父子进程就可以通过这个pipe来传递数据，并在子进程中执行指定的命令，同时保持对终端的控制权。



### formatIDMappings

func formatIDMappings(mappings []SysProcIDMap) ([]byte, error)

formatIDMappings函数用于将一个SysProcIDMap类型的切片序列化为一个用于execve系统调用的字节数组，以便在启动容器时使用。这个字节数组包含了两个整数，每个整数用来表示一个命名空间的ID，然后是新或旧命名空间的ID映射。这些ID映射指定了在启动容器时哪个进程实例中的命名空间应该继承到容器中。

基本上，SysProcIDMap表示进程的命名空间ID映射，将容器进程的命名空间映射到其父进程的命名空间。这些ID映射确定哪个命名空间将被转移到容器中。

formatIDMappings函数将这些ID映射转换为用于execve系统调用的字节数组，以便在使用Linux namespaces时在容器中正确地设置套接字地址。



### writeIDMappings

writeIDMappings是一个函数，用于将进程的用户和组ID映射写入/proc/<pid>/uid_map和/proc/<pid>/gid_map文件。这些文件是Linux内核用于实现用户命名空间的一部分，用于将进程的用户和组ID映射到不同的ID空间。

在用户命名空间中，进程可以使用不同的用户和组ID，在不同的命名空间中对同一个ID进行操作，进程看到的是完全不同的用户或组ID。这些ID映射文件允许进程将ID映射到该命名空间的不同范围内，因此其看起来就像是在不同的命名空间中执行。

writeIDMappings函数的作用是将进程映射ID写入文件中，该函数首先打开/proc/<pid>/uid_map和/proc/<pid>/gid_map文件，然后写入映射信息。每行写入一个映射：原始ID、映射ID、映射范围。如果成功映射，则返回nil，否则返回错误。如果映射范围未被指定，则操作系统假定映射范围为1，即一个用户或组ID映射到一个不同的ID。

writeIDMappings函数是Linux内核中实现用户命名空间的一个重要部分，它允许进程在不同的命名空间中使用不同的用户和组ID，并让进程看起来就像是在不同的命名空间中执行。这使得Linux内核更加灵活，因为它允许进程在进程隔离的环境中执行，同时确保安全性和资源隔离。



### writeSetgroups

writeSetgroups函数在执行execve系统调用时被调用，通过向进程的辅助组ID列表（supplementary group ID list）中写入新的辅助组ID列表。在Linux中，每个进程都有一个主要的组ID，用于决定进程访问权限，同时还可以拥有多个辅助组ID，这些辅助组ID也可以影响进程的访问权限。writeSetgroups函数将新的辅助组ID列表写入进程的asu_groups字段中，以便子进程能够获取正确的权限。

该函数先将辅助组的数量通过一个二进制序列写入到一个缓冲区中，并将该缓冲区传递给管道。然后通过循环，将每个辅助组的ID逐个写入到缓冲区中，并传递给管道。最后，管道中的数据将被写入到asu_groups字段中，供子进程使用。

需要注意的是，如果进程不需要改变辅助组ID，这个函数将不会被调用。此外，如果操作系统的版本低于2.6.23，这个函数也不会被调用。



### writeUidGidMappings

writeUidGidMappings函数是在子进程中运行的，它的作用是将用户和组的映射写入到/proc/PID/uid_map和/proc/PID/gid_map文件中，实现用户和组的隔离。

具体来说，当一个进程使用clone或unshare创建一个新的namespace时，它需要映射父进程中的用户和组到新的namespace中，以确保在新的namespace中只有指定的用户和组是可见的。writeUidGidMappings函数就是用来实现这个映射的。

在实现过程中，writeUidGidMappings函数首先会检查是否需要进行映射，如果不需要则直接返回。如果需要进行映射，则通过调用syscall.Syscall函数来把映射写入到/proc/PID/uid_map和/proc/PID/gid_map文件中。具体来说，它会将映射写入到一个缓冲区中，再通过syscall.Syscall函数来写入到文件中。

总的来说，writeUidGidMappings函数的作用是实现用户和组的隔离，在Linux中有着重要的作用，它可以保证用户和组在不同的namespace中具有不同的身份，以提高系统的安全性和可靠性。



