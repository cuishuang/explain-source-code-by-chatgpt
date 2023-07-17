# File: get.go

get.go文件是Go语言命令行工具（go tool）中的一个子命令，对应的子命令为“get”，用于从指定的远程仓库获取依赖包并自动解析其相互依赖关系，并将其安装在本地的$GOPATH/pkg目录之下。

具体来说，get子命令可以完成以下操作：

1. 下载指定包的源码，并自动解析其依赖关系。
2. 编译指定包的源码，并生成相应的可执行文件或库文件。
3. 安装指定包的可执行文件或库文件到本地$GOPATH/bin和$GOPATH/pkg目录下，以便后续使用。
4. 根据指定的tag、branch或commit ID下载指定的版本。

除此之外，get子命令还支持一些扩展的功能，如：

1. 支持同时下载多个包，并自动解决它们之间的依赖关系。
2. 支持从非官方仓库中下载包，如Github、Bitbucket等。
3. 支持对指定的包进行版本控制，并支持通过-v参数查看具体的版本信息。
4. 支持对下载过的包进行更新、升级、删除等管理操作。

总之，get.go文件在Go语言编译工具中具有重要的作用，它是使用Go语言开发的必要工具之一，也是Go语言项目管理的重要组成部分。




---

### Var:

### CmdGet





### HelpGopathGet





### getD





### getF





### getT





### getU





### getFix





### getInsecure





### downloadCache





### downloadRootCache





## Functions:

### init





### runGet





### downloadPaths





### download





### downloadPackage





### selectTag





### checkImportPath





