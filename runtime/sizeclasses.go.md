# File: sizeclasses.go

sizeclasses.go文件是Go语言运行时系统的源码之一，其作用是定义了一组内存块大小等级（称为大小等级表），用于在运行时分配等大小的内存块。这些大小等级表是通过在运行过程中缓存一定数量的常见内存块大小来提高内存分配的效率。

具体来说，sizeclasses.go中定义了三个大小等级表：TinySizeClass、SmallSizeClass和LargeSizeClass。TinySizeClass定义了从0到15个字节的内存块大小，SmallSizeClass定义了从16字节到32768字节的内存块大小，LargeSizeClass定义了大于32768字节的内存块大小。每个大小等级表中，内存块大小是以2的幂次方的方式递增的，以便易于计算。

除了定义大小等级表外，sizeclasses.go还定义了与内存分配相关的常量，例如内存块的对齐方式、缓存池中每个大小等级表缓存的内存块数量、堆中每次增长的内存块数量等。这些常量均可通过修改sizeclasses.go中的变量值来调整内存分配的性能和行为。

总之，sizeclasses.go的作用是提高Go语言运行时系统在动态内存分配方面的性能和效率，使其更加适合高效的并发编程。




---

### Var:

### class_to_size

在Go语言的运行时中，内存分配器需要考虑的一个重要因素就是如何高效地利用内存。为了实现这个目标，运行时会更具不同大小对象的需求，提前将一定范围内的大小进行分类，并为每个大小段预留一定大小的内存池，以便快速分配和回收内存块。

而在运行时中，大小分类的具体实现就是通过class_to_size变量来完成的。这个变量是一个数组，其中每个元素表示了一种内存块的大小，以字节为单位。数组的下标则是该内存块所属的大小分类，即对应的内存池编号。如：

```go
var class_to_size = []uint16{
    8, 16, 32, 48, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320, 384, 448, 512, 640, 768, 896,
    1024, 1280, 1536, 1792, 2048, 2560, 3072, 3584, 4096, 5120, 6144, 7168, 8192, 10240, 12288,
    14336, 16384, 20480, 24576, 28672, 32768, 40960, 49152, 57344, 65536, 81920, 98304, 114688,
    131072, 163840, 196608, 229376, 262144, 327680, 393216, 458752, 524288, 655360, 786432,
    917504, 1048576, 1310720, 1572864, 1835008, 2097152, 2621440, 3145728, 3670016, 4194304,
    5242880, 6291456, 7340032, 8388608, 10485760, 12582912, 14680064, 16777216,
}
```

通过这个数组，内存分配器就可以快速地确定所需内存大小所属的分类，从而直接从相应的内存池中分配所需内存。这样可以避免在每次分配内存时进行复杂的计算，提高内存分配效率。因此，class_to_size变量对于运行时中内存分配器的高效运行非常重要。



### class_to_allocnpages

在Go语言的运行时系统中，sizeclasses.go文件中的class_to_allocnpages变量用于存储每个 size class 能够容纳的对象所需的页数。该变量是一个映射，键为 size class 的编号，值为相应 size class 可分配的对象需要的页数。

在 Go 语言中，每个对象都有一个对应的 size class，size class 确定了对象的大小，并且能够帮助运行时系统进行内存分配。每个 size class 包含一组大小相似的对象，例如，位于同一 size class 的对象大小可能都在 16 个字节到 32 个字节之间。为了有效使用内存，运行时系统会用一些技巧来尽可能地将对象放置在连续的内存中，并且在进行内存分配时，尽可能地减少空间浪费。

对于每个 size class，运行时系统需要知道要分配多少内存页面，以确保分配足够的内存。这就是 class_to_allocnpages 变量的作用。它存储了每个 size class 所需的页面数，以便运行时系统可以快速地查找并分配内存。如果没有这个变量，运行时系统需要在执行内存分配时动态计算页数，这将导致分配延迟和运行时开销的增加。



### class_to_divmagic

在Go语言的垃圾回收器中，内存分配的大小需要精确控制，以避免内存的浪费和分配效率的下降。为了实现这个目的，Go语言使用了大小分类（size classes）的技术，将特定范围内的内存大小归为一类，并为每个类分配固定大小的堆块，从而避免了动态分配内存所带来的运行时开销和碎片问题。

而class_to_divmagic这个变量，则是在将要分配内存的大小转换为对应的size class时，用到的一个映射表。它是一个数组，每个元素代表一个内存大小范围所对应的size class编号，这个映射是通过计算内存大小的log2值，再进行一些复杂的位运算得到的。此外，class_to_divmagic这个变量还包含了一些用于内存对齐的魔数（magic number），用于在分配内存时进行大小调整。这样，当应用程序需要分配内存时，可以很快地确定所需要的size class，并从对应的堆块中分配内存，大大提高了内存分配的效率和减少了内存碎片。



### size_to_class8

在Go语言中，内存分配是一个非常重要的任务，因为这直接关系到程序的性能和效率。在运行时中，存在一个名为size_to_class8的变量，它是一个长度为98的数组，该数组定义了在不同大小的对象上分配内存时应该采用哪些大小的堆块。

具体来说，size_to_class8的作用是将对象的大小映射到堆块的大小。在Go语言中，内存分配器会根据传入的对象大小通过查找size_to_class8数组来确定应该使用哪个堆块大小。这个数组中存储了堆块大小的信息，对于每个不同的堆块大小，对应的数组元素是一个结构体，包含了该堆块的信息，如大小和对齐方式等。

使用size_to_class8数组的好处是可以在分配内存时快速且准确地确定应该使用哪个堆块大小。这大大提高了内存分配的效率和性能，也降低了内存分配器的复杂性。

总之，size_to_class8变量在Go语言的运行时中具有非常重要的作用，它是整个内存分配器的核心组件之一，用于确定所需内存大小并选择合适的堆块进行分配。



### size_to_class128

size_to_class128变量在Go语言的runtime包中的sizeclasses.go文件中，是一个用于将内存块大小映射到与其大小相匹配的内存块类别的映射表。

具体地说，这个映射表将内存块的大小按照2的幂次方（从0到47）分成了48个不同的大小类别（称为“size class”），其中第0个大小类别（对应大小为0）是特殊的，其余47个类别的大小依次为8字节、16字节、32字节、64字节、...、4194304字节。每个类别中包含了多个大小相同的内存块，每个内存块的大小都等于该类别大小的两倍的幂次方。

这个映射表的作用是让程序在申请一块大小为s的内存块时，能够快速地找到与其大小最接近的内存块类别，并从相应的类别中分配内存块。这样做的好处是可以避免内存的浪费和碎片化，提高内存使用效率。



