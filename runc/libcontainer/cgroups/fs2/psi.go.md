# File: runc/libcontainer/cgroups/fs2/psi.go

在runc项目中，runc/libcontainer/cgroups/fs2/psi.go文件的作用是监控容器中的Pressure Stall Information（PSI）数据。PSI是一种从内核获取的性能指标，用于检测资源压力对系统性能的影响。

该文件中的statPSI函数用于读取并解析容器中的PSI数据。该函数通过读取文件系统中的相关文件（如/proc/[pid]/psi/fs_status）来获取容器内部的PSI数据。

函数parsePSIData用于解析从statPSI函数获取的PSI数据。它从字符串解析器中解析出各个PSI指标（如CPU、内存和磁盘等）的压力级别和各个窗口的计数器。该函数的作用是将这些解析出的数据保存在PSIStats结构体中，以便在后续的代码中使用。

总的来说，这个文件的作用是获取容器中的PSI数据，并将其解析为可用的结构以供后续使用。这些PSI数据可以帮助监控和调优容器的性能。

