# File: text/unicode/norm/readwriter.go

在Go的text包中，text/unicode/norm/readwriter.go文件的作用是提供了用于Unicode规范化（Normalization）的读写功能。

在该文件中，定义了两个结构体：normWriter和normReader。这两个结构体实现了io.Writer和io.Reader接口，用于对数据进行Unicode规范化的写入和读取操作。

normWriter结构体的作用是实现unicode规范化的写入操作。它包含一个嵌入式字段io.Writer，用于实现数据的写入功能。它还包含一个字段normForm，用于存储Unicode规范化的形式，可以是NormalizationFormC、NormalizationFormD、NormalizationFormKC或NormalizationFormKD。

normReader结构体的作用是实现unicode规范化的读取操作。它包含一个嵌入式字段io.Reader，用于实现数据的读取功能。它还包含一个字段normForm，用于存储Unicode规范化的形式，可以是NormalizationFormC、NormalizationFormD、NormalizationFormKC或NormalizationFormKD。

在这两个结构体中，有一些函数用于实现规范化的读写操作：

1. Write方法：该方法实现了io.Writer接口的Write方法，用于将规范化后的Unicode数据写入到底层的io.Writer中。

2. Close方法：该方法实现了io.Closer接口的Close方法，用于关闭底层的io.Writer。

3. Writer方法：该方法返回一个io.Writer接口，用于写入规范化的Unicode数据。

4. Read方法：该方法实现了io.Reader接口的Read方法，用于从底层的io.Reader中读取规范化后的Unicode数据。

5. Reader方法：该方法返回一个io.Reader接口，用于读取规范化的Unicode数据。

这些函数和方法的作用是为了提供方便的接口，使开发者可以在不直接操作底层数据流的情况下，进行Unicode规范化的读写操作。

