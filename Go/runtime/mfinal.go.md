# File: mfinal.go

mfinal.go 文件是 Go 运行时的一部分，主要功能是在 GC 扫描完毕之后清理和释放未使用的内存，并处理一些与内存相关的细节。

具体地说，mfinal.go 中的重点是 finalizer 的处理。finalizer 是一个可以在 GC 回收一个对象之前执行的函数，用于释放对象所占用的系统资源和执行其他清理操作，以便程序可以更好地管理内存和其他系统资源。Go 语言的 finalizer 机制是基于垃圾回收器的引用计数实现的，而 mfinal.go 在垃圾回收器中起到了关键的作用，它负责管理 finalizer 相关的数据结构。

mfinal.go 中的代码需要同时处理两个相关的数据结构：对于每个对象，都需要维护一个包含所有 finalizer 的链表，同时每个 finalizer 都需要维护一个包含所有关联对象的链表。因此，mfinal.go 中包含了以下主要函数：

1. addfinalizer：这个函数用于将一个 finalizer 添加到某个对象的 finalizer 链表中。添加一个 finalizer 需要为对象分配一些额外的内存来存储 finalizer 相关的信息。同时，为了避免出现循环引用导致的内存泄漏，addfinalizer 还需要检查 finalizer 中是否引用了对象自身，以及对象是否引用了 finalizer 中的其他对象。

2. runfinq：当 GC 扫描完所有对象之后，就会调用这个函数来执行所有 finalizer。runfinq 会遍历所有对象的 finalizer 链表，并依次执行每个 finalizer。在执行 finalizer 之前，需要将对象的 finalizer 链表清空，以避免在 finalizer 中再次添加 finalizer 导致链表无限增长。

3. runAllFinaizers：这个函数用于在程序关闭时释放所有对象和 finalizer 占用的系统资源。它的原理是将所有 finalizer 按照引用关系排序，并调用它们进行清理操作。需要注意的是，如果某个 finalizer 无法执行或者执行出错，整个清理过程会终止，可能导致一些资源无法被释放。

总之，mfinal.go 的作用是为 Go 语言提供了一套强大的 finalizer 机制，它使得程序能够更好地管理内存和其他系统资源，并且保证了 finalizer 的执行顺序和正确性。




---

### Var:

### fingStatus

在go语言的runtime包中，mfinal.go文件中有一个名为fingStatus的变量。该变量用于检测并控制堆对象的最终化状态，即垃圾回收器在回收之前需要对该对象进行的操作。

具体来说，当一个对象被创建并分配在堆上时，它可能需要在稍后被垃圾回收器进行最终化处理，比如执行析构函数或清除某些资源。这些最终化操作可能需要一定的时间和资源，并且可能会影响垃圾回收器的性能，因此需要进行控制。

fingStatus变量就是用于控制对象的最终化状态的。它包含三种状态：fingNotFinalized、fingBeingFinalized和fingFinalized。fingNotFinalized表示该对象还没有被标记为需要最终化操作；fingBeingFinalized表示该对象正在进行最终化操作；fingFinalized表示该对象已经完成最终化操作。

当一个对象需要最终化操作时，垃圾回收器会将其状态设置为fingBeingFinalized，并调用相应的最终化函数；最终化函数执行完毕后，将其状态设置为fingFinalized。在垃圾回收器中，通过检查fingStatus变量的值，可以判断对象是否需要进行最终化操作，避免了重复执行最终化操作的问题。

综上所述，fingStatus变量在go语言运行时中起到了控制对象最终化状态的作用，是垃圾回收和内存管理的关键之一。



### finlock

在go/src/runtime中mfinal.go这个文件中，finlock变量表示是否允许在某个goroutine执行finalizer。

在Go语言中，finalizer是一种用于在垃圾回收时执行特定操作的机制。当一个对象在垃圾回收时被删除时，如果它有finalizer，则会在从内存中删除对象之前执行特定的处理操作。而这些特定的操作可能会占用一定的时间和资源，可能会引起不可预期的结果。

为了避免这种不可预期的结果，Go语言在执行finalizer时默认是会将锁定当前goroutine，即其它goroutine无法执行finalizer。而finlock变量则是用于控制是否允许在某个goroutine执行finalizer。

当finlock为0时，表示当前goroutine可以执行finalizer。当finlock为1时，表示当前goroutine不可执行finalizer。

