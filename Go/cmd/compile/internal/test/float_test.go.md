# File: float_test.go

float_test.go是Go语言标准库中测试浮点数相关函数的文件。它主要包括以下几个方面的测试：

1.测试浮点数的转换函数：包括ParseFloat、ParseInt等函数，用于测试将浮点数转换为其他类型的函数是否正确。

2.测试常用的算术函数：包括Add、Sub、Mul等函数，用于测试四则运算以及乘方运算是否正确。

3.测试特殊值的处理：包括NaN、Infinity等值的处理，用于测试在处理浮点数时对于这些特殊值的正确处理方式。

4.测试精度：由于浮点数在计算机中的存储与表示存在误差，所以需要测试浮点数函数的精度是否符合预期。

float_test.go的作用是确保Go语言标准库中浮点数相关函数的正确性、鲁棒性和稳定性，从而保证Go语言在处理浮点数时能够达到预期的精度和正确性。




---

### Var:

### snan32bitsVar





### snan64bitsVar





### sinkFloat





## Functions:

### compare1





### compare2





### TestFloatCompare





### TestFloatCompareFolded





### cvt1





### cvt2





### cvt3





### cvt4





### cvt5





### cvt6





### cvt7





### cvt8





### cvt9





### cvt10





### cvt11





### cvt12





### f2i64p





### ip64





### TestFloatConvert





### TestFloatConvertFolded





### TestFloat32StoreToLoadConstantFold





### TestFloatSignalingNaN





### TestFloatSignalingNaNConversion





### TestFloatSignalingNaNConversionConst





### BenchmarkMul2





### BenchmarkMulNeg2





