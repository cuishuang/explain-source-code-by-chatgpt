# File: map.go

go/src/sync中的map.go文件实现了一个并发安全的哈希表(Map)，对于多线程或者并发操作场景下的数据读写、添加、删除等操作提供保护。Map在Go语言中是一种常用的数据结构，可以通过键值对存储和获取数据，Map中的键(key)必须是可比较的类型，值(value)可以是任意类型。

map.go中主要实现了以下功能：

1.初始化Map：在初始化Map时，会根据Map的大小来分配足够的内存空间。

2.Map读写：通过Lock和Unlock来保护Map的读写操作，避免数据并发访问时出现竞争条件。对于读操作，使用一个只读的锁来保护Map，可以让多个协程同时读取数据，提高操作效率；对于写操作，使用一个写锁来保证Map的一致性，保证写操作的原子性。

3.删除键值对：使用读写锁来保护删除操作和数据访问，防止出现数据竞争的情况。

Map中使用的读写锁(sync.RWMutex)是一种线程安全的锁，使用其可以实现在一个时间点写锁和多个时间点读锁的互斥性。在写锁被获取时，读锁和写锁都不能被获取，只有在写锁释放之后，其他锁才能被获取。在读锁被获取时，其他的读锁仍然可以被获取，但写锁不能被获取。这样可以保证在同一个时间点，只有一个协程在对Map进行写操作，提高了程序的并发性和安全性。

总之，map.go在Go语言中的底层实现中起到了保护并发操作下Map数据一致性的重要作用。




---

### Var:

### expunged

在sync/map.go中，expunged是标记被删除的元素的特殊值。当元素被删除时，它的值被设置为expunged，而不是被真正删除。这是为了避免在调整散列表大小时产生的竞争状态，因为其他goroutine可能正在尝试读取这个元素。只有在重新调整散列表大小时才会真正删除标记为expunged的元素。

expunged的作用是：

1.在删除元素时，将其标记为被删除而不是立即删除，以避免竞争状态。

2.在重新调整散列表大小时，真正地删除标记为expunged的元素。

3.在调用delete方法时，与nil不同的特殊值，因为它可以区分某些元素实际上为nil的情况。

总之，expunged提供了一种删除元素的安全方式，以避免竞争条件，并且允许在重新调整散列表大小时正确定义哪些元素应该被删除。






---

### Structs:

### Map

Map结构体在Go语言的runtime中，用于定义键-值对集合的数据结构，即哈希表。它是实现Go语言中map数据类型的基础结构。

Map结构体的定义如下：

```
type hmap struct {
    count     int     // map中实际元素数量
    B         uint8   // 桶到溢出桶的比率
    hash0     uint32  // hash种子，用于哈希函数
    buckets   unsafe.Pointer  // 桶数组指针
    oldbuckets unsafe.Pointer  // 扩容前的桶数组指针
    nevacuate uintptr // 扩容标记，0表示未扩容，1表示正在扩容
    extra *mapextra  // 附加信息指针
}
```

Map结构体的成员解释如下：

1. count：map中实际元素数量，也就是map的大小。
2. B：map中哈希桶到溢出桶的最大比率，默认为6。当哈希桶到溢出桶的比率超过该值，哈希表会进行扩容操作。
3. hash0：哈希种子，用于哈希函数。
4. buckets：桶数组指针，指向哈希表的所有桶。
5. oldbuckets：扩容前的桶数组指针，当哈希表进行扩容时，新建一个更大的桶数组，将原桶数组中的元素重新哈希到新桶数组中，此时需要用到该指针。
6. nevacuate：扩容标记，0表示未扩容，1表示正在扩容。
7. extra：附加信息指针，用于存储其他额外的信息，如当前使用的哈希函数等。

Map结构体中的buckets字段指向一个桶数组，桶数组中的每一个元素都是一个链表的头节点，用于存放哈希冲突的元素。当进行查找时，先通过哈希函数计算出键值对应的桶索引，然后在对应桶的链表中查找元素。当元素数量超过一定阈值时，桶数组会进行扩容操作，将元素重新哈希到更大的桶数组中，并将扩容前的桶数组指针存储在oldbuckets字段中。在扩容过程中，如果有新元素加入，会先将该元素插入到oldbuckets中，等到扩容完成后再将该元素移动到新的桶数组中。

