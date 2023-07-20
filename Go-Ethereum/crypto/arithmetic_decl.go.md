# File: crypto/bls12381/arithmetic_decl.go

在go-ethereum项目中，`crypto/bls12381/arithmetic_decl.go`文件定义了BLS12-381曲线上的有限域运算的结构和操作函数。

该文件中的`mul`结构体表示BLS12-381曲线上有限域中的元素（Field Element），它包含两个`fr.Element`类型的字段x和y。mul结构体的目的是封装有限域的运算，提供各种运算操作的函数。

以下是`mul`结构体中的函数以及它们的作用：

- `init`：对mul对象进行初始化，设置默认值。
- `square`：计算mul对象的平方。
- `neg`：计算mul对象的负元素。
- `add`：将两个mul对象相加并返回结果。
- `addAssign`：将一个mul对象和另一个mul对象相加并将结果存储在原始mul对象中。
- `ladd`：将两个mul对象相加并返回结果，其中一个mul对象是一个常量。
- `laddAssign`：将一个mul对象和一个常量相加并将结果存储在原始mul对象中。
- `double`：将mul对象乘以2并返回结果。
- `doubleAssign`：将mul对象乘以2并将结果存储在原始mul对象中。
- `ldouble`：将mul对象乘以一个常量2并返回结果。
- `sub`：将两个mul对象相减并返回结果。
- `subAssign`：将一个mul对象和另一个mul对象相减并将结果存储在原始mul对象中。
- `lsubAssign`：将一个mul对象和一个常量相减并将结果存储在原始mul对象中。
- `_neg`：计算mul对象的相反数。
- `mulNoADX`：对两个mul对象进行乘法运算并返回结果，这个函数没有使用ADX指令集。
- `mulADX`：对两个mul对象进行乘法运算并返回结果，这个函数使用了ADX指令集。

这些函数实现了有限域上的各种运算操作，用于计算BLS12-381曲线上的加法、减法、乘法和取反等运算。

