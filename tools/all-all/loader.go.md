# File: tools/go/loader/loader.go

在Golang的Tools项目中，`tools/go/loader/loader.go`文件的作用是实现了一个通用程序加载器，用于加载和解析Go程序的源代码。这个文件中定义了一系列的数据结构和函数，用于完成从源代码到程序加载的过程。

下面分别介绍文件中的一些重要的结构体和函数：

ignoreVendor变量：用于控制是否忽略`vendor/`目录下的代码解析。

Config结构体：表示程序的加载配置，包含了一些加载选项和相应的setter方法。

PkgSpec结构体：表示一个要加载的包的规范，其中包含了包的导入路径和使用的标记。

Program结构体：表示一个加载的Go程序，包含了所有相关的包和导入信息。

PackageInfo结构体：表示一个已加载的包的信息，包含了包名、导入路径、语法树和导入的依赖包信息等。

importer接口：定义了一个将Go包导入到程序中的方法。

findpkgKey结构体：表示一个包的导入键，用于唯一标识一个包。

findpkgValue结构体：表示一个包的导入值，包含了导入路径和实际导入的包。

importInfo结构体：表示一个包的导入信息，包含了导入路径和导入的包的信息等。

importError结构体：表示导入包过程中的错误。

byImportPath结构体：用于根据导入路径进行包的排序。

closure结构体：表示一组包的闭包，即所有依赖的包和自身的包组成的集合。

String方法：返回一个结构体的字符串表示。

appendError方法：将错误追加到错误列表中。

fset变量：表示一个文件集合，用于对源代码进行解析。

ParseFile函数：解析Go源文件，返回文件的抽象语法树。

FromArgs函数：根据命令行参数创建一个程序配置。

CreateFromFilenames函数：根据文件名创建一个程序配置。

CreateFromFiles函数：根据源文件创建一个程序配置。

ImportWithTests函数：导入Go包及其测试包并返回程序。

Import函数：导入一个Go包并返回程序。

addImport函数：将一个导入信息添加到一个包的导入集合中。

PathEnclosingInterval函数：返回包含给定位置的语法树节点路径。

InitialPackages函数：返回通过标记指定的包的导入路径。

Package函数：返回一个包的导入路径和其指定的导入标记。

awaitCompletion函数：等待所有包的加载完成。

Complete函数：根据指定的配置完成程序的加载和解析。

Load函数：加载和解析一个Go程序的源代码。

Len函数：返回一个结构体的长度。

Less函数：比较两个结构体的大小。

Swap函数：交换两个结构体的位置。

markErrorFreePackages函数：标记所有不包含错误的包。

build函数：根据配置构建一个程序。

parsePackageFiles函数：根据给定的文件集合解析包文件。

doImport函数：导入一个包。

findPackage函数：根据给定的导入路径和配置查找并返回包。

importAll函数：导入所有当前包的依赖包。

findPath函数：根据导入路径查找包。

startLoad函数：开始加载一个包。

load函数：加载一个包并返回其导入信息。

addFiles函数：将文件添加到一个包中。

newPackageInfo函数：创建一个包的信息。

