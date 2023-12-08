# File: text/unicode/norm/input.go

在Go的text/unicode/norm/input.go文件中，包含了一些有关Unicode字符规范化的函数和结构体。

1. 结构体定义:
   - inputBytes: 表示输入字符串的字节数组。
   - inputString: 表示输入字符串的字符串形式。
   - setBytes: 表示字符集的字节数组。
   - setString: 表示字符集的字符串形式。
   - _byte: 表示已读取的字节数。
   - skipASCII: 标记是否跳过ASCII字符规范化。
   - skipContinuationBytes: 标记是否跳过继续字节的规范化。
   - appendSlice: 标记是否可以追加字符片段。
   - copySlice: 标记是否需要复制字符片段。
   - charinfoNFC: 包含了一些关于字符规范化的信息，用于NFC规范化。
   - charinfoNFKC: 包含了一些关于字符规范化的信息，用于NFKC规范化。
   - hangul: 存储Unicode Hangul字符的相关信息。

2. 函数定义:
   - inputString() string: 返回inputBytes的字符串形式。
   - setInputBytes([]byte): 设置inputBytes的值。
   - setInputString(string): 设置inputBytes的字符串形式。
   - setBytes() ([]byte, error): 返回setString的字节数组形式。
   - setString(string): 设置setString的值。
   - skipASCII() bool: 返回是否允许跳过ASCII字符规范化的标记。
   - skipContinuationBytes() bool: 返回是否允许跳过继续字节规范化的标记。
   - appendSlice() bool: 返回是否允许追加字符片段的标记。
   - copySlice() bool: 返回是否需要复制字符片段的标记。
   - charinfoNFC(rune) *properties: 返回指定rune的字符规范化信息，用于NFC规范化。
   - charinfoNFKC(rune) *properties: 返回指定rune的字符规范化信息，用于NFKC规范化。
   - hangul(decomposition, offset rune) (lo, hi rune): 返回指定分解和偏移量的Hangul字符。

这些函数和结构体的作用是为了提供一些基本的输入和处理功能，用于Unicode字符规范化的过程。他们允许设置输入和字符集的不同形式，标记是否需要跳过ASCII字符规范化或继续字节规范化，以及提供了一些与字符规范化相关的信息和处理方法。

