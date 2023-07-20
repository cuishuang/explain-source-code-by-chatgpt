# File: core/rawdb/accessors_trie.go

在go-ethereum项目中，core/rawdb/accessors_trie.go文件的作用是提供trie数据结构的访问器函数，该文件实现了与trie数据的读取、写入和删除相关的操作。

首先，让我们来了解一下hasherPool变量及其作用。在accessors_trie.go中，hasherPool是一个全局变量，它是一个对象池，用于管理NodeHasher对象的重用。NodeHasher是一个用于计算trie节点hash的辅助工具。通过将NodeHasher对象存储在hasherPool中，可以避免在每次需要计算hash时都创建和销毁新的对象，从而提高性能。

nodeHasher结构体在accessors_trie.go中定义了几个字段和方法，用于计算trie节点的hash。它是NodeHasher接口的一个实现，用于对多层trie进行哈希。

接下来，让我们逐个介绍函数的作用：

- newNodeHasher: 创建并返回一个新的NodeHasher对象。利用hasherPool变量从对象池中获取NodeHasher对象，如果对象池为空，则创建新的对象。

- returnHasherToPool: 将不再使用的NodeHasher对象放回hasherPool中，以便下次可以重用。

- hashData: 计算给定数据的哈希值。

- ReadAccountTrieNode: 从数据库中读取账户trie树中指定哈希的节点数据。

- HasAccountTrieNode: 检查数据库中是否存在账户trie树中指定哈希的节点。

- WriteAccountTrieNode: 将账户trie树中的节点数据写入数据库。

- DeleteAccountTrieNode: 从数据库中删除账户trie树中指定哈希的节点。

- ReadStorageTrieNode: 从数据库中读取存储trie树中指定哈希的节点数据。

- HasStorageTrieNode: 检查数据库中是否存在存储trie树中指定哈希的节点。

- WriteStorageTrieNode: 将存储trie树中的节点数据写入数据库。

- DeleteStorageTrieNode: 从数据库中删除存储trie树中指定哈希的节点。

- ReadLegacyTrieNode: 从数据库中读取遗留trie树中指定哈希的节点数据。

- HasLegacyTrieNode: 检查数据库中是否存在遗留trie树中指定哈希的节点。

- WriteLegacyTrieNode: 将遗留trie树中的节点数据写入数据库。

- DeleteLegacyTrieNode: 从数据库中删除遗留trie树中指定哈希的节点。

- HasTrieNode: 检查数据库中是否存在指定哈希的trie节点。

- ReadTrieNode: 从数据库中读取指定哈希的trie节点数据。

- WriteTrieNode: 将指定哈希的trie节点数据写入数据库。

- DeleteTrieNode: 从数据库中删除指定哈希的trie节点。

这些函数提供了对trie数据结构进行操作的接口。通过这些函数，可以实现在区块链中读取、写入和删除trie数据的功能，从而支持以太坊区块链的核心功能。

