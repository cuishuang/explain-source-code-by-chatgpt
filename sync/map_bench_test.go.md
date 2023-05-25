# File: map_bench_test.go

map_bench_test.go 文件是 Go 语言标准库中 `sync` 包中的一个测试文件，在该文件中，实现了对于不同并发的情况下使用 `sync` 包中的 `Map` 类型所进行的性能测试。

具体来说，该文件主要用于考察 `sync.Map` 在高并发环境下的性能表现，包括读取、写入、删除等操作的性能指标。同时，该文件还结合了 Go 语言内置的 `testing` 框架，使用 benchmark 方式进行测试，可以更加全面而准确地评估 `sync.Map` 的性能表现。

在测试过程中，该文件主要涉及到多个测试函数和基准函数，其中测试函数主要用于校验 `sync.Map` 的基本功能是否正常，而基准函数则用于对 `sync.Map` 进行高并发下的性能测试。在测试过程中，测试函数会执行一些基本的操作，如 put、get 和 delete 等，并验证 `sync.Map` 的返回结果是否正确。而在基准测试过程中，会模拟多个并发读写操作，同时记录操作时间等数据，以此来评估 `sync.Map` 在不同并发条件下的性能表现。

总之，map_bench_test.go 文件充分利用了 Go 语言在性能测试方面的优势，通过全面而精确的测试方式，帮助我们更好地了解了 `sync.Map` 类型在高并发环境下的性能表现和优化空间。




---

### Structs:

### bench

在Go语言的runtime包中，map_bench_test.go文件负责测试map的性能。其中，bench结构体是用于定义map测试的基准数据的结构体。

bench结构体中包含以下字段：

- name：基准测试的名称。
- size：基准测试的map大小。
- load：map的填充因子（load factor）。
- ncpu：并发测试的goroutine数量。
- equal：比较两个map是否相等的函数。

通过定义bench结构体中的这些字段，我们可以灵活地设置不同的基准数据，测试map在不同情况下的性能表现。例如，可以设置map的大小不同、填充因子不同、并发测试的goroutine数量不同等等。

在运行map的基准测试时，我们可以通过对bench结构体实例进行修改来设置不同的测试场景，从而得到更全面的测试结果。



## Functions:

### benchMap

map_bench_test.go这个文件中的benchMap函数是一个基准测试函数，用于测试并比较不同实现方式的映射(Map)结构在插入、删除和读取操作方面的性能。

具体来说，benchMap函数会使用标准库的map和sync.Map两种不同的映射实现方式，并对它们进行以下基准测试：

1. 插入操作：对每种映射实现方式，重复执行一定次数的插入操作，并统计完成这些操作所需的时间。

2. 删除操作：对每种映射实现方式，先插入一定数量的键值对，再重复执行一定次数的删除操作，并统计完成这些操作所需的时间。

3. 读取操作：对每种映射实现方式，先插入一定数量的键值对，再重复执行一定次数的读取操作，并统计完成这些操作所需的时间。其中，读取操作分为并发和串行两种方式进行测试。

最终，benchMap函数会将不同实现方式的测试结果打印出来，并对它们进行性能比较和分析。

通过基准测试，可以更好地了解不同实现方式在不同场景下的性能表现，为我们选择合适的数据结构提供参考和依据。



### BenchmarkLoadMostlyHits

BenchmarkLoadMostlyHits函数在sync.Map类型的键值对中模拟了大部分是命中的情况下的加载性能测试。具体来说，测试函数首先创建一个长度为10000的sync.Map对象，然后填充其中的键值对。其中，80%的键值对会被填充成相同的值，也就是命中的情况。接着，函数会循环100次，每次循环都会遍历整个sync.Map对象，并将每个键对应的值加载到本地变量中。最后，测试函数会通过Go的testing包提供的Benchamark函数来对上述操作的性能进行测试。

具体来说，对于每次测试，BenchmarkLoadMostlyHits函数会记录每个遍历操作所花费的时间，并计算出平均时间和单次遍历的平均时间。这些数据可以用来评估在大部分键值对被命中的情况下，sync.Map类型的性能表现如何。如果测试结果显示性能表现良好，那么就可以说明在这种情况下，sync.Map是一种非常适合使用的数据类型。



### BenchmarkLoadMostlyMisses

该函数是对 Go 语言标准库中 sync 包中的 Map 类型进行基准测试的一个函数之一，主要测试该数据结构在大量的“未命中”（即访问不存在的键）场景下的读操作性能。