综上所述，Map结构体在Go语言中是实现哈希表的基础结构。它通过桶数组和链表来存储键值对，并且支持动态扩容操作。



### readOnly

`readOnly` 结构体是 Go 中内置 `map` 的一部分，它被用来存储 `map` 的只读部分。当 `map` 被声明为只读时，所有的变量都将被标记为只读状态，这意味着无法修改其中的任何值。

`readOnly` 结构体的定义如下：

```go
type readOnly struct {
    m       map[interface{}]*entry
    amended bool
    _       [64]byte // 防止不必要的分配
}
```

其中，`m` 表示 `map` 中的键值对集合，`amended` 则表示 `map` 是否被修改过，`_` 用来避免持有 `readOnly` 实例时的内存分配。

当只读 `map` 需要被修改时，会先将其转换为普通的 `map`，然后执行修改操作。这个过程称为「拷贝-on-write」，它确保了只有在需要修改时才会进行深拷贝操作，从而提高了程序的性能。

`readOnly` 在 Go 中还有其他的应用，例如可以用来构建并发安全的缓存（类似于 `sync.Map` 的实现方式）。

总之，`readOnly` 结构体是 Go 中内置 `map` 的重要组成部分，并提供了 `map` 的只读特性，从而增加了程序的安全性和性能。



### entry

entry结构体是用来表示哈希表中的一个键值对的，它的定义如下：

type entry struct {
    p unsafe.Pointer // *maptype
    // The first two fields are hash and key. Both are used to
    // check equality of keys during a map read operation.
    // 数据占位，实际数据存放在下一行
    // 注意 hash 和 key 是2者的和的异或值
    // uintptr是Go语言中一个无符号整数类型，大小和int相同
    // 可以在任何指针被存储或传递的地方使用
    // 它表示一个指针的比特位模式
    // 如果指针值被填充到一个指针大小的数据结构中，则应使用unsafe.Pointer
    // 但是uintptr大小就是一个指针的大小，可以在任何整数类型变量中保存它，
    // 因为指针大小在各个架构下都是固定的
    // hash 取自 B+树 如何维护key唯一性？
    // key 由用户传入。
    // 对于已存在的 map[key] = value 赋值操作来说，map 不会重新开辟一个新的bucket来存储value，
    // 只会用新的value替换旧的value，所以 key 不能变
    // 
    // hash: key 的哈希值
    // key: 键值
    // 当正在进行 map 读操作时，map 会比较 hash 和 key 的值是否与被查找的键值一致。
    // 如果一致，则返回相应的值；否则，将继续查找下一个元素，直到找到相应的元素或所有元素都被查找过为止。
    // 对于已存在的 map[key] = value 赋值操作来说，map 不会重新开辟一个新的bucket来存储value，
    // 只会用新的value替换旧的value，所以 key 不能变
    hash    uint32
    key     unsafe.Pointer
    // 值，跟在entry后面，也是占位
    // 值有可能是map的
    // interface{}类型中[具体的数据类型,值]，[类型信息，值]存储方式
    // 支持向接口变量中存储的数据类型为 nil
    // 如果 map 的 key 类型不是引用类型，那么这个指针就指向这个 value，
    // 否则，它将是一个指向存储在“heap”上的指针，这个指针指向 value 的底层数据结构
    value   unsafe.Pointer // *Type
}

entry包含了3个指针类型的字段：

- p：指向map类型元数据的指针。
- key：作为map的键值，根据需要转换为具体的类型。
- value：作为map的值，同样需要根据需要转换为具体的类型。

除此之外，entry还包含了一个哈希值（hash）作为键值对的唯一标识，用于在进行查找、插入、删除等操作时快速定位目标元素的位置。

