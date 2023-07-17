# File: runtime.go

runtime.go是Go语言的运行时库中的一个文件，它包含了一些底层的运行时实现，用于支持Go语言程序的执行。该文件中包含了许多与内存管理、调度、垃圾回收、锁等相关的函数和结构体，是Go语言的核心部分之一。

具体来说，runtime.go的作用有以下几个方面：

1. 内存管理：runtime.go中定义了许多内存管理相关的函数和结构体，如go:mcache、go:mcentral、go:mheap等等，用于管理分配和回收内存、堆栈等资源，确保程序正常运行。

2. 调度：runtime.go也包含了与调度相关的函数和结构体，如go:g、go:p、go:work、go:runq等等。调度器是Go语言并发编程的基础，负责协调和分配线程等各种资源，确保程序的并发执行。

3. 垃圾回收：由于Go语言使用了自动垃圾回收机制，因此runtime.go中也包含了与垃圾回收相关的函数和结构体，如go:memstats、go:gchelper、go:gcwork、go:findObject等等。这些函数和结构体用于跟踪并管理垃圾对象，确保及时地回收不再使用的内存。

4. 锁：Go语言中的锁机制也是常见的并发编程模式，因此runtime.go中也包含了与锁相关的函数和结构体，如go:mutex、go:semaphore、go:cond等等。这些函数和结构体用于管理临界区，防止并发访问产生竞争条件。

总之，runtime.go是Go语言核心部分之一，提供了底层的支持，使得Go语言能够在多核多线程环境下高效地运行。虽然这些函数和结构体大部分都不需要用户手动调用，但了解这些底层实现对于理解和优化程序依然非常重要。




---

### Var:

### writeBarrier





### x86HasPOPCNT





### x86HasSSE41





### x86HasFMA





### armHasVFPv4





### arm64HasATOMICS





## Functions:

### newobject





### mallocgc





### panicdivide





### panicshift





### panicmakeslicelen





### panicmakeslicecap





### throwinit





### panicwrap





### gopanic





### gorecover





### goschedguarded





### goPanicIndex





### goPanicIndexU





### goPanicSliceAlen





### goPanicSliceAlenU





### goPanicSliceAcap





### goPanicSliceAcapU





### goPanicSliceB





### goPanicSliceBU





### goPanicSlice3Alen





### goPanicSlice3AlenU





### goPanicSlice3Acap





### goPanicSlice3AcapU





### goPanicSlice3B





### goPanicSlice3BU





### goPanicSlice3C





### goPanicSlice3CU





### goPanicSliceConvert





### printbool





### printfloat





### printint





### printhex





### printuint





### printcomplex





### printstring





### printpointer





### printuintptr





### printiface





### printeface





### printslice





### printnl





### printsp





### printlock





### printunlock





### concatstring2





### concatstring3





### concatstring4





### concatstring5





### concatstrings





### cmpstring





### intstring





### slicebytetostring





### slicebytetostringtmp





### slicerunetostring





### stringtoslicebyte





### stringtoslicerune





### slicecopy





### decoderune





### countrunes





### convI2I





### convT





### convTnoptr





### convT16





### convT32





### convT64





### convTstring





### convTslice





### assertE2I





### assertE2I2





### assertI2I





### assertI2I2





### panicdottypeE





### panicdottypeI





### panicnildottype





### ifaceeq





### efaceeq





### fastrand





### makemap64





### makemap





### makemap_small





### mapaccess1





### mapaccess1_fast32





### mapaccess1_fast64





### mapaccess1_faststr





### mapaccess1_fat





### mapaccess2





### mapaccess2_fast32





### mapaccess2_fast64





### mapaccess2_faststr





### mapaccess2_fat





### mapassign





### mapassign_fast32





### mapassign_fast32ptr





### mapassign_fast64





### mapassign_fast64ptr





### mapassign_faststr





### mapiterinit





### mapdelete





### mapdelete_fast32





### mapdelete_fast64





### mapdelete_faststr





### mapiternext





### mapclear





### makechan64





### makechan





### chanrecv1





### chanrecv2





### chansend1





### closechan





### typedmemmove





### typedmemclr





### typedslicecopy





### selectnbsend





### selectnbrecv





### selectsetpc





### selectgo





### block





### makeslice





### makeslice64





### makeslicecopy





### growslice





### unsafeslicecheckptr





### panicunsafeslicelen





### panicunsafeslicenilptr





### unsafestringcheckptr





### panicunsafestringlen





### panicunsafestringnilptr





### mulUintptr





### memmove





### memclrNoHeapPointers





### memclrHasPointers





### memequal





### memequal0





### memequal8





### memequal16





### memequal32





### memequal64





### memequal128





### f32equal





### f64equal





### c64equal





### c128equal





### strequal





### interequal





### nilinterequal





### memhash





### memhash0





### memhash8





### memhash16





### memhash32





### memhash64





### memhash128





### f32hash





### f64hash





### c64hash





### c128hash





### strhash





### interhash





### nilinterhash





### int64div





### uint64div





### int64mod





### uint64mod





### float64toint64





### float64touint64





### float64touint32





### int64tofloat64





### int64tofloat32





### uint64tofloat64





### uint64tofloat32





### uint32tofloat64





### complex128div





### getcallerpc





### getcallersp





### racefuncenter





### racefuncexit





### raceread





### racewrite





### racereadrange





### racewriterange





### msanread





### msanwrite





### msanmove





### asanread





### asanwrite





### checkptrAlignment





### checkptrArithmetic





### libfuzzerTraceCmp1





### libfuzzerTraceCmp2





### libfuzzerTraceCmp4





### libfuzzerTraceCmp8





### libfuzzerTraceConstCmp1





### libfuzzerTraceConstCmp2





### libfuzzerTraceConstCmp4





### libfuzzerTraceConstCmp8





### libfuzzerHookStrCmp





### libfuzzerHookEqualFold





### addCovMeta