该基准测试函数大致的测试流程如下：

1. 首先生成一个指定大小的整型切片 keys，该切片中的元素是从 0 到 n-1 的整数，其中 n 是 Map 的大小；
2. 接着生成一个指定概率（通常是 1-AbsentRatio）的伪随机数切片 rands，切片长度与 keys 相同；
3. 根据 AbsentRatio 和 rands 中的数值生成一个字符串切片 absentKeys，其中 absentKeys 中的元素是由 rands 中的每个数值生成的随机字符串，用作访问 Map 时的键；
4. 创建一个 Map 实例 m，并将 keys 中的每个元素的字符串形式作为键，将其值设置为 nil；
5. 进行一次显式的 GC；
6. 开始计时，往该 Map 中顺序地插入 absentKeys 中的所有元素，并在每个元素插入后，随机地访问该 Map 中的 5 个键。这些键中，恰好有一个键在 Map 中已存在，其余都不存在（即“未命中”的情况）。上述插入和访问操作分别进行 NumIter 次。
7. 计算在所有访问操作中平均每个操作的耗时，并将该值与参考值进行比较，输出测试结果。

该基准测试函数的主要作用是用于评估 Map 类型在大量“未命中”场景下的读访问性能，对于用户来说，除了可以通过运行该测试来了解该类型性能外，还可以作为性能优化过程中的一个参考指标，帮助用户判断是否需要针对不同的使用场景采用不同的 Map 实现或者其他的数据结构。



### BenchmarkLoadOrStoreBalanced

BenchmarkLoadOrStoreBalanced是一个基准测试函数，用于测试在多个goroutine之间使用Map类型的并发性能。该函数的作用是：

1. 创建一个Map类型的变量m，并向其中添加一些键值对。

2. 启动多个goroutine，每个goroutine都会对m进行并发访问，并调用LoadOrStore方法，该方法会在Map中查找指定的key，如果存在则返回对应的value，否则会在Map中插入一个新的键值对并返回value。

3. 使用sync.WaitGroup等待所有的goroutine都执行完毕。

4. 通过性能测试评估并发访问Map的效率和吞吐量，进而验证Map类型的并发安全性和性能表现。该测试函数的名称LoadOrStoreBalanced暗示了该测试旨在测试对Map的并发访问负载分布是否平衡，即每个goroutine是否能够获得大致相等的访问机会，进而避免出现热点问题。



### BenchmarkLoadOrStoreUnique

BenchmarkLoadOrStoreUnique是一个基准测试函数，它的作用是测试sync.Map的LoadOrStoreUnique方法的性能。LoadOrStoreUnique方法是sync.Map提供的一个非标准方法，它在写入时检查key是否已经存在于map中，如果存在则返回已经存在的value，否则插入新的key-value对并返回新的value。这个方法的作用是确保只有一个goroutine可以写入同一个key。

BenchmarkLoadOrStoreUnique函数的基准测试过程如下：
1. 随机生成一组key-value对，key是int类型的0~9999之间的随机数，value是一个字符串。
2. 创建一个sync.Map实例。
3. 启动多个goroutine，并发执行以下步骤：
   a. 从key-value对中随机选择一组key-value。
   b. 调用LoadOrStoreUnique方法插入key-value对或获取已存在的value。
4. 记录并输出goroutine数和操作次数下的平均操作时间和操作次数。

这个基准测试函数的目的是测试LoadOrStoreUnique方法的性能，并且通过测试各种并发度和操作次数的情况下的性能表现，可以帮助开发者了解在不同负载下sync.Map的性能表现。



### BenchmarkLoadOrStoreCollision

BenchmarkLoadOrStoreCollision是一个基准测试函数，用于测试多个goroutine同时尝试使用LoadOrStore方法向一个共享的map中存储键值对时的性能表现。由于多个goroutine可能会同时尝试存储相同的键值对，因此存在冲突的可能性，从而需要使用mutex进行同步。该基准测试函数通过增加goroutine的数量，来测试LoadOrStore在高并发场景下的性能表现。函数执行时，先创建一个包含指定数量键值对的map对象，然后启动指定数量的goroutine，并让它们同时使用LoadOrStore向map中存储键值对。在统计测试结果时，会记录执行时间，并计算每秒钟可以处理的键值对数量（ops/s）。该基准测试函数的目的是评估并发访问共享map的性能表现，以便优化Go语言标准库中sync包中的相关方法。



