# File: malloc_test.go

malloc_test.go是Go语言标准库中runtime包的一个单元测试文件。该文件主要用于对runtime包中的内存分配器（malloc）进行各种测试，包括：

1. 内存分配和释放的正确性测试
2. 并发场景下内存分配和释放的正确性测试
3. 大内存分配和释放的性能测试
4. 内存池的性能测试

该文件中主要的函数包括TestLargeAlloc、TestConcurrentFree、TestParitalFull和BenchmarkPool，这些函数分别测试了上述场景下的内存分配和释放的正确性和性能。

通过这些测试，可以确保runtime包中的内存分配器在各种场景下都能正常工作，并具有一定的性能优势。同时也为Go语言使用者提供了一些内存控制和优化的技巧。




---

### Var:

### testMemStatsCount

在Go语言的运行时库（runtime）中，malloc_test.go文件是用于测试内存分配和回收机制的文件。testMemStatsCount是在测试中使用的变量，主要作用是记录并比较分配或释放内存的次数。

在测试用例中，使用了一些内存分配和回收的函数，然后使用testMemStatsCount变量来记录分配或释放内存的次数。这样就可以检查这些函数是否按照预期分配和释放了内存或者是否达到了预计的次数。

通过检查testMemStatsCount变量来验证内存分配和回收的次数，有助于检测是否存在内存泄漏或者内存浪费的问题。这对于保障应用程序的稳定性、性能和可靠性非常重要，因此testMemStatsCount在测试过程中扮演着非常重要的角色。



### arenaCollisionSink

在 Go 的 runtime 包中，malloc_test.go 文件中的 arenaCollisionSink 变量是一个空的结构体，它主要用于测试内存分配器，特别是在多线程环境下使用。这个变量用于检测内存分配器是否会造成堆空间内存碎片化，因为内存碎片化可能会导致内存分配器效率下降。

arenaCollisionSink 变量是在测试中使用的，主要用于记录堆中超过 2MB 的内存碎片的数量。这些内存碎片可能会导致内存分配器效率下降，因此需要使用 arenaCollisionSink 变量进行记录和分析。当有内存碎片发生时，它们的数量将被累加到 arenaCollisionSink 变量中。如果 arenaCollisionSink 变量的值增加，则表示可能存在内存碎片问题，需要进一步进行排查和分析。

总之，arenaCollisionSink 变量的作用是在内存分配器测试环境中监视内存碎片，以便进一步优化内存分配器的效率和性能。



### n

在go/src/runtime/malloc_test.go文件中，n变量是用于测试内存分配器（malloc）性能的输入参数。它代表要分配的对象数量。

具体来说，n变量在函数BenchmarkMalloc(b *testing.B)中使用。该函数会循环多次分配内存，并计算总共花费的时间。循环中每次分配内存都是使用指定数量的n个对象。

通过修改n的值，可以测试不同大小的内存分配情况，从而优化分配器的性能。通常的做法是逐步调整n的值，找到最优的分配策略。






---

### Structs:

### obj12

在Go语言中，obj12是runtime包用来描述内存块的结构体，它包含了一些重要的字段以及方法，来维护和管理内存块。其中，比较重要的字段包括：

1. size：表示内存块的大小，单位为字节；
2. next：指向下一个未分配的内存块；
3. mark：用于标记内存块是否被使用。

此外，obj12还有一些方法，用于初始化、分配以及回收内存块等操作。

在malloc_test.go这个文件中，obj12主要用于测试内存分配和回收的功能。通过模拟多个协程同时分配和回收内存块，来测试内存管理的正确性和性能。同时，obj12也可以用于性能分析和优化，可以根据它来了解和优化内存分配和使用的情况。



### acLink

acLink是一个结构体，用于表示一个AC自动机中的节点。AC自动机是一种字符串匹配算法，常用于文本搜索和模式匹配。

在malloc_test.go文件中，acLink结构体用于测试runtime包中的mheap对象中的tcache缓存机制。在测试中，acLink结构体被当作一个简单的对象，用于存储随机生成的指针和指针大小。

具体而言，acLink结构体定义如下：

type acLink struct {
    next  *[2]acLink // 子节点
    ptr   unsafe.Pointer // 随机生成的指针
    size  uintptr // 随机生成的指针大小
}

