# File: accounts/abi/method.go

在go-ethereum项目中，accounts/abi/method.go文件的作用是定义了与智能合约交互的ABI（Application Binary Interface）方法。

FunctionType结构体定义了方法类型，其中包含了可选的4个常量：FunctionTypeCall、FunctionTypeTransaction、FunctionTypeEstimateGas和FunctionTypeDeploy。它们分别用于标识不同类型的方法，如只调用、发送交易、估算耗费的gas和部署。

Method结构体表示智能合约中的函数或事件。它包含函数名（Name）、函数类型（FunctionType）、输入参数（Inputs）、输出参数（Outputs）等信息。

NewMethod函数用于创建一个新的Method实例，它接收函数名、函数类型和输入参数，并返回创建的Method对象。

String方法用于将Method对象转换为字符串形式，方便打印和显示。

IsConstant是一个便捷的方法，可以判断方法是否为常量（即是否仅查询，且不改变状态），它通过检查Method的FunctionType是否为FunctionTypeCall来判断。

IsPayable也是一个便捷的方法，用于检查方法是否接受以太币作为支付。它通过检查Method的FunctionType是否为FunctionTypeTransaction并且其中一个输入参数的类型是否为"address"来判断。