### BenchmarkLoadAndDeleteBalanced

BenchmarkLoadAndDeleteBalanced函数是一个针对并发读写平衡的map的性能测试函数。该函数首先创建一个并发读写安全的map，并向其中插入1000000个键值对。然后，它使用多个goroutine并发地对map进行读取和删除操作。具体来说，这些goroutine会从map中随机选择一个键，然后同时执行读取和删除这个键的操作。这样可以测试map在高并发读写时的性能表现。

通过该函数的性能测试，我们可以评估编写并发程序时使用并发读写安全的map的效果，并找出其中的性能瓶颈和改进点。



### BenchmarkLoadAndDeleteUnique

BenchmarkLoadAndDeleteUnique这个函数是在测试并发情况下，map的Load和Delete操作的性能。它的主要作用是通过对Map中元素的删除操作进行测试，来确定Map数据结构并发读写的性能表现和瓶颈，以及在删除元素时是否会出现资源竞争和数据安全问题。

该函数的测试用例首先会初始化一个IntMap类型的Map对象，并往里面放入大量的数据元素。接着，它会模拟并发读写访问Map对象，并在并发执行的过程中，随机选择其中的若干个元素进行删除操作。

该测试用例会对删除元素操作进行计时，统计Map对象的并发操作性能，包括并发读写的吞吐量、延迟、竞争等情况。通过这些测试数据，可以分析并发操作Map对象的性能瓶颈，进而优化程序代码，提高其性能和稳定性。

总之，BenchmarkLoadAndDeleteUnique这个函数是Sync包中一个重要的性能测试工具，可以帮助开发人员更好地了解和优化Map并发读写的性能表现，确保程序性能和数据安全。



### BenchmarkLoadAndDeleteCollision

BenchmarkLoadAndDeleteCollision是一个基准测试函数，用于测试在并发负载下对map的连续加载和删除操作的性能。具体而言，它创建了一个初始大小为10000的map，并在其中添加1000000个元素，其中有一定的概率会发生哈希碰撞。然后，它启动了一定数量的goroutine，每个goroutine都会连续执行随机的加载和删除操作。测试的目的在于测量并发负载下map的性能表现。

这个函数的主要作用包括：

1. 测试map的并发性能：这个函数通过创建多个goroutine来模拟并发负载下的操作，以便测试map在高负载情况下的性能表现。

2. 测试map的哈希碰撞处理能力：这个函数在添加元素时，有一定的概率会发生哈希碰撞，以便测试map在存在碰撞情况下的性能表现。

3. 提供性能基准测试数据：这个函数测量了map在高负载条件下的性能，并提供数字化的结果，可以用作优化map相关算法和数据结构的基准。



### BenchmarkRange

BenchmarkRange函数是一个基准测试函数，它的作用是测试goroutine并发读取和遍历sync.Map类型的映射的性能。

具体来说，BenchmarkRange函数首先创建一个sync.Map类型的映射，并向其中插入一定数量的键值对。然后，它启动多个goroutine并发读取和遍历该映射，每个goroutine都会执行一定数量的读取和遍历操作。每个读取操作都会随机选择一个键，并读取其对应的值；每个遍历操作都会遍历整个映射并累加所有键值对的值。最后，函数计算总共执行了多少次读取和遍历操作，并输出执行的速度。

该测试函数的目的是评估sync.Map类型的映射的读取和遍历操作的性能，以及测试在并发读取和遍历下，该映射的表现如何。



### BenchmarkAdversarialAlloc

BenchmarkAdversarialAlloc是在sync/map_bench_test.go文件中定义的一个基准测试函数。它的作用是测试map的写入性能，同时也考察了map的内存分配和GC对性能的影响。

具体来说，该函数测试具有以下特征的map写入性能：

1. 键为int类型，值为[]byte类型，初始容量为5000。
2. 循环向map中随机写入10000个键值对，其中键和值都是随机生成的。
3. 每个写入操作都会调用runtime.GC()函数强制垃圾回收，模拟大量内存分配和释放的场景。

BenchmarkAdversarialAlloc的执行过程中，会通过多次迭代和计时，计算出实际执行时间和内存分配情况，从而得出map的写入性能。这个函数可以通过go test命令运行。



### BenchmarkAdversarialDelete

