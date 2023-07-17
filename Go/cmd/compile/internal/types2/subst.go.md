# File: subst.go

subst.go是Go语言标准库中的一个文件，它的作用是实现Go源码中的字符串替换功能。

具体来说，subst.go中的代码实现了一个Replace函数，该函数可以将一个字符串中的某个子串替换成另一个指定的字符串。该函数的原型如下：

```
func Replace(s, old, new string, n int) string
```

其中，s是待替换的原始字符串，old是要被替换的子串，new是替换后的字符串，n是指定替换次数，如果n小于0表示替换所有。

该函数内部使用了strings包中的一些函数，如Index和Join等，来实现基本的字符串操作。

当Go语言源代码需要进行字符串替换操作时，通常会直接使用subst.go中的Replace函数，而无需重新实现。这也体现了Go语言的简便性和高效性。




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