其中，next字段表示该节点的两个子节点（0和1），ptr字段表示随机生成的指针，size字段表示生成指针的大小。

在测试中，acLink结构体被用于生成一组随机指针，存储在tcache缓存中。该测试主要验证tcache缓存机制是否能够正确地缓存和释放这些指针，从而提高内存分配和释放的效率。

当然，acLink结构体本身并不是AC自动机算法的关键组成部分，仅仅是一个用于测试的中间结构体。



### LargeStruct

LargeStruct是在malloc_test.go文件中用于测试mallocgc函数的一个结构体。该结构体包含一个大型数组，用于测试在分配连续大块内存时mallocgc函数的性能和正确性。由于该结构体的大小超过了堆的缓存大小，因此在分配内存时需要使用mallocgc函数。

在该测试中，首先通过使用defaultAllocator（即堆）分配多个小的对象，然后再通过使用mallocgc分配LargeStruct结构体来测试mallocgc函数是否能够正确地扩展堆以容纳更大的对象。测试还会比较在使用默认分配器和使用mallocgc函数时对象分配的速度和堆的使用情况。

通过测试这些性能和正确性的指标，可以验证mallocgc函数的正确性和性能。此外，如果有必要，可以修改分配器的实现方式以优化内存分配的性能和可靠性。



## Functions:

### TestMemStats

TestMemStats是Go语言运行时包(runtime)中的一个测试函数，测试内存统计信息的准确性。它会调用runtime.ReadMemStats()函数获取当前内存使用情况，并将内存使用情况与预期值进行比较，以确保内存统计信息的正确性。

在TestMemStats函数中，会使用四个 goroutine 并发分配和释放内存，以模拟多线程环境下的内存使用情况，并通过runtime.Gosched()函数强制使当前 goroutine 让出 CPU，以增加并发场景的复杂性。程序还会通过设置 runtime.MemProfileRate 和 runtime.MemProfileThreshold 两个参数，来控制内存 profile 的生成方式，进一步提高内存使用场景的复杂性。

测试完成后，程序会输出当前内存使用情况的统计信息，包括内存总量、堆内存使用情况、栈内存使用情况、垃圾回收次数等。最后通过断言语句来比较内存使用情况和预期值是否相等，如果测试失败，则会输出相应的错误信息。

TestMemStats函数的作用是确保runtime包中的内存统计功能能够正常工作，为Go语言程序的内存管理提供参考依据，以确保程序的内存使用情况可以得到有效的监控和管理。



### TestStringConcatenationAllocs

TestStringConcatenationAllocs是一个测试函数，它用于检测字符串连接操作中的内存分配情况。

在该测试中，测试函数会创建一定数量的字符串，并使用“+”操作符将它们连接起来。然后，使用runtime.MemStats函数获取系统内存占用情况，可以得到在字符串连接过程中，分配了多少字节的内存。

这个测试能够帮助开发人员了解字符串连接在大数据量场景下的性能和内存开销问题。如果字符串连接过程中分配的内存过多，可能会导致系统内存不足、内存泄漏等问题。

通过这个测试能够优化字符串操作的性能，减少内存分配开销，提高系统稳定性和健壮性。



### TestTinyAlloc

TestTinyAlloc函数是runtime包中malloc_test.go文件中的一个测试函数。它用于测试内存分配器的tiny alloc（小对象内存池分配器）实现是否正确。

tiny alloc是runtime包的一个内存分配器，它专门用于分配小对象（大小小于等于16字节的对象）。tiny alloc分配器的实现方式是维护一个对象池的链表，每次分配时从链表中取出一个空闲对象，如果没有空闲对象则会从堆中申请一块内存，并将其划分为多个小对象。

TestTinyAlloc函数的测试用例主要分为以下几个步骤：

1.创建一个GC上下文。

2.设置最大分配大小为16字节。

3.从tiny alloc中分配一些对象，并将它们添加到一个slice中。

4.检查分配的对象数量是否正确。

5.将对象还回tiny alloc内存池。

6.重新分配同样数量的对象。

7.检查分配的对象数量是否仍然正确。

如果TestTinyAlloc测试函数没有报错，则说明tiny alloc的内存池分配器实现是正确的。