BenchmarkAdversarialDelete是一个基准测试函数，用于测试在高并发删除情况下，对map进行删除操作的性能。在该函数中，会启动多个goroutine并发地向map中添加随机的key-value对，然后再启用多个goroutine并发地从map中删除这些随机的key-value对，其中删除操作可能会发生竞争条件（例如多个goroutine同时尝试删除同一个key），从而测试map在高并发环境下的性能和稳定性。

该函数的详细流程如下：

1. 首先定义了一个包含10000个键值对的map，并启动了400个goroutine并发地向其中添加随机的键值对。
2. 然后等待一段时间，让添加操作执行完毕，等待时间为5秒。
3. 接着启动400个goroutine并发地从map中删除随机的键值对。在删除过程中，可能会发生竞争条件。
4. 使用benchmark机制统计删除操作的平均执行时间，并输出到控制台中。

通过在高并发情况下进行测试，可以更好地评估和优化map的性能，避免潜在的竞争问题。



### BenchmarkDeleteCollision

BenchmarkDeleteCollision函数是一个基准测试函数，主要用于测试在存在冲突情况下删除操作的性能表现。具体来说，该函数会生成一个指定大小的sync.Map实例，并在其中插入一定数量的键值对，使得这些键值对构成哈希冲突，然后对这个sync.Map实例进行多次删除操作，统计删除操作的平均耗时。

在具体实现中，BenchmarkDeleteCollision函数使用了Go语言自带的testing包和sync.Map类型，以及伪随机数生成器rand包。该函数首先通过命令行参数获取要测试的sync.Map实例的大小和每次测试的删除操作次数，然后使用rand包生成一定数量的键值对并插入到sync.Map实例中，以模拟存在冲突的情况。接下来，BenchmarkDeleteCollision函数重复多次进行删除操作，并使用testing包提供的计时和统计功能，最终输出测试结果，包括每次删除操作的平均耗时和总共执行的删除操作次数。

通过对BenchmarkDeleteCollision函数的运行结果进行分析，我们可以了解到在存在哈希冲突的情况下，sync.Map实例的删除操作的性能表现，进而优化代码，提高程序性能。



### BenchmarkSwapCollision

BenchmarkSwapCollision是go/src/sync/map_bench_test.go中的一个性能测试函数，用于测试map在交换引起碰撞的情况下的性能表现。

该函数的作用是创建一个map，然后向其中插入一定数量的键值对。接着，该函数会持续对这个map进行随机的key-value交换操作，以模拟交换引起的碰撞。最后，该函数会计算map在进行一定数量的交换操作后的性能表现，包括插入、查找和删除操作的耗时（ns/op）和内存分配（allocs/op）等指标。

通过这个性能测试函数，我们可以了解到map在交换引起碰撞的情况下的性能表现，以及在不同的运行条件下的性能差异，例如不同的key-value数量、不同的并发情况等。这有助于我们优化应用程序中的map使用，以提高应用程序的性能和稳定性。



### BenchmarkSwapMostlyHits

BenchmarkSwapMostlyHits函数是Go语言sync包中的map类型的性能测试函数之一。具体来说，该函数测试的是在高读取/低写入环境下，使用sync.Map类型的Map对象进行数据交换操作和数据读取操作的性能。

该函数的实现过程是：首先创建一个包含一定数量（默认为10000）随机字符串的Map对象，然后通过一个循环不断对该Map对象进行插入删除操作（默认为1000000次），其中插入操作的键和值也是随机生成的字符串，而删除操作的键则是随机选择已有的键。在每次操作前，还会随机生成一个目标键和一个目标值。在循环的每次迭代中，函数会执行以下两个操作：

1. 随机选择一个已有键，将其对应的值交换为目标值。
2. 随机选择一个已有键，读取其对应的值。

循环结束后，该函数会输出性能测试的结果：每秒钟执行的插入操作数、删除操作数、交换操作数、读取操作数和总操作数，以及每个操作的平均执行时间和内存分配大小。

该函数的目的是测试sync.Map类型的Map对象的性能，并提供性能优化的参考。在高读取/低写入的情况下，使用该类型的Map对象可以有效提高并发性能，但需要注意使用场景和并发安全问题。



### BenchmarkSwapMostlyMisses

BenchmarkSwapMostlyMisses是一个基准测试函数，用于测试在Map Swap中进行大多数未命中操作的情况下的性能表现。Map Swap是一个用于交换Map的值的函数，它通过将键的旧值替换为新值来完成操作。