entry的设计遵循了Go语言的编程哲学，即尽量减少指针的使用，尽量使用值类型，使得数据结构更加紧凑、高效。同时，使用unsafe.Pointer作为指针类型，可以避免类型转换的开销。



## Functions:

### newEntry

newEntry函数的作用是创建一个新的MapEntry对象，并将键值对添加到MapEntry中。

MapEntry表示在哈希表中存储的键值对，包含一个键和一个指向数据的指针。newEntry函数首先创建一个新的MapEntry对象，然后将键复制到MapEntry中，并将MapEntry指向数据的指针设置为空。最后，它返回新创建的MapEntry对象。

具体实现如下：

```
func newEntry(h int, key interface{}, p unsafe.Pointer) *entry {
    return &entry{p, uint64(h), key}
}
```

其中，参数h表示哈希值，key表示键，p表示指向数据的指针。函数返回一个指向新创建的entry对象的指针。



### loadReadOnly

loadReadOnly是sync包中map实现的私有方法，它的作用是在只读模式下，从map中读取指定key的value值。

具体实现过程如下：

1. 首先获取当前bucket的只读读锁，防止其他并发协程写操作导致读写冲突。

2. 然后根据传入的key值，计算key的hash值，找到对应的bucket。

3. 然后在该bucket中遍历所有链表节点，查找符合key值的节点，找到后返回其value值。

4. 最后释放读锁并返回value值。

loadReadOnly方法的主要作用是提供了对map的并发安全读取，同时在只读模式下，避免对map进行写操作，从而避免了读写冲突。 这样能够保证在多协程并发读取的情况下，能够安全地从map中读取数据，而不会出现数据错误的情况。



### Load

Load函数是sync包中的一个方法，主要用于获取Map对象中指定键值所对应的value值。具体来讲，该函数接受一个参数key，返回key对应的value值。如果key不存在于Map中，那么返回Map的value类型的零值。

Load函数的实现主要涉及了对Map对象的内部结构进行访问。Map对象会被分割成若干个桶(bucket)，每个桶中存储了一组key-value映射。当我们使用Load函数时，会先根据key计算出对应的桶号(bucket index)，然后访问对应的桶中的元素。如果桶中没有对应的key，则返回value类型的零值。

需要注意的是，Load函数只能获取Map对象中已存在的key-value映射，无法新增或删除Map中的元素。要想新增或删除Map中的元素，需要使用Map对象的其他方法，如Store、LoadOrStore、Delete等。



### load

load函数在sync/map.go文件中主要用于Load和LoadOrStore方法中，用于获取指定key的键值对。具体作用如下：

1. 获取Map指定key的键值对：load函数通过key值搜索并获取Map中对应的键值对，如果Map中不存在该key，则返回nil。

2. 支持并发读取：load函数支持多个goroutine并发读取Map中的数据，保证并发安全。

3. 原子操作：load函数对于读操作是原子操作，是线程安全的，防止读取数据时出现锁竞争等问题。

load函数主要由以下几个参数构成：

1. p：Map指向内存地址的指针。

2. bucket：表示该key对应的桶，如果该桶不存在，则需要重新生成。

3. keyHash：表示该key的哈希值，根据哈希值可以快速定位到对应的桶。

4. key：表示需要查找的key。

通过load函数，我们可以很方便地获取Map中指定key的值，实现了快速查找和并发安全读取。



### Store

在Go语言中，`sync`包提供了一些同步原语，例如`Mutex`、`WaitGroup`和`Cond`等，以帮助多个goroutine协同工作。`sync.Map`是该包中的一个特殊类型，它提供了在并发场景下安全地访问、更新和遍历map的功能。

`Store`是`sync.Map`结构体中的一个方法，它的作用是将指定的key-value对存储到map中。具体地说，`Store`的函数签名如下：

```
func (m *Map) Store(key, value interface{})
```

参数`key`和`value`分别是要存储的键和值。`Store`函数可以在并发环境中安全地使用，也就是说，它能够保证在多个goroutine同时访问`sync.Map`时不会出现竞争条件和数据竞争导致的错误或异常。

