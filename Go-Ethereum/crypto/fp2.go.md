# File: crypto/bls12381/fp2.go

在go-ethereum项目中，crypto/bls12381/fp2.go文件的作用是实现了有限域Fp2和其上的基本运算。下面逐个介绍每个结构体和函数的作用：

- fp2Temp结构体：表示Fp2的临时变量。用于临时存储运算结果，以减少内存分配。

- fp2结构体：表示Fp2的实例。它由两个Fp类型的元素(a, b)组成，表示复数a + bi。

- newFp2Temp()函数：创建一个新的Fp2临时变量。

- newFp2()函数：创建一个新的Fp2实例。

- fromBytes()函数：将字节切片转换为Fp2实例。

- toBytes()函数：将Fp2实例转换为字节切片。

- new()函数：创建并返回一个新的Fp2实例，其中a和b为Fp类型的参数。

- zero()函数：返回零元素的Fp2实例。

- one()函数：返回单位元素的Fp2实例。

- add()函数：计算两个Fp2实例的和，并返回一个新的Fp2实例。

- addAssign()函数：将两个Fp2实例相加，并将结果存储在第一个实例中。

- ladd()函数：计算两个Fp2实例的和，并返回一个新的Fp2实例。

- double()函数：计算一个Fp2实例的加倍值，并返回一个新的Fp2实例。

- doubleAssign()函数：将一个Fp2实例加倍，并将结果存储在原实例中。

- ldouble()函数：计算一个Fp2实例的加倍值，并返回一个新的Fp2实例。

- sub()函数：计算两个Fp2实例的差，并返回一个新的Fp2实例。

- subAssign()函数：将两个Fp2实例相减，并将结果存储在第一个实例中。

- neg()函数：计算一个Fp2实例的负值，并返回一个新的Fp2实例。

- mul()函数：计算两个Fp2实例的乘积，并返回一个新的Fp2实例。

- mulAssign()函数：将两个Fp2实例相乘，并将结果存储在第一个实例中。

- square()函数：计算一个Fp2实例的平方，并返回一个新的Fp2实例。

- squareAssign()函数：计算一个Fp2实例的平方，并将结果存储在原实例中。

- mulByNonResidue()函数：将一个Fp2实例乘以特定的非剩余元素，并返回一个新的Fp2实例。

- mulByB()函数：将一个Fp2实例乘以特定的常数b，并返回一个新的Fp2实例。

- inverse()函数：计算一个Fp2实例的逆，并返回一个新的Fp2实例。

- mulByFq()函数：将一个Fp2实例与一个Fp实例相乘，并返回一个新的Fp2实例。

- exp()函数：计算一个Fp2实例的指数运算，并返回一个新的Fp2实例。

- frobeniusMap()函数：计算Fp2实例的Frobenius映射，并返回一个新的Fp2实例。

- frobeniusMapAssign()函数：计算Fp2实例的Frobenius映射，并将结果存储在原实例中。

- sqrt()函数：计算一个Fp2实例的平方根，并返回一个新的Fp2实例。

- isQuadraticNonResidue()函数：检查一个Fp2实例是否为二次非剩余元素。

