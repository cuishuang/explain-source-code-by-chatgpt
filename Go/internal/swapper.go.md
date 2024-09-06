# File: swapper.go

swapper.go是go语言中的内部包internal中的一个文件，其作用是提供一些底层的交换函数和工具函数，用于支持一些高级特性的实现，比如chan类型、sync包中的对象和一些数据结构等。

具体来说，swapper.go中包含了以下几种函数：

- func SwapInt32(addr *int32, new int32) int32：用于原子交换一个int32类型的变量的值，返回修改前的值；
- func SwapInt64(addr *int64, new int64) int64：用于原子交换一个int64类型的变量的值，返回修改前的值；
- func SwapUint32(addr *uint32, new uint32) uint32：用于原子交换一个uint32类型的变量的值，返回修改前的值；
- func SwapUint64(addr *uint64, new uint64) uint64：用于原子交换一个uint64类型的变量的值，返回修改前的值；
- func SwapUintptr(addr *uintptr, new uintptr) uintptr：用于原子交换一个uintptr类型的变量的值，返回修改前的值；
- func SwapPointer(addr unsafe.Pointer, new unsafe.Pointer) unsafe.Pointer：用于原子交换一个unsafe.Pointer类型的变量的值，返回修改前的值；
- func CopySlice(to, fr unsafe.Pointer, elemsize, n uintptr)：用于将一个unsafe.Pointer类型的切片从一个地址复制到另一个地址，其中elemsize表示每个元素的大小，n表示要复制的元素的个数。

除了上述函数外，swapper.go中还包含了一些工具函数，例如：

- func RaceDisable() bool：用于关闭Go语言的竞争检测器，以提高性能；
- func RaceEnable()：用于开启Go语言的竞争检测器，以保证程序的正确性。

总的来说，swapper.go提供了一些底层的交换函数和工具函数，这些函数可以被其他高级特性和数据结构使用，是go语言的内部实现的重要组成部分。

## Functions:

### Swapper

Swapper是一个通用函数，用于在容器中交换两个元素的位置。它是在sort包中使用的，但它实际上定义在内部包internal中，这是为了强制它只能被Go标准库使用，不能被外部包导入。

该函数的类型为func(i, j int)，方法签名之所以使用这样的参数类型，是因为它与sort.Slice()方法兼容。sort.Slice()方法允许用户对容器的任何元素应用一个自定义函数，并将其作为参数传递给安排这些元素的算法。Swapper方法可以被认为是sort.Slice()函数的一部分，它用来协调交换容器元素的请求。

当你在sort包的内部函数中需要交换两个元素的位置时，你可以直接调用Swapper i，j。

总之，Swapper是一个通用函数，它用来交换一个容器中的两个元素的位置。它是通过sort.Slice()方法实现的，它可以被认为是sort.Slice()函数的一部分，并且是在内部包internal中定义的，这使得它不能被外部包导入和使用。