总之，TestTinyAlloc函数的作用是测试tiny alloc内存分配器的正确性，以确保runtime包的内存管理系统能够正确地进行小对象的分配和回收。



### TestTinyAllocIssue37262

TestTinyAllocIssue37262这个func是用来测试和解决issue #37262的。该issue描述了在某些条件下，tiny alloc会失败导致GC频繁触发的问题。这个func主要测试了以下几个方面：

1. 验证在没有问题的情况下，tiny alloc是否正常工作。
2. 模拟出tiny alloc失败的情况，并观察GC的行为。
3. 验证在fix issue后，是否能够正确处理tiny alloc的问题。

该func使用了一组测试数据，在不同的条件下运行test函数进行测试，具体包括以下几个步骤：

1. 大量使用内存并分配小对象。此时，tiny alloc应该能够正常工作。
2. 通过手动设置frees向runtime模拟tiny alloc失败的情况。此时，tiny alloc应该会失败，并触发GC。
3. 验证GC的行为，包括GC的次数、回收的内存大小等。
4. 重复第一步，以验证修复后的tiny alloc是否能够正常工作。

通过这些测试，可以验证和解决issue #37262，确保tiny alloc的稳定性和可靠性。



### TestPageCacheLeak

TestPageCacheLeak函数的主要作用是测试在运行时（runtime）中是否存在页面缓存泄漏的情况。其测试过程包括以下步骤：

1. 创建一个文件，然后读取一些数据并将其存储在页面缓存中。
2. 强制垃圾回收机制（GC），以确保所有页面缓存都得到释放。
3. 重新读取文件中的一些数据，并再次将其存储在页面缓存中。
4. 强制垃圾回收机制（GC），然后检查是否存在任何页面缓存泄漏的情况。

通过这个测试函数，可以确保runtime中的Memory Allocator能够正确地创建、管理和释放页面缓存，并避免泄漏和其他相关问题。同时，这也有助于提高程序的性能和可靠性。



### TestPhysicalMemoryUtilization

TestPhysicalMemoryUtilization函数是runtime/malloc包中的某个测试函数，它的作用是测试内存分配器在分配和释放大量内存时的物理内存利用率。该函数会在测试中分配大量的内存，并在完成测试后释放这些内存。在测试过程中，该函数会检查系统中的物理内存利用率是否与预期一致，以此验证内存分配器的内存管理机制是否能够最大化地利用系统内存资源。

在测试开始时，TestPhysicalMemoryUtilization函数会调用runtime.MemStats.Fill将系统内存状态信息存储在一个结构体中，然后分配大量的内存（默认为1GB），并在完成测试后释放这些内存，再次调用MemStats.Fill填充内存状态信息并计算物理内存利用率（由“Sys”字段表示）。然后，函数会检查实际物理内存利用率是否与预期一致，即分配的内存量和实际利用的物理内存量之间是否存在不可忽略的差异。

通过这个测试，可以验证内存分配器在分配和释放大量内存时是否正常，并且可以计算内存分配器的物理内存利用效率，确保其最大化地利用系统的内存资源。这对于保障系统的稳定性和性能非常重要。



### TestScavengedBitsCleared

TestScavengedBitsCleared函数是Go语言运行时包（runtime）中，在malloc_test.go文件中的一个测试函数。其主要作用是测试垃圾收集器是否已经正确清理并收回悬空对象（Scavenged bits）。

我们知道，Go语言中的垃圾收集器采用了标记-清除（Mark and Sweep）算法。在这种算法的具体实现过程中，当前正在运行的程序会在内存中划分出一块空间作为堆区域。程序在执行过程中会不断地从操作系统中向堆区域中请求内存空间，而随着程序执行过程的不断进行，堆区域中的内存空间就会被不断地占用和释放。

在垃圾收集器运行的过程中，需要追踪和标记出所有可达（reachable）的对象，而那些不再可达的对象则会被判定为垃圾（garbage）并被回收。但是，在进行垃圾收集的过程中，可能会出现一些对象已经被程序释放，但是垃圾收集器却没有及时清理和回收的情况。这就会导致堆区的空间被占用了很多，但实际上是无法使用的，从而影响系统的整体性能。

