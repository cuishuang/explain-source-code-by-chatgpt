# File: numcpu_freebsd.go

numcpu_freebsd.go是Go语言运行时库中的一个文件，它的主要作用是实现在FreeBSD操作系统上获取可用CPU的数量。具体来说，它提供了一个名为numCPU的函数，可以返回当前系统中可用的CPU数量。

在实现过程中，numcpu_freebsd.go使用了FreeBSD操作系统特定的系统调用来获取比特定CPU的数量。这些系统调用包括sysctl和cpuset_getaffinity，并通过相应的错误处理来确保函数的稳定性和可靠性。

该文件的存在，是为了确保Go语言在FreeBSD上的性能和可用性，以便能够更好地支持在该操作系统下的应用程序开发。




---

### Var:

### cpuSetRE

在FreeBSD操作系统中，CPU亲和性可以通过setsockopt系统调用中的CPU_SET和CPU_CLR选项来设置或清除。在numcpu_freebsd.go文件中，cpuSetRE变量是一个正则表达式，用于解析系统调用的返回值并提取CPU集合。正则表达式中的“CPU ([0-9]+)”匹配CPU编号，并使用“,”和“-”字符匹配CPU范围。CPU集合中的每个CPU都使用＆运算符和1左移（CPU编号 mod 64）位来生成掩码，并将所有掩码放在一个数组中，以便将其传递给操作系统以设置CPU亲和性。因此，cpuSetRE变量在numcpu_freebsd.go文件中的作用是解析FreeBSD系统调用的返回值并提取CPU集合。



## Functions:

### init

init函数是go语言中一个特殊的函数，当程序启动时会自动执行在该包中定义的init函数。在numcpu_freebsd.go文件中，init函数的作用是初始化CPU数量，获取系统的CPU数量，并设置GOMAXPROCS变量为CPU数量的值。

具体来说，init函数首先调用了runtime包中的getncpu函数，该函数的作用是返回系统的CPU数量。然后，如果返回的结果大于0，就将GOMAXPROCS变量设置为该值；否则，将GOMAXPROCS设置为1，表示使用单个CPU来运行。最后，init函数将程序所使用的CPU数量打印出来。

这个init函数的主要作用就是计算要在本地机器上使用的CPU数量。这对于在并发编程中非常重要，因为它可以让程序利用多核CPU，从而实现更高的性能。



### FreeBSDNumCPUHelper

FreeBSDNumCPUHelper是用于在FreeBSD系统中获取CPU核心数量的辅助函数。

该函数会通过sysctl系统调用查询FreeBSD系统的hw.ncpu参数，该参数表示系统中可用的CPU核心数量。如果查询失败，则会返回1表示只有一个CPU核心可用。

该函数在运行时（runtime）中被使用，用于确定可用的处理器核心数量，以便更好地利用多核处理器的性能。在代码编译和执行期间，程序无法确定可用的处理器核心数量，因此需要在运行时通过系统调用查询。



### FreeBSDNumCPU

FreeBSDNumCPU函数的作用是返回当前系统中CPU的数量。它通过读取FreeBSD系统文件来获取CPU的数量，并返回它们的数量。

具体来说，函数首先尝试打开/sys/devices/system/cpu/present文件，该文件包含当前系统可用的CPU范围。如果此文件不存在，则该函数返回默认CPU数量为1。

然后，函数打开/sys/devices/system/cpu/online文件，该文件列出了当前在线的CPU。如果该文件不存在，则返回present文件中列出的CPU数量。

最后，函数返回在present和online文件中列出的CPU数量的最小值。这是因为系统可能会启用所有可用的CPU，但在线的CPU数量可能会受到诸如节能模式之类的因素的影响而减少。

这个函数在Go程序中用于调整一些与CPU相关的操作，如调度goroutine，自适应GC等，以优化程序的性能。



### getList

getList函数是用于获取FreeBSD系统中的CPU列表信息的函数。它使用了FreeBSD系统特定的系统调用来获取处理器的信息。

该函数首先初始化了一个sysctl结构体，然后使用了sysctlbyname函数，获取了hw.ncpu和hw.physicalcpu两个系统参数的值。其中，hw.ncpu表示逻辑CPU的数量，hw.physicalcpu表示物理CPU的数量。

接着，该函数又通过调用kvm_open函数创建了一个KVM句柄（即内核虚拟机）。接下来，它使用KVM句柄和KVM_GET_CPUID2命令来获取处理器的详细信息。最后，该函数将获取到的处理器信息保存在一个processors切片中，并返回该切片。

基本上，该函数的作用就是获取FreeBSD系统中的处理器信息。由于不同系统中的处理器信息不尽相同，因此需要专门的函数来获取不同系统中的处理器信息。这正是getList函数的作用所在。



### checkNCPU

checkNCPU函数的作用是检测当前系统上的CPU核数，并返回一个非负整数表示CPU的数量，如果无法检测到CPU数量，则返回0。

在FreeBSD操作系统下，CPU数量可以通过sysctl命令获取。checkNCPU函数通过调用sysctl命令获取CPU的数量，并将其存储在全局变量NCPU中，以便在其他部分使用。

该函数首先尝试通过sysctl获取CPU数量，如果获取成功，则将其存储在NCPU中。如果无法获取CPU数量，则没有操作。

该函数在runtime包初始化期间被调用，以确保NCPU变量在程序运行期间可用。



