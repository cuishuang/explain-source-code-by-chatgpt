# File: unify.go

unify.go这个文件的作用是实现了Go语言中类型的归一化，即将不同但具有相同底层类型的类型转换为同一类型。该文件中定义了一个私有结构体unifier，其中包含两个字典map[typePair]typePair和map[type.Type]typePair，分别用于存储类型对(typePair)和类型(type.Type)的映射关系。

类型归一化的主要目的是消除不必要的类型转换，并使得代码更加清晰简洁。在Go语言中，虽然带有相同底层类型的类型可以互相转换，但是它们并不相等，这可能导致程序中存在大量不必要的类型转换。unify.go中的归一化机制就是为了解决这个问题。

unifier结构体中的两个字典map的作用如下：

1. map[typePair]typePair：将不同但具有相同底层类型的类型进行归一化。通过将类型对(typePair)作为键，将同一底层类型的类型映射到一个公共类型(typePair)上。这样，就可以消除不必要的类型转换，并实现类型的归一化。

2. map[type.Type]typePair：建立类型和归一化类型之间的映射关系。这样，即使在之前没有出现的类型，也可以找到其相应的归一化类型。

通过类型归一化机制，程序可以避免不必要的类型转换，从而优化代码性能和可读性。




---

### Structs:

### unifier





### typeParamsById





## Functions:

### newUnifier





### unify





### tracef





### String





### Len





### Less





### Swap





### join





### asTypeParam





### setHandle





### at





### set





### unknowns





### inferred





### nify