为了解决这个问题，Go语言的运行时会在堆区域中维护一个类似“黑白”标记的机制，即把所有的对象全部标记为白色，然后标记出所有的可达对象为黑色。在进行垃圾回收的过程中，收回所有未被标记为黑色的白色对象。

而对于Scavenged bits（悬空对象），则是指在GC开始前，已经被释放的对象上的位（bit），它们会在GC运行过程中被重用来记录对象活动性。在线程调度和内存放大时（比如mcache和mcentral的缓存），可能会把释放的对象分配给其他goroutine并使用这些位。这些位可能被错误地用于GC剩余扫描来标记堆内部的对象。

因此，TestScavengedBitsCleared函数具体的测试方案是先造一些垃圾对象，然后手动释放这些对象，在进行垃圾回收时，应该能够正确清理掉已经被释放的对象上的Scavenged Bits，并将这部分内存空间归还给操作系统。如果测试函数能够正确地检测出Scavenged Bits没有被清理的情况，那么就意味着垃圾回收器存在一些问题，需要进一步进行调试和修复。



### TestArenaCollision

TestArenaCollision是在Go的runtime包中的malloc_test.go文件中定义的一个函数，用于测试Golang程序在多个goroutine下使用内存分配器时是否会发生内存冲突。

为了支持并发编程，Golang使用了一种称为goroutine的轻量级线程。当多个goroutine同时访问同一内存时，可能会发生一些问题，如内存泄漏、内存冲突等。TestArenaCollision函数就是用于测试并发环境下内存分配器的安全性。

测试过程可简述如下：TestArenaCollision函数中启动多个goroutine，每个goroutine都会使用malloc函数申请内存，并随机写入一些数据。写入完成后，goroutine会释放该内存。但是，为了模拟内存冲突，TestArenaCollision函数中会手动修改Golang运行时的arena的头和尾指针，从而使得一个goroutine申请的内存区域与另一个goroutine申请的内存区域发生了冲突。最后，TestArenaCollision函数检查每个goroutine写入的数据是否正确，如果正确则表示内存分配器没有发生冲突，否则就会报错。

总之，TestArenaCollision函数用于测试Golang程序在多个goroutine下使用内存分配器的安全性，是Golang内存管理的一个重要测试用例。



### BenchmarkMalloc8

BenchmarkMalloc8是一个基准测试函数，旨在衡量在分配大小为8字节的对象时malloc函数的性能。在这个函数中，通过循环调用malloc分配一定数量的8字节大小的对象，然后计算分配所花费的时间。该函数的作用是测试malloc函数在特定情况下的性能，并提供用于优化分配过程的参考基准。

具体来说，该函数中使用了Go语言中的benchmarking框架，它会对函数进行多次迭代调用，并测量每次调用的执行时间。这样可以消除一些偶然性，以便得出更准确的执行时间数据。同时在测试前会为每次迭代随机生成一个种子数，以消除随机因素对测试结果的干扰。

在测试过程中，基准测试函数会设置一个分配的数量，这取决于测试环境的内存容量。基于这个设置，函数会在一个循环中调用malloc分配8字节大小的对象，并计算分配所需的时间。最后，基准测试框架会输出每次迭代的执行时间，以及平均执行时间和标准差等统计数据。这些数据可以用于比较不同传参下malloc函数的性能表现，并帮助开发人员进行优化。



### BenchmarkMalloc16

BenchmarkMalloc16是一个由Go语言的testing包提供的基准测试函数。它的作用是测试内存分配的性能，具体来说，它测试分配大小为16字节的对象的速度。

在函数内部，它首先计算出需要进行多少轮测试。然后，每轮测试中会分别进行1、2、4、8、16、32、64、128、256、512、1024、2048、4096、8192、16384个16字节对象的分配和释放，每次分配和释放都会计时，同时记录分配和释放的次数以及总共分配和释放的字节数。最后，根据这些数据，计算出每秒钟可以进行多少次分配和释放操作，以及每秒钟可以分配和释放多少字节。

通过这个测试函数，我们可以得到一些关于内存分配性能的重要信息，例如每秒钟可以进行多少次分配和释放操作，每秒钟可以分配和释放多少字节。这些信息可以帮助我们更好地了解和优化代码的内存使用情况。



### BenchmarkMallocTypeInfo8