因此，在实现finalizer时，需要在唯一允许执行finalizer的goroutine中加锁、解锁。同时，我们也可以通过调整finlock的值来控制goroutine对finalizer的执行。



### fing

变量fing在mfinal.go文件中的作用是表示需要扫描的对象的索引。具体来说，fing是一个指向下一次GC扫描时需要开始扫描的对象的指针。

在Go语言中，通过标记-清除算法垃圾回收时会先标记所有可达对象，然后清除未被标记的对象。标记是从根节点开始，递归遍历对象图，标记所有可达对象。在递归遍历时，需要记录已经标记的对象，以防止重复标记。因此，Go语言使用了三色标记算法，将所有对象分为白色、黑色和灰色三种状态，在遍历对象时，将待遍历的对象放入灰色队列，标记完成后，将所有黑色对象清除。

在mfinal.go文件中，针对灰色对象的扫描需要从上一次GC结束的位置开始，并扫描到下一次GC开始时灰色对象队列的末尾。因此需要使用变量fing记录灰色对象队列的起始位置。在扫描过程中，每扫描一个对象，都会将fing指向下一个待扫描的对象，直到扫描完成。

总之，变量fing在mfinal.go文件中的作用是帮助标记-清除算法垃圾回收器定位下一次需要扫描的对象。



### finq

在Go的运行时环境中，mfinal.go文件中的finq变量是一个指向待清理的finalizer对象队列的指针。finalizer对象是一种特殊的对象，其在被垃圾回收器回收之前被添加到finq队列中，当它被回收时，其关联的finalizer函数将被调用。

当一个对象在不再被引用时，垃圾回收器将对该对象进行垃圾回收，但它可能仍然需要执行一些附加的清理工作，例如释放一些资源或关闭文件等。这样的情况下，开发人员可以使用finalizer来指定在对象被回收时要执行的这些额外的清理工作。

finq变量的作用是保存所有等待被清理的对象的finalizer函数和对象的指针，它充当了一个中间缓冲区，以便在垃圾回收器的清理过程中为这些对象的finalizer函数提供一个单独的执行线程。



### finc

在go语言的运行时(runtime)中，mfinal.go文件中的finc变量用于存储当前goroutine（协程）执行完毕后需要执行的函数。具体来说，当一个goroutine执行完毕，会调用goexit函数，这个函数会检查当前goroutine是否有需要执行的函数（通常是一些清理工作），如果有，就会调用这些函数，执行完毕后返回到主线程。

finc是一个func()类型的函数，类型定义在sys_call_windows.go等系统调用文件中。在goroutine执行时，可以通过进一步的调用来将需要执行的函数注册到finc中，这些函数将会在goroutine执行完毕时被执行。举个例子，在需要进行一些资源清理的时候，可以将清理函数注册到finc中，这样在goroutine执行完毕后会自动释放资源并执行清理函数，避免资源泄漏等问题。

总之，finc变量的作用是为goroutine提供一个嵌入式的清理机制，保证goroutine执行完毕后可以进行相应的资源清理工作，保证代码的安全与可靠性。



### finptrmask

finptrmask是一个指针的bit mask，它用于标识一个对象是否包含指向堆上分配对象的指针。它的作用是帮助垃圾收集器识别堆上对象的引用。

在GC期间，GC器遍历栈和全局变量寻找每个指向堆上对象的指针，它会标记这些指针所指向的对象为活动对象。但是，某些对象可能不包含指向堆上对象的指针，例如纯粹的数值对象或字符串，对于这些对象，GC将会跳过它们，不进行扫描。

通过使用指针bitmap，GC先查找对象的指针bitmap，而不是直接扫描所有对象。如果对象的指针bitmap中包含一个bit，则GC将扫描该对象，并在对堆进行扫描时处理它。如果bit未设置，则表明对象不包含指向堆上的指针，因此GC跳过该对象。

因此，finptrmask的作用在于为GC提供有效的筛选机制，以提高GC执行效率。



### allfin

在go/src/runtime中mfinal.go文件中，allfin这个变量是一个链表，用于记录所有的finalizer函数。finalizer函数是在一个对象被垃圾回收器释放时被调用的函数（这个函数需要传入这个对象的指针作为参数），它通常用于执行某些资源清理动作，比如关闭一个文件，释放一个锁等。

