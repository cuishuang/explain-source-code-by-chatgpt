# File: text/internal/colltab/weighter.go

weighter.go文件是text项目中text/internal/colltab包的一部分，它实现了用于Unicode字符权重计算的功能。Unicode字符权重是根据字符的默认排序顺序定义的，用于在文本排序和比较操作中判断字符的相对位置。

该文件中定义了weighter接口及其两个实现结构体：SimpleWeighter和ComplexWeighter。

1. weighter接口：该接口定义了计算字符权重的方法。具体而言，该接口包括以下方法：
   - Weight(r rune) (int64, error)：计算给定Unicode字符的权重。该方法返回一个int64类型的权重值和一个错误值。如果计算权重时发生错误，例如传入了无效的Unicode字符，则会返回错误。
   - WeightString(s string) (int64, error)：计算给定字符串的整体权重。该方法将字符串拆分为Unicode字符，并计算每个字符的权重，最后返回所有字符的累积权重之和。

2. SimpleWeighter结构体：该结构体实现了weighter接口。SimpleWeighter的计算方法基于字符的rune值，并将其转化为int64类型的权重。它没有依赖于具体的语言或区域设置，而是按照字符的Unicode值进行排序。

3. ComplexWeighter结构体：该结构体也实现了weighter接口。ComplexWeighter的计算方法更为复杂，它利用Unicode规范中指定的Unicode字符属性来计算权重。根据不同的字符属性（如General_Category、Canonical_Combining_Class等），给定字符的权重会有所不同。此结构体的实现更适合进行多语言的排序和比较操作。

综上所述，weighter.go文件中定义了计算字符权重的接口和两个具体实现结构体。这些结构体可用于进行Unicode字符的排序和比较操作，以实现文本处理和语言特定的规则。

