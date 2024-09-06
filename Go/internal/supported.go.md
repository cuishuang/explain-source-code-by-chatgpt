# File: supported.go

supported.go文件在go/src/internal目录下，其作用是列举和定义Go语言中支持的操作系统和架构等相关信息，并提供了一些函数用于判断某个系统或架构是否被支持。

具体来说，supported.go文件中定义了以下内容：

1. archList：包含所有Go语言支持的架构列表，如"amd64"、"386"、"arm"等。

2. goosList：包含所有Go语言支持的操作系统列表，如"linux"、"darwin"、"windows"等。

3. buildTags：是一组用于编译条件指令的标签，用于在编译代码时选择特定的操作系统或架构。

4. func RuinsOnArch(arch string) bool：用于判断某个架构是否被支持。

5. func RunsOnGOOS(goos string) bool：用于判断某个操作系统是否被支持。

6. func MatchFile("", name string, fi os.FileInfo) (bool, error)：用于判断一个文件是否符合当前操作系统和架构的命名规则。

这些定义和函数可以在Go语言的源代码中被使用，以帮助程序员确定在特定的操作系统和架构上编译Go程序的正确方式。通过这些信息和函数，开发者可以编写出在各种操作系统和架构上运行良好的Go程序。

## Functions:

### RaceDetectorSupported

RaceDetectorSupported是Go语言内置工具的一部分，它的作用是判断当前的操作系统和架构是否支持竞争检测器。在Go语言中，“竞争”指的是同时对共享的内存中的同一个变量进行读写操作时可能出现的问题。通过竞争检测器可以检测到这些问题，从而帮助开发人员在编写代码的时候避免出现竞争问题。

在supported.go文件中，RaceDetectorSupported函数会根据当前的操作系统和架构信息，判断是否支持竞争检测器。如果支持，返回true，否则返回false。具体来说，判断的逻辑是：

- 如果当前操作系统不是linux、solaris、freebsd、darwin或windows其中之一，则不支持竞争检测器。
- 如果当前程序不是以-race标志编译的，则不支持竞争检测器。
- 如果当前架构不是386、amd64、arm、arm64、ppc64、ppc64le或mips64其中之一，则不支持竞争检测器。

如果以上三个条件都符合，则该函数返回true，表示当前操作系统和架构支持竞争检测器。这个函数能够帮助Go语言开发人员判断是否可以使用竞争检测器来避免代码中的竞争问题，从而提高程序代码的质量。



### MSanSupported

MSanSupported是一个函数，它的作用是检测系统或编译器是否支持MemorySanitizer（简称MSan）。

MSan是一个内存动态检测工具，它通过在运行时检测内存访问错误来帮助发现和修复内存错误，如使用未初始化的数据、访问已释放的内存、使用已释放的指针等。对于使用C/C++等语言编写的代码，MSan可以帮助开发人员提高代码质量和可靠性。

由于MSan需要在编译时链接特定的库文件，因此需要检查系统或编译器是否支持MSan。在supported.go文件中，MSanSupported函数通过以下步骤检测当前系统或编译器是否支持MSan：

1. 检查当前操作系统是否为Linux或FreeBSD，因为MSan仅在这些系统上受支持。
2. 检查当前编译器是否支持MSan编译选项。如果支持，则返回true；否则返回false。
3. 检查当前系统是否有msan启动脚本，并尝试执行该脚本以获取msan是否可用的结果。

如果以上步骤中的任何一步返回true，则MSanSupported函数返回true，表示当前系统或编译器支持MSan；否则返回false，表示不支持。

总之，MSanSupported函数可以帮助开发人员确定当前系统或编译器是否支持MemorySanitizer，并在必要时采取相应的措施以确保代码的质量和可靠性。



### ASanSupported

ASanSupported这个func的作用是判断当前操作系统是否支持ASan（AddressSanitizer）检测工具，并返回是否支持。

AddressSanitizer是一种内存错误检测工具，可以通过检测指针和heap分配来捕获危险的访问、内存泄漏和其他一些与内存相关的问题。在C和C++等语言中，内存错误是常见的导致程序崩溃和安全漏洞的主要原因之一。ASan对于发现和修正这些问题非常有用。

ASanSupported函数的实现方式因操作系统而异。对于Linux系统，它使用go env命令查找LIBRARY_PATH变量中的链接器文件夹，判断是否存在libasan.so库文件来判断是否支持ASan。对于其他系统，如Windows，目前尚不支持ASan，因此始终返回false。

ASanSupported函数的结果对Go语言内部的一些初始化和编译时期的决策起着重要作用，如在构建Go二进制文件时，会判断是否需要加入ASan相关的选项来支持ASan检测。因此，对于需要使用ASan工具的项目，在构建之前需要确保运行环境支持ASan。



### FuzzSupported

