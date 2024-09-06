# File: cpu_wasm.go

cpu_wasm.go文件是为了支持WebAssembly平台而编写的CPU信息收集和处理代码。WebAssembly是一种新型的底层应用程序技术，可以在多种平台上运行，包括Web浏览器、桌面应用程序等。

在cpu_wasm.go文件中，定义了几个函数，用于获取WebAssembly平台上的CPU架构信息、CPU核心数、CPU频率等信息。这些函数包括getCPUArch()、getCPUCount()和getCPUMHz()等。

由于WebAssembly平台有其独特的限制，比如没有直接访问系统硬件的权限，cpu_wasm.go文件通过JavaScript代码与Web浏览器交互，从而能够获取到WebAssembly平台上的CPU信息。例如，通过JavaScript中的navigator.hardwareConcurrency属性就可以获取到CPU核心数。

cpu_wasm.go文件的主要作用是在Go语言程序中提供了与WebAssembly平台交互的功能。可以在Go语言程序中使用这些函数获取WebAssembly平台上的CPU信息，从而更好地了解程序的运行环境，优化程序的性能。

## Functions:

### doinit

文件`cpu_wasm.go`记录的是Wasm平台下的CPU信息：

```go
//go:build js && goexperiment.regabi
// +build js,goexperiment.regabi

package internal

import (
    "runtime/internal/sys"
    "unsafe"
)

var cpuid_init_once_flag uint32

var cpuid_once_init_data [4]uint32

func cpuid(reg uint32) (eax, ebx, ecx, edx uint32)

//go:linkname doinit runtime.doinit
func doinit() {
    if cpuid_init_once_flag == 0 {
        eax, ebx, ecx, edx := cpuid(0x40000000)
        if ebx == 0x4E585355 { // UXN"XN"
            // We are running on Wasm platform. Print "Wasm" to stdout
            // with stripped "XN".
            e := make([]byte, 4)
            *(*uint32)(unsafe.Pointer(&e[0])) = ebx
            print("Wasm ", string(e[:3]), "\n")
            // Copy cpuid data to safe memory.
            cpuid_once_init_data[0] = eax
            cpuid_once_init_data[1] = ebx
            cpuid_once_init_data[2] = ecx
            cpuid_once_init_data[3] = edx
        }
        cpuid_init_once_flag = 1
    } else {
        eax, ebx, ecx, edx = cpuid_once_init_data[0],
            cpuid_once_init_data[1], cpuid_once_init_data[2], cpuid_once_init_data[3]
    }

    cpu.init(eax, ebx, ecx, edx)
}

type cpuinfo struct {
    typ     sys.CPUType
    family  uint8
    model   uint8
    stepping uint8
    feature [2]uint32
}

var cpu cpuinfo
```

其中，`doinit`函数的作用是在程序启动时初始化CPU信息（通过调用`cpuid()`函数获取CPU的型号、家族、模型、特性等信息），并将其存储在全局变量`cpu`中以便后续程序中使用。

具体来说，当程序第一次调用`doinit`时，它会检查变量`cpuid_init_once_flag`是否为0来确定是否已经初始化CPU信息。如果`cpuid_init_once_flag`为0，它会调用`cpuid()`函数获取CPU信息，并将其存储到全局变量`cpu`中。如果`cpuid_init_once_flag`不为0，则说明已经初始化了CPU信息，此时`doinit`会直接读取全局变量`cpu`中存储的信息。

需要注意的是，`doinit`函数最终被链接到了`runtime.doinit`函数，也就是说，当程序运行时，除了`internal`包外的所有包的`init()`函数都会在`runtime.doinit`函数中调用，进而触发`doinit`函数的调用。因此，当程序运行时，CPU信息会在程序启动时自动初始化。



