# File: build.go

build.go是Go语言中命令行工具"go build"的源代码文件。它的作用是根据指定的Go源代码文件或目录进行编译和链接，生成可执行文件。

该文件实现了"Go build"命令的主要功能，包括识别传入的参数，解析编译选项，调用编译器进行编译和链接等。具体来说，它包含以下几个主要函数：

1. main函数：程序的入口函数，用于初始化和解析命令行参数，并调用其他函数实现编译操作。

2. compile函数：根据传入的Go源代码文件或目录，调用编译器编译生成目标文件。

3. link函数：将编译生成的目标文件链接成可执行文件。

4. buildID函数：根据编译器生成的版本信息，生成可执行文件的唯一标识符。

5. toolID函数：根据编译器的名称和版本号，生成工具链的唯一标识符。

除了这些核心函数之外，build.go文件还包含了其他一些辅助函数，例如处理错误信息、打印编译日志等。总体来说，它是Go语言编译工具链中非常重要的一部分，为用户提供了简单方便的编译和构建方式。




---

### Var:

### goarch





### gorootBin





### gorootBinGo





### gohostarch





### gohostos





### goos





### goarm





### go386





### goamd64





### gomips





### gomips64





### goppc64





### goroot





### goroot_final





### goextlinkenabled





### gogcflags





### goldflags





### goexperiment





### workdir





### tooldir





### oldgoos





### oldgoarch





### oldgocache





### exe





### defaultcc





### defaultcxx





### defaultpkgconfig





### defaultldso





### rebuildall





### noOpt





### isRelease





### vflag





### okgoarch





### okgoos





### clangos





### oldtool





### unreleased





### deptab





### depsuffix





### gentab





### installed





### installedMu





### unixOS





### runtimegen





### cleanlist





### timeLogEnabled





### timeLogMu





### timeLogFile





### timeLogStart





### toolchain





### cgoEnabled





### broken





### firstClass





## Functions:

### find





### xinit





### compilerEnv





### compilerEnvLookup





### rmworkdir





### chomp





### findgoversion





### isGitRepo





### setup





### mustLinkExternal





### install





### startInstall





### runInstall





### packagefile





### matchtag





### shouldbuild





### copyfile





### dopack





### clean





### cmdenv





### timelog





### toolenv





### cmdbootstrap





### wrapperPathFor





### goInstall





### appendCompilerFlags





### goCmd





### checkNotStale





### needCC





### checkCC





### defaulttarg





### cmdinstall





### cmdclean





### cmdbanner





### banner





### cmdversion





### cmdlist





### IsRuntimePackagePath





### setNoOpt