BenchmarkMallocTypeInfo8是一个基准测试函数，用于测试分配和释放指定大小和类型的对象所需的时间。具体来说，该函数测试分配和释放大小为8字节的对象所需的时间，并且对象类型为固定的typeInfo类型。

该函数首先使用runtime.ReadMemStats函数获取系统内存使用情况，然后使用循环来执行以下步骤：

1. 记录开始时间；
2. 分配一个大小和类型为typeInfo的8字节对象；
3. 记录结束时间和分配的内存指针；
4. 释放分配的内存；
5. 重复以上步骤多次，计算分配和释放所有对象的总时间和总内存使用情况；
6. 输出测试结果。

通过这个基准测试，可以对系统的内存分配和释放性能进行评估和优化。BenchmarkMallocTypeInfo8函数是Go语言runtime库中的一个重要部分，对于实现高效的内存管理和提高程序性能至关重要。



### BenchmarkMallocTypeInfo16

BenchmarkMallocTypeInfo16是一个Go语言中的基准测试函数（benchmark function），它是用来对Go语言GC的内存分配（malloc）以及类型信息（type information）的性能进行基准测试的。该函数在runtime/malloc.go文件中定义。

具体来说，BenchmarkMallocTypeInfo16会在不同的场景下分别测试内存分配和类型信息初始化的性能，以便评估其中的性能瓶颈。其中，测试场景包括：

1. 用于测试分配大量小对象的性能，通过调用runtime.MallocTiny函数来进行内存分配；
2. 用于测试分配单个较大对象的性能，通过调用runtime.MallocGC函数来进行内存分配；
3. 用于测试新类型的类型信息初始化的性能，通过调用runtime.TypeForSizes函数来获取一个新类型的类型信息。

在基准测试过程中，BenchmarkMallocTypeInfo16会多次重复调用上述的测试场景，并记录每次操作的执行时间，最后生成一个报告，其中包括每个测试场景的平均执行时间、标准差、中位数等统计信息，以及比较不同场景下的性能差异等数据。这些数据可以用于评估GC的内存分配和类型信息性能的优化效果，以及为后续的优化工作提供参考。



### BenchmarkMallocLargeStruct

BenchmarkMallocLargeStruct是一个基准测试函数，它的主要作用是测试在分配大的struct时malloc和calloc的效率对比。

在这个函数中，首先定义了一个较大的结构体LargeStruct，然后使用testing.B对象来运行benchmark。在循环中，使用runtime.MemStats来获取内存分配信息，然后调用malloc和calloc函数分配LargeStruct所需的内存。分别测试了malloc和calloc分别调用1,2,4,8,16,32,64,128,256次的时间。

对于分配大的struct时，由于结构体内存变量非常大，因此需要分配大块的内存，这使得分配的时间非常昂贵。在此背景下使用不同的内存分配算法可以对性能有相应的影响。在这个function中的benchmark实际上是比较malloc和calloc两种分配算法哪一个对于分配这样大结构体更高效。

通常，malloc比calloc更高效，因为前者只需分配内存而不必清零，后者则需要分配并清零内存。 但是，对于大的struct，由于结构体非常大，清零所有内存可能会良好的提升性能。因此，对于大型结构体，可以使用calloc以获得更好的性能。



### BenchmarkGoroutineSelect

BenchmarkGoroutineSelect是一个性能测试函数，用于测试在多个goroutine之间使用select语句时的性能表现。在Go语言中，通过select语句可以选择不同的通道进行读写操作，可以用于在多个goroutine之间进行协作和同步。

BenchmarkGoroutineSelect函数中首先创建了一个通道channel，然后启动了多个goroutine，在这些goroutine中通过select语句从channel中读取数据，并将读取到的数据进行累加。最后，调用testing.Benchmark函数进行性能测试，记录了每轮测试的耗时和内存分配情况，用于评估select语句在goroutine协作中的性能表现。

该函数的作用是提供一个标准化的测试环境，用于测试在多个goroutine之间使用select语句时的性能和功能表现。通过这个测试函数，可以评估select语句在实际应用中的性能和可靠性，从而为程序的优化和改进提供参考。



### BenchmarkGoroutineBlocking

