# File: cfg_test.go

cfg_test.go是Go语言标准库中cmd包下的一个测试文件，主要用于测试cmd包中与配置文件相关的函数和方法的正确性。其具体作用如下：

1.测试配置文件的解析

cfg_test.go中包含了TestReadConfigFile函数，用于测试读取、解析配置文件的函数是否能够正确地读取和解析配置文件，并将解析结果正确地存储到Config结构体中。

2.测试配置项的读取

cfg_test.go中包含了TestLookup函数，用于测试读取配置项的方法Lookup是否能够正确地读取到所需的配置项，并返回相应的值。此外，还测试了LookupString、LookupInt、LookupFloat和LookupBool等相关方法。

3.测试配置项的设置

cfg_test.go中包含了TestSet函数，用于测试设置配置项的方法Set是否能够正确地将指定的配置项设置为指定的值，并将修改后的配置信息正确地保存到指定的文件中。

总之，cfg_test.go是Go语言标准库中cmd包的测试文件，主要用于测试与配置文件相关的函数和方法的正确性和可靠性。

## Functions:

### writeFile





### writePkgConfig





### writeOutFileList





### runPkgCover





### TestCoverWithCfg





