# File: crypto/ecies/params.go

在go-ethereum项目中，crypto/ecies/params.go文件定义了一系列ECIES（Elliptic Curve Integrated Encryption Scheme，椭圆曲线综合加密方案）相关的参数和结构体，以及用于操作这些参数和结构体的函数。

- DefaultCurve：默认的椭圆曲线参数，用于生成ECIES密钥对。
- ErrUnsupportedECDHAlgorithm：表示不支持的ECDH算法错误。
- ErrUnsupportedECIESParameters：表示不支持的ECIES参数错误。
- ErrInvalidKeyLen：表示无效的密钥长度错误。
- ECIES_AES128_SHA256, ECIES_AES192_SHA384, ECIES_AES256_SHA256, ECIES_AES256_SHA384：预定义的不同加密算法和摘要算法组合的ECIES参数。
- paramsFromCurve：根据给定的椭圆曲线参数生成对应的ECIES参数。

ECIESParams结构体表示ECIES的参数，包括加密算法、摘要算法、共享密钥长度等。具体包含的字段有：
- Cipher: 加密算法
- Hash: 摘要算法
- PublicKeyLen: 共享密钥长度
- IvLen: 初始化向量长度
- Overhead: 密文的额外长度（不包括加密后的数据本身）

AddParamsForCurve函数根据给定的椭圆曲线参数向包内的全局参数列表中添加对应的ECIES参数。
ParamsFromCurve函数根据给定的椭圆曲线参数获取对应的ECIES参数。
pubkeyParams函数用于根据指定的ECIES参数返回一个接口，该接口包含了生成公钥和私钥的方法。

