# File: eth/catalyst/simulated_beacon_api.go

在go-ethereum项目中，`simulated_beacon_api.go`文件的作用是实现了模拟的Beacon Chain API，它是一个用于模拟Beacon链的API，供开发者进行测试和验证。

该文件中定义了几个API结构体，其中主要包括三个结构体：`BeaconBlockResponse`、`BeaconCommittee`和`BeaconValidatorsResponse`。

- `BeaconBlockResponse`结构体用于表示模拟的Beacon区块的响应信息。它包含了区块的高度、区块的hash值、区块的时间戳等信息。

- `BeaconCommittee`结构体用于表示模拟的Beacon共识委员会信息。它包含了共识委员会的编号、共识委员会的成员列表等信息。

- `BeaconValidatorsResponse`结构体用于表示模拟的Beacon验证者的响应信息。它包含了验证者的公钥、验证者的状态、验证者在共识委员会中的角色等信息。

另外，`simulated_beacon_api.go`文件中还定义了一些函数，包括`AddWithdrawal`和`SetFeeRecipient`等。

- `AddWithdrawal`函数用于模拟添加提款请求。它接收提款者的地址和提款金额作为参数，并将提款请求添加到模拟的Beacon链中。

- `SetFeeRecipient`函数用于设置手续费接收地址。它接收手续费接收地址作为参数，并将该地址设置为模拟的Beacon链的手续费接收地址。

这些函数的作用是为开发者提供一些模拟行为，以供测试和验证各种场景下的功能和逻辑。