FuzzSupported函数是Go语言中的内部函数，用于确定当前运行Go程序的平台是否支持使用fuzz测试工具进行测试。

fuzz测试是一种随机化测试技术，它通过随机生成输入数据来发现程序中的潜在漏洞或错误。在Go语言中，可以使用go-fuzz工具进行fuzz测试，但该工具只支持特定版本的Go语言和特定的操作系统/体系结构。为了确保在运行fuzz测试时不会发生错误或异常，需要确保当前运行程序的平台支持fuzz测试。

FuzzSupported函数就是用于判断当前平台是否支持fuzz测试的函数。它首先检查当前Go语言版本是否支持fuzz测试，如果是则继续检查操作系统和体系结构是否支持fuzz测试。如果当前平台支持fuzz测试，则返回true，否则返回false。

具体实现方式可以见代码：

```go
// FuzzSupported reports whether the current platform supports fuzzing.
func FuzzSupported() bool {
    // supportedGoVersion为最低支持fuzz测试的版本号
    const supportedGoVersion = 0x010d // We require Go 1.13.
    if version.Version < supportedGoVersion {
        return false
    }

    // 检查操作系统和体系结构是否支持fuzz测试
    switch runtime.GOOS {
    case "android", "ios":
        // Android and iOS are supported only if built for testing, aka race mode.
        return race.Enabled
    case "freebsd", "linux", "netbsd", "openbsd", "dragonfly":
        switch runtime.GOARCH {
        case "amd64", "arm64", "ppc64le", "s390x":
            return true
        default:
            return false
        }
    case "darwin":
        switch runtime.GOARCH {
        case "amd64":
            return true
        default:
            return false
        }
    case "windows":
        switch runtime.GOARCH {
        case "amd64":
            return true
        default:
            return false
        }
    }

    return false
}
```

总结来说，FuzzSupported函数的主要作用是检查当前平台是否支持fuzz测试，以确保在运行fuzz测试时不会发生错误或异常。



### FuzzInstrumented

FuzzInstrumented函数是用于启用模糊测试的。在Go语言的标准库中，支持使用Fuzz测试来在特定输入上测试函数的行为。Fuzz测试会生成随机的输入，并将它们提供给待测函数。在测试过程中，模糊测试引擎会使用一系列固定规则来生成随机输入。通过在不同的随机生成的输入上测试代码，我们可以发现并解决一些隐藏的错误和漏洞。

FuzzInstrumented函数会使用Go语言的fuzz包来对函数进行模糊测试的检测。它接收两个参数:

- data []byte：即作为输入提供给待测函数的随机数据。
- goFuzz bool：指示是否应该以Go模糊测试方式运行。

当goFuzz参数为true时，FuzzInstrumented函数会调用fuzz包中的Fuzz方法来测试函数的行为。Fuzz方法的参数是一个回调函数，它是待测函数的包装函数。在Fuzz方法内部，会生成随机输入并将其作为回调函数的输入。随机输入会被分成多个块，以测试函数的行为，同时还允许测试程序的稍后状态检查。Fuzz测试会尝试将大量随机化的输入提供给已编译的二进制文件，以验证是否能够实现正确的行为。

因此，FuzzInstrumented函数允许在Go语言的标准库中启用Fuzz测试，从而帮助发现隐藏的错误和漏洞。



### MustLinkExternal

在Go中，MustLinkExternal函数是一个内部函数，用于链接外部库。

具体来说，当Go代码需要访问一个外部库时，它需要通过链接器将其连接到输出二进制文件中。然而，有时链接过程可能会失败（例如，由于缺少必需的库文件或符号）。在这种情况下，程序将无法启动。

MustLinkExternal函数的作用是在编译时验证外部库是否可用，并在链接时将其连接到输出文件中。如果链接失败，则会引发一个panic，以提醒开发人员将缺失的库文件或符号添加到编译环境中。

这个函数的使用场景通常发生在Go的一些内部包中，例如internal/syscall/windows和internal/poll。这些包需要使用Windows的系统调用和异步I/O等功能，这些功能需要链接Windows下的DLL文件。使用MustLinkExternal可以确保这些DLL文件在链接时已经可用，以避免运行时错误。

总之，MustLinkExternal函数是Go的一个重要特性，它可以保证外部库的可用性，并提高程序的稳定性和可靠性。



### BuildModeSupported

BuildModeSupported函数是一个用于检查当前编译器是否支持指定的BuildMode的辅助函数。

在Go语言中，BuildMode定义了编译器使用哪种模式生成代码。常见的BuildMode有以下几种：

- "exe": 生成可执行文件
- "c-archive": 生成C语言静态库
- "c-shared": 生成C语言动态库
- "shared": 生成Go语言动态库
- "plugin": 生成Go语言插件

