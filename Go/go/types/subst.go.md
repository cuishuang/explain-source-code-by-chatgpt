# File: subst.go

subst.go是Go语言标准库中的一个文件，主要涉及了字符串替换的相关操作。其作用是实现字符串中子串的替换操作，例如将字符串中的某个字符或者子串替换为另一个字符或者子串。具体而言，subst.go定义了一个Replace函数，该函数可以接收源字符串、需要替换的子串、替换为的子串以及替换的次数作为参数，实现了在源字符串中将指定的子串替换为目标子串的功能。

Replace函数支持的替换操作还包括了正则表达式的替换，只需要将需要替换的子串用正则表达式表示即可。此外，由于Replace函数设计为可重入的，因此用户可以在替换的过程中逐步修改源字符串，实现更加灵活的操作。

除了Replace函数之外，subst.go还定义了一些其他的函数和常量，例如SplitAfterN函数、SplitN函数和MaxUint64等。这些函数和常量与替换操作相关，可以进一步扩展Replace函数的功能。




---

### Structs:

### substMap





### subster





## Functions:

### makeSubstMap





### makeRenameMap





### empty





### lookup





### subst





### typ





### typOrNil





### var_





### substVar





### tuple





### varList





### func_





### substFunc





### funcList





### typeList





### termlist





### replaceRecvType