在Go语言中，当goroutine需要申请内存时，通常会调用runtime.mallocgc函数来进行内存分配。因为内存分配是一个相对比较耗时的操作，如果在调用该函数时出现阻塞，将会影响程序的性能。因此，runtime包中提供了BenchmarkGoroutineBlocking函数来测试goroutine阻塞时的内存分配性能。

该函数实现了一个基准测试，测试的对象是在有大量goroutine时内存分配的性能。具体来说，该测试会创建若干个goroutine，并在每个goroutine中不断的分配和释放内存，直到测试时间达到预设的时间为止。该函数会记录每次内存分配的耗时，并计算出平均每次分配所需的时间，用以评估在有大量goroutine时内存分配的性能。

通过该函数的测试结果，可以更好地了解在有大量goroutine时，内存分配的性能瓶颈在哪里，从而进行相应的优化。



### BenchmarkGoroutineForRange

BenchmarkGoroutineForRange这个函数是一个基准测试函数，用于测试同时启动多个 Goroutine 的性能。

函数中首先创建了一个管道，然后启动了1000个 Goroutine，每个 Goroutine 都会向管道发送一个整数。

接着使用 for range 语句遍历管道中的所有整数，并做简单的计算，最终用计时器记录了遍历管道所用的时间。

通过这个基准测试函数，我们可以测试当前计算机环境下同时启动多个 Goroutine 的性能，并与其他计算机或不同的运行时环境进行比较，从而提高程序性能。



### benchHelper

在 Go 语言中，`benchHelper` 是一个测试函数辅助函数，主要用于解决在进行基准测试时从缓存中读取千兆字节的数据时导致 CPU 频率急剧下降，从而影响测试结果的问题。

在 `mallock_test.go` 文件中，`benchHelper` 函数的具体实现如下：

```
func benchHelper(n int, b *testing.B) {
    if b.N < n {
        // the loop bellow is going to be very slow so refuse
        // to run the benchmark at all
        return
    }
    // fake cache heat-up
    x := make([]byte, 64<<10)
    for i := range x {
        x[i] = byte(i)
    }
    i := 0
    b.ResetTimer()
    for b.Next() {
        p := make([]byte, n)
        p[i] = byte(0xAA)
        i = (i + 1) % n
    }
    b.SetBytes(int64(n))
}
```

可以看到，在进行基准测试之前，`benchHelper` 函数首先会判断测试的循环次数 `b.N` 是否小于参数 `n`，如果是，就直接退出函数。如果不是，函数会先创建一个大小为 64 kB 的字节数组，并将其填充为递增的字节序列，以模拟热身缓存。

接下来，函数会在测试循环开始之前调用 `b.ResetTimer()` 方法，以重置计时器。然后，在测试循环中调用 `b.Next()` 方法以执行一次测试迭代。每次迭代中，函数先创建一个大小为 `n` 的字节数组 `p`，并将其第 `i` 个字节赋值为 0xAA。最后，函数使用 `(i + 1) % n` 计算出下一个要赋值的字节位置，并将其保存到变量 `i` 中。

在上述迭代完成后，`benchHelper` 函数会通过 `b.SetBytes` 方法将字节数 `n` 作为参数传递给测试对象，以保证测试结果的正确性。



### BenchmarkGoroutineIdle

BenchmarkGoroutineIdle是一个基准测试函数，用于测试在空闲goroutine中分配内存的性能。它主要包含以下几个部分：

1. 准备工作：通过参数设置和内存分配来初始化测试环境。

2. 基准测试代码：循环执行go func(){}()，并在其中进行内存分配和释放操作，记录测试耗时。

3. 结果验证：通过比较测试结果，判断内存分配性能的差异。

具体来说，BenchmarkGoroutineIdle在循环中执行了10000次go func(){}()，即启动了10000个空闲goroutine。每个goroutine中再循环执行10000次memoryAllocator函数，该函数会分配一个大小为64字节的内存块，并将其释放。在每次循环结束后，会记录当前时间，并计算出相邻两次循环的时间差。

通过对比执行分配和释放操作的时长，可以得知系统中分配内存的性能。对于不同的系统，性能可能会有所差异，但这个基准测试可以帮助开发人员了解到系统内存分配的大致情况，并且可以通过修改参数或者代码来进行性能调优。