BuildModeSupported函数接受一个BuildMode参数，并返回一个bool类型的值，表示当前编译器是否支持指定的BuildMode。如果支持，返回true，否则返回false。

对于不同的编译器，BuildModeSupported函数的实现也不同。在Linux和macOS系统下，它会检查编译器是否支持cgo，因为cgo是生成动态库和插件的必要组件。在Windows系统下，则会检查编译器是否支持PE/COFF格式。

通过BuildModeSupported函数，我们可以在编写跨平台的Go代码时，根据不同的编译器特性来进行代码分支，从而确保生成的可执行文件或库在不同平台上都能够正常运行。



### InternalLinkPIESupported

InternalLinkPIESupported是一个函数，位于Go语言标准库中的go/src/internal/supported.go文件中。该函数的作用是检查当前操作系统是否支持链接PIE（Position Independent Executable）的功能。

链接PIE是一种编译选项，它使得编译后的可执行文件可以在不同的内存地址上进行加载并运行，这可以提高可执行文件的安全性和可移植性。InternalLinkPIESupported函数的作用是检查当前操作系统是否支持链接PIE选项，如果支持，则返回true，否则返回false。

在函数内部，它首先使用go/build包中的Context函数从运行时环境中获取操作系统的名称（例如linux、darwin、windows等）。然后，它通过调用当前操作系统的命令行工具（例如gcc或ld）来编译一个简单的测试程序，并尝试在编译过程中使用PIE选项来编译这个程序。如果编译成功，则说明当前操作系统支持PIE选项，函数返回true。

总的来说，InternalLinkPIESupported函数的作用是帮助Go语言编译器了解当前操作系统的特性，从而更加智能地生成可执行文件。



### DefaultPIE

在go/src/internal/supported.go文件中，DefaultPIE是一个函数，它的作用是根据当前操作系统的架构和编译器类型返回适当的值。

在许多操作系统中，程序可以以两种不同的方式编译：静态链接和动态链接。静态链接将程序中需要的所有库与程序本身一起编译成一个可执行文件，而动态链接则将程序与使用的库分开编译。在动态链接中，库被加载到内存中，程序通过链接库的地址来访问库中的函数。

对于动态链接的程序，可以使用可执行文件的位置无关代码（PIC）功能，以便将其与其他共享库在内存中加载时进行合并。这是通过使用装载可执行程序（PIE）来实现的。PIE文件是一种特殊的位置无关可执行文件，可以让程序能够在加载时进行地址重定向，而无需事先知道它们将在内存中的位置。

DefaultPIE函数通过检查goenv.GOARCH和goenv.Compiler的值，以确定操作系统架构和编译器类型，然后返回适当的值来确定是否在构建时将PIE标志设置为默认值。如果操作系统和编译器支持PIE，则返回true；否则返回false。

总之，DefaultPIE函数是用于确定是否将PIE标志设置为默认值的函数，以实现更好的内存管理，并提高程序的安全性和性能。



### CgoSupported

该文件中的CgoSupported函数主要用于检测是否支持CGO。

CGO（C to Go）是Go语言中用于调用C语言库的工具。使用CGO可以让Go程序使用C语言编写的动态库等程序。

在实现中，CgoSupported函数首先检查环境变量GODEBUG中是否设置了cgocheck，默认为“0”，如果为“1”或“2”则输出警告信息。然后它使用assembly代码尝试调用C函数并将结果保存到变量中。如果可以成功执行该操作，则说明CGO可用，并且该函数将返回true，否则返回false。

该函数的作用是为了在程序开始运行时检查是否支持CGO，以便在程序中使用CGO时进行必要的报错或处理，以确保程序能够正常运行。在某些情况下，如果CGO不可用或不能使用，需要使用其他方式调用C语言库，在这种情况下，可以根据该函数的返回值进行相应的处理。



### ExecutableHasDWARF

ExecutableHasDWARF函数是Go语言的一个内部方法，用于检查指定的二进制文件中是否包含DWARF调试信息。DWARF是一种调试信息格式，用于存储程序的调试信息，包括源代码、变量、函数和行号等信息。通过这些信息，开发人员可以在程序崩溃或出现异常时快速诊断和修复问题。

ExecutableHasDWARF函数接受一个文件路径作为参数，并使用debug/dwarf包来解析该文件中的DWARF调试信息。如果DWARF调试信息存在，则返回true；否则返回false。通过检查二进制文件是否包含调试信息，可以帮助开发人员了解是否需要在程序崩溃或出现异常时记录调试信息，以便更好地定位问题。

此外，ExecutableHasDWARF函数还可以在编译器生成二进制文件时进行优化。例如，如果二进制文件中不包含DWARF调试信息，则编译器可以对代码进行更多的优化以提高程序的性能。因此，ExecutableHasDWARF函数的结果可能会影响编译器生成的二进制文件的代码质量。



