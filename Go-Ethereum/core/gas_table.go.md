# File: core/vm/gas_table.go

在go-ethereum项目中，`core/vm/gas_table.go`文件的作用是定义了EVM（以太坊虚拟机）的各种操作和指令的所需的燃气（gas）消耗。燃气是以太坊网络中衡量计算资源消耗的单位，每个操作和指令都需要一定数量的燃气来执行。

以下是一些相关变量和函数的详细介绍：

变量：
1. `gasCallDataCopy`：用于存储`CALLDATACOPY`操作的燃气消耗数量。
2. `gasCodeCopy`：用于存储`CODECOPY`操作的燃气消耗数量。
3. `gasMcopy`：用于存储内存复制操作的燃气消耗数量。
4. `gasExtCodeCopy`：用于存储`EXTCODECOPY`操作的燃气消耗数量。
5. `gasReturnDataCopy`：用于存储`RETURNDATACOPY`操作的燃气消耗数量。
6. `gasReturn`：用于存储`RETURN`操作的燃气消耗数量。
7. `gasRevert`：用于存储`REVERT`操作的燃气消耗数量。
8. `gasMLoad`：用于存储内存加载操作的燃气消耗数量。
9. `gasMStore8`：用于存储单字节内存存储操作的燃气消耗数量。
10. `gasMStore`：用于存储内存存储操作的燃气消耗数量。
11. `gasCreate`：用于存储`CREATE`操作的燃气消耗数量。

函数：
1. `memoryGasCost`：计算内存扩展和备份操作的燃气消耗。
2. `memoryCopierGas`：计算内存复制操作的燃气消耗。
3. `gasSStore`：计算状态存储操作的燃气消耗。
4. `gasSStoreEIP2200`：使用 EIP-2200 惯例计算状态存储操作的燃气消耗。
5. `makeGasLog`：计算日志操作的燃气消耗。
6. `gasKeccak256`：计算Keccak256哈希操作的燃气消耗。
7. `pureMemoryGascost`：计算纯内存操作的燃气消耗。
8. `gasCreate2`：计算`CREATE2`操作的燃气消耗。
9. `gasCreateEip3860`：使用 EIP-3860 惯例计算`CREATE`操作的燃气消耗。
10. `gasCreate2Eip3860`：使用 EIP-3860 惯例计算`CREATE2`操作的燃气消耗。
11. `gasExpFrontier`：计算 EVM 中EXP操作的燃气消耗（早期版本）。
12. `gasExpEIP158`：计算 EVM 中EXP操作的燃气消耗（根据EIP-158）。
13. `gasCall`：计算`CALL`操作的燃气消耗。
14. `gasCallCode`：计算`CALLCODE`操作的燃气消耗。
15. `gasDelegateCall`：计算`DELEGATECALL`操作的燃气消耗。
16. `gasStaticCall`：计算`STATICCALL`操作的燃气消耗。
17. `gasSelfdestruct`：计算`SELFDESTRUCT`操作的燃气消耗。

这些变量和函数定义了以太坊虚拟机执行各种操作时所需的燃气消耗，是以太坊网络中的计算资源衡量单位。

