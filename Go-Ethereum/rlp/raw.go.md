# File: rlp/raw.go

在go-ethereum项目中，rlp/raw.go文件的作用是实现RLP编解码器。RLP（Recursive Length Prefix）是一种编码格式，主要用于以紧凑的方式序列化数据。该文件包含了实现将数据编码为RLP格式和将RLP格式解码为数据的函数。

ratValueType变量是一个枚举类型，表示RLP编码的数据类型。它包括String、Bytes、List和Int四种类型。

RawValue结构体是一个帮助结构体，用于保存解码的原始字节数组和数据类型。

- StringSize函数用于计算给定字符串的RLP编码所需的字节数。
- BytesSize函数用于计算给定字节数组的RLP编码所需的字节数。
- ListSize函数用于计算给定列表的RLP编码所需的字节数。
- IntSize函数用于计算给定整数的RLP编码所需的字节数。
- Split函数用于将原始的RLP字节数组按照编码的规则拆分为各个子项。
- SplitString函数用于将原始的RLP字节数组解码为字符串类型。
- SplitUint64函数用于将原始的RLP字节数组解码为Uint64类型。
- SplitList函数用于将原始的RLP字节数组解码为列表类型。
- CountValues函数用于计算原始的RLP字节数组中包含的子项个数。
- readKind函数用于读取原始的RLP字节数组中的数据类型。
- readSize函数用于读取原始的RLP字节数组中的数据长度。
- AppendUint64函数用于将Uint64类型的数据编码为RLP格式并追加到字节数组中。

这些函数共同实现了将数据编码为RLP格式或将RLP格式解码为数据的功能，提供了对RLP编解码的支持。