当调用`Store`方法时，如果给定的key在`sync.Map`中已经存在，那么它对应的value将会被更新成新的值。如果给定的key在`sync.Map`中不存在，那么将会创建一个新的key-value对并将其存储到map中。

值得注意的是，`sync.Map`类型不支持使用普通的索引`[]`语法来访问或更新其元素。因此，我们必须使用`Load`和`Store`等方法来实现对map的操作。

综上，`Store`方法是`sync.Map`类型提供的一个关键方法，它能够在并发环境下安全地存储数据，并保证在多个goroutine之间不会出现竞态条件和数据竞争的问题。



### tryCompareAndSwap

tryCompareAndSwap是sync包中实现的一个函数，用于原子地比较并交换某个变量的值。具体作用是：

在map.go中，使用tryCompareAndSwap函数来实现了对map数据结构的并发安全的访问和操作。在Map的代码中，使用了一个叫bucket的结构来实现对map元素的存储和访问，bucket中有一个bmap引用，指向一个bmap类型的结构体，用于存储key/value对。在多个goroutine中并发访问同一个map时，需要确保对bucket以及bmap结构体的并发访问的安全，tryCompareAndSwap函数就是用来保证这一点的。

当多个goroutine同时访问同一个map时，可能会同时对同一个bucket或者同一个bmap结构体进行读写操作，为了保证并发的安全性，需要使用tryCompareAndSwap函数进行原子操作，保证在一次操作中只有一个goroutine能够成功访问这个bucket或者bmap结构体，其他goroutine则需要等待直到前一个操作完成。

其具体实现是通过使用Go语言内置的原子操作函数来实现的，具体操作是：先将原始值原子读取出来，进行某些操作后，再将新值进行原子地比较并交换。而在tryCompareAndSwap函数中，如果原始值与预期值相同，就进行替换新值操作，并返回true表示成功；如果原始值与预期值不同，则返回false，代表操作失败。

这样的实现方式可以保证map数据结构的高效性和安全性，保证每个goroutine都可以正确地访问和操作map中的元素。



### unexpungeLocked

unexpungeLocked函数的主要作用是将一个已删除的元素从map中恢复，并将它重新插入到map中。

当删除一个键值对时，sync.Map不会立即从底层map中删除这个键值对。相反，它会将这个键值对的状态标记为被删除，即将entry的p字段设置为nil。这可以避免使用锁定来删除元素。但是，如果未删除的元素填满了map的桶，并且后续元素需要添加到这个桶中，则需要将已删除元素从桶中移除，以便为新元素腾出空间。

在这种情况下，unexpungeLocked函数将在未填满的键值对中查找删除状态的元素，并将其恢复为未删除状态。然后，它可以将它添加到一个新的、空的桶中，并将桶的数量增加1。

unexpungeLocked函数也可以用于恢复在负载很低的情况下删除的元素，这些元素可能会导致map不均衡。当有删除的元素时，如果map中元素的数量不足1/8的大小，则将map中未标记为删除的元素迁移到新的桶中。

总之，unexpungeLocked函数可以帮助恢复已删除的元素，并维护map的平衡性。



### swapLocked

在sync/map.go文件中，swapLocked函数用于交换两个映射项的键值对。该函数被其他函数调用，例如Store和LoadAndDelete。

具体来说，该函数接受两个参数：旧键值对（old）和新键值对（new）。该函数首先检查旧键值对是否存在，如果存在，则将其从映射中删除。随后，新键值对将被插入到映射中。最后，如果存在旧键值对，该函数还会返回该项的值。

通常，swapLocked函数用于实现基本的同步原语，例如互斥锁和读写锁。这为多个goroutine访问相同映射时提供了线程安全性。 整个同步过程在该函数内部完成，因此其他函数可以直接调用swapLocked函数，以方便地使用映射安全操作。



### LoadOrStore

LoadOrStore是一个在Map中查找给定键的值，如果键存在，则返回该键的值，否则将给定值存储在该键下并返回该值。它的函数签名如下：

```go
func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
```

其中，key是要查找或存储的键，value是要存储的值。如果键存在，则返回该键的值和true，表明该值已经存在；否则，存储给定值并返回该值和false。

