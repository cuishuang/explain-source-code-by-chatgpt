# File: light/postprocess.go

在go-ethereum项目中，light/postprocess.go文件的作用是处理轻客户端的后处理逻辑。主要包括索引配置和相关的索引器、处理函数等。

DefaultServerIndexerConfig、DefaultClientIndexerConfig、TestServerIndexerConfig、TestClientIndexerConfig是默认的索引器配置和测试的索引器配置。

errNoTrustedCht、errNoTrustedBloomTrie、errNoHeader是一些错误变量，表示没有信任的累计哈希树(CHT)、布隆过滤器(Bloom Trie)或者头块的错误。

IndexerConfig结构体是用来配置ChtIndexer和BloomTrieIndexer的参数。

ChtNode结构体用于存储累计哈希树(CHT)的节点数据。

ChtIndexerBackend结构体是累计哈希树(CHT)索引器的后端实现。

BloomTrieIndexerBackend结构体是布隆过滤器(Bloom Trie)索引器的后端实现。

GetChtRoot函数用于获取累计哈希树(CHT)的根节点。

StoreChtRoot函数用于存储累计哈希树(CHT)的根节点。

NewChtIndexer函数用于创建一个新的累计哈希树(CHT)索引器。

fetchMissingNodes函数用于获取缺失的节点数据。

Reset函数用于重置索引器的状态。

Process函数用于处理区块数据，包括累计哈希树(CHT)和布隆过滤器(Bloom Trie)的处理。

Commit函数用于提交索引器的更改。

Prune函数用于修剪索引器的数据。

GetBloomTrieRoot函数用于获取布隆过滤器(Bloom Trie)的根节点。

StoreBloomTrieRoot函数用于存储布隆过滤器(Bloom Trie)的根节点。

NewBloomTrieIndexer函数用于创建一个新的布隆过滤器(Bloom Trie)索引器。

这些函数主要是用于处理和操作累计哈希树(CHT)和布隆过滤器(Bloom Trie)的相关数据，以及索引器的创建、存储和数据处理等操作。

