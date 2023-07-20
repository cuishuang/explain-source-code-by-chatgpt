# File: core/types/transaction_signing.go

core/types/transaction_signing.go这个文件在go-ethereum项目中的作用是处理以太坊交易的签名相关逻辑。

ErrInvalidChainId是一个定义了链ID无效的错误变量。

big8是一个用于表示8的大整数。

sigCache是一个用于缓存签名的结构体，用于在签名过程中防止重复签名。

Signer是一个接口，定义了对交易进行签名的方法。

cancunSigner、londonSigner、eip2930Signer、EIP155Signer、HomesteadSigner、FrontierSigner是实现了Signer接口的结构体，分别对应不同的以太坊协议版本。

MakeSigner是根据指定的链ID创建对应的签名器。

LatestSigner是根据给定的链头和链ID获取最新的签名器。

LatestSignerForChainID是根据给定的链ID获取最新的签名器。

SignTx是对交易进行签名的方法。

SignNewTx是对新交易进行签名的方法。

MustSignNewTx是对新交易进行签名的方法，如果出现错误则会抛出panic。

Sender是根据交易和签名获取发送者地址的方法。

NewCancunSigner是创建一个Cancun协议版本的签名器的方法。

Equal是判断两个签名是否相等的方法。

SignatureValues是将签名分成R、S和V值的结构体。

Hash是计算交易哈希的方法。

NewLondonSigner是创建一个London协议版本的签名器的方法。

NewEIP2930Signer是创建一个EIP2930协议版本的签名器的方法。

ChainID是表示链ID的类型。

NewEIP155Signer是创建一个EIP155协议版本的签名器的方法。

decodeSignature是解码签名的方法。

recoverPlain是使用明文签名恢复公钥的方法。

deriveChainId是根据链ID计算父链ID的方法。