下面是一个示例，说明如何使用LoadOrStore函数：

```go
package main

import (
    "sync"
    "fmt"
)

func main() {
    var m sync.Map
    v1, loaded := m.LoadOrStore("foo", "bar")
    fmt.Println(v1, loaded) // bar false
    v2, loaded := m.LoadOrStore("foo", "baz")
    fmt.Println(v2, loaded) // bar true
}
```

上面的示例中，首先使用LoadOrStore函数将"bar"存储在键"foo"下。在第二次调用LoadOrStore函数时，它尝试将"baz"存储在同一键下，但是由于该键已经存在，LoadOrStore函数返回先前存储的值"bar"和true，表明该键已经存在。

LoadOrStore函数是一种原子操作，它可以安全地在多个goroutine之间使用，可以用于实现一些常见的并发模式，例如单实例模式、结果缓存等。



### tryLoadOrStore

tryLoadOrStore函数是sync.Map中的一个非导出函数，用于在map中查找一个key的值，如果找到则返回值，否则在map中存储一个键值对并返回输入的value。

具体流程如下：

1. 使用readLoad操作读取map中key的值，如果读取到了，就直接返回该值。

2. 如果readLoad操作读取不到key的值，就要使用Store操作往map中添加一个键值对。这里需要用到cas操作，避免竞态条件的发生。如果cas操作成功，说明其他协程没有先于当前协程往map里添加该键值对，此时可以返回输入的value。

3. 如果cas操作失败，说明其他协程已经添加了该键值对，那么就需要重新进行读取操作，直到返回值为止。

tryLoadOrStore函数主要用于并发安全地往map里存储键值对，同时也能保证多个协程同时读取map的时候，能够正确地获取到需要的值。



### LoadAndDelete

LoadAndDelete函数是sync包中的一个方法，它的作用是异步获取并删除指定键的值。函数签名如下：

```go
func (m *Map) LoadAndDelete(key interface{}) (value interface{}, loaded bool)
```

参数说明：

- m *Map: 要操作的 Map 对象指针。
- key interface{}: 要删除的键。

函数返回两个值：

- value interface{}: 如果键存在，value表示该键对应的值；如果键不存在，value为nil。
- loaded bool: 表示该键是否存在，true表示存在，false表示不存在。

该方法可以同时完成获取和删除操作，避免了多次调用 Load 和 Delete 方法产生的竞态条件。调用该方法后，如果该键存在，则会返回对应的值并将其从 Map 中删除；如果该键不存在，则只返回nil。因此，调用该方法的结果既能完成获取，又能完成删除的操作，是一种比较高效的方式。



### Delete

Delete这个func用于删除map中的指定key以及对应的值。其定义如下：

```go
func (m *Map) Delete(key interface{})
```

该函数接受一个参数key，表示要删除的键。

当Delete被调用时，它首先通过读写锁来保持对map的互斥访问。然后，它会从map中删除指定key以及对应的值。最后，它释放锁。如果删除的key不存在，则Delete函数不会执行任何操作。

该函数在多线程环境下是安全的，因为它在操作之前获取了互斥锁，并在操作完成后释放了锁。这确保了在同一时间只有一个线程可以访问Map。

总之，Delete函数是map类型的一个内置方法，用于从映射中删除一个指定的键及其相关的值。它带有一个参数，即表示要删除的键。如果找不到该键，Delete函数不执行任何操作。



### delete

delete函数是用来删除map中指定key-value对的函数。

delete函数需要传入两个参数：第一个参数是一个map，第二个参数是要删除的key。

如果map中存在这个key，则delete函数会删除对应的key-value对，并返回true；如果map中不存在这个key，则delete函数不进行任何操作，直接返回false。

delete函数的执行是原子操作，即在执行期间，不会有其他的goroutine会访问这个map。但是，delete函数不会强制同步map中的元素，因此对map的修改可能不会立即反映到其他goroutine中。

需要注意的是，当从一个nil的map中删除元素时，会产生一个panic异常。因此，在删除之前要确保map不是nil。