在该函数中，首先创建一个包含1000个键值对的Map，然后在一个循环中，执行1000次随机生成键值对并使用Map Swap将其添加到Map中。其中，80%的情况下，这些键值对的键都是Map中不存在的，因此可以模拟大多数未命中操作的情况。

最后，该函数使用Go的testing包中的Benchmark函数来执行这个基准测试，并记录Map Swap在执行这个过程中所需的时间。通过这个基准测试函数，可以评估Map Swap在大多数未命中操作的情况下的性能表现，以及用于Map Swap的同步原语的性能如何影响其执行时间。



### BenchmarkCompareAndSwapCollision

BenchmarkCompareAndSwapCollision是一个基准测试函数，用于测试在多个goroutine同时对同一项进行比较交换操作时的性能。在该函数中，首先创建了一个包含10000个key的map，并使用sync.Map类型对其进行初始化。然后启动了20个goroutine，每个goroutine都会以一定的速度不断地尝试通过调用sync.Map的CompareAndSwap函数对一个指定的key进行比较交换操作，其中key的值一直被设定为0。

在这个过程中，由于多个goroutine同时对同一项进行操作，因此在一定概率下会出现竞争状态，即多个goroutine同时尝试修改同一个key的值，从而导致一些goroutine的比较交换操作失败。通过这个基准测试函数，可以测试在这种竞争状态下，使用sync.Map的CompareAndSwap函数的性能表现。函数的返回值是每秒钟可以执行的比较交换操作的次数。

在实际应用中，由于并发访问的问题，操作同一项数据时可能会出现竞争状态，从而导致操作失败，甚至引发死锁等问题。因此在实现并发程序时，需要考虑如何有效地避免竞争状态。sync.Map是Go语言中提供的一种并发安全的map类型，可以通过使用其提供的原子操作函数，如CompareAndSwap，来避免竞争状态，保证并发安全。BenchmarkCompareAndSwapCollision就是对sync.Map的性能进行测试，验证其在处理高并发和竞争状态下的实际可用性。



### BenchmarkCompareAndSwapNoExistingKey

BenchmarkCompareAndSwapNoExistingKey是一个基准测试函数，用于基准测试使用sync.Map的CompareAndSwap方法来尝试使用不存在的键添加一个键值对时的性能。

在该测试中，函数首先创建一个值类型为int的sync.Map。然后，在每个基准测试迭代中，将使用10个并发goroutine运行CompareAndSwap函数。每个goroutine都会同时尝试通过使用不同的键添加键值对到sync.Map中。然后，基准测试函数会通过调用b.N次该操作来测量整个操作的平均执行时间。

该测试旨在评估sync.Map的性能如何处理并发尝试同时使用不同键向map添加键值对的情况。在这种基准测试中，函数会通过尝试使用不存在的键添加键值对来模拟相同的情况。

通过进行这样的基准测试，可以确定在同时运行多个goroutine时，sync.Map的实现是否能够有效地处理并发访问，以及与标准的map实现相比是否具有更好的性能。



### BenchmarkCompareAndSwapValueNotEqual

BenchmarkCompareAndSwapValueNotEqual函数的作用是测试在使用比较并交换操作CAS（CompareAndSwap）时，如果要修改的值与旧值不匹配（值不相等），性能会受到怎样的影响。

该函数使用Go语言内置的基准测试框架（Benchmark）进行测试，其测试逻辑如下：

1. 创建一个map集合和一个随机生成的字符串数组。
2. 迭代字符串数组，将其作为key，并将一个初始值作为value，插入到map集合中。
3. 随机选择map集合中的一个key，生成一个新的随机值作为新的value。
4. 使用比较并交换操作CAS，尝试修改选定key对应的value的值。在此函数中，修改的值与旧值不匹配，即这个操作预计会失败。
5. 以此进行一定的次数的测试，并统计测试结果。

测试的目的是为了查看当修改的值与旧值不匹配时，比较并交换操作CAS的性能表现。如果每次都失败，那么该操作会导致大量的开销，直到找到一个匹配的值为止。在这种情况下，性能会相对较差。因此，可以通过这个函数测试某些情况下，使用CAS操作的效果如何，并确认这个方法在改进性能方面是否有作用。



### BenchmarkCompareAndSwapMostlyHits

