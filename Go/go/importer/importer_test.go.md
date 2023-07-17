# File: importer_test.go

importer_test.go是Go语言的一个测试文件，用于测试go/importer包中的各种功能和方法。 

importer_test.go包含了一系列测试用例，这些测试用例用于检查importer包中各个函数、方法的正确性和稳定性。

具体来说，importer_test.go主要用于测试以下功能：

1. importPath：该函数用于将import路径转换为本地磁盘路径。
2. newImporter：该函数用于创建一个新的importer对象。
3. lookup：该方法用于查找给定名称的对象，并返回相应的类型信息。
4. pkgDir：该函数用于获取文件所在目录的包对象。
5. importFrom：该方法用于将给定的其他程序包导入到当前程序包中。

测试文件中的每个测试用例都会使用Go语言的标准测试工具来测试importer包中的不同函数和方法，并将测试结果输出到控制台。测试用例验证每个函数执行的正确性，并测试函数性能和可靠性。

总之，importer_test.go是Go语言importer模块的一个重要测试文件，用于测试和验证importer模块中的各项功能是否能正常工作。




---

### Var:

### importerTests








---

### Structs:

### importerTest





## Functions:

### runImporterTest





### TestGoxImporter





### gccgoPath





### TestObjImporter





