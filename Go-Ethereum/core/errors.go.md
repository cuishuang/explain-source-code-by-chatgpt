# File: core/vm/errors.go

在go-ethereum项目中，core/vm/errors.go文件是虚拟机执行时发生错误时使用的包含错误类型和错误消息的定义文件。

- ErrOutOfGas: 表示执行过程中的燃料耗尽错误。
- ErrCodeStoreOutOfGas: 表示编码存储尝试使用的燃料不足错误。
- ErrDepth: 表示执行深度超过了最大允许深度的错误。
- ErrInsufficientBalance: 表示账户余额不足以支付交易燃料的错误。
- ErrContractAddressCollision: 表示创建合约时遇到地址冲突的错误。
- ErrExecutionReverted: 表示执行过程中遇到了回滚错误。
- ErrMaxInitCodeSizeExceeded: 表示合约初始化代码的大小超过了最大允许大小的错误。
- ErrMaxCodeSizeExceeded: 表示合约代码的大小超过了最大允许大小的错误。
- ErrInvalidJump: 表示无效的跳转指令错误。
- ErrWriteProtection: 表示试图写入保护内存的错误。
- ErrReturnDataOutOfBounds: 表示返回的数据超出了允许的边界的错误。
- ErrGasUintOverflow: 表示燃料计算溢出的错误。
- ErrInvalidCode: 表示无效的合约代码错误。
- ErrNonceUintOverflow: 表示nonce计数溢出的错误。
- errStopToken: 表示执行过程遇到STOP操作码的错误。

ErrStackUnderflow、ErrStackOverflow、ErrInvalidOpCode是表示执行期间栈操作错误的结构体。

Error函数是用于创建错误的辅助函数，根据给定的错误类型和错误详细信息创建一个错误实例并返回。

