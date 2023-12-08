# File: text/internal/colltab/contract.go

在Go的text项目中，text/internal/colltab/contract.go文件的作用是实现了与排序和比较相关的Unicode字符合同的表示和解析。

该文件定义了以下结构体：

1. ContractTrieSet：表示Unicode字符合同的Trie集合。它用于存储和查询字符合同Trie，以检查两个字符是否属于同一合同。
2. ctScanner：表示字符合同的扫描器。它用于遍历字符合同的各个子合同。
3. ctScannerString：表示字符合同字符的扫描器。它用于遍历字符合同中的Unicode字符。

以下是这些结构体的详细介绍：

1. ContractTrieSet：ContractTrieSet是一个包含ContractTrie的集合。通过使用位图来表示TrieSet的内容，它提供了高效的查找和插入操作。
2. ctScanner：ctScanner是用于扫描字符合同的子合同的结构体。它提供了从字符合同主体中提取子合同的功能。
3. ctScannerString：ctScannerString是用于扫描字符合同的Unicode字符的结构体。它允许以字符串形式访问字符合同。
   
以下是一些函数的功能介绍：

1. scanner：scanner函数是根据合同字符串构建一个ctScanner对象。该函数接收一个合同字符串作为输入，并返回一个ctScanner对象，该对象可以通过MoveNext方法来遍历字符合同的子合同。
2. scannerString：scannerString函数是根据合同字符串构建一个ctScannerString对象。该函数接收一个合同字符串作为输入，并返回一个ctScannerString对象，该对象可以通过MoveNext方法来遍历字符合同的Unicode字符。
3. result：result函数根据提供的ctScanner和字符合同实例，返回布尔值。该函数检查两个字符是否属于同一合同。
4. scan：scan函数根据提供的待扫描字符串和字符合同实例，返回一个ctScanner对象。该函数接收一个字符串和字符合同作为输入，并返回一个ctScanner对象，该对象可以通过MoveNext方法来遍历字符串的字符，并忽略不在字符合同中的字符。

这些函数和结构体的组合提供了合同解析和比较功能，用于确定字符是否属于同一字符合同。这在排序和比较字符串时非常有用，因为它们可以确保正确处理不同语言和脚本中的排序规则。

