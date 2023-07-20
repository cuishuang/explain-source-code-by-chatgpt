# File: accounts/keystore/account_cache.go

在go-ethereum项目中，accounts/keystore/account_cache.go文件的作用是管理以太坊账户的缓存，并提供对这些账户的操作和查询方法。

accountsByURL是一个映射关系，用于存储以太坊账户的URL和accountCache结构体的对应关系。

AmbiguousAddrError是一个自定义错误类型，表示存在多个与给定地址匹配的账户。

accountCache是一个结构体，用于表示以太坊账户的缓存。它包含了账户的地址、密钥文件路径等信息，并提供对这些信息的读写方法。

Len方法用于返回accountCache的长度（账户数）。

Less方法用于比较两个accountCache结构体的地址。

Swap方法用于交换两个accountCache结构体的位置。

Error方法用于返回AmbiguousAddrError的错误信息。

newAccountCache函数创建并返回一个新的accountCache实例。

accounts函数返回存储在accountCache中的所有以太坊账户。

hasAddress函数检查缓存中是否存在给定地址的账户。

add函数将一个以太坊账户添加到缓存中。

delete函数从缓存中删除指定地址的账户。

deleteByFile函数从缓存中删除与给定密钥文件路径匹配的账户。

watcherStarted函数检查是否已启动账户监视器。

removeAccount函数从缓存中移除指定地址的账户。

find函数根据给定地址查找对应的accountCache结构体。

maybeReload函数检查由于文件修改而导致的账户更新，并进行相应的修改。

close函数用于关闭accountCache实例。

scanAccounts函数用于扫描指定目录下的密钥文件，将其中的账户添加到缓存中。

