# File: complexnums.go

complexnums.go文件是Go语言标准库中一个实现复数运算的文件。该文件定义了一个名为`complex`的类型，用于表示复数，并提供了一系列与复数相关的操作函数，包括复数的四则运算、取模、求幅角、求实部和虚部等。具体函数列表如下：

1. `func Complex(x, y float64) complex128`
   根据实部x和虚部y构造一个复数

2. `func Conj(z complex128) complex128`
   返回一个指定复数的共轭复数

3. `func Abs(z complex128) float64`
   返回一个指定复数的模

4. `func Phase(z complex128) float64`
   返回一个指定复数的幅角

5. `func Rect(r, θ float64) complex128`
   根据模r和相位角θ构造一个复数

6. `func Polar(r, θ float64) (x, y float64)`
   根据模r和相位角θ计算出该复数的实部和虚部

7. `func Real(z complex128) float64`
   返回一个指定复数的实部

8. `func Imag(z complex128) float64`
   返回一个指定复数的虚部

9. `func Cmplx(x, y float64) complex128`
   根据实部x和虚部y构造一个复数（可以处理不同类型的x和y）

除此之外，`complexnums.go`文件还通过Operator()函数为`complex`类型提供了相应的操作符运算函数重载，包括加减乘除、指数运算、等于和不等于等。通过对复数的支持，Go语言在处理数学或者物理模拟等任务时变得更加方便和高效。

