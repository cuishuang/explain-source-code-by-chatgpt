# File: cmd/kubeadm/app/util/pubkeypin/pubkeypin.go

在Kubernetes项目中，cmd/kubeadm/app/util/pubkeypin/pubkeypin.go文件是用于公钥固定(pin)验证的工具包。公钥固定是一种安全机制，用于确保客户端与服务器之间通信的安全性，以防止中间人攻击。

`supportedFormats`是一个切片，包含了支持的公钥固定的格式。目前该切片包含了`SHA256`和`SHA1`两种格式。

`Set`是一个结构体，表示一组公钥固定。它包含了固定的公钥和可选的固定格式。

- `NewSet`函数用于创建一个新的公钥固定集合。
- `Allow`函数用于为公钥固定集合添加公钥和固定格式。
- `CheckAny`函数用于检查给定的证书公钥是否与集合中的任何一个公钥固定匹配。
- `Empty`函数用于检查集合是否为空。
- `Hash`函数用于计算给定的证书公钥的散列值。
- `allowSHA256`函数用于创建一个仅支持SHA256格式的公钥固定集合。
- `checkSHA256`函数用于检查给定证书公钥是否与SHA256格式的任一公钥固定匹配。

这些函数和结构体一起提供了一种方便的方法来处理和验证公钥固定，以确保安全的通信。