### trySwap

在sync/map.go中，trySwap()函数用于原子地检查键值对是否存在于地图之中，如果存在则返回它的原始值并更新为新值。 这个函数是为了支持Concurrent Map操作而设计的。

这个函数具有以下特性：

- 它会检查该键的存在，并在存在时返回原始值，否则返回nil。
- 它将使用内置的原子操作来防止并发访问。在存在时，它会尝试原子交换键值对的值。
- 将新值存储到其他变量，并返回旧值以启用进一步处理。

trySwap()函数的作用是确保Concurrent Map操作是安全和一致的。它允许多个goroutine同时并发地读写映射，而不必担心数据的一致性和并发问题。



### Swap

在Go语言中，map是一种常见的数据结构。在多协程并发编程中，当多个协程需要同时对map进行操作时，就可能导致map的数据出现不一致的情况。为了避免这种情况的发生，Go语言中提供了sync包，提供了一些线程安全的map操作方法，其中就包含了Swap这个方法。

Swap函数用于原子地交换map中给定键的值。在多协程并发编程中，使用Swap方法可以避免数据一致性问题。如果在两个协程中同时修改同一个键的值，那么只有一个协程能够成功修改，另一个协程则会返回修改前的值。这样可以保证map中每个键的值都是一致的。

Swap方法的函数签名如下：

```
func (m *Map) Swap(key interface{}, value interface{}) interface{}
```

它接受两个参数，第一个参数是键，第二个参数是值。它首先检查键值对是否存在，如果存在则将旧值返回，否则返回nil。然后，将键对应的值替换为新的值，并返回旧的值。在交换值的过程中，Swap方法使用了互斥锁来确保原子性。



### CompareAndSwap

CompareAndSwap是sync包中map.go文件中的一个方法，它是用来将某个键值对的值进行替换的。

具体来说，CompareAndSwap方法会接受三个参数：键、期望值以及新值。它会检查给定键的值是否等于期望值，如果相等，则将改键的值替换为新值，并返回true；否则不进行替换，返回false。

这个方法的作用通常是在多个Goroutine共享同一个Map时，避免出现竞争条件。由于Map是一种非线程安全的数据结构，因此在多个线程同时操作一个Map时，可能会导致数据错乱、数据丢失等问题。通过使用CompareAndSwap方法将读写操作串行化，可以保证数据的正确性。

总之，CompareAndSwap是Map类型中实现原子读写的一个关键方法，通过它的使用，可以避免出现竞争条件从而保证程序的正确性。



### CompareAndDelete

`CompareAndDelete`函数是`sync.Map`类型中提供的一个用于删除指定键值对的方法。这个函数根据给定的键和值，比较它们是否与当前`sync.Map`中对应的键和值都匹配。如果键和值都匹配，则删除该键值对，并返回`true`，否则不删除并返回`false`。该函数实现中使用了`unsafe.Pointer`和`atomic`包，以确保并发安全性。

具体来说，`CompareAndDelete`函数的主要作用如下：

1. 根据给定的键和值构造一个`interface{}`类型的对象，并使用`unsafe.Pointer`将其转换为指向`sync.Map.internal`字段的指针。

2. 在`sync.Map.internal`字段中找到键所对应的桶，遍历该桶中的所有节点，比较节点的键和值是否与给定的键和值匹配。

3. 如果找到匹配的节点，使用`atomic.CompareAndSwapPointer`原子操作将该节点的`next`字段指向的节点设置为新的节点，然后将该节点从桶中移除，并返回`true`。如果没有找到匹配的节点，什么也不做，并返回`false`。

总之，`CompareAndDelete`函数提供了一种并发安全的方法来删除`sync.Map`中的键值对。它可以在多个goroutine同时访问`sync.Map`时使用，以确保不会出现并发访问和竞争条件的问题。



### Range

Range这个func是用来遍历和访问一个map中所有的键值对的。

具体来说，Range是一个map类型的方法，调用该方法可以针对该map中的每一个键值对，都执行一个回调函数。

Range函数的函数签名如下所示：

