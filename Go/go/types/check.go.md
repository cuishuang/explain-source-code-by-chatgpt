# File: check.go

check.go文件是Go语言工具链中的一个重要文件之一，它主要包含了一系列用于检查代码规范、语法正确性以及代码性能和可移植性等方面的工具函数和示例。

具体来说，check.go文件中定义了一些常用的语法检查工具函数，例如CheckFile函数用于检查单个源文件的语法正确性，CheckPackage函数用于检查整个包的语法正确性，以及Fprint和Print等用于输出错误信息的打印函数等。

此外，check.go文件中还包含了各种性能和可移植性检查工具函数的实现，例如Walk函数用于遍历整个包的所有Go文件并生成语法树，RangeDecls函数用于遍历语法树并查找所有的变量和函数声明，以及Sizeof检查函数用于检查各种Go语言类型的字节数等等。

总体来说，check.go文件的作用是为开发人员提供一系列方便、快捷、准确的代码检查和性能调试工具，可以帮助开发人员快速发现并修复代码问题，提高代码的质量和可维护性。




---

### Var:

### nopos





### errBadCgo








---

### Structs:

### exprInfo





### environment





### importKey





### dotImportKey





### action





### actionDesc





### Checker





### cleaner





### posVers





### bailout





## Functions:

### lookup





### describef





### addDeclDep





### brokenAlias





### validAlias





### isBrokenAlias





### rememberUntyped





### later





### push





### pop





### needsCleanup





### NewChecker





### initFiles





### handleBailout





### Files





### checkFiles





### processDelayed





### cleanup





### record





### recordUntyped





### recordTypeAndValue





### recordBuiltinType





### recordCommaOkTypes





### recordInstance





### instantiatedIdent





### recordDef





### recordUse





### recordImplicit





### recordSelection





### recordScope