这个链表的作用是对所有的finalizer函数进行管理，包括添加、删除和执行。当一个对象被标记为垃圾时，垃圾回收器会执行它对应的finalizer函数。同时，当一个finalizer函数被添加到allfin链表时，它会被分配到一个finalizer队列中等待执行。这个队列的作用是避免频繁地调用finalizer函数，它会将需要执行的函数放到一个后台goroutine中异步执行，以提高性能。

需要注意的是，因为finalizer函数是在垃圾回收过程中异步执行的，所以不能保证它们会按照添加的顺序执行，也不能保证它们会按照预期执行。因此，在使用finalizer函数时需要格外小心，最好只在极少数情况下使用它们。



### finalizer1

在Go语言中，finalizer1是一个用于管理垃圾回收期间处理任意类型对象的finalizers的结构体。它的作用是在垃圾回收时，调用已经注册的finalizers来执行某些清理操作，比如关闭打开的文件、释放动态分配的内存等等。

finalizer1结构体包含了一个待清理的对象的地址和一个用于执行清理操作的函数指针。这个函数指针指向一个用户定义的清理函数，在垃圾回收时被调用。

finalizer1结构体也维护了一个链表，用于管理已经注册的finalizers。具体来说，当用户调用了runtime.SetFinalizer函数来注册一个finalizer时，finalizer1就会创建一个新的finalizer1结构体，并将它添加到链表中。当一个对象需要在垃圾回收时执行清理操作时，垃圾回收器会遍历finalizer链表，并依次执行所有已注册的finalizers。

除了finalizer1结构体之外，mfinal.go文件还包含了一些与垃圾回收相关的其他结构体和函数。这些结构体和函数一起组成了垃圾回收器的实现。






---

### Structs:

### finblock

在Go语言中，每个Goroutine都有一个对应的M（Machine的意思），它负责运行该Goroutine。当Goroutine结束时，M也需要关闭它自己的所有资源。finblock结构体就是用于这个目的的。

一个finblock表示一个在M上等待关闭的资源块，它包括两个重要的字段：waiters和sema。waiters存储了所有等待关闭该资源块的Goroutines，sema是一个计数信号量，表示当前有多少个Goroutine等待关闭该资源块。

当一个M需要关闭某个资源块时，它会调用finblock的Add方法将自己加入waiters列表，并利用sema信号量阻塞等待所有等待者都到齐。当所有等待者都到齐时，M会处理该资源块的关闭，将sema信号量释放，所有等待者都会得到通知继续执行。

总之，finblock的作用就是管理M上的资源关闭。它可以确保所有的等待者在所有其他等待者都已经到齐之后才开始关闭资源，既保证了并发性，又保证了资源安全。



### finalizer

在Go语言中，finalizer是一种机制，用于在对象被垃圾回收时执行一些清理操作。finalizer结构体是用于表示一个对象的finalizer函数的结构体。

在mfinal.go文件中，finalizer结构体包含以下字段：

- fn：一个函数指针，指向对象的finalizer函数。
- arg：一个空接口，表示finalizer函数的参数。
- nret：一个int值，表示finalizer函数的返回值的数量。
- ptrType：一个类型指针，表示对象的类型。

finalizer结构体的作用是将finalizer函数与对象关联起来。当对象被垃圾回收时，GC会检查该对象是否有finalizer函数。如果有，就会在垃圾回收时执行该函数。这个机制可以用于一些需要进行资源回收或清理的场景，比如关闭文件、释放锁等等。

finalizer结构体还有一个重要的作用是防止内存泄漏。如果一个对象的finalizer函数没有被执行，那么这个对象所占用的内存就不会被释放，从而导致内存泄漏。因此，在使用finalizer函数时，需要确保该函数执行的可靠性和效率。同时，需要注意不要过度依赖finalizer函数，应该尽量避免使用finalizer函数来进行资源管理。



## Functions:

### lockRankMayQueueFinalizer

lockRankMayQueueFinalizer函数的作用是加锁，并将一个对象的finalizer函数加入到finalizequeue队列中。finalizequeue队列中的函数会在垃圾回收时被执行，清理对象。

具体来说，当一个对象的引用计数为零时，意味着它已经不再被任何其他对象所引用，这时就可以通过垃圾回收机制来回收它的内存空间。但是，在回收对象之前，可以先执行与之关联的finalize函数，清理一些必要的资源，防止内存泄露。