func (m Map) Range(f func(key Key, value Value) bool)

其中，Map类型实际上是一个sync.Map类型，Key和Value都是interface{}类型的变量，f是一个回调函数，用于处理每一个键值对。

在调用Range函数时，需要传入一个回调函数f，在这个回调函数中，我们可以读取、操作map中的键和值，或者做一些其他我们需要做的事情。

Range函数的一个重要特点是，它是并发安全的。也就是说，即使有多个goroutine在同时访问同一个map，也不会引起线程安全问题。

同时，Range函数还有一个返回值，它是一个bool类型的值。如果该返回值为false，Range函数会停止遍历map，否则，Range函数会继续遍历map中的键值对。

总之，Range函数是一个非常有用的函数，可以帮助我们在并发程序中遍历map并进行一些操作。同时，它还能保证并发安全和遍历按顺序进行。



### missLocked

missLocked这个函数是在map中进行读取操作时，如果没有找到对应的键值对，则调用的函数。该函数的作用是在读取操作的过程中，如果遇到了尚未被初始化的桶，则会创建一个新的桶，并将其插入到正在使用的桶链表中。如果找到了对应的桶，则将其加锁，以保证线程安全。

具体来说，missLocked函数首先会检查map是否达到最大负载因子，如果达到则会进行重新哈希操作。接着，它会遍历map中所有的桶，查找是否存在对应的键值对。如果没有找到，则会将新的键值对插入到正在使用的桶链表中，并将桶加锁。最后返回新插入的键值对。

这个函数的作用是确保在读取操作中，即使map中不存在对应的键值对，也能够安全地进行插入操作，并保证线程安全。



### dirtyLocked

sync/map.go中的dirtyLocked函数用于将map的deferredDeletes，dirty和pendingWrite字段设置为nil，表示map之前的所有修改都已经被应用到了底层的数据结构中。这个函数会被在一些条件下调用，例如当map的锁定状态被改变、gc标记时、以及map存储空间动态增长时。

具体来说，dirtyLocked函数的作用如下：

1. 如果deferredDeletes不为空，它会遍历deferredDeletes中的每个元素，并将其删除。deferredDeletes是一个记录待删除元素的缓存，在map的主体中，如果一个元素被标记为删除，那么它会被添加到deferredDeletes中暂时缓存起来，以保证删除操作不会导致迭代器的失效。

2. 如果dirty不为空，它对底层的数据结构进行更新。dirty是记录所有要写回到底层数据结构的修改的缓存，当底层数据结构需要重新构建时，会将所有的修改写回到底层数据结构中。dirtyLocked将dirty中的每个元素写回到底层数据结构中，清空dirty，表示之前的修改已经被应用到底层。

3. 如果pendingWrite不为空，也将其清空。pendingWrite是一个记录刚被插入但还未暴露给用户的新元素的缓存，以保证它们不会被之前的操作修改，直到新元素被读取或者删除。pendingWrite的清空表示之前的操作已经在pendingWrite中记录了新增元素的操作，它们已经被应用到底层的数据结构中，新元素已经可以被安全地操作了。

由于map的底层结构是一个哈希表，需要在并发环境下保证同步，dirtyLocked函数的正确调用是保证map数据结构一致性和线程安全性的关键所在。



### tryExpungeLocked

tryExpungeLocked是一个内部函数，其目的是尝试从当前Map中移除被标记为删除的Entry，以减少Map中Entry的数量。该函数实现了ExpungeLocked函数中的核心逻辑。

具体来说，tryExpungeLocked的作用如下：

1. 尝试获取锁，如果锁已经被其他goroutine持有，则返回false。
2. 遍历当前Map，并检查每个Entry是否被标记为删除。
3. 如果某个Entry被标记为删除，则将其从当前Map中移除，并将删除标记清除。
4. 这样，被删除的Entry所占用的内存就可以被垃圾回收机制回收。

总的来说，tryExpungeLocked函数的作用是将Map中被标记为删除的Entry从Map中移除，并释放其占用的内存，以便更好地管理Map中的Entry。



