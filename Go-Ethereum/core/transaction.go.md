# File: core/types/transaction.go

在go-ethereum项目中，core/types/transaction.go文件是用于定义和操作以太坊交易的数据结构和方法。

- ErrInvalidSig：无效的交易签名错误
- ErrUnexpectedProtection：意外保护错误
- ErrInvalidTxType：无效的交易类型错误
- ErrTxTypeNotSupported：不支持的交易类型错误
- ErrGasFeeCapTooLow：Gas费用上限过低错误
- errShortTypedTx：短类型交易错误

以下是一些重要的结构体和其作用：
- Transaction：表示以太坊的交易，包含交易的所有字段和方法。
- TxData：交易的数据字段
- Transactions：表示一组交易
- TxByNonce：根据Nonce排序的交易集合
- TxWithMinerFee：带有矿工费用的交易集合
- TxByPriceAndTime：根据价格和时间排序的交易集合
- TransactionsByPriceAndNonce：根据价格和Nonce排序的交易集合

以下是一些重要的方法和其作用：
- NewTx：创建一个新的交易
- EncodeRLP：将交易编码为RLP格式的字节串
- encodeTyped：将TypedData编码为字节串
- MarshalBinary：将交易序列化为字节流
- DecodeRLP：从RLP字节串解码交易
- UnmarshalBinary：从字节流解码交易
- decodeTyped：从字节串解码TypedData
- setDecoded：设置解码标志
- sanityCheckSignature：验证交易签名的合法性
- isProtectedV：判断交易是否为受保护的V值
- Protected：返回交易是否受保护的V值
- Type：返回交易类型
- ChainId：返回交易的链ID
- Data：返回交易的数据
- AccessList：返回交易的访问列表
- Gas：返回交易的Gas限制
- GasPrice：返回交易的Gas价格
- GasTipCap：返回交易的Gas小费上限
- GasFeeCap：返回交易的Gas费用上限
- BlobGas：返回交易的BlobGas值
- BlobGasFeeCap：返回交易的BlobGas费用上限
- BlobHashes：返回交易的BlobHashes
- Value：返回交易的价值
- Nonce：返回交易的Nonce
- To：返回交易的接收地址
- Cost：返回交易的成本
- RawSignatureValues：返回交易的原始签名值
- GasFeeCapCmp：比较交易的Gas费用上限
- GasFeeCapIntCmp：将交易的Gas费用上限与整数值进行比较
- GasTipCapCmp：比较交易的Gas小费上限
- GasTipCapIntCmp：将交易的Gas小费上限与整数值进行比较
- EffectiveGasTip：返回有效的Gas小费上限
- EffectiveGasTipValue：返回有效的Gas小费上限值
- EffectiveGasTipCmp：比较有效的Gas小费上限
- EffectiveGasTipIntCmp：将有效的Gas小费上限与整数值进行比较
- BlobGasFeeCapCmp：比较交易的BlobGas费用上限
- BlobGasFeeCapIntCmp：将交易的BlobGas费用上限与整数值进行比较
- Hash：计算交易的哈希值
- Size：返回交易的大小
- WithSignature：为交易设置签名
- Len：返回交易集合的长度
- EncodeIndex：将索引编码为字节串
- TxDifference：比较两个交易的差异
- HashDifference：比较两个交易哈希的差异
- Less：比较两个交易的大小
- Swap：交换两个交易的位置
- NewTxWithMinerFee：创建带有矿工费用的新交易
- Push：向交易集合中添加一个新交易
- Pop：从交易集合中弹出最后一个交易
- NewTransactionsByPriceAndNonce：创建根据价格和Nonce排序的交易集合
- Peek：返回交易集合中最后一个交易
- Shift：从交易集合中弹出第一个交易
- copyAddressPtr：复制地址指针的副本