为了能够正确执行finalize函数，需要在对象的引用计数为零时立即将其加入到finalizequeue队列中，并在垃圾回收时执行队列中的函数。由于finalize函数的执行可能存在对象的依赖关系，为了保证正确性，需要加锁排队执行。

lockRankMayQueueFinalizer函数通过加锁并将对象添加到finalizequeue队列中，实现了finalize函数的管理。



### queuefinalizer

queuefinalizer函数的作用是将待回收的对象和其关联的finalizer函数添加到一个队列中，等待垃圾回收器进行回收。

具体来说，当程序中的某个对象不再需要时，我们可以通过runtime.SetFinalizer函数为其设置一个finalizer函数。这个finalizer函数会在垃圾回收器回收对象时被调用。但是，由于GO语言的垃圾收集器是基于标记-清除算法实现的，在回收某个对象时，如果这个对象有关联的finalizer函数，则这个对象不能被直接回收，而需要将其变为一个特殊的对象（也就是finalizer对象），并将该对象添加到一个finalizer队列中。等待下一个垃圾回收场景来处理。

queuefinalizer函数就是用来管理这个finalizer队列的。它会将finalizer对象添加到队列中，同时会将其与相应的待回收对象和finalizer函数相关联。当垃圾回收器扫描到这个finalizer对象时，它会调用相应的finalizer函数，完成对象的清理操作。

需要注意的是，由于finalizer函数是在垃圾回收器线程中执行的，因此不能对finalizer函数执行繁重的操作，否则将影响垃圾回收的效率。因此，我们应该尽量避免使用finalizer函数，或者将其设置为异步的方式执行。



### iterate_finq

iterate_finq是一个函数类型为func(f *funcQueueItem) bool，用于迭代funcQueue中的每个函数节点。具体来说，它会遍历funcQueue，并对其中的每个funcQueueItem调用传入的函数f，如果f返回false则停止迭代，否则继续遍历。

在runtime中，funcQueue是用于保存待执行的defer函数的队列。iterate_finq的作用是在某些特殊情况下处理队列中未执行的defer函数，如在需要废弃M（一种可执行goroutine的线程）时，需要先执行其上的defer函数以保证资源的正确释放等。

总之，iterate_finq函数是用于迭代funcQueue中每个节点并执行一些特殊操作的工具函数。



### wakefing

mfinal.go文件定义了在程序退出时执行的一些操作，其中包括实现所有goroutine的停止和清理。

wakefing函数是其中一个函数，它的主要作用是通知所有不处于阻塞（例如在sleep或select等待）的goroutine，它们应该退出或重新执行操作。函数中的操作包括：

1. 设置全局状态，表明goroutine需要退出；
2. 将所有运行状态的goroutine加入到唤醒列表中；
3. 唤醒正在等待p的goroutine；
4. 当正在运行的goroutine使用完当前的p后，唤醒下一个等待的goroutine；

这些操作确保了所有在运行或等待中的goroutine都能够被及时通知到程序退出，从而正确地停止和清理所有资源。



### createfing

createfing函数的作用是创建一个新的fing结构体，并返回该结构体的指针。该结构体用于管理goroutine的调度，包括分配M（machine）和P（processor）等资源，以及调度goroutine的执行等。

具体地说，createfing函数首先从空闲list（即idlefing）中获取一个空闲的fing结构体。如果没有空闲的结构体，则会调用newfing函数创建一个新的结构体。接着，函数会初始化该结构体的各个字段，包括调度器（sched）和相关锁（lock、glock），以及与M、P等资源相关的状态等。最后，函数返回该结构体的指针，使得该结构体可以被进一步使用和管理。

总的来说，createfing函数是runtime调度器中的一个重要组成部分，用于创建和管理goroutine的调度和执行。它通过管理fing结构体，实现对M、P等资源的分配和调度，保证线程安全和高效的运行。



### finalizercommit

`finalizercommit`是Go语言运行时系统中的一个函数，其主要用途是将被标记为待销毁的对象从待终止对象列表中删除，并执行与该对象关联的finalizer函数。

在Go语言中，`finalizer`是一种能够定义在对象上的函数，在该对象被垃圾回收时自动执行。这个特性可用于确保资源被正确释放，例如关闭文件、释放网络连接等。

