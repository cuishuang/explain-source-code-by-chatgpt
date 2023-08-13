# File: tsdb/chunkenc/bstream.go

在Prometheus项目中，tsdb/chunkenc/bstream.go文件的作用是实现基于bit流的编码和解码。该文件定义了几个结构体和函数，用于处理位级别的数据流。

1. 结构体：
- bstream：该结构体用于读取和写入位数据，底层使用字节数组存储。
- bit：表示一个位，取值为0或1。
- bstreamReader：用于从bstream中读取位数据。

2. 函数：
- bytes：将位数据流转换为字节数组。
- writeBit：向bstream中写入一个位。
- writeByte：向bstream中写入一个字节。
- writeBits：向bstream中写入多个位。
- newBReader：创建一个新的bstreamReader实例。
- readBit：从bstreamReader中读取一个位。
- readBitFast：通过缓存读取一个位，优化读取性能。
- readBits：从bstreamReader中读取多个位。
- readBitsFast：通过缓存读取多个位，优化读取性能。
- ReadByte：从bstreamReader中读取一个字节。
- loadNextBuffer：加载下一个字节的数据到bstreamReader缓冲区。

这些函数和结构体一起完成了对位数据的编码和解码操作。通过bstream和bstreamReader，可以将数据按位存储，并以高效的方式进行读取和写入。其他辅助函数如writeBit、readBit等用于简化位级别的操作，而bytes则用于将位数据转换为字节数组。