BenchmarkCompareAndSwapMostlyHits是一个基准测试函数，用于测试sync.Map中的比较并交换操作的性能。在该函数中，会创建一个map，然后往里面随机加入100000个键值对。然后通过循环，对这个map进行大量的读写操作，其中大部分都是对已经存在的键进行读写操作，只有个别操作是对不存在的键进行读写操作。

该函数的主要作用有以下几个：

1. 测试sync.Map中比较并交换操作的性能：在该函数中，会对map中已经存在的键进行大量的读写操作，这些操作都会涉及到比较并交换操作。通过这些操作，可以测试sync.Map中比较并交换操作的性能。

2. 模拟真实场景中的读写操作：在实际场景中，很少有程序会对不存在的键进行读写操作，大多数操作都是对已经存在的键进行读写操作。因此，通过该函数可以模拟真实场景中的读写操作，更加真实地测试sync.Map的性能。

3. 提供性能优化的参考：通过该函数的测试结果，可以分析出sync.Map中比较并交换操作的性能瓶颈所在，并进行相应的性能优化。



### BenchmarkCompareAndSwapMostlyMisses

BenchmarkCompareAndSwapMostlyMisses函数在并发环境下对于一个由sync.Map实现的Map进行读写比较和交换操作的性能进行评测。

具体来说，该函数首先创建一个sync.Map实例并初始化一个包含100000个键值对的Map。这些键值对是通过随机生成的键和值对填充而成的。接着，该函数在一个尽可能大的GOMAXPROCS值下运行100个并发goroutines，每个goroutine都会对Map进行10000次读写比较和交换操作。这些操作是随机生成的，对于95%的情况，操作将未命中Map元素并返回false，而在5%的情况下，操作将命中Map元素并使用"比较并交换"操作进行更新或插入。

最终，BenchmarkCompareAndSwapMostlyMisses函数会测量整个测试的耗时，并将其与对于同样操作的其他Map的耗时进行比较，以评估sync.Map的性能。



### BenchmarkCompareAndDeleteCollision

BenchmarkCompareAndDeleteCollision这个函数是在测试sync.Map的性能时使用的，它的主要作用是模拟在多个goroutine中对同一个key进行CompareAndDelete操作的场景，并测试这种操作的效率和可靠性。

具体来说，函数会创建一定数量的goroutine，每个goroutine会不断地往sync.Map中存入key和value，然后在随机的时间间隔内，会对同一个key进行CompareAndDelete操作。这种操作会以一定概率发生冲突，即两个或多个goroutine同时对同一个key进行操作，此时只有一个goroutine能够成功，其他的goroutine需要重试。

通过对比测试结果，可以看出sync.Map在多个goroutine并发访问的情况下，对CompareAndDelete操作的处理效率和正确性。这对于使用sync.Map的开发者来说是非常有价值的，可以帮助他们更好地评估和优化自己应用中的并发性能。



### BenchmarkCompareAndDeleteMostlyHits

BenchmarkCompareAndDeleteMostlyHits是一个性能测试函数，用于测试在大部分情况下删除操作都会命中的情况下，sync.Map中的CompareAndDelete()方法的性能表现。

该函数会创建一个包含固定数量元素的sync.Map，然后会对其进行一定数量的读写操作，并且约95%的删除操作会命中现有的元素。

通过对比执行前后的时间差，可以评估sync.Map在这样的场景中的性能表现，并与其他实现相比较。

该测试函数可以帮助开发者了解sync.Map的在高并发访问和大量删除操作的场景下的性能表现，以及相应的优化空间。



### BenchmarkCompareAndDeleteMostlyMisses

BenchmarkCompareAndDeleteMostlyMisses是一个基准测试函数。它的目的是测试实现了sync.Map的映射表的性能和正确性，当大多数Get和Delete操作都会失败时的情况。

在这个测试中，会生成一个包含10000个键的映射表，大部分键值对将不会被存储。接着，会进行一系列的Get和Delete操作，这些操作会随机选择映射表中的键。由于大部分键值对都不存在，Get和Delete操作将会失败。

测试函数将记录下每个操作所需要的时间以及操作的准确性。准确性指的是操作是否按照期望方式执行。例如，如果某个Get操作期望返回键值对不存在，但实际上返回了值，那么这个操作就不准确。

测试的结果可以用来评估实现sync.Map的映射表的性能和正确性，以及在大多数键值对都不存在时，是否能够正确地处理Get和Delete操作。这对于数据密集型应用中的高性能并发访问非常重要。