在执行完对象的Finalizer函数后，需要将其从待终止对象列表中删除，以免继续占用内存资源。这就是`finalizercommit`函数的作用。具体来说，`finalizercommit`会将待销毁对象从待终止对象列表中删除，并将其中的对象添加到下一个周期的终止列表中，以便在下一个垃圾回收周期中继续进行垃圾回收。

总之，`finalizercommit`函数在Go语言运行时系统中具有非常重要的作用，它能够确保被标记为待销毁的对象在执行完Finalizer函数后能够被及时删除，从而保证整个系统的效率与可靠性。



### runfinq

mfinal.go文件中的runfinq函数主要用于执行待清理的对象的finalizer函数。

在Go语言中，垃圾回收器会自动回收不再被引用的对象。但是有些对象在被回收之前需要执行一些额外的操作，比如释放资源、关闭文件等等。这些额外的操作可以通过定义finalizer函数来实现。在引用计数为0时，垃圾回收器会将待清理对象添加到待清理链表中，等待下一次垃圾回收时执行finalizer函数。runfinq函数就是负责对待清理链表中的对象逐个执行finalizer函数的。

具体实现上，runfinq函数会首先把mheap中的finalizerwait队列中的对象放入待清理链表中，然后遍历待清理链表，对每个对象执行其finalizer函数。执行完finalizer函数后，该对象会从待清理链表中移除。具体的流程如下：

1. 从mheap.finalizerwait队列中把待清理对象转移到mheap.finqueue链表中；
2. 遍历mheap.finqueue链表，对其中每个对象执行finalizer函数；
3. 执行完finalizer函数后，从mheap.finqueue链表中移除该对象。

需要注意的是，由于finalizer函数的执行是在垃圾回收器中进行的，所以不能保证finalizer函数会在程序退出前一定能全部执行完毕。因此，在实现finalizer函数时不要依赖该函数一定会在程序退出前执行。



### SetFinalizer

SetFinalizer是Go语言中垃圾收集器的一个重要机制。它能够让程序员为一个对象指定一个回调函数，并在这个对象被垃圾收集器回收时自动调用这个回调函数。这个回调函数可以用来进行必要的资源清理操作，比如关闭文件、断开网络连接等。

具体来说，SetFinalizer函数的作用是为一个对象设置一个回调函数。这个回调函数会在对象被垃圾收集器回收时自动调用。SetFinalizer函数有两个参数，第一个参数是需要设置回调函数的对象，第二个参数是回调函数。这个回调函数可以是任意的普通函数、对象方法或闭包。回调函数的参数必须指向被回收的对象。

Go语言中的垃圾收集器是自动的，并且是基于引用计数的。当一个对象的引用计数为0时，会被垃圾收集器回收。如果一个对象设置了回调函数，那么在垃圾收集器回收这个对象时，会先调用回调函数，然后再回收这个对象。

需要注意的是，在SetFinalizer函数中设置的回调函数并不是一定会被调用的。由于垃圾收集器的实现策略不同，在某些情况下回调函数可能无法调用。因此，在使用SetFinalizer函数时，程序员需要保证回调函数只进行必要的资源清理操作，而不能依赖于它的调用。

总之，SetFinalizer函数是Go语言中一个有用的机制，能够帮助程序员管理资源，减少泄漏和错误。但是，需要慎重使用，避免滥用和过度依赖。



### KeepAlive

mfinal.go文件中的KeepAlive是一个函数，其作用是用于避免出现垃圾回收时被误回收的情况。在Go语言中，垃圾回收器通常使用三色标记法来识别和回收不再使用的内存空间。在这个过程中，垃圾回收器会扫描所有在堆上分配的对象，将其标记为活动对象或非活动对象。非活动对象就是待回收的垃圾对象。

如果一个活动对象和一个非活动对象A相互引用，那么如果A被认为是垃圾，那么这个活动对象将不能被正确的释放。这个时候就需要使用KeepAlive来避免活动对象被错误地释放掉。

实现方式如下：在KeepAlive函数中，传入一个需要保持存活的对象，这个对象会被KeepAlive函数保持引用，直到函数返回为止，这样在这个对象被GC回收时，即使这个对象和其他非活动对象相互引用，也不会被销毁掉。有了这个函数，程序员就可以更加自由地使用Go语言的垃圾回收机制，保证程序的正确执行。

总之，KeepAlive函数用于在垃圾回收发生时，确保对象没有被错误地回收。保持活动对象的引用，直到函数返回时结束。



