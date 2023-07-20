# File: core/vm/contracts.go

在go-ethereum项目的core/vm/contracts.go文件中，包含了与预编译合约相关的功能实现。

以下是变量的作用：

- PrecompiledContractsHomestead、PrecompiledContractsByzantium、PrecompiledContractsIstanbul、PrecompiledContractsBerlin、PrecompiledContractsCancun、PrecompiledContractsBLS：这些变量用于定义预编译合约的函数实现列表，分别对应以太坊的不同硬分叉版本，存储了不同硬分叉版本下的所有预编译合约的函数实现。
- PrecompiledAddressesCancun、PrecompiledAddressesBerlin、PrecompiledAddressesIstanbul、PrecompiledAddressesByzantium、PrecompiledAddressesHomestead：这些变量用于定义不同硬分叉版本下的预编译合约的地址列表，存储了以太坊不同硬分叉版本下的所有预编译合约的地址。
- big0、big1、big3、big4、big7、big8、big16、big20、big32、big64、big96、big480、big1024、big3072、big199680：这些变量用于存储常用的大整数。

以下是结构体的作用：

- PrecompiledContract：这个结构体用于存储预编译合约的信息，包括合约的代码和合约函数的操作数栈数量。
- ecrecover、sha256hash、ripemd160hash、dataCopy、bigModExp、bn256AddIstanbul、bn256AddByzantium、bn256ScalarMulIstanbul、bn256ScalarMulByzantium、bn256PairingIstanbul、bn256PairingByzantium、blake2F、bls12381G1Add、bls12381G1Mul、bls12381G1MultiExp、bls12381G2Add、bls12381G2Mul、bls12381G2MultiExp、bls12381Pairing、bls12381MapG1、bls12381MapG2、kzgPointEvaluation：这些结构体用于存储对应的预编译合约的函数实现。

以下是函数的作用：

- init：该函数用于初始化预编译合约相关的数据。
- ActivePrecompiles：该函数返回指定硬分叉版本下所有活跃的预编译合约地址。
- RunPrecompiledContract：该函数用于执行预编译合约的函数，将指定的操作数栈和内存传递给预编译合约的实现函数，并返回结果。
- RequiredGas：该函数根据执行预编译合约所需的操作数和内存大小计算执行所需的燃料。
- Run：该函数用于执行智能合约，包括调用了以太坊的预编译合约的部分。
- modexpMultComplexity：该函数返回使用模幂乘算法计算指定位数的乘方所需的复杂度。
- newCurvePoint、newTwistPoint：这些函数用于创建椭圆曲线点。
- runBn256Add、runBn256ScalarMul、runBn256Pairing：这些函数用于执行bn256相关的操作，如点加法、标量乘法和计算对。
- decodeBLS12381FieldElement：该函数用于解码BLS12-381域上的元素。
- kzgToVersionedHash：该函数用于将KZG验证的结果转换为版本化哈希。

